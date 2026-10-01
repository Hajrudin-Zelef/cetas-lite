---
id: collect-261001-general-networking/general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026-5
title: "Téléchargement de l'ISO Proxmox VE 9.1"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["agent", "amd", "datacenter", "gpu", "intel", "memory", "nvidia"]
source: docs/RAG/collect-261001-general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026.md
source_anchor: ""
source_lines: [366, 457]
sha256: a21ad9830f49fc2243f8513ad203dcefbafec9b804bbe9ca6dfe5b8dc7483291
---

# Téléchargement de l'ISO Proxmox VE 9.1

| Critère | Proxmox VE 9.1 | VMware ESXi 8.0 U3 | Hyper-V Server 2025 | 
|---|---|---|---|
| Modèle de licence | AGPLv3, abo facultatif | Abonnement obligatoire (Broadcom) | Inclus avec Windows Server | 
| Prix entrée (CPU/an) | ~115 € HT (Community) | VVF/VCF bundle, sur devis | Inclus Windows Datacenter | 
| Hyperviseur | KVM + LXC | VMkernel propriétaire | Hyper-V propriétaire | 
| Filesystem natif | ZFS, ext4, XFS, BTRFS | VMFS6, vSAN | NTFS, ReFS, S2D | 
| Stockage HCI | Ceph Squid intégré | vSAN (licence séparée) | Storage Spaces Direct | 
| Cluster max | 32 nœuds | 96 nœuds (vCenter) | 64 nœuds | 
| Conteneurs natifs | LXC + OCI | vSphere Pods (Tanzu) | Windows Containers | 
| Sauvegarde | PBS 4.1 inclus | Veeam, vSphere Replication | Veeam, Azure Backup | 
| API automatisation | REST + Ansible + Terraform | vSphere API + PowerCLI | WMI + PowerShell | 

## Pièges classiques et erreurs fréquentes à éviter

Cinq pièges reviennent dans 80 % des dépannages que les administrateurs Proxmox français rencontrent en première installation, selon les fils du forum officiel et de la communauté r/Proxmox. Les anticiper économise des heures de diagnostic.

### Piège n°1 : ZFS sur RAID matériel

Si vous laissez votre carte PERC ou MegaRAID en mode RAID5/6, ZFS ne verra qu’un seul disque virtuel et perdra toute capacité d’auto-réparation. Conséquence : un secteur silencieusement corrompu sur le RAID matériel ne sera jamais détecté par les checksums ZFS car celui-ci croira lire un disque sain. Basculez systématiquement le contrôleur en mode HBA, IT ou JBOD passthrough.

### Piège n°2 : pas de QEMU Guest Agent

Sans agent installé dans la VM, le bouton *Shutdown* envoie un signal ACPI que de nombreuses configurations ignorent. Résultat : extinction brutale et risque de perte de données. Pire, les sauvegardes en mode snapshot ne peuvent pas freezer le filesystem invité, ce qui produit des images potentiellement incohérentes. Installez le paquet `qemu-guest-agent` sur tous les invités Linux et le service *QEMU Guest Agent* depuis le virtio-win.iso sur Windows.

### Piège n°3 : MTU incohérent sur Corosync

Si certains nœuds du cluster utilisent un MTU 1500 et d’autres 9000 (jumbo frames pour Ceph), le bus Corosync peut perdre des paquets et déclencher des fences intempestifs. Réservez un VLAN ou un switch dédié à Corosync, fixez le MTU explicitement (`ip link set vmbr0 mtu 1500`) et activez deux liens redondants (`--link0` et `--link1`).

### Piège n°4 : cache ZFS ARC qui dévore la RAM

Par défaut, ZFS utilise jusqu’à 50 % de la RAM totale comme cache ARC. Sur un nœud avec 32 Go destiné à 28 Go de VM, cela génère du swap et de la latence. Limitez l’ARC en éditant `/etc/modprobe.d/zfs.conf` avec `options zfs zfs_arc_max=4294967296` (4 Go), puis lancez `update-initramfs -u && reboot`.

### Piège n°5 : sauvegardes sur le même stockage que les VM

Stocker les vzdump dans `local` ou dans le même pool que les VM est une fausse sécurité : la perte du SSD système emporte aussi les sauvegardes. Configurez impérativement un PBS distant ou au minimum un partage NFS sur un NAS différent. Pour le hors site, programmez une synchronisation PBS-to-PBS via `proxmox-backup-manager sync-job create`.

## Astuces avancées pour optimiser les performances

Au-delà de l’installation par défaut, plusieurs leviers permettent d’extraire un gain de performance significatif. Sur un nœud Xeon Gold 6526Y avec 256 Go de RAM, l’activation des **hugepages** 1 Go pour une VM PostgreSQL 17 a réduit la latence de 18 % en lecture aléatoire et de 31 % en écriture séquentielle dans nos tests internes du laboratoire. Activez via `qm set 100 --hugepages 1024 --memory 16384`.

