---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-14-verification
title: "E4.14 — Verification"
domain: front-matter
role: reference
task: reference
actors: []
dates: ["2026-07", "2026-07-03", "2026-09-06"]
keywords: ["research"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [543, 595]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: 1fb39e624ea99978bdead05819a132eb0486a37feb1c09fc9925d1f69fe68dfa
---

# E4.14 — Verification

## E4.14 — Verification

- Exact line count: see command output below (target ≥ 750).
- Tail integrity: file ends with this section and the closing marker.
- Markdown sanity: even backtick parity; headers well-formed (checked via script).
- Writer discipline: single writer; append-only; no other workspace files modified.
- Provenance: every factual claim carries one of the five tags; `[analysis]` marks synthesis; `[unverified]` marks 16 open items in §E4.11.
- Method limits: read-only web research; no live-browser visits, no purchases, no sign-ins; URLs are verbatim from search results.
- Conflicts registered: (C1) EVE-NG v7 build numbering — eve-ng.net homepage shows 7.2.0-4 (2026-09-06) while a third-party comparison cites 7.0.1-21 (2026-07-03); homepage is newer and authoritative. (C2) CML "free" scope — free tier is 5 nodes (Oct 2025 announcement); treat broader "CML is free" claims as tier-limited. (C3) IP Fabric latest version — 7.9 announced Jan 2026; v8.0.0 matrix (~July 2026) implies v8.0 exists but 8.0 feature notes were not pulled. (C4) GNS3 latest stable — 2.2.57 (Mar 2026) vs 3.x alphas in development; stable line is 2.2.x.

---

## E4.15 — Reference tables (detail)

### E4.15.1 Batfish question catalog (examples)

| Question | What it proves | Phase |
|---|---|---|
| `traceroute` / reachability | A can reach B for given header space | pre/post |
| `bidirectionalReachability` | Return path exists | pre/post |
| `undefinedReferences` | No dangling ACL/prefix-list/route-map refs | pre |
| `unusedStructures` | Dead config detection | pre |
| `compareFilters` | ACL equivalence across devices/versions | pre |
| `differentialReachability` | Only intended reachability changed | pre (diff) |
| `bgpSessionStatus` | All BGP sessions Established | post |
| `bgpPeerConfiguration` | Peer config symmetry | pre |
| `ospfSessionCompatibility` | OSPF adjacency compatibility | pre |
| `evpnL3VniReachability` | EVPN L3VNI reachability (where modeled) | pre/post |
| `ipOwners` | IP ownership consistency | pre |
| `detectLoops` | No forwarding loops | pre |
| `multipathConsistency` | ECMP consistency | pre |
| `reachability` under failure | Survives single link/device failure | pre |
| `namedStructures` | Inventory of defined structures | audit |

- Table is `[analysis]` synthesis of the README question catalog; exact question names should be verified in pybatfish docs before scripting `[unverified]`.

### E4.15.2 SuzieQ table inventory (representative)

| Table | Content |
|---|---|
| `device` | Inventory, version, model |
| `interfaces` | State, MTU, speed |
| `lldp` | Neighbor discovery |
| `bgp` | Sessions, prefixes |
| `ospf` | Neighbors, LSAs |
| `routes` | RIB entries |
| `arpnd` | ARP/ND tables |
| `macs` | MAC table |
| `evpnVni` | EVPN/VXLAN state (platform-dependent) |
| `topology` | Derived topology |

- Representative list `[official]` (docs); platform support varies `[unverified]`.

