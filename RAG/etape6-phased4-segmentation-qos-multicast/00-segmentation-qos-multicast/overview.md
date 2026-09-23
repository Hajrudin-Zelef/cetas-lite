---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/overview
title: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["ethernet", "research"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [1, 45]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: f7175b88f9955a03270fc4c9a45367db43b6770e49b112f0e2158595b70d049a
---

# Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services

**Scope:** Data-center fabric segmentation (VLAN design at scale, VRF, microsegmentation), QoS/DCB (classification, PFC/ETS/DCBX, ECN, RoCEv2 lossless tuning), multicast (PIM-SM/SSM, RP design, VXLAN overlay multicast, IGMP), and fabric network services (DHCP relay, DNS, NTP/PTP).
**Date covered through:** 2026-09-22
**Method:** Read-only web research via browser_search / browser_open. No live-browser visits, no purchases, no form submissions. Nothing sent externally.
**Provenance legend:** `[official]` = vendor/official documentation or datasheet · `[vendor-reported]` = vendor blog/whitepaper/marketing claim · `[independent]` = independent testing/review (e.g., ServeTheHome, Packet Pushers) · `[secondary]` = press/news/aggregators · `[unverified]` = cannot be corroborated. Every factual claim below carries a tag. Figures are snapshots as of the stated date, not guarantees.
**Editorial rule:** This file is append-only; sections are added in waves. If a later wave corrects an earlier claim, the correction is noted in place with date rather than silently edited.
**Note on non-comparables:** Vendor QoS/DCB defaults, buffer sizes, and "lossless" claims are NOT directly comparable across vendors; they are recorded per-vendor with their own test conditions.

---

## Verification log (gaps / conflicts / unverified)

*(Filled at the end of each wave and consolidated at file completion.)*

### Wave 1 — VLAN design at scale

## Wave 1 — VLAN design at scale

### 1.1 The 4094-segment ceiling
- IEEE 802.1Q VLAN IDs are 12-bit; usable range 1–4094 (VLAN 0 and 4095 reserved) `[official]` — source: Pluribus VXLAN doc, crawled 2026-09-22; Cisco Q-in-VNI deployment guide. https://techdocassets.pluribusnetworks.com/netvisor/nv1_521/UnderstandingVXLAN-basedBridgeDo.html · https://www.cisco.com/c/en/us/td/docs/dcn/whitepapers/q-in-vni-over-vxlan-fabric-deployment-guide.pdf
- The ceiling becomes a blocker in multi-tenant DCs because each tenant needs independently assignable VLAN IDs and MAC domains without conflicts `[secondary]` — same sources; MAC table pressure compounds it (N VMs × M servers can exceed hardware MAC tables, per Pluribus).
- Pre-VXLAN mitigation: 802.1ad/802.1QinQ (VLAN stacking) multiplies the space to 4094×4094 ≈ 16M hierarchical identifiers — e.g., outer tag = customer, inner tag = customer service — usable in DC/cloud-exchange handoffs `[vendor-reported]` — Pluribus; Cisco Catalyst 9500 EVPN VXLAN config guide documents Q-in-VNI (S-VLAN carrying multiple C-VLANs) for carrier/L2-transparent designs `[official]`. https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst9500/software/release/17-15/configuration_guide/vxlan/b_1715_bgp_evpn_vxlan_9500_cg/configuring_layer2_overlay_with_Q-in-VNI.html
- Limitation noted by Cisco: overlapping MAC entries between multiple C-VLANs under a single EVPN-mapped S-VLAN is not supported `[official]` — same Catalyst 9500 guide.

### 1.2 VXLAN VNI as the scale answer
- VXLAN (RFC 7348) encapsulates L2 Ethernet in L3 UDP (destination port 4789) via VTEPs; the VNI is 24-bit → ~16.7M logical segments `[official]` — RFC 7348; Cisco BGP EVPN VXLAN overview (cisco.com, ie9300 config guide). https://www.cisco.com/c/en/us/td/docs/switches/lan/cisco_ie9300/software/Configuration_Guide/BGP-EVPN-VXLAN/b-bgp-evpn-vxlan-cg/bgp_evpn_vxlan_overview.pdf
- VXLAN adds ~50 bytes of encapsulation overhead; underlay should support jumbo frames (MTU 9000+) or path-MTU discovery must be handled `[secondary]` — darnodo notebook (GitHub, updated 26 days before 2026-09-22). https://github.com/darnodo/notebook/blob/HEAD/content/documentation/VXLAN/beginners/vxlan-for-beginners.en.md

### 1.3 VLAN↔VNI mapping conventions
- VLAN ID is 12-bit and locally significant on the wire; VNI is 24-bit and globally significant in the overlay; config maps local VLANs to global VNIs (e.g., `vxlan vlan 100 vni 10100`) `[secondary]` — alukacs03 containerlab (updated ~2026-07). https://github.com/alukacs03/clauntainerlab/blob/HEAD/labs/29-vxlan-data-plane/README.md
- Common human-readable convention: VNI = VLAN × 100 (VLAN 100 → VNI 10100) `[secondary]` — same source. Deterministic mapping (VNI = 10000 + VLAN) also common; no standard mandates either — convention only, flag as `[unverified]` as a universal rule.
- NetPilot multi-tenant EVPN example (2026-09): tenant VRFs get unique route-targets + L3 VNIs (50100/50200/50300), each tenant carrying two L2 VNIs (10101/10102 etc.) with symmetric IRB `[secondary]` — netpilot-labs example prompts. https://github.com/netpilot-labs/example-prompts/blob/HEAD/data-center/multi-tenant-vrfs-evpn.md

### 1.4 Segmentation strategy patterns
- Per-tenant: one VRF + L3 VNI per tenant, strict route-target separation prevents cross-tenant leakage; shared-services VRF added later for DNS/DHCP/SIEM `[secondary]` — NetPilot pattern; Cisco ACI docs describe equivalent (tenant VRFs + shared-services tenant) `[official]`. https://www.ciscopress.com/articles/article.asp?p=2928191&seqNum=4
- Per-application/per-environment: separate L2 VNIs (or VLANs in smaller fabrics) per app tier and per environment (prod/staging/dev) — widely documented design pattern, exact VNI numbering schemes are operator conventions `[secondary]` — multiple design guides; treat specific numbering claims as operator-specific.
- Hybrid cloud: QinQ + VXLAN combined (Pluribus Adaptive Cloud Fabric) for double-tagged handoffs to cloud-exchange providers `[vendor-reported]` — Pluribus.
- Gap: no public, independently verified dataset was found on how large enterprises actually allocate their 16M VNI space in practice (sparse vs dense allocation) — recorded as open item `[unverified]`.

### Wave 1 verification
- Sources: 8 (Cisco ×3 official, Pluribus ×2 vendor-reported, GitHub ×3 secondary). No conflicts found. One gap logged (VNI allocation practice in the wild).

---

