---
id: collect-261001-general-networking/general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026-4
title: "Windows (PowerShell)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-general-networking/veracrypt-chiffrer-un-disque-en-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [145, 227]
sha256: dad5f90aad236961b2f83bb144deb953f9ec8ab605f629f559fba2948a04be2f
---

# Windows (PowerShell)

Un fichier-clé (keyfile) est un fichier quelconque (image, document, fichier audio) utilisé en complément ou à la place d’un mot de passe pour déverrouiller un volume. Sans ce fichier précis, le volume reste indéchiffrable même avec le bon mot de passe. Dans la fenêtre de montage, cliquez sur **Use keyfile**, ajoutez un ou plusieurs fichiers depuis **Add Files**, puis cochez l’option correspondante lors du montage.

L’usage le plus courant consiste à stocker le fichier-clé sur une clé USB distincte de l’ordinateur : sans cette clé physiquement branchée, impossible de monter le volume, même en connaissant le mot de passe par cœur. C’est une forme simple d’authentification à deux facteurs « fait maison ». La contrepartie est évidente : perdre ce fichier-clé sans en avoir conservé de copie de sauvegarde équivaut à perdre l’accès aux données chiffrées de façon définitive, aussi certainement qu’en oubliant le mot de passe. Conservez toujours une copie du fichier-clé sur un support de sauvegarde séparé et hors ligne.

## Étape 11 : Automatiser le montage en ligne de commande

VeraCrypt propose une interface en ligne de commande complète, utile pour scripter le montage et le démontage plutôt que de passer par l’interface graphique à chaque fois. Voici la syntaxe de base sous Linux et macOS :

```
# Monter un volume avec mot de passe et fichier-clé
veracrypt --text --keyfiles="/media/usb/cle.key" --pim=0 \
  --mount coffre.hc /mnt/coffre
# VeraCrypt demandera alors :
Enter password for coffre.hc:
# Démonter proprement
veracrypt --text --dismount /mnt/coffre
# Lister les volumes actuellement montés
veracrypt --text --list
```
Le paramètre `--pim` (Personal Iterations Multiplier) permet d’augmenter manuellement le nombre d’itérations de dérivation de clé au-delà de la valeur par défaut, pour renforcer la résistance à la force brute au prix d’un déverrouillage plus lent. Sur un script automatisé destiné à tourner sans surveillance, ajoutez l’option `--non-interactive` associée à un mot de passe fourni par une variable d’environnement plutôt qu’en clair dans le script, afin d’éviter qu’il n’apparaisse dans l’historique des commandes ou dans les journaux système.

## Étape 12 : Sauvegarder l’en-tête de volume pour éviter la perte de données

Chaque volume VeraCrypt commence par un en-tête chiffré qui contient les clés nécessaires pour déchiffrer le reste du volume. Si cet en-tête est endommagé, par exemple à cause de secteurs défectueux ou d’une coupure de courant pendant une écriture, l’intégralité du volume peut devenir illisible, même avec le bon mot de passe. VeraCrypt propose justement un outil de sauvegarde d’en-tête, accessible via **Tools > Backup Volume Header** dans l’interface graphique.

Générez cette sauvegarde juste après la création de chaque volume important, et conservez le fichier résultant sur un support totalement distinct du volume lui-même : un en-tête de secours stocké sur le même disque que le volume qu’il protège ne sert à rien en cas de panne physique du support. En cas de corruption, la fonction **Restore Volume Header** permet de restaurer l’accès à partir de cette sauvegarde, à condition de connaître toujours le mot de passe d’origine.

## Étape 13 : Mesurer les performances avec le benchmark intégré

VeraCrypt intègre un outil de benchmark accessible via **Tools > Benchmark**, qui mesure la vitesse de chiffrement et de déchiffrement de chaque algorithme directement sur votre matériel. Des tests communautaires partagés sur le forum SourceForge du projet donnent un ordre de grandeur utile, à titre indicatif puisque les résultats varient selon le processeur et la charge du système au moment du test.

