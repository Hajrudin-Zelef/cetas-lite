---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/1-6-underlay-requirements
title: "1.6 Underlay requirements"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [148, 192]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: d6fada53ad503499b1e57faacd523b974d21d91d84087392ea563ce036145476
---

# 1.6 Underlay requirements

### 1.6 Underlay requirements

For a VXLAN overlay to work, the underlay must provide [secondary][vendor-reported]:

1. **IP reachability between VTEP loopbacks** — via OSPF, IS-IS, eBGP, or static
   routes. eBGP with private ASNs per device (or per tier) is the dominant
   modern DC underlay design (see wave-1 file); OSPF/IS-IS remain common in
   enterprise DCs.
2. **MTU headroom** — underlay MTU ≥ overlay MTU + 50 (IPv4) or + 70 (IPv6).
   Typical: underlay 9216, overlay 9000/1500. Mismatch = silent drops of
   large frames [independent].
3. **ECMP** — equal-cost multipath so tunnel entropy (UDP sport) spreads load.
   Note: some platforms hash only on outer IP + UDP ports; if the UDP source
   port entropy is weak, polarization can occur [secondary][unverified].
4. **Multicast (only for flood-and-learn BUM)** — PIM-SM/Bidir with RP placement
   planned; not needed at all with EVPN ingress replication [secondary].
5. **No STP dependency** — the underlay is L3; STP should be absent or confined
   to legacy edges. BPDU guard/filter on host ports remains relevant
   [secondary].

### 1.7 BUM handling options (data-plane view)

BUM = broadcast, unknown-unicast, multicast [secondary]:

| Method | Mechanism | Underlay need | Scale note |
|---|---|---|---|
| Underlay multicast | one mcast group per VNI (or shared) | PIM in underlay | group state grows with VNIs; classic RFC 7348 model [official] |
| Ingress (head-end) replication | ingress VTEP unicasts N copies | none extra | replication load on ingress; EVPN Type-3 automates the peer list [secondary] |
| Assisted replication | designated replicator VTEPs | none extra | offloads replication from low-end leaves [vendor-reported] |
| Static | manually configured peer VTEPs | none | lab/small only [secondary] |

### 1.8 Gaps and conflicts — Wave 1

- Exact reserved-VNI ranges are vendor-specific; no cross-vendor standard found
  this wave — flagged as gap.
- Claims about "analyst consensus" on flood-and-learn being legacy are
  [unverified]; vendor design guides uniformly recommend EVPN, which is
  [vendor-reported] fact, not independent proof of deployment share.
- Linux default UDP port 8472 vs IANA 4789 interop note is [independent]/
  [secondary] from lab writeups; verify against current kernel behavior before
  relying on it.

---
---

