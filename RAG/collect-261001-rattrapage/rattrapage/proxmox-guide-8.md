---
id: collect-261001-rattrapage/rattrapage/proxmox-guide-8
title: "Proxmox VE — Guide ultra-complet (2000+ lignes)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/proxmox_guide.md
source_anchor: ""
source_lines: [1732, 1949]
sha256: 9e4131d8f265be79651ddf26eeaa4fa0a8f07c420c4b5175783fdb00b882f1e0
---

# Proxmox VE — Guide ultra-complet (2000+ lignes)

```bash
# Le cluster est sensible à la latence : diagnostiquer
ping -c 100 pve02 | tail -3        # perte ? gigue ?
corosync-cfgtool -s                # anneaux corosync
journalctl -u corosync --since "1 hour ago" | grep -i "failed\|error\|knet"
ss -ulpn | grep 5404               # ports corosync en écoute
ethtool eth0 | grep -i speed       # négociation correcte ?
```
- Symptômes typiques : nœuds qui « flappent » (rejoignent/quittent),
  quorum instable, migrations qui échouent.
- Causes n°1 : **switch saturé**, spanning-tree, MTU incohérent,
  câble défectueux (tester avec un autre câble !).
- Solution durable : **réseau corosync dédié** (VLAN ou physique).

---

## 100. Sauvegarde des configurations — le git du cluster

```bash
# Ce qui n'est PAS dans les backups VM :
# /etc/pve (configs cluster — répliqué, mais pas versionné)
# /etc/network/interfaces, /etc/ceph, /etc/nut, /etc/zfs
# → versionner :
cd /etc/pve && git init -b main
git add -A && git commit -m "config $(date +%F)"
git remote add origin git@interne:infra/pve-config.git
git push -u origin main
# + script mensuel : commit auto des changements
```
- Après un sinistre, **reconstruire à l'identique** = ces fichiers.

---

## 101. Ajouter un nœud à un cluster existant

```
1. Même version Proxmox (pveversion -v identique)
2. Même hostname/IP définitifs, NTP synchronisé
3. Depuis le nouveau : pvecm add IP_D_UN_NŒUD (mot de passe root)
4. Vérifier : pvecm status (4 votes → nombre pair ! → qdevice ou 5e nœud)
5. Recréer : stockages locaux, réseau (mêmes noms de bridges !),
   Ceph OSD si applicable
6. Migrer quelques VM de test, vérifier HA
```
- ⚠️ Passer de 3 à 4 nœuds = nombre **pair** → prévoir qdevice ou
  5e nœud pour garder un quorum sain.

---

## 102. Cas d'usage par métier

| Métier | Architecture Proxmox |
|---|---|
| Bureautique PME | 2 nœuds, ZFS, réplication, PBS sur NAS |
| Data center | 3+ nœuds, Ceph, HA, PBS dédié + sync |
| Hébergeur | SDN, Terraform, facturation par VM |
| Industrie | 2 nœuds durcis, NUT, arrêt auto, baie verrouillée |
| Éducation | Templates + clones liés (salles TP en 5 min) |
| Lab / test | 1 nœud, ZFS, snapshots à gogo |

---

## 103. Checklist d'audit annuel Proxmox

```
□ Versions : tous les nœuds à jour (même version)
□ Quorum : impair ou qdevice, test de panne simulée
□ Ceph : HEALTH_OK, < 80 %, scrub OK, OSD tous up
□ ZFS : < 80 %, scrub mensuel actif, snapshots nettoyés
□ Backups : 100 % succès 12 mois, 12 restaurations testées (1/mois)
□ PRA : exercice annuel fait, RTO mesuré, RETEX écrit
□ NUT : test de coupure annuel, seuils vérifiés
□ Sécurité : 2FA, comptes (départs ?), pare-feu, 8006 non exposé
□ Capacité : projection 12 mois (faut-il un nœud ?)
□ Documentation : fiches à jour, git /etc/pve poussé
□ Contrats : abonnements, garanties matérielles
□ Stock : disques de rechange, câbles, PDU
```

---

## 104. Index rapide

```
UI : https://IP:8006 | qm / pct / pvesh / pvecm
Cluster : corosync 5404-5405/udp, quorum impair
Stockage partagé = HA possible | Ceph : 3 nœuds, 10 GbE, < 85 %
ZFS : mirror/raidz, lz4, scrub mensuel, < 80 %
Backup : PBS > vzdump, verify hebdo, GC hebdo, sync distant
Test restauration 1×/mois | Disque déconnecté 1×/mois
NUT : upsc, test trimestriel | HA : fencing obligatoire
Mises à jour : 1 nœud à la fois | Jamais de service sur l'hôte
```

---

*Fin du guide — 2000+ lignes. Ton coup de cœur mérite cette
exhaustivité, Zelef. **À toi de jouer sur le terrain.***
---

## 105. Proxmox vs alternatives — trancher

