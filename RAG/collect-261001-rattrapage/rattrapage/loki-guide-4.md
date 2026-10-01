---
id: collect-261001-rattrapage/rattrapage/loki-guide-4
title: "Grafana Loki — Le guide complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["apache", "attention", "latency"]
source: docs/RAG/collect-261001-rattrapage/loki_guide.md
source_anchor: ""
source_lines: [706, 1008]
sha256: 53ed6c781936fdcf710f1f25a8e7041295e43132543cea54caf84fa3e9939054
---

# Grafana Loki — Le guide complet

[Service]
Type=simple
User=promtail
Group=promtail
ExecStart=/usr/local/bin/promtail -config.file=/etc/promtail/promtail.yaml
Restart=on-failure
RestartSec=5
NoNewPrivileges=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/var/lib/promtail /tmp
# Lecture des logs système :
ReadOnlyPaths=/var/log /var/lib/docker/containers
PrivateTmp=true
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
```

> ⚠️ Avec `ProtectSystem=strict`, listez explicitement les chemins de logs
> dans `ReadOnlyPaths`, sinon Promtail ne verra aucun fichier.

---

## 17. Promtail : configuration de base

`/etc/promtail/promtail.yaml` minimal et fonctionnel :

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /var/lib/promtail/positions.yaml   # où Promtail en est dans chaque fichier

clients:
  - url: http://loki.local:3100/loki/api/v1/push
    # batchwait / batchsize : compromis latence/débit
    batchwait: 1s
    batchsize: 1048576   # 1 Mo

scrape_configs:
  - job_name: system
    static_configs:
      - targets: [localhost]
        labels:
          job: syslog
          host: srv01
          __path__: /var/log/syslog
```

Points clés :

