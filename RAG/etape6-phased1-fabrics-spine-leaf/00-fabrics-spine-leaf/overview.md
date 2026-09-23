---
id: etape6-phased1-fabrics-spine-leaf/00-fabrics-spine-leaf/overview
title: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
domain: step-6-phase-d1-data-center-fabric-architectures-spine-leaf-
role: deep-dive
task: reference
actors: ["Broadcom", "Nvidia"]
dates: ["2026-09-22"]
keywords: ["benchmarks", "compute", "cost", "distribution", "gpu", "hyperscaler", "latency", "nvidia", "pricing", "research", "training"]
source: docs/RAG/etape6_phaseD1_fabrics_spine_leaf.md
source_anchor: ""
source_lines: [1, 68]
section: "Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs"
sha256: 5244350677e31e52c6c62995d91fc6eaa2a6ea9528e1c1ddc7e2ddf1a9753c1d
---

# Step 6 Phase D1 — Data-Center Fabric Architectures: Spine/Leaf, Clos, Vendor Reference Designs

**Scope:** Spine/leaf and Clos topology design principles, oversubscription ratios, scaling math, failure domains, cabling implications; core/distribution/access vs spine-leaf fit and migration; vendor reference designs (Cisco Nexus/ACI, Arista EOS/CloudVision, NVIDIA Spectrum/Spectrum-X + Cumulus/SONiC, Dell Enterprise SONiC/SmartFabric OS10, HPE Aruba CX, Juniper Apstra/QFX, whitebox/community SONiC); fabric-level underlay essentials (numbered vs unnumbered, cabling plans); real deployments with dates and costs where published. Deep BGP/EVPN-VXLAN protocol detail is covered in sibling Phase D files; this file keeps fabric-level essentials only.

**Date / currency:** Research current through 2026-09-22.

**Method:** Read-only web research via browser_search/browser_open. No live-browser visits, no sign-ins, nothing sent externally. Single writer; append-only waves.

**Provenance legend:** [official] = vendor/operator documentation or announcement · [vendor-reported] = vendor marketing/benchmarks not independently verified · [independent] = third-party testing or journalism · [secondary] = aggregators, blogs, analyst summaries · [unverified] = plausible but not confirmed — treat as lead, not fact.

**Open items / gaps:** tracked in the final section; flagged inline as **[gap]**.

---

## Wave 1 — Spine/leaf & Clos fundamentals (design principles, oversubscription, scaling math, failure domains, cabling)

### 1.1 What a Clos fabric is and why data centers converged on it

The leaf-spine architecture is a two-tier implementation of a Clos network, a multistage switching topology first described by Charles Clos in 1953 for telephone circuit switching [secondary]. The modern data-center form has two layers: leaf switches (top-of-rack, connecting servers/storage) and spine switches (interconnecting every leaf), with every leaf connected to every spine and no inter-leaf or inter-spine links in the basic design [secondary].

Key properties that drove adoption over classic three-tier (core/distribution/access) designs: predictable latency (every server-to-server path traverses exactly two hops: leaf → spine → leaf), full bisection bandwidth when built non-blocking, and horizontal scale-out by adding leaf/spine pairs rather than forklift chassis upgrades [secondary]. Equal-Cost Multi-Path (ECMP) routing across all leaf–spine links distributes load; flow-level hashing means a single flow uses one path but aggregate traffic spreads across all available paths [secondary].

The three-stage Clos (leaf–spine) suits a single pod; larger fabrics add a third tier (super-spine), making a five-stage Clos where each pod's spines connect to all super-spines — the standard hyperscaler pattern for scaling beyond one pod's radix limits [secondary].

### 1.2 Oversubscription ratios: what they mean and what operators use

Oversubscription ratio = total downlink (server-facing) capacity ÷ total uplink (fabric-facing) capacity on a leaf [secondary]. A 3:1 ratio means 300 Gbps of server ports share 100 Gbps of uplink — standard for general enterprise compute where not all servers burst simultaneously [secondary].

