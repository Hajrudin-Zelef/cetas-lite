---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-21-timeline-netbox-nautobot-january-2024-september-2026
title: "Wave 21 — Timeline: NetBox & Nautobot, January 2024 → September 2026 `[official][vendor-reported]`"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["CoreWeave", "Microsoft"]
dates: ["2024-01", "2024-01-31", "2025-01-06", "2025-04", "2025-04-02", "2025-04-15", "2025-05-28", "2025-06-30", "2025-07", "2026-02", "2026-02-10", "2026-06-30", "2026-07", "2026-08-16", "2026-08-25", "2026-09", "2026-09-02", "2026-09-15", "2026-09-22", "2026-10-13"]
keywords: ["acquisition", "agent", "copilot", "packaging", "pricing", "research", "series b", "valuation"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [703, 750]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 857017993cae1ea96ee56dfd9d3159b0f4741c4f6c82b9471b87a203f9d1abfb
---

# Wave 21 — Timeline: NetBox & Nautobot, January 2024 → September 2026 `[official][vendor-reported]`

## Wave 21 — Timeline: NetBox & Nautobot, January 2024 → September 2026 `[official][vendor-reported]`

Chronology reconstructed from the events cited in this file.

- **2024-01-31** — NetBox Labs enters strategic partnerships with leading network discovery and assurance providers to reduce network-automation adoption barriers (GlobeNewswire) `[vendor-reported]`.
- **2024 (mid)** — NetBox Labs Series B: $35M led by NGP (SiliconANGLE); company positioning: "central nervous system of AI data centers"; 18,000+ GitHub stars, 300+ contributors cited in pitch deck (startupfundraising teardown) `[vendor-reported][secondary]`.
- **2024 (late)** — NetBox v4.1: plugins catalog in UI, circuit groups, VLAN ID ranges, nested modules, rack types `[official]`.
- **2025-01-06** — NetBox v4.5.0: Python 3.12–3.14 required; GraphQL/API/token/cable-termination breaking changes `[official]`.
- **2025-01** — Nautobot 2.4: Apps marketplace page in core, wireless data models, Kubernetes job execution, event publication `[official]`.
- **2025-04-02** — NetBox Labs announces **NetBox Assurance** (general availability): managed drift detection, Diode SDK, Discovery agent (GlobeNewswire) `[vendor-reported]`.
- **2025-04-15** — Network to Code launches commercial **OS Upgrades** and **Operational Compliance** apps (ACCESS Newswire; 80% upgrade-time reduction claimed on pre-release deployments) `[vendor-reported]`.
- **2025-05-28** — Network to Code launches **NautobotGPT**, the first AI co-pilot for Nautobot users (ACCESS Newswire) `[vendor-reported]`.
- **2025-06** — nautobot-app-netbox-importer v2.2.0: pydantic v2.11.5, `--test-input` fixture option `[official]`.
- **2025-06-30** — Golden Config v3.0.7: Nautobot 3.x support, no more empty Git commits on unchanged configs `[official]`.
- **2025 (Nov)** — Nautobot 3.0: Bootstrap 5 UI, approval workflows, load-balancer/VPN models, Data Validation Engine in core `[official]`.
- **2025 (fall)** — Nautobot powers SCinet 2025 at the Supercomputing conference (named high-performance deployment) `[vendor-reported]`.
- **2026-01** — phpIPAM v1.8.x line supports PHP 7.2–8.5 `[official]`; teemIP team announces formation of teemIP SAS (July 2026 per vendor posts) `[vendor-reported]`.
- **2026-02-10** — NetBox Labs announces **general availability of NetBox Copilot** (GlobeNewswire) `[vendor-reported]`.
- **2026-04** — Nautobot 3.1: job console logging, custom-field scoping, async list/search, bulk rename; PostgreSQL ≥ 14; Django 5.2 `STORAGES` migration `[official]`.
- **2026 (mid, ~June)** — NetBox Labs announces **Infrastructure Intelligence Platform**; total raised **>$55M**; backers NGP, Notable Capital, Flybridge, IBM, Salesforce, Two Sigma; SI partners WWT, AHEAD, Presidio; 10,000+ organizations; named customers ARM, CoreWeave, J.P. Morgan, Kaiser Permanente, Riot Games (bizwire.eu) `[vendor-reported]`.
- **2026-06-30** — NetBox v4.6.4: 1C8P:8C1P breakout profile, enum dropdowns, `dns_name` in events `[official]`.
- **2026-07** — Nautobot 3.2: breakout cables, IP Address Range model, Job Cancel, live search/typeahead, modules hierarchy refinements `[official]`; Device Onboarding v5.4, SSoT v4.6 `[official]`.
- **2026-08-16** — phpIPAM v1.8.2 (latest stable) `[official]`; **2026-08-25** — NetBox v4.6.9: InfiniBand 4X types, SFP112, GraphQL cable prefetch `[official]`.
- **2026-09-02** — **NetBox v4.7.0**: ltree replaces django-mptt (PG 15+ required), cooling modeling, channelized subinterfaces, port mappings, background bulk writes, pre-rendered config contexts, custom scripts deprecated from core `[official]`.
- **2026-09-15** — **NetBox v4.7.1**: ltree trigger restore fix, transceiver module-type profiles, InfiniBand 2X (HDR/NDR/XDR) interface types `[official]`.
- **2026-09-22** — Research cutoff for this file.
- **2026-10-13 (upcoming)** — NetBox Evolve inaugural conference, Kennedy Space Center, Orlando — 10 years of open-source NetBox `[vendor-reported]`.

*End of file. 21 waves; append-only; provenance tags audited; no invented URLs, SKUs, versions, or figures.*

## Wave 22 — Follow-up research seeds (post-cutoff work for a later pass)

These are concrete, answerable queries left open by the gaps log — not filler.

1. **NetBox v4.8:** check `docs/release-notes/` in netbox-community/netbox for any release after 4.7.1; if found, record the ltree follow-ups and custom-script plugin name.
2. **NetBox custom-scripts plugin:** locate the dedicated plugin named by the v4.7 deprecation notice (name, repo, release date).
3. **NetBox Labs round history:** reconcile "$35M Series B (2024)" vs "more than $55M raised (2026)" — find the intervening round(s), lead investors, and any valuation.
4. **NetBox Cloud + Assurance pricing:** pull public pricing pages or plan tiers; confirm the free-tier limits and Assurance entity-pricing bands.
5. **NetBox Copilot vs NetBox Operator:** confirm whether the July 2025 "Operator" press referred to the same product GA'd as "Copilot" in February 2026.
6. **NTC commercial Apps pricing/packaging:** confirm how OS Upgrades and Operational Compliance are sold (bundle vs per-app), list prices, and the source of the 80% upgrade-time claim.
7. **Nautobot Cloud pricing and feature gates:** document tiers and what separates Cloud from self-hosted apps.
8. **ChatOps 2026 release:** verify the latest ChatOps release version and date beyond the April 2025 community pin.
9. **RackTables / openDCIM 2026 currency:** confirm latest releases and maintenance status.
10. **Device42 post-Freshworks (2026):** verify current product direction and NetBox/CMDB positioning after the 2024 acquisition.
11. **GitHub metrics refresh:** pull current star/contributor counts for netbox-community/netbox and nautobot/nautobot to replace the 2024 figures.
12. **NetBox Evolve 2026 (Oct 13):** after the event, capture key announcements (Kennedy Space Center, Orlando).

*End of file. Waves 1–8 and 10–22 (21 sections; numbering skips Wave 9 as originally written); append-only; provenance tags audited; no invented URLs, SKUs, versions, or figures.*
