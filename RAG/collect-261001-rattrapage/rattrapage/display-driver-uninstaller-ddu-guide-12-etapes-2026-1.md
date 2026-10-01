---
id: collect-261001-rattrapage/rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026-1
title: "Calculer l'empreinte SHA-256 de l'archive (PowerShell)"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "arr", "gpu", "intel", "mai", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/display-driver-uninstaller-ddu-guide-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 50]
sha256: d06f6dd8ed18815e3c13f21024cbf1e863eaf87123c558ec48e425a1188d7975
---

# Calculer l'empreinte SHA-256 de l'archive (PowerShell)

Un écran noir au démarrage, des micro-saccades qui apparaissent après une mise à jour, une carte graphique qui refuse de tenir son overclock : neuf fois sur dix, le coupable n’est pas le matériel mais un pilote mal désinstallé. **Display Driver Uninstaller**, plus connu sous l’acronyme **DDU**, est l’outil de référence pour effacer la moindre trace d’un pilote graphique NVIDIA, AMD ou Intel avant d’en installer un neuf. Ce tutoriel vous guide pas à pas, de la sauvegarde jusqu’à la réinstallation propre, avec les commandes PowerShell, les commutateurs en ligne de commande et les pièges spécifiques à Windows 11 24H2.

Comptez environ 30 minutes pour un nettoyage complet en mode sans échec. À l’heure où le prix des cartes graphiques flambe en Europe, savoir entretenir proprement la vôtre vaut largement l’investissement en temps. Toutes les versions et données de ce guide ont été vérifiées au 17 août 2026, en tenant compte de la sortie de **DDU 18.1.5.6** le 17 juillet 2026 sur le site officiel Wagnardsoft, désormais la référence à jour.

## Qu’est-ce que Display Driver Uninstaller (DDU) ?

**Display Driver Uninstaller** est un utilitaire gratuit développé par Wagnard (studio Wagnardsoft) qui supprime intégralement les pilotes graphiques et leurs résidus : entrées de registre, fichiers, services, dossiers temporaires et clés de configuration que le désinstalleur classique de Windows laisse derrière lui. Là où le Panneau de configuration retire le paquet principal, DDU racle jusqu’aux dernières traces, ce qui en fait la méthode la plus fiable pour repartir d’une base saine.

La version stable au moment de la rédaction est **DDU 18.1.5.6**, publiée le 17 juillet 2026 sur le site officiel Wagnardsoft, qui a succédé coup sur coup à la 18.1.5.3 (10 mai 2026) puis à la 18.1.5.4 (29-30 mai 2026). Il s’agit d’une mise à jour de maintenance : amélioration de l’arrêt et de la suppression des services, corrections sur le programme d’installation et ajout de nouveaux commutateurs en ligne de commande (`-cleanrealtek` et `-cleansoundblaster`) pour la partie audio. L’outil reste extrêmement léger – environ 1,2 Mo pour la version portable et 1,6 Mo pour l’installateur – et ne nécessite aucune installation : il se lance directement depuis son dossier.

DDU prend en charge les trois fabricants de GPU : **NVIDIA**, **AMD** et **Intel** (y compris les cartes Intel Arc). Il sait aussi nettoyer certains pilotes audio (Realtek, Sound Blaster) souvent installés en même temps que les pilotes HDMI/DisplayPort. Côté système, il fonctionne de Windows 7 à Windows 11. Téléchargeable depuis le site officiel Wagnardsoft et son miroir Guru3D, son code source est consultable sur GitHub.

## Pourquoi nettoyer ses pilotes GPU en 2026 ?

Réinstaller par-dessus l’ancien pilote suffit dans la majorité des cas, mais plusieurs situations imposent un nettoyage complet avec **Display Driver Uninstaller**. La première est le **changement de fabricant** : passer d’une carte NVIDIA à une Radeon AMD, ou inversement, laisse cohabiter deux jeux de pilotes qui se disputent les mêmes fichiers système. Sans DDU, les conflits sont quasi garantis (écran noir, code 43, pilote qui plante).

La deuxième situation concerne les **pilotes corrompus**. Après une mise à jour ratée, une coupure de courant en pleine installation ou une infection nettoyée à la hâte, le pilote graphique peut se retrouver dans un état instable. Les symptômes typiques : artefacts à l’écran, plantages dans les jeux, écrans bleus pointant vers `nvlddmkm.sys` (NVIDIA) ou `amdkmdag.sys` (AMD), ou encore l’application NVIDIA/Adrenalin qui refuse de s’ouvrir.

