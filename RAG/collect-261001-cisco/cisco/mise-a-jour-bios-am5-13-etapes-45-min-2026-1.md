---
id: collect-261001-cisco/cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026-1
title: "verification-bios-am5.ps1"
domain: cisco
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "attention"]
source: docs/RAG/collect-261001-cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [1, 52]
sha256: ec74b0249b57a4d0ca480762406f22d408edae32b682c0dd62c52f4d0af1628e
---

# verification-bios-am5.ps1

Ouvrir son PC pour la première fois après avoir changé de processeur, et tomber sur un écran noir. Ça arrive à des milliers d’utilisateurs AM5 chaque mois : un Ryzen 9000X3D flambant neuf posé sur une carte mère B650 achetée en 2023, dont le firmware date d’avant la sortie de la puce. Résultat, aucun signal vidéo, aucun bip, et souvent un retour au magasin qui aurait pu être évité en dix minutes.

La mise à jour du BIOS, ou plus précisément de l’UEFI qui l’a techniquement remplacé, reste l’étape la plus survolée d’un montage PC. Elle conditionne pourtant la compatibilité avec les nouveaux processeurs, la stabilité de la RAM DDR5 et parfois la sécurité de la machine. Sur le socket AM5, une plateforme qu’AMD a positionnée pour durer plusieurs générations comme nous le détaillions dans notre analyse de Zen 6 et de la longévité de l’AM5, cette opération devient même récurrente : chaque nouvelle génération de Ryzen, chaque kit mémoire un peu exotique, chaque correctif de sécurité repasse par une mise à jour du firmware.

Ce tutoriel détaille les 13 étapes pour flasher le BIOS d’une carte mère AM5 en toute sécurité, avec ou sans processeur installé, en couvrant les quatre grandes marques du marché (ASUS, MSI, Gigabyte, ASRock) ainsi que les erreurs qui transforment une opération de routine en carte mère hors service. Comptez environ 45 minutes, marge de dépannage comprise.

## Pourquoi mettre à jour le BIOS de votre carte mère AM5 en 2026

Le BIOS, pour Basic Input/Output System, est le tout premier logiciel exécuté au démarrage d’un PC, avant même le système d’exploitation. Sur les cartes mères modernes, il a été remplacé techniquement par l’UEFI (Unified Extensible Firmware Interface), plus rapide et plus riche en fonctionnalités, mais l’appellation BIOS a survécu dans le langage courant. C’est d’ailleurs celle que gardent la plupart des fabricants dans leurs propres menus et documentations.

Trois raisons poussent concrètement les utilisateurs d’AM5 à mettre à jour leur firmware en 2026.

La première est la compatibilité processeur. AMD a lancé plusieurs générations de Ryzen sur le même socket AM5 depuis 2022, et chaque nouvelle puce nécessite un microcode que les cartes mères plus anciennes n’embarquent pas d’origine. Gigabyte a par exemple publié la révision **AGESA 1.1.7.0 Patch A** pour rendre ses cartes X670, B650 et A620 compatibles avec le démarrage des Ryzen 9000, un cas typique où l’absence de mise à jour empêche purement et simplement le PC de démarrer, sans le moindre message d’erreur pour orienter l’utilisateur.

La deuxième raison est la compatibilité mémoire. La DDR5 reste plus exigeante que la DDR4 sur la validation des profils EXPO et XMP, et les fabricants publient régulièrement des révisions de firmware qui élargissent la liste des kits pris en charge ou stabilisent les fréquences élevées. Dans un contexte où les prix de la RAM ont explosé, comme nous le rapportions dans notre couverture de la pénurie et de la hausse de +171 % sur les modules DDR5, faire fonctionner correctement un kit déjà acheté compte plus que jamais plutôt que d’en racheter un neuf en cas d’incompatibilité.

