---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r5715-review-bc780bb6-1
title: "fr-review-dell-poweredge-r5715-review-bc780bb6"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Microsoft"]
dates: []
keywords: ["amd", "datacenter", "ethernet", "gpu"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r5715-review-bc780bb6.md
source_anchor: ""
source_lines: [1, 55]
sha256: 959ee2d8f8769bfe0b3b34039d7ac280917844a14ede7fc5a4269cb156dcaccb
---

# fr-review-dell-poweredge-r5715-review-bc780bb6

Le PowerEdge R5715 est le deuxième modèle de la gamme PowerEdge de 17e génération de Dell, dédiée aux PME. Il se distingue de son homologue 1U par des priorités différentes. Alors que le R4715 privilégie la densité de calcul et l'efficacité du nombre de cœurs par unité de rack, le R5715 est conçu autour de la capacité de stockage et de l'extensibilité des E/S dans un format 2U mono-socket. Les lecteurs de notre test du R4715 retrouveront les fondamentaux de la plateforme : même famille de processeurs AMD EPYC de 5e génération, même architecture mémoire DDR5 à 24 emplacements et même système de gestion iDRAC 10. Seule la tâche confiée au R5715 diffère légèrement.
Notre serveur de test était équipé d'un processeur AMD EPYC 9015, le modèle 8 cœurs de la gamme Turin, associé à 384 Go de mémoire DDR5 et une configuration RAID 1 BOSS pour le démarrage. Nos tests ont porté sur le fond de panier de stockage 3.5 pouces à 12 baies du R5715, un domaine où le 9015 prend tout son sens. Les charges de travail telles que le partage de fichiers, la sauvegarde de données et la vidéosurveillance en point de vente n'ont pas besoin de 32 cœurs ; elles privilégient la densité de stockage, un débit soutenu et une gestion fiable. Le 9015 permet de limiter la consommation d'énergie et les coûts de licence, tout en offrant jusqu'à 288 To de capacité de stockage brute dans un seul nœud 2U.
La R5715 augmente le nombre de ports PCIe Gen5 à quatre, contre trois pour la R4715, et ajoute un port réseau OCP 3.0 supplémentaire, offrant ainsi une plus grande flexibilité pour répondre à la demande croissante d'E/S. Les deux plateformes prennent en charge les protocoles 100 GbE et 400 GbE via PCIe AIC, ce qui les rend parfaitement adaptées aux environnements exigeant une large bande passante. Cependant, aucune des deux ne prend officiellement en charge la connectivité Fibre Channel. De plus, aucune ne prend en charge les GPU ni les DPU. Elles fonctionnent avec les mêmes alimentations de 800 W et 1100 W, disponibles en versions Platinum et Titanium, avec une redondance tolérante aux pannes et un refroidissement par air.
Spécifications Dell PowerEdge R5715
Le tableau ci-dessous met en évidence les spécifications physiques et matérielles de la plateforme Dell PowerEdge R5715.
| Spécifications | Dell PowerEdge R5715 | 
|---|---|
| Processeur |  | 
| Processeur | Un processeur AMD EPYC série 9005 de 5e génération, jusqu'à 32 cœurs | 
| Facteur de forme | Serveur rack 2U | 
| Mémoire |  | 
| Emplacements DIMM | 24 emplacements DIMM DDR5 | 
| Mémoire maximale | 1.5 To (jusqu'à 64 Go par DIMM) | 
| Vitesse de Mémoire | Jusqu'à 5200 MT / s | 
| Type de mémoire | Modules RDIMM ECC DDR5 enregistrés uniquement | 
| Stockage |  | 
| Contrôleurs internes (RAID) | PERC H365i, H965i | 
| Soufflet interne | BOSS-N1 DC-MHS | 
| HBA externes | N/D | 
| Baies d'entraînement avant | 12 ports SAS/SATA 3.5 pouces 16 ports SAS/SATA 2.5 pouces | 
| Tuning Moteur |  | 
| Alimentations | Platine 800 W, 1100 W Titane 800 W, 1100 W FTR pris en charge | 
| Refroidissement et ventilateurs |  | 
| Options de refroidissement | refroidissement par air | 
| Ventilateurs | Jusqu'à six ventilateurs enfichables à chaud | 
| Dimensions |  | 
| Hauteur | 86.8 mm (3.41 pouces) | 
| Largeur | 482.0 mm (18.97 pouces) | 
| Profondeur (avec lunette) | 802.4 mm (31.59 pouces) | 
| Profondeur (sans lunette) | 801.51 mm (31.55 pouces) | 
| Biseau | Lunette métallique en option | 
| Réseautage et expansion |  | 
| Options de réseau OCP | 2 cartes réseau OCP 3.0 (en option), 1 GbE, 10 GbE, 25 GbE Emplacement 4 : 1×16 Gen5 OCP 3.0 Emplacement 10 : 1×16 Gen5 OCP 3.0 | 
| Carte réseau intégrée | Port Ethernet BMC dédié 1 Gb | 
| Carte réseau PCIe AIC | 100 GbE et 400 GbE ; NDR VPI (400 GbE) | 
| Emplacements PCIe | Jusqu'à 4 emplacements PCIe Gen5 (connecteurs x16) Emplacement 2 : 1×16 Gen5 Pleine hauteur Emplacement 3 : 1×16 Gen5 Pleine hauteur Emplacement 7 : 1×16 Gen5 Pleine hauteur Emplacement 9 : 1×16 Gen5 Pleine hauteur | 
| Options GPU | N/D | 
| Ports |  | 
| Ports avant | 1 port USB 2.0 Type-A (KVM LCP en option) 1 port USB 2.0 Type-C (HÔTE/BMC Direct) 1 port MiniDisplayPort (KVM LCP en option) | 
| Ports arrière | 2x USB 3.1 type A 1x VGA Port Ethernet BMC dédié 1 Gb | 
| Ports internes | 1x USB 3.1 type A | 
| Direction |  | 
| Gestion intégrée | iDRAC10, iDRAC Direct, API RESTful iDRAC avec Redfish, interface de ligne de commande RACADM, module sans fil Quick Sync 2 | 
| Logiciel OpenManage | OpenManage Enterprise (OME), OME Power Manager, OME Services, OME Update Manager, OME APEX AIOps Observability, OME Integration for VMware vCenter, OME Integration for Microsoft System Center, OpenManage Integration for Windows Admin Center | 
| Outils | IPMI | 
| intégrations | Intégrations OpenManage : Collections Red Hat Ansible, fournisseurs Terraform | 
| La Gestion du changement | Gestionnaire de référentiel Dell, Mise à jour du système Dell, Catalogues d'entreprise, Utilitaire de mise à jour du serveur (SUU) | 
| Sécurité |  | 
| Caractéristiques de sécurité | Firmware signé cryptographiquement, chiffrement des données au repos (SED avec gestion de clés locale ou externe), démarrage sécurisé, vérification sécurisée des composants (contrôle d'intégrité matérielle), effacement sécurisé, racine de confiance sur le silicium, verrouillage du système (nécessite iDRAC10 Enterprise ou Datacenter), TPM 2.0 certifié FIPS/CC-TCG, détection d'intrusion dans le châssis, virtualisation chiffrée sécurisée AMD (SEV), chiffrement sécurisé de la mémoire AMD (SME) | 
| Systèmes d'exploitation et hyperviseurs |  | 
| Systèmes d'exploitation/hyperviseurs pris en charge | Serveur Ubuntu LTS de Canonical, Microsoft Windows Server avec Hyper-V, Red Hat Enterprise Linux, SUSE Linux Enterprise Server, VMware ESXi | 
Le Dell PowerEdge R5715 est un serveur rack 2U monoprocesseur basé sur la plateforme AMD EPYC série 9005 de 5e génération. Conçu comme une plateforme de stockage performante pour les entreprises exigeant une capacité élevée et une connectivité E/S fiable sans le surcoût d'une architecture biprocesseur, le R5715 cible les charges de travail telles que les bases de données, les partages de fichiers, les sauvegardes et la virtualisation, où un seul processeur EPYC puissant gère la charge plus efficacement que deux processeurs de génération précédente. Nous avons également utilisé ce châssis pour notre comparatif de consommation énergétique entre disques durs et mémoire flash Micron 6600 ION , en remplaçant huit disques durs de 30 To par un SSD de 245 To. Prenant en charge jusqu'à 288 To de stockage brut et disposant de quatre emplacements d'extension PCIe Gen5, le R5715 offre des performances bien supérieures à celles de sa catégorie de prix.
Panneau extérieur et avant
Le R5715 est livré avec un cadre métallique optionnel arborant le motif hexagonal emblématique de Dell. Ce cadre s'emboîte parfaitement sur le châssis et laisse apparaître les commandes du panneau avant sur le côté droit : un bouton d'alimentation, un port USB 2.0 Type-C pour un accès direct au BMC, un port iDRAC Direct et un bouton d'identification du système. Sans le cadre, le châssis mesure 3.41 cm de hauteur, 18.97 cm de largeur et 31.55 cm de profondeur, ce qui lui permet de s'intégrer dans les racks 2U standard. La qualité de fabrication est irréprochable, avec des loquets de baie de disques sans outil et des clips de fixation bleus utilisés systématiquement sur les composants internes pour un accès rapide.
Configuration du stockage
