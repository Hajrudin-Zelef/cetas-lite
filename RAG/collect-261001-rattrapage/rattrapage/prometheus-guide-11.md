---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-11
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [2262, 2515]
sha256: b32b46a5a5e0a0af590d984b00528ef26ccd9a17f7e1efe51fea4c35b6c063cf
---

# Guide Prometheus — Supervision métrique complète

1. Le snapshot **reste dans** `snapshots/` après création : il faut le supprimer
   à la main après archivage, sinon le disque se remplit.
2. `delete_series` ne libère l'espace qu'après **compaction** (~2 h) et ne touche
   pas aux snapshots. Pour vraiment purger, il faut supprimer les blocs.

> **Règle 3-2-1** : 3 copies, 2 supports, 1 hors site. Le snapshot seul sur la
> même machine n'est pas une sauvegarde.

## 75. Rétention : dimensionner le disque

```bash
--storage.tsdb.retention.time=30d
--storage.tsdb.retention.size=50GB
```

Estimation : Prometheus compresse à ~**1 à 2 octets par échantillon**.

```
Espace ≈ nb_séries × (86400 / scrape_interval) × jours × 1,5 octet
```

Exemple : 100 000 séries, scrape 15 s, 30 jours :

```
100 000 × 5760 échantillons/jour × 30 j × 1,5 o ≈ 26 Go
```

Commandes de diagnostic :

```promql
# Nombre de séries actives (l'indicateur n°1 à surveiller)
prometheus_tsdb_head_series

# Taille WAL / blocs
prometheus_tsdb_wal_segment_current
prometheus_tsdb_size_retention_total_bytes
```

