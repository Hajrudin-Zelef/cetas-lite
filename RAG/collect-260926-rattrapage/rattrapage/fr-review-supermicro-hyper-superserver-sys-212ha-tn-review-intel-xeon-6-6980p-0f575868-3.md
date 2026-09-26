---
id: collect-260926-rattrapage/rattrapage/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868-3
title: "fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "gpu", "valuation"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868.md
source_anchor: ""
source_lines: [84, 128]
sha256: ce1ae90b7a5f4595e3af311da90589579c94fcd4554177e2a6cd26ca6ac50357
---

# fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868

y-cruncher est une application de benchmarking et de test de résistance populaire lancée en 2009. Ce test est multithread et évolutif, calculant Pi et d'autres constantes jusqu'à des milliers de milliards de chiffres. Plus vite c'est mieux dans ce test. Ce logiciel a été fantastique pour tester les plates-formes à nombre de cœurs élevé et montrer les avantages de calcul entre les plates-formes à un ou deux sockets.
Le test Y-Cruncher met en évidence les différences de performances marquées entre les processeurs Granite Rapids et Sierra Forest, en soulignant notamment le rôle de l'architecture du processeur et de la conception du cœur. Le Supermicro SYS-212HA-TN, équipé du processeur Intel Xeon 128P à 6980 cœurs (basé sur P-Core), a constamment surpassé ses concurrents, excellant dans les calculs à haut débit. Le Lenovo ThinkSystem SR630 V4, équipé de deux processeurs Intel Xeon 144E à 6780 cœurs (basés sur E-Core), a fourni de bons résultats mais est resté en retrait par rapport au SYS-212HA-TN. Le Supermicro SYS-112H-TNRT, doté d'un seul processeur Xeon 6780E, a été désavantagé en raison de sa configuration à un seul processeur.
| y-cruncher (0.8.5.9) (plus bas est mieux) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Hyper Super SYS-212HA-TN (1x Intel Xeon 6980P, 384 Go de RAM) | 
|---|---|---|---|
| 1 milliard | 8.757 secondes | 5.997 secondes | 2.965 secondes | 
| 2.5 milliard | 24.928 secondes | 17.573 secondes | 8.081 secondes | 
| 5 milliard | 53.489 secondes | 37.793 secondes | 17.650 secondes | 
| 10 milliard | 113.727 secondes | 81.046 secondes | 38.170 secondes | 
| 25 milliard | 308.218 secondes | 220.025 secondes | 108.123 secondes | 
| 50 milliard | 674.299 secondes | 476.826 secondes | 238.4429 secondes | 
Geekbench 6
Geekbench 6 est un outil d'évaluation des performances multiplateforme qui mesure les performances globales d'un système. Le navigateur Geekbench permet de comparer n'importe quel système à cet outil.
Le processeur mono-cœur Lenovo a obtenu un score de 1,173 112, tandis que le Supermicro SYS-1,154H-TNRT mono-processeur a obtenu un score de 6780 212, ce qui démontre l'efficacité du Xeon 6980E dans les charges de travail mono-thread. Le Supermicro SYS-2,059HA-TN, équipé du Xeon XNUMXP, a quant à lui obtenu un score de XNUMX XNUMX, exploitant ses P-Cores pour des performances mono-thread exceptionnelles.
Lors du test multi-cœur, le SYS-112H-TNRT a obtenu un score de 15,167 212, surpassant de peu le SYS-15,055HA-TN à 630 4, tandis que le Lenovo ThinkSystem SR13,868 V630 à double socket a obtenu un score de 4 288. Geekbench a eu du mal à évoluer pleinement avec les XNUMX E-Cores du SRXNUMX VXNUMX. Ce goulot d'étranglement met en évidence les limites de certaines applications lorsqu'elles sont confrontées à un nombre élevé de cœurs.
| Geekbench 6 (Plus c'est mieux) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Hyper Super SYS-212HA-TN (1x Intel Xeon 6980P, 384 Go de RAM) | 
|---|---|---|---|
| Processeur monocœur | 1,154 | 1,173 | 2,059 | 
| Processeur multicœur | 15,167 | 13,868 | 15,055 | 
Cinebench R23
L'outil de référence Cinebench R23 évalue les performances du processeur d'un système en restituant une scène 3D complexe à l'aide du moteur Cinema 4D. Il mesure les performances monocœur et multicœur, offrant une vue complète des capacités du processeur dans la gestion des tâches de rendu 3D.
Le tableau ci-dessous met en évidence les résultats du test Cinebench R23. Le Lenovo ThinkSystem SR630 V4 a excellé en performances multi-cœurs et mono-cœurs, obtenant respectivement 99,266 894 et 112 points. Le Supermicro SYS-92,516H-TNRT a obtenu des résultats admirables avec 888 6780 points multi-cœurs et 212 points mono-cœurs, démontrant une forte efficacité avec un seul Xeon 128E. Le Supermicro SYS-6980HA-TN, alimenté par le Xeon 76,617P à 1,479 cœurs, a obtenu XNUMX XNUMX points multi-cœurs et XNUMX XNUMX points mono-cœurs, démontrant ainsi ses performances exceptionnelles en mode monothread.
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Hyper Super SYS-212HA-TN (1x Intel Xeon 6980P, 384 Go de RAM) | 
|---|---|---|---|
| Processeur multicœur | 92,516 pts | 99,266 pts | 76,617 pts | 
| Processeur monocœur | 888 pts | 894 pts | 1,479 pts | 
| Rapport PM | 104.20 x | 111.00 x | 51.79 x | 
Cinebench 2024
Cinebench 2024 étend les capacités de référence de R23 en ajoutant une évaluation des performances du GPU. Il continue de tester les performances du processeur mais inclut également des tests qui mesurent la capacité du GPU à gérer les tâches de rendu.
Le tableau ci-dessous met en évidence les résultats du test Cinebench 2024. Le Lenovo ThinkSystem SR630 V4, équipé de deux processeurs Xeon 6780E, a obtenu 2,884 112 points au test multicœur, tandis que le Supermicro SYS-6780H-TNRT, également équipé d'un seul Xeon 2,565E, a enregistré 53 XNUMX points. Les deux systèmes ont obtenu XNUMX points au test monocœur, ce qui reflète des performances par cœur similaires.
Le Supermicro SYS-212HA-TN, équipé d'un seul processeur Xeon 6980P avec un nombre élevé de P-Cores de 128, a obtenu 6,078 87 points lors du test multicœur. Ce système a également excellé en performances monocœur, obtenant 6980 points, bénéficiant de l'architecture améliorée et des P-Cores haute fréquence du 69.69P. Son évolutivité des performances multicœurs a donné lieu à un ratio MP de 54.43x, surpassant largement le ratio de 48.38x du Lenovo et de 112x du Supermicro SYS-XNUMXH-TNRT.
Une fois de plus, cela met en évidence l'avantage du nombre élevé de P-Core du Xeon 6980P, permettant au Supermicro SYS-212HA-TN de surpasser la configuration à double processeur du Lenovo Think System SR630 V4 et le Supermicro SYS-112H-TNRT à processeur unique dans les tests de performance multicœurs et monocœurs.
| Cinebench R24 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Hyper Super SYS-212HA-TN (1x Intel Xeon 6980P, 384 Go de RAM) | 
|---|---|---|---|
| Processeur multicœur | 2,565 pts | 2,884 pts | 6,078 pts | 
| Processeur monocœur | 53 pts | 53 pts | 87 pts | 
| Rapport PM | 48.38 x | 54.43 x | 69.69 x | 
Test de vitesse Blackmagic RAW
Le Blackmagic RAW Speed Test est un outil d'analyse comparative des performances conçu pour mesurer les capacités d'un système à gérer la lecture et l'édition vidéo à l'aide du codec Blackmagic RAW. Il évalue la capacité d'un système à décoder et à lire des fichiers vidéo haute résolution, en fournissant des fréquences d'images pour le traitement basé sur le CPU et le GPU.
Le Supermicro SYS-212HA-TN a surpassé le Think System SR630 V4 et le Supermicro SYS-112H-TNRT avec un score total de 148 FPS.
| Test de vitesse Blackmagic RAW (plus c'est élevé, mieux c'est) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Hyper Super SYS-212HA-TN (1x Intel Xeon 6980P, 384 Go de RAM) | 
|---|---|---|---|
| FPS CPU | FPS 116 | FPS 120 | FPS 148 | 
7-Zip
L'outil de référence de mémoire intégré de l'utilitaire populaire 7-Zip mesure les performances du processeur et de la mémoire d'un système pendant les tâches de compression et de décompression, indiquant dans quelle mesure le système peut gérer les opérations gourmandes en données.
Résultats clés (plus c'est élevé, mieux c'est)
