---
id: collect-250926-servers-hardware/servers-hardware/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43-2
title: "fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/clean4/fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43.md
source_anchor: ""
source_lines: [25, 89]
sha256: 069184717e16459e340b8a64133fedb89bb895b525bf8146b1edd7e4f334992a
---

# fr-review-intel-xeon-6-review-sierra-forest-6780e-6766e-65f60f43

Prenant en charge les dernières technologies telles que DDR5, PCIe 5.0 et CXL 2.0, les processeurs Xeon 6 garantissent que les centres de données sont prêts pour les avancées futures et peuvent s'intégrer de manière transparente aux solutions matérielles et logicielles émergentes.
Des solutions modernes pour des problèmes modernes
Les processeurs Intel Xeon 6 Sierra Forest représentent une avancée majeure dans l'architecture des centres de données. En combinant hautes performances, efficacité énergétique et évolutivité, ils répondent aux besoins multiformes des centres de données modernes, ouvrant la voie à une efficacité opérationnelle améliorée, à des coûts réduits et à la capacité de répondre aux demandes croissantes de l'IA et d'autres applications à forte intensité de calcul.
Carte des références Intel Xeon 6 E-core
Toutes les références Sierra Forest comportent des E-Cores avec 88 voies PCIe Gen5/CXL.
| SKU | Coloré | GHz de base | TurboGHz | Turbo-GHz maximum | Mo de cache L3 | Watts TDP | Max Scala. | Vitesse mémoire DDR5 1DPC | 
|---|---|---|---|---|---|---|---|---|
| 6780E | 144 | 2.2 | 3.0 | 3.0 | 108 | 330 | 2S | 6400 | 
| 6766E | 144 | 1.9 | 2.7 | 2.7 | 108 | 250 | 2S | 6400 | 
| 6756E | 128 | 1.8 | 2.6 | 2.6 | 96 | 225 | 2S | 6400 | 
| 6746E | 112 | 2.0 | 2.7 | 2.7 | 96 | 250 | 2S | 5600 | 
| 6740E | 96 | 2.4 | 3.2 | 3.2 | 96 | 250 | 2S | 6400 | 
| 6731E | 96 | 2.2 | 3.1 | 3.1 | 96 | 250 | 1S | 5600 | 
| 6710E | 64 | 2.4 | 3.2 | 3.2 | 96 | 205 | 2S | 5600 | 
Tests de performances Intel Xeon 6
Notre laboratoire a échantillonné deux ensembles de processeurs pour cette revue, le 6780E et le 6766E. Intel a fourni une plate-forme serveur QCT pour les tests. Nous souhaitons noter quelques mises en garde clés dans nos données. La plate-forme serveur elle-même était instable dans bon nombre de nos environnements de test. Par exemple, Windows Server 2022 ne fonctionnerait pas correctement, nous avons donc utilisé 2025.
Les processeurs sont des échantillons précoces et non des processeurs neufs dans la boîte. En tant que tel, nous n’avons pas pu exécuter toute notre batterie de tests et n’avons pas eu le temps avant la levée de l’embargo de résoudre correctement certains des problèmes, de performances et de stabilité que nous constatons. En tant que telles, les données suivantes doivent être considérées comme instructives et non définitives. Nous attendrons la livraison des plates-formes par les constructeurs OEM d'Intel avant de porter un jugement plus concluant sur les capacités des processeurs Sierra Forest.
Intel nous a expédié un système Quanta QuantaGrid D55Q-2U comme plate-forme de test pour présenter les nouveaux processeurs. Ce serveur propose un fond de panier à connexion directe NVMe de 24 baies prenant en charge les SSD U.2 Gen4/Gen5.
Notre configuration de test comprenait 16 x 16 Go de RAM DDR5-6400 et des modules Micron MTC10F1084S1RC64BDY.
Xeon 6780E et Xeon 6766E
- Quanta QuantaGrid D55Q-2U
- 256GB DDR5 6400MHz
Xeon Platine 8592+
- Dell Poweredge R760
- 1 To DDR5 4800 XNUMX MHz
Xeon Gold 6430
- Dell Poweredge R760
- 1 To DDR5 4800 XNUMX MHz
Xeon Platine 8480+
- HP ML350 génération 11
- 256GB DDR5 4800MHz
Avec les tout premiers gremlins de la plate-forme de test en jeu, les nouveaux processeurs Xeon 6780E et 6766E ont été testés dans un environnement Windows Server 2025, tandis que les anciens processeurs ont été testés dans Windows Server 2022.
Mélangeur OptiX 4.0
Dans Blender OptiX, nous avons 3 tests différents : Monster, Junkshop et Classroom. Chez Monster, nous avons vu les 6780E et 6766E dépasser les 8592+ de 21 % et 14 %, avec une différence de 18 % entre les deux. Sur Junkshop, le 6780E était en avance de 8592 % sur le 10.5+, le 6766E étant derrière le 8592+ de seulement 0.35 %. Il y avait une différence de 10.8 % entre le 6780E et le 6766E sur Junkshop. Pour la partie salle de classe, les 6780E et 6766E ont devancé les 8592+ de 20 % et 22 % avec une différence de 10 % entre les 6780E et 6766E.
| Processeur Blender 4.0 | 2x Xeon 6780E (256 Go DDR5) | 2x Xeon 6766E (256 Go DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 
|---|---|---|---|---|---|
| Monster | 1410.463 | 1297.715 | 1115.057 | 943.300 | 540.039 | 
| Brocanteur | 862.418 | 777.716 | 780.408 | 627.662 | 361.066 | 
| Salle de classe | 696.543 | 628.960 | 556.550 | 475.144 | 278.228 | 
Cinebench R23
Pour Cinebench R23, les performances multicœurs ont montré le 6780E et le 6766E derrière le 8592+ de 38 % et 42 %, respectivement. Pour les performances monocœur, le 6780E était 24 % derrière le 8592+ et le 6766E était 31 % derrière. La différence entre le 6780E et le 6766E est tombée à environ 5 % pour le score multicœur et à 18 % pour le monocœur.
| Cinebench R23 | 2x Xeon 6780E (256 Go DDR5) | 2x Xeon 6766E (256 Go DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 
|---|---|---|---|---|---|
| Processeur multicœur | 67,984 | 64,326 | 110,498 | 79,164 | 69,663 | 
| Processeur monocœur | 873 | 793 | 1,144 | 1,461 | 1,022 | 
| Rapport PM | 77.91x | 81.10x | 96.63x | 54.20x | 68.17x | 
Cinebench 2024
Pour Cinebench 2024, le score multicœur a connu une baisse d'environ 55 % du 8592+ au 6780E et une baisse de 61 % du 8592+ au 6766E. Cela présentait également une différence de 13 % entre le 6780E et le 6766E sur Multi-Core. Pour les scores monocœur, les 6780E et 6766E n'avaient qu'une différence de 5 %, les 6780E et 6766E étant respectivement 37 % et 34 % derrière le 8592+.
| Cinebench R23 | 2x Xeon 6780E (256 Go DDR5) | 2x Xeon 6766E (256 Go DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| Processeur multicœur | 2,687 | 2,347 | 6,001 | 3,746 | 4,699 | 
| Processeur monocœur | 43 | 45 | 68 | 59 | 76 | 
| Rapport PM | 62.85x | 52.65x | 88.48x | 63.22x | 61.44x | 
Y-Cruncher
Y-cruncher est une application populaire d'analyse comparative et de tests de résistance lancée en 2009. Ce test est multithread et évolutif, calculant Pi et d'autres constantes jusqu'à des milliards de chiffres. Plus vite, c'est mieux dans ce test.
Les nouveaux 6780E et 6766E fonctionnent un peu plus lentement que l'Emerald Rapids 8592+, mais ce ne sont pas exactement des concurrents directs. Le 6780E fonctionnait environ 13 % plus vite que le Xeon Gold 6430 pour le test 1 milliard, mais environ 39 % plus lentement que le Xeon Platinum 8592+. Le 6766E fonctionnait 19 % plus lentement que le Gold 6439 et 42 % plus lentement que le Platinum 8592+.
| Y-Cruncher (plus bas est mieux) | Xeon 6780E (256 Go DDR5) | Xeon 6766E (256 Go DDR5) | 2x Xeon Platinum 8592+(ER) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Gold 6430(SR) (R760 – 1 To DDR5 4800 XNUMX MHz) | 2x Xeon Platinum 8480+(SR) (ML350 G11 – 256 Go DDR5 4400 XNUMX MHz) | 
|---|---|---|---|---|---|
| 1 milliard | 6.927 secondes | 7.254 secondes | 4.239 secondes | 6.060 secondes | 5.136 secondes | 
| 2.5 milliard | 17.898 secondes | 19.507 secondes | 11.466 secondes | 16.896 secondes | 13.768 secondes | 
| 5 milliard | 38.454 secondes | 41.116 secondes | 25.325 secondes | 36.843 secondes | 29.889 secondes | 
| 10 milliard | 81.146 secondes | 87.403 secondes | 54.921 secondes | 80.574 secondes | 65.194 secondes | 
| 25 milliard | 217.530 secondes | 238.813 secondes | 156.923 secondes | 229.017 secondes | 186.841 secondes | 
| 50 milliard | 565.913 secondes | 502.245 secondes | N/D | N/D | N/D | 
7-Zip
