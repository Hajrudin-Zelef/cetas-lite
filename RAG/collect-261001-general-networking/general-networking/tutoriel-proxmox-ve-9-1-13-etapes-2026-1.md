---
id: collect-261001-general-networking/general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026-1
title: "Téléchargement de l'ISO Proxmox VE 9.1"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel"]
dates: []
keywords: ["amd", "distribution", "gpu", "intel", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 29]
sha256: 2a49179b5b60280b534c7a85cbdd731938c8749ca7f0f6a704496384d23f6a40
---

# Téléchargement de l'ISO Proxmox VE 9.1

Avec la fin annoncée des licences perpétuelles VMware par Broadcom et l’envolée des coûts vSphere — parfois multipliée par 10 selon les rapports d’utilisateurs européens en 2025 — l’**installation Proxmox** est devenue le réflexe numéro un des administrateurs systèmes français à la recherche d’une alternative open source. Proxmox VE 9.1, publié le 19 novembre 2025 par Proxmox Server Solutions GmbH, embarque le noyau Linux 6.17, QEMU 10.0.2, ZFS 2.3.3 et Ceph Squid 19.2.3, le tout sur une base Debian 13 « Trixie » héritée de la version majeure 9.0 sortie le 5 août 2025. Depuis, Proxmox Server Solutions a publié le 21 mai 2026 la version 9.2, bâtie sur Debian 13.5 « Trixie » avec le noyau Linux 7.0, QEMU 11.0, LXC 7.0 et ZFS 2.4 selon le guide d’installation 2026 du Devaubree Blog, et référencée sous la version 9.2-1 (et 9.2-1-arm64 pour les architectures ARM) sur la page officielle de téléchargement ; la documentation Proxmox — mise à jour en août 2026 — confirme que la 9.2 reste la version courante de la branche, preuve de la cadence soutenue de la 9.x. Ce tutoriel détaille en 13 étapes comment installer, configurer et exploiter un hyperviseur Proxmox VE 9.1 en production, depuis le téléchargement de l’ISO jusqu’au cluster multi-nœuds avec sauvegardes automatisées vers un Proxmox Backup Server 4.1.

Publié le 26 avril 2026, ce guide cible aussi bien les administrateurs qui migrent depuis VMware ESXi que les autodidactes qui découvrent la virtualisation — une démarche que confirme le blog francophone Oleks IT, qui a détaillé le 4 mai 2026 sa propre installation de Proxmox VE 9 avec les versions mises à jour de QEMU 9.x et LXC 6.x. Le guide anglophone de Mr.PlanB, publié le même mois, arrive à la même conclusion : la version en cours est Proxmox VE 9.2-1, sortie en mai 2026 et toujours bâtie sur Debian 13 « Trixie ». Vous y trouverez les commandes complètes, des extraits de configuration prêts à copier, les pièges classiques (BIOS UEFI, ZFS sur RAID matériel, IOMMU pour le passthrough GPU) et un projet final : un cluster à trois nœuds avec haute disponibilité, stockage Ceph répliqué et sauvegardes incrémentielles vers un PBS dédié. Comptez environ 90 minutes pour la première installation et 3 à 4 heures pour le projet complet.

## Pourquoi choisir Proxmox VE 9.1 en 2026

Proxmox Virtual Environment est un hyperviseur open source bare-metal qui combine deux technologies de virtualisation Linux : **KVM** pour les machines virtuelles complètes et **LXC** pour les conteneurs système. La distribution est éditée à Vienne par Proxmox Server Solutions GmbH depuis 2008, sous licence AGPLv3. Contrairement à VMware ESXi, qui depuis le rachat par Broadcom en novembre 2023 a vu ses licences perpétuelles supprimées et ses tarifs renégociés, Proxmox conserve son modèle *open core* : le code reste libre, les abonnements ne couvrent que le support et l’accès au dépôt « Enterprise » stable.

