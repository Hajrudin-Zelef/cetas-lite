---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-13
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["asic", "cpo", "dci", "dsp", "ethernet", "gpu", "intel", "lpo", "nvidia", "optics", "serdes"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1799, 1936]
sha256: 8854afc50d182ccd74a7eb22d1974983b4ca007164b75b632713187394e26d8d
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Lecture** : le rythme est d'un doublement tous les ~2 ans (51,2T →
102,4T → 204,8T). **Ne pas acheter « pour 10 ans »** : le cycle réseau IA
est de **3-4 ans**. Budgéter en conséquence (leasing ou amortissement
accéléré).

## 103. PCIe Gen6 / Gen7 — le bus suit

| Génération | Débit x16 | NIC associée | Statut |
|---|---|---|---|
| Gen5 | 512 Gb/s | 400G (CX-7, Thor 2) | Déployé ✅ |
| Gen6 | 1024 Gb/s | 800G (CX-8, BF-4) | ✅ (serveurs 2025-2026) |
| Gen7 | 2048 Gb/s | 1,6T (CX-9) | ⏳ Spéc finalisée, produits 2027+ (⚠️) |

**Point de vigilance** : une NIC 1,6T sur PCIe Gen6 x16 serait bridée —
le 1,6T **exige** le Gen7 (ou du x32, inexistant). Vérifier la carte mère
**avant** d'acheter des NIC 1,6T en 2027.

## 104. Conclusion : les 10 décisions qui comptent

1. **Séparer** backend IA (1:1, lossless) et front-end (3:1) — toujours.
2. **Choisir le débit par le joule/bit**, pas par le prix catalogue.
3. **Valider le PCIe** (génération × lanes électriques) avant chaque NIC.
4. **Figer le triplet** FW/driver/OS par fabric.
5. **Nettoyer la fibre** (procédure affichée, kit par salle).
6. **Superviser le pré-FEC**, pas seulement le link up/down.
7. **Exiger UEC-ready** sur les AO 2026+ (assurance anti-lock-in).
8. **Chiffrer l'énergie** (switch + optiques + NIC + PUE) dans chaque projet.
9. **Prévoir le CPO** : pas d'engagement pluggable 800G à 7 ans.
10. **Cycle 3-4 ans** : le réseau IA ne s'amortit plus sur 7 ans.

## 105. Glossaire — 40 termes

| Terme | Définition |
|---|---|
| AEC | Active Electrical Cable — cuivre actif, amplifie le signal (3-6 m). |
| AOC | Active Optical Cable — fibre active avec modules intégrés. |
| ASIC | Circuit intégré dédié (ex. Tomahawk, Spectrum) — le cœur du switch. |
| BDP | Bandwidth-Delay Product — dimensionne les buffers (débit × RTT). |
| Breakout | Diviser un port rapide en N ports lents (ex. 400G → 4×100G). |
| CPO | Co-Packaged Optics — optiques intégrées au package du switch. |
| Cut-through | Commutation dès la lecture de l'en-tête (~600 ns) vs store-and-forward. |
| DAC | Direct Attach Cable — cuivre passif Twinax, courte portée. |
| DCQCN | Contrôle de congestion pour RoCEv2 (ECN + CNP). |
| DPU | Data Processing Unit — NIC + ARM + accélérateurs. |
| DSP | Digital Signal Processor — dans les modules optiques classiques. |
| ECN | Explicit Congestion Notification — marquage anti-congestion. |
| ECMP | Répartition sur N chemins de coût égal (hash par flux). |
| EVPN | BGP qui distribue MAC/IP pour les overlays VXLAN. |
| FEC | Forward Error Correction — Reed-Solomon, rend le PAM4 viable. |
| IPU | Infrastructure Processing Unit — nom Intel du DPU. |
| Incast | Many-to-one : N émetteurs vers 1 récepteur (all-reduce). |
| LPO | Linear Pluggable Optics — optique sans DSP, basse conso. |
| MPO/MTP | Connecteur multi-fibres (12/16/24). |
| NDR / XDR | InfiniBand 400G / 800G. |
| NOS | Network Operating System (EOS, NX-OS, SONiC, Cumulus). |
| NRZ / PAM4 | Modulations 1 bit / 2 bits par symbole. |
| OCP 3.0 | Format de carte réseau pour serveurs hyperscale. |
| OSFP / QSFP-DD | Formats de modules 400/800G (OSFP = meilleure thermique). |
| Overlay / Underlay | Réseau virtuel (VXLAN) / réseau physique (BGP) qui le porte. |
| PFC | Priority Flow Control — pause par file (lossless). |
| PUE | Power Usage Effectiveness — efficacité énergétique du DC. |
| Rail-optimized | 1 plan réseau par indice de NIC GPU. |
| RDMA | Accès mémoire distant sans CPU (zéro-copy). |
| RoCEv2 | RDMA over Converged Ethernet (UDP 4791). |
| SerDes | Serializer/Deserializer — l'étage électrique du port. |
| SONiC | NOS open-source (Microsoft, Linux Foundation). |
| Spectrum-X | Pile Ethernet IA fermée de NVIDIA. |
| Spraying (packet) | Répartition à la granularité paquet (vs par flux). |
| SR-IOV | Virtualisation PCIe : 1 carte → N fonctions virtuelles. |
| SuperNIC | NIC IA avec fonctions réseau avancées (NVIDIA). |
| UEC / UET | Ultra Ethernet Consortium / son transport. |
| VTEP | VXLAN Tunnel EndPoint — extrémité de tunnel. |
| VXLAN | Overlay L2 sur L3 (16M de segments). |
| ZR | Optique cohérente longue portée (80 km+, DCI). |

