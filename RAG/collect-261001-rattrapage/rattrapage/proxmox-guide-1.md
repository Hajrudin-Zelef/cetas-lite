---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-1
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "datacenter", "gpu", "intel", "memory", "open source"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [1, 236]
sha256: 06edf401787c20b949b7351f1b8d6a2ebde210cc4268be8a09fb1f7452239dfe
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

> Tout sur Proxmox VE, l'hyperviseur open source : installation,
> ZFS, Ceph, cluster, HA, sauvegardes (vzdump + Proxmox Backup
> Server), réseau, sécurité, GPU passthrough, cloud-init,
> Terraform, API, NUT (onduleur), dépannage — **angle chef de
> service systèmes**.
>
> Basé sur Proxmox VE 8.x/9.x (Debian 12/13). Les menus exacts
> bougent légèrement entre versions : **validez sur votre version**
> (`pveversion -v`).

---

## 1. Proxmox VE — pourquoi c'est le bon choix

- **Hyperviseur bare-metal** basé sur Debian : KVM (VM) + LXC
  (conteneurs) dans une seule interface web.
- **Gratuit et open source** (AGPLv3) ; abonnement Enterprise
  optionnel (support + dépôts stables).
- **Tout intégré** : clustering, HA, Ceph, pare-feu, sauvegardes,
  SDN — sans licence par socket comme VMware.
- Parfait pour PME/ETI africaines : un seul outil, matériel
  standard, pas de coûts cachés.

---

## 2. Architecture — comprendre avant d'installer

```
┌─────────────────────────────────────────────┐
│              Proxmox VE (Debian)            │
│  ┌──────────┐  ┌──────────┐  ┌────────────┐  │
│  │ KVM (VM) │  │ LXC      │  │ Ceph / ZFS │  │
│  │ Windows, │  │ Linux    │  │ Stockage   │  │
│  │ Linux    │  │ léger    │  │ distribué  │  │
│  └──────────┘  └──────────┘  └────────────┘  │
│  Cluster (corosync) │ HA │ Firewall │ Backup│
└─────────────────────────────────────────────┘
```
- **KVM** : virtualisation complète (tout OS, dont Windows).
- **LXC** : conteneurs système Linux (10× plus légers, démarrage
  en secondes, snapshots instantanés).
- **Règle** : un service = une VM ou un conteneur (jamais de
  services en dur sur l'hôte).

---

## 3. Prérequis matériels

| Usage | CPU | RAM | Disques | Réseau |
|---|---|---|---|---|
| Lab / test | 4 cœurs | 16 Go | 1 SSD | 1 GbE |
| PME (10–20 VM) | 8–16 cœurs | 64 Go | 2 SSD (ZFS mirror) + HDD | 2× 1 GbE |
| Production | 16+ cœurs | 128 Go+ | SSD NVMe + Ceph | 2× 10 GbE |
| Cluster Ceph | 3 nœuds min | 64 Go+/nœud | SSD par OSD | 10 GbE dédié Ceph |

- **CPU** : Intel VT-x/AMD-V + **EPT/NPT** (vérifier dans le BIOS).
- **RAM** : la ressource n°1 — toujours prévoir 20 % de marge.
- **Disques** : SSD obligatoire pour l'OS et les VM (un HDD =
  des VM au ralenti).
- **Éviter** : cartes RAID hardware avec cache sans batterie
  (ZFS veut du JBOD/HBA).

---

## 4. Installation

1. Télécharger l'ISO (proxmox.com), vérifier le SHA256.
2. Boot → Install Proxmox VE → disque cible.
3. **Options avancées (F2)** : choisir **ZFS (RAID1)** si 2 disques.
4. Pays, mot de passe root, **IP fixe + FQDN** (ex. `pve01.local`).
5. Reboot → interface web : `https://IP:8006` (root + mot de passe).

### 4.1 Installation sur Debian existante
```bash
# Possible (dépôt Proxmox sur Debian 12/13), mais l'ISO est plus propre
```

---

## 5. Premier boot — checklist

```
□ Web UI accessible (https://IP:8006), certificat (ACME §33)
□ Dépôts : no-subscription (lab) ou enterprise (prod) — §34
□ apt update && apt full-upgrade && reboot si kernel
□ Stockage : vérifier (local, local-lvm / ZFS)
□ Réseau : vmbr0 sur la bonne interface, IP, DNS, NTP
□ Alertes email configurées (§37)
□ Sauvegarde : PBS ou vzdump planifié (§18-19)
□ Utilisateurs : pas de root partagé (§29), 2FA (§30)
□ Pare-feu datacenter activé (§32)
□ NUT (onduleur) si applicable (§68)
□ Fiche d'exploitation (§76)
```

---

## 6. L'interface web — tour du propriétaire

- **Datacenter** : vue globale, cluster, stockage, pare-feu, ACME.
- **Nœud** : résumé, shell, mises à jour, disques, syslog.
- **VM/Conteneur** : console (noVNC/SPICE), matériel, options,
  snapshots, backup, pare-feu.
