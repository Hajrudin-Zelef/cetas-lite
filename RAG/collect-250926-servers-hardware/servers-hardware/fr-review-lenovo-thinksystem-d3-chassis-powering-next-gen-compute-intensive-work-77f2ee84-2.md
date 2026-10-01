---
id: collect-250926-servers-hardware/servers-hardware/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84-2
title: "fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["dram", "energy", "gpu", "intel"]
source: docs/RAG/clean4/fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84.md
source_anchor: ""
source_lines: [35, 79]
sha256: 8754ef94b41370b308802e0fd9c4396509238e0e1d0b0e118102973cd5f12071
---

# fr-review-lenovo-thinksystem-d3-chassis-powering-next-gen-compute-intensive-work-77f2ee84

Quatre ventilateurs de refroidissement de 40 mm à remplacement simple se trouvent derrière les processeurs et aspirent l'air par l'avant. Il est rare dans les serveurs multi-nœuds de placer les ventilateurs de refroidissement dans le nœud lui-même, c'est pourquoi Lenovo a réalisé quelque chose d'unique ici. Cela signifie que les nœuds ne partagent pas de fans, ce qui les rend plus indépendants. Cette conception rend les nœuds les plus efficaces car leurs exigences en matière de charge de travail changent les unes par rapport aux autres.
Maintenant, regardons le SD550 V3. Ce serveur se distingue du SD530 V3 en façade en proposant six disques SSD de 2.5 pouces (SAS, SATA ou NVMe). Le panneau avant est par ailleurs le même que celui du SD530 V3. L'arrière du serveur est similaire au SD530 V3, mais le SD550 V3 ajoute un emplacement PCIe x16 Gen 5 pleine hauteur/demi-longueur, permettant encore plus d'options de mise en réseau et de GPU.
En interne, le SD550 V3 dispose de deux sockets et de huit emplacements DIMM par socket, tout comme le SD530 V3. La différence est que le SD550 V3 possède des dissipateurs thermiques et des ventilateurs de processeur beaucoup plus grands (trois, et non quatre), de sorte qu'il fonctionnera plus frais et aura une efficacité énergétique supérieure. Les deux processeurs peuvent atteindre 350 W/64 cœurs/3.9 GHz sans pénalité pour l'installation d'un deuxième processeur. Le SD550 V3 prend également en charge un adaptateur RAID interne.
Spécifications du Lenovo ThinkSystem SD530 V3
| Composants | Spécifications | 
| Type de machine | 7DD3 – 3 ans de garantie 7DDA – 1 an de garantie | 
| Facteur de forme | Nœud de calcul 1U demi-largeur. | 
| Boîtier pris en charge | Châssis ThinkSystem D3, hauteur 2U ; jusqu'à 4 serveurs par châssis. | 
| Processeur | Un ou deux processeurs Intel Xeon Scalable de 5e génération (anciennement nommés « Emerald Rapids »). Avec un processeur installé, prend en charge des processeurs jusqu'à 64 cœurs, des vitesses de cœur jusqu'à 3.9 GHz et des valeurs TDP jusqu'à 350 W. Avec deux processeurs installés, prend en charge des processeurs jusqu'à 32 cœurs, des vitesses de cœur jusqu'à 3.9 GHz et des valeurs TDP jusqu'à 205 W. | 
| Chipset | Chipset Intel C741 « Emmitsburg », faisant partie de la plate-forme nommée « Eagle Stream ». | 
| Mémoire | 16 emplacements DIMM avec deux processeurs (8 emplacements DIMM par processeur) par nœud. Chaque processeur dispose de huit canaux mémoire, avec 1 DIMM par canal (DPC). Les RDIMM Lenovo TruDDR5 et 3DS sont pris en charge jusqu'à 5600 XNUMX MHz | 
| Mémoire persistante | Non pris en charge | 
| Mémoire maximale | Jusqu'à 1 To de mémoire système (en utilisant soit 8 RDIMM 128DS de 3 Go, soit 16 RDIMM de 64 Go) | 
| Protection de la mémoire | ECC, SDDC, nettoyage de patrouille/à la demande, défaut limité, parité de commande d'adresse DRAM avec relecture, nouvelle tentative d'erreur ECC non corrigée de DRAM, ECC sur puce, vérification et nettoyage d'erreur ECC (ECS), réparation après package | 
| Baies de lecteur |  | 
| Stockage interne maximal | 30.72 To avec 2 disques SSD E15.36.S EDSFF NVMe de 3 To | 
| Contrôleur de stockage | 2x ports NVMe intégrés (Intel VROC NVMe en option pour RAID) | 
| Baies de lecteur optique | Il n'y a pas de baies internes ; utilisez une clé USB externe. | 
| Baies de lecteur de bande | Il n'y a pas de baies internes. Utilisez une clé USB externe. | 
| Interfaces réseau | Emplacement OCP 3.0 SFF dédié avec interface hôte PCIe 5.0 x16. Prend en charge une variété d'adaptateurs à 2 et 4 ports avec une connectivité réseau 1, 10, 25 ou 100 GbE. En option, un port peut être partagé avec le processeur de gestion XClarity Controller 2 (XCC2) pour la prise en charge Wake-on-LAN et NC-SI. | 
| Emplacements PCIe | Un emplacement PCIe 5.0 x16 avec un facteur de forme discret | 
| Prise en charge du GPU | Prend en charge 1x GPU simple largeur | 
| Ports | Avant : un port VGA pour la vidéo, un port USB 3.2 G1 (5 Gb/s), un port de diagnostic externe et un port série DB9 pour la connectivité locale Arrière : un port MiniDP pour la vidéo, un port USB 3.2 G1 (5 Gb/s), 1 port USB 2.0 (également pour la gestion locale XCC), 1 port de gestion des systèmes RJ-45 1GbE pour la gestion à distance XCC | 
| Refroidissement | 4 ventilateurs à double rotor de 40 mm à remplacement simple avec redondance de rotor N+1 | 
| Source d'alimentation | Fourni par le châssis D3. | 
| Pièces remplaçables à chaud | Variateurs | 
| Gestion des systèmes | Panneau de commande avec LED d'état. Combiné de diagnostic externe en option avec écran LCD. Gestion intégrée XClarity Controller 2 (XCC2) basée sur le contrôleur de gestion de carte mère (BMC) ASPEED AST2600, la fourniture d'infrastructure centralisée XClarity Administrator, les plug-ins XClarity Integrator et la gestion centralisée de l'alimentation du serveur XClarity Energy Manager - XCC Platinum pour activer les fonctions de contrôle à distance et d'autres fonctionnalités. | 
| Vidéo | Graphiques embarqués avec 16 Mo de mémoire avec accélérateur matériel 2D, intégrés au contrôleur de gestion XClarity Controller 2. Deux ports vidéo, VGA avant et Mini DisplayPort arrière. Les deux ports peuvent être utilisés simultanément si vous le souhaitez. La résolution maximale des deux ports est de 1920×1200 à 60 Hz. | 
| Sécurité | Mot de passe à la mise sous tension, mot de passe administrateur, Trusted Platform Module (TPM), prenant en charge TPM 2.0. | 
| Systèmes d'exploitation pris en charge | Microsoft Windows Server, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi et Ubuntu Server. Pour plus de détails, y compris d'autres systèmes d'exploitation certifiés par le fournisseur ou testés, consultez la section de prise en charge du système d'exploitation. | 
| Garantie limitée | Unité remplaçable par le client de trois ans et garantie limitée sur site avec 9×5 le jour ouvrable suivant (NBD). | 
| Service et support | Des mises à niveau de service facultatives sont disponibles via les services Lenovo : temps de réponse de 4 heures ou 2 heures, temps de réparation de 6 heures, extension de garantie de 1 an ou 2 ans, support logiciel pour le matériel Lenovo et certaines applications tierces. | 
| Température ambiante | Jusqu'à ASHRAE Classe A2 : 10°C – 35°C (50°F – 95°F) | 
| Dimensions | Largeur : 222 mm (8.7 pouces), hauteur : 41 mm (1.6 pouces), profondeur 908 mm (35.7 pouces) | 
| Poids | Maximum : 7.6 kg (16.76 lb) | 
Spécifications du Lenovo ThinkSystem SD550 V3
| Composants | Spécifications | 
| Type de machine | 7DD2 – 3 ans de garantie 7DD9 – 1 ans de garantie | 
| Facteur de forme | Nœud de calcul 2U demi-largeur. | 
| Boîtier pris en charge | Châssis ThinkSystem D3, hauteur 2U ; jusqu'à 2 serveurs SD550 V3 par châssis. | 
| Processeur | Un ou deux processeurs Intel Xeon Scalable de 5e génération (anciennement nommés « Emerald Rapids »). Prend en charge les processeurs jusqu'à 64 cœurs, des vitesses de cœur jusqu'à 3.9 GHz et des valeurs TDP jusqu'à 350 W. | 
| Chipset | Chipset Intel C741 « Emmitsburg », faisant partie de la plateforme nommée « Eagle Stream » | 
| Mémoire | 16 emplacements DIMM avec deux processeurs (8 emplacements DIMM par processeur) par nœud. Chaque processeur dispose de huit canaux mémoire, avec 1 DIMM par canal (DPC). Les RDIMM Lenovo TruDDR5 et 3DS sont pris en charge jusqu'à 5600 XNUMX MHz | 
| Mémoire persistante | Non pris en charge | 
| Mémoire maximale | Jusqu'à 2 To en utilisant 16 RDIMM 128DS de 3 Go | 
