---
id: etape6-trackd-firewalls/00-firewalls/2-pfsense-netgate
title: "2. pfSense (Netgate)"
domain: step-6-track-d-fortigate-pfsense-opnsense-firewalls-network-
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-01", "2026-01-26", "2026-01-28", "2026-01-29", "2026-04-01", "2026-05-27", "2026-06-15", "2026-07", "2026-07-01", "2026-07-15", "2026-08-13"]
keywords: ["aws", "license", "pricing", "research", "throughput"]
source: docs/RAG/etape6_trackD_firewalls.md
source_anchor: ""
source_lines: [86, 138]
section: "Step 6 — Track D: FortiGate + pfSense + OPNsense (Firewalls & Network Security)"
sha256: e3c210ebe1a5c47355b53001aed12de92e76f7aebf03712299388af32470b38f
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

