---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-5
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [876, 1092]
sha256: 94102e4852cd4612de1e61db1c5038111ee9574b9eb33a0d08752ba76cb3c033
---

# Guide Prometheus — Supervision métrique complète

⚠️ `delta()`/`deriv()` sont pour les **gauges**. Sur un compteur, utilisez
`increase()`/`rate()`.

## 35. histogram_quantile() et les latences

```promql
# p50, p95, p99 de la latence HTTP par route, sur 5 min
histogram_quantile(0.50, sum by (route, le) (rate(http_request_duration_seconds_bucket[5m])))
histogram_quantile(0.95, sum by (route, le) (rate(http_request_duration_seconds_bucket[5m])))
histogram_quantile(0.99, sum by (route, le) (rate(http_request_duration_seconds_bucket[5m])))
```

Points critiques :

1. **Toujours** `sum by (le, ...)` **avant** `histogram_quantile` (agréger les buckets).
2. Le label `le` doit être conservé dans le `by`.
3. `rate()` **à l'intérieur**, sur les buckets (ce sont des compteurs).

```promql
# Taux de requêtes lentes (> 1 s) : sans quantile, souvent plus parlant
sum by (route) (rate(http_request_duration_seconds_bucket{le="+Inf"}[5m])
-
sum by (route) (rate(http_request_duration_seconds_bucket{le="1"}[5m]))
```

## 36. Fonctions temporelles

```promql
time()                          # timestamp unix actuel
hour()                          # heure UTC (0-23) — attention au fuseau !
day_of_week()                   # 0=dimanche … 6=samedi (UTC)
days_in_month()
month()  year()  minute()
timestamp(metrique)              # timestamp du dernier échantillon
```

⚠️ Ces fonctions sont en **UTC**. Pour des alertes "heures ouvrées" en Europe/Paris,
décaler manuellement : `hour(time() + 2*3600)` (été) — fragile aux changements
d'heure. Mieux : gérer les horaires dans Alertmanager (time_intervals).

```promql
# Âge des données : détecter un exporter figé
time() - timestamp(node_time_seconds)
```

## 37. Manipulation de labels : label_replace, label_join

```promql
# Extraire le nom court depuis instance="srv-web01:9100" -> host="srv-web01"
label_replace(up, "host", "$1", "instance", "([^:]+):.*")

# Combiner deux labels : src + dst -> paire="src->dst"
label_join(probe_success, "paire", "->", "src", "dst")

# Renommer proprement pour Grafana
label_replace(
  node_filesystem_avail_bytes,
  "serveur", "$1", "instance", "([^:]+):.*"
)
```

Signature : `label_replace(vecteur, "dst", "$1", "src", "regex")`.
Si la regex ne matche pas, le label destination n'est pas créé (pas d'erreur).

## 38. absent() : détecter ce qui n'existe pas

```promql
# Retourne 1 si AUCUNE série ne matche (ex. : plus aucun worker en vie)
absent(up{job="workers"})

# Alerte "aucune sauvegarde réussie depuis 26 h"
absent_over_time(backup_success[26h])

# Combiner : alerter si le job de backup n'émet plus RIEN
absent(up{job="backup"}) or (up{job="backup"} == 0)
```

⚠️ `absent()` retourne un vecteur **vide** (pas de 0) quand des séries existent :
c'est son intérêt (rien à alerter) mais ça surprend dans les dashboards.

## 39. Agrégations : sum, avg, min, max, count, stddev…

```promql
sum by (datacenter) (rate(http_requests_total[5m]))   # trafic par DC
avg by (job) (node_load1)                             # charge moyenne par job
max by (instance) (node_cpu_temperature)              # point chaud par serveur
count by (job) (up)                                   # nombre de cibles par job
count by (job) (up == 0)                              # nombre de cibles DOWN par job
stddev by (service) (latence)                         # dispersion
```

- `by (labels)` : groupe **en gardant** ces labels (tout le reste est jeté).
- `without (labels)` : groupe **en jetant** ces labels (garde tout le reste).

```promql
# Équivalents :
sum by (datacenter, job) (rate(http_requests_total[5m]))
sum without (instance, cpu, mode) (rate(http_requests_total[5m]))
```

