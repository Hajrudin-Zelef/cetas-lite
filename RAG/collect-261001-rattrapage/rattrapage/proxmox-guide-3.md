---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-3
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["agent", "arr", "datacenter", "gpu", "intel", "memory", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [491, 761]
sha256: c26c9669e443b89e9baa786e3ed9a9607852e1622a8dce9ce20ca36c3934a317
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

- **3 niveaux** : Datacenter → Nœud → VM/CT (le plus spécifique gagne).
- Groupes de sécurité : règles réutilisables (ex. « web » : 80/443).
```bash
pve-firewall status
```
- Activer au niveau datacenter (`Firewall: yes`), puis règles par VM.
- **Options** : anti-spoofing, log des refus (debug).

---

## 27. Sécurité — checklist Proxmox

```
□ Dépôts enterprise (prod) ou no-subscription (lab) — pas test en prod
□ Root : mot de passe fort, pas partagé, 2FA sur les admins
□ UI : certificat ACME/Let's Encrypt (§33), pas exposée sur Internet
       (ou bastion/VPN + 2FA + fail2ban)
□ Pare-feu datacenter + règles par VM
□ Mises à jour : hebdo (fenêtre), kernel + reboot planifié
□ Sauvegardes : PBS + test de restauration mensuel
□ Logs : syslog distant
□ NUT : arrêt propre sur coupure (§68)
□ Accès SSH : clés, pas de mot de passe
```

---

## 28. Mises à jour

```bash
apt update && apt full-upgrade
# Via UI : Nœud → Mises à jour → Rafraîchir → Mettre à niveau
```
- **Dépôts** :
  - `pve-enterprise` : stable, avec abonnement.
  - `pve-no-subscription` : gratuit, légèrement moins testé.
  - `pve-test` : jamais en prod.
- **Kernel** : reboot planifié après maj (fenêtre mensuelle).
- **Cluster** : un nœud à la fois (migrer les VM d'abord !).

---

## 29. Certificat TLS de l'UI (ACME)

- Datacenter → ACME → Ajouter compte → Commander le certificat.
- Fini le warning « certificat non valide » — et le HTTPS propre
  pour les accès distants.

---

## 30. Alertes email

- Datacenter → Options → Email d'alerte (SMTP du nœud).
- `/etc/postfix/main.cf` : relayhost vers votre SMTP.
- Tester : une alerte simulée (débrancher un câble réseau… en lab).

---

## 31. Monitoring — voir le parc

### 31.1 Intégré
- Résumé par nœud/VM : CPU, RAM, disque, réseau (graphes).
- `pvesh get /nodes/pve01/status`.

### 31.2 Zabbix / Prometheus
- Agent Zabbix dans les VM + template Proxmox (API).
- **pve-exporter** (Prometheus) : métriques de tout le cluster →
  Grafana (le tableau de bord du chef de service).

### 31.3 Ce qu'on supervise
- État nœuds, quorum, Ceph HEALTH, espace stockages.
- Sauvegardes : succès/échec (alerte immédiate).
- Températures, SMART des disques.

---

## 32. GPU passthrough (PCI)

```
1. BIOS : VT-d/IOMMU activé
2. Kernel : intel_iommu=on iommu=pt (GRUB, update-grub, reboot)
3. dmesg | grep -e IOMMU   (vérifier)
4. VM → Matériel → Ajouter → Périphérique PCI → le GPU
   (cocher "Primary GPU", ROM-Bar selon carte)
5. Pilote NVIDIA dans l'invité (+ masquer la virtualisation
   : args: -cpu host,+kvm_pv_unhalt pour le code 43)
```
- Usage : VDI, transcodage, IA/ML.
- ⚠️ Le GPU est **dédié** à la VM (plus dispo pour l'hôte).

---

## 33. USB passthrough

- VM → Matériel → Ajouter → Périphérique USB (par ID vendeur/produit).
- Usage : clé de licence, **onduleur USB pour NUT** (§68).

---

## 34. cloud-init en détail (Proxmox)

```bash
qm set 200 --ciuser admin --cipassword '***' \
  --sshkeys ~/.ssh/id_rsa.pub \
  --ipconfig0 ip=192.168.1.20/24,gw=192.168.1.1 \
  --nameserver 192.168.1.1 --searchdomain local
qm set 200 --cicustom "user=local:snippets/user.yaml"
```
- `snippets/` : user-data avancé (packages, runcmd, users…).
- Le template + cloud-init = **VM prête en 60 secondes**.

---

## 35. Terraform — l'infra as code

```hcl
resource "proxmox_vm_qemu" "web" {
  name        = "srv-web-01"
  target_node = "pve01"
  clone       = "template-ubuntu"
  cores       = 2
  memory      = 4096
  ipconfig0   = "ip=192.168.1.21/24,gw=192.168.1.1"
}
```
- Provider : Telmate/bpg. Token API : Datacenter → Permissions →
  API Tokens.
- **Idéal** : recréer tout le parc depuis du code (après sinistre).

---

## 36. API et CLI

```bash
pvesh get /nodes                    # API en CLI
pvesh create /nodes/pve01/qemu -vmid 100 -name test -memory 2048
qm list ; qm config 100 ; qm status 100
pct list ; pct config 300
pvesm status                        # stockages
```
- Scripts : tout ce que fait l'UI passe par l'API → automatisable
  (curl + token).

---

## 37. Hooks et planification

- **Hook scripts** vzdump : `--script` (notifier, monter/démonter…).
- **Jobs planifiés** : Datacenter → Sauvegarde (cron intégré).
- Exemple : snapshot avant maj auto :
```bash
# cron sur le nœud
0 1 * * 0 qm snapshot 100 avant-maj-hebdo && apt ... 
```

---

## 38. Disques — opérations

```bash
qm resize 100 scsi0 +20G              # agrandir (étendre ensuite dans l'OS)
qm move-disk 100 scsi0 ceph-pool      # changer de stockage (à chaud si partagé)
qm set 100 --delete scsi1             # supprimer (détacher d'abord !)
```
- Réduire un disque : **dangereux** (réduire dans l'OS d'abord,
  backup avant).
- `discard` : l'espace libéré dans la VM est rendu au stockage.

---

## 39. Stockages réseau — NFS, iSCSI, SMB

```bash
# NFS : Datacenter → Stockage → Ajouter → NFS (IP, export)
# iSCSI : portail, target → LUN → LVM sur iSCSI
# SMB/CIFS : pour ISO et backups (pas pour les disques VM)
```
- NFS : `vers=4`, jumbo frames si 10 GbE.
- iSCSI + multipath pour la redondance.

---

## 40. Performance — tuning KVM

```
CPU : type=host, pas de surallocation abusive (ratio 2-4:1 max)
RAM : ballooning ON, KSM ON (pages partagées entre VM identiques)
DISQUE : VirtIO SCSI, cache=none (Ceph/ZFS), iothread=1, discard=on
RÉSEAU : VirtIO, multiqueue (queues = vCPU) sur 10 GbE
```
- `ksmtuned` : automatique sur Proxmox.
- NUMA : activer si VM > 8 vCPU (topologie alignée sur l'hôte).

---

## 41. Invités Windows

1. ISO Windows + ISO **virtio-win** (fedorapeople.org).
2. Pendant l'install : charger le pilote (disque non détecté).
3. Après : installer `virtio-win-guest-tools` (agent + pilotes).
4. UEFI/OVMF + TPM pour Windows 11.
---

## 42. Maintenance d'un nœud — sans couper le service

```
1. Migrer les VM vers les autres nœuds (ou HA auto)
2. Vérifier : qm list (vide), ceph -s (si Ceph)
3. apt update && apt full-upgrade
4. Reboot si kernel
5. Vérifier le retour : pvecm status, services, Ceph HEALTH_OK
6. Re-migrer les VM (ou laisser l'équilibrage)
```
- **Jamais** : rebooter 2 nœuds Ceph en même temps.

---

## 43. Remplacement d'un nœud mort

1. `pvecm delnode pve03` (depuis un nœud sain, quorum OK).
2. Nouveau serveur : même hostname/IP, installer Proxmox.
3. `pvecm add` pour rejoindre.
4. Recréer les stockages locaux si besoin, restaurer les VM
   depuis PBS.
- **Leçon** : PBS + configs documentées = un nœud se remplace
  en 2 h.

---

## 44. SMART et ZED — les disques parlent

```bash
smartctl -a /dev/sda | grep -i "reallocated\|pending\|wear"
# SSD : Media_Wearout_Indicator / Percentage Used
```
- **ZED** (ZFS Event Daemon) : notifie les erreurs ZFS par mail.
- Remplacer un disque **avant** qu'il meure : SMART en alerte +
  erreurs ZFS = commande immédiate (délai Afrique !).

---

## 45. Logs — où chercher

```bash
journalctl -u pveproxy      # interface web
journalctl -u pvedaemon     # API
/var/log/syslog            # général
/var/log/kern.log
dmesg | grep -i "error\|fail"
# Cluster : /var/log/corosync/corosync.log
# Ceph : ceph -s, ceph health detail
```

---

## 46. Dépannage — 15 cas terrain

**Cas 1 — L'UI ne répond plus** : `systemctl status pveproxy`,
`pveproxy` restart, disque plein ? (`df -h`), certificat expiré ?

**Cas 2 — Quorum perdu** : `pvecm status` → si 1 nœud sur 3 :
réseau corosync ? switch ? Ne **jamais** forcer le quorum sans
comprendre (`pvecm expected 1` = mesure de dernier recours).

**Cas 3 — VM bloquée** : `qm stop 100` (dur), `qm unlock 100`
(verrou), vérifier le stockage.

**Cas 4 — Migration qui échoue** : stockage partagé ? même nom de
bridge ? assez de RAM sur la cible ? réseau 1 GbE trop lent
(prendre le temps).