# PARTIE K — APPROFONDISSEMENTS TERRAIN

## 106. Supervision d'un fabric IA : télémétrie, compteurs, seuils

Un fabric IA ne se supervise pas comme un LAN : on surveille des
**distributions** (latence p99, BER), pas des états binaires.

| Sonde | Compteur / métrique | Seuil 🟡 | Seuil 🔴 | Outil |
|---|---|---|---|---|
| NIC | pre-FEC BER | > 1e-9 | > 1e-6 | ethtool -S, telemetry |
| NIC | FEC uncorrectable | > 0 | > 0 | idem |
| NIC | CNP reçus/s | variation > 50 % | storm | compteurs RoCE |
| NIC | retransmissions | > 0,01 % | > 0,1 % | — |
| Switch | drops par file | > 0 | > 100/s | SNMP/gNMI |
| Switch | ECN marks/s | > 10 % paquets | > 30 % | — |
| Switch | PFC pauses/s | > 0 | > 10/s | ⛔ = deadlock imminent |
| Switch | buffer occupancy p99 | > 70 % | > 90 % | LANZ (Arista), telemetry |
| Optique | Température | > 70 °C | > 80 °C | DOM |
| Optique | Rx power vs sensibilité | marge < 3 dB | marge < 1 dB | DOM |
| Fabric | latence p99 leaf↔leaf | > 5 µs | > 20 µs | sondes actives |

**Stack recommandée** : telemetry streaming (gNMI/gRPC) → collecteur
(Telegraf/Prometheus) → Grafana + alertes. Le SNMP polling à 5 min est
**aveugle** aux microbursts (µs-ms) : pour l'IA, il faut du streaming à
la seconde ou des histogrammes embarqués (LANZ chez Arista).

**Règle** : 1 dashboard « santé fabric » avec 6 graphes max (BER, drops,
ECN, PFC, température optiques, latence p99). Le reste en drill-down.
Un NOC qui regarde 200 graphes n'en regarde aucun.

## 107. Tests de réception d'un fabric : la méthode avant la signature

Ne jamais signer la réception d'un fabric IA sans ces tests :

| Test | Outil | Critère de passage (⚠️) |
|---|---|---|
| Bande passante paire à paire | `ib_write_bw` / `qperf` | ≥ 95 % du débit nominal |
| All-reduce NCCL 8→N nœuds | `nccl-tests` (all_reduce_perf) | busbw ≥ 80 % du théorique |
| Latence | `ib_write_lat` | < 2 µs (RoCE), < 1 µs (IB) |
| ECMP : tous les spines utilisés | compteur par lien sous charge | écart < 10 % entre liens |
| Perte d'un spine | débrancher 1 spine sous charge | < 5 % de perte de busbw, convergence < 1 s |
| Perte d'un leaf | idem | trafic rerouté, pas de blackhole |
| 72 h de burn-in | charge NCCL continue | 0 erreur FEC non corrigée, 0 flap |
| Bruit thermique | charge max 1 h | aucune optique > 75 °C |

**Procédure** : tester **d'abord à vide** (1 flux), **puis en all-to-all**
(tous les nœuds). Un fabric qui passe à vide et échoue en all-to-all a un
problème d'ECMP/CC — c'est le cas le plus fréquent (P39, §122).
**Documenter les résultats** : c'est la baseline. Toute régression future
se mesure contre elle.

## 108. NCCL et collectives : ce que le réseau doit garantir

NCCL (NVIDIA Collective Communications Library) = la bibliothèque qui fait
communiquer les GPU entre eux. Le réseau doit garantir :

| Collective | Pattern | Exigence réseau |
|---|---|---|
| All-Reduce | Tous ↔ tous (anneau/arbre) | Bande passante bissection max, 1:1 |
| All-Gather | Tous reçoivent tout | Idem |
| Broadcast | 1 → tous | Faible latence |
| Point-to-point | Paires | RDMA, zéro-copy |

