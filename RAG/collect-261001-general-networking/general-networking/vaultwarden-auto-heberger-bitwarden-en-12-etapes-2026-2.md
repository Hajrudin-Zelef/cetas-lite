---
id: collect-261001-general-networking/general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026-2
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026.md
source_anchor: ""
source_lines: [38, 127]
sha256: a4906a4b624f324d0005bb909270370ac674a5af1388528d18a86793b78794e3
---

# Mise à jour complète du système

Conséquence directe : votre serveur Vaultwarden ne stocke que des données chiffrées et un hachage d’authentification. Même un attaquant disposant d’un accès complet à la base SQLite ne peut pas déchiffrer vos mots de passe sans le mot de passe maître, qui n’y figure nulle part. C’est l’architecture *zero-knowledge* : le fournisseur – ici, vous-même – n’a techniquement pas connaissance des secrets. Cela explique aussi pourquoi un mot de passe maître perdu est définitivement perdu, et pourquoi le choix du KDF (étape 10) compte tant : il détermine le coût d’une attaque par force brute si la base venait à fuir.

Ce modèle ne dispense pas de sécuriser le serveur. Un attaquant capable d’altérer le code servi au navigateur pourrait théoriquement intercepter le mot de passe maître au moment de la saisie – d’où l’importance du HTTPS, des en-têtes de sécurité et des mises à jour. Mais il garantit qu’une simple fuite de la base de données, scénario le plus fréquent, ne compromet pas directement vos identifiants. C’est tout l’intérêt de coupler une cryptographie éprouvée à une infrastructure que vous maîtrisez.

## Architecture du projet et prérequis

L’architecture cible est volontairement minimaliste et robuste. Deux conteneurs Docker cohabitent sur un réseau interne : **Vaultwarden**, qui écoute en HTTP sur le port 80 à l’intérieur du réseau Docker, et **Caddy**, le reverse proxy qui termine le TLS et obtient automatiquement un certificat Let’s Encrypt. Seuls les ports 80 et 443 sont exposés à l’extérieur, et le port 80 ne sert qu’à la redirection HTTPS et au défi ACME. Le coffre n’est jamais accessible en clair.

Pourquoi le HTTPS est-il obligatoire et non optionnel ? Parce que les clients Bitwarden s’appuient sur la *Web Crypto API*, qui n’est disponible que dans un contexte sécurisé (HTTPS ou localhost). Sans certificat valide, l’interface web et le déverrouillage échouent. Tenter d’exposer Vaultwarden en HTTP simple est l’erreur la plus fréquente des tutoriels anciens. Voici les prérequis précis, avec leurs versions :

- **Un VPS** sous**Ubuntu 24.04 LTS** (1 vCPU et 1 Go de RAM suffisent largement), de préférence chez un hébergeur européen.
- **Docker Engine 27+** et le plugin**Docker Compose v2** (commande`docker compose` ).
- **Vaultwarden 1.36.0** (image`vaultwarden/server` ) et**Caddy 2.11.4** (image`caddy` ).
- **Un nom de domaine** que vous contrôlez, avec accès à la zone DNS pour créer un enregistrement`A` (et`AAAA` en IPv6).
- **Les ports 80 et 443 ouverts** en entrée sur le pare-feu et chez l’hébergeur.
- Un accès **SSH** avec un compte non-`root` disposant de`sudo` .
- **fail2ban** et**sqlite3** installés (nous les ajoutons dans les étapes).

Si Docker n’est pas encore en place, suivez d’abord notre guide dédié pour installer Docker sur Ubuntu 24.04 en 13 étapes, puis revenez ici. Tout au long du tutoriel, remplacez `coffre.exemple.fr` par votre domaine réel et adaptez les chemins à votre environnement.

## Étape 1 – Préparer et durcir le VPS Ubuntu 24.04

Connectez-vous en SSH puis mettez le système à jour. Nous créons un utilisateur dédié, activons le pare-feu UFW en n’ouvrant que le strict nécessaire, et installons les outils de base. Ce socle de durcissement est valable pour n’importe quel service exposé, pas seulement Vaultwarden.

