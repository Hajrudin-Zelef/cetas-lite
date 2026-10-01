---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7-2
title: "fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7.md
source_anchor: ""
source_lines: [33, 38]
sha256: d40a26638b642a362783e600c40e49b6c16121cc096d79750c6fc8cff1218b05
---

# fr-review-dell-poweredge-r770-review-modular-powerful-and-ai-ready-68c2acf7

S'inscrivant dans la lignée de la série R7x0, le R770 offre de nombreuses options de configuration pour répondre à divers besoins de déploiement. Une première significative pour cette gamme est le choix entre une configuration d'E/S arrière traditionnelle et une configuration d'E/S avant accessible en couloir froid, offrant une plus grande flexibilité pour s'adapter aux différentes configurations de centres de données et aux exigences de maintenance. Les options de stockage sont tout aussi polyvalentes, allant des nœuds de calcul avec un stockage local minimal, voire nul, aux configurations haute densité prenant en charge jusqu'à 40 disques E3.S pour les charges de travail centrées sur le stockage.
Pour répondre aux besoins croissants en calcul accéléré, notamment pour l'IA et le HPC, le R770 offre de solides capacités d'extension. Selon la configuration du châssis et de la carte d'extension, le serveur peut accueillir jusqu'à six cartes PCIe Gen 5 x16 pleine hauteur et pleine longueur (FHFL). De plus, il prend en charge l'installation de deux GPU double largeur, ce qui en fait une plateforme performante pour un large éventail de tâches. La flexibilité réseau est assurée par des emplacements mezzanine OCP 3.0, prenant en charge des cartes x8 ou x16 selon la configuration.
Dell a également apporté plusieurs améliorations de conception visant à améliorer la facilité d'entretien et la fiabilité. L'évolution de la carte Boot Optimized Storage Solution (BOSS) en est un parfait exemple. Auparavant relié par câbles et intégré à la carte PCIe, le contrôleur BOSS du R770 est désormais une carte standardisée OCP qui s'interface directement avec la carte mère, éliminant ainsi la complexité du câblage. Ce nouveau contrôleur BOSS intègre également des disques NVMe M.2 plus rapides et des dissipateurs thermiques pour garantir des températures de fonctionnement et des performances optimales pour les périphériques de démarrage. Autre amélioration subtile, mais pratique pour les techniciens : le remplacement des cavaliers traditionnels par des commutateurs DIP plus conviviaux pour des fonctions telles que l'effacement de la NVRAM.
Le changement architectural le plus profond est l'adoption complète de la norme OCP DC MHS. Dell avait déjà intégré des éléments OCP dans les générations précédentes, notamment en adoptant des emplacements pour cartes réseau OCP 3.0. Le R770 va encore plus loin. Les composants clés sont désormais conformes aux spécifications OCP, notamment les modules processeurs hôtes (HPM), communément appelés carte mère, qui incluent des composants tels que les emplacements pour cartes graphiques, désormais des connecteurs M-XIO. Le connecteur M-XIO offre une interface standardisée pour les cartes riser, améliorant ainsi la flexibilité et l'évolutivité. L'iDRAC est également implémenté comme OCP DC-SCM (Server Control Module).
De plus, le R770 intègre le nouveau connecteur d'alimentation PICPWR pour les connexions de périphériques tels que les GPU et les fonds de panier. Ce connecteur constitue une avancée significative, simplifiant l'alimentation et intégrant la surveillance de l'alimentation en ligne.
		
