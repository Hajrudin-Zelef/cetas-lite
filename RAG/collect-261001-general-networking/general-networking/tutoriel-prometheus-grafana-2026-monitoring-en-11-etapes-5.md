---
id: collect-261001-general-networking/general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes-5
title: "Docker version 27.5.1, build 9f9e405"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "latency", "memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes.md
source_anchor: ""
source_lines: [572, 709]
sha256: 05a69372454f008ca26853fd002d8048fe67a9c3678e09ea88abc768f927b67d
---

# Docker version 27.5.1, build 9f9e405

```
# app/requirements.txt
flask==3.1.0
prometheus-client==0.22.0
# app/app.py
from flask import Flask, request
from prometheus_client import Counter, Histogram, Gauge, generate_latest
import time
import random
app = Flask(__name__)
# Métriques custom
REQUEST_COUNT = Counter(
    'app_requests_total',
    'Nombre total de requêtes HTTP',
    ['method', 'endpoint', 'status']
)
REQUEST_LATENCY = Histogram(
    'app_request_duration_seconds',
    'Latence des requêtes en secondes',
    ['method', 'endpoint'],
    buckets=[0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0]
)
ACTIVE_REQUESTS = Gauge(
    'app_active_requests',
    'Nombre de requêtes en cours de traitement'
)
@app.before_request
def before_request():
    request.start_time = time.time()
    ACTIVE_REQUESTS.inc()
@app.after_request
def after_request(response):
    latency = time.time() - request.start_time
    REQUEST_COUNT.labels(
        method=request.method,
        endpoint=request.path,
        status=response.status_code
    ).inc()
    REQUEST_LATENCY.labels(
        method=request.method,
        endpoint=request.path
    ).observe(latency)
    ACTIVE_REQUESTS.dec()
    return response
@app.route('/')
def index():
    time.sleep(random.uniform(0.01, 0.1))
    return {'status': 'ok', 'timestamp': time.time()}
@app.route('/api/data')
def api_data():
    time.sleep(random.uniform(0.05, 0.5))
    return {'data': [1, 2, 3], 'count': 3}
@app.route('/metrics')
def metrics():
    return generate_latest(), 200, {'Content-Type': 'text/plain'}
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
```
Les trois types de métriques Prometheus utilisés ici couvrent 90 % des besoins de monitoring applicatif. Le **Counter** (`app_requests_total`) ne peut qu'augmenter et compte les événements cumulés. L'**Histogram** (`app_request_duration_seconds`) mesure la distribution des latences dans des buckets prédéfinis — indispensable pour calculer les percentiles (P50, P95, P99). Le **Gauge** (`app_active_requests`) peut monter et descendre, idéal pour les valeurs instantanées comme le nombre de connexions actives.

Ajoutez la cible dans `prometheus.yml` puis rechargez :

```
# Ajouter dans prometheus/prometheus.yml → scrape_configs
  - job_name: 'flask-app'
    static_configs:
      - targets: ['app:8080']
    metrics_path: '/metrics'
    scrape_interval: 10s
# Requêtes PromQL pour votre application
# Taux de requêtes par seconde
rate(app_requests_total[5m])
# Latence P95
histogram_quantile(0.95, rate(app_request_duration_seconds_bucket[5m]))
# Taux d'erreurs (5xx)
sum(rate(app_requests_total{status=~"5.."}[5m])) / sum(rate(app_requests_total[5m])) * 100
```
## Étape 11 : Optimiser les Performances de Prometheus en Production

En production, une stack Prometheus Grafana mal configurée peut devenir un goulet d'étranglement. Voici les optimisations essentielles validées par les équipes SRE en 2026, basées sur les améliorations du TSDB de Prometheus 3.11.

| Paramètre | Valeur par défaut | Recommandation production | Impact | 
|---|---|---|---|
| storage.tsdb.retention.time | 15d | 30d à 90d | Rétention des données historiques | 
| storage.tsdb.retention.size | Illimité | 50GB à 200GB | Limite l'espace disque utilisé | 
| storage.tsdb.min-block-duration | 2h | 2h (ne pas modifier) | Taille minimale des blocs TSDB | 
| storage.tsdb.max-block-duration | 36h | 36h (par défaut optimal) | Compaction automatique des blocs | 
| query.max-concurrency | 20 | 10 à 30 | Requêtes PromQL simultanées | 
| query.timeout | 2m | 30s à 60s | Timeout des requêtes longues | 
| scrape_interval | 1m | 15s à 30s | Fréquence de collecte | 

