---
id: collect-261001-cisco/cisco/fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9-1
title: "fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9"
domain: cisco
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9.md
source_anchor: ""
source_lines: [1, 56]
sha256: 6187cb2b20d96aee9cf217c8354464c9078fa4637dcc0f3e5ce3d89bf6c4c348
---

# fr-review-ubiquiti-enterprise-nas-review-16-bays-dual-25gbe-and-zfs-for-3999-25cfd5e9

L'Enterprise NAS (ENAS) est la première tentative d'Ubiquiti pour proposer une solution de stockage d'entreprise plus légère. Ce châssis 3U comprend 16 baies pour disques remplaçables à chaud, un processeur Arm Neoverse N2 à huit cœurs, 64 Go de mémoire ECC, deux ports SFP28 25 GbE et deux alimentations redondantes remplaçables à chaud de 550 W, le tout pour un prix public conseillé de 3 999 $. À noter : la boutique Ubiquiti applique actuellement une majoration pour la mémoire au moment du paiement, en raison de la hausse des prix de cette dernière. Le prix final est donc supérieur au prix catalogue jusqu'à ce que le coût des composants se stabilise. Le système de fichiers utilisé est ZFS, les baies acceptent tous les disques et tous les logiciels sont inclus. Aucune licence, aucun déblocage de fonctionnalités ni aucun contrat de support ne sont requis.
Le prix d'un NAS nu peut paraître élevé au premier abord, mais il offre de nombreux avantages. La double connectivité 25 GbE en standard est un atout majeur que la concurrence ne peut égaler : le RS4021xs+ 16 baies de Synology est livré avec deux ports RJ45 10 GbE et intègre la connectivité 25 GbE via une carte d'extension PCIe. Le constat est similaire pour les NAS rackables de capacité équivalente, où la connectivité réseau supérieure à 10 GbE est une option payante. Ajoutez à cela une alimentation redondante et de la mémoire ECC, et Ubiquiti répond aux critères qui distinguaient traditionnellement les NAS pour PME des baies d'entrée de gamme pour entreprises. ENAS est disponible dès maintenant directement auprès d'Ubiquiti (lien affilié).
L'ENAS vient couronner une expansion de gamme menée à un rythme remarquable. Ubiquiti s'est forgé une solide réputation grâce à ses points d'accès et passerelles, et son offre réseau continue de se développer. Rien que l'année dernière, nous avons testé les points d'accès UniFi E7 et E7 Campus WiFi 7, la passerelle fibre Cloud Gateway et le routeur Dream Router 7. Le stockage est un domaine plus récent. La gamme a débuté avec l' UNAS Pro , un boîtier 2U à sept baies vendu 499 $, destiné principalement aux clients fidèles d'UniFi. Elle s'est ensuite enrichie avec l' UNAS 2 compact et l' UNAS Pro 8 , qui ont ajouté une alimentation redondante et un cache NVMe. L'ENAS est le premier modèle de la famille à intégrer ZFS, la mémoire ECC, l'iSCSI natif et des ports d'extension SAS. Il est important de noter que les systèmes de fichiers diffèrent suffisamment pour qu'Ubiquiti précise que les données UNAS ne peuvent pas être restaurées directement sur un ENAS ; seuls les utilisateurs, les groupes et les paramètres sont transférés.
Le segment sur lequel ENAS s'implante n'est pas vide. Synology et QNAP, entre autres, dominent le marché des NAS pour PME depuis vingt ans, et leurs plateformes offrent un accès logiciel bien plus complet : écosystèmes d'applications, conteneurs, suites de surveillance et outils de sauvegarde performants qu'UniFi Drive ne propose pas à l'heure actuelle. Cependant, les acteurs historiques ont ouvert la voie à de nouveaux concurrents. Synology a passé la majeure partie de l'année 2025 à appliquer une politique de compatibilité qui limitait de fait ses unités de la série 2025 Plus aux seuls disques validés, principalement de marque Synology, avant de revenir sur sa décision avec DSM 7.3 à l'automne. Dans ce contexte, un système ZFS à 16 baies, prônant l'ouverture des disques et sans licence logicielle, représente un produit parfaitement positionné.
Le profil de l'acheteur idéal est facile à identifier : une petite ou moyenne entreprise, ou le fournisseur de services gérés qui l'exploite, disposant déjà d'un parc UniFi de taille moyenne à importante. Dans ces environnements, ENAS s'intègre à la même console que les commutateurs, passerelles, points d'accès et caméras, hérite de Site Manager pour une visibilité multisite et s'intègre au même modèle d'identité, UniFi Endpoint gérant l'accès distant aux fichiers. Le stockage n'est plus un fournisseur distinct avec une interface utilisateur et un renouvellement séparés. Pour les entreprises qui ont besoin d'un stockage de fichiers et de blocs simple et fiable et qui maîtrisent déjà l'interface UniFi, la continuité opérationnelle repose autant sur le produit que sur le matériel.
Une précision concernant notre unité de test : nous disposons d’ENAS en laboratoire depuis novembre 2025, et la révision matérielle de production a modifié la configuration des E/S arrière, passant de 10 GbE à la double interface 25 GbE disponible dès aujourd’hui. Tout ce qui est présenté et testé ici correspond au matériel commercialisé actuellement, et le logiciel a évolué rapidement depuis la réception de notre unité ; nous y reviendrons.
Spécifications des NAS d'entreprise Ubiquiti
| Spécifications | NAS d'entreprise Ubiquiti | 
|---|---|
| Marché |  | 
| Dimensions | 481.4 × 480 × 132 mm (19 × 18.9 × 5.2 pouces) | 
| Capacité de stockage | 16 baies de lecteur 2.5/3.5 pouces 2 baies M.2 NVMe | 
| Interface réseau | 2 × 25G SFP28 (25G/10G/1G) 1 × 10GbE RJ45 (10G/5G/2.5G/1G/100M) | 
| Port d'extension | 2 × SFF-8644 (24G) | 
| Redondance de l'alimentation | Appareils | 
| Facteur de forme | 3U montage en rack | 
| Hardware |  | 
| Assistance Drive | 16 disques durs/SSD de 2.5/3.5 pouces 2 SSD M.2 NVMe 2 ports d'extension | 
| Budget de puissance maximal pour les entraînements | 450W | 
| Consommation maximale | 550W | 
| Méthode de puissance | Modules d'alimentation remplaçables à chaud, à double entrée CA | 
| Alimentation | 2 modules d'alimentation AC/DC 550 W remplaçables à chaud | 
| Processeur | Processeur ARM N2 à huit cœurs cadencé à 2.4 GHz | 
| Mémoire | 64GB | 
| Direction | Ethernet | 
| Interface RF | Bluetooth 4.1 | 
| Poids | 16.1 kg (lb 35.5) | 
| Matériau du boîtier | Acier SGCC | 
| Matériel de montage | Acier SGCC | 
| Profondeur de rack prise en charge | Les rails supportent des supports à quatre montants de 600 mm (23.6 pouces) avec des trous carrés (9.5 × 9.5 mm). Profondeur des poteaux de 600 à 1066 mm (23.6 à 42 pouces) | 
| Façade | Cadre 3U (écran tactile de 4.7 pouces, LED RGBW) 3U Bezel Lite (vierge) Les deux sont optionnels | 
| LED |  | 
| LED d'état | Ethernet, SFP28, disque dur, système, port d'extension, CRPS | 
| Respect de l'environnement |  | 
| Température de fonctionnement | -5 ° C à 40 ° C (° F à 23 104 ° F) | 
| humidité d'exploitation | 5% à 95% sans condensation | 
| Conforme à la NDAA | Oui | 
| Certifications | FCC, CE, IC | 
| Logiciels |  | 
| Protocoles de fichiers pris en charge | NFS, PME | 
| Protocoles de blocs pris en charge | iSCSI | 
| Types de RAID | Miroir, RAID-Z1, RAID-Z2, RAID-Z3 | 
| Groupes RAID | Multiple | 
| Prise en charge des disques de secours | Appareils | 
| Lecteurs personnels et partagés | Appareils | 
| Cache SSD | Appareils | 
| Capacité maximale des SSD NVMe | 8 Tio | 
| Cryptage des fichiers | Appareils | 
| Prise en charge de la sauvegarde | Serveur UNAS distant, serveur CIFS/SMB, services cloud | 
| Services de sauvegarde dans le cloud | Google Drive, OneDrive, Dropbox, Amazon S3, Backblaze B2, Wasabi | 
| Instantanés | Appareils | 
| Partager des liens | Appareils | 
| Sauvegarde Time Machine | Appareils | 
| Assistance pour l'application client | Appareils | 
| Groupes d'utilisateurs | Appareils | 
Construction et conception
