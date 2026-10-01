---
id: collect-261001-general-networking/general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes-6
title: "Docker version 27.5.1, build 9f9e405"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "latency", "memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes.md
source_anchor: ""
source_lines: [710, 841]
sha256: fb11fbc31d91be227e3cca2efb8425869451cacd38d88fe59c48aaa12fdf0fad
---

# Docker version 27.5.1, build 9f9e405

**Problème 1 : « Get http://prometheus:9090: dial tcp: lookup prometheus: no such host » dans Grafana.** Cause : Grafana et Prometheus ne sont pas sur le même réseau Docker. Solution : vérifiez que les deux services partagent le même réseau dans `docker-compose.yml` et que le nom du service correspond à l'URL de la datasource.

```
# Diagnostic réseau Docker
docker network ls
docker network inspect monitoring-stack_monitoring
# Vérifier la connectivité depuis le conteneur Grafana
docker exec grafana wget -qO- http://prometheus:9090/-/healthy
# Expected: Prometheus Server is Healthy.
```
**Problème 2 : Cible « DOWN » dans Prometheus Targets.** Causes possibles : le service cible n'expose pas de métriques, le port est incorrect, ou le chemin `metrics_path` est mauvais. Diagnostic :

```
# Vérifier que le endpoint /metrics est accessible
curl -s http://localhost:9100/metrics | head -5
# HELP go_gc_duration_seconds A summary of pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
# Si erreur : vérifier que le conteneur est bien en cours d'exécution
docker compose ps node-exporter
docker compose logs node-exporter --tail=20
```
**Problème 3 : « error loading config file: parsing YAML file prometheus.yml » au démarrage de Prometheus.** Les fichiers YAML sont sensibles à l'indentation. Utilisez des espaces (jamais des tabulations) et validez votre configuration avant de relancer :

```
# Valider la configuration Prometheus
docker run --rm -v $(pwd)/prometheus:/etc/prometheus prom/prometheus:v3.11.0 promtool check config /etc/prometheus/prometheus.yml
# Valider les règles d'alerte
docker run --rm -v $(pwd)/prometheus:/etc/prometheus prom/prometheus:v3.11.0 promtool check rules /etc/prometheus/alert-rules.yml
```
**Problème 4 : Grafana affiche « No data » malgré des cibles UP.** Vérifiez que la source de données Prometheus est correctement configurée dans Grafana (Settings → Data Sources). Le champ URL doit être `http://prometheus:9090` (nom du service Docker), pas `http://localhost:9090` qui pointe vers le conteneur Grafana lui-même.

**Problème 5 : Prometheus consomme trop de RAM.** Identifiez les métriques à haute cardinalité avec ces requêtes :

```
# Top 10 des métriques par nombre de séries
curl -s http://localhost:9090/api/v1/status/tsdb | python3 -c "
import json, sys
data = json.load(sys.stdin)['data']
print('Séries actives:', data['headStats']['numSeries'])
print('\nTop métriques par cardinalité:')
for item in data['seriesCountByMetricName'][:10]:
    print(f'  {item[\"name\"]}: {item[\"value\"]} séries')
"
```
**Problème 6 : Les alertes ne sont jamais envoyées par email.** Vérifiez la configuration SMTP d'Alertmanager. Gmail nécessite un « App Password » (pas votre mot de passe habituel) et le port 587 avec TLS. Testez avec : `docker exec alertmanager amtool check-config /etc/alertmanager/alertmanager.yml`.

**Problème 7 : « permission denied » sur les volumes montés.** Les conteneurs Prometheus et Grafana s'exécutent avec des UID spécifiques. Prometheus utilise UID 65534 (nobody) et Grafana UID 472. Corrigez les permissions :

```
# Corriger les permissions pour Prometheus
sudo chown -R 65534:65534 ./prometheus_data/
# Corriger les permissions pour Grafana
sudo chown -R 472:472 ./grafana_data/
# Alternative : utiliser user dans docker-compose.yml
services:
  prometheus:
    user: "65534:65534"
```
**Problème 8 : Les dashboards provisionnés ne se mettent pas à jour.** Si vous modifiez un fichier JSON de dashboard et qu'il ne se met pas à jour dans Grafana, vérifiez que `allowUiUpdates: true` est défini dans le fichier de provisioning des dashboards. Notez que les modifications faites via l'interface Grafana seront écrasées au prochain redémarrage si le provisioning est actif.

