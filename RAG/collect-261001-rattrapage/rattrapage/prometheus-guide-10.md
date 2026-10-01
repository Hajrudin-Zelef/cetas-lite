---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-10
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention", "datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [2048, 2261]
sha256: e5a52ce4266aab3e859c353f8265232ce33bc59c8a292f3f95574b00133107da
---

# Guide Prometheus — Supervision métrique complète

      - alert: DisqueCritique
        expr: |
          100 * (1 - node_filesystem_avail_bytes{fstype!~"tmpfs|overlay|squashfs"}
                     / node_filesystem_size_bytes) > 95
        for: 10m
        labels: {severity: critical, equipe: infra}
        annotations:
          summary: "Disque {{ $labels.mountpoint }} à {{ $value | printf \"%.1f\" }} % sur {{ $labels.instance }}"

      - alert: DisqueSatureSous3Jours
        expr: predict_linear(node_filesystem_avail_bytes{fstype!~"tmpfs|overlay|squashfs"}[6h], 3*24*3600) < 0
        for: 1h
        labels: {severity: warning, equipe: infra}
        annotations:
          summary: "Tendance : disque {{ $labels.mountpoint }} plein sous 3 j sur {{ $labels.instance }}"

      - alert: InodesEpuises
        expr: 100 * (1 - node_filesystem_files_free / node_filesystem_files) > 85
        for: 30m
        labels: {severity: warning, equipe: infra}
        annotations:
          summary: "Inodes {{ $labels.mountpoint }} à {{ $value | printf \"%.1f\" }} % sur {{ $labels.instance }}"

      # --- Mémoire / CPU ---
      - alert: MemoireCritique
        expr: node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes < 0.1
        for: 10m
        labels: {severity: critical, equipe: infra}
        annotations:
          summary: "RAM disponible < 10 % sur {{ $labels.instance }}"

      - alert: ChargeElevee
        expr: node_load5 > 2 * count by (instance) (node_cpu_seconds_total{mode="idle"})
        for: 15m
        labels: {severity: warning, equipe: infra}
        annotations:
          summary: "Load5 ({{ $value | printf \"%.1f\" }}) > 2× CPU sur {{ $labels.instance }}"

      # --- Certificats ---
      - alert: CertificatExpireBientot
        expr: (probe_ssl_earliest_cert_expiry - time()) / 86400 < 21
        for: 1h
        labels: {severity: warning, equipe: infra}
        annotations:
          summary: "Certificat {{ $labels.instance }} expire dans {{ $value | printf \"%.0f\" }} j"

      # --- Onduleurs ---
      - alert: UpsSurBatterie
        expr: xupsInputSource == 2 or upsAdvInputLineStatus == 2
        for: 2m
        labels: {severity: critical, equipe: energie}
        annotations:
          summary: "Onduleur {{ $labels.instance }} SUR BATTERIE !"
          description: "Coupure secteur détectée. Vérifier groupe électrogène."

      - alert: UpsBatterieFaible
        expr: xupsBatteryCapacity < 30 or upsAdvBatteryCapacity < 30
        for: 5m
        labels: {severity: critical, equipe: energie}
        annotations:
          summary: "Batterie onduleur {{ $labels.instance }} à {{ $value }} %"

      - alert: UpsBatterieAChanger
        expr: xupsBatteryNeedsReplacement == 1 or upsAdvBatteryReplaceIndicator == 1
        for: 1h
        labels: {severity: warning, equipe: energie}
        annotations:
          summary: "Batterie à remplacer sur {{ $labels.instance }} — planifier intervention"
```

## 70. Règles d'alerting : les 10 erreurs à ne pas commettre

| # | Erreur | Correction |
|---|---|---|
| 1 | Pas de `for:` | Toujours `for: 5m` minimum |
| 2 | Seuil sur compteur brut (`http_requests_total > 1000`) | `rate(...[5m]) > X` |
| 3 | Fenêtre trop courte (`rate(...[1m])`) | `[5m]` minimum en alerte |
| 4 | Alerte sans label `severity` | Le routage en dépend |
| 5 | Pas de `summary`/`description` | Annotations obligatoires |
| 6 | Même seuil warning/critical | warning < critical, `for` plus long en warning |
| 7 | Alerter sur `up` sans `for` | Un scrape raté isolé ≠ panne |
| 8 | Oublier `or` entre deux MIB (Eaton/APC) | Les deux variantes dans l'expr |
| 9 | Comparaison UTC non décalée | Utiliser les time_intervals d'Alertmanager |
| 10 | Jamais tester | `promtool test rules` + exercice trimestriel |

---

---

# PARTIE V — Fédération, stockage distant, sauvegarde, HA

---

## 71. Fédération : agréger plusieurs Prometheus

La **fédération** permet à un Prometheus "global" de scraper les métriques
(agrégées) d'autres Prometheus "locaux" (par site, par datacenter).

Sur le Prometheus **global** :

```yaml
  - job_name: 'federate-dc-paris'
    honor_labels: true            # garder job/instance d'origine !
    metrics_path: '/federate'
    params:
      'match[]':
        - '{job="node-linux"}'           # quelles séries rapatrier…
        - '{__name__=~"job:.*"}'         # …dont les recording rules
        - '{__name__=~"instance:.*"}'
    static_configs:
      - targets: ['prom-dc-paris.lan:9090']
