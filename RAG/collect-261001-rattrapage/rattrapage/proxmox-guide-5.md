---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-5
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "benchmark", "memory"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [994, 1234]
sha256: be81027e2133f63e92873b00eae92360250592b1162a7021df3b2f1768605ba2
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

### 55.2 Cache tiering / classes
- Classes de périphériques : `ceph osd crush set-device-class`
  (ssd/hdd) → pools « chauds » (SSD) et « froids » (HDD).

### 55.3 Réparation
```bash
ceph pg dump | grep -v active+clean   # PG non sains
ceph osd pool set vm-pool pg_autoscale_mode on
rados bench -p vm-pool 30 write       # test de perf (pool de test !)
```

---

## 56. Réseau avancé

### 56.1 Open vSwitch (alternative au bridge Linux)
- Fonctionnalités : sFlow, OpenFlow, bonds avancés.
- Nœud → Réseau → Créer → OVS Bridge/IntPort.

### 56.2 LACP — checklist
```
□ Switch : port-channel en mode LACP actif
□ Proxmox : bond 802.3ad, miimon 100
□ Tester : débrancher un câble → pas de coupure
□ Débit : iperf3 entre nœuds (proche du théorique ?)
```

### 56.3 MTU 9000 (jumbo)
- Sur le réseau Ceph/backup : +10–20 % de débit.
- **Partout ou nulle part** (un maillon à 1500 = fragmentation).

---

## 57. SDN avancé — multi-site

- **EVPN** : étendre un VXLAN sur plusieurs sites (routé).
- **IPAM** : plus de tableur d'IP, le DHCP SDN attribue.
- Cas : 2 sites reliés (VPN/fibre) avec les mêmes VNets.

---

## 58. Sécurité avancée

- **CIS Benchmark** : durcir l'hôte Debian (mot de passe GRUB,
  `/boot` en lecture seule, auditd).
- **Ports** : 8006 (UI) jamais exposé brut → VPN/bastion.
- **Clés API** : rotation annuelle, périmètre minimal.
- **Mises à jour** : fenêtre mensuelle, jamais en retard de 2 versions.
- **Sauvegarde des configs** : `/etc/pve` (versionné en git !),
  `/etc/network/interfaces`, `/etc/ceph`.

```bash
cd /etc/pve && git init && git add -A && git commit -m "init"
# + push vers un dépôt privé après chaque changement majeur
```

---

## 59. API — exemples curl

```bash
# Token : PVEAPIToken=user@pam!token=xxxx
curl -sk -H "Authorization: PVEAPIToken=root@pam!mon-token=xxx" \
  https://pve01:8006/api2/json/nodes
curl -sk -H "Authorization: ..." \
  https://pve01:8006/api2/json/nodes/pve01/qemu | jq '.data[].name'
# Démarrer une VM :
curl -sk -X POST -H "Authorization: ..." \
  https://pve01:8006/api2/json/nodes/pve01/qemu/100/status/start
```
---

## 60. Terraform — aller plus loin

```hcl
# main.tf — parc complet
resource "proxmox_vm_qemu" "srv" {
  count       = 3
  name        = "srv-app-0${count.index + 1}"
  target_node = "pve01"
  clone       = "tpl-ubuntu-2404"
  full_clone  = true
  cores       = 2
  memory      = 4096
  scsihw      = "virtio-scsi-pci"
  disk {
    size    = "30G"
    storage = "ceph-rbd"
  }
  network {
    model  = "virtio"
    bridge = "vmbr0"
    tag    = 10
  }
  ipconfig0 = "ip=192.168.10.2${count.index + 1}/24,gw=192.168.10.1"
  sshkeys   = file("~/.ssh/id_rsa.pub")
}
```
- `terraform plan/apply/destroy` : le parc devient du code.
- **State** : backend distant (pas en local seul).

---

## 61. Ansible + Proxmox

```yaml
# Après Terraform : configurer les VM
- hosts: proxmox_vms
  become: true
  roles:
    - base        # ssh, fail2ban, zabbix-agent, unattended-upgrades
    - web         # nginx + TLS
    - monitoring
```
- Inventaire dynamique : plugin `community.general.proxmox`
  (les VM apparaissent automatiquement).

---

## 62. Monitoring avancé — Grafana

- **pve-exporter** → Prometheus → Grafana : dashboard « Proxmox
  Cluster » (CPU/RAM/disque par nœud/VM, Ceph, backups).
- Alertes : VM down, Ceph HEALTH_WARN, backup failed, disque > 85 %.
- Le dashboard affiché au bureau = la vitrine du service.

