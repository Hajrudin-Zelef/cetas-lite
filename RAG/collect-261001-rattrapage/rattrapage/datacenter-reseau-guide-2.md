---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-2
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: ["2025-06-11", "2025-10-14", "2026-09-27"]
keywords: ["asic", "blackwell", "chiplet", "cpo", "ethernet", "gpu", "intel", "lpo", "nvidia", "rubin", "serdes"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [115, 243]
sha256: 898fa81ea7695afa780d029a2d5fb950c319f13ce259d10a7591031da8b5416b
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Intel E810 (détail vérifié le 27/09/2026)** : contrôleur E810-CAM1,
PCIe 4.0 x16, 2 ports QSFP28 100G (10/25/50/100G auto), SR-IOV 256 fonctions
virtuelles, VMDq, VXLAN/GENEVE/NVGRE en offload, RDMA iWARP + RoCEv2, DPDK,
IEEE 1588 PTP, DCB (PFC/ETS). Conso relevée : 15 W (version 2 ports, fiche
distributeur) à 27,1 W avec AOC — l'optique compte, voir §89.
C'est la NIC 100G « enterprise » de référence quand on ne fait pas d'IA.

## 5. Le 200 GbE : le palier de transition

Le 200G (QSFP56, 4×50G PAM4) a été un palier court : il sert surtout en
**mode dégradé** des NIC 400G (ConnectX-7 en 200G) ou en breakout
(1×400G → 2×200G, 1×800G → 4×200G).

Cas d'usage réel : stockage NVMe-oF où 100G est juste et 400G surdimensionné,
et vieux fabrics InfiniBand HDR (200G) en migration vers Ethernet.
**En achat neuf 2026 : on ne cible plus le 200G natif**, on achète du 400G
capable de négocier 200G.

## 6. Le 400 GbE : le backend IA standard (2024-2026)

Le 400G est le débit du backend des clusters GPU Hopper (H100/H200) :
**1 NIC 400G par GPU** (ratio 1:1, rail-optimized, voir §39) ou 1 NIC pour
2 GPU sur les designs économiques.

| NIC 400G | PCIe | Format | Distinction clé | Statut |
|---|---|---|---|---|
| NVIDIA ConnectX-7 | Gen5 x16 | QSFP112 / OSFP | RDMA/RoCEv2, la référence backend Hopper | ✅ vérifié 27/09/2026 |
| Broadcom Thor 2 (BCM957608) | Gen5 x16 | QSFP112 | 5 nm, faible conso, RoCE + apports UEC (ACK sélectif, multipath niveau paquet) | ✅ vérifié 27/09/2026 |
| Intel 400G | — | — | ❌ **Non trouvée au 27/09/2026** : Intel n'a toujours pas de NIC 400G sur le marché (confirmé par la presse spécialisée) | ❌ |
| NVIDIA BlueField-3 SuperNIC | Gen5 x16 | 2×400G ou 1×800G | ConnectX-7 + ARM : réordonne les paquets, contrôle de congestion HW, télémétrie par flux | ✅ vérifié 27/09/2026 |

**Broadcom Thor 2 (détail vérifié le 27/09/2026)** : ASIC 5 nm présenté comme
« purpose-built AI », 400 Gb/s, PCIe 5.0 x16, 8 lanes SerDes 100G PAM4/50G/25G
NRZ, DAC passif jusqu'à ~5 m vers switch Tomahawk 5, optiques linéaires
basse conso (LPO, voir §28). Disponible en carte, chiplet ou IP.
C'est le challenger « ouvert » de ConnectX-7 pour les fabrics Broadcom/Arista.

**Point d'architecte** : à 400G, le CPU ne peut plus traiter les paquets en
logiciel (trop d'interruptions, trop de copies). Les **offloads** (§9) et le
**RDMA** (§10) ne sont plus optionnels : sans eux, 2×400G sur un serveur =
CPU saturé avant la moitié du débit.

## 7. Le 800 GbE : la génération 2025-2026

Le 800G (8×100G PAM4, OSFP ou QSFP-DD800) est le débit du backend des
clusters Blackwell (B200/B300) et du scale-out Rubin.

| NIC 800G | PCIe | Format | Distinction clé | Statut |
|---|---|---|---|---|
| NVIDIA ConnectX-8 SuperNIC | **Gen6** x16 (jusq. 48 lanes) | OSFP224 1×800G ou 2×112 | In-Network Computing, MPI_Alltoall, QoS + contrôle de congestion, Socket Direct 16 lanes, OCP 3.0 / CEM | ✅ vérifié 27/09/2026 |
| Broadcom Thor Ultra | (PCIe Gen5/6 ⚠️) | 800G | **Premier NIC conforme UEC 1.0** (annoncé 14/10/2025), 1 seul flux 800G (pas 2×400G), RDMA modernisé : retransmission sélective, CC émetteur/récepteur programmable | ✅ vérifié 27/09/2026 |
| NVIDIA ConnectX-9 SuperNIC | — | — | **1,6 Tb/s** scale-out, associé à l'architecture Rubin (2ᵉ sem. 2026) | ✅ annoncé, voir §99 |

**ConnectX-8 (détail vérifié le 27/09/2026)** : dévoilé à COMPUTEX 2025,
firmware 40.46.x (sept. 2025). Au-delà du 800G, c'est un **commutateur PCIe**
: dans les serveurs RTX PRO, il remplace plusieurs switchs PCIe discrets,
donnant 400 Gb/s par GPU (ratio 2:1 GPU:NIC) et ~2× de perf NCCL all-to-all.
Supporte le chiffrement inline (TLS/IPsec) sans perte de débit (moteurs
crypto dédiés). Écosystème DOCA 2.0.

**Lecture énergétique** : le 800G double le débit pour ~+50-70 % de conso NIC
(⚠️ ordre de grandeur) — le **joule par bit baisse**. C'est l'argument à
porter en comité : monter en débit, c'est aussi de l'efficacité énergétique
(voir §89-92).

## 8. PCIe requis par débit — le tableau qui évite les bridages

Débits théoriques **unidirectionnels** par génération (x16) :

| PCIe | Gb/s théo. x16 | Débit NIC max sain | Marge utile |
|---|---|---|---|
| Gen3 x16 | 128 | 25G (100G bridé ~50 %) | — |
| Gen4 x16 | 256 | 100G, 200G limite | 200G = 78 % du bus : OK mais sans marge |
| Gen5 x16 | 512 | 400G | 78 % : le standard backend IA |
| Gen6 x16 | 1024 | 800G | 78 % : ConnectX-8, BlueField-4 ✅ |

Règles pratiques :
- **Toujours x16 électriques** pour ≥100G : un slot mécanique x16 câblé x8
  (fréquent sur serveurs 1U denses) divise par 2. Vérifier `lspci` :
  `LnkSta: Speed 32GT/s, Width x16` (Gen5).
- **Bifurcation** : pour 2× NIC 200G sur un slot x16, la carte mère doit
  supporter la bifurcation x8/x8 — vérifier dans le BIOS/fiche serveur.
- **Socket Direct** (NVIDIA) : carte auxiliaire 16 lanes qui relie la NIC aux
  deux sockets CPU sans passer par un switch PCIe — réduit la latence NUMA.
  ✅ supporté par ConnectX-8 (vérifié 27/09/2026).
- **CXL** : ne remplace pas le PCIe pour la NIC en 2026 ; le CXL.mem sert à
  l'extension mémoire, pas au réseau. ❌ Pas de NIC CXL réseau trouvée au
  27/09/2026 — ne pas confondre avec les annonces CPO.

## 9. Offloads : ce que la NIC fait à la place du CPU

À partir de 100G, chaque feature ci-dessous vaut des cœurs CPU entiers :

| Offload | Ce qu'il fait | Gain typique | Disponible sur |
|---|---|---|---|
| Checksum (Tx/Rx) | Calcule IP/TCP/UDP en HW | ~1-2 cœurs à 100G | Toutes |
| TSO (TCP Segmentation) | Découpe les gros segments en HW | −30-50 % CPU à 100G | Toutes |
| GRO/LRO (Rx) | Réassemble côté réception | Idem côté Rx | Toutes |
| SR-IOV | 256 VF PCIe vers les VM (bypass hyperviseur) | Latence ÷2-3 vs vSwitch | E810 (256 VF ✅), CX-6/7 |
| VMDq / ADQ (Intel) | Files Rx par VM/application | QoS par workload | E810 ✅ |
| VXLAN/GENEVE/NVGRE | Encap/decap overlay en HW | Indispensable EVPN/VXLAN à 100G+ | E810, CX-6/7, Thor 2 |
| DPDK / XDP | Bypass kernel, poll-mode | 10-100 Mpps par cœur | E810, CX, Thor |
| TLS/IPsec inline | Chiffre en HW à plein débit | Sécurité sans taxe CPU | ConnectX-8 ✅, BF-3 |
| RoCEv2 | RDMA sur UDP en HW | Zéro-copy, CPU ~0 | CX-6/7/8, Thor 2, E810 (iWARP+RoCEv2 ✅) |

**Test de réception** : `ethtool -k ethX` doit montrer `tx-checksumming: on`,
`scatter-gather: on`, `tcp-segmentation-offload: on`. Si un offload est
« off [fixed] », c'est le driver ou le firmware — mettre à jour avant
d'incriminer le réseau (P2, §107).

## 10. RDMA : RoCEv2, iWARP, InfiniBand natif — qui fait quoi

| Transport | Support | Latence typ. | Lossless requis ? | Écosystème IA |
|---|---|---|---|---|
| InfiniBand natif (NDR 400G, XDR 800G) | IB | ~0,6-1 µs | Oui natif (crédits) | NCCL/MPI par défaut, 20 ans de maturité |
| RoCEv2 (UDP 4791) | Ethernet | ~1-2 µs | Oui (PFC+ECN) à régler | Standard backend Ethernet IA |
| iWARP (TCP/SCTP) | Ethernet | ~2-5 µs | Non (TCP gère) | Stockage, moins HPC |
| UET (Ultra Ethernet) | Ethernet | ~1 µs visé | **Non** (conçu lossy) | ✅ UEC 1.0 du 11/06/2025, Thor Ultra |

**DCQCN** (Data Center Quantized Congestion Notification) = l'algorithme de
contrôle de congestion de RoCEv2 : la NIC marque les paquets (ECN), le
récepteur renvoie des CNP, l'émetteur réduit son débit. Mal réglé, c'est la
cause n°1 des contre-performances RoCE (voir §45-46).
**Point 2026** : UEC supprime l'exigence du lossless (livraison non ordonnée
+ retransmission sélective), ce qui change le dimensionnement des buffers
(voir §54).

## 11. Critères de choix d'une NIC — matrice de décision

