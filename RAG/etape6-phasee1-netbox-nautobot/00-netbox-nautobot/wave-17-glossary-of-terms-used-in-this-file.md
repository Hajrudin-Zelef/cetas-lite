---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-17-glossary-of-terms-used-in-this-file
title: "Wave 17 — Glossary of terms used in this file"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2025-05-28", "2025-07", "2026-02", "2026-02-10", "2026-10-13"]
keywords: ["agent", "agentic", "copilot", "pricing"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [522, 572]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 670d856902878a93c35121985072f0c5b530fe928e1006655451963fd4ecf628
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

