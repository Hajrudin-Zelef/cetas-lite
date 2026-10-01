---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes-1
title: "Mise à jour des dépôts et installation des dépendances"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Nvidia"]
dates: []
keywords: ["agentic", "agents", "gpu", "mai", "mcp", "nvidia", "open source"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes.md
source_anchor: ""
source_lines: [1, 59]
sha256: 8b128b540e7a744a45270d7b1ed7baa8dcc37c4841231bcb2feb7087de5ed6f8
---

# Mise à jour des dépôts et installation des dépendances

**Publié le 03 avril 2026** – Docker Compose s’est imposé comme l’outil de référence pour orchestrer plusieurs conteneurs sur un seul hôte. En avril 2026, le format Compose est passé sous l’égide du *Compose Specification*, supporté nativement par Docker Engine et compatible avec des plateformes tierces comme Kubernetes (via Kompose) ou Podman. Ce tutoriel Docker Compose vous guide pas à pas, en 13 étapes concrètes, pour bâtir une **stack de production complète** : application Python (FastAPI), base de données PostgreSQL, cache Redis, reverse proxy Nginx, monitoring Prometheus/Grafana, healthchecks, secrets, profiles et déploiement sécurisé.

À l’issue de ce guide, vous disposerez d’un `docker-compose.yml` production-ready, d’un fichier `compose.override.yml` pour le développement local avec *Compose Watch*, et d’une compréhension fine des pièges à éviter (volumes orphelins, ordre de démarrage, gestion des secrets, mauvaise utilisation de `depends_on`). Comptez environ 60 à 90 minutes pour parcourir l’ensemble si vous tapez chaque commande, et 30 minutes pour lire et copier-coller. Tous les exemples sont testés sur Ubuntu 24.04 LTS et macOS 15 (Apple Silicon), avec Docker Desktop ou Docker Engine standalone.

## Pourquoi Docker Compose Reste Incontournable en 2026

Malgré la domination de Kubernetes dans les déploiements à grande échelle, **Docker Compose** demeure l’outil de prédilection pour le développement local, les environnements de test, l’intégration continue et même certains déploiements de production mono-hôte. Le format Compose a évolué : depuis 2024, il fait partie du *Compose Specification* géré par la communauté open source, garantissant une portabilité bien au-delà de l’écosystème Docker. Docker Docs a mis à jour sa page officielle Docker Compose en mai 2026 pour documenter cette intégration native de Compose v2+ dans la CLI Docker, et le 3 septembre 2026, Docker a publié sur GitHub Releases la version **v5.5.1** – selon la page de cycle de vie de Versio.io (mise à jour en septembre 2026), cette ligne v5, toujours non-LTS, compte déjà 16 versions publiées depuis son lancement (la première remontant à début décembre 2025) – ouvrant une nouvelle ligne stable que ce tutoriel prend pour référence. Le moteur de rendu officiel est désormais réécrit en Go et intégré comme plugin de la CLI Docker (`docker compose`, sans tiret).

Le sondage Stack Overflow Developer Survey 2025 place Docker au premier rang des outils les plus appréciés des développeurs (76,7 % d’usage, 78 % d’avis favorables). Selon les statistiques publiques de Docker Hub, plusieurs centaines de milliards de pulls cumulés ont été enregistrés en 2025, et le dépôt GitHub de Compose dépasse 35 000 étoiles. En production, des sociétés comme PayPal, Spotify et Atlassian utilisent encore Compose pour des microservices internes ou des outils annexes, prouvant qu’il s’agit bien plus qu’un simple outil de poste de travail. L’écosystème s’étend même désormais à l’IA : dans son annonce « agentic apps » du 10 juillet 2025, Docker a étendu Compose avec un support natif des agents IA et une intégration MCP Gateway, déjà disponible sur au moins deux grands clouds publics.

Les principaux apports de Docker Compose en 2026 sont la suppression définitive du champ `version:` (déprécié depuis v2.27, désormais ignoré par le runtime), l’intégration native de *Compose Watch* pour le rechargement à chaud, le support de *Bake* pour des builds parallèles ultra-rapides, la prise en charge des **secrets** via fichiers chiffrés, des **profiles** conditionnels et un support officiel des GPU NVIDIA via `deploy.resources.reservations.devices`. Ce tutoriel couvre l’ensemble de ces nouveautés.

## Prérequis : Versions et Environnement Recommandés

Avant de démarrer, assurez-vous de disposer des outils suivants. Les versions indiquées sont celles validées au 1<sup>er</sup> avril 2026. Toutes les commandes ont été testées sous Ubuntu 24.04, Debian 12, Fedora 40, macOS 15 Sequoia et Windows 11 (via WSL2 Ubuntu).

| Outil | Version Minimale | Version Recommandée 2026 | Vérification | 
|---|---|---|---|
| Docker Engine | 20.10 | 27.x ou supérieur | `docker --version` | 
| Docker Compose plugin | v2.20 | v2.31+ (latest) | `docker compose version` | 
| Système Linux | Kernel 4.x | Kernel 6.x (Ubuntu 24.04) | `uname -r` | 
| RAM disponible | 2 Go | 8 Go pour stack complète | `free -h` | 
| Espace disque | 5 Go | 20 Go (images + volumes) | `df -h /var/lib/docker` | 
| curl / git | n’importe quelle | récente | `curl --version` /`git --version` | 

Sous Linux, installez Docker via le dépôt officiel (le paquet `docker.io` de Debian/Ubuntu est souvent en retard) en installant explicitement le paquet `docker-compose-plugin` pour obtenir Compose v2 – c’est d’ailleurs la démarche suivie par le « Docker Compose Getting Started Guide », mis à jour par Tutorials.technology en mars 2026. Sur macOS et Windows, privilégiez **Docker Desktop** (gratuit pour usage personnel et entreprises de moins de 250 employés ou moins de 10 M$ de revenus annuels selon la licence Docker) ; la version **4.56.0**, livrée par Docker en janvier 2026, embarque désormais Compose v5 avec un nouveau SDK Go, en remplacement des installations historiquement limitées à Compose v2. Une alternative open source viable est **Rancher Desktop** ou **Colima**, qui exposent une CLI Docker compatible.

## Étape 1 : Installer Docker Engine et Docker Compose

Sous Ubuntu 24.04 LTS, l’installation officielle se fait en quelques lignes via le script `get.docker.com` ou, mieux, via les dépôts *apt* signés. La méthode *apt* facilite les mises à jour de sécurité automatiques.

```
# Mise à jour des dépôts et installation des dépendances
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg lsb-release
# Ajout de la clé GPG officielle Docker
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \
  sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
# Ajout du dépôt stable
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
# Installation de Docker Engine + Compose plugin + Buildx
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io \
  docker-buildx-plugin docker-compose-plugin
# Permission sans sudo (relogguez-vous après)
sudo usermod -aG docker $USER
newgrp docker
# Vérification
docker --version
docker compose version
```
La commande `docker compose version` doit retourner au minimum `Docker Compose version v2.39.2` – la version publiée le 4 août 2025 dans les notes de version officielles, celle-là même que cite le tutoriel Dokploy d’août 2025 en exigeant un Docker Engine/CLI en version 28.3.3 pour un déploiement optimal – voire la toute dernière **v5.5.1**, sortie le 3 septembre 2026, si vous avez mis à jour récemment ; entre les deux, la documentation Docker Engine 28.x signale aussi des builds embarquant Compose v2.39.1. Si vous voyez l’ancienne commande `docker-compose` (avec tiret) en Python, il s’agit de la **v1, dépréciée depuis juin 2023**. N’utilisez plus que `docker compose` (sans tiret), implémenté en Go, plus rapide et activement maintenu : une alerte de sécurité SUSE publiée en juin 2025 documentait justement un correctif pour un paquet `docker-compose` v1.x, rappelant que cette branche ne reçoit plus aucun correctif officiel.

## Étape 2 : Comprendre le Compose Specification

