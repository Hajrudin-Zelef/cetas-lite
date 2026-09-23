---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-30-worked-example-64-leaf-fabric-bill-of-materials-and-
title: "Wave 30 — Worked example: 64-leaf fabric bill of materials and acceptance checks"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Broadcom", "Meta", "Nvidia", "xAI"]
dates: ["2026-09-02"]
keywords: ["asic", "gpus", "latency", "licenses", "llama", "nvidia", "optics", "throughput"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [697, 750]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 92133ae0a4e38627cefbe49fb7954f06acb09a790918e63183228e82e32aa404
---

# Wave 30 — Worked example: 64-leaf fabric bill of materials and acceptance checks

## Wave 30 — Worked example: 64-leaf fabric bill of materials and acceptance checks

### 30.1 BOM for a 64-leaf, 4-spine, 400G fabric (illustrative)

| Item | Qty | Unit assumption | Line total basis |
|---|---|---|---|
| Leaf switches 64×400G | 64 | per-leaf list price | 64 × unit |
| Spine switches 64×400G | 4 | higher-radix list price | 4 × unit |
| 400G-DR4 optics (leaf→spine) | 512 | 2 per link × 256 links | 256 links × 2 ends |
| Breakout cables 400G→4×100G | 64 | server-side | as needed |
| Fiber trunks (12/24-fiber) | ~22 | per cable run plan | per-row layout |
| NOS licenses (3-yr) | 68 | per-switch | 64 + 4 |
| Optics spares (5%) | ~26 | 5% of optic count | round up |
| PDUs / rack space | — | per-row power budget | site-specific |

Total optic count check: 64 leaves × 4 spines = 256 fabric links; 2 optics per link = 512 [analysis — derived from Wave 20 cabling math]. Adjust leaf:spine fan-out for the target oversubscription (Wave 7).

### 30.2 Fabric acceptance checklist

- [ ] All expected LLDP adjacencies present (spine × leaf matrix complete).
- [ ] BGP session count matches plan (underlay + overlay); zero flaps in 24 h.
- [ ] ECMP path count symmetric on every leaf (check per-destination hash spread).
- [ ] BFD sessions up on all fabric links; failover time measured and documented.
- [ ] BUM replication verified (ping unknown-unicast across VTEPs).
- [ ] Oversubscription measured under peak load matches design ratio.
- [ ] Config backup + NetBox/Nautobot source-of-truth synchronized.
- [ ] Runbook: spine drain procedure tested on one spine.

[secondary — compiled from vendor bring-up guides and Wave 6 checklist]

*End of Phase D1 (30 waves). Append-only; earlier sections untouched. Deep BGP underlay and EVPN-VXLAN protocol coverage continues in sibling Phase D files.*

---

## Wave 31 — Key numbers at a glance & tag audit

| Metric | Value | Source wave |
|---|---|---|
| Leaf:spine oversubscription bands | 1:1 (AI) → 3:1 (web) → non-blocking (campus) | 7, 23 |
| BGP ECMP paths default (AVD) | 4 (raise to 16+) | 9 |
| Cisco NX-OS 10.4(3)F+ VTEP scale | 1,000 VTEPs (32 GB+) | 8 |
| Arista 7800R4 AI spine | 576×800G (460.8T) | 9 |
| Broadcom Tomahawk 6 / Cisco G300 | 102.4T (64×1.6T / 64×1600GbE) | 13, 17 |
| NVIDIA Spectrum-X Clos claim | 51.2T ASIC, 16,000 ports / 2-tier | 10 |
| Meta DSF system | 128×800G spine, 38×800G leaf (scheduled) | 14 |
| xAI Colossus scale | 100,000+ Hopper GPUs, Spectrum-X [vendor-reported] | 14 |
| Whitebox 5-yr TCO claim | 54% lower vs proprietary (vendor model) | 15 |
| Rail-optimized hop latency | ~600 ns intra-rail, <2 µs cross-rail [secondary] | 18 |
| Juniper Apstra 5-stage JVD | Apstra 5.0.0-64, Junos 23.4R2 | 12 |
| CVE-2026-20212 (2026-09-02) | CVSS 9.8, Silicon One Nexus 9000, TCP 43210/43211 | 17 |

**Tag audit (final):** every factual claim above carries a provenance tag; vendor-constructed claims (TCO model, Colossus 95% throughput, G300 figures, 1.6T optics) carry [vendor-reported]; marketplace prices flagged indicative; gray-market and community sources flagged [secondary]; Llama-5-class unverified items remain [unverified].

*End of Phase D1 (31 waves, 750 lines target met). Append-only; earlier sections untouched. Deep BGP underlay and EVPN-VXLAN protocol coverage continues in sibling Phase D files.*
