---
id: collect-250926-servers-hardware/servers-hardware/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee-1
title: "fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["gpu", "intel"]
source: docs/RAG/clean4/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee.md
source_anchor: ""
source_lines: [1, 52]
sha256: 4f5ccc0455b5c8b02ebbe6daa1074717298e9e3e7881312f52ed50af1cc15ab0
---

# fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee

La communauté des entreprises Virtual Desktop Infrastructure (VDI) réclame des solutions plus robustes. Alors que les entreprises s’efforcent d’améliorer leur efficacité et leur expérience utilisateur, le rôle des accélérateurs matériels dédiés est devenu essentiel. La série Data Center GPU Flex d'Intel occupe une place importante dans cet espace, offrant des solutions sur mesure pour les jeux en nuage, les médias, le VDI et l'accélération graphique au sein des centres de données.
Que sont les GPU Intel Flex Series ?
L'offre d'Intel repose principalement sur deux produits phares : les GPU Flex 140 et Flex 170. La Flex Series 140, une carte PCIe Gen4 à profil bas, intègre deux GPU, chacun doté de huit cœurs Xe et de 6 Go de mémoire GDDR6. Cette configuration est idéale pour gérer jusqu'à 12 sessions VDI et répond aux besoins des utilisateurs ayant des exigences graphiques modérées.
Pour les applications plus gourmandes en graphiques, la Flex Series 170 évolue avec un seul nœud GPU doté de 32 cœurs Xe et de 16 Go de mémoire GDDR6, le tout sur une carte PCIe pleine taille, offrant une puissance de feu substantielle pour les tâches haute résolution.
|  | GPU pour centre de données Intel Flex 170 | GPU pour centre de données Intel Flex 140 | 
|---|---|---|
| Essentiels |  |  | 
| Microarchitecture | Xe HPG | Xe HPG | 
| Options embarquées disponibles | Non | Non | 
| Conditions d'utilisation | Serveur/Entreprise | Serveur/Entreprise | 
| Cas d'utilisation | Cloud Computing | Cloud Computing | 
| Spécifications du processeur graphique |  |  | 
| xe-couleurs | 32 | 16 | 
| Tranches de rendu | 8 | 4 | 
| Unités de lancer de rayons | 32 | 16 | 
| Moteurs d'extensions matricielles Intel® Xe (Intel® XMX) | 512 | 256 | 
| Unités d'exécution | 512 | 256 | 
| Horloge dynamique Graphics Max | 2050 MHz | 1950 MHz | 
| Extensions matricielles Intel® Xe (Intel® XMX) Horloge dynamique maximale | 1950 MHz | 1600 MHz | 
| TBP | 150 W | 75 W | 
| Caractéristiques de la mémoire |  |  | 
| Taille de la mémoire | 16 GB | 12 GB | 
| Type de mémoire | GDDR6 | GDDR6 | 
| Interface mémoire graphique | 256 Bits | 192 Bits | 
| Bande passante de la mémoire graphique | 576 GB / s | 336 GB / s | 
| Technologies prises en charge |  |  | 
| Lancer de rayons | Oui | Oui | 
| Prise en charge d'une API | Oui | Oui | 
| Prise en charge d'OpenVINO™ | Oui | Oui | 
| Prise en charge de DirectX* | DirectX 12 Ultime | DirectX 12 Ultime | 
| Assistance Vulkan* | 1.3 | 1.3 | 
| Prise en charge d'OpenGL* | Jusqu'à 4.6 | Jusqu'à 4.6 | 
| Prise en charge d'OpenCL* | 3 | 3 | 
| Moteurs de codecs multiformats | 2 | 4 | 
| Fonctionnalité |  |  | 
| Encodage/décodage matériel H.264 | Oui | Oui | 
| Encodage/décodage matériel H.265 (HEVC) | Oui | Oui | 
| Encodage/Décodage AV1 | Oui | Oui | 
| Flux binaire et décodage VP9 | Oui | Oui | 
Le fondement de ces unités de traitement graphique est le Xe-core, le Flex Series 170 comportant deux fois plus de cœurs et de tranches de rendu que le Flex 140. Cela trace une voie directe vers des prouesses de traçage de rayons doublées, renforçant la capacité du Flex 170 à gérer un rendu complexe. tâches et simulations facilement.
Sur le plan de la mémoire, nous voyons également le Flex Series 170 prendre de l'avance avec une allocation GDDR16 de 6 Go, surclassant les 140 Go du Flex 12. Cette mémoire supplémentaire, combinée à une interface plus large de 256 bits, offre une bande passante graphique culminant à 576 Go/s, contre 336 Go/s pour son homologue.
Le débat sur l'efficacité persiste concernant la puissance thermique de conception (TDP) des GPU. Le TDP de 170 watts du Flex Series 150 indique l'accent mis sur les performances, tandis que le TDP de 140 watts du Flex Series 75 souligne son penchant pour les applications économes en énergie et économes en énergie.
La prise en charge de technologies de pointe telles que Ray Tracing, oneAPI et OpenVINO est un fil conducteur commun aux deux modèles, garantissant une plate-forme de développement évolutive. Parallèlement, la compatibilité DirectX 12 Ultimate prend en charge des graphiques ultra-réalistes, un clin d'œil aux applications croisées potentielles dans des domaines tels que la visualisation professionnelle et les jeux en nuage.
La magie des GPU Intel Flex : SR-IOV
La magie du GPU Flex Series vient du SR-IOV. Si vous le connaissez, n'hésitez pas à passer à la section suivante ; Il n'y a rien de nouveau ici. Sinon, ou si vous avez besoin d’un rappel, attachez votre ceinture ; c'est des trucs sympas.
La virtualisation d'E/S à racine unique (SR-IOV) est une technologie qui améliore la gérabilité et l'efficacité des environnements virtualisés en permettant à un seul périphérique physique, tel qu'une carte d'interface réseau (NIC) ou une unité de traitement graphique (GPU), d'apparaître comme plusieurs appareils virtuels distincts. Ceci est particulièrement utile dans les centres de données pour améliorer les performances des machines virtuelles (VM) et maximiser l'utilisation des ressources matérielles sous-jacentes.
La technologie SR-IOV repose sur deux concepts fondamentaux : les fonctions physiques (PF) et les fonctions virtuelles (VF). Le PF est l'interface principale du périphérique physique et gère la fonctionnalité SR-IOV, y compris la création et la gestion des VF. Ces VF sont des versions allégées du PF, dotées des ressources nécessaires au déplacement des données mais avec des capacités de configuration réduites. Chaque VF peut être directement attribué à une VM, offrant ainsi un accès direct et hautes performances aux capacités du périphérique sans la surcharge typique des périphériques virtualisés.
Lorsque SR-IOV est utilisé avec des GPU, il permet à chaque machine virtuelle (VM) d'accéder directement à une partie des ressources du GPU. Cet accès direct est facilité grâce aux fonctions virtuelles (VF), qui sont des représentations légères du GPU qui peuvent être attribuées individuellement aux machines virtuelles.
Bien que les VF soient contrôlés par les machines virtuelles auxquelles ils sont attribués, le PF conserve le contrôle global, gérant les ressources et appliquant les politiques au niveau des appareils. Cette configuration est inestimable dans les scénarios où les performances et une faible latence sont critiques, comme dans les environnements VDI complexes, les tâches de calcul hautes performances et les services Web à grande échelle, améliorant considérablement l'efficacité opérationnelle des systèmes virtualisés.
Cette configuration permet aux machines virtuelles de contourner les méthodes traditionnelles de partage de ressources basées sur un hyperviseur, réduisant ainsi les frais généraux et améliorant les performances. Les tâches gourmandes en GPU, telles que le rendu 3D, le traitement vidéo ou les applications d'apprentissage automatique, peuvent entraîner une latence considérablement plus faible, une utilisation plus efficace des ressources GPU et une amélioration des performances globales dans les environnements virtualisés.
Les GPU Intel Flex offrent une accélération VDI « gratuite »
Le titre de la section n’est pas une blague. La série Intel Data Center GPU Flex arrive sur la scène des accélérateurs avec un avantage significatif dans les déploiements VDI : aucun coût de licence pour la configuration des configurations de GPU virtuel (vGPU). Tirant parti de la virtualisation GPU basée sur la virtualisation d'E/S à racine unique (SR-IOV) susmentionnée, cette série Intel Flex élimine les barrières financières traditionnelles associées au provisionnement de vGPU. L'absence de frais de licence réduit les coûts d'installation initiaux et diminue les dépenses opérationnelles continues, permettant ainsi des économies significatives à long terme.
