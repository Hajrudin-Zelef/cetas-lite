---
id: collect-261001-rattrapage/rattrapage/loki-guide-12
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "memory"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [2835, 3071]
sha256: 5fd01e3fab20dd0764f5202c8ca43fc12293a9ec01467a048665259d992c60b4
---

# Grafana Loki — Le guide complet

**Symptôme** : le disque grossit malgré `retention_period: 744h`.
**Causes** : `compactor.retention_enabled: false` (oubli classique !),
compactor non démarré (`-target` partiel), ou `retention_delete_delay`
pas encore écoulé. Détail section 79.

### Cas 9 — Logs en retard de plusieurs minutes

**Symptôme** : les logs arrivent dans Grafana avec 2-5 min de retard.
**Causes** : `batchwait` trop haut, horloge désynchronisée (NTP !),
ingester saturé.
**Remède** : `batchwait: 1s`, vérifier `chrony`, dimensionner.

### Cas 10 — Montée de version cassée (schéma)

**Symptôme** : après upgrade, erreurs `schema mismatch` ou requêtes vides
sur l'historique.
**Cause** : entrée `schema_config` modifiée au lieu d'ajoutée.
**Remède** : restaurer l'ancienne entrée + ajouter la nouvelle avec
`from:` futur (section 66). Ne jamais réécrire l'historique des schémas.

---

## 71. Erreur 1 : `too many outstanding requests`

**Message typique (côté Promtail)** :
```
error sending batch: too many outstanding requests
```

**Ce que ça veut dire** : le client Promtail a plus de requêtes en vol
que Loki n'en accepte (`max_outstanding_per_tenant`, défaut 2048), ou
Loki répond trop lentement et la file sature.

**Causes** :
- pic d'ingestion brutal (redémarrage d'app qui vide son buffer) ;
- Loki sous-dimensionné ou disque lent (flush bloqué) ;
- réseau saturé entre Promtail et Loki.

**Remèdes** :
1. Côté Promtail, augmenter la taille des batchs (moins de requêtes) :
   ```yaml
   clients:
     - url: http://loki.local:3100/loki/api/v1/push
       batchsize: 2097152   # 2 Mo au lieu d'1 Mo
       batchwait: 2s
   ```
2. Côté Loki, vérifier que l'ingester flushe bien (disque, `flush_op_timeout`).
3. Si récurrent : scaler (plus de RAM/disque, ou mode `write` dédié).

---

## 72. Erreur 2 : `entry out of order`

**Message** : `rpc error: entry out of order` (HTTP 400 au push).

**Ce que ça veut dire** : pour un flux donné, Loki a reçu une ligne avec
un timestamp **antérieur** à la dernière ligne déjà ingérée. Loki exige
l'ordre chronologique par flux (les chunks sont append-only).

**Causes classiques** :
- deux Promtail lisent le **même fichier** (ex : fichier partagé en NFS,
  conteneur lu par 2 agents) ;
- horloge qui recule (VM restaurée depuis un snapshot, NTP qui corrige
  un gros décalage) ;
- `positions.yaml` corrompu → relecture d'anciennes lignes.

**Remèdes** :
```yaml
limits_config:
  # Tolère un léger désordre (défaut : 0 = aucun)
  max_out_of_order_time: 30s
```
> ⚠️ N'augmentez que si la cause est comprise (ex : logs multi-threadés
> avec timestamps proches). Le vrai fix est presque toujours : **une
> seule source par fichier**.

---

## 73. Erreur 3 : `per-stream rate limit exceeded`

**Message** : `per-stream rate limit exceeded (limit: 3MB/sec, burst: 15MB/sec)`.

**Ce que ça veut dire** : **un seul flux** dépasse le débit autorisé.
C'est le garde-fou anti-« un flux qui écrase tout » (souvent le symptôme
d'un label à trop forte cardinalité… ou d'une application qui logue en
boucle).

**Diagnostic** :
```logql
# Quel flux consomme le plus ? (top débit par job/host)
topk(10, sum by (job, host) (bytes_rate({job=~".+"}[5m])))
```

**Remèdes** :
1. Trouver la cause : bug applicatif (log en boucle) ? label trop fin ?
2. Ajuster si légitime (gros job batch qui logue beaucoup) :
   ```yaml
   limits_config:
     per_stream_rate_limit: 10MB
     per_stream_rate_limit_burst: 40MB
   ```
3. Si c'est un label à forte cardinalité : le supprimer (section 26),
   pas augmenter la limite.

---

## 74. Erreur 4 : `label value too long` / `max label value length`

**Message** : `label value too long` ou `max label value length exceeded`.

