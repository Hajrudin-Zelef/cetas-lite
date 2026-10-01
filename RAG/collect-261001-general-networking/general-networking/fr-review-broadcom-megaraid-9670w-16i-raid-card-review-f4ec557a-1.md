---
id: collect-261001-general-networking/general-networking/fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a-1
title: "fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Broadcom"]
dates: []
keywords: ["amd"]
source: docs/RAG/collect-261001-general-networking/fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a.md
source_anchor: ""
source_lines: [1, 32]
sha256: 86b2418c223d95d1cbd8e7e8c37c49884d6efbac5bccf3b37d86b2f4751acce4
---

# fr-review-broadcom-megaraid-9670w-16i-raid-card-review-f4ec557a

La série MegaRAID 9600 est un adaptateur de stockage de troisième génération qui prend en charge les disques SATA, SAS et NVMe, conçu pour offrir les meilleures performances et disponibilité des données possibles pour les serveurs de stockage. Par rapport à la génération précédente, la série 9600 offre une augmentation de la bande passante par 2, une augmentation des IOP par 4, une réduction de la latence d'écriture par 25 et une augmentation des performances par 60 lors des reconstructions. Comme d'habitude, la famille de cartes 9600 comprend un certain nombre de configurations. Dans cette revue, nous examinons le Broadcom MegaRAID 9670W-16i, qui prend en charge 16 ports internes.
Le MegaRAID 9670W-16i est basé sur le SAS4116W RAID-on-Chip (RoC), qui est un facteur clé dans l'amélioration des performances à tous les niveaux. Les utilisateurs peuvent connecter jusqu'à 240 appareils SAS/SATA ou 32 appareils NVMe par contrôleur avec l'interface x16 PCIe Gen 4.0.
La série 9600 offre également une prise en charge du démarrage sécurisé matériel et de l'attestation SPDM, une protection de l'équilibre et des performances pour RAID 0, 1, 5, 6, 10, 50 et 60 et JBOD. La protection du cache flash CacheVault est une option pour ceux qui veulent plus de protection.
Pourquoi un RAID matériel pour NVMe ?
Le RAID matériel a été l'option incontournable pour offrir un stockage résilient bien avant que la barbe des administrateurs informatiques d'aujourd'hui ne soit grise. Mais comme le stockage est devenu plus rapide, en particulier les SSD NVMe, les cartes RAID ont eu du mal à suivre le rythme. Ainsi, lorsque Broadcom nous a présenté une critique du MegaRAID 9670W-16i, nous étions un peu dubitatifs. La vérité est qu'il y a un coût de performance pour une carte RAID, et nous les avons évités pendant de nombreuses années en conséquence. Cela dit, la valeur de ce qu'offre le RAID matériel est indéniable.
Pour les environnements qui n'offrent pas d'options RAID logicielles, notamment VMware ESXi, les clients ne peuvent pas facilement agréger le stockage ou protéger le stockage avec RAID. Bien que vSAN puisse être facilement implémenté au niveau du cluster, il ne peut pas être utilisé pour un nœud ESXi autonome à la périphérie. Ici, les clients peuvent vouloir les avantages de rassembler plusieurs SSD en RAID pour une plus grande banque de données ou une certaine résilience des données.
Même sous Windows, qui propose des espaces de stockage pour des serveurs individuels, certains types de RAID logiciels, tels que RAID5/6, sont durement touchés. Dans le passé, le RAID matériel était une solution efficace pour combler l'écart de performances des périphériques SAS et SATA, le MegaRAID 9670W vise à changer cela pour les périphériques NVMe.
Banc d'essai Broadcom MegaRAID 9670W-16i
Pour cet examen, nous avons travaillé avec Micron, Supermicro et Broadcom pour assembler un banc d'essai. Le serveur est un système Supermicro AS-1114S-WN10RT 1U avec processeur AMD Milan 7643 et 128 Go de DDR4. À l'intérieur de ce système se trouve le 9670W-16i, câblé à deux JBOD NVMe à 8 baies. À l'intérieur de chacun se trouvent 8 SSD Micron 7450 formatés à une capacité de 6.4 To.
Pour mesurer les performances des disques via l'adaptateur MegaRAID 9670W-16i, les tests de performances ont été répartis dans les configurations suivantes. Le premier comprenait une configuration JBOD mesurant chaque disque en dehors des configurations RAID (mais toujours via le HBA), RAID10, puis RAID5. Ces configurations ont été parcourues à travers un processus de script pour préconditionner le flash, exécuter les tests pour lesquels elles ont été conditionnées et passer au mélange de préconditionnement/charge de travail suivant. Ce processus, au total, mesurait environ 16 heures.
- Durée totale du test d'environ 16 heures dans cet ordre :
- Préconditionnement séquentiel (~2:15)
- Tests séquentiels sur 16x JBOD, 2x 8DR10, 2x 8DR5 (~2 heures)
- Préconditionnement aléatoire – 2 parties (~4:30)
- Tests optimaux aléatoires sur 16x JBOD, 2x 8DR10, 2x 8DR5 (~3 heures)
- Tests de reconstruction aléatoires sur 1x 16DR10, 1x 16DR5 (~2:30)
- Latence d'écriture aléatoire pour une reconstruction optimale pour 1x 16DR5 (~1:40)
La première section des mesures de performances se concentre sur la bande passante via la carte en modes JBOD, RAID10 et RAID5. Avec le MegaRAID 9670W-16i offrant une largeur d'emplacement x16 PCIe Gen4, ses performances de pointe seront d'environ 28 Go/s dans une direction, et c'est là que l'emplacement Gen4 culmine. En comparaison, un SSD U.2 Gen4 se connecte via une connexion x4 et peut culminer à environ 7 Go/s, et c'est là que la plupart des disques d'entreprise peuvent dépasser les charges de travail de lecture.
Cela dit, le MegaRAID 9670W sature complètement l'emplacement auquel il est connecté. En ce qui concerne les performances de lecture, la configuration JBOD arrive avec 28.3 Go/s avec RAID10, et RAID5 arrive juste en dessous, mesurant 28 Go/s. Lorsque nous nous concentrons sur les performances d'écriture, la ligne de base JBOD est de 26.7 Go/s, tandis que la configuration RAID10 est arrivée avec 10.1 Go/s et RAID5 à 13.2 Go/s. Lorsque nous examinons une répartition 50/50 du trafic de lecture et d'écriture simultanés, la configuration JBOD mesurait 41.6 Go/s, RAID10 à 19.6 Go/s et RAID5 à 25.8 Go/s.
| Charge de travail | JBOD (Mo/s) | RAID 10 – Optimal (Mo/s) | RAID 5 – Optimal (Mo/s) | 
|---|---|---|---|
| Lectures séquentielles maximales | 28,314 | 28,061 | 28,061 | 
| Écritures séquentielles maximales | 26,673 | 10,137 | 13,218 | 
| Maximum 50:50 lectures séquentielles : écritures | 41,607 | 19,639 | 25,833 | 
Lorsque nous nous concentrons sur les transferts aléatoires en petits blocs, nous constatons que le MegaRAID 9670W résiste assez bien en termes de performances de lecture par rapport au chiffre de référence JBOD de 7 millions d'IOPS. Cette vitesse a chuté à environ la moitié (3.2 millions d'IOPS) lors d'une opération de reconstruction si un SSD tombait en panne dans le groupe RAID. En ce qui concerne les performances d'écriture aléatoire, la ligne de base JBOD mesurait 6.3 millions d'IOPS contre 2.2 millions de RAID10 et 1 million de RAID5. Ces chiffres n'ont pas connu de baisse considérable lorsqu'un SSD a été défaillant du groupe et que la carte RAID a été forcée de se reconstruire. Dans cette situation, RAID10 n'a pas changé, bien que RAID5 soit passé de 1M à 788k IOPS.
Dans la charge de travail 4K OLTP avec un mélange de performances de lecture et d'écriture, la ligne de base JBOD mesurait 7.8 millions d'IOPS contre RAID10 avec 5.6 millions d'IOPS et RAID5 avec 2.8 millions d'IOPS. Lors d'une reconstruction, RAID10 est passé de 5.6 M à 2.4 M IOPS, et RAID5 est passé de 2.8 M à 1.8 M IOPS.
| Charge de travail | JBOD | RAID 10 – Optimal | RAID 5 – Optimal | RAID 10 – Reconstruction | RAID 5 – Reconstruction | 
|---|---|---|---|---|---|
| Lectures aléatoires de 4 Ko (IOP) | 7,017,041 | 7,006,027 | 6,991,181 | 3,312,304 | 3,250,371 | 
| Écritures aléatoires de 4 Ko (IOP) | 6,263,549 | 2,167,101 | 1,001,826 | 2,182,173 | 788,085 | 
| OLTP 4 Ko (IOP) | 7,780,295 | 5,614,088 | 2,765,867 | 2,376,036 | 1,786,743 | 
Un autre aspect important des performances RAID est le comportement du stockage entre les conditions optimales et les performances de reconstruction en cas de panne d'un disque. Si les performances ou la latence devaient subir un coup dur, la réactivité des applications peut devenir un problème. À cette fin, nous nous sommes concentrés sur la latence d'écriture aléatoire RAID5 4K en modes optimal et de reconstruction. Sur l'ensemble du spectre, la latence est restée assez similaire, ce qui est exactement ce que vous voulez voir dans un système de stockage d'environnement de production.
