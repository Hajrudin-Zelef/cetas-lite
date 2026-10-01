---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf-1
title: "fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "datacenter", "ethernet", "gpu", "intel", "nvidia"]
source: docs/RAG/clean4/fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf.md
source_anchor: ""
source_lines: [1, 55]
sha256: 8408e755c5dd6fe1483c2340832d0bed9ec18d2969023eaa53ec07189b715dca
---

# fr-review-dell-poweredge-r4715-review-1u-amd-epyc-built-for-the-midmarket-9cd432cf

La gamme PowerEdge de 17e génération de Dell est déjà bien implantée, et avec les R4715 et R5715, elle cible désormais plus spécifiquement les PME et les entreprises de taille intermédiaire. Ces deux serveurs monoprocesseurs reposent sur la même architecture AMD EPYC de 5e génération que l'ensemble de la gamme PowerEdge de 17e génération. Ils sont optimisés pour les organisations où le nombre de cœurs adapté, la gestion des licences et la simplicité d'utilisation priment sur le débit maximal. Le R4715 est un modèle 1U plus compact, conçu pour la virtualisation, les bases de données à grande échelle et les déploiements en périphérie de réseau. Le R5715 adopte un format 2U, offrant davantage de baies de disques et d'extension PCIe, et convient aux configurations où la capacité de stockage et les performances d'E/S sont essentielles.
Le serveur R4715 est conçu pour les organisations exécutant des charges de travail de virtualisation, de bases de données à grande échelle et de calcul en périphérie, pour lesquelles l'efficacité des licences et la simplicité d'utilisation sont essentielles. Notre modèle de test était équipé d'un processeur AMD EPYC 9335, le processeur haut de gamme à 32 cœurs disponible sur cette plateforme, associé à 384 Go de mémoire DDR5 et une configuration de démarrage RAID 1 BOSS. Outre ses 32 cœurs, le 9335 offre 128 Mo de cache L3 et un TDP de 210 W. Les utilisateurs n'ayant pas besoin d'autant de cœurs peuvent opter pour les modèles EPYC 9255 (24 cœurs), EPYC 9135 (16 cœurs) ou EPYC 9015 (8 cœurs).
Quelques points importants à préciser d'emblée concernant la plateforme : le R4715 ne prend pas en charge les GPU ni les DPU. Il ne s'agit pas d'un oubli. Cette plateforme est conçue spécifiquement pour les charges de travail axées sur le processeur, et Dell a fait des compromis délibérés afin de limiter le coût des composants et l'encombrement. Pour les charges de travail nécessitant des accélérateurs, les modèles R6715 et R7715 sont plus adaptés.
Le R4715 offre un châssis compact refroidi par air, avec jusqu'à trois emplacements PCIe Gen5, 24 emplacements DDR5 RDIMM et des options de stockage flexibles de 3.5 et 2.5 pouces, y compris U.2 NVMe. De plus, il intègre iDRAC10 avec OpenManage Enterprise et une sécurité matérielle via une racine de confiance au niveau du silicium. Les options réseau incluent 25 GbE via OCP 3.0, 100 GbE et 400 GbE via PCIe AIC. Broadcom, Intel et NVIDIA complètent l'écosystème de cartes réseau ; à noter qu'aucune connectivité Fibre Channel n'est officiellement prise en charge. Il fonctionne avec des alimentations de 800 W ou 4715 1100 W, disponibles en versions Platinum ou Titanium, et prend en charge la redondance tolérante aux pannes. Pour un serveur d'entrée de gamme, les fonctionnalités de gestion et de sécurité pour entreprises sont globalement satisfaisantes.
Spécifications Dell PowerEdge R4715
Le tableau ci-dessous met en évidence les spécifications physiques et matérielles de la plateforme Dell PowerEdge R4715.
| Spécifications | Dell PowerEdge R4715 | 
|---|---|
| Processeur |  | 
| Processeur | Un processeur AMD EPYC série 9005 de 5e génération, jusqu'à 32 cœurs | 
| Facteur de forme | Serveur rack 1U | 
| Mémoire |  | 
| Emplacements DIMM | 24 emplacements DIMM DDR5 | 
| Mémoire maximale | 1.5 To (jusqu'à 64 Go par DIMM) | 
| Vitesse de Mémoire | Jusqu'à 5200 MT / s | 
| Type de mémoire | Modules RDIMM ECC DDR5 enregistrés uniquement | 
| Stockage |  | 
| Contrôleurs internes (RAID) | PERC H365i, H965i | 
| Soufflet interne | BOSS-N1 DC-MHS | 
| HBA externes | N/D | 
| Baies d'entraînement avant | 4x 3.5 pouces SAS 8 ports SAS/SATA 2.5 pouces 8 ports U.2 NVMe G4 | 
| Tuning Moteur |  | 
| Alimentations | Platine 800 W, 1100 W Titane 800 W, 1100 W FTR pris en charge | 
| Refroidissement et ventilateurs |  | 
| Options de refroidissement | refroidissement par air | 
| Ventilateurs | Jusqu'à quatre ensembles (module à double ventilateur) de ventilateurs remplaçables à chaud | 
| Dimensions |  | 
| Hauteur | 42.8 mm (1.68 pouces) | 
| Largeur | 482.0 mm (18.97 pouces) | 
| Profondeur (avec lunette) | 816.921 mm (32.16 pouces) | 
| Profondeur (sans lunette) | 815.141 mm (32.09 pouces) | 
| Biseau | Lunette métallique en option | 
| Réseautage et expansion |  | 
| Options de réseau OCP | 2 cartes réseau OCP 3.0 (en option), 1 GbE, 10 GbE, 25 GbE Emplacement 2 : 1×16 Gen5 OCP 3.0 Emplacement 5 : 1×16 Gen5 OCP 3.0 | 
| Carte réseau intégrée | Port Ethernet BMC dédié 1 Gb | 
| Carte réseau PCIe AIC | 100 GbE et 400 GbE ; NDR VPI (400 GbE) | 
| Emplacements PCIe | Jusqu'à 3 emplacements PCIe Gen5 (connecteurs x16) Emplacement 1 : 1×16 Gen5 pleine hauteur ou profil bas Emplacement 2 : 1×16 Gen5 Low Profile ou 1×16 OCP3.0 Emplacement 4 : 1×16 Gen5 pleine hauteur ou profil bas | 
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
Dell Conception et installation du PowerEdge R4715
Le R4715 est un serveur rack 1U mesurant 1.68 mm de hauteur, 18.97 mm de largeur et 32.09 mm de profondeur sans la façade métallique optionnelle (32.16 mm avec la façade). La face avant comprend un bouton d'alimentation, un bouton d'identification du système, un port USB 2.0 Type-A (utilisé avec le module KVM LCP optionnel), un port USB 2.0 Type-C pour l'accès direct à iDRAC et un port MiniDisplayPort optionnel pour la même configuration KVM. L'ouverture des baies de disques se fait sans outil sur l'ensemble du châssis.
Configuration du stockage
