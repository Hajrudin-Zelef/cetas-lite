---
id: collect-261001-general-networking/general-networking/comment-monter-un-pc-gamer-12-etapes-150-min-2026-3
title: "Rechercher les pilotes GPU disponibles via winget"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["gpu", "amd", "attention", "intel"]
source: docs/RAG/collect-261001-general-networking/comment-monter-un-pc-gamer-12-etapes-150-min-2026.md
source_anchor: ""
source_lines: [83, 130]
sha256: 5d36f1b10d7a740479e46212f22d3e7d2ed83f36ba69e986849f280e714a147b
---

# Rechercher les pilotes GPU disponibles via winget

Repérez dans le manuel de la carte mère les emplacements à utiliser en priorité pour un fonctionnement en double canal : il s’agit en général des slots A2 et B2 (le deuxième et le quatrième en partant du processeur), pas des deux premiers. Ouvrez les clips à une extrémité, alignez l’encoche du module avec le détrompeur du slot, et appuyez fermement et uniformément aux deux extrémités jusqu’au clic des deux clips. Une barrette mal enclenchée est l’une des causes les plus fréquentes d’un PC qui ne démarre pas.

À ce stade, la RAM tourne à sa fréquence de base JEDEC (souvent 4 800 ou 5 200 MT/s), pas encore à sa vitesse annoncée sur la boîte. L’activation du profil XMP (Intel) ou EXPO (AMD) se fait plus tard, dans le BIOS. Pour un kit DDR5-6000 CL30, qui représente le compromis prix/performance le plus courant sur les configurations 2026, mieux vaut aussi vérifier au préalable sa présence dans la liste de compatibilité mémoire de la carte mère, en particulier sur plateforme AM5.

## Étape 5 : installer le stockage SSD NVMe M.2

Repérez le slot M.2 le plus proche du processeur, généralement le plus rapide (PCIe 5.0 sur les cartes récentes, jusqu’à environ 14 000 Mo/s en lecture séquentielle, contre environ 7 000 Mo/s pour du PCIe 4.0). Retirez la vis ou le clip de maintien du dissipateur M.2 s’il y en a un, insérez le SSD à environ 30 degrés dans le connecteur, puis abaissez-le et fixez-le avec sa vis. Le guide d’installation SSD de Crucial détaille la même logique de montage, transposable d’un format M.2 à l’autre.

Un point technique à vérifier dans le manuel de la carte mère : certains slots M.2 secondaires partagent leurs voies PCIe avec un port d’extension ou avec le second slot M.2. Installer un SSD dans le mauvais emplacement peut désactiver un port SATA ou réduire la bande passante d’un autre composant sans qu’aucun message d’erreur ne s’affiche. Si vous avez déjà suivi une procédure similaire, par exemple pour installer un SSD M.2 sur PS5, le geste de manipulation du connecteur reste identique, seule la vérification de compatibilité change.

## Étape 6 : fixer la carte mère dans le boîtier

Installez d’abord le cache I/O (I/O shield) fourni avec la carte mère dans la découpe arrière du boîtier, en appuyant sur ses quatre coins jusqu’au clic. Attention aux languettes métalliques internes, elles coupent facilement. Vérifiez ensuite que les entretoises (standoffs) déjà présentes dans le boîtier correspondent au format de la carte mère (ATX, mATX ou ITX) : une entretoise oubliée sous une zone sans trou de vis peut créer un court-circuit contre le dos de la carte.

Positionnez la carte mère en l’inclinant légèrement pour faire coïncider ses ports avec le cache I/O, puis abaissez-la à plat. Vissez d’abord la vis centrale pour stabiliser la carte, puis les autres en suivant un ordre en croix, sans jamais forcer sur un tournevis électrique qui pourrait riper et rayer la carte.

Profitez de cette étape, carte mère encore accessible sans composant autour, pour repérer les passe-câbles à l’arrière du plateau (les petits caoutchoucs ou découpes prévues derrière le boîtier). Faire transiter dès maintenant les câbles d’alimentation par ces passages, plutôt qu’une fois la carte graphique installée, évite d’avoir à retravailler tout le cheminement des câbles en fin de montage.

## Étape 7 : installer l’alimentation (PSU) et préparer le câblage

