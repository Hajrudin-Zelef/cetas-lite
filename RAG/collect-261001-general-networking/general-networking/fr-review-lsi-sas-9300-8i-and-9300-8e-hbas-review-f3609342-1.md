---
id: collect-261001-general-networking/general-networking/fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342-1
title: "fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Microsoft", "Samsung"]
dates: []
keywords: ["benchmark", "benchmarks", "intel", "nand"]
source: docs/RAG/collect-261001-general-networking/fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342.md
source_anchor: ""
source_lines: [1, 40]
sha256: 1f4bd57780984eee5919883bde97bc21821c5e6f5ad9a089dbe2928156abb064
---

# fr-review-lsi-sas-9300-8i-and-9300-8e-hbas-review-f3609342

Les HBA LSI SAS 9300 sont spécialement conçus pour des performances élevées dans les serveurs haut de gamme qui se connectent à des boîtiers de stockage à grande échelle ainsi que pour la connectivité des disques internes dans les serveurs et stations de travail 1U/2U. Au cœur de la gamme SAS 9300 HBA se trouve le contrôleur d'E/S LSI SAS 3008, qui tire parti des dernières avancées de la technologie SAS et PCI Express tout en atteignant plus d'un million d'IOP à partir d'un seul IOC. De plus, chaque HBA de la gamme 1 prend en charge 9300 ou 8 ports SAS individuels, qui fonctionnent à 4 Gb/s et sont rétrocompatibles avec les générations PCIe et SAS précédentes grâce à la négociation automatique.
Les HBAS 9300 de LSI sont basés sur les contrôleurs SAS architecturés Fusion-MPT. Cette architecture utilise la technologie Fusion-MPT (Message Passing Technology) de LSI, dans laquelle chaque contrôleur dispose de processeurs PowerPC intégrés pour offrir un déchargement maximal du processeur hôte et permet à LSI de publier un seul pilote de système d'exploitation binaire pour faire fonctionner n'importe quel contrôleur ou adaptateur Fusion-MPT. En conséquence, LSI indique qu'il permettra des performances élevées, un développement logiciel réduit et une mise sur le marché plus rapide.
Pour cet examen, nous nous concentrerons sur les modèles SAS LSI SAS 9300-8i et LSI SAS 9300-8e 12 Gb / s, ce dernier offrant 2 × 4 ports externes Mini-SAS HD SFF8643 tandis que le premier dispose de 2 × 4 Mini-SAS Ports internes HD SFF8644. En ce qui concerne les applications, le 9300-8i prend en charge les disques SAS et SATA et est spécialement conçu pour les utilisateurs qui ont besoin de bande passante pour les applications critiques dans les serveurs et stations de travail 1U/2U. Les HBA 9300-8e sont idéaux dans les situations où les applications nécessitent une connectivité de stockage externe, donnant aux serveurs de toute taille la possibilité de connecter plus de 1000 12 périphériques finaux SAS ou SATA dans des boîtiers externes. Il dispose également d'un stockage évolutif pour prendre en charge les performances et les applications gourmandes en capacité qui utilisent SAS 6 Gb/s, mais peuvent également prendre en charge les disques SATA XNUMX Gb/s.
Les HBA LSI SAS 9300-8i et LSI SAS 9300-8e sont disponibles au prix public de 290 $ et 405 $ respectivement et sont couverts par une garantie de 3 ans qui comprend une assistance technique avancée gratuite ainsi qu'une option de remplacement avancée.
Spécifications LSI SAS 9300-8i/e
- Modèle : LSI SAS 9300-8i (LSI003450), LSI SAS 9300-8e (LSI00343)
- Dimensions physiques : Profil bas (2.6" x 6.0")
- Supports : Pleine hauteur et profil bas ventilé
- Type de bus hôte : x8 voies PCI Express® 3.0
- Ports internes : 8
- Support de câble : Cuivre passif
- Taux de transfert de données : 12 Gb/s conforme à la norme SAS 3.0
- Périphériques pris en charge : 1024 périphériques SAS/SATA non RAID
- Prise en charge du système d'exploitation : Microsoft Windows, Linux (SuSE, Red Hat), Solaris, VMware, FreeBSD
- Prise en charge de Fusion MPT : Fusion MPT™ 2.5
- Contrôleur d'E/S : LSI SAS3008
- Connecteurs internes (LSI SAS 9300-8i uniquement) : 2 Mini-SAS HD SFF8643
- Connecteurs externes (LSI SAS 9300-8e uniquement) : 2 Mini-SAS HD SFF8644
- MTBF : >2,800,000 XNUMX XNUMX heures
- Température de fonctionnement : 0°C à 55°C
- Humidité en fonctionnement:% 5 à 90% sans condensation
- Tension de fonctionnement : +12 V +/-8 % ; 3.3V +/-9%
- Alimentation PCI : 13 W, 14.5 W
- Garantie : 3 ans ; support technique avancé gratuit, option de remplacement avancé
Conception et construction
Le LSI SAS 9300-8i a une conception très similaire à celle du Supermicro LSI SAS3008 HBA, y compris son facteur de forme HHHL PCIe et la disposition des ports à l'arrière de la carte. De plus, comme nous l'avons mentionné ci-dessus, il est équipé exactement du même contrôleur que le Supermicro, le LSI SAS3008/Fusion MPT 2.5. Les seules différences significatives sont la couleur différente du circuit imprimé ainsi qu'une position de connecteur de port légèrement différente.
Les adaptateurs HBA LSI 12 Gb/s SAS 9300 sont une carte à profil bas de 6.6" × 2.7" et peuvent s'adapter à des serveurs montés en rack 1U/2U ainsi qu'à des stations de travail avec un facteur de forme à profil bas.
Au cœur de la carte se trouve un dissipateur thermique de bonne taille recouvrant le contrôleur d'E/S Fusion MPT 2.5 LSI SAS3008. Les connecteurs de bord de carte PCIe x8 pour son interface PCIe 3.0 sont également visibles.
Les utilisateurs peuvent connecter jusqu'à 1024 appareils SAS et SATA avec ses 8 ports SAS et SATA 9300 Gb/s externes et internes (8-9300e et 8-12i respectivement) et fournir jusqu'à 12 Gb/s SAS et jusqu'à 6 Gb/s Performances SATA sur 8 voies de connectivité PCIe 3.0. L'une des caractéristiques les plus importantes sont leurs 2 Mini-SAS HD (SFF8643 et SFF8644).
Analyse de la charge de travail des applications
Nous avons testé le HBA Supermicro LSI SAS3008 avec les disques professionnels suivants :
- HGST Ultrastar SSD800MM (400 Go, contrôleur Intel DB29AA11B0 comarqué, NAND MLC 25 nm, SAS 12 Gb/s)
- Toshiba PX02SM (unités de 400 Go et 800 Go, contrôleur TC58NC9036GTC comarqué Marvell, Toshiba 24nm eMLC NAND, 12Gb/s SAS)
- Toshiba PX02SS (400 Go, contrôleur Marvell comarqué TC58NC9036GTC, Toshiba 24nm eMLC NAND, 12Gb/s SAS)
- Toshiba PX03SN (800 Go, contrôleur comarqué Marvell TC58NC9036GTC, Toshiba 19 nm MLC NAND, 12 Gb/s SAS)
- Seagate1200 (400 Go, contrôleur Marvell, 21nm Samsung eMLC NAND, 12Gb/s SAS)
Pour comprendre les caractéristiques de performance des périphériques de stockage d'entreprise répertoriés (avec le HBA LSI), il est essentiel de modéliser l'infrastructure et les charges de travail des applications trouvées dans les environnements de production en direct. En tant que tels, nos premiers tests de performances comprennent les performances OLTP de MySQL via SysBench et les performances OLTP de Microsoft SQL Server avec une charge de travail TCP-C simulée.
Le test de base de données Percona MySQL via SysBench mesure les performances d'une activité OLTP avec un groupe de serveurs Lenovo ThinkServer RD630 comme clients de base de données, l'environnement de base de données étant stocké sur un seul disque. Nous examinerons le nombre moyen de transactions par seconde (TPS) pour ces disques. Bien que Percona et MariaDB utilisent les API d'application Fusion-io compatibles avec la mémoire flash dans les versions les plus récentes de leurs bases de données, nous testerons chaque périphérique dans leurs modes de stockage par blocs traditionnels.
Le HBA LSI a permis au Toshiba PX02SS de disposer de la meilleure référence MySQL à ce jour, avec un nombre total d'IOPS impressionnant de plus de 2,150 03. Le PX800SN et le SSDXNUMXMM sont sur sa piste, prenant respectivement la deuxième et la troisième place.
Le protocole de test OLTP Microsoft SQL Server de StorageReview utilise la version préliminaire actuelle du benchmark TPC-C (Transaction Processing Performance Council), un benchmark de traitement transactionnel en ligne qui simule les activités rencontrées dans des environnements applicatifs complexes. Le benchmark TPC-C est plus représentatif que les benchmarks de performance synthétiques des performances réelles, permettant d'évaluer au mieux les points forts et les goulots d'étranglement de l'infrastructure de stockage dans les environnements de bases de données. Notre protocole SQL Server utilise une base de données SQL Server de 685 Go (échelle 3 000) et mesure les performances transactionnelles et la latence sous une charge de 30 000 utilisateurs virtuels.
