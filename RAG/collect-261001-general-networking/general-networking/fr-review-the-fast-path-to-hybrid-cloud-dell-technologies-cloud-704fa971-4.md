---
id: collect-261001-general-networking/general-networking/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971-4
title: "fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971.md
source_anchor: ""
source_lines: [57, 66]
sha256: 80db4b0b83d4e9f19e1eee6893e44f8106f7f18d21e39820dfe170289a0838a7
---

# fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971

Une étape cruciale lors de la mise à niveau du logiciel consiste à s'assurer que le matériel sous-jacent est compatible avec la nouvelle version des composants VMware SDDC. Une fois que nous avons vérifié et téléchargé les bits de mise à niveau, nous avons ensuite vérifié la compatibilité de chaque hôte. Il s'agissait d'un processus manuel consistant à se connecter à l'iDRAC des 8 hôtes et à vérifier les révisions du micrologiciel sur chacun des lecteurs, contrôleurs de stockage et adaptateurs de carte d'interface réseau installés. Une fois que nous avons eu notre liste, nous avons ensuite dû accéder au site VMware Hardware Compatibility List et rechercher et vérifier manuellement chacun des composants.
Pour effectuer la mise à niveau réelle des composants VMware SDDC pour passer de VVD 5.1.1 à 5.1.2, les composants logiciels suivants sont mis à niveau :
- Dispositifs de contrôleur de service de plate-forme
- Appareils vCenter Server
- Service de téléchargement de vSphere Update Manager
- Hôtes ESXi
La mise à jour des clusters de gestion et de charge de travail a duré 12 h 59 min 16 s. Seuls les composants logiciels VMware ont été mis à jour lors de ce test LCM.
Ce rapport est parrainé par Dell Technologies. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
[1] Offre valable pour certaines solutions préconfigurées. Contactez votre représentant commercial pour plus d'informations. Offre non valable pour les commandes de plus de 1 000 instances, le stockage hybride, certains composants vRealize (vRA, vRO) et certaines autres fonctionnalités. L'approbation de crédit du client, l'étude de site et le cahier des charges de configuration doivent être complétés avant la validation de la commande. La disponibilité des produits, les délais de livraison, les jours fériés et d'autres facteurs peuvent impacter le délai de déploiement. Le déploiement comprend la livraison, l'installation standardisée et la configuration matérielle et logicielle. Offre valable uniquement aux États-Unis, au Royaume-Uni, en France et en Allemagne.
[2] D'après une analyse interne de Dell Technologies, novembre 2019
