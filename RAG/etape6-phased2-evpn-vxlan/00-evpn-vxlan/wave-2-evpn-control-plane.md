---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/wave-2-evpn-control-plane
title: "Wave 2 — EVPN control plane"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: ["2026-04-23"]
keywords: ["ethernet", "parameters"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [193, 245]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: a3b68ae12635984789438bbabd63e06886ba2792f0914cc300c00dcd146c0f65
---

# Wave 2 — EVPN control plane

## Wave 2 — EVPN control plane

### 2.1 BGP EVPN as the overlay control plane

EVPN (Ethernet VPN), RFC 7432 [official], is a BGP address-family control plane
that replaces flood-and-learn with control-plane-driven MAC/IP advertisement.
It is transport-agnostic: the same control plane runs over MPLS, VXLAN
(RFC 8365 [official]), NVGRE, or SR-MPLS [secondary].

MP-BGP encoding parameters [secondary]:

- AFI 25 (L2VPN), SAFI 70 (EVPN); standard BGP TCP port 179.
- Routes carried in MP_REACH_NLRI / MP_UNREACH_NLRI.
- Design rule repeated across vendor guides: use iBGP route reflectors for
  `l2vpn evpn`; do not full-mesh beyond a handful of PEs; run redundant RRs
  (single RR = single point of failure for the whole fabric) [secondary].

What EVPN buys over flood-and-learn [vendor-reported][independent]:

- ARP/ND suppression (Type-2 with IP) — kills the ARP flood problem.
- Host mobility signaling (MAC mobility extended community, sequence numbers).
- All-active multihoming with DF election and aliasing.
- Multi-tenancy via RD/RT per EVI/VRF.

### 2.2 EVPN route types 1–5 (RFC 7432) + 6–8 (RFC 9251)

| Type | Name | Standards ref | Purpose |
|---|---|---|---|
| 1 | Ethernet Auto-Discovery (A-D), per-ES and per-EVI | RFC 7432 §7.1 | Multihoming: fast convergence (mass MAC withdrawal), aliasing, split-horizon label advertisement [official] |
| 2 | MAC/IP Advertisement | RFC 7432 §7.2 | The workhorse: advertises host MAC, optionally with IP; drives L2 reachability and ARP/ND suppression [official] |
| 3 | Inclusive Multicast Ethernet Tag (IMET) | RFC 7432 §7.3 | PE discovery per bridge domain; builds the BUM flooding list (ingress-replication peer set) [official] |
| 4 | Ethernet Segment | RFC 7432 §7.4 | Redundancy-group discovery; carries inputs to DF election among PEs sharing an ESI [official] |
| 5 | IP Prefix | RFC 7432 §7.5 / RFC 9136 | Advertises IP prefixes without MAC binding; used for symmetric-IRB inter-subnet routing and injecting external prefixes into the fabric [official] |
| 6 | Selective Multicast Ethernet Tag | RFC 9251 [official] | Selective (IGMP-driven) multicast; avoids flooding all BUM to uninterested PEs |
| 7 | Multicast Join Synch | RFC 9251 [official] | Synchronizes IGMP joins across multihomed PEs |
| 8 | Multicast Leave Synch | RFC 9251 [official] | Synchronizes IGMP leaves across multihomed PEs |

Notes:

- Type-1 has two flavors: per-ES (advertised with ESI, lists EVIs on the
  segment) and per-EVI (MAX_ET tag). Implementations must keep them
  semantically distinct (the rustbgpd implementation, 2026-04-23, explicitly
  models this) [secondary].
- Type-2 carries an optional second MPLS label field that symmetric IRB
  reuses per RFC 9135 to signal the L3VNI [official][secondary].
- Type-5 wire formats: IPv4 34-byte and IPv6 58-byte forms per RFC 9136
  [secondary].
- Types 6–8 (RFC 9251, "Internet Multicast Gateway"-era EVPN multicast
  optimizations) exist on the wire, but vendor support in DC VXLAN fabrics is
  sparse as of 2026 — most DC designs still use IMET + ingress replication or
  assisted replication for BUM [unverified — support matrix not collected this
  wave; flagged as gap].

