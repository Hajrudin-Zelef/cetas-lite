---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/7-4-device42-commercial-dcim-cmdb-vendor-reported-secondary
title: "7.4 Device42 (commercial DCIM/CMDB) `[vendor-reported][secondary]`"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: ["California", "CoreWeave", "Nvidia"]
dates: ["2025-07-14", "2026-05-24"]
keywords: ["agent", "agentic", "apache", "datacenter", "ethernet", "funding", "nvidia", "pricing", "research", "revenue", "series b"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [277, 320]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 1ab9578c06525e67c4abc3728295dd9ae47e7a4e25a852adca16c203c977837e
---

# 7.4 Device42 (commercial DCIM/CMDB) `[vendor-reported][secondary]`

- **RackTables** (racktables.org): long-standing open-source datacenter asset management (rack space, IP addresses, network connections); GPL. Still referenced in 2026 but with visibly slower release cadence than NetBox/Nautobot — exact latest release not re-verified in this wave (see gaps log).
- **openDCIM**: open-source DCIM focused on power/space/asset tracking; community-maintained. Currency of 2026 releases not verified in this wave (see gaps log).

### 7.4 Device42 (commercial DCIM/CMDB) `[vendor-reported][secondary]`

- Device42 is a commercial agentless discovery + DCIM/CMDB platform (acquired by **Freshworks** in 2024 — widely reported; verify before citing) `[secondary]`. It competes in the DCIM/discovery space where NetBox Cloud + Discovery/Assurance now play; typical enterprise DCIM evaluation set: Device42, Sunbird dcTrack, Nlyte, Modius — NetBox vs commercial DCIM comparisons usually hinge on source-of-truth-first vs discovery-first architecture `[unverified]` on 2026 pricing.

### 7.5 NetBox vs commercial DCIM framing `[secondary]`

- NetBox Labs' product motion (NetBox Cloud + Discovery agent + Assurance drift detection, airgapped Enterprise) directly targets the DCIM/discovery overlap: source of truth + observed state reconciliation, which is the traditional DCIM value proposition `[vendor-reported]` on portfolio direction.
- NetBox's differentiators in DCIM evaluations: Apache 2.0 core, API-first, plugin ecosystem, config-context-driven automation, cooling modeling (v4.7, 2026); gaps classically cited: no native discovery (addressed via Discovery/Assurance add-ons), no built-in monitoring/alerting, rack elevation is visualization-only `[secondary]`.

---

## Wave 8 — Real deployments, adoption signals, and company metrics

### 8.1 NetBox Labs — company and adoption metrics `[vendor-reported][secondary]`

- **$35M Series B (2024):** NetBox Labs raised a $35M Series B in 2024 with a 12-slide deck positioning the company as the "commercial steward" of NetBox, the "central nervous system" for modern AI-driven infrastructure. Deck claims: **18,000+ GitHub stars**, **300+ contributors**, presence in "thousands of production environments" `[vendor-reported]`. Source: https://startupfundraising.com/library/articles/netbox-labs-pitch-deck-teardown
- SiliconANGLE (2025-07-14): NetBox Cloud used by **"thousands" of enterprises** as network source of truth, including "several dozen Fortune 500 organizations", data-center hyperscalers, AI scale-ups, and government agencies globally `[vendor-reported]`. Source: https://siliconangle.com/2025/07/14/netbox-labs-central-nervous-system-ai-data-centers-gets-35m-funding/
- **CoreWeave** (publicly traded AI data-center operator) — named customer. Jim Julson, head of networks: "We're building dozens of new AI data centers every year, and they're full of complex infrastructure… NetBox is crucial for accelerating our timelines with automation. Deploying our infrastructure even a month sooner as the result of these efficiencies directly impacts our revenue, and NetBox enables streamlined operations and automation once infrastructure is in production." `[vendor-reported]`
- Product portfolio per the same report: NetBox Discovery (network/device discovery), NetBox Assurance (drift detection/remediation), **NetBox Operator** — described as "an agentic AI operations tool that accelerates automation for networking teams" `[vendor-reported]`.

### 8.2 Named NetBox deployments

- **Hagerty Insurance** (case study, netboxlabs.com) `[vendor-reported]`: NetBox introduced Fall 2022 as a rack-space planning PoC; evolved into comprehensive documentation/planning for two primary data centers in an active-active metro cluster (Metro Ethernet connectivity), predominantly Cisco networking; remote-first team uses it to track physical configurations and plan deployments "down to the patch cables, patch panel ports, and even power cables". Source: https://netboxlabs.com/customer-stories/hagerty/
- **CENIC** (California research & education network) `[official]`: evaluated multiple DCIM platforms, selected NetBox as NSoT; modeled ~500 associate circuits migrating to Juniper MX10008 routers across six of seven backbone node sites (400 Gbps handoffs); deployed via **NetBox Cloud** (SaaS) with NetBox Labs managing hosting/backups/upgrades. Source: https://cenic.org/blog/cenic-network-automation-update-new-network-source-of-truth-tool-implemented
- **Insight Global** (NetBox Heroes podcast, Michael Kutka) `[vendor-reported]`: NetBox as IPAM and source of truth for equipment; revived during MPLS→SD-WAN transition to track circuit providers; APIs used to push/pull configuration per office. Source: https://netboxlabs.com/blog/netbox-heroes-michael-kutka/
- **SFMIX** (San Francisco internet exchange, public docs) `[independent]`: NetBox as source of truth for intended IXP network configurations; uses tagging + custom fields for cross-administrative-domain objects (peering LANs, participants, ports); VM deployment via Ansible; Device-Type-Library-Import for device types. Source: https://github.com/sfmix/sfmix/blob/HEAD/documentation/netbox.md
- **Inngest** (engineering blog) `[independent]`: NetBox for hardware inventory management in their data-center move off cloud; ETL funnel (Ansible facts, interfaces, placement) into NetBox as source of truth; NetBox data feeds libvirt-based VM placement/scheduling. Source: https://github.com/inngest/website/blob/HEAD/content/blog/event-driven-infrastructure-inventory-inngest-netbox.mdx
- **Homelab-scale (2026):** community member running **NetBox 4.6.1** on K3s (deployed 2026-05-24) as source of truth for devices/VMs/LXCs/IPs/VLANs/prefixes/sites/clusters, written by Terraform via the `e-breuninger/netbox` provider — evidence of the GitOps+NetBox pattern in 2026 `[secondary]`. Source: https://github.com/xiiisins/homelab/blob/HEAD/docs/services/netbox.md

### 8.3 Nautobot — company and adoption metrics `[vendor-reported]`

- Network to Code 2025 performance (announced early 2026): **52% recurring revenue growth in 2025**, second consecutive year of double-digit organic growth across SaaS and services; **Guidepost Growth Equity expanded its investment**; new CEO leadership (Paul Brady) fueling product-led growth and commercialization. Investment supports commercial software built on Nautobot plus AI-powered capabilities. Source: https://www.lelezard.com/en/news-22091711.html
- Scale claims: **50,000+ Nautobot installs globally**; supports "some of the world's largest and most complex enterprise networks, **such as NVIDIA**", delivering automation across global production environments `[vendor-reported]`. Advanced AI capabilities delivered through **NautobotGPT** `[vendor-reported]`.

### 8.4 Named Nautobot deployments

- **SCinet 2025** (Supercomputing Conference, Nov 16–21, 2025, St. Louis) `[vendor-reported]`: Nautobot selected as the source of truth and automation platform for the high-capacity temporary network connecting supercomputers and research networks; NTC supplying the platform plus hands-on engineering; automated repetitive design tasks — cabling and patching, hardware deployment, IP schema configuration; first time SCinet used a commercial network automation platform end-to-end (design → planning → deployment → teardown). Source: https://www.morningstar.com/news/accesswire/1104015msn/network-to-code-powers-scinet-2025-the-high-performance-network-behind-the-supercomputing-conference
- **Kubernetes deployment pattern** (NTC blog, 2025) `[official]`: documented Nautobot-on-Kubernetes workflow (Helm/flux, custom Docker image with Golden Config + Nornir plugin via requirements.txt, `nautobot_config.py` mounted) — representative of production deployment practice. Source: https://networktocode.com/blog/deploying-nautobot-to-kubernetes-03/

---

