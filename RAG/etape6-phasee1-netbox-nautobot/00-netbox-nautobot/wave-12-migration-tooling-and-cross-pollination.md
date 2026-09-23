---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-12-migration-tooling-and-cross-pollination
title: "Wave 12 — Migration tooling and cross-pollination"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2021-02-20", "2025-05-28", "2025-06-13", "2026-01-06", "2026-02-10", "2026-04", "2026-06-30", "2026-07", "2026-08-16", "2026-08-25", "2026-09-02", "2026-09-15", "2026-09-22"]
keywords: ["agent", "apache", "copilot", "governance", "license", "open source", "revenue", "training"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [380, 456]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: b9ecb122e7fd4a306e19eded6d8d316dff9fe2a2a84564c6d76ff1748d3b028e
---

# Wave 12 — Migration tooling and cross-pollination

## Wave 12 — Migration tooling and cross-pollination

### 12.1 nautobot-app-netbox-importer (first-party NetBox→Nautobot migration app) `[official]`

- Nautobot ships an official migration app: **nautobot-app-netbox-importer** — "Nautobot plugin to simplify data migration from NetBox", built on **diffsync**. Docs: https://docs.nautobot.com/projects/netbox-importer/en/latest/
- Repository: https://github.com/nautobot/nautobot-app-netbox-importer — created 2021-02-20; ~24 stars, 12 forks (as crawled); topics: diffsync, nautobot, nautobot-plugin, netbox. Sources: https://awesome.ecosyste.ms/projects/github.com%2Fnautobot%2Fnautobot-app-netbox-importer ; https://github.com/nautobot/nautobot-app-netbox-importer/releases/tag/v2.2.0
- **v2.2.0 (2025-06-13):** pydantic v2.11.5 bump; `--test-input` option for `import-netbox` to use test fixtures; contributors @gsnider2195, @smk4664, @snaselj, @HanlinMiao.
- Compatibility: supports Nautobot 3.0+ (min 3.0.0, max <4.0.0), Python 3.13 (as of the 2026 cookie updates).
- Operational notes from its release history: sample data generation uses already-sampled IDs; test failures converted to warnings for certain cases — signals of the practical friction in field mapping (status objects vs strings, nested structures) also noted by Infrahub's migration docs `[official][secondary]`.

### 12.2 Cross-pollination between ecosystems `[secondary]`

- **nautobot-topology-views** (rsp2k/nautobot-topology-views): a port of NetBox's `netbox-topology-views` to Nautobot — draws topology maps from Nautobot cable data, filters on name/location/tag/device role, exports to draw.io XML or PNG; requires Nautobot 3.1.3+, Python 3.12+, Django 5.2. Evidence that the plugin/app ecosystems track each other's popular extensions. Source: https://github.com/rsp2k/nautobot-topology-views
- **Infrahub adapters** (OpsMill): dedicated NetBox and Nautobot adapters for migrating either system into Infrahub; migration playbook phases (pilot workflows → incremental sync → cutover → optional retirement); docs explicitly state many teams reach a stable side-by-side state and stay there indefinitely `[official]`. Source: https://github.com/opsmill/infrahub-sync/blob/HEAD/docs/docs/migrating-from-netbox-or-nautobot.mdx
- Community learning: `nautobot/100-days-of-nautobot` includes explicit compare/contrast notes of Nautobot Jobs vs NetBox functionality for NetBox-versed users (Day002) `[official]`.

### 12.3 Practical migration considerations `[secondary]`

- Schema friction points documented across sources: Nautobot status fields are objects vs NetBox status strings; differing nested structures; custom-field semantics; cable/path modeling differences; the importer app's fixture-based testing exists precisely because of these.
- Direction asymmetry: the mature, maintained path is **NetBox → Nautobot** (official importer app); no equivalent first-party **Nautobot → NetBox** importer was found in this wave — teams moving that direction rely on CSV/API export-import or SSoT syncs `[unverified]` on completeness.

---

## Wave 13 — AI assistants and professional services

### 13.1 NautobotGPT vs NetBox Copilot (the 2025–2026 AI race) `[vendor-reported]`

- **NautobotGPT** — launched **2025-05-28** (New York, ACCESS Newswire) `[vendor-reported]`. "AI-Powered Co-Pilot for Network Engineers to Accelerate Automation"; billed as "the first and only AI assistant built for network teams using its open source Network Source of Truth and Automation Platform, Nautobot". Capabilities: on-demand Nautobot expert; turns natural-language questions into Nautobot Jobs; creates, edits, and troubleshoots Nautobot Jobs without writing code; AI-powered traceback analysis with step-by-step guidance; answers grounded in "proprietary Nautobot insights and practices". Positioning: "fastest path from network idea to automation — reducing testing, creation and troubleshooting tasks from hours to minutes". Source: https://www.accessnewswire.com/newsroom/en/computers-technology-and-internet/network-to-code-launches-nautobotgpt-ai-powered-co-pilot-for-netw-1032322
- **NetBox Copilot** — GA **2026-02-10** `[vendor-reported]`. Interactive AI agent embedded in the NetBox platform; grounded in the customer's NetBox infrastructure data ("comprehensive semantic map"); inquiry + workflow execution; natural-language data-quality checks, change investigation, impact assessment, capacity planning, Ansible playbook generation; enterprise governance. Source: https://www.globenewswire.com/news-release/2026/02/10/3235154/0/en/NetBox-Labs-Announces-General-Availability-of-NetBox-Copilot-Delivering-Enterprise-Ready-AI-Grounded-in-Accurate-Infrastructure-Data.html
- Comparison notes `[secondary]`: NautobotGPT shipped ~9 months earlier and centers on **job creation/troubleshooting** (automation authoring); NetBox Copilot centers on **grounded inquiry + workflow execution** (operational Q&A with governance). Both vendors frame the source of truth as the grounding that makes AI trustworthy in production — a notable 2026 positioning convergence. Independent head-to-head evaluation not found (see gaps log).

### 13.2 Professional services and community programs `[vendor-reported][secondary]`

- **Network to Code:** professional services (network automation consulting, Nautobot implementation), training programs, and the commercial Apps portfolio (OS Upgrades, Operational Compliance); 52% recurring revenue growth in 2025 across SaaS and services; Guidepost Growth Equity backing; CEO Paul Brady `[vendor-reported]`. Sources: https://www.lelezard.com/en/news-22091711.html ; https://www.morningstar.com/news/accesswire/1157492msn/network-to-code-launches-commercial-network-automation-software-with-nautobot-31
- **NetBox Labs:** professional services, NetBox Cloud onboarding, enterprise support tiers; community programs include the **NetBox Heroes** podcast/blog series spotlighting practitioners (e.g. Michael Kutka / Insight Global) and the canonical plugins catalog `[vendor-reported]`. Source: https://netboxlabs.com/blog/netbox-heroes-michael-kutka/
- Hiring/market signal: both platforms appear as desired skills in network-automation job postings; NetBox's larger community footprint (18k+ stars, 300+ contributors per 2024 deck) vs Nautobot's 50k+ installs claim — different denominators (stars vs installs), not directly comparable `[unverified]` on current 2026 numbers.

---

## Wave 14 — Consolidated reference tables

### 14.1 Version and platform-requirements matrix (as of 2026-09-22) `[official]`

| Platform | Version | Released | Python | PostgreSQL | Other deps | Headline changes |
|---|---|---|---|---|---|---|
| NetBox | 4.7.1 | 2026-09-15 | 3.12–3.14 | 15+ | Redis 6.0+, Django 6.1 | ltree-trigger restore fix; transceiver module-type profile; IB 2X (HDR/NDR/XDR) types |
| NetBox | 4.7.0 | 2026-09-02 | 3.12–3.14 | 15+ | Redis 6.0+, Django 6.1 | Cooling model; channelized subinterfaces; port_mappings; module bay types; background bulk writes; pre-rendered config contexts; ltree replaces django-mptt |
| NetBox | 4.6.9 | 2026-08-25 | 3.12–3.14 | 14+ (15 recommended) | Redis 6.0+ | InfiniBand 4X types; 100GBase-X-SFP112; GraphQL cable prefetch |
| NetBox | 4.6.4 | 2026-06-30 | 3.12–3.14 | 14+ | Redis 6.0+ | 1C8P:8C1P breakout profile; enum dropdowns; dns_name in events |
| NetBox | 4.5.0 | 2026-01-06 | 3.12–3.14 | 14+ | Redis 5+ | render_config permission; token plaintext retrieval removed; cable-terminations read-only |
| NetBox | 4.1 | late 2024 | 3.10+ | 14+ | — | Circuit groups; VLAN ID ranges; nested modules; rack types; plugins catalog in UI |
| Nautobot | 3.2 | July 2026 | 3.12+ (3.14 in cookie) | 14+ | Django 5.2 | Breakout cables; IP range model; job cancel; live search/typeahead |
| Nautobot | 3.1 | April 2026 | 3.12+ | 14+ (drops 12/13) | Django 5.2, STORAGES | Job console logging; custom field scoping; async loading; bulk rename |
| Nautobot | 3.0 | Nov 2025 | 3.10+ | 12+ | Bootstrap 5 | UI refresh; approval workflows; LB/VPN models; Data Validation Engine in core |
| Nautobot | 2.4 | Jan 2025 | 3.8+ | 12+ | — | Apps marketplace page; wireless models; K8s job execution; event publication |
| phpIPAM | 1.8.2 | 2026-08-16 | 7.2–8.5 | MySQL 8.0+/MariaDB 10.2.1+ | — | Latest stable IPAM release |

### 14.2 Capability matrix (condensed) `[official][secondary]`

| Capability | NetBox (v4.7) | Nautobot (3.2) |
|---|---|---|
| License | Apache 2.0 | Apache 2.0 |
| IPAM | Prefixes, IP ranges (new), VRFs, VLANs w/ multi-range groups, FHRP | Prefixes, IP addresses, IP range model (3.2), VRFs, VLANs |
| DCIM | Racks/rack types, devices, modules (nested), cables, power + **cooling (4.7)** | Locations tree, devices, modules, racks, cables, cloud models |
| Jobs/automation engine | Custom scripts **deprecated** → plugin; event rules/webhooks; background bulk API | Native Jobs (schedule, approve, K8s); Data Validation Engine; event publication |
| Git integration | Git data sources (via ecosystem) | Git repositories as data source (core); Golden Config GitOps |
| Config generation | Config templates (Jinja2) | Golden Config app (backup/intended/compliance); Jinja render API |
| Drift detection | NetBox Assurance (commercial add-on) | Operational Compliance app (commercial, 2026) |
| AI assistant | NetBox Copilot (GA 2026-02-10) | NautobotGPT (launched 2025-05-28) |
| ChatOps | Via ecosystem | ChatOps app (Slack/Teams/Webex) |
| Discovery | NetBox Discovery agent + Diode SDK | Device Onboarding app; SSoT integrations |
| Marketplace | Plugins catalog in UI (v4.1+) | Apps Marketplace page (core since 2.4) |
| Managed SaaS | NetBox Cloud (free tier advertised) | Nautobot Cloud |
| Migration path | — | Official NetBox→Nautobot importer app (diffsync) |

---