```

Bonnes pratiques :

- Ne fédérer que des **métriques agrégées** (recording rules `job:...`,
  `instance:...`), jamais les séries brutes → sinon le global explose.
- `honor_labels: true` **obligatoire** (sinon tout devient `job="federate"`).
- `external_labels: {datacenter: 'dc-paris'}` côté local pour distinguer.

```
 [Prometheus dc-paris] ──┐
 [Prometheus dc-lyon]  ──┼──federate──► [Prometheus GLOBAL] ──► Grafana
 [Prometheus usine]    ──┘                    │
                                              └──► Alertmanager global
```

## 72. Remote write : envoyer vers un stockage distant

`remote_write` envoie les échantillons **au fil de l'eau** vers un stockage
long terme (Thanos, Mimir/Cortex, VictoriaMetrics, InfluxDB…) :

```yaml
remote_write:
  - url: 'https://thanos-receive.lan/api/v1/receive'
    name: 'thanos-central'
    queue_config:
      capacity: 10000
      max_shards: 50
      min_shards: 1
      max_samples_per_send: 5000
      batch_send_deadline: 5s
    write_relabel_configs:       # n'envoyer que l'utile !
      - source_labels: [__name__]
        regex: 'go_.*|process_.*'
        action: drop
```

Points d'attention :

- `write_relabel_configs` : **filtrer avant envoi** (coût réseau/stockage).
- Le remote_write a un **buffer local** (WAL) : en cas de coupure réseau,
  il rejoue (jusqu'à la taille du buffer, puis il jette).
- Métrique de santé : `prometheus_remote_storage_shards`,
  `prometheus_remote_storage_samples_dropped_total`.

## 73. Remote read : lire depuis le stockage distant

```yaml
remote_read:
  - url: 'https://thanos-query.lan/api/v1/read'
    name: 'thanos-central'
    read_recent: true     # les données récentes sont lues en local (plus rapide)
```

`read_recent: true` : Prometheus lit d'abord sa TSDB locale pour le récent,
puis le distant pour l'ancien. Sans ça, tout passe par le réseau.

⚠️ Remote read = **dépendance réseau** pour Grafana. En cas de coupure,
prévoir un dashboard "local only" ou accepter la dégradation.

## 74. Sauvegarde : snapshots TSDB via l'API

Activer l'API admin au démarrage :

```bash
--web.enable-admin-api
```

Créer un snapshot (cohérent, sans arrêter Prometheus) :

```bash
curl -X POST http://localhost:9090/api/v1/admin/tsdb/snapshot
# {"status":"success","data":{"name":"20260926T233000Z-abcdef1234"}}
# -> snapshot dans /var/lib/prometheus/snapshots/20260926T233000Z-abcdef1234
```

Script de sauvegarde quotidienne (`/usr/local/bin/backup-prometheus.sh`) :

```bash
#!/bin/bash
set -euo pipefail
PROM="http://127.0.0.1:9090"
DEST="/backup/prometheus"
RETENTION_JOURS=7

SNAP=$(curl -s -X POST "$PROM/api/v1/admin/tsdb/snapshot" | grep -o '[0-9TZ-]*-[a-f0-9]*' | head -1)
[ -z "$SNAP" ] && { echo "ERREUR: snapshot vide"; exit 1; }

# Compresser + chiffrer éventuellement, puis copier hors machine
tar -czf "$DEST/prometheus-$SNAP.tar.gz" -C /var/lib/prometheus "snapshots/$SNAP"

# Nettoyer le snapshot source (il reste sur disque sinon !)
curl -s -X POST "$PROM/api/v1/admin/tsdb/delete_series?match[]={__name__=~\".+\"}" > /dev/null || true
# ^ non : delete_series ne supprime pas les snapshots. Nettoyage manuel :
rm -rf "/var/lib/prometheus/snapshots/$SNAP"

# Purge des vieilles archives
find "$DEST" -name 'prometheus-*.tar.gz' -mtime +$RETENTION_JOURS -delete
echo "Sauvegarde OK : $DEST/prometheus-$SNAP.tar.gz"
```

⚠️ **Deux pièges** :

