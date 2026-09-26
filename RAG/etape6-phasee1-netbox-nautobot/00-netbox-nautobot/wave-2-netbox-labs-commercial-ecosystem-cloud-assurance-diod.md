---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/wave-2-netbox-labs-commercial-ecosystem-cloud-assurance-diod
title: "Wave 2 — NetBox Labs commercial ecosystem: Cloud, Assurance, Diode, Discovery"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: []
dates: ["2025-04-02", "2026-09-22"]
keywords: ["agent", "agents", "cost", "funding", "pricing", "research", "valuation"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [78, 139]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 5872f7912dff6486a96c709ae2767bd91e1b2d0482ec2d5132ac3d6ac29a6531
---

# Wave 2 — NetBox Labs commercial ecosystem: Cloud, Assurance, Diode, Discovery

## Wave 2 — NetBox Labs commercial ecosystem: Cloud, Assurance, Diode, Discovery

### 2.1 NetBox Assurance (drift detection) — announced 2025-04-02 `[vendor-reported]`

Source: https://Www.Globenewswire.Com/news-release/2025/04/02/3054215/0/en/NetBox-Labs-Announces-Availability-of-NetBox-Assurance.html

- NetBox Assurance is a managed drift-detection platform addressing "the gap between what you think your network looks like and what it actually looks like" (Richard Boucher, Senior Product Manager, NetBox Labs).
- Positioning: Day 1 (transformation/automation initiatives) and Day 2 (operational control — detect and remediate drift, reduce downtime risk, shrink attack surface).
- Built on the foundation of **Diode**, a source-available project from NetBox Labs offering a subset of Assurance functionality. Both share a common API for data ingestion through the **Diode SDK**, which provides an alternative interface for sending data to NetBox with built-in idempotence, automatic ordering, and other capabilities for high-performance integrations.
- Availability: optional add-on for both NetBox Cloud and NetBox Enterprise; tiered pricing based on volume of ingested entities per month (public price points not disclosed in the announcement — sales contact) `[vendor-reported][unverified]` on exact tiers.
- The **NetBox Discovery** observability agent is included with NetBox Assurance; agent extensions available in Standard and Premium bundles for Professional and Enterprise tiers.
- Portfolio context (from the same announcement): SaaS and self-managed enterprise-grade NetBox, airgapped installations, NetBox Discovery, and "a growing suite of AI features that enable AI-driven network and infrastructure management" `[vendor-reported]`.

### 2.2 NetBox Assurance in NetBox Cloud (managed) `[vendor-reported]`

Source: https://netboxlabs.com/blog/netbox-assurance-in-netbox-cloud-fully-managed-drift-detection-platform/

- Cloud deployment: zero platform infrastructure management (automatic scaling as discovery data grows, built-in resiliency, seamless upgrades); instant access for distributed teams without VPN; dev/staging environments (Professional tier and above); shared org-wide drift visibility; integrated user management through NetBox Cloud access controls; SOC 2 compliant infrastructure; encrypted data transmission for Discovery agent communications; automated security updates.
- Once activated, Assurance appears in NetBox Cloud navigation; Discovery agents and Diode integrations deployed on-prem connect to the managed Assurance platform for continuous drift detection.

### 2.3 Discovery & assurance partnerships (Jan 2024) `[vendor-reported]`

Source: https://www.globenewswire.com/news-release/2024/01/31/2820957/0/en/NetBox-Labs-Enters-Strategic-Partnerships-with-Leading-Network-Discovery-and-Assurance-Providers-to-Reduce-Adoption-Barriers-for-Network-Automation.html

- NetBox Labs partnered with **Forward Networks**, **IP Fabric**, and **Slurp'it**: bundled integrations with NetBox Cloud for network assurance and validation workflows.
- Claimed customer benefits: zero-cost Day-1 discovery via Slurp'it and IP Fabric to populate NetBox Cloud's data model; ongoing assurance/validation integrations (Forward Networks, IP Fabric, Slurp'it) with NetBox Cloud as the network source of truth; compliance-check insight by augmenting partner platform data with NetBox Cloud's model.

### 2.4 NetBox Labs company notes

- NetBox Labs is the commercial steward of NetBox; hosts the canonical plugins catalog (https://netboxlabs.com/netbox-plugins/) integrated into the product UI since v4.1 `[official]`.
- Funding/valuation figures: not captured in this research wave — flagged as an open item (see gaps log). Company described in press as supporting "the network and infrastructure community with a growing portfolio of innovative products that span the network operations, observability, security, and automation space" `[vendor-reported]`.

---

## Wave 3 — NetBox plugins ecosystem

### 3.1 Canonical catalog

- Canonical plugins catalog hosted by NetBox Labs: https://netboxlabs.com/plugins/ (also referenced as https://netboxlabs.com/netbox-plugins/); integrated natively into the NetBox UI since v4.1 so users can browse plugins and check for updates without leaving the product `[official]`.
- Public demo instance: https://demo.netbox.dev/ `[official]`; free NetBox Cloud tier advertised at https://netboxlabs.com/products/free-netbox-cloud/ `[vendor-reported]` (plan limits not captured — see gaps log).

### 3.2 Popular plugins (named in the NetBox README) `[official]`

- **NetBox Branching** (netboxlabs/netbox-branching) — work with isolated, mergeable branches of NetBox data (Git-like workflows for the source of truth).
- **NetBox Custom Objects** (netboxlabs/netbox-custom-objects) — define entirely new object types directly in the UI.
- **NetBox DNS** (sys4/netbox-plugin-dns) — manage DNS zones and records as an authoritative source of truth.
- **NetBox BGP** (netbox-community/netbox-bgp) — document and manage BGP sessions and routing policies.
- Community wiki of contributions: https://github.com/netbox-community/netbox/wiki/Community-Contributions `[official]`.

### 3.3 Plugin framework notes (v4.x)

- Plugins add models, views, and integrations on top of core; v4.7 plugin API additions: `GenericObjectChoiceField`/`GenericObjectFormMixin`, custom Jinja filters + template context injection, extension of core GraphQL types, custom Event Rule action types via `EventRuleAction` subclassing `[official]`.
- MPTT-backed `NestedGroupModel` deprecated in favor of `NestedLtreeGroupModel` — plugin authors with hierarchical models must migrate `[official]`.
- Core custom scripts are deprecated in favor of a **dedicated plugin**, to be removed in v5.0 — significant ecosystem signal: scripting moves out of core `[official]`.
- Community plugin development has modernized around `uv` for installs/CI and ruff for linting (observed across plugin changelogs, e.g. jsenecal/netbox-notices, jsenecal/netbox-sqids in 2026) `[secondary]`.

---

## Wave 4 — Nautobot core: version line and platform features (2023–2026)

### 4.1 Version timeline (latest first, as of 2026-09-22)

