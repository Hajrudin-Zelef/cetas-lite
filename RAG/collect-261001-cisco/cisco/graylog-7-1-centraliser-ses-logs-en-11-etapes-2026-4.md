---
id: collect-261001-cisco/cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026-4
title: "Génère un secret de session Graylog (obligatoire, 16+ caractères)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "open source"]
source: docs/RAG/collect-261001-cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026.md
source_anchor: ""
source_lines: [252, 321]
sha256: 7d61c781269e239835103a1c51efc26ce15a6fc9969974cd8c882f154f8010ba
---

# Génère un secret de session Graylog (obligatoire, 16+ caractères)

```
curl -X POST http://localhost:9000/api/events/definitions \
  -u admin:MonMotDePasse!2026 \
  -H "Content-Type: application/json" \
  -H "X-Requested-By: cli" \
  -d '{
    "title": "Brute-force SSH detecte",
    "priority": 3,
    "config": {
      "type": "aggregation-v1",
      "query": "event_type:ssh_auth_failure",
      "group_by": ["source_ip"],
      "series": [{"function": "count"}],
      "conditions": {"expression": {"expr": ">", "left": {"expr": "number-ref", "ref": "count()"}, "right": {"expr": "number", "value": 10}}},
      "search_within_ms": 300000,
      "execute_every_ms": 60000
    }
  }'
```
C’est précisément ce type de règle qui exploite les capacités de détection comportementale natives ajoutées dans la branche 7.1, sans avoir besoin d’écrire de corrélation manuelle complexe comme c’était le cas dans les versions antérieures à 6.0.

## Étape 9 : Rétention des index, data tiering et volumétrie

Sans politique de rétention, un index Graylog grossit indéfiniment jusqu’à saturer le disque. Dans “System > Indices > Default index set”, configurez une rotation basée sur la taille (par exemple 5 Go par index) ou sur le temps (rotation quotidienne), puis définissez le nombre maximum d’index conservés avant suppression automatique. Une configuration courante pour une PME conserve 30 jours de logs “chauds” consultables instantanément, puis archive au-delà.

Le data tiering introduit en version 6.0 et toujours actif en 7.1 permet de basculer automatiquement les index les plus anciens vers un stockage “warm” moins coûteux (disque plus lent ou stockage objet), tout en les gardant interrogeables sans réimport manuel. Sur un volume de 50 Go de logs par jour, cette fonction peut diviser par dix le coût de stockage à un an par rapport à tout conserver sur du SSD rapide.

## Étape 10 : Durcir l’instance avec TLS, RBAC et pare-feu

Une instance Graylog qui centralise des logs de sécurité devient elle-même une cible de choix. Trois mesures de durcissement s’imposent avant toute mise en production.

D’abord, activez TLS sur l’interface web en générant un certificat (Let’s Encrypt ou une autorité interne) et en le déclarant via la variable `GRAYLOG_HTTP_EXTERNAL_URI=https://logs.monentreprise.fr/` couplée à un reverse proxy Nginx qui termine le TLS. Ensuite, créez des rôles distincts via “System > Roles” : un analyste SOC n’a pas besoin des droits d’administration système, seulement d’un accès en lecture aux dashboards et à la recherche. Enfin, restreignez au pare-feu l’accès aux ports 12201 et 1514 aux seules IP internes autorisées à envoyer des logs.

```
# Exemple avec UFW sur le serveur Graylog
sudo ufw allow from 10.0.0.0/24 to any port 12201 proto udp
sudo ufw allow from 10.0.0.0/24 to any port 1514 proto udp
sudo ufw allow from 10.0.0.0/24 to any port 9000 proto tcp
sudo ufw deny 12201
sudo ufw deny 1514
sudo ufw enable
```
## Étape 11 : Sauvegarder, mettre à jour et superviser Graylog

Trois éléments doivent être sauvegardés régulièrement : la base MongoDB (configuration, utilisateurs, dashboards), les snapshots OpenSearch (les données de logs elles-mêmes), et le fichier `docker-compose.yml` avec son `.env`. Un script cron simple suffit pour la partie MongoDB.

```
#!/bin/bash
# backup-graylog-mongo.sh
DATE=$(date +%Y%m%d)
docker exec graylog-mongodb mongodump --archive=/tmp/backup-$DATE.gz --gzip
docker cp graylog-mongodb:/tmp/backup-$DATE.gz /backups/graylog/mongo-$DATE.gz
find /backups/graylog/ -name "mongo-*.gz" -mtime +30 -delete
```
Pour la mise à jour vers une future version 7.1.x, comme la 7.1.9 apparue sur la page de téléchargement officielle le 2 septembre 2026, ne modifiez jamais le tag d’image en production sans avoir d’abord testé la migration sur un environnement de préproduction, et sauvegardez systématiquement avant de lancer `docker compose pull && docker compose up -d`. Évitez en revanche de basculer directement sur la bêta 7.2.0-beta.1, mise en ligne le 31 août 2026, tant qu’elle n’a pas atteint une disponibilité générale stable. Consultez toujours la calendrier de fin de vie des versions Graylog avant une montée de version majeure, car les exigences MongoDB et OpenSearch changent d’une branche à l’autre.

## Conservation des logs et conformité RGPD

Centraliser des logs de sécurité soulève une question juridique que beaucoup d’administrateurs négligent : ces journaux contiennent souvent des adresses IP, des identifiants de connexion, parfois des noms d’utilisateur complets, autant de données à caractère personnel au sens du RGPD. La CNIL recommande une durée de conservation proportionnée à la finalité du traitement, généralement limitée à six mois pour des logs de connexion à finalité de sécurité, sauf obligation légale contraire ou enquête en cours.

Concrètement, cela signifie que la politique de rétention configurée à l’étape 9 n’est pas qu’un paramètre technique de gestion d’espace disque : c’est aussi un engagement de conformité. Documentez la durée choisie, la justification métier associée, et assurez-vous que les index archivés en tier “warm” restent purgés au terme de la période retenue plutôt que conservés indéfiniment par simple oubli. Pour les entités soumises à la directive NIS2, cette traçabilité documentée fait partie des éléments vérifiés lors d’un contrôle par l’ANSSI, qui recommande depuis 2026 l’usage d’outils open source justement pour garder une maîtrise complète de la chaîne de traitement des données, sans dépendre d’un hébergeur tiers hors UE.

## Projet complet : une architecture de supervision prête pour une PME

En assemblant les onze étapes précédentes, voici l’architecture complète obtenue à la fin de ce tutoriel : un serveur Graylog 7.1.8 conteneurisé qui reçoit les logs Syslog des pare-feu et switches, les logs GELF des applications Docker, et les journaux systeme de chaque serveur Linux via rsyslog. Un pipeline extrait automatiquement les tentatives de connexion suspectes, un dashboard centralise la vue d’ensemble, et une alerte prévient l’équipe dès qu’un seuil de brute-force est franchi. La rétention des logs suit une politique à 30 jours “chauds” avec data tiering au-delà, et l’accès est protégé par TLS, des rôles RBAC et un pare-feu restrictif.

Cette base peut ensuite recevoir les résultats d’un scan OpenVAS/Greenbone pour croiser vulnérabilités connues et tentatives d’exploitation réelles, ou les résultats d’un scan réseau Nmap pour documenter l’inventaire des services exposés. Pour une entreprise soumise à la directive NIS2, cette architecture répond directement aux exigences de journalisation et de détection d’incident prévues par le texte.

## Erreurs fréquentes à éviter