La plupart des boîtiers modernes logent l’alimentation en bas, ventilateur orienté vers le bas si le boîtier dispose d’une grille d’aération sous le PSU, ou vers le haut dans le cas contraire. Faites passer les câbles nécessaires par les passages prévus à l’arrière du boîtier avant de brancher quoi que ce soit : anticiper le cheminement des câbles à ce stade facilite énormément la suite.

Connectez le câble ATX 24 broches à la carte mère, puis le connecteur d’alimentation processeur (EPS, 4+4 ou 8 broches selon la carte) situé près du socket CPU. Ce connecteur est souvent le plus facile à oublier une fois la carte graphique en place, mieux vaut donc le brancher avant l’étape suivante.

```
Ordre de branchement conseillé pour l'alimentation :
1. ATX 24 broches ......... vers la carte mère
2. EPS 4+4 ou 8 broches ... vers l'alimentation du processeur
3. SATA ................... vers les disques additionnels, si besoin
4. PCIe 12V-2x6 ou 6+2 .... vers la carte graphique (étape suivante)
5. Câbles façade / ventilateurs ... en tout dernier
```
Sur une alimentation certifiée ATX 3.1, le connecteur natif 12V-2×6 remplace avantageusement les anciens adaptateurs multi-broches pour les cartes graphiques récentes, avec un seul câble capable de délivrer jusqu’à 600 W. Vérifiez qu’il est enfoncé jusqu’au clic : un connecteur d’alimentation GPU mal enclenché reste l’une des causes les plus documentées de redémarrages sous charge.

## Étape 8 : installer la carte graphique (GPU)

Repérez le slot PCIe x16 principal, généralement le plus proche du processeur. Retirez les caches d’extension à l’arrière du boîtier correspondant à l’emplacement de la carte : deux caches pour une carte double slot, trois pour une carte triple slot, ce qui couvre la majorité des modèles actuels comme la RTX 5070 Ti ou la RX 9070 XT. Ouvrez le clip de rétention à l’extrémité du slot, présentez la carte bien à plat au-dessus du connecteur, puis appuyez uniformément sur les deux extrémités jusqu’à entendre le clic du clip.

Vissez le bracket arrière de la carte au boîtier, puis branchez le ou les câbles d’alimentation PCIe. Une fois la carte graphique installée, basculez systématiquement les câbles d’affichage (HDMI ou DisplayPort) sur ses propres sorties vidéo, jamais sur celles de la carte mère : ces dernières sont désactivées dès qu’un GPU dédié est détecté, et un écran resté branché dessus restera noir sans qu’aucune panne ne soit en cause.

Les cartes graphiques haut de gamme actuelles dépassent souvent 1,3 kg avec leurs radiateurs massifs à trois ventilateurs, un poids suffisant pour faire fléchir légèrement le PCB dans le temps si la carte reste sans soutien. Un support anti-affaissement (GPU brace), vendu séparément ou parfois fourni avec le boîtier, cale la carte par en dessous et évite toute tension prolongée sur le connecteur PCIe. Ce n’est pas indispensable pour un modèle plus compact et plus léger, mais cela devient pertinent dès qu’une carte triple ventilateurs entre dans la configuration.

## Étape 9 : brancher les connecteurs de façade et les ventilateurs

C’est l’étape la plus fastidieuse, mais aussi celle où les erreurs de branchement ont le moins de conséquences graves. Sur la carte mère, un bloc de broches regroupe les connecteurs façade : Power SW (bouton d’allumage, sans polarité), Reset SW, Power LED et parfois HDD LED. Le manuel de la carte mère indique toujours le schéma exact, broche par broche : gardez-le ouvert pendant cette étape plutôt que de vous fier à la mémoire.

Deux connecteurs méritent une attention particulière. Le connecteur USB 3.0 façade (bloc bleu de 19 ou 20 broches) est plus fragile et plus facile à mal orienter que les autres. Un branchement à l’envers peut endommager la carte mère, donc vérifiez le détrompeur avant de forcer quoi que ce soit. Les connecteurs RGB, ensuite, existent en deux standards non interchangeables : 12V à 4 broches pour du RGB classique, 5V à 3 broches pour de l’ARGB adressable. Brancher un périphérique 5V sur un en-tête 12V peut l’endommager instantanément.

