---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/2-6-distributed-anycast-gateway
title: "2.6 Distributed anycast gateway"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [307, 360]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: c9e101b076fa968910c78913028d8d93234de6a49b7a3204063debed2a8c6437
---

# 2.6 Distributed anycast gateway

### 2.6 Distributed anycast gateway

Every leaf that hosts a VLAN configures the **same gateway IP and same virtual
MAC** on its SVI/IRB interface, so a host's default-gateway ARP entry stays
valid no matter which leaf it attaches to (critical for VM mobility)
[vendor-reported][secondary].

- Cisco: `fabric forwarding anycast-gateway-mac`; Arista: virtual IP/MAC on
  SVI; Cumulus/SONiC + FRR: anycast SVI MAC; Dell OS10: virtual gateway
  address — terminology differs, semantics identical [vendor-reported].
- The anycast MAC is typically advertised as a **sticky** Type-2 route so
  mobility logic never treats it as a moved host [vendor-reported].
- Inter-subnet traffic with anycast GW + symmetric IRB: host ARPs for gateway,
  leaf routes at ingress into L3VNI — optimal, single gateway hop
  [secondary].

### 2.7 Gaps and conflicts — Wave 2

- RFC 9251 Type 6/7/8 support in shipping DC platforms: no vendor support
  matrix collected — gap; treat "supported" claims as [unverified] until
  checked per platform/version.
- Cisco's NCS 5500/IOS-XR ESI-label behavior change (SHG label moved from low
  20 bits to high 20 bits of the extended community per RFC 7432 adherence):
  mixed old/new software on the same ES drops BUM or, worse, can loop in
  corner cases — Cisco's own guidance is to minimize mixed-version windows
  [official]. Interop note: peer PEs from other vendors with non-adherent SHG
  encoding need the same caution [official].

---
---

## Wave 3 — EVPN multihoming

### 3.1 Concept and modes

EVPN multihoming lets one CE (host, server, or downstream switch) attach to 2+
PEs (VTEPs) via the same Ethernet Segment (ES), identified by a shared ESI.
Contrast with legacy MC-LAG/vPC: EVPN multihoming is standards-based
(RFC 7432), control-plane driven, and works across vendors [official][secondary].

Two modes [official: RFC 7432 §14]:

- **All-active**: all PEs forward unicast simultaneously (ECMP from remote PEs,
  LACP hashing on the CE side). Only BUM is restricted to the DF. This is the
  DC-fabric default — effectively "vPC without peer-link" [vendor-reported].
- **Single-active**: only the DF forwards unicast and BUM for a given EVI;
  standby PE takes over on failure. Used where the CE cannot do LACP (e.g.,
  active/standby NICs, some appliances) [vendor-reported].

A third behavior, **single-flow-active** (per-flow rather than per-VLAN,
from the EVPN L2-gateway-protocol draft), exists for gateway integrations;
remote PEs must disable ESI-label processing and aliasing when they see it
[secondary: IETF draft-ietf-bess-evpn-l2gw-proto].

