---
id: collect-261001-cisco/cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026-2
title: "verification-bios-am5.ps1"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [53, 122]
sha256: 4a8072f59a02ffe9eca8681078bded1df2a65d2dc37ed0d22eeca0d62ed2665d
---

# verification-bios-am5.ps1

`Get-CimInstance Win32_BIOS | Select-Object Manufacturer, SMBIOSBIOSVersion, ReleaseDate`
Vous pouvez aussi consulter cette information directement dans le BIOS lui-même, en appuyant sur Suppr ou F2 au démarrage (la touche exacte s’affiche brièvement à l’écran pendant les deux premières secondes du POST), puis en repérant le numéro de version, presque toujours situé en haut à droite de l’écran principal ou dans un onglet « Main ».

Comparez ensuite ce numéro à la dernière version disponible sur le site du fabricant. Les conventions de nommage varient énormément d’une marque à l’autre : MSI utilise des suites comme « 7D78v1C », ASUS des formats numériques comme « 3421 », Gigabyte des versions comme « F23 ». Ne vous fiez jamais à un simple tri alphabétique ou numérique, référez-vous systématiquement à la date de publication indiquée sur la page de support officielle.

## Étape 3 : vérifier la compatibilité CPU et RAM sur la liste QVL

C’est l’étape que la majorité des utilisateurs sautent, et c’est justement celle qui évite la panne la plus fréquente : le nouveau processeur acheté qui ne démarre pas une fois installé.

Chaque fabricant publie deux documents distincts sur la page produit de sa carte mère : la liste de support CPU, souvent nommée « CPU Support List », et la liste de qualification mémoire, la fameuse QVL pour Qualified Vendor List. La première indique, pour chaque processeur, la version BIOS minimale requise pour démarrer. La seconde fait de même pour les kits de RAM testés en usine à une fréquence donnée.

Le cas le plus documenté sur AM5 concerne justement les Ryzen 9000. D’après un guide technique consacré à la question, la révision **AGESA ComboAM5 PI 1.1.7.0** constitue le minimum pour qu’un Ryzen 9000 termine simplement sa séquence de démarrage, tandis que la révision **1.3.0.1**, publiée en 2025, ajoute la prise en charge des variantes X3D de cette même génération. Sans cette information, un utilisateur qui installe un Ryzen 9000X3D sur une carte flashée en usine avant la sortie de la puce se retrouve face à un écran noir.

Pour la RAM, consultez la QVL même si votre kit affiche fièrement un profil EXPO. Un kit non testé peut fonctionner normalement, tourner en dessous de sa fréquence annoncée, ou plus rarement empêcher le démarrage. Notre couverture de la validation par MSI de modules DDR5 chinois CXMT à 8 200 MT/s sur AM5 illustre bien à quel point ces listes de compatibilité évoluent au fil des mises à jour BIOS, parfois plusieurs mois après la sortie commerciale d’un kit mémoire.

## Étapes 4 à 6 : télécharger, vérifier et préparer le fichier BIOS

Une fois le modèle, la version actuelle et la compatibilité confirmés, la préparation du fichier se déroule en trois temps.

**Étape 4, le téléchargement.** Rendez-vous exclusivement sur le site officiel du fabricant, jamais sur un site tiers ou un forum même si le fichier semble identique. Recherchez votre référence exacte, en incluant la révision matérielle notée à l’étape 1, puis téléchargez la version la plus récente listée dans l’onglet BIOS de la page support.

**Étape 5, la vérification d’intégrité.** Un fichier corrompu pendant le téléchargement peut passer inaperçu jusqu’au moment du flash, où il provoque un échec en plein milieu de l’opération. Si le fabricant publie un hash de contrôle SHA-256 sur sa page de téléchargement, vérifiez-le avant d’aller plus loin :

`certutil -hashfile "C:\Downloads\B650-TOMAHAWK-WIFI.zip" SHA256`
Comparez la chaîne obtenue avec celle publiée par le fabricant, caractère par caractère. Si elles ne correspondent pas, retéléchargez le fichier plutôt que de tenter le flash avec un doute.