⚠️ `without` est dangereux si un nouveau label à haute cardinalité apparaît :
préférez `by` explicite dans les recording rules.

## 40. topk(), bottomk(), count_values(), group()

```promql
topk(5, 100 * (1 - node_filesystem_avail_bytes / node_filesystem_size_bytes))
# -> les 5 filesystems les plus pleins

bottomk(3, node_memory_MemAvailable_bytes)
# -> les 3 serveurs avec le moins de RAM libre

count_values("instances", up{job="node-linux"})
# -> nombre de cibles par valeur de up : {instances="42"} 1, {instances="3"} 0

group by (datacenter) (up)
# -> 1 par datacenter ayant au moins une cible (test d'existence)
```

⚠️ `topk`/`bottomk` : le nombre de séries retournées peut varier → compliqué
pour l'alerting (préférer un seuil avec `>`).

## 41. Subqueries : des requêtes dans des requêtes

Permettent d'appliquer une fonction de range sur le **résultat** d'une requête :

```promql
# Dérivée du débit sur 1 h, calculée sur des points à 5 min
deriv(rate(http_requests_total[5m])[1h:5m])

# Moyenne sur 1 h du p95 glissant à 5 min
avg_over_time(
  histogram_quantile(0.95, sum by (le) (rate(http_request_duration_seconds_bucket[5m])))[1h:5m]
)

# Tendance : le trafic augmente-t-il depuis ce matin ?
rate(http_requests_total[5m])[8h:5m]
```

Syntaxe : `<requête>[durée:pas]`. Le pas (`5m`) doit être ≥ l'intervalle de scrape.

⚠️ Les subqueries sont **coûteuses** : à réserver aux dashboards occasionnels,
jamais dans des alertes évaluées toutes les 15 s.

## 42. Recording rules : pré-calculer pour aller vite

Les **recording rules** évaluent une requête à intervalle régulier et stockent
le résultat comme une **nouvelle métrique**. Indispensable pour :

- Accélérer les dashboards (requêtes lourdes pré-calculées),
- Simplifier les alertes,
- Réduire la cardinalité (agréger tôt).

`/etc/prometheus/rules/recording.yml` :

```yaml
groups:
  - name: node-recording
    interval: 30s
    rules:
      # CPU usage % par instance (pré-calculé)
      - record: instance:cpu_usage_percent:30s
        expr: |
          100 * (1 - avg by (instance) (
            rate(node_cpu_seconds_total{mode="idle"}[5m])
          ))

      # Trafic HTTP par service et par code
      - record: service:http_requests:rate5m
        expr: |
          sum by (service, status) (rate(http_requests_total[5m]))

      # Espace disque libre minimum par nœud
      - record: instance:disk_free_bytes:min
        expr: |
          min by (instance) (node_filesystem_avail_bytes{fstype!~"tmpfs|overlay|squashfs"})
```

Convention de nommage : `niveau:metrique:operation` (`job:`, `instance:`…).
Le suffixe indique la fenêtre si pertinent.

⚠️ Une recording rule **crée des séries** : ne pas en abuser sur des requêtes
à haute cardinalité (voir section 79).

## 43. Alerting rules : la syntaxe des alertes

`/etc/prometheus/rules/alertes.yml` :

```yaml
groups:
  - name: infra-alertes
    interval: 30s
    rules:
      - alert: InstanceDown
        expr: up == 0
        for: 5m                 # doit être vrai pendant 5 min avant de tirer
        labels:
          severity: critical
          equipe: infra
        annotations:
          summary: "Instance {{ $labels.instance }} DOWN (job {{ $labels.job }})"
          description: >
            La cible {{ $labels.instance }} ne répond plus depuis plus de 5 minutes.
            Runbook: https://wiki.lan/runbooks/instance-down
          dashboard: "https://grafana.lan/d/infra?var-instance={{ $labels.instance }}"
```

- `expr` : la condition PromQL (vecteur non vide = alerte).
- `for` : durée pendant laquelle la condition doit rester vraie → évite les
  alertes sur un micro-à-coup. États : `pending` → `firing`.
- `labels` : routage Alertmanager (severity, équipe…).
- `annotations` : texte humain (templating `{{ $labels.x }}`, `{{ $value }}`).

## 44. Tester et débugger ses requêtes

Méthode en 4 étapes :

