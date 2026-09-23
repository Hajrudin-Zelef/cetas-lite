---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-8-multicast-in-vxlan-overlays-igmp
title: "Wave 8 — Multicast in VXLAN overlays + IGMP"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [207, 256]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 39d60ab92a11501a111d688d170d165e00e279ec3cfec99f735e1d792f06558a
---

# Wave 8 — Multicast in VXLAN overlays + IGMP

## Wave 8 — Multicast in VXLAN overlays + IGMP

### 8.1 BUM handling in VXLAN: two underlay options
- Each VNI maps to a multicast group; BUM (broadcast/unknown-unicast/multicast) in the VNI is encapsulated with that group as the outer destination IP and replicated by the underlay `[secondary]` — Cisco community VXLAN/EVPN design notes (sheynshield repo). https://github.com/shayan-heydarikhah/sheynshield/blob/HEAD/Cisco/Design/VXLAN_EVPN.md
- Option A — underlay multicast: VTEPs join per-VNI groups; replication done by network hardware; more scalable in large fabrics; requires multicast-capable underlay `[secondary]` — Cisco community multi-site deployment doc. https://community.cisco.com/kxiwq67737/attachments/kxiwq67737/discussions-tetration-analytics/735/1/VXLAN%20EVPN%20Multi-Sites%20Deployment.pdf?inline=true
- Option B — ingress (head-end) replication: source VTEP unicasts one copy per remote VTEP; no underlay multicast needed; simpler to deploy but bandwidth-inefficient as VTEP count grows; remote VTEPs statically configured in flood-and-learn mode, auto-discovered via EVPN Type-3 (IMET) routes in EVPN mode `[secondary]` — same sources; alukacs03 containerlab lab 29. https://github.com/alukacs03/clauntainerlab/blob/HEAD/labs/29-vxlan-data-plane/README.md
- Operational guidance `[official]` — Cisco TRM whitepaper (Nexus 9000): do NOT mix — underlay multicast for some L2 VNIs + ingress replication for others is "not advisable" (unpredictable forwarding, operational complexity); pick one BUM technique fabric-wide; migrating IR→multicast underlay later is traffic-impacting (different forwarding databases must be reprogrammed). If overlay multicast routing will ever be needed, deploy underlay multicast from day one. http://www.cisco.com/c/en/us/td/docs/dcn/whitepapers/tenant-routed-multicast-in-nexus9000-vxlan-bgp-evpn-fabrics.html
- Assisted Replication (newer EVPN mechanism): ingress PE sends one copy to designated replicators (often spines) that fan out — reduces ingress-PE CPU/bandwidth burden in dense overlays `[secondary]` — dn.org EVPN multicast analysis. https://dn.org/multicast-in-evpn-networks-router-cast-vs-assisted-replication/?utm_source=rss&utm_medium=rss&utm_campaign=multicast-in-evpn-networks-router-cast-vs-assisted-replication

### 8.2 Tenant Routed Multicast (TRM)
- Without TRM, overlay multicast rides as BUM in the underlay and cannot cross subnets; TRM (BGP-EVPN based) routes multicast between sources/receivers on different subnets/VTEPs `[official]` — Cisco BGP EVPN VXLAN overview. https://www.cisco.com/c/en/us/td/docs/switches/lan/cisco_ie9300/software/Configuration_Guide/BGP-EVPN-VXLAN/b-bgp-evpn-vxlan-cg/bgp_evpn_vxlan_overview.pdf
- TRM requirements per Cisco (Nexus 9000) `[official]`: IPv4 multicast underlay with PIM-SM (SSM and BiDir NOT supported in the underlay; ingress replication in the underlay NOT supported); overlay supports PIM-SSM and PIM-SM (BiDir not supported); RP for the overlay must be reachable inside the tenant VRF; 224.0.0.0/24 link-local block is bridged, not routed; check source TTL settings so traffic is not dropped at FHR/LHR. http://www.cisco.com/c/en/us/td/docs/dcn/whitepapers/tenant-routed-multicast-in-nexus9000-vxlan-bgp-evpn-fabrics.html
- Two multicast domains exist in a TRM fabric: underlay (default VRF) and overlay (tenant VRF) — each needs its own RP `[official]` — same whitepaper.
- Arista EOS 4.36.2F documents EVPN multicast VRF leaking (Sept 2026) — egress PE SMET signaling, PEG-as-RP option, MSDP sync requirement `[official]` — recorded in Wave 2, cross-reference. https://www.arista.com/en/um-eos/eos-evpn-multicast-vrf-leaking?searchword=eos%20evpn

