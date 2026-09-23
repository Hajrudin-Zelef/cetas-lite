---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-2-arista-eos-vxlan1-interface-mlag-esi
title: "5.2 Arista EOS — Vxlan1 interface + MLAG/ESI"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [583, 629]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: b7811098b4ea5fdafb54004745b0071e4b275a5e06e6d73f5531fec64f6566c7
---

# 5.2 Arista EOS — Vxlan1 interface + MLAG/ESI

### 5.2 Arista EOS — Vxlan1 interface + MLAG/ESI

Arista's model: `interface Vxlan1` with `vxlan source-interface Loopback`,
`vxlan vlan <id> vni <vni>` bindings; BGP EVPN under `router bgp` with
`vlan <id>` / `vrf <name>` sub-contexts, RD/RT auto [official: EOS 4.36.0F
VXLAN config guide].

- **Multi-agent routing is mandatory**: `service routing protocols model
  multi-agent` (reboot required) — EVPN does not function without it
  [official][vendor-reported].
- Gateway: `ip virtual-router mac-address` + `ip virtual-router address` on
  SVIs (VARP heritage); virtual-VTEP for VXLAN routing [official].
- Multihoming: two models — **MLAG** (Arista's vPC equivalent; MLAG pair shares
  anycast VTEP IP on loopback1 while keeping unique loopback0 for BGP) and
  **EVPN all-active / single-active multihoming** (ESI-based, EOS 4.36-era
  documented) [official][vendor-reported].
- EOS troubleshooting entry points: `show bgp evpn summary`,
  `show recirc-channel` (VXLAN routing on MLAG VTEPs needs per-peer
  recirc-channel config) [vendor-reported].
- Arista's AVD (Ansible) generates complete symmetric-IRB L3LS fabrics —
  the de-facto automation path [secondary].
- CVX (CloudVision eXchange): VXLAN control-service (VCS) as an alternative
  controller-driven control plane to BGP EVPN [vendor-reported].

Sanitized pattern (EOS leaf):

```
service routing protocols model multi-agent
!
interface Loopback0
   ip address 192.0.2.21/32
!
interface Vxlan1
   vxlan source-interface Loopback0
   vxlan udp-port 4789
   vxlan vlan 100 vni 10100
!
router bgp 65001
   neighbor 192.0.2.1 remote-as 65000
   neighbor 192.0.2.1 update-source Loopback0
   address-family evpn
      neighbor 192.0.2.1 activate
   vlan 100
      rd auto
      route-target both auto
```

