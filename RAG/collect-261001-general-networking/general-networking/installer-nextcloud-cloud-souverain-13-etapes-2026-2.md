---
id: collect-261001-general-networking/general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026-2
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "attention"]
source: docs/RAG/collect-261001-general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026.md
source_anchor: ""
source_lines: [46, 166]
sha256: 3b4ade8c611e90a8c1ecfbdfb3baf95fbc60ab298e3ae57c205ba86c8cd9bcb4
---

# Mise à jour complète du système

```
# Mise à jour complète du système
sudo apt update && sudo apt upgrade -y
# Outils essentiels
sudo apt install -y curl git ufw fail2ban
# Pare-feu : on ouvre uniquement SSH, HTTP et HTTPS
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow OpenSSH
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable
sudo ufw status verbose
```
Le paquet `fail2ban` protège votre accès SSH contre les attaques par force brute en bannissant automatiquement les adresses IP suspectes. C’est une protection essentielle dès qu’un serveur est exposé sur Internet. Vérifiez ensuite que votre fuseau horaire est correct (important pour les tâches planifiées et les journaux) avec `timedatectl set-timezone Europe/Paris`. Créez enfin un utilisateur dédié non-root si vous n’en avez pas déjà un, et désactivez l’authentification SSH par mot de passe au profit des clés.

**Sortie attendue** de `ufw status verbose` : vous devez voir les règles 22, 80 et 443 en « ALLOW IN », et la politique par défaut « deny (incoming) ». Si le port 443 n’apparaît pas, Caddy ne pourra pas obtenir de certificat à l’étape 7.

## Étape 2 – Installer Docker Engine et Docker Compose v2

Nous utilisons le dépôt officiel Docker plutôt que la version d’Ubuntu, généralement plus ancienne. Le script ci-dessous ajoute la clé GPG et le dépôt Docker, puis installe le moteur et le plugin Compose v2 (la commande moderne `docker compose` sans tiret, à ne pas confondre avec l’ancien `docker-compose`).

```
# Installation officielle de Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
# Permettre à votre utilisateur de lancer Docker sans sudo
sudo usermod -aG docker $USER
newgrp docker
# Vérification des versions installées
docker --version
docker compose version
```
**Sortie attendue** : `Docker version 27.x.x` et `Docker Compose version v2.x.x`. Le script officiel `get.docker.com` installe toujours la dernière version stable, le plugin Compose v2 et le démarrage automatique du service. Activez-le explicitement au démarrage avec `sudo systemctl enable --now docker`. Si la commande `docker compose version` renvoie une erreur « command not found », c’est que le plugin n’est pas installé : exécutez `sudo apt install docker-compose-plugin`.

Pour aller plus loin sur l’installation et le durcissement de Docker, notre guide d’installation de Docker sur Ubuntu 24.04 détaille la configuration du démon, la rotation des journaux et la gestion des volumes – des réglages qui valent la peine d’être appliqués sur un serveur de production.

## Étape 3 – Structurer le projet et les variables d’environnement

Une architecture propre facilite les sauvegardes et la réversibilité. Créez un répertoire dédié et un fichier `.env` qui centralisera tous vos secrets – jamais codés en dur dans le fichier Compose. Ce découplage est une bonne pratique de sécurité et simplifie la portabilité de votre installation vers un autre hébergeur français.

```
# Arborescence du projet
mkdir -p ~/nextcloud-souverain
cd ~/nextcloud-souverain
# Génération de mots de passe robustes (à copier dans .env)
openssl rand -base64 32   # pour POSTGRES_PASSWORD
openssl rand -base64 32   # pour REDIS_PASSWORD
openssl rand -base64 32   # pour NEXTCLOUD_ADMIN_PASSWORD
```
Créez ensuite le fichier `.env` avec un éditeur (`nano .env`) et collez-y vos valeurs. Remplacez les mots de passe par ceux générés ci-dessus et adaptez le domaine et l’e-mail.

