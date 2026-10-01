---
id: collect-250926-servers-hardware/servers-hardware/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84-3
title: "fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "dram", "energy", "ethernet", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84.md
source_anchor: ""
source_lines: [80, 124]
sha256: 66ef27bd6d6d174f09dd155dec18d5f2ab66bc9b8d7aada04660ee3ed64422d1
---

# fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84

| Protection de la mémoire | ECC, SDDC, nettoyage de patrouille/à la demande, défaut limité, parité de commande d'adresse DRAM avec relecture, nouvelle tentative d'erreur ECC non corrigée de DRAM, ECC sur puce, vérification et nettoyage d'erreur ECC (ECS), réparation après package | 
| Baies de lecteur |  | 
| Stockage interne maximal | 92.16 To avec 6 disques SSD SAS/SATA 15.36 pouces de 2.5 To 92.16 To avec 6 disques SSD NVMe 15.36 pouces 2.5 To | 
| Contrôleur de stockage | Ports NVMe intégrés (Intel VROC NVMe en option pour RAID) Ports SATA intégrés Prise en charge de l'adaptateur RAID interne (CFF) pour la prise en charge des disques SAS/SATA | 
| Baies de lecteur optique | Pas de baies internes ; utilisez une clé USB externe. | 
| Baies de lecteur de bande | Pas de baies internes. Utilisez une clé USB externe. | 
| Interfaces réseau | Emplacement OCP 3.0 SFF dédié avec interface hôte PCIe 5.0 x16. Prend en charge une variété d'adaptateurs à 2 et 4 ports avec une connectivité réseau 1, 10, 25 ou 100 GbE. En option, un port peut être partagé avec le processeur de gestion XClarity Controller 2 (XCC2) pour la prise en charge Wake-on-LAN et NC-SI. | 
| Emplacements PCIe | Un ou deux emplacements PCIe x16 pleine hauteur demi-longueur (FHHL), en fonction du nombre de processeurs installés : Configurations à 1 processeur :  Configurations à 2 processeur :  | 
| Prise en charge du GPU | Prend en charge 2x GPU simple largeur | 
| Ports | Avant : un port VGA pour la vidéo, un port USB 3.2 G1 (5 Gb/s), un port de diagnostic externe et un port série DB9 pour la connectivité locale Arrière : un port MiniDP pour la vidéo, un port USB 3.2 G1 (5 Gb/s), 1 port USB 2.0 (également pour la gestion locale XCC), 1 port de gestion des systèmes RJ-45 1GbE pour la gestion à distance XCC | 
| Refroidissement | 3 ventilateurs à double rotor de 60 mm à remplacement simple avec redondance de rotor N+1 | 
| Source d'alimentation | Fourni par le châssis D3. | 
| Pièces remplaçables à chaud | Variateurs | 
| Gestion des systèmes | Panneau de commande avec LED d'état. Combiné de diagnostic externe en option avec écran LCD. Gestion intégrée XClarity Controller 2 (XCC2) basée sur le contrôleur de gestion de carte mère (BMC) ASPEED AST2600, la fourniture d'infrastructure centralisée XClarity Administrator, les plug-ins XClarity Integrator et la gestion centralisée de l'alimentation du serveur XClarity Energy Manager - XCC Platinum en option pour activer les fonctions de contrôle à distance et d'autres fonctionnalités. . | 
| Vidéo | Graphiques embarqués avec 16 Mo de mémoire avec accélérateur matériel 2D, intégrés au contrôleur de gestion XClarity Controller 2. Deux ports vidéo (VGA avant et Mini DisplayPort arrière) ; les deux peuvent être utilisés simultanément si vous le souhaitez. La résolution maximale des deux ports est de 1920×1200 à 60 Hz. | 
| Sécurité | Mot de passe à la mise sous tension, mot de passe de l'administrateur, Trusted Platform Module (TPM), prenant en charge TPM 2.0. | 
| Systèmes d'exploitation pris en charge | Microsoft Windows Server, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi et Ubuntu Server. Pour plus de détails, y compris d'autres systèmes d'exploitation certifiés par le fournisseur ou testés, consultez la section de prise en charge du système d'exploitation. | 
| Garantie limitée | Unité remplaçable par le client de trois ans et garantie limitée sur site avec 9×5 le jour ouvrable suivant (NBD). | 
| Service et support | Des mises à niveau de service facultatives sont disponibles via les services Lenovo : temps de réponse de 4 heures ou 2 heures, temps de réparation de 6 heures, extension de garantie de 1 an ou 2 ans, support logiciel pour le matériel Lenovo et certaines applications tierces. | 
| Température ambiante | Jusqu'à ASHRAE Classe A2 : 10°C – 35°C (50°F – 95°F) | 
| Dimensions | Largeur : 222 mm (8.7 pouces), hauteur : 82 mm (3.2 pouces), profondeur 898 mm (35.4 pouces) | 
| Poids | Maximum : 11.76 kg (25.93 lb) | 
Présentation du Lenovo ThinkSystem SD535 V3
Le ThinkSystem SD530 V3 est une version AMD du ThinkSystem SD530 V3 basé sur Intel. La principale différence est qu'il ne prend en charge qu'un seul processeur, mais ce n'est pas nécessairement un inconvénient puisque sa seule puce AMD EYPIC série 9004 peut comporter jusqu'à 128 cœurs, ce qui est plus que ce que le SD530 V3 peut prendre en charge même lorsqu'il est équipé de deux processeurs. Le ThinkSystem SD535 V3 peut prendre en charge plus de cœurs au total que le SD530 V3 ; ce dernier n'en prend en charge que 64, même avec deux processeurs installés. Le SD535 V3 prend également en charge plus de mémoire par socket, avec 12 emplacements DIMM prenant en charge 1.5 To en utilisant des RDIMM 128DS de 3 Go.
Ici, nous pouvons voir le socket CPU unique et les emplacements DIMM environnants. Le refroidissement par air du SD535 V3 est le même que celui du SD530 V3, avec quatre ventilateurs arrière de 40 mm à remplacement simple. Le stockage intégré comprend deux disques de démarrage M.2. Un adaptateur SAS/SATA RAID est également pris en charge pour les configurations avec six SSD frontaux de 2.5 pouces.
Les ports avant et arrière du SD535 V3 sont identiques à ceux du SD530 V3, avec USB, série, VGA et un port de diagnostic pour le contrôleur XClarity ; de même, à l'arrière, vous trouverez Ethernet de gestion à distance, deux ports USB-A, une sortie vidéo mini-DisplayPort, un OCP 3.0 et un emplacement PCIe PCIe x16 Gen 5.
Spécifications du Lenovo ThinkSystem SD535 V3
| Composants | Spécifications | 
| Type de machine | 7DD1 – 3 ans de garantie 7DD8 – 1 ans de garantie | 
| Facteur de forme | Nœud de calcul 1U demi-largeur. | 
| Boîtier pris en charge | Châssis ThinkSystem D3, hauteur 2U ; jusqu'à 4 serveurs par châssis. | 
| Processeur | Un processeur AMD EPYC 4 de 9004e génération. Prend en charge les processeurs jusqu'à 128 cœurs, les vitesses de cœur jusqu'à 4.1 GHz et les valeurs TDP jusqu'à 400 W. | 
| Chipset | Non applicable (les fonctions du hub du contrôleur de plateforme sont intégrées au SOC du processeur) | 
| Mémoire | 12 emplacements DIMM avec 12 canaux mémoire du processeur (1 DIMM par canal, DPC). Les RDIMM Lenovo TruDDR5 et 3DS sont pris en charge jusqu'à 4800 XNUMX MHz | 
| Mémoire persistante | Non pris en charge | 
| Mémoire maximale | Jusqu'à 1.5 To de mémoire système grâce à 12 modules RDIMM 128DS de 3 Go | 
| Protection de la mémoire | ECC, SDDC, nettoyage de patrouille/à la demande, défaut limité, parité de commande d'adresse DRAM avec relecture, nouvelle tentative d'erreur ECC non corrigée de DRAM, ECC sur puce, vérification et nettoyage d'erreur ECC (ECS), réparation après package | 
| Baies de lecteur |  | 
| Stockage interne maximal |  | 
| Contrôleur de stockage | Ports NVMe intégrés (pas de prise en charge du RAID) Ports SATA intégrés (pas de prise en charge du RAID) Prise en charge d'un adaptateur RAID interne (CFF) ou d'un adaptateur PCIe pour la prise en charge des disques SAS/SATA | 
| Baies de lecteur optique | Pas de baies internes ; utilisez une clé USB externe. | 
| Baies de lecteur de bande | Pas de baies internes. Utilisez une clé USB externe. | 
| Interfaces réseau | Emplacement OCP 3.0 SFF dédié avec interface hôte PCIe 5.0 x16. Prend en charge une variété d'adaptateurs à 2 et 4 ports avec une connectivité réseau 1, 10, 25 ou 100 GbE. Un port peut être partagé avec le processeur de gestion XClarity Controller 2 (XCC2) pour la prise en charge Wake-on-LAN et NC-SI. | 
| Emplacements PCIe | Un emplacement PCIe 5.0 x16 avec un facteur de forme discret | 
| Prise en charge du GPU | Prend en charge 1x GPU simple largeur | 
