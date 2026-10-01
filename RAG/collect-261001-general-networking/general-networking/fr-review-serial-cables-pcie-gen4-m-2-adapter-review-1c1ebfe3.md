---
id: collect-261001-general-networking/general-networking/fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3
title: "fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["amd", "arr"]
source: docs/RAG/collect-261001-general-networking/fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3.md
source_anchor: ""
source_lines: [1, 13]
sha256: c4b9c71a9177d847c778381bb8a0262de545ce4dbe5e3f9fcc442286c6c8cacf
---

# fr-review-serial-cables-pcie-gen4-m-2-adapter-review-1c1ebfe3

L'adaptateur Serial Cables PCIe Gen4 m.2 (PCI4-AD-x4M2-04-G4) permet aux laboratoires de test et aux utilisateurs d'ajouter un slot Gen4 supplémentaire PCIe x4 aux ordinateurs de bureau et serveurs avec une carte mère prise en charge. Traditionnellement, les adaptateurs m.2 offrent aux entreprises une alternative de lecteur de démarrage économique pour les serveurs (sans occuper un emplacement de montage avant) ; cependant, cet adaptateur hautement personnalisable peut être utilisé pour bien plus que cela.
L'adaptateur Serial Cables Gen4 est compatible avec pratiquement toutes les tailles de SSD m.2, car il comporte des trous plaqués de facteur de forme 2230 (30 mm), 2242, 2260, 2280 et pleine longueur 22110 (110 mm). À l'heure actuelle, la plupart des SSD Gen4 m.2 sont disponibles dans une taille de facteur de forme (2280), mais une gamme de supports permet aux utilisateurs d'ajouter d'autres cartes non Gen4 de différentes tailles. Bien que la carte adaptateur soit rétrocompatible avec les disques Gen3, vous ne verrez bien sûr que les vitesses Gen3. Cela vaut également pour la carte mère hôte. Si vous ajoutez cette carte sur une carte qui ne prend pas en charge Gen4, le lecteur sera limité à la bande passante Gen3 de l'emplacement PCIe.
En ce qui concerne la commutation des disques, ce type est unique et polyvalent. Là où la plupart des cartes grand public utilisent une vis pour maintenir un SSD m.2 pour des moyens plus "permanents", cette carte comporte différentes formes d'onglets à dégagement rapide. L'une de ces méthodes est une petite fonction à ressort qui met facilement en place un lecteur m.2. L'autre est une languette en plastique sur un point de pivot (en forme de médiator); installez simplement le lecteur dans l'emplacement m.2, appuyez légèrement dessus, puis faites glisser la languette en plastique sur le lecteur pour le maintenir en place. L'élimination de l'exigence d'un tournevis sera certainement bien accueillie par ceux qui ont des cas d'utilisation qui impliquent de changer constamment de lecteur, car il peut être fastidieux et ennuyeux d'installer des lecteurs m.2 sur des cartes adaptateurs traditionnelles.
En haut de l'adaptateur, vous verrez une gamme de broches familières, notamment CLKREQ # (signal de demande d'horloge), WAKE # (fonctionnalité de réveil) et PREST # (utilisé pour spécifier les paramètres de tension d'alimentation). Vous remarquerez également un interrupteur tactile (c'est-à-dire un bouton qui démarre ou arrête un flux de courant le long d'un circuit) situé en bas à gauche du PCB. Cela n'a aucune fonction pour le moment, car Serial Cables a ajouté ce bouton pour un client spécifique qui faisait du développement SSD. Recherchez les mises à jour du micrologiciel pour activer cette fonctionnalité à l'avenir.
Performances
Pour démontrer les performances de l'adaptateur Serial Cables PCIe Gen4 m.2, nous avons utilisé la Lenovo P620 (une station de travail puissante dotée d'un processeur AMD Threadripper PRO avec prise en charge PCIe Gen4) et exécuté le test de performance Blackmagic avec les configurations suivantes :
- équipé l'adaptateur d'un SSD Samsung 980 Pro PCIe Gen4 et installé sur l'un des slots PCIe P620 de la station de travail
- installé le Samsung 980 Pro directement sur l'emplacement natif à l'intérieur du poste de travail
L'objectif était de montrer que l'adaptateur respecte et ou dépasse l'emplacement intégré d'un poste de travail de niveau 1. En utilisant l'adaptateur de câbles série à l'intérieur du P620, le 980 Pro a atteint 5.29 Go/s en lecture et 4.36 Go/s en écriture.
Ces résultats étaient pratiquement identiques aux vitesses enregistrées lorsque le disque était installé directement sur la carte, puisque le Samsung Pro affichait 5.28 Go/s en lecture et 4.34 Go/s en écriture.
Conclusion
Bien que nous utilisions simplement l'adaptateur Serial Cables PCIe Gen4 m.2 comme pass-through pour les SSD M.2 NVMe vers un système hôte qui prend en charge Gen4, vous pouvez faire beaucoup plus avec lui dans un laboratoire de test ou un laboratoire d'ingénierie qui travaille avec beaucoup de lecteurs. En tant que tel, l'adaptateur Serial Cables n'est certainement pas pour tout le monde, mais il est bien mieux construit que les adaptateurs ordinaires de 15 m.2 que vous pourriez trouver chez votre détaillant en ligne préféré.
Il s'agit d'une carte adaptateur extrêmement niche qui peut être configurée pour accomplir une gamme de tâches hautement techniques. Couplant tout cela avec ses performances qui ont dépassé les vitesses embarquées, il n'y a pas beaucoup de cartes comme celle-ci.
