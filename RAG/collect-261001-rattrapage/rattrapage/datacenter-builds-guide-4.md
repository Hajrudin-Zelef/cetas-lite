---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-4
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD"]
dates: ["2026-09-27"]
keywords: ["capex", "compute", "ethernet", "nand", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [511, 688]
sha256: c6825392513483465bd22ac28a1c7b0ef78693ccca92aec70e7fa59af5cce5e3
---

# Datacenter Builds — Le guide des BOMs

1. **Acheter des cœurs sans bande passante mémoire** : 160 cœurs avec 384 Go
   de RAM = famine mémoire. Ratio plancher : 4 Go/cœur.
2. **Bi-socket pour « la sécurité »** : en 2026, un mono-socket 96 cœurs bat un
   bi-socket 2× 32 cœurs en perf/prix/watt. Le 2P ne se justifie que par la RAM
   (> 1,5 To) ou les lanes PCIe.
3. **Oublier la licence par socket/cœur** : VMware, certaines DB commerciales
   se paient au cœur — un 192-cœurs peut coûter plus cher en licences qu'en
   matériel. Vérifiez avant de commander.
4. **HPC sur Ethernet 10G « parce qu'on a déjà les switchs »** : 90 % des
   déceptions MPI viennent du réseau, pas du CPU.
5. **Ventilos 1U à 15 000 tr/min dans un bureau** : 70 dB. Le 1U va en salle
   ventilée, jamais sous un bureau.

---

## 26. Check-list d'achat Compute (à joindre au devis)

- [ ] CPU : suffixe P (1P) ou F (fréquence) — justifié par écrit
- [ ] RAM : 12 DIMM identiques, 1 DPC, ECC Registered, fréquence cible
- [ ] Stockage : endurance DWPD adaptée (1 / 3 / 10)
- [ ] NICs : 2 ports min, optiques/DAC chiffrés dans le devis
- [ ] PSU : 2× redondantes, dimensionnées à P_max × 1,25
- [ ] BMC/IPMI : licence incluse (iDRAC Enterprise / iLO Advanced se paient !)
- [ ] Garantie : 3 ans NBD mini, 5 ans pour l'amortissement comptable
- [ ] Firmware : version livrée notée, plan de MAJ au burn-in

---

## 27. Workload 2 — Bases de données : principes

Une DB se dimensionne sur **trois axes** : RAM (working set), IOPS écriture
(endurance), latence (NVMe local, pas de SAN lent).
- **OLTP** (PostgreSQL, MySQL) : fréquence CPU haute (suffixe F), RAM = working
  set + 30 %, NVMe mixed-use 3 DWPD, fsync rapide.
- **Data warehouse / OLAP** : cœurs nombreux, RAM énorme, débit séquentiel.
- Règle d'or : **shared_buffers = 25 % RAM**, le reste au cache OS (Linux fait
  très bien le cache).
- Toujours **2 nœuds mini** (primaire + réplica synchrone) : une DB seule est
  une panne en attente.

---

## 28. PostgreSQL OLTP S — BOM (500 Go actifs, 5 000 TPS)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 1P EPYC, 8 baies NVMe | 2U 1P nu | 1 | 3 800 € |
| EPYC 9375F (32c, 3,85 GHz, 320 W) | fréquence pour OLTP | 1 | 3 893 € |
| RAM 12× 64 Go = 768 Go | DDR5 ECC | 12 | 16 560 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 4× NVMe 3,84 To mixed-use 3 DWPD (data) | PM9A5 ou équiv. | 4 | ≈ 4 200 € × 4 = 16 800 € |
| 2× NVMe 1,92 To mixed-use (WAL dédié) | — | 2 | ≈ 2 400 € × 2 = 4 800 € |
| NIC 2× 25 GbE (réplication) | E810 | 1 | 450 € |
| PSU 2× 1 600 W | incluses | — | — |
| **TOTAL** | | | **≈ 46 665 €** |
| P_max / réaliste | 320+120+150+50+150 | | **≈ 0,79 kW / 0,6 kW** |

Le WAL sur disques dédiés : +20 à 40 % de TPS sur écritures intenses.
C'est le meilleur euro/perf d'une DB.

---

## 29. PostgreSQL OLTP M — BOM (2 To actifs, 20 000 TPS)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 2P EPYC, 12 baies | 2U 2P nu | 1 | 4 200 € |
| EPYC 9475F (48c, 3,65 GHz, 360 W) | 2× | 2 | 4 790 € × 2 = 9 580 € |
| RAM 24× 64 Go = 1,5 To (12/canaux × 2 CPU) | DDR5 ECC | 24 | 33 120 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 8× NVMe 7,68 To mixed-use 3 DWPD (data) | — | 8 | ≈ 5 500 € × 8 = 44 000 € |
| 2× NVMe 3,84 To (WAL) | — | 2 | ≈ 4 200 € × 2 = 8 400 € |
| NIC 2× 100 GbE (réplication + backup) | CX6-DX | 1 | 1 100 € |
| PSU 2× 2 000 W | incluses | — | — |
| **TOTAL** | | | **≈ 100 760 €** |
| P_max / réaliste | 720+240+240+80+150 | | **≈ 1,43 kW / 1,05 kW** |

