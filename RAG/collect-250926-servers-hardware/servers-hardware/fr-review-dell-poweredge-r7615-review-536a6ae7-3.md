---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r7615-review-536a6ae7-3
title: "fr-review-dell-poweredge-r7615-review-536a6ae7"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "dram", "gpu", "intel", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r7615-review-536a6ae7.md
source_anchor: ""
source_lines: [70, 121]
sha256: ca978e26085a283cf68ed76e00474a7e1fa5ede92fc2c95bdc8a3580c3319da3
---

# fr-review-dell-poweredge-r7615-review-536a6ae7

Au milieu de la carte mère se trouve le processeur, qui est le processeur AMD EPYC 9354. De chaque côté du processeur/dissipateur thermique se trouvent 12 emplacements DIMM DDR5 (qui prennent en charge jusqu'à 3 To RDIMM et des vitesses allant jusqu'à 4800 16 MT/s). Dans notre cas, nous ne disposons que de XNUMX Go de RAM.
En passant aux emplacements d'extension, vous verrez 2 emplacements x8 pleine hauteur, 2 emplacements x16 profil bas et 2 emplacements x16 double largeur, tous prenant en charge PCIe Gen5. Cette variété permet une large gamme de cartes d'extension, vous donnant la flexibilité d'ajouter n'importe quoi, des cartes réseau supplémentaires aux contrôleurs de stockage haut de gamme.
En ce qui concerne les blocs d'alimentation, le PowerEdge R7615 offre une gamme de choix. En fonction des besoins électriques de votre version spécifique, vous pouvez opter pour des alimentations simples ou doubles, enfichables à chaud et entièrement redondantes, avec des puissances pouvant aller jusqu'à 2400 XNUMX W. Cette polyvalence garantit que le serveur peut être personnalisé pour répondre à des besoins d'alimentation spécifiques, ce qui est crucial pour les configurations hautes performances.
Performances Dell PowerEdge R7615
Comme nous l'avons indiqué ci-dessus, la configuration de notre serveur Dell PowerEdge R7615 est aussi basique que possible. Cela signifie que nous n’attendons pas beaucoup de performances lors de notre analyse comparative. Le système R7615 de notre examen comprend un processeur AMD EPYC 9354 (3.25 GHz, 32 cœurs/64 threads, 256 Mo de cache) et 16 Go de RAM DDR5.
Bien que le serveur soit conçu pour gérer deux GPU, notre système n’en possède aucun. Il est également limité sur la DRAM, nous avons donc adapté notre approche d'analyse comparative pour nous concentrer sur les tests qui exploitent fortement les capacités du processeur, ce qui signifie que nous n'effectuerons pas d'analyses comparatives telles que ESRI, Luxmark, SPECworkstation 3, 7zip et SPECviewperf 2020. Au lieu de cela, notre L'évaluation se concentrera sur des benchmarks qui mettent principalement l'accent sur le processeur pour montrer clairement ses performances dans les tâches gourmandes en calcul.
Mixeur OptiX
Le premier est le test Blender, une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
|  | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cœurs, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430 (32 cœurs, 3.7 GHz max)) | 
| Blender OptiX version 3.6 (CPU) (Échantillons par minute ; plus c'est élevé, mieux c'est) | Score | Score | 
| Monster | 397 | 279 | 
| Brocanteur | 245 | 177 | 
| Salle de classe | 200 | 135 | 
Lors de nos tests d'analyse comparative utilisant Blender OptiX version 3.6, la configuration de base du Dell PowerEdge R7615, doté d'un processeur AMD EPYC 9354P et de 16 Go de RAM DDR5, a donné des résultats modestes. Le serveur a noté 397 échantillons par minute sur le test « Monster », 245 échantillons par minute sur « Junkshop » et 200 échantillons par minute sur « Classroom ».
| Blender OptiX version 4.0 (CPU) (Échantillons par minute ; plus c'est élevé, mieux c'est) | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cœurs, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430 (32 cœurs, 3.7 GHz max)) | 
| Monster | 375 | 952 | 
| Brocanteur | 248 | 613 | 
| Salle de classe | 198 | 437 | 
Lors de l'analyse comparative de la version 4.0, le serveur Dell PowerEdge R7615 a obtenu 375 échantillons par minute dans le test « Monster », 248 échantillons par minute dans « Junkshop » et 198 échantillons par minute dans « Classroom ». Ces résultats reflètent une capacité de rendu stable, étroitement alignée sur les performances observées dans la version précédente de Blender.
Test de vitesse Blackmagic RAW
Nous avons également commencé à exécuter le test de vitesse RAW de Blackmagic, qui teste les performances de lecture vidéo. De plus, ce test est davantage un test hybride regroupant à la fois le CPU et le GPU dans un scénario réel pour le décodage RAW.
| Test de vitesse Blackmagic RAW (Plus c'est haut, mieux c'est) | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cœurs, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| CPU 8K | 59 | FPS 99 | 
| CUDA 8K | N/D | N/D | 
Le HPE ProLiant DL320 a démontré des performances modestes lors du Blackmagic RAW Speed Test. Dans les scénarios de lecture vidéo et de décodage RAW, il n’a atteint que 59 FPS lors du test CPU 8K. Ce serveur ne disposant pas de GPU, la partie 8K CUDA de ce test n’a pas été évaluée.
Geekbench 6
Geekbench 6 est un outil d'évaluation multiplateforme mesurant les performances globales d'un système. Toutefois, il serait intéressant d'analyser les performances monocœur et multicœur, ainsi que les résultats du benchmark OpenCL. Un score élevé indique de meilleures performances. Précisons que nous n'avons examiné que les résultats du processeur, ce serveur ne disposant pas de carte graphique.
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
| Geekbench 6 | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cœurs, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| CPU Benchmark - Monocœur | 2,070 | 1,732 | 
| Référence CPU - Multi-Core | 12,049 | 12,792 | 
| Référence GPU – OpenCL | N/D | N/D | 
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs.
| Cinebench R23 | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cœurs, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| Processeur (multicœur) (points) | 52,401 | 38,707 | 
| Processeur (monocœur) (points) | 1,343 | 1,245 | 
| Rapport PM | 39.03x | 31.10x | 
Cinebench 2024
Voici les résultats de la version 2024 de Cinebench, en termes de performances CPU et GPU.
| Cinebench R23 | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cœurs, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| Processeur (multicœur) (points) | 2,127 | 2,014 | 
| Processeur (monocœur) (points) | 83 | 71 | 
| Rapport PM | 25.61x | 28.29x | 
croque-y
y-cruncher est un programme multithread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des milliards de chiffres. Depuis son lancement en 2009, y-cruncher est devenu une application d'analyse comparative et de test de stress populaire auprès des overclockeurs et des passionnés de matériel. Plus vite, c'est mieux dans ce test.
| croque-y (Temps de calcul total) | Dell PowerEdge R7615 (AMD EPYC 9354P, 32 cœurs, 3.25 GHz) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| 1 milliard de chiffres (secondes) | 36.686 | 21.452 | 
| 2.5 milliard de chiffres (secondes) | 98.68 | 50.418 | 
| 10 milliard de chiffres (secondes) |  | 131.135 | 
Dans le benchmark y-cruncher, conçu pour tester les performances informatiques, notre version Dell PowerEdge R7615 a présenté ses capacités en matière de traitement des nombres. Il a complété le calcul d'un milliard de chiffres en 1 secondes et a mis 36.686 secondes pour traiter 98.68 milliards de chiffres. Ces temps reflètent la modeste maîtrise du serveur dans la gestion de tâches de calcul intensives, démontrant de solides performances pour sa configuration de base.
Inférence UL Procyon AI (CPU)
