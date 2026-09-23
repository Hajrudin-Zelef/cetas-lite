---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/5-8-gaps-wave-5
title: "5.8 Gaps — Wave 5"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["agent", "ethernet", "nvidia"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [746, 821]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 24de212f7f155c29bdaefe45346cb9bb01c45846a169d7e5b4cdaaf392bd5ae1
---

# 5.8 Gaps — Wave 5

### 5.8 Gaps — Wave 5

- Juniper native config patterns (set-format) not collected in depth — gap.
- Cisco ESI-based (non-vPC) multihoming support matrix per NX-OS release:
  [unverified].
- NVIDIA SONiC EVPN specifics: gap.
- No systematic vendor-published multi-vendor interop matrix found; the
  Juniper NCE Arista-Junos guide is the best available [official] evidence.

---
---

## Wave 6 — Troubleshooting and operations

### 6.1 Layered troubleshooting method

Work bottom-up; most EVPN-VXLAN outages are underlay or MTU problems wearing
an overlay disguise [secondary][independent]:

1. **Underlay**: VTEP loopback reachability (`ping`), BGP underlay session
   state, MTU (ping with DF-bit and large size end-to-end).
2. **BGP EVPN control plane**: session Established, expected route types
   received (Type-2 for hosts, Type-3 for BUM peers, Type-5 for prefixes).
3. **Data plane**: VNI provisioned on both VTEPs, NVE/Vxlan interface up,
   MAC learned (local or via Type-2), flood list populated for BUM.
4. **Multihoming specifics**: ESI match, DF election state, MH-uplink flags.

### 6.2 Verification commands per vendor

**Cisco NX-OS** [official: VXLAN config guides 10.1–10.6; Cat 9600
troubleshooting guide]:

- `show nve peers` — data-plane peer VTEPs and state.
- `show nve vni` — VNI provisioning and BUM method.
- `show bgp l2vpn evpn` — EVPN RIB; per-route-type filters available.
- `show l2route evpn mac-ip all detail` — remote MAC next-hops with labels.
- `show l2route evpn imet all detail` — egress VNI per remote peer.
- `show forwarding adjacency nve platform` — symmetric/asymmetric NVE
  adjacencies.
- `show nve peers control-plane-vni peer-ip <ip>` — egress/downstream VNI per
  adjacency.
- `show bgp evi <l2-evi>` — VRF ↔ L2VNI association.

**Arista EOS** [vendor-reported: Arista community troubleshooting guides]:

- `show bgp evpn summary` — session state, prefixes received.
- `show bgp evpn route-type mac-ip detail` / `... imet` — control-plane view.
- `show vxlan address-table` — data-plane MAC↔VTEP bindings.
- `show mac address-table` — local learning check.
- `show bgp evpn host-flap` — MAC blacklisted due to flapping.
- `show recirc-channel` — VXLAN routing on MLAG VTEPs (per-peer config needed).
- Log bundle for TAC: `show tech-support`, `show tech-support extended evpn`,
  `show agent log`, `show logging system`.
- Known bugs to keep on the radar (Arista-published): BUG-603591 (flood lists
  deleted for some VLANs → BUM/ARP blackhole); BUG-629099 (single-active
  multihoming: DF won't forward VXLAN-encapsulated packets from a
  non-DF segment peer to a single-active port) [vendor-reported].

**NVIDIA Cumulus / FRR** [official]:

- `vtysh -c 'show bgp l2vpn evpn'` / `show bgp l2vpn evpn route` — EVPN RIB.
- `bridge fdb show` / `ip neigh show` — kernel data-plane state.
- `net show bgp` (NCLU) / `nv show` (NVUE) equivalents.
- Check `evpn mh uplink` present on all fabric-facing interfaces when using
  EVPN multihoming [official].

**Dell OS10** [official]:

- `show virtual-network`, `show nve vni`, `show bgp l2vpn evpn`.
- `show vlt brief` / VLT domain state for multihoming pairs.

**Juniper Junos** [secondary]:

- `show route table bgp.evpn.0` (with `match-prefix 4:*` etc. for route types),
  `show evpn database`, `show ethernet-switching table`.

