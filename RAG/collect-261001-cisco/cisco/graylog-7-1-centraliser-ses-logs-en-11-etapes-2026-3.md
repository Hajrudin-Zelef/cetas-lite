---
id: collect-261001-cisco/cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026-3
title: "Génère un secret de session Graylog (obligatoire, 16+ caractères)"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026.md
source_anchor: ""
source_lines: [118, 251]
sha256: 766a05235f47f6df2822b92feb6e15173cc932af305287574b4e303adece7287
---

# Génère un secret de session Graylog (obligatoire, 16+ caractères)

```
services:
  mongodb:
    image: mongo:8.0
    container_name: graylog-mongodb
    restart: unless-stopped
    volumes:
      - mongodb_data:/data/db
    networks:
      - graylog-net
  opensearch:
    image: opensearchproject/opensearch:2.19.5
    container_name: graylog-opensearch
    restart: unless-stopped
    environment:
      - "OPENSEARCH_JAVA_OPTS=-Xms2g -Xmx2g"
      - "discovery.type=single-node"
      - "plugins.security.disabled=true"
      - "bootstrap.memory_lock=true"
    ulimits:
      memlock:
        soft: -1
        hard: -1
    volumes:
      - opensearch_data:/usr/share/opensearch/data
    networks:
      - graylog-net
  graylog:
    image: graylog/graylog:7.1.8
    container_name: graylog-server
    restart: unless-stopped
    depends_on:
      - mongodb
      - opensearch
    environment:
      - GRAYLOG_PASSWORD_SECRET=${GRAYLOG_PASSWORD_SECRET}
      - GRAYLOG_ROOT_PASSWORD_SHA2=${GRAYLOG_ROOT_PASSWORD_SHA2}
      - GRAYLOG_HTTP_EXTERNAL_URI=http://localhost:9000/
      - GRAYLOG_MONGODB_URI=mongodb://mongodb:27017/graylog
      - GRAYLOG_ELASTICSEARCH_HOSTS=http://opensearch:9200
    ports:
      - "9000:9000"
      - "12201:12201/udp"
      - "12201:12201/tcp"
      - "1514:1514/udp"
    volumes:
      - graylog_data:/usr/share/graylog/data
    networks:
      - graylog-net
networks:
  graylog-net:
    driver: bridge
volumes:
  mongodb_data:
  opensearch_data:
  graylog_data:
```
Notez l’usage de l’image `graylog/graylog:7.1.8` et non un tag `latest` : figer la version évite qu’une mise à jour majeure surprenne la production un lundi matin.

## Étape 3 : Générer les secrets et démarrer la pile Graylog

Une fois le fichier en place, le démarrage se fait en une seule commande. Le premier lancement prend plusieurs minutes le temps qu’OpenSearch initialise ses index internes.

```
docker compose up -d
docker compose ps
docker compose logs -f graylog
```
Résultat attendu dans `docker compose ps` une fois la pile stabilisée :

```
NAME                 STATUS
graylog-mongodb      Up 2 minutes
graylog-opensearch    Up 2 minutes (healthy)
graylog-server        Up 1 minute
```
Si le conteneur `graylog-server` redémarre en boucle, c’est presque toujours parce qu’il a démarré avant qu’OpenSearch soit prêt à répondre. Patientez deux minutes supplémentaires avant de conclure à un problème de configuration.

## Étape 4 : Première connexion et configuration initiale de l’interface

Ouvrez un navigateur sur `http://IP_DU_SERVEUR:9000`. L’identifiant par défaut est `admin`, et le mot de passe est celui saisi en clair à l’étape 2 avant son hachage SHA-256 (dans notre exemple, `MonMotDePasse!2026`). Changez-le immédiatement depuis le menu “System > Users” si ce serveur est exposé au-delà d’un réseau de test.

La première tâche dans l’interface consiste à vérifier que l’index par défaut fonctionne, via “System > Indices”. Vous devez voir l’index set “Default index set” en statut vert, preuve que la communication avec OpenSearch est opérationnelle. C’est aussi le bon moment pour activer l’authentification à deux facteurs sur le compte admin, une pratique désormais recommandée par l’ANSSI pour tout accès à un outil de sécurité critique. Ceux qui utilisent déjà un dispositif de conformité NIS2 retrouveront ici une des exigences de traçabilité des accès imposées par la directive.