Le **passthrough PCIe** permet d’attribuer un GPU NVIDIA RTX 6000 Ada ou une carte réseau Mellanox ConnectX-6 directement à une VM, sans pénalité de virtualisation. Activez d’abord `intel_iommu=on iommu=pt` dans `/etc/default/grub`, regénérez avec `update-grub`, identifiez le PCIe ID via `lspci -nn`, puis attachez : `qm set 100 --hostpci0 01:00,pcie=1,x-vga=1`. Sur les workloads IA, c’est la seule façon d’exposer un GPU complet sans MIG ou vGPU.

Pour les bases de données critiques, désactivez la **balloon** : `qm set 100 --balloon 0`. La RAM dynamique sacrifie la prévisibilité des temps de réponse. À la place, dimensionnez précisément la mémoire et ajoutez du swap ZFS dans la VM. Activez aussi le *NUMA awareness* sur les machines bi-socket (`--numa 1 --sockets 2`) : sans cela, une VM peut traverser le QPI Intel à chaque accès RAM, dégradant la latence de 25 à 40 ns par opération.

## Sécurisation : firewall, 2FA et certificats Let’s Encrypt

Une installation Proxmox exposée sans durcissement est une cible de choix pour les scans automatisés. Trois mesures au minimum sont indispensables avant toute mise en production : activer le firewall intégré, exiger l’authentification à deux facteurs et remplacer le certificat auto-signé par un certificat Let’s Encrypt valide.

```
# Activation du firewall au niveau Datacenter
pve-firewall start
# Autoriser SSH et l'interface web depuis votre IP fixe
cat > /etc/pve/firewall/cluster.fw <<EOF
[OPTIONS]
enable: 1
[RULES]
IN ACCEPT -source 198.51.100.42 -p tcp -dport 22 -log nolog
IN ACCEPT -source 198.51.100.42 -p tcp -dport 8006 -log nolog
IN ACCEPT -p tcp -dport 22 -log warning
IN DROP -log warning
EOF
# Activation 2FA TOTP pour root
pveum user modify root@pam --append-tfa totp
# Generation d'un certificat Let's Encrypt via le wizard integre
pvenode acme account register default [email protected]
pvenode config set --acme domains=pve01.lab.example.fr
pvenode acme cert order
# Le certificat est renouvele automatiquement tous les 60 jours
systemctl status pve-daily-update.timer
```
Pour les déploiements multi-tenants, créez des realms dédiés (LDAP vers Active Directory ou Microsoft Entra ID via OIDC) et appliquez des permissions granulaires : un utilisateur *backup-operator* n’a besoin que du privilège `VM.Backup`, pas de `VM.Allocate`. La granularité monte jusqu’au niveau VM individuelle via le système de pools.

## Tableau des tarifs des abonnements Proxmox VE

| Niveau | Prix CPU/an HT | Tickets support | SLA | Cas d’usage | 
|---|---|---|---|---|
| Aucun (No-Subscription) | 0 € | Forum communautaire | Aucun | Lab, formation, dev | 
| Community | ~115 € à 120 € | 3 tickets/an | Dépôt Enterprise stable | PME, sites secondaires | 
| Basic | ~355 € à 370 € | 3 tickets/an | Réponse 1 jour ouvré | Production non critique | 
| Standard | ~550 € | 10 tickets/an | Réponse 4 heures ouvrées | Production critique | 
| Premium | ~1 100 € | Illimité | Réponse 2 heures, 24/7 | Mission-critical 24/7 | 

Tous les niveaux donnent accès au même code source et aux mêmes fonctionnalités logicielles : la différence porte uniquement sur le support et le dépôt utilisé. La règle de Proxmox impose de souscrire un abonnement de niveau identique sur tous les nœuds d’un même cluster, faute de quoi vous mélangez des dépôts Enterprise et No-Subscription, source de bugs subtils. Pour un cluster homogène à 3 nœuds bi-socket en niveau Basic, comptez 3 × 2 × 365 € = 2 190 € HT par an, soit cinq à dix fois moins qu’un VVF Broadcom équivalent selon les négociations rapportées par CRN et The Register en 2025.

## Dépannage et résolution des erreurs courantes

Voici les huit incidents les plus fréquents rencontrés en première installation et leurs solutions vérifiées sur Proxmox VE 9.1.

### Erreur 1 : « unable to create kvm » au lancement de VM

Cause : virtualisation matérielle désactivée dans le BIOS. Vérifiez avec `egrep -c '(vmx|svm)' /proc/cpuinfo` : un résultat de 0 confirme le problème. Solution : redémarrez en BIOS, activez Intel VT-x ou AMD-V, sauvegardez et rebootez.

