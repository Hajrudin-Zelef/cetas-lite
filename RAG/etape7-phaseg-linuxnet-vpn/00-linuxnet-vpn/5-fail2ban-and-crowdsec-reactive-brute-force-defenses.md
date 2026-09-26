---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/5-fail2ban-and-crowdsec-reactive-brute-force-defenses
title: "5. fail2ban and CrowdSec: reactive brute-force defenses"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["agent", "agents", "incident", "intel"]
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [168, 197]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: 08c688a341f84e64903e7bb7e895c33265cff3765fee4aa544a9b019c955586c
---

# 5. fail2ban and CrowdSec: reactive brute-force defenses

## 5. fail2ban and CrowdSec: reactive brute-force defenses

### 5.1 Positioning

- fail2ban and CrowdSec are reactive, log-driven controls: they watch authentication logs and ban offending IPs. Practitioner comparisons are explicit that key-only SSH with root/password login disabled is the foundational control; fail2ban/CrowdSec are supplemental, not primary, defenses [secondary: https://github.com/homesecexplorer/videos/blob/HEAD/Fail2banVsCrowdsec.md retrieved 2026-09-22].
- fail2ban is the classic single-host tool: "jails" pair a log filter with a ban action (typically an iptables/nftables rule) for a fixed duration [secondary: https://github.com/vietanhdev/bulwark/blob/HEAD/docs/articles/fail2ban-vs-crowdsec-vs-denyhosts.md retrieved 2026-09-22].
- CrowdSec is the collaborative successor model: a local agent parses logs against "scenarios", remediation is applied by "bouncers" (firewall, nginx, Cloudflare, and other integrations), and signals can be shared with a community blocklist so participants benefit from others' detections [secondary: https://github.com/davidjameshowell/mailcow_crowdsec retrieved 2026-09-22].

### 5.2 Comparison matrix

| Dimension | fail2ban | CrowdSec |
|---|---|---|
| Detection scope | local host logs [secondary] | local logs plus optional community threat intel [secondary] |
| Ban mechanism | firewall rule insert/delete (jails) [secondary] | bouncers for firewall, WAF, CDN layers [secondary] |
| Configuration style | INI jails and filters [secondary] | YAML scenarios, collections, hub [secondary] |
| Ecosystem | mature, huge filter library [secondary] | growing hub of scenarios and bouncers [secondary] |
| Resource profile | lightweight Python daemon [secondary] | heavier (Go agent plus optional local API) [secondary] |

- A practical integration example is CrowdSec protecting a Mailcow email stack via its firewall bouncer [secondary: https://github.com/davidjameshowell/mailcow_crowdsec retrieved 2026-09-22].
- Hestia Control Panel's commit history shows hosting panels adopting CrowdSec alongside fail2ban-era tooling, reflecting the broader shift toward collaborative banning [secondary: https://github.com/hestiare/hestia/commit/cb1e1be1705c4a1bcb541e20314829f309fe0835 retrieved 2026-09-22].

### 5.3 Operational guidance

- Do not run overlapping fail2ban jails and CrowdSec scenarios on the same log sources without coordination: double bans create confusing, hard-to-expire firewall states [secondary: https://github.com/homesecexplorer/videos/blob/HEAD/Fail2banVsCrowdsec.md retrieved 2026-09-22].
- Whitelist management IPs, monitoring hosts, and VPN egress ranges before enabling aggressive banning; self-inflicted lockouts of the admin's own static IP are the most reported operational incident with both tools [secondary].
- On OpenSSH 9.8+, `PerSourcePenalties` provides in-daemon connection rate limiting and penalties, which one practitioner reference describes as making fail2ban "largely optional" for SSH specifically — though it does not replace log-driven banning for other services [secondary: https://github.com/fuyutarow/dotfiles/blob/HEAD/agents/skills/securing-remote-access/SKILL.md retrieved 2026-09-22].
- Neither tool compensates for password authentication left enabled: with passwords on, banning only slows credential stuffing rather than stopping it [secondary].

---

