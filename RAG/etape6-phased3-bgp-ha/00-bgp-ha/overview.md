---
id: etape6-phased3-bgp-ha/00-bgp-ha/overview
title: "Phase D3 — BGP underlay and high availability in the data center"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["benchmarks", "nvidia", "research"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [1, 18]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: af4666e000536ae39225d597abd2670e37b2f0fdeed3b200a4a58ca15d04d073
---

# Phase D3 — BGP underlay and high availability in the data center

**Scope:** BGP as the data-center underlay protocol (eBGP vs iBGP designs, unnumbered BGP/RFC 5549, ASN planning, ECMP, timers/tuning, graceful restart, BFD integration); route policy (prefix lists, route maps, communities, RPKI); multi-chassis LAG technologies (Cisco vPC, Arista MLAG, NVIDIA Cumulus CLAG, Dell VSX/VLT equivalents, Aruba VSX, Juniper MC-LAG); redundancy and convergence (FHRP/VRRP/HSRP/GLBP, anycast gateway, ISSU/NSF, link failure detection, convergence benchmarks).

**Research date:** 2026-09-22 (current through this date). **Method:** read-only web research (browser_search, browser_open); no live-browser visits; nothing sent externally. No identifiers guessed. **Deliverable discipline:** append-only; each wave is a numbered section added sequentially.

## Provenance legend

- `[official]` — vendor documentation, RFCs, standards bodies, official release notes.
- `[vendor-reported]` — vendor blogs, white papers, vendor-published benchmarks (not independently reproduced).
- `[independent]` — third-party publications, community lab write-ups with reproducible configs.
- `[secondary]` — press, analysts, aggregators summarizing a primary source.
- `[unverified]` — single-source claims not corroborated, or lab-only observations.

**Known limitations of this file:** convergence benchmarks are vendor-reported or lab-scale and not directly comparable across vendors; RPKI-in-DC adoption claims are thin and mostly anecdotal; some per-version notes depend on documentation crawled in 2026 and should be re-checked against the current release train before deployment. Gaps are flagged inline with `**Gap:**`.

---

