---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-3-nvidia-cumulus-linux-frr-and-sonic
title: "5.3 NVIDIA — Cumulus Linux (FRR) and SONiC"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: actor-profile
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "asic", "memory"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [630, 696]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 121f8631ef1c87b541b449e10b3d642b68d9aa3d03cbce5bac15613ef8af5adb
---

# 5.3 NVIDIA — Cumulus Linux (FRR) and SONiC

### 5.3 NVIDIA — Cumulus Linux (FRR) and SONiC

**Cumulus Linux** (FRR-based; docs cited are 4.x-era — 5.x is current as of
2026, verify deltas) [official]:

- Single VXLAN device model; EVPN configured in FRR (`/etc/frr/frr.conf`)
  via NCLU (`net add ...`), NVUE (`nv set ...`), or vtysh.
- Symmetric IRB: per-tenant VXLAN interface carrying the L3VNI + SVI in tenant
  VRF + VRF↔L3VNI mapping; Type-5 origination via
  `advertise ipv4 unicast` under `router bgp <asn> vrf <name>` /
  `address-family l2vpn evpn` [official].
- Gateway SVI advertisement: `advertise-svi-ip` (all or per-VNI)
  [official].
- EVPN multihoming: `es-sys-mac` on bonds + `evpn mh uplink` on fabric-facing
  interfaces (**mandatory** — without MH uplinks, split-horizon/DF filters
  don't apply and BUM duplicates/loops back to the ES, causing MAC flaps)
  [official].
- ASIC note: centralized/symmetric/prefix-based routing on Spectrum needs
  Spectrum-A1 or later; Trident II+ border-leaf workaround documented
  [official].

**NVIDIA Spectrum-X / SONiC**: NVIDIA's newer fabrics run SONiC with FRR EVPN;
configuration follows the FRR/SONiC data model rather than NCLU [vendor-reported]
[unverified — SONiC EVPN specifics not collected this wave; gap].

Sanitized pattern (FRR/Cumulus-style):

```
router bgp 65001 vrf RED
  address-family l2vpn evpn
    advertise ipv4 unicast
  exit-address-family
!
interface swp51
  evpn mh uplink
```

### 5.4 Dell — SmartFabric OS10 and Enterprise SONiC

**OS10** (10.5.x documented) [official]:

- `nve` with `source-interface loopback0`; `virtual-network <id>` +
  `vxlan-vni`; VLAN↔VN mapping switch-scoped or port-scoped.
- EVPN auto-EVI mode + `disable-rt-asn` (auto-RT without ASN — required when
  leaves use per-leaf ASNs; only when all VTEPs are OS10) [official].
- Anycast gateway: `ip virtual-router mac-address` + `ip virtual-router
  address` on `interface virtual-network` [official].
- Symmetric IRB: `evpn` → `router-mac` + per-VRF `vni` (L3VNI) +
  `route-target auto`; Type-5 origination via `advertise ipv4 connected`
  [official].
- Multihoming via **VLT** (Dell's MLAG): VLT pair shares loopback0 IP as one
  logical VTEP; iBGP unnumbered peering between VLT peers over vlan4000;
  `uplink-state-group` for UFD [official].
- BGP unnumbered (RFC 5549) is Dell's preferred underlay peering in examples
  [official].

**Enterprise SONiC** (Dell Technologies Info Hub use-case guidebook)
[vendor-reported]:

- Symmetric vs asymmetric IRB guidance: symmetric recommended for scale
  (VTEPs don't carry every tenant VLAN in RIB memory); asymmetric OK for
  small/medium DCs.
- Best practices: L3 demarcation leaf→spine; MC-LAG at leaf; storm-control for
  BUM; unnumbered BGP leaf-spine; neighbor/ARP suppression on tenant VLANs;
  link-state tracking; `max-med on-startup` for graceful convergence
  [vendor-reported].

