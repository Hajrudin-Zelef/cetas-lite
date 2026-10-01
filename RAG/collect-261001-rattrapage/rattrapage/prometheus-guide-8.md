---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-8
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [1589, 1840]
sha256: 5d424cd5911442f88bdb4f699609fd665c5151a8356acf7d9e0d64ab89275d6c
---

# Guide Prometheus — Supervision métrique complète

```yaml
# dans snmp.yml (généré), section auths :
auths:
  v3-securise:
    version: 3
    username: monitor
    security_level: authPriv
    auth_protocol: SHA
    auth_password: "xxx"
    priv_protocol: AES
    priv_password: "xxx"
```

```yaml
# prometheus.yml
  - job_name: 'snmp-ups'
    metrics_path: /snmp
    params:
      module: [ups_eaton]
      auth: [v3-securise]
    static_configs:
      - targets: ['192.168.10.20', '192.168.10.21']
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: '127.0.0.1:9116'
```

> **Lien avec le métier** : couplez ces métriques avec le guide onduleurs
> (seuils de floating, loi d'Arrhenius : +10 °C = durée de vie / 2). Une alerte
> `xupsBatteryTemperature > 30` prolongée justifie une visite technique.

## 57. Exporters SQL : MySQL et PostgreSQL

**mysqld_exporter** (`:9104`) :

```bash
# Fichier d'authentification (jamais le mot de passe en ligne de commande !)
sudo tee /etc/mysqld_exporter/.my.cnf > /dev/null <<'EOF'
[client]
user=exporter
password=MotDePasseLong
host=127.0.0.1
EOF
sudo chmod 600 /etc/mysqld_exporter/.my.cnf
sudo chown mysqld_exporter:mysqld_exporter /etc/mysqld_exporter/.my.cnf
```

```sql
-- Utilisateur MySQL dédié (privilèges minimaux)
CREATE USER 'exporter'@'localhost' IDENTIFIED BY 'MotDePasseLong';
GRANT PROCESS, REPLICATION CLIENT, SELECT ON *.* TO 'exporter'@'localhost';
```

```promql
# Connexions utilisées vs max
mysql_global_status_threads_connected / mysql_global_variables_max_connections

# Requêtes lentes (toujours > 0 à surveiller en tendance)
increase(mysql_global_status_slow_queries[1h])

# Replication : retard en secondes (0 = OK)
mysql_slave_status_seconds_behind_master

# InnoDB : lectures disque vs buffer pool
rate(mysql_global_status_innodb_buffer_pool_reads[5m])
```

**postgres_exporter** (`:9187`) — via `DATA_SOURCE_NAME` :

```bash
DATA_SOURCE_NAME="postgresql://exporter:mdp@localhost:5432/postgres?sslmode=disable"
```

```promql
# Connexions vs max
pg_stat_database_numbackends / pg_settings_max_connections

# Taux de hit du cache (doit rester > 99 %)
sum(rate(pg_stat_database_blks_hit[5m])) /
  (sum(rate(pg_stat_database_blks_hit[5m])) + sum(rate(pg_stat_database_blks_read[5m])))

# Transactions wraparound : âge max (alerte si > 1 milliard)
max(pg_stat_database_datfrozenxid_age)

# Locks en attente
pg_locks_count{mode="exclusive", granted="false"} > 0
```

---

---

# PARTIE IV — Alertmanager et l'alerting

---

## 58. Alertmanager : architecture et concepts

Prometheus **évalue** les règles et envoie les alertes **en état firing** à
Alertmanager. C'est Alertmanager qui gère tout le reste :

```
Prometheus ──► Alertmanager ──┬──► email (SMTP)
    ▲                        ├──► Telegram (bot API)
    │ (renvoie tant          ├──► webhook (GLPI, script…)
    │  que firing)           ├──► Slack / Teams / PagerDuty…
                             └──► (silences, inhibitions, regroupements)
```

Concepts :

- **Déduplication** : une alerte identique reçue 10 fois = 1 notification.
- **Regroupement (grouping)** : "42 instances down" → 1 seul message groupé,
  pas 42 SMS à 3 h du matin.
- **Routage (routes)** : arbre de décision selon les labels
  (`severity=critical` → astreinte, `severity=warning` → email équipe).
- **Inhibition** : "si le datacenter est down, ne pas m'alerter sur chaque serveur".
- **Silences** : coupure volontaire et tracée des notifications (maintenance).

⚠️ Alertmanager **ne génère aucune alerte** : sans règles côté Prometheus,
il ne se passe rien. Et Prometheus sans Alertmanager = alertes visibles
uniquement dans l'UI (inutile en production).

## 59. Alertmanager : installation

```bash
AM_VERSION="0.28.0"
cd /tmp
wget "https://github.com/prometheus/alertmanager/releases/download/v${AM_VERSION}/alertmanager-${AM_VERSION}.linux-amd64.tar.gz"
tar xzf "alertmanager-${AM_VERSION}.linux-amd64.tar.gz"
sudo cp "alertmanager-${AM_VERSION}.linux-amd64/alertmanager" /usr/local/bin/
sudo cp "alertmanager-${AM_VERSION}.linux-amd64/amtool" /usr/local/bin/
sudo mkdir -p /etc/alertmanager /var/lib/alertmanager
sudo cp "alertmanager-${AM_VERSION}.linux-amd64/alertmanager.yml" /etc/alertmanager/
sudo useradd --no-create-home --shell /usr/sbin/nologin alertmanager
sudo chown -R alertmanager:alertmanager /etc/alertmanager /var/lib/alertmanager
```

`/etc/systemd/system/alertmanager.service` :

```ini
[Unit]
Description=Alertmanager
Wants=network-online.target
After=network-online.target

[Service]
User=alertmanager
Group=alertmanager
Type=simple
Restart=always
ExecStart=/usr/local/bin/alertmanager \
  --config.file=/etc/alertmanager/alertmanager.yml \
  --storage.path=/var/lib/alertmanager \
  --web.listen-address=127.0.0.1:9093 \
  --cluster.listen-address=""     # vide = pas de cluster (voir section 76)

[Install]
WantedBy=multi-user.target
```

Vérification : `amtool check-config /etc/alertmanager/alertmanager.yml`
puis `curl -X POST http://127.0.0.1:9093/-/reload` après modification.

## 60. Configuration Alertmanager : squelette

`/etc/alertmanager/alertmanager.yml` :

```yaml
global:
  resolve_timeout: 5m            # "résolue" si plus reçue pendant 5 min
  smtp_smarthost: 'smtp.lan:587'
  smtp_from: 'alertes@entreprise.lan'
  smtp_auth_username: 'alertes@entreprise.lan'
  smtp_auth_password: 'xxx'       # idéalement via fichier (voir section 84)
  smtp_require_tls: true

# Arbre de routage (évalué de haut en bas, première route qui matche)
route:
  group_by: ['alertname', 'datacenter']
  group_wait: 30s                 # attend 30 s pour grouper les alertes liées
  group_interval: 5m              # nouvel envoi pour le même groupe toutes les 5 min
  repeat_interval: 4h             # répéter la notification toutes les 4 h si toujours firing
  receiver: 'equipe-infra-email'
  routes:
    - matchers:
        - severity = "critical"
      receiver: 'astreinte-telegram'
      continue: false
    - matchers:
        - severity = "warning"
      receiver: 'equipe-infra-email'

receivers:
  - name: 'equipe-infra-email'
    email_configs:
      - to: 'infra@entreprise.lan'
        headers:
          Subject: '[{{ .Status | toUpper }}] {{ .GroupLabels.alertname }}'

  - name: 'astreinte-telegram'
    telegram_configs:
      - bot_token: 'XXX'
        chat_id: -123456789
        message: |
          {{ if eq .Status "firing" }}🔥{{ else }}✅{{ end }}
          <b>{{ .GroupLabels.alertname }}</b>
          {{ range .Alerts }}{{ .Annotations.summary }}
          {{ end }}
        parse_mode: HTML
```

## 61. Le routage : arbre de décision

Les `routes` forment un arbre : chaque alerte descend jusqu'à la première
feuille qui matche. `continue: true` permet de continuer vers les routes sœurs.

```yaml
route:
  receiver: 'defaut'
  group_by: ['alertname']
  routes:
    # 1. Critiques infra -> Telegram astreinte + email
    - matchers: ['severity = "critical"', 'equipe = "infra"']
      receiver: 'astreinte'
      continue: true            # ...et aussi la route suivante
    - matchers: ['severity = "critical"']
      receiver: 'direction-email'
    # 2. Onduleurs -> équipe énergie (même en warning)
    - matchers: ['equipement =~ "ups.*"']
      receiver: 'equipe-energie'
    # 3. Heures ouvrées uniquement pour les warnings non critiques
    - matchers: ['severity = "warning"']
      receiver: 'equipe-infra-email'
      active_time_intervals: ['heures-ouvrees']

time_intervals:
  - name: heures-ouvrees
    time_intervals:
      - weekdays: ['monday:friday']
        times:
          - start_time: '08:00'
            end_time: '18:00'
```

> Les `time_intervals` (noms au pluriel dans la config !) permettent le
> "pas d'alerte warning la nuit" sans bidouille PromQL dépendante de l'UTC.

