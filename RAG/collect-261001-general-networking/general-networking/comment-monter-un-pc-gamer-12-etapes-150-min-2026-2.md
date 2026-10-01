---
id: collect-261001-general-networking/general-networking/comment-monter-un-pc-gamer-12-etapes-150-min-2026-2
title: "Rechercher les pilotes GPU disponibles via winget"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "intel", "nvidia"]
source: docs/RAG/collect-261001-general-networking/comment-monter-un-pc-gamer-12-etapes-150-min-2026.md
source_anchor: ""
source_lines: [38, 82]
sha256: 155406a515e1679bb21acbf7f9f6bc007ad54afa091100b4ee439dbce0121422
---

# Rechercher les pilotes GPU disponibles via winget

| Composant | Entrée de gamme (~950 €) | Milieu de gamme (~1 450 €) | Haut de gamme (~2 100 €+) | 
|---|---|---|---|
| Processeur | AMD Ryzen 5 9600X | AMD Ryzen 7 9800X3D | AMD Ryzen 7 9800X3D | 
| Carte graphique | Nvidia RTX 5060 | Nvidia RTX 5070 | RTX 5070 Ti ou AMD RX 9070 XT | 
| Mémoire RAM | 16 Go DDR5-5600 | 32 Go DDR5-6000 CL30 | 32 Go DDR5-6000 CL30 | 
| Stockage | 1 To NVMe Gen4 | 1 To NVMe Gen4 | 2 To NVMe Gen5 | 
| Alimentation | 550 W, 80+ Bronze, ATX 3.0 | 750 W, 80+ Gold, ATX 3.1 | 850 W, 80+ Gold, ATX 3.1 | 
| Refroidissement | Ventirad à air | Ventirad haut de gamme ou AIO 240 mm | AIO 280-360 mm | 
| Boîtier | ATX moyen tour, façade aérée | ATX moyen tour, façade mesh | ATX grand tour, façade mesh | 
| Usage visé | 1080p élevé/ultra | 1440p élevé/ultra | 1440p ultra / 4K élevé | 

Ces trois paliers reprennent les familles de composants les plus documentées à la mi-2026, en s’appuyant sur les guides de référence publiés par Newegg et sur les repères déjà utilisés dans notre couverture des prix composants. Ces repères bougent vite : Tom’s Hardware relevait en septembre 2026 une RTX 5070 tombée à 799 $ au plus bas contre un MSRP de 549 $, tandis que la RTX 5060 Ti 16 Go, pourtant pensée pour l’entrée de gamme, oscillait déjà entre 550 et 800 $ selon SetupLevel. Un point de vigilance : sur la configuration d’entrée, mieux vaut ne pas descendre sous 16 Go de RAM en 2026, la plupart des jeux récents consommant déjà 12 à 14 Go en arrière-plan avec Windows 11.

Sur l’achat lui-même, mieux vaut privilégier des enseignes qui acceptent les retours sans frais et qui indiquent clairement l’état du composant (neuf scellé, reconditionné, déstockage). Une carte graphique d’occasion vendue sans historique d’usage reste un pari risqué, en particulier sur les modèles qui ont pu servir au minage par le passé : demandez systématiquement les températures observées et la durée d’utilisation avant d’acheter d’occasion.

## Bien choisir son boîtier et sa carte mère avant l’achat

Deux composants ne figurent pas toujours en tête de liste alors qu’ils conditionnent tout le reste du montage : le boîtier et la carte mère. Le format de la carte mère (ATX, mATX ou ITX) doit correspondre à ce que le boîtier annonce officiellement, pas seulement « sembler rentrer ». Vérifiez aussi la longueur maximale de carte graphique supportée, exprimée en millimètres : une RTX 5070 Ti ou une RX 9070 XT dépasse fréquemment les 300 mm, une donnée à comparer avec la fiche technique du boîtier avant l’achat plutôt qu’après réception.

Sur la ventilation, une façade perforée ou en mesh laisse circuler beaucoup plus d’air qu’une façade pleine en verre trempé, au prix d’un peu plus de bruit ambiant. Si un refroidissement liquide AIO est prévu, vérifiez également la compatibilité radiateur du boîtier : 240 mm se loge presque partout, 280 et 360 mm demandent un boîtier annoncé compatible, en façade ou en partie haute selon les modèles.

