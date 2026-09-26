---
id: etape6-trackd-firewalls/00-firewalls/8-collection-metadata
title: "8. Collection metadata"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Apple", "Glasswing", "Microsoft", "OpenAI"]
dates: ["2026-04", "2026-05-06", "2026-06-03", "2026-06-16", "2026-07", "2026-09-22"]
keywords: ["advisory", "agent", "agentic", "agents", "cyber", "cybersecurity", "exploit", "mcp", "model context protocol", "research", "revenue", "throughput"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [210, 263]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: 9d3199f2da62da87756802b9e82b431a568ad9d5c3525f1274b66c7b155d1292
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

