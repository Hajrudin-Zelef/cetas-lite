---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r770ap-review-84b71047-1
title: "fr-review-dell-poweredge-r770ap-review-84b71047"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "distribution", "ethernet", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r770ap-review-84b71047.md
source_anchor: ""
source_lines: [1, 47]
sha256: 38678912c768d4a4b1e4af20072d8339083a63a3ca3bfb49b93e56cf68e952ef
---

# fr-review-dell-poweredge-r770ap-review-84b71047

Le Dell PowerEdge R770AP n'est pas un serveur à usage général, et c'est précisément son but. Alors que la plupart des plateformes 2U biprocesseurs privilégient la flexibilité, le R770AP fait l'impasse sur tout, sacrifiant la prise en charge GPU, les options de stockage mixtes et la capacité mémoire brute au profit de la densité de cœurs, de la bande passante mémoire et de la stabilité d'exécution maximales disponibles dans la gamme Intel actuelle de Dell. C'est un serveur conçu autour d'une architecture de processeur spécifique pour une catégorie particulière de charges de travail, et il assume pleinement ses limitations.
Pour comprendre sa raison d'être, il faut commencer par la plateforme sur laquelle il s'appuie. La gamme PowerEdge R7x0 de Dell a toujours été le serveur Intel 2U le plus polyvalent de la marque, le PowerEdge R7725, équipé d'un processeur AMD , jouant un rôle équivalent du côté EPYC. Le PowerEdge R770 perpétue cette tradition Intel en prenant en charge les processeurs Xeon 6 à cœurs P et E, les accélérateurs GPU, le stockage mixte SAS/SATA/NVMe, jusqu'à 8 To de mémoire répartis sur 32 emplacements DIMM, et une capacité d'extension PCIe Gen5 suffisante pour couvrir tous les besoins, de la virtualisation à l'inférence IA.
Le PowerEdge R770AP n'est pas ce serveur.
L'appellation « AP » signifie « Performances avancées », mais elle ne rend pas pleinement compte des différences entre ces deux machines. Alors que le R770 utilise la puce Intel Granite Rapids-SP sur socket LGA 4710 avec 8 canaux mémoire et jusqu'à 86 cœurs de performance (P-cores), le R770AP adopte la plateforme Granite Rapids-AP sur socket LGA 7529, offrant jusqu'à 128 cœurs de performance par socket (120 cœurs dans notre configuration de test) et 12 canaux mémoire DDR5. Cette distinction est au cœur de la stratégie globale d'Intel pour sa gamme Xeon 6 série 6900 : les processeurs 6900P sur plateforme AP représentent le nec plus ultra des puces serveur Intel, conçues spécifiquement pour les charges de travail où les performances par cœur, la bande passante mémoire et la stabilité d'exécution priment sur la flexibilité globale de la configuration du serveur.
L'architecture étendue Xeon 6 d'Intel divise le centre de données en deux catégories. Les processeurs à cœur E privilégient la densité et l'efficacité énergétique pour les charges de travail cloud-native et évolutives, telles que les microservices et la distribution de contenu. Les processeurs à cœur P ciblent les tâches de calcul intensives où des performances constantes par thread sont essentielles : simulations HPC, analyses en temps réel, bases de données volumineuses en mémoire et calculs financiers sensibles à la latence. La série 6900P se positionne au sommet de cette gamme de processeurs à cœur P, combinant le plus grand nombre de cœurs disponibles avec une bande passante mémoire à 12 canaux, jusqu'à 96 lignes PCIe Gen5 par socket, jusqu'à 6 liens UPI 2.0 et des pools de cache L3 atteignant 504 Mo sur les modèles haut de gamme comme l'Intel Xeon 6978P. L'objectif architectural n'est pas seulement le débit brut, mais aussi un débit prévisible, minimisant les fluctuations d'ordonnancement et la variabilité des accès mémoire qui dégradent les performances dans les environnements critiques en termes de latence.
Le R770AP est l'incarnation même de cette philosophie chez Dell. Il se débarrasse de tout ce qui est superflu sur la plateforme Granite Rapids-AP : la prise en charge des GPU est totalement supprimée, les options de stockage SAS et SATA sont remplacées par des configurations exclusivement NVMe (jusqu'à 16 SSD NVMe Gen5 de 2.5 pouces ou jusqu'à 32 SSD NVMe E3.S Gen5, selon la configuration), la capacité mémoire est plafonnée à 3 To répartis sur 24 emplacements DIMM (12 par socket, 1 DPC pour une vitesse maximale par canal), et l'extension PCIe est réduite à cinq emplacements Gen5 x16 et deux cartes réseau 3.0 OCP. Il en résulte une plateforme 2U à deux sockets optimisée pour la densité de calcul, la bande passante mémoire et le comportement déterministe exigé par les charges de travail telles que le trading haute fréquence, l'analyse des risques en temps réel et la simulation massivement parallèle.
Notre modèle de test associe deux processeurs Intel Xeon 6978P, chacun doté de 120 cœurs de performance cadencés à 2.1 GHz (fréquence de base) et à 3.2 GHz (mode turbo sur tous les cœurs), et embarque 3 To de mémoire DDR5-6400 répartis sur les 24 emplacements DIMM. Comparé au R770, équipé de deux processeurs Xeon 6787P (86 cœurs chacun, 8 canaux mémoire et 2 To de DDR5), le R770AP offre 39.5 % de cœurs supplémentaires et 50 % de canaux mémoire en plus. Reste à savoir si ces avantages architecturaux se traduisent par des gains de performance concrets et si les compromis de la plateforme sont justifiés pour les charges de travail ciblées par Dell et Intel.
Spécifications du Dell PowerEdge R770AP
Le tableau ci-dessous présente les spécifications de configuration physique et prise en charge pour la plateforme Dell PowerEdge R770AP.
| Spécifications | Dell PowerEdge R770AP | 
|---|---|
| Processeur |  | 
| Processeur | Deux processeurs Intel® Xeon® 6 série 6900 avec cœurs P, jusqu'à 128 cœurs chacun. | 
| Mémoire |  | 
| Emplacements DIMM | 24 emplacements DIMM DDR5 | 
| Mémoire maximale | 3 TB | 
| Vitesse de Mémoire | Jusqu'à 6400 MT / s | 
| Type de mémoire | Modules RDIMM ECC DDR5 enregistrés uniquement | 
| Stockage |  | 
| Contrôleurs de stockage (RAID) | Face avant (interne) du PERC H975i DC-MHS | 
| Soufflet interne | BOSS-N1 DC-MHS : HWRAID 1, 2 SSD M.2 NVMe ou USB | 
| Baies d'entraînement avant | Jusqu'à 16 SSD NVMe G5 x4 de 2.5 pouces (capacité maximale de 245.76 To) Jusqu'à 16 SSD NVMe G5 x2 de 2.5 pouces (capacité maximale de 245.76 To) Jusqu'à 32 SSD NVMe EDSFF E3.S Gen5 (capacité maximale de 491.52 To) | 
| Baies de transmission arrière | N/D | 
| Tuning Moteur |  | 
| Alimentations | 1500 W Titane, 100-120 LLAC ou 200-240 HLAC, 240 V CC, redondance remplaçable à chaud 1800 W Titane, 200-240 HLAC, 240 VDC, redondance remplaçable à chaud 2400 W Titane, 100-120 LLAC ou 200-240 HLAC, 240 V CC, redondance remplaçable à chaud 3200 W Titane, 200-220 HLAC ou 220.1-240 HLAC, 240 VDC, redondance remplaçable à chaud 3200 W Titane, 277 Vca et HVDC, redondance remplaçable à chaud* | 
| Refroidissement et ventilateurs |  | 
| Options de refroidissement | refroidissement par air | 
| Ventilateurs | Jusqu'à 6 ventilateurs remplaçables à chaud | 
| Facteur de forme et dimensions |  | 
| Facteur de forme | Serveur rack 2U | 
| Hauteur | 86.8 mm (3.42 pouces) | 
| Largeur | 482 mm (19.0 pouces) | 
| Profondeur (avec lunette) | 802.40 mm (31.59 pouces) | 
| Profondeur (sans lunette) | 801.51 mm (31.56 pouces) | 
| Biseau | Lunette métallique en option | 
| Réseautage et expansion |  | 
| Options de réseau OCP | Jusqu'à deux cartes réseau OCP NIC 3.0 Emplacement 4 : 1×8 ou 1×16 Gen5 OCP 3.0 Emplacement 10 : 1×16 Gen5 OCP 3.0 | 
| Carte réseau intégrée | Port Ethernet BMC dédié 1 Gb | 
| Emplacements PCIe | Jusqu'à 5 emplacements PCIe Gen5 (connecteurs x16) Emplacement 2 : 1×16 Gen5, pleine hauteur, demi-longueur Emplacement 3 : 1×16 Gen5, pleine hauteur/profil bas, demi-longueur Emplacement 5 : 1×16 Gen5, pleine hauteur, demi-longueur Emplacement 7 : 1×16 Gen5, pleine hauteur, demi-longueur Emplacement 9 : 1×16 Gen5, pleine hauteur/profil bas, demi-longueur | 
| Options GPU | N/D | 
| Ports |  | 
| Ports avant | 1x USB 2.0 de type C | 
| Ports arrière | 1 port Ethernet BMC dédié 2x USB 3.1 type A 1x VGA | 
| Ports internes | 1x USB 3.1 type A | 
| Direction |  | 
| Gestion intégrée | iDRAC10, iDRAC Direct, API RESTful iDRAC avec Redfish, interface de ligne de commande RACADM, module de service iDRAC | 
| Sécurité |  | 
