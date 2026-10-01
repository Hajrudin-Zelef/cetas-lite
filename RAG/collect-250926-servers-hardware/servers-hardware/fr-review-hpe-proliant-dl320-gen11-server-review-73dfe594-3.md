---
id: collect-250926-servers-hardware/servers-hardware/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594-3
title: "fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["benchmark", "benchmarks", "dram", "gpu", "intel", "nvidia", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594.md
source_anchor: ""
source_lines: [79, 133]
sha256: 9238205afd9fe9ab4a803edc2f5a0fd475ab794da9a186ef4ebd4970cc35a3c0
---

# fr-review-hpe-proliant-dl320-gen11-server-review-73dfe594

Les capacités de stockage du DL320 Gen11, qui prennent en charge une combinaison de disques SFF et LFF, suggèrent une polyvalence, répondant à une gamme d'exigences de stockage allant des SSD haute vitesse aux disques durs de plus grande capacité. Les contrôleurs de stockage du DL360 Gen11, tels que le HPE MR408i-o Gen11, se concentrent sur des solutions de stockage hautes performances pour prendre en charge des applications plus gourmandes en données.
Performances HPE ProLiant DL320 Gen11
Notre version de serveur HPE ProLiant DL320 Gen11 présente une configuration de base adaptée aux PME. En tant que tel, nous n’attendons pas beaucoup de performances lors de notre analyse comparative. Cela dit, de nombreuses tâches dans lesquelles ces serveurs sont déployés ne sont pas particulièrement intensives, les données de performances sont donc instructives en matière d'évolutivité.
À la base, notre examen DL320 Gen11 comprend le processeur Intel Xeon-G 32 à 6430 cœurs et 32 Go de RAM DDR5 (via deux clés de 16 Go) et comprend deux SSD SATA RI SFF BC MV de 1.92 To. Il dispose également de deux blocs d'alimentation FS Platinum Hot-Plug LH de 800 W pour une alimentation fiable et efficace.
Pour compléter la version, un contrôleur de stockage HPE MR408i-o Gen11 SPDM, qui offre une prise en charge RAID, une protection améliorée des données et un adaptateur BASE-T 1 ports de 4 Go pour gérer l'activité réseau.
Le serveur ne dispose pas de GPU et sa configuration est légère en DRAM. Nous avons donc adapté notre approche d'analyse comparative pour nous concentrer sur les tests qui exploitent fortement les capacités du processeur, ce qui signifie que nous n'effectuerons pas d'analyses comparatives telles que ESRI, Luxmark, SPECworkstation 3, 7zip, et SPECviewperf 2020. Au lieu de cela, notre évaluation se concentrera sur des benchmarks qui mettent principalement l'accent sur le processeur pour montrer clairement ses performances dans les tâches gourmandes en calcul.
Mixeur OptiX
Le premier est le test Blender, une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
| HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430 (32 cœurs, 3.7 GHz max)) |  | 
| Blender OptiX version 3.6 (processeur) (Échantillons par minute ; plus c'est élevé, mieux c'est) | Score | 
| Monster | 279 | 
| Brocanteur | 177 | 
| Salle de classe | 135 | 
Lors des tests de référence pour la version 3.5 de Blender, le HPE ProLiant DL320 a obtenu des résultats variables dans différents scénarios de rendu avec son processeur Intel Xeon-G 4 de 6430e génération. Pour la scène « Monster », il a enregistré un score de 279 échantillons par minute. La scène « Junkshop » (qui implique un niveau de détail et de complexité modéré) a donné un score de 177 échantillons par minute. Enfin, dans la scène « Classroom », qui pourrait être la plus exigeante des trois en raison de détails complexes et de calculs d'éclairage plus nombreux, le serveur a enregistré 135 échantillons par minute. Ces résultats démontrent qu'il n'est pas destiné aux tâches de rendu complexes telles que configurées.
| HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430 (32 cœurs, 3.7 GHz max)) |  | 
| Blender OptiX version 4.0 (processeur) (Échantillons par minute ; plus c'est élevé, mieux c'est) |  | 
| Monster | 952 | 
| Brocanteur | 613 | 
| Salle de classe | 437 | 
Dans la version 4.0 de Blender, le HPE ProLiant DL320 a fourni des résultats améliorés, affichant 952 échantillons par minute pour la scène « Monster », 613 pour « Junkshop » et 437 pour « Classroom ».
Test de vitesse Blackmagic RAW
Nous avons également commencé à exécuter le test de vitesse RAW de Blackmagic, qui teste les performances de lecture vidéo. De plus, ce test est davantage un test hybride regroupant à la fois le CPU et le GPU dans un scénario réel pour le décodage RAW.
| Test de vitesse Blackmagic RAW (Plus c'est haut, mieux c'est) | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| CPU 8K | FPS 99 | 
| CUDA 8K | N/D | 
Le HPE ProLiant DL320 a démontré des performances décentes lors du Blackmagic RAW Speed Test. Dans les scénarios de lecture vidéo et de décodage RAW, il a atteint 99 FPS lors du test du processeur 8K, ce qui indique qu'il peut gérer efficacement certains traitements vidéo haute résolution. Comme nous n'avons pas de GPU dans cette version, nous n'avons pas évalué la partie 8K CUDA de ce test.
Geekbench 6
Geekbench 6 est un outil d'évaluation multiplateforme mesurant les performances globales d'un système. Toutefois, il serait intéressant d'analyser les performances monocœur et multicœur, ainsi que les résultats du benchmark OpenCL. Un score élevé indique de meilleures performances. Précisons que nous n'avons examiné que les résultats du processeur, ce serveur ne disposant pas de carte graphique.
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
| Geekbench 6 | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| CPU Benchmark - Monocœur | 1,732 | 
| Référence CPU - Multi-Core | 12,792 | 
| Référence GPU – OpenCL | N/D | 
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs.
| Cinebench R23 | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| Processeur (multicœur) (points) | 38,707 | 
| Processeur (monocœur) (points) | 1,245 | 
| Rapport PM | 31.10x | 
Cinebench 2024
Voici les résultats de la version 2024 de Cinebench, en termes de performances CPU et GPU.
| Cinebench R23 | HPE ProLiant DL320 (Intel 4th Processeur Gen Xeon-G 6430, 32 cœurs, 3.7 GHz maximum) | 
| Processeur (multicœur) (points) | 2,014 | 
| Processeur (monocœur) (points) | 71 | 
| Rapport PM | 28.29x | 
croque-y
y-cruncher est un programme multithread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des milliards de chiffres. Depuis son lancement en 2009, y-cruncher est devenu une application d'analyse comparative et de test de stress populaire auprès des overclockeurs et des passionnés de matériel. Plus vite, c'est mieux dans ce test.
| y-cruncher (temps de calcul total) |  | 
| 1 milliard de chiffres (secondes) | 21.452 | 
| 2.5 milliard de chiffres (secondes) | 50.418 | 
| 10 milliard de chiffres (secondes) | 131.135 | 
Les résultats du benchmark y-cruncher, tout en indiquant les limites de performances de la version bas de gamme, peuvent toujours être pertinents pour les petites et moyennes entreprises (PME) dans des contextes spécifiques. Ce type de système pourrait être tout à fait adéquat pour les PME dont les besoins informatiques ne sont pas centrés sur des tâches informatiques de haute intensité.
Conclusion
Le HPE ProLiant DL320 Gen11 est un serveur 1U 1P hautement adaptable et efficace, compte tenu de la configuration appropriée, conçu pour répondre aux exigences évolutives des charges de travail Edge AI, des environnements ROBO (Remote Office/Branch Office) et des applications PME.
En son cœur se trouve un seul processeur évolutif Intel Xeon de 4e génération, prenant en charge jusqu'à 32 cœurs et 270 W, ainsi qu'une capacité de mémoire allant jusqu'à 2 To à 4800 5 MT/s et des emplacements PCIe Gen4. HPE propose également quelques configurations intéressantes, dont une prenant en charge une douzaine de disques durs et une autre avec la prise en charge de 4x GPU NVIDIA LXNUMX, ce qui est idéal pour l'inférence.
