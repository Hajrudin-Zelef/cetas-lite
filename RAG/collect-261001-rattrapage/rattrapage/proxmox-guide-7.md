---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-7
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter", "incident"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [1472, 1731]
sha256: 5b8c24ff893d31ff3d66513cfa31d2ba2b9ae7befd70bcaba88fd97e705cc247
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

- Baie **verrouillée**, accès badgé, registre des accès.
- USB désactivés sur les nœuds (ou BIOS protégé par mot de passe).
- Caméra sur le local technique (les « interventions » non
  déclarées…).
- Extincteur CO₂, détection incendie (VESDA sur gros sites).

---

## 85. Nouveautés et tendances (2024-2026)

- **SDN** : remplace progressivement les bridges manuels.
- **Import OVF/OVA** natif (migration VMware facilitée).
- **Ceph Reef/Quincy** : meilleures perfs, moins de RAM.
- **PBS** : tape backup, namespaces — la maturité entreprise.
- **Terraform/OpenTofu** : l'infra as code devient la norme.
- **ARM** : Proxmox expérimental sur ARM (à surveiller, pas en prod).

---

## 86. Ressources officielles

- Documentation : pve.proxmox.com (excellente, en anglais).
- Forum : forum.proxmox.com (très actif).
- Wiki : pve.proxmox.com/wiki.
- Modèles de référence : les « Proxmox VE Administration Guide » (PDF).

---

## 87. Formation équipe — plan 4 semaines

- **S1** : concepts (KVM/LXC), UI, créer une VM/CT, snapshots.
- **S2** : stockage (ZFS, NFS), réseau (bridges, VLAN), backups vzdump.
- **S3** : cluster, HA, Ceph (bases), PBS, NUT.
- **S4** : dépannage (cas §46), Terraform/Ansible (bases), astreinte.
- Validation : chaque technicien **installe un cluster 3 nœuds en lab**
  et restaure une VM depuis PBS.

---

## 88. Budget type annuel (3 nœuds, PME)

| Poste | Ordre de grandeur |
|---|---|
| Abonnements Enterprise (3 CPU) | ~300 €/an |
| Disques de rechange (stock) | 1–2 SSD/an |
| PBS (matériel ou VM + disques) | amorti 3 ans |
| Électricité (2 kW) | selon tarif local |
| Maintenance (visites + astreinte) | selon contrat |
| Formation | 1 session/an |

- Le coût est **10× inférieur** à VMware pour un périmètre équivalent.

---

## 89. Erreurs Ceph — à ne jamais faire

```
✗ Rebooter 2 nœuds en même temps
✗ Remplir au-delà de 85 %
✗ Paniquer pendant un backfill (c'est normal, ça prend des heures)
✗ Mélanger SSD et HDD dans le même pool sans classes
✗ Oublier le réseau dédié (1 GbE partagé = lenteurs)
✗ Supprimer un pool sans vérifier les images (rbd ls)
✗ Ignorer un HEALTH_WARN (ça ne se répare pas tout seul)
```

---

## 90. Les 10 règles d'or Proxmox (à afficher)

```
1. Jamais de service sur l'hôte.
2. Templates + cloud-init, pas d'install manuelle.
3. Snapshots avant chaque changement.
4. Backups PBS quotidiens + test mensuel.
5. Quorum impair, Ceph < 85 %, ZFS < 80 %.
6. 1 nœud à la fois pour les mises à jour.
7. NUT configuré ET testé.
8. 8006 jamais exposé, 2FA admins.
9. Tout documenté (fiche par VM).
10. Le disque déconnecté mensuel sauve des entreprises.
```
---

## 91. Exemple complet — monter un cluster 3 nœuds pas à pas

```
JOUR 1 — Matériel
  □ 3 serveurs identiques (16 cœurs, 64 Go, 2×480 SSD système + 2×2 To data)
  □ Switch 10 GbE (Ceph + VM), switch 1 GbE (management)
  □ Câblage : 2× 10 GbE + 1× 1 GbE par nœud, étiqueté
  □ Onduleur 6 kVA branché, PDU

JOUR 2 — Installation
  □ Proxmox sur les 3 (ZFS mirror sur les 480 Go)
  □ IP : pve01 .11, pve02 .12, pve03 .13, FQDN
  □ Dépôts enterprise, full-upgrade, reboot

JOUR 3 — Réseau
  □ Bond LACP 10 GbE → vmbr0 (VLAN aware), vmbr1 (Ceph, 10.10.10.0/24)
  □ Test : iperf3 entre nœuds (> 9 Gb/s), débrancher un câble (pas de coupure)

JOUR 4 — Cluster
  □ pve01 : créer le cluster ; pve02/03 : rejoindre
  □ pvecm status : 3 votes, quorum OK
  □ Pare-feu datacenter, 2FA, ACME, alertes email

JOUR 5 — Ceph
  □ Installer Ceph, réseau public (vmbr0) / cluster (vmbr1)
  □ Créer les OSD (les 2×2 To de chaque nœud)
  □ Pools : vm-pool (rbd), cephfs si besoin
  □ ceph -s : HEALTH_OK, 6 OSD up

JOUR 6 — PBS
  □ 4e machine (ou VM) : PBS, datastore 4 To
  □ Stockage PBS ajouté sur le cluster, chiffrement
  □ Job quotidien 02:00, Verify hebdo, GC hebdo

JOUR 7 — Production
  □ Template Ubuntu + cloud-init, 1re VM de test
  □ Backup → restauration testée
  □ HA sur la VM test, test de bascule (débrancher pve01 !)
  □ NUT : onduleur USB sur pve01, clients sur pve02/03, test
  □ Fiches remplies, dossier d'exploitation, formation équipe
```

