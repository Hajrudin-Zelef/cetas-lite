---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-2
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Broadcom", "Intel", "Nvidia", "Samsung", "United States"]
dates: ["2026-09-27"]
keywords: ["amd", "blackwell", "capex", "compute", "dram", "ethernet", "gpu", "hbm3", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [149, 331]
sha256: de067ef82919f26a4359aafba944d7206362f85a510facca0a1a61fb294c3304
---

# Datacenter Builds — Le guide des BOMs

| SSD | Capacité | Endurance | Prix constaté |
|---|---|---|---|
| Samsung PM9A3 U.2 | 960 Go | 1 DWPD | ≈ 877 $ |
| Samsung PM9A3 U.2 | 3,84 To | 1,3 DWPD | ≈ 3 300 $ |
| Samsung PM9A3 U.2 | 7,68 To | 1 DWPD | ≈ 4 256 $ |
| Samsung PM9A3 U.2 | 15,36 To | 1 DWPD | ≈ 2 900–3 000 $ (marché gris, à vérifier) |

Choix d'endurance (critique, voir sections 34 et 48) :
- **Read-intensive (1 DWPD)** : Ceph, backup, boot. Le PM9A3 convient.
- **Mixed-use (3 DWPD)** : DB, virtualisation dense. Gamme PM9A5 / équivalent
  Micron 7450 PRO — ≈ +40 à 60 % vs read-intensive (à vérifier au devis).
- **Write-intensive (10+ DWPD)** : ZIL/SLOG, journaux. Petit volume, cher.

Formats :
- **U.2 2,5"** : standard, hot-swap, le choix par défaut.
- **E3.S / EDSFF** : densité supérieure (1U 24+ baies), remplace progressivement
  l'U.2 sur les plateformes 2025–2026. Vérifiez la dispo du châssis.
- **M.2** : boot uniquement (2× en RAID 1), jamais pour la donnée en prod.

---

## 7. NICs : 25/100/400 GbE, OCP 3.0

| NIC | Usage | Prix indicatif |
|---|---|---|
| 2× 25 GbE SFP28 (Intel E810 / Broadcom) | Compute, Ceph front | ≈ 350–500 € (à vérifier) |
| 2× 100 GbE QSFP28 (Mellanox CX6-DX) | Ceph backend, HPC | ≈ 900–1 300 € (à vérifier) |
| 2× 400 GbE QSFP112 (CX7) | Inter-GPU, backbone IA | ≈ 2 500–4 000 € (non trouvé au 27/09/2026 — devis requis) |
| OCP 3.0 2× 25 GbE | Standard serveurs récents | inclus ou ≈ 250 € |

Règles :
- **RoCE v2** (RDMA sur Ethernet) suffit pour Ceph et le stockage ; **InfiniBand**
  reste roi pour le HPC/GPU à grande échelle (section 21).
- Toujours **2 NICs ou 1 dual-port** : séparez front/back sur Ceph, management
  sur port dédié BMC.
- Comptez les **optiques et DAC** dans la BOM : 30–80 € le DAC 25G 3 m,
  150–400 € l'optique 100G SR4 (à vérifier). Sur 40 serveurs, c'est un budget
  à part entière.

---

## 8. GPU : panorama d'achat 2026

| GPU | VRAM | TDP | Prix constaté (mars–sept 2026) |
|---|---|---|---|
| NVIDIA L40S | 48 Go GDDR6 | 350 W | 8 610–8 900 $ (MSRP 8 352 $) |
| NVIDIA RTX PRO 6000 Blackwell | 96 Go GDDR7 | 600 W | 9 450–9 800 $ (MSRP 7 673 $ workstation) |
| NVIDIA RTX 6000 Ada | 48 Go GDDR6 | 300 W | 7 400–7 800 $ |
| NVIDIA H100 SXM 80 Go | 80 Go HBM2e | 700 W | 25 000–35 000 $ neuf (18 000–22 000 $ occasion) |
| NVIDIA H200 141 Go | 141 Go HBM3e | 700 W | 30 000–40 000 $ |
| NVIDIA B200 | 192 Go HBM3e | 1 000 W | 30 000–40 000 $ |
| AMD MI300X | 192 Go HBM3 | 750 W | 10 000–12 000 $ (standalone) |
| AMD MI325X | 256 Go HBM3e | 750 W | non trouvé au 27/09/2026 (prix public) |
| AMD MI355X | 288 Go HBM3e | 1 000 W | non trouvé au 27/09/2026 (prix public) |
| NVIDIA L4 | 24 Go GDDR6 | 72 W | 2 500–8 000 $ |

Tous prix vérifiés le 27/09/2026 sauf mentions. Le marché est en allocation :
**ne promettez jamais un délai GPU sans écrit du distributeur.**

---

## 9. Alimentation serveurs : Titanium, redondance, 200–240 V

- Standard : **2× PSU redondantes 80 PLUS Titanium** (96 %+ de rendement à 50 %).
- Puissances courantes : 1 600 W (1U/2U compute), 2 000–2 600 W (2U 4-GPU),
  3 000 W+ (nœuds 8-GPU HGX).
- **Branchez toujours en 220–240 V** : un PSU 2 600 W ne sort que 1 000–1 200 W
  en 110 V. En France c'est du 230 V mono ou 400 V tri — pas de sujet, mais
  vérifiez les PDU US si vous importez du matériel.
- Dimensionnement PSU : P_max_serveur × 1,25, réparti sur 2 PSU (N+1).
  Exemple : serveur 1 400 W → 2× 1 600 W (chaque PSU doit pouvoir porter seule
  la charge en cas de panne de l'autre… en pratique on tolère N+1 à 80 %).

---

## 10. Formule n° 1 — IOPS : le dimensionnement stockage

```
IOPS_requis = (IOPS_lecture × %_lecture) + (IOPS_écriture × %_écriture × pénalité_RAID)
```

Pénalités d'écriture : miroir = 2, RAID 5/6 ≈ 4, Ceph réplica 3 ≈ 3 (chaque
écriture client = 3 écritures disques + overhead réseau).

Exemple terrain — 200 VMs bureautiques :
- 50 IOPS/VM en pic (boot storm : 100) → 200 × 50 = 10 000 IOPS mixtes 70/30.
- Écritures : 3 000 × 3 (Ceph replica 3) = 9 000 IOPS disques.
- Total disques : 7 000 + 9 000 = 16 000 IOPS.
- Un NVMe PM9A3 fait ≈ 900 k IOPS lecture / 200 k écriture : **2 NVMe suffisent
  en IOPS**, mais la capacité dictera le nombre réel (section 12).

**Leçon : en 2026, avec le NVMe, c'est presque toujours la capacité qui
dimensionne, pas les IOPS.** Sauf DB OLTP extrême et VDI en boot storm.

---

## 11. Formule n° 2 — Débit réseau et disques

```
Débit_requis = volume / fenêtre_temps × marge
```

Exemple — backup de 20 To en 8 h :
- 20 To = 160 Tb ; / 28 800 s = 5,5 Gb/s théoriques.
- × 1,5 (marge, overhead, dédup) ≈ **8–10 Gb/s** → NIC 25 GbE recommandée.

Exemple — rebuild Ceph : un OSD de 7,68 To qui rebuild à 500 Mo/s met
≈ 4 h 20. Pendant ce temps le cluster est dégradé : prévoyez la bande
passante backend en conséquence (2× 25 G mini, 2× 100 G confortable).

Tableau de conversion rapide :

| Débit | To/heure |
|---|---|
| 1 Gb/s | 0,45 To/h |
| 10 Gb/s | 4,5 To/h |
| 25 Gb/s | 11 To/h |
| 100 Gb/s | 45 To/h |

---

## 12. Formule n° 3 — RAM : le working set + 30 %

```
RAM = working_set_chaud × 1,3 + OS/hyperviseur + marge_croissance
```

- **DB** : working set = taille des tables+index fréquemment lus. 80 % des
  requêtes touchent 20 % des données : dimensionnez pour ces 20 %.
- **Virtualisation** : somme des RAM VMs × (1 − taux_déduplication_KSM ≈ 0,85)
  + 16 Go hyperviseur.
- **Ceph OSD** : 4 Go de base + 1 Go/To (BlueStore), arrondi à 5 Go/To.
- **VDI** : 4–8 Go par utilisateur selon profil (section 58).

Avec la DRAM à 1 500 $/64 Go, **chaque Go inutile coûte ≈ 23 $**. Mesurez le
working set réel (pg_stat, monitoring) avant d'acheter.

---

## 13. Formule n° 4 — Watts : la plus importante du guide

```
P_serveur ≈ Σ TDP_CPU + Σ TDP_GPU + P_RAM + P_disques + P_NICs + P_base
```

Avec :
- P_RAM ≈ 8–10 W par DIMM DDR5 (12 DIMM ≈ 110 W).
- P_disque NVMe ≈ 12–20 W en charge (× nombre).
- P_base (carte mère, BMC, ventilos) ≈ 80–150 W selon châssis.
- **Facteur de charge réel** : un serveur ne tourne jamais à 100 % de tous ses
  TDP simultanément. Appliquez un **facteur de coïncidence de 0,7–0,8** pour
  l'électrique, mais dimensionnez les PSU et le refroidissement à 1,0.

Exemple — serveur 2× EPYC 9655P + 12 NVMe + 2× 100G :
- 2× 400 + 12× 15 + 2× 25 + 110 + 120 = **1 260 W** max théorique.
- Charge réaliste : ≈ 900 W. PSU : 2× 1 600 W. PDU : compter 1,3 kW.

**C'est cette formule qui alimente tout le dimensionnement onduleur
(partie D, sections 155–160).** Faites le tableau une fois, réutilisez-le partout.

---

## 14. Formule n° 5 — TCO simplifié sur 5 ans

```
TCO_5ans = CAPEX_matériel + 5 × (électricité + maintenance + licences)
```

- Électricité : P_moyenne_kW × 8 760 h × prix_kWh (0,15–0,25 €/kWh pro en France
  en 2026 — à vérifier sur votre contrat).
- Exemple : serveur à 900 W moyens, 0,20 €/kWh → 900 × 8,76 × 0,20 ≈
  **1 577 €/an**, soit **7 885 € sur 5 ans**. À comparer au CAPEX : sur un
  serveur à 15 000 €, l'énergie représente **un tiers du TCO**.
- Maintenance : 8–12 % du CAPEX/an en contrat constructeur NBD, ou pièces
  détachées en stock pour l'autonomie.

**Un serveur « pas cher » qui consomme 400 W de plus coûte 3 500 € de plus
en électricité sur 5 ans.** Le TCO tue les fausses économies.

---

## 15. Règles de marge terrain (à appliquer à toutes les BOMs)