## Étape 5 : Configurer les inputs GELF et Syslog

Graylog ne reçoit aucun log tant qu’aucun “input” n’est déclaré. Un input est un point d’écoute réseau qui accepte un format de message précis. Rendez-vous dans “System > Inputs”, sélectionnez “GELF UDP” dans la liste déroulante, puis cliquez sur “Launch new input”. Renseignez le port 12201 et laissez l’adresse d’écoute sur 0.0.0.0 pour accepter les connexions de tout le réseau interne.

Répétez l’opération avec un input “Syslog UDP” sur le port 1514, utile pour les équipements réseau (pare-feu, switches) qui ne parlent que le protocole Syslog historique et ne savent pas produire du GELF. Une fois les deux inputs actifs, l’écran “System > Inputs” doit afficher un statut “Running” en vert pour chacun, avec un compteur de messages qui reste à zéro tant qu’aucune source n’envoie de données.

## Étape 6 : Envoyer les logs Linux et applicatifs vers Graylog

La méthode la plus simple pour un serveur Linux existant consiste à reconfigurer rsyslog pour qu’il duplique ses journaux vers Graylog en plus du stockage local. Éditez `/etc/rsyslog.d/60-graylog.conf` sur chaque machine à surveiller.

```
# /etc/rsyslog.d/60-graylog.conf
# Envoie tous les logs systeme vers Graylog en Syslog UDP
*.* @IP_DU_SERVEUR_GRAYLOG:1514
# Puis redemarrer le service
# sudo systemctl restart rsyslog
```
Pour les conteneurs Docker, la méthode recommandée passe par le driver de logging GELF natif, qui évite d’installer un agent supplémentaire dans chaque conteneur.

```
docker run -d \
  --log-driver=gelf \
  --log-opt gelf-address=udp://IP_DU_SERVEUR_GRAYLOG:12201 \
  --name mon-application \
  mon-image:latest
```
Pour Nginx, ajoutez simplement une règle rsyslog dédiée qui capte les fichiers `access.log` et `error.log` via un module imfile, ou redirigez directement le format de log Nginx vers syslog dans `nginx.conf` avec la directive `error_log syslog:server=IP_DU_SERVEUR_GRAYLOG:1514 warn;`. Une fois ces sources connectées, retournez sur “Search” dans l’interface Graylog : les messages doivent apparaître en quelques secondes, classés par ordre chronologique inverse.

## Étape 7 : Créer des pipelines et extracteurs pour structurer les logs

Un log brut sous forme de texte libre est difficile à interroger. Les pipelines Graylog appliquent des règles de traitement qui extraient des champs structurés (adresse IP source, code de statut HTTP, nom d’utilisateur) directement exploitables dans les recherches et les dashboards. Rendez-vous dans “System > Pipelines” pour créer une règle qui repère les tentatives de connexion SSH échouées.

```
rule "extraire echec SSH"
when
  contains(to_string($message.message), "Failed password")
then
  let extracted_ip = regex("from ([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3})", to_string($message.message));
  set_field("source_ip", extracted_ip["0"]);
  set_field("event_type", "ssh_auth_failure");
end
```
Une fois cette règle attachée au pipeline par défaut et connectée au flux “All messages”, chaque tentative de connexion SSH échouée génère automatiquement un champ `source_ip` et un champ `event_type` filtrables. C’est la base d’une détection de brute-force efficace, le même principe que celui utilisé par CrowdSec pour bloquer les attaques, mais ici appliqué à des fins d’investigation a posteriori plutôt que de blocage en temps réel.

## Étape 8 : Construire des dashboards et alertes de sécurité

Un dashboard Graylog agrège plusieurs widgets (graphiques, compteurs, cartes) sur un seul écran. Depuis “Dashboards > Create dashboard”, ajoutez un widget de type “Aggregation” filtré sur `event_type:ssh_auth_failure`, groupé par `source_ip`, pour visualiser en un coup d’œil quelles adresses martèlent vos serveurs.

Les alertes se configurent dans “Alerts > Event Definitions”. Créez une définition qui déclenche une notification dès que plus de 10 échecs SSH proviennent de la même IP sur une fenêtre de 5 minutes. Graylog 7.1 permet de router cette alerte vers un webhook, un e-mail SMTP, ou directement vers un canal Slack ou Microsoft Teams via une intégration HTTP.