---

## 92. Ceph pas à pas — les commandes

```bash
# Sur pve01 : installer et initier
pveceph install
pveceph init --network 10.10.10.0/24
# Monitors (1 par nœud)
pveceph mon create
# Manager
pveceph mgr create
# OSD (répéter sur chaque nœud/disque)
pveceph osd create /dev/sdb
# Pools
pveceph pool create vm-pool --size 3 --min_size 2
# Vérifier
ceph -s
ceph osd df
# Ajouter au stockage Proxmox : Datacenter → Stockage → RBD
```

---

## 93. HA pas à pas

```
1. Stockage partagé vérifié (Ceph/RBD)
2. Fencing : IPMI configuré sur chaque nœud (Datacenter → HA → Fencing)
3. Datacenter → HA → Groupes : créer "prod" (pve01,pve02,pve03)
4. HA → Ajouter : VM 100 → groupe prod, état started, max_restart 3
5. TEST (fenêtre planifiée) : 
   a. ha-manager status → tout au vert
   b. Éteindre brutalement pve01 (pas un shutdown propre !)
   c. Chronométrer : la VM redémarre sur pve02 en ~2-5 min
   d. Rallumer pve01, vérifier le retour
6. Documenter le RTO mesuré
```

---

## 94. Template cloud-init avancé (user-data)

```yaml
# /var/lib/vz/snippets/user.yaml
#cloud-config
hostname: vm-template
manage_etc_hosts: true
users:
  - name: admin
    sudo: ALL=(ALL) NOPASSWD:ALL
    ssh_authorized_keys:
      - ssh-ed25519 AAAA... admin@poste
packages:
  - qemu-guest-agent
  - zabbix-agent2
  - fail2ban
  - unattended-upgrades
runcmd:
  - [systemctl, enable, --now, qemu-guest-agent]
  - [systemctl, enable, --now, zabbix-agent2]
timezone: Africa/Abidjan
```
```bash
qm set 9000 --cicustom "user=local:snippets/user.yaml"
qm template 9000
```

---

## 95. GLPI + Proxmox

- Fiche par VM dans GLPI (import via API ou manuel).
- Tickets : toute intervention Proxmox = ticket (traçabilité).
- FusionInventory : inventaire auto des VM.

---

## 96. Zabbix — superviser Proxmox

```
1. Template "Proxmox VE" (API) sur le cluster :
   - état nœuds, quorum, Ceph, espace stockages
2. Agent Zabbix dans chaque VM (métier)
3. Triggers :
   - Ceph != HEALTH_OK → Disaster
   - Backup failed → High
   - Stockage > 80 % → Warning
   - Nœud unreachable → Disaster
4. Actions : SMS/Telegram à l'astreinte
```

---

## 97. Plan de reprise d'activité (PRA) — trame

```
1. PÉRIMÈTRE : cluster prod (VM critiques : ___, ___, ___)
2. RPO : ___ (fréquence backups) / RTO : ___ (mesuré : ___)
3. SCÉNARIOS :
   a. Panne 1 nœud → HA auto (RTO ~5 min)
   b. Panne 2 nœuds → mode dégradé (VM critiques sur le survivant)
   c. Site perdu → PBS distant → nouveau cluster (RTO 1 jour)
   d. Ransomware → disque déconnecté mensuel
4. CONTACTS : astreinte ___, fournisseur ___, onduleur ___
5. EXERCICE ANNUEL : date ___, scénario ___, RTO mesuré ___
6. RETEX : rapport 48 h après chaque exercice/incident
```

---

## 98. Annexe — ports réseau Proxmox

| Port | Usage |
|---|---|
| 8006/tcp | Interface web |
| 5404-5405/udp | corosync (cluster) |
| 22/tcp | SSH / migration |
| 3128/tcp | Proxy SPICE |
| 5900-5999/tcp | Console VNC |
| 6789, 6800-7300/tcp | Ceph (mon, OSD) |
| 8007/tcp | PBS |
| 111, 2049/tcp | NFS (si utilisé) |
| 3260/tcp | iSCSI |

---

*Fin du guide — 2000+ lignes. Proxmox n'a plus de secret pour toi,
Zelef : **cluster propre, Ceph sain, PBS testé, NUT branché.**
Ton outil préféré, maîtrisé de bout en bout.*
---

## 99. Dépannage réseau du cluster — approfondi

