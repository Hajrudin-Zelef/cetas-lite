---
id: collect-261001-general-networking/general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026-3
title: "Identifier la clé USB (attention à bien cibler le bon disque)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attribution"]
source: docs/RAG/collect-261001-general-networking/truenas-monter-un-nas-diy-en-13-etapes-2026.md
source_anchor: ""
source_lines: [95, 166]
sha256: b60e52381c8d5a5dfdcc1735c89397587ee084d7c8d73ec1d11aeb166a484c6d
---

# Identifier la clé USB (attention à bien cibler le bon disque)

Au premier démarrage, la console texte de TrueNAS affiche l’adresse IP attribuée par DHCP sur votre réseau local. Notez-la, puis ouvrez un navigateur depuis un autre PC connecté au même réseau et rendez-vous sur cette adresse en HTTPS. Le certificat sera auto-signé au départ, votre navigateur affichera donc un avertissement de sécurité : c’est normal, acceptez l’exception temporairement (on réglera un vrai certificat à l’étape sécurité).

Connectez-vous avec le compte root et le mot de passe défini pendant l’installation. Dans le menu Network, configurez une adresse IP statique plutôt que le DHCP par défaut : un NAS qui change d’adresse IP à chaque redémarrage du routeur casse tous vos partages réseau et vos scripts de sauvegarde. Réservez l’adresse via votre box/routeur (réservation DHCP par adresse MAC) plutôt que de la figer en dur dans TrueNAS, ce qui évite les conflits en cas de changement de réseau.

## Étape 6 : Créer votre premier pool de stockage ZFS

Rendez-vous dans Storage puis Pools et cliquez sur Create Pool. L’assistant graphique liste tous les disques physiques détectés. Sélectionnez ceux que vous voulez inclure (jamais le disque de démarrage), choisissez la topologie décidée à l’étape 3 (RAIDZ2 dans notre exemple), nommez votre pool (par exemple « tank », convention historique héritée de ZFS/Solaris) et validez. TrueNAS initialise le pool en quelques secondes, contrairement à un RAID matériel classique qui peut prendre des heures.

Une fois le pool créé, planifiez immédiatement un scrub périodique (Data Protection > Scrub Tasks). Le scrub relit l’intégralité des données du pool et corrige les erreurs de checksum détectées grâce à la redondance ZFS, un peu comme un contrôle technique régulier. La fréquence recommandée par la documentation communautaire tourne autour d’une fois par mois pour un usage domestique, et toutes les deux semaines pour un pool qui stocke des données professionnelles critiques. Un scrub sur un pool de 16 To avec des disques mécaniques classiques prend généralement entre 6 et 12 heures, pendant lesquelles les performances du NAS baissent légèrement mais restent utilisables.

Avant de charger vos données, lancez un test de santé sur chaque disque neuf avec `badblocks` ou un test S.M.A.R.T. long via l’interface TrueNAS (Storage > Disks > sélectionner le disque > Manual Test > Long). Cette étape prend plusieurs heures sur des disques de 10 To et plus, mais elle révèle les disques défectueux avant qu’ils ne contiennent des données critiques.

```
# Vérifier l'état du pool ZFS après création
zpool status tank
# Exemple de sortie attendue pour un pool RAIDZ2 sain
  pool: tank
 state: ONLINE
config:
        NAME        STATE     READ WRITE CKSUM
        tank        ONLINE       0     0     0
          raidz2-0  ONLINE       0     0     0
            sda     ONLINE       0     0     0
            sdb     ONLINE       0     0     0
            sdc     ONLINE       0     0     0
            sdd     ONLINE       0     0     0
errors: No known data errors
```
## Étape 7 : Configurer les partages SMB, NFS et les utilisateurs

Avec le pool créé, il faut maintenant définir des datasets (des sous-volumes du pool) et les partager sur le réseau. Dans Datasets, créez-en un par usage : « documents », « medias », « backups », par exemple. Chaque dataset peut avoir ses propres réglages de compression, quotas et permissions, ce qui évite qu’un usage sature tout le pool.

