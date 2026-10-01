---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-3
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["amd", "capex", "compute", "dram", "gpu", "intel"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [332, 510]
sha256: 791550f4b45c0582a24b5347b71ab2654ad856e6f8c07f7df3a9a0eb32a9f615
---

# Datacenter Builds — Le guide des BOMs

1. **Capacité** : +30 % de croissance à 3 ans minimum sur stockage et RAM.
2. **Puissance** : +25 % sur PSU et PDU vs P_max calculée.
3. **Réseau** : un uplink ne dépasse jamais 70 % en nominal.
4. **IOPS** : +50 % vs pic mesuré (les pics non mesurés existent).
5. **Délais** : commandez GPU et DRAM 3–6 mois avant le besoin en 2026.
6. **Pièces** : 1 disque et 1 PSU de rechange par grappe homogène de 10 nœuds.

---

## 16. Châssis de référence utilisés dans ce guide

| Châssis | Format | Usage guide | Prix nu constaté |
|---|---|---|---|
| ASUS RS501A-E12 / Supermicro AS-1115CS-TNR | 1U 1P EPYC | Compute, firewall, Ceph MON | ≈ 4 500–5 000 € (vérifié le 27/09/2026) |
| 2U 8–12 baies (type Supermicro 2029P) | 2U 2P EPYC | DB, virtualisation, backup | ≈ 3 500–4 500 € (à vérifier) |
| 2U 24 baies NVMe (type ASRock Rack 2U24E) | 2U 2P EPYC | Ceph NVMe | ≈ 4 300 $ (relevé ancien — à vérifier) |
| ASUS ESC4000A-E12 | 2U 4× GPU double-slot | Inférence 2–4 GPU | ≈ 4 320 £ (~5 000 €) nu (vérifié le 27/09/2026) |
| 4U 36 baies 3,5" (type Supermicro 847) | 4U | Ceph hybride, backup | ≈ 4 000–5 500 € (à vérifier) |
| 8-GPU HGX (Supermicro/Dell XE) | 8U | Inférence lourde | devis uniquement (non trouvé au 27/09/2026 en prix public) |

> Les BOMs ci-dessous utilisent ces châssis comme base « nu » (sans CPU/RAM/disques).

---

# PARTIE B — BOMs PAR WORKLOAD

---

## 17. Workload 1 — Compute / HPC : principes

Le « compute » générique (virtualisation, Kubernetes, batch) se dimensionne au
**ratio cœurs/RAM/bande passante** :
- Virtualisation/K8s : 4–8 Go de RAM par cœur.
- HPC : 2–4 Go par cœur + interconnect rapide (InfiniBand ou RoCE).
- 1× EPYC 64 cœurs + 512 Go–1 To RAM couvre 80 % des besoins PME.

Deux tailles ci-dessous : **S** (mono-socket 32 cœurs, le cheval de labour) et
**M** (bi-socket 96 cœurs, nœud HPC/virt dense).

---

## 18. Compute S — BOM complète (1× EPYC 32 cœurs, 384 Go RAM)

| Composant | Référence type | Qté | Prix unit. | Total |
|---|---|---|---|---|
| Châssis 1U 1P EPYC 9005, 10 baies NVMe | Supermicro AS-1115CS-TNR nu | 1 | 4 585 € | 4 585 € |
| CPU AMD EPYC 9355P (32c/64t, 280 W) | 100-000001694 (P = 1P) | 1 | 2 891 € | 2 891 € |
| RAM DDR5-5600 ECC RDIMM 32 Go | 12× 32 Go = 384 Go | 12 | ≈ 750 € | 9 000 € |
| SSD boot M.2 NVMe 960 Go (RAID 1) | 2× 960 Go entreprise | 2 | 180 € | 360 € |
| SSD data U.2 NVMe 3,84 To (PM9A3) | MZQL23T80 | 4 | ≈ 3 035 € | 12 140 € |
| NIC 2× 25 GbE SFP28 | Intel E810-XXVDA2 | 1 | 450 € | 450 € |
| DAC 25G 3 m + optiques | — | 4 | 50 € | 200 € |
| PSU 2× 1 600 W Titanium redondantes | incluses châssis | — | — | — |
| Rails + kit | — | 1 | 120 € | 120 € |
| **TOTAL** | | | | **≈ 29 750 €** |

Prix vérifiés le 27/09/2026 sauf RAM/SSD/NIC (ordres de grandeur 2026 — à vérifier
au devis, la DRAM bouge vite).

---

## 19. Compute S — consommation électrique estimée

| Poste | Calcul | Watts |
|---|---|---|
| CPU EPYC 9355P | 1× 280 W TDP | 280 |
| RAM 12 DIMM | 12 × 9 W | 108 |
| 4× NVMe data + 2× M.2 | 6 × 15 W | 90 |
| NIC 25G | 1 × 25 W | 25 |
| Base (carte mère, BMC, ventilos 1U) | — | 120 |
| **P_max théorique** | | **≈ 623 W** |
| **P_réaliste (facteur 0,75)** | | **≈ 470 W** |

PSU : 2× 1 600 W → largement couvertes. **Compter 0,65 kW par nœud** pour la PDU.
Sur 10 nœuds : 6,5 kW/rack — de l'air cooling standard.

---

## 20. Compute M — BOM complète (2× EPYC 96 cœurs, 768 Go RAM)

| Composant | Référence type | Qté | Prix unit. | Total |
|---|---|---|---|---|
| Châssis 2U 2P EPYC, 12 baies | 2U 12× 3,5"/2,5" nu | 1 | 4 000 € | 4 000 € |
| CPU AMD EPYC 9655P (96c/192t, 400 W) | 2× | 2 | 5 208 € | 10 416 € |
| RAM DDR5-5600 ECC RDIMM 64 Go | 12× 64 Go = 768 Go | 12 | ≈ 1 380 € | 16 560 € |
| SSD boot 2× M.2 960 Go RAID 1 | — | 2 | 180 € | 360 € |
| SSD data U.2 NVMe 7,68 To (PM9A3) | MZ-QL27T600 | 6 | ≈ 3 915 € | 23 490 € |
| NIC 2× 25 GbE SFP28 (front) | Intel E810-XXVDA2 | 1 | 450 € | 450 € |
| NIC 2× 100 GbE QSFP28 (stockage/HPC) | Mellanox CX6-DX | 1 | 1 100 € | 1 100 € |
| PSU 2× 2 000 W Titanium | incluses | — | — | — |
| Câbles/optiques | — | lot | — | 600 € |
| **TOTAL** | | | | **≈ 56 975 €** |

Note : la RAM (16 560 €) coûte plus cher que les 2 CPU réunis. Bienvenue en 2026.

---

## 21. Compute M — consommation et comparatif S/M

| Poste | Calcul | Watts |
|---|---|---|
| 2× CPU EPYC 9655P | 2 × 400 W | 800 |
| RAM 12 DIMM 64 Go | 12 × 10 W | 120 |
| 6× NVMe 7,68 To | 6 × 18 W | 108 |
| NICs (25G + 100G) | 25 + 2 × 30 W | 85 |
| Base 2U | — | 150 |
| **P_max théorique** | | **≈ 1 263 W** |
| **P_réaliste (0,75)** | | **≈ 950 W** |

| Critère | Compute S | Compute M |
|---|---|---|
| Cœurs / threads | 32 / 64 | 192 / 384 |
| RAM | 384 Go | 768 Go |
| Stockage data | 15,4 To bruts | 46 To bruts |
| Prix | ≈ 29 750 € | ≈ 56 975 € |
| P_max | 0,62 kW | 1,26 kW |
| €/cœur | ≈ 930 € | ≈ 297 € |
| Usage type | Virt légère, K8s, services | Virt dense, HPC, build |

**Le M est 3× moins cher au cœur.** Si vous virtualisez, le M gagne presque
toujours au TCO — sauf si vous n'avez besoin que de quelques VMs.

---

## 22. HPC petite grappe : l'interconnect décide de tout

Pour 4–8 nœuds HPC, le réseau est plus important que le CPU :
- **InfiniBand NDR (400 Gb/s)** : latence ≈ 1 µs, le choix MPI sérieux.
  Switch 8–16 ports NDR : 15 000–30 000 € (à vérifier — devis requis).
- **RoCE v2 sur 100/200 GbE** : latence 2–5 µs, 3× moins cher, suffisant pour
  80 % des codes (et pour Ceph).
- Ne faites pas de HPC sérieux sur du 25 GbE : le réseau devient le goulot
  avant les CPU.

BOM nœud HPC type (×4) : Compute M (section 20) + NIC IB NDR 400G
(≈ 2 500 €, à vérifier) − NIC 100G. Soit ≈ 58 500 €/nœud + switch IB.
**Total grappe 4 nœuds : ≈ 260 000 €**, P_max ≈ 5,5 kW (un demi-rack, air OK).

---

## 23. Cas chiffré : grappe Proxmox VE 3 nœuds HA

Hypothèses : 60 VMs bureautiques (4 vCPU / 16 Go), stockage Ceph intégré
ou ZFS répliqué, HA Proxmox.

| Poste | Détail | Total |
|---|---|---|
| 3× Compute M (section 20) | 56 975 € × 3 | 170 925 € |
| Switch 2× 25G + 2× 100G (ToR, redondant) | 2× ≈ 8 000 € | 16 000 € |
| Câblage (DAC, optiques, PDU déjà comptées) | lot | 2 500 € |
| Licences Proxmox VE (Community→Standard, 3 CPU-socket/an) | 3× ≈ 340 €/an | ≈ 1 020 €/an |
| **CAPEX** | | **≈ 189 500 €** |
| P_max grappe | 3 × 1,26 kW | **≈ 3,8 kW** |

Densité : 60 VMs × 16 Go = 960 Go RAM utile ; 3 nœuds × 768 Go = 2 304 Go bruts,
soit N+1 avec marge (perte d'un nœud : 1 536 Go restants > 960 × 1,15).
C'est le dimensionnement à retenir : **toujours N+1 sur la ressource critique**.

---

## 24. Kubernetes bare-metal : build worker type

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 1U 1P EPYC | AS-1115CS-TNR nu | 1 | 4 585 € |
| EPYC 9455P (48c/96t) | — | 1 | 3 406 € |
| RAM 12× 64 Go = 768 Go | DDR5 ECC | 12 | 16 560 € |
| 2× NVMe 3,84 To (OS + ephemeral) | PM9A3 | 2 | 6 070 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| **TOTAL/worker** | | | **≈ 31 070 €** |
| P_max / réaliste | | | **≈ 0,85 kW / 0,64 kW** |

Ratio 2026 constaté : **8 Go RAM par cœur** pour du K8s générique
(microservices Java/Go). Montez à 12–16 Go/cœur pour l'IA et la data.

---

## 25. Pièges terrain — Compute / HPC