### 8.3 IGMP snooping & querier
- IGMP snooping constrains L2 multicast flooding to ports with actual receivers; without it, multicast is flooded like broadcast within the VLAN `[secondary]` — standard references.
- In EVPN fabrics, IGMP snooping on VTEPs interacts with EVPN Type-6/7/8 routes (selective multicast Ethernet tag / join sync) — EVPN can carry IGMP state in BGP so remote VTEPs only send group traffic where receivers exist `[secondary]` — EVPN multicast design literature.
- Querier: one device per VLAN must send IGMP general queries; in anycast-gateway fabrics the querier role must be deterministic (else duplicate queries / state churn) — operator design point; per-vendor knobs differ `[secondary]`.
- Gap: per-vendor IGMP-snooping/EVPN Type-6-8 support matrix for 2026 platforms not compiled — open item `[unverified]`.

### Wave 8 verification
- Sources: 7. Cisco's explicit "don't mix BUM modes" guidance captured verbatim. TRM protocol-support matrix (SM/SSM yes, BiDir no) recorded as platform-specific (Nexus 9000), not universal.

---

## Wave 9 — Network services: DHCP relay, DNS, NTP/PTP

### 9.1 DHCP relay in EVPN fabrics
- Problem: in EVPN fabrics with distributed anycast gateways, a DHCP Discover flooded across the stretched L2 is relayed by EVERY DAG-enabled leaf → DHCP server overload, wrong-DAG route installs, outages `[official]` — Cisco IOS XE Segment Routing/EVPN-IRB guides (NCS 4200, ASR 920): prior to IOS XE 17.5.1 DHCP relay was not supported on EVPN; from 17.6.1 specialized handling supports DHCPv4/v6 with DAGs on the same VRF, and the FIRST-HOP device handles DHCP to avoid fabric-wide floods. https://www.cisco.com/c/en/us/td/docs/routers/ncs4200/configuration/guide/segment-routing/17-1-1/b-segment-routing-17-1-ncs4200/m-evpn-irb.html
- Correct design (NX-OS, documented live on Nexus 9000v NX-OS 10.5.3) `[secondary]` — enizaksoy lab (2026-07): relay configured on EVERY leaf carrying the VLAN, with a per-leaf unique loopback as relay source-interface (advertised in the VRF) so replies route deterministically; `ip dhcp relay information option` (option 82) + `information option vpn` (link-selection / VRF-aware server allocation); centralized server reachable via VRF-aware relay. https://github.com/enizaksoy/cisco-vxlan-evpn-multi-as-lab/blob/HEAD/docs/dhcp-relay-deep-dive.md · config pattern: https://deliabtech.com/blogs/data-center/configure-dhcp-in-vxlan-evpn-fabric/
- Security: every relayed Discover is punted to CPU (giaddr rewrite is software) — a DoS vector; layered defense: edge `storm-control broadcast`, DHCP snooping on untrusted ports, `ip dhcp snooping verify mac-address`, CoPP `strict` profile with DHCP policers `[secondary]` — same lab.
- DHCP snooping role: drops rogue OFFER/ACK from hosts; the anycast-gateway MAC (0000.1111.2222 on NX-OS) identifies gateway-originated packets in captures `[secondary]` — same lab.

### 9.2 DNS in the DC
- Shared-services pattern: DNS (like DHCP/AD/SIEM) lives in a shared-services VRF/tenant, reached from tenant VRFs via selective route leaking or ACI "shared subnet" constructs `[official]/[secondary]` — Cisco Press ACI; NetPilot pattern (Waves 1–2).
- Design points: anycast DNS (same IP on multiple servers, routed to nearest) for resilience; DNS64/NAT64 only in v6-transition designs; keep DNS in the underlay-independent services block so fabric re-convergence does not orphan name resolution `[secondary]` — standard practice; no single authoritative vendor doc found — recorded as practice, not spec.
- Gap: no current public data on DNS-anycast deployment prevalence in enterprise DCs — `[unverified]`.

### 9.3 NTP/PTP timing
- NTP: hierarchical stratum model, ms-level accuracy — sufficient for logging/correlation; every fabric device + hypervisor should sync to redundant stratum-1/2 sources `[secondary]` — standard practice.
- PTP (IEEE 1588): sub-microsecond sync via boundary clocks / transparent clocks on switches; required for financial trading (MiFID II timestamping), 5G fronthaul, and some AI-collective telemetry `[secondary]` — widely documented; Arista EOS and NX-OS both document PTP boundary/transparent clock modes `[official]` (per-vendor docs, not deep-linked here).
- DC relevance: PTP in the underlay for telemetry correlation; NTP everywhere as baseline; never mix unauthenticated NTP across security zones without NTS/authentication where required `[secondary]`.
- Gap: per-platform PTP profile support (G.8275.1 vs default E2E/P2P) matrix not compiled — open item `[unverified]`.

### Wave 9 verification
- Sources: 5. Cisco 17.5.1/17.6.1 DHCP-relay version boundary captured verbatim. DNS/PTP sections are practice-level; flagged where no authoritative doc was consulted.

---

