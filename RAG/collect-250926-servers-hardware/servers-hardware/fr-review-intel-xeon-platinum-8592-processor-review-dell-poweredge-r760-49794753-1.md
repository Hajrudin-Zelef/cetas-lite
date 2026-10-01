---
id: collect-250926-servers-hardware/servers-hardware/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753-1
title: "fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "benchmark", "gpu", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753.md
source_anchor: ""
source_lines: [1, 78]
sha256: 6a91a715f28fbf322ebba0be53a2ebcb5052460dcbec1350083a4c9fa13cbf7a
---

# fr-review-intel-xeon-platinum-8592-processor-review-dell-poweredge-r760-49794753

Nous avons déjà testé le Dell PowerEdge R760 et l'avons trouvé, comme toujours chez Dell, extrêmement performant. Suite au récent lancement des processeurs Intel Xeon Scalable de 5e génération , nous avons pu nous procurer un ensemble de processeurs haut de gamme Intel Xeon Platinum 8592+. Nous nous sommes dit : quoi de mieux que de moderniser l'un de nos PowerEdge R760 pour combiner le meilleur des deux mondes et voir si la dernière innovation d'Intel mérite de faire des envieux dans le monde de l'entreprise ?
Spécifications Intel Xeon Platinum 8592+
Nous avons commencé avec la configuration fournie par Dell dans notre R760, qui était composée de deux processeurs Intel 6430 et de 1 To de RAM DDR5 4800, avant d'échanger les processeurs 8592+. Bien que le Xeon 6340 ne soit pas le haut de gamme de la famille Sapphire Rapids, nous allons quand même le comparer au Xeon 8592+. Nous alignons ces deux-là dans cette revue non pas pour les présenter comme des concurrents directs, mais pour montrer le potentiel des plates-formes existantes avec la sortie des nouveaux processeurs Emerald Rapids.
Nous avons apporté les chiffres du Sapphire Rapids Platinum 8480+ ainsi que ceux d'un HP ML350 à titre de comparaison avec le Xeon 8592+ pour obtenir des chiffres plus précis. Le 8480+ du ML350 Ge11 est à une place parmi les meilleurs processeurs de la gamme Sapphire Rapids. Il convient de noter que nous examinons ici uniquement les processeurs, sans comparer les serveurs eux-mêmes les uns aux autres.
|  | Processeur Intel Xeon Platinum 8592+ (R760) | Processeur Intel Xeon Gold 6430 (R760) | Processeur Intel Xeon Platinum 8480+ (ML350 G11) | 
|---|---|---|---|
| Génération | 5e, Rapides d'Émeraude | 4ème, Rapides Saphir | 4ème, Rapides Saphir | 
| Coloré | 64 | 32 | 56 | 
| Threads | 128 | 64 | 112 | 
| Fréquence de base | 1.9 GHz | 2.10 GHz | 2.0 GHz | 
| Turbo Fréquence | 3.9 GHz | 3.4 GHz | 3.8 GHz | 
| Cache | 320 MB | 60 MB | 105 MB | 
| Vitesse Intel UPI | 20 GT / s | 16 GT / s | 16 GT / s | 
| Liens UPI maximum | 4 | 3 | 4 | 
| TDP | 350 W | 270 W | 350 W | 
| Date de lancement | Q4 '23 | Q1 '23 | Q1 '23 | 
| Taille maximale de la mémoire | 4TB |  |  | 
| Types de mémoire | DDR5 5600 1 MT/s (XNUMX DPC) | Jusqu'à DDR5 4400 MT/s (1DPC et 2DPC) | Jusqu'à DDR5 4800 1 MT/s XNUMXDPC Jusqu'à DDR5 4400 2 MT/s XNUMXDPC | 
| Nombre maximum de canaux mémoire | 8 |  |  | 
| Révision PCIe | 5 |  |  | 
| Nombre maximum de voies PCIe | 80 |  |  | 
| Cœurs hautement prioritaires | 20 @ 2.1 GHz | 12 @ 2.2 GHz | 16 @ 2.1 GHz | 
| Cœurs à faible priorité | 44 @ 1.7 GHz | 20 @ 1.8 GHz | 40 @ 1.7 GHz | 
| Taille EPC maximale par défaut pour Intel SGX | 512 GB | 128 GB | 512 GB | 
| Page Produit | Lien | Lien | Lien | 
Performances du processeur Intel Xeon Platinum 8592+
Les spécifications système testées de haut niveau pour chaque processeur sont les suivantes. Chaque système possédait deux des processeurs répertoriés.
Xeon Platine 8592+
- Dell Poweredge R760
- 1 To DDR5 4800 XNUMX MHz
Xeon Gold 6430
- Dell Poweredge R760
- 1 To DDR5 4800 XNUMX MHz
Xeon Platine 8480+
- HP ML350 génération 11
- 256GB DDR5 4800MHz
Mixeur OptiX
Le premier est Blender OptiX, une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire CLI Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Les Xeon 8592+ connaissent un bon départ avec plus du double des performances des Xeon 6430, reflétant les données indiquées sur les fiches techniques. Les Xeon 8480+ se rapprochent cependant de 81 points du Xeon 8592+ en classe.
| Processeur Blender 4.0 | 2 Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Monster | 1115.057 | 540.039 | 943.300 | 
| Brocanteur | 780.408 | 361.066 | 627.662 | 
| Salle de classe | 556.550 | 278.228 | 475.144 | 
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs. Dans ce test, nous constatons environ le double des résultats sur les Emerald Rapids Xeon 8592+ par rapport aux Sapphire Rapids 6430. Nous avons constaté de meilleures performances monocœur du Xeon 8480+ que du Xeon 8592+, mais cela n’est pas vraiment surprenant car le Xeon 8480+ avait une vitesse d’horloge de base plus élevée. Le Xeon 8592+ a pris une longueur d'avance dans le test multicœur.
| Cinebench R23 | 2 Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Processeur multicœur | 110,498 | 69,663 | 79,164 | 
| Processeur monocœur | 1,144 | 1,022 | 1,461 | 
| Rapport PM | 96.63x | 68.17x | 54.20x | 
Cinebench 2024
Voici les résultats CPU pour la version 2024 de Cinebench.
Nous avons constaté un modèle de performances similaire entre les trois processeurs et Cinebench R23.
| Cinebench 2024 | 2 Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Processeur multicœur | 6,001 | 3,746 | 4,699 | 
| Processeur monocœur | 68 | 59 | 76 | 
| Rapport PM | 88.48x | 63.22x | 61.44x | 
Geekbench 6
Geekbench 6 est un outil d'évaluation multiplateforme mesurant les performances globales d'un système. Plus le score est élevé, meilleures sont les performances. Geekbench propose un test de performances du GPU, mais sans GPU, seuls les résultats du CPU sont disponibles.
*Nous avons eu des problèmes avec les Emerald Rapids Xeon 8592+ lors de ce test et celui-ci n'a pas abouti. Nous y reviendrons donc lorsque nous pourrons obtenir des chiffres de performances. En attendant, nous n'avons de numéros que pour les Xeon 6430 et Xeon 8480+*.
| Geekbench 6 | 2 Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| Benchmark CPU – Single Core | N/D | 1,488 | 1,939 | 
| Benchmark du processeur – Multicœur | N/D | 16,054 | 15,218 | 
Y-Cruncher
Y-cruncher est une application populaire d'analyse comparative et de tests de résistance lancée en 2009. Ce test est multithread et évolutif, calculant Pi et d'autres constantes jusqu'à des milliards de chiffres. Plus vite, c'est mieux dans ce test.
Nous constatons que les résultats suivent le même schéma que les autres tests, le Xeon 6430 étant à l'arrière et le Xeon 8480+ juste derrière les chiffres du Xeon 8592+.
| Y-Cruncher (plus bas est mieux) | 2 Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|
| 1 milliard | 4.239 secondes | 6.060 secondes | 5.136 secondes | 
| 2.5 milliard | 11.466 secondes | 16.896 secondes | 13.768 secondes | 
| 5 milliard | 25.325 secondes | 36.843 secondes | 29.889 secondes | 
| 10 milliard | 54.921 secondes | 80.574 secondes | 65.194 secondes | 
| 25 milliard | 156.923 secondes | 229.017 secondes | 186.841 secondes | 
Compression à 7 zips
L'utilitaire populaire 7-Zip dispose d'un test de mémoire intégré qui démontre très bien les performances du processeur. Dans ce test, nous l'exécutons avec une taille de dictionnaire de 128 Mo lorsque cela est possible. Comme prévu, nous constatons toujours de meilleurs résultats sur les processeurs Xeon 8592+.
