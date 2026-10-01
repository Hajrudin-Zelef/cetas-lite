---
id: collect-261001-general-networking/general-networking/etape6-trackd-firewalls-2
title: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "CISA", "Intel", "Nvidia"]
dates: ["2025-12", "2026-01", "2026-01-15", "2026-01-22", "2026-01-26", "2026-04-01", "2026-05-27", "2026-06", "2026-08-13", "2026-09", "2026-09-17"]
keywords: ["agents", "cybersecurity", "exploit", "intel", "nvidia", "pricing", "research", "zero-day"]
source: docs/RAG/collect-261001-general-networking/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [59, 109]
sha256: 4c9639d04e8ce53acb03ea2eb9c190d1f2bbdd45cafc10d4f331e854ed405f29
---

# Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)

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

## 2. pfSense (Netgate)

### 2.1 Release timeline 2026 (pfSense Plus, all on FreeBSD 16.0-CURRENT)

| Version | Released | Base OS | Config Rev | Notes |
|---|---|---|---|---|
| 25.11.1 | 2026-01-26 | FreeBSD 16.0-CURRENT | 24.2 | Jan 2026 patch release |
| 26.03 | 2026-04-01 | FreeBSD 16.0-CURRENT | 24.5 | Major 2026 release |
| 26.03.1 | 2026-05-27 | FreeBSD 16.0-CURRENT | 24.5 | Patch (IPv6 NDP regression reported) |
| 26.07 | 2026-08-13 | FreeBSD 16.0-CURRENT | 24.6 | Nexus controller era begins |
| 26.10 | TBD | FreeBSD 16.0-CURRENT | 24.6 | Not yet released as of Sep 22, 2026 |

**[official — https://docs.netgate.com/pfsense/en/latest/releases/versions.html]**

- **pfSense Plus 26.07 (August 13, 2026)** is the headline 2026 release. It advances the **Netgate Nexus controller** — a new **Go-based controller replacing the legacy PHP GUI** and serving as the modern foundation for pfSense Plus **[official — https://forum.netgate.com/topic/201119/netgate-releases-pfsense-plus-software-version-26.07]**.
- 26.07 exclusive Nexus features: **CoreDNS** (high-performance integrated DNS, via the new Netgate `rexdns` plugin), **Threatgate** (bulk address/domain lists for firewall rules, aliases and CoreDNS groups), and **Snort 3** (multi-threading support, faster rule syntax) available only via the Nexus GUI **[official]**.
- 26.07 security: **critical WireGuard fix for CVE-2026-58085** plus other security enhancements **[official]**.
- Other 26.07 fixes: DHCP, DNS Resolver, DynamicDNS, gateways/monitoring, IPsec, VXLAN interfaces, OpenVPN, firewall rules/NAT, traffic shaper, wireless **[official]**.
- **pfSense CE 2.9.0** exists on **FreeBSD 16.0-CURRENT** (pfREST 2.10) — reported via community/GitHub compatibility matrices, not the main Netgate release notes; treat as **[secondary]**.
- Netgate policy: new pfSense releases track **FreeBSD-CURRENT**, not RELEASE — official blog *"pfSense software is moving ahead"* confirmed they will stay on CURRENT with no return to RELEASE **[secondary — https://forum.netgate.com/topic/199339/latest-pfsense-release-25-11-uses-freebsd-16-official-release-is-december-2027]**.
- Known issue: **IPv6 NDP regression on pfSense Plus 26.03.1** (Redmine #16919, Jun 30, 2026) **[official — Redmine, https://redmine.pfsense.org/projects/pfsense-plus/issues.pdf]**.

### 2.2 Pricing and licensing 2026

