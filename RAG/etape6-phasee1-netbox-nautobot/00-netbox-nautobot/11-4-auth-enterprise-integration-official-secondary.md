---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/11-4-auth-enterprise-integration-official-secondary
title: "11.4 Auth & enterprise integration `[official][secondary]`"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2021-02-20", "2025-05-28", "2025-06-13", "2026-02-10", "2026-09-22"]
keywords: ["agent", "copilot", "governance", "open source", "revenue", "training"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [369, 422]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 7782812ff4cca652244c85384e3b0ab54e88d3f7ff9aa13d476527f85056a874
---

# 11.4 Auth & enterprise integration `[official][secondary]`

- Reference deployment: Docker Compose and Kubernetes (Helm); documented K8s workflow with Flux, custom image layering apps (Golden Config + Nornir plugin via `requirements.txt`), and `nautobot_config.py` management. Nautobot 2.4 added a Kubernetes Job Execution and Job Queue data model `[official]`. Source: https://networktocode.com/blog/deploying-nautobot-to-kubernetes-03/
- **Nautobot 3.1 upgrade notes:** requires PostgreSQL ≥ 14 (drops 12.x/13.x); Django 5.2 unified `STORAGES` setting replaces `DEFAULT_FILE_STORAGE`/`STATICFILES_STORAGE` and Nautobot's `STORAGE_BACKEND`/`STORAGE_CONFIG`/`JOB_FILE_IO_STORAGE`; Python 3.14 supported in the app cookiecutter `[official]`.
- Jobs run on workers with scheduling, approval workflows (3.0), console logging (3.1), and cancel (3.2) — the operational surface for automation is the Jobs UI/API rather than shell scripts `[official]`.

### 11.4 Auth & enterprise integration `[official][secondary]`

- NetBox: SSO via python-social-auth (OIDC/SAML; v6.0/v5.1 in v4.7 — retest before upgrading), LDAP, remote-auth headers; object-level permissions with constraints; `render_config` permission (v4.5+) `[official]`.
- Nautobot: SSO, approval workflows for jobs/changes (3.0), Data Validation Engine in core (3.0), secrets groups for credential management `[official]`.

---

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

