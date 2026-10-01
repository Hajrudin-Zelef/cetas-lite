---
id: collect-261001-general-networking/general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026-3
title: "Télécharger l'ISO et la somme SHA-256"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "distribution", "memory"]
source: docs/RAG/collect-261001-general-networking/proxmox-9-1-cluster-ha-99-9-uptime-en-13-etapes-2026.md
source_anchor: ""
source_lines: [130, 261]
sha256: 91b1ebb00b756682c9440b1b5f5f04d3c89881fa5eefce252ed85d0b35550da4
---

# Télécharger l'ISO et la somme SHA-256

```
# Créer une VM Debian 13 avec OVMF UEFI
qm create 100 \
  --name debian13-web \
  --memory 4096 \
  --cores 2 \
  --cpu host \
  --net0 virtio,bridge=vmbr0 \
  --bios ovmf \
  --machine q35 \
  --efidisk0 tank-vms:1,format=raw,efitype=4m,pre-enrolled-keys=1 \
  --scsihw virtio-scsi-single \
  --scsi0 tank-vms:32,iothread=1,ssd=1,discard=on \
  --ide2 local:iso/debian-13.0.0-amd64-netinst.iso,media=cdrom \
  --boot order='scsi0;ide2;net0' \
  --agent enabled=1,fstrim_cloned_disks=1 \
  --ostype l26
# Démarrer la VM
qm start 100
# Ouvrir la console noVNC depuis l'interface web ou en CLI
qm terminal 100
```
La console s’ouvre dans un onglet noVNC et présente l’écran d’installation Debian. Suivez la procédure habituelle, n’oubliez pas d’activer « SSH server » et « standard system utilities » dans tasksel. Après installation, exécutez apt install qemu-guest-agent dans la VM : cet agent permet à Proxmox de communiquer avec l’invité (graceful shutdown, fstrim coordonné, gel pour snapshots cohérents). Vérifiez ensuite depuis l’hôte avec qm guest cmd 100 ping qui doit répondre « OK ».

## Étape 6 : Créer un container LXC léger

Les containers LXC sont la véritable arme secrète de Proxmox. Pour une application Linux, ils consomment dix fois moins de RAM qu’une VM équivalente, démarrent en moins d’une seconde et profitent directement du système de fichiers ZFS pour des snapshots instantanés. Idéal pour héberger Nginx, MariaDB, Redis, ou même une instance n8n autohébergée.

Téléchargez d’abord un template via la commande pveam. Le dépôt officiel propose des images turnkey-linux pré-configurées (LAMP, Nextcloud, Mediawiki, etc.) et des modèles « standard » de chaque distribution. Pour cet exemple, on utilise Debian 13 standard, environ 90 Mio.

```
# Mettre à jour la liste des templates LXC
pveam update
# Rechercher Debian 13
pveam available | grep debian-13
# Télécharger le template
pveam download local debian-13-standard_13.0-1_amd64.tar.zst
# Créer le container avec ID 200
pct create 200 \
  local:vztmpl/debian-13-standard_13.0-1_amd64.tar.zst \
  --hostname nginx-front \
  --password "$(openssl rand -base64 18)" \
  --memory 1024 \
  --cores 1 \
  --rootfs tank-vms:8 \
  --net0 name=eth0,bridge=vmbr0,ip=192.168.1.50/24,gw=192.168.1.1 \
  --features nesting=1,keyctl=1 \
  --unprivileged 1 \
  --onboot 1
# Démarrer et entrer dans le container
pct start 200
pct enter 200
```
Une fois à l’intérieur du container, vous obtenez un shell root Debian 13 vierge. Installez Nginx en deux commandes, c’est suffisant pour servir un site statique en HTTPS via Let’s Encrypt. La nesting=1 active la possibilité d’exécuter Docker à l’intérieur du LXC – pratique mais à utiliser avec parcimonie car cela multiplie les couches de sécurité. Le mode « unprivileged 1 » est obligatoire en production : il mappe les UID root du container vers une plage non-privilégiée de l’hôte, ce qui empêche l’évasion en cas de faille kernel.

## Étape 7 : Configurer le réseau et le pont vmbr0

Le réseau dans Proxmox repose sur des bridges Linux standards déclarés dans /etc/network/interfaces. Le bridge par défaut s’appelle vmbr0 et il est attaché à votre première carte physique. C’est par lui que toutes les VM et tous les containers sortent vers le LAN. En production, on ajoute un second bridge vmbr1 dédié au stockage iSCSI ou NFS, parfois un vmbr2 sur un VLAN d’administration séparé.