- **Raccourcis** : double-clic = console ; clic droit = actions.
- **Shell** : chaque nœud a un terminal web (root) — pratique,
  mais préférez SSH + clés.

---

## 7. Stockage — les concepts

| Type | Usage | Partagé ? |
|---|---|---|
| **local** (dir) | ISO, templates, backups vzdump | Non |
| **local-lvm** / **local-zfs** | Disques VM/LXC | Non |
| **NFS** | ISO, backups, (VM si rapide) | Oui |
| **Ceph/RBD** | Disques VM en cluster | Oui |
| **CephFS** | Fichiers partagés | Oui |
| **iSCSI** | LUN SAN | Oui |
| **PBS** | Sauvegardes dédupliquées | Oui |

- **Règle d'or** : HA et migration à chaud exigent un stockage
  **partagé** (Ceph, NFS…).

---

## 8. ZFS en détail — le filesystem qui pardonne

### 8.1 Créer un pool
```bash
zpool create -f tank mirror /dev/sdb /dev/sdc
# ou RAIDZ : zpool create -f tank raidz /dev/sdb /dev/sdc /dev/sdd
zpool status tank
```

### 8.2 Les réglages qui comptent
```bash
zfs set compression=lz4 tank        # gratuit, souvent +30 % d'espace
zfs set atime=off tank              # moins d'écritures
zfs set recordsize=16K tank/vm-disks # adapté aux VM
```

### 8.3 Snapshots et clones
```bash
zfs snapshot tank/vm-100-disk-0@sauve-avant-maj
zfs rollback tank/vm-100-disk-0@sauve-avant-maj
zfs list -t snapshot
```

### 8.4 Surveillance
```bash
zpool status -x          # santé (doit dire "all pools are healthy")
zpool iostat -v 5
zfs list
# Scrub mensuel (intégré à Proxmox) : détecte la corruption silencieuse
```

### 8.5 Règles ZFS
- **ashift=12** pour les disques 4K (défaut sur les ISO récentes).
- **Jamais plus de 80 %** de remplissage (les perfs s'effondrent).
- **RAM** : ZFS adore la RAM (ARC) — 1 Go de RAM par To de
  stockage en règle grossière.
- **SLOG** (SSD) : pour le sync intensif (bases de données).
- **L2ARC** : cache lecture SSD (optionnel).

---

## 9. LVM-thin — l'alternative simple

- `local-lvm` : thin provisioning (les VM ne consomment que ce
  qu'elles utilisent).
- Snapshots possibles (moins souples que ZFS).
- Bien pour un nœud unique sans besoin ZFS.

---

## 10. Réseau — bridges, bonds, VLAN

### 10.1 Le bridge vmbr0 (base)
```
Internet/LAN ──► eth0 ──► vmbr0 ──► VM (tap)
```
- Créé à l'install. Les VM branchées sur vmbr0 sont sur le LAN.

### 10.2 Bond (2 cartes)
- Nœud → Réseau → Créer → Linux Bond (mode 802.3ad/LACP ou
  active-backup) → vmbr0 sur le bond.
- **LACP** : configurer aussi le switch !

### 10.3 VLAN
- Bridge « VLAN aware » : cocher sur vmbr0 → tag par VM
  (Matériel → Réseau → Tag VLAN).
- Séparer : management, VM, Ceph, backup (4 VLAN = propre).

### 10.4 En CLI (`/etc/network/interfaces`)
```
auto vmbr0
iface vmbr0 inet static
    address 192.168.1.11/24
    gateway 192.168.1.1
    bridge-ports eth0
    bridge-stp off
    bridge-fd 0
```
- ⚠️ Toujours garder un accès (IPMI) quand on touche au réseau.

---

## 11. SDN (Software Defined Network)

- Datacenter → SDN : zones (VLAN, QinQ, VXLAN, EVPN), VNets,
  sous-réseaux, DHCP, IPAM.
- **VXLAN** : réseau overlay multi-sites sans VLAN trunk.
- Pour les infras qui grandissent : l'IPAM intégré évite le
  tableur d'IP.

---

## 12. Créer une VM KVM

### 12.1 Via l'UI
Créer VM → Nom, OS (ISO), Système (BIOS **OVMF/UEFI** pour les OS
récents), Disque (**VirtIO SCSI**, discard), CPU (cores, type
**host**), Mémoire (**ballooning**), Réseau (**VirtIO**).

### 12.2 Via CLI
```bash
qm create 100 --name srv-web --memory 4096 --cores 2 \
  --net0 virtio,bridge=vmbr0 --scsihw virtio-scsi-pci
qm importdisk 100 /var/lib/vz/template/iso/ubuntu.iso local-lvm
qm set 100 --scsi0 local-lvm:vm-100-disk-0 --boot order=scsi0
qm start 100
```

