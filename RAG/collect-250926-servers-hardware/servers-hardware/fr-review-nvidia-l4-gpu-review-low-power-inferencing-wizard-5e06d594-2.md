---
id: collect-250926-servers-hardware/servers-hardware/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594-2
title: "fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "benchmark", "benchmarks", "intel", "mlperf", "open source"]
source: docs/RAG/clean4/fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594.md
source_anchor: ""
source_lines: [75, 145]
sha256: 2b2a9f441973516dc31409d98ef8f2dddb37eacaebd6feb61d7233f3b116c5b5
---

# fr-review-nvidia-l4-gpu-review-low-power-inferencing-wizard-5e06d594

Lors de nos tests avec deux L4 dans le Dell T560, nous avons observé cette mise à l’échelle quasi linéaire des performances pour les benchmarks Resnet50 et BERT K99. Cette mise à l'échelle témoigne de l'efficacité des GPU L4 et de leur capacité à fonctionner en tandem sans pertes significatives dues à la surcharge ou à l'inefficacité.
| Dell PowerEdge T560 2x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Serveur | 24,407.50 | 
| Resnet50 – Hors ligne | 25,463.20 | 
| BERT K99 – Serveur | 1,801.28 | 
| BERT K99 – Hors ligne | 1,904.10 | 
La mise à l'échelle linéaire cohérente à laquelle nous avons assisté avec deux GPU NVIDIA L4 s'étend de manière impressionnante aux configurations comportant quatre unités L4. Cette mise à l'échelle est particulièrement remarquable car le maintien des gains de performances linéaires devient de plus en plus difficile avec chaque GPU ajouté en raison de la complexité du traitement parallèle et de la gestion des ressources.
| Dell PowerEdge T560 4x NVIDIA L4 | Score | 
|---|---|
| Resnet50 – Serveur | 48,818.30 | 
| Resnet50 – Hors ligne | 51,381.70 | 
| BERT K99 – Serveur | 3,604.96 | 
| BERT K99 – Hors ligne | 3,821.46 | 
Ces résultats sont fournis à titre indicatif uniquement et ne constituent pas des résultats officiels ou compétitifs de MLPerf. Pour consulter la liste complète des résultats officiels, veuillez vous rendre sur la page des résultats MLPerf.
En plus de valider l'évolutivité linéaire des GPU NVIDIA L4, nos tests en laboratoire mettent en lumière les implications pratiques du déploiement de ces unités dans différents scénarios opérationnels. Par exemple, la cohérence des performances entre les modes serveur et hors ligne dans toutes les configurations avec les GPU L4 révèle leur fiabilité et leur polyvalence.
Cet aspect est particulièrement pertinent pour les entreprises et les institutions de recherche où les contextes opérationnels varient considérablement. De plus, nos observations sur l’impact minimal des goulots d’étranglement d’interconnexion et l’efficacité de la synchronisation GPU dans les configurations multi-GPU fournissent des informations précieuses pour ceux qui cherchent à faire évoluer leur infrastructure d’IA. Ces informations vont au-delà de simples chiffres de référence, offrant une compréhension plus approfondie de la manière dont un tel matériel peut être utilisé de manière optimale dans des scénarios du monde réel, guidant de meilleures décisions architecturales et stratégies d'investissement dans l'infrastructure d'IA et HPC.
NVIDIA L4 – Performances des applications
Nous avons comparé les performances du nouveau NVIDIA L4 à celles des NVIDIA A2 et NVIDIA T4 qui l'ont précédé. Pour présenter cette amélioration des performances par rapport aux modèles précédents, nous avons déployé les trois modèles sur un serveur de notre laboratoire, avec Windows Server 2022 et les derniers pilotes NVIDIA, en exploitant l'ensemble de notre suite de tests GPU.
Ces cartes ont été testées sur un Dell Poweredge R760 avec la configuration suivante :
- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- Windows Server 2022
- Pilote NVIDIA 538.15
- ECC désactivé sur toutes les cartes pour un échantillonnage 1x
Alors que nous lançons les tests de performances entre ce groupe de trois GPU d'entreprise, il est important de noter les différences de performances uniques entre les modèles A2 et T4 précédents. Lorsque l'A2 est sorti, il offrait des améliorations notables telles qu'une consommation d'énergie inférieure et un fonctionnement sur un emplacement PCIe Gen4 x8 plus petit, au lieu du plus grand emplacement PCIe Gen3 x16 requis par l'ancien T4. Dès le départ, cela lui a permis de s'intégrer dans davantage de systèmes, en particulier avec le plus petit encombrement nécessaire.
Mélangeur OptiX 4.0
Blender OptiX est une application de modélisation 3D open source. Ce test peut être exécuté à la fois pour le CPU et le GPU, mais nous n'avons effectué que le GPU comme la plupart des autres tests ici. Ce benchmark a été exécuté à l'aide de l'utilitaire CLI Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
| Mixeur 4.0 (Plus haut, c'est mieux) | Nvidia L4 | Nvidia A2 | Nvidia T4 | 
|---|---|---|---|
| GPU Blender CLI – Monstre | 2,207.765 | 458.692 | 850.076 | 
| GPU Blender CLI – Junkshop | 1,127.829 | 292.553 | 517.243 | 
| GPU Blender CLI – Salle de classe | 1,111.753 | 262.387 | 478.786 | 
Test de vitesse Blackmagic RAW
Nous testons les CPU et les GPU avec le RAW Speed Test de Blackmagic qui teste les vitesses de lecture vidéo. Il s’agit plutôt d’un test hybride incluant les performances du CPU et du GPU pour le décodage RAW réel. Ceux-ci sont affichés sous forme de résultats séparés, mais nous nous concentrons ici uniquement sur les GPU, donc les résultats du CPU sont omis.
| Test de vitesse Blackmagic RAW (Plus haut, c'est mieux) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| CUDA 8K | FPS 95 | FPS 38 | FPS 53 | 
GPU Cinebench 2024
Cinebench 2024 de Maxon est une référence de rendu CPU et GPU qui utilise tous les cœurs et threads du processeur. Encore une fois, puisque nous nous concentrons sur les résultats du GPU, nous n’avons pas exécuté les parties CPU du test. Des scores plus élevés sont meilleurs.
| Cinebench 2024 (Plus haut, c'est mieux) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| GPU | 15,263 | 4,006 | 5,644 | 
GPU PI
GPUPI 3.3.3 est une version de l'utilitaire d'analyse comparative léger conçu pour calculer π (pi) en milliards de décimales à l'aide de l'accélération matérielle via les GPU et les CPU. Il exploite la puissance de calcul d'OpenCL et de CUDA, qui comprend des unités de traitement centrales et graphiques. Nous avons exécuté CUDA uniquement sur les 3 GPU et les chiffres ici sont le temps de calcul sans temps de réduction ajouté. Plus bas, c'est mieux.
| Temps de calcul du GPU PI en secondes (Plus bas, c'est mieux) | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
|---|---|---|---|
| GPUPI v3.3 – 1B | 3.732s | 19.799s | 7.504s | 
| GPUPI v3.3 – 32B | 244.380s | 1,210.801s | 486.231s | 
Alors que les résultats précédents ne portaient que sur une seule itération de chaque carte, nous avons également eu l'occasion d'examiner un déploiement de 5 cartes NVIDIA L4 à l'intérieur du Dell PowerEdge T560.
| Temps de calcul du GPU PI en secondes (Plus bas, c'est mieux) | Dell PowerEdge T560 (2x Xeon Gold 6448Y) avec 5x NVIDIA L4 | 
|---|---|
| GPUPI v3.3 – 1B | 0 s 850 ms | 
| GPUPI v3.3 – 32B | 50 s 361 ms | 
Banc Octane
OctaneBench est un utilitaire d'analyse comparative pour OctaneRender, un autre moteur de rendu 3D avec prise en charge RTX similaire à V-Ray.
| Octane (plus haut est mieux) |  |  |  |  | 
| Scène | Noyau | Nvidia L4 | Nvidia A2 | NVIDIA T4 | 
| Intérieur | Canaux d'information | 15.59 | 4.49 | 6.39 | 
|  | Eclairage direct | 50.85 | 14.32 | 21.76 | 
|  | Suivi de chemin | 64.02 | 18.46 | 25.76 | 
| L'idée | Canaux d'information | 9.30 | 2.77 | 3.93 | 
|  | Eclairage direct | 39.34 | 11.53 | 16.79 | 
|  | Suivi de chemin | 48.24 | 14.21 | 20.32 | 
| VTT | Canaux d'information | 24.38 | 6.83 | 9.50 | 
|  | Eclairage direct | 54.86 | 16.05 | 21.98 | 
|  | Suivi de chemin | 68.98 | 20.06 | 27.50 | 
| Skybox | Canaux d'information | 12.89 | 3.88 | 5.42 | 
|  | Eclairage direct | 48.80 | 14.59 | 21.36 | 
|  | Suivi de chemin | 54.56 | 16.51 | 23.85 | 
| Score total |  | 491.83 | 143.71 | 204.56 | 
Carte graphique Geekbench 6
