---
id: etape6-trackd-firewalls/00-firewalls/8-collection-metadata
title: "8. Collection metadata"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Apple", "Glasswing", "Microsoft", "OpenAI", "United States"]
dates: ["2026-04", "2026-05-06", "2026-06-03", "2026-06-16", "2026-06-30", "2026-07", "2026-07-15", "2026-08-13", "2026-08-20", "2026-09-18", "2026-09-22"]
keywords: ["advisory", "agent", "agentic", "agents", "aws", "cyber", "cybersecurity", "exploit", "licenses", "mcp", "model context protocol", "pricing"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [210, 295]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: 3deb0c06536aa18737171fd375abb634ca58533de93a351b261adc12de2a62ce
---

# 8. Collection metadata

## 8. Collection metadata

- **Collected:** September 22, 2026.
- **Sources:** Fortinet official docs/earnings (docs.fortinet.com, GlobeNewswire), Netgate forum + docs (forum.netgate.com, docs.netgate.com), OPNsense official docs (docs.opnsense.org) + Deciso PR, BleepingComputer-via-TechRadar, Dark Reading, InfoRiskToday, EOTISEC analytical report, market research aggregators, CSO Online, SecurityBrief, Linuxiac, cybersecuritynews.com.
- **Conventions:** all monetary figures in USD unless noted; prices are Sep 2026 snapshots; every version number carries a release date where verified.

---

## 9. Supplementary research pass — additional findings (second track-D agent, 2026-09-22)

### 9.1 FortiSOC — agentic-AI SOC platform details (announced June 16, 2026)

- **FortiSOC** launched June 16, 2026 (Sunnyvale, Calif.), cloud-delivered (SaaS) platform unifying SIEM, SOAR, threat intelligence, UEBA, case management, and ITDR **[official — https://www.fortinet.com/corporate/about-us/newsroom/press-releases/2026/fortinet-launches-fortisoc-unified-soc-platform-powered-by-agentic-ai]**.
- Agentic layer branded **FortiAI-Assist**: autonomous alert investigation, cross-asset/identity correlation, AI-generated playbooks, **Model Context Protocol (MCP)**-powered agent coordination across alerts, investigations, threat hunting, cases and response **[official]**.
- Built-in detection content/playbooks from Fortinet's own global SOC, with **monthly content updates** **[official]**.
- Previewed at **Fortinet Accelerate 2026** as the successor architecture to FortiAnalyzer/FortiSIEM/FortiSOAR/FortiTIP, which continue to be sold for customers preferring them **[secondary — https://investingnews.com/fortinet-advances-its-security-operations-platform-with-unified-soc-agentic-ai-and-expanded-endpoint-security/]**.
- FortiGuard SOC-as-a-Service enhancements: third-party log sources, FortiNDR and FortiCNAPP telemetry **[secondary]**.
- Positions Fortinet directly against Palo Alto Networks, CrowdStrike, Microsoft Sentinel in XDR/SIEM **[independent — https://cyber.netsecops.io/articles/fortinet-launches-unified-fortisoc-platform-with-agentic-ai/]**.

### 9.2 FortiOS 8.0 release-notes confirmations (official docs)

- **FortiAI-Assist** is documented as Feature ID 1058390 in the FortiOS 8.0.0 release notes: RAG-enhanced docs support, automated diagnostic analysis, CLI Code Lab with natural-language config generation; uses FortiAI with monthly FortiCare Premium token allotments (top-up possible) or customer OpenAI keys **[official — https://fortinetweb.s3.amazonaws.com/docs.fortinet.com/v2/attachments/d94b0a41-2dee-11f1-b31d-02356ffb40d9/fortios-v8.0.0-release-notes.pdf]**.
- New `http_authd` daemon for administrative authentication; GUI-driven FortiAP packet sniffer and PCAP export; granular federated-upgrade failure reporting **[official]**.

### 9.3 FortiGate 1200G / 3500G / 400G spec confirmations

- **FortiGate 1200G**: 397 Gbps firewall throughput, 10G/25G/100G connectivity, availability expected Q3 2026; integrates FortiGuard AI-Powered Security Services + FortiAI **[secondary — https://www.csoonline.com/article/4202566/fortinets-new-fortigate-platform-converges-firewall-sase-technologies-2.html; https://www.itp.net/cybersecurity/fortinet-launches-fortigate-1200g-to-advance-firewall-and-sase-convergence]**.
- **FortiGate 3500G**: ~595 Gbps firewall throughput, ~179M concurrent sessions (data-center scale, 400Gb connectivity) **[secondary — partner-published; vendor-derived]**.
- **FortiGate 400G**: ~164 Gbps firewall throughput, ~28M concurrent sessions (enterprise edge) **[secondary — partner-published; vendor-derived]**.

### 9.4 CVE-2026-35616 — FortiClient EMS authentication bypass ("FortiBleed" exploitation thread)

- Actively exploited FortiClient EMS auth-bypass flaw, patched **April 2026**; Fortinet advisory linked in PoC documentation **[secondary — https://github.com/intelseclab/poc-archive/blob/HEAD/pocs/network/2026-07-03_cve-2026-35616-forticlient-ems-auth-bypass/README.md]**.
- Attack technique: enumerate client-cert CA names via `openssl s_client`, forge a self-signed certificate, replay forged `X-SSL-CLIENT-VERIFY` header against `/api/v1/fabric_device_auth/*` and `/api/v1/fortigate/*` endpoints — no real mTLS handshake required **[secondary]**.
- **Remediation: upgrade FortiClient EMS to 7.4.7 or later**; interim: restrict management-plane access to trusted admin networks **[secondary]**.
- Linked to the "FortiBleed" credential-theft campaign and **INC Ransom / Lynx ransomware** operations (BleepingComputer, July 2026); EKZ Stealer artifacts noted downstream **[secondary — https://www.bleepingcomputer.com/news/security/fortibleed-credential-theft-campaign-linked-to-lynx-ransomware/]**.

### 9.5 CVE-2024-21762 — Thai ISP (3BB) intrusion (observed June 3, 2026)

- Hunt.io discovered an exposed operator toolkit (92.63.180[.]133:8888; 298 files / 30 subdirs) targeting a **FortiGate 60F SSL-VPN** (mail.3bb.co[.]th:10443) **[independent — https://cyberpress.org/hackers-exploit-fortigate-ssl-vpn-flaw-to-breach-thai-isp-and-deploy-meshcentral-backdoor/; https://gbhackers.com/hackers-exploit-fortigate-ssl-vpn-flaw/]**.
- Attackers fingerprinted FortiOS builds, staged controlled service-crash tests to distinguish vulnerable appliances, then exploited **CVE-2024-21762** (SSL-VPN out-of-bounds write → unauthenticated RCE) via `/remote/hostcheck_validate` ROP payloads; reverse shell to 92.63.180[.]133:9443; attempted firmware download to tailor ROP gadgets to the target's model/firmware **[independent]**.
- Post-exploitation: **MeshCentral** deployed as C2 (group "TH-3BB"), agents running with root privileges **[independent]**.

### 9.6 Fortinet Q1 2026 context and financial framing

- Q1 2026 (reported May 6, 2026): beat estimates; billings **+31% YoY**; announced FortiGate 3500G/400G alongside; collaboration noted with **Anthropic (Project Glasswing)** and OpenAI Group on AI security (scope undisclosed) **[secondary — https://siliconangle.com/2026/05/06/fortinet-shares-surge-first-quarter-earnings-beat-raised-outlook/]**.
- Q1 guidance at the time raised FY2026 outlook to revenue **$7.71–7.87B**, billings $8.80–9.10B, adjusted EPS $3.10–3.16 **[secondary]**.
- Q2 2026 segment detail (per earnings coverage): Secure Networking billings **+34%** (~66% of billings); Unified SASE **+35%** (~24%); AI-driven SecOps **+25%** (~10%); OT **+55%+**; **FortiSASE billings +100%+ YoY**; total billings **$2.37B (+33.4%)**; revenue **$2.05B (+25.6–26%)**; product revenue $773M (+52%); service revenue $1.27B (+14%); deferred revenue $7.68B **[independent — https://www.zacks.com/stock/news/2964411/fortinet-q2-earnings-revenues-beat-estimates-increase-yy; https://infotechlead.com/security/fortinet-q2-2026-revenue-jumps-26-to-2-05-bn-as-ai-data-center-sase-and-ot-security-demand-accelerates-97371]**.
- Combined "SASE Firewall" (NGFW + SD-WAN + Unified SASE + AI/quantum) grew **34%** and represented ~90% of billings; Fortinet's TAM claims: $148B by 2029 (SASE Firewall), $69B Secure Networking, $79B Unified SASE, $167B AI-driven SecOps **[vendor-reported]**.
- Stock: up ~12% in the month after Q2 results (Aug 28, 2026); traded near all-time highs ~$153–$174 with 52-week range $73.55–$174.05; ~96% one-year gains; forward P/E ~40–60 **[secondary — https://www.zacks.com/stock/news/2981878/why-is-fortinet-ftnt-up-12-since-last-earnings-report; https://tickeron.com/blogs/fortinet-ftnt-and-palo-alto-networks-panw-cybersecurity-sector-comparison-16826/; https://www.ad-hoc-news.de/boerse/corporate-news/fortinet-inc-stock-trades-near-record-high-as-analysts-back-innovation/70109912]**.
- NYC **Innovation Hub** opened 2026 (training labs, demos, executive briefing, Fortinet Cloud PoP) **[secondary]**.

### 9.7 pfSense additional confirmations

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
