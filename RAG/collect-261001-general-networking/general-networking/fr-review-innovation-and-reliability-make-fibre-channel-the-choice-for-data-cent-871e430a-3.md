---
id: collect-261001-general-networking/general-networking/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a-3
title: "fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a.md
source_anchor: ""
source_lines: [45, 67]
sha256: fbc462bff95f7f973f164dcb849ad4c3a35fb86e5e5dfb5b392fb2af3823a944
---

# fr-review-innovation-and-reliability-make-fibre-channel-the-choice-for-data-cent-871e430a

Fibre Channel est considérée comme la solution de connectivité de stockage la plus fiable du marché, avec une tradition d'améliorations progressives. Les technologies de serveur et de stockage poussent la demande pour une plus grande bande passante SAN. La capacité d'application et de stockage, les baies de stockage de 32 Go et 64 Go prenant en charge les SSD et NVMe, la virtualisation des serveurs et les déploiements multi-cloud prouvent la valeur des Fibre Channels car ils offrent un débit plus élevé, une latence plus faible et des vitesses de liaison plus élevées, le tout avec des performances prévisibles.
Marvell a récemment annoncé l'introduction de ses tout nouveaux HBA 64GFC. Il s'agit notamment de la série QLE2870 de HBA FC à un, deux et quatre ports qui doublent la bande passante disponible, fonctionnent sur un bus PCIe 4.0 plus rapide et prennent simultanément en charge FC et FC-NVMe, idéales pour les entreprises critiques à l'épreuve du temps. applications.
NVMe livre!
Il ne fait aucun doute que les périphériques NVMe offrent un accès en lecture et en écriture extrêmement rapide. La discussion porte donc sur la connexion de ces périphériques NVMe à des réseaux à haut débit sans considérer qu'il s'agit toujours de périphériques de stockage nécessitant une livraison garantie. La technologie qui a été développée comme méthode de livraison sans perte est Fibre Channel. Dans de nombreux tests effectués par des leaders de l'industrie, NVMe-oF et NVMe/FC ont obtenu de meilleurs résultats lorsque la technologie sous-jacente était Fibre Channel.
Les baies Flash permettent des performances de stockage de blocs plus rapides dans les charges de travail virtualisées à haute densité et réduisent le temps de réponse des applications gourmandes en données. Tout cela semble plutôt bien, sauf si l'infrastructure réseau ne peut pas fonctionner au même niveau que les baies de stockage flash.
Le stockage basé sur Flash exige une infrastructure déterministe à faible latence. D'autres architectures de réseau de stockage augmentent souvent la latence, créant des goulots d'étranglement et une congestion du réseau. Plus de paquets doivent être envoyés lorsque cela se produit, créant encore plus de congestion. Avec le contrôle de flux basé sur le crédit de Channel, les données peuvent être livrées aussi rapidement que le tampon de destination peut les recevoir, en supprimant les paquets ou en forçant les retransmissions.
Nous avons publié une analyse approfondie en début d'année de l'approche FC-NVMe de Marvell. Pour en savoir plus, consultez l'article « Marvell renforce son engagement envers la technologie FC-NVMe ».
Fonctionnalités principales de stockage NVMe introduites pour la première fois dans vSphere 7.0
Un blog VMware a décrit NVMe over Fabrics (NVMe-oF) comme une spécification de protocole qui connecte les hôtes au stockage flash à haut débit via des structures réseau utilisant le protocole NVMe. VMware a introduit NVMe-oF dans vSphere 7 U1. Le blog VMware a indiqué que les résultats de référence ont montré que Fibre Channel (FC-NVMe) surpassait systématiquement le SCSI FCP dans les environnements virtualisés vSphere, offrant un débit plus élevé et une latence plus faible. La prise en charge de NVMe sur TCP/IP a été ajoutée à vSphere 7.0 U3.
Sur la base de l'adoption croissante de NVMe, VMware a ajouté la prise en charge du stockage NVMe partagé à l'aide de NVMe-oF. Compte tenu de la faible latence inhérente et du débit élevé, les industries tirent parti de NVMe pour les charges de travail d'IA, de ML et d'informatique. En règle générale, NVMe utilisait un bus PCIe local, ce qui rendait difficile la connexion à une baie externe. À l'époque, l'industrie avait proposé des options de connectivité externe pour NVMe-oF basées sur IP et FC.
Dans vSphere 7, VMware a ajouté la prise en charge du stockage NVMe partagé à l'aide de NVMe-oF avec NVMe sur FC et NVMe sur RDMA.
Les matrices continuent d'offrir des vitesses plus élevées tout en maintenant la livraison garantie sans perte requise par les réseaux de stockage. vSphere 8.0 prend en charge 64GFC, la vitesse FC la plus rapide à ce jour.
vSphere 8.0 fait progresser son objectif NVMe, en donnant la priorité à Fibre Channel
Les vVols ont été l'objectif principal de l'ingénierie du stockage VMware au cours des dernières versions, et avec vSphere 8.0, le stockage central a ajouté la prise en charge des vVols dans NVMe-oF. Dans un premier temps, VMware ne prendra en charge que FC mais continuera à valider et à prendre en charge d'autres protocoles NVMe-oF.
Il s'agit d'une nouvelle spécification vVols, pour le framework VASA/VC. Pour en savoir plus, consultez la documentation VASA 4.0/vVols 3.0 .
vSphere 8 continue d'ajouter des fonctionnalités et des améliorations et a récemment augmenté les espaces de noms pris en charge à 256 et les chemins à 2K pour NVMe-FC et TCP. Une autre fonctionnalité, la prise en charge des commandes de réservation pour les périphériques NVMe, a été ajoutée à vSphere. Les commandes de réservation permettent aux clients d'utiliser la capacité VMDK en cluster avec Microsoft WSFC avec les banques de données NVMe-oF.
Simple à configurer et simple à gérer !
Fibre Channel a une autre efficacité intégrée : la découverte automatique. Lorsqu'un périphérique FC est connecté au réseau, il est automatiquement découvert et ajouté à la matrice s'il dispose des informations d'identification nécessaires. La carte des nœuds est mise à jour et le trafic peut traverser la fibre. Il s'agit d'un processus simple sans intervention d'un administrateur.
Il y a plus de frais généraux lors de la mise en œuvre de NVMe/TCP. Étant donné que NVMe/TCP ne dispose pas d'un mécanisme de découverte automatique, ESXi a ajouté la prise en charge du service de découverte NVMe. La prise en charge avancée du service de découverte NVMe-oF dans ESXi permet la découverte dynamique du service de découverte NVMe conforme aux normes. ESXi utilise le service mDNS/DNS-SD pour obtenir des informations telles que l'adresse IP et le numéro de port des services de découverte NVMe-oF actifs sur le réseau. ESXi envoie une requête DNS multidiffusion (mDNS) demandant des informations aux entités fournissant le service de découverte (NVMe) (DNS-SD). Si une telle entité est active sur le réseau (sur lequel la requête a été envoyée), elle enverra une réponse (monodiffusion) à l'hôte avec les informations demandées, c'est-à-dire l'adresse IP et le numéro de port où le service est exécuté.
Conclusion
Fibre Channel a été spécialement conçu pour transporter le stockage de blocs et, comme indiqué dans cet article, il s'agit d'une structure fiable, à faible latence, sans perte et hautes performances. Pour être clair, des progrès sont faits pour améliorer l'utilisation de TCP pour le trafic de stockage sur certains réseaux à haut débit importants. Mais le fait demeure que TCP n'est pas un réseau sans perte, et la retransmission des données est toujours un problème.
Famille de produits Marvell FC
Ce rapport est parrainé par Marvell. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
