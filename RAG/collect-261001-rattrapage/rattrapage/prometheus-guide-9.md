---
id: collect-261001-rattrapage/rattrapage/prometheus-guide-9
title: "Guide Prometheus — Supervision métrique complète"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/prometheus_guide.md
source_anchor: ""
source_lines: [1841, 2047]
sha256: 20e951fa3096f06073276427d34999c737d62e05dbfded2f8d7839b78dd46f27
---

# Guide Prometheus — Supervision métrique complète

## 62. Receiver email : bien le configurer

```yaml
receivers:
  - name: 'equipe-infra-email'
    email_configs:
      - to: 'infra@entreprise.lan'
        from: 'alertes@entreprise.lan'
        smarthost: 'smtp.lan:587'
        auth_username: 'alertes@entreprise.lan'
        auth_password_file: '/etc/alertmanager/smtp_pass'  # mieux que en clair
        require_tls: true
        headers:
          Subject: '[{{ .Status | toUpper }}:{{ .CommonLabels.severity }}] {{ .CommonLabels.alertname }} ({{ .Alerts | len }} instance(s))'
        html: |
          <h2>{{ .CommonLabels.alertname }} — {{ .Status }}</h2>
          <table>
          {{ range .Alerts }}
            <tr><td>{{ .Labels.instance }}</td><td>{{ .Annotations.summary }}</td></tr>
          {{ end }}
          </table>
          <p><a href="{{ .ExternalURL }}">Voir dans Alertmanager</a></p>
```

⚠️ Toujours inclure `{{ .Alerts | len }}` ou la liste des instances : un email
"InstanceDown" sans dire **laquelle** est inutile.

## 63. Receiver Telegram

```yaml
  - name: 'astreinte-telegram'
    telegram_configs:
      - bot_token_file: '/etc/alertmanager/telegram_token'  # fichier, pas en clair
        chat_id: -123456789        # ID du groupe (négatif) ou de l'utilisateur
        message: |
          {{ if eq .Status "firing" }}🔥 <b>ALERTE</b>{{ else }}✅ <b>RÉSOLUE</b>{{ end }}
          <b>{{ .CommonLabels.alertname }}</b> [{{ .CommonLabels.severity }}]
          {{ range .Alerts -}}
          • <code>{{ .Labels.instance }}</code> : {{ .Annotations.summary }}
          {{ end }}
          <a href="{{ .ExternalURL }}">Alertmanager</a>
        parse_mode: HTML
        disable_notifications: false
```

Créer le bot via `@BotFather`, récupérer le token, ajouter le bot au groupe,
récupérer le `chat_id` via `https://api.telegram.org/bot<TOKEN>/getUpdates`.

## 64. Receiver webhook : brancher GLPI, scripts, etc.

```yaml
  - name: 'glpi-webhook'
    webhook_configs:
      - url: 'https://glpi.lan/plugins/alertes/prometheus'
        http_config:
          basic_auth:
            username: 'prometheus'
            password_file: '/etc/alertmanager/glpi_pass'
        send_resolved: true
        max_alerts: 10
```

Le webhook reçoit un JSON avec `status`, `alerts[]` (labels, annotations,
`startsAt`, `endsAt`). Côté GLPI/script : créer un ticket à `firing`, le clore
à `resolved` (corrélation par `alertname` + `instance` = clé du ticket).

Exemple minimal de script récepteur (Python) :

```python
from flask import Flask, request
app = Flask(__name__)

@app.route("/hook", methods=["POST"])
def hook():
    data = request.json
    for alert in data["alerts"]:
        key = (alert["labels"]["alertname"], alert["labels"].get("instance", "?"))
        if data["status"] == "firing":
            creer_ticket(key, alert["annotations"].get("summary", ""))
        else:
            clore_ticket(key)
    return "ok"
```

## 65. Inhibition : couper le bruit en cascade

Si le switch d'un site tombe, inutile de recevoir 30 alertes "serveur down" :
**l'alerte parente inhibe les enfants**.

```yaml
inhibit_rules:
  # Si un site est injoignable, inhiber les alertes unitaires de ce site
  - source_matchers:
      - alertname = "SiteDown"
    target_matchers:
      - severity =~ "critical|warning"
    equal: ['site']              # même site des deux côtés

  # Si Prometheus ne scrape plus, inhiber les alertes "métier" basées sur le scrape
  - source_matchers:
      - alertname = "PrometheusScrapeDown"
    target_matchers:
      - alertname =~ "Instance.*"
    equal: ['job']
```

