---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-c6615-server-review-e5a753e4-2
title: "fr-review-dell-poweredge-c6615-server-review-e5a753e4"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["benchmark", "gpu", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-c6615-server-review-e5a753e4.md
source_anchor: ""
source_lines: [44, 120]
sha256: d77f6baf1d96e366cd79347c80b12d9723901085ea44e0c9af19beb51ac891e9
---

# fr-review-dell-poweredge-c6615-server-review-e5a753e4

Bien qu'une grande partie de cet examen se concentre sur les performances globales au niveau du système, nous avons légèrement abordé les deux types de stockage sur ce système avec des charges de travail aux quatre coins. Notre premier test s'est concentré sur le groupe SSD de démarrage BOSS RAID1.
| Dell BOSS RAID1 | Lire les performances | Performances d'écriture | 
|---|---|---|
| Séquentiel 1 Mo Q32/4T | 2,963MB / s | 1,067MB / s | 
| Aléatoire 4K Q32/8T | 600,786 0.426 IOPS (XNUMX ms) | 249,819 1.024 IOPS (XNUMX ms) | 
Ensuite, nous avons examiné un seul SSD Gen5 E3.S, qui comprenait le SSD KIOXIA CM7 de 7.68 To à lecture intensive dans notre système de test.
| KIOXIA 7.68 To CM7-R | Lire les performances | Performances d'écriture | 
|---|---|---|
| Séquentiel 1 Mo Q32/4T | 13,736MB / s | 7,089MB / s | 
| Aléatoire 4K Q32/8T | 931,671 0.266 IOPS (XNUMX ms) | 768,739 0.329 IOPS (XNUMX ms) | 
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs. Voici les résultats pour toutes les puces EPYC.
Dans Cinebench R23, les quatre nœuds se situaient autour de 74,000 3 sur la partie multicœur, le nœud 75,000 se glissant dans 1 4. Les quatre nœuds sont restés beaucoup plus proches pour les scores monocœur, les nœuds 1,088 et 3 étant de 8 2. Le nœud 5 n’avait que XNUMX points de retard et le nœud XNUMX avait XNUMX points d’avance. Dans l’ensemble, tous les nœuds ne présentaient que des écarts de performances mineurs, typiques des différents processeurs, même s’ils appartiennent tous au même modèle.
| Cinebench R23 | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Processeur multicœur | 74,877 | 74,961 | 75,011 | 74,745 | 74,898.5 | 
| Processeur monocœur | 1,088 | 1,093 | 1,084 | 1,088 | 1,088.25 | 
| Rapport PM | 64.84 | 68.60 | 69.17 | 68.70 | 67.83 | 
Cinebench 2024
Cinebench 2024 de Maxon est une référence de rendu CPU et GPU qui utilise tous les cœurs et threads du processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Étant donné que ces nœuds n'ont pas de GPU, nous n'avons que les numéros multicœurs et monocœurs.
Dans Cinebench 2024, tous les nœuds sont restés proches les uns des autres, avec une variance minime sur les parties multicœur et monocœur. Les performances moyennes étaient de 4,509 67.25 points pour le multicœur et de 66.98 points pour le monocœur, avec un ratio MP de XNUMX.
| Cinebench 2024 | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Processeur multicœur | 4,544 | 4,577 | 4,436 | 4,481 | 4,509.5 | 
| Processeur monocœur | 68 | 68 | 65 | 68 | 67.25 | 
| Rapport PM | 66.79 | 67.23 | 68.21 | 65.69 | 66.98 | 
Processeur Geekbench 6
Geekbench 6 est un benchmark multiplateforme qui mesure les performances globales d'un système. Ce test comprend une partie dédiée au processeur et une autre au processeur graphique, mais comme ces nœuds ne possèdent pas de GPU, nous ne disposons que des résultats du processeur. Plus le score est élevé, meilleures sont les performances.
Dans Geekbench, nous avons vu des chiffres serrés jusqu'à ce que nous arrivions au nœud 3, qui s'est légèrement replié sur les monocœurs et les multicœurs. La moyenne entre tous les nœuds était de 1,687 19,319.5 en monocœur et de XNUMX XNUMX en multicœur.
| Processeur Geekbench 6 | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Single-Core | 1,707 | 1,708 | 1,625 | 1,708 | 1,687 | 
| Multi-Core | 19,544 | 19,234 | 18,999 | 19,501 | 19,319.5 | 
Processeur Blender 4.0
La prochaine étape est Blender OptiX, une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire CLI Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Les nœuds C6615 ont enregistré des chiffres assez cohérents. Les scores moyens étaient de 591.79 sur Monster, 415.88 sur Junkshop et 311.74 sur Classroom.
| Processeur Blender 4.0 | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Monster | 595.23 | 593.51 | 584.35 | 594.07 | 591.79 | 
| Brocanteur | 415.26 | 415.11 | 418.05 | 415.08 | 415.88 | 
| Salle de classe | 308.57 | 312.91 | 312.69 | 312.78 | 311.74 | 
Processeur Blender 4.1
Blender OptiX 4.1 apporte de nouvelles fonctionnalités, telles que le débruitage accéléré par GPU, la rationalisation du processus de rendu et la réduction du temps nécessaire aux tâches de débruitage. Malgré ces progrès, les améliorations globales des performances dans les scores de référence par rapport à la version 4.0 sont minimes, indiquant seulement de légères améliorations en termes d'efficacité.
Encore une fois, nous constatons des chiffres cohérents dans tous les domaines, avec des moyennes de 587.22 sur Monster, 420.20 sur Junkshop et 306.60 sur Classroom.
| Processeur Blender 4.1 | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Monster | 590.46 | 590.58 | 584.76 | 583.08 | 587.22 | 
| Brocanteur | 418.38 | 416.71 | 426.73 | 419.03 | 420.20 | 
| Salle de classe | 306.86 | 304.81 | 308.95 | 305.79 | 306.60 | 
Compression à 7 zips
L'utilitaire populaire 7-Zip dispose d'un test de mémoire intégré qui démontre les performances du processeur. Dans ce test, nous l'exécutons sur une taille de dictionnaire de 128 Mo lorsque cela est possible.
Des scores équitables ont été observés dans tous les nœuds. Dans les scores totaux, nous avons constaté une utilisation totale du processeur de 5,778.75 4.355 %, une note/utilisation totale de 252 GIPS et une note totale de XNUMX GIPS.
| Processeur Blender 4.1 | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
| Compression |  |  |  |  |  | 
| Utilisation actuelle du processeur | 5,548 % | 5,549 % | 5,633 % | 5,585 % | 5,578.75 % | 
| Note actuelle/utilisation | 4.256 GIPS | 4.210 GIPS | 4.156 GIPS | 4.177 GIPS | 4.20 GIPS | 
| Courant | 236.158 GIPS | 233.626 GIPS | 234.092 GIPS | 233.285 GIPS | 234.290 GIPS | 
| Utilisation résultante du processeur | 5,536 % | 5,537 % | 5,601 % | 5,553 % | 5,556.75 % | 
| Évaluation/utilisation résultante | 4.193 GIPS | 4.202 GIPS | 4.172 GIPS | 4.168 GIPS | 4.184 GIPS | 
| Note résultante | 232.118 GIPS | 232.631 GIPS | 233.691 GIPS | 231.443 GIPS | 232.470 GIPS | 
| Décompression |  |  |  |  |  | 
| Utilisation actuelle du processeur | 5,973 % | 6,027 % | 5,992 % | 6,014 % | 6,001.5 % | 
| Note actuelle/utilisation | 4.543 GIPS | 4.501 GIPS | 4.565 GIPS | 4.509 GIPS | 4.530 GIPS | 
| Courant | 271.343 GIPS | 271.287 GIPS | 273.507 GIPS | 271.196 GIPS | 271.833 GIPS | 
| Utilisation résultante du processeur | 5,997 % | 6,015 % | 5,999 % | 5,990 % | 6,000.25 % | 
| Évaluation/utilisation résultante | 4.537 GIPS | 4.519 GIPS | 4.550 GIPS | 4.499 GIPS | 4.526 GIPS | 
| Note résultante | 272.066 GIPS | 271.775 GIPS | 272.946 GIPS | 269.509 GIPS | 271.574 GIPS | 
| Note totale |  |  |  |  |  | 
| Utilisation totale du processeur | 5,767 % | 5,776 % | 5,800 % | 5,772 % | 5,778.75 % | 
| Note totale/utilisation | 4.365 GIPS | 4.360 GIPS | 4.361 GIPS | 4.333 GIPS | 4.355 GIPS | 
| Note totale | 252.092 GIPS | 252.203 GIPS | 253.318 GIPS | 250.476 GIPS | 252.022 GIPS | 
Test de vitesse brute Blackmagic
Nous utilisons le Raw Speed Test de Blackmagic pour évaluer la façon dont les machines effectuent le décodage RAW réel. Ce test peut intégrer à la fois l’utilisation du CPU et du GPU, mais nous testerons uniquement l’utilisation du CPU.
Les quatre nœuds ont présenté des performances extrêmement proches, avec une moyenne de 119.75 FPS.
| Test de vitesse brute Blackmagic | Nœud 1 | Nœud 2 | Nœud 3 | Nœud 4 | Normale | 
|---|---|---|---|---|---|
