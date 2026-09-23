---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-22-vendor-reference-design-comparison-matrix-synthesis
title: "Wave 22 — Vendor reference-design comparison matrix (synthesis)"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["datacenter", "dci", "gpu", "gpus", "nvidia"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [529, 573]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 1018405307196f75764c70076fb89acaa60a4c9be862678eaae0415feec784e7
---

# Wave 22 — Vendor reference-design comparison matrix (synthesis)

## Wave 22 — Vendor reference-design comparison matrix (synthesis)

| Dimension | Cisco (NX-OS EVPN) | Cisco (ACI) | Arista (EOS/AVD) | NVIDIA (Cumulus/SONiC) | Dell (ENT SONiC / OS10+SFS) | Aruba CX | Juniper (Apstra) |
|---|---|---|---|---|---|---|---|
| Underlay default | eBGP numbered | IS-IS (fabric) | eBGP / RFC5549 unnumbered | eBGP unnumbered | eBGP | eBGP multi-AS | eBGP |
| Overlay | EVPN-VXLAN | VXLAN + OpFlex/COOP | EVPN-VXLAN | EVPN-VXLAN | EVPN-VXLAN | EVPN-VXLAN | EVPN-VXLAN |
| Controller required? | No | Yes (APIC) | No (CV optional) | No (NetQ optional) | No (SFS orchestrated) | No (Fabric Composer optional) | Yes (Apstra) |
| Automation model | Ansible/NDO | APIC API | AVD (Ansible data model) | NVUE/Ansible | SFS auto / Ansible | Ansible/FC | Intent blueprint |
| Multi-homing | vPC | vPC | MLAG | MLAG | VLT | VSX | ESI/MC-LAG |
| 800G spine (2026) | G300-based N9000 | G300-based N9000 | 7800R4 (576×800G) | Spectrum-X 800G | Z9664F-RT class | CX 9300 (32×400G) | QFX5230-64CD class |
| 3-stage/5-stage | ACI multi-pod | ACI multi-pod | AVD multi-stage | Super-spine | Manual | VSX pairs | JVD 5-stage |
| Best fit | Standards-based DC | Policy-governed DC | Cloud/AI at scale | Open/AI fabrics | Midmarket/private cloud | Campus-to-DC | Intent-managed DC |

[analysis — compiled from vendor docs cited in Waves 8–12, 17]

### 22.1 Arista 7368X4 modular leaf (2026)

The 7368X4 is a 4U, 8-slot modular leaf: 12.8 Tbps system, line cards spanning 25G/100G/400G (up to 128×100G or 32×400G per system), hot-swappable cards mixing legacy 10/25G breakouts and 400G spine uplinks in one chassis — aimed at large spine-leaf fabrics and AI-ready clusters with mixed server generations [secondary — https://www.ad-hoc-news.de/boerse/ueberblick/why-arista-s-7368x4-system-is-quietly-reshaping-data-center-leaf/69586724].

### 22.2 Hierarchical EVPN: Arista EVPN Gateway

Arista's EVPN Gateway whitepaper addresses flat-fabric EVPN state scale (MAC-IP routes, flood lists, next-hops across hundreds of leaves): divide the topology into self-contained EVPN domains with gateways rewriting EVPN routes and presenting themselves as next-hop, extended across DCs as DCI with VXLAN or MPLS encapsulation — standards-based (IETF) for multivendor interop [official — http://sit.www.arista.com/assets/data/pdf/Whitepapers/EVPN-Data-Center-EVPN-Gateway-for-Hierarchical-Multi-Domain-EVPN-and-DCI-WP.pdf]. AVD's dual-dc-l3ls example implements this pattern [official — https://github.com/aristanetworks/avd/blob/HEAD/ansible_collections/arista/avd/examples/dual-dc-l3ls/README.md].

### 22.3 Overlay peering best practice (Arista ATD lab)

Arista's own lab guides document the recommended split: **underlay peering on physical interfaces, overlay (EVPN) peering on loopbacks with eBGP-multihop** — spines act as EVPN route servers; BGP standard + extended communities enabled (EVPN uses extended communities for signaling) [official — https://github.com/aristanetworks/atd-public/blob/HEAD/topologies/dual-datacenter/labguides/source/l2l3evpn.rst].

---

## Wave 23 — Oversubscription planning worksheet (worked profiles)

| Profile | Leaf config | Downlink | Uplink | Ratio | Spine plan | Notes |
|---|---|---|---|---|---|---|
| AI/GPU backend | 32×800G | 16×800G (GPUs) | 16×800G | 1:1 | Rail-optimized, 1:1 | Non-blocking; separate frontend fabric |
| NVMe-oF storage | 48×25G + 8×100G | 1,200G | 800G | 1.5:1 | Lossless (PFC+ECN) | Arista/Pure validated: 7050SX3 leaves, 7050CX3 spines, BGP EVPN [official — https://www.purestorage.com/content/dam/pdf/en/white-papers/wp-arista-pure-deploying-nvme-of-enterprise-leaf-spine-architecture.pdf] |
| General virtualization | 48×25G + 6×100G | 1,200G | 600G | 2:1 | 3:1 max | Most common enterprise band |
| Web frontend / north-south | 48×25G + 4×100G | 1,200G | 400G | 3:1 | Oversubscribed OK | Cyclical traffic assumed |
| Campus aggregation | 48×1G + 4×25G | 48G | 100G | 0.48:1 | Effectively non-blocking | Access economics differ |

ECMP essentials for fabrics: hash on L3+L4 (include UDP source port for RoCEv2); default AVD ECMP paths = 4 — raise for large fabrics (16+ paths verified on Cisco ACI per-leaf); confirm symmetric hashing to avoid polarization [secondary/official].

BFD baselines seen in reference designs: Juniper JVD enables BFD on underlay and overlay eBGP; Cumulus/fabricator enables BFD on spine-leaf links; typical targets are sub-second detection (300ms × 3 is a common starting point, tune per platform) [official/secondary].

---