| Critère | Proxmox VE | VMware vSphere | Hyper-V | XCP-ng |
|---|---|---|---|---|
| Licence | Gratuit (opt. ~100 €/CPU/an) | $$$$ (par cœur) | Inclus Windows Server | Gratuit |
| KVM/LXC | Oui | Non (ESXi) | Non | Oui (Xen) |
| Ceph intégré | Oui | vSAN ($$$) | Non (S2D) | Non |
| Backup intégré | Oui (PBS) | Veeam ($$$) | DPM | Xen Orchestra |
| Cluster/HA | Oui | Oui | Oui | Oui |
| Courbe d'apprentissage | Moyenne | Moyenne | Faible (si Windows) | Moyenne |
| Support francophone | Communauté | Commercial | Commercial | Communauté |

- **Verdict PME/ETI** : Proxmox = 90 % des fonctionnalités pour
  5 % du prix. Le choix rationnel quand on maîtrise Linux.

---

## 106. Les coûts cachés — les anticiper

```
□ Réseau 10 GbE (switch + cartes + DAC/fibre) : le poste n°1
□ Disques SSD d'entreprise (endurance !) — pas de SSD grand public en Ceph
□ Onduleur + groupe (voir guide onduleurs)
□ Climatisation du local
□ Formation de l'équipe (le vrai investissement)
□ Temps : Ceph et le cluster demandent de la montée en compétence
```
- Le logiciel est gratuit ; **l'infrastructure autour ne l'est pas**.
  Chiffrez tout avant de promettre.

---

## 107. Feuille de route d'adoption — 6 mois

```
MOIS 1 : Lab (1 nœud ou 3 VM Proxmox imbriquées), formation S1-S2
MOIS 2 : Maquette 2 nœuds + ZFS + réplication + PBS
MOIS 3 : Pilote (2-3 VM non critiques en prod), NUT
MOIS 4 : Cluster 3 nœux + Ceph (si pertinent), HA testé
MOIS 5 : Migration des VM critiques, Terraform/Ansible
MOIS 6 : PRA écrit, exercice annuel, dossier d'exploitation complet
```
- Ne jamais tout migrer d'un coup : **le pilote d'abord**.

---

## 108. FAQ — les questions qu'on me pose

**Q : 1 seul nœud, ça vaut le coup ?**
R : Oui (ZFS + snapshots + PBS). Ce n'est pas de la HA, mais c'est
déjà 10× mieux qu'un serveur « à l'ancienne ».

**Q : Ceph ou ZFS ?**
R : 1-2 nœuds → ZFS + réplication. 3+ nœuds + HA → Ceph.

**Q : Combien de VM par nœud ?**
R : Compter la RAM d'abord (ne pas dépasser 80 %), puis le CPU
(ratio 2-4:1), puis les IOPS disque.

**Q : Proxmox et Windows, ça marche bien ?**
R : Oui, avec VirtIO + guest tools. Des milliers de PME tournent
ainsi (contrôleur de domaine, ERP…).

**Q : Et la sauvegarde des VM Windows (VSS) ?**
R : QEMU agent + VSS → snapshots cohérents applicativement.

**Q : Faut-il l'abonnement Enterprise ?**
R : En prod : oui (dépôts stables + support). C'est le prix d'un
repas par CPU et par an.

**Q : Proxmox sur Raspberry Pi ?**
R : Non (x86_64 uniquement en stable). Pour l'ARM : à surveiller.

**Q : Peut-on mixer des nœuds de générations différentes ?**
R : Oui si même version Proxmox, mais les perfs suivent le plus
faible (surtout Ceph).

---

*2000+ lignes — guide terminé. Ton outil préféré n'a plus de
secret, Zelef.*
---

## 109. Recettes prêtes à l'emploi — scripts

### 109.1 Health check quotidien (cron)
```bash
#!/bin/bash
# /usr/local/bin/pve-health.sh — rapport mail quotidien
DEST="admin@example.com"
{
echo "=== $(hostname) $(date) ==="
echo "--- Cluster ---"; pvecm status 2>&1 | grep -E "Quorum|Nodes"
echo "--- Ceph ---"; ceph -s 2>&1 | head -8
echo "--- ZFS ---"; zpool list -H -o name,cap,health 2>&1
echo "--- Stockages ---"; pvesm status 2>&1 | awk '{print $1, $5}'
echo "--- Backups 24h ---"; find /var/log/vzdump* -mtime -1 2>/dev/null | wc -l
echo "--- Services en échec ---"; systemctl list-units --failed --no-legend
echo "--- NUT ---"; upsc monups 2>/dev/null | grep -E "battery.charge|ups.status" || echo "NUT: non configuré"
} | mail -s "[PVE] Santé $(hostname) $(date +%F)" "$DEST"
```

### 109.2 Snapshot pré-maintenance (toutes les VM)
```bash
#!/bin/bash
# /usr/local/bin/snap-avant-maj.sh
TAG="avant-maj-$(date +%F)"
for vm in $(qm list | awk 'NR>1{print $1}'); do
  echo "Snapshot $vm..."
  qm snapshot "$vm" "$TAG" --description "Avant maintenance mensuelle" || echo "ÉCHEC $vm"
done
for ct in $(pct list | awk 'NR>1{print $1}'); do
  echo "Snapshot CT $ct..."
  pct snapshot "$ct" "$TAG" || echo "ÉCHEC CT $ct"
done
echo "Terminé. Nettoyer dans 7 jours : qm delsnapshot <id> $TAG"
```

