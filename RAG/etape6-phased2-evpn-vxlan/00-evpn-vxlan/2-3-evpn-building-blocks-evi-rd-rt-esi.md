---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/2-3-evpn-building-blocks-evi-rd-rt-esi
title: "2.3 EVPN building blocks: EVI, RD, RT, ESI"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [246, 306]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 4f2296ab740f3c291ab33490fafdcb4ae6b4cc74cf3499d92132e96af7e6e30f
---

# 2.3 EVPN building blocks: EVI, RD, RT, ESI

### 2.3 EVPN building blocks: EVI, RD, RT, ESI

- **EVI** (EVPN Instance): the EVPN-side representation of a broadcast domain
  (VLAN/BD ↔ L2VNI) or VRF (L3VNI). VLAN-aware, VLAN-based, and VLAN-bundle
  service models come from RFC 7432; VXLAN fabrics overwhelmingly use the
  VLAN-based model (one EVI per VLAN/VNI) [official][secondary].
- **RD** (Route Distinguisher): makes routes unique per EVI/VRF (Type 0/1/2
  encodings per RFC 4364). Auto-derived or manual; auto mode common in
  Dell OS10, Cisco, Arista [vendor-reported].
- **RT** (Route Target): import/export policy — which PEs import which EVIs.
  Auto-RT derivation is the norm in DC fabrics [vendor-reported].
- **ESI** (Ethernet Segment Identifier): 10-byte opaque value identifying a
  multihomed segment shared by 2+ PEs. ESI 0 = single-homed. ESI types include
  operator-configured (Type 0), LACP-based (Type 1), and others per RFC 7432
  §5 [official][secondary].

### 2.4 MAC/IP advertisement, mobility, and ARP suppression

- Type-2 with MAC only → L2 reachability. Type-2 with MAC+IP → also feeds
  ARP/ND suppression tables: the leaf answers ARP for known IPs locally
  instead of flooding [vendor-reported].
- **MAC mobility**: the MAC Mobility extended community carries a sequence
  number; on host move, the new PE advertises a higher sequence; remote PEs
  update. **Sticky MAC** marks infrastructure MACs (e.g., anycast gateway,
  firewall) as immobile so a spoofed move cannot hijack them [official:
  RFC 7432 §15][vendor-reported].
- Duplicate-detection: vendors implement duplicate-MAC detection (e.g., N moves
  in M seconds → freeze) to catch loops/misconfigurations; thresholds and
  behaviors differ per vendor — interop caution [vendor-reported][unverified].

### 2.5 IRB — Integrated Routing and Bridging

IRB lets VTEPs bridge (intra-subnet) and route (inter-subnet). A packet from
subnet A to subnet B crosses a MAC VRF (bridge table) then an IP VRF (routing
table). Two models [official: Cisco config guides; vendor-reported]:

**Asymmetric IRB** ("bridge-route-bridge"):

- Ingress VTEP does bridge + route; egress VTEP does bridge only.
- The ingress leaf rewrites to the destination host's real MAC and encapsulates
  directly into the destination's **L2VNI** — so the ingress VTEP must have the
  destination L2VNI configured even with no local hosts on it.
- Return traffic uses a different VNI than forward traffic (asymmetric path).
- Scaling wall: with N subnets per tenant, every routing VTEP needs all N
  L2VNIs. Simple, but does not scale [vendor-reported][independent].

**Symmetric IRB** ("bridge-route-route-bridge") — the production standard:

- Ingress VTEP: bridge → route into a tenant-wide **L3VNI**; egress VTEP: route
  → bridge out the local L2VNI. Both sides route.
- Each VTEP configures only locally-attached L2VNIs plus the shared L3VNI per
  tenant VRF — far better VNI scaling [vendor-reported].
- Signaling: Type-5 routes advertise prefixes; Type-2 carries the **RMAC**
  (Router MAC extended community, RFC 9135) so the egress knows the ingress
  router's MAC for inner destination rewrite; the second label field in
  Type-2 carries the L3VNI [official][secondary].
- Dell OS10 10.5.1 documents symmetric IRB with eBGP underlay + EVPN overlay,
  unique L3VNI/RD/RT per tenant VRF, and border-leaf gateway roles [official].
- Asymmetric IRB is still the default/simple option on several platforms and
  fine for small fabrics; symmetric is recommended at scale [vendor-reported].

