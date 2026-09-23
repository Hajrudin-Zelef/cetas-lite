---
id: etape6-trackd-firewalls/00-firewalls/overview
title: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: reference
actors: ["Apple", "CISA", "Intel", "Nvidia"]
dates: ["2025-12", "2026-01", "2026-01-15", "2026-01-22", "2026-02-01", "2026-03-10", "2026-05-06", "2026-06", "2026-07-28", "2026-07-29", "2026-09", "2026-09-17", "2026-09-22"]
keywords: ["advisory", "agent", "agents", "asic", "benchmarks", "chiplet", "cybersecurity", "disaggregated", "exploit", "foundry", "governance", "intel"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [1, 85]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: cf7efa24d7901e960017ea23290297ac29a2aa73fe066ec7fff1c9e147b53779
---

# Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)

**Project:** RAG data-collection, Step 6 (Réseau & sécurité)
**Coverage window:** February 1, 2026 → September 22, 2026
**Status:** Research snapshot. Prices and version figures are dated snapshots as of September 22, 2026.

### Provenance legend
- **[official]** — vendor's own blog, docs, release notes, earnings statement, or advisory.
- **[vendor-reported]** — figure claimed by the vendor (benchmarks, market share) without independent audit.
- **[independent]** — reputable third-party press (BleepingComputer, The Register, Dark Reading, Reuters, IDC, Gartner) or independent measurement.
- **[secondary]** — lower-tier press, blogs, analyst summaries, community trackers; useful but unverified.
- **[unverified]** — single-source or conflicting claims; treat as uncertain.

---

## 1. FORTINET / FortiGate

### 1.1 FortiOS 8.0 — the headline OS release (announced March 10, 2026)

- Fortinet launched **FortiOS 8.0** at its **Accelerate 2026 conference in Las Vegas** on **March 10, 2026** — the new OS release powering the Fortinet Security Fabric **[secondary — https://www.stocktitan.net/news/FTNT/fortinet-introduces-forti-os-8-0-to-expand-secure-networking-with-iv0y66jrnxym.html; https://www.coffeefranchisehub.com/archives/42978]**.
- FortiOS 8.0 groups its updates into three pillars: **AI-driven security, next-generation SASE, and quantum-safe protection** **[secondary — https://securitybrief.co.uk/story/fortinet-s-fortios-8-0-adds-ai-quantum-safe-tools]**.
- **AI usage controls / governance** (new in 8.0):
  - **FortiView for AI**: attack-surface and shadow-AI visibility — shows how AI applications/services are used across the organization, distinguishing sanctioned vs unsanctioned tools **[secondary]**.
  - **AI-aware application control** for generative-AI usage; visibility into **Model Context Protocol (MCP) and agent-to-agent (A2A) interactions** to reduce blind spots **[secondary]**.
  - Enhanced data loss prevention (DLP) powered by **optical character recognition (OCR)** **[secondary]**.
  - AI agents across the Fortinet Security Fabric enabling conversational troubleshooting and configuration in firewall and SD-WAN environments **[secondary]**.
- **SASE in 8.0**: **FortiSASE Outpost** and sovereign deployment options for data residency; unified SD-WAN bundles; multipath IPsec tunnels **[secondary]**.
- **Quantum-safe**: quantum-resilient cryptographic controls, hybrid post-quantum cryptography (PQC) SSL deep inspection **[secondary]**.
- **ZTNA in 8.0** (from the official FortiOS 8.0 release notes): a FortiGate can now act as a **ZTNA service connector**, reverse-proxying service connections upstream to the ZTNA Edge (e.g., FortiPAM or FortiProxy acting as ZTNA access proxy); configuration simplification/modularization (`firewall.access-proxy` split into `ztna.traffic-forward-proxy`, `ztna.web-proxy`, `ztna.web-portal`; new `ztna.destination` object); IPv6 security posture tags/groups for dual-stack ZTNA policies; tag sharing with FortiClient EMS **[official — https://docs.fortinet.com/document/fortigate/8.0.0/fortios-release-notes/743723/new-features-or-enhancements]**.
- Ken Xie quote (CEO): *"FortiOS 8.0 reflects more than 25 years of continued innovation at the intersection of networking and security. As organizations embrace AI, cloud, and increasingly encrypted environments, a unified operating system is essential to reduce complexity, improve visibility, and ensure security can scale without slowing the business."* **[secondary]**

### 1.2 New FortiGate hardware in 2026 (G series)

**FortiGate 3500G + FortiGate 400G — announced May 6, 2026**
- New additions to the **FortiGate G series**, positioned to secure AI workloads and encrypted traffic in data centers and enterprise edges **[official/via GlobeNewswire — https://www.globenewswire.com/news-release/2026/05/06/3288892/0/en/Fortinet-Expands-FortiGate-G-Series-to-Secure-AI-from-the-Data-Center-to-Modern-Enterprise-Edges.html]**.
- Powered by Fortinet's **NP7 and SP5** custom ASICs and FortiOS **[official]**.

**FortiGate 1200G — announced July 28, 2026**
- Midrange G-series appliance: **10G, 25G and 100G connectivity**; **397 Gbps firewall throughput**; aimed at campus, data center and hybrid environments inspecting large volumes of encrypted traffic **[secondary — https://www.csoonline.com/article/4202566/fortinets-new-fortigate-platform-converges-firewall-sase-technologies-2.html; https://www.itp.net/cybersecurity/fortinet-launches-fortigate-1200g-to-advance-firewall-and-sase-convergence]**.
- Fills the gap between campus/edge/branch appliances (e.g., 400G) and flagship data-center platforms (e.g., 3500G) **[secondary]**.
- Ships with **FortiOS 8.0** and can operate as a **FortiSASE Outpost**: the 1200G becomes a local SASE point of presence (POP) inside customer-controlled environments (on-prem sites, private DCs, colocation), enforcing security locally while staying centrally managed from the cloud — targeting data-sovereignty and latency-sensitive workloads **[secondary]**.
- **Availability expected in Q3 2026** **[secondary]**.
- Integrates **FortiGuard AI-Powered Security Services** and **FortiAI** for real-time threat intelligence and AI-assisted SecOps **[secondary]**.

### 1.3 Intel SP6 collaboration (July 29, 2026)

- Fortinet announced a **strategic collaboration with Intel to develop Fortinet Security Processor 6 (SP6)**, combining Fortinet's proprietary security-processor expertise with Intel's design, development, packaging and manufacturing capabilities, strengthening supply-chain resilience and diversity **[official — via Q2 2026 earnings, https://www.globenewswire.com/news-release/2026/07/29/3335555/0/en/fortinet-reports-strong-second-quarter-2026-financial-results.html]**.
- Secondary reporting (TradingView/UDIS): Intel Foundry will co-develop and manufacture the SP6 ASIC; Fortinet is reportedly **the first external cybersecurity customer for Intel Foundry and the first external customer on the Intel 4 node**, with production at **Fab 34 in Ireland**; the SP6 moves from the monolithic 7nm design of the SP5 to a **disaggregated chiplet architecture** for better yields and lower unit costs **[secondary — https://www.tradingview.com/chart/FTNT/vj3CqrYb-Why-Is-Intel-Building-Fortinet-s-Next-Firewall-Chip/; unverified in part — Intel Foundry customer specifics not independently confirmed in this research]**.

### 1.4 Financials — Q2 2026 results (July 29, 2026)

- Revenue **$2.05B (+26% YoY)**; product revenue **$773M (+52% YoY)**; billings **$2.37B (+33% YoY)**; GAAP operating margin 34%; non-GAAP operating margin **38%**; GAAP EPS $0.82 (+44%); non-GAAP EPS $0.90 (+41%); operating cash flow $1.04B; free cash flow $966M **[official — https://www.globenewswire.com/news-release/2026/07/29/3335555/0/en/fortinet-reports-strong-second-quarter-2026-financial-results.html]**.
- FY2026 revenue guidance raised to **19% YoY growth**; Q3 2026 EPS guidance $0.83–0.87 (non-GAAP); FY2026 EPS guidance $3.41–3.47 **[official/secondary — https://www.americanbankingnews.com/2026/09/17/newedge-advisors-llc-increases-holdings-in-fortinet-inc-ftnt.html]**.
- FY2025 context: total revenue **$6.80B**, net income $1.85B, free cash flow $2.21B **[secondary — StockTitan 10-K reference]**.
- CEO Ken Xie framing: customers value Fortinet's **"SASE Firewall"** — firewall, SD-WAN and SASE integrated on a single FortiOS, powered by FortiASIC, with sovereign deployment options **[official]**.

### 1.5 Market position 2026

- Fortinet is the leader in firewall **unit shipments**: **over 50% share by unit shipments** (Beaver Research, Jan 2026, citing IDC) — one secondary source cites ~48% **[secondary — https://medium.com/@beaverresearch/discounted-firewall-king-fortinets-high-quality-growth-at-a-bargain-0bf76c98d817; https://www.tradingview.com/chart/FTNT/vj3CqrYb-Why-Is-Intel-Building-Fortinet-s-Next-Firewall-Chip/]**.
- **750,000+ customers** worldwide; ~80% of the Fortune 100 and 72% of the Global 2000 use Fortinet products **[secondary, vendor-cited]**.
- Named a **Leader in the 2026 Gartner Magic Quadrant for Hybrid Mesh Firewall** **[secondary — via MarketBeat headline roundup, https://www.marketbeat.com/instant-alerts/filing-fortinet-inc-ftnt-shares-bought-by-newedge-advisors-llc-2026-09-17/]**.
- Secure Networking addressable market estimated by Fortinet at ~$86B by 2027 (~9% annual growth); Universal SASE ~$36B by 2027 (~20% growth) — Fortinet's own market sizing, [vendor-reported] **[secondary — https://www.stocktitan.net/news/FTNT/fortinet-sharpens-business-focus-on-core-growth-areas-to-extend-5ukuetb3r4or.html]**.

### 1.6 Vulnerabilities and advisories 2026

**CVE-2026-24858 — FortiCloud SSO improper cryptographic signature verification (zero-day, actively exploited)**
- Disclosed **late January 2026** (Fortinet CISO blog post; patched FortiOS **7.4.11**, with 7.6.6 and 8.0.0 follow-ups) **[independent — https://www.cybersecurity-help.cz/blog/5202.html]**.
- Exploitation observed **starting January 15, 2026** (Arctic Wolf); attackers created new local admin accounts, granted VPN access, and exfiltrated firewall configurations even on fully patched devices; two malicious FortiCloud accounts disabled on **January 22, 2026** **[independent — https://www.inforisktoday.com/fortinet-locks-down-forticloud-sso-amid-zero-day-attacks-a-30612]**.
- Root cause: an attacker with a FortiCloud account and a registered device could log into **other devices registered under different accounts** when FortiCloud SSO was enabled. Fortinet clarified it impacts **FortiCloud SSO only** (not third-party SAML IdP or FortiAuthenticator) **[independent — https://Www.Darkreading.com/vulnerabilities-threats/fortinet-new-zero-day-malicious-sso-logins]**.
- Initially suspected to be a bypass of the December 2025 patch for **CVE-2025-59718** (FortiCloud SSO authentication bypass via crafted SAML messages; Arctic Wolf reported active exploitation in December 2025; added to CISA KEV), plus companion **CVE-2025-59719**. Fortinet confirmed CVE-2026-24858 is a **separate, new vulnerability**, not a patch bypass **[independent — https://www.inforisktoday.com/fortinet-locks-down-forticloud-sso-amid-zero-day-attacks-a-30612; https://www.techradar.com/pro/security/fortinet-fortigate-devices-hit-in-automated-attacks-which-create-rogue-accounts-and-steal-firewall-data]**.
- CVE-2026-24858 was added to the **CISA Known Exploited Vulnerabilities (KEV)** catalog **[independent — Dark Reading]**.
- Workaround while unpatched: disable FortiCloud SSO admin login (`set admin-forticloud-sso-login disable`) **[independent — TechRadar via BleepingComputer]**.

**"FortiBleed" credential-compromise reporting (unverified as a zero-day)**
- An EOTISEC analytical report (EOTISEC-2026-044, June 2026) drew on findings by researcher Volodymyr Diachenko, Kevin Beaumont, Hudson Rock and SOCRadar describing a large validated credential set tied to Fortinet devices. The report explicitly states **no Fortinet zero-day was confirmed as of its issue date** **[secondary — https://sveoti.net/wp-content/uploads/2026/06/EOTISEC-2026-044_FortiBleed.pdf]**. Treat the credential set as [unverified]; treat "no confirmed zero-day" as the report's own assessment **[unverified]**.
- Separately, a September 2026 underground forum post claimed a **private FortiGate RCE 1-day for sale**; analysis (UNDERCODE NEWS) concluded it is an exploit-sale claim with **no CVE, root cause, affected build range, or independent reproduction** — not evidence of a new zero-day **[secondary — https://undercodenews.com/fortigate-under-pressure-underground-1-day-rce-sale-raises-fresh-alarm-for-enterprise-vpn-security-video/; unverified]**.

### 1.7 FortiAI / AI security features (context)

- FortiGuard AI-Powered Security Services and **FortiAI** integrated into the FortiGate 1200G for real-time threat intelligence, automated protection and AI-assisted SecOps **[secondary — ITP.net]**.
- Prior 2026-adjacent AI launches cited by StockTitan Argus: **Secure AI Data Center blueprint with Arista** deployed at MPS; **FortiGate VM integrated on NVIDIA BlueField-3 DPUs** for AI workloads; FortiGate 3800G for Secure AI Data Center; AI-powered workspace/email security suite (Jun 4, 2026) **[secondary]**.
- FortiOS 8.0 AI agents enable conversational troubleshooting and configuration across firewall and SD-WAN **[secondary]**.

---