```
# Exemple complet /etc/network/interfaces pour un nœud avec 2 NIC
auto lo
iface lo inet loopback
# Interface physique 1 – LAN principal
auto enp3s0
iface enp3s0 inet manual
# Interface physique 2 – stockage
auto enp4s0
iface enp4s0 inet manual
# Bridge LAN (VM et containers)
auto vmbr0
iface vmbr0 inet static
    address 192.168.1.10/24
    gateway 192.168.1.1
    bridge-ports enp3s0
    bridge-stp off
    bridge-fd 0
    bridge-vlan-aware yes
    bridge-vids 2-4094
# Bridge stockage 10 GbE (Ceph, NFS)
auto vmbr1
iface vmbr1 inet static
    address 10.0.0.10/24
    bridge-ports enp4s0
    bridge-stp off
    bridge-fd 0
    mtu 9000
```
Appliquez la configuration sans redémarrer avec ifreload -a (le service ifupdown2 est installé par défaut sur Proxmox 9.1). Le mtu 9000 active les jumbo frames, indispensable sur le réseau Ceph car cela divise par six le nombre de paquets pour un transfert OSD. Pensez à activer le même MTU sur votre switch et sur tous les nœuds. Une asymétrie de MTU entre deux interfaces du même VLAN provoque des fragmentations silencieuses et des chutes de performance jusqu’à 70 %.

## Étape 8 : Bâtir un cluster Proxmox à trois nœuds

Un cluster Proxmox apporte la gestion centralisée, la migration à chaud et la haute disponibilité. Le quorum Corosync exige un nombre impair de nœuds – trois est le minimum pratique. Si vous n’avez que deux machines physiques, ajoutez un QDevice : il s’agit d’un petit conteneur sur un Raspberry Pi qui sert d’arbitre, sans héberger de VM, et coûte 35 euros au lieu d’un troisième serveur. Pour ce tutoriel, on suppose trois nœuds nommés pve1, pve2 et pve3, sur le réseau 192.168.1.10, 11 et 12.

```
# Sur pve1 – créer le cluster
pvecm create homelab-cluster --link0 192.168.1.10
# Sur pve2 – rejoindre le cluster
pvecm add 192.168.1.10 --link0 192.168.1.11
# Sur pve3 – rejoindre le cluster
pvecm add 192.168.1.10 --link0 192.168.1.12
# Vérifier le quorum
pvecm status
# Sortie attendue
# Quorum information
# Date:             Fri Apr 17 14:32:11 2026
# Quorum provider:  corosync_votequorum
# Nodes:            3
# Node ID:          0x00000001
# Ring ID:          1.4f
# Quorate:          Yes
```
Le mot de passe demandé est celui du compte root de pve1. Une fois les trois nœuds joints, l’interface web de n’importe quel d’entre eux affiche les trois sous Datacenter, et vous pouvez piloter une VM depuis l’interface de pve3 même si elle tourne sur pve1. La configuration partagée vit dans /etc/pve, un système de fichiers Corosync qui se réplique entre tous les nœuds. Ne touchez jamais ce dossier en édition directe sur un nœud isolé : vous risquez un split-brain. Préférez la console web ou les commandes officielles qui passent par l’API.

## Étape 9 : Activer la haute disponibilité et la migration à chaud

La haute disponibilité (HA) dans Proxmox s’appuie sur le service pve-ha-lrm et un groupe HA défini au niveau du datacenter. Quand un nœud tombe, les VM marquées « HA » redémarrent automatiquement sur un nœud survivant en moins de soixante secondes. Pour que cela fonctionne, les disques virtuels doivent vivre sur un stockage partagé : Ceph, NFS, ou ZFS avec réplication asynchrone (option moins idéale car non-temps-réel).

```
# Créer un groupe HA avec priorité décroissante
ha-manager groupadd webgroup --nodes "pve1:3,pve2:2,pve3:1"
# Ajouter la VM 100 à la HA
ha-manager add vm:100 --group webgroup --max_restart 3 --max_relocate 2
# Vérifier l'état HA
ha-manager status
# Migrer manuellement à chaud la VM 100 vers pve2
qm migrate 100 pve2 --online
```
La migration à chaud transfère la mémoire vive de la VM par incréments tout en la maintenant en activité. Selon nos tests sur 4 Gio de RAM et un lien 10 GbE, la coupure réseau effective dure entre 280 et 450 millisecondes – assez court pour qu’une connexion TCP ouverte ne soit pas réinitialisée. Si vous migrez à travers un lien 1 GbE, comptez 8 à 12 secondes de coupure visible. La commande qm migrate accepte aussi l’option –with-local-disks pour transférer les disques vers un stockage local distant, utile si vous mélangez ZFS local et Ceph.

## Étape 10 : Mettre en place Proxmox Backup Server 4.2