---

## 30. PostgreSQL Data Warehouse — BOM (20 To, analytique)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 2P EPYC, 24 baies NVMe | type 2U24E nu | 1 | ≈ 3 970 € |
| EPYC 9655P (96c, 400 W) | 2× | 2 | 5 208 € × 2 = 10 416 € |
| RAM 24× 128 Go = 3 To | DDR5 ECC | 24 | ≈ 2 750 € × 24 = 66 000 € |
| 16× NVMe 7,68 To read-intensive (PM9A3) | MZ-QL27T600 | 16 | 3 915 € × 16 = 62 640 € |
| NIC 2× 100 GbE | CX6-DX | 1 | 1 100 € |
| PSU 2× 2 000 W | incluses | — | — |
| **TOTAL** | | | **≈ 144 125 €** |
| P_max / réaliste | 800+240+290+60+150 | | **≈ 1,54 kW / 1,1 kW** |

Capacité brute : 123 To. En production analytique (compression colonne
≈ 3–5×), cela porte **300–500 To logiques**. L'OLAP se paie en RAM, pas en CPU.

---

## 31. MySQL / MariaDB : ce qui change

Même BOM que PostgreSQL S/M, avec 3 différences :
1. **InnoDB buffer pool = 70–80 % RAM** (vs 25 % shared_buffers + cache OS
   pour PG) : prévoyez +15 % de RAM à working set égal.
2. Le WAL s'appelle redo log : mêmes disques dédiés recommandés.
3. Réplication : un réplica asynchrone suffit souvent (vs synchrone en PG
   pour le RPO zéro) — divisez le coût du 2ᵉ nœud par 1,5 si l'async est
   acceptable métier.

---

## 32. Dimensionnement RAM d'une DB — la formule qui évite le sur-achat

```
RAM_cible = (taille_tables_chaudes + taille_index_chauds) × 1,3
```

Méthode :
1. `pg_total_relation_size` trié par accès (`pg_stat_user_tables`) : identifiez
   les 20 % de tables qui servent 80 % des requêtes.
2. Additionnez tables + index de ce top 20 %.
3. × 1,3 (marge + OS + connexions : 10 Mo par connexion, 500 connexions = 5 Go).

Exemple : 300 Go chauds → 390 Go → **384 Go suffisent**, pas 768 Go.
Économie 2026 : 384 Go = 6 modules × 1 380 € ≈ **8 280 € économisés**.
Mesurez d'abord, achetez ensuite.

---

## 33. Stockage DB : l'endurance décide du modèle de SSD

| Profil écriture | DWPD requis | Modèle type | Surcoût vs 1 DWPD |
|---|---|---|---|
| Lecture 90 % (reporting) | 1 | PM9A3 | base |
| OLTP mixte 70/30 | 3 | PM9A5 / Micron 7450 PRO | +40–60 % (à vérifier) |
| WAL / redo seul | 3–10 | Optane-like / Z-NAND (non trouvé au 27/09/2026 en neuf — marché de niche) | — |
| TempDB / tri | 3 | mixed-use | +40–60 % |

Calcul d'endurance : `DWPD_requis = (Go_écrits/jour) / capacité_To / 1000`.
Exemple : 2 To écrits/jour sur 4× 3,84 To = 15,4 To → 0,13 DWPD → 1 DWPD suffit
largement. **La plupart des DB n'ont pas besoin de 3 DWPD** : calculez avant
de payer le surcoût.

---

## 34. Vector DB pour l'IA — en bref (pgvector, Qdrant, Milvus)

Pour le RAG de Zelef et les applis IA : pas besoin d'un build dédié au début.

| Option | Quand | Sizing indicatif |
|---|---|---|
| **pgvector** (extension PG) | < 5 M vecteurs, RAG classique | Même serveur que la DB, +20 % RAM |
| **Qdrant** | 5–100 M vecteurs, filtrage riche | 1 nœud Compute S (section 18) par 20 M vecteurs 768-d |
| **Milvus** | > 100 M vecteurs, distribué | 3 nœuds Compute S minimum |

Règle : **1 M vecteurs 768 dimensions ≈ 3 Go RAM** (float32) / 1,5 Go en
quantifié. 10 M vecteurs = 30 Go → tient dans n'importe quel serveur du guide.
Ne construisez une infra vectorielle dédiée qu'au-delà de 50 M vecteurs.

---

## 35. Cas chiffré : PostgreSQL 2 To actifs, HA synchrone

| Poste | Détail | Total |
|---|---|---|
| 2× PG OLTP M (section 29) | 100 760 € × 2 | 201 520 € |
| 1× témoin/quorum (petit 1U) | — | 8 000 € |
| Switch 25G redondant (déjà au §23 si mutualisé) | — | — |
| Licences | 0 € (open source) | 0 € |
| **CAPEX** | | **≈ 209 500 €** |
| P_max | 2 × 1,43 kW + 0,3 | **≈ 3,2 kW** |
| Coût/To actif utile | 209 500 / 2 | **≈ 105 k€/To** |

C'est le prix de la donnée critique : la redondance synchrone double le coût,
mais le RPO zéro n'a pas de prix en production.

---

## 36. Sauvegarde des DB : ne l'oubliez pas dans la BOM

