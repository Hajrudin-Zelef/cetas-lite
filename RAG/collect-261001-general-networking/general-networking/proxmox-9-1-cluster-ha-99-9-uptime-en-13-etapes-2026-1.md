---
id: collect-261001-general-networking/general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026-1
title: "Télécharger l'ISO et la somme SHA-256"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel"]
dates: []
keywords: ["amd", "distribution", "gpu", "intel", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 31]
sha256: a2f8c858611281af93aa3e0caf5af9641014719287011f155c9c67812831ccc2
---

# Télécharger l'ISO et la somme SHA-256

**Proxmox** s’est imposé en 2026 comme l’alternative numéro un à VMware ESXi pour les homelabs, PME et hébergeurs européens cherchant à fuir le modèle d’abonnement Broadcom. Ce tutoriel Proxmox VE 9.1 vous guide pas à pas, du téléchargement de l’ISO à la mise en cluster trois nœuds avec haute disponibilité, en couvrant ZFS 2.3.3, QEMU 10.0.2, LXC 6.0.4, Ceph Squid 19.2.3 et la nouvelle SDN EVPN. Comptez environ 90 minutes pour suivre les 13 étapes sur un serveur dédié ou un mini-PC type Intel N100 ; publié le 17 avril 2026, ce guide reflète l’état exact de la stack au lendemain de la mise à jour Proxmox Backup Server 4.2-1, dont l’ISO restait encore listée comme la dernière version stable de PBS sur le site officiel en avril 2026. Depuis, Proxmox Server Solutions a officiellement annoncé la disponibilité de VE 9.2 dans ses communications le 28 mai 2026 – une version sortie le 21 mai 2026 sous forme d’un ISO allégé à 1,71 Gio, bâtie sur un tout nouveau noyau Linux 7.0 et dotée d’un équilibreur de charge dynamique pour les clusters, selon ComputingForGeeks – mais les 13 étapes ci-dessous restent la meilleure porte d’entrée avant de basculer vers cette version plus récente.

Proxmox VE 9.1-1, sorti le 19 novembre 2025, s’appuie sur Debian 13.2 « Trixie » avec un noyau Linux 6.17, une évolution par rapport à VE 9.0 – sortie le 5 août 2025 avec le socle kernel 6.14.8, QEMU 10.0.2 et ZFS 2.3.3 détaillé plus loin dans ce guide, et présentée par l’éditeur autrichien comme le jalon marquant ses 20 ans d’innovation – auquel la 9.1 ajoute surtout le support de la virtualisation imbriquée (nested virtualization) et d’Intel TDX pour isoler des VM confidentielles ; le tracker endoflife.date référence d’ailleurs déjà la 9.2, publiée en mai 2026, comme la version stable la plus récente de toute la branche 9.x – une variante arm64 de cet ISO 9.2-1, pesant 1,50 Gio, a même été mise à jour le 5 août 2026 pour les architectures ARM. Côté virtualisation, il combine KVM, LXC et un système de stockage natif ZFS qui couvrent à la fois les machines virtuelles Windows, les containers Linux ultra-légers et les pools redondants. Pour l’écosystème français – administrations qui migrent dans le cadre de la stratégie de souveraineté numérique de la DINUM et PME confrontées à la hausse des prix VMware – la promesse est claire : une stack 100 % open source, pas de licence par cœur, et un panneau web complet sur le port 8006.

## Pourquoi choisir Proxmox VE 9.1 en 2026

Le choix d’un hyperviseur en 2026 se résume désormais à trois variables : le coût des licences, la maîtrise des données et la richesse fonctionnelle. Proxmox VE coche les trois cases. Là où VMware vSphere Foundation impose un abonnement minimum à 350 USD par cœur et par an depuis le rachat Broadcom de novembre 2023, Proxmox VE reste gratuit sans limitation fonctionnelle. Le dépôt « Enterprise » payant n’apporte qu’un support officiel et des paquets retardés pour stabilité, pas de fonctionnalités exclusives. Les administrateurs français rencontrés sur la liste pve-user signalent que la bascule d’un cluster ESXi de 24 cœurs vers Proxmox VE divise la facture annuelle par 14 environ, en éliminant aussi vCenter et vSAN – d’autant qu’un abonnement Proxmox démarre à seulement 115 euros par cœur et par an selon le communiqué de presse officiel accompagnant le lancement de VE 9.0 le 5 août 2025, contre un plancher de 350 USD chez Broadcom.