---

## 63. ISO et templates — gestion

- Stockage `local` : ISO (trier par OS), templates LXC
  (`pveam update && pveam download local ubuntu-24.04-standard`).
- **Nettoyer** : les ISO de 2019 ne servent plus.
- Templates versionnés : `tpl-ubuntu-2404-v3` (date dans la note).

---

## 64. Windows Server — bonnes pratiques

```
□ VirtIO (disque/réseau) + guest tools + QEMU agent
□ Désactiver la mise en veille, pare-feu Windows adapté
□ Activation : KMS ou clé par VM (licences !)
□ Antivirus : exclusions des dossiers applicatifs
□ Windows Update : fenêtre (pas pendant les backups)
□ Ballooning : OK (driver virtio-balloon)
```

---

## 65. Docker sur Proxmox — où le mettre ?

- **Dans une VM** : le plus propre (isolation, snapshots).
- **Dans un LXC** : possible (privilégié + nesting), plus léger,
  mais moins isolé.
- **Jamais sur l'hôte** : l'hôte reste un hyperviseur pur.

---

## 66. Kubernetes sur Proxmox

- 3 VM (control-plane) + N workers, Terraform + Ansible
  (kubespray) ou Talos.
- Stockage : CSI Ceph RBD (le combo gagnant).
- Commencer petit : k3s dans 3 VM pour les premiers clusters.

---

## 67. Multi-sites — les options

1. **Cluster étendu** : corosync exige < 5 ms de latence → fibre
   noire/métro uniquement.
2. **Clusters séparés + PBS sync** : le standard (RPO = fréquence
   de sync).
3. **SDN EVPN** : mêmes réseaux logiques sur 2 sites.

---

## 68. Migration depuis VMware / Hyper-V

```bash
# Méthode : exporter en OVF/OVA → importer
qm importovf 100 export.ova local-lvm
# Puis : installer les pilotes VirtIO, QEMU agent, nettoyer les
# anciens tools VMware, réactiver Windows (changement de matériel)
```
- Tester **une VM pilote** avant de tout migrer.
- Licences Windows : vérifier la transférabilité.

---

## 69. Erreurs classiques des débutants

1. **Installer sur 1 disque sans RAID** → le disque meurt, tout meurt.
2. **ZFS rempli à 95 %** → perfs effondrées, snapshots impossibles.
3. **Backups jamais testés** → le grand classique.
4. **Cluster à 2 nœuds sans qdevice** → split-brain à la première panne.
5. **Exposer le 8006 sur Internet** → attaqué en quelques heures.
6. **Root partagé** → on ne sait plus qui a fait quoi.
7. **Pas de NUT** → coupure = corruption.
8. **Surallocation CPU 10:1** → tout rame, personne ne comprend.
9. **Mélanger les dépôts** (enterprise + test) → système instable.
10. **Ne pas documenter** → le savoir part avec le départ.

---

## 70. Pense-bête de poche

```
pvecm status | ceph -s | zpool status -x | pvesm status
qm list | pct list | ha-manager status
vzdump --help | pbs : verify hebdo, GC hebdo, sync distant
Backups testés 1×/mois | NUT testé 1×/trimestre
Dépôts : enterprise (prod) | Mises à jour : 1 nœud à la fois
Quorum impair | Ceph < 85 % | ZFS < 80 %
8006 : jamais sur Internet brut | 2FA admins
```

---

## 71. Questions pièges

1. Pourquoi le cluster se fige sans quorum ? → Sécurité : éviter
   le split-brain (2 cerveaux qui décident différemment).
2. Snapshot = sauvegarde ? → Non (même stockage, pas de 3-2-1).
3. HA sans stockage partagé ? → Impossible (la VM doit être
   accessible depuis l'autre nœud).
4. Pourquoi tester NUT en vrai ? → Un scénario non testé ne marche pas.
5. ZFS : pourquoi < 80 % ? → Le copy-on-write a besoin d'espace
   libre pour rester performant.
6. Linked clone en prod ? → Non (dépend du template).
7. Pourquoi `host` comme type CPU ? → Perfs max (flags CPU natifs).
8. Fencing : pourquoi obligatoire en HA ? → Éviter 2 instances de
   la même VM (corruption disque).
9. PBS vs vzdump ? → Dedup, incrémental, chiffrement, file restore.
10. Ransomware et backups connectés ? → Chiffrés aussi → disque
    déconnecté mensuel.
---

## 72. Glossaire

