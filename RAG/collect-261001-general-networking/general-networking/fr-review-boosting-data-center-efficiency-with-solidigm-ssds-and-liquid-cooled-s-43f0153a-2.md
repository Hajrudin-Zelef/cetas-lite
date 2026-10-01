---
id: collect-261001-general-networking/general-networking/fr-review-boosting-data-center-efficiency-with-solidigm-ssds-and-liquid-cooled-s-43f0153a-2
title: "fr-review-boosting-data-center-efficiency-with-solidigm-ssds-and-liquid-cooled-s-43f0153a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-general-networking/fr-review-boosting-data-center-efficiency-with-solidigm-ssds-and-liquid-cooled-s-43f0153a.md
source_anchor: ""
source_lines: [16, 36]
sha256: 2aafc42341f255cbb12e78c99d21d41b49e4e50d3276a9c06f5af00847f2de19
---

# fr-review-boosting-data-center-efficiency-with-solidigm-ssds-and-liquid-cooled-s-43f0153a

Nous avons utilisé la surveillance de l’alimentation embarquée de Dell à l’intérieur du système de gestion embarqué iDRAC9 du serveur pour surveiller l’alimentation au niveau du système.
Nous nous sommes concentrés sur les charges de travail de bande passante de lecture et d'écriture séquentielles, en utilisant une taille de bloc de 128 Ko sur chaque lecteur, puis nous avons mesuré les performances globales sur les 24 SSD. Il convient de noter que cette configuration particulière du Dell PowerEdge R760 avec 24 baies NVMe exploite un commutateur PCIe par rapport aux baies NVMe à connexion directe. Ainsi, la bande passante totale mesurée sature les voies de commutation PCIe disponibles avant d'atteindre les lecteurs. Cela a un impact sur les performances de lecture totales que nous avons mesurées par rapport à la fiche technique du Soldigim P5536, mais les vitesses d'écriture globales étaient toutes inférieures à cette limite.
|  | Puissance totale | Vitesse d'écriture | Lire Go/s | Watts Au dessus de la base | Watts/entraînement (avec surcharge du système) | 
|---|---|---|---|---|---|
| Pas de lecteur au ralenti | 462 | - | - | - | - | 
| Disques de ralenti installés | 594 | - | - | 132 | 5.5 | 
| Lecture séquentielle 24x PS0 | 858 | - | 109GB / s | 396 | 16.5 | 
| Lecture séquentielle 24x PS1 | 858 | - | 105GB / s | 396 | 16.5 | 
| Lecture séquentielle 24x PS2 | 759 | - | 79.8GB / s | 297 | 12.375 | 
| Écriture séquentielle 24x PS0 | 1089 | 82.5GB / s | - | 627 | 26.125 | 
| Écriture séquentielle 24x PS1 | 825 | 34.4GB / s | - | 363 | 15.125 | 
| Écriture séquentielle 24x PS2 | 726 | 17.3GB / s | - | 264 | 11 | 
En revenant sur notre article sur les avantages de la conversion d’une plateforme refroidie par air en refroidissement liquide direct, nous avons constaté une légère augmentation des performances des processeurs, mais nous avons également économisé 200 W d’énergie. L’énergie est une denrée précieuse dans la nouvelle vague de serveurs centrés sur l’IA qui consacrent souvent toutes les ressources disponibles aux GPU et aux processeurs haut de gamme. Dans un centre de données dont le budget énergétique est limité ou proche de la limite du refroidissement par air, le passage au DLC permet d’acheter un budget énergétique qui permet au serveur d’être rempli de plus de SSD pour la même empreinte énergétique qu’un serveur refroidi par air.
Une économie d’énergie de 200 W peut s’avérer très utile en termes de densité de stockage. Cette économie vous permet de doubler l’empreinte de stockage de 12 à 24 SSD dans un serveur refroidi par liquide par rapport à un serveur refroidi par air si vous avez des charges de travail orientées vers des charges de travail à lecture intensive. Avec le Solidigm D5-P5336, ce serveur à 24 baies a augmenté sa capacité de stockage de 737 To à 1,474 24 To grâce à la boucle liquide. Si la charge de travail est lourde en écriture, vous pourrez équiper le serveur d’environ huit SSD supplémentaires. Cependant, ces chiffres concernent les modes d’alimentation de base, donc si vous êtes prêt à réduire les performances d’écriture du haut de gamme, vous pouvez facilement équiper votre serveur de XNUMX SSD avec une charge de travail lourde en écriture avec des performances réduites.
Conclusion
Grâce à nos tests des SSD Solidigm D5-P5336, nous avons pu constater que la gestion des états d'alimentation NVMe peut avoir un impact significatif sur l'efficacité énergétique sans affecter considérablement les performances. Les opérateurs de centres de données qui cherchent à maximiser l'efficacité énergétique peuvent exploiter ces états d'alimentation pour atteindre une plus grande densité de stockage ou réduire les coûts opérationnels, en particulier dans les environnements centrés sur l'IA où l'énergie est primordiale. Les SSD haute densité de Solidigm sont bien placés pour cela, offrant une excellente efficacité téraoctet par watt, en particulier avec les technologies modernes de refroidissement liquide.
Nos résultats révèlent que même de légers ajustements des états d'alimentation peuvent générer des économies d'énergie significatives, ce qui peut s'avérer crucial dans les environnements limités par la disponibilité de l'énergie. L'optimisation de la consommation électrique globale des serveurs améliore la densité de stockage et favorise des opérations de centre de données plus durables.
La gestion de l'alimentation devient de plus en plus critique à mesure que les serveurs modernes sont poussés à leurs limites, en particulier dans les charges de travail pilotées par l'IA. L'association du refroidissement liquide et des options de gestion efficaces des SSD offre une voie à suivre pour les centres de données qui cherchent à faire évoluer les performances et la densité de stockage sans dépasser les budgets énergétiques.
Vous pourrez voir la démonstration complète de ces technologies en direct à l'OCP 2024. Nous montrerons comment le refroidissement liquide et les SSD de Solidigm peuvent être les pierres angulaires de l'efficacité énergétique dans le centre de données moderne.
Solutions de stockage Solidigm
Ce rapport est sponsorisé par Solidigm. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
