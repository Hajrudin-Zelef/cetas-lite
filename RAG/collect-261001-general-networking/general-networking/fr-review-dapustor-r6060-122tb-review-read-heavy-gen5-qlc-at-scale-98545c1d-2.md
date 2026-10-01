---
id: collect-261001-general-networking/general-networking/fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d-2
title: "fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d.md
source_anchor: ""
source_lines: [56, 80]
sha256: efbc2fcc623a747b88524d54f18675495e3c3fde2cbd55748f87865e4c0773ba
---

# fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d

Lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Le DapuStor R6060 122.88 To a atteint 11 554,0 Mo/s au test de lecture séquentielle 128K, ce qui le place à nouveau en deuxième position derrière le Micron 6550 ION 61.44 To à 13 979,7 Mo/s.
Après ces deux modèles, les performances ont chuté brutalement pour le reste du groupe de comparaison : le Solidigm P5336 de 61.44 To a atteint 7 132,3 Mo/s, le Solidigm P5336 de 122.88 To 7 121,6 Mo/s et le DapuStor J5060 de 61.44 To 7 126,8 Mo/s. Seul le R6060, hormis le modèle Micron, a dépassé les 11 Go/s lors de ce test.
Latence de lecture séquentielle de 128 K (IODepth 64 / NumJobs 1)
Le DapuStor R6060 de 122.88 To a enregistré une latence de lecture séquentielle de 692.1 µs (128 Ko), se classant ainsi deuxième du groupe. Le Micron 6550 ION de 61.44 To a affiché la latence la plus faible (571.9 µs), tandis que le Solidigm P5336 de 122.88 To a enregistré la plus élevée (1123.0 µs). Les deux modèles Solidigm P5336 étaient quasiment identiques : 1123.0 µs pour la version 122.88 To et 1121.3 µs pour la version 61.44 To. Le DapuStor J5060 se situait entre les deux avec une latence de 1122.1 µs.
Le R6060 a conservé une belle avance sur le reste du peloton, même si Micron a toujours décroché le meilleur résultat.
Écriture aléatoire 64K
Le DapuStor R6060 de 122.88 To a affiché des performances quasi stables lors du test d'écriture aléatoire à 64 Ko. Il a débuté à 3 477,9 Mo/s et 55 600 IOPS avec un ratio de 1/1, puis a bondi à 3 915,2 Mo/s et 62 600 IOPS avec un ratio de 1/2. À partir de là, ses performances sont restées quasiment inchangées. La majeure partie du test s'est déroulée entre 3 913,7 Mo/s et 3 916,9 Mo/s, avec un résultat maximal de 3 916,9 Mo/s et 62 700 IOPS avec un ratio de 4/1. Même avec des paramètres plus exigeants, il a maintenu 3 914,6 Mo/s avec un ratio de 32/4, 3 913,8 Mo/s avec un ratio de 16/8 et 3 914,0 Mo/s avec un ratio de 32/8. Donc, en dehors du point de départ inférieur 1/1, le R6060 est resté essentiellement sur un plateau fixe pendant toute la durée du balayage.
En comparant les autres disques du tableau, le Micron 6550 ION 61.44 To a affiché des performances nettement supérieures, passant d'environ 2.4 Go/s à un peu plus de 10.3 Go/s. Derrière le R6060, le Solidigm P5336 122.88 To s'est maintenu autour de 3.0 Go/s, le DapuStor J5060 61.44 To a oscillé autour de 2.8 Go/s, et le Solidigm P5336 61.44 To a varié entre 2.5 et 2.6 Go/s. Le R6060 a ainsi conservé la deuxième place, devançant largement les autres disques non Micron.
Latence d'écriture aléatoire de 64 K
La latence a suivi la même tendance stable : 18 µs à 1/1, 31 µs à 1/2 et 2/1, 63 µs à 1/4, 2/2 et 4/1, puis 127 µs à 2/4, 4/2 et 8/1. Avec l’augmentation de la charge, elle est passée à 255 µs à 1/8, 2/8, 8/2 et 16/1, puis à 510 µs à 4/8, 8/4 et 32/1. Pour les charges les plus élevées, elle a atteint 1 021 µs à 8/8, 16/4 et 32/2, 2 043 µs à 16/8 et 32/4, et un pic de 4 087 µs à 32/8.
Le disque Micron a affiché la latence la plus faible, de loin, avec environ 1 600 µs à 32 bits/s. Le R6060 a terminé deuxième avec 4 087 µs. Le Solidigm P5336 de 122.88 To a atteint environ 5 100 µs, le J5060 environ 5 500 µs et le Solidigm P5336 de 61.44 To un peu plus de 6 000 µs. Ainsi, même si le R6060 était loin d'égaler le Micron en termes de latence, il a tout de même devancé les trois autres disques lorsque le test s'est intensifié.
Lecture aléatoire 64K
Le SSD DapuStor R6060 de 122.88 To a affiché une courbe de performance moins régulière en lecture aléatoire 64 Ko, mais ses résultats finaux étaient les meilleurs. Il a démarré à 381.4 Mo/s et 6 100 IOPS à 1/1, puis a atteint 748.2 Mo/s à 1/2, 1 007,1 Mo/s à 2/2, 1 431,8 Mo/s à 1/4 et 2 343,8 Mo/s à 2/4. Le débit a continué d'augmenter jusqu'à 2 767,5 Mo/s à 1/8, puis 3 668,5 Mo/s à 4/4, 4 060,2 Mo/s à 16/2, 4 750,0 Mo/s à 2/8, 6 433,3 Mo/s à 8/4 et 7 495,3 Mo/s à 4/8. De là, il a atteint 8 428,0 Mo/s à 32/2, 9 827,5 Mo/s à 16/4, 10 782,5 Mo/s à 8/8, 12 798,2 Mo/s à 16/8 et a culminé à 13 274,8 Mo/s à 32/8.
Comme l'indique le graphique ci-dessous, le Micron 6550 ION 61.44 To est resté en tête pendant une grande partie du test, atteignant un débit légèrement inférieur à 13.0 Go/s, bien que le R6060 ait réussi à terminer juste au-dessus du Micron à la toute fin.
Latence de lecture aléatoire de 64 K
Bien que plus instable que les autres disques en début de test, le R6060 a affiché de solides performances en fin de test. Il a commencé à 163 µs à 1/1, puis a mesuré 167 µs à 1/2, 174 µs à 1/4, 180 µs à 1/8, 217 µs à 2/8, 220 µs à 2/4, 249 µs à 2/2, 260 µs à 2/1, 287 µs à 4/8, 290 µs à 4/4 et 352 µs à 8/4. Les fréquences moyennes sont devenues moins nettes, avec 446 µs à 8/1, 505 µs à 16/2, 531 µs à 32/2, 574 µs à 16/1, 595 µs à 32/1, 700 µs à 16/8, 738 µs à 32/4 et un pic de 1 285 µs à 32/8.
Sur le graphique, Micron est resté globalement en dessous et a terminé autour de 1 200 µs à 32 bits/s. Le R6060 a terminé à 1 285 µs, ce qui était encore beaucoup plus bas que le J5060 et les deux disques Solidigm, qui ont tous dépassé les 2 200 µs à l'extrémité.
Lecture aléatoire 16K
Le SSD DapuStor R6060 de 122.88 To a obtenu d'excellents résultats au test de lecture aléatoire 16K après une augmentation progressive de la charge de travail. Il a démarré à 9 600 IOPS à 1/1, puis a progressé jusqu'à 18 000 à 2/1, 37 200 à 1/4, 58 200 à 8/1 et 72 800 à 1/8. Il a ensuite continué à progresser, atteignant 112 800 à 16/1, 133 300 à 4/4, 138 100 à 2/8 et 140 300 à 1/16, avant de connaître une accélération significative avec l'augmentation de la charge de travail. Il a ainsi atteint 211 600 à 32/1, 246 900 à 8/4, 256 500 à 4/8 et 261 200 à 2/16. Avec les paramètres les plus lourds, il a maintenu 436.0K à 16/4, 447.7K à 8/8 et 456.1K à 4/16, avant de terminer à 659.3K à 32/4, 671.3K à 16/8, 679.7K à 8/16, 784.5K à 32/8, 786.2K à 16/16 et de culminer à 817.7K IOPS à 32/16.
Face aux autres disques testés, le R6060 s'est avéré l'un des plus performants, même si Micron a conservé la première place. Le Micron 6550 ION 61.44 To a atteint environ 860 000 IOPS, devançant ainsi le R6060 en hautes performances. Cela dit, le R6060 a largement dominé les deux disques Solidigm P5336 pendant la majeure partie du test et a terminé nettement au-dessus d'eux.
Latence de lecture aléatoire de 16 K
La latence du R6060 est restée satisfaisante pendant la majeure partie du test, jusqu'à la fin. Elle a débuté à 104 µs sur un rapport signal/bruit de 1:1, puis s'est établie à 110 µs sur un rapport signal/bruit de 2:1 et à 124 µs sur un rapport signal/bruit de 4:1. Sur les réglages de puissance faible et moyenne, elle est restée stable, avec notamment 107 µs à 1/4, 109 µs à 1/8, 113 µs à 1/16, 115 µs à 2/8, 120 µs à 4:4, 130 µs à 8:4, 137 µs à 8:1 et 141 µs à 16:1. À mesure que la charge augmentait, la latence augmentait progressivement à 143 µs à 8/8, 148 µs à 16/4, 151 µs à 32/1, 189 µs à 8/16 et 193 µs à 16/8, avant de grimper plus rapidement aux combinaisons les plus lourdes, atteignant 196 µs à 32/4, 330 µs à 16/16, 336 µs à 32/8 et culminant à 642 µs à 32/16.
Cela a permis au R6060 de conserver de bonnes performances pendant la majeure partie du test, malgré une légère hausse en fin de parcours. Le Micron 6550 ION a quant à lui conservé la meilleure courbe de latence globale et est resté nettement en dessous sur l'ensemble du graphique.
Écriture aléatoire 16K
