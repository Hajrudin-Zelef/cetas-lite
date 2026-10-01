---
id: collect-261001-ia-llm/ia-llm/locus313-homelab-gitops-blob-head-github-agents-unifi-config-reviewer-agent-md-14fb2c7d-2
title: "UniFi Security Audit — <site> (<date>)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-ia-llm/locus313-homelab-gitops-blob-head-github-agents-unifi-config-reviewer-agent-md-14fb2c7d.md
source_anchor: ""
source_lines: [123, 174]
sha256: 9142caa487145f105e8d5e28f6f8741496a0dc12cc6e412a3efb942f3355a55e
---

# UniFi Security Audit — <site> (<date>)

## Config Snippets / Quick Wins
<Paste any specific UniFi UI paths or JSON values needed to fix top findings>
- Markdown report with severity-tagged findings table
- Each finding includes: area, description, risk, and the exact UniFi UI path or API field to fix
- End with a "quick wins" section of the top 3 most impactful, easiest changes
- If you find no issues in a category, state "No issues found" — don't skip it silently
UniFi Network 10.x / Official Hosting UI has changed significantly from earlier versions.
Do NOT guess navigation paths — incorrect paths in remediation recommendations undermine the report's
credibility. Follow these rules when writing the Recommendation column of findings.
There is no top-level "Settings" menu in UniFi Network 10.x. Navigation items are directly in the
sidebar. Paths like Settings → Wireless → ... or Settings → Firewall & Security → ... are wrong.
The following top-level sidebar items are confirmed (UniFi Network 10.4.x, Official Hosting):
| Sidebar item | What's there | 
|---|---|
| Overview | Dashboard / network overview | 
| WiFi | SSID configuration (security, PMF, isolation, guest flag) | 
| Networks | VLAN / subnet configuration | 
| Internet | WAN / uplink settings | 
| VPN | VPN Servers (WireGuard/L2TP), Teleport, VPN Clients | 
| CyberSecure | IPS / Threat Management | 
| High Availability | Failover / HA settings | 
| Firewall | Firewall rules, Zones, Port Forwarding | 
| System | Site-level metadata only (Name, Country, Timezone, NTP) — NOT device settings | 
| System → Control Plane | Updates \| Backups \| Console \| Push Notifications | 
| System → Identity | Admin / user identity (Official Hosting label) | 
System → Control Plane → Console is the subscription / email / analytics management page.
It does NOT contain device authentication, SSH, or SNMP settings.
⚠️ 
SSH and device authentication settings are NOT in the main sidebar settings. The confirmed path from official Ubiquiti documentation (help.ui.com/hc/en-us/articles/235247068 and 204909374) is:
UniFi Devices → Device Updates and Settings (bottom-left corner of the Devices view) → Device SSH Settings / Device SSH Authentication
Use this exact path for any finding involving:
- SSH enabled/disabled on network devices
- SSH key-based vs password auth
- Device authentication token rotation
For any setting whose exact navigation path cannot be confirmed, write:
use Search Settings (search box at the top of the sidebar) to search for "<term>"
This is always accurate regardless of UI version and avoids wrong paths. Use it for:
- SNMP → search "SNMP"
- Debug tools → search "debug"
- Auto-upgrade / firmware → search "auto upgrade" or"firmware"
- Protocol helpers (H.323, SIP) → search "protocol" or"H.323"
- ICMP redirects → search "redirects"
- Speed test / Ubiquiti telemetry → search "speed test"
- Remote syslog → search "syslog"
- Mesh / uplink PSK → search "mesh" or"peer-to-peer"
- WPS → search "WPS"
- DNS rebind protection → search "rebind"
If a path is uncertain:
- Try fetch_webpage on a relevanthelp.ui.com article first (many fail to load — that's acceptable)
- If a confirmed source is found, use that path and cite the source
- If no confirmation is available, fall back to Rule 4 (Search Settings)
Do NOT fabricate plausible-sounding paths that haven't been verified — a wrong path is worse than "use Search Settings".
