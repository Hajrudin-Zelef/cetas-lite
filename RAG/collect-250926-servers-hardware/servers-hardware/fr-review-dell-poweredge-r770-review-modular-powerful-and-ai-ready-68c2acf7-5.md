---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-5
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "intel", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [107, 163]
sha256: bc4e010b6fa419292fbe569aa5db2348b1b4cb9c529e816420fdbac9c70e7d60
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

Les résultats du benchmark Procyon AI Computer Vision démontrent également d'excellentes performances d'inférence IA. Le système a obtenu de faibles temps d'inférence, avec MobileNet V3 à 20.64 ms et ResNet 50 à 22.42 ms. Inception V4 et DeepLab ont fonctionné respectivement à 65.23 ms et 41.37 ms, gérant efficacement des charges de travail de vision plus complexes. YOLO V3, un modèle clé de détection d'objets, a traité en 37.80 ms, ce qui le rend particulièrement adapté aux applications d'IA en temps réel. REAL-ESRGAN, un modèle de super-résolution à forte intensité de calcul, a enregistré 1,159.22 81 ms, ce qui nous a valu un score global de XNUMX en AI Computer Vision.
| Vision par ordinateur IA (une durée plus courte est meilleure) (un score plus élevé est meilleur) | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | 
|---|---|
| Temps d'inférence moyen MobileNet V3 | 20.64 ms | 
| Temps d'inférence moyen ResNet 50 | 22.42 ms | 
| Temps d'inférence moyen Inception V4 | 65.23 ms | 
| Temps d'inférence moyen de DeepLab | 41.37 ms | 
| Temps d'inférence moyen YOLO V3 | 37.80 ms | 
| Temps d'inférence moyen REAL-ESRGAN | 1,159.22 ms | 
| Score global de vision par ordinateur IA | 81 | 
Hammer DB TPROC-C
Nous avons également évalué les performances de quatre bases de données open source populaires (MariaDB 11.4.4, MySQL 8.4.4, MySQL 5.7.44 et PostgreSQL 17.2) à l'aide du benchmark HammerDB TPROC-C pour simuler les charges de travail OLTP sur 500 entrepôts.
MariaDB s'est imposée comme la solution la plus performante, notamment dans les configurations à double socket, où elle a évolué efficacement et atteint le débit de transactions le plus élevé. MySQL 8.4.4 a enregistré des améliorations notables par rapport à l'ancienne version 5.7.44, mettant en évidence les améliorations des versions récentes. PostgreSQL 17.2 a fourni des performances constantes, mais a affiché un léger retard par rapport à MariaDB et MySQL 8.4.4. MariaDB a fourni 3.15 millions de TPM sur un seul socket et 5.8 millions de TPM sur deux sockets, surpassant les autres dans les deux scénarios.
Tableau de comparaison des performances (Transactions par minute, TPM)
| Moteur de base de données | TPM à socket unique | TPM à double socket | 
|---|---|---|
| MariaDB 11.4.4 | 3,150,000 | 5,800,000 | 
| MySQL 8.4.4 | 2,850,000 | 5,150,000 | 
| PostgreSQL 17.2 | 2,700,000 | 4,900,000 | 
| MySQL 5.7.44 | 2,300,000 | 4,250,000 | 
Malgré la puissance matérielle du R770, avec ses 86 cœurs par processeur (un mélange de cœurs à haute et basse priorité), aucune base de données n'a enregistré de gains de performances significatifs lorsqu'elle était répartie sur les deux sockets. Cela reflète la préférence générale des bases de données open source pour l'exécution sur un seul socket, en raison d'une meilleure localisation du cœur et d'une latence mémoire réduite.
Compte tenu de ces résultats, le R770 est plus adapté à l'exécution de plusieurs instances de bases de données dans un environnement virtualisé qu'à la mise à l'échelle d'une seule instance. L'architecture du système est idéale pour prendre en charge une charge de travail de bases de données mixtes à haute densité, exploitant à la fois les cœurs de performance et d'efficacité pour assurer un débit constant sur plusieurs instances.
7-Zip
L'outil de référence de mémoire intégré de l'utilitaire populaire 7-Zip mesure les performances du processeur et de la mémoire d'un système pendant les tâches de compression et de décompression, indiquant dans quelle mesure le système peut gérer les opérations gourmandes en données.
Lors du benchmark 7-Zip, le système Dell a obtenu une note supérieure (266.425 GIPS) à celle de Lenovo (224.313 GIPS) pour les tâches de compression, le système Dell affichant une utilisation du processeur légèrement inférieure. Cependant, Lenovo a surpassé Dell en décompression, avec une note supérieure (288.457 GIPS contre 256.154 GIPS) et une utilisation du processeur légèrement supérieure. Dell a obtenu une note globale légèrement supérieure (261.290 GIPS), démontrant une meilleure efficacité globale pour les tâches de compression et de décompression.
| Compression 7-Zip et Décompression | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 Go de RAM) | 
|---|---|---|
| Compression – Utilisation actuelle du processeur | 5267 % | 5064 % | 
| Compression – Courant nominal/utilisation | 5.061 GIPS | 4.341 GIPS | 
| Compression – Courant nominal | 266.591 GIPS | 219.840 GIPS | 
| Compression – Utilisation du processeur résultante | 5270 % | 5156 % | 
| Compression – Évaluation/utilisation résultante | 5.056 GIPS | 4.350 GIPS | 
| Compression – Évaluation résultante | 266.425 GIPS | 224.313 GIPS | 
| Décompression – Utilisation actuelle du processeur | 5623 % | 6184 % | 
| Décompression – Évaluation/utilisation actuelle | 4.586 GIPS | 4.688 GIPS | 
| Décompression – Courant nominal | 257.909 GIPS | 289.879 GIPS | 
| Décompression – Utilisation du processeur résultante | 5627 % | 6205 % | 
| Décompression – Évaluation/utilisation résultante | 4.553 GIPS | 4.649 GIPS | 
| Décompression – Évaluation résultante | 256.154 GIPS | 288.457 GIPS | 
| Total – Utilisation totale du processeur | 5448 % | 5681 % | 
| Total – Note totale/Utilisation | 4.804 GIPS | 4.500 GIPS | 
| Total – Note totale | 261.290 GIPS | 256.385 GIPS | 
croque-y
y-cruncher est une application de benchmarking et de test de résistance populaire lancée en 2009. Ce test est multithread et évolutif, calculant Pi et d'autres constantes jusqu'à des milliers de milliards de chiffres. Plus vite c'est mieux dans ce test. Ce logiciel a été fantastique pour tester les plates-formes à nombre de cœurs élevé et montrer les avantages de calcul entre les plates-formes à un ou deux sockets.
Les résultats du benchmark Y-Cruncher montrent un écart de performances significatif entre le Dell PowerEdge R770, équipé de processeurs P-core, et le Lenovo ThinkSystem SR630 V4 équipé de processeurs E-core, notamment à mesure que la taille du jeu de données augmente. Il s'agit moins de déterminer quel système est le meilleur que de comparer les types de processeurs sous cette charge de travail.
Pour les calculs de moindre envergure, le système Dell était déjà en tête, calculant 1 milliard de chiffres de Pi en 2.753 secondes, tandis que le système Lenovo prenait plus du double, soit 5.997 secondes. À mesure que la charge de travail augmentait, l'écart se creusait. À 10 milliards de chiffres, le système Dell a terminé en 34.873 secondes, soit moins de la moitié du temps de 81.046 secondes du système Lenovo. Au-delà des 50 milliards de chiffres, Dell a conservé son avance, réalisant la tâche en 221.255 secondes, contre 476.826 secondes pour Lenovo, soit une vitesse de 53 % pour Dell.
À 100 milliards de chiffres, Lenovo n'a pas pu terminer le test en raison de sa configuration actuelle de 512 Go de RAM. Avec 2 To de RAM, Dell a géré la charge de travail efficacement, terminant en 491.737 secondes.
| Y-cruncher (une durée plus courte est meilleure) | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 Go de RAM) | 
|---|---|---|
| 1 milliard | 2.753 secondes | 5.997 secondes | 
| 2.5 milliard | 7.365 secondes | 17.573 secondes | 
| 5 milliard | 16.223 secondes | 37.793 secondes | 
| 10 milliard | 34.873 secondes | 81.046 secondes | 
| 25 milliard | 99.324 secondes | 220.025 secondes | 
| 50 milliard | 221.255 secondes | 476.826 secondes | 
| 100 milliard | 491.737 secondes |  | 
Mixeur OptiX
