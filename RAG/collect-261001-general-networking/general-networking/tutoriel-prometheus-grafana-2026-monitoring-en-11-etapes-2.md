---
id: collect-261001-general-networking/general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes-2
title: "Docker version 27.5.1, build 9f9e405"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention"]
source: docs/RAG/collect-261001-general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes.md
source_anchor: ""
source_lines: [62, 251]
sha256: 6a9485649225d8e56359629436fdc4326442ed90b21861ca2eb86429a0dbb6c6
---

# Docker version 27.5.1, build 9f9e405

```
# docker-compose.yml
services:
  prometheus:
    image: prom/prometheus:v3.11.0
    container_name: prometheus
    volumes:
      - ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml:ro
      - ./prometheus/alert-rules.yml:/etc/prometheus/alert-rules.yml:ro
      - prometheus_data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--storage.tsdb.retention.time=30d'
      - '--web.enable-lifecycle'
      - '--web.enable-remote-write-receiver'
    ports:
      - "9090:9090"
    networks:
      - monitoring
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--spider", "-q", "http://localhost:9090/-/healthy"]
      interval: 30s
      timeout: 5s
      retries: 3
  grafana:
    image: grafana/grafana:11.6.0
    container_name: grafana
    volumes:
      - grafana_data:/var/lib/grafana
      - ./grafana/provisioning:/etc/grafana/provisioning:ro
      - ./grafana/dashboards:/var/lib/grafana/dashboards:ro
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=SecurePass2026!
      - GF_USERS_ALLOW_SIGN_UP=false
      - GF_SERVER_ROOT_URL=http://localhost:3000
    ports:
      - "3000:3000"
    networks:
      - monitoring
    depends_on:
      prometheus:
        condition: service_healthy
    restart: unless-stopped
  node-exporter:
    image: prom/node-exporter:v1.9.0
    container_name: node-exporter
    volumes:
      - /proc:/host/proc:ro
      - /sys:/host/sys:ro
      - /:/rootfs:ro
    command:
      - '--path.procfs=/host/proc'
      - '--path.sysfs=/host/sys'
      - '--path.rootfs=/rootfs'
      - '--collector.filesystem.mount-points-exclude=^/(sys|proc|dev|host|etc)($$|/)'
    ports:
      - "9100:9100"
    networks:
      - monitoring
    restart: unless-stopped
  alertmanager:
    image: prom/alertmanager:v0.28.1
    container_name: alertmanager
    volumes:
      - ./alertmanager/alertmanager.yml:/etc/alertmanager/alertmanager.yml:ro
    ports:
      - "9093:9093"
    networks:
      - monitoring
    restart: unless-stopped
networks:
  monitoring:
    driver: bridge
volumes:
  prometheus_data:
  grafana_data:
```
Points clés de cette configuration : le flag `--storage.tsdb.retention.time=30d` conserve 30 jours de données (ajustez selon votre espace disque). Le flag `--web.enable-lifecycle` permet de recharger la configuration Prometheus sans redémarrer le conteneur via `curl -X POST http://localhost:9090/-/reload`. Les volumes nommés (`prometheus_data`, `grafana_data`) garantissent la persistance des données entre les redémarrages.

Le `depends_on` avec `condition: service_healthy` assure que Grafana ne démarre qu’une fois Prometheus opérationnel. Sans cette vérification, Grafana tenterait de se connecter à une source de données inexistante au démarrage.

## Étape 3 : Configurer Prometheus et les Cibles de Scraping

Le fichier `prometheus.yml` définit les cibles que Prometheus va scraper (interroger) pour collecter les métriques. Prometheus utilise un modèle pull : il va chercher les métriques sur chaque cible à intervalles réguliers, contrairement aux systèmes push comme StatsD ou Datadog Agent.

