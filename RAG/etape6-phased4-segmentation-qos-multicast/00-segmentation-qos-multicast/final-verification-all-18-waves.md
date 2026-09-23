---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/final-verification-all-18-waves
title: "Final verification (all 18 waves)"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["asic", "ethernet", "mcp"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [535, 611]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: cd6736827225e8426f8ac77944a4a2a214eaa84812b9817a1ccb5962368de7e5
---

# Final verification (all 18 waves)

## Final verification (all 18 waves)

- Scope-to-wave map (extended): Wave 10 EVPN/segmentation control plane · Wave 11 VRF deep-dive · Wave 12 QoS deep-dive · Wave 13 congestion-control algorithms + UEC · Wave 14 PIM/mVPN deep-dive · Wave 15 services deep-dive · Wave 16 segmentation practice · Wave 17 multicast practice · Wave 18 QoS practice.
- New open items: per-ASIC buffer/PFC calculator availability `[unverified]`; PCI artifact lists vary by QSA (practice note).
- Conflicts carried forward: C1 (anycast-RP/MSDP platform variance) · C2 (Ethernet-vs-IB parity as vendor perspective) · C3 (self-reported lab numbers) · C4 (QoS defaults non-comparable). New: C5 — RoCE DSCP marking values (26 vs 48) are operator conventions with no universal standard; document per site.
- Provenance: all factual claims tagged; meta/verification lines ("Sources:", "Note:") are editorial and untagged by design.

*End of Phase D4 — 18 waves, append-only, complete.*

## Wave 19 — Config walkthroughs: EVPN multi-tenant VRF (NX-OS)

The following is a structural walkthrough synthesized from sourced NX-OS references — read as a template, validate against the running OS version before use `[secondary]` — eoprede network-mcp NX-OS EVPN reference; NetPilot multi-tenant pattern.

```
! Step 1 — underlay features
feature bgp
feature nv overlay
feature vn-segment-vlan-based

! Step 2 — EVPN per-VNI stitching (L2 VNI 10100 = tenant A VLAN 100)
evpn
  vni 10100 l2
    rd auto
    route-target import auto
    route-target export auto

! Step 3 — BGP EVPN overlay (eBGP underlay example)
router bgp 65001
  router-id 10.255.0.1
  neighbor 10.0.0.0/24 remote-as 65000
    update-source loopback0
    address-family l2vpn evpn
      send-community both
      route-map NEXT-HOP-UNCHANGED out   ! required: EVPN next-hop = origin VTEP

route-map NEXT-HOP-UNCHANGED permit 10
  set ip next-hop unchanged

! Step 4 — tenant VRF with L3 VNI (symmetric IRB)
vrf context TENANT-A
  vni 900100
  rd auto
  address-family ipv4 unicast
    route-target import auto evpn
    route-target export auto evpn

vlan 900
  vn-segment 900100
interface Vlan900
  description L3VNI-FOR-TENANT-A
  vrf member TENANT-A
  ip forward        ! no IP needed on the L3VNI SVI

interface nve1
  member vni 900100 associate-vrf
```

- Key checks: `send-community both` present on every EVPN session; spine route-reflector clients configured; L3VNI SVI carries `ip forward` only `[secondary]` — same source.
- Segmentation reading of this config: tenant isolation lives in `route-target import/export auto` per VRF — two tenants with distinct auto-RTs never exchange routes; leaking = adding explicit import RTs `[secondary]` — NetPilot pattern (Wave 2).
- DHCP relay on the tenant SVI (from Wave 9, deliabtech/NX-OS pattern):

```
interface Vlan30
  vrf member green
  ip address 10.10.30.1/24
  fabric forwarding mode anycast-gateway
  ip dhcp relay address 10.10.20.100
  ip dhcp relay source-interface loopback3   ! per-leaf unique, in vrf green
ip dhcp relay information option              ! option 82
ip dhcp relay information option vpn          ! VRF-aware link-selection
```

### Wave 19 verification
- Sources: 3. Config blocks are templates synthesized from dated references, labeled as such; not claimed as tested on any specific NX-OS release.

---

