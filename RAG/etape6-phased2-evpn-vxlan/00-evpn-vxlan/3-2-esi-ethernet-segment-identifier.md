---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/3-2-esi-ethernet-segment-identifier
title: "3.2 ESI (Ethernet Segment Identifier)"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["ethernet", "decode"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [361, 412]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 6e727d49b06830e6d26b656089cfdcdcdb6b4f47f1581bb15d449611bfa6cbd4
---

# 3.2 ESI (Ethernet Segment Identifier)

### 3.2 ESI (Ethernet Segment Identifier)

- 10 bytes: 1-byte type + 9-byte value [official].
- Types: 0x00 manual/operator-configured, 0x01 LACP-derived (from LACP system
  ID/port), 0x02 STP, 0x03 MAC-based, 0x05 router-ID-derived; ESI all-zeros =
  single-homed [official][secondary].
- **Peering PEs MUST share the same ESI** or multihoming procedures break;
  ESI mismatch is a classic misconfiguration (traffic blackholes or loops)
  [official][independent].
- Practical: DC fabrics usually configure ESI manually per MLAG/server pair or
  derive from LACP; document the scheme — ESI is the join key for every other
  procedure in this wave [secondary].

### 3.3 DF election (Designated Forwarder)

The DF is the PE responsible for forwarding **BUM into the ES** (and unicast
to the CE in single-active mode). Election inputs ride Type-4 (Ethernet
Segment) routes: each PE advertises its ESI + originating router IP
[official].

- Default algorithm (RFC 7432): **modulo/service-carving** — PEs ordered by
  originating IP; DF for VLAN V = PE at index (V mod N). DF is elected
  **per-VLAN/EVI**, so different VLANs on the same ES land on different leaves
  — the lowest IP does not win everything. Tunable to preference-based
  election on most platforms [official][secondary].
- RFC 8584 [official] defines an extensible DF-election framework; newer
  algorithms (e.g., HRW/highest-random-weight) address service-carving's
  unfairness on PE failure/addition (all VLANs re-carve). Vendor support for
  non-default algorithms varies — check per platform/version [unverified].
- Non-DF PEs **filter** BUM toward the ES; without this, the CE receives
  duplicate broadcasts. In virtualized lab NOS (e.g., cEOS containers), the
  election state is observable but data-plane filtering may not be enforced —
  lab-vs-hardware caveat [secondary].

### 3.4 Split-horizon (loop prevention)

When BUM is flooded among PEs, a frame arriving from the fabric must not be
sent back out another PE of the **same ES** — that would loop through the CE.
Two mechanisms [official][secondary]:

- **ESI-label (split-horizon label) filtering** (MPLS heritage, adapted): each
  PE advertises a split-horizon label in the Type-1 per-ES route (ESI Label
  extended community); ingress PE pushes it; egress PEs of the same ES drop
  frames carrying their own ES's label. Cisco's NCS/IOS-XR note (Wave 2):
  encoding moved from low-20 to high-20 bits per RFC 7432 adherence — mixed
  versions mis-decode and can loop in corner cases [official].
- **Local bias** (VXLAN all-active, RFC 8365): a PE receiving BUM from the
  fabric with its own ESI does not forward it to the local ES; instead, the
  ingress PE sends BUM only to the DF-capable set and each PE applies
  local-bias filtering. Simpler than labels; standard for VXLAN all-active
  [official: RFC 8365][secondary].

