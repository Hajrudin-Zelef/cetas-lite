---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/wave-5-vendor-implementations-and-interop
title: "Wave 5 — Vendor implementations and interop"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["dci"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [527, 582]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 195af1cb8f83a935dca253635ec4b90e3957c21ef1a9feba66b1bb905e85b4c9
---

# Wave 5 — Vendor implementations and interop

## Wave 5 — Vendor implementations and interop

### 5.1 Cisco NX-OS (Nexus 9000) — VTEP/NVE model

Cisco's DC EVPN-VXLAN implementation centers on the NVE interface
(`interface nve1`) with BGP EVPN control plane [official]:

- VTEP source is a loopback; `source-interface loopback0` under `nve`.
- Member VNIs attached per-VLAN; EVPN auto-RT supported (4-byte ASN
  auto-RT since NX-OS 9.2(1)) [official].
- **Documented limits** (NX-OS 10.4/10.5/10.6 guides) [official]:
  - VNI 16777215 is reserved — do not configure.
  - Any SVI for a VLAN extended over VXLAN **must** use anycast gateway mode
    (`fabric forwarding anycast-gateway-mac`); other modes unsupported.
  - Routing-protocol adjacencies over anycast-gateway SVIs are not supported.
  - `vpc orphan-ports suspend` recommended for single-attached/routed devices
    on vPC VTEPs.
  - Ingress replication + multicast underlay on N9K-X9736C-FX3 line cards
    since 10.5(2)F.
  - Known cosmetic issue: Type-2 mobility sequence mismatch between vPC peers
    (one peer seq K, other seq 0) — no functional impact [official].
- Multihoming on NX-OS is primarily **vPC-based** (Cisco's MLAG) rather than
  standards ESI-LAG; vPC VTEPs share a virtual VTEP IP. ESI-based EVPN
  multihoming exists on newer releases/platforms — check per release
  [vendor-reported][unverified].
- ACI is a separate, policy-driven fabric (not CLI EVPN-VXLAN); ACI multipod/
  multisite provides DCI within the ACI model. Interop between ACI and
  NX-OS EVPN-VXLAN fabrics is via border-leaf / L3Out-style handoffs, not a
  unified control plane [vendor-reported][secondary].

Sanitized pattern (NX-OS leaf) — documentation addresses only:

```
feature bgp
feature nv overlay
feature vn-segment-vlan-based
!
fabric forwarding anycast-gateway-mac 0001.0001.0001
!
interface loopback0
  ip address 192.0.2.11/32
!
interface nve1
  no shutdown
  source-interface loopback0
  member vni 10100
    ingress-replication protocol bgp
!
router bgp 65001
  neighbor 192.0.2.1 remote-as 65000
    update-source loopback0
    address-family l2vpn evpn
      send-community
      send-community extended
```

