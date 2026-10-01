---
id: collect-261001-general-networking/general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes-1
title: "Docker version 27.5.1, build 9f9e405"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "distribution", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-prometheus-grafana-2026-monitoring-en-11-etapes.md
source_anchor: ""
source_lines: [1, 61]
sha256: 3c2b872ea994304ce2b1ceec2ce03a047865c9c55e52373be27d88b9b15d07ce
---

# Docker version 27.5.1, build 9f9e405

Prometheus Grafana est la stack de monitoring open source la plus déployée au monde : plus de 90 % des clusters Kubernetes certifiés CNCF l’utilisent en 2026. Le succès commercial suit la même courbe : Grafana Labs, l’éditeur derrière Grafana, a vu son chiffre d’affaires récurrent annuel atteindre 400 millions de dollars dès septembre 2025 avec plus de 7 000 clients payants, et sa valorisation a grimpé à 9 milliards de dollars sur la même période, contre 6 milliards seize mois plus tôt. Ce tutoriel vous guide de l’installation à la production, avec Docker Compose, PromQL, alertes et dashboards complets. Vous aurez un système de monitoring fonctionnel en moins de 60 minutes.

Prometheus a depuis fait entrer sa branche LTS 3.13 en statut « supported » en juillet 2026, avec un maintien assuré jusqu’en juillet 2027, tandis que la branche stable a continué d’avancer : la version 3.14.0 est sortie en août 2026, suivie dès septembre 2026 d’une première release candidate 3.15.0-rc.0, confirmant le rythme de publication d’environ une nouvelle version toutes les six semaines. Ces versions consolident les optimisations TSDB, le Remote Write 2.0 et le support natif OpenTelemetry pour les équipes qui préfèrent la stabilité. Côté visualisation, Grafana OSS est passé en version majeure 13 : la 13.0.2 est sortie le 9 juin 2026, suivie de la 13.1.1 le 21 juillet 2026 (deuxième version de maintenance de la branche 13.x), avec des dashboards IA assistés et une gestion unifiée des alertes — un duo Prometheus 3.14.0 / Grafana 13.2.0 ayant même été repris tel quel par Anaconda dans son Data Science & AI Workbench 5.9.0.2 en août 2026, signe de l’adoption de cette nouvelle branche bien au-delà de l’écosystème CNCF. Ensemble, ces deux outils couvrent la collecte de métriques, la visualisation, l’alerting et l’analyse de performance — le tout gratuitement.

Ce guide cible les développeurs, DevOps et SRE francophones qui veulent mettre en place un monitoring professionnel. Chaque étape inclut du code testé, les erreurs courantes et leurs solutions. À la fin, vous disposerez d’un projet Docker Compose complet avec Prometheus, Grafana, Node Exporter et Alertmanager.

## Prérequis : Versions et Configuration Minimale

Avant de commencer l’installation de Prometheus Grafana, vérifiez que votre environnement respecte les prérequis suivants. Ce tutoriel a été testé sur Ubuntu 24.04 LTS et Debian 12, mais fonctionne sur toute distribution Linux avec Docker installé. Évitez les versions intermédiaires déjà en fin de support comme Prometheus 3.2.1 (sortie le 26 février 2025, support de sécurité arrêté le 31 mars 2025), 3.3.1 (sortie le 2 mai 2025, support arrêté dès le 27 mai 2025), ou même l’ancienne branche LTS 3.5 dont le support de sécurité s’est arrêté en juillet 2026 : privilégiez désormais la branche LTS 3.13 (supportée jusqu’en juillet 2027) ou la dernière version stable 3.14.0.

| Composant | Version minimale | Version recommandée | Rôle | 
|---|---|---|---|
| Docker Engine | 24.0 | 27.5+ | Conteneurisation | 
| Docker Compose | 2.20 | 2.34+ | Orchestration multi-conteneurs | 
| Prometheus | 3.0 | 3.11.0 | Collecte et stockage des métriques | 
| Grafana | 11.0 | 11.6.0 | Visualisation et dashboards | 
| Node Exporter | 1.8 | 1.9.0 | Métriques système (CPU, RAM, disque) | 
| Alertmanager | 0.27 | 0.28.1 | Routage et notification des alertes | 
| RAM disponible | 2 Go | 4 Go+ | Prometheus consomme ~1,5 Go pour 100K séries | 
| Espace disque | 10 Go | 50 Go+ | Rétention par défaut : 15 jours | 

Vérifiez vos versions avec ces commandes :

```
docker --version
# Docker version 27.5.1, build 9f9e405
docker compose version
# Docker Compose version v2.34.0
# Vérifier les ressources disponibles
free -h
df -h /var/lib/docker
```
Si Docker n’est pas installé, suivez le tutoriel Docker pour débutants avant de continuer. Pour les utilisateurs macOS ou Windows, Docker Desktop 4.38+ intègre Docker Compose v2 nativement.

## Étape 1 : Créer la Structure du Projet Prometheus Grafana

Une bonne organisation du projet facilite la maintenance et le versionnement. Créez l’arborescence suivante qui sépare clairement la configuration de chaque composant. Cette structure est celle recommandée par la communauté CNCF pour les déploiements Docker Compose de Prometheus Grafana ; pour un déploiement natif Kubernetes, la même communauté maintient le Prometheus Operator, dont la version v0.90 (18 mars 2026) marque déjà au moins la huitième branche mineure publiée depuis janvier 2025, preuve d’un rythme de développement soutenu.

```
# Créer le répertoire du projet
mkdir -p monitoring-stack/{prometheus,grafana/{provisioning/datasources,provisioning/dashboards,dashboards},alertmanager}
cd monitoring-stack
# Vérifier la structure
tree .
# .
# ├── alertmanager/
# ├── grafana/
# │   ├── dashboards/
# │   └── provisioning/
# │       ├── dashboards/
# │       └── datasources/
# └── prometheus/
```
Chaque répertoire a un rôle précis. Le dossier `prometheus/` contient le fichier de configuration principal et les règles d’alerte. Le dossier `grafana/provisioning/` permet de pré-configurer les sources de données et les dashboards au démarrage — c’est ce qu’on appelle le provisioning automatique. Le dossier `alertmanager/` contient la configuration du routage des alertes vers Slack, email ou PagerDuty.

Cette approche « Infrastructure as Code » garantit que votre stack de monitoring est reproductible. Vous pouvez versionner l’ensemble dans Git et déployer la même configuration sur plusieurs environnements (développement, staging, production) sans modification manuelle.

## Étape 2 : Configurer Docker Compose pour la Stack Complète

Le fichier Docker Compose orchestre les quatre conteneurs de votre stack Prometheus Grafana. Cette configuration utilise les versions officielles les plus récentes à ce jour — Prometheus 3.13.1 LTS (10 juillet 2026), désormais dépassée en amont par la 3.14.0 (août 2026) et même une release candidate 3.15.0-rc.0 (septembre 2026) pour les équipes qui suivent le edge, et Grafana 13.1.1 (21 juillet 2026) — et inclut les bonnes pratiques de sécurité : réseau dédié, volumes persistants et healthchecks.

