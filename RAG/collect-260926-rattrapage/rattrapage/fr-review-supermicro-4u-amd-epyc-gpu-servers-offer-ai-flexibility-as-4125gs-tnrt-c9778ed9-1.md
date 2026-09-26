---
id: collect-260926-rattrapage/rattrapage/fr-review-supermicro-4u-amd-epyc-gpu-servers-offer-ai-flexibility-as-4125gs-tnrt-c9778ed9-1
title: "fr-review-supermicro-4u-amd-epyc-gpu-servers-offer-ai-flexibility-as-4125gs-tnrt-c9778ed9"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "OpenAI"]
dates: []
keywords: ["amd", "gpu", "distribution", "dram", "hbm3", "intel", "nvidia"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-supermicro-4u-amd-epyc-gpu-servers-offer-ai-flexibility-as-4125gs-tnrt-c9778ed9.md
source_anchor: ""
source_lines: [1, 29]
sha256: 496de63b250588773743d07cc1d20bc3a6bd0f103b46debabf94c05c84dbb956
---

# fr-review-supermicro-4u-amd-epyc-gpu-servers-offer-ai-flexibility-as-4125gs-tnrt-c9778ed9

Supermicro propose depuis longtemps des serveurs GPU dans plus de formes et de tailles que nous n'avons le temps d'aborder dans cette revue. Aujourd'hui, nous examinons leur relativement nouveau serveur GPU 4U refroidi par air qui prend en charge deux processeurs AMD EPYC série 9004, PCIe Gen5 et un choix de huit cartes GPU supplémentaires double largeur ou 12 cartes GPU simples largeur. Bien que Supermicro propose également des variantes de ces serveurs basées sur Intel, la famille AS-4125GS-TNRT basée sur AMD est la seule de cette classe à prendre en charge les GPU NVIDIA H100 et AMD Instinct Mi210.
Le serveur GPU Supermicro AS-4125GS-TNRT possède quelques autres points forts matériels tels que la mise en réseau 10GbE intégrée, la gestion hors bande, 9 emplacements FHFL PCIe Gen5, 24 baies 2.5″, dont quatre NVMe, et le reste SATA/SAS. Il existe également 4 alimentations redondantes de 2000 2 W au niveau du titane. Sur la carte mère, il y a un seul emplacement M.XNUMX NVMe pour le démarrage.
Avant d'aller trop loin dans cette voie, il convient également de mentionner que Supermicro propose deux autres variantes de la configuration du serveur AS-4125GS-TNRT. Bien qu'ils utilisent la même carte mère, l'AS-4125GS-TNRT1 est une configuration à socket unique avec un commutateur PCIe prenant en charge jusqu'à 10 GPU double largeur et 8 baies SSD NVMe. L'AS-4125GS-TNRT2 est une configuration biprocesseur qui revient plus ou moins à la même chose, toujours avec le switch PCIe.
Quelle que soit la configuration, le Supermicro AS-4125GS-TNRT est incroyablement flexible grâce à sa conception et sa capacité à sélectionner des modèles avec le commutateur PCIe. Ce style de serveur GPU est populaire car il permet aux organisations de démarrer petit et de développer, de mélanger et d'associer des GPU pour différents besoins, ou de faire tout ce qu'elles veulent. Les systèmes GPU à socket offrent la possibilité de mieux regrouper les GPU pour les grosses charges de travail d'IA, mais les systèmes de cartes d'extension sont imbattables en termes de flexibilité des charges de travail.
De plus, même si cela peut paraître un blasphème pour certains, les serveurs GPU de carte d'extension Supermicro peuvent même être utilisés avec des cartes AMD et NVIDIA dans le même boîtier ! Halètement, si vous voulez, mais de nombreux clients ont compris que certaines charges de travail préfèrent un Instinct, tandis que d'autres comme le GPU NVIDIA. Enfin, bien que moins populaires que les serveurs GPU remplis à craquer, il convient de mentionner que ces emplacements ne sont que des emplacements PCIe ; il n'est pas déraisonnable d'imaginer des scénarios dans lesquels les clients pourraient préférer les FPGA, les DPU ou toute autre forme d'accélérateur dans cette plate-forme. Encore une fois, la flexibilité est le principal avantage de cette conception.
Aux fins de notre examen, le Supermicro AS-4125GS-TNRT était livré en version barebone, prêt à ajouter du CPU, de la DRAM, du stockage et, bien sûr, des GPU. Nous avons travaillé avec Supermicro pour emprunter 4 GPU NVIDIA H100 pour cet examen.
Spécifications Supermicro AS-4125GS-TNRT
| Spécifications techniques |  | 
| Processeur | Processeurs Dual Socket SP5 jusqu'à 128C/256T chacun | 
| Mémoire | Jusqu'à 24 RDIMM/LRDIMM ECC DDR256 de 4800 Go à 5 XNUMX MHz (Mémoire totale de 6 To) | 
| GPU |  | 
| Slots d'extension | 9x emplacements PCIE 5.0 x16 FHFL | 
| Alimentations | 4 alimentations redondantes de 2000 XNUMX W | 
| Networking | 2 x 10 GbE | 
| Stockage |  | 
| Carte mère | Super H13DSG-O-CPU | 
| Direction |  | 
| Sécurité |  | 
| chassis Taille | 4U | 
Supermicro AS-4125GS-TNRT Examen de la configuration
Nous avons configuré notre système à partir de Supermicro comme barebones, bien qu'ils le vendent en grande partie comme un système configuré. Une fois arrivé au laboratoire, la première chose que nous avons faite a été de l'équiper d'une paire de processeurs AMD EPYC 9374F 32c 64t. Ceux-ci ont été sélectionnés pour leur vitesse d’horloge élevée et leurs performances multicœurs respectables.
Pour les accélérateurs, nous avions toute une gamme de choix, allant des anciens coprocesseurs Intel Phi aux dernières cartes PCIe H100 en passant par les GPU de station de travail RTX 6000 ada haut de gamme. Notre objectif était d’équilibrer la puissance de calcul brute avec l’efficacité et la polyvalence. Finalement, nous avons décidé de commencer avec quatre GPU NVIDIA RTX A6000, puis de passer à quatre cartes NVIDIA H100 PCIe pour nos premiers tests. Cette combinaison démontre la flexibilité de la plateforme Supermicro et des cartes accélératrices NVIDIA.
Le RTX A6000, principalement conçu pour les performances dans les charges de travail gourmandes en graphiques, excelle également dans les applications IA et HPC grâce à son architecture Ampere. Il offre 48 Go de mémoire GDDR6, ce qui le rend idéal pour gérer de grands ensembles de données et des simulations complexes. Ses 10,752 336 cœurs CUDA et XNUMX Tensor permettent un calcul accéléré, ce qui est crucial pour nos tests d'IA et d'apprentissage profond.
D'autre part, les cartes NVIDIA H100 PCIe sont les dernières cartes livrées de la gamme d'architecture Hopper, conçues principalement pour les charges de travail d'IA. Chaque carte comprend un nombre impressionnant de 80 milliards de transistors, 80 Go de mémoire HBM3 et le révolutionnaire Transformer Engine, conçu pour les modèles d'IA comme le GPT-4. Les cœurs Tensor de 100e génération et les instructions DPX du H4 améliorent considérablement les tâches d'inférence et de formation de l'IA.
En intégrant ces GPU dans notre système barebones Supermicro, nous nous sommes concentrés sur une gestion thermique et une distribution d'énergie optimales, compte tenu de la consommation d'énergie et de la génération de chaleur substantielles de ces composants haut de gamme. Le châssis Supermicro, bien que ne prenant pas officiellement en charge une telle configuration, s'est avéré suffisamment polyvalent pour s'adapter à notre configuration. Pour contrôler les performances thermiques des A6000, nous avons dû les espacer d'une largeur de carte en raison de la conception du ventilateur à cage d'écureuil, mais les H100 peuvent être emballés avec leurs ailettes de refroidissement passives passthrough.
Notre suite d'analyse comparative comprenait un mélange de cas d'utilisation spécifiques au HPC et à l'IA. Celles-ci allaient des charges de travail d'analyse comparative traditionnelles aux tâches de formation et d'inférence d'IA utilisant des modèles de réseaux neuronaux convolutifs. Notre objectif était de pousser ces accélérateurs dans leurs limites, en évaluant leurs performances brutes, leur efficacité, leur évolutivité et leur facilité d'intégration avec notre serveur Supermicro A+.
Tests GPU Supermicro AS-4125GS-TNRT
Dans le cadre de nos travaux en laboratoire sur un modèle CNN de base, nous avons commencé par tester les GPU phares de NVIDIA en effectuant un entraînement de niveau station de travail sur une paire de GPU RTX8000 plus anciens mais très performants.
Lors de notre analyse des performances de l'IA, nous avons observé une progression remarquable mais attendue des capacités, passant du NVIDIA RTX 8000 à quatre GPU RTX A6000 et enfin à quatre cartes NVIDIA H100 PCIe. Cette progression a mis en valeur la puissance brute de ces accélérateurs et l'évolution des accélérateurs NVIDIA au cours des dernières années, alors que l'accent est de plus en plus mis sur les charges de travail d'IA.
