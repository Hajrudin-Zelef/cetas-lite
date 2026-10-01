---
id: collect-261001-general-networking/general-networking/fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf-1
title: "fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "intel"]
source: docs/RAG/collect-261001-general-networking/fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf.md
source_anchor: ""
source_lines: [1, 42]
sha256: 8ff516f688ef2827ccb3293fb39ead61ab954409956149ac18ea3ca9ef10d559
---

# fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf

Avec le lancement de TrueNAS SCALE 22.12.2 plus tôt cette année, l'équipe iX a clairement l'intention de parler un peu plus à l'entreprise. Une partie de la messagerie mise à l'échelle provient de nouveaux systèmes matériels et de la prise en charge de la gestion des NVDIMM et des améliorations des fonctionnalités HA et de réplication. Mais une autre poussée majeure est venue autour de la prise en charge des applications, faisant des applications TrueNAS SCALE une affaire beaucoup plus importante.
iX a expédié un nouveau Mini R, configuré avec TruNAS SCALE, et nous a demandé de fouiller un peu pour voir ce que nous pensions du catalogue d'applications amélioré et des intégrations.
TrueNAS SCALE Bluefin, version 22.12.2, introduit un nouvel ensemble d'applications dans le catalogue TrueNAS. Ces applications sont classées en trois trains en fonction de leurs niveaux de support et de maintenance. Le train « Communauté » se compose d'applications fournies par iXsystems ou la communauté, iX examinant et modérant les modifications, mais ne fournissant pas de support direct.
Le train "officiel" contient des applications régulièrement testées et entretenues par les ingénieurs de TrueNAS, avec des problèmes importants résolus rapidement. Les applications du train "Communauté" peuvent passer au train "Officiel" si elles s'avèrent de haute qualité et largement utilisées.
Le train "Enterprise" est exclusif aux appliances TrueNAS Enterprise et comprend des applications avec des fonctionnalités de niveau entreprise, soigneusement testées, maintenues et documentées pour les cas d'utilisation critiques. Les éditeurs de logiciels peuvent contacter iXsystems pour discuter de l'inclusion de leurs applications dans ce train.
En ce qui concerne le matériel sous-jacent, le TrueNAS Mini R est un NAS rackable plus grand au sein de la famille Mini ; vous trouverez plus d’informations sur le site Web de TrueNAS.
Bien que physiquement plus grand que les autres Minis, le TrueNAS Mini R utilise cette empreinte pour le stockage. Le serveur 2U à faible profondeur offre 12 baies de stockage SATA 3.5″ pouvant être utilisées avec un mélange de disques SSD et de disques durs. Le système commence à 1848 $ et est assez configurable en fonction des besoins des clients.
Spécifications du TrueNAS Mini R
| Fonctionnalité | TrueNAS Mini R | 
|---|---|
| Châssis | Boîtier à 12 baies - Conception silencieuse pour le bureau : 45 dB en veille, 52 dB en crête | 
| Baies de disques | 12 baies de lecteur SATA 3.5" remplaçables à chaud (adaptateurs 3.5" à 2.5" en option disponibles) | 
| Capacité brute maximale | Jusqu'à 216 TB | 
| Processeur | Processeur Intel octa-core C3758 | 
| Mémoire | 32 Go DDR4 avec ECC (extensible à 64 Go) | 
| RAID | OpenZFS : bande (RAID0), miroir multidisque (RAID10), parité RAIDZ1 (RAID5), RAIDZ2 (RAID6) et RAIDZ3 (triple parité) | 
| Gestion de disque | Disques remplaçables à chaud, Bad Block Scan + HDD SMART, prise en charge du montage ISO, chiffrement de disque accéléré par le matériel | 
| Réseau | Standard : 2 ports LAN Ethernet RJ45 1/10GBaseT Port IPMI RJ45 dédié (gestion du matériel à distance) 2 x carte d'extension SFP+ 10G (en option) | 
| Ports USB | 1 port USB 3.0 (arrière) 2 ports USB 2.0 (arrière) | 
| Cache de lecture/écriture | (Facultatif) Améliorez les performances en ajoutant un cache de lecture hautes performances dédié (L2ARC) ou en ajoutant un cache d'écriture hautes performances dédié (ZIL/SLOG) | 
| Extension PCIe | 1 x PCI Express 3.0 x 4 | 
| Alimentation | 100 V à 240 V CA, 50/60 HZ, monophasé | 
| Consommation électrique (maximale) | Sans disque : 63 W, avec disques et carte d'extension 10 G : 167 W | 
| Gestion de l'énergie | Mise sous/hors tension à distance (IPMI), réponse au signal de l'onduleur et alertes | 
| Interface de contrôle utilisateur | Navigateur Web et gestion du matériel à distance (IPMI) | 
| Dimensions (L x P x H) | 17.2″ x 21″ x 3.5″ / 437 × 533 × 89 mm | 
| Poids (pas de disques) | 18.7 livres / 8.5 kg | 
| Garantie limitée | Garantie d'un an incluse - Garantie de 1 ans en option à l'achat. La garantie logicielle nécessite un enregistrement sur portail.ixsystems.com. | 
| Accessoires | Guide de configuration de base ; 2 clés de lunette ; sac de vis HDD ; 2 câbles en cuivre cat7 de 6 m ; Facette; 4 pieds en caoutchouc adhésifs ; Kit de rails courts : profondeur de montage en rack de 19″ à 26.6″ ; Kit de rails longs en option disponible : profondeur de montage en rack de 26.5″ à 36.4″ | 
ÉCHELLE TrueNAS avec LINUX
TrueNAS SCALE sur le Mini R offre suffisamment de puissance pour exécuter un partage ZFS entièrement déployé avec la possibilité d'exécuter plusieurs conteneurs. Bien que le processeur ne soit pas conçu pour être une centrale électrique (processeur Intel Atom C3758 à 2.20 GHz), il s'agit d'une puce à huit cœurs à huit threads avec un cache L16 de 2 Mo.
Il s'agit d'une puce 14 nm basée sur l'architecture Intel Denverton, lancée au troisième trimestre 2017, qui affiche une faible consommation (TDP de 25 W) et prend en charge la mémoire vive DDR4 ECC (qu'elle exploite pleinement). Notre modèle de test est équipé de 64 Go de RAM. Pour répondre aux besoins de traitement et d'accélération, le TrueNAS requiert une quantité importante de RAM, pouvant atteindre 5 Go par To pour la déduplication. La configuration minimale est de 8 Go de RAM pour un maximum de huit disques, et il faut ajouter 1 Go par disque supplémentaire. Pour une configuration de base, il est conseillé d'opter pour 16 Go ou 32 Go de RAM.
Notre expérience a montré que le processeur était plus que suffisant pour exécuter TrueNAS et suivre plusieurs conteneurs et même certaines machines virtuelles.
TrueNAS SCALE offre toujours toutes les fonctionnalités ZFS et de partage de TrueNAS CORE tout en ajoutant une nouvelle prise en charge des applications conteneurisées via Docker et le catalogue d'applications organisé.
Nous avons même obtenu une installation complète de Windows Server 2022 pour installer et fonctionner dans l'espace de virtualisation. Certes, c'était un peu plus rugueux qu'un hyperviseur plus mature.
Nous avons constaté que les pilotes de Windows Server 2022 devaient être chargés à partir d'un ISO monté séparément, mais il s'intégrait bien aux ressources embarquées et était immédiatement en ligne.
Applications TrueNAS SCALE
Les applications TrueNAS SCALE sont très similaires à celles prises en charge par TrueNAS CORE. La sélection de l'application dans la liste et son déploiement sont à peu près identiques à l'expérience de déploiement et de gestion de TrueNAS CORE.
Nous avons essayé de déployer "Home Assistant" et "Grafana", et c'est aussi simple que de choisir, cliquer et déployer.
Une fois lancé, il affiche le numéro de port de l'application en cours d'exécution et vous pouvez vous connecter via un navigateur sur le réseau. Nous devions simplement ouvrir un navigateur, entrer l'adresse IP du NAS et le numéro de port de l'application, et chaque nouvelle application nous accueillait avec un écran de configuration et de configuration.
Ce fut une expérience presque identique à CORE et très similaire à de nombreuses autres offres NAS grand public.
Cependant, ce qui est nouveau, c'est la possibilité d'ajouter des conteneurs non répertoriés dans la liste des applications. Le déploiement de conteneurs Docker était simple si vous avez une certaine expérience avec Docker.