Pour partager un dataset avec des machines Windows ou macOS, utilisez SMB (menu Shares > Windows Shares SMB). Pour des machines Linux ou des NAS Synology/QNAP en réplication, NFS reste plus léger et plus rapide. Créez d’abord des comptes utilisateurs dédiés dans Credentials > Local Users : évitez d’utiliser le compte root pour l’accès quotidien aux partages, une mauvaise pratique de sécurité classique.

```
# Monter un partage NFS TrueNAS depuis un client Linux
sudo mount -t nfs 192.168.1.50:/mnt/tank/medias /mnt/nas-medias
# Ajouter le montage de façon permanente dans /etc/fstab
192.168.1.50:/mnt/tank/medias  /mnt/nas-medias  nfs  defaults  0  0
```
## Étape 8 : Automatiser les snapshots et la réplication

Un des plus gros avantages de ZFS face à un système de fichiers classique, ce sont les snapshots quasi instantanés et sans coût de performance notable. Dans Data Protection > Periodic Snapshot Tasks, planifiez une tâche par dataset important, avec une rétention adaptée : par exemple des snapshots horaires conservés 24h, des snapshots journaliers conservés 30 jours, et des snapshots mensuels conservés un an pour les données sensibles.

Les snapshots protègent contre la suppression accidentelle ou le ransomware, mais pas contre une panne matérielle totale du NAS lui-même : ils vivent sur le même pool. Pour une vraie protection, configurez une réplication ZFS (Data Protection > Replication Tasks) vers un second TrueNAS, un NAS distant, ou un service cloud compatible S3. C’est le seul mécanisme qui vous protège si le boîtier entier tombe en panne, brûle ou se fait voler.

## Étape 9 : Installer des applications Docker (Plex, Nextcloud, Immich)

TrueNAS Community Edition intègre depuis la version 24.10 un nouveau backend d’applications basé sur Docker, qui a remplacé l’ancien système Kubernetes des premières versions de SCALE. La migration s’est faite automatiquement pour les utilisateurs qui mettaient à jour depuis une version antérieure. La documentation officielle décrit ce système comme l’« Applications Market », accessible depuis le menu Apps, avec un catalogue par défaut baptisé « TRUENAS » qui regroupe des applications validées par iXsystems : Plex, Nextcloud, MinIO, Immich, Home Assistant et bien d’autres.

Comme le résume l’article de Computing for Geeks consacré aux options de NAS DIY : « TrueNAS is the free Community Edition, formerly called SCALE, is built on ZFS », une distinction utile à garder en tête face à des concurrents comme Unraid, qui n’utilisent pas ZFS en natif.

### Le catalogue TrueNAS Apps

Pour installer une application, allez dans Apps > Discover Apps, choisissez par exemple Nextcloud, et configurez le chemin de stockage vers un dataset dédié (créez un dataset « apps » séparé de vos données personnelles). L’interface graphique gère le déploiement du container, l’attribution du port réseau et les volumes de stockage sans passer par la ligne de commande. Pour des besoins plus personnalisés, TrueNAS accepte aussi l’import de fichiers docker-compose classiques :

```
# Exemple simplifié d'un fichier compose pour une app custom sur TrueNAS
version: "3.9"
services:
  monapp:
    image: monimage/exemple:latest
    ports:
      - "8080:80"
    volumes:
      - /mnt/tank/apps/monapp:/config
    restart: unless-stopped
```
## Étape 10 : Sécuriser l’accès à votre NAS

Un NAS mal sécurisé exposé sur Internet est une cible facile. Trois réglages sont non négociables. D’abord, activez l’authentification à deux facteurs (2FA) sur le compte root et tous les comptes admin, via System Settings > General > Two-Factor Authentication. Ensuite, remplacez le certificat auto-signé par un certificat valide : TrueNAS supporte l’intégration ACME (Let’s Encrypt) directement dans System Settings > Certificates, à condition d’avoir un nom de domaine pointant vers votre NAS.

