---
id: collect-250926-servers-hardware/servers-hardware/fr-review-dell-poweredge-c6615-server-review-e5a753e4-1
title: "fr-review-dell-poweredge-c6615-server-review-e5a753e4"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "China", "Intel", "Microsoft"]
dates: []
keywords: ["amd", "datacenter", "distribution", "ethernet", "intel"]
source: docs/RAG/clean4/fr-review-dell-poweredge-c6615-server-review-e5a753e4.md
source_anchor: ""
source_lines: [1, 43]
sha256: bd58b4f834ab8305a3e67155d383916c5f635b019487083559b901bc167f44c0
---

# fr-review-dell-poweredge-c6615-server-review-e5a753e4

La plate-forme Dell PowerEdge C-Series dispose d'un châssis 2U prenant en charge quatre serveurs dans la catégorie Dell Modular Infrastructure. En fonction de la charge de travail, le système C-series peut être configuré avec deux types de nœuds différents : un nœud AMD C6615 à socket unique ou un nœud Intel C6620 à double socket.
Notre examen se concentrera sur le châssis de la série C, qui comporte quatre nœuds AMD EPYC à socket unique connectés à un fond de panier de disque E8.S PCIe Gen3 à 5 baies.
Du point de vue du stockage, la plate-forme peut être configurée avec un fond de panier de disque SFF de 2.5 pouces, qui prend en charge jusqu'à 24 SSD NVMe ou un support Gen5 en exploitant un fond de panier E8.S à 3 baies. En interne, ces disques sont connectés directement à chaque nœud, avec une répartition égale sur les quatre serveurs. Par exemple, dans la configuration à 24 baies, chaque nœud voit six disques ; dans la configuration à 8 baies, chaque nœud voit deux disques.
Le châssis C6600 offre des alimentations et un refroidissement redondants partagés pour les quatre nœuds installés, bien qu'au-delà de cela, chaque nœud soit géré indépendamment. Ainsi, contrairement à un châssis lame géré avec un portail de gestion de châssis, il s'agit plutôt de quatre petits serveurs PowerEdge sous un même toit métallique. Chaque nœud C6615 dispose de connexions réseau dédiées, d'une interface iDRAC et d'emplacements PCIe pour l'extension.
Spécifications du nœud Dell PowerEdge C6615
| Spécifications C6615 |  | 
|---|---|
| Processeur | Un processeur AMD EPYC avec jusqu'à 64 cœurs | 
| Mémoire | 6 emplacements DIMM DDR5, prend en charge RDIMM de 576 Go (6 x 96 Go) maximum, vitesses allant jusqu'à 4800 XNUMX MT/s | 
| Contrôleurs de stockage | Contrôleurs internes (RAID) : PERC H755N, PERC H355 Démarrage interne : sous-système de stockage optimisé au démarrage (NVMe BOSS-N1) : HWRAID 1, 2 x SSD M.2 HBA SAS 12 Gbit/s internes (non RAID) : HBA355i RAID logiciel : S160 | 
| Disponibilité | Disques durs et blocs d'alimentation redondants enfichables à chaud | 
| Baies de disques | Baies avant : Jusqu'à 16 disques SAS/SATA (HDD/SSD) de 2.5 pouces, 61 To maximum Jusqu'à 16 disques SATA/NVMe de 2.5 pouces, 15.36 To maximum sur une configuration de fond de panier universel Jusqu'à 16 x 2.5 pouces sur fond de panier NVMe Jusqu'à 8 x E3.s sur le fond de panier de disque dur SSD NVMe | 
| Remplacement à chaud, alimentations redondantes | 3200 277 W 336 VCA ou XNUMX VCC 2800 200 W Titane 240-240 VCA ou XNUMX VCC Platine 2400 100 W 240-240 VCA ou XNUMX VCC 1800 200 W Titane 240-240 VCA ou XNUMX VCC | 
| Dimensions | Hauteur – 40.0 mm (1.57 pouces) Largeur – 174.4 mm (6.86 pouces) Profondeur – 549.7 mm (21.64 pouces), 561.3 mm (22.10 pouces) – SAS/SATA ou NVMe ou E3.S ou configuration universelle | 
| Poids | 3.7 kg (8.15 livres) | 
| Gestion intégrée | IDRAC9 IDRAC Direct API RESTful IDRAC avec Redfish Module de service IDRAC | 
| Logiciel OpenManage | Plug-in CloudIQ pour PowerEdge OpenManage Entreprise Intégration OpenManage Enterprise pour VMware Vcenter Intégration OpenManage pour Microsoft System Center Intégration d'OpenManage avec le centre d'administration Windows Plug-in OpenManage Power Manager Plug-in de service OpenManage Plug-in OpenManage Update Manager | 
| intégrations | BMC TrueSight Microsoft System Center Intégration d'OpenManage avec ServiceNow Intégration d'OpenManage avec le centre d'administration Windows Plug-in OpenManage Power Manager Plug-in de service OpenManage Plug-in OpenManage Update Manager | 
| Sécurité | Virtualisation cryptée sécurisée AMD (SEV) Chiffrement de mémoire sécurisé AMD (SME) Micrologiciel signé cryptographiquement Chiffrement des données au repos (SED avec gestion de clé locale ou externe) Vérification des composants Secure BootSecured (vérification de l'intégrité du matériel) Secure Erase Racine de confiance en silicone Verrouillage du système (nécessite IDRAC9 Enterprise ou Datacenter) TPM 2.0 FIPS, certifié CC-TCG, TPM 2.0 China NationZ | 
| Carte réseau intégrée | 1 x 1 Go | 
| Ports arrière | 1 x USB 3.0 1 port Ethernet IDRAC 1 port IDRAC Direct (Micro-AB USB) 1 x Mini DisplayPort | 
| Emplacements PCIE | Jusqu'à 2 emplacements PCIe x16 Gen5 Low-Profile 1 x OCP 3.0 x16 Gen5 | 
| Système d'exploitation et hyperviseurs | Serveur canonique Ubuntu LTS Serveur Microsoft Windows avec Hyper-V Red Hat Enterprise Linux Serveur d'entreprise SUSE Linux VMware ESXi/vSAN | 
Construire et concevoir
Le châssis Dell PowerEdge C6600 et les nœuds C6615 offrent une option informatique exceptionnellement dense pour les scénarios de déploiement qui doivent minimiser l'espace physique utilisé dans un environnement de montage en rack. Cela convient aux solutions hyperconvergées fonctionnant dans un environnement en cluster, nécessitant plusieurs nœuds ou des charges de travail lourdes qui ne nécessitent pas la consommation de 4U ou 8U via les conceptions de serveurs traditionnelles 1U ou 2U. Le châssis a une empreinte 2U avec une profondeur de 30 pouces. Le poids du châssis peut y monter en fonction de la configuration finale. Dell indique un poids maximum d'une configuration C16 à 6600 baies avec tous les disques installés à 93.69 livres.
L'avant du système est assez basique par rapport aux autres plates-formes PowerEdge, sans beaucoup de marque Dell. Ce type de serveur n'offre pas le cadre PowerEdge standard mais place les disques et les entrées de ventilateur au premier plan. L'avant de la version E3.S C6600 comporte huit SSD Gen5 NVMe au milieu, flanqués d'entrées de ventilateur de refroidissement.
Les oreilles latérales du châssis contiennent des boutons d'alimentation dédiés pour chaque nœud et des boutons d'information indiquant l'état ou les problèmes de ce nœud.
Chaque nœud C6615 dispose d'une disposition de ports condensée à l'arrière du châssis par rapport à un serveur 1U ou 2U traditionnel. Les ports incluent USB, iDRAC, un connecteur d'affichage et un port de service USB.
Pour la mise en réseau, un emplacement OCP est disponible pour différentes options d'interface (la nôtre dispose d'une carte réseau 25GbE à quatre ports), et deux emplacements PCIe sont également disponibles. Les emplacements OCP et double PCIe offrent une interface Gen5.
L'ouverture du châssis PowerEdge C6600 vous donne une visibilité sur la disposition de la manière dont le refroidissement, la distribution de l'alimentation et les chemins d'E/S des disques sont gérés. Le câblage PCIe/SAS du fond de panier du disque est acheminé directement vers chaque nœud via des raccords à connexion rapide qui transmettent également les données et l'alimentation.
En fonction de la configuration interne de chaque nœud, les connexions des lecteurs se connectent directement à la carte mère ou à une carte PERC pour les options RAID matérielles.
Hormis le refroidissement et l’alimentation, les nœuds ne partagent aucune autre ressource.
Performances du Dell PowerEdge C6615
Spécifications des nœuds testés
Nos quatre nœuds C6615 ont des configurations identiques. Nous les comparerons et montrerons les performances moyennes sur les nœuds.
- 1 processeur AMD EPYC 8534P 64 cœurs
- 6 x 96 Go DDR5 4800 576 Mo/s (XNUMX Go)
- Windows Server standard 2022
- SSD de démarrage Dell RAID1 BOSS
- 2 SSD PCIe Gen5 E3.S
Lors de nos tests de performances, les nœuds ont fonctionné en parallèle pour donner un score global prenant en compte les ressources d'alimentation et de refroidissement partagées.
Performances de stockage
Chacun des quatre nœuds Dell Power Edge C6615 comprend un SSD BOSS RAID1 pour le démarrage et deux baies E3.S pour les SSD d'entreprise Gen5. Bien que la carte BOSS ne soit pas en reste, elle offre un profil de performances très différent de celui des SSD E3.S.
