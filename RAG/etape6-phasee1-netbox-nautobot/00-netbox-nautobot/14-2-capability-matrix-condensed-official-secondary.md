---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/14-2-capability-matrix-condensed-official-secondary
title: "14.2 Capability matrix (condensed) `[official][secondary]`"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2025-05-28", "2026-01-06", "2026-02-10", "2026-04", "2026-06-30", "2026-07", "2026-08-16", "2026-08-25", "2026-09-02", "2026-09-15"]
keywords: ["agent", "apache", "copilot", "license"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [423, 456]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 9362619c4855a9979d2a1f2edadc6ba4cb554fcd3bc9ca16d3c8710461ce4197
---

# 14.2 Capability matrix (condensed) `[official][secondary]`

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

