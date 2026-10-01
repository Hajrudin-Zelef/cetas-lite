---
id: collect-261001-general-networking/general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes-3
title: "Docker version 27.5.1, build 9f9e405"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["memory", "valuation"]
source: docs/RAG/collect-261001-general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes.md
source_anchor: ""
source_lines: [252, 426]
sha256: 1a30892d606ba7b247827673d567b953be3e3e10b36f4fd8d042b0c444a55f8c
---

# Docker version 27.5.1, build 9f9e405

```
# Démarrer tous les services en arrière-plan
docker compose up -d
# Vérifier que tous les conteneurs sont en cours d'exécution
docker compose ps
# Sortie attendue :
# NAME             IMAGE                        STATUS                    PORTS
# alertmanager     prom/alertmanager:v0.28.1     Up 2 minutes             0.0.0.0:9093->9093/tcp
# grafana          grafana/grafana:11.6.0        Up 2 minutes (healthy)   0.0.0.0:3000->3000/tcp
# node-exporter    prom/node-exporter:v1.9.0     Up 2 minutes             0.0.0.0:9100->9100/tcp
# prometheus       prom/prometheus:v3.11.0       Up 2 minutes (healthy)   0.0.0.0:9090->9090/tcp
# Vérifier les logs pour détecter d'éventuelles erreurs
docker compose logs --tail=20 prometheus
docker compose logs --tail=20 grafana
```
Vérifiez le bon fonctionnement de chaque composant :

| Service | URL | Vérification | Identifiants par défaut | 
|---|---|---|---|
| Prometheus | http://localhost:9090 | Status → Targets : toutes les cibles « UP » | Aucun (ajoutez basic_auth en production) | 
| Grafana | http://localhost:3000 | Login → Explore → Prometheus disponible | admin / SecurePass2026! | 
| Node Exporter | http://localhost:9100/metrics | Texte brut avec métriques système | Aucun | 
| Alertmanager | http://localhost:9093 | Interface web avec status OK | Aucun | 

Accédez à `http://localhost:9090/targets` dans votre navigateur. Vous devriez voir deux cibles avec le statut « UP » : `prometheus` et `node-exporter`. Si une cible affiche « DOWN », vérifiez les logs du conteneur correspondant avec `docker compose logs [service]`.

## Étape 6 : Maîtriser PromQL — Le Langage de Requête de Prometheus

PromQL (Prometheus Query Language) est le cœur de l’exploitation de Prometheus Grafana. Ce langage fonctionnel permet d’interroger les séries temporelles stockées par Prometheus. Maîtriser PromQL est essentiel pour créer des dashboards pertinents et des alertes précises.

Ouvrez l’interface Prometheus à `http://localhost:9090/graph` pour tester ces requêtes en temps réel.

### Requêtes PromQL Essentielles pour le Monitoring Système

```
# Utilisation CPU en pourcentage (moyenne sur 5 minutes)
100 - (avg by(instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
# Utilisation mémoire en pourcentage
(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100
# Espace disque utilisé en pourcentage
(1 - (node_filesystem_avail_bytes{mountpoint="/"} / node_filesystem_size_bytes{mountpoint="/"})) * 100
# Trafic réseau entrant (octets par seconde)
rate(node_network_receive_bytes_total{device!="lo"}[5m])
# Trafic réseau sortant (octets par seconde)
rate(node_network_transmit_bytes_total{device!="lo"}[5m])
# Nombre de processus en cours d'exécution
node_procs_running
# Load average sur 1 minute
node_load1
# Uptime du système en heures
(time() - node_boot_time_seconds) / 3600
# IOPS disque (lectures par seconde)
rate(node_disk_reads_completed_total[5m])
# Latence moyenne d'écriture disque
rate(node_disk_write_time_seconds_total[5m]) / rate(node_disk_writes_completed_total[5m])
```
La fonction `rate()` est la plus utilisée en PromQL. Elle calcule le taux de croissance par seconde d’un compteur sur une fenêtre de temps. Sans `rate()`, un compteur comme `node_cpu_seconds_total` affiche une valeur toujours croissante qui n’est pas directement exploitable. La fenêtre `[5m]` signifie que le calcul s’effectue sur les 5 dernières minutes, ce qui lisse les variations ponctuelles.

