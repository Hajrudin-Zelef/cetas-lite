---
id: etape6-trackd-firewalls/00-firewalls/3-4-opnsense-business-edition-deciso-commercial
title: "3.4 OPNsense Business Edition (Deciso commercial)"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2026-01-15", "2026-04-17", "2026-06-16", "2026-07-15"]
keywords: ["research"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [139, 162]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: 27abba3b48e608732ad23adcf0e873970dd94c5905c93eb4e48d64c378331703
---

# 3.4 OPNsense Business Edition (Deciso commercial)

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

