---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-6
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["capex", "dram", "gpu", "incident", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [886, 1074]
sha256: 80e87c3e86806df6650942af53f1b0b6af988729a4e4ab51a1b19e4473d02239
---

# Datacenter Builds — Le guide des BOMs

| Critère | U.2 2,5" | E3.S (EDSFF) |
|---|---|---|
| Densité 1U | 10–12 baies | 24–32 baies |
| Hot-swap | oui | oui |
| Prix SSD | référence | ≈ +10–15 % (à vérifier) |
| Dispo châssis 2026 | large | croissante (Supermicro, Dell) |
| Conseil | valeur sûre | si densité 1U requise |

Pour un nouveau build Ceph 2026 : **regardez l'E3.S** si le châssis existe chez
votre intégrateur, sinon restez en U.2 sans regret.

---

## 50. Scrub, deep-scrub et conso cachée

- Le **scrub hebdo** et le **deep-scrub mensuel** lisent tous les disques :
  prévoyez +10–15 % de charge disques/réseau en fenêtre de scrub.
- Planifiez-les **hors heures de backup** : scrub + backup + rebuild
  simultanés = la recette d'un incident.
- Les NVMe consomment 2× plus en écriture qu'au repos : la P_réaliste de la
  section 42/43 suppose 50 % d'écriture — en scrub (lecture), comptez −15 %.

---

## 51. Pièges terrain — Ceph NVMe

1. **Mélanger des SSD de tailles différentes** dans un même pool : Ceph
   répartit par OSD, le plus petit disque limite tout le monde.
2. **Back réseau en 10G « pour commencer »** : le premier rebuild sature tout
   et les clients rament. 25G mini, 100G dès 8 nœuds.
3. **Oublier les MON dédiés en NVMe** : des MON sur SATA = latence de cluster.
4. **Remplir à 90 %+** « temporairement » : il n'y a rien de plus permanent que
   du temporaire en stockage.
5. **Firmware SSD hétérogènes** : un lot avec un vieux firmware peut faire du
   garbage-collection au mauvais moment. Uniformisez au burn-in.
6. **Sous-dimensionner la RAM** : BlueStore + RocksDB mangent 5 Go/To. En
   dessous, c'est le swap et la mort lente.

---

## 52. Ce qu'il faut retenir — Ceph NVMe

- 5 nœuds OSD mini en prod, réplica 3, 85 % max de remplissage.
- Back réseau = 2× le débit de rebuild ; 100G dès que ça compte.
- Coût 2026 : ≈ 3 000–3 600 €/To utile (la DRAM et les SSD dominent).
- 25 kW pour 500 To utiles : prévoyez le refroidissement dès la BOM.

---

## 53. Workload 4 — Ceph hybride : principes

Le Ceph hybride combine **HDD 3,5" haute capacité** (données froides/tièdes) et
**NVMe** (cache, DB/WAL BlueStore, pools chauds). C'est l'architecture
prix/capacité pour l'objet, la vidéo, le backup secondaire.
- Ratio classique : **1 To de NVMe pour 10–20 To de HDD**.
- Les HDD sont en 2026 des **20–24 To CMR** (Seagate Exos / WD Ultrastar).
- Comptez **120–180 IOPS par HDD** : le dimensionnement IOPS limite le nombre
  de HDD par nœud avant la capacité.

---

## 54. Nœud OSD hybride S — BOM (12× 20 To HDD + 4× 3,84 To NVMe)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 4U 36 baies 3,5", 1P EPYC | type 4U36 nu | 1 | 4 800 € |
| EPYC 9355P (32c) | — | 1 | 2 891 € |
| RAM 12× 32 Go = 384 Go | DDR5 ECC | 12 | 9 000 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 12× HDD 20 To CMR (Exos X20) | — | 12 | ≈ 380 € × 12 = 4 560 € |
| 4× NVMe 3,84 To (cache/WAL) | PM9A3 | 4 | 3 035 € × 4 = 12 140 € |
| NIC 2× 25 GbE (front) | E810 | 1 | 450 € |
| NIC 2× 25 GbE (back) | E810 | 1 | 450 € |
| PSU 2× 1 600 W | incluses | — | — |
| **TOTAL** | | | **≈ 34 650 €** |
| P_max / réaliste | 280+110+60+90+50+150 | | **≈ 0,74 kW / 0,55 kW** |

Capacité brute : 240 To HDD + 15 To NVMe. Les HDD consomment peu (≈ 8–10 W
chacun en charge) mais **vibrent** : ne mélangez jamais HDD et NVMe sans
isolation mécanique correcte (le châssis 4U est conçu pour).

---

## 55. Nœud OSD hybride M — BOM (24× 20 To HDD + 6× 7,68 To NVMe)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 4U 36 baies, 2P EPYC | 4U36 2P nu | 1 | 5 500 € |
| EPYC 9555P (64c) | 2× | 2 | 4 830 € × 2 = 9 660 € |
| RAM 24× 32 Go = 768 Go | DDR5 ECC | 24 | 18 000 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 24× HDD 20 To CMR | Exos X20 | 24 | 380 € × 24 = 9 120 € |
| 6× NVMe 7,68 To (cache/WAL) | PM9A3 | 6 | 3 915 € × 6 = 23 490 € |
| NIC 2× 25 GbE (front) | E810 | 1 | 450 € |
| NIC 2× 100 GbE (back) | CX6-DX | 1 | 1 100 € |
| PSU 2× 2 000 W | incluses | — | — |
| **TOTAL** | | | **≈ 67 680 €** |
| P_max / réaliste | 720+240+190+140+85+150 | | **≈ 1,53 kW / 1,1 kW** |

Capacité brute : 480 To HDD + 46 To NVMe par nœud.

---

## 56. Ratios de cache : combien de NVMe pour combien de HDD

| Workload | Ratio NVMe:HDD | Exemple (nœud M) |
|---|---|---|
| Objet S3 froid (backup, archives) | 1:20 | 24 To NVMe / 480 To HDD ✓ |
| Objet tiède (VM images) | 1:10 | 48 To NVMe / 480 To HDD |
| RBD bloc actif | 1:5 ou full NVMe | — (prenez du full NVMe, §41) |

Le NVMe sert à : BlueStore DB/WAL (10–30 Go par OSD HDD), cache Ceph
(read/write), et pools « chauds » séparés. **Ne mettez jamais le WAL BlueStore
sur le HDD lui-même** : c'est −50 % de perfs d'écriture.

---

## 57. Cas chiffré : 2 Po utiles hybrides, EC 4+2

Besoin brut : 2 000 / 0,85 / (4/6) = **3 530 To bruts** → 8 nœuds M
(8 × 480 = 3 840 To HDD).

| Poste | Détail | Total |
|---|---|---|
| 8× nœuds hybrides M | 67 680 € × 8 | 541 440 € |
| 3× MON/MGR (section 44) | 40 455 € | 40 455 € |
| 2× switch 25G/100G | — | 25 000 € |
| Câblage | lot | 8 000 € |
| **CAPEX** | | **≈ 614 900 €** |
| P_max | 8 × 1,53 + 3 × 0,45 + 1 | **≈ 14,6 kW** |
| **Coût/To utile** | 614 900 / 2 000 | **≈ 307 €/To** |

**307 €/To vs 3 090 €/To en full NVMe : un facteur 10.** C'est pour ça que
l'hybride existe. Réservez le NVMe à ce qui a besoin de latence.

---

## 58. Comparatif coût NVMe vs hybride (€/To utile, 2026)

| Architecture | €/To utile | Latence typique | Usage |
|---|---|---|---|
| Full NVMe réplica 3 | ≈ 3 100 € | < 1 ms | RBD, DB, VM |
| Full NVMe EC 4+2 | ≈ 1 600 € (à vérifier) | 1–3 ms | Objet perf |
| Hybride EC 4+2 | ≈ 310 € | 5–20 ms | Objet froid, backup |
| HDD seul EC 8+3 (à vérifier) | ≈ 180 € | 10–30 ms | Archives |

---

## 59. Pièges terrain — Ceph hybride

1. **HDD SMR « pas cher »** : le SMR (shingled) est inutilisable en Ceph
   (réécritures catastrophiques). **CMR uniquement** — vérifiez la fiche.
2. **Rebuild HDD de 20 To** : à 200 Mo/s = **28 heures** par disque. Pendant ce
   temps, un 2ᵉ disque qui lâche = perte de données en réplica 3 dégradé.
   Prévoyez des spares et surveillez le SMART.
3. **Vibrations** : 24 HDD qui vibrent dans un rack mal fixé tuent les perfs
   de tous les disques. Serrez les rails, évitez les racks bas de gamme.
4. **Température HDD** : 25–40 °C idéal, 50 °C = vieillissement accéléré.
   Les 4U 36 baies ont besoin d'un flux d'air sérieux en façade.
5. **Mélanger CMR et NVMe dans le même CRUSH bucket** sans device class :
   Ceph répartit au hasard → perfs imprévisibles. **Device classes
   `hdd`/`ssd` obligatoires.**

---

## 60. Ce qu'il faut retenir — Ceph hybride

- 1 To NVMe pour 10–20 To HDD, WAL BlueStore sur NVMe, CMR uniquement.
- ≈ 310 €/To utile en EC 4+2 : 10× moins cher que le full NVMe.
- Le rebuild d'un 20 To prend 28 h : spares + surveillance SMART.
- Device classes Ceph : séparez `hdd` et `ssd` dès le jour 1.

---

## 61. Workload 5 — VDI : principes

La VDI se dimensionne **par utilisateur**, pas par serveur :
- **CPU** : 2–4 vCPU par utilisateur (bureautique 2, CAO 4–8).
- **RAM** : 4–8 Go par utilisateur (bureautique 4, CAO 16).
- **IOPS** : 20–50 par utilisateur, **100+ au boot storm** (8h55, tout le monde
  allume son poste en même temps).
- **GPU** : vGPU NVIDIA (vPC pour bureau, vWS pour CAO) quand l'appli l'exige.

Deux builds : **S** (200 users bureautiques) et **M** (400 users dont CAO avec vGPU).

---

## 62. Tableau des profils utilisateurs VDI

