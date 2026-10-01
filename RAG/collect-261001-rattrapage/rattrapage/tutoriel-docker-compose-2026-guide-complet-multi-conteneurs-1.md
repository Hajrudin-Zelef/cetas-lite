---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs-1
title: "Vérifier les versions installées"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["diffusion", "mai", "packaging"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs.md
source_anchor: ""
source_lines: [1, 75]
sha256: b58bc48a999d5515dfdedff587eb9db5602714401550b292989721a0d40ff9d9
---

# Vérifier les versions installées

Docker Compose est devenu l’outil incontournable pour orchestrer des applications multi-conteneurs. Depuis la sortie de Docker Compose v5.0.0 “Mont Blanc” fin 2025 jusqu’à la v5.5.0, publiée le 17 août 2026 avec la prise en charge native des conteneurs d’initialisation (init containers), l’outil n’a cessé de gagner en maturité : maîtriser sa dernière génération est plus que jamais essentiel pour tout développeur ou DevOps. Ce tutoriel complet vous guide pas à pas, de l’installation à la mise en production, en construisant un projet fonctionnel avec une API Python, une base de données PostgreSQL, un cache Redis et un proxy Nginx.

Que vous soyez un développeur qui lance `docker run` manuellement pour chaque conteneur ou un ingénieur DevOps cherchant à standardiser vos environnements, ce guide vous permettra de comprendre Docker Compose en profondeur et de l’utiliser efficacement dans vos projets professionnels. Chaque étape inclut des exemples de code fonctionnels, des sorties attendues et des pièges courants à éviter.

## Prérequis : Versions et Outils Nécessaires

Avant de commencer ce tutoriel Docker Compose, assurez-vous de disposer des outils suivants installés sur votre machine. Chaque version a été testée et validée pour garantir la compatibilité avec les exemples de ce guide. Docker Desktop, disponible sur Windows, macOS et Linux, embarque nativement le plugin Docker Compose – la version 4.79.0, sortie le 22 juin 2026, embarquait déjà Docker Compose v5.4.0 et son nouveau workflow de réconciliation, ce qui simplifie considérablement l’installation et la maintenance des environnements.

| Outil | Version minimale | Version recommandée (mars 2026) | Remarque | 
|---|---|---|---|
| Docker Engine | 24.0 | 29.1 | Inclut les dernières optimisations de performance | 
| Docker Compose | 2.25.0 | v2.40+ / v5.0.0 | Plugin intégré à Docker CLI | 
| Docker Desktop | 4.30 | 4.40+ | Optionnel mais recommandé sur Windows/macOS | 
| Python | 3.10 | 3.12 | Pour notre API FastAPI | 
| Git | 2.30 | 2.44+ | Pour cloner le projet exemple | 
| VS Code / Éditeur | – | Dernière version | Extension Docker recommandée | 

Pour vérifier que Docker et Docker Compose sont correctement installés, exécutez les commandes suivantes dans votre terminal. Si Docker Compose n’est pas disponible en tant que plugin, vous devrez peut-être le mettre à jour. Depuis la migration vers Compose V2, la commande utilise un espace au lieu d’un tiret : `docker compose` remplace l’ancien `docker-compose`. Cette distinction est fondamentale et source de nombreuses erreurs chez les débutants.

```
# Vérifier les versions installées
docker --version
# Sortie attendue : Docker version 29.1.x, build xxxxxxx
docker compose version
# Sortie attendue : Docker Compose version v2.40.x
# Si vous avez Docker Desktop, tout est déjà inclus
# Sinon, installer le plugin Compose manuellement :
sudo apt-get update
sudo apt-get install docker-compose-plugin
```
Un point critique souvent négligé : assurez-vous que votre utilisateur fait partie du groupe `docker` pour éviter de devoir utiliser `sudo` à chaque commande. Sur Linux, exécutez `sudo usermod -aG docker $USER` puis déconnectez-vous et reconnectez-vous. Sur Docker Desktop (Windows/macOS), cela est géré automatiquement.

## Comprendre l’Architecture Docker Compose en 2026

Docker Compose repose sur un fichier YAML déclaratif – `compose.yaml` (le nom recommandé depuis 2024, remplaçant `docker-compose.yml`) – qui décrit l’ensemble de votre stack applicative. Chaque service correspond à un conteneur, et Compose gère automatiquement le réseau, les volumes et les dépendances entre services. En 2026, avec la sortie de Docker Compose v5.0.0 “Mont Blanc” fin 2025, l’écosystème a connu une refonte majeure incluant un support SDK officiel, la délégation des builds à Docker Bake et le support des ressources OCI et Git distantes. Cette lignée v5 s’est enrichie rapidement : la v5.1.4 du 20 mai 2026 a précédé la v5.2.0 du 23 juin 2026, puis la v5.4.0, publiée le 3 août 2026, a introduit un nouveau workflow de réconciliation pour les volumes et les réseaux, réduisant les dérives de configuration entre l’état déclaré et l’état réel de la stack.

L’architecture de Docker Compose se compose de quatre éléments fondamentaux. Les **services** définissent les conteneurs à exécuter, avec leur image, leurs ports, leurs variables d’environnement et leurs dépendances. Les **réseaux** permettent aux conteneurs de communiquer entre eux de manière isolée, créant des segments réseau virtuels qui protègent vos applications. Les **volumes** assurent la persistance des données en dehors du cycle de vie des conteneurs, ce qui est essentiel pour les bases de données. Enfin, les **configurations** et **secrets** permettent de gérer les paramètres sensibles de manière sécurisée.

La branche v2 a vécu son dernier chapitre avec la v2.40.3, publiée le 30 octobre 2025 – la dernière release de la lignée v2 avant le grand saut vers v5 – et intégrée à Debian experimental (paquet 2.40.3-1) le 11 mars 2026 par la Debian Go Packaging Team, preuve de sa diffusion jusque dans les distributions Linux. La v5.0.0 a ensuite migré vers les modules Moby et mis à jour containerd vers la version 2.2.2, renforçant la stabilité globale de l’outil. Ces évolutions font de Docker Compose un outil de plus en plus mature, utilisé par des millions de développeurs à travers le monde pour standardiser leurs environnements de développement et de production.

## Étape 1 : Initialiser la Structure du Projet

Commençons par créer la structure complète de notre projet. Nous allons construire une application web composée de quatre services : une API REST Python avec FastAPI, une base de données PostgreSQL 16, un cache Redis 7 et un reverse proxy Nginx. Cette architecture reflète une stack de production réaliste que vous retrouverez dans de nombreuses entreprises en 2026.

```
# Créer la structure du projet
mkdir -p compose-tutorial/{api,nginx,db}
cd compose-tutorial
# Créer les fichiers nécessaires
touch compose.yaml
touch api/main.py api/requirements.txt api/Dockerfile
touch nginx/nginx.conf
touch db/init.sql
touch .env
# Structure finale :
# compose-tutorial/
# ├── compose.yaml
# ├── .env
# ├── api/
# │   ├── Dockerfile
# │   ├── main.py
# │   └── requirements.txt
# ├── nginx/
# │   └── nginx.conf
# └── db/
#     └── init.sql
```
Cette organisation respecte les bonnes pratiques Docker en séparant chaque service dans son propre répertoire. Le fichier `compose.yaml` à la racine orchestre l’ensemble, tandis que chaque service dispose de son contexte de build isolé. Le fichier `.env` centralise les variables d’environnement sensibles, évitant ainsi de les coder en dur dans le fichier Compose. Notez que nous utilisons le nom `compose.yaml` – le format officiellement recommandé en 2026 – plutôt que l’ancien `docker-compose.yml`.

**Piège courant n°1 :** Ne placez jamais votre fichier `.env` dans un dépôt Git public. Ajoutez-le immédiatement à votre `.gitignore`. Une erreur fréquente chez les développeurs débutants est de versionner des mots de passe de base de données ou des clés API. En 2026, les outils de scanning de secrets comme GitGuardian détectent automatiquement ces fuites, mais la prévention reste la meilleure approche.

## Étape 2 : Créer l’API FastAPI et son Dockerfile

Notre API Python utilise FastAPI, le framework web le plus populaire en France en 2026, pour créer des endpoints REST performants. FastAPI offre une validation automatique des données, une documentation OpenAPI intégrée et des performances asynchrones natives, ce qui en fait le choix idéal pour notre démonstration Docker Compose.

