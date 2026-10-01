---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-22
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Samsung"]
dates: ["2026-09-27"]
keywords: ["datacenter", "capex", "compute", "incident", "nand"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [3629, 3824]
sha256: 93aa54ea0c5e1cc21b4fc82c57e7ce7ea78a65ecef8e67c6aa1ff522213bf6aa
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

RGW supporte les règles de lifecycle (transition vers EC froid,
expiration). Configure dès la création des buckets applicatifs :
`logs/` → expiration 90 j, `backups/` → transition EC 30 j. Sans
politique, les buckets ne font que grossir — c'est mathématique.

## 376. Le monitoring des coûts (FinOps stockage)

Tag par projet/pool : Go/mois, IOPS, €. Facturation interne au To
avec les vrais coûts (TCO, pas CAPEX). Effet : les équipes nettoient
spontanément quand elles voient la facture. Le stockage « gratuit »
est toujours gaspillé.

## 377. Les erreurs humaines : la défense en profondeur

`rm`, `rbd snap purge`, mauvaise règle CRUSH : la défense =
confirmations (`--yes-i-really-mean-it`), RBAC (pas d'admin à tout
le monde), snapshots avant opérations, et **délai de grâce** (corbeille,
soft-delete RGW). On ne supprime jamais vite en production.

## 378. La documentation vivante

La doc qui n'est pas relue meurt. Rituel : à chaque incident, la
section concernée est relue et corrigée. Responsable doc nommé,
revue trimestrielle de 30 min. Une doc de 50 pages à jour vaut mieux
que 500 pages obsolètes.

## 379. La veille technologique : le rituel

30 min/mois : release notes Ceph, roadmaps NAND (Samsung/Solidigm/
KIOXIA/Micron), specs SNIA/EDSFF/CXL, Greensheet marché. C'est ce
rituel qui t'évite d'acheter de l'U.2 en 2027 ou de rater les 245 To.
Ce guide est un instantané au 27/09/2026 : **re-vérifie les prix,
les dispos et les roadmaps à chaque projet**.

## 380. Conclusion : l'esprit du guide

Le stockage datacenter 2026, c'est : du NVMe E3.S partout où la
latence compte, du HDD CMR là où le To compte, du logiciel (Ceph)
entre les deux, et de l'énergie en fil rouge. Les formules de ce
guide tiennent sur deux pages (section 339) ; tout le reste, c'est
de la discipline d'exécution : burn-in, monitoring, spares, game
days, doc. Les clusters qui survivent 5 ans sans incident majeur
ne sont pas les plus chers — ce sont les mieux opérés.

---

**FIN DU GUIDE — 380 sections numérotées, vérifications web au 27/09/2026.**
**Fichier : `~/workspace/user/files/datacenter_stockage_guide.md`**
**Rappel : prix en allocation tendue (sept. 2026) — devis datés exigés, faits marqués « à vérifier » à valider.**
**Guides liés : `proxmox_guide.md`, `onduleurs_ups_guide.md`, `zabbix_guide.md`, `debian_ubuntu_guide.md`.**

---

# PARTIE X — ANNEXES PRATIQUES

## 381. Commandes HBA : storcli (mémo)

```
storcli /c0 show                    # état contrôleur
storcli /c0 show all                # tout (long)
storcli /c0/eall/sall show          # tous les disques
storcli /c0/e0/s4 show              # détail un disque
storcli /c0/e0/s4 start locate      # LED localisation
storcli /c0/e0/s4 stop locate
storcli /c0 show rebuild            # rebuilds en cours
```
En script : parse la sortie `show all` et pousse les compteurs
d'erreurs vers Zabbix. Une erreur média qui augmente = RMA.

## 382. Commandes NVMe : nvme-cli (mémo)

```
nvme smart-log /dev/nvme0           # santé (percentage_used !)
nvme error-log /dev/nvme0           # erreurs
nvme id-ctrl /dev/nvme0             # identité, firmware
nvme format /dev/nvme0 --ses=2      # crypto erase (SED)
nvme set-feature /dev/nvme0 -f 0x0c  # états de puissance
```
`percentage_used` > 70 % = planifier le remplacement (section 180).

## 383. Commandes SMART : smartctl (mémo)

```
smartctl -a /dev/sdX                # tout
smartctl -t short /dev/sdX          # test court
smartctl -t long /dev/sdX           # test long (heures)
smartctl -l selftest /dev/sdX       # résultats
smartd                              # daemon de surveillance
```
Attributs critiques : 5, 187, 188, 197, 198 (section 166).

## 384. Commandes Ceph : le mémo quotidien

```
ceph -s ; ceph health detail        # santé
ceph osd df tree                    # OSD, PG, remplissage
ceph osd pool autoscale-status      # PG vs cible
ceph pg stat                        # états PG
ceph osd out <id>                   # sortir un OSD
ceph balancer status                # équilibrage
rados bench -p <pool> 30 write      # bench (pool dédié !)
ceph daemon osd.<id> perf dump       # perfs fines
```

## 385. Commandes réseau : le mémo

```
ethtool -S ethX | grep -i 'crc\|drop\|err'   # erreurs
ping -M do -s 8972 <ip>                      # MTU 9000 bout-en-bout
lspci -vv | grep -A2 LnkSta                  # PCIe négocié
ibstat / ibv_devices                         # RDMA si présent
```

## 386. Matrice câbles : que brancher où

| De → Vers | Câble | Débit max |
|---|---|---|
| HBA 9500-16e → JBOD 12G | SFF-8644 ↔ SFF-8644 | 4×12G par câble |
| HBA 9600-16e → JBOD 24G | SFF-8674 ↔ SFF-8674 | 4×24G par câble |
| HBA 9500-16i → backplane | SFF-8654 ↔ SFF-8654/8643* | selon brochage |
| Backplane NVMe direct | MCIO / SlimSAS x4 | Gen5 x4 par baie |

\* 8654↔8643 : uniquement avec adaptateur/câble au bon brochage
(SFF-9402) — vérifie la doc du châssis (section 218).

## 387. Matrice disques : que mettre où

| Disque | Baie OK | Usage |
|---|---|---|
| HDD SAS 12G | SAS 12G/24G | capacitaire, dual-path |
| HDD SATA | SAS/SATA | capacitaire simple |
| SSD SATA | SAS/SATA | boot, db petit cluster |
| SSD NVMe U.2 | U.2 / U.3 | existant Gen4 |
| SSD NVMe U.3 | U.3 (UBM) | mixte tri-mode |
| SSD E3.S | E3.S | **neuf 2026+** |
| SSD E1.S | E1.S | 1U dense |

## 388. Glossaire complémentaire (20 termes)

| Terme | Définition |
|---|---|
| AOC | Active Optical Cable : câble optique actif (10-100 m). |
| AWG | Jauge du câble (plus petit = plus gros = plus loin). |
| Backplane | Fond de panier : connectique des baies. |
| Blast radius | Périmètre d'impact d'une panne. |
| BBU / CacheVault | Protection du cache RAID (batterie/supercondensateur). |
| CMR | Conventional Magnetic Recording (HDD datacenter). |
| COW | Copy-On-Write (snapshots, clones). |
| DDT | Dedup Table (ZFS) : ~5 Go RAM par To dédupliqué. |
| ECN / PFC | Contrôle de congestion / de flux (RoCE lossless). |
| FDP | Flexible Data Placement (NVMe, section 211). |
| GFS | Grandfather-Father-Son (rétention backup). |
| LACP | Agrégation de liens (802.3ad). |
| MLAG | LACP multi-châssis (2 switchs). |
| OCP | Open Compute Project (specs SSD 2.x). |
| PLP | Power Loss Protection (condensateurs SSD). |
| RPO / RTO | Perte max / reprise max (PRA). |
| SED | Self-Encrypting Drive (TCG Opal). |
| SMR | Shingled Magnetic Recording (à bannir en RAID). |
| WORM | Write Once Read Many (archivage, Object Lock). |
| ZNS | Zoned Namespaces (NVMe, section 346). |

## 389. FAQ achats (1/2)

**Q : 12G ou 24G pour du HDD neuf ?**
R : 12G. Le 24G n'apporte rien à des HDD à 300 Mo/s (section 44).
**Q : U.2 ou E3.S pour du neuf ?**
R : E3.S (section 66) — sauf contrainte de châssis existant.
**Q : Combien de PG pour démarrer ?**
R : Laisse l'autoscaler faire, pools en `bulk: true` (section 110).
**Q : RAID matériel ou HBA pour Ceph ?**
R : HBA IT, toujours (section 29).

## 390. FAQ achats (2/2)

**Q : Quelle endurance pour du block.db ?**
R : 3 DWPD TLC minimum, jamais de QLC (sections 170, 222).
**Q : Faut-il du 100G pour Ceph ?**
R : Oui en full NVMe (cluster séparé), 25G suffit en HDD (section 114).
**Q : Quel onduleur pour 4 nœuds NVMe ?**
R : 10 kVA mini (section 189), avec le pic de spin-up si HDD (200).
**Q : MinIO ou Ceph RGW ?**
R : MinIO = simple/mono-site ; RGW = critique/multi-sites (150).

## 391. Tableau : qui consomme quoi (rappel énergie)

| Équipement | Conso typique | Source |
|---|---|---|
| HDD 20-30 To | 8-10 W actif | ordre de grandeur |
| NVMe Gen5 7,68 To | 17-18 W actif / 5 W idle | vérifié 27/09/2026 |
| HBA 9500-16e | 8,74 W | vérifié 27/09/2026 |
| NIC 100G | ~20 W | ordre de grandeur |
| Switch 32×100G | ~300 W | ordre de grandeur |
| Nœud 60 baies complet | ~980 W au mur | calcul section 184 |
| Nœud 12 NVMe complet | ~800 W au mur | calcul section 185 |