Côté carte mère, le chipset détermine surtout les fonctionnalités disponibles plutôt que la performance brute. Sur plateforme AMD AM5, les chipsets d’entrée (B650, B850) couvrent l’essentiel des besoins d’un PC gamer, tandis que les chipsets X670 ou X870 ajoutent davantage de voies PCIe et de ports USB, utiles surtout en cas de stockage multiple ou de cartes d’extension. Côté Intel LGA1851, la logique est similaire entre B860 et Z890, ce dernier étant surtout pertinent pour l’overclocking mémoire poussé. Vérifiez enfin le nombre de slots M.2 disponibles : deux emplacements suffisent à la plupart des usages, un troisième devient utile si vous prévoyez d’ajouter du stockage sans toucher à la carte graphique.

## Étape 1 : préparer le poste de travail et vérifier la compatibilité

Dégagez une table stable, si possible non recouverte de moquette (l’électricité statique s’y accumule davantage). Sortez le boîtier, la carte mère encore dans son emballage antistatique, et gardez chaque composant dans sa boîte jusqu’au moment de l’installer : cela évite les chocs et limite les risques de perdre une vis. Avant de sortir le tournevis, reprenez la liste de compatibilité une dernière fois : socket du CPU conforme à la carte mère, kit mémoire présent dans la QVL, alimentation disposant des connecteurs PCIe adaptés à la carte graphique choisie (12V-2×6 natif ou adaptateur fourni).

Touchez une partie métallique non peinte du boîtier avant chaque manipulation de composant, ou portez un bracelet antistatique relié à une masse. Ce geste simple évite l’essentiel des décharges électrostatiques susceptibles d’endommager une carte mère ou un processeur.

## Étape 2 : installer le processeur (CPU) sur la carte mère

Posez la carte mère à plat sur son emballage antistatique. Soulevez le levier de rétention du socket, repérez le petit triangle ou l’encoche gravée sur un coin du processeur, et alignez-le avec le repère correspondant sur le socket. Le CPU doit se poser sans aucune pression : s’il ne descend pas naturellement, c’est qu’il n’est pas aligné, jamais qu’il faut forcer. Une fois posé, refermez le levier jusqu’au clic.

La manipulation diffère légèrement selon la plateforme. Sur socket AMD AM5, les broches se trouvent sur le socket de la carte mère et non sur le processeur, ce qui rend le CPU plus difficile à endommager mais le socket plus fragile. Sur LGA1851 côté Intel, c’est l’inverse : les broches sont sur le socket, recouvertes d’un capot de protection à retirer seulement à la pose du processeur. Le guide officiel d’Intel rappelle d’ailleurs de ne jamais toucher les broches du socket, même hors tension.

## Étape 3 : appliquer la pâte thermique et fixer le refroidissement

Si le ventirad ou l’AIO fourni n’a pas de pâte thermique pré-appliquée sur son embase en cuivre, déposez un point de la taille d’un petit pois au centre du processeur. La pression de fixation du dissipateur suffit à l’étaler uniformément. Inutile de l’étaler soi-même au doigt ou avec une carte : ce geste introduit plus souvent des bulles d’air qu’il n’en retire.

Pour un ventirad à air, fixez le support selon le schéma du manuel, en serrant les vis en croix (jamais une à fond avant l’autre) pour répartir la pression uniformément. Pour un refroidissement liquide AIO, la procédure est différente : orientation du radiateur, sens de la pompe, purge du circuit. Nous détaillons chaque sous-étape, y compris le choix entre montage en intake ou en exhaust, dans notre guide dédié à l’installation d’un watercooling AIO.

Le sens de montage des ventilateurs de boîtier compte presque autant que le refroidisseur lui-même. La configuration la plus courante place des ventilateurs en admission (intake) en façade avant, pour apporter de l’air frais directement sur la carte graphique et le processeur, et des ventilateurs en extraction (exhaust) à l’arrière et sur le dessus du boîtier, là où l’air chaud s’accumule naturellement. Une flèche moulée sur le cadre en plastique de chaque ventilateur indique le sens de rotation des pales, pratique pour vérifier l’orientation avant de refermer le boîtier plutôt qu’après le premier test de stress.

## Étape 4 : installer la mémoire RAM

