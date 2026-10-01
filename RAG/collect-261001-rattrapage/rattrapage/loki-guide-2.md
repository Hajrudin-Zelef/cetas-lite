---
id: collect-261001-rattrapage/rattrapage/loki-guide-2
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "aws", "memory"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [204, 405]
sha256: 4e3a3d7b5b243be6e3cd6e945d6c7e97f680b4af701e699558d5d9f192e360a4
---

# Grafana Loki — Le guide complet

## 4. Pas d'index plein texte — la différence avec Elasticsearch

Elasticsearch construit un **index inversé** : pour chaque mot de chaque
ligne, il sait dans quels documents il apparaît. Puissant, mais coûteux en
disque, en RAM et en CPU à l'ingestion.

Loki fait le choix inverse :

1. À l'ingestion : il n'indexe que les **labels**. Le texte brut est
   compressé tel quel (gzip) dans les chunks.
2. À la requête : il sélectionne d'abord les flux via les labels
   (rapide, index petit), **puis** il filtre le texte avec `|=`, `|~`
   en décompressant les chunks concernés.

Implications concrètes :

- **Requête rapide** = sélecteur de labels sélectif
  (`{job="nginx", host="web01"}` sur 1 heure).
- **Requête lente** = sélecteur trop large (`{job=~".+"}` sur 30 jours)
  + filtre texte : Loki doit décompresser des montagnes de chunks.
- **Bonne pratique** : concevez vos labels pour que 90 % des requêtes
  touchent un petit nombre de flux.

> ⚠️ **Avertissement cardinalité** : chaque combinaison unique de labels
> crée un flux en mémoire dans l'ingester. 10 labels à 10 valeurs chacun
> = potentiellement 10 milliards de flux. C'est le crash assuré.
> Voir sections 25 à 27.

---

## 5. Architecture : vue d'ensemble

En mode distribué, Loki se compose de plusieurs composants. En mode
**monolithique** (recommandé pour démarrer), un seul binaire fait tout.

```
                    ┌─────────────┐
  Promtail ────────►│ Distributeur│──┐
  (agents)          └─────────────┘  │
                                     ▼
                    ┌─────────────┐  ┌──────────┐   ┌──────────────┐
                    │Query-frontend│◄─│ Querier  │◄──│  Ingester    │
                    └──────┬──────┘  └──────────┘   └──────┬───────┘
                           │                               │
                           ▼                               ▼
                    ┌─────────────┐                 ┌──────────────┐
                    │   Grafana   │                 │   Stockage   │
                    └─────────────┘                 │ (chunks+index)│
                                                    └──────────────┘
```

| Composant | Rôle |
|---|---|
| Distributeur | reçoit les logs (push depuis Promtail), valide, route vers les ingesters |
| Ingester | écrit les lignes dans les chunks en mémoire, flushe vers le stockage |
| Querier | exécute les requêtes LogQL (chunks en mémoire + stockage) |
| Query-frontend | cache, parallélise et déduplique les requêtes (optionnel mais recommandé) |
| Ruler | évalue les règles d'alerting/enregistrement sur les logs |
| Compactor | compacte l'index, applique la rétention, déduplique |

**Mode monolithique** : un seul processus `loki` avec `-target=all`
fait tourner les 6 rôles. Parfait jusqu'à ~quelques To de logs/jour.

---

## 6. Le distributeur (distributor)

Le distributeur est la **porte d'entrée** des logs :

1. Il reçoit les requêtes HTTP `POST /loki/api/v1/push` de Promtail
   (ou de tout client compatible).
2. Il **valide** : labels présents ? taille des lignes OK ? rate limiting
   par tenant respecté ?
3. Il **répartit** les flux vers les ingesters via un anneau de hachage
   cohérent (hash ring) — le même flux va toujours vers les mêmes ingesters.
4. Il applique le **rate limiting** global et par flux
   (`ingestion_rate_mb`, `per_stream_rate_limit`).

Réglages importants (dans `limits_config`) :

```yaml
limits_config:
  ingestion_rate_mb: 16        # débit max par distributeur (Mo/s)
  ingestion_burst_size_mb: 32  # pic autorisé
  per_stream_rate_limit: 5MB   # par flux (défaut 3 Mo/s en 3.x)
  per_stream_rate_limit_burst: 20MB
  max_streams_per_user: 0      # 0 = illimité ; à fixer en prod (ex: 10000)
  max_line_size: 256kB         # lignes plus longues = rejetées (tronquées si truncate)
  max_line_size_truncate: true # tronquer plutôt que rejeter (recommandé)
```