La troisième raison, moins visible mais tout aussi réelle, concerne la sécurité. Un bulletin ASUS relayé par la communauté fin 2025 mentionnait des correctifs visant une vulnérabilité de dépassement de mémoire dans le composant PeCoffLoader, ainsi qu’une faille de vérification de signature de microcode AMD, corrigées via la révision **AGESA ComboAM5 PI 1.2.0.3**.

Pour un socket dont la longévité annoncée dépasse largement celle de la génération précédente, la mise à jour BIOS n’est donc plus une opération ponctuelle réservée aux passionnés d’overclocking. Elle devient un geste de maintenance normal, au même titre qu’une mise à jour Windows.

## Prérequis : ce qu’il vous faut avant de commencer

Avant de lancer la moindre mise à jour, réunissez les éléments suivants. Un flash BIOS interrompu reste la cause numéro un de carte mère rendue inutilisable, et la majorité des cas se résume à un prérequis oublié plutôt qu’à un vrai problème matériel.

**Côté matériel :** une clé USB de 4 à 32 Go formatée exclusivement en FAT32 (le format exFAT ou NTFS n’est reconnu par aucune interface de flash BIOS grand public), une alimentation électrique stable, et si possible un second appareil comme un smartphone pour consulter les instructions pendant l’opération, l’écran principal étant indisponible une bonne partie du temps.

**Côté logiciel :** Windows 10 ou 11 à jour pour identifier votre matériel, l’utilitaire CPU-Z en dernière version disponible pour confirmer le modèle exact de carte mère, et l’invite de commandes PowerShell intégrée à Windows, en version 5.1 ou supérieure et préinstallée depuis Windows 10. Aucune installation tierce n’est nécessaire pour l’essentiel de la procédure.

**Côté informations :** notez précisément la référence complète de votre carte mère, révision comprise. Les fabricants publient parfois des firmwares distincts pour une même gamme selon la révision de carte (Rev 1.0, Rev 1.1…), et flasher le mauvais fichier figure parmi les erreurs les plus fréquentes détaillées plus bas. Vérifiez aussi l’état de la pile CMOS si votre carte mère a plus de deux ans, une pile faible pendant un flash augmentant le risque de coupure inopinée.

Enfin, prévoyez 45 minutes sans interruption. Le flash proprement dit ne dure que 3 à 10 minutes selon les fabricants, mais l’identification du matériel, le téléchargement, la vérification et la reconfiguration post-update représentent l’essentiel du temps réel passé sur l’opération.

## Étape 1 : identifier le modèle exact de votre carte mère

La méthode la plus fiable ne consiste pas à se fier de mémoire au nom inscrit sur la boîte, mais à lire l’information directement via un outil logiciel. Si le PC démarre encore normalement, ouvrez une invite PowerShell en administrateur et exécutez la commande suivante.

`Get-CimInstance Win32_BaseBoard | Select-Object Manufacturer, Product, Version`
Exemple de sortie obtenue sur une carte B650 :

```
Manufacturer                         Product                    Version
------------                         -------                    -------
Micro-Star International Co., Ltd.   MAG B650 TOMAHAWK WIFI     1.0
```
Ce résultat donne le fabricant, le nom commercial exact et la révision matérielle, les trois informations nécessaires pour télécharger le bon firmware. Si Windows ne démarre pas, l’alternative consiste à ouvrir le boîtier et lire l’étiquette sérigraphiée directement sur le circuit imprimé, généralement située entre les slots PCIe et les ports SATA, ou à photographier la référence avec un smartphone avant de refermer le boîtier.

Attention aux variantes de nommage. Une même carte mère peut exister en version simple, WiFi ou « Plus », avec des firmwares non interchangeables. Un ASUS ROG Strix B650E-F Gaming WiFi et sa version sans le suffixe WiFi ne partagent pas nécessairement le même fichier BIOS, même si les deux cartes se ressemblent visuellement.

## Étape 2 : repérer la version BIOS actuellement installée

Une fois le modèle confirmé, il faut connaître la version de firmware déjà en place, pour savoir si une mise à jour est réellement nécessaire et éviter de flasher une version identique ou plus ancienne. Toujours en PowerShell :

