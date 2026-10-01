---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-14
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom"]
dates: ["2026-09-27"]
keywords: ["datacenter", "capex", "compute", "ethernet", "gpu", "nand", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [2222, 2360]
sha256: 08abd19080d08310b65f0c9e638bcb72ab072edf7812b374d232e6df0b49d9a8
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

| Terme | Définition |
|---|---|
| AFR | Annualized Failure Rate : % de disques morts par an. |
| Backfill | Recopie des PG sur les OSD survivants après une panne (Ceph). |
| BER | Bit Error Rate : taux d'erreur de lecture (10⁻¹⁴ à 10⁻¹⁷). |
| Block.db / block.wal | Métadonnées RocksDB et journal de BlueStore (Ceph). |
| BlueStore | Backend de stockage Ceph (remplace Filestore, obsolète). |
| Burn-in | Tests intensifs 48-72 h avant mise en production. |
| Cache tiering | Pool SSD devant pool HDD (déprécié dans Ceph). |
| CRUSH | Algorithme de placement pseudo-aléatoire des données (Ceph). |
| CXL | Compute Express Link : mémoire extensible sur PCIe. |
| DataBolt2 | Techno Broadcom : 24G amont avec disques 12G/6G aval. |
| Device class | Classe CRUSH (hdd, ssd, nvme) pour tiering par règles. |
| DWPD | Drive Writes Per Day : endurance normalisée sur 5 ans. |
| EC (Erasure Coding) | k fragments + m parités (Reed-Solomon), surcoût m/(k+m). |
| EDSFF | Enterprise & Datacenter Standard Form Factor (E1.S/E1.L/E3.S/E3.L). |
| EEDP / DIF | Protection des données de bout en bout (blocs 520 B). |
| Expander | « Switch » SAS : 1 port amont, N ports disques. |
| FDP | Flexible Data Placement : l'hôte guide le placement NVMe. |
| HAMR | Heat-Assisted Magnetic Recording : HDD 30 To+ (Seagate). |
| HBA | Host Bus Adapter : expose les disques bruts (mode IT). |
| IR mode | Firmware RAID intégré (MegaRAID) — pas pour Ceph/ZFS. |
| IT mode | Firmware JBOD pur — le mode Ceph/ZFS. |
| JBOD | Châssis de disques sans CPU, piloté par HBA externe. |
| LACP | Agrégation de liens Ethernet (802.3ad). |
| mClock | Ordonnanceur QoS de Ceph (client vs recovery). |
| MTBF | Mean Time Between Failures (ex. 2,5 M h sur Exos). |
| Multipath | Un seul device logique pour plusieurs chemins physiques. |
| NVMe-oF | NVMe over Fabrics : NVMe sur réseau (Ethernet/IB). |
| OSD | Object Storage Daemon : 1 par disque (Ceph). |
| Over-provisioning | NAND cachée au-delà de la capacité annoncée. |
| PG | Placement Group : unité de placement/réplication (Ceph). |
| PLC | Penta-Level Cell (5 bits/cellule) — pas de produit en 2026. |
| PLP | Power Loss Protection : condensateurs anti-coupure (SSD). |
| PUE | Power Usage Effectiveness : énergie totale / énergie IT. |
| QLC | Quad-Level Cell (4 bits/cellule) : capacitaire, ~0,5-1 DWPD. |
| RPO / RTO | Perte de données max / durée de reprise max acceptées. |
| SES | SCSI Enclosure Services : monitoring du JBOD. |
| SFF-8654 | Connecteur SlimSAS 24G (interne). |
| SFF-8674 | Connecteur SAS 24G (externe). |
| SMART | Self-Monitoring, Analysis and Reporting Technology (disques). |
| TBW / PBW | Tera/PetaBytes Written : endurance totale garantie. |
| TLC | Triple-Level Cell (3 bits/cellule) : le standard datacenter. |
| Tri-mode | SAS + SATA + NVMe sur le même contrôleur/baie. |
| UBM | Universal Bay Management : négociation tri-mode (SFF-TA-1005). |
| URE | Unrecoverable Read Error : erreur de lecture irrécupérable. |

---

# PARTIE P — QUIZ

## 235. Quiz : 10 questions

1. Quelle est la différence entre IT mode et IR mode, et lequel faut-il pour Ceph ?
2. Un lien SAS-4 « 24G » négocie en pratique à quel débit effectif par voie ?
3. Pourquoi le RAID 5 est-il déconseillé sur des disques de 20 To ?
4. Calcule le nombre de PG pour 96 OSD en réplica 3 (cible 100 PG/OSD).
5. Quelle taille minimale de block.db pour 40 HDD de 20 To en RGW ?
6. Combien d'OSD HDD au maximum par NVMe de block.db ?
7. Un nœud 60 baies HDD tire ~980 W au mur. Quelle puissance à climatiser avec un PUE de 1,5 ?
8. Pourquoi le spin-down est-il contre-productif sur un cluster Ceph actif ?
9. Quelle est la différence physique entre U.2 et U.3 ?
10. Cite deux raisons de préférer l'erasure coding 8+3 au réplica 3 pour de l'objet froid.

## 236. Quiz : réponses

1. IT = JBOD pur (disques bruts) ; IR = RAID matériel intégré. Pour Ceph/ZFS : **IT toujours** (le logiciel gère redondance et checksums).
2. **22,5 Gb/s** effectifs par voie (24G après encodage — vérifié sur fiches 9600).
3. Risque d'**URE pendant le rebuild** (BER 10⁻¹⁴ sur 20 To lus) → rebuild qui échoue → perte du volume ; + write hole.
4. (100 × 96) / 3 = 3200 → **4096** (puissance de 2) → 128 PG/OSD, dans la fourchette 100-200.
5. 40 × 20 To × 4 % (RGW) = **32 To** de block.db.
6. **≤ 15** OSD HDD par NVMe (4-5 par SSD SATA).
7. 980 W × 1,5 = **~1 470 W** à climatiser.
8. Ceph/ZFS réveillent les disques en permanence (scrub, heartbeat) → cycles incessants = latence + usure mécanique.
9. **Même connecteur physique** ; U.3 ajoute la négociation **tri-mode** (SAS/SATA/NVMe) via UBM côté backplane.
10. Surcoût **37,5 % vs 200 %** (donc ~2× moins de To bruts à acheter) et tolérance de **3 pannes distribuées** sur le cluster.

---

**Fin du guide — vérifications web au 27/09/2026. Prix et disponibilités : allocation tendue, devis datés exigés.**

---

# PARTIE Q — CAS CHIFFRÉS DÉTAILLÉS (RETOURS TERRAIN)

## 237. Cas n°1 : backup 1 Po utile (froid, EC)

Hypothèses : 1 Po utile (1 000 To), backup Veeam/PBS, fenêtre 8 h,
rétention GFS. Protection : EC 8+3 (37,5 %). Disques : Exos M 30 To
(vérifié : ~19 $/To, 27/09/2026).
Brut = 1000 × 1,375 × 1,15 (marge) ≈ 1 581 To → **54 HDD de 30 To**
(1 620 To bruts). Nœuds : 6 nœuds × 9 HDD (k+m = 11 ≤ 6 nœuds ? Non !
EC 8+3 exige au moins 11 domaines : **12 nœuds × 5 HDD** ou 6 nœuds
avec 2 racks/failure-domains — voir piège 240). Prenons **11 nœuds ×
5 HDD** = 55 disques. DB : 55 × 30 To × 2,5 % = 41 To → 6 NVMe 7,68 To
(1 par nœud + 1 spare). Réseau : 2×25G/nœud. Fenêtre backup : 1 Po /
8 h = 35 Go/s — **irréaliste** en 25G : revois la fenêtre (incrémentiel
quotidien ~5 % = 50 To = 1,8 Go/s OK) ou monte en 100G.

## 238. Cas n°1 (suite) : TCO 5 ans

| Poste | Calcul | Montant |
|---|---|---|
| 55 HDD 30 To × ~570 € | 31 350 € | 31 k€ |
| 6 NVMe 7,68 To × ~3 500 € | 21 000 € | 21 k€ |
| 11 nœuds (châssis+CPU+RAM) | 11 × 4 000 € | 44 k€ |
| Réseau 25G (switchs + NIC) | — | ~15 k€ |
| CAPEX total | | **~111 k€** |
| Énergie : 11 × 400 W × 8760 × 1,5 × 0,20 € | | ~11,6 k€/an → 58 k€/5 ans |
| **TCO 5 ans** | | **~170 k€ → ~170 €/To utile** |

Ordres de grandeur (**à vérifier** sur devis). L'énergie = 1/3 du TCO :
le HDD « pas cher » se paie en kWh.

## 239. Cas n°2 : cluster IA training (débit extrême)

Hypothèses : nourrir 8 GPU, besoin 200 Go/s lecture séquentielle,
dataset 500 To. Disques : NVMe Gen5 E3.S 15,36 To (7 Go/s soutenables
chacun en lecture). Disques = 200 / 7 ≈ 29 → **32 NVMe** (marge).
Nœuds : 2 nœuds × 16 NVMe ? Non : blast radius + PCIe. **4 nœuds ×
8 NVMe** (Gen5 x4 = 8×16 = 128 voies PCIe par nœud — vérifie le
bilan PCIe de la carte mère !). Réseau : 200 Go/s = 1,6 Tb/s →
**4× 400G par nœud** ou NVMe-oF direct. Protection : réplica 2
(données reproductibles depuis l'objet froid) ou EC 4+2 local.
Alternative : **WEKA/Lustre** (section 157-158) — à chiffrer contre
Ceph, car à ce débit le tuning Ceph est un projet en soi.

## 240. Piège du cas n°1 : k+m vs nombre de nœuds

EC 8+3 = 11 fragments sur 11 **domaines de panne** distincts. Avec
6 nœuds, CRUSH ne peut pas placer 11 fragments sur 6 hosts sans en
mettre 2 sur le même host → perte d'un host = 2 fragments perdus,
il ne reste que 9 ≥ 8 (k) : ça tient **mathématiquement** mais la
marge fond. Règle : **k+m ≤ nombre de failure domains**, ou accepte
explicitement la dégradation. D'où 11 nœuds au cas n°1 — ou EC 4+2
(6 domains) sur 6 nœuds.

## 241. Cas n°3 : PME, Proxmox + Ceph HCI (3 nœuds)

