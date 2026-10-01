---
id: collect-261001-general-networking/general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026-3
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026.md
source_anchor: ""
source_lines: [128, 275]
sha256: dc37236e3f5eb8bab04669d215c6a02992a14ee9e80b7aa73e28b7ea22489eb5
---

# Mise à jour complète du système

```
# /opt/vaultwarden/docker-compose.yml
services:
  vaultwarden:
    image: vaultwarden/server:1.36.0
    container_name: vaultwarden
    restart: unless-stopped
    env_file: .env
    volumes:
      - ./vw-data:/data
    networks:
      - vw-net
  caddy:
    image: caddy:2.11.4
    container_name: caddy
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
      - "443:443/udp"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
      - ./caddy-data:/data
      - ./caddy-config:/config
    networks:
      - vw-net
networks:
  vw-net:
```
Le mapping `443:443/udp` active HTTP/3 (QUIC) côté Caddy. L’épinglage des versions d’image (`:1.36.0` plutôt que `:latest`) est une bonne pratique : il rend vos déploiements reproductibles et vous évite une mise à jour majeure non maîtrisée. Le tableau suivant récapitule les variables d’environnement essentielles que nous allons définir dans le fichier `.env`.

| Variable | Valeur recommandée | Rôle | 
|---|---|---|
| `DOMAIN` | `https://coffre.exemple.fr` | URL canonique complète, indispensable au bon fonctionnement | 
| `SIGNUPS_ALLOWED` | `false` (après création) | Verrouille les inscriptions publiques | 
| `SIGNUPS_VERIFY` | `true` | Exige une vérification e-mail (nécessite SMTP) | 
| `INVITATIONS_ALLOWED` | `true` | Autorise l’ajout d’utilisateurs par invitation | 
| `ADMIN_TOKEN` | Hachage Argon2 | Protège le panneau d’administration | 
| `LOG_FILE` | `/data/vaultwarden.log` | Active la journalisation pour fail2ban | 
| `LOG_LEVEL` | `warn` | Niveau de verbosité des journaux | 
| `PUSH_ENABLED` | `false` (par défaut) | Notifications push mobiles (optionnel) | 

## Étape 5 – Générer l’ADMIN_TOKEN Argon2 et le fichier .env

Le panneau d’administration de Vaultwarden permet de gérer les utilisateurs et la configuration. Le protéger par un jeton en clair serait risqué : depuis les versions récentes, Vaultwarden accepte un hachage **Argon2id** au format PHC, bien plus sûr. Générez-le directement avec le binaire embarqué dans l’image :

```
# Générer un hachage Argon2id pour le jeton admin
docker run --rm -it vaultwarden/server:1.36.0 /vaultwarden hash
# Saisissez un mot de passe fort, deux fois.
# Sortie (exemple) :
# Generating an Argon2id PHC string...
# $argon2id$v=19$m=65540,t=3,p=4$bWluZWQ...$cXVlbHF1ZWNob3Nl...
```
Copiez l’intégralité de la chaîne `$argon2id$...`. Créez ensuite le fichier `.env` dans `/opt/vaultwarden`. Point crucial et source d’erreur n°1 : dans un fichier lu par Docker Compose, **chaque `$` du hachage doit être doublé en `$$`**, faute de quoi Compose tentera d’interpréter la chaîne comme des variables et corrompra le jeton.

```
# /opt/vaultwarden/.env
DOMAIN=https://coffre.exemple.fr
# Inscriptions : laissez "true" UNIQUEMENT le temps de créer le 1er compte
SIGNUPS_ALLOWED=true
SIGNUPS_VERIFY=false
INVITATIONS_ALLOWED=true
# Jeton admin Argon2id – NOTEZ les $ doublés en $$
ADMIN_TOKEN=$$argon2id$$v=19$$m=65540,t=3,p=4$$bWluZWQ...$$cXVlbHF1ZWNob3Nl...
# Journalisation (requise pour fail2ban)
LOG_FILE=/data/vaultwarden.log
LOG_LEVEL=warn
# SMTP (optionnel) – décommentez et adaptez pour activer la vérification e-mail
# SMTP_HOST=smtp.exemple.fr
# [email protected]
# SMTP_PORT=587
# SMTP_SECURITY=starttls
# [email protected]
# SMTP_PASSWORD=motdepasse_smtp
```
Restreignez les permissions du fichier : `chmod 600 .env`. Il contient le hachage du jeton admin et, éventuellement, des identifiants SMTP. Si vous n’avez pas de serveur SMTP, laissez `SIGNUPS_VERIFY=false` : la vérification par e-mail sera simplement désactivée, sans bloquer le service.

