---
id: collect-261001-general-networking/general-networking/fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b-1
title: "fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b"
domain: general-networking
role: reference
task: reference
actors: ["Intel", "Microsoft", "Samsung"]
dates: []
keywords: ["benchmark", "benchmarks", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b.md
source_anchor: ""
source_lines: [1, 47]
sha256: 49428e136971530c6dd6475f3f6a63a7354e42bb4df76241a0727f64f6d86286
---

# fr-review-lsi-sas-9207-8i-e-host-bus-adapter-review-f7ef8c4b

Bien que les adaptateurs de bus hôte LSI SAS 9207-8i/e existent depuis quelques années maintenant, ils sont toujours la norme de facto pour les solutions SAS2, qui a été immédiatement adaptée lors de la sortie de PCIe 3.0. Depuis la création des HBA et des contrôleurs compatibles SAS 6 Gb/s, LSI a été son plus grand promoteur, ayant fabriqué des appareils incroyablement réussis depuis le début de la gamme 9200. Avec ces adaptateurs HBA, le LSI a joué un rôle déterminant dans l'intégration du 6 Gb/s dans le courant dominant.
Cela dit, les LSI SAS 9207-8i et 9207-8e ne font pas exception à ce fait en raison de leur équilibre entre prix abordable, performances et évolutivité interne. Les HBA LSI sont conçus pour les applications de stockage de serveur de milieu de gamme, y compris le stockage hiérarchisé et la sauvegarde et la restauration avec la capacité du 9207-8i de connecter jusqu'à 256 périphériques SAS et SATA grâce à ses 8 ports SAS 6 Gb/s internes, tandis que le 9207- 8e prend en charge jusqu'à 1024 SAS ou SATA avec ses 8 ports externes. Le 8i et le 8e sont tous deux équipés du contrôleur d'E/S SAS LSISAS2308 6 Gb/s ainsi que d'un processeur PowerPC double cœur 800 MHz.
En ce qui concerne l'encombrement, les HBA LSI peuvent s'intégrer dans des serveurs et des stations de travail montés en rack 1U/2U avec un facteur de forme compact et prennent en charge toutes les applications critiques qui utilisent la bande passante de la connectivité PCIe 3.0.
Spécifications techniques
- Modèle : LSI SAS 9207-8i (LSI00302)
- Contenu : HBA, deux câbles Mini-SAS vers SAS x4 (CBL-SAS8087OCF-06M) et CD de documentation
- Supports : Pleine hauteur et profil bas
- Type de bus hôte : x8 voies PCI Express 3.0
- Dimensions physiques : Profil bas (2.6" x 6.6")
- Ports internes : 8
- Support de câble : Cuivre passif
- Taux de transfert de données : 6 Gb/s conforme à la norme SAS 2.1
- Périphériques pris en charge : 256 périphériques SAS/SATA non RAID
- Prise en charge du système d'exploitation : Microsoft Windows, Linux (SuSE, Red Hat), Solaris, VMware, FreeBSD
- Prise en charge de Fusion MPT : Fusion MPT 2.0
- Contrôleur d'E/S : LSI SAS2308
- Connecteurs internes : 2 Mini-SAS SFF8087
- MTBF : > 2,000,000 XNUMX XNUMX heures
- Température de fonctionnement: 0 ° C à 55 ° C
- Humidité en fonctionnement:% 5 à 90% sans condensation
- Tension de fonctionnement : +12 V +/-8 % ; 3.3V +/-8%
- Alimentation PCI : 9.8 W typique, débit d'air 200 LFM
- Garantie : 3 ans ; avec option de remplacement avancée
Conception et construction
Comme nous l'avons mentionné ci-dessus, les HBA LSI SAS 9207-8i/e sont des facteurs de forme à profil bas qui s'intègrent dans n'importe quel rack de serveur 1U et 2U, offrant 8 voies de SAS 6 Gb/s ainsi que 8 voies de PCI Express 3.0 avec 8 Gb/s. . La ligne 9207 est sensiblement différente des nouveaux modèles LSI (tels que la série 93XX), y compris une nuance de vert différente ainsi qu'une forme différente.
Au cœur des 9207-8i et 9207-8e se trouve le contrôleur SAS LSISAS2308 6Gb/s, qui est recouvert d'un grand dissipateur thermique. Le contrôleur LSI est intégré aux processeurs PowerPC, ce qui permet un déchargement maximal du processeur hôte et donne à LSI la possibilité de publier un seul pilote de système d'exploitation binaire pour faire fonctionner n'importe quel contrôleur ou adaptateur Fusion MPT ; une caractéristique très attrayante pour les entreprises.
Les deux connecteurs mini-SAS internes SFF-9207 (SFF8) du 4-8087i pour son interface PCIe 8087 sont situés à l'extrémité de la carte, ce qui fournit des taux de transfert de données SAS et SATA de 3.0, 1.5 et 3 Gb/s par voie avec la capacité d'atteindre plus de 6 650 IOP. Le 9207-8i peut prendre en charge jusqu'à 256 périphériques non RAID.
Le 9207-8e est équipé de deux (x4) connecteurs mini-SAS externes (SFF8088) situés à l'avant gauche de la carte, avec la possibilité d'offrir également des taux de transfert de données SAS et SATA de 1.5, 3 et 6 Gb/s par voie à 650 9207 IOP. Le 8-1024e peut prendre en charge jusqu'à XNUMX XNUMX périphériques non RAID.
Analyse de la charge de travail des applications
Pour évaluer ses performances lors d'une charge de travail applicative, nous examinerons le test du Toshiba HK3R2, où le contrôleur LSI SAS 9207-8i a été utilisé. Comme vous le verrez ci-dessous, ses performances sont impressionnantes.
Notre environnement de base de données MarkLogic NoSQL nécessite des groupes de quatre disques SSD d'une capacité utile d'au moins 200 Go, car la base de données NoSQL nécessite environ 650 Go d'espace pour ses quatre nœuds de base de données. Notre protocole utilise un hôte SCST et présente chaque SSD dans JBOD, avec un alloué par nœud de base de données. Le test se répète sur 24 intervalles, nécessitant entre 30 et 36 heures au total. MarkLogic enregistre la latence moyenne totale ainsi que la latence d'intervalle pour chaque SSD.
Dans nos tests de latence moyenne globale utilisant notre référence de base de données MarkLogic NoSQL, le HK3R2 a affiché une latence moyenne globale de 2.122 ms lors de l'utilisation du 9207-8i.
Le test de performance suivant consiste à évaluer les performances d'une base de données Percona MySQL OLTP à l'aide de SysBench . Dans cette configuration, nous avons utilisé un groupe de serveurs Lenovo ThinkServer RD630 comme clients de base de données, l'environnement étant stocké sur un seul disque. Pour cette analyse, nous nous concentrerons uniquement sur ses performances en termes de transactions moyennes par seconde (TPS).
Tirant parti du HBA 9207-8i, le SSD Toshiba HK3R2 a poursuivi ses meilleures performances de sa catégorie, atteignant près de 1,700 XNUMX TPS. Ces résultats en font le grand gagnant parmi ses comparables.
Le protocole de test OLTP Microsoft SQL Server de StorageReview utilise la version préliminaire actuelle du benchmark TPC-C (Transaction Processing Performance Council's Benchmark C), un benchmark de traitement transactionnel en ligne qui simule les activités rencontrées dans des environnements applicatifs complexes. Le benchmark TPC-C est plus représentatif que les benchmarks de performance synthétiques des performances réelles, permettant d'évaluer avec précision les points forts et les goulots d'étranglement de l'infrastructure de stockage dans les environnements de bases de données. Notre protocole SQL Server utilise une base de données SQL Server de 685 Go (échelle 3 000). Nous analyserons ses performances en termes de latence avec une charge de 30 000 unités virtuelles (VU).
Ici, le Toshiba HK3R2 affichait un impressionnant 12.0 ms (à égalité avec le Samsung 845DC Pro) pour la meilleure latence moyenne globale.
Analyse synthétique de la charge de travail d'entreprise
Pour mesurer les performances synthétiques d'entreprise du contrôleur HBA LSI SAS 9207-8i/e, nous prendrons comme référence le châssis d'extension de stockage JBOD iXsystems Titan 316J sur lequel il a été utilisé avec les disques suivants :
- Toshiba MK01GRRB (147 Go, 15,000 6.0 tr/min, XNUMX Gb/s SAS)
- Toshiba MBF2600RC (600 Go, 10,000 6.0 tr/min, XNUMX Gb/s SAS)
- Hitachi Ultrastar 7K4000 (4 To, 7,200 6.0 tr/min, XNUMX Gb/s SATA)
La plateforme de tests d'entreprise était basée sur un serveur Lenovo ThinkServer RD630 . Le ThinkServer RD630 est configuré avec :
- 2 x Intel Xeon E5-2620 (2.0 GHz, cache de 15 Mo)
- Windows Server 2008 R2 SP1 64 bits, Windows Server 2012 64 bits et CentOS 6.3 64 bits
- Jeu de puces Intel C602
- Mémoire – 16 Go (2 x 8 Go) 1333 Mhz DDR3 enregistrés RDIMM
- HBA LSI 9207 SAS/SATA 6.0Gb/s
