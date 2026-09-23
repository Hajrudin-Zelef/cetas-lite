---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/6-3-common-failure-modes
title: "6.3 Common failure modes"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["agent", "dci", "nvidia"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [822, 879]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: 8abbb4e9f97f21c837633f36a71ac9641a2f0914de23b98e74960ac7478f0f32
---

# 6.3 Common failure modes

### 6.3 Common failure modes

| Symptom | Likely cause | Check |
|---|---|---|
| Small pings OK, transfers hang | Underlay MTU < overlay+50/70 | DF-bit ping sweep; raise underlay MTU [independent] |
| No BUM (ARP fails) but unicast OK | Empty flood list / missing Type-3 / multicast RP issue | IMET routes received; `show vxlan flood` equivalents; BUG-603591 on Arista [vendor-reported] |
| Duplicate packets at multihomed host | Split-horizon/DF filtering not applied; missing `evpn mh uplink` (Cumulus); ESI mismatch | Type-4 DF state; ESI equality on peers [official][vendor-reported] |
| MAC flapping between VTEPs | Same VTEP IP twice; looped ES; host genuinely moving | `show bgp evpn host-flap`; unique loopbacks [vendor-reported] |
| Gateway MAC moves / hijack | Missing sticky-MAC on anycast gateway | Sticky community on gateway Type-2 [vendor-reported] |
| Inter-subnet works one way | Asymmetric IRB with missing L2VNI on ingress; RMAC mismatch in symmetric | L3VNI provisioned both sides; `router-mac` consistency (Dell) [vendor-reported] |
| EVPN session won't establish | AFI/SAFI mismatch; TTL (eBGP multihop); missing multi-agent mode (Arista) | `show bgp evpn summary`; capability negotiation [vendor-reported] |
| BUM storm across fabric | Storm-control absent; flood-and-learn leaking | storm-control config; ARP suppression state [secondary] |
| Slow convergence on link failure | Mass-withdrawal not firing; DF re-election churn | Type-1 withdrawal on ES down; DF preference tuning [secondary] |
| Mixed-vendor BUM drop | ESI-label/SHG encoding mismatch (Cisco high-20-bit change) | Software versions aligned per ES [official] |

### 6.4 Operational best practices (recap)

- Prefer **ingress replication** for BUM unless underlay multicast is already
  a solved problem in your network [secondary].
- Enable **ARP/ND suppression** wherever the platform supports it
  [vendor-reported].
- **Storm-control** on tenant VLANs (Dell Enterprise SONiC guide explicitly
  recommends) [vendor-reported].
- `max-med on-startup` / graceful BGP convergence knobs during maintenance
  [vendor-reported].
- Link-state / uplink tracking so edge ports follow fabric uplink state
  (Cumulus `evpn mh uplink`; Dell `uplink-state-group`; Cisco vPC
  orphan-port suspend) [official].
- Monitor: EVPN session state, Type-2/3/5 route counts per VNI, MAC table
  utilization, flood-list sizes, DF election changes [secondary].

### 6.5 Gaps — Wave 6

- Per-vendor telemetry/streaming (gNMI paths for EVPN state): not collected —
  belongs partly to Phase E (automation/observability); flagged for cross-link.
- Arista bug IDs cited are vendor-published; fix versions not verified this wave.

---

## Coverage audit and open items

**Scope coverage:** VXLAN encap (Wave 1) ✓ · EVPN control plane + route types
1–5 (+6–8 noted) (Wave 2) ✓ · IRB symmetric/asymmetric + anycast GW (Wave 2) ✓ ·
multihoming ESI/DF/split-horizon/aliasing (Wave 3) ✓ · DCI multi-site + OTV
(Wave 4) ✓ · vendors Cisco/Arista/NVIDIA/Dell/Aruba/Juniper + interop
(Wave 5) ✓ · troubleshooting per vendor (Wave 6) ✓.

**Open items / gaps:** RFC 9251 Type 6/7/8 shipping support matrix · RFC 8584
DF-algorithm support per platform · Juniper native config patterns ·
NVIDIA SONiC EVPN specifics · Cisco ESI (non-vPC) multihoming matrix ·
systematic multi-vendor interop matrix · EVPN gNMI telemetry paths (→Phase E) ·
reserved-VNI ranges per vendor · independent OTV→EVPN migration case studies.

**Provenance audit:** every factual claim above carries a tag; `[unverified]`
marks single-source or unconfirmed items; no identifiers, SKUs, or URLs were
guessed — all URLs are verbatim from search results.

*End of Phase D wave 2. File is append-only; earlier waves untouched.*