Logique : `source` (l'alerte grave) **inhibe** `target` (les alertes dérivées)
quand les labels `equal` correspondent. L'alerte inhibée reste visible dans
l'UI (barrée), elle n'est juste pas notifiée.

## 66. Silences : maintenance propre et tracée

Un **silence** coupe les notifications pour un sélecteur de labels, avec début,
fin, auteur et commentaire. À créer dans l'UI (`:9093` → Silences → New Silence)
ou via `amtool` :

```bash
# Silence de 2 h sur toutes les alertes du serveur en maintenance
amtool silence add \
  --alertmanager.url=http://127.0.0.1:9093 \
  --comment="Maintenance planifiée - changement disques" \
  --author="zelef" \
  --duration=2h \
  instance="srv-db01:9100"

# Lister / expirer
amtool silence query --alertmanager.url=http://127.0.0.1:9093
amtool silence expire <id> --alertmanager.url=http://127.0.0.1:9093
```

> **Règle d'équipe** : toute maintenance = un silence **avec commentaire et
> durée limitée**. Un silence "oublié" qui court encore 6 mois après = des
> alertes perdues. Auditez les silences actifs chaque semaine.

## 67. Regroupement : group_by, group_wait, repeat_interval

```yaml
route:
  group_by: ['alertname', 'datacenter', 'severity']
  group_wait: 30s        # à la 1re alerte d'un groupe, attendre 30 s (rafale initiale)
  group_interval: 5m     # si d'autres alertes rejoignent le groupe, notifier toutes les 5 min max
  repeat_interval: 4h    # tant que ça sonne, rappeler toutes les 4 h
```

Bien choisir `group_by` :

- Trop fin (`instance`) → N notifications pour N serveurs (bruit).
- Trop large (rien) → un seul message fourre-tout illisible.
- Bon compromis : `['alertname', 'datacenter']` ou `['alertname', 'equipe']`.

⚠️ `repeat_interval: 4h` sur du critical, `12h` ou `24h` sur du warning.
Un `repeat_interval` trop court = harcèlement ; trop long = on oublie l'alerte.

## 68. Bonnes pratiques d'alerting (l'anti-bruit)

1. **Alerter sur les symptômes, pas sur les causes** : "le site est lent"
   (latence p95) plutôt que "CPU à 82 %". Le CPU à 82 % n'est pas un problème
   en soi.
2. **`for:` toujours** : 5 min minimum pour le critical, 15-30 min pour le warning.
   Zéro `for` = alerte sur un à-coup.
3. **Seuils avec marge** : alerter à 85 % disque (warning) puis 95 % (critical),
   pas à 99 %.
4. **Chaque alerte = une action** : si personne ne sait quoi faire quand elle
   sonne, c'est une mauvaise alerte. Ajouter un **runbook** en annotation.
5. **Severities limitées** : `critical` (astreinte, nuit), `warning` (équipe,
   heures ouvrées). Pas de `info` notifié (dashboard uniquement).
6. **Pas d'alerte sur des compteurs bruts** : toujours `rate()`/`increase()`.
7. **Tester** : `promtool test rules` + déclencher volontairement 1×/trimestre.
8. **Revue trimestrielle** : toute alerte ignorée 3 fois = seuil à revoir ou
   alerte à supprimer. Le bruit tue la vigilance.

## 69. Exemples d'alertes production (prêtes à l'emploi)

`/etc/prometheus/rules/alertes.yml` :

```yaml
groups:
  - name: alertes-infra
    interval: 30s
    rules:
      # --- Disponibilité ---
      - alert: InstanceDown
        expr: up == 0
        for: 5m
        labels: {severity: critical, equipe: infra}
        annotations:
          summary: "Cible DOWN : {{ $labels.instance }} (job {{ $labels.job }})"
          description: "Aucun scrape réussi depuis 5 min. Vérifier réseau/service."
          runbook: "https://wiki.lan/runbooks/instance-down"

      - alert: BlackboxProbeFailed
        expr: probe_success == 0
        for: 5m
        labels: {severity: critical, equipe: infra}
        annotations:
          summary: "Sonde {{ $labels.job }} en échec sur {{ $labels.instance }}"

      # --- Disques ---
      - alert: DisqueBientotPlein
        expr: |
          100 * (1 - node_filesystem_avail_bytes{fstype!~"tmpfs|overlay|squashfs"}
                     / node_filesystem_size_bytes) > 85
        for: 30m
        labels: {severity: warning, equipe: infra}
        annotations:
          summary: "Disque {{ $labels.mountpoint }} à {{ $value | printf \"%.1f\" }} % sur {{ $labels.instance }}"

