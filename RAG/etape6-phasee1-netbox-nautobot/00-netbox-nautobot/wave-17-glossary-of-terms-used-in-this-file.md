---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-17-glossary-of-terms-used-in-this-file
title: "Wave 17 — Glossary of terms used in this file"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2025-05-28", "2025-07", "2026-02", "2026-02-10", "2026-10-13"]
keywords: ["agent", "agentic", "copilot", "governance", "latency", "pricing", "training"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [522, 619]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 90f6e554ddb0da4bb71ecc32ce236c5316e93866a16f8d3803c500d7994622ec
---

# Wave 17 — Glossary of terms used in this file

## Wave 17 — Glossary of terms used in this file

- **Source of truth (SoT):** the authoritative system where the intended/desired state of the network is modeled, validated, and stored; automation, monitoring, and provisioning consume it via API. Both NetBox and Nautobot claim this role `[official]`.
- **SSoT (Single Source of Truth app):** Nautobot's official integration framework that synchronizes Nautobot with external systems (ServiceNow, Infoblox, SolarWinds, DNAC, NetBox, etc.) bidirectionally using diffsync `[official]`.
- **diffsync:** Network to Code's open-source diff/comparison-and-sync engine; powers both SSoT integrations and the NetBox→Nautobot importer `[official]`.
- **Diode:** NetBox Labs' ingestion pipeline — Diode SDK + netboxlabs-diode-netbox-plugin for idempotent, schema-validated data ingestion into NetBox; used by NetBox Assurance and the Discovery agent `[official]`.
- **NetBox Assurance:** commercial NetBox Labs add-on (2025) that continuously reconciles NetBox's intended state against observed network state (drift detection), including a discovery agent `[vendor-reported]`.
- **NetBox Copilot:** NetBox Labs' enterprise AI assistant, GA 2026-02-10 — grounded in the customer's NetBox data, executes workflows, performs data-quality checks and capacity planning, generates Ansible playbooks `[vendor-reported]`.
- **NautobotGPT:** Network to Code's AI co-pilot for Nautobot (launched 2025-05-28) — turns natural language into Nautobot Jobs, troubleshoots Jobs, AI-powered traceback analysis `[vendor-reported]`.
- **Golden Config:** Nautobot app for GitOps network configuration management — config backups, intended-config rendering, source-of-truth aggregation, compliance reporting `[official]`.
- **ChatOps:** Nautobot app that brings network automation commands into Slack/Microsoft Teams/Webex/cisco chat; bridges chat, automation, and NetDevOps collaboration `[official]`.
- **nornir-nautobot:** plugin making Nautobot the inventory source for Nornir automation tasks `[official]`.
- **NetBox Operator / NetBox Copilot:** agentic-AI offerings reported by NetBox Labs; "Operator" appeared in July 2025 coverage, "Copilot" in the February 2026 GA announcement — relationship unverified `[unverified]`.
- **Config contexts:** NetBox's hierarchical JSON/YAML data rendered against devices/VMs for config templating; v4.7 pre-renders and caches them, with an invalidation task on save/delete `[official]`.
- **Event rules:** NetBox's successor mechanism to webhooks for reacting to object changes (v4.7 added snapshot-aware rules and improved webhook robustness) `[official]`.
- **MPTT → ltree:** NetBox v4.7 replaced the django-mptt tree implementation with the PostgreSQL `ltree` extension for hierarchical models (locations, sites) — requires PG 15+ and a maintenance-window migration `[official]`.
- **IPAM:** IP address management (prefixes, IPs, VRFs, VLANs). **DCIM:** data center infrastructure management (racks, devices, cabling, power).
- **InfiniBand HDR/NDR/XDR / SFP112:** interconnect form factors added as interface types in NetBox v4.6.9/v4.7.1 (2026) — evidence of continued hardware-tracking currency `[official]`.
- **Nautobot Cloud:** Network to Code's managed Nautobot offering (details/pricing not captured in this wave) `[unverified]` on specifics.
- **NetBox Cloud:** NetBox Labs' managed SaaS NetBox, with a free tier advertised and enterprise SLAs `[vendor-reported]`.
- **Module bay types:** NetBox v4.7 DCIM model for module bays; **module types:** modular device components; **port mappings:** v4.7 model for multi-protocol (e.g., transceiver) application services `[official]`.
- **Breakout cables:** Nautobot 3.2 DCIM feature modeling channelized/breakout connections; NetBox 4.6.4 added a 1C8P:8C1P breakout profile `[official]`.
- **Custom scripts (NetBox):** deprecated from core in v4.7, moving to a dedicated plugin; removed from core in v5.0 (planned) `[official]`.
- **Approval workflows (Nautobot):** Nautobot 3.0 feature — human approval gates for Jobs and changes `[official]`.
- **Data Validation Engine:** Nautobot's in-core validation framework (integrated in 3.0) for enforcing data-quality rules on objects `[official]`.
- **Live search/typeahead:** Nautobot 3.2 UI feature for instant object search `[official]`.
- **IP Address Range:** Nautobot 3.2 model for contiguous ranges not tied to prefixes `[official]`.
- **NetBox Evolve:** NetBox Labs' inaugural conference (2026-10-13, Kennedy Space Center) marking 10 years of open-source NetBox `[vendor-reported]`.

---

## Wave 18 — API and automation-engine deep dive (release-note-grounded)

### 18.1 NetBox REST API patterns (v4.7, Sept 2026) `[official]`

- **Async bulk writes:** `?background=true` query parameter moves bulk REST writes to a background worker, returning 202 instead of holding the request — designed for large automation pushes (Ansible/Terraform providers) `[official]`.
- **Per-object errors:** bulk operations now return granular, per-object error details instead of failing the whole batch opaquely — important for idempotent automation reconciliation loops `[official]`.
- **Token handling:** v4.5 removed plaintext token retrieval from the API (tokens are hashed server-side); automation must capture the token at creation — breaking change for onboarding scripts that expected to read tokens back `[official]`.
- **Cable model API:** v4.5 made cable-terminations read-only in the REST API (they're derived from the cable object) — consumers that wrote terminations directly must migrate `[official]`.
- **GraphQL (v4.6.9):** queryset `prefetch_related()` optimization for cables; GraphQL had seen broader API changes in v4.5 — integrators pinned to specific NetBox minor versions should regression-test GraphQL queries on upgrade `[official]`.
- **Event system:** v4.7 added snapshot-aware event rules and improved webhook resilience; webhooks enqueued while workers restart can fail with `TypeError` — the officially documented upgrade order is: let workers drain before upgrading `[official]`.

### 18.2 NetBox plugin ecosystem mechanics `[official][secondary]`

- Plugins are first-class: NetBox ships an in-app **plugins catalog** (since v4.1, late 2024); plugins add models, API endpoints, UI views, and background tasks. Canonical examples include Diode integration, netbox-topology-views, LDAP/SSO extensions `[official]`.
- v4.7 introduced `plugin_dependencies` in `PLUGINS_CONFIG` so plugins can declare load-order dependencies on other plugins — relevant for composite stacks (e.g., Diode + Assurance + discovery) `[official]`.
- Plugins with their own `ltree` models need their own migration steps on v4.7 upgrade (mirroring core's mptt→ltree change) `[official]`.
- Lifecycle caveat: core custom scripts are being retired (v4.7 deprecation → dedicated plugin → v5.0 removal), so automation previously embedded in core must be re-hosted as Jobs-style plugins or external orchestration before the v5.0 cutover `[official]`.

### 18.3 Nautobot automation-engine patterns `[official]`

- **Jobs as the execution primitive:** Nautobot Jobs (Python classes) are scheduled, approved (approval workflows since 3.0), and run on distributed workers — including Kubernetes Job Execution (2.4) and console logging (3.1). Approval gates make Jobs the sanctioned vehicle for changes that need human sign-off `[official]`.
- **Secrets groups:** credentials used by Jobs and integrations are stored in secrets groups (with providers) rather than in code or environment — the audit-relevant surface for credential hygiene `[official]`.
- **Event publication (2.4+):** Nautobot publishes events that external systems can consume; combined with webhooks and event rules this gives both platforms comparable reactive automation surfaces, with Nautobot's being worker/Jobs-native `[official]`.
- **Config context rendering API:** Golden Config and Nautobot's Jinja rendering APIs let automation fetch intended configs per device — the Nautobot counterpart to NetBox's config-context rendering pipeline (which v4.7 pre-renders) `[official]`.
- **Job cancel (3.2):** running Jobs can be cancelled from the UI/API — an operational control that matters for long-running provisioning and compliance sweeps `[official]`.

### 18.4 Coexistence patterns with orchestrators (Phase E2 bridge) `[secondary]`

- Reference architectures across the collected material share one shape: **SoT (NetBox/Nautobot) → orchestrator (Ansible/Nornir/Terraform) → devices**, with telemetry and drift (Assurance/Diode, SSoT, Operational Compliance) closing the loop back into the SoT `[secondary]`.
- NetBox's v4.7 async bulk writes and pre-rendered config contexts directly reduce the latency of the SoT→orchestrator leg; Nautobot's Jobs and approval workflows keep the change itself inside the SoT's governance boundary — the two platforms optimize different legs of the same pipeline `[secondary]`.

---

## Wave 19 — Selection guide: when the collected evidence favors which platform `[secondary]`

This section synthesizes only facts already recorded in this file; it introduces no new claims.

### 19.1 Choose NetBox when…

- Your primary need is a **mature, strict data model for inventory truth**: NetBox's decade-long DCIM/IPAM model (with 2026 additions — cooling, channelized subinterfaces, port mappings, transceiver profiles) is the deepest pure-inventory schema in the set `[official]`.
- You want the **largest community surface**: 18,000+ stars / 300+ contributors (2024 figures, latest captured), a broad plugin ecosystem with an in-app catalog, and extensive third-party deployment recipes (Docker, Podman, K3s, linuxserver images) `[secondary]`.
- Your automation is **orchestrator-led (Ansible/Terraform)**: NetBox's API-first design plus v4.7 async bulk writes and pre-rendered config contexts are built for external renderers; Terraform providers and Ansible modules treat NetBox as the data authority `[official][secondary]`.
- You need **enterprise AI-grounded operations**: NetBox Copilot (GA 2026-02-10) and the Assurance/Diode drift pipeline are the commercial add-ons NetBox Labs is investing in (with $55M+ raised and SI partners WWT/AHEAD/Presidio) `[vendor-reported]`.
- You prefer **managed SaaS with a free entry tier**: NetBox Cloud advertises a free tier `[vendor-reported]`.

### 19.2 Choose Nautobot when…

- You want **automation and validation inside the SoT**: native Jobs (schedule/approve/cancel), the in-core Data Validation Engine (3.0), and Golden Config's GitOps compliance loop keep change governance inside one platform `[official]`.
- Your stack is **Python/Nornir-centric**: nornir-nautobot, SSoT diffsync integrations, and ChatOps give a coherent Network-to-Code-native automation fabric `[official]`.
- You need **built-in Git integration and chat-driven operations**: Git repos as data sources, Golden Config, and ChatOps are core/near-core, not third-party add-ons `[official]`.
- You are **migrating from NetBox**: the official `nautobot-app-netbox-importer` (diffsync-based) exists specifically for this direction; community and vendor docs note field-mapping friction (status objects vs strings) so plan a pilot phase `[official]`.
- You value **Network to Code services**: implementation, training, and the commercial Apps (OS Upgrades, Operational Compliance) with professional-services backing `[vendor-reported]`.

### 19.3 Coexistence is a documented pattern

- OpsMill's migration docs (Infrahub adapters for both NetBox and Nautobot) explicitly advise: pilot critical workflows, run incremental syncs, cut over gradually, and **stay side-by-side indefinitely if that is stable** — many teams do. SSoT/Diode-style sync tooling exists for both directions `[official][secondary]`.
- Practical implication: choosing one SoT does not require retiring the other on day one; the sync and importer tooling makes a phased migration or permanent federation the low-risk default `[secondary]`.

### 19.4 Decision checklist (derived from gaps)

- Verify current **pricing** for NetBox Cloud tiers, Assurance entity pricing, Nautobot commercial Apps, and Nautobot Cloud — none were public in this wave `[unverified]`.
- Verify **support SLAs and data-residency** for each SaaS offering before committing regulated workloads `[unverified]`.
- Plan the **v4.7 upgrade window** (ltree migration locks, webhook drain) if selecting NetBox self-hosted; plan **PG ≥ 14 and Django STORAGES migration** for Nautobot 3.1+ upgrades `[official]`.
- For Nautobot adoption, budget for the **Python/Django skill level** Jobs require; for NetBox, budget for the **external orchestrator** that NetBox deliberately does not include `[official][secondary]`.

---

