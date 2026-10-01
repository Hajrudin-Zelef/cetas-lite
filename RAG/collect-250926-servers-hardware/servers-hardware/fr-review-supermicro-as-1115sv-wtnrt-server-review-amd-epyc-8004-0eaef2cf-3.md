---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf-3
title: "fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "gpu", "inference", "valuation"]
source: docs/RAG/clean4/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf.md
source_anchor: ""
source_lines: [122, 179]
sha256: 08f3072146f0bae2d705b377860929a692405664419e7aa749d57550860a51cb
---

# fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf

| Processeur (monocœur) (Points) | 69 | 81 | 
| GPU | 3,634 | N/D | 
| Rapport PM | 57.27 | 60.47x | 
Évaluation du processeur Geekbench
Geekbench 6 est un outil d'évaluation des performances multiplateforme qui mesure les performances globales d'un système. Toutefois, il serait intéressant de comparer les performances monocœur et multicœur, ainsi que les performances OpenCL. Un score élevé indique de meilleures performances.
Ici, le système a démontré des performances globales solides en termes de mesures CPU et GPU. Le benchmark des processeurs monocœurs a enregistré un score de 1,718 19,455, tandis que le test des processeurs multicœurs a montré 35,764 XNUMX. De plus, les performances du GPU ont été évaluées à l'aide du framework OpenCL, avec un score de XNUMX XNUMX, indiquant une solide puissance de traitement graphique adaptée à diverses applications gourmandes en calcul.
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
| Geekbench 6 | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| CPU Benchmark - Monocœur | 1,718 | 2,055 | 
| Référence CPU - Multi-Core | 19,455 | 22,868 | 
| Référence GPU – OpenCL | 35,764 | N/D | 
croque-y
y-cruncher est un programme multi-thread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des billions de chiffres. Depuis son lancement en 2009, il est devenu une application populaire d'analyse comparative et de test de résistance pour les overclockeurs et les passionnés de matériel.
En termes de résultats, nous avons des résultats allant de 1 milliard à 25 milliards pour le Supermicro 1115SV-WTNRT.
| y-cruncher (Temps de calcul total) (plus le niveau est bas, mieux c'est) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| 1 milliard de chiffres (secondes) | 9.989 secondes | 7.274 secondes | 
| 2.5 milliard de chiffres (secondes) | 24.974 secondes | 17.055 secondes | 
| 5 milliard de chiffres (secondes) | 52.117 secondes | 34.336 secondes | 
| 10 milliard de chiffres (secondes) | 110.483 secondes | 71.336 secondes | 
| 25 milliard de chiffres (secondes) | 303.372 secondes | 196.695 secondes | 
| 50 milliard de chiffres (secondes) | N/D | 439.435 secondes | 
Inférence UL Procyon AI
UL Procyon AI Inference est conçu pour évaluer les performances d'une station de travail dans des applications professionnelles. Il est important de noter que ce test n'exploite pas les capacités de plusieurs processeurs. Plus précisément, cet outil évalue la capacité de la station de travail à gérer les tâches et les flux de travail basés sur l'IA, fournissant une analyse détaillée de son efficacité et de sa rapidité de traitement des algorithmes et applications d'IA complexes.
Les résultats couvrent une gamme de modèles d'IA, indiquant la polyvalence du serveur dans la gestion de différents types de charges de travail d'IA.
Les temps d'inférence vont de très rapides (4.11 ms pour MobileNet V3) à beaucoup plus lents (2,131.81 130 ms pour Real-ESRGAN), ce qui indique une grande variation dans la façon dont le système gère les différentes tâches d'IA. Compte tenu de la complexité et de la variété des fonctions, un score de XNUMX suggère que même si le système fonctionne de manière adéquate dans un large éventail d'opérations basées sur l'IA, il pourrait y avoir des limites lors de la gestion de tâches plus gourmandes en ressources comme Real-ESRGAN.
| Temps d'inférence moyens UL Procyon (le plus faible est le mieux) | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | 
| Mobile Net V3 | 4.11 ms | 
| ResNet 50 | 8.40 ms | 
| Création V4 | 30.27 ms | 
| Deep Lab V3 | 30.87 ms | 
| YOLO V3 | 45.66 ms | 
| Réel-ESRGAN | 2,131.81 ms | 
| Note globale | 130 | 
Compression à 7 zips
L'utilitaire populaire 7-Zip dispose d'un test de mémoire intégré qui démontre les performances du processeur. Dans ce test, nous l'exécutons avec une taille de dictionnaire de 128 Mo lorsque cela est possible. Encore une fois, nous n'avons que les résultats de l'ASG-1115S-NE316R pour ce test.
|  | Supermicro 1115SV-WTNRT (64C EPYC 8534P, 194 Go DDR5) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
| Compression |  |  | 
| Utilisation actuelle du processeur | 5,596 % | 3,574 % | 
| Courant nominal/utilisation | 4.387 | 5.755 GIPS | 
| Courant | 245.508 | 205.670 GIPS | 
| Utilisation résultante du processeur | 5,615 % | 3,572 % | 
| Évaluation/utilisation résultante | 4.383 | 5.733 GIPS | 
| Note résultante | 246.126 | 204.758 GIPS | 
| Décompression |  |  | 
| Utilisation actuelle du processeur | 6,198 % | 3,810 % | 
| Courant nominal/utilisation | 4.224 GIPS | 5.753 GIPS | 
| Courant | 261.832 GIPS | 219.191 GIPS | 
| Utilisation résultante du processeur | 6,057 % | 3,803 % | 
| Évaluation/utilisation résultante | 4.382 GIPS | 5.779 GIPS | 
| Note résultante | 265.381 GIPS | 219.768 GIPS | 
| Note totale |  |  | 
| Utilisation totale du processeur | 5,836 % | 3,687 % | 
| Note totale/utilisation | 4.383 GIPS | 5.756 GIPS | 
| Note totale | 255.753 GIPS | 219.768 GIPS | 
Conclusion
Le Supermicro Server AS-1115SV-WTNRT se distingue comme un serveur robuste monté en rack 1U, capable de gérer des applications à forte demande, notamment la virtualisation, la gestion de bases de données et l'informatique de pointe. Équipé d'un processeur AMD EPYC 8004 « Siena », l'AS-1115SV-WTNRT apparaît comme une option attrayante pour les fournisseurs de services soucieux d'améliorer l'efficacité et les performances des centres de données tout en maîtrisant les coûts.
La série AMD EPYC 8004 excelle dans les plates-formes à socket unique, offrant de nombreux cœurs à un coût inférieur à celui des modèles à cœurs plus élevés, ainsi qu'une faible consommation d'énergie (à partir de 70 watts seulement). Cela les rend particulièrement adaptés aux centres de données denses où l’espace et la puissance sont limités. De plus, ces processeurs prennent en charge six canaux de mémoire DDR5 dans un encombrement compact, permettant des configurations grand public qui maintiennent à la fois vitesse et efficacité.
L'intégration de ces processeurs dans des plates-formes telles que le Supermicro Server AS-1115SV-WTNRT s'aligne sur les tendances actuelles en matière d'exploitation rentable des centres de données tout en évoluant efficacement pour répondre aux demandes futures. Les fournisseurs de services qui exploitent ces processeurs connaîtront certainement des améliorations notables en termes d’efficacité opérationnelle et une réduction du coût total de possession (TCO). Cela dit, l'AS-1115SV-WTNRT offre une solution efficace et évolutive qui répond parfaitement aux défis informatiques modernes, en offrant flexibilité et redondance dans les options de stockage.
