---
id: collect-261001-general-networking/general-networking/fr-review-dell-perc13-70dd67e6-1
title: "fr-review-dell-perc13-70dd67e6"
domain: general-networking
role: reference
task: reference
actors: ["Broadcom"]
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-general-networking/fr-review-dell-perc13-70dd67e6.md
source_anchor: ""
source_lines: [1, 43]
sha256: e07dc7afd99e1584547eb35d6d40f84b9ac6a07990047ee834ddd2900c9851f0
---

# fr-review-dell-perc13-70dd67e6

Le contrôleur RAID Dell H975i, issu de la série PERC13, représente la plus importante avancée de l'entreprise en matière de RAID matériel depuis plus de dix ans. Si Dell a régulièrement mis à jour sa gamme PERC, celles-ci étaient essentiellement incrémentielles, axées sur le réglage des contrôleurs et l'amélioration de la bande passante au fur et à mesure de l'évolution des générations PCIe. Cependant, l'architecture sous-jacente restait liée aux technologies SATA et SAS, qui ont défini le RAID d'entreprise pendant des années. Le PERC H975i rompt définitivement ce cycle. Basé sur la gamme de chipsets SAS51xx de Broadcom, ce contrôleur marque une transition décisive vers une conception axée sur le flash et le NVMe natif. En prenant exclusivement en charge les disques NVMe et en éliminant la prise en charge des technologies HDD et SATA traditionnelles, le H975i intègre une approche avant-gardiste de l'infrastructure de stockage, optimisée pour les exigences de hautes performances et de faible latence des charges de travail modernes, gourmandes en données et axées sur l'IA.
Points clés à retenir
- RAID NVMe Flash-first : Le PERC13 H975i s'éloigne entièrement de SAS/SATA, construit sur Broadcom SAS51xx pour une architecture native NVMe et prête pour l'IA.
- Grand saut générationnel : Le PCIe Gen5 x16 avec jusqu'à 16 disques NVMe par contrôleur (32 avec deux) a fourni 52.5 Go/s et 12.5 M IOPS par contrôleur lors des tests, avec des gains par rapport au PERC12, notamment +88 % de bande passante en lecture, +318 % de bande passante en écriture, +31 % d'IOPS en lecture 4K et +466 % d'IOPS en écriture 4K.
- Ajustement du serveur d'IA : La conception intégrée à l'avant libère les emplacements PCIe arrière pour les GPU, raccourcit les exécutions MCIO et permet un canal de stockage dédié par accélérateur pour un débit plus stable et plus déterministe sans surcharge du processeur.
- Résilience sous stress : Le cache protégé par un supercondensateur et les reconstructions plus rapides réduisent le temps jusqu'à 10 min/Tio tout en maintenant des performances élevées pendant les reconstructions (jusqu'à 53.7 Go/s en lecture, 68 Go/s en écriture, 17.3 M/5.33 M 4K IOPS).
- Sécurité de bout en bout : Racine de confiance matérielle, identité de périphérique SPDM et cryptage à spectre complet qui couvre les lecteurs, les données en vol et le cache du contrôleur.
Le PERC H975i offre des performances et des innovations architecturales inégalées. S'appuyant sur une interface hôte PCIe Gen 5 x16 et prenant en charge jusqu'à 16 disques NVMe (32 disques NVMe par système avec deux contrôleurs), le H975i a atteint lors de nos tests un débit maximal remarquable de 52.5 Go/s et 12.5 millions d'IOPS par contrôleur. Cela représente une multiplication par près de deux des performances dans toutes les catégories clés par rapport au PERC2, qui atteignait 12 millions d'IOPS et 6.9 Go/s de débit. Au-delà des performances brutes, le PERC27 intègre un mécanisme de protection du cache par supercondensateur (remplaçant les systèmes traditionnels alimentés par batterie), garantissant l'intégrité des données sans compromettre la fiabilité opérationnelle. S'appuyant sur les fonctionnalités de sécurité de son prédécesseur, le H13i étend désormais les capacités de chiffrement à spectre complet, chiffrant les données dans le cache et offrant une protection complète, en transit comme au repos.
Le PERC H975i se présente comme un accélérateur de stockage spécialement conçu pour répondre aux exigences de calcul sans précédent des charges de travail d'IA. Il offre à la fois une densité et des performances élevées, ainsi qu'un stockage à faible latence sans surcharge CPU. En pratique, l'association d'une carte RAID PCIe Gen5 capable de saturer une interface x16 avec un GPU Gen5 confère à chaque accélérateur son propre pipeline de stockage dédié. Cela simplifie la topologie PCIe/NUMA, prévient les effets de voisinage bruyant et isole les tâches de reconstruction ou d'arrière-plan du domaine d'E/S du GPU concerné.
En adaptant cette architecture à deux cartes RAID et deux GPU, vous préservez des performances linéaires tout en évitant les conflits sur les voies partagées ou les caches. Il en résulte une bande passante d'entrée plus stable pour l'entraînement et l'inférence gourmands en données (lots volumineux, remaniements rapides, lectures rapides des points de contrôle), avec des distributions de latence plus serrées sous charge et lors des reconstructions. Cette architecture ne se contente pas d'augmenter les valeurs de pointe ; elle rend le débit plus déterministe, ce qui est précisément ce dont les serveurs d'IA multi-GPU ont besoin pour maintenir une utilisation élevée.
Spécifications des Dell PERC12 H965i et PERC13 H975i
| Fonctionnalité | PERC12 H965i Avant | PERC13 H975i Avant | 
|---|---|---|
| Niveaux de RAID | 0, 1, 5, 6, 10, 50, 60 | 0, 1, 5, 6, 10, 50, 60 | 
| Non RAID (JBOD) | Oui | Oui | 
| Type de bus hôte | PCIe Gen4x16 | PCIe Gen5x16 | 
| Gestion des bandes latérales | I2C, PCIe VDM | I2C, PCIe VDM | 
| Boîtiers par port | N'est pas applicable | N'est pas applicable | 
| Processeur/jeu de puces | RAID sur puce Broadcom, SAS4116W | RAID sur puce Broadcom, SAS5132W | 
| Pack d'énergie / Alimentation de secours | expert | Supercondensateur | 
| Sécurité de la gestion des clés locales | Oui | Oui | 
| Gestionnaire de clés d'entreprise sécurisé | Oui | Oui | 
| Profondeur de la file d'attente du contrôleur | 8,192 | 8,192 | 
| Cache non volatile | Oui | Oui | 
| Mémoire cache | 8 Go DDR4 3200 MT/s | Cache RAID intégré | 
| Fonctions de cache | Écriture différée, lecture anticipée, écriture directe, toujours écriture différée, pas de lecture anticipée | Écriture différée, écriture directe, toujours écriture différée, pas de lecture anticipée | 
| Disques virtuels complexes maximum | 64 | 16 | 
| Nombre maximal de disques virtuels simples | 240 | 64 | 
| Nombre maximal de groupes de disques | 64 | 32 | 
| Nombre maximal de disques virtuels par groupe de disques | 16 | 8 | 
| Nombre maximal de périphériques de secours | 64 | 8 | 
| Périphériques remplaçables à chaud pris en charge | Oui | Oui | 
| Configuration automatique (principal et exécution unique) | Oui | Oui | 
| Moteur XOR matériel | Oui | Oui | 
| Extension de capacité en ligne | Oui | Oui | 
| Disque de secours dédié et global | Oui | Oui | 
| Types de lecteur pris en charge | NVMe Gen3 et Gen4 | NVMe Gen3, Gen4 et Gen5 | 
| Taille de l'élément de bande VD | 64KB | 64KB | 
| Prise en charge NVMe PCIe | Gen4 | Gen5 | 
| Configuration maximale des disques NVMe | 8 lecteurs par contrôleur | 16 lecteurs par contrôleur | 
| Tailles de secteur prises en charge | 512B, 512e, 4Kn | 512B, 512e, 4Kn | 
| Prise en charge du démarrage du stockage | UEFI uniquement | UEFI uniquement | 
Le contrôleur frontal PERC13 H975i des serveurs Dell PowerEdge est conçu pour une intégration transparente à l'architecture système. Contrairement aux cartes d'extension traditionnelles qui occupent les emplacements PCIe arrière, le H975i se connecte directement au fond de panier du disque avant et s'interface avec les connecteurs MCIO avant de la carte mère via des interfaces PCIe 5.0 dédiées. Cette conception intégrée préserve les emplacements PCIe arrière pour les GPU hautes performances et les extensions PCIe supplémentaires, tout en réduisant considérablement la longueur des câbles. Cela contribue à préserver l'intégrité du signal, rendant le système plus fiable et plus facile à entretenir. Il en résulte une configuration interne plus épurée et une meilleure circulation de l'air pour les déploiements denses et gourmands en ressources de calcul.