Troisième cas, fréquent chez les joueurs : le **retour à une version antérieure**. Quand un nouveau pilote dégrade les performances ou casse la compatibilité d’un jeu précis, revenir proprement à la version précédente exige d’effacer la nouvelle d’abord. Enfin, si vous prévoyez d’overclocker votre carte, partir de pilotes propres élimine une variable parasite dans le diagnostic d’instabilité. Pour la suite, notre guide MSI Afterburner pour overclocker son GPU prend logiquement le relais une fois vos pilotes réinstallés.

Avec la pénurie qui touche aussi bien la mémoire que les cartes graphiques – voir notre dossier sur les prix de la RAM en 2026 – conserver son GPU actuel en parfait état logiciel est devenu un réflexe d’économie. Un pilote propre, c’est quelques images par seconde gagnées sans dépenser un euro.

## DDU face aux autres méthodes de désinstallation

Pourquoi recourir à **Display Driver Uninstaller** alors que Windows et les fabricants proposent leurs propres outils ? Parce qu’aucun d’eux ne va aussi loin. Le désinstalleur de Windows (Paramètres > Applications) retire le paquet principal mais laisse intacts des dizaines d’entrées de registre, de services et de fichiers. Le Gestionnaire de périphériques, avec sa case « Supprimer le pilote », fait à peine mieux. Quant aux options « installation propre » intégrées aux installeurs NVIDIA et AMD, elles ne nettoient que les paramètres de leur propre marque, en mode normal, et ratent les résidus profonds.

| Méthode | Fabricants couverts | Profondeur de nettoyage | Mode sans échec | Coût | 
|---|---|---|---|---|
| Désinstalleur Windows | Tous (paquet principal) | Faible – laisse registre et fichiers | Non | Gratuit, intégré | 
| Gestionnaire de périphériques | Tous (partiel) | Faible à moyenne | Non | Gratuit, intégré | 
| Installation propre NVIDIA | NVIDIA uniquement | Moyenne (paramètres NVIDIA) | Non | Gratuit | 
| AMD Cleanup Utility | AMD uniquement | Élevée pour AMD | Recommandé | Gratuit | 
| **Display Driver Uninstaller** | **NVIDIA, AMD, Intel (+ audio)** | **Maximale – registre, fichiers, services** | **Recommandé** | **Gratuit** | 

Le verdict est sans appel : pour un nettoyage réellement complet, multi-fabricants et automatisable en ligne de commande, **DDU** reste l’outil le plus exhaustif. L’AMD Cleanup Utility constitue une bonne alternative si vous restez exclusivement chez AMD, mais dès que vous changez de marque ou que vous voulez effacer un pilote audio HDMI récalcitrant, Display Driver Uninstaller s’impose. C’est aussi le seul à offrir un mode silencieux scriptable, ce qui en fait le favori des techniciens et des ateliers de réparation.

## Prérequis : versions, téléchargements et sauvegardes

Avant de lancer **DDU**, rassemblez le nécessaire. La règle d’or : **téléchargez vos nouveaux pilotes GPU AVANT de nettoyer les anciens**. Une fois DDU passé, votre carte tournera sur le pilote d’affichage basique de Microsoft (résolution réduite, pas d’accélération), et il vaut mieux ne pas dépendre du réseau à ce moment-là. Vérifiez aussi que **.NET Framework 4.8** ou une version supérieure est installé : Wagnardsoft l’exige depuis la version 18.1.5.2 du 11 avril 2026, faute de quoi l’exécutable refuse de se lancer. Le tableau suivant récapitule les éléments à préparer.

| Élément | Version / source | Remarque | 
|---|---|---|
| Display Driver Uninstaller | 18.1.5.4 (29 mai 2026) | Site Wagnardsoft ou miroir Guru3D uniquement | 
| Système d’exploitation | Windows 7 à Windows 11 (24H2) | Compte administrateur requis | 
| Décompresseur | 7-Zip 24.x ou WinRAR | L’archive DDU est auto-extractible (7-Zip SFX) | 
| Pilote NVIDIA | GeForce / Studio (le plus récent) | Téléchargé depuis le site officiel NVIDIA | 
| Pilote AMD | Adrenalin Edition | Téléchargé depuis le site officiel AMD | 
| Pilote Intel | Arc & Graphics | Pour iGPU et cartes Intel Arc | 
| Espace disque | ~1,5 Go libre | Pour DDU + paquet pilote décompressé | 

