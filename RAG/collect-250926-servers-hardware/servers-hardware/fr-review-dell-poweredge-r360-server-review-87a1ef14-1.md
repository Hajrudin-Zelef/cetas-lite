---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r360-server-review-87a1ef14-1
title: "fr-review-dell-poweredge-r360-server-review-87a1ef14"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "ethernet", "gpu", "intel", "nvidia", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r360-server-review-87a1ef14.md
source_anchor: ""
source_lines: [1, 64]
sha256: 07b0386a2be48e4bd6145d22502bd024d38d5c9adf8a6c3481b1cecf683cca24
---

# fr-review-dell-poweredge-r360-server-review-87a1ef14

Le PowerEdge R360 est un serveur rack 1U à socket unique pour les applications Near Edge et PME. Il est destiné à offrir une valeur maximale et est alimenté par des processeurs Intel Xeon série E-2400, une mémoire DDR5 et la prise en charge PCIe Gen5.
Le PowerEdge R360 s'améliore par rapport au PowerEdge R350 précédent en offrant un processeur plus puissant, une plus grande bande passante mémoire, un stockage de démarrage supérieur (BOSS N-1 contre BOSS S-2) et la prise en charge d'un GPU de centre de données d'entrée NVIDIA A40 de 60 à 2 W. (Le PowerEdge R350 ne prenait pas en charge les GPU.)
Le PowerEdge R360 prend en charge quatre disques de 3.5 pouces ou huit disques de 2.5 pouces et deux emplacements PCIe Gen5 pour le stockage et l'extension. La version tour de ce serveur, 4.5U PowerEdge T360, offre plus de stockage avec huit disques 3.5 pouces et quatre emplacements PCIe (trois Gen4 et un Gen5). Le PowerEdge R360 propose également deux options RAID matérielles, exploitant le PERC H355 ou H755 via une carte de montage dédiée.
Spécifications Dell PowerEdge R360
| Processeur | Un processeur Intel Xeon série E-2400 avec jusqu'à 8 cœurs ou Un processeur Intel Pentium G7400/G7400T avec 2 cœurs | 
| Mémoire |  | 
| Contrôleurs de stockage |  | 
| Baies de disques | Baies avant :  | 
| Alimentations |  | 
| Options de refroidissement | refroidissement par air | 
| Ventilateurs | Jusqu'à 4 ventilateurs | 
| Dimensions |  | 
| Facteur de forme | Serveur rack 1U | 
| Gestion intégrée |  | 
| Biseau | Lunette de sécurité | 
| Logiciel OpenManage |  | 
| Mobilité | OuvrirGérer Mobile | 
| Intégrations OpenManage |  | 
| Sécurité |  | 
| Carte réseau intégrée | 2x LOM 1GbE | 
| Ports | Ports avant  Ports arrière  Ports internes :  | 
| PCIe |  | 
| Système d'exploitation et hyperviseurs |  | 
Conception Dell PowerEdge R360
Le PowerEdge R360 est un serveur rack 1U aux dimensions standard (1.68 x 18.97 x 24.57 pouces [HWD, avec cadre]). Les ports avant incluent un port USB Micro-AB pour iDRAC Direct et un port USB 2.0. Un cadre de sécurité pour empêcher l'accès à ces ports est facultatif.
À l’arrière se trouvent une prise Ethernet iDRAC dédiée, un autre USB 2.0, un USB 3.2 Gen1 et des ports VGA et série. En interne, le PowerEdge R360 dispose d’un seul port USB 3.2 Gen1. L'emplacement BOSS-N1 pour les lecteurs de démarrage M.2 dédiés est également visible ici, ce qui signifie que le stockage régulier n'a pas besoin d'être utilisé à cette fin. Les deux alimentations du serveur sont également visibles de l'arrière ; il arbore une paire d'unités redondantes remplaçables à chaud de 600 W Platinum ou 700 W Titanium.
En regardant à l'intérieur, la disposition de base du PowerEdge R360 place le processeur à l'arrière ; même le Xeon E-2488 haut de gamme se contente d'un refroidisseur passif. Le processeur a un TDP de seulement 95 watts. Quatre emplacements DIMM DDR5 à sa gauche prennent en charge 128 Go de RAM (4x 32 Go). La mémoire enregistrée (ECC) n'est pas prise en charge. Les deux emplacements d'extension PCIe se trouvent à côté du dissipateur thermique du processeur.
Quatre ventilateurs à remplacement simple le long de la ligne centrale assurent le refroidissement de l'air. Le carénage de circulation d'air à l'intérieur du châssis dirige l'air à travers le processeur, bien qu'il comporte une seule découpe pour diriger l'air à travers le matériel PCIe installé.
Les baies de disques sont en avance sur celles-ci ; le PowerEdge R360 peut accueillir quatre baies de 3.5 pouces ou huit baies de 2.5 pouces. Dell propose plusieurs contrôleurs RAID et configurations RAID 0, 1, 5, 6 ou 10 en usine.
Dell PowerEdge R360 contre PowerEdge R260
Le PowerEdge R360 est essentiellement une version pleine taille du PowerEdge R260 à faible profondeur . Ces deux serveurs sont destinés aux applications en périphérie de réseau et aux PME.
Dell propose également des versions tour de ces serveurs ; le PowerEdge T360 (dont nous publierons prochainement un test) est la version tour du PowerEdge R360 et prend en charge des disques de stockage supplémentaires et des emplacements d'extension. Il existe aussi le PowerEdge T160, plus léger . Ces serveurs sont tous conçus pour les applications PME et offrent des options de stockage adaptées plutôt que des capacités de processeur ou de mémoire importantes, avec un seul socket CPU et quatre emplacements DIMM.
Gestion et maintenance Dell PowerEdge R360
Consultez notre test du PowerEdge R260 pour une analyse approfondie du contrôleur d'accès à distance intégré Dell (iDRAC). L'iDRAC offre des fonctionnalités complètes de gestion à distance, permettant aux administrateurs de surveiller, de mettre à jour et de dépanner le serveur où qu'ils soient.
Performances Dell PowerEdge R360
Nous avons testé le PowerEdge R360 configuré comme suit :
- Windows Server standard 2022
- Intel Xeon E-2488 (8 cœurs/16 threads)
- 128 Go de mémoire DDR5-3600 ECC
- 2 disques SSD RAID 480 de 1 Go (système d'exploitation) ; Disque dur de 16 To de 3.5 pouces
- GPU Nvidia A2
Notre configuration possède le processeur le plus puissant, la quantité maximale de mémoire et un GPU, elle est donc aussi performante que possible. Étant donné que les configurations des disques de stockage peuvent varier considérablement, la plupart de nos tests se concentreront sur le processeur.
Lors de nos tests de performance, le PowerEdge R360 sera comparé au PowerEdge R260 , que nous avons testé avec le même processeur mais deux fois moins de RAM (64 Go). Notre autre modèle comparable est la tour PowerEdge T360, équipée d'une puce Xeon E-2414 (4 cœurs/4 threads) plus modérée et de 32 Go de RAM.
Processeur Blender 4.0
Blender est une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur. Étonnamment, le PowerEdge R260 a montré de bien meilleures performances que le PowerEdge R360. Comme on pouvait s'y attendre, le PowerEdge T360 quadricœur n'était pas en lice.
| Blender 4.0 (échantillons par minute, plus c'est haut, mieux c'est) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| Monster | 79.07 | 99.64 | 36.93 | 
| Brocanteur | 53.93 | 66.65 | 23.31 | 
| Salle de classe | 40.87 | 51.19 | 18.75 | 
Test de vitesse du processeur Blackmagic RAW
Nous avons également commencé à exécuter le test de vitesse RAW de Blackmagic, qui teste la lecture vidéo. Ces serveurs ne seraient probablement pas utilisés pour la lecture vidéo, mais le PowerEdge R260 s'est encore une fois révélé plus rapide que le PowerEdge R360 ici.
| Test de vitesse Blackmagic RAW (Plus c'est haut, mieux c'est) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| CPU 8K | 50 images/s | 65 images/s | 23 images/s | 
Compression à 7 zips
Le test de mémoire intégré dans l'utilitaire populaire 7-Zip montre un delta plus étroit entre le PowerEdge R360 et le PowerEdge R260, bien que ce dernier ait toujours l'avantage. Le PowerEdge T360 était encore une fois loin d'être proche, avec un processeur largement inférieur.
| 7-Zip Compression Benchmark (Plus c'est haut, mieux c'est) | Dell PowerEdge R360 | Dell PowerEdge R260 | Dell PowerEdge T360 | 
| Utilisation actuelle du processeur | 1,366 % | 1,336 % | 330 % | 
| Courant nominal/utilisation | 4.999 GIPS | 5.892 GIPS | 8.894 GIPS | 
| Courant | 68.279 GIPS | 78.716 GIPS | 29.364 GIPS | 
| Utilisation résultante du processeur | 1,361 % | 1,342 % | 331 % | 
| Évaluation/utilisation résultante | 5.035 GIPS | 5.871 GIPS | 8.885 GIPS | 
| Note résultante | 68.504 GIPS | 78.809 GIPS | 29.430 GIPS | 
| Décompression |  |  |  | 
| Utilisation actuelle du processeur | 1,597 % | 1,591 % | 394 % | 