Typical published planning figures [secondary]:
- General enterprise virtualization: 3:1 to 5:1 leaf oversubscription.
- Storage-heavy or backup fabrics: often 1:1 (non-blocking) or 2:1.
- AI/GPU training fabrics: 1:1 non-blocking, frequently with a separate dedicated backend fabric for GPU-to-GPU traffic (e.g., NVIDIA's rail-optimized designs) distinct from the frontend management/storage fabric.
- Hyperscale web frontends: published examples range from 3:1 up toward 12:1 at the aggregation layers of very large fabrics, with traffic engineering compensating [unverified — ratio claims vary by source and era].

A 48×25G + 6×100G leaf (1,200 Gbps down, 600 Gbps up) is a textbook 2:1 design [secondary]. Modern 64×100G or 32×400G leaves follow the same arithmetic; the ratio, not the absolute speed, is the design variable.

### 1.3 Scaling math: radix, port counts, and maximum fabric size

Fabric scale is bounded by switch radix (port count). For a two-tier fabric with S spine switches of radix R_s and L leaf switches of radix R_l [secondary]:
- Each leaf uses S ports toward spines, leaving (R_l − S) server-facing ports.
- Number of spines S ≤ R_s (one link per leaf per spine needs a spine port per leaf; L ≤ R_s).
- Servers supported ≈ L × (R_l − S), with L ≤ R_s.

Example: 32-port spines (S=32) and 64-port leaves (R_l=64): up to 32 leaves × 32 server ports = 1,024 servers in a non-oversubscribed (1:1) fabric [secondary]. Oversubscription relaxes the per-leaf uplink count: with 4:1 oversubscription, a 64-port leaf uses ~13 uplink ports, supporting far more leaves per fabric but L is still capped by spine radix (each spine needs one port per leaf).

For 800G-era hardware: 51.2T chips (e.g., Broadcom Tomahawk 5) provide 64×800G; 102.4T chips (Tomahawk 6, announced) provide 128×800G or 256×400G [vendor-reported]. A single 64×800G spine tier with 64-port 800G leaves theoretically supports ~4,096 server ports non-blocking — the radix jump is what makes single-pod 8k-GPU fabrics plausible [secondary]. **[gap]** — exact shipping dates and independent power/thermal figures for Tomahawk 6-based systems were not confirmed in this wave.

### 1.4 Failure domains and resilience properties

A spine failure removes 1/S of fabric capacity, not connectivity: every leaf still reaches every other leaf via remaining spines, so the blast radius of one spine failure is capacity degradation, not partition [secondary]. A leaf failure affects only its directly attached servers — the smallest practical failure domain in the design [secondary].

Design consequences [secondary]:
- No single point of failure at the fabric layer without resorting to chassis redundancy; ECMP reconverges around a failed spine automatically.
- Maintenance (e.g., spine software upgrade) drains one spine at a time; fabrics are routinely run at N−1 or N−2 spine capacity.
- Link failures are absorbed by ECMP; the standard practice is to run fabrics with headroom (e.g., plan for ≤50–70% steady-state link utilization) so a spine loss does not congest survivors.
- Split-brain risk lives in the control plane (BGP sessions), not the topology — hence BFD (Bidirectional Forwarding Detection) on underlay links for sub-second failure detection, covered at protocol depth in the sibling BGP file.

### 1.5 Cabling implications of leaf-spine

Full-mesh leaf–spine cabling grows as L×S links — e.g., 32 leaves × 8 spines = 256 fabric links plus server links [secondary]. In-row/middle-of-row placement shortens DAC runs; end-of-row designs push fiber. Practical rules of thumb [secondary]:
- ≤3 m: passive DAC is cheapest (see Phase C for DAC/AOC pricing).
- 3–30 m: active copper (AEC) or AOC.
- >30 m / structured cabling: fiber with appropriate transceivers; MPO/MTP trunks for high-density spine interconnects (see Phase C).

Cabling plans for Clos fabrics are deterministic: leaf i connects port j to spine j (or a documented permutation), which makes automation and fault isolation straightforward and is why "cabling plan" artifacts are standard in vendor reference designs [secondary]. Structured-cabling vendors publish pre-terminated MPO trunk kits specifically dimensioned for leaf-spine counts [vendor-reported].

---