Prometheus 3.11 a optimisé les opérateurs binaires PromQL, réduisant le temps d’évaluation de 25 à 40 % sur les requêtes complexes impliquant des jointures entre séries temporelles. Pour les environnements avec plus de 100 000 séries actives, cette amélioration se traduit par des dashboards Grafana nettement plus réactifs.

## Étape 7 : Créer un Dashboard Grafana pour le Monitoring Système

Plutôt que de créer un dashboard manuellement via l’interface Grafana, nous allons utiliser le provisioning JSON. Cette méthode est reproductible, versionnable dans Git et partageable entre équipes. Créez le fichier suivant dans `grafana/dashboards/` :

```
{
  "dashboard": {
    "title": "Monitoring Système - Node Exporter",
    "uid": "node-exporter-system",
    "timezone": "Europe/Paris",
    "refresh": "30s",
    "time": { "from": "now-1h", "to": "now" },
    "panels": [
      {
        "title": "Utilisation CPU (%)",
        "type": "timeseries",
        "gridPos": { "h": 8, "w": 12, "x": 0, "y": 0 },
        "targets": [{
          "expr": "100 - (avg by(instance) (rate(node_cpu_seconds_total{mode=\"idle\"}[5m])) * 100)",
          "legendFormat": "{{instance}}"
        }],
        "fieldConfig": {
          "defaults": {
            "unit": "percent",
            "thresholds": {
              "steps": [
                { "color": "green", "value": null },
                { "color": "yellow", "value": 70 },
                { "color": "red", "value": 90 }
              ]
            }
          }
        }
      },
      {
        "title": "Utilisation Mémoire (%)",
        "type": "gauge",
        "gridPos": { "h": 8, "w": 6, "x": 12, "y": 0 },
        "targets": [{
          "expr": "(1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)) * 100",
          "legendFormat": "RAM"
        }],
        "fieldConfig": {
          "defaults": {
            "unit": "percent",
            "min": 0,
            "max": 100,
            "thresholds": {
              "steps": [
                { "color": "green", "value": null },
                { "color": "yellow", "value": 75 },
                { "color": "red", "value": 90 }
              ]
            }
          }
        }
      },
      {
        "title": "Espace Disque (%)",
        "type": "gauge",
        "gridPos": { "h": 8, "w": 6, "x": 18, "y": 0 },
        "targets": [{
          "expr": "(1 - (node_filesystem_avail_bytes{mountpoint=\"/\"} / node_filesystem_size_bytes{mountpoint=\"/\"})) * 100",
          "legendFormat": "Disque /"
        }],
        "fieldConfig": {
          "defaults": {
            "unit": "percent",
            "min": 0,
            "max": 100,
            "thresholds": {
              "steps": [
                { "color": "green", "value": null },
                { "color": "yellow", "value": 70 },
                { "color": "red", "value": 85 }
              ]
            }
          }
        }
      },
      {
        "title": "Trafic Réseau (octets/s)",
        "type": "timeseries",
        "gridPos": { "h": 8, "w": 12, "x": 0, "y": 8 },
        "targets": [
          {
            "expr": "rate(node_network_receive_bytes_total{device!=\"lo\"}[5m])",
            "legendFormat": "Entrant - {{device}}"
          },
          {
            "expr": "rate(node_network_transmit_bytes_total{device!=\"lo\"}[5m])",
            "legendFormat": "Sortant - {{device}}"
          }
        ],
        "fieldConfig": { "defaults": { "unit": "Bps" } }
      },
      {
        "title": "Load Average",
        "type": "stat",
        "gridPos": { "h": 8, "w": 12, "x": 12, "y": 8 },
        "targets": [
          { "expr": "node_load1", "legendFormat": "1 min" },
          { "expr": "node_load5", "legendFormat": "5 min" },
          { "expr": "node_load15", "legendFormat": "15 min" }
        ]
      }
    ]
  },
  "overwrite": true
}
```
Enregistrez ce fichier sous `grafana/dashboards/node-exporter.json`, puis redémarrez Grafana : `docker compose restart grafana`. Le dashboard apparaît automatiquement dans le dossier « Monitoring » de Grafana.

Les seuils de couleur (thresholds) sont configurés selon les bonnes pratiques SRE de Google : vert sous 70 % d’utilisation CPU, jaune entre 70 et 90 %, rouge au-dessus de 90 %. Pour la mémoire, les seuils sont légèrement plus permissifs car Linux utilise activement le cache mémoire, ce qui gonfle artificiellement l’utilisation affichée.

