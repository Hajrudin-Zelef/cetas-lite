---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-4-nautobot-core-version-line-and-platform-features-2023
title: "Wave 4 — Nautobot core: version line and platform features (2023–2026)"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["JFrog", "Microsoft"]
dates: ["2024-03", "2024-08", "2025-01", "2025-04", "2025-04-10", "2025-11", "2026-04", "2026-04-15", "2026-06-30", "2026-07", "2026-07-01", "2026-07-28", "2026-08-06", "2026-08-13", "2026-09-22"]
keywords: ["apache", "distribution"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [136, 208]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 93aa9f4150f25594b02e19f3d8a9caeecbc63c2401bfdb20aedc3e1f3a58becb
---

# Wave 4 — Nautobot core: version line and platform features (2023–2026)

## Wave 4 — Nautobot core: version line and platform features (2023–2026)

### 4.1 Version timeline (latest first, as of 2026-09-22)

- **Nautobot 3.2 — July 2026** `[official]` — latest core release as of 2026-09-22. Headline features: Breakout Cables; IP Address Range Model; Job Cancel; Model Search Typeahead; Live Search; Homepage Stickiness; Modules Hierarchy Refinements. Source: https://docs.nautobot.com/projects/core/en/latest/release-notes/
- **Nautobot 3.1 — April 2026** `[official]` — Job Console Logging; Custom Field Scoping; Dependent Object Creation and Search; Tree List Views on Location and Prefix; Async Loading on all List Views; Async Loading on Global Search; Enhanced Configurable Columns; Bulk Rename Across All UI ViewSets.
  - Upgrade notes: requires PostgreSQL ≥ 14 (drops 12.x/13.x) via the Django 5.2 dependency upgrade; Django `DEFAULT_FILE_STORAGE`/`STATICFILES_STORAGE` replaced by unified `STORAGES` setting (Nautobot-specific `STORAGE_BACKEND`, `STORAGE_CONFIG`, `JOB_FILE_IO_STORAGE` merged into it); app cookie adds Python 3.14 support. Source: https://github.com/madtechservices/nautobot/blob/HEAD/nautobot/docs/release-notes/version-3.1.md and https://github.com/nautobot/cookiecutter-nautobot-app/blob/HEAD/docs/admin/release_notes/version_3.1.md
- **Nautobot 3.0 — November 2025** `[official]` — major UI refresh (Bootstrap 5 migration, new navigation, global search); Approval Workflow system for multi-stage job and change approvals; Load Balancer Models; VPN Models; Device Uniqueness Updates; **Data Validation Engine integrated into core**; Enhanced Search and updated Saved Views; ECharts integration for custom charting; Updated GraphQL UI and API.
- **Nautobot 2.4 — January 2025** `[official]` — Virtual Device Context Data Models; Wireless Data Models; **Apps Marketplace Page**; Event Publication Framework; Jinja2 Template Rendering REST API; Kubernetes Job Execution and Job Queue Data Model; UI Component Framework.
- **Nautobot 2.3 — August 2024** `[official]` — Cloud Models; Device Modules Model; Dynamic Group Enhancements; Interface Roles; Object Metadata Framework & Models; Saved Views; Worker Status Page.
- **Nautobot 2.2 — March 2024** `[official]` — Contacts & Teams Models; Controller Model; DeviceFamily Model; Prefix and VLAN Many Locations; Software Image File and Software Version Models; Syntax Highlighting.

### 4.2 Core data-model and automation features (Nautobot 3.x)

- Source-of-truth models: locations (tree views), devices, device types/modules, racks, cables, interfaces (incl. breakout cables in 3.2), IPAM (prefixes, IP addresses, new IP Address Range model in 3.2, VRFs, VLANs), circuits, wireless, cloud models, VPN models, load balancer models, virtual device contexts, software versions/image files, controllers, contacts & teams `[official]`.
- **Jobs framework** (core differentiator vs NetBox): Python jobs with scheduling, approval workflows (3.0), job console logging (3.1), job cancel (3.2), Kubernetes job execution (2.4) `[official]`.
- **Data Validation Engine in core** (3.0): declarative validation rules enforced on data changes `[official]`.
- **Event Publication Framework** (2.4): publish change events to external systems `[official]`.
- APIs: REST, GraphQL (updated UI/API in 3.0), Jinja2 template rendering REST API (2.4) `[official]`.
- UI: Bootstrap 5 refresh (3.0), async loading everywhere (3.1), live search/typeahead (3.2), saved views, customizable home page panels, ECharts custom charting `[official]`.

### 4.3 Network to Code commercial launch with Nautobot 3.1 (2026-04-15) `[vendor-reported]`

Source: https://www.morningstar.com/news/accesswire/1157492msn/network-to-code-launches-commercial-network-automation-software-with-nautobot-31

- **Network to Code** (NTC, described as "a leader in AI-powered network automation", CEO Paul Brady, announcement dateline New York City) announced general availability of its **first major wave of commercial software applications** with Nautobot 3.1, expanding the platform with new product bundles and applications for automating network change at scale, validating results, and retaining operational evidence across multi-vendor environments.
- Two flagship commercial applications at the center: **OS Upgrades** and **Operational Compliance** — together they help teams execute change, validate results before/after upgrades, identify operational drift, and retain evidence. Across pre-release customer deployments, the applications **reduced upgrade time by 80%** while delivering complete compliance records `[vendor-reported]` (customer names and sample sizes not disclosed — see gaps log).
- Positioning: "an important step in our open-core strategy and a more practical path for customers to turn automation into operational outcomes" (Paul Brady). The announcement frames enterprise pain as repeatable/safe/provable change across thousands of devices, operational risk, audit needs, and manual verification burden.
- Tooling signal: the Nautobot app cookiecutter added a `nautobot-app-commercial` template (2026-08-06) for creating commercial (licensed) Nautobot Apps distributed through **NTC's private Artifactory repository** — confirming the commercial-apps distribution channel `[official]`.

### 4.4 Nautobot App ecosystem (2026 releases) `[official]`

- **nautobot-app-device-onboarding v5.4** (2026-07-01/27): onboard switch stacks as Virtual Chassis objects (Cisco IOS/IOS-XE); stricter device matching (hostname + serial by default); `Sync VRF to Prefix` toggle. Source: https://github.com/nautobot/nautobot-app-device-onboarding/blob/HEAD/docs/admin/release_notes/version_5.4.md
- **nautobot-app-ssot v4.6** (2026-07-28): new **Cisco SD-WAN (vManage) integration** (devices, device types, software versions, interfaces, IPs, VRFs); parallel loading of Systems of Record (30–50% faster, opt-in in v4.1); integrations cover LibreNMS, Arista CloudVision, Infoblox, Meraki, ServiceNow, SolarWinds. Source: https://github.com/nautobot/nautobot-app-ssot/blob/HEAD/docs/admin/release_notes/version_4.6.md
- **nornir-nautobot v4.4.0** (2026-08-13): Jinja template rendering with `substitute_lines` in `get_config`; dependency updates. Source: https://github.com/nautobot/nornir-nautobot/releases/tag/v4.4.0
- **nautobot-app-circuit-maintenance v3.1** (2026-03/04): Django 5.2 compatibility (removed django-cryptography). Source: https://github.com/nautobot/nautobot-app-circuit-maintenance/blob/HEAD/docs/admin/release_notes/version_3.1.md
- **nautobot-app Golden Config** and **nautobot-app ChatOps** are the long-standing flagship automation apps (config generation/compliance/remediation; ChatOps integrations for Slack/MS Teams/Mattermost/Webex) — current 2026 version status to be confirmed in the next wave `[secondary]` (see gaps log).

---

## Wave 5 — Nautobot Apps: marketplace and flagship automation apps

### 5.1 Apps marketplace

- Nautobot has shipped an **Apps Marketplace Page** in core since v2.4 (January 2025) for discovering installable apps `[official]`.
- Representative 3.x app stack (community-pinned example, April 2025, updated for 3.x line): plugin-nornir 3.2.0, device-lifecycle-mgmt 4.1.1, ssot 4.2.2, golden-config 3.0.5, **chatops 4.0.0** with Slack/Teams/Webex/Ansible/Arista extras; also data-validation (2.x only — now core), firewall-models, floorplan, bgp-models, secrets-providers, design-builder. ChatOps was bundled again on 3.x images after being dropped from 1.x/2.x stacks over dependency clashes `[secondary]`. Source: https://github.com/netdevops-it/netdevops.it.github.io/blob/HEAD/docs/blog/posts/2025/2025-04-10-nautobot-docker-apps.md
- App development tooling: cookiecutter `nautobot-app` template (v3.1.4, 2026-08-06) with Python 3.14 support and a `nautobot-app-commercial` template for licensed apps distributed via NTC's private Artifactory `[official]`.
- Software Lifecycle Policy published at https://networktocode.com/company/legal/software-lifecycle-policy/ — e.g. Golden Config supports the previous major only for Nautobot LTM 2.4 users with critical bug/security fixes `[official]`.

### 5.2 Golden Config (nautobot-app-golden-config) `[official]`

- Purpose: NetDevOps approach to golden configuration and configuration compliance. Four key use cases: (1) Configuration Backups — Nornir process connects to devices, optionally parses out lines/secrets, backs up config to a Git repo; (2) Intended Configuration — Nornir generates config from Jinja templates in Git combined with GraphQL-sourced data, stored to a Git repo; (3) Source of Truth Aggregation — per-device GraphQL query producing the data structure used for generation; (4) Configuration Compliance — Nornir comparison of actual (backups) vs intended CLI configs. The features are decoupled: e.g. RANCID/Oxidized backups can feed compliance without using the backup job.
- Maintainers: Ken Celenza (@itdependsnetworks), Jeff Kala (@jeffkala).
- Nautobot features used: Dynamic Groups, Jobs, Job Buttons, Secret Groups, Git Repositories, Git as a Data Source, GraphQL Saved Queries.
- Versions: **v3.0.7 (2026-06-30)** — Nautobot 3.0 compatibility, min Nautobot 3.0, Python 3.13, Bootstrap 5.3 UI, Apache ECharts rendering; backup/intended jobs no longer create empty Git commits when configs are unchanged. v2.5 (2025-09): DeepDiff upgrade for CVE-2025-58367, heir_config v2→v3, remediation settings manageable via Git data sources. v2.4: first iteration of Config Plans with post-processing (view/approve config plans, pre-deployment rendering of post-processed config).
- Sources: https://github.com/nautobot/nautobot-app-golden-config/blob/HEAD/docs/admin/release_notes/version_3.0.md ; https://github.com/nautobot/nautobot-app-golden-config/blob/HEAD/docs/user/app_overview.md ; https://pypi.org/project/nautobot-golden-config/1.2.0/

### 5.3 ChatOps (nautobot-app-chatops) `[official][secondary]`

- ChatOps app enables chatbot-driven network operations from Slack, Microsoft Teams, Webex (and Mattermost historically), with command integrations including Ansible and Arista-specific commands; v4.0.0 targets Nautobot 3.x with Slack/Teams/Webex/Ansible/Arista extras `[secondary]`.
- Classic use cases: on-call engineers run Nautobot lookups, device commands, and automation jobs from chat; approval-gated workflows from chat. Source: https://github.com/netdevops-it/netdevops.it.github.io/blob/HEAD/docs/blog/posts/2025/2025-04-10-nautobot-docker-apps.md
- Note: exact latest ChatOps version/date not re-verified in this wave — see gaps log `[unverified]` for currency.

### 5.4 Other notable apps `[official][secondary]`

- **SSoT (Single Source of Truth)** v4.6 (2026-07-28): sync framework between Nautobot and external systems — Cisco SD-WAN/vManage (new 2026), Infoblox, ServiceNow, Arista CloudVision, LibreNMS, Meraki, SolarWinds; parallel loading (30–50% faster).
- **Device Onboarding** v5.4 (2026-07): network-driven discovery into Nautobot (IOS/IOS-XE virtual chassis support).
- **Device Lifecycle Management**: hardware/software lifecycle tracking.
- **BGP Models, Firewall Models, Floor Plan** (rack elevation visualization), **Design Builder** (declarative design-as-code), **Secrets Providers**.
- **Circuit Maintenance**: maintenance notification ingestion.

---

