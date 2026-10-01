---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-2
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "memory"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [237, 490]
sha256: 8333e920ab2292560646c70d0010673d623f27332881c572fb65f1033e1d2539
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

### 12.3 Les bons réglages
- **Disque** : VirtIO SCSI + `discard` + `iothread`.
- **CPU** : type `host` (perfs), activer NUMA si > 8 vCPU.
- **Agent QEMU** : installer dans l'invité (`qemu-guest-agent`) →
  IP visible, shutdown propre, snapshots cohérents.
---

## 13. Templates VM + cloud-init — déployer en 2 minutes

```bash
# 1. Créer la VM, installer l'OS, qemu-guest-agent, cloud-init
# 2. Nettoyer : apt clean, cloud-init clean, vider les clés SSH
qm template 100   # convertit en template
# 3. Cloner :
qm clone 100 200 --name srv01 --full
qm set 200 --ciuser admin --sshkeys ~/.ssh/id_rsa.pub \
  --ipconfig0 ip=192.168.1.20/24,gw=192.168.1.1
qm start 200
```
- **Full clone** : copie indépendante. **Linked clone** : basé sur
  le template (rapide, mais dépendant).
- En production : full clone + cloud-init = standard.

---

## 14. Conteneurs LXC

### 14.1 Créer un CT
```bash
# Télécharger un template : CT Templates → ubuntu-24.04-standard
pct create 300 local:vztmpl/ubuntu-24.04-standard_24.04-1_amd64.tar.zst \
  --hostname ct-web --cores 2 --memory 2048 --rootfs local-lvm:8 \
  --net0 name=eth0,bridge=vmbr0,ip=192.168.1.30/24,gw=192.168.1.1
pct start 300 && pct enter 300
```

### 14.2 Privilégié vs non privilégié
- **Non privilégié** (défaut) : plus sûr (UID mappés).
- **Privilégié** : nécessaire pour Docker-dans-LXC, NFS…
  (cocher « nesting » pour Docker).

### 14.3 LXC vs VM — que choisir ?
| | LXC | VM |
|---|---|---|
| Démarrage | secondes | dizaines de secondes |
| Surcoût | ~nul | quelques % |
| OS | Linux uniquement | Tout (Windows…) |
| Snapshots | instantanés | rapides (qcow2/ZFS) |
| Isolation | correcte | forte |

---

## 15. Snapshots — VM et CT