## Conseils Avancés pour les Experts Prometheus Grafana

Ces techniques avancées sont utilisées par les équipes SRE des plus grandes entreprises tech en 2026. Elles permettent de passer d'un monitoring basique à une observabilité de niveau production.

### Recording Rules pour Accélérer les Dashboards

Les recording rules pré-calculent des requêtes PromQL coûteuses et stockent le résultat comme une nouvelle métrique. Au lieu de recalculer `rate(http_requests_total[5m])` à chaque chargement de dashboard, la recording rule le calcule une seule fois toutes les 15 secondes.

```
# prometheus/recording-rules.yml
groups:
  - name: node_recording_rules
    interval: 15s
    rules:
      - record: instance:node_cpu_utilization:ratio
        expr: 1 - avg by(instance) (rate(node_cpu_seconds_total{mode="idle"}[5m]))
      - record: instance:node_memory_utilization:ratio
        expr: 1 - (node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes)
      - record: job:app_request_rate:5m
        expr: sum by(job) (rate(app_requests_total[5m]))
      - record: job:app_request_latency_p95:5m
        expr: histogram_quantile(0.95, sum by(job, le) (rate(app_request_duration_seconds_bucket[5m])))
```
Utilisez ensuite `instance:node_cpu_utilization:ratio` dans vos dashboards Grafana au lieu de la requête complète. Le temps de rendu passe de 2-5 secondes à quelques millisecondes pour les environnements avec des milliers de séries. La convention de nommage `level:metric:operations` est le standard CNCF recommandé.

### Fédération et Multi-Cluster Monitoring

Pour les architectures multi-environnements, la fédération Prometheus permet à un Prometheus central de scraper des métriques agrégées depuis des Prometheus locaux. Ajoutez cette configuration au Prometheus central :

```
# Sur le Prometheus central
scrape_configs:
  - job_name: 'federate-production'
    honor_labels: true
    metrics_path: '/federate'
    params:
      'match[]':
        - '{__name__=~"instance:.*"}'
        - '{__name__=~"job:.*"}'
    static_configs:
      - targets: ['prometheus-prod.internal:9090']
        labels:
          cluster: 'production'
  - job_name: 'federate-staging'
    honor_labels: true
    metrics_path: '/federate'
    params:
      'match[]':
        - '{__name__=~"instance:.*"}'
    static_configs:
      - targets: ['prometheus-staging.internal:9090']
        labels:
          cluster: 'staging'
```
La fédération fonctionne mieux avec les recording rules : ne fédérez que les métriques pré-agrégées (préfixées `instance:` ou `job:`) pour limiter le trafic réseau et la charge sur le Prometheus central.

## Comparaison : Prometheus Grafana vs Alternatives en 2026

Avant de vous engager dans l'écosystème Prometheus Grafana, il est utile de comprendre comment cette stack se positionne face aux alternatives en 2026. Chaque solution a ses forces selon le contexte d'utilisation.

| Critère | Prometheus + Grafana | Datadog | Victoria Metrics | Elastic Observability | 
|---|---|---|---|---|
| Coût (100 hôtes) | Gratuit (self-hosted) | ~1 800 €/mois | Gratuit (self-hosted) | Gratuit (self-hosted) | 
| Stockage longue durée | Via Thanos/Cortex | Natif (15 mois) | Natif (illimité) | Natif (ILM) | 
| Courbe d'apprentissage | Moyenne (PromQL) | Faible (UI guidée) | Faible (compatible PromQL) | Élevée (KQL + ECS) | 
| Scalabilité | Verticale (+ fédération) | Illimitée (SaaS) | Horizontale native | Horizontale native | 
| Écosystème exporters | 900+ exporters | 750+ intégrations | Compatible Prometheus | Beats + agents | 
| Support OpenTelemetry | Natif (v3.0+) | Natif | Natif | Natif | 
| Alerting | Alertmanager | Intégré + ML | Compatible Alertmanager | Watcher + ML | 

Prometheus Grafana reste le choix dominant pour les équipes qui privilégient le contrôle total, la portabilité et l'absence de vendor lock-in. Selon la CNCF, Prometheus est utilisé par 92 % des organisations cloud-native en 2026. Sa compatibilité avec plus de 900 exporters couvre pratiquement toutes les technologies existantes.

