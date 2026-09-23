---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/wave-4-fabric-level-underlay-essentials-numbered-vs-unnumber
title: "Wave 4 — Fabric-level underlay essentials (numbered vs unnumbered, cabling plans)"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Google", "Meta", "Microsoft", "Nvidia", "xAI"]
dates: []
keywords: ["benchmarks", "cost", "ethernet", "gpu", "hyperscaler", "license", "nvidia", "pricing"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [158, 203]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 8f9f4271b3cd6fcd313e08aaa6f989e94171b91bcfaf6ea072d3c1507efd42a2
---

# Wave 4 — Fabric-level underlay essentials (numbered vs unnumbered, cabling plans)

## Wave 4 — Fabric-level underlay essentials (numbered vs unnumbered, cabling plans)

Deep protocol mechanics (BGP path selection, EVPN route types, VXLAN encapsulation) belong to sibling Phase D files. Fabric-level essentials:

- **Numbered vs unnumbered underlay links:** Numbered = /31 or /30 IP per point-to-point link (more addresses consumed, easier traceroute/snmp per-link). Unnumbered = no IP on the link; BGP sessions form via IPv6 link-local (RFC 5549) with IPv4 routes exchanged over the v6 session — the modern default in Arista AVD, NVIDIA Cumulus, and SONiC reference designs because it eliminates underlay IPAM for fabric links [secondary]. Trade-off: unnumbered complicates some monitoring tools that expect per-interface IPs [independent].
- **Loopback-based peering:** BGP sessions are typically established between loopbacks (not interface addresses) for stability across link flaps; eBGP multihop or directly-connected with loopback next-hop handling per vendor design [secondary].
- **BFD:** sub-second liveness on underlay links is standard practice in reference designs (Arista, Cisco, Cumulus all document BFD with BGP for fabrics) [official].
- **Cabling plans:** deterministic leaf-i↔spine-j mappings published per fabric; automation generates both config and cable labels from the same source of truth (NetBox/Nautobot — see Phase E) [secondary].

---

## Wave 5 — Real deployments: scale examples, references, costs

- **Microsoft Azure / SONiC:** the original hyperscale SONiC deployment; Microsoft has publicly discussed managing one of the world's largest SONiC fleets across Azure regions (tens of thousands of devices) [secondary]. Specific current port counts are not officially published — **[gap]**.
- **Meta:** public engineering posts describe multi-tier Clos fabrics (FBOSS/Aggregator generations: F16/Rack, HPR) scaling to tens of thousands of servers per fabric region; Meta's AI fabrics use rail-optimized designs with dedicated backend networks [secondary].
- **Google:** Jupiter/Aurora generations of in-house Clos fabrics; published papers describe 1+ Pb/s bisection fabrics [secondary]. Dates: Jupiter-era publications ~2015; current-generation details are proprietary — **[gap]**.
- **xAI Colossus (Memphis):** publicly reported as ~100k+ GPU cluster built in 2024 with Ethernet-based backend networking (NVIDIA Spectrum-X) per industry reporting [secondary]. Exact fabric topology not officially detailed — **[unverified]** on specifics.
- **Enterprise reference scales:** vendor CVDs routinely validate 100–200 leaf nodes per fabric (Cisco ACI, Arista AVD examples); production enterprise fabrics of 20–60 leaves are the common published case-study band [vendor-reported][secondary].
- **Costs:** public per-fabric cost figures are rare. Data points: whitebox 32×400G leaf switches list in the low tens of thousands USD (vendor/retailer pricing, 2025–2026) vs branded equivalents often 2–4× higher [independent][secondary]. Full-fabric TCO comparisons are vendor-collateral territory — **[gap]** flagged; treat any single-vendor TCO claim as [vendor-reported].

---

## Wave 6 — Design decision checklist (synthesis)

1. Traffic profile first: east-west ratio determines whether leaf-spine is justified; measure before designing [independent].
2. Pick oversubscription per tier from workload (1:1 AI/GPU backend; 2–3:1 general virtualization; higher for north-south-heavy) — not from vendor defaults [independent].
3. Choose underlay style (unnumbered eBGP is the current community default) and NOS before hardware; NOS dictates automation tooling [secondary].
4. Size spine radix for 3–5 year leaf growth; spines are the hardest tier to replace [secondary].
5. Separate fabrics for AI backend vs frontend where GPU collectives exist; converged fabrics compromise both [vendor-reported][secondary].
6. Plan N−1 spine capacity and ≤70% steady-state link utilization [secondary].
7. Generate cabling plans and configs from one source of truth (see Phase E) — hand-built fabrics drift [independent].

---

## Coverage audit & open items

**Scope-to-wave map:** Clos/spine-leaf principles (Wave 1) · oversubscription & scaling math (Wave 1) · failure domains & cabling (Wave 1) · 3-tier vs spine-leaf fit & migration (Wave 2) · Cisco/Arista/NVIDIA/Dell/Aruba/Juniper/whitebox reference designs (Wave 3) · underlay essentials (Wave 4) · real deployments & costs (Wave 5) · decision checklist (Wave 6).

**Provenance audit:** all factual claims above carry tags; vendor scale/cost claims are tagged [vendor-reported] or [secondary] as appropriate; nothing in this file is presented as independently benchmarked unless tagged [independent].

**Open items / [gap]s carried forward:** (1) Tomahawk 6 independent power/thermal/shipping data; (2) hyperscaler current-generation fabric specifics (Azure/Jupiter-class counts); (3) Dell Enterprise SONiC license pricing; (4) HPE/Juniper NOS consolidation roadmap; (5) independent 800G fabric benchmarks; (6) published fabric TCO comparisons; (7) Cisco ACI current-version max leaf counts vs release notes.

*End of Phase D1. Deep BGP underlay/EVPN-VXLAN protocol coverage continues in sibling Phase D files.*

---

