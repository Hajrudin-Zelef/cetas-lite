---
id: collect-261001-general-networking/general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026-3
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026.md
source_anchor: ""
source_lines: [167, 305]
sha256: a4a38d2fe8d4c73982c40921af29ca9d6c5fea8ccfdb9a43ab903d77ad257943
---

# Mise à jour complète du système

```
# docker-compose.yml (partie 3/3 : application Nextcloud)
  app:
    image: nextcloud:33-apache
    restart: unless-stopped
    depends_on:
      db:
        condition: service_healthy
      redis:
        condition: service_healthy
    volumes:
      - nc_data:/var/www/html
    environment:
      POSTGRES_HOST: db
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      REDIS_HOST: redis
      REDIS_HOST_PASSWORD: ${REDIS_PASSWORD}
      NEXTCLOUD_ADMIN_USER: ${NEXTCLOUD_ADMIN_USER}
      NEXTCLOUD_ADMIN_PASSWORD: ${NEXTCLOUD_ADMIN_PASSWORD}
      NEXTCLOUD_TRUSTED_DOMAINS: ${NC_DOMAIN}
      OVERWRITEPROTOCOL: https
      OVERWRITECLIURL: https://${NC_DOMAIN}
    networks:
      - nc_internal
      - nc_proxy
```
Notez que le conteneur `app` est connecté à **deux réseaux** : `nc_internal` pour dialoguer avec la base et le cache, et `nc_proxy` pour recevoir le trafic du reverse proxy. Nextcloud n’expose ainsi **aucun port directement** sur l’hôte – tout passe par Caddy. C’est l’architecture la plus sûre : un seul point d’entrée chiffré. Le volume `nc_data` contient à la fois le code, la configuration (`config/config.php`) et les données utilisateurs.

## Étape 7 – Reverse proxy Caddy et HTTPS automatique

Caddy est le composant qui rend ce tutoriel particulièrement adapté à la souveraineté : il obtient et renouvelle **automatiquement** les certificats TLS via Let’s Encrypt, sans aucune configuration manuelle de Certbot. Ajoutons le service Caddy et finalisons le fichier Compose avec la déclaration des volumes et des réseaux.

```
# Fin du docker-compose.yml : reverse proxy + volumes + réseaux
  caddy:
    image: caddy:2-alpine
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile:ro
      - caddy_data:/data
      - caddy_config:/config
    depends_on:
      - app
    networks:
      - nc_proxy
volumes:
  db_data:
  redis_data:
  nc_data:
  caddy_data:
  caddy_config:
networks:
  nc_internal:
  nc_proxy:
```
Créez maintenant le fichier `Caddyfile` dans le même répertoire. Sa concision est remarquable : quelques lignes suffisent pour un reverse proxy HTTPS complet, avec les redirections recommandées par Nextcloud pour la découverte des services (CalDAV/CardDAV).

```
# Caddyfile
cloud.votreentreprise.fr {
    reverse_proxy app:80
    # Redirections recommandées par Nextcloud
    redir /.well-known/carddav /remote.php/dav 301
    redir /.well-known/caldav /remote.php/dav 301
    # En-tête de sécurité HSTS
    header Strict-Transport-Security "max-age=15768000;"
    encode gzip
}
```
Remplacez le domaine par le vôtre (il doit correspondre à `NC_DOMAIN`). Caddy détecte automatiquement le nom de domaine, contacte Let’s Encrypt et installe le certificat dès le premier démarrage. Aucune intervention n’est requise, et le renouvellement est géré en arrière-plan tous les 60 jours.

## Étape 8 – Lancer la pile et terminer l’installation web

Tout est en place. Lancez l’ensemble des conteneurs en arrière-plan et surveillez les journaux pour vérifier le bon déroulement, en particulier l’obtention du certificat par Caddy. Notez que cette étape correspond, côté documentation officielle, au flux du **Web Installer** (`setup-nextcloud.php`) dont Nextcloud GmbH a mis à jour la documentation le 17 juillet 2026 – utile si vous préférez suivre les instructions officielles en parallèle de ce tutoriel Docker.

