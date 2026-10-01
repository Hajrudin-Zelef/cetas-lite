---
id: collect-261001-general-networking/general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes-7
title: "Docker version 27.5.1, build 9f9e405"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes.md
source_anchor: ""
source_lines: [842, 929]
sha256: 73233ef421a7a4b3272556d65e6b9efb517e79296c1abd17ea4060a722e0e7b3
---

# Docker version 27.5.1, build 9f9e405

En revanche, pour les équipes qui préfèrent un service managé sans maintenance opérationnelle, des solutions comme Datadog ou Grafana Cloud offrent une expérience clé en main. Le choix dépend du ratio coût total de possession (TCO) vs temps d'ingénierie disponible pour la maintenance.

## Projet Complet : Stack de Monitoring Docker Compose

Voici le récapitulatif de tous les fichiers du projet. Clonez ou recréez cette structure pour disposer d'une stack Prometheus Grafana complète et fonctionnelle en moins de 5 minutes.

```
# Structure finale du projet
monitoring-stack/
├── docker-compose.yml              # Orchestration des 4 services
├── prometheus/
│   ├── prometheus.yml              # Configuration principale + scrape targets
│   ├── alert-rules.yml             # 6 règles d'alerte (CPU, RAM, disque, prédictive)
│   └── recording-rules.yml         # Recording rules pour performances
├── grafana/
│   ├── provisioning/
│   │   ├── datasources/
│   │   │   └── prometheus.yml      # Auto-configuration Prometheus
│   │   └── dashboards/
│   │       └── dashboards.yml      # Auto-import des dashboards
│   └── dashboards/
│       └── node-exporter.json      # Dashboard système complet
├── alertmanager/
│   └── alertmanager.yml            # Routage alertes + notifications email
└── app/                            # (Optionnel) Application Flask exemple
    ├── app.py
    ├── requirements.txt
    └── Dockerfile
# Lancer le projet complet
cd monitoring-stack
docker compose up -d
# Vérifier le statut
docker compose ps
curl -s http://localhost:9090/-/healthy
curl -s http://localhost:3000/api/health
# Accéder aux interfaces
# Prometheus : http://localhost:9090
# Grafana    : http://localhost:3000 (admin / SecurePass2026!)
# Alertmanager : http://localhost:9093
```
Ce projet couvre les bases solides d'un monitoring de production. Pour aller plus loin, vous pouvez ajouter Loki pour la centralisation des logs, Tempo pour le tracing distribué, et ainsi constituer la stack LGTM complète (Loki, Grafana, Tempo, Mimir) qui est la référence open source de l'observabilité en 2026.

### Articles Connexes

Pour approfondir les sujets abordés dans ce tutoriel, consultez nos guides connexes :

- Tutoriel Docker pour Débutants 2026 — Si vous débutez avec Docker et les conteneurs
- Maîtriser Docker Compose : Applications Multi-Conteneurs — Approfondissez l'orchestration Docker Compose
- Grafana vs Datadog 2026 — Comparaison détaillée des plateformes d'observabilité
- Déployer avec Kubernetes et Helm — Passez au niveau supérieur avec Kubernetes
- Tutoriel Flask Python 2026 — Créez l'application Python monitorée dans ce guide
- Configurer Nginx comme Reverse Proxy — Sécurisez l'accès à Prometheus et Grafana

## FAQ : Questions Fréquentes sur Prometheus Grafana

**Quelle est la différence entre Prometheus et Grafana ?**

Prometheus est un système de collecte et de stockage de métriques : il scrape les cibles, stocke les données dans son TSDB et évalue les règles d'alerte. Grafana est un outil de visualisation : il se connecte à Prometheus (et d'autres sources de données) pour afficher les métriques sous forme de graphiques, tableaux et dashboards. Ensemble, ils forment une stack de monitoring complète où Prometheus est le « backend » et Grafana le « frontend ».

**Combien de RAM faut-il pour Prometheus en production ?**

La consommation mémoire de Prometheus dépend directement du nombre de séries temporelles actives. Comptez 2 à 3 Ko par série active. Pour un environnement typique avec 50 000 séries (environ 50 serveurs avec Node Exporter et quelques applications), prévoyez 2 à 4 Go de RAM. Pour 500 000 séries et plus, envisagez 8 à 16 Go et utilisez les recording rules pour réduire la charge des requêtes.

**Prometheus peut-il stocker des données sur plusieurs mois ?**

Oui, mais ce n'est pas recommandé en stockage local. Prometheus 3.11 supporte des rétentions longues via `--storage.tsdb.retention.time`, mais pour des données au-delà de 90 jours, utilisez des solutions de stockage longue durée comme Thanos (stockage objet S3/GCS) ou Grafana Mimir. Ces solutions compressent les données historiques et offrent une rétention illimitée à moindre coût.

**Comment monitorer Kubernetes avec Prometheus Grafana ?**

La méthode standard est le Prometheus Operator (kube-prometheus-stack), un chart Helm qui déploie automatiquement Prometheus, Grafana, Alertmanager et les exporters nécessaires (kube-state-metrics, node-exporter). Il configure aussi le service discovery Kubernetes pour détecter automatiquement les nouveaux pods et services. L'installation se fait en une commande : `helm install monitoring prometheus-community/kube-prometheus-stack`.

**Est-ce que Prometheus supporte OpenTelemetry ?**

Depuis Prometheus 3.0 (fin 2024), le support OpenTelemetry est natif. Prometheus peut ingérer des métriques au format OTLP directement via le Remote Write receiver, sans conversion supplémentaire. Cela permet d'utiliser les SDK OpenTelemetry dans vos applications tout en conservant Prometheus comme backend de stockage. C'est la direction recommandée par la CNCF pour 2026.

**Comment sécuriser Prometheus et Grafana en production ?**

Trois niveaux de sécurité sont recommandés. Niveau 1 : activez le TLS et le basic_auth natifs de Prometheus via le fichier `web.yml`. Niveau 2 : placez Prometheus et Grafana derrière un reverse proxy (Nginx, Traefik) avec authentification OAuth2 ou LDAP. Niveau 3 : isolez la stack monitoring dans un réseau Docker dédié et n'exposez que Grafana publiquement, en configurant Prometheus avec `--web.listen-address=0.0.0.0:9090` uniquement sur le réseau interne.

**Quelle est la différence entre rate() et irate() en PromQL ?**

`rate()` calcule le taux moyen de croissance par seconde sur toute la fenêtre de temps spécifiée (ex : 5 minutes). `irate()` calcule le taux instantané entre les deux derniers points de données uniquement. Utilisez `rate()` pour les alertes et les dashboards de tendance (plus stable), et `irate()` pour les dashboards de détail en temps réel où vous voulez voir les pics instantanés. Ne mélangez jamais `irate()` avec des alertes — cela provoque du flapping.

**Comment migrer de Prometheus 2.x vers 3.x ?**

La migration vers Prometheus 3.x nécessite quelques ajustements. Les données TSDB de Prometheus 2.x sont compatibles avec la version 3.x, mais certains flags CLI dépréciés ont été supprimés. Avant la migration, exécutez `promtool check config` pour identifier les paramètres obsolètes. Les principales incompatibilités concernent la nouvelle UI (l'ancienne UI est supprimée), le changement de format des alertes, et les feature flags désormais activés par défaut. Consultez la documentation officielle Prometheus pour le guide de migration complet.

*Dernière mise à jour : 07 avril 2026. Testé avec Prometheus 3.11.0, Grafana 11.6.0, Node Exporter 1.9.0 et Alertmanager 0.28.1 sur Ubuntu 24.04 LTS.*
