---
id: etape6-trackd-firewalls/00-firewalls/9-8-opnsense-additional-confirmations
title: "9.8 OPNsense additional confirmations"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: reference
actors: ["AWS", "United States"]
dates: ["2026-06-30", "2026-07-15", "2026-08-13", "2026-08-20", "2026-09-18", "2026-09-22"]
keywords: ["advisory", "aws", "cybersecurity", "licenses", "mcp", "pricing", "research", "throughput", "valuation"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [264, 295]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: fe6a47fa3b9041aec0890a7bb65dd8b4e823602088f94411e26865c44e83e83b
---

# 9.8 OPNsense additional confirmations

- **pfSense Plus 26.07 (2026-08-13)** and **CE 2.9.0 (2026-08-20)** dates confirmed via GitHub MCP-server compatibility PR **[independent — https://github.com/gensecaihq/pfsense-mcp-server/pull/89]**.
- CE 2.9.0 base OS **FreeBSD 16.0-CURRENT** (lab-verified by a third-party project; same pfREST v2.10 contract across CE 2.8.1/2.9.0 and Plus 26.07) **[independent]**.
- Forum friction (Sept 2026): users whose Plus licenses (e.g., on SG2100) no longer renew discussing downgrade to free CE 2.9.0; "lifetime" era of the 2022 limited-time Plus promotion confirmed over **[secondary — https://forum.netgate.com/topic/201119/netgate-releases-pfsense-plus-software-version-26.07]**.
- **Netgate 1100**: $269 (TAC Lite included); TAC Professional 1-yr $49 / 2-yr $98 add-ons **[secondary — https://shop.netgate.com/products/1100-pfsense]**.
- **Netgate TAC on AWS**: Pro **$399/yr** (24h email SLA), Enterprise **$799/yr** (4h email/phone SLA), claimed 98% satisfaction **[vendor-reported — https://aws.amazon.com/marketplace/pp/prodview-lstvktrnia26i]**.
- Practitioner comparison (Wundertech, ~Aug 2026): OPNsense on fixed Jan/July cadence with bi-weekly security updates vs pfSense's irregular, Netgate-hardware-prioritized cadence; pfSense still recommended for complete beginners due to docs/community; OPNsense preferred interface; Zenarmor gives OPNsense a layer-7 story pfSense lacks **[independent — https://www.wundertech.net/pfsense-vs-opnsense/]**.

### 9.8 OPNsense additional confirmations

- **26.7.4 released September 15–16, 2026** — latest supported release at research date; 26.1 EOL July 15, 2026 **[official — https://forum.opnsense.org/Administrative/Announcements; https://en.wikipedia.org/wiki/OPNsense]**.
- Point releases: 26.7.1 (Jul 23), 26.7.2 (Aug 13), 26.7.3 (Sep 3), 26.7.4 (Sep 15) **[official — forum announcements]**.
- **Business Edition 26.4.2 released September 18, 2026**; BE 26.4 (Jun 3), 26.4.1 (Jul 23) **[official — forum announcements]**.
- **Nordian partners with Deciso** to scale European cybersecurity innovation — announced June 30, 2026 **[official — forum announcement]**.
- Ongoing 26.7-series issues: Tailscale-behind-NAT control-plane connectivity loss after 26.7.1_1 (GitHub issue opnsense/plugins#5613, Aug 2026, open); ports/tooling upgrade hiccups (Aug 2026) **[secondary]**.

### 9.9 Market figures (additional)

- NGFW software-solutions market: **$10.22B in 2026** → $11.16B (2027) per one vendor-estimate; North America $2.46B (2026), US $1.71B; Europe $1.75B (2026) **[secondary — https://www.fortunebusinessinsights.com/next-generation-firewall-market-112358; https://datainsightsmarket.com referenced estimates]**.
- Gartner **2026 Magic Quadrant for Hybrid Mesh Firewall** (Sept 2026): Fortinet, Palo Alto Networks, Check Point named Leaders; Fortinet highest on Ability to Execute, second consecutive year **[secondary — https://www.ad-hoc-news.de/boerse/corporate-news/fortinet-inc-stock-trades-near-record-high-as-analysts-back-innovation/70109912]**.
- 2026 buyer guidance (GBHackers): FortiGate safest broad price-performance pick; Palo Alto Networks where security depth outranks budget; Check Point where tested prevention accuracy matters; Zscaler for cloud-first; treat any NGFW's patching SLA as part of the purchase decision given edge devices are primary initial-access targets **[independent — https://gbhackers.com/ngfw-solutions-compared-features-pricing/]**.

### 9.10 Open verification items (supplementary)

1. 3500G/400G throughput specs (~595/164 Gbps) are partner-published, not from a Fortinet datasheet.
2. Plus 26.07 / CE 2.9.0 release dates derived from GitHub PR + forum, not the Netgate blog body.
3. FortiSOC real-world adoption/price figures not located.
4. CVE-2026-35616 FortiBleed-to-ransomware linkage rests on PoC-archive and press, not a Fortinet advisory on the campaign.
5. Stock/valuation figures are snapshots dated Jul–Sep 2026 and are time-sensitive.

- **Collected:** September 22, 2026.
- **Sources:** Fortinet official docs/earnings (docs.fortinet.com, GlobeNewswire), Netgate forum + docs (forum.netgate.com, docs.netgate.com), OPNsense official docs (docs.opnsense.org) + Deciso PR, BleepingComputer-via-TechRadar, Dark Reading, InfoRiskToday, EOTISEC analytical report, market research aggregators, CSO Online, SecurityBrief, Linuxiac, cybersecuritynews.com.
- **Conventions:** all monetary figures in USD unless noted; prices are Sep 2026 snapshots; every version number carries a release date where verified.
