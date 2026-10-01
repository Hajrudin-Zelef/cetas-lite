---
id: collect-261001-general-networking/general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes-4
title: "Docker version 27.5.1, build 9f9e405"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes.md
source_anchor: ""
source_lines: [427, 571]
sha256: 9e8f09605edcb572dc9e7057705a28747a561e1ed4ccdee2269e2c8aaf50ebd3
---

# Docker version 27.5.1, build 9f9e405

Grafana 13.1.1 — livrée le 21 juillet 2026, depuis complétée par la 13.2.0 que l’on retrouve notamment packagée avec Prometheus 3.14.0 dans l’Anaconda Data Science & AI Workbench 5.9.0.2 en août 2026 — propose aussi d’importer des dashboards communautaires depuis grafana.com. Le dashboard ID 1860 (Node Exporter Full) est le plus populaire avec plus de 15 millions de téléchargements. Pour l’importer : Dashboards → Import → entrez l’ID 1860 → sélectionnez Prometheus comme source de données.

## Étape 8 : Configurer les Règles d’Alerte Prometheus

Les alertes transforment votre monitoring passif en système proactif. Prometheus évalue les règles d’alerte à chaque `evaluation_interval` (15s dans notre configuration) et envoie les alertes déclenchées à Alertmanager.

```
# prometheus/alert-rules.yml
groups:
  - name: system_alerts
    rules:
      # Alerte CPU élevé
      - alert: HighCpuUsage
        expr: 100 - (avg by(instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100) > 85
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "CPU élevé sur {{ $labels.instance }}"
          description: "L'utilisation CPU dépasse 85% depuis 5 minutes. Valeur actuelle : {{ $value | printf \"%.1f\" }}%"
      # Alerte mémoire critique
      - alert: HighMemoryUsage
        expr: (1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100 > 90
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "Mémoire critique sur {{ $labels.instance }}"
          description: "Utilisation mémoire supérieure à 90%. Valeur : {{ $value | printf \"%.1f\" }}%"
      # Alerte espace disque
      - alert: DiskSpaceRunningOut
        expr: (1 - (node_filesystem_avail_bytes{mountpoint="/"} / node_filesystem_size_bytes{mountpoint="/"})) * 100 > 80
        for: 10m
        labels:
          severity: warning
        annotations:
          summary: "Espace disque faible sur {{ $labels.instance }}"
          description: "Le disque / est utilisé à {{ $value | printf \"%.1f\" }}%"
      # Prédiction d'espace disque (alerte prédictive)
      - alert: DiskWillFillIn24Hours
        expr: predict_linear(node_filesystem_avail_bytes{mountpoint="/"}[6h], 24*3600) < 0
        for: 30m
        labels:
          severity: critical
        annotations:
          summary: "Disque plein dans moins de 24h sur {{ $labels.instance }}"
          description: "Au rythme actuel, le disque sera plein dans les prochaines 24 heures"
      # Cible Prometheus indisponible
      - alert: TargetDown
        expr: up == 0
        for: 2m
        labels:
          severity: critical
        annotations:
          summary: "Cible {{ $labels.job }} indisponible"
          description: "{{ $labels.instance }} est DOWN depuis plus de 2 minutes"
      # Alertmanager injoignable
      - alert: AlertmanagerDown
        expr: absent(up{job="alertmanager"})
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "Alertmanager est indisponible"
```
La clause `for: 5m` est cruciale : elle définit la durée pendant laquelle la condition doit rester vraie avant que l'alerte ne se déclenche. Sans cette clause, un pic de CPU de 2 secondes déclencherait une alerte — ce qu'on appelle du « flapping ». La règle `DiskWillFillIn24Hours` utilise `predict_linear()`, une fonction PromQL avancée qui effectue une régression linéaire pour prédire quand l'espace disque sera épuisé.

Après avoir créé ce fichier, rechargez la configuration Prometheus sans redémarrage :

```
# Recharger la configuration Prometheus à chaud
curl -X POST http://localhost:9090/-/reload
# Vérifier que les règles sont bien chargées
curl -s http://localhost:9090/api/v1/rules | python3 -m json.tool | head -30
# Sortie attendue :
# {
#   "status": "success",
#   "data": {
#     "groups": [
#       {
#         "name": "system_alerts",
#         "rules": [...]
#       }
#     ]
#   }
# }
```
## Étape 9 : Configurer Alertmanager pour les Notifications

Alertmanager reçoit les alertes de Prometheus et les route vers les canaux de notification appropriés. Sa force réside dans le groupage (regrouper plusieurs alertes similaires en une seule notification), l'inhibition (supprimer des alertes redondantes) et le silence (désactiver temporairement des alertes pendant une maintenance).

Remplacez le fichier `alertmanager/alertmanager.yml` minimal par cette configuration complète :

```
# alertmanager/alertmanager.yml
global:
  resolve_timeout: 5m
  smtp_smarthost: 'smtp.gmail.com:587'
  smtp_from: '[email protected]'
  smtp_auth_username: '[email protected]'
  smtp_auth_password: 'votre-app-password'
  smtp_require_tls: true
route:
  group_by: ['alertname', 'severity']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 4h
  receiver: 'email-team'
  routes:
    # Alertes critiques → notification immédiate
    - match:
        severity: critical
      receiver: 'email-critical'
      group_wait: 10s
      repeat_interval: 1h
    # Alertes warning → notification groupée
    - match:
        severity: warning
      receiver: 'email-team'
      group_wait: 1m
receivers:
  - name: 'email-team'
    email_configs:
      - to: '[email protected]'
        send_resolved: true
  - name: 'email-critical'
    email_configs:
      - to: '[email protected]'
        send_resolved: true
inhibit_rules:
  # Si une alerte critical existe, inhiber les warning du même job
  - source_match:
      severity: 'critical'
    target_match:
      severity: 'warning'
    equal: ['alertname', 'instance']
```
Le paramètre `group_wait: 30s` attend 30 secondes avant d'envoyer la première notification, ce qui permet de regrouper plusieurs alertes qui se déclenchent simultanément. Le `repeat_interval: 4h` évite de spammer votre boîte mail en réenvoyant la même alerte toutes les 4 heures tant que le problème persiste.

La règle d'inhibition est une bonne pratique : si votre serveur est complètement down (alerte `critical`), il est inutile de recevoir aussi l'alerte `warning` de CPU élevé — l'alerte critique couvre déjà le problème.

## Étape 10 : Monitorer une Application avec des Métriques Custom

Node Exporter fournit les métriques système, mais pour monitorer votre application vous devez exposer des métriques métier. Voici un exemple complet avec une application Python Flask qui expose des métriques Prometheus. Ajoutez ce service à votre `docker-compose.yml` :