```
# Démarrage de toute la pile
docker compose up -d
# Vérifier que les 4 conteneurs tournent
docker compose ps
# Suivre les journaux (Ctrl+C pour quitter)
docker compose logs -f caddy app
```
**Sortie attendue** de `docker compose ps` : quatre services (`db`, `redis`, `app`, `caddy`) avec le statut « running » ou « healthy ». Dans les journaux Caddy, recherchez la ligne `certificate obtained successfully`. Ouvrez ensuite votre navigateur sur `https://cloud.votreentreprise.fr` : vous devez voir l’assistant de connexion Nextcloud avec un cadenas valide. Comme nous avons pré-rempli les identifiants admin via le fichier `.env`, l’installation de la base de données se fait automatiquement au premier accès – connectez-vous simplement avec le compte `admin` et son mot de passe.

Si la page affiche une erreur « Accès interdit » ou « domaine non fiable », c’est que `NEXTCLOUD_TRUSTED_DOMAINS` n’a pas été pris en compte – nous corrigeons cela à l’étape 9. Si vous obtenez une erreur de certificat, vérifiez que votre DNS pointe bien vers le serveur et que le port 80 est ouvert.

## Étape 9 – Régler les domaines de confiance et le cron avec occ

`occ` est l’outil en ligne de commande de Nextcloud, l’équivalent d’un couteau suisse pour l’administration. On l’exécute à l’intérieur du conteneur `app`, en tant qu’utilisateur du serveur web `www-data`. Voici les commandes essentielles de post-installation.

```
# Vérifier l'état de l'instance
docker compose exec -u www-data app php occ status
# Ajouter / confirmer un domaine de confiance
docker compose exec -u www-data app php occ config:system:set \
  trusted_domains 1 --value=cloud.votreentreprise.fr
# Définir la région du téléphone par défaut (France)
docker compose exec -u www-data app php occ config:system:set \
  default_phone_region --value=FR
# Activer le mode cron pour les tâches d'arrière-plan
docker compose exec -u www-data app php occ background:cron
```
Le passage en mode `cron` est important : par défaut, Nextcloud exécute ses tâches d’arrière-plan (nettoyage, notifications, indexation) à chaque chargement de page (AJAX), ce qui est peu fiable. Pour un vrai système cron, ajoutez une tâche planifiée sur l’hôte qui appelle le script toutes les 5 minutes :

```
# Éditer la crontab de l'utilisateur
crontab -e
# Ajouter cette ligne (exécution toutes les 5 minutes)
*/5 * * * * docker compose -f /home/UTILISATEUR/nextcloud-souverain/docker-compose.yml exec -T -u www-data app php cron.php
```
Rendez-vous ensuite dans **Paramètres d’administration → Vue d’ensemble** dans l’interface web. Cette page liste les avertissements de configuration. À ce stade, le cache mémoire (Redis), le verrouillage transactionnel et la région téléphonique doivent être au vert. Toute alerte restante y est documentée avec un lien vers la solution.

## Étape 10 – Activer l’édition collaborative avec Collabora Online

Pour remplacer véritablement Google Docs ou Microsoft Office 365, il faut l’édition collaborative de documents bureautiques directement dans le navigateur. Collabora Online Development Edition (CODE) fournit cette capacité, en s’appuyant sur le moteur LibreOffice. Ajoutez ce service à votre `docker-compose.yml`.

```
# Service Collabora à ajouter dans docker-compose.yml
  collabora:
    image: collabora/code
    restart: unless-stopped
    environment:
      aliasgroup1: https://cloud.votreentreprise.fr:443
      extra_params: --o:ssl.enable=false --o:ssl.termination=true
    cap_add:
      - MKNOD
    networks:
      - nc_proxy
```
Ajoutez le routage dans le `Caddyfile` pour exposer Collabora sur un sous-domaine (créez d’abord l’enregistrement DNS `office.votreentreprise.fr`), puis relancez la pile avec `docker compose up -d`. Dans Nextcloud, installez l’application « Nextcloud Office » depuis l’App Store intégré, allez dans **Paramètres d’administration → Nextcloud Office**, choisissez « Utiliser votre propre serveur » et indiquez l’URL `https://office.votreentreprise.fr`. Vous pourrez désormais créer et co-éditer des documents .odt, .docx, .xlsx et .pptx à plusieurs, en temps réel, sans qu’aucune donnée ne quitte votre infrastructure. C’est la pierre angulaire d’une bureautique réellement souveraine.

## Étape 11 – Durcir la sécurité : 2FA, SSO et en-têtes

