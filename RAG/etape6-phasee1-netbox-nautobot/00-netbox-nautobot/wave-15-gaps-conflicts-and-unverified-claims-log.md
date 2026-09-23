---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-15-gaps-conflicts-and-unverified-claims-log
title: "Wave 15 — Gaps, conflicts, and unverified claims log"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["CoreWeave", "JFrog", "Microsoft"]
dates: ["2024-07", "2025-04", "2025-07", "2026-02", "2026-07", "2026-09-22", "2026-10", "2026-10-13"]
keywords: ["acquisition", "agentic", "agents", "apache", "benchmark", "copilot", "funding", "gpu", "lean", "open source", "packaging", "pricing"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [457, 521]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: d8193b653a5c3d7ef9b6c2ebb0219cf90916ca68ac79d99e4492bf9f20680189
---

# Wave 15 — Gaps, conflicts, and unverified claims log

## Wave 15 — Gaps, conflicts, and unverified claims log

### 15.1 Open items (not found / not verified in this research wave)

1. **NetBox Labs funding beyond Series B:** the $35M Series B (2024) is documented; any 2025–2026 raises, valuation, or revenue figures were not captured — open `[unverified]`.
2. **NetBox Cloud pricing:** no public list prices found; Assurance add-on is "tiered pricing based on ingested entities per month" (tiers undisclosed); free Cloud tier limits not captured `[unverified]`.
3. **NetBox "Operator" vs "Copilot" naming:** July 2025 press called the agentic AI tool "NetBox Operator"; February 2026 GA announcement calls it "NetBox Copilot" — rename vs distinct product unconfirmed `[unverified]`.
4. **Nautobot commercial Apps pricing:** OS Upgrades / Operational Compliance list prices and packaging not disclosed; "80% upgrade time reduction" claim lacks disclosed customer names/sample sizes `[unverified]` on specifics.
5. **Nautobot Cloud pricing/feature gates:** not captured `[unverified]`.
6. **Exact ChatOps latest version/date:** v4.0.0 cited from a community app-stack pin (April 2025); 2026 release currency not re-verified `[unverified]`.
7. **RackTables / openDCIM 2026 release currency:** not re-verified `[unverified]`.
8. **Device42 post-Freshworks-acquisition status (2026):** acquisition widely reported in 2024; current product direction not verified in this wave `[unverified]`.
9. **Independent NetBox-vs-Nautobot benchmark:** no independent performance/scale head-to-head found; comparisons are vendor-adjacent or practitioner opinion `[unverified]` as objective data.
10. **Nautobot→NetBox migration tooling:** no first-party importer found; completeness of community approaches unknown `[unverified]`.
11. **Current GitHub star/contributor counts (Sept 2026):** the 18,000+ stars / 300+ contributors figures are from the 2024 Series B deck; current counts not pulled `[unverified]` as current.
12. **Nautobot 3.2 patch releases after July 2026:** not enumerated `[unverified]`.
13. **NetBox v4.8:** not observed as of 2026-09-22 (absence of evidence, not evidence of absence).

### 15.2 Conflicts and cautions registered

- **C1 — Slashdot "$7,500/year" for Nautobot:** aggregator-listed price, not vendor-confirmed; NTC's actual commercial packaging (2026) is apps/bundles via private Artifactory — treat the Slashdot figure as unreliable `[unverified]`.
- **C2 — "Thousands of enterprises" / "50,000+ installs":** both are vendor-reported with different denominators (enterprises vs installs); not comparable, not independently audited.
- **C3 — Anycast-RP/MSDP-style conflicts:** none in this file's scope; protocol conflicts belong to Phase D files.
- **C4 — NetBox custom scripts:** core deprecation (v4.7, removal in v5.0) in favor of a dedicated plugin is official, but the plugin's name/home was not captured — open item.
- **C5 — "Forked in 2021" history:** the 2026 video states Nautobot forked from NetBox in 2021 to add app framework/jobs/GraphQL; widely repeated and consistent with project history `[secondary]` — accepted with the usual single-practitioner-source caution.

### 15.3 Method notes

- All URLs in this file are verbatim from search-result `Full URLs` sections or fetched pages; no URLs were constructed or guessed.
- Release-note facts are `[official]` (GitHub release-notes docs, vendor docs/press); company metrics and product claims from press releases are `[vendor-reported]`; practitioner comparisons, homelab docs, and engineering blogs are `[secondary]`; single-source or uncorroborated items are `[unverified]`.
- No live-browser visits, no forms, no sign-ins; nothing was sent externally.

*End of Phase E1 — NetBox & Nautobot. File is append-only; earlier waves untouched.*

## Wave 16 — NetBox Labs in 2026: platform, funding, and community signals

### 16.1 "Infrastructure Intelligence Platform" announcement (mid-2026) `[vendor-reported]`

- In mid-2026 (~June, ~103 days before 2026-09-22), NetBox Labs announced an **Infrastructure Intelligence Platform** with new capabilities "spanning entire network and infrastructure lifecycle". Source: https://bizwire.eu/netbox-labs-announces-infrastructure-intelligence-platform-with-new-capabilities-spanning-entire-network-and-infrastructure-lifecycle/
- CEO and cofounder **Kris Beevers** framing: *"Ten years ago, NetBox solved a critical problem by giving infrastructure teams an authoritative source of truth. But today, enterprises of all maturity levels need more than just a source of truth. They need a system of record that delivers a trusted, continuously updated understanding of infrastructure that both humans and AI can operate against safely and confidently. Whether just starting out with agents or well on your way to an automated infrastructure management stack, NetBox Labs is the foundational platform for running infrastructure at any scale."*
- The same statement commits to continued open-source investment alongside commercial expansion: "NetBox Labs will always invest in open source, while continuing to quickly expand our commercial platform".
- Note the positioning evolution visible across this file: 2025-2026 announcements moved from "central nervous system of the AI data center" (July 2025) to "infrastructure intelligence platform" (2026) — NetBox Labs increasingly frames NetBox as infrastructure data that AI operates against, not just network inventory `[vendor-reported]`.

### 16.2 Funding and investor picture (2026) `[vendor-reported]`

- NetBox Labs has **raised more than $55 million** from investors (cumulative, per the 2026 announcement) — vs the $35M Series B reported in July 2024. The delta between the two figures implies further raises in 2024–2026, but the rounds/timing were not itemized in this source; earlier summary item 1 in the gaps log remains open on the breakdown `[vendor-reported][unverified]` on round specifics.
- Named backers: **NGP, Notable Capital, Flybridge, IBM, Salesforce, Two Sigma**. Strategic system-integrator partnerships: **WWT, AHEAD, Presidio** (global SI ecosystem) `[vendor-reported]`.
- Named enterprise customers (2026 announcement): **ARM, CoreWeave, J.P. Morgan, Kaiser Permanente, Riot Games**; platform "trusted by 10,000+ organizations for more than a decade". Headquartered in New York City `[vendor-reported]`.
- Cross-reference: CoreWeave (AI cloud / GPU fleet) and SFMIX (public peering exchange) also appear in named case-study material — the AI-data-center story and the public-infrastructure story both lean on NetBox as source of truth `[vendor-reported][secondary]`.

### 16.3 NetBox Evolve conference (October 2026) `[vendor-reported]`

- To mark the 10th birthday of open source NetBox (first released 2016), NetBox Labs hosts its inaugural community/customer/partner conference **NetBox Evolve on October 13, 2026, at the Kennedy Space Center, Orlando, Florida** (netboxevolve.com) `[vendor-reported]`.

### 16.4 Design philosophy and application stack (official docs) `[official]`

- Official introduction docs state three design tenets:
  1. **Replicate the real world** — strict data model (e.g., IP addresses assigned to interfaces, not devices; an interface may hold multiple IPs).
  2. **Serve as a "Source of Truth"** — NetBox holds the **desired state**, not operational state; automated import of live network state is **strongly discouraged**; all data should be human-vetted before entry so downstream tools can populate with high confidence.
  3. **Keep it simple** — the 80% solution is favored over a complete but complex one; low learning curve, lean codebase.
- Official application stack: HTTP service (nginx or Apache) → WSGI (gunicorn or uWSGI) → Django/Python → **PostgreSQL 15+** → task queue **Redis/django-rq**. NetBox does not talk to network nodes directly; it makes data available programmatically to automation, monitoring, and assurance tools (separation of duties; swap tools without changing the data authority) `[official]`. Source: https://github.com/netbox-community/netbox/blob/HEAD/docs/introduction.md
- README (2026): "successor to legacy IPAM and DCIM applications... central source of truth for the modern network" `[official]`. Source: https://github.com/netbox-community/netbox/blob/HEAD/README.md

---

