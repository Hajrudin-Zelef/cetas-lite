---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-14
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "latency", "memory"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [2926, 3111]
sha256: fea86d1cbffe2d3a93710ae1a9e0cf9c8831c624e033c79e83cddb108fbbc819
---

# Guide Prometheus — Supervision métrique complète

```yaml
- alert: PortSwitchDown
  expr: ifOperStatus == 2 and ifAdminStatus == 1
  for: 15m
  labels: {severity: warning}
  annotations:
    summary: "Port {{ $labels.ifDescr }} DOWN sur {{ $labels.instance }}"
```

> Exclure les ports d'accès volontairement vides via un label ou un silence
> permanent documenté — sinon : bruit.

## 96. Cas pratique 3 : supervision des onduleurs du site

Objectif : ne plus jamais découvrir une coupure par les utilisateurs.

1. Activer SNMP sur chaque onduleur (v3 authPriv, section 56).
2. Job `snmp-ups` + job `blackbox-icmp` vers les IP des onduleurs
   (si le ping ne répond plus, l'onduleur est injoignable : alerte critique).
3. Alertes : `UpsSurBatterie` (2 min), `UpsBatterieFaible`, `UpsBatterieAChanger`
   (section 69).

Dashboard "Énergie" :

```promql
xupsBatteryCapacity or upsAdvBatteryCapacity                 # % batterie
xupsOutputLoad or upsAdvOutputLoad                            # % charge
xupsBatteryRunTimeRemaining or upsAdvBatteryRunTimeRemaining  # autonomie (min)
xupsBatteryTemperature                                       # température
```

Scénario à tester **trimestriellement** : coupure secteur simulée → vérifier
que l'alerte `UpsSurBatterie` arrive en < 3 min sur le canal d'astreinte, et
que l'autonomie affichée correspond au besoin (lien avec le plan de continuité).

## 97. Cas pratique 4 : supervision d'une application web

Objectif : SLI/SLO sur une application : disponibilité, latence, erreurs.

Instrumentation (exemple Python/prometheus_client) :

```python
from prometheus_client import Counter, Histogram, start_http_server

REQUESTS = Counter("app_requests_total", "Requêtes", ["route", "status"])
LATENCY = Histogram("app_request_duration_seconds", "Latence", ["route"],
                    buckets=[0.05, 0.1, 0.25, 0.5, 1, 2.5, 5])

start_http_server(8000)   # endpoint /metrics pour Prometheus
```

SLI en PromQL :

```promql
# Disponibilité 30 j (objectif 99,9 %)
100 * sum(rate(app_requests_total{status!~"5.."}[30d]))
    / sum(rate(app_requests_total[30d]))

# Latence p95 par route
histogram_quantile(0.95, sum by (route, le) (rate(app_request_duration_seconds_bucket[5m])))

# Budget d'erreur restant (SLO 99,9 % sur 30 j)
1 - (sum(rate(app_requests_total{status=~"5.."}[30d])) / sum(rate(app_requests_total[30d])))
```

Alerte "budget d'erreur qui fond vite" (multi-fenêtres, anti-bruit) :

```yaml
- alert: BudgetErreurCritique
  expr: |
    sum(rate(app_requests_total{status=~"5.."}[1h])) / sum(rate(app_requests_total[1h])) > 0.005
    and
    sum(rate(app_requests_total{status=~"5.."}[5m])) / sum(rate(app_requests_total[5m])) > 0.005
  for: 10m
  labels: {severity: critical}
```

## 98. Cas pratique 5 : capacity planning trimestriel

Objectif : répondre à "quand faudra-t-il acheter du disque / de la RAM ?".

```promql
# Croissance disque : Go consommés par jour, par serveur (sur 30 j)
-1 * deriv(node_filesystem_avail_bytes{fstype!~"tmpfs|overlay"}[30d]) * 86400 / 1e9

# Jours restants avant saturation (tendance 30 j)
node_filesystem_avail_bytes{fstype!~"tmpfs|overlay"} /
  (-1 * deriv(node_filesystem_avail_bytes[30d]) * 86400)

# Top 10 des croissances les plus rapides
topk(10,
  -1 * deriv(node_filesystem_avail_bytes{fstype!~"tmpfs|overlay|squashfs"}[30d]) * 86400 / 1e9
)
```

Méthode :

1. Exporter le résultat chaque trimestre (API `/api/v1/query` → CSV).
2. Croiser avec les projets à venir (nouvelle appli = +X Go).
3. Présenter : "serveur X saturé dans 4 mois au rythme actuel → prévoir extension
   au budget N+1".

> C'est ce cas d'usage qui justifie la rétention 90 j – 1 an sur les métriques
> d'infrastructure (ou un Thanos/Mimir, section 77-78).

## 99. Checklist de mise en production

### Installation
- [ ] Binaires officiels, checksums vérifiés
- [ ] Utilisateurs dédiés non-root
- [ ] Services systemd avec `Restart=always` et `MemoryMax`
- [ ] `promtool check config` OK avant chaque reload
- [ ] `--web.enable-lifecycle` actif (reload sans restart)

### Configuration
- [ ] `scrape_timeout` < `scrape_interval` partout
- [ ] `external_labels` renseignés (`datacenter`, `replica`)
- [ ] `sample_limit` / `target_limit` en garde-fous
- [ ] `metric_relabel_configs` : drop de `go_*`, `process_*` si inutiles
- [ ] `honor_labels: true` sur la fédération

### Alerting
- [ ] Alertmanager en cluster (si paire Prometheus)
- [ ] Chaque alerte a `severity`, `for`, `summary`, `runbook`
- [ ] Routage testé (`amtool config routes test`)
- [ ] Inhibitions configurées (panne site → pas de spam)
- [ ] Silences documentés pour les maintenances
- [ ] Canaux testés de bout en bout (email + Telegram/webhook)

### Données & résilience
- [ ] Rétention dimensionnée (calcul section 75)
- [ ] Snapshots quotidiens + copie hors machine (3-2-1)
- [ ] Paire HA ou plan de restauration testé
- [ ] Méta-monitoring : `PrometheusDown` supervisé de l'extérieur

### Sécurité
- [ ] Écoute locale ou reverse proxy TLS + auth
- [ ] Secrets dans des fichiers (`password_file`), perms 640
- [ ] SNMPv3 authPriv sur les équipements
- [ ] API admin et `/-/quit` restreintes par IP

### Exploitation
- [ ] Dashboard "Santé supervision" (targets, séries, règles)
- [ ] Revue trimestrielle des alertes (bruit → seuils)
- [ ] Exercice : panne simulée 1×/trimestre
- [ ] Documentation : où sont les configs, comment reloader, qui est d'astreinte

## 100. Pense-bête de poche PromQL

```promql
# ── Santé ──────────────────────────────────────────────
up == 0                                                    # cibles down
count by (job) (up == 0)                                   # down par job

# ── CPU ────────────────────────────────────────────────
100 * (1 - avg by (instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])))
node_load5 > 2 * count by (instance) (node_cpu_seconds_total{mode="idle"})

# ── Mémoire ────────────────────────────────────────────
100 * (1 - node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)

# ── Disque ─────────────────────────────────────────────
100 * (1 - node_filesystem_avail_bytes / node_filesystem_size_bytes)
100 * (1 - node_filesystem_files_free / node_filesystem_files)   # inodes !
predict_linear(node_filesystem_avail_bytes[6h], 3*24*3600) < 0

# ── Réseau ─────────────────────────────────────────────
rate(node_network_receive_bytes_total{device!~"lo|veth.*|docker.*"}[5m]) * 8 / 1e6
increase(node_network_receive_errors_total[1h]) > 0

# ── Compteurs génériques ───────────────────────────────
rate(metrique_total[5m])        # débit/sec
increase(metrique_total[1h])    # total sur 1 h
irate(metrique_total[5m])       # débit instantané

# ── Histogrammes ───────────────────────────────────────
histogram_quantile(0.95, sum by (le) (rate(http_request_duration_seconds_bucket[5m])))

# ── Temps ──────────────────────────────────────────────
metrique offset 1h             # il y a 1 h
rate(metrique_total[5m] offset 7d)   # semaine dernière
avg_over_time(metrique[1h])    # moyenne glissante

# ── Absence ────────────────────────────────────────────
absent(up{job="backup"})        # le job n'émet plus rien

