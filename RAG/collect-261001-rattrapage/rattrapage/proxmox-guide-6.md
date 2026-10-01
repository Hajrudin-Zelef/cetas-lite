---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-6
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "datacenter", "incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [1235, 1471]
sha256: 9d59738348cd2a89aff25b62d83828febccdf74d15c064db723ebaeb16e2101e
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

| Terme | Signification |
|---|---|
| PVE | Proxmox Virtual Environment |
| PBS | Proxmox Backup Server |
| KVM | Virtualisation complète (VM) |
| LXC | Conteneurs système Linux |
| qm / pct | CLI pour VM / conteneurs |
| pvesh | CLI pour l'API |
| corosync | Protocole de cluster (quorum) |
| qdevice | Arbitre externe (quorum à 2 nœuds) |
| HA | Haute disponibilité (redémarrage auto) |
| Fencing | Isolation d'un nœud défaillant |
| OSD | Disque Ceph |
| PG | Placement Group (Ceph) |
| RBD | Block device Ceph |
| ZFS | Filesystem (snapshots, checksums) |
| ARC | Cache RAM de ZFS |
| Scrub | Vérification d'intégrité ZFS |
| vzdump | Outil de sauvegarde intégré |
| Template | Modèle de VM/CT |
| cloud-init | Config auto des VM au boot |
| NUT | Network UPS Tools |
| RPO / RTO | Perte max / durée max de reprise |

---

## 73. Fiche VM/CT — modèle (une par machine)

```
ID : 100  Nom : srv-web-01  Type : VM (KVM)
Nœud : pve01  Cluster : prod
OS : Ubuntu 24.04  Rôle : serveur web intranet
CPU : 2  RAM : 4 Go  Disque : 30 Go (ceph-rbd)
Réseau : vmbr0, VLAN 10, IP 192.168.10.21/24
Responsable : ___  Mise en service : ___
Backup : PBS quotidien 02:00  HA : oui (groupe prod)
Notes : ___
```

---

## 74. Fiche nœud — modèle

```
Nœud : pve01  IP : 192.168.1.11  Rôle : hyperviseur prod
Matériel : ___ (CPU/RAM/disques, n° série, garantie jusqu'au ___)
OS : Proxmox VE ___  Dépôt : enterprise
Stockage : ZFS tank (mirror 2×1 To), Ceph (4 OSD)
Réseau : bond0 (eth0+eth1 LACP), vmbr0, VLAN aware
NUT : oui (client du serveur nut sur pve02)
Dernière maj : ___  Prochaine fenêtre : ___
```

---

## 75. Fiche site Proxmox — modèle

```
SITE : ___
Cluster : ___ (nœuds : ___, quorum : ___)
Stockage : ___ (Ceph/ZFS/NFS, % utilisé : ___)
PBS : ___ (datastore ___, sync distant : ___)
Onduleur : ___ kVA (NUT : oui/non, testé le ___)
Groupe : ___ kVA (testé le ___)
Supervision : ___  Alertes : ___
Contrat : ___  Échéance : ___
Dossier d'exploitation : ☐ à jour le ___
```

---

## 76. KPI — piloter le service virtualisation

- **Disponibilité** des VM critiques (objectif 99,9 %).
- **RPO/RTO** réels (mesurés sur le dernier test).
- **Sauvegardes** : % succès, dernier test de restauration.
- **Capacité** : % RAM/CPU/stockage par nœud (alerte à 80 %).
- **Ceph** : HEALTH_OK permanent (tout WARN = incident).
- **Patchs** : délai de mise à jour après release (objectif < 30 j).
- **Rapport mensuel** : 1 page (dispo, incidents, capacité,
  backups, risques).

---

## 77. Contrats de maintenance — virtualisation

- **Périmètre** : hyperviseurs, Ceph, PBS, réseau, onduleur+NUT.
- **Niveaux** : Bronze (1 visite/an), Silver (trimestriel +
  supervision), Gold (mensuel + astreinte 4 h + pièces).
- **Inclure** : test de restauration annuel **contractuel**,
  exercice NUT annuel, rapport de capacité semestriel.
- **Chiffrage** : visites + astreinte + pièces (disques de
  rechange !) + marge.

---

## 78. Gestion d'équipe — virtualisation

- **1 référent Proxmox** (cluster, Ceph, PBS) + techniciens
  (VM courantes, backups, tickets).
- **Règles** : pas de clic manuel répétitif (Ansible/Terraform),
  toute VM a une fiche, tout changement = ticket GLPI.
- **Astreinte** : avec runbook (fiches réflexe §70 + alertes).
- **Formation** : 1 admin formé « Ceph » minimum (c'est le
  composant le plus délicat).

---

## 79. Script d'audit mensuel Proxmox

