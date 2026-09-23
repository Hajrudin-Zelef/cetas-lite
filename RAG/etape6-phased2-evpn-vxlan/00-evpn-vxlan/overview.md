---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/overview
title: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["benchmarks", "dci", "ethernet", "nvidia", "research", "training"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [1, 53]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: e55c8d24a5f5e8e5546e4cdbaae6f08b0c80a9e45e5de3e37d1d991eff02d341
---

# Step 6 — Phase D wave 2: EVPN-VXLAN overlay

**Scope:** VXLAN encapsulation fundamentals, EVPN control plane (route types 1–5+),
IRB symmetric/asymmetric, anycast gateway, EVPN multihoming (ESI, DF election,
split-horizon, aliasing), DCI (EVPN multi-site, OTV comparison), vendor
implementations and interop (Cisco Nexus/ACI, Arista EOS, NVIDIA Cumulus/SONiC,
Dell OS10/Enterprise SONiC, Aruba CX, Juniper), troubleshooting and operations.

**Research date:** 2026-09-22 (current through this date unless noted)

**Method:** read-only web research via search engine snippets and page fetches;
no live-browser visits, no sign-ins, nothing sent externally. No identifiers
guessed. Config examples are generic/sanitized patterns only (loopbacks in
192.0.2.0/24 documentation space, ASNs from documentation ranges).

**Provenance legend:** `[official]` = vendor documentation / IETF RFC / standards
body; `[vendor-reported]` = vendor blog, datasheet claim, or release note;
`[independent]` = third-party technical verification (labs, benchmarks, analysts);
`[secondary]` = community wiki, training blogs, forum consensus; `[unverified]` =
single-source or unconfirmed claim. Gaps and conflicts are flagged explicitly.

**Related files:** `etape6_phaseD1_clos_design.md` (wave 1: spine/leaf, BGP, MC-LAG),
`etape6_phaseD3_security_segmentation.md` (wave 3: security/segmentation) —
see ROADMAP.md.

---

## Wave 1 — VXLAN encapsulation fundamentals

### 1.1 What VXLAN is and why it exists

VXLAN (Virtual Extensible LAN), specified in RFC 7348 [official], encapsulates
complete Layer 2 Ethernet frames inside Layer 3 UDP packets so virtual L2
segments can ride across a standard IP underlay. The problems it solves, in
order of pain for data-center operators [secondary][independent]:

- 12-bit VLAN IDs cap isolation at 4094 segments — far too few for multi-tenant
  clouds with thousands of tenants.
- Spanning tree blocks redundant links in large L2 domains, wasting half the
  fabric bandwidth operators paid for.
- MAC tables overflow when every leaf learns every host in one big L2 domain.
- VM/workload migration needs L2 adjacency across L3 boundaries.

VXLAN answers with a 24-bit VXLAN Network Identifier (VNI): 16,777,216 possible
segments (not a round 16 million) [official: RFC 7348 §3]. An overlay lets the
underlay be a fully routed network where every link forwards — no STP blocking
[secondary].

Key negative fact operators must internalize: **VXLAN encapsulates; it does not
encrypt.** A capture on the underlay shows the VXLAN header and the inner frame
in cleartext. If confidentiality is required, the underlay must be trusted or
separately encrypted (e.g., IPsec/WireGuard under the tunnels) [independent].