- `positions.yaml` : **ne jamais le supprimer en production** sauf pour
  forcer une relecture complète (il mémorise l'offset de chaque fichier).
- `batchwait` : Promtail attend jusqu'à 1 s pour regrouper les lignes avant
  d'envoyer. Baissez à `500ms` si la latence d'ingestion compte.
- Un `clients` peut pointer vers plusieurs Loki (HA) : ajoutez des entrées.

---

## 18. `scrape_configs` et `static_configs`

Chaque `scrape_config` définit **quoi lire, où l'envoyer, quels labels
ajouter**. C'est le pendant des `scrape_configs` de Prometheus.

```yaml
scrape_configs:
  # 1. Logs système
  - job_name: syslog
    static_configs:
      - targets: [localhost]
        labels:
          job: syslog
          host: srv01
          __path__: /var/log/syslog

  # 2. Plusieurs fichiers dans un seul job (globs)
  - job_name: nginx
    static_configs:
      - targets: [localhost]
        labels:
          job: nginx
          host: srv01
          __path__: /var/log/nginx/*.log

  # 3. Plusieurs chemins avec des labels différents
  - job_name: apps
    static_configs:
      - targets: [localhost]
        labels:
          job: app-paiement
          env: prod
          __path__: /opt/paiement/logs/app.log
      - targets: [localhost]
        labels:
          job: app-facturation
          env: prod
          __path__: /opt/facturation/logs/app.log

  # 4. Journal systemd (journald)
  - job_name: journal
    journal:
      path: /var/log/journal
      max_age: 12h
      labels:
        job: systemd-journal
        host: srv01
    relabel_configs:
      - source_labels: ['__journal__systemd_unit']
        target_label: unit
      - source_labels: ['__journal_priority_keyword']
        target_label: level
```

---

## 19. `__path__` : globs, pièges et bonnes pratiques

`__path__` accepte les globs (`*`, `**`) mais quelques règles évitent les
mauvaises surprises :

| Règle | Exemple | Pourquoi |
|---|---|---|
| Un job = un type de log | `job: nginx` → `/var/log/nginx/*.log` | labels cohérents |
| Éviter `**` trop large | ❌ `/var/log/**/*.log` | ramasse tout, labels incohérents |
| Fichiers stables | ✔ `/var/log/nginx/access.log` | la rotation logrotate garde le nom |
| Attention aux liens symboliques | vérifier `follow_symlinks` | par défaut Promtail ne suit pas toujours |

**Rotation des logs (logrotate)** : Promtail gère nativement la rotation
par renommage (`access.log` → `access.log.1`) grâce au fichier `positions`.
Configurez logrotate en mode `copytruncate` **ou** `create` — les deux
fonctionnent, mais `copytruncate` peut dupliquer quelques lignes lors de
la copie (acceptable pour des logs).

```bash
# /etc/logrotate.d/nginx — exemple compatible Promtail
/var/log/nginx/*.log {
    daily
    rotate 14
    compress
    delaycompress
    missingok
    notifempty
    create 0640 www-data adm
    sharedscripts
    postrotate
        [ -f /var/run/nginx.pid ] && kill -USR1 $(cat /var/run/nginx.pid)
    endscript
}
```

---

## 20. Parser logfmt

Le format **logfmt** (`clé=valeur`) est le format structuré le plus simple.
Beaucoup d'applications modernes (et Loki lui-même) l'utilisent.

Ligne d'exemple :

```
ts=2026-09-26T10:00:01Z level=error msg="connexion refusée" host=db01 retries=3
```

Pipeline Promtail :

```yaml
scrape_configs:
  - job_name: monapp
    static_configs:
      - targets: [localhost]
        labels:
          job: monapp
          __path__: /opt/monapp/app.log
    pipeline_stages:
      - logfmt:
          mapping:
            ts: timestamp
            level: level
            msg: message
      - labels:
          level: level        # promeut 'level' en vrai label Loki
      - timestamp:
          source: timestamp
          format: RFC3339
      - output:
          source: message
```

Ce que fait chaque étape :
1. `logfmt` : découpe la ligne en champs nommés.
2. `labels` : `level` devient un **label indexé** (pratique pour
   `{job="monapp", level="error"}`).
3. `timestamp` : utilise le timestamp du log plutôt que l'heure de réception.
4. `output` : remplace la ligne brute par le champ `message` (optionnel).

> ⚠️ Ne promeuvez en labels que des champs à **faible cardinalité**
> (`level`, `host`). Jamais `msg`, `retries` ou un ID.

---

## 21. Parser JSON

Pour les applications qui loguent en JSON (de plus en plus courant) :

```json
{"ts":"2026-09-26T10:00:01Z","level":"error","msg":"timeout","service":"api","latency_ms":1500}
```

```yaml
    pipeline_stages:
      - json:
          expressions:
            timestamp: ts
            level: level
            message: msg
            service: service
      - labels:
          level:
          service:
      - timestamp:
          source: timestamp
          format: RFC3339
      - output:
          source: message
```

Variante **sans réécriture** (on garde la ligne JSON brute, utile pour
`json` dans les requêtes LogQL à la volée — voir section 32) :

```yaml
    pipeline_stages:
      - json:
          expressions:
            level: level
      - labels:
          level:
      # pas de stage 'output' : la ligne reste le JSON d'origine
```

**Conseil** : parser à l'ingestion (Promtail) = requêtes plus rapides et
moins chères. Parser à la requête (`| json` en LogQL) = flexibilité sans
redéployer Promtail. Pour les champs stables et critiques (`level`,
`service`), parsez à l'ingestion.

---

## 22. Parser regex

Quand le format est fixe mais ni JSON ni logfmt (ex : logs Apache/nginx
personnalisés), la regex fait le travail.

Format d'exemple (nginx) :

```
192.168.1.10 - - [26/Sep/2026:10:00:01 +0000] "GET /api/users HTTP/1.1" 500 1234 "-" "curl/8.0"
```

```yaml
    pipeline_stages:
      - regex:
          expression: '^(?P<ip>\S+) \S+ \S+ \[(?P<time>[^\]]+)\] "(?P<method>\S+) (?P<path>\S+) \S+" (?P<status>\d{3}) (?P<size>\d+|-)'
      - labels:
          method:
          status:
      - timestamp:
          source: time
          format: "02/Jan/2006:15:04:05 -0700"
```

> 💡 Testez vos regex avec `|~` en LogQL ou sur regex101.com **avant** de
> les déployer. Une regex qui ne matche pas = des champs vides silencieux.
> Ajoutez un stage de debug temporaire :
>
> ```yaml
>       - template:
>           source: debug
>           template: 'ip={{ .ip }} status={{ .status }}'
>       - output:
>           source: debug
> ```

---

## 23. Multiline : gérer les stack traces Java et Python

Par défaut, Promtail envoie **une ligne = un log**. Les stack traces
(Java, Python) s'étalent sur plusieurs lignes : il faut les **regrouper**
avec le stage `multiline`.

Exemple Python :

```
2026-09-26 10:00:01 ERROR app: échec du traitement
Traceback (most recent call last):
  File "/opt/app/main.py", line 42, in run
    process()
ValueError: invalid literal
```

