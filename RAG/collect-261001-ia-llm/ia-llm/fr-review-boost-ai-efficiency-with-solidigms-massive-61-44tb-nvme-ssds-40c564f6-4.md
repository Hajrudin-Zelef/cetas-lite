---
id: collect-261001-ia-llm/ia-llm/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6-4
title: "fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6"
domain: ia-llm
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["gpu", "nand", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6.md
source_anchor: ""
source_lines: [113, 146]
sha256: e0b26540b264376612333e9be43dc785c3fc509444ce9990546f10d9a4acd25b
---

# fr-review-boost-ai-efficiency-with-solidigms-massive-61-44tb-nvme-ssds-40c564f6

| Stockage->CPU->GPU | 32 | 25.12613 | 159296.2804 | 
| Stockage->GPU (GDS) | 4 | 25.057484 | 19946.07198 | 
| Stockage->GPU (GDS) | 16 | 25.044871 | 79770.6007 | 
| Stockage->GPU (GDS) | 32 | 25.031055 | 159478.8246 | 
| Stockage->PAGE_CACHE->CPU->GPU | 16 | 24.493948 | 109958.4447 | 
| Stockage->PAGE_CACHE->CPU->GPU | 32 | 24.126103 | 291792.8345 | 
| Stockage->GPU (GDS) | 1 | 23.305366 | 5362.611458 | 
| Stockage->PAGE_CACHE->CPU->GPU | 4 | 21.906704 | 22815.52797 | 
| Stockage->CPU->GPU | 1 | 15.27233 | 8182.667969 | 
| Stockage->PAGE_CACHE->CPU->GPU | 1 | 6.016992 | 20760.22778 | 
Écrire correctement toute application pour interagir avec le stockage est primordial et doit être pris en compte car les entreprises souhaitent maximiser leur investissement GPU.
GPU direct
En isolant les performances du GPU Direct uniquement dans tous les tests, nous pouvons avoir une idée générale de la façon dont la technologie NVIDIA brille.
| Type d'E / S | Type de transfert | Threads | Taille de l'ensemble de données (Kio) | Taille des E/S (Ko) | Débit (Gio/sec) | Latence moyenne (usecs) | 
|---|---|---|---|---|---|---|
| ÉCRIRE | GPUD | 8 | 777,375,744 | 1024 | 12.31 | 634.55 | 
| LIS | GPUD | 8 | 579,439,616 | 1024 | 9.30 | 840.37 | 
| RANDÉCRITURE | GPUD | 8 | 751,927,296 | 1024 | 12.04 | 648.67 | 
| RANDIRER | GPUD | 8 | 653,832,192 | 1024 | 10.50 | 743.89 | 
| ÉCRIRE | GPUD | 8 | 69,929,336 | 4 | 1.12 | 27.32 | 
| LIS | GPUD | 8 | 37,096,856 | 4 | 0.59 | 51.52 | 
| RANDÉCRITURE | GPUD | 8 | 8,522,752 | 4 | 0.14 | 224.05 | 
| RANDIRER | GPUD | 8 | 21,161,116 | 4 | 0.34 | 89.99 | 
| RANDÉCRITURE | GPUD | 8 | 57,083,336 | 4 | 0.91 | 33.42 | 
| RANDIRER | GPUD | 8 | 27,226,364 | 4 | 0.44 | 70.07 | 
Réflexions de clôture
Puisque cet article se concentre sur le Solidigm 61.44 To P5336, prenons du recul et abordons le débat TLC vs QLC autour des performances par rapport à la capacité. Lorsque l'on regarde d'autres produits du portefeuille Solidigm, comme la gamme D7, qui utilise TLC 3D NAND, la capacité est limitée en échange de performances. Lors de nos tests, en particulier avec les disques Solidigm de 61.44 To, nous avons constaté des performances de débit globales qui peuvent maintenir correctement les GPU alimentés en données à faibles latences. Nous entendons des retours d'ODM et d'OEM concernant la demande de stockage de plus en plus proche du GPU, et le disque Solidigm D5-P5336 semble faire l'affaire. Comme il existe généralement un nombre limité de baies NVMe disponibles sur les serveurs GPU, les disques denses Solidigm figurent en tête de liste pour le stockage sur serveur GPU local.
En définitive, la capacité de stockage massive offerte par ces disques, associée aux GPU, ne constitue qu'une partie de la solution ; leurs performances restent indispensables. En additionnant les performances d'un seul disque sur plusieurs, on constate qu'un débit suffisant est disponible, même pour les tâches les plus exigeantes. Dans le cas d'une configuration RAID 0 à quatre disques utilisant GDSIO, le débit total en écriture peut atteindre 12.31 Gio/s, et en lecture, 25.13 Gio/s.
Ce niveau de débit est plus que suffisant pour les tâches d'IA les plus exigeantes, telles que la formation de grands modèles d'apprentissage en profondeur sur des ensembles de données massifs ou l'exécution d'inférences en temps réel sur des flux vidéo haute résolution. La possibilité d'augmenter les performances en ajoutant davantage de disques à la matrice RAID0 en fait un choix incontournable pour les applications d'IA où un accès rapide et efficace aux données est crucial.
Cependant, il est important de noter que les configurations RAID0, bien qu'offrant des performances élevées, n'offrent aucune redondance des données. Il est donc essentiel de mettre en œuvre des stratégies de sauvegarde et de protection des données appropriées pour éviter toute perte de données en cas de panne de disque.
Une autre considération unique dans les centres de données aujourd'hui est la puissance. Alors que les serveurs d’IA consomment plus d’énergie que jamais et ne montrent aucun signe de ralentissement, la puissance totale disponible est l’un des plus gros goulots d’étranglement pour ceux qui cherchent à intégrer des GPU dans leurs centres de données. Cela signifie que l’accent est encore plus mis sur l’économie de chaque watt possible. Si vous pouvez obtenir plus de To par watt, nous abordons quelques processus de réflexion intéressants autour du coût total de possession et des coûts d'infrastructure. Même retirer ces disques du serveur GPU et les placer dans un serveur de stockage à l'échelle du rack peut fournir un débit massif avec des capacités extrêmes.
L'intégration de disques SSD QLC Solidigm D5-P5336 de 61.44 To avec des serveurs d'IA à emplacements limités NVMe représente une avancée significative pour relever les défis de stockage des charges de travail d'IA modernes. Leur densité extrême, leurs caractéristiques de performance et leur rapport To/watt les rendent idéaux pour les phases de préparation, de formation et de réglage des données, ainsi que d'inférence. En optimisant l'utilisation des voies PCIe et en fournissant des solutions de stockage haute capacité, ces SSD permettent à l'AI Factory moderne de se concentrer sur le développement et le déploiement de modèles plus sophistiqués et plus précis, stimulant ainsi l'innovation dans le domaine de l'IA.
Page Lenovo ThinkSystem SR675 V3
Ce rapport est sponsorisé par Solidigm. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