## Étape 6 – Mettre en place le reverse proxy HTTPS avec Caddy

Caddy est idéal pour ce rôle car il obtient et renouvelle les certificats Let’s Encrypt sans configuration manuelle. Le `Caddyfile` tient en quelques lignes : il déclare le domaine, active la compression et transmet tout le trafic au conteneur Vaultwarden sur son port 80 interne.

```
# /opt/vaultwarden/Caddyfile
coffre.exemple.fr {
    encode gzip zstd
    # Transmet HTTP et WebSocket vers Vaultwarden
    reverse_proxy vaultwarden:80 {
        header_up X-Real-IP {remote_host}
    }
    # En-têtes de sécurité
    header {
        Strict-Transport-Security "max-age=31536000; includeSubDomains; preload"
        X-Content-Type-Options "nosniff"
        X-Frame-Options "DENY"
        Referrer-Policy "no-referrer"
    }
}
```
Un détail capital concernant les **WebSockets** : depuis Vaultwarden 1.29, les notifications temps réel (synchronisation instantanée entre appareils) transitent par le même port HTTP que le reste de l’application. Vous n’avez donc *plus* besoin du port 3012 séparé que réclamaient les anciens guides. La directive `reverse_proxy vaultwarden:80` suffit : Caddy détecte et relaie automatiquement la mise à niveau WebSocket. Ajouter une configuration 3012 obsolète provoquerait des erreurs inutiles.

## Étape 7 – Démarrer la stack et vérifier le déploiement

Tout est en place. Lancez les deux conteneurs en arrière-plan, puis surveillez les journaux pour confirmer que Caddy obtient bien son certificat et que Vaultwarden démarre.

```
cd /opt/vaultwarden
docker compose up -d
# Suivre les journaux des deux services
docker compose logs -f
# Vérifier l'état des conteneurs
docker compose ps
```
Dans les journaux de Caddy, recherchez une ligne indiquant `certificate obtained successfully`. Côté Vaultwarden, vous verrez le démarrage du serveur Rocket. Testez ensuite l’accès depuis l’extérieur en interrogeant le point de santé intégré :

```
# Point de santé : renvoie l'heure UTC du serveur si tout va bien
curl -s https://coffre.exemple.fr/alive
# Sortie attendue (exemple) :
# "2026-06-26T09:14:03.512Z"
# Version exposée par l'API
curl -s https://coffre.exemple.fr/api/version
# Sortie attendue (exemple) :
# "1.36.0"
```
Ouvrez enfin `https://coffre.exemple.fr` dans un navigateur : la page d’inscription Bitwarden doit s’afficher avec un cadenas valide. Si le certificat est marqué comme non sécurisé, c’est presque toujours un problème de DNS (étape 3) ou de port 80 bloqué empêchant le défi ACME.

## Étape 8 – Créer le premier compte et verrouiller les inscriptions

Sur la page d’accueil, cliquez sur « Créer un compte » et enregistrez votre utilisateur principal avec une adresse e-mail et un **mot de passe maître** long et unique. Ce mot de passe est la clé de tout votre coffre : il n’est jamais transmis au serveur et ne peut pas être réinitialisé sans indice de récupération. Notez-le hors ligne pendant la phase d’installation.

Une fois le compte créé, **fermez immédiatement les inscriptions publiques**. Sans cela, n’importe qui découvrant votre URL pourrait créer un compte. Modifiez le `.env` et redémarrez la pile :

```
# Dans /opt/vaultwarden/.env, passez :
# SIGNUPS_ALLOWED=false
sed -i 's/^SIGNUPS_ALLOWED=true/SIGNUPS_ALLOWED=false/' /opt/vaultwarden/.env
# Appliquer le changement
docker compose up -d
# Vérifier l'état du service après redémarrage
docker compose ps
```
À partir de là, l’ajout de nouveaux utilisateurs (membres de la famille, collègues) se fera par **invitation** depuis le panneau d’administration ou une organisation, grâce à `INVITATIONS_ALLOWED=true`. C’est le modèle recommandé : aucune inscription anonyme, chaque accès est nominatif et tracé.

## Étape 9 – Connecter les clients Bitwarden officiels

L’intérêt de Vaultwarden est de pouvoir utiliser les applications Bitwarden officielles. Dans chaque client – extension navigateur, application de bureau, mobile ou CLI – il faut indiquer l’URL de votre serveur auto-hébergé *avant* de se connecter.

