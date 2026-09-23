---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-24-deployment-notes-and-field-patterns
title: "Wave 24 — Deployment notes and field patterns"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Broadcom", "Google", "Meta", "Nvidia", "xAI"]
dates: []
keywords: ["benchmarks", "cost", "cpo", "ethernet", "license", "nvidia", "pricing", "serdes", "throughput"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [574, 628]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: ebb9ff430b4e7c9465484e3a318040e1d82cce12bfbbe8efea9db56b272e91c9
---

# Wave 24 — Deployment notes and field patterns

## Wave 24 — Deployment notes and field patterns

- **NVMe-oF on leaf-spine (Arista + Pure Storage):** two-tier fabric, leaves as EVPN L3 VTEPs, spines as EVPN route servers; lossless via PFC + ECN with DSCP-mapped transmit queue; edge ports L2 access with SVI gateways [official — https://www.purestorage.com/content/dam/pdf/en/white-papers/wp-arista-pure-deploying-nvme-of-enterprise-leaf-spine-architecture.pdf].
- **Campus extension:** Arista's campus design guide applies the same leaf-spine + EVPN-VXLAN pattern to wired campus (4 leaves + 2 spines sample), with /31 p2p links, eBGP underlay, iBGP between leaves, MLAG leaf pairs [official — https://manuals.plus/m/afc4d8af60bb9487a58a67b736c26465c47bb80e17fa671bb8d59731376a17e7_optim.pdf].
- **Multi-vendor pods:** community demo topologies mix Cisco/Arista/Dell/Edgecore pods under one super-spine pair (N9K-C9336C-FX2 class) with per-pod "personalities" — feasible because eBGP/EVPN are standards-based; operational cost is the real tax [secondary — https://github.com/t0m3kz/infrahub-demo/blob/HEAD/data/demos/01_data_center/dc5/README.md].
- **Maintenance discipline:** drain one spine at a time; run N−1 spine capacity; keep steady-state link utilization ≤70%; validate ECMP spread after every change [secondary].

---

## Wave 25 — Final coverage audit (supersedes Wave 21)

**Scope-to-wave map (complete, 25 waves):** Clos/spine-leaf principles, oversubscription, scaling math (1, 7, 23) · failure domains, cabling plans (1, 20) · 3-tier vs spine-leaf, migration (2) · Cisco NX-OS/ACI verified scale + 2026 debate + G300/Nexus One/Hyperfabric + CVE-2026-20212 (3, 8, 17) · Arista hardware/AVD/EVPN-GW/7368X4 (3, 9, 22) · NVIDIA Spectrum-X/Cumulus (3, 10) · Dell SFS/OS10 (3, 11) · Aruba CX/Fabric Composer (3, 11) · Juniper Apstra 5-stage JVD (3, 12) · whitebox/SONiC economics (3, 15) · underlay essentials: numbered/unnumbered, RFC 5549, BFD (4, 10, 11, 23) · multi-homing (19) · silicon radix: Tomahawk 6, G300 (7, 13, 17) · AI topologies: rail-only/rail-optimized, ROD/RUD (18) · hyperscalers: Meta DSF/RoCE, xAI Colossus, RSC (5, 14) · vendor comparison matrix (22) · deployment notes (24) · decision checklist (6).

**Provenance audit:** claims carry [official]/[vendor-reported]/[independent]/[secondary]/[unverified]; vendor-constructed models and gray-market prices flagged; derived rows marked [analysis].

**Open items / [gap]s (final, 10):** (1) Google Jupiter-class current specifics; (2) Azure SONiC fleet counts; (3) HPE/Juniper NOS roadmap; (4) independent 800G fabric benchmarks; (5) Dell Enterprise SONiC license pricing; (6) Cisco ACI per-release max leaf counts; (7) xAI Colossus exact topology [unverified]; (8) Tomahawk 6 independent power/thermal; (9) CVE-2026-20212 remediation → security phase; (10) Nexus One/Hyperfabric independent reviews.

*End of Phase D1 (25 waves). Append-only; earlier sections untouched. Deep BGP underlay and EVPN-VXLAN protocol coverage continues in sibling Phase D files.*

---

## Wave 26 — Switch-silicon generational table (radix context for fabric design)

| Silicon | Vendor | Throughput | SerDes | Era | Notes |
|---|---|---|---|---|---|
| Tomahawk 3 | Broadcom | 12.8T | 50G | 2018 | 32×400G class |
| Tomahawk 4 | Broadcom | 25.6T | 100G | 2020 | 64×400G |
| Tomahawk 5 | Broadcom | 51.2T | 100G | 2022 | 64×800G; powers many 2024–25 800G switches |
| Tomahawk 6 | Broadcom | 102.4T | 224G | 2025–26 | 128×800G / 64×1.6T; Davisson CPO variant (Waves 13, 7) |
| Silicon One G100 | Cisco | 25.6T | — | 2019–20 | First Silicon One |
| Silicon One G200 | Cisco | 51.2T | — | 2022–23 | Nexus 9000 AI fabrics |
| Silicon One G300 | Cisco | 102.4T | — | Feb 2026 | 64×1.6T; N9000/8000 systems (Wave 17) |
| Spectrum-2 | NVIDIA/Mellanox | 6.4T | 50G | 2019 | — |
| Spectrum-3 | NVIDIA/Mellanox | 12.8T | 100G | 2020 | Cumulus/SONiC workhorse |
| Spectrum-4 | NVIDIA | 51.2T | 100G | 2023 | Spectrum-X AI platform (Wave 10) |
| Jericho3-AI | Broadcom | — | — | 2024 | Meta DSF leaf (7700R4C-38PE) |
| Ramon3 | Broadcom | — | — | 2024 | Meta DSF spine (7720R4-128PE, 102.4T system) |

Throughput figures [vendor-reported]; eras approximate from launch announcements [secondary]. Radix doubling every ~2–3 years is what keeps collapsing fabric tiers (fewer tiers for the same server count).

## Wave 27 — Fabric glossary (terms used in this file)

- **ASN** — Autonomous System Number; private range 64512–65535 (and 4200000000–4294967295) used per RFC 7938-style fabrics: unique AS per leaf, shared AS per spine tier [secondary].
- **BFD** — Bidirectional Forwarding Detection; sub-second link liveness for BGP [official].
- **BUM** — Broadcast, Unknown-unicast, Multicast; the traffic class overlays must replicate [secondary].
- **Clos** — multistage non-blocking topology (Charles Clos, 1953); leaf-spine is a 2-stage folded Clos [secondary].
- **ECMP** — Equal-Cost Multi-Path; flow-hashed load sharing across equal paths [secondary].
- **ERB** — Edge-Routed Bridging; routing at the leaf (vs CRB, centrally-routed) [official].
- **ESI** — Ethernet Segment Identifier; EVPN multi-homing construct [official].
- **IRB** — Integrated Routing and Bridging; symmetric vs asymmetric models [official].
- **Oversubscription** — downlink BW ÷ uplink BW per leaf (Wave 7) [official].
- **Radix** — switch port count; bounds fabric scale (Wave 1.3) [secondary].
- **RFC 5549** — IPv4 NLRI with IPv6 next-hop; enables unnumbered eBGP underlays [official].
- **VTEP / VNI** — VXLAN Tunnel EndPoint / VXLAN Network Identifier [secondary].

