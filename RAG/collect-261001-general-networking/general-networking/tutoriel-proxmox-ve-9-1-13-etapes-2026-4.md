---
id: collect-261001-general-networking/general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026-4
title: "Téléchargement de l'ISO Proxmox VE 9.1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-proxmox-ve-9-1-13-etapes-2026.md
source_anchor: ""
source_lines: [230, 365]
sha256: c66d376b2ebc2d965d9b64b9ac4b8101b87f98c42f1be8816577239a180a660e
---

# Téléchargement de l'ISO Proxmox VE 9.1

## Étape 10 : Stockage hyper-convergé avec Ceph Squid

Ceph Squid 19.2.3 est intégré nativement dans Proxmox VE 9.1. Sur un cluster à trois nœuds avec au moins un disque dédié par nœud (idéalement un NVMe enterprise PLP), vous pouvez bâtir un pool RBD répliqué qui survit à la perte d’un hôte sans interruption. L’installation se fait via l’interface web : *Datacenter > Ceph > Install Ceph*, sélection de la version Squid (19.2), repository No-Subscription. La procédure se répète sur chaque nœud.

```
# Installation Ceph en CLI sur chaque noeud
pveceph install --repository no-subscription --version squid
# Initialisation du cluster Ceph (sur pve01 uniquement)
pveceph init --network 192.168.30.0/24
# Creation des moniteurs (sur les 3 noeuds)
pveceph mon create
# Creation des managers
pveceph mgr create
# Ajout des OSD (un disque NVMe dedie par noeud)
ceph-volume lvm zap /dev/nvme3n1 --destroy
pveceph osd create /dev/nvme3n1
# Creation d'un pool RBD replique (taille=3, min_size=2)
pveceph pool create vm-pool --size 3 --min_size 2 --pg_num 128
# Ajout dans Proxmox comme stockage RBD
pvesm add rbd vm-ceph --pool vm-pool --content images,rootdir --krbd 0
```
Avec un pool en taille 3, chaque écriture est répliquée trois fois sur trois OSD différents : la perte d’un hôte ou d’un disque ne provoque aucune perte de données. La latence est plus élevée qu’un ZFS local (3 à 5 ms en lecture, 10 à 30 ms en écriture sur 10 GbE), mais la disponibilité justifie le compromis pour les charges critiques. Comptez 25 GbE minimum pour 100 OSD ou plus.

## Étape 11 : Sauvegardes incrémentielles avec Proxmox Backup Server 4.1

Proxmox Backup Server (PBS) 4.1, publié le 26 novembre 2025, est l’outil dédié de sauvegarde déduplicquée. PBS 4.1 partage la même base Debian 13 que PVE 9 et supporte la déduplication par blocs variables, le chiffrement client AES-256-GCM, la vérification automatique d’intégrité (verify jobs) et la synchronisation entre datastores. Installez PBS sur une machine séparée — physique ou virtuelle dans un autre cluster — pour respecter la règle 3-2-1 (3 copies, 2 supports, 1 hors site).

```
# Cote PBS : creation d'un datastore et d'un utilisateur
proxmox-backup-manager datastore create main /mnt/datastore/main
proxmox-backup-manager user create backup@pbs --password $(openssl rand -base64 16)
proxmox-backup-manager acl update /datastore/main DatastoreAdmin --auth-id backup@pbs
# Recuperation de l'empreinte du serveur
proxmox-backup-manager cert info | grep Fingerprint
# Cote PVE : ajout du datastore
pvesm add pbs pbs-main \
  --server 192.168.10.20 \
  --datastore main \
  --username backup@pbs \
  --password "MOTDEPASSEDEFINI" \
  --fingerprint "AB:CD:EF:..."
# Test de sauvegarde manuelle d'une VM
vzdump 100 --storage pbs-main --mode snapshot --compress zstd
# Creation d'un job de sauvegarde planifie (interface web ou CLI)
cat > /etc/pve/jobs.cfg.d/daily-backup.cfg <<EOF
vzdump: daily-backup
        schedule sat 02:00
        all 1
        storage pbs-main
        mode snapshot
        compress zstd
        notes-template "{{guestname}} backup"
        prune-backups keep-daily=7,keep-weekly=4,keep-monthly=6,keep-yearly=2
EOF
```
Le mode *snapshot* avec QEMU Guest Agent permet une sauvegarde à chaud sans interruption ni gel notable de la VM (généralement moins de 200 ms de fsfreeze). Le mode *stop* reste réservé aux invités sans agent. Activez systématiquement les *verify jobs* hebdomadaires : ils relisent l’intégralité des chunks et détectent les corruptions silencieuses du stockage avant qu’elles ne touchent une restauration.

