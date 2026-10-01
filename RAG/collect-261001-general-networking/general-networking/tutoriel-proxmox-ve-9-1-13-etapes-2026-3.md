---
id: collect-261001-general-networking/general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026-3
title: "Téléchargement de l'ISO Proxmox VE 9.1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr", "memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026.md
source_anchor: ""
source_lines: [87, 229]
sha256: e6ed6b95cf3b3872470772570f22a633a14b5440aa54d7cc6143dd6c6f045e9d
---

# Téléchargement de l'ISO Proxmox VE 9.1

L’écran d’accueil affiche immédiatement un bandeau rouge : *« No valid subscription »*. Sans abonnement payant, Proxmox interroge par défaut le dépôt Enterprise auquel vous n’avez pas accès, ce qui bloque les mises à jour. Désactivez ce dépôt et activez le dépôt « No-Subscription » qui contient les mêmes paquets, simplement testés moins longtemps.

```
# Desactivation du depot Enterprise
sed -i 's/^deb/#deb/' /etc/apt/sources.list.d/pve-enterprise.list
sed -i 's/^deb/#deb/' /etc/apt/sources.list.d/ceph.list
# Ajout du depot No-Subscription
echo "deb http://download.proxmox.com/debian/pve trixie pve-no-subscription" \
  > /etc/apt/sources.list.d/pve-no-subscription.list
# Mise a jour et upgrade complet
apt update
apt full-upgrade -y
# Redemarrage si un nouveau noyau a ete installe
reboot
```
Pour supprimer définitivement le bandeau, vous pouvez patcher le fichier `/usr/share/javascript/proxmox-widget-toolkit/proxmoxlib.js` (recherchez la chaîne « No valid subscription »). Cette modification doit être réappliquée après chaque mise à jour majeure : elle n’est ni recommandée ni interdite, juste un confort visuel.

## Étape 6 : Créer votre première machine virtuelle KVM

Téléchargez d’abord une ISO d’invité dans le stockage `local`. Cliquez sur *local (pve01)* dans la barre latérale, puis sur l’onglet *ISO Images* et bouton *Download from URL*. Pour Ubuntu 24.04.2 LTS Server par exemple, indiquez `https://releases.ubuntu.com/24.04.2/ubuntu-24.04.2-live-server-amd64.iso`. Le téléchargement direct depuis l’interface évite les transferts via votre poste local et utilise la bande passante du serveur.

Cliquez ensuite sur *Create VM* en haut à droite. L’assistant guide en sept onglets : General (ID 100, nom `web01`), OS (sélection de l’ISO Ubuntu), System (machine type q35, BIOS OVMF UEFI, contrôleur SCSI VirtIO SCSI single, agent QEMU coché), Disks (50 Go en VirtIO Block sur stockage local-zfs, format raw ou qcow2, cache none, IO thread coché, SSD emulation), CPU (4 cores, type host pour les meilleures performances, flags AES si chiffrement disque), Memory (8192 Mo avec ballooning), Network (bridge vmbr0, modèle VirtIO, firewall coché). Validez par *Finish*.

```
# Creation equivalente en CLI
qm create 100 \
  --name web01 \
  --ostype l26 \
  --machine q35 \
  --bios ovmf \
  --efidisk0 local-zfs:1,format=raw,efitype=4m,pre-enrolled-keys=1 \
  --scsihw virtio-scsi-single \
  --scsi0 local-zfs:50,iothread=1,ssd=1,discard=on \
  --cdrom local:iso/ubuntu-24.04.2-live-server-amd64.iso \
  --boot order=scsi0;ide2 \
  --cores 4 \
  --sockets 1 \
  --cpu host \
  --memory 8192 \
  --balloon 4096 \
  --net0 virtio,bridge=vmbr0,firewall=1 \
  --agent enabled=1
# Demarrage de la VM
qm start 100
# Console serie dans le terminal
qm terminal 100
```
Après installation de l’OS invité, n’oubliez pas d’installer le QEMU Guest Agent dans la VM. Sous Ubuntu : `apt install qemu-guest-agent && systemctl enable --now qemu-guest-agent`. Cet agent permet à Proxmox de récupérer l’IP de la VM, de freezer le filesystem pendant les snapshots et de déclencher un arrêt propre via le bouton *Shutdown*.

## Étape 7 : Conteneurs LXC et nouveautés OCI de Proxmox 9.1

Les conteneurs LXC partagent le noyau de l’hôte : ils consomment dix fois moins de RAM qu’une VM équivalente et démarrent en moins d’une seconde. Proxmox VE 9.1 introduit une fonctionnalité majeure : les **templates OCI**, qui permettent de consommer directement les images Docker Hub ou GitHub Container Registry via LXC, en plus du format historique `.tar.zst` de Linux Containers.

