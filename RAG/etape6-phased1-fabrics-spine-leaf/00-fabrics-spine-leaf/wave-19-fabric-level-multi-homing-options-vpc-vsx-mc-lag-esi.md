---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-19-fabric-level-multi-homing-options-vpc-vsx-mc-lag-esi
title: "Wave 19 — Fabric-level multi-homing options (vPC / VSX / MC-LAG / ESI)"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Google", "Nvidia", "xAI"]
dates: []
keywords: ["acquisition", "benchmarks", "cost", "ethernet", "hyperscaler", "license", "nvidia", "pricing"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [483, 528]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 5d314b38c7c4bb25f131b8c755c77bcca1505b658dea3740da89888fbdc057c8
---

# Wave 19 — Fabric-level multi-homing options (vPC / VSX / MC-LAG / ESI)

## Wave 19 — Fabric-level multi-homing options (vPC / VSX / MC-LAG / ESI)

Deep protocol detail belongs to sibling files; fabric-level comparison:

| Mechanism | Vendor | Scope | Notes |
|---|---|---|---|
| vPC | Cisco NX-OS | Pair of switches, one logical LAG | 64 total ports verified per ACI guide; peer-link + keepalive design |
| VSX | Aruba AOS-CX | Active-active dual chassis | Live upgrades; required on Aruba DCN spines/leaves (Arch III) |
| MC-LAG / MLAG | Arista EOS, Cumulus | Active-active LAG across two switches | Arista MLAG + anycast VTEP in AVD examples |
| ESI LAG | EVPN standards | Multi-homing via Ethernet Segment Identifier | Juniper JVD validates ESI LAG with LACP; vendor-neutral |

[official/secondary — sources in Waves 8–12]

Design rule: dual-home every server (or accept the leaf as a failure domain); for storage/AI, prefer ESI or MLAG over single-homing; keep the multi-homing domain within one leaf pair to bound blast radius [secondary].

---

## Wave 20 — Deterministic cabling plan: worked example

For a 4-spine × 8-leaf fabric with 100G fabric links (example):

| Leaf | Spine 1 | Spine 2 | Spine 3 | Spine 4 |
|---|---|---|---|---|
| Leaf-1 | Et1 → Sp1-Et1 | Et2 → Sp2-Et1 | Et3 → Sp3-Et1 | Et4 → Sp4-Et1 |
| Leaf-2 | Et1 → Sp1-Et2 | Et2 → Sp2-Et2 | … | … |
| … | … | … | … | … |
| Leaf-8 | Et1 → Sp1-Et8 | Et2 → Sp2-Et8 | Et3 → Sp3-Et8 | Et4 → Sp4-Et8 |

Rule: **leaf-i uses port-group j toward spine-j; spine-j uses port i toward leaf-i** — fully deterministic, generatable from NetBox/Nautobot (see Phase E), printable as cable labels [analysis from standard practice — secondary]. For rail-optimized AI fabrics the rule becomes **NIC-N → Leaf-N** (Wave 18).

Cabling-plan artifacts to generate per fabric: port map, cable labels (both ends), optic/DAC type per link, and a test matrix (link → expected LLDP neighbor) for bring-up validation [secondary].

---

## Wave 21 — Final coverage audit & open items

**Scope-to-wave map (complete):** Clos/spine-leaf principles, oversubscription, scaling math (Waves 1, 7) · failure domains, cabling (Waves 1, 20) · 3-tier vs spine-leaf, migration (Wave 2) · Cisco Nexus/ACI (Waves 3, 8, 17) · Arista EOS/CloudVision/AVD (Waves 3, 9) · NVIDIA Spectrum/Spectrum-X/Cumulus/SONiC (Waves 3, 10) · Dell SONiC/SFS/OS10 (Waves 3, 11) · HPE Aruba CX (Waves 3, 11) · Juniper Apstra/QFX (Waves 3, 12) · whitebox/community SONiC (Waves 3, 15) · underlay essentials: numbered/unnumbered, RFC 5549, BFD, cabling plans (Waves 4, 10, 11, 20) · multi-homing (Wave 19) · Tomahawk 6 / G300 silicon radix (Waves 7, 13, 17) · AI backend topologies: rail-only/rail-optimized, ROD/RUD (Wave 18) · hyperscaler deployments (Waves 5, 14) · economics (Wave 15) · decision checklist (Wave 6).

**Provenance audit:** all factual claims carry tags; vendor scale/cost claims tagged [vendor-reported] or [secondary]; marketplace prices flagged indicative; analysis-labeled rows are derived, not sourced.

**Open items / [gap]s (final):** (1) Google Jupiter-class current fabric specifics; (2) Azure SONiC current fleet counts; (3) HPE/Juniper post-acquisition NOS roadmap; (4) independent 800G fabric benchmarks; (5) Dell Enterprise SONiC license pricing; (6) Cisco ACI per-release max leaf counts (APIC 5.2.x guides only); (7) xAI Colossus exact topology [unverified]; (8) Tomahawk 6 independent power/thermal figures; (9) CVE-2026-20212 remediation/patch coverage — hand to security phase; (10) Nexus One / Hyperfabric independent operational reviews.

*End of Phase D1 (21 waves). File is append-only; earlier sections untouched. Deep BGP underlay and EVPN-VXLAN protocol coverage continues in sibling Phase D files.*

---

