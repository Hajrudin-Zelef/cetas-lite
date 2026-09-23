---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-13-roce-congestion-control-algorithms-ultra-ethernet
title: "Wave 13 — RoCE congestion control algorithms & Ultra Ethernet"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Google", "Huawei", "Nvidia"]
dates: []
keywords: ["ethernet", "agents", "nvidia", "parameters", "research"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [380, 425]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 98c7ee11cfbce8a393ab0939d8900ccbbea2f285fcd10df49c7723f744be46ac
---

# Wave 13 — RoCE congestion control algorithms & Ultra Ethernet

## Wave 13 — RoCE congestion control algorithms & Ultra Ethernet

### 13.1 DCQCN (the deployed standard)
- DC-QCN: end-to-end congestion control for RoCEv2 combining ECN marking at switches, CNP feedback from receivers, and rate-based sender reaction (additive increase / multiplicative decrease on CNP receipt) — the reaction point (RP) lives in the NIC `[secondary]` — widely documented; Kar lab tuned it in Wave 6.
- NIC support: Mellanox/NVIDIA ConnectX implements DCQCN in firmware; parameters (alpha/gamma, rate decrease/increase factors, min/max rates) are NIC-driver tunables — defaults are NVIDIA-chosen, not IEEE-standardized `[vendor-reported]` — NVIDIA/Mellanox docs; FS.com ZTR whitepaper (Wave 6).
- Interaction rule: DCQCN must engage BEFORE PFC thresholds — otherwise PFC masks the congestion signal and DCQCN never learns `[secondary]` — Kar lab (Wave 6).

### 13.2 Alternatives and research directions
- TIMELY: RTT-gradient-based congestion control (Google) — reacts to delay rather than ECN marks; suited to environments where ECN tuning is impractical `[secondary]` — research literature (Mittal et al., SIGCOMM 2015).
- HPCC: high-precision congestion control using in-band telemetry (INT) — sender computes precise rate from per-hop queue data; needs INT-capable switches `[secondary]` — research literature (Li et al., SIGCOMM 2019).
- Swift (Google): delay-based, deployed in Google's datacenters as DCTCP/TIMELY successor `[secondary]` — research literature.
- Positioning: DCQCN is what ships in NICs today; TIMELY/HPCC/Swift are reference points for evaluating vendor "AI fabric" claims — ask which algorithm a vendor's NIC actually implements `[secondary]`.

### 13.3 Ultra Ethernet Consortium (UEC) — status note
- UEC is developing an Ethernet-based transport for AI/HPC with a new congestion-control and reliability story intended to succeed RoCEv2's PFC/ECN tuning burden; as of the Kar lab note (2026-05), "UEC reliable mode is 2027 onward — RoCEv2 is now and for the foreseeable installed base" `[secondary]` — Kar notebook; UEC public materials.
- Implication for this file: all lossless-tuning guidance here targets the RoCEv2 installed base; UEC migration planning is a 2027+ topic — flagged, not covered `[unverified]` as a deployment reality.

### Wave 13 verification
- Sources: 5. Research algorithms labeled as literature, not shipping defaults. UEC timing presented as forward-looking with explicit flag.

---

## Wave 14 — Multicast deep-dive: PIM mechanics, BSR/Auto-RP, mVPN

### 14.1 PIM-SM mechanics
- Neighbor discovery via PIM Hellos; DR (Designated Router) election per LAN — highest IP wins; the DR registers sources and forwards IGMP joins upstream `[official]` — RFC 7761 (PIM-SM).
- Shared tree (*,G) → source tree (S,G) switchover: default threshold (often 0 kbps = switch immediately, or traffic-rate based) moves receivers off the RP onto the shortest path `[official]` — RFC 7761.
- RPF check: multicast forwarding requires the packet to arrive on the interface toward the source/RP per the unicast table — RPF failure = drop; this is why RP reachability in the right VRF matters (Wave 8 TRM) `[official]` — RFC 7761.
- PIM timers that operators tune: Hello interval/holdtime, Join/Prune interval, register suppression, SPT switchover threshold, keepalive — defaults are conservative; data-center TRM designs often leave them default `[secondary]`.

### 14.2 RP discovery mechanics
- Static RP: configured on every router; simplest and deterministic; Juniper-recommended with anycast-RP `[official]` — Juniper docs (Wave 7).
- Auto-RP (Cisco): candidate-RPs announce via 224.0.1.39, mapping agents distribute via 224.0.1.40 — Cisco-proprietary `[official]` — Cisco docs.
- BSR (Bootstrap Router, PIMv2): candidate-BSRs elect a BSR that floods RP-set — standards-based `[official]` — RFC 7761.
- Guidance: in DC fabrics, static anycast-RP is the norm (small, controlled RP set); BSR/Auto-RP belong to larger enterprise/SP domains `[secondary]` — operator practice.

### 14.3 Inter-domain multicast & mVPN
- MSDP (RFC 3618): SA exchange between domains/RPs; mesh-groups for scale; SA filters and SA-limits as hygiene `[official]` — RFC 3618; Huawei docs (Wave 7).
- mVPN (multicast VPN, RFC 6513/6514): carrying customer multicast over a provider core — Rosen GRE MDT (default + data MDTs) vs MLDP (label-switched P2MP LSPs); NG-mVPN uses BGP (same SAFI family thinking as EVPN) `[official]` — RFC 6513/6514.
- DC relevance: mVPN is primarily an SP construct; DC fabrics use TRM/EVPN-multicast instead — do not conflate the two in designs `[secondary]`.

### Wave 14 verification
- Sources: 6 (RFC 7761, RFC 3618, RFC 6513/6514, Juniper, Cisco). DC-vs-SP scoping note recorded to prevent design-category errors.

---