---

## 7. L'ingester

L'ingester est le **cœur chaud** de Loki :

- il maintient les chunks **en mémoire** pour les flux actifs ;
- il répond aux requêtes sur les données récentes (non encore flushées) ;
- il flushe les chunks vers le stockage selon :
  - `chunk_idle_period` (défaut 30 min) : aucun nouveau log depuis X ;
  - `max_chunk_age` (défaut 2 h) : chunk trop vieux ;
  - `chunk_target_size` (~1,5 Mo) : chunk assez gros.

Points de vigilance :

- **Mémoire** : proportionnelle au nombre de flux actifs. Surveillez
  `loki_ingester_memory_streams`.
- **WAL** (write-ahead log) : depuis Loki 2.4, l'ingester écrit un journal
  sur disque (`wal.dir`) pour survivre à un redémarrage sans perdre les
  chunks en mémoire. **Activez-le toujours en production.**

```yaml
ingester:
  wal:
    enabled: true
    dir: /var/lib/loki/wal
  chunk_idle_period: 30m
  max_chunk_age: 2h
  chunk_target_size: 1572864   # ~1,5 Mo
  max_chunk_age: 2h
```

---

## 8. Le querier et le query-frontend

**Querier** : exécute les requêtes LogQL. Il interroge :
1. les ingesters (données récentes en mémoire),
2. le stockage (chunks flushés + index).

**Query-frontend** (fortement recommandé en production) :
- **met en cache** les résultats (`results_cache`) ;
- **découpe** les grosses requêtes en sous-requêtes parallèles
  (`split_queries_by_interval: 30m`) ;
- **limite** la parallélisation (`max_outstanding_per_tenant`) ;
- **rejette** les requêtes trop gourmandes avant qu'elles ne tuent les queriers.

```yaml
query_range:
  split_queries_by_interval: 30m
  results_cache:
    cache:
      embedded_cache:
        enabled: true
        max_size_mb: 100

limits_config:
  max_query_parallelism: 32
  max_outstanding_per_tenant: 2048
  max_query_length: 721h        # 30 jours max par requête
  max_query_lookback: 721h
```

---

## 9. Le stockage : filesystem vs stockage objet

Loki sépare **chunks** (les logs compressés) et **index** (labels → chunks).

| Backend | Chunks | Index | Usage |
|---|---|---|---|
| `filesystem` | disque local | BoltDB local | lab, petits déploiements |
| `s3` (+ compatible : MinIO) | objet | TSDB ou BoltDB | **production recommandée** |
| `gcs` / `azure` | objet | TSDB | clouds GCP / Azure |
| `tsdb` (index) | — | fichiers TSDB sur objet | **recommandé depuis 2.8** |

**Recommandation production** : chunks sur **S3 (ou MinIO on-premise)** +
index **TSDB**. Le filesystem ne passe pas à l'échelle (un seul nœud,
pas de réplication) et complique la sauvegarde.

Exemple de configuration S3/MinIO :

```yaml
storage_config:
  tsdb_shipper:
    active_index_directory: /var/lib/loki/tsdb-index
    cache_location: /var/lib/loki/tsdb-cache
  aws:
    s3: s3://access_key:secret_key@minio.local:9000/loki-chunks
    s3forcepathstyle: true
```

> 💡 Avec MinIO on-premise, vous gardez un stockage objet S3-compatible
> sans dépendre d'un cloud public. Idéal pour une infra d'entreprise.

---

## 10. Installation sur Debian/Ubuntu : le binaire

Loki est distribué en binaire statique unique. Procédure manuelle propre :

```bash
# 1. Variables
LOKI_VERSION="3.3.2"
ARCH="amd64"   # ou arm64

# 2. Téléchargement (vérifiez la version sur github.com/grafana/loki/releases)
cd /tmp
curl -sSLO "https://github.com/grafana/loki/releases/download/v${LOKI_VERSION}/loki-linux-${ARCH}.zip"
curl -sSLO "https://github.com/grafana/loki/releases/download/v${LOKI_VERSION}/promtail-linux-${ARCH}.zip"

