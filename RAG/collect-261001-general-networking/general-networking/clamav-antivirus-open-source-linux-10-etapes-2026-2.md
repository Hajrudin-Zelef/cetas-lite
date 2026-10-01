---
id: collect-261001-general-networking/general-networking/clamav-antivirus-open-source-linux-10-etapes-2026-2
title: "Debian / Ubuntu"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "distribution"]
source: docs/RAG/collect-261001-general-networking/clamav-antivirus-open-source-linux-10-etapes-2026.md
source_anchor: ""
source_lines: [51, 181]
sha256: 1ec59718c5f9373bb22c78a4d5163e61d479e986033262bafbee911bb8867368
---

# Debian / Ubuntu

```
# Debian / Ubuntu
sudo apt update
sudo apt install clamav clamav-daemon clamav-freshclam clamav-milter -y
# RHEL / Rocky Linux / AlmaLinux (dépôt EPEL requis)
sudo dnf install epel-release -y
sudo dnf install clamav clamav-update clamd clamav-milter -y
# Fedora
sudo dnf install clamav clamav-update clamd clamav-milter -y
```
Sur Debian et Ubuntu, l’installation du paquet `clamav-daemon` déclenche automatiquement un premier téléchargement de la base de signatures et tente de démarrer le service `clamav-freshclam`. Sur RHEL et dérivés, cette étape est manuelle : il faudra lancer freshclam une première fois vous-même, ce que nous verrons à l’étape suivante. Comptez entre 5 et 15 minutes pour cette première installation, l’essentiel du temps étant consommé par le téléchargement initial de la base virale, qui pèse plusieurs centaines de mégaoctets.

## Étape 2 : vérifier la version installée et le statut du cycle de vie

Avant d’aller plus loin, vérifiez quelle version vient d’être installée. Les dépôts des distributions LTS (Debian stable, Ubuntu LTS) livrent parfois une version plus ancienne que la dernière release amont.

```
clamscan --version
clamd --version
```
Exemple de sortie attendue avec une installation à jour :

`ClamAV 1.4.6/27671/Mon Aug 10 08:12:33 2026`
Le premier nombre après la version (ici 27671) correspond au numéro de la base de signatures, et la date qui suit indique quand cette base a été téléchargée pour la dernière fois. Si votre distribution propose une version antérieure à 1.4, comparez-la à la table du cycle de vie plus bas dans cet article : certaines branches comme 0.103 ou 1.2 ont déjà atteint leur fin de support de sécurité, et il devient urgent de migrer.

Si votre paquet système est trop ancien, deux options s’offrent à vous : activer un dépôt tiers plus récent (backports Debian, dépôt EPEL à jour pour RHEL) ou compiler ClamAV depuis les sources publiées sur le dépôt GitHub Cisco-Talos. Pour la majorité des déploiements en PME, s’en tenir aux paquets de la distribution avec des mises à jour régulières du système suffit amplement.

## Étape 3 : configurer freshclam pour la mise à jour des signatures

Un antivirus sans base de signatures à jour ne sert à rien. freshclam se configure via le fichier `/etc/clamav/freshclam.conf` (Debian/Ubuntu) ou `/etc/freshclam.conf` (RHEL). Ouvrez-le et vérifiez les paramètres suivants :

```
# /etc/clamav/freshclam.conf
DatabaseMirror database.clamav.net
DatabaseDirectory /var/lib/clamav
UpdateLogFile /var/log/clamav/freshclam.log
LogFileMaxSize 5M
LogTime yes
Checks 12
NotifyClamd /etc/clamav/clamd.conf
```
La directive `Checks 12` indique à freshclam de vérifier les mises à jour douze fois par jour (soit toutes les deux heures), un rythme raisonnable pour ne pas manquer les signatures publiées en urgence tout en restant respectueux des miroirs officiels. La ligne `NotifyClamd` est cruciale : elle indique à freshclam de signaler à clamd qu’il doit recharger la base en mémoire dès qu’une nouvelle version est téléchargée, sans quoi le démon continuerait à tourner avec une base obsolète jusqu’à son prochain redémarrage.

Lancez ensuite une première synchronisation manuelle pour vérifier que tout fonctionne :

```
sudo systemctl stop clamav-freshclam
sudo freshclam
sudo systemctl start clamav-freshclam
sudo systemctl enable clamav-freshclam
```
Une exécution réussie affiche une sortie proche de celle-ci, avec le téléchargement séquentiel des trois fichiers de la base virale :

