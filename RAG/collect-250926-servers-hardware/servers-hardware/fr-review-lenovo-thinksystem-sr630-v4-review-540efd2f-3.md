---
id: collect-250926-servers-hardware/servers-hardware/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f-3
title: "fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "gpu", "intel", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f.md
source_anchor: ""
source_lines: [90, 129]
sha256: 695001414a71b9130393d4754c350c50bea2e6a9527e5508481424a9d91635a8
---

# fr-review-lenovo-thinksystem-sr630-v4-review-540efd2f

Tout d'abord, nous allons passer au test Blender, une application de modélisation 3D open source. Ce test a été exécuté à l'aide de l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Les tests Blender OptiX révèlent des informations intéressantes sur les systèmes testés. Le Lenovo ThinkSystem SR630 V4 a excellé avec Blender 4.2.0, en fournissant 1,432 569 échantillons par minute dans la scène Monster, tandis que l'Intel Ice Lake Server a atteint 1 et le Supermicro Hyper 112U 4.0H-TN (exécutant Blender 781) a atteint 914. Lenovo a signalé 403 échantillons dans la scène Junkshop, avec 514 pour l'Intel Ice Lake Server et 657 pour Supermicro. La scène Classroom a montré que Lenovo avait 280 échantillons, Ice Lake avec 371 et Supermicro avec XNUMX.
| Processeur de mélangeur | Supermicro Hyper 1U 112H-TN (1x Xeon 6780E, 512 Go DDR5) Mixeur 4.0 | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) Mixeur 4.2 | Serveur Intel Ice Lake (2 x Intel Xeon 8380, 512 Go)Mixeur 4.2 | 
| Monster | 781.42 | 1432.09 | 569.10 | 
| Brocanteur | 514.658 | 914.75 | 403.96 | 
| Salle de classe | 370.52 | 656.68 | 280.86 | 
Geekbench 6
Geekbench 6 est un outil d'évaluation des performances multiplateforme qui mesure les performances globales d'un système. Le navigateur Geekbench permet de comparer n'importe quel système à cet outil.
Le test mono-cœur de Lenovo a obtenu un score de 1,173 1,154, tandis que le Supermicro à double processeur a enregistré 15,167 13,868, ce qui met en évidence son potentiel pour les tâches qui dépendent de l'efficacité de chaque cœur. Dans le test multi-cœur, Supermicro a enregistré 144 288 et Lenovo 630 4. Toutes les applications n'ont pas bien résisté à l'augmentation du nombre de cœurs. Geekbench a eu du mal à passer de 1,668 à 17,409 cœurs sur la plate-forme Lenovo SRXNUMX VXNUMX à double socket. Les anciens processeurs Ice Lake ont affiché des scores mono-cœur et multi-cœur plus élevés, atteignant respectivement XNUMX XNUMX et XNUMX XNUMX.
| Geekbench 6 (Plus c'est haut, mieux c'est) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Intel Ice Lake (2 x Intel Xeon 8380, 512 Go) | 
| Processeur monocœur | 1,154 | 1,173 | 1,668 | 
| Processeur multicœur | 15,167 | 13,868 | 17,409 | 
Cinebench R23
L'outil de référence Cinebench R23 évalue les performances du processeur d'un système en restituant une scène 3D complexe à l'aide du moteur Cinema 4D. Il mesure les performances monocœur et multicœur, offrant une vue complète des capacités du processeur dans la gestion des tâches de rendu 3D.
Le tableau ci-dessous montre que les systèmes Supermicro et Lenovo ont enregistré de bons résultats. Le Lenovo ThinkSystem SR630 V4 a surpassé les performances en multi-cœur et en mono-cœur, avec respectivement 99,266 894 et 1 points. Le Supermicro Hyper 112U 92,516H-TN a enregistré 888 74,020 et 6780 points. Le processeur supplémentaire a eu quelques avantages dans ce benchmark, mais les chiffres n'ont pas doublé, passant d'un à deux processeurs. Les processeurs Ice Lake ont enregistré XNUMX XNUMX, avec une mise à l'échelle limitée par rapport aux doubles Xeon XNUMXE de Lenovo.
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Intel Ice Lake (2 x Intel Xeon 8380, 512 Go) | 
| Processeur multicœur | 92,516 | 99,266 pts | 74,020 XNUMX points | 
| Processeur monocœur | 888 pts | 894 pts | 1,059 XNUMX points | 
| Rapport PM | 104.20 x | 111.00 x | 69.87 x | 
Cinebench 2024
Cinebench 2024 étend les capacités de référence de R23 en ajoutant une évaluation des performances du GPU. Il continue de tester les performances du processeur mais inclut également des tests qui mesurent la capacité du GPU à gérer les tâches de rendu.
Les résultats de la version 2024 de Cinebench ont raconté une histoire similaire. Ici, le Lenovo ThinkSystem SR630 V4 avec ses deux processeurs 6780E a reflété l'avantage sur le Supermicro Hyper 1U 112H-TN monosocket en termes de performances CPU multicœurs avec un score de 2,884 2,565 points. Le Supermicro a rapporté 8380 4,131 points. Les processeurs Intel Ice Lake 8380 ont démontré sa puissance avec un score multicœur de 61 XNUMX. Pour les tests monocœurs, l'Intel Ice Lake XNUMX a obtenu XNUMX points.
| Cinebench R23 | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Intel Ice Lake (2 x Intel Xeon 8380, 512 Go) | 
| Processeur multicœur | 2,565 pts | 2,884 pts | 4,131 pts | 
| Processeur monocœur | 53 pts | 53 pts | 61 pts | 
| Rapport PM | 48.38 x | 54.43 x | 68.22 x | 
croque-y
y-cruncher est une application de benchmarking et de test de résistance populaire lancée en 2009. Ce test est multithread et évolutif, calculant Pi et d'autres constantes jusqu'à des milliers de milliards de chiffres. Plus vite c'est mieux dans ce test. Ce logiciel a été fantastique pour tester les plates-formes à nombre de cœurs élevé et montrer les avantages de calcul entre les plates-formes à un ou deux sockets.
Dans les tests de performance de y-cruncher, le ThinkSystem SR360 V4 à double socket a mis 5.997 secondes pour calculer Pi à 1 milliard de chiffres. Le processeur Intel Xeon 6780E a mis 8.757 secondes. Pour calculer 50 milliards de chiffres, un seul processeur a eu besoin de 674.299 secondes, contre 476.826 secondes pour les processeurs doubles. Bien que toutes les charges de travail ne répondent pas bien au nombre élevé de cœurs des nouveaux processeurs à e-core, y-cruncher n'a eu aucun mal à les exploiter. L'ancien Intel Ice Lake Server a terminé le calcul de Pi à 1 milliard de chiffres en 7.074 secondes lors du premier test à 1 milliard de chiffres, tout en atteignant 617.828 secondes en 50 milliards de chiffres.
| y-cruncher (0.8.5.9) (plus bas est mieux) | Supermicro Hyper 1U 112H-TN (Xeon 6780E, 512 Go DDR5) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E, 512 Go) | Serveur Intel Ice Lake (2 x Intel Xeon 8380, 512 Go) | 
| 1 milliard | 8.757 secondes | 5.997 secondes | 7.074 secondes | 
| 2.5 milliard | 24.928 secondes | 17.573 secondes | 19.203 secondes | 
| 5 milliard | 53.489 secondes | 37.793 secondes | 42.300 secondes | 
| 10 milliard | 113.727 secondes | 81.046 secondes | 93.886 secondes | 
| 25 milliard | 308.218 secondes | 220.025 secondes | 272.679 secondes | 
| 50 milliard | 674.299 secondes | 476.826 secondes | 617.828 secondes | 
Test de vitesse Blackmagic RAW
Le Blackmagic RAW Speed Test est un outil d'analyse comparative des performances conçu pour mesurer les capacités d'un système à gérer la lecture et l'édition vidéo à l'aide du codec Blackmagic RAW. Il évalue la capacité d'un système à décoder et à lire des fichiers vidéo haute résolution, en fournissant des fréquences d'images pour le traitement basé sur le CPU et le GPU.
Le Lenovo ThinkSystem SR630 V4 a obtenu des résultats légèrement supérieurs à ceux des systèmes Supermicro et Ice Lake, avec un score de 120 FPS en CPU 8K, ce qui en fait un excellent choix pour les tâches de lecture et d'édition vidéo. Le Supermicro Hyper 1U 112H-TN a obtenu 116 FPS en CPU 8K. Bien que le Lenovo à double socket ait obtenu de meilleurs résultats, ce n'était pas énorme, compte tenu de la configuration à double socket avec 116 FPS (CPU 8K) et 0 FPS (GPU 8K). Le serveur Intel Ice Lake était au même niveau que le Lenovo et Supermicro avec 116 FPS dans le benchmark CPU 8K. Cependant, il n'a pas enregistré de score GPU en raison de l'absence d'un GPU dédié.
7-Zip
