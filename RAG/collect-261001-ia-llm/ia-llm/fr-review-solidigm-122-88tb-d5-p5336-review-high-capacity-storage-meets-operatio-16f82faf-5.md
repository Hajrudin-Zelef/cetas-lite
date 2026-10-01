---
id: collect-261001-ia-llm/ia-llm/fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf-5
title: "fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-ia-llm/fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf.md
source_anchor: ""
source_lines: [125, 149]
sha256: 425a5250dc00731bdfd7ffb400ce63f2dd90943b3e5598bb296dd6f847ea6cce
---

# fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf

Lors du test d'écriture aléatoire de 64 Ko à faible concurrence (1-1), les trois disques affichent des performances similaires. Le Micron 6550 affiche 2,485.97 39,780 Mo/s et 5336 122.88 IOPS. Le Solidigm P2,429.93 38,880 To suit avec 5336 61.44 Mo/s et 2,412.90 38,610 IOPS, tandis que le PXNUMX XNUMX To est légèrement en retrait avec XNUMX XNUMX Mo/s et XNUMX XNUMX IOPS.
À mesure que la charge de travail évolue vers une concurrence plus élevée (32-8), les différences deviennent plus marquées. Le modèle 122.88 To atteint 3,121.54 49,950 Mo/s et 61.44 2,654.46 IOPS, surpassant le modèle 42,470 To, qui atteint 17.6 122 Mo/s et 6550 10,070.71 IOPS. Cela représente une augmentation de 161,130 % du débit pour le disque de XNUMX To, plus grande capacité, et une évolutivité plus efficace sous une pression d'écriture aléatoire plus importante. Le Micron XNUMX se démarque nettement avec XNUMX XNUMX Mo/s et XNUMX XNUMX IOPS.
Latence d'écriture aléatoire de 64 K
À faible charge (1-1), les trois disques affichent une latence identique de 0.025 ms. Sous la charge de travail (32-8), le Solidigm P5336 122.88 To enregistre 5.121 ms, contre 6.026 ms pour le P5336 61.44 To. Cela se traduit par une réduction de 15 % de la latence pour le modèle 122 To de plus grande capacité. Le Micron 6550 maintient une latence nettement inférieure à 1.588 ms, affichant une meilleure réactivité en cas de forte simultanéité d'écriture aléatoire.
Lecture aléatoire 64K
Lors du test de lecture aléatoire de 64 K à charge minimale (1-1), le Micron 6550 atteint 482.09 Mo/s et 7,710 5336 IOPS. Le Solidigm P61.44 299.40 To suit avec 4,790 Mo/s et 5336 122.88 IOPS, tandis que le P274.04 4,390 To affiche 8.5 Mo/s et 61.44 XNUMX IOPS, soit une baisse de performances de XNUMX % par rapport au modèle XNUMX To à cette profondeur.
En forte concurrence (32-8), le P5336 122.88 To offre 7,124.69 113,995 Mo/s et 61.44 7,125 IOPS, tandis que le 114,000 To atteint un débit proche de 6550 13,153.64 Mo/s et environ 210,460 XNUMX IOPS. À ce niveau, il n'y a pas de différence de performances significative entre les deux capacités. Le Micron XNUMX continue d'évoluer, atteignant XNUMX XNUMX Mo/s et XNUMX XNUMX IOPS.
Latence de lecture aléatoire de 64 K
À une profondeur (1-1) et avec un nombre de tâches égal, le Micron 6550 affiche la latence la plus faible, soit 0.129 ms. Le Solidigm P5336 61.44 To suit avec 0.208 ms, tandis que le P5336 122.88 To affiche une latence légèrement supérieure, soit 0.228 ms, ce qui représente une augmentation de 9.6 % pour cette capacité supérieure. Sous une charge plus importante (32-8), les deux modèles Solidigm affichent une latence identique, soit 2.245 ms, ne montrant aucun avantage lié à cette capacité accrue. Le Micron 6550 maintient une latence bien inférieure, soit 1.217 ms, pour cette exécution.
Écriture aléatoire 16K
À faible charge (1-1), le Solidigm P5336 122.88 To offre un débit de 549.14 Mo/s et 35,145 5336 IOPS. Le P61.44 1,036.53 To affiche des performances nettement supérieures avec 66,338 122.88 Mo/s et 47 61.44 IOPS. Le modèle 6550 To affiche ainsi une bande passante et des IOPS inférieurs d'environ 856.61 % à ceux du modèle 54,823 To. Le Micron XNUMX se situe entre les deux disques Solidigm, avec XNUMX Mo/s et XNUMX XNUMX IOPS.
À haute concurrence (32-8), le modèle 122.88 To maintient un débit de 549.14 Mo/s et 35,145 1 IOPS, sans aucune variation par rapport à ses performances 1-61.44. La version 2,542.36 To, quant à elle, atteint 162,711 363 Mo/s et 122 6550 IOPS, soit une augmentation de 10,295.66 % du débit par rapport au disque 658,922 To. Le Micron XNUMX domine le classement général, atteignant XNUMX XNUMX Mo/s et XNUMX XNUMX IOPS.
Latence d'écriture aléatoire de 16 K
À (1-1), le Solidigm P5336 122.88 To affiche une latence de 0.028 ms, tandis que le P5336 61.44 To est plus rapide à 0.015 ms. Le Micron 6550 se situe entre les deux à 0.018 ms. Cela se traduit par une latence 86 % supérieure pour le modèle 122 To par rapport au Solidigm 61 To à charge minimale. Sous une charge élevée (32-8), le modèle 122.88 To maintient un temps de réponse stable de 0.028 ms, indiquant une absence de mise à l'échelle. Le modèle 61.44 To atteint 1.572 ms, reflétant une pression accrue en simultanéité, mais aussi un gain de débit significatif. Le Micron 6550 reste efficace à 0.388 ms, affichant une meilleure réactivité en cas de sollicitation maximale en écriture aléatoire.
Lecture aléatoire 16K
À charge minimale (1-1), le Micron 6550 atteint 188.80 Mo/s et 12,083 5336 61.44 IOPS. Le Solidigm P126.55 8,100 To suit avec 5336 Mo/s et 122.88 125.87 IOPS, tandis que le P8,060 0.5 To enregistre XNUMX Mo/s et XNUMX XNUMX IOPS. Les deux disques Solidigm affichent des performances quasiment identiques, avec une différence inférieure à XNUMX %, ce qui indique l'absence d'avantage lié à la capacité à cette profondeur.
À une concurrence plus élevée (32-16), le Micron 6550 atteint 13,053.35 835,420 Mo/s et 61.44 7,063.02 IOPS. Le Solidigm 452,030 To atteint 122.88 6,855.59 Mo/s et 438,760 2.9 IOPS, légèrement devant le modèle XNUMX To, qui atteint XNUMX XNUMX Mo/s et XNUMX XNUMX IOPS. Cela représente une baisse de débit de XNUMX % pour le disque Solidigm, plus grand, sous cette charge de travail.
Latence de lecture aléatoire de 16 K
À (1-1), le Micron 6550 enregistre la latence la plus faible, soit 0.082 ms. Les deux modèles Solidigm P5336 suivent avec respectivement 0.123 ms et 0.124 ms, avec une différence de moins de 1 %. Sous forte charge (32-16), le Micron maintient une mise à l'échelle efficace à 0.612 ms, tandis que le Solidigm 61.44 To atteint 1.132 ms et le 122.88 To atteint 1.165 ms. Cela représente une augmentation de 2.9 % de la latence pour le disque Solidigm de plus grande capacité par rapport au modèle 61 To, ce qui indique une légère baisse d'efficacité en cas de concurrence maximale.
Lecture aléatoire 4K
À charge minimale (1-1), le Micron 6550 offre 57.71 Mo/s et 14,770 5336 IOPS. Le Solidigm P61.44 38.21 To suit avec 9,782 Mo/s et 5336 122.88 IOPS, tandis que le P37.93 9,710 To est légèrement en retrait avec 1 Mo/s et XNUMX XNUMX IOPS. La différence entre les deux modèles Solidigm est inférieure à XNUMX %, ce qui ne montre aucun avantage de la capacité accrue à faible profondeur.
Avec une concurrence accrue (32-16), le Micron 6550 atteint 7,787.27 1.99 Mo/s et 61.44 million d'IOPS. Le Solidigm 3,799.77 To atteint 972,743 122.88 Mo/s et 3,643.64 932,770 IOPS, légèrement plus que le Solidigm 4.1 To, qui affiche 4.1 XNUMX Mo/s et XNUMX XNUMX IOPS. Cela représente une baisse de XNUMX % du débit et une baisse de XNUMX % des IOPS pour le modèle de plus grande capacité en cas de charge de lecture aléatoire maximale.
Latence de lecture aléatoire de 4 K
À (1-1), le Micron 6550 affiche la latence la plus faible, soit 0.067 ms. Les deux modèles Solidigm P5336 atteignent 0.102 ms, sans amélioration de la latence avec l'augmentation de la capacité. Sous forte charge (32-16), le Micron conserve une efficacité élevée, à 0.260 ms. Le P5336 61.44 To enregistre 0.525 ms, tandis que le modèle 122.88 To augmente légèrement à 0.546 ms, soit une augmentation de 4 % de la latence qui reflète une baisse d'efficacité très minime pour le disque de plus grande capacité dans des conditions de lecture aléatoire maximales.
Stockage direct du GPU