| Algorithme | Débit observé (MB/s) | Remarque | 
|---|---|---|
| AES (simple) | Environ 650 à 1 000+ | Le plus rapide, accélération matérielle AES-NI quasi systématique | 
| Serpent (simple) | Environ 390 à 560 | Pas d’accélération matérielle dédiée | 
| Twofish (simple) | Environ 380 | Performance proche de Serpent | 
| Serpent(AES) [cascade] | Environ 350 | Cascade à deux algorithmes | 
| Twofish(Serpent) [cascade] | Environ 290 | Cascade à deux algorithmes | 
| Serpent(Twofish(AES)) [cascade] | Environ 280 | Triple cascade, sécurité maximale théorique | 

Sur un SSD NVMe récent avec AES-NI actif, la différence de vitesse ressentie au quotidien entre un disque chiffré avec AES et un disque non chiffré reste minime pour un usage bureautique courant : l’accélération matérielle absorbe l’essentiel du surcoût de calcul. L’écart se creuse nettement dès que vous activez une cascade à deux ou trois algorithmes, ou sur du matériel plus ancien dépourvu d’AES-NI, où le chiffrement logiciel pur peut diviser le débit par cinq ou plus.

## Projet complet : un coffre-fort numérique avec sauvegarde automatisée

Pour mettre en pratique l’ensemble des étapes précédentes, voici un projet complet et fonctionnel : un conteneur VeraCrypt dédié aux documents confidentiels, déverrouillé par un mot de passe combiné à un fichier-clé stocké sur une clé USB, avec un script qui monte le volume, synchronise un dossier source vers le coffre chiffré, puis démonte proprement le tout. Ce script peut ensuite être appelé automatiquement par une tâche planifiée (cron sous Linux/macOS, Planificateur de tâches sous Windows avec WSL).

```
#!/bin/bash
# backup-chiffre.sh — synchronise un dossier vers un coffre VeraCrypt chiffré
set -euo pipefail
VOLUME="/home/user/coffre.hc"
MOUNT_POINT="/mnt/coffre"
KEYFILE="/media/usb/cle.key"
SOURCE_DIR="/home/user/Documents/confidentiel"
LOG_FILE="/home/user/backup-chiffre.log"
if [ ! -f "$KEYFILE" ]; then
  echo "$(date '+%F %T') - Fichier-cle introuvable, abandon" >> "$LOG_FILE"
  exit 1
fi
mkdir -p "$MOUNT_POINT"
veracrypt --text --keyfiles="$KEYFILE" --pim=0 --protect-hidden=no \
  --mount "$VOLUME" "$MOUNT_POINT" --non-interactive \
  --password="$VERACRYPT_PASSWORD"
if mountpoint -q "$MOUNT_POINT"; then
  rsync -av --delete "$SOURCE_DIR/" "$MOUNT_POINT/"
  sync
  veracrypt --text --dismount "$MOUNT_POINT"
  echo "$(date '+%F %T') - Sauvegarde chiffree terminee avec succes" >> "$LOG_FILE"
else
  echo "$(date '+%F %T') - Echec du montage du volume VeraCrypt" >> "$LOG_FILE"
  exit 1
fi
```
Le mot de passe est ici lu depuis la variable d’environnement `VERACRYPT_PASSWORD`, définie séparément (par exemple dans un fichier lu uniquement par root, avec des permissions restreintes à 600) plutôt qu’écrit en clair dans le script. Pour planifier ce script toutes les nuits à 2h du matin via cron, ajoutez une ligne à la crontab de l’utilisateur concerné :

```
# crontab -e
0 2 * * * VERACRYPT_PASSWORD="$(cat /root/.backup-pass)" /home/user/backup-chiffre.sh
```
Ce projet combine les briques essentielles vues dans ce tutoriel : un conteneur chiffré (étape 3), une authentification renforcée par fichier-clé (étape 10), et l’automatisation en ligne de commande (étape 11). Vous pouvez l’étendre en ajoutant une rotation de sauvegardes horodatées, ou en synchronisant le résultat vers un second support physique conservé hors site.

## 5 erreurs fréquentes à éviter avec VeraCrypt

