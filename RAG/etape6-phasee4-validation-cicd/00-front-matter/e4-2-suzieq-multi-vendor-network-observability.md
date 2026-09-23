---
id: etape6-phasee4-validation-cicd/00-front-matter/e4-2-suzieq-multi-vendor-network-observability
title: "E4.2 — SuzieQ: multi-vendor network observability"
domain: front-matter
role: reference
task: reference
actors: ["AWS"]
dates: ["2023-06", "2025-02", "2026-01", "2026-09"]
keywords: ["aws", "benchmark", "funding", "pricing", "research", "series a", "series b"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [63, 132]
section: "Phase E4 — Network Validation, Observability & CI/CD"
sha256: 0f96945a363dc1c061f324e77a3be77bebe8eeaa17193aaa0fb83dd004aa8af6
---

# E4.2 — SuzieQ: multi-vendor network observability

## E4.2 — SuzieQ: multi-vendor network observability

### E4.2.1 Project identity and company

- SuzieQ is a network observability platform by Stardust Systems (CEO Dinesh Dutt); ships as a free open-source version plus an Enterprise edition deployed in production by multiple customers `[official]` — https://github.com/gelbox/suzieq.
- Gartner named Stardust Systems a "Cool Vendor" for making network automation easy for enterprises `[vendor-reported]` (SuzieQ README).
- Repository `gelbox/suzieq` last updated ~160 days before research date (GitHub crawl metadata); active but not daily-committed `[secondary]` (GitHub).

### E4.2.2 Architecture: poll → normalize → store → analyze

- SuzieQ collects data from devices/systems across the network (pollers), normalizes it into a vendor-independent schema, stores it, and exposes analysis on top `[official]` (README) — https://suzieq.readthedocs.io/en/latest/poller/.
- Design goal: think of the network holistically; demonstrate systematic network analysis rather than box-by-box CLI `[official]` (README).
- Audience: network engineers and designers (not only operators) `[official]` (README).

### E4.2.3 Features and use cases

- Documented use cases: network documentation, troubleshooting, fabric-wide visibility, network refresh/redesign, low/no-code validation, audits and compliance, proactive health checks `[secondary]` — https://podm8.com/episodes/heavy-networking-from-packet-pushers/hn719-meet-suzieq-the-network-observability-application-RgRWOrajV2 (Packet Pushers Heavy Networking HN719, guest Dinesh Dutt).
- Polling frequency: SuzieQ can pull device data as frequently as every 90 seconds, enabling near-immediate drift detection between observed and intended state `[vendor-reported]` — https://netboxlabs.com/blog/netbox-labs-stardust-systems-network-observability-partnership/.
- Multi-vendor by design; lightweight and fast `[secondary]` (HN719).

### E4.2.4 NetBox integration (2026)

- NetBox Labs × Stardust Systems partnership: "NetBox Sync" in SuzieQ Enterprise 2.8 populates NetBox with discovered data — devices, hostnames, device types, serial numbers, platforms, platform versions, interfaces, MTUs, IPs, prefixes, VRFs `[official]` — https://netboxlabs.com/blog/netbox-labs-stardust-systems-network-observability-partnership/.
- The integration is bidirectional: NetBox intended state enriches SuzieQ observed-state alerting, cutting troubleshooting time; combined use in pre- and post-change testing `[official]` (same source).
- Part of NetBox Labs' "big-tent" discovery ecosystem alongside Forward Networks, IP Fabric, and Slurp'it (announced January 2026) `[official]` (same source).

### E4.2.5 Gaps

- Current open-source release version number not captured in research `[unverified]` — check https://github.com/gelbox/suzieq/releases.
- Enterprise edition pricing is not public `[unverified]`.
- Supported NOS/platform list changes over time; verify against the poller docs before committing to coverage claims `[unverified]` — https://suzieq.readthedocs.io/en/latest/poller/.

---

## E4.3 — IP Fabric: automated network assurance platform

### E4.3.1 Company and funding

- IP Fabric founded 2015 in Prague by Pavel Bykov (CEO/co-founder), Roman Aprias (CTO/co-founder), Miroslav Hybl (engineering co-founder) `[secondary]` — https://yespress.io/ip-fabric.
- Funding: ~$30.7M total raised; Series B $25M led by One Peak (June 2023) with Senovo and Presto Ventures; earlier seed/Series A backed by Senovo, Presto Ventures, Credo Ventures (amounts undisclosed) `[secondary]` (yespress.io).
- Customer base includes Fortune 50 companies; named reference customers include Red Hat, Air France, Major League Baseball, Airbus, Avast, HCL, A1 Telekom Austria `[vendor-reported]` (yespress.io).
- Positioning: "automated network assurance platform" / digital twin; built by network people frustrated with existing tooling limits `[vendor-reported]` (yespress.io).

### E4.3.2 Version 7.9 (January 2026) — unified cloud visibility

- IP Fabric 7.9 released January 2026: unifies cloud, on-premises, and hybrid network assurance `[vendor-reported]` — https://www.dbta.com/Editorial/News-Flashes/IP-Fabric-Boosts-Unified-Cloud-Visibility-to-Hasten-Digital-Transformation-173106.aspx.
- New: expanded Azure and GCP visibility (Azure Firewall, Private Link, Private Endpoints, multi-project GCP discovery), hybrid pathing with enriched cloud metadata and GCP Interconnect support, improved IPv6 support across major vendors, workflow improvements `[vendor-reported]` (dbta.com).
- Cloud-native constructs modeled as first-class objects in the digital twin (Azure Firewall, GCP topologies), enabling traffic-flow tracking through cloud security boundaries and private connectivity paths `[secondary]` — https://dedirock.com/blog/enhancing-hybrid-environment-visibility-key-features-of-ip-fabric-7-9/.
- IPv6 path analysis across dual-stack environments; API scalability enhancements for autonomous operations `[secondary]` (dedirock.com).
- Compliance angle: worked with financial institutions so cloud firewall policies mirror on-premises (PCI-DSS context) `[secondary]` — https://itmagazine.com/2026/01/17/enhancing-hybrid-environment-visibility-whats-new-in-ip-fabric-7-9/.

### E4.3.3 Version history trail (2025)

- Version 7.2 (2025): security posture assurance — firewall discovery and simulation; ServiceNow extension for CMDB accuracy and dependency mapping `[secondary]` (yespress.io).
- Version 7.0 (February 2025): 160+ automated intent-verification checks; multi-view dashboards; shareable snapshots (functional network simulations/"digital twins"); exportable network diagrams (Visio); expanded cloud discovery (AWS Direct Connect Transit VIF, Transit Gateway insights); SD-WAN visibility (Silver Peak, Viptela); auto-discovery of security tech (Check Point, Palo Alto, Stormshield); advanced routing data (exact BGP routes advertised to neighbors) `[vendor-reported]` — https://Www.Globenewswire.com/news-release/2025/02/10/3023193/0/en/IP-Fabric-7-0-Transforms-Cloud-and-Edge-Innovation-Across-Hybrid-Networks.html.
- 2020: Python SDK and integrations (ServiceNow, Ansible, NetBox) `[secondary]` (yespress.io).

### E4.3.4 Capabilities summary

- Auto-discovery of multi-vendor networks; intent verification with predefined checks flagging misconfigurations/inconsistencies at scale; end-to-end path analysis (incl. hybrid cloud); network snapshots as shareable simulations; security/compliance validation `[vendor-reported]` (globenewswire 7.0 release).
- Integrations: ServiceNow (CMDB/dependency mapping), Ansible, NetBox `[secondary]` (yespress.io).

### E4.3.5 Pricing signals and gaps

- Pricing is not published on the sources found; enterprise sales motion implied by Fortune-50 customer base `[unverified]`.
- Independent head-to-head benchmark vs Forward Networks or Batfish not found in research `[unverified]`.
- Exact 7.9.x patch level as of September 2026 not captured `[unverified]`.

---