**Ce que ça veut dire** : une valeur de label dépasse `max_label_value_length`
(défaut 2048 caractères). Typique quand on promeut un **champ libre**
(message d'erreur, URL, stack trace) en label via le stage `labels:`.

**Remède** : ne pas mettre ce champ en label. Le parser en champ
(`| json`, `| regexp`) ou en structured metadata suffit :

```yaml
    pipeline_stages:
      - json:
          expressions:
            error_msg: msg
      # ❌ INTERDIT :
      # - labels:
      #     error_msg:
      # ✅ autorisé :
      - structured_metadata:
          error_msg:
```

Si la limite bloque un label légitime (hostname très long…), augmentez-la
avec parcimonie :
```yaml
limits_config:
  max_label_value_length: 4096
```

---

## 75. Erreur 5 : `schema mismatch` / erreurs de schéma

**Message** : `schema mismatch: schema config for ... not found` ou requêtes
vides sur certaines périodes après un changement de config.

**Ce que ça veut dire** : Loki ne trouve pas de schéma couvrant la période
demandée. Cause quasi unique : une entrée `schema_config` a été **modifiée**
ou **supprimée** au lieu d'ajouter une nouvelle entrée.

**Règle d'or** (répétée section 66) :
> Les entrées `schema_config` sont **immuables et cumulatives**.
> On ajoute, on ne modifie jamais, on ne supprime qu'après expiration
> complète de la rétention.

**Remède** : restaurer l'entrée d'origine depuis la sauvegarde de
`/etc/loki/loki.yaml`, redémarrer, vérifier :
```bash
curl -s http://localhost:3100/loki/api/v1/status/buildinfo
# + tester une requête sur l'ancienne période
```

---

## 76. Erreur 6 : HTTP 429 `too many requests`

**Message** : le push Promtail reçoit `429 Too Many Requests`.

**Ce que ça veut dire** : le **rate limiting global** (`ingestion_rate_mb`)
est dépassé. Contrairement à l'erreur 3 (par flux), ici c'est le **tenant
entier** qui pousse trop vite.

**Diagnostic** :
```logql
# Débit d'ingestion actuel vs limite
sum(rate(loki_distributor_bytes_received_total[5m])) / 1024 / 1024
# → comparer à ingestion_rate_mb (défaut 16 Mo/s en monolithique)
```

**Remèdes** :
1. Vérifier qu'il ne s'agit pas d'un pic temporaire (redémarrage massif).
2. Augmenter si le dimensionnement le permet :
   ```yaml
   limits_config:
     ingestion_rate_mb: 32
     ingestion_burst_size_mb: 64
   ```
3. Si le 429 est **constant** : c'est un problème de capacité, pas de
   limite — scaler le stockage et les ingesters.

---

## 77. Erreur 7 : explosion de cardinalité

**Symptômes** : RAM de l'ingester en croissance continue, `loki_ingester_memory_streams`
qui ne redescend jamais, OOM-kill, requêtes qui timeout.

**Diagnostic** :
```bash
# Nombre de flux actifs
curl -s http://localhost:3100/metrics | grep loki_ingester_memory_streams
# Séries par requête rejetées ?
curl -s http://localhost:3100/metrics | grep loki_discarded_samples_total
```

**Plan d'action** :
1. **Stopper l'hémorragie** : baisser temporairement `max_streams_per_user`
   (les nouveaux flux sont rejetés, l'existant continue).
2. **Identifier** : quel job a explosé ?
   ```logql
   sum by (job) (count_over_time({job=~".+"}[15m]))
   ```
   puis inspecter les labels de ce job (`/loki/api/v1/label/.../values`).
3. **Corriger** : retirer le label à forte cardinalité côté Promtail
   (section 27), redéployer.
4. **Prévenir** : `max_label_names_per_series`, revue des `pipeline_stages`
   à chaque nouveau job.

---

## 78. Erreur 8 : `query timeout`

**Message** : `query timeout` (ou timeout côté Grafana après 60-120 s).

**Ce que ça veut dire** : le querier n'a pas fini dans le temps imparti
(`query_timeout`, défaut 5 min côté Loki ; souvent moins côté proxy/Grafana).

**Checklist** :
1. La requête est-elle raisonnable ? (plage, sélecteur — section 39)
2. `query_timeout` et `max_query_length` cohérents ?
   ```yaml
   limits_config:
     query_timeout: 5m
     max_query_length: 721h
   ```
3. Le query-frontend découpe-t-il ? (`split_queries_by_interval: 30m`)
4. Les queriers ont-ils assez de CPU/RAM ? (requêtes = CPU + décompression)
5. Côté nginx : `proxy_read_timeout 300s;` (section 57)