**Étape 6, la préparation de la clé USB.** Formatez-la en FAT32 via l’invite de commandes, la méthode la plus fiable pour éviter les échecs liés à l’explorateur Windows sur les clés de plus de 32 Go :

```
diskpart
list disk
select disk 1
clean
create partition primary
format fs=fat32 quick
assign
exit
```
Remplacez « disk 1 » par le numéro exact de votre clé USB, visible dans la liste affichée par la commande « list disk ». Une erreur de sélection à cette étape efface le mauvais disque, vérifiez donc deux fois la taille affichée avant de valider quoi que ce soit.

Décompressez ensuite le fichier téléchargé et copiez le fichier .bin ou .cap directement à la racine de la clé, sans le placer dans un sous-dossier. Certaines méthodes de flash, comme le M-Flash de MSI, exigent en plus un renommage précis documenté dans le manuel de la carte mère.

## Étapes 7 à 9 : flasher sans CPU installé grâce au BIOS FlashBack

L’avantage majeur des cartes mères AM5 récentes tient dans leur capacité à flasher le BIOS sans processeur, sans RAM et sans carte graphique installés. Une aubaine quand le CPU tout juste acheté n’est justement pas reconnu par le firmware d’origine, ce qui serait sinon un problème difficile à résoudre sans emprunter un processeur compatible.

Chaque fabricant a sa propre appellation et sa propre procédure, mais le principe reste identique : brancher l’alimentation sans démarrer le PC, insérer la clé USB dans un port spécifique du panneau arrière, puis maintenir un bouton dédié jusqu’à ce qu’une diode commence à clignoter.

### ASUS BIOS FlashBack

Sur les cartes ASUS, branchez le câble d’alimentation 24 broches et le connecteur CPU 8 broches sans allumer le PC. Insérez la clé USB dans le port identifié « BIOS » à l’arrière du boîtier, jamais un port USB classique. Maintenez le bouton BIOS FlashBack environ trois secondes jusqu’à ce que la diode commence à clignoter. Le clignotement indique que le flash est en cours, son arrêt signale la fin de l’opération. La procédure complète est détaillée dans le guide officiel ASUS pour les cartes AM5.

### MSI Flash BIOS Button

La procédure MSI suit la même logique : fichier à la racine d’une clé FAT32 insérée dans le port BIOS dédié, alimentation branchée, PC éteint, puis pression sur le bouton Flash BIOS à l’arrière du boîtier. Selon la documentation officielle MSI pour les cartes AM5, l’opération dure généralement entre 5 et 10 minutes, la diode clignotante confirmant sa progression jusqu’à extinction complète.

### Gigabyte Q-Flash Plus

Gigabyte utilise le même principe sous le nom Q-Flash Plus, disponible en particulier sur les cartes X670, B650 et A620 compatibles avec la révision AGESA 1.1.7.0 Patch A mentionnée plus haut. Le fichier BIOS doit être renommé « GIGABYTE.bin » avant d’être copié sur la clé, une exigence propre à cette marque qui fait échouer la procédure si elle n’est pas respectée à la lettre.

### ASRock BIOS Flashback

ASRock propose une fonction équivalente sur ses cartes AM5 haut et moyen de gamme, avec un bouton dédié et un port USB identifié. La documentation exacte varie sensiblement d’un modèle à l’autre chez ce fabricant, il est donc recommandé de consulter le manuel PDF spécifique à votre référence plutôt que de suivre une procédure générique trouvée en ligne.

Dans les quatre cas, ne débranchez jamais l’alimentation et ne retirez jamais la clé USB tant que la diode clignote. Une interruption à ce stade est la cause la plus documentée de carte mère rendue inutilisable après une tentative de mise à jour.

## Étapes 10 à 13 : flash classique depuis le BIOS existant et reconfiguration

Si votre PC démarre déjà normalement avec l’ancien firmware, la méthode « classique » depuis l’interface UEFI reste plus rapide que le FlashBack.

