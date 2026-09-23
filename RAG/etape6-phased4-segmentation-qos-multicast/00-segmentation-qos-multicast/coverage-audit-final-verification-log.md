---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/coverage-audit-final-verification-log
title: "Coverage audit & final verification log"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["asic", "ethernet"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [257, 330]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 971d77c78cc7d280cf3b3a925b49b9a470fa2cba51569f64690df832ddee2cbe
---

# Coverage audit & final verification log

## Coverage audit & final verification log

### Scope-to-wave map
| Scope item | Covered in |
|---|---|
| VLAN design at scale, 4094 limit, QinQ | Wave 1 §§1.1–1.2 |
| VLAN-VNI mapping, segmentation strategies | Wave 1 §§1.3–1.4 |
| VRF-lite, leaking, multi-tenancy patterns | Wave 2 |
| Microsegmentation, ACLs, PBR | Wave 3 |
| Zero-trust DC concepts | Wave 3 §3.5 |
| QoS classification/marking (DSCP/CoS) | Wave 4 §4.1 |
| Queuing/scheduling | Wave 4 §4.2 |
| DCB: PFC, ETS, DCBX | Wave 5 |
| ECN, RoCEv2 lossless tuning | Wave 6 |
| Vendor QoS models | Wave 6 §6.4 |
| PIM-SM/SSM, RP placement, MSDP, Anycast-RP | Wave 7 |
| Multicast in VXLAN (IR vs underlay), TRM | Wave 8 §§8.1–8.2 |
| IGMP snooping/querier | Wave 8 §8.3 |
| DHCP relay, DNS, NTP/PTP | Wave 9 |

### Consolidated open items / gaps
1. VNI allocation practice in production fabrics (sparse vs dense) — `[unverified]`.
2. ACL TCAM scale table across 2026 DC ASICs — `[unverified]`.
3. Cisco TrustSec/SGT deep dive deferred to a future pass.
4. Queue-count/buffer-size table per 2026 ASIC (Spectrum-4/5, Tomahawk 5/6, Jericho3-AI) — `[unverified]`.
5. PFC-watchdog/deadlock defaults per vendor — `[unverified]`.
6. Per-vendor IGMP-snooping/EVPN Type-6–8 support matrix — `[unverified]`.
7. Per-platform PTP profile support matrix — `[unverified]`.
8. DNS-anycast deployment prevalence — `[unverified]`.

### Conflicts registered
- C1: Anycast-RP without MSDP — supported on Nexus (PIM variant) and Junos (with/without MSDP), NOT on IOS-XE IPv4 (MSDP required). Platform-specific; never generalize.
- C2: "Ethernet vs InfiniBand for AI" performance parity — vendor-adjacent claim (intelligentvisibility, pro-Ethernet); recorded as perspective, not consensus.
- C3: rdma-ai-cluster GitHub repo's "96% drop reduction / 2× bandwidth" — self-reported, downgraded to `[unverified]`.
- C4: QoS defaults (queues, maps, buffers) are NOT comparable across vendors — recorded per-vendor only.

### Provenance audit
- All factual claims carry one of `[official]`, `[vendor-reported]`, `[independent]`, `[secondary]`, `[unverified]`. No identifiers guessed. URLs used verbatim from search results only.

*End of Phase D4. File complete: 9 waves, append-only.*

## Wave 10 — EVPN control plane and segmentation (route types, RD/RT, IRB)

### 10.1 EVPN NLRI route types (RFC 7432 and extensions)
| Type | Name | Carries / purpose |
|---|---|---|
| 1 | Ethernet Auto-Discovery (A-D) | ESI, Ethernet Tag; per-ES and per-EVI; aliasing + fast mass-withdraw on multihoming failure |
| 2 | MAC/IP Advertisement | Host MAC + optional IP/ARP binding, ESI, router-MAC; L2 reachability and ARP/ND suppression |
| 3 | Inclusive Multicast Ethernet Tag (IMET) | VTEP membership per VNI; builds the ingress-replication BUM flood list |
| 4 | Ethernet Segment | ESI + originating VTEP; ES discovery, Designated Forwarder election, split-horizon |
| 5 | IP Prefix | Subnet/prefix routes independent of MAC; symmetric-IRB inter-subnet routing |
| 6 | Selective Multicast (SMET) | IGMP-derived group membership; forward only where receivers exist |
| 7/8 | IGMP join/leave sync | State sync between PEs for multi-homed CEs (incl. DF) |

`[official]` — Cisco NCS 5500 IOS XR 7.11.x L2VPN/EVPN guide (per-type usage table, incl. note: with EVPN IRB, host /32 via RT-2, subnet /24 via RT-5). https://www.cisco.com/c/en/us/td/docs/iosxr/ncs5500/vpn/711x/configuration/guide/b-l2vpn-cg-ncs5500-711x/evpn-features.html `[secondary]` — artofinfra EVPN best practices; hon95 wiki; lolyu EVPN book notes. https://github.com/danjonesio/docs-dot-artofinfra/blob/HEAD/docs/general/evpn.md

### 10.2 Why this matters for segmentation
- Type-2 MAC/IP advertisement is what lets the fabric suppress ARP/ND flooding: the control plane already knows the binding, so no broadcast is needed — a segmentation-adjacent win (less BUM = smaller blast radius and less snooping load) `[secondary]` — artofinfra.
- Type-3 IMET is the membership list that makes ingress replication work without underlay multicast `[secondary]` — Cisco community; Wave 8.
- Type-5 is the L3-segmentation primitive: tenant prefixes advertised per-VRF with per-tenant RTs; the RT is what keeps tenant A's /24 out of tenant B's table `[secondary]` — NetPilot multi-tenant example (Wave 2); eoprede NX-OS reference.
- Vendor-agnostic best practices `[secondary]` — artofinfra: always use BGP EVPN in production (never flood-and-learn); use iBGP route reflectors for `l2vpn evpn` rather than full mesh; run redundant RRs (single RR = SPOF for the whole fabric).
- eBGP-underlay nuance: EVPN next-hop is the destination VTEP; transit eBGP speakers must use `next-hop unchanged` (or equivalent) for the EVPN AF or remote leaves black-hole the overlay `[secondary]` — lolyu EVPN notes; enizaksoy multi-AS lab (Wave 2).

### 10.3 RD/RT design for tenants
- Route Distinguisher: makes routes unique across VRFs (same prefix in two tenants); `rd auto` derives it from VNI + router-ID `[secondary]` — eoprede NX-OS reference.
- Route Target: decides which local VRF imports a route; `route-target import/export auto` derives from VNI — preferred in most DC designs; manual RTs needed for leaking/shared-services patterns `[secondary]` — same source.
- Multi-tenancy rule of thumb: one RT pair per tenant VRF for strict isolation; additional import RTs only where leaking is explicitly designed (shared services, internet egress) `[secondary]` — NetPilot; Cisco Press ACI (Wave 2).
- `send-community both` (or `extended`) is mandatory on EVPN BGP sessions — without extended communities, RTs are stripped and routes are rejected `[secondary]` — eoprede reference.

### Wave 10 verification
- Sources: 6. No conflicts. EVPN route-type table cross-checked across Cisco official + 3 independent secondary sources.

---

