---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-6
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "gpu", "intel", "nvidia", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [164, 212]
sha256: d6a822649ccdeba743bce460f895f4474d8491e23196424169b03be89fec841f
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

Une application de modélisation 3D open source. Ce benchmark a été réalisé avec l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Les résultats du benchmark Blender montrent un net avantage en termes de performances pour le Dell PowerEdge R770 par rapport au Lenovo ThinkSystem SR630 V4, notamment en termes de rendu CPU. Dans le test « CPU Monster », Dell a atteint 1,706.002 19 échantillons par minute, soit une avance de 1,432.09 % sur les 1,169.370 914.75 échantillons par minute de Lenovo. Le test « CPU Junkshop » a encore accentué cet écart, le Dell atteignant 28 791.475 échantillons par minute, surpassant de 656.68 % les 20 échantillons par minute de Lenovo. De même, Dell a enregistré XNUMX échantillons par minute dans le test « CPU Classroom », tandis que Lenovo était à la traîne avec XNUMX échantillons par minute, soit une différence de XNUMX %.
L'absence de GPU dans le système Lenovo signifiait également qu'il ne pouvait pas participer au rendu basé sur le GPU, où le NVIDIA L4 de Dell affichait un score de 1,895.71 950.42 échantillons/min pour Monster, 968.43 échantillons/min et un score Classroom de XNUMX échantillons/min.
| Analyse comparative du CPU Blender | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 Go de RAM) | 
|---|---|---|
| Monstre CPU (Blender 4.3) | 1,706.002 XNUMX échantillons/min | 1432.09 XNUMX échantillons/min | 
| Magasin de récupération de processeurs (Blender 4.3) | 1,169.370 XNUMX échantillons/min | 914.75 XNUMX échantillons/min | 
| Classe CPU (Blender 4.3) | 791.475 XNUMX échantillons/min | 656.68 XNUMX échantillons/min | 
| Monstre GPU (Blender 4.3) | 1,895.712 XNUMX échantillons/min | (pas de GPU) | 
| Magasin de GPU (Blender 4.3) | 950.424 XNUMX échantillons/min | (pas de GPU) | 
| Classe GPU (Blender 4.3) | 968.432 XNUMX échantillons/min | (pas de GPU) | 
Cinebench R23
L'outil de référence Cinebench R23 évalue les performances du processeur d'un système en restituant une scène 3D complexe à l'aide du moteur Cinema 4D. Il mesure les performances monocœur et multicœur, offrant une vue complète des capacités du processeur dans la gestion des tâches de rendu 3D.
Dans Cinebench R23, les résultats du benchmark mettent en évidence des différences notables de performances CPU entre le Dell PowerEdge R770 et le Lenovo ThinkSystem SR630 V4, notamment en termes de nombre de cœurs par processeur. Le Lenovo ThinkSystem SR630 V4, équipé de deux processeurs Intel Xeon 2E (6780 cœurs par processeur), a surpassé le Dell au test CPU multicœur avec un score de 144 99,266 points, contre 74,710 288 points pour le Dell. Cet écart reflète l'avantage de Lenovo dans les charges de travail multithread, grâce à son nombre de cœurs plus élevé (2 cœurs au total) par rapport aux deux processeurs Intel Xeon 6787P du Dell (86 cœurs par processeur), ce qui limite ses performances multicœurs.
Lors du test CPU Single-Core, Dell a obtenu de meilleurs résultats avec un score de 1,272 894 points, surpassant les XNUMX points de Lenovo, soulignant l'efficacité supérieure du processeur monothread de Dell malgré son nombre de cœurs inférieur.
| Cinebench R23 | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 Go de RAM) | 
|---|---|---|
| Processeur multicœur | 74,710 pts | 99,266 pts | 
| Processeur monocœur | 1,272 pts | 894 pts | 
| Rapport PM | 58.74 x | 111.00 x | 
Cinebench 2024
Cinebench 2024 étend les capacités de référence de R23 en ajoutant une évaluation des performances du GPU. Il continue de tester les performances du processeur mais inclut également des tests qui mesurent la capacité du GPU à gérer les tâches de rendu.
Dans ce benchmark mis à jour, le Dell PowerEdge R770 a obtenu 12,996 630 points pour les performances du GPU, soulignant sa capacité à gérer les tâches de rendu accélérées par le GPU. Le Lenovo ThinkSystem SR4 VXNUMX n'a pas de GPU dédié et n'a donc pas enregistré de score GPU.
Lors du test CPU multicœur, le Lenovo a obtenu 2,884 2,831 points, légèrement devant les 71 53 points de Dell, ce qui témoigne d'un léger avantage en termes de performances multicœurs. En CPU monocœur, le Dell a surpassé le Lenovo, obtenant XNUMX points, contre XNUMX points pour Lenovo, ce qui témoigne des performances supérieures du Dell en monocœur malgré un nombre de cœurs réduit.
| Cinebench R24 | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 Go de RAM) | 
|---|---|---|
| Score GPU | 12,996 pts |  | 
| Processeur multicœur | 2,831 pts | 2,884 pts | 
| Processeur monocœur | 71 pts | 53 pts | 
| Rapport PM | 39.77 x | 54.43 x | 
Geekbench 6
Geekbench 6 est un outil d'évaluation des performances multiplateforme qui mesure les performances globales d'un système. Le navigateur Geekbench permet de comparer n'importe quel système à cet outil.
Les résultats du benchmark Geekbench 6 montrent des différences de performances évidentes entre le Dell PowerEdge R770 et le Lenovo ThinkSystem SR630 V4. Lors du test CPU monocœur, le Dell a surpassé le Lenovo avec un score de 1,797 1,173, tandis que ce dernier a obtenu 53 XNUMX, soit une amélioration de XNUMX % des performances monocœur pour le Dell.
Lors du test CPU multicœur, Dell a de nouveau dominé avec 15,880 13,868 points, tandis que Lenovo a obtenu 14 6787 points, ce qui confère à Dell un avantage de XNUMX % en performances multicœurs. Cela suggère que les processeurs Intel Xeon XNUMXP de Dell offrent une puissance de calcul globale supérieure, notamment pour les tâches nécessitant plusieurs cœurs.
Le test GPU OpenCL a encore plus mis en évidence l'avantage de Dell, avec un score de 148,730 4 grâce au GPU NVIDIA LXNUMX.
| Geekbench 6 (Plus c'est mieux) | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 Go de RAM) | 
|---|---|---|
| Processeur monocœur | 1,797 | 1,173 | 
| Processeur multicœur | 15,880 | 13,868 | 
| Score OpenCL du GPU | 148,730 | (pas de GPU) | 
Test de vitesse Blackmagic RAW
Le Blackmagic RAW Speed Test est un outil d'analyse comparative des performances conçu pour mesurer les capacités d'un système à gérer la lecture et l'édition vidéo à l'aide du codec Blackmagic RAW. Il évalue la capacité d'un système à décoder et à lire des fichiers vidéo haute résolution, en fournissant des fréquences d'images pour le traitement basé sur le CPU et le GPU.
Lors du test basé sur le processeur, le Dell PowerEdge R770 a atteint 141 ips, surpassant le Lenovo ThinkSystem SR630 V4, qui a obtenu 120 ips. Cela indique que le système Dell gère le traitement vidéo basé sur le processeur plus efficacement que le Lenovo. Lors du test basé sur le processeur graphique, le Dell PowerEdge R770 a obtenu 157 ips, bénéficiant de la présence d'un processeur graphique NVIDIA.
| Test de vitesse Blackmagic RAW (plus c'est élevé, mieux c'est) | Dell PowerEdge R770 (2 processeurs Intel Xeon 6787P \| 2 To de RAM) | Lenovo ThinkSystem SR630 V4 (2 x Intel Xeon 6780E \| 512 Go de RAM) | 
|---|---|---|
| FPS CPU | FPS 141 | FPS 120 | 
| FPS CUDA | FPS 157 | 0 FPS (pas de GPU) | 
Test de vitesse du disque Blackmagic
Le test de vitesse du disque Blackmagic évalue les vitesses de lecture et d'écriture d'un disque, en évaluant ses performances, en particulier pour les tâches de montage vidéo. Il aide les utilisateurs à s'assurer que leur stockage est suffisamment rapide pour le contenu haute résolution, comme la vidéo 4K ou 8K.
