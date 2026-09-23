---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-5-aruba-cx-hpe-aruba-networking
title: "5.5 Aruba CX (HPE Aruba Networking)"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["dci"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [697, 745]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 24e1c466b940b2dadca6e5ae2c61184b25237bf53e3583e557cb5ce4b5c1933f
---

# 5.5 Aruba CX (HPE Aruba Networking)

### 5.5 Aruba CX (HPE Aruba Networking)

- EVPN-VXLAN on DC platforms: 8325/8360/8400/9300/CX-10000; full iBGP/eBGP
  underlays, L2VNI + L3VNI symmetric IRB [vendor-reported][secondary].
- Anycast gateway = **active-gateway** (`interface vlan N` /
  `active-gateway ip` + `active-gateway mac`) [secondary].
- Multihoming via **VSX** (Virtual Switching Extension): two-chassis cluster,
  ISL + keepalive, shared system-mac/active-gateway; Aruba's validated DCN
  workflow pairs eBGP EVPN multi-AS VXLAN **with VSX** (note: that workflow
  explicitly does not configure centralized L3 gateway) [official:
  developer.arubanetworks.com].
- AOS-CX CLI is EOS-inspired (multi-line stanzas, `exit` boundaries) —
  relevant for automation/migration [secondary].
- Multi-fabric DCI: border-leader VTEP concept (1 per site, e.g., 8325) peering
  via eBGP to border VTEPs; tunnel-to-tunnel forwarding with loop prevention
  between fabrics [official: HPE techdocs AOS-CX 10.13].
- Config shape: `interface vxlan 1` / `vni <VNI>` / `vlan <VLAN>` plus `evpn`
  stanza with per-VLAN RD/RT [secondary].

### 5.6 Juniper (QFX / Junos)

- Juniper documents EVPN-VXLAN on QFX (edge-routed bridging overlays,
  ESI-LAG multihoming) — Junos `show route table bgp.evpn.0` displays
  Type-1..5 in `1:`/`2:`/`3:`/`4:`/`5:` prefix format [secondary].
- **Interop proof**: Juniper NCE guide "EVPN VXLAN Interoperability Between
  Arista EOS and Junos OS" demonstrates a mixed fabric: Arista MLAG leaves
  (shared loopback1 anycast VTEP) + Junos ESI-LAG leaves interoperating —
  the strongest multi-vendor EVPN-VXLAN interop evidence collected this wave
  [official: juniper.net NCE].
- Cisco↔Arista, Arista↔Cumulus interop is lab-demonstrated in community
  fabrics; systematic vendor-published interop matrices remain rare — gap
  [secondary][unverified].

### 5.7 Interop notes (cross-vendor)

What must align for multi-vendor EVPN-VXLAN [secondary][independent]:

1. UDP port 4789 (watch legacy 8472 defaults).
2. RT/RD conventions — auto-RT schemes differ; safest is explicit RTs or a
   documented auto scheme used fabric-wide.
3. ESI format agreement for multihoming (manual ESI values must match).
4. DF election algorithm — default RFC 7432 modulo everywhere unless all
   vendors support the same RFC 8584 alternative.
5. Anycast-gateway MAC consistency (same virtual MAC on every vendor's leaf).
6. BUM method: agree on ingress replication (simplest cross-vendor) vs
   underlay multicast.
7. MTU: 50/70-byte overhead on every platform.
8. ESI-label/SHG encoding adherence (Cisco's high-20-bit change — Wave 2).