## Étape 12 : Migrer une VM depuis VMware ESXi

Depuis Proxmox VE 8.2, l’interface intègre un **ESXi Import Wizard** qui permet de connecter directement un vCenter ou un hôte ESXi et d’importer ses VM sans passer par un export OVF. La fonctionnalité a été stabilisée et étendue dans la 9.1 : prise en charge des disques thin et thick, conversion automatique du contrôleur LSI vers VirtIO SCSI, conservation des UUID pour préserver les licences Windows. Selon les rapports de la communauté, l’opération sur une VM Windows Server 2022 de 200 Go prend 15 à 25 minutes sur un lien 10 GbE.

```
# Ajout d'un hote ESXi comme stockage source (interface : Datacenter > Storage > Add > ESXi)
# ou en CLI :
pvesm add esxi esxi-source \
  --server 192.168.40.50 \
  --username root \
  --password "MOTPASSE_ESXI" \
  --skip-cert-verification 1
# Listing des VM sur l'ESXi source
qm importovf --help
# Pour l'import via wizard : interface web > esxi-source > VM > Import
# Import en CLI (alternative manuelle)
qm importdisk 110 \
  /mnt/pve/esxi-source/ha-datacenter/srv-app01/srv-app01.vmdk \
  local-zfs --format qcow2
# Modification du controleur apres import
qm set 110 --scsihw virtio-scsi-single \
            --scsi0 local-zfs:vm-110-disk-0 \
            --boot order=scsi0 \
            --net0 virtio,bridge=vmbr0
qm start 110
```
Pour Windows, deux étapes manuelles dans la VM source avant la migration accélèrent le processus : installer les VirtIO drivers depuis le ISO `virtio-win.iso` (Fedora) et désactiver les VMware Tools. Sans drivers VirtIO, Windows démarrera en BSOD INACCESSIBLE_BOOT_DEVICE après import. Pour Linux, la plupart des distributions modernes embarquent les drivers VirtIO dans leur initramfs et démarrent sans manipulation.

## Étape 13 : Activer la haute disponibilité et automatiser via Ansible

Le HA Manager de Proxmox redémarre automatiquement les VM critiques sur un autre nœud si l’hôte d’origine plante. La configuration se fait via *Datacenter > HA* ou via la commande `ha-manager`. Précondition : les VM doivent résider sur un stockage partagé (Ceph, NFS ou iSCSI), faute de quoi le HA Manager ne pourrait pas démarrer la VM ailleurs.

```
# Creation d'un groupe HA
ha-manager groupadd lab-ha --nodes "pve01:2,pve02:1,pve03:1" --restricted
# Ajout d'une VM en HA avec demarrage automatique
ha-manager add vm:100 --group lab-ha --max_relocate 3 --state started
# Verification
ha-manager status
# Test : extinction brutale du noeud hebergeant la VM 100
# (on attend ~120 s pour le fencing puis le redemarrage ailleurs)
```
Pour automatiser la création de VM, le module `community.general.proxmox_kvm` d’Ansible reste la référence. Il interroge l’API REST sur le port 8006 et orchestre les opérations idempotentes. Voici un playbook minimal qui crée trois VM web en parallèle :

```
# inventory.ini
[proxmox]
pve01.lab.example.fr ansible_user=root
# playbook.yml
---
- name: Provision web VMs on Proxmox
  hosts: proxmox
  gather_facts: no
  vars:
    api_user: root@pam
    api_token_id: ansible
    api_token_secret: "{{ lookup('env', 'PVE_TOKEN') }}"
  tasks:
    - name: Create VMs
      community.general.proxmox_kvm:
        api_user: "{{ api_user }}"
        api_token_id: "{{ api_token_id }}"
        api_token_secret: "{{ api_token_secret }}"
        api_host: "{{ inventory_hostname }}"
        node: pve01
        clone: ubuntu-24-template
        newid: "{{ 200 + item }}"
        name: "web0{{ item }}"
        cores: 2
        memory: 4096
        net:
          net0: "virtio,bridge=vmbr0,firewall=1"
        state: present
      loop: [1, 2, 3]
```
Générez un token API depuis l’interface (*Datacenter > Permissions > API Tokens*), exportez le secret dans la variable d’environnement `PVE_TOKEN`, puis lancez `ansible-playbook -i inventory.ini playbook.yml`. L’idempotence garantit que relancer le playbook ne recrée pas les VM existantes : il vérifie l’ID avant d’agir.

## Comparaison Proxmox VE 9.1 vs VMware ESXi 8.0 vs Hyper-V 2025

