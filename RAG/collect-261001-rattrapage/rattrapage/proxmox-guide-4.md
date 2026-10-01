---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-4
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [762, 993]
sha256: aab7391ad9c6d8606cda9baa7d362ce5870e31d639cbf81b7fec45c05aae753c
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

**Cas 5 — Ceph HEALTH_WARN** : `ceph health detail` → OSD down ?
placement groups en backfill (normal après ajout) ? disque plein ?

**Cas 6 — Backup qui échoue** : espace sur le stockage de backup ?
VM avec snapshot coincé ? PBS joignable (empreinte cert) ?

**Cas 7 — Plus de snapshots possibles (ZFS)** : pool > 80 % →
nettoyer, `zfs list -t snapshot` (les vieux !).

**Cas 8 — Conteneur LXC sans réseau** : bridge, VLAN tag, DHCP
dans le CT (`/etc/network/interfaces` du CT).

**Cas 9 — Windows sans disque à l'install** : pilote virtio non
chargé → ISO virtio-win + « Charger un pilote ».

**Cas 10 — HA qui ne bascule pas** : stockage partagé ? fencing
configuré ? `ha-manager status` (erreurs ?).

**Cas 11 — Corosync qui flappe** : latence réseau (ping entre
nœuds), MTU, switch saturé → réseau cluster dédié.

**Cas 12 — PBS : dedup qui n'avance plus** : Garbage Collect pas
planifié → le datastore grossit. Planifier GC hebdo.

**Cas 13 — Clé API qui ne marche plus** : token révoqué ? expiration ?
permissions du token (`pveum`).

**Cas 14 — Nœud qui ne rejoint pas le cluster** : heure désync
(chrony !), hostname résolu ?, version identique ?, pare-feu
(ports corosync 5404-5405/udp) ?

**Cas 15 — Après coupure électrique, les VM ne démarrent pas** :
ordre de démarrage (délai), stockage Ceph pas encore HEALTH_OK
(attendre), NUT mal configuré (§68).

---

## 47. Cas pratiques commentés

**Cas A — PME 15 VM, 2 nœuds**
2 serveurs (64 Go, ZFS mirror SSD), réplication toutes les 15 min,
PBS sur NAS, pas de Ceph (surcoût inutile). HA manuel (migration).
Coût maîtrisé, RPO 15 min.

**Cas B — Data center 3 nœuds Ceph**
3 nœuds (128 Go, 4 SSD OSD, 10 GbE dédié Ceph), pools rbd size=3,
HA sur les VM critiques, PBS dédié + sync hors site. RPO ~0, RTO
quelques minutes.

**Cas C — Ransomware un lundi matin**
Les VM chiffrées, mais : PBS avec **rétention immuable** + disque
externe mensuel déconnecté → restauration en 1 journée. Sans le
disque déconnecté, le chiffrement aurait touché les backups
connectés.

**Cas D — Disque SSD mort un vendredi soir**
ZFS mirror : la VM continue. `zpool status` DEGRADED → disque de
rechange en stock → `zpool replace` → resilver le week-end. Lundi :
personne n'a rien vu.

**Cas E — Montée en charge web**
Template + cloud-init + Terraform : 5 nouvelles VM en 10 min
pour absorber le pic. Puis `terraform destroy` après.

---

## 48. Dimensionnement — 3 architectures types

### 48.1 Site PME (budget serré)
- 2 nœuds (8 cœurs, 64 Go, 2×1 To SSD ZFS mirror).
- Réplication 15 min, PBS sur NAS, pas de HA auto.
- ~15 VM/conteneurs.

### 48.2 Production (standard)
- 3 nœuds (16 cœurs, 128 Go, 2×480 Go SSD système + 4×2 To Ceph).
- 10 GbE (Ceph + VM), HA, PBS dédié + sync distant.
- ~50 VM.

### 48.3 Gros site
- 5+ nœuds, Ceph NVMe, 25 GbE, 2 PBS (sync), Terraform, monitoring
  Grafana, NUT + groupe.

---

## 49. Bonnes pratiques production — les 20 règles

```
1. Jamais de service sur l'hôte (tout en VM/CT)
2. Nommage : pve01/02/03, vm-100-nom, ct-300-nom
3. Tags Proxmox par environnement (prod/test)
4. Notes sur chaque VM (rôle, responsable)
5. Templates + cloud-init (pas d'install manuelle)
6. Snapshots avant chaque changement
7. Backups PBS quotidiens + test mensuel
8. Mises à jour : fenêtre mensuelle, 1 nœud à la fois
9. Monitoring : Zabbix/Prometheus + alertes
10. Pare-feu : datacenter + VM
11. 2FA admins, 1 compte/personne
12. Quorum : nombre impair ou qdevice
13. Ceph : jamais > 85 %, jamais 2 nœuds rebootés ensemble
14. ZFS : scrub mensuel, < 80 %
15. NUT : arrêt propre sur coupure
16. Documentation : fiche par VM (§76)
17. Ansible/Terraform : pas de clics manuels répétitifs
18. Logs distants
19. Stock de disques de rechange (délai Afrique !)
20. Exercice de restauration annuel (comme un incendie)
```
---

