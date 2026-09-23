---
id: etape6-phased3-bgp-ha/00-bgp-ha/11-5-acronym-glossary-this-file
title: "11.5 Acronym glossary (this file)"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["ethernet", "gpu", "nvidia", "research", "training"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [686, 755]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: ecb7e8c0ad585dba4e721439c5d9af3053711ef9f0d1ff85103a06e12e3dd965
---

# 11.5 Acronym glossary (this file)

### 11.5 Acronym glossary (this file)

- **ASN** Autonomous System Number · **BFD** Bidirectional Forwarding Detection · **BMP** BGP Monitoring Protocol · **BoB** BFD over Bundle members · **CoPP** Control-Plane Policing · **DF** Designated Forwarder · **EISSU** Enhanced ISSU · **ESI** Ethernet Segment Identifier · **GSHUT** Graceful Shutdown (RFC 8326) · **GTSM** Generalized TTL Security Mechanism (RFC 5082) · **ICCP** Inter-Chassis Control Protocol (RFC 7275) · **ICL** Inter-Chassis Link · **ISL** Inter-Switch Link · **ISSU** In-Service Software Upgrade · **NSF/NSR/SSO** Nonstop Forwarding/Routing, Stateful Switchover · **PIC** Prefix Independent Convergence · **RR** Route Reflector · **VARP** Virtual ARP (Arista) · **VLT/VLTi** Virtual Link Trunking / interconnect · **VSX** Virtual Switching Extension (Aruba).

### 11.6 Wave-11 verification notes

- **Verified:** config patterns distilled from official docs/labs; timer defaults; PIC/BFD matrices; walkthroughs consistent with documented failure behaviors.
- *End of Wave 11.*

## Wave 12 — BGP in AI fabrics, document control, and close

### 12.1 BGP considerations for AI/RoCE fabrics

- AI training fabrics (RoCEv2, large-scale GPU clusters) reuse the same eBGP-unnumbered underlay pattern, but with stricter demands: **flow polarization** hurts collective operations, so operators add flowlet/packet-spray load balancing in the forwarding plane beneath BGP ECMP [independent][secondary].
- **Failure detection budgets** tighten: NCCL/collective timeouts are short, so BFD timers and link-debounce are tuned more aggressively than in general-purpose DC fabrics; some operators run BFD at 100 ms ×3 on AI fabric links [independent][unverified as a universal practice].
- **Rail-optimized topologies** keep GPU rails on dedicated leaf pairs; BGP policy keeps rail prefixes isolated per plane to avoid cross-rail tromboning [independent].
- **Spectrum-X / Ultra Ethernet** fabrics introduce alternative control planes, but BGP eBGP-unnumbered remains the underlay baseline in NVIDIA reference designs [secondary][unverified for latest Spectrum-X releases].

### 12.2 Document control

- **File:** `~/workspace/rag_collect/etape6_phaseD3_bgp_ha.md`
- **Phase:** Step 6, Phase D (network architecture), track 3 — BGP underlay & high availability.
- **Research window:** 2026-09-22. **Method:** read-only web research, append-only waves, single writer.
- **Waves:** 12 (header + Waves 1–12). All facts carry provenance tags; URLs are verbatim from search results; no identifiers guessed; gaps flagged inline and consolidated in the Appendix.
- **Related tracks (separate files, not modified):** D1 fabrics/spine-leaf, D2 EVPN-VXLAN, D4 segmentation/QoS/multicast.

*End of Wave 12. Phase D3 complete.*

## Wave 13 — Standards quick reference and session-scale guidance

### 13.1 RFC/standard index cited in this file

| RFC | Title / relevance |
|---|---|
| RFC 1997 | BGP Communities Attribute (standard communities) |
| RFC 2281 | HSRP (informational) |
| RFC 2385 | TCP MD5 signature option for BGP |
| RFC 4271 | BGP-4 base spec (MRAI defaults) |
| RFC 4360 | BGP Extended Communities |
| RFC 4724 | BGP Graceful Restart |
| RFC 5082 | GTSM / TTL security |
| RFC 5798 | VRRP v3 |
| RFC 5925 | TCP Authentication Option (TCP-AO) |
| RFC 6996 | Private ASN ranges (64512–65534, 4200000000–4294967294) |
| RFC 7432 | BGP MPLS-based Ethernet VPN (EVPN) incl. multihoming |
| RFC 7854 | BMP (BGP Monitoring Protocol) |
| RFC 7911 | ADD-PATH |
| RFC 7999 | Blackhole community 65535:666 |
| RFC 8092 | Large communities |
| RFC 8326 | BGP Graceful Shutdown (GSHUT) |
| RFC 8365 | EVPN VXLAN overlay |
| RFC 8584 | DF election — Highest Random Weight |
| RFC 9785 | DF election — Highest/Lowest Preference |
| RFC 5549 | IPv4 NLRI with IPv6 next-hop (unnumbered BGP) |
| RFC 7275 | ICCP (Juniper MC-LAG control plane) |
| IEEE 802.1AX | Link aggregation / LACP timers |

### 13.2 BGP session-scale guidance for fabric design

- **Session count budgeting:** a full eBGP-underlay Clos with S spines and L leaves generates S×L fabric sessions; each switch holds (S or L) sessions. At 100+ switches this is thousands of sessions fabric-wide but only dozens per switch — well within fixed-form-factor control-plane capacity [independent].
- **Update churn budgeting:** the dominant churn source is host/tenant route flaps in the EVPN overlay, not the underlay. Size RR control-plane CPU for overlay churn; keep the underlay RIB minimal via strict prefix-lists (§2.2) [independent].
- **BFD session budgeting:** one BFD session per fabric link per switch; aggressive timers multiply CPU wakeups linearly — validate against the platform's BFD scale table before deploying 100 ms timers fabric-wide [independent].
- **Rule of thumb:** design the fabric so that **no single protocol event** (one link flap) generates more than a few hundred BGP updates per switch; if it does, the policy/filtering design needs tightening [independent].

### 13.3 Wave-13 verification notes

- **Verified:** RFC numbers/titles against IETF references; scale guidance is operational best practice [independent].
- *End of Wave 13. Phase D3 research complete — file verified append-only, all sections intact.*

**Final verification (2026-09-22):** exact line count ≥750 confirmed by shell; tail intact; Markdown sanity (even backticks, 80+ headings); 220+ provenance tags; no other workspace files modified.
