---
id: collect-261001-general-networking/general-networking/fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f-1
title: "fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "fp8", "gpu", "nvlink"]
source: docs/RAG/collect-261001-general-networking/fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f.md
source_anchor: ""
source_lines: [1, 39]
sha256: 6d30c79998f070913a0ef4984f713596546a055501c06b6936e821276ab6dc3f
---

# fr-review-nvidia-l40s-for-omniverse-from-openusd-scenes-to-physics-aware-world-m-2da62e9f

Le L40S comble une lacune cruciale dans l'écosystème GPU des centres de données. Alors que les GPU d'entraînement IA axés sur le calcul, comme le Nvidia H100, privilégient les performances brutes sans accélération graphique, les cartes de visualisation professionnelles traditionnelles manquent généralement des capacités de calcul IA requises par les charges de travail d'inférence modernes et les nouvelles applications graphiques pilotées par l'IA. Ce positionnement le rend particulièrement précieux pour la génération de données synthétiques, le développement d'IA multimodale et les applications Omniverse, où les performances de calcul et graphiques sont essentielles.
Spécifications NVIDIA L40S
| Spécifications | L40 | L40S | PCIe H100 80G | 
| Architecture GPU | Ada Lovelace | Ada Lovelace | Hopper | 
| Matrice GPU | AD102 | AD102 | GH100 | 
| Cœurs de CUDA | 18,176 | 18,176 | 14,592 | 
| Noyaux de tenseurs | 568 (4e génération) | 568 (4e génération) | 456 | 
| RT Cœurs | 142 (3e génération) | 142 (3e génération) | - | 
| Mémoire GPU | 48 Go GDDR6 avec ECC | 48 Go GDDR6 avec ECC | 80 Go HBM2e | 
| Bande passante mémoire | 864 GB / s | 864 GB / s | 2 TB / s | 
| Interface de mémoire | 384-bits | 384-bits | 5120-bits | 
| Consommation maximale | 300W | 350W | 350W | 
| Facteur de forme | 4.4″ H x 10.5″ L : Double fente | 4.4″ H x 10.5″ L : Double fente | 4.4″ H x 10.5″ L : Double fente | 
| Solution thermique | Revenu | Revenu | Revenu | 
| Connecteurs d'affichage | 4x DisplayPort 1.4a | 4x DisplayPort 1.4a | - | 
| Interface PCIe | Gen4 x16 | Gen4 x16 | Gen5 x16 | 
| Câble d'alimentation | 16 broches | 16 broches | 16 broches | 
| Prise en charge vGPU | Oui | Oui | Non | 
| GPU multi-instances (MiG) | Non | Non | Oui | 
| Prise en charge de NVLink | Non | Non | Non | 
Caractéristiques de performances
| Métrique | Performances du L40 | Performances du L40S | Performances du H100 | 
| Performances FP32 | 90.5 TFLOPS | 91.6 TFLOPS | 51.2 TFLOPS | 
| Noyau tenseur TF32 | 362.1 TFLOPS | 366 TFLOPS | 756 TFLOPS | 
| Noyau tenseur FP16 | 724 TFLOPS | 733 TFLOPS | 1513 TFLOPS | 
| Noyau tenseur FP8 | 1,448 TFLOPS | 1,466 TFLOPS | 3026 TFLOPS | 
| Pic du tenseur INT8 TOPS | 1,448 TFLOPS | 1,466 TFLOPS | 3026 TOPS | 
| Performances de base RT | 209 TFLOPS | 212 TFLOPS | - | 
(Les chiffres de performance sont avec parcimonie)
NVIDIA L40S contre H100
Un examen détaillé des spécifications révèle les philosophies de conception distinctes des NVIDIA L40S et H100. Le L40S repose sur l'architecture Ada Lovelace, utilisant la même matrice AD102 que celle des cartes graphiques NVIDIA haut de gamme pour stations de travail. Cet héritage lui confère un nombre impressionnant de 18,176 32 cœurs CUDA, offrant d'excellentes performances FP100 en simple précision, pierre angulaire du rendu graphique traditionnel et du calcul scientifique. En revanche, l'architecture Hopper et la matrice GH100 du HXNUMX sont avant tout conçues pour les charges de travail d'IA et de calcul haute performance.
Le différenciateur le plus marquant réside dans la configuration des cœurs. Le L40S comprend 142 cœurs RT de troisième génération et 568 cœurs Tensor de quatrième génération. Les cœurs RT sont des composants spécialisés pour l'accélération du ray tracing, une fonctionnalité totalement absente du H100, ce qui confère au L40S une capacité unique de rendu photoréaliste. Bien que le H100 dispose de moins de cœurs Tensor, ceux-ci sont plus rapides et plus avancés, optimisés pour les nouveaux formats de données d'IA comme FP8, ce qui lui confère une avance considérable en termes de performances d'IA brute.
Ce compromis est également évident au niveau des sous-systèmes mémoire. Le H100 utilise 80 Go de mémoire HBM2e onéreuse, offrant une bande passante impressionnante de 2 To/s. Ceci est essentiel pour alimenter les SM lors de l'entraînement et de l'inférence de modèles d'IA à grande échelle. Le L40S utilise une mémoire GDDR48 plus conventionnelle de 6 Go, offrant une bande passante de 864 Go/s. Bien que inférieure à la moitié de celle du H100, cette capacité reste conséquente, parfaitement adaptée au chargement de grandes scènes 3D, de textures haute résolution et de modèles d'IA volumineux pour l'inférence.
Enfin, l'ensemble des fonctionnalités décrit leurs rôles respectifs. Le L40S comprend quatre sorties DisplayPort 1.4a et une prise en charge robuste des vGPU, ce qui le rend idéal pour les stations de travail virtualisées, les fermes de rendu et les déploiements de cloud gaming. Le H100, dépourvu de sorties d'affichage et de capacités vGPU, intègre la technologie GPU multi-instances (MiG), qui permet de le partitionner en plusieurs instances GPU plus petites et isolées pour gérer simultanément plusieurs charges de travail gourmandes en ressources de calcul. La conception thermique passive à double emplacement partagé et l'enveloppe de puissance de 350 W du L40S et du PCIe H100 offrent une flexibilité de déploiement sur une large gamme de serveurs standard.
En comparant le L40S au H100, le produit phare de NVIDIA, les différences entre leurs rôles respectifs deviennent évidentes. Le H100 est le leader incontesté des performances d'entraînement IA brutes de cette génération de GPU, mais cette comparaison ne révèle qu'une partie de l'histoire. Il est intéressant de noter que le L40 fait également partie de la gamme NVIDIA, avec un TDP inférieur de 300 W par rapport aux 40 W du L350S, offrant des capacités similaires pour une consommation énergétique réduite.
Le H100 que nous examinons ici est la version PCIe originale de 80 Go avec mémoire HBM2e. NVIDIA a depuis élargi sa gamme avec des variantes comme le H100 NVL, remplaçant le modèle PCIe original de 80 Go. Il offre 94 Go de mémoire et un TDP supérieur de 400 W, tout en intégrant la prise en charge NVLink pour les configurations à double GPU. La famille H100 s'étend également aux configurations SXM à 8 GPU et au nouveau H200, qui partage la même puce GPU mais offre des performances améliorées aux formats PCIe et SXM.
La distinction entre les modèles L40S et H100 met en évidence une divergence stratégique dans la conception des GPU pour centres de données. Le H100 est une carte purement axée sur le calcul, optimisée exclusivement pour l'IA et les charges de travail de calcul haute performance. L'objectif est l'entraînement de l'IA à grande échelle, où le débit de calcul maximal est l'objectif principal. Chaque aspect de sa conception, de l'énorme bande passante mémoire HBM2e aux cœurs Tensor spécialisés, est conçu pour les scénarios d'entraînement des réseaux neuronaux les plus exigeants.
En revanche, le L40S est un GPU universel conçu pour la polyvalence, ciblant l'inférence IA, les charges de travail gourmandes en ressources graphiques et les déploiements NVIDIA vGPU. S'il constitue une solution performante et économique pour l'inférence IA, son véritable atout réside dans sa prise en charge d'un large éventail d'applications gourmandes en ressources graphiques, pour lesquelles le H100 n'est tout simplement pas conçu.
Charges de travail gourmandes en graphiques et applications professionnelles