⚠️ Si `retention.size` est atteinte, Prometheus **purge les blocs les plus
anciens sans prévenir** (c'est son fonctionnement normal, pas un bug).

## 76. Haute disponibilité : la paire de Prometheus

Le pattern standard : **2 Prometheus identiques** qui scrapent les mêmes cibles
(scrape redondant), avec `external_labels` différents :

```yaml
# prom-01
global:
  external_labels: {replica: 'prom-01', datacenter: 'dc-paris'}

# prom-02 : identique sauf replica: 'prom-02'
```

Côté Alertmanager : **cluster** (gossip) pour dédupliquer :

```bash
# alertmanager-01
--cluster.listen-address=0.0.0.0:9094
--cluster.peer=alertmanager-02.lan:9094

# alertmanager-02 : miroir
```

Et Prometheus pointe vers les deux :

```yaml
alerting:
  alertmanagers:
    - static_configs:
        - targets: ['am-01.lan:9093', 'am-02.lan:9093']
```

Schéma :

```
 [Prometheus 01] ──┐
                   ├──► [Alertmanager cluster] ──► notifications (1 seule fois)
 [Prometheus 02] ──┘
```

⚠️ Sans cluster Alertmanager, chaque Prometheus enverrait ses propres
notifications → **doublons**. Le cluster gossip déduplique.

Limite : les deux Prometheus stockent en double (2× disque). Pour aller plus
loin : Thanos / Mimir.

## 77. Haute disponibilité : Thanos (introduction)

**Thanos** ajoute au-dessus de Prometheus : stockage objet illimité (S3),
requête globale multi-clusters, déduplication des réplicas, downsampling.

```
 [Prometheus 01] ──sidecar──┐
                            ├──► [Thanos Store / S3] ──► [Thanos Query] ──► Grafana
 [Prometheus 02] ──sidecar──┘         (rétention infinie)
```

Composants :

| Composant | Rôle |
|---|---|
| **Sidecar** | Upload les blocs TSDB vers S3, proxy les requêtes |
| **Store Gateway** | Lit l'historique depuis S3 |
| **Compactor** | Compacte + downsample (5 min, 1 h) pour les longues périodes |
| **Query** | Point d'entrée unique : fusionne sidecars + store |
| **Ruler** | Évalue les règles sur les données globales |

Quand l'adopter : multi-sites, besoin de rétention > 90 jours, ou quand la
paire de Prometheus ne suffit plus. Pour un seul site avec 30 jours de
rétention, **la paire suffit** — ne pas sur-ingénierer.

## 78. Haute disponibilité : Mimir / Cortex (introduction)

**Grafana Mimir** (fork de Cortex) : TSDB distribuée, compatible remote_write /
remote_read, multi-tenant, haute dispo native.

```
 [Prometheus × N] ──remote_write──► [Mimir] ──► Grafana (1 an+ de rétention)
```

| | Paire Prometheus | Thanos | Mimir |
|---|---|---|---|
| Complexité | Faible | Moyenne | Élevée |
| Rétention longue | Non | Oui (S3) | Oui (S3) |
| Requête globale | Non (fédération manuelle) | Oui | Oui |
| Idéal pour | 1 site, < 6 mois | Multi-sites, S3 dispo | Gros volumes, multi-tenant |

> **Conseil de chef de service** : commencez par **une paire bien configurée +
> snapshots quotidiens**. Migrez vers Thanos/Mimir quand le besoin (rétention
> longue, multi-sites) est avéré, pas avant.

---

# PARTIE VI — Cardinalité, sécurité

---

## 79. La cardinalité : le piège n°1 de Prometheus

La **cardinalité** = le nombre de séries temporelles uniques. C'est **le**
facteur limitant de Prometheus : chaque série coûte de la RAM (~1-2 Ko en head)
et ralentit les requêtes.

```
1 métrique × 10 valeurs de label A × 10 valeurs de label B = 100 séries
1 métrique × 10 000 user_id × 100 chemins = 1 000 000 000 séries  💥
```

Ordres de grandeur :

| Séries actives | État |
|---|---|
| < 1 M | Confortable |
| 1 – 5 M | OK avec RAM adaptée (8-16 Go) |
| 5 – 10 M | Zone rouge : optimiser |
| > 10 M | OOM probable, requêtes lentes |

Surveiller : `prometheus_tsdb_head_series` + page **Status → TSDB**
(top 10 des métriques par nombre de séries).

## 80. Cardinalité : ce qu'il ne faut JAMAIS faire

```python
# ❌ CATASTROPHIQUE : un label par utilisateur / par email / par IP client
http_requests_total{user_id="48293", email="jean.dupont@x.lan"}
login_attempts_total{client_ip="10.0.4.128"}

# ❌ Très mauvais : chemins d'URL non normalisés (1 série par URL !)
http_requests_total{path="/produit/12893", path="/produit/12894", ...}

# ❌ Mauvais : ID de requête, UUID, timestamp en label
job_duration_seconds{job_id="a3f8c9e2-..."}

# ✅ BON : labels à cardinalité bornée et connue
http_requests_total{method="GET", status="200", route="/produit/:id"}
#                         (5 valeurs)  (10)        (50 routes) = 2 500 séries max
```

**Règle d'or** : un label doit avoir un nombre de valeurs **borné, petit et
prévisible** (< 100 en général). Tout identifiant unique (user_id, IP client,
UUID, chemin brut) est **interdit** en label.

Autres pièges :

- `labelmap` qui promeut tous les labels Kubernetes (`__meta_kubernetes_*`)
  sans filtrer.
- Exporter applicatif qui expose un label `version` incluant le hash de build.
- `info`-métriques (`node_uname_info`, `target_info`) : 1 série par cible,
  c'est OK, mais ne pas les multiplier.

## 81. Cardinalité : diagnostiquer et corriger

Diagnostic :

```promql
# Top 10 des métriques les plus lourdes (par nombre de séries)
topk(10, count by (__name__) ({__name__=~".+"}))

# Nombre de séries d'une métrique suspecte
count({__name__="http_requests_total"})

# Quelles valeurs de label explosent ?
count by (path) ({__name__="http_requests_total"})
```

Corrections (par ordre de préférence) :

1. **Côté application** : supprimer/normaliser le label fautif (vrai fix).
2. **`metric_relabel_configs: drop`** : jeter la métrique ou le label
   avant ingestion (section 21).
3. **`labeldrop`** : supprimer juste le label à haute cardinalité, garder la métrique :

```yaml
metric_relabel_configs:
  - regex: 'user_id|client_ip|job_id'
    action: labeldrop
```

4. **Agrégation en amont** : recording rule qui `sum by` sans le label fautif,
   puis drop de la métrique brute.

⚠️ `labeldrop` après ingestion massive ne récupère pas la RAM immédiatement :
le head se vide au fil des compactions. Le vrai fix reste côté émetteur.

## 82. Sécurité : authentification sur l'interface web

Prometheus natif n'a **pas d'authentification** (voulu : simplicité). Options :

**a) Basic auth native (depuis Prometheus 2.x, fichier web.yml)** :

```yaml
# /etc/prometheus/web.yml
basic_auth_users:
  admin: '$2y$05$...'     # hash bcrypt généré par htpasswd
  lecteur: '$2y$05$...'
```

```bash
htpasswd -nBC 10 "" | tr -d ':\n'   # génère le hash (paquet apache2-utils)
```

```bash
--web.config.file=/etc/prometheus/web.yml
```

**b) Reverse proxy (recommandé)** : nginx/Apache devant, avec auth + TLS +
ACL IP. Voir section 83.

**c) Ne pas exposer** : `web.listen-address=127.0.0.1:9090` + accès via SSH
ou VPN uniquement. Le plus simple est souvent le plus sûr.

## 83. Sécurité : TLS et reverse proxy

