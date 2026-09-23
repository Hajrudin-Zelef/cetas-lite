---
id: etape6-phased3-bgp-ha/00-bgp-ha/wave-5-file-verification-and-coverage-audit
title: "Wave 5 — File verification and coverage audit"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["benchmarks", "hyperscaler"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [295, 341]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: ea043ed35ce0645449ec7567071b912b3a469de99404a128276c34bff9eedc9e
---

# Wave 5 — File verification and coverage audit

## Wave 5 — File verification and coverage audit

### 5.1 Verification (2026-09-22)

- Line count, tail integrity, and Markdown sanity checked via shell after final append (exact count reported by the writer in the final report).
- All waves append-only; header and earlier sections untouched after writing.
- Every factual claim carries a provenance tag; URLs use verbatim full URLs from search results; no identifiers guessed.

### 5.2 Coverage audit against the Phase D3 scope

| Scope item | Covered in |
|---|---|
| eBGP vs iBGP designs | §1.1 |
| Unnumbered BGP (RFC 5549) | §1.3 |
| ASN planning (private ASNs) | §1.2 |
| ECMP | §1.4 |
| BGP timers/tuning | §1.5 |
| Graceful restart | §1.6 |
| BFD integration | §1.5, §1.7 |
| Prefix lists / route maps | §2.1, §2.2 |
| Communities | §2.3 |
| RPKI status in DC | §2.4 |
| Cisco vPC | §3.1 |
| Arista MLAG | §3.2 |
| Cumulus CLAG | §3.3 |
| Dell VLT | §3.4 |
| Aruba VSX | §3.5 |
| Juniper MC-LAG | §3.6 |
| MC-LAG comparison | §3.7 |
| FHRP (VRRP/HSRP/GLBP) | §4.1 |
| Anycast gateway | §4.2 |
| ISSU/NSF | §4.3 |
| Link failure detection | §4.4 |
| Convergence benchmarks | §4.5 |

### 5.3 Consolidated open items / gaps

1. Production ASN-scheme survey (hyperscaler/enterprise).
2. BFD scale at aggressive timers in production fabrics.
3. Quantified GR benefit in Clos fabrics (lab evidence mixed).
4. RPKI-at-DC-border adoption survey.
5. ISSU support matrix for Cumulus/Dell/Juniper MC-LAG.
6. Independent cross-vendor convergence benchmarks; standardized methodology.
7. Sub-second MC-LAG failover claims are vendor-reported and scenario-specific.

*End of Wave 5. Phase D3 file complete (append-only).*