```
ClamAV update process started at Mon Sep 3 09:14:02 2026
main.cvd database is up to date (version: 62, sigs: 6647427)
daily.cvd updated (version: 27671, sigs: 2043112)
bytecode.cvd updated (version: 335, sigs: 92)
Database updated (8690631 signatures) from database.clamav.net
```
Si la commande échoue avec un message de type `ERROR: Can't connect to port 443`, le souci vient presque toujours d’un pare-feu sortant trop restrictif ou d’un proxy d’entreprise non déclaré. Nous détaillerons ce cas précis dans la section dépannage.

## Étape 4 : configurer et démarrer le démon clamd

Le fichier de configuration principal de clamd se trouve dans `/etc/clamav/clamd.conf`. Les paramètres par défaut conviennent pour un premier test, mais quelques ajustements s’imposent avant une mise en production.

```
# /etc/clamav/clamd.conf
LocalSocket /var/run/clamav/clamd.ctl
FixStaleSocket yes
User clamav
MaxThreads 12
MaxConnectionQueueLength 200
StreamMaxLength 100M
MaxScanSize 400M
MaxFileSize 100M
ScanArchive yes
ArchiveBlockEncrypted no
ExcludePath ^/proc
ExcludePath ^/sys
LogFile /var/log/clamav/clamd.log
LogFileMaxSize 10M
LogTime yes
LogRotate yes
```
Deux paramètres méritent une attention particulière en production. `MaxThreads` contrôle le nombre de scans simultanés : sur un serveur mail à fort trafic, une valeur trop basse crée une file d’attente qui ralentit la livraison des messages, tandis qu’une valeur trop haute peut saturer la RAM sur une petite instance. `StreamMaxLength` et `MaxFileSize` définissent respectivement la taille maximale d’un flux reçu via le socket et la taille maximale d’un fichier scanné : au-delà, ClamAV ignore le fichier plutôt que de risquer une consommation mémoire excessive, un compromis à connaître si votre organisation échange régulièrement de gros documents.

Démarrez et activez le service :

```
sudo systemctl start clamav-daemon
sudo systemctl enable clamav-daemon
sudo systemctl status clamav-daemon
```
Le premier démarrage peut prendre 30 à 60 secondes, le temps que clamd charge intégralement la base de signatures en mémoire. Sur une machine avec 2 Go de RAM, surveillez la consommation mémoire à ce moment précis : c’est le pic le plus élevé du cycle de vie normal du service.

## Étape 5 : lancer un premier scan avec clamscan et le fichier de test EICAR

Avant de connecter ClamAV à quoi que ce soit de sensible, validez que la détection fonctionne réellement. Le standard de l’industrie pour ce test est le fichier EICAR, une chaîne de caractères inoffensive reconnue par tous les moteurs antivirus du marché comme un signal de test, sans être un vrai malware.

```
curl -o /tmp/eicar.com https://secure.eicar.org/eicar.com
clamscan /tmp/eicar.com
```
Une détection réussie affiche :

```
/tmp/eicar.com: Win.Test.EICAR_HDB-1 FOUND
----------- SCAN SUMMARY -----------
Known viruses: 8690631
Engine version: 1.4.6
Scanned directories: 0
Scanned files: 1
Infected files: 1
Data scanned: 0.00 MB
Data read: 0.00 MB (ratio 0.00:1)
Time: 8.234 sec (0 m 8 s)
Start Date: 2026:09:03 09:22:14
End Date:   2026:09:03 09:22:22
```
Si `Infected files: 1` apparaît, votre installation détecte correctement. Testez ensuite un scan récursif sur un répertoire réel, par exemple un dossier de documents partagés, pour valider les performances sur un volume représentatif :

`clamscan -r --infected --remove=no /srv/documents/`
L’option `--infected` limite l’affichage aux fichiers infectés (utile sur un gros volume), et `--remove=no` vous protège d’une suppression accidentelle tant que vous validez encore votre configuration. Supprimez immédiatement le fichier EICAR de test une fois vos vérifications terminées.

## Étape 6 : activer la surveillance en temps réel avec clamonacc

clamscan et clamd via socket fonctionnent en mode « à la demande » : quelqu’un ou quelque chose doit déclencher le scan. Pour une protection continue d’un répertoire critique, comme un partage Samba de dépôt de fichiers, clamonacc s’appuie sur fanotify pour intercepter chaque création ou modification de fichier et déclencher un scan automatique.

Activez d’abord le support on-access dans `clamd.conf` :

