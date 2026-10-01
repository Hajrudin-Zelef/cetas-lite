---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-5
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["benchmark", "capex"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [689, 885]
sha256: e5e854d892d3f285f604d5bb354830a4271a65cf74ce2372f56b6a6cddaabf38
---

# Datacenter Builds — Le guide des BOMs

- **pgBackRest / Percona XtraBackup** vers le serveur backup (section 97+) :
  full hebdo + incr quotidienne + WAL archivés en continu.
- Dimensionnement : taille_base × (1 + 6 × taux_incr + 1) × rétention.
  Exemple : base 2 To, incr 5 %/j, rétention 30 j → 2 × (1 + 0,3 + 1) ≈ 4,6 To
  sur le repo backup. **Chiffrez toujours le repo** (le backup contient toute
  la donnée en clair sinon).
- Restaurez **tous les mois** en test : un backup non testé n'existe pas.

---

## 37. Monitoring DB : les 5 métriques qui comptent

1. **Cache hit ratio** (> 99 % en OLTP, sinon RAM insuffisante).
2. **Checkpoints** : fréquence et durée (WAL/disques trop lents si > 30 s).
3. **Bloat** : tables/index gonflés → VACUUM/OPTIMIZE planifié.
4. **Réplication lag** : < 1 s en synchrone, alerte au-delà.
5. **Température SSD** : les NVMe throttlent à 70–75 °C — vérifiez le flux
   d'air du châssis (les baies avant des 2U denses chauffent).

---

## 38. Pièges terrain — Bases de données

1. **Mettre la DB sur Ceph sans QoS** : la latence réseau tue les TPS.
   DB critique = NVMe local, toujours.
2. **RAID 5/6 sur SSD pour une DB** : la pénalité d'écriture ×4 + le garbage
   collection = latence imprévisible. Miroir ou RAID 10.
3. **Oublier le WAL dédié** : le meilleur euro/perf, et le plus oublié.
4. **Sur-RAM « au cas où »** en 2026 : 384 Go inutiles = 8 280 € brûlés.
5. **Tester la perf sur des données vides** : un benchmark sur table vide ne
   vaut rien. Testez à volumétrie réelle.
6. **Réplica synchrone sur un lien > 2 ms** : chaque commit attend le réplica —
   la latence applicative explose. Synchrone = même salle, fibre courte.

---

## 39. Tableau récapitulatif — Builds DB

| Build | Cible | Prix | P_max | RAM | Stockage |
|---|---|---|---|---|---|
| PG OLTP S | 500 Go / 5 kTPS | ≈ 46,7 k€ | 0,79 kW | 768 Go | 15,4 To MU + WAL |
| PG OLTP M | 2 To / 20 kTPS | ≈ 100,8 k€ | 1,43 kW | 1,5 To | 61 To MU + WAL |
| PG Warehouse | 20 To analytique | ≈ 144,1 k€ | 1,54 kW | 3 To | 123 To RI |
| HA synchrone (2× M) | RPO zéro | ≈ 209,5 k€ | 3,2 kW | — | — |

MU = mixed-use 3 DWPD, RI = read-intensive 1 DWPD.

---

## 40. Ce qu'il faut retenir — DB

- CPU **F** (fréquence), RAM = working set × 1,3, WAL dédié, NVMe local.
- L'endurance se **calcule** (Go écrits/jour) — ne payez pas 3 DWPD par défaut.
- En 2026, la RAM domine le prix : mesurez avant d'acheter.
- Vector DB : commencez par pgvector, Qdrant au-delà de 5 M vecteurs.

---

## 41. Workload 3 — Ceph full NVMe : principes

Ceph full NVMe = le stockage unifié (bloc RBD, objet S3, fichier CephFS) à
haute performance. Règles d'architecture :
- **Minimum 3 nœuds OSD** (5+ recommandé pour la résilience), + 3 MON/MGR
  (colocalisés ou dédiés).
- **Réplica 3** par défaut (survie à 1 panne) ; **EC 4+2** pour l'objet froid
  (économe en capacité, pas pour le bloc à latence faible).
- **Réseau séparé** : front (clients) et back (réplication/rebuild) — le back
  doit encaisser 2× le débit d'un rebuild complet.
- 1 OSD par SSD (pas de partitionnement), BlueStore.

---

## 42. Nœud OSD NVMe S — BOM (8× 7,68 To = 61 To bruts)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 12 baies NVMe, 1P EPYC | 2U 12× U.2 nu | 1 | 4 200 € |
| EPYC 9355P (32c) — Ceph aime les cœurs | — | 1 | 2 891 € |
| RAM 12× 32 Go = 384 Go (5 Go/To × 61 To + base) | DDR5 ECC | 12 | 9 000 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 8× NVMe 7,68 To PM9A3 (OSD) | MZ-QL27T600 | 8 | 3 915 € × 8 = 31 320 € |
| NIC 2× 25 GbE (front) | E810 | 1 | 450 € |
| NIC 2× 100 GbE (back/rebuild) | CX6-DX | 1 | 1 100 € |
| PSU 2× 1 600 W | incluses | — | — |
| **TOTAL** | | | **≈ 49 320 €** |
| P_max / réaliste | 280+110+150+85+150 | | **≈ 0,78 kW / 0,6 kW** |

---

## 43. Nœud OSD NVMe M — BOM (16× 7,68 To = 123 To bruts)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 24 baies NVMe, 2P EPYC | type 2U24E nu | 1 | ≈ 3 970 € |
| EPYC 9555P (64c) | 2× | 2 | 4 830 € × 2 = 9 660 € |
| RAM 24× 32 Go = 768 Go | DDR5 ECC | 24 | 18 000 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 16× NVMe 7,68 To PM9A3 (OSD) | MZ-QL27T600 | 16 | 3 915 € × 16 = 62 640 € |
| NIC 2× 25 GbE (front) | E810 | 1 | 450 € |
| NIC 2× 100 GbE (back) | CX6-DX | 1 | 1 100 € |
| PSU 2× 2 000 W | incluses | — | — |
| **TOTAL** | | | **≈ 96 180 €** |
| P_max / réaliste | 720+240+290+85+150 | | **≈ 1,49 kW / 1,1 kW** |

---

## 44. Nœuds MON/MGR — BOM (×3)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 1U 1P EPYC, 4 baies | AS-1115CS-TNR nu | 1 | 4 585 € |
| EPYC 9124 (16c, 200 W) | — | 1 | 749 € |
| RAM 6× 32 Go = 192 Go | DDR5 ECC | 6 | 4 500 € |
| 2× NVMe 1,92 To (RocksDB MON) | — | 2 | ≈ 1 600 € × 2 = 3 200 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| **TOTAL / nœud** | | | **≈ 13 485 €** |
| **× 3 nœuds** | | | **≈ 40 455 €** |
| P_max / nœud | | | **≈ 0,45 kW** |

Les MON stockent RocksDB : mettez du NVMe correct, pas du SATA. Un MON lent =
cluster lent (les OSD attendent les maps).

---

## 45. Réseau Ceph : le dimensionner sans se tromper

| Lien | Débit recommandé | Pourquoi |
|---|---|---|
| Front (clients → OSD) | 2× 25 GbE par nœud | pics clients |
| Back (réplication) | 2× 100 GbE par nœud (M) / 2× 25 GbE (S) | rebuild × N OSD |
| MON | 2× 10/25 GbE | faible |

Règle : **bande passante back ≥ (débit_rebuild_1_OSD × nb_OSD_en_rebuild)**.
Avec 16 OSD à 500 Mo/s = 8 Go/s = 64 Gb/s → 2× 100 GbE = 200 Gb/s, confortable.
En 25 GbE × 2 = 50 Gb/s : le rebuild sature le back — acceptable en S, pas en M.

Switches : 2× ToR 100G 32 ports (redondance MLAG) ≈ 12 000–18 000 € pièce
(à vérifier). Comptez aussi **un réseau out-of-band** pour les MON.

---

## 46. Formule capacité utile Ceph

```
Utile_réplica3 = brut × 0,85 (marge remplissage) / 3
Utile_EC4+2    = brut × 0,85 × 4/6
```

Exemples :
- 5× nœuds S (5 × 61 To = 307 To bruts) en réplica 3 → **87 To utiles**.
- 5× nœuds M (5 × 123 To = 615 To bruts) en réplica 3 → **174 To utiles**.
- Mêmes 615 To en EC 4+2 (pool objet) → **348 To utiles**.

Ne remplissez jamais un pool au-delà de **85 %** : au-delà, Ceph ne peut plus
rééquilibrer et les perfs s'effondrent. C'est une limite dure, pas un conseil.

---

## 47. Cas chiffré : 500 To utiles full NVMe, réplica 3

Besoin brut : 500 / 0,85 × 3 = **1 765 To bruts** → 15 nœuds M (15 × 123 To).

| Poste | Détail | Total |
|---|---|---|
| 15× nœuds OSD M | 96 180 € × 15 | 1 442 700 € |
| 3× MON/MGR | 40 455 € | 40 455 € |
| 2× switch 100G 32p (back) | 15 000 € × 2 | 30 000 € |
| 2× switch 25G 48p (front) | 8 000 € × 2 | 16 000 € |
| Câblage/optiques | lot | 15 000 € |
| **CAPEX** | | **≈ 1 544 000 €** |
| P_max | 15 × 1,49 + 3 × 0,45 + 4 × 0,4 | **≈ 25 kW** |
| **Coût/To utile** | 1 544 000 / 500 | **≈ 3 090 €/To** |

25 kW = **un rack haute densité** (air limite) ou 2 racks standard.
C'est ici que le refroidissement (partie D) entre en jeu.

---

## 48. Cas chiffré : 150 To utiles, 5 nœuds S

| Poste | Détail | Total |
|---|---|---|
| 5× nœuds OSD S | 49 320 € × 5 | 246 600 € |
| 3× MON/MGR (mutualisables si petit) | 40 455 € | 40 455 € |
| 2× switch 25G/100G | — | 20 000 € |
| Câblage | lot | 5 000 € |
| **CAPEX** | | **≈ 312 000 €** |
| P_max | 5 × 0,78 + 3 × 0,45 + 0,8 | **≈ 6 kW** |
| **Coût/To utile** (87 To, §46) | 312 000 / 87 | **≈ 3 590 €/To** |

Le petit cluster coûte **plus cher au To** (MON fixes, switches fixes) :
l'effet d'échelle joue dès 8–10 nœuds.

---

## 49. E3.S vs U.2 : le format de 2026

