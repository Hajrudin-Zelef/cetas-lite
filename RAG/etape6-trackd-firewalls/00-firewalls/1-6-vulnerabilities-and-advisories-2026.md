---
id: etape6-trackd-firewalls/00-firewalls/1-6-vulnerabilities-and-advisories-2026
title: "1.6 Vulnerabilities and advisories 2026"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: reference
actors: ["Apple", "CISA", "Intel", "Nvidia"]
dates: ["2025-12", "2026-01", "2026-01-15", "2026-01-22", "2026-06", "2026-09", "2026-09-17"]
keywords: ["agents", "cybersecurity", "exploit", "intel", "nvidia", "research", "zero-day"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [59, 85]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: e37379ee612ca893c0361f340fa969d3fa4df2bb2a9bfdeb37a5dec85c8b18db
---

# 1.6 Vulnerabilities and advisories 2026

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

