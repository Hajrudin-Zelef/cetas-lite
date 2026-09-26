---
id: collect-260926-rattrapage/rattrapage/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868-1
title: "fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "agent", "arr", "valuation"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868.md
source_anchor: ""
source_lines: [1, 42]
sha256: a00bd8cc19dbb396372059a12c1f218d6978af8691afdeb3c4a708c53042f659
---

# fr-review-supermicro-hyper-superserver-sys-212ha-tn-review-intel-xeon-6-6980p-0f575868

Supermicro nous a présenté le Hyper SuperServer SYS-212HA-TN, un système présenté comme un serveur 2U polyvalent conçu pour une évolutivité et des performances de niveau entreprise. Doté des processeurs Intel® Xeon® 6 série 6900, ce serveur est présenté comme capable d'exceller dans une large gamme de charges de travail, notamment la virtualisation, l'inférence IA, l'apprentissage automatique, le stockage défini par logiciel, le cloud computing, la gestion de base de données et le stockage haute densité. Cette évaluation examinera dans quelle mesure il tient sa promesse de polyvalence et de hautes performances dans des scénarios réels.
Spécifications du serveur Hyper SuperServer SYS-212HA-TN de Supermicro
| Catégories | DÉTAILS | 
|---|---|
| SKU de produit | SuperServer SYS-212HA-TN | 
| Carte mère | Super X14SBH-AP | 
| Processeur | Processeurs Intel® Xeon® série 7529 à socket unique BR (LGA-6900) avec cœurs P, jusqu'à 128 C/256 T, jusqu'à 504 Mo de cache, prend en charge les processeurs TDP jusqu'à 500 W (refroidissement par air) | 
| Mémoire système | Nombre d'emplacements : 12 emplacements DIMM, Mémoire maximale (1DPC) : Jusqu'à 3 To 8800MT/s ECC DDR5 MRDIMM, Jusqu'à 3 To 6400MT/s ECC DDR5 RDIMM, Tension de la mémoire : 1.1 V | 
| Appareils embarqués | Prise en charge NVMe RAID 0/1/5/10 (clé RAID Intel® VROC requise), système sur puce, connectivité réseau via AIOM | 
| Entrée / Sortie | LAN : 1 port LAN BMC dédié RJ45 1 GbE, USB : 2 ports USB 3.2 Gen1 (arrière), Vidéo : 1 port VGA | 
| BIOS système | AMI 64 Mo SPI Flash | 
| Direction | SuperCloud Composer, Supermicro Server Manager (SSM), Super Diagnostics Offline (SDO), Supermicro Thin-Agent Service (TAS), SuperServer Automation Assistant (SAA) | 
| Sécurité | TPM 2.0, Silicon Root of Trust (RoT), micrologiciel signé cryptographiquement, démarrage sécurisé, récupération automatique du micrologiciel | 
| Châssis | Montage en rack 2U, CSE-HS201-R000NFP | 
| Dimensions et poids | Hauteur : 3.5″ (88.9 mm), largeur : 17.2″ (437 mm), profondeur : 31.74″ (806.2 mm), poids brut : 75 lb (34 kg), poids net : 45 lb (20.5 kg) | 
| Panneau avant | LED : activité du disque dur, activité LAN1, état de l'alimentation, informations système, boutons : marche/arrêt, bouton UID | 
| Slots d'extension | Option A : 1 PCIe 5.0 x16, 1 PCIe 5.0 x8, 1 PCIe 5.0 x16 double largeur, 1 emplacement PCIe 5.0 x16 AIOM ; Option B : 2 PCIe 5.0 x8, 1 PCIe 5.0 x8 HHFL, 1 PCIe 5.0 x16 double largeur, 1 emplacement PCIe 5.0 x16 AIOM | 
| Baies de lecteur/stockage | 8 baies de lecteur NVMe*/SAS*/SATA* 2.5″ remplaçables à chaud à l'avant, option A : 24 baies de lecteur SAS*/SATA* 2.5″ remplaçables à chaud à l'avant, 2 emplacements M.2 NVMe | 
| Refroidissement du système | 6 ventilateurs contrarotatifs 60x60x56mm | 
| Alimentation | 2 alimentations redondantes de niveau Titanium de 1200 1300 W, 1600 2000 W, 2600 XNUMX W, XNUMX XNUMX W, XNUMX XNUMX W | 
| Environnement d'exploitation | Température de fonctionnement : 10°C ~ 35°C, Température hors fonctionnement : -40°C à 70°C, Humidité relative de fonctionnement : 8% à 90% | 
Configurations de processeur prises en charge
Ce serveur prend en charge les processeurs Intel Xeon de 6e génération (Granite Rapids), optimisés spécifiquement pour les processeurs à cœurs P. Il exploite pleinement les cœurs P dédiés aux performances pour une puissance de calcul accrue. Contrairement à l' Hyper SuperServer 1U 112H-TN , précédemment testé et compatible avec les processeurs Intel Xeon de 6e génération (Sierra Forest) à cœurs E, ce châssis est conçu spécifiquement pour les processeurs à cœurs P, garantissant ainsi des performances maximales grâce à l'architecture à cœurs P.
Vous trouverez ci-dessous la liste des processeurs compatibles avec cette unité :
| SKU | Noyaux / Threads | Fréquence de base (GHz) | Tous les cœurs Turbo (GHz) | Turbo max (GHz) | TDP (Watts) | Cache (Mo) | Évolutivité maximale | Vitesse de la mémoire DDR5 | Vitesse de la mémoire MRDIMM | 
|---|---|---|---|---|---|---|---|---|---|
| 6980P | 128 | 2.0 | 3.2 | 3.9 | 500 | 504 | 2S | 6400 | 8800 | 
| 6979P | 120 | 2.1 | 3.2 | 3.9 | 500 | 504 | 2S | 6400 | 8800 | 
| 6972P | 96 | 2.4 | 3.5 | 3.9 | 500 | 480 | 2S | 6400 | 8800 | 
| 6952P | 96 | 2.1 | 3.2 | 3.9 | 400 | 480 | 2S | 6400 | 8800 | 
| 6960P | 72 | 2.7 | 3.8 | 3.9 | 500 | 432 | 2S | 6400 | 8800 | 
Mémoire
Ce serveur est conçu pour les charges de travail hautes performances et prend en charge jusqu'à 12 emplacements DIMM remplis, ce qui permet une capacité de mémoire maximale de 3 To lorsqu'il est configuré avec 1DPC. Il offre des performances exceptionnelles avec la prise en charge des modules ECC DDR5 MRDIMM à des vitesses allant jusqu'à 8800 5 MT/s ou des modules ECC DDR6400 RDIMM jusqu'à XNUMX XNUMX MT/s, ce qui le rend idéal pour les applications hautes performances et gourmandes en mémoire.
Stockage
En ce qui concerne le stockage, ce serveur offre des configurations polyvalentes pour répondre aux différents besoins des applications. Par défaut, il prend en charge huit baies de lecteur NVMe/SAS/SATA 2.5″ échangeables à chaud à l'avant, offrant des capacités de stockage à grande vitesse. Le système s'étend à 24 baies pour les demandes de stockage plus importantes. De plus, le système comprend deux emplacements M.2 NVMe (M-key 2280/22110/25110) pour les SSD hautes performances, avec VROC requis pour les configurations RAID. Comme indiqué dans la liste des pièces en option, des configurations spécifiques peuvent nécessiter des contrôleurs de stockage ou des câbles supplémentaires. Cette flexibilité garantit l'évolutivité et l'adaptabilité pour diverses exigences de stockage. Il n'y a pas de prise en charge EDSFF aujourd'hui, mais la page de configuration nous amène à penser que Supermicro pourrait offrir une prise en charge plus poussée des facteurs de forme des lecteurs à l'avenir.
Networking
Le SYS-212HA-TN dispose d'un emplacement PCIe 5.0 x16 AIOM (compatible OCP 3.0) pour les cartes d'interface réseau. Il offre également des options d'extension telles que 1 emplacement PCIe 5.0 x16 FH/10.5″ L double largeur, 1 emplacement PCIe 5.0 x8 FHFL et 1 emplacement PCIe 5.0 x16 HHFL double largeur. Cette flexibilité permet une intégration transparente des solutions réseau à large bande passante pour répondre aux exigences du cloud, de l'IA et d'autres applications gourmandes en données.
Tuning Moteur
Pour garantir une fiabilité et une disponibilité maximales, Supermicro propose des configurations de serveur avec 2 blocs d'alimentation redondants 1+1 pour une fiabilité et une disponibilité améliorées. Ces blocs d'alimentation vont de 1200 2000 W à 120 240 W et prennent en charge des plages d'entrée de 44 V/65 V CA ou de -96 à -XNUMX V CC, répondant ainsi à divers besoins d'application. Disponibles en niveaux d'efficacité Titanium (XNUMX %) et Gold, ils offrent d'excellentes performances énergétiques, équilibrant efficacité opérationnelle et rentabilité pour diverses charges de travail.
Conception et fabrication du serveur Hyper SuperServer SYS-212HA-TN de Supermicro
Le SYS-212HA-TN est un rack 2U standard conçu pour les environnements de serveur classiques tout en conservant un profil compact. Ses dimensions sont de 3.5 mm (88.9″) de hauteur, 17.2 mm (437″) de largeur et 31.74 mm (806.2″) de profondeur, ce qui garantit la compatibilité avec les racks de serveur standard. Les dimensions de l'emballage sont de 9.96″ (H) x 26.46″ (L) x 43.31″ (P), et l'unité pèse 45 kg (20.5 lb) net, avec un poids d'expédition brut de 75 kg (34 lb). Construit avec un châssis argenté pratique, complété par un panneau avant noir et une conception d'oreille de rack Supermicro traditionnelle
Panneau avant