La deuxième raison tient à l’intégration native de ZFS, Ceph et LXC. Aucune autre distribution hyperviseur ne livre dans un seul installateur un système de fichiers transactionnel avec compression et déduplication, un stockage objet distribué prêt pour la production, et une couche container nettement plus légère que Docker pour les charges Linux. La version 9.1 ajoute également une couche SDN supportant VXLAN et EVPN, ce qui permet de bâtir des réseaux multi-tenants sans appareil tiers – un cas d’usage typique chez les hébergeurs français qui revendent du VPS ; Proxmox VE 9.2, disponible depuis mai 2026, complète encore cette couche SDN avec un nouveau fabric WireGuard natif.

Enfin, la communauté francophone est particulièrement active. Le forum officiel compte plus de 100 000 sujets, et des intégrateurs comme Proxmox France ou Tranquil IT proposent des contrats de support en français à des tarifs bien en deçà de ceux pratiqués par Broadcom. Pour un homelab personnel, vous pouvez démarrer sur un mini-PC Intel N100 à 180 euros et obtenir une plateforme qui héberge confortablement dix VM et vingt conteneurs LXC. Pour un projet PME, un cluster de trois nœuds Dell PowerEdge T350 d’occasion suffit à délivrer de la haute disponibilité avec migration à chaud, le tout pour moins de 4 000 euros.

## Prérequis matériels et logiciels

Avant de lancer l’installation, vérifiez méthodiquement votre matériel. Proxmox VE 9.1 exige un processeur 64 bits avec virtualisation matérielle activée dans le BIOS. Cela correspond à Intel VT-x sur les Core et Xeon, et AMD-V sur les Ryzen et EPYC. Pour le passthrough PCI ou GPU – utile si vous prévoyez d’attribuer une RTX 4090 à une VM de poste de travail – il vous faut en plus IOMMU (Intel VT-d ou AMD-Vi). Sur les cartes mères grand public, ces options s’activent généralement sous le menu « Advanced CPU Configuration » du BIOS UEFI.

Côté mémoire vive, le minimum théorique est de 2 Gio mais aucun usage sérieux n’est possible en dessous de 16 Gio. La règle empirique consiste à compter 2 Gio pour Proxmox lui-même, 8 Gio pour ZFS ARC si vous utilisez un pool de plus de 4 To, puis la somme de la RAM attribuée à chaque VM ou container. Le disque système doit dépasser 32 Gio mais on recommande aujourd’hui un SSD NVMe de 256 Gio pour Proxmox plus un pool ZFS séparé pour les disques virtuels. Une carte réseau gigabit suffit pour un nœud unique, deux NIC sont indispensables en cluster.

| Composant | Minimum | Recommandé homelab | Production PME | Cluster HA | 
|---|---|---|---|---|
| CPU | Intel/AMD 64 bits, VT-x/AMD-V | Intel N100 4 cœurs | Xeon Silver ou EPYC 7313 | Xeon Gold 6338 ou EPYC 7543 | 
| RAM | 2 Gio | 16 Gio DDR4 | 64 Gio DDR4 ECC | 128 à 256 Gio DDR5 ECC | 
| Disque système | 32 Gio | SSD NVMe 256 Gio | 2× SSD 480 Gio en miroir ZFS | 2× NVMe 960 Gio + Ceph | 
| Stockage VM | Inclus système | SSD SATA 1 To | 4× HDD 4 To en RAIDZ2 | Ceph 12 OSD minimum | 
| Réseau | 1× 1 GbE | 2× 2,5 GbE | 2× 10 GbE LACP | 4× 10 GbE + 2× 25 GbE Ceph | 
| Onduleur | Facultatif | Recommandé 600 VA | Obligatoire 1500 VA | Redondant 3000 VA | 

Côté logiciel, vous n’aurez besoin que d’une image ISO officielle, d’une clé USB d’au moins 4 Gio et d’un outil de gravure comme Balena Etcher, Rufus ou la commande dd sous Linux. La présente version 9.1-1 pèse précisément 1,83 Gio, contre 1,71 Gio pour l’ISO 9.2-1 mis à jour le 21 mai 2026 selon la page de téléchargement officielle de Proxmox. N’utilisez jamais l’image marquée 8.4-1 comme base d’un nouveau cluster : selon endoflife.date, toute la branche Proxmox VE 8 (lancée en juin 2023) arrive en fin de support le 31 août 2026, soit quelques semaines seulement après la publication de ce guide, et elle ne reçoit déjà plus de nouvelles fonctionnalités. Le hash SHA-256 officiel de la 9.1-1 commence par 6d8f5afc78c0c66812d7272cde7c8b98 – vérifiez-le impérativement avant d’écrire sur la clé.

## Étape 1 : Télécharger et vérifier l’ISO Proxmox VE 9.1