```
# Mise à jour complète du système
sudo apt update && sudo apt full-upgrade -y
# Outils nécessaires aux étapes suivantes
sudo apt install -y ufw fail2ban sqlite3 curl ca-certificates gnupg
# Pare-feu : on n'ouvre que SSH, HTTP et HTTPS
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow OpenSSH
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw --force enable
sudo ufw status verbose
```
La sortie de `ufw status verbose` doit confirmer que seules les règles 22 (OpenSSH), 80 et 443 sont autorisées en entrée. fail2ban démarre automatiquement avec une prison `sshd` par défaut, ce qui protège déjà votre accès SSH contre le bourrage d’identifiants. Nous ajouterons une prison spécifique à Vaultwarden à l’étape 12. Pensez également à désactiver l’authentification SSH par mot de passe au profit des clés, et à interdire la connexion `root` directe – une bonne pratique de base avant d’exposer le moindre service.

## Étape 2 – Installer Docker Engine et Docker Compose

Vaultwarden se déploie le plus proprement via Docker. Installez Docker Engine depuis le dépôt officiel afin d’obtenir une version récente (27 ou supérieure) et le plugin Compose v2. N’utilisez pas le paquet `docker.io` d’Ubuntu, souvent en retard de plusieurs versions.

```
# Dépôt officiel Docker
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \
  sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] \
  https://download.docker.com/linux/ubuntu $(. /etc/os-release; echo $VERSION_CODENAME) stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io \
  docker-buildx-plugin docker-compose-plugin
# Autoriser votre utilisateur à piloter Docker sans sudo
sudo usermod -aG docker "$USER"
newgrp docker
# Vérification
docker --version
docker compose version
```
Les deux commandes de vérification doivent renvoyer des versions valides. Si `docker compose version` échoue mais que `docker-compose` (avec un tiret) existe, vous utilisez l’ancienne v1 : préférez systématiquement la v2 intégrée. Pour aller plus loin sur l’orchestration multi-conteneurs, notre guide Traefik pour un reverse proxy HTTPS détaille une alternative à Caddy que nous évoquerons en fin d’article.

## Étape 3 – Configurer le DNS et pointer le domaine

Caddy ne peut obtenir un certificat Let’s Encrypt que si votre domaine pointe déjà vers l’IP publique du VPS. Dans la zone DNS de votre registraire, créez un enregistrement `A` faisant correspondre `coffre.exemple.fr` à l’adresse IPv4 du serveur. Ajoutez un enregistrement `AAAA` si vous disposez d’une IPv6. Attendez la propagation, puis vérifiez depuis votre machine locale :

```
# Récupérer l'IP publique du VPS (depuis le serveur)
curl -s https://api.ipify.org; echo
# Vérifier la résolution DNS (depuis votre poste)
dig +short coffre.exemple.fr
# Sortie attendue : l'IPv4 de votre VPS, par exemple
# 203.0.113.42
```
Tant que `dig` ne renvoie pas l’IP exacte du VPS, n’allez pas plus loin : l’émission du certificat échouerait et Caddy entrerait dans une boucle de tentatives. La propagation DNS prend généralement quelques minutes, parfois jusqu’à une heure selon le TTL configuré. Profitez-en pour vérifier chez votre hébergeur qu’aucun pare-feu externe ne bloque les ports 80 et 443 – c’est une cause d’échec classique sur les VPS dont le panneau d’administration filtre par défaut.

## Étape 4 – Structurer le projet et écrire docker-compose.yml

Créez un répertoire de projet propre. Toutes les données persistantes (base SQLite, pièces jointes, certificats) y résideront via des montages de volumes, ce qui simplifie les sauvegardes.

```
sudo mkdir -p /opt/vaultwarden
sudo chown "$USER":"$USER" /opt/vaultwarden
cd /opt/vaultwarden
mkdir -p vw-data caddy-data caddy-config
```
Rédigez ensuite le fichier `docker-compose.yml`. Vaultwarden n’expose aucun port directement vers l’hôte : seul Caddy publie 80 et 443. La communication entre les deux conteneurs passe par un réseau Docker interne nommé `vw-net`.