La consommation mémoire de Prometheus suit une règle approximative : comptez 2 à 3 Ko par série temporelle active. Pour un environnement Kubernetes avec 50 pods et 200 métriques par pod, cela représente 10 000 séries, soit environ 20-30 Mo de RAM. Un cluster de 500 pods avec des métriques détaillées peut facilement atteindre 500 000 séries et nécessiter 1 à 1,5 Go de RAM dédiée à Prometheus.

Pour les environnements très larges, envisagez des solutions de stockage longue durée comme Thanos ou Cortex qui s'intègrent nativement avec Prometheus via le Remote Write 2.0 (désormais stable dans Prometheus 3.11). Ces solutions permettent une rétention illimitée avec un coût de stockage objet (S3, GCS) bien inférieur au stockage local SSD.

```
# Configuration production optimisée dans docker-compose.yml
  prometheus:
    image: prom/prometheus:v3.11.0
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--storage.tsdb.retention.time=90d'
      - '--storage.tsdb.retention.size=100GB'
      - '--query.max-concurrency=15'
      - '--query.timeout=45s'
      - '--web.enable-lifecycle'
      - '--web.enable-remote-write-receiver'
      - '--enable-feature=native-histograms'
    deploy:
      resources:
        limits:
          memory: 4G
          cpus: '2.0'
        reservations:
          memory: 2G
          cpus: '1.0'
```
Le flag `--enable-feature=native-histograms` active les histogrammes natifs de Prometheus 3.x, qui réduisent de 10x le nombre de séries temporelles générées par un histogramme classique. Au lieu de créer une série par bucket (10-20 séries), un seul histogramme natif stocke toute la distribution dans une série unique. C'est l'une des optimisations les plus significatives de Prometheus 3.

## 5 Erreurs Courantes avec Prometheus Grafana et Comment les Éviter

Après avoir installé des centaines de stacks Prometheus Grafana, voici les erreurs les plus fréquentes qui affectent les équipes DevOps en 2026. Chacune peut coûter des heures de débogage si elle n'est pas identifiée rapidement.

**Erreur 1 : Utiliser irate() au lieu de rate() pour les alertes.** La fonction `irate()` calcule le taux instantané entre les deux derniers points de données, ce qui la rend extrêmement volatile. Dans une règle d'alerte, cela provoque du flapping constant. Utilisez toujours `rate()` avec une fenêtre d'au moins `[5m]` pour les alertes. Réservez `irate()` uniquement aux dashboards où vous voulez voir les variations instantanées.

**Erreur 2 : Ne pas configurer la rétention par taille.** Sans `--storage.tsdb.retention.size`, Prometheus peut remplir votre disque en quelques jours si le nombre de séries temporelles explose (par exemple, un label avec une cardinalité élevée). Configurez toujours une limite de taille en plus de la limite temporelle.

**Erreur 3 : Labels à haute cardinalité.** Ajouter un label `user_id` ou `request_id` à une métrique crée une nouvelle série temporelle pour chaque valeur unique. Avec 100 000 utilisateurs, cela génère 100 000 séries par métrique, consommant 200 à 300 Mo de RAM supplémentaire. Utilisez les logs (Loki) pour les données à haute cardinalité, pas les métriques Prometheus.

**Erreur 4 : Oublier send_resolved dans Alertmanager.** Sans `send_resolved: true`, vous recevez l'alerte quand le problème commence mais jamais quand il est résolu. L'équipe d'astreinte doit alors vérifier manuellement si le problème est toujours actif, ce qui est une source majeure de fatigue d'alerte.

**Erreur 5 : Exposer Prometheus sans authentification.** Par défaut, Prometheus n'a aucune authentification. En production, utilisez un reverse proxy Nginx ou Traefik avec basic_auth ou OAuth2. Prometheus 3.x supporte nativement le TLS et le basic_auth via le fichier `web.yml`, sans proxy supplémentaire.

## Dépannage : 8 Problèmes Fréquents et Leurs Solutions

Cette section couvre les problèmes les plus fréquemment rencontrés lors du déploiement d'une stack Prometheus Grafana, avec des commandes de diagnostic précises et des solutions testées.