```
# Mise a jour du catalogue de templates communautaires
pveam update
# Liste des templates Debian disponibles
pveam available --section system | grep debian
# Telechargement du template Debian 13
pveam download local debian-13-standard_13.0-1_amd64.tar.zst
# Creation d'un conteneur LXC
pct create 200 local:vztmpl/debian-13-standard_13.0-1_amd64.tar.zst \
  --hostname db01 \
  --cores 2 \
  --memory 2048 \
  --swap 512 \
  --rootfs local-zfs:8 \
  --net0 name=eth0,bridge=vmbr0,ip=dhcp,firewall=1 \
  --features nesting=1,keyctl=1 \
  --unprivileged 1 \
  --onboot 1 \
  --password $(openssl rand -base64 12)
# Demarrage et entree dans le conteneur
pct start 200
pct enter 200
```
Le flag `--unprivileged 1` est crucial en production : il fait tourner le conteneur avec un mapping UID/GID décalé, de sorte qu’un root échappé ne soit pas root sur l’hôte. Le flag `nesting=1` est nécessaire si vous comptez exécuter Docker dans le conteneur ou imbriquer un autre LXC. Pour les images OCI, la syntaxe est légèrement différente :

```
# Import d'une image OCI depuis un registre (Proxmox 9.1+)
pct create 201 oci://docker.io/library/nginx:1.27 \
  --hostname web-oci \
  --cores 1 \
  --memory 512 \
  --rootfs local-zfs:4 \
  --net0 name=eth0,bridge=vmbr0,ip=dhcp \
  --unprivileged 1
```
## Étape 8 : Configurer le stockage ZFS, NFS et iSCSI

Proxmox supporte une dizaine de backends de stockage : Directory, LVM, LVM-Thin, ZFS, ZFS over iSCSI, NFS, CIFS/SMB, RBD (Ceph), Cephfs, GlusterFS et BTRFS. Le choix dépend de l’usage. Pour un nœud isolé, ZFS local suffit. Pour un cluster sans Ceph, partagez un volume NFS ou iSCSI depuis un NAS Synology, TrueNAS ou QNAP. Pour un cluster hyper-convergé, Ceph est la référence, mais demande au moins 3 nœuds et 10 GbE.

```
# Creation d'un pool ZFS sur deux disques NVMe en miroir
zpool create -o ashift=12 \
  -O compression=lz4 \
  -O atime=off \
  -O xattr=sa \
  -O recordsize=128k \
  tank mirror /dev/nvme1n1 /dev/nvme2n1
# Verification de l'etat du pool
zpool status tank
# Declaration du pool dans Proxmox
pvesm add zfspool tank-zfs --pool tank --content images,rootdir
# Ajout d'un partage NFS depuis un NAS
pvesm add nfs nas-iso \
  --server 192.168.10.50 \
  --export /volume1/proxmox-iso \
  --content iso,vztmpl \
  --options vers=4.1
# Verification
pvesm status
```
Pour les VM Windows et SQL Server, activez l’option `recordsize=64k` sur le dataset ZFS dédié : la valeur par défaut de 128 Ko amplifie l’écriture (write amplification) sur des bases de données qui écrivent par blocs de 8 Ko. Pour les VM Linux avec ext4 et LVM, gardez 128 Ko. Pour les caches PostgreSQL ou MySQL, créez un dataset à 16 Ko aligné sur la taille de page du moteur.

## Étape 9 : Construire un cluster à trois nœuds

Un cluster Proxmox autorise jusqu’à 32 nœuds reliés par Corosync (bus de cluster) et synchronisés via la base pmxcfs (un FUSE filesystem distribué). À partir de trois nœuds, vous bénéficiez du quorum (majorité), de la migration à chaud, de la haute disponibilité (HA Manager) et du stockage répliqué Ceph. Avec deux nœuds, prévoyez un QDevice externe (Raspberry Pi, conteneur tiers) pour éviter le split-brain.

```
# Sur le premier noeud (pve01) : creation du cluster
pvecm create lab-cluster --link0 192.168.20.11
# Verification
pvecm status
# Sur les deuxieme (pve02) et troisieme (pve03) noeuds :
# rejoindre le cluster via l'IP du premier
pvecm add 192.168.20.11 \
  --link0 192.168.20.12 \
  --use_ssh
# Saisir le mot de passe root du premier noeud quand demande
# Verification finale depuis n'importe quel noeud
pvecm nodes
pvecm status
# Sortie attendue :
# Membership information
# ----------------------
# Nodeid      Votes Name
#      1          1 pve01 (local)
#      2          1 pve02
#      3          1 pve03
```
Une fois le cluster formé, l’interface web d’un nœud affiche l’arborescence complète avec les trois hôtes. Tout nœud peut piloter l’ensemble du cluster : les fichiers de configuration sont synchronisés en temps réel via pmxcfs sous `/etc/pve`. Réservez idéalement un sous-réseau dédié pour Corosync (paramètre `--link0`) avec une latence inférieure à 5 ms. Pour la résilience, ajoutez un `--link1` sur un second VLAN.

