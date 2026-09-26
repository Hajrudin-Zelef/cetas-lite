---
id: etape6-phased3-bgp-ha/00-bgp-ha/7-4-layer-3-over-mc-lag-pairs
title: "7.4 Layer 3 over MC-LAG pairs"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["benchmarks"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [403, 426]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: e9824c2f498b2e9aa54db3e256c214bb78065567a204625b47c9b93f99b14645
---

# 7.4 Layer 3 over MC-LAG pairs

- MC-LAG pairs present **one logical bridge** to STP: both peers share the LACP system ID and typically synchronize STP state over the inter-switch link (Aruba VSX syncs STP states via ISLP; Dell VLT processes BPDUs on the primary) [official].
- Best practice: the MC-LAG pair should be STP root for its attached L2 domains; peer-link/ISL ports are usually STP edge or excluded to avoid blocking the inter-switch link [independent].
- Dell VLT: RSTP/RPVST+ supported, MSTP not supported on the documented OS10EE train; enable STP before VLT [official].
- In EVPN fabrics, STP is often **disabled on fabric links entirely** (L3 routed underlay has no loops) and confined to server-facing edge ports [independent].

### 7.4 Layer 3 over MC-LAG pairs

- Running L3 (SVIs, routing adjacencies) on MC-LAG peers requires care: both peers must make consistent forwarding decisions, or traffic trombones over the peer link [independent].
- **Per-vendor answers:**
  - **Cisco vPC:** peer-gateway + HSRP active/active for FHRP; L3 peer-router features for routing over vPC [official].
  - **Arista MLAG:** VARP for anycast gateway; L3 peering over MLAG links supported with consistent config [official].
  - **Dell VLT:** VLT peer routing — both peers route for the VLT VLAN, avoiding peer-link tromboning [official — Dell VLT guides].
  - **Aruba VSX:** active-gateway; L3 databases built independently per switch [official].
  - **Juniper MC-LAG:** VRRP with active/active via ICCP coordination; IRB over MC-AE [official].
- **Gap:** detailed L3-over-MC-LAG convergence numbers per vendor not found in public docs.

### 7.5 Wave-7 verification notes and open items

- **Verified:** EVPN-MH procedures (DF election, split-horizon, aliasing); Netris recommendation; STP sync behavior; per-vendor L3-over-MC-LAG features.
- **Open:** L3-over-MC-LAG convergence benchmarks.
- *End of Wave 7.*

---