```
# prometheus/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s
  scrape_timeout: 10s
rule_files:
  - "alert-rules.yml"
alerting:
  alertmanagers:
    - static_configs:
        - targets:
          - alertmanager:9093
scrape_configs:
  # Prometheus se scrape lui-même
  - job_name: 'prometheus'
    static_configs:
      - targets: ['prometheus:9090']
        labels:
          environment: 'production'
  # Métriques système via Node Exporter
  - job_name: 'node-exporter'
    static_configs:
      - targets: ['node-exporter:9100']
        labels:
          environment: 'production'
          instance_type: 'linux'
  # Exemple : scraper une application Python/Go
  # - job_name: 'my-app'
  #   metrics_path: '/metrics'
  #   static_configs:
  #     - targets: ['app:8080']
```
Le paramètre `scrape_interval: 15s` est le standard recommandé. Un intervalle plus court (5s) augmente la précision mais consomme significativement plus de ressources : chaque série temporelle occupe environ 1-2 octets par échantillon dans le TSDB. Pour 10 000 séries à 5s, cela représente environ 3,5 Go par jour contre 1,2 Go à 15s.

La section `alerting` connecte Prometheus à Alertmanager. Lorsqu’une règle d’alerte est déclenchée, Prometheus envoie la notification à Alertmanager qui se charge du routage, du groupage et de la déduplication. Nous configurerons les règles d’alerte à l’étape 8.

## Étape 4 : Provisionner Grafana avec la Source de Données Prometheus

Le provisioning automatique de Grafana évite la configuration manuelle via l’interface web. En plaçant des fichiers YAML dans le répertoire `provisioning/`, Grafana configure automatiquement les sources de données et les dashboards au démarrage. C’est indispensable pour les déploiements reproductibles — d’autant plus depuis la sortie le 30 juillet 2026 du plugin Grafana Prometheus datasource v13.1.7 par Grafana Labs, qui exige désormais Grafana 12.3.0 ou une version supérieure d’après le changelog GitHub officiel du projet.

```
# grafana/provisioning/datasources/prometheus.yml
apiVersion: 1
datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://prometheus:9090
    isDefault: true
    editable: false
    jsonData:
      timeInterval: '15s'
      queryTimeout: '60s'
      httpMethod: POST
      exemplarTraceIdDestinations:
        - name: traceID
          datasourceUid: tempo
```
L’option `access: proxy` signifie que Grafana fait les requêtes vers Prometheus côté serveur, pas côté navigateur. C’est plus sécurisé car Prometheus n’a pas besoin d’être exposé publiquement. Attention toutefois si votre configuration reposait sur l’authentification SigV4 ou Azure AD intégrée : la documentation Grafana publiée en août 2026 confirme que Grafana 13 a retiré ces deux méthodes d’authentification du datasource Prometheus natif, au profit de plugins d’authentification dédiés. L’option `httpMethod: POST` reste recommandée depuis Prometheus 3.0 pour les requêtes PromQL longues qui dépasseraient la limite d’URL en GET.

Créez ensuite le fichier de provisioning des dashboards :

```
# grafana/provisioning/dashboards/dashboards.yml
apiVersion: 1
providers:
  - name: 'default'
    orgId: 1
    folder: 'Monitoring'
    type: file
    disableDeletion: false
    updateIntervalSeconds: 30
    allowUiUpdates: true
    options:
      path: /var/lib/grafana/dashboards
      foldersFromFilesStructure: false
```
Ce fichier indique à Grafana de charger automatiquement tous les fichiers JSON présents dans `/var/lib/grafana/dashboards`. Le paramètre `updateIntervalSeconds: 30` vérifie les modifications toutes les 30 secondes, ce qui permet de mettre à jour les dashboards sans redémarrer Grafana.

## Étape 5 : Lancer la Stack et Vérifier le Fonctionnement

Avant de lancer la stack Prometheus Grafana, créez un fichier de configuration minimal pour Alertmanager. Sans ce fichier, le conteneur refusera de démarrer.

```
# alertmanager/alertmanager.yml
global:
  resolve_timeout: 5m
route:
  group_by: ['alertname', 'severity']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 1h
  receiver: 'default'
receivers:
  - name: 'default'
    webhook_configs:
      - url: 'http://localhost:5001/'
        send_resolved: true
```
Lancez maintenant l’ensemble de la stack :

