---
id: collect-261001-ia-llm/ia-llm/fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64-1
title: "fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64"
domain: ia-llm
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["arr", "ethernet"]
source: docs/RAG/collect-261001-ia-llm/fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64.md
source_anchor: ""
source_lines: [1, 59]
sha256: 44bf4b3631ab39780e4bbd58f409b301b1fa8e2c14119ddaf9c302178d6b2f28
---

# fr-review-ubiquiti-unas-pro-review-streamlined-storage-for-unifi-enthusiasts-b6d79f64

L'Ubiquiti UNAS Pro est l'un des derniers ajouts à l'écosystème croissant d'appareils en réseau d'Ubiquiti. Il offre une solution NAS hautes performances qui s'intègre parfaitement à la plateforme UniFi. Conçu pour les particuliers et les professionnels, l'UNAS Pro allie des options de stockage flexibles, des fonctionnalités robustes de protection des données et une expérience utilisateur simplifiée grâce à la nouvelle application UniFi Drive. Ce système rackable 2U offre une solution polyvalente pour le stockage de fichiers, les sauvegardes et l'accès à distance, et s'adresse aux utilisateurs recherchant une plateforme de stockage gérée localement et sans abonnement.
Au moment de la rédaction de cet avis, l' UNAS Pro est proposé au prix de 499 $ sur la boutique Ubiquiti (lien affilié).
Spécifications UNAS Pro
| Mécaniques | Spécifications UNAS Pro | 
| Dimensions | (17.4 x 12.8 x 3.4″) 442 x 325 x 87 mm | 
| Poids | Sans supports de montage : 9.2 kg (20.3 lb) Avec supports de montage : 9.5 kg (20.8 lb) | 
| Matériel de clôture | Acier SGCC | 
| Matériel de montage | Acier SGCC | 
| Hardware |  | 
| Processeur | Processeur ARM® Cortex®-A57 quadricœur à 1.7 GHz | 
| La mémoire système | 8 GB | 
| Interface de gestion | Ethernet dans la bande | 
| Interface réseau | (1) port GbE RJ45 (1) port SFP+ 1/10 GbE | 
| Interface RF | Bluetooth 4.1 | 
| Tuning Moteur |  | 
| Méthode d'alimentation | (1) Entrée CA universelle, 100-240 V CA, 3 A max., 50/60 Hz (1) Entrée CC USP-RPS, 11.5 V CC, 13.91 A | 
| Source d'alimentation | AC/DC, interne, 200 W | 
| Plage de tension prise en charge | 100—240 V CA | 
| Max. consommation d'énergie | 160W | 
| Environnemental |  | 
| Température ambiante de fonctionnement | -5 à 40 ° C (23 à 104 ° F) | 
| Humidité ambiante de fonctionnement | 5 à 95 % sans condensation | 
| Certifications | FCC, CE, IC | 
| Exigences de la demande |  | 
| Lecteur UniFi | Version 1.16.0 et versions ultérieures | 
UNAS Pro Construction et conception
L'UNAS Pro partage une conception similaire à celle de l'UNVR Pro, avec un format rack 2U et une construction robuste en acier SGCC. Ses dimensions compactes (442 x 325 x 87 mm) la rendent idéale pour les baies réseau et serveurs. L'UNAS Pro offre une mémoire système totale de 8 Go, garantissant des performances accrues pour la gestion de volumes de données et de cache plus importants. Elle fonctionne sous UniFi Drive, un logiciel conçu pour un stockage et une gestion efficaces des fichiers, contrairement à UniFi Protect, présent sur l' UNVR Pro.
La face avant de l'UNAS Pro est équipée de sept baies de disques durs 2.5/3.5 pouces, compatibles avec les disques durs et SSD, chacune dotée de voyants LED pour une surveillance en temps réel de l'alimentation, de l'activité et des erreurs du disque. Un écran tactile de 1.3 pouce permet une gestion intuitive des appareils. Les options de connectivité incluent une liaison montante SFP+ 1/10G, un port RJ1 45 GbE pour une mise en réseau flexible et un commutateur de réinitialisation.
L'arrière de l'unité est équipé de deux ventilateurs pour assurer un refroidissement optimal des disques. Il dispose également d'un port d'alimentation C13 verrouillable pour des connexions sécurisées et d'un port USP-RPS pour une alimentation redondante, améliorant ainsi la fiabilité et la disponibilité du système.
Protection et sécurité des données
UNAS Pro assure la sécurité des fichiers grâce à plusieurs niveaux de protection, notamment la redondance RAID, le chiffrement des disques et les snapshots automatiques. Les configurations RAID offrent différents niveaux de tolérance aux pannes, notamment :
- RAID 5 (protection de base) : nécessite un minimum de trois disques.
- RAID 6 (protection moyenne) : nécessite un minimum de quatre disques pour une protection contre les pannes de deux disques.
- RAID 10 (protection supérieure) : nécessite au moins quatre disques pour une redondance améliorée.
Les lecteurs chiffrés sécurisent davantage les données et se verrouillent automatiquement au redémarrage ou à la mise hors tension du système. Les administrateurs doivent ressaisir la clé de chiffrement pour accéder aux fichiers après le redémarrage.
Capacités de sauvegarde et de capture instantanée
Le système prend en charge jusqu'à 4,096 XNUMX instantanés sur tous les disques, avec des limites individuelles de :
- 128 instantanés par disque personnel
- 256 instantanés par lecteur partagé
Les instantanés réguliers permettent aux utilisateurs de restaurer efficacement les versions précédentes des fichiers, minimisant ainsi les risques de perte de données.
Partage et accès aux fichiers
UNAS Pro prend en charge plusieurs protocoles de partage de fichiers, notamment SMB, CIFS et NFS, garantissant un accès fluide sur différents systèmes d'exploitation. De plus, l'intégration avec Active Directory (AD) et LDAP permet aux administrateurs d'importer des utilisateurs et de gérer efficacement les identifiants des services de fichiers.
Les utilisateurs peuvent accéder directement à UniFi Drive via SMB depuis leur ordinateur, offrant ainsi une solution simple et efficace pour gérer leurs fichiers. L'application Identity Endpoint, disponible sur ordinateur et appareils mobiles, offre une connexion fluide et sans licence.
Options de sauvegarde dans le cloud et à distance
UNAS Pro permet des sauvegardes pour :
- Autres appareils UNAS Pro distants
- Appareils via SMB/CIFS
- Google Drive pour la redondance basée sur le cloud
- Prise en charge de OneDrive
Prise en charge de macOS Time Machine
Pour les utilisateurs de macOS, UNAS Pro prend en charge Apple Time Machine, permettant des sauvegardes automatisées et sans effort. Des instructions d'installation détaillées sont disponibles dans le guide officiel.
Couverture étendue des soins d'assurance-chômage
Pour les utilisateurs recherchant une protection supplémentaire, UI Care propose une extension de garantie facultative à 99 $ par unité. Cette extension prolonge la période de protection de remplacement à cinq ans et comprend des avantages exclusifs tels que :
- Services RMA prioritaires : recevez un appareil de remplacement avant de retourner l'original.
- Livraison de retour gratuite : une étiquette d'expédition prépayée est fournie pour des retours sans tracas.
- Couverture étendue : protection de remplacement de cinq ans pour une tranquillité d’esprit totale.
Avec UI Care, les utilisateurs bénéficient de remplacements plus rapides et de temps d'arrêt minimisés, ce qui en fait un complément précieux pour ceux qui comptent sur un stockage et une protection des données ininterrompus.
Présentation de l'application UniFi Drive
L'application UniFi Drive étend l'écosystème Ubiquiti en proposant une solution de stockage simplifiée qui s'intègre directement aux appareils UniFi. Conçue dans un souci de simplicité et de sécurité, l'application offre un stockage local hors cloud sans abonnement récurrent. Qu'il s'agisse de gérer des sauvegardes, de partager des fichiers ou de protéger des données critiques, UniFi Drive offre des options de stockage flexibles tout en exploitant l'interface familière d'UniFi.