```
# Fichier .env
# --- Base de données PostgreSQL ---
POSTGRES_DB=nextcloud
POSTGRES_USER=nextcloud
POSTGRES_PASSWORD=COLLEZ_VOTRE_MDP_POSTGRES
# --- Cache Redis ---
REDIS_PASSWORD=COLLEZ_VOTRE_MDP_REDIS
# --- Compte administrateur Nextcloud ---
NEXTCLOUD_ADMIN_USER=admin
NEXTCLOUD_ADMIN_PASSWORD=COLLEZ_VOTRE_MDP_ADMIN
# --- Domaine et messagerie ---
NC_DOMAIN=cloud.votreentreprise.fr
[email protected]
```
Sécurisez immédiatement ce fichier : `chmod 600 .env`. Il contient toutes vos clés et ne doit être lisible que par votre utilisateur. Si vous versionnez votre configuration avec Git, ajoutez impérativement `.env` à votre `.gitignore` pour ne jamais publier vos secrets.

## Étape 4 – Configurer le service PostgreSQL

Nextcloud recommande PostgreSQL pour les déploiements de production : il est plus performant que SQLite et plus robuste que MySQL sur les grosses bases. Nous allons commencer à construire le fichier `docker-compose.yml`. Voici le premier service, la base de données, avec un volume persistant pour ne jamais perdre les données même si le conteneur est recréé.

```
# docker-compose.yml (partie 1/3 : base de données)
services:
  db:
    image: postgres:16-alpine
    restart: unless-stopped
    volumes:
      - db_data:/var/lib/postgresql/data
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${POSTGRES_USER}"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - nc_internal
```
Deux points méritent l’attention. Le `healthcheck` garantit que Nextcloud ne démarrera qu’une fois la base réellement prête à accepter des connexions, évitant les erreurs de connexion au premier lancement. Le réseau `nc_internal` est un réseau Docker privé : la base de données n’est **jamais exposée** sur Internet, elle n’est accessible que par les autres conteneurs. C’est un principe de sécurité fondamental – aucune raison d’ouvrir le port 5432 vers l’extérieur. Nous définirons les volumes et réseaux à la fin du fichier, à l’étape 7.

## Étape 5 – Ajouter le cache Redis pour les performances

Redis joue deux rôles critiques dans Nextcloud : il sert de cache mémoire (memcache) pour accélérer l’application et de backend de verrouillage transactionnel des fichiers (transactional file locking), qui évite la corruption lors d’accès concurrents. Sur une installation multi-utilisateurs, son absence provoque des lenteurs et des avertissements dans le panneau d’administration. Ajoutons-le au fichier Compose.

```
# docker-compose.yml (partie 2/3 : cache Redis)
  redis:
    image: redis:7-alpine
    restart: unless-stopped
    command: redis-server --requirepass ${REDIS_PASSWORD}
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "-a", "${REDIS_PASSWORD}", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - nc_internal
```
L’option `--requirepass` protège Redis par mot de passe, indispensable même sur un réseau interne. Comme PostgreSQL, Redis reste confiné au réseau privé `nc_internal` et n’expose aucun port public. Nous brancherons Nextcloud sur ce cache via des variables d’environnement à l’étape suivante, ce qui activera automatiquement le verrouillage de fichiers transactionnel – l’un des réglages les plus impactants sur la stabilité d’un **auto-hébergement Nextcloud** en équipe.

## Étape 6 – Déployer le conteneur Nextcloud

Voici le cœur du système. Nous utilisons l’image `nextcloud:33-apache`, qui embarque PHP 8.3, le serveur web Apache et tous les modules requis. Les variables d’environnement connectent Nextcloud à PostgreSQL et à Redis, et pré-remplissent le compte administrateur pour une installation automatisée. Le paramètre `OVERWRITEPROTOCOL=https` est crucial : il indique à Nextcloud qu’il est servi derrière un reverse proxy en HTTPS, ce qui évite les redirections cassées et les avertissements de sécurité.

