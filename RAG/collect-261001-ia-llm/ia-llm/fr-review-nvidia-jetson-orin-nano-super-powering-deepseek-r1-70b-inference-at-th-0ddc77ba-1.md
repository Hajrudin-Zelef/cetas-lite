---
id: collect-261001-ia-llm/ia-llm/fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba-1
title: "fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba"
domain: ia-llm
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["inference", "nvidia", "ethernet", "gpu", "nand", "valuation"]
source: docs/RAG/collect-261001-ia-llm/fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba.md
source_anchor: ""
source_lines: [1, 41]
sha256: fc621c159fe7875c03d68878bd6723225baeb124f4ac0205c0dd038f0de5bd0d
---

# fr-review-nvidia-jetson-orin-nano-super-powering-deepseek-r1-70b-inference-at-th-0ddc77ba

Le Jetson Orin Nano Super est une centrale de calcul compacte qui apporte des capacités d'IA sophistiquées aux appareils de pointe. Il allie performances, prix abordable et options d'intégration solides, ce qui en fait un candidat idéal pour le prototypage et le développement de produits commerciaux. Qu'il soit utilisé dans des kits robotiques ou intégré dans des machines plus grandes, sa conception flexible permet aux ingénieurs de déployer l'IA dans des scénarios qui exigent efficacité et faible consommation d'énergie - pour seulement 249 $.
La plateforme Jetson est spécialement conçue pour les déploiements en périphérie, garantissant que les projets dans des environnements avec un espace ou une puissance limités puissent toujours exploiter les performances de l'IA haut de gamme. Avec un format évolutif et des options de connectivité étendues, elle offre une passerelle vers des solutions innovantes en matière de robotique, de surveillance intelligente et même de conservation de la faune.
La carte Jetson Orin Nano Super est réputée pour la réalisation de projets nécessitant une intelligence artificielle embarquée, que ce soit dans des kits robotiques traditionnels utilisant une programmation classique ou dans des configurations plus avancées intégrant des frameworks tels que ROS (Robot Operating System). Disponible sous forme de kit de développement complet et de carte fille SoC autonome, elle s'intègre facilement à une vaste gamme de produits et de machines. Cette polyvalence la rend populaire pour des applications allant des petits projets éducatifs aux déploiements industriels à grande échelle.
Spécifications du kit de développement Jetson Orin Nano Super
Le Jetson Orin Nano Super intègre des fonctionnalités impressionnantes dans un format compact. Le processeur Arm Cortex-A6AE à 78 cœurs constitue une base solide pour le calcul, tandis que le GPU NVIDIA Ampere à 1024 cœurs avec Tensor Cores accélère diverses charges de travail, notamment les tâches d'apprentissage profond et de vision par ordinateur. Avec 67 TOPS (Tera Operations Per Second) de performances d'IA et une mémoire LPDDR8 à large bande passante de 5 Go, cette plate-forme est conçue pour effectuer des opérations complexes en périphérie.
| Spécifications | DÉTAILS | 
|---|---|
| Processeur | Processeur Arm Cortex-A6AE v78 8.2 bits à 64 cœurs, 3 Mo L2 + 4 Mo L3 | 
| GPU | GPU d'architecture NVIDIA Ampere à 1024 cœurs avec 32 cœurs Tensor | 
| Performances de l'IA | 67 TOPS | 
| Mémoire | 8 Go 128 bits LPDDR5 102 Go/s | 
| Stockage | Prise en charge SSD NVMe 16 Go eMMC 5.1, microSD, M.2 Key M 1x emplacement M.2 Key M avec x4 PCIe Gen3 1x emplacement M.2 Key M avec x2 PCIe Gen3 | 
| Networking | Gigabit Ethernet 1x | 
| Écran | 1x HDMI, 1x eDP 1.4 | 
| Connectivité | 4 ports USB 3.2 de type A, 1 port USB de type C | 
| Alimentation | La prise jack cylindrique CC accepte une alimentation de 7 V à 20 V | 
| Caméra | 2x connecteurs de caméra MIPI CSI | 
| Expansion | Connecteurs d'extension GPIO à 40 broches | 
| Consommation d'énergie | 7W – 25W configurable | 
| Système d'exploitation | Linux basé sur Ubuntu avec NVIDIA JetPack SDK | 
| Dimensions | 103mm x x 90.5mm 34.77mm | 
Les options de connectivité sont nombreuses, ce qui rend le Nano Super extrêmement polyvalent pour de nombreuses applications. Quatre ports USB 3.2 Type-A et un port USB Type-C vous permettent de connecter facilement une gamme de périphériques, des périphériques de stockage externes aux périphériques d'entrée ou aux capteurs. Le Gigabit Ethernet intégré garantit une mise en réseau fiable, tandis que les deux connecteurs de caméra MIPI CSI permettent l'intégration de deux caméras. Cette fonctionnalité est particulièrement avantageuse pour les applications nécessitant une perception de la profondeur, essentielle dans la robotique et les systèmes autonomes où une cartographie environnementale précise est essentielle.
Les capacités de stockage incluent 16 Go de mémoire eMMC 5.1, une carte microSD et un double SSD M.2 NVMe via des emplacements dédiés avec connectivité PCIe Gen3. Cela offre un espace de stockage suffisant pour les systèmes d'exploitation, les logiciels et les ensembles de données et prend en charge les transferts de données à haut débit nécessaires aux analyses en temps réel et aux tâches d'inférence d'IA. De plus, l'inclusion d'interfaces HDMI et eDP 1.4 permet au Nano Super de prendre en charge les écrans, ce qui le rend idéal pour les applications de type kiosque ou l'affichage numérique.
Pousser le nano-super à ses limites : LLM Inference at the Edge
Notre travail avec le Nano Super s'est concentré sur l'exploration de son potentiel pour effectuer des tâches de développement d'IA, en particulier l'inférence de modèles de langage volumineux (LLM). Nous avons reconnu que les limitations de la mémoire embarquée compliquaient l'exécution de modèles avec des milliards de paramètres, nous avons donc mis en œuvre une approche innovante pour contourner ces contraintes. En règle générale, les 8 Go de mémoire graphique du Nano Super limitent ses capacités à des modèles plus petits, mais nous avons cherché à exécuter un modèle 45 fois plus grand que ce qui conviendrait traditionnellement.
Nous avons amélioré le stockage du Nano Super en intégrant le SSD Solidigm D5-P5336 de 122.88 To , récemment lancé , un disque NVMe à très haute capacité conçu pour les environnements de centres de données, afin de soutenir cette tâche ambitieuse.
Le SSD Solidigm D5-P5336 de 122 To est une solution de stockage révolutionnaire pour les charges de travail gourmandes en données, notamment dans les domaines de l'IA et des centres de données. Voici ses spécifications détaillées :
- Capacités: 122.88TB
- Technologie:NAND à quatre niveaux (QLC)
- Interface: PCIe x4 de 4e génération
- Performances:Jusqu'à 15 % d'amélioration sur les charges de travail gourmandes en données par rapport aux modèles précédents
- Facteur de forme:U.2 Environ la taille d'un jeu de cartes
- Cas d'usage:Idéal pour la formation de l'IA, la collecte de données, la capture multimédia et le transcodage
Indicateurs de performance
- Vitesse de lecture / écriture séquentielle:Jusqu'à 7.1 Go/s (lecture) et 3.3 Go/s (écriture)
- Performances aléatoires:Jusqu'à 1,269,000 XNUMX XNUMX IOPS
Mesures de la durée de vie
- Endurance:Le SSD Solidigm 122 To est conçu pour les charges de travail gourmandes en données et offre une endurance élevée. Vous pouvez utiliser le Estimateur d'endurance SSD Solidigm pour calculer la durée de vie prévue en fonction de charges de travail spécifiques.
Mesures de puissance
- To par watt=122 To25 W=4.88 To/WTo par watt=25 W122 To=4.88 To/W. Avec ces mesures de puissance, ce disque offre environ 4.88 téraoctets de stockage par watt d'énergie consommée, soulignant son efficacité pour les applications gourmandes en données.
Le Nano Super comprend deux baies NVMe M.2, que nous avons testées dans le cadre de cette évaluation. Les deux emplacements offrent une connexion PCIe Gen3, avec un emplacement de 30 mm prenant en charge 2 voies PCIe et un emplacement de 80 mm prenant en charge 4 voies PCIe complètes. Nous avons utilisé l'emplacement de 80 mm associé à un câble de dérivation pour fournir la plus grande bande passante au SSD QLC Solidigm D5-P5336 122 To. Notre câble d'alimentation USB-C n'était pas prêt pour la démonstration, nous avons donc utilisé une alimentation ATX qui fournissait 12 V et 3.3 V au lecteur U.2.