## 50. NUT — l'onduleur parle à Proxmox (votre double casquette !)

### 50.1 Principe
```
Onduleur (USB/SNMP) → NUT (serveur) → upsmon (clients)
  Coupure → sur batterie → batterie basse → ARRÊT PROPRE
  des VM/CT puis des nœuds, AVANT la fin de l'autonomie.
```

### 50.2 Installation (serveur NUT sur un nœud ou un CT)
```bash
apt install nut-server nut-client
# /etc/nut/ups.conf
[monups]
  driver = usbhid-ups
  port = auto
  desc = "Easy UPS 40 kVA"
# /etc/nut/upsd.users
[admin]
  password = ***
  actions = SET
  instcmds = ALL
# /etc/nut/upsmon.conf
MONITOR monups@localhost 1 admin *** master
SHUTDOWNCMD "/sbin/shutdown -h +0"
# /etc/nut/nut.conf : MODE=netserver (serveur) / netclient (autres nœuds)
systemctl enable --now nut-server nut-monitor
upsc monups   # doit afficher les valeurs (charge %, batterie %, autonomie)
```

### 50.3 Seuils
- `upssched` : à « batterie basse » ou après X minutes sur batterie →
  script d'arrêt.
- **Règle** : déclencher l'arrêt quand il reste **2× le temps
  d'arrêt** des VM (ex. : arrêt = 5 min → seuil à 10 min restantes).

---

## 51. Arrêt automatique sur coupure — scénario complet

```bash
#!/bin/bash
# /usr/local/bin/coupure.sh — appelé par NUT (upssched)
/usr/bin/logger "COUPURE : arrêt propre des VM"
# 1. Arrêter les VM non critiques (ordre : applis → BDD → infra)
/usr/sbin/qm shutdown 101 --timeout 120 &
/usr/sbin/qm shutdown 102 --timeout 120 &
wait
# 2. Arrêter les conteneurs
/usr/sbin/pct shutdown 300 &
wait
# 3. Les nœuds s'arrêtent (upsmon SHUTDOWNCMD)
/sbin/shutdown -h +1 "Arret onduleur"
```

### 51.1 Au retour du courant
- BIOS : **restauration sur AC** (Power On) → les nœuds redémarrent.
- Proxmox : VM en « démarrer au boot » (Options → Start at boot,
  avec **délai/ordre** : infra d'abord, applis ensuite).
- Ordre type : PBS → contrôleur de domaine/DNS → BDD → applis.

### 51.2 Tester !
- **Test trimestriel** : débrancher l'onduleur (fenêtre planifiée),
  chronométrer, vérifier l'arrêt propre et le redémarrage.
- Un scénario NUT jamais testé = un scénario qui ne marche pas.

---

## 52. Politique de sauvegarde détaillée

```
QUOTIDIEN (02:00) : toutes les VM/CT → PBS (rétention 7j)
HEBDO (dimanche)  : vérification PBS (Verify job)
MENSUEL           : sync PBS → site distant + disque externe déconnecté
TEST              : restauration 1 VM/mois (tournant), rapport
EXCLUSIONS        : caches, /tmp, données re-générables (documentées)
```

### 52.1 Fenêtres et impact
- Mode snapshot : impact ~nul (quelques secondes de freeze I/O).
- Éviter les backups pendant les batchs lourds.
- **Jumbo** : étaler (VM critiques à 02:00, le reste à 03:00).

---

## 53. Restauration — procédures pas à pas

### 53.1 Restaurer une VM complète (PBS)
PBS → Datastore → Backups → [VM] → Restaurer → nouvel ID →
vérifier le réseau (MAC, IP) → démarrer → tester l'appli.

### 53.2 Restaurer un fichier (file restore)
PBS → File Restore → monter le backup → naviguer → télécharger
le fichier. **Sans restaurer toute la VM** (le cas n°1 en prod).

### 53.3 Reprise après sinistre (site perdu)
1. Nouveau cluster (même version).
2. PBS distant → restaurer les VM critiques d'abord (ordre §51.1).
3. Reconfigurer réseau/DNS.
4. **Objectif** : RTO < 1 jour (d'où l'exercice annuel).

---

## 54. PBS avancé

- **Namespaces** : isoler les clients (multi-sites, multi-équipes).
- **Sync jobs** : pull/push entre PBS (bande passante limitée :
  `rate-in`).
- **Tape backup** : LTO pour l'archivage longue durée (10 ans).
- **Prune** : politiques par groupe (garder les mensuels 1 an…).
- **Clés de chiffrement** : une par client, stockées hors site
  (coffre + papier).

---

## 55. Ceph avancé

### 55.1 CRUSH map
- Règles : séparer SSD/HDD, répartir par rack/salle (`crush
  location`).
- `ceph osd crush move` : placer les OSD correctement.

