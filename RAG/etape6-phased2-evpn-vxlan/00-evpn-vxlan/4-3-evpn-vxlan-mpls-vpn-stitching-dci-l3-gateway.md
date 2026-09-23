---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/4-3-evpn-vxlan-mpls-vpn-stitching-dci-l3-gateway
title: "4.3 EVPN-VXLAN ↔ MPLS-VPN stitching (DCI L3 gateway)"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["dci", "ethernet"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [479, 526]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 2f23e78b1bb2ff8b937ab1644eff9c5938669e391ab0e4731fea5906705456fd
---

# 4.3 EVPN-VXLAN ↔ MPLS-VPN stitching (DCI L3 gateway)

### 4.3 EVPN-VXLAN ↔ MPLS-VPN stitching (DCI L3 gateway)

For DC-to-legacy-SP interconnection, Cisco ASR 9000 (IOS XR 24.x) documents a
**DCI L3 gateway** stitching EVPN-VXLAN to MPLS L3VPN (VPNv4/v6) [official]:

- Two RT sets per VRF: stitching RTs (toward the DC EVPN neighbor) and normal
  RTs (toward the L3VPN neighbor); the DCI router stitches them.
- Use cases: multi-tenant DC-to-DC over an SP core; DC-to-legacy-PE CE
  connectivity [official].

### 4.4 OTV comparison

OTV (Overlay Transport Virtualization, Cisco proprietary) was the 2010s DCI
standard [vendor-reported][secondary]:

| Aspect | OTV | EVPN multi-site |
|---|---|---|
| Standard | Cisco proprietary | RFC 7432 + vendor extensions (multi-vendor) [official] |
| Control plane | IS-IS based overlay hellos/adjacencies | BGP EVPN [official] |
| Multihoming | active/standby edge devices | all-active with ESI/DF [official] |
| Failure detection | aggressive timers, fate-sharing risks | BGP-driven, per-route withdrawal [secondary] |
| ARP handling | ARP suppression via ND cache | Type-2 ARP suppression, standardized [official] |
| Trajectory | feature-frozen; Cisco positions EVPN multi-site as the replacement | current strategic DCI [vendor-reported] |

Migration note: OTV→EVPN-multi-site migrations are documented in Cisco
migration guides; dual-operation during migration needs careful VLAN/VNI
mapping to avoid loops [vendor-reported][unverified — no independent migration
case study collected this wave].

### 4.5 Nokia I-ES (interconnect Ethernet segments) — note

Nokia SR OS documents **Interconnect Ethernet Segments (I-ES)** per
draft-ietf-bess-dci-evpn-overlay: virtual ESs letting dual-BGP-instance DC
gateways (EVPN-MPLS + EVPN-VXLAN) provide redundancy in VXLAN access networks,
with single-active/all-active, ESI-label split-horizon, DF election, aliasing
and backup per RFC 7432. Positioned as superior to older multi-homed anycast
DCGW designs (which risk BUM duplication with mLDP in the WAN) [official:
Nokia SR OS 23.7 docs].

### 4.6 Gaps — Wave 4

- Arista multi-domain DCI specifics and interop notes vs Cisco BGW: not
  collected in depth — gap for a future pass.
- No independent OTV→EVPN migration case studies found this wave.

---
---

