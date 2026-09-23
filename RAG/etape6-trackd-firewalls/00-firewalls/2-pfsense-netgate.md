---
id: etape6-trackd-firewalls/00-firewalls/2-pfsense-netgate
title: "2. pfSense (Netgate)"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: reference
actors: ["AWS", "Microsoft"]
dates: ["2026-01", "2026-01-15", "2026-01-26", "2026-01-28", "2026-01-29", "2026-04-01", "2026-04-17", "2026-05-27", "2026-06-15", "2026-06-16", "2026-07", "2026-07-01", "2026-07-15", "2026-08-13"]
keywords: ["aws", "license", "pricing", "research", "throughput"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [86, 162]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: 32e2e395944954fdb8b9b429310425c1dc51e85379cec8f8a22459d60b1e6619
---

# 2. pfSense (Netgate)

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

- **pfSense Plus software (3rd-party hardware, Netgate Installer): starting at $129/yr** — no feature or throughput upcharges **[official — https://www.netgate.com/pricing-pfsense-plus]**.
- Cloud (AWS/Azure marketplaces): **$0.08–$0.40/hr** for the pfSense Plus software component; free 30-day trials **[official]**.
- Netgate appliances include pfSense Plus at no charge; appliances start at **$189** (e.g., Netgate 1100 at $269 with TAC Lite) **[official]**.
- TAC support: **TAC Pro $399/yr**, **TAC Enterprise $799/yr** (Netgate 1100 options: TAC Pro 1-yr $49 / 2-yr $98 as appliance add-ons) **[official/secondary — https://aws.amazon.com/marketplace/pp/prodview-lstvktrnia26i; https://shop.netgate.com/products/1100-pfsense]**.
- **Multi-instance management (MIM) entitlements**: Netgate Nexus multi-instance management is sold **per managed device**; the controller can manage the instance it resides on by default; additional entitlements expire/renew **yearly** — this generated community backlash in 2026 about the end of the "lifetime free" 2022 Home+Lab promotion terms **[official via forum/FAQ + community — https://forum.netgate.com/topic/201119/netgate-releases-pfsense-plus-software-version-26.07; https://forum.netgate.com/topic/201061/a-new-release-candidate-for-pfsense-plus-software-version-26.07-available/18]**.
- Context: the free Home+Lab download was discontinued years earlier (abuse by third-party vendors); CE remains fully free **[secondary — https://techunwrapped.com/netgate-reverses-course-and-pfsense-plus-homelab-will-never-be-free-again/]**.

### 2.3 Netgate company/product news 2026

- No new flagship appliance model launches were verified in this research window (Netgate 1100/2100/4100/6100/8200 lineup continues; 6100 HA deployments discussed in 26.07 forum threads) **[secondary/unverified]**.
- Nexus is the strategic direction: Netgate is rebuilding management around the Go-based Nexus controller, with MIM (multi-instance management) as the monetized management plane **[official — forum announcements]**.

---

## 3. OPNsense (Deciso)

### 3.1 Release model

- OPNsense follows a **YY.M cadence** with a codename per major: majors in **January and July** each year, with bi-weekly-ish minor patches. Major upgrades are sequential (no skipping) **[secondary — https://github.com/lindolfojunior/skills-evolver/blob/HEAD/skills/opnsense-firewall/SKILL.md]**.
- License: clear and stable **2-Clause BSD** **[official — https://docs.opnsense.org/releases/CE_26.7.html]**.

### 3.2 OPNsense CE 26.1 "Witty Woodpecker" (released ~January 28, 2026)

- Released late January 2026; upgrade path from 25.7 unlocked **January 29, 2026** **[official — https://docs.opnsense.org/releases/CE_26.1.html]**.
- Headline features: almost a **full firewall MVC/API experience** (automation rules promoted to the new rules GUI); **Suricata 8 with inline inspection mode using "divert"**; IPv6 reliability/feature improvements; **router advertisements MVC/API**; full code shell-command escaping revamp; **default IPv6 mode now using Dnsmasq** for client connectivity; Unbound blocklist source selection; **automatic host discovery service** **[official]**.
- Patch series through July 2026: **26.1.10 (June 15, 2026)** — 3 core security issues, FreeBSD security advisories, source-NAT migration prep, FreeBSD 15.1 support groundwork, MVC rules GUI fixes; **26.1.11 (July 1, 2026)** — security advisories incl. config-line injection via GUI text fields, stored XSS in GPS init string, OpenVPN path traversal in common_name, plus FreeBSD security advisories (VM use-after-free, execve TOCTOU, OpenZFS, KTLS, TCP RACK, POSIX shm, audit, iconv, libalias); hotfixes **26.1.11_3 / _5 / _6 / _10** through early July **[official — https://docs.opnsense.org/releases/CE_26.1.html]**.

### 3.3 OPNsense CE 26.7 "Xenial Xenops" (released July 15, 2026)

- Released **July 15, 2026** — *"Building the foundation for the next generation of open network security"* (Deciso PR via PR Newswire) **[official/secondary — https://docs.opnsense.org/releases/CE_26.7.html; https://tools.prnewswire.com:443/en-us/live/19304/release/20260715UN05123?filter=19304]**.
- Headline changes:
  - **Interface assignments and gateway groups via MVC/API** — first implementation of the new Interface Assignments framework; establishes the foundation for future **API-driven interface management** (per Deciso CIO Ad Schellevis) **[official — https://linuxiac.com/opnsense-26-7-open-source-firewall-released-powered-by-freebsd-15-1/]**.
  - **Firewall rules now default to MVC/API**; **outbound NAT → source NAT migration assistant** **[official]**.
  - **Captive portal IPv6 support**; **Kea DDNS, custom options, dynamic prefix delegation** (dynamic prefix delegation distributes changing IPv6 prefixes downstream) **[official — https://linuxiac.com/opnsense-26-7-open-source-firewall-released-powered-by-freebsd-15-1/]**.
  - Under the hood: **FreeBSD 15.1**, **OpenSSL 3.5**, **OpenVPN 2.7**, **PHP 8.5**, **Python 3.13** **[official]**.
  - Security hardening: multiple **stored XSS fixes**, shell-safe escaping, legacy exec() removal continuing **[official]**.
- Hotfix **26.7.1_1** issued: default-gateway-switch alarm trigger fix, pecl-mcrypt 1.0.9 **[official]**.
- 25.7-series maintenance continued in parallel: **25.7.11 (January 15, 2026)** introduced the **host discovery service** (hostwatch 1.0.4, enabled by default, automatic MAC registry for IPv4/IPv6 feeding MAC-type aliases and captive portal) **[secondary — https://cybersecuritynews.com/opnsense-firewall-25-7-11-released/]**.

### 3.4 OPNsense Business Edition (Deciso commercial)

- **Business Edition 26.4 released April 17, 2026**, following Community Edition 26.1 "Witty Woodpecker" **[secondary — https://www.thefreelibrary.com/OPNsense+Business+Edition+26.4+Released.-a0883212425]**.
- BE 26.4 features: full MVC/API experience (automation rules in new rules GUI), Suricata divert inline inspection, IPv6 RA MVC/API, shell-escaping revamp, Dnsmasq default IPv6 mode, Unbound blocklist source selection, automatic host discovery, captive portal IPv6, **enhanced OpenVPN user portal with custom configurations**, **improved OpenID Connect** (identity-claim visibility; Microsoft Entra ID integration), **OPNcentral** opt-out for automatic login, **OPNWAF mod_status** diagnostics **[official — https://github.com/opnsense/docs/blob/HEAD/source/releases/BE_26.4.rst]**.
- **BE 26.4.1 (June 16, 2026)**: based on community 26.1.9 with reliability improvements **[official — https://github.com/opnsense/docs/blob/HEAD/source/releases/BE_26.4.rst]**.
- Prior BE 25.10 (Feb 9, 2026, 25.10.2): revamped frontend grid UI, experimental GUI privilege separation, firewall automation GUI, OIDC, captive portal backend rewrite, FreeBSD 14.3 **[official — https://github.com/opnsense/docs/blob/HEAD/source/releases/BE_25.10.rst]**.

### 3.5 Community and adoption notes

- The release-notes text notes "over 11 and a half years" of continuous development and credits a rising volume of core-code security reports from the community in mid-2026 (reporters acknowledged: **lujiefsi**, **Jonas Ampferl of Hacking Cult**, **Jan Kahmen of turningpoint**, **chrstnth**, **haxorton**, **Greelan**, **Konstantinos Spartalis**, **Etienne Girault**) — Deciso said it was improving its report-handling process after 26.7 **[official — https://docs.opnsense.org/releases/CE_26.1.html]**.
- No verified 2026 download/user-count figures were located in this research; Deciso positions OPNsense for home-to-enterprise use with the Community Edition free and the Business Edition as the commercial tier **[unverified — figures]**.

---