La version 9.1 publiée le 19 novembre 2025 introduit des nouveautés majeures pour les exploitants. Les conteneurs LXC peuvent désormais consommer des images au format **OCI** (Open Container Initiative), ce qui rapproche Proxmox de l’écosystème Docker et Kubernetes. La virtualisation imbriquée a été affinée : un administrateur peut activer ou désactiver les extensions VMX/SVM par machine virtuelle, condition nécessaire pour tester un cluster Kubernetes, un nœud Hyper-V ou un autre Proxmox dans une VM. L’état TPM est désormais stocké dans un fichier qcow2 distinct, ce qui simplifie les sauvegardes et la migration des VM Windows 11. Enfin, le sous-système SDN (Software-Defined Networking), dont les bases — SDN Fabrics et les snapshots thick-LVM — avaient été posées dès la 9.0 du 5 août 2025 sur noyau Linux 6.14, reçoit des améliorations majeures dans la 9.1, notamment la gestion fine des VLAN par zone et l’intégration avec des contrôleurs externes ; la feuille de route 9.x s’est poursuivie avec la 9.2 (noyau Linux 7.0), qui ajoute un équilibreur de charge dynamique pour le HA Manager, selon le calendrier de fonctionnalités publié par ComputingForGeeks.

En face, Proxmox VE 8 — dont la dernière mouture, la 8.4, publiée le 9 avril 2025 sur une base Debian 12.10 et noyau Linux 6.8.12-9, demeure selon endoflife.date la version de référence de la branche 8.x — n’est pris en charge que jusqu’au 31 août 2026, après quoi seules les versions 9.x recevront des mises à jour de sécurité. Les utilisateurs qui n’ont pas encore migré disposent donc d’une fenêtre resserrée pour basculer, d’autant que Proxmox documente désormais un chemin de mise à niveau direct de la 8.4 vers la 9.2 via les commandes `pve8to9` puis `apt full-upgrade`, détaillé par Datazone le 21 mai 2026. La rolling release suit Debian : Proxmox VE 9 est aligné sur Debian 13.2 Trixie, ce qui garantit dix ans de support LTS via le projet upstream.

## Prérequis matériels et logiciels

Avant de lancer l’installation, vérifiez que votre matériel répond aux exigences. Proxmox VE 9.1 reste accessible : la documentation officielle parle d’un minimum de 2 Go de RAM et d’un processeur 64 bits, mais ces valeurs ne couvrent que l’hyperviseur lui-même. Le communiqué de Proxmox Server Solutions du 5 août 2025, accompagnant la sortie de la 9.0 sur base Debian 13 Trixie, confirmait déjà que l’ISO s’installe sans difficulté sur serveur bare-metal, sans dépendance à un hyperviseur hôte préexistant. En production, les administrateurs expérimentés recommandent au moins 16 Go de RAM, un processeur récent avec extensions de virtualisation (Intel VT-x ou AMD-V), 256 Go de stockage SSD et une carte réseau gigabit. Pour exploiter Ceph ou ZFS, ajoutez 1 Go de RAM par To de stockage et privilégiez les disques SSD enterprise avec capacité PLP (Power Loss Protection).

| Composant | Minimum | Recommandé | Production | 
|---|---|---|---|
| CPU | 64 bits, VT-x/AMD-V | 4 cœurs récents | 16+ cœurs Xeon Scalable / EPYC | 
| RAM | 2 Go | 16 Go ECC | 64 à 512 Go ECC | 
| Disque système | 32 Go HDD | SSD 256 Go | 2 × NVMe enterprise en miroir ZFS | 
| Disque VM | SSD SATA | NVMe consommateur | NVMe enterprise PLP, RAID-Z2 ou Ceph | 
| Réseau | 1 GbE | 2 × 1 GbE bond | 2 × 25 GbE pour Ceph + 2 × 10 GbE management | 
| BIOS | UEFI ou Legacy | UEFI + Secure Boot off | UEFI + IOMMU activé pour passthrough | 

Côté logiciel, vous aurez besoin d’un poste client avec un navigateur récent (Firefox 128 ESR ou Chromium 130+), un client SSH (OpenSSH sous Linux/macOS, PuTTY ou Windows Terminal sous Windows) et l’outil `balenaEtcher` ou `dd` pour graver l’ISO sur clé USB. Les commandes de ce tutoriel ont été testées sur un MacBook Pro M2, un poste Ubuntu 24.04 LTS et un Windows 11 24H2. Aucun compte payant n’est requis : la communauté gratuite suffit pour démarrer, l’abonnement Community Edition à environ 115 € HT par CPU et par an n’apporte que l’accès au dépôt stable.

## Étape 1 : Télécharger l’ISO Proxmox VE 9.1

