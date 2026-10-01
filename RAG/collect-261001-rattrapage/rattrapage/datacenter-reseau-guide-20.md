---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-20
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Google", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["datacenter", "asic", "compute", "cost", "cpo", "dci", "distribution", "dsp", "ethernet", "intel", "latency", "lpo"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [2852, 3048]
sha256: 676a7d82f3c4e355c037b3dc9a431282a0eed1fdd5f9b3c71e46fb07e0acb011
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Template test de réception** (un par fabric) :
```
Fabric : __________    Date : __________    Opérateur : __________
[ ] ib_write_bw ≥ 95 % : ____ Gb/s mesurés
[ ] NCCL all-reduce busbw ≥ 80 % : ____
[ ] Latence < 2 µs : ____ µs
[ ] ECMP ±10 % : max ____ / min ____
[ ] Perte 1 spine : ____% perte, ____ s convergence
[ ] Burn-in 72 h : ____ erreurs FEC, ____ flaps
[ ] Thermique : max optique ____°C
Réserve(s) : __________
Signature : __________
```

## 145. Lexique des acronymes — 100 entrées

| Acronyme | Signification |
|---|---|
| ACI | Application Centric Infrastructure (Cisco) |
| ACL | Access Control List |
| ADQ | Application Device Queues (Intel) |
| AEC | Active Electrical Cable |
| AOC | Active Optical Cable |
| ASIC | Application-Specific Integrated Circuit |
| ASN | Autonomous System Number |
| BDP | Bandwidth-Delay Product |
| BER | Bit Error Rate |
| BFD | Bidirectional Forwarding Detection |
| BGP | Border Gateway Protocol |
| BiDi | Bidirectional (1 fibre, 2 longueurs d'onde) |
| BOM | Bill of Materials |
| Bpps | Billion packets per second |
| CEM | Card Electromechanical (PCIe) |
| CNP | Congestion Notification Packet |
| CPO | Co-Packaged Optics |
| DAC | Direct Attach Cable |
| DCB | Data Center Bridging |
| DCI | Data Center Interconnect |
| DCQCN | Data Center Quantized Congestion Notification |
| DOM | Digital Optical Monitoring |
| DPU | Data Processing Unit |
| DR | 500 m (optique) |
| DSP | Digital Signal Processor |
| DWDM | Dense Wavelength Division Multiplexing |
| ECMP | Equal-Cost Multi-Path |
| ECN | Explicit Congestion Notification |
| EDFA | Erbium-Doped Fiber Amplifier |
| EOS | Extensible Operating System (Arista) |
| ETS | Enhanced Transmission Selection |
| EVPN | Ethernet VPN |
| FEC | Forward Error Correction |
| FR | 2 km (optique) |
| FW | Firmware |
| GENEVE | Generic Network Virtualization Encapsulation |
| gNMI | gRPC Network Management Interface |
| HCA | Host Channel Adapter (InfiniBand) |
| IB | InfiniBand |
| iWARP | Internet Wide Area RDMA Protocol |
| LACP | Link Aggregation Control Protocol |
| LANZ | Latency Analyzer (Arista) |
| LPO | Linear Pluggable Optics |
| LR | 10 km (optique) |
| LRO | Linear Receive Optics |
| MAC | Media Access Control / Moves-Adds-Changes |
| MACsec | Media Access Control Security |
| MDA | Main Distribution Area |
| MPO | Multi-fiber Push-On |
| MSA | Multi-Source Agreement |
| MTU | Maximum Transmission Unit |
| NCCL | NVIDIA Collective Communications Library |
| NDR | 400G InfiniBand |
| NOS | Network Operating System |
| NRZ | Non-Return-to-Zero |
| NVMe-oF | NVMe over Fabrics |
| OCP | Open Compute Project |
| OFED | OpenFabrics Enterprise Distribution |
| OM3/4/5 | Optical Multimode (grades) |
| ONIE | Open Network Install Environment |
| OOB | Out-Of-Band |
| OS2 | Optical Single-mode (grade) |
| OSFP | Octal Small Form-factor Pluggable |
| OVS | Open vSwitch |
| PAM4 | Pulse Amplitude Modulation (4 niveaux) |
| PFC | Priority Flow Control |
| PoP | Point of Presence |
| PTP | Precision Time Protocol |
| PUE | Power Usage Effectiveness |
| Qbb/Qaz | Normes 802.1 (PFC/ETS) |
| QSFP | Quad Small Form-factor Pluggable |
| RDMA | Remote Direct Memory Access |
| RoCE | RDMA over Converged Ethernet |
| RTT | Round-Trip Time |
| SFP | Small Form-factor Pluggable |
| SLA | Service Level Agreement |
| SNMP | Simple Network Management Protocol |
| SONiC | Software for Open Networking in the Cloud |
| SR | Short Reach |
| SR-IOV | Single-Root I/O Virtualization |
| STP | Spanning Tree Protocol |
| SWDM | Short Wavelength Division Multiplexing |
| TAC | Technical Assistance Center |
| ToR | Top-of-Rack |
| TSO/GRO | TCP Segmentation / Generic Receive Offload |
| UEC | Ultra Ethernet Consortium |
| UET | Ultra Ethernet Transport |
| VMDq | Virtual Machine Device Queues |
| VNI | VXLAN Network Identifier |
| VRF | Virtual Routing and Forwarding |
| VTEP | VXLAN Tunnel EndPoint |
| VXLAN | Virtual Extensible LAN |
| WDM | Wavelength Division Multiplexing |
| XDR | 800G InfiniBand |
| ZR | 80 km+ (optique cohérente) |
| ZTP | Zero-Touch Provisioning |

## 146. Timeline — 25 ans d'Ethernet datacenter

| Année | Jalon |
|---|---|
| 1999 | Gigabit Ethernet (1000BASE-SX/LX) : le DC passe au gigabit |
| 2002 | 10G (802.3ae) : agrégation |
| 2010 | 40/100G (802.3ba) : QSFP+, le DC moderne naît |
| 2012 | Spine-leaf / Clos : Google, puis tous les hyperscalers |
| 2014 | 25/50G (802.3by) : le 25G remplace le 10G |
| 2015 | RoCEv2 : le RDMA sur Ethernet devient standard |
| 2017 | 200/400G (802.3bs), PAM4 : la modulation change |
| 2018 | SONiC open-source (Microsoft) : le whitebox décolle |
| 2020 | 400G en volume, DAC/AOC 400G |
| 2021 | BlueField-2/3 : le DPU entre au DC |
| 2023 | **UEC fondé** (07/2023) : l'industrie veut un Ethernet IA ouvert |
| 2024 | 800G (802.3df), Spectrum-X, Thor 2 |
| 2025 | **ConnectX-8**, **Thor Ultra** (UEC 1.0, 10/2025), **BlueField-4** (10/2025), **Tomahawk 6-Davisson CPO** (10/2025), **UEC 1.0** (06/2025) |
| 2026 | OFC : Broadcom Taurus (400G/lane), ère 200T annoncée ; VIAVI valide UET |
| 2027+ | 1,6T (ConnectX-9, Rubin), CPO en volume, UEC multi-vendeurs |

## 147. Comparatif NOS — EOS, NX-OS, SONiC, Cumulus

| Critère | Arista EOS | Cisco NX-OS | SONiC (durci) | NVIDIA Cumulus |
|---|---|---|---|---|
| Modèle | Propriétaire | Propriétaire | Open-source | Propriétaire (Linux) |
| Maturité IA | ✅ Etherlink | ⚠️ (G200 récent) | ⚠️ (selon distribution) | ✅ Spectrum-X |
| Automation | eAPI, CloudVision | NX-API, NDFC | gNMI, standard | Ansible, NCLU |
| ZTP | ✅ | ✅ | ✅ | ✅ |
| Telemetry | LANZ, streaming | Streaming | Streaming | Streaming |
| Coût licence | Élevé (⚠️) | Élevé (⚠️) | Faible (support à part) | Moyen (⚠️) |
| Courbe d'apprentissage | Faible (CLI Cisco-like) | Faible | Haute (Linux + SAI) | Moyenne (Linux) |
| Idéal pour | Enterprise, IA | Existant Cisco | Whitebox, hyperscale-like | Fabrics NVIDIA |

**Règle** : on ne change pas de NOS comme de chemise — le coût de formation
et de réécriture d'automation dépasse l'économie de licence. Choisir pour
10 ans, pas pour le prix de l'année 1.

## 148. Checklist d'audit sécurité du fabric (trimestrielle)

- [ ] Comptes : pas de `admin/admin`, AAA partout, revue des accès
- [ ] Firmware : versions à jour **dans la politique figée** (pas « latest »)
- [ ] Images signées : secure boot actif sur switchs et DPU
- [ ] Management : VRF/OOB séparé, ACLs infra, pas de SSH ouvert sur le data
- [ ] SNMP : v3 uniquement (pas de v2c « public » qui traîne)
- [ ] Logs : centralisés, 1 an de rétention, alertes sur login anormal
- [ ] Configs : sauvegardées, versionnées (Git), diff avant/après chaque change
- [ ] Ports : 802.1X ou MAC lockdown sur les ports serveurs sensibles
- [ ] DCI : chiffrement actif et vérifié (pas seulement « configuré »)
- [ ] Plan de réponse : qui appeler, quel TAC, quel spare — affiché au NOC

## 149. Supervision — exemples d'alertes Prometheus

```yaml
# BER pré-FEC : alerte si > 1e-9 pendant 10 min
- alert: HighPreFecBer
  expr: pre_fec_ber > 1e-9
  for: 10m
  labels: { severity: warning }
  annotations: { summary: "BER pré-FEC élevée sur {{ $labels.iface }}" }

# Erreurs FEC non corrigées : critique immédiate
- alert: FecUncorrectable
  expr: increase(fec_uncorrectable_total[5m]) > 0
  labels: { severity: critical }

# Pauses PFC : signe de deadlock
- alert: PfcStorm
  expr: rate(pfc_pause_frames_total[5m]) > 10
  labels: { severity: critical }

# Température optique
- alert: OpticHot
  expr: optic_temp_celsius > 70
  for: 15m
  labels: { severity: warning }

# Marge optique faible (Tx - Rx vs sensibilité)
- alert: OpticLowMargin
  expr: (optic_rx_dbm - optic_rx_sensitivity_dbm) < 3
  for: 30m
  labels: { severity: warning }