```bash
qm snapshot 100 avant-maj --description "Avant kernel 6.8"
qm rollback 100 avant-maj
qm delsnapshot 100 avant-maj
pct snapshot 300 avant-maj
```
- Snapshots **avec RAM** : possible (état complet), plus lourds.
- ⚠️ Les snapshots ne sont pas des sauvegardes (même stockage !).
- Nettoyer les vieux snapshots (ils ralentissent et mangent
  l'espace).

---

## 16. Sauvegardes vzdump — la base

### 16.1 Configurer
Datacenter → Sauvegarde → Ajouter :
- Nœud, stockage (NFS/PBS), planification (`02:00`),
- Mode : **snapshot** (sans interruption), compression **zstd**,
- Rétention : garder 7 quotidiens, 4 hebdo.

### 16.2 CLI
```bash
vzdump 100 --mode snapshot --compress zstd --storage nas-backup
vzdump 100-110 --all 0   # tout le nœud
```

### 16.3 Restaurer
```bash
qmrestore /var/lib/vz/dump/vzdump-qemu-100.vma.zst 200
pct restore 300 /var/lib/vz/dump/vzdump-lxc-300.tar.zst
```
- Restaurer **sous un autre ID** pour tester sans écraser.

---

## 17. Proxmox Backup Server (PBS) — le niveau pro

### 17.1 Pourquoi PBS plutôt que vzdump ?
- **Déduplication** (blocs) : 10 VM Ubuntu = ~1,2× l'espace d'une.
- **Incrémentiel** réel, chiffrement côté client, vérification
  d'intégrité, **restauration de fichiers** (file restore).
- S'installe sur une machine dédiée (ou VM).

### 17.2 Mise en place
```bash
# Sur le PBS : Datastore → Ajouter (disque dédié)
# Sur PVE : Stockage → Ajouter → Proxmox Backup Server
#   (empreinte du certificat à vérifier !)
```
- Jobs : Datacenter → Sauvegarde → stockage PBS.
- **Rétention** : keep-daily=7, keep-weekly=4, keep-monthly=6.
- **Vérification** : job « Verify » hebdo (corruption détectée).
- **Sync** : répliquer vers un 2ᵉ PBS hors site (règle 3-2-1).
- **Garbage Collect** : planifié (nettoie les blocs orphelins).

### 17.3 Chiffrement
- Clé générée côté PVE, **sauvegardée hors site** (sans la clé,
  les backups sont perdus).

---

## 18. Stratégie de sauvegarde complète (chef de service)

```
NIVEAU 1 : PBS local (quotidien, rétention 7j/4s/6m)
NIVEAU 2 : Sync PBS → site distant (hebdo)
NIVEAU 3 : Dump mensuel → disque externe déconnecté (ransomware)
TEST : restauration 1 VM/mois (planning tournant)
```
- **Ransomware** : un backup connecté en permanence n'est pas un
  backup. Le disque déconnecté mensuel sauve des entreprises.

---

## 19. Cluster — assembler les nœuds

### 19.1 Créer
Nœud1 : Datacenter → Cluster → Créer (nom).
Nœud2/3 : Rejoindre (IP du nœud1 + mot de passe root).
```bash
pvecm status     # état
pvecm nodes      # membres
```

### 19.2 Règles
- **Noms d'hôtes uniques**, IP fixes, même version Proxmox.
- Réseau cluster **dédié** si possible (corosync est sensible à
  la latence).
- **Nombre impair** de nœuds (ou qdevice) pour le quorum.

### 19.3 Quorum
- Quorum = majorité des votes. Sans quorum : **le cluster se fige**
  (pas de démarrage/migration) — c'est une sécurité, pas un bug.
- 2 nœuds seuls = pas de quorum en cas de panne → ajouter un
  **qdevice** (un petit Debian/Raspberry comme arbitre).

---

## 20. Haute disponibilité (HA)

### 20.1 Activer
Datacenter → HA → Ajouter (VM/CT) → groupe, état « started ».
```bash
ha-manager status
```

### 20.2 Exigences
- Stockage **partagé** (Ceph/RBD, NFS…).
- **Fencing** configuré (IPMI/iDRAC) : sans fencing, pas de HA
  fiable (risque de split-brain).
- Tester : débrancher un nœud → les VM redémarrent ailleurs
  (quelques minutes).

### 20.3 Groupes et priorités
- Groupes : `groupe-prod` (nœuds 1,2,3), priorités par VM.
- **Répartition** : ne pas tout mettre en HA (les VM critiques
  d'abord).

---

## 21. Migration

```bash
qm migrate 100 pve02 --online   # à chaud (stockage partagé)
qm migrate 100 pve02            # à froid
```
- À chaud : quelques secondes de coupure (RAM à copier).
- Sans stockage partagé : migration avec disques (plus longue).
- **Avant maintenance d'un nœud** : migrer les VM (ou HA).

---

## 22. Ceph — le stockage distribué

### 22.1 Principes
- Chaque nœud apporte des disques (OSD) → un seul pool logique.
- Réplication ×3 : un nœud ou des disques peuvent mourir.
- **Exigences** : 3+ nœuds, **10 GbE dédié**, SSD de préférence.

### 22.2 Installation
PVE → Ceph → Installer → configurer le réseau public/cluster →
Créer les OSD (un par disque) → Créer les pools (rbd, cephfs).

### 22.3 Pools et dimensionnement
```bash
ceph osd pool create vm-pool 128 128
# PG : ~100 par OSD (calculateur : ceph osd pool autoscale)
ceph -s    # HEALTH_OK ?
```
- `size=3, min_size=2` : le standard (2 copies min pour écrire).

### 22.4 Maintenance Ceph
```bash
ceph -s ; ceph osd tree ; ceph df
ceph osd out osd.3        # sortir un OSD (remplacement disque)
ceph osd crush reweight ...
# Rebalance auto après ajout/retrait (peut durer des heures)
```
- **Ne jamais** : remplir à plus de 85 %, rebooter 2 nœuds à la
  fois, paniquer pendant un rebalance (c'est normal).

---

## 23. Réplication de stockage (sans Ceph)

- Datacenter → Réplication : réplique les disques ZFS locaux
  vers un autre nœud (toutes les 15 min, configurable).
- **Plan B** : migration rapide en cas de panne (quelques minutes
  de perte), sans Ceph.
- Idéal : 2 nœuds, ZFS local, petits budgets.

---

## 24. Utilisateurs, permissions, ACL

- Datacenter → Permissions : utilisateurs (pveum), groupes, rôles.
```bash
pveum user add admin2@pam
pveum acl modify /vms/100 --users admin2@pam --roles PVEVMUser
```
- **Rôles** : Administrator, PVEVMAdmin, PVEVMUser, PVEDatastoreUser…
- **Jamais de root partagé** : un compte par admin.
---

## 25. 2FA et LDAP/AD

### 25.1 TOTP (2FA)
Datacenter → Utilisateurs → [user] → TFA → Ajouter (QR code).
- Exiger pour les admins (surtout l'UI exposée).

### 25.2 LDAP / Active Directory
Datacenter → Accès → Ajouter → LDAP/AD (serveur, base DN, bind).
- Les comptes du domaine se connectent à Proxmox avec leurs
  identifiants habituels.

---

## 26. Pare-feu intégré (pve-firewall)

