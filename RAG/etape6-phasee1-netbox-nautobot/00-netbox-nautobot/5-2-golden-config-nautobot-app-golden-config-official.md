---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/5-2-golden-config-nautobot-app-golden-config-official
title: "5.2 Golden Config (nautobot-app-golden-config) `[official]`"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["JFrog", "Microsoft"]
dates: ["2025-01", "2025-04", "2025-04-10", "2026-06-30", "2026-07-28", "2026-08-06"]
keywords: ["apache"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [180, 208]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: efae93f7b45169b361c05d99e8498be25078ffc2c9d7aa617f8e1c670a38e803
---

# 5.2 Golden Config (nautobot-app-golden-config) `[official]`

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