```bash
#!/bin/bash
# /usr/local/bin/audit-pve.sh
echo "=== AUDIT PVE $(date +%F) ==="
pvecm status | grep -A5 "Quorum"
echo "[CEPH]"; ceph -s | head -12
echo "[ZFS]"; zpool list -H -o name,cap,health
echo "[STOCKAGE]"; pvesm status
echo "[VM]"; qm list | tail -n +2 | wc -l; echo "conteneurs :"; pct list | tail -n +2 | wc -l
echo "[HA]"; ha-manager status 2>/dev/null | head -5
echo "[BACKUPS]"; ls -la /var/log/vzdump* 2>/dev/null | tail -3
echo "[MAJ]"; apt list --upgradable 2>/dev/null | grep -c pve
echo "[NUT]"; upsc monups 2>/dev/null | grep -E "battery.charge|ups.load|battery.runtime" || echo "NUT non configuré"
```
- À lancer via Ansible sur tous les nœuds, sortie centralisée.

---

## 80. Mémo commandes — tout en une page

```bash
# Cluster
pvecm status | pvecm nodes | pvecm add IP | pvecm delnode nom
# VM
qm list | qm start|stop|shutdown|reset|destroy 100
qm snapshot|rollback|delsnapshot 100 nom
qm clone 100 200 --name x --full
qm migrate 100 pve02 --online
qm resize 100 scsi0 +20G | qm move-disk 100 scsi0 stockage
qm set 100 --memory 4096 --cores 2
qm template 100
# Conteneurs
pct list | pct start|stop|shutdown 300 | pct enter 300
pct snapshot|rollback 300 nom
# Stockage
pvesm status | pvesm add nfs ...
zpool status -x | zfs list -t snapshot
ceph -s | ceph osd tree | ceph df
# Backup
vzdump 100 --mode snapshot --compress zstd --storage pbs
qmrestore /chemin/vzdump... 200
# Divers
pve-firewall status | ha-manager status
pvesh get /nodes/pve01/status
upsc monups | upsmon -c fsd   # NUT : test d'arrêt
```
---

## 81. Checklists imprimables

### 81.1 Installation d'un nœud
```
□ Matériel vérifié (VT-x, disques, RAM, 2 cartes réseau)
□ BIOS : virtualisation, IOMMU, boot UEFI, power-on AC
□ ISO vérifiée (SHA256), installée (ZFS mirror si 2 disques)
□ IP fixe, FQDN, DNS, NTP
□ Dépôts configurés (enterprise/no-subscription)
□ apt full-upgrade, reboot
□ Stockage vérifié, réseau (vmbr0, bond si prévu)
□ Alertes email, certificat ACME
□ Utilisateurs (pas de root partagé), 2FA
□ Pare-feu datacenter
□ Intégré au cluster (si prévu)
□ PBS : stockage ajouté, job de backup créé
□ NUT (si onduleur)
□ Fiche nœud remplie (§74)
```

### 81.2 Maintenance mensuelle (par nœud)
```
□ pvecm status (quorum), ceph -s (HEALTH_OK), zpool status -x
□ Espace : pvesm status (< 80 %), PBS datastore (< 85 %)
□ Backups : 100 % succès sur 30 j, 1 restauration testée
□ Mises à jour : fenêtre planifiée si nécessaire
□ Logs : erreurs (journalctl -p err --since "30 days ago" | head)
□ Températures, SMART (smartctl -a, usure SSD)
□ NUT : upsc (charge %, autonomie)
□ Rapport mensuel (§76)
```

---

## 82. Cas pratiques supplémentaires

**Cas F — Corosync qui flappe après ajout d'un switch**
Le nouveau switch a le spanning-tree lent → corosync perd des
paquets → nœuds qui s'excluent. → Portfast sur les ports Proxmox,
ou réseau cluster dédié. Leçon : le réseau du cluster est sacré.

**Cas G — Restauration d'un fichier effacé par un utilisateur**
PBS → File Restore → le fichier est de retour en 5 min, sans
arrêter la VM. Le cas qui justifie PBS à lui seul.

**Cas H — Montée de version Proxmox 8 → 9**
Un nœud à la fois : migrer les VM, `apt full-upgrade`, reboot,
vérifier (cluster, Ceph, HA), passer au suivant. Jamais les 3
d'un coup.

**Cas I — Panne de courant sans NUT (leçon apprise)**
Coupure 2 h, les nœuds s'arrêtent brutalement, 1 VM corrompue
(disque). → NUT installé la semaine suivante + test trimestriel.
Ne pas attendre la leçon pour agir.

**Cas J — Espace Ceph à 90 %**
Alertes ignorées 2 semaines → Ceph passe en read-only → VM
figées. → Nettoyage d'urgence (vieux snapshots, backups),
ajout d'OSD. Règle : alerte à 75 %, action à 80 %, jamais 85 %.

---

## 83. Proxmox + énergie — dimensionner l'électricité

```
Par nœud (16 cœurs, 128 Go, 4 SSD) : ~300–500 W en charge
3 nœuds + switch + PBS : ~1,5–2 kW
→ Onduleur : 3 kVA minimum (10 min), 6–10 kVA conseillé
→ Groupe : si site critique (voir guide onduleurs)
Prise : circuit dédié 16 A par baie, PDU avec mesure
Sonde température : alerte à 30 °C dans la baie
```
- Le volet **énergies** et le volet **systèmes** se rejoignent ici :
  un cluster sans onduleur, c'est une épave en attente de coupure.

---

## 84. Sécurité physique

