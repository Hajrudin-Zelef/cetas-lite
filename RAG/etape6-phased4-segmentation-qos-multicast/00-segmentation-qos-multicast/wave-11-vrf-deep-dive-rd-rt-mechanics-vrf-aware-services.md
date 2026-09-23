---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-11-vrf-deep-dive-rd-rt-mechanics-vrf-aware-services
title: "Wave 11 — VRF deep-dive: RD/RT mechanics, VRF-aware services"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [331, 379]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 7e963d801a240780a4c5b5c88cf47e1f53e7f54a9a1963da70eeef6d12731fb4
---

# Wave 11 — VRF deep-dive: RD/RT mechanics, VRF-aware services

## Wave 11 — VRF deep-dive: RD/RT mechanics, VRF-aware services

### 11.1 VRF-lite vs MPLS VPN (RFC 4364)
- VRF-lite: per-interface VRF assignment on one box, no MPLS labels; the DC-fabric norm for ToR/leaf multi-tenancy `[secondary]` — netlab docs (Wave 2).
- RFC 4364 BGP/MPLS IP VPNs: PE routers hold per-VRF RIBs; MP-BGP VPNv4/VPNv6 carries RD+RT+labeled prefixes between PEs — the SP/WAN heritage that EVPN's RT mechanics descend from `[official]` — RFC 4364 (IETF).
- EVPN-VXLAN replaces the MPLS label with the VNI but keeps RD/RT semantics — which is why RT design transfers directly between the two worlds `[secondary]` — design literature.

### 11.2 Leaking patterns in detail
- Static leaking: `ip route vrf A <prefix> <next-hop> vrf B` style — simple, no protocol, scales poorly `[secondary]`.
- BGP RT leaking: export RT of VRF A imported by VRF B (route-target import/export with route-maps to filter) — the scalable pattern; Cisco NX-OS caps at 1000 leaked prefixes by default and blocks leaking INTO the default VRF `[official]` — Cisco NX-OS L3 virtualization guide (Wave 2).
- Shared-services VRF: tenants leak only the service prefixes (DNS/DHCP/AD resolvers), services VRF imports nothing back except return paths as needed — least-privilege leaking `[secondary]` — NetPilot; Cisco Press ACI.
- Internet-egress VRF: single shared VRF on border leaves, tenants leak a default or selected prefixes toward it; DSVNI/asymmetric-IRB evidence exists on NX-OS 10.5.3 `[secondary]` — enizaksoy lab (Wave 2).
- VRF-select / policy-based VRF selection: classify ingress traffic (ACL/DSCP) into a VRF — used for guest/partner segmentation without extra interfaces `[secondary]` — vendor docs (Cisco VRF Select).
- Operational warning: leaking is a deliberate hole in tenant isolation — every leaked prefix should be justified, filtered by route-map, and logged in the design doc; audit leaked-prefix counts in monitoring `[secondary]` — operator practice.

### 11.3 VRF-aware services
- DHCP relay, DNS, NTP, SNMP, syslog, TACACS/RADIUS, and management must each be VRF-aware (source-interface in the right VRF, server reachability via leaking or shared VRF) — a common source of "works in default VRF, broken in tenant VRF" outages `[secondary]` — operator practice; DHCP deep-dive in Wave 9.
- netlab platform matrix reminder (2026-09): VRF-aware loopbacks and leaking supported on Arista EOS, Aruba AOS-CX, Cisco IOS/XE/XR/NX-OS, Cumulus, Dell OS10, FRR, Junos, RouterOS 7, SR Linux, VyOS — with per-platform caveats `[secondary]` — ipspace netlab (Wave 2).

### Wave 11 verification
- Sources: 5 + RFC 4364. No new conflicts. Leaking presented as exception-to-isolation with audit guidance.

---

## Wave 12 — QoS deep-dive: MQC walkthrough, Nexus system classes, Arista model

### 12.1 Modular QoS CLI (MQC) structure
- Three policy-map types on NX-OS: `type qos` (classification/marking), `type queuing` (scheduling/bandwidth/WRED), `type network-qos` (PFC enable, MTU per class) — applied via `service-policy` input/output/system `[official]` — Cisco NX-OS QoS guides.
- Class-maps match on ACL, DSCP, CoS, IP precedence, or packet-length; `match-any`/`match-all` semantics `[official]` — Cisco.
- Trust boundary principle: trust DSCP on inter-switch/fabric links; mark/classify at the access edge; never trust host markings inside a multi-tenant DC without validation `[secondary]` — design practice.

### 12.2 Nexus system classes (structural)
- NX-OS defines QoS system classes with default behaviors; the `network-qos` class carrying storage/RDMA traffic gets PFC enabled and jumbo MTU; the default class stays lossy `[official]` — Cisco NX-OS QoS guides.
- PFC configuration is per network-qos class per priority — enabling PFC on the wrong class (or on all classes) is a classic misconfiguration that pauses best-effort traffic `[secondary]` — operator practice; FS.com `mlnx_qos` dump in Wave 6 shows the per-priority enable bitmap that must be audited.
- Non-comparable warning repeated: default class counts, queue depths, and buffer carve-outs differ between Nexus 9000 generations — always read the platform-specific guide.

### 12.3 Arista EOS QoS model (structural)
- EOS maps traffic to traffic-classes (TCs) via class-maps/policy-maps; per-TC ECN thresholds and WRED; PFC enabled per 802.1p priority; DCB application TLVs via LLDP for RoCE/iSCSI auto-classification `[vendor-reported]` — Arista TOI/Config guides.
- AI-fabric convention: RoCE on TC 3 with PFC + ECN, control traffic on TC 6/7, best-effort on TC 0 — convention, not standard; verify per deployment `[secondary]` — NetPilot pattern (Wave 6).

### 12.4 ECN marking math (conceptual)
- WRED-with-ECN: below min-threshold → no mark; between min and max → mark with increasing probability; above max → mark (or drop) all; the switch marks the CE codepoint (RFC 3168) instead of dropping `[official]` — RFC 3168; Dell OS10 `random-detect ecn` example in Wave 6 shows green/yellow/red color thresholds.
- Tuning tension: thresholds too low → underutilization (senders back off early); too high → PFC fires first and ECN never gets to act — the Kar lab (Wave 6) demonstrated the well-tuned middle (ECN before PFC) cutting tail latency by an order of magnitude `[secondary]`.

### Wave 12 verification
- Sources: 6. Structural comparison only; no numeric cross-vendor claims made. Trust-boundary and PFC-class warnings recorded as operator practice.

---

