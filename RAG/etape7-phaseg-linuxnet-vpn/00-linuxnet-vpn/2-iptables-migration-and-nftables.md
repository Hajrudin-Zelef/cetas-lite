---
id: etape7-phaseg-linuxnet-vpn/00-linuxnet-vpn/2-iptables-migration-and-nftables
title: "2. iptables migration and nftables"
domain: step-7g-linux-networking-and-access-nat-firewalls-ssh-vpns-a
role: deep-dive
task: reference
actors: []
dates: ["2026-07", "2026-09-22"]
keywords: []
source: docs/RAG/etape7_phaseG_linuxnet_vpn.md
source_anchor: ""
source_lines: [63, 118]
section: "Step 7G — Linux Networking and Access: NAT, Firewalls, SSH, VPNs, and Security Monitoring"
sha256: d384570ddf1e084e88b4a77f8098414f8b020ea6b0f82047da4415a040392881
---

# 2. iptables migration and nftables

## 2. iptables migration and nftables

### 2.1 Why nftables replaced iptables

- nftables is the successor framework to iptables/ip6tables/arptables/ebtables, providing a single unified syntax for IPv4, IPv6, ARP, and bridge filtering instead of four separate toolsets [secondary: https://dargslan.com/blog/nftables-vs-iptables-2026-comparison-migration-guide retrieved 2026-09-22].
- Version 1.1.6 of nftables is listed as the current target on modern distributions as of mid-2026, and it is the default backend on modern distros [secondary: https://github.com/yaleizhou/learn-skills.dev/blob/HEAD/data/skills-md/iuliandita/skills/networking/SKILL.md, July 2026 table, retrieved 2026-09-22].
- Key architectural advantages over iptables: atomic rule-set updates (the whole ruleset is committed in one transaction), native sets and maps for high-cardinality matching (IP sets without external ipset), a single rule table for both address families (`inet` family), and a more expressive bytecode VM in the kernel [secondary: https://dargslan.com/blog/nftables-vs-iptables-2026-comparison-migration-guide retrieved 2026-09-22].
- Most distributions ship an `iptables-nft` compatibility shim, so legacy `iptables` commands keep working while being translated to nftables rules underneath; this is why mixed environments appear functional but should still be migrated to native syntax for maintainability [secondary: https://dargslan.com/blog/nftables-vs-iptables-2026-comparison-migration-guide retrieved 2026-09-22].

### 2.2 Migration workflow

- The recommended migration path is: inventory existing iptables rules (`iptables-save`), convert them with `iptables-translate` (for individual rules) or `iptables-restore-translate` (for full rulesets), review the generated nftables output, load it with `nft -f`, and then disable the legacy tools [secondary: https://dargslan.com/blog/nftables-vs-iptables-2026-comparison-migration-guide retrieved 2026-09-22].
- The `iptables-translate` output is a starting point, not a finished ruleset: it produces literal translations that miss opportunities to use native sets, maps, and verdict dictionaries, so a manual cleanup pass is expected [secondary: https://dargslan.com/blog/nftables-vs-iptables-2026-comparison-migration-guide retrieved 2026-09-22].
- Common migration gotchas: chain priority values differ from iptables table ordering (raw/mangle/nat/filter map to nftables hook priorities), the `filter` table default policies need explicit restatement, and logging rules should be rate-limited with `limit rate` to avoid log floods [secondary: https://dargslan.com/blog/nftables-vs-iptables-2026-comparison-migration-guide retrieved 2026-09-22].
- A representative nftables base ruleset for a server (from operational guides) uses `table inet filter` with `input`/`forward`/`output` base chains, established/related acceptance, loopback acceptance, ICMP rate limiting, SSH gating, and a final drop — the direct structural equivalent of a typical iptables INPUT policy-drop setup [secondary: https://github.com/armourinfosec/linux-administration-and-server-hardening/blob/HEAD/Security-Firewall-and-Monitoring/nftables-Configuration-and-Management.md retrieved 2026-09-22].

### 2.3 nftables operational commands

```bash
nft list ruleset                 # full active ruleset
nft -f /etc/nftables.conf        # atomic load of a ruleset file
nft flush ruleset                # drop everything (careful on remote hosts)
nft add element inet filter ssh_allow { 203.0.113.7 }   # dynamic set update
nft monitor trace                # packet tracing for debugging drops
```

- Dynamic sets with timeouts (`flags timeout`) are the native replacement for fail2ban-style temporary bans at the nftables layer: an address can be added with an expiry instead of inserting and later deleting rules [secondary].
- `nft monitor trace` is the primary packet-path debugger, showing which chain and rule handled a packet — far more informative than iptables' counters alone [secondary: https://github.com/armourinfosec/linux-administration-and-server-hardening/blob/HEAD/Security-Firewall-and-Monitoring/nftables-Configuration-and-Management.md retrieved 2026-09-22].

---

## 3. UFW and firewalld

### 3.1 Positioning

- UFW (Uncomplicated Firewall) is the default human-friendly frontend on Debian and Ubuntu; firewalld is the default zone-based frontend on RHEL, Rocky Linux, AlmaLinux, and Fedora [secondary: https://github.com/mfahad710/devops_documentations/blob/HEAD/Linux/Firewall-UFW-Firewalld.md retrieved 2026-09-22].
- Both are frontends: UFW historically drove iptables and on modern systems drives nftables via the compatibility layer, while firewalld natively manages nftables rulesets through its zone/policy model [secondary: https://blog.hofstede.it/linux-firewalls-how-to-actually-secure-a-cloud-server-iptables-nftables-firewalld-ufw/ retrieved 2026-09-22].
- UFW's model is a simple ordered rule list (`ufw allow 22/tcp`, `ufw deny from 203.0.113.0/24`); firewalld's model is zones (trust levels assigned to interfaces/sources) plus services and rich rules, which fits multi-homed servers better [secondary: https://github.com/peterbamuhigire/linux-skills/blob/HEAD/07-security-and-hardening/linux-firewall-ssl/SKILL.md retrieved 2026-09-22].

### 3.2 Comparison matrix

| Dimension | UFW | firewalld |
|---|---|---|
| Default on | Debian/Ubuntu family [secondary] | RHEL/Rocky/Alma/Fedora family [secondary] |
| Mental model | ordered allow/deny rules [secondary] | zones + services + rich rules [secondary] |
| Runtime changes | rules apply immediately; `--dry-run` available [secondary] | `--permanent` vs runtime split; `firewall-cmd --reload` [secondary] |
| nftables native | via compatibility translation [secondary] | native nftables backend [secondary] |
| Learning curve | minimal [secondary] | moderate (zones, policies) [secondary] |
| Best fit | single-purpose servers, homelab [secondary] | multi-interface servers, dynamic networks [secondary] |

- A recurring operational warning: enabling UFW with `ufw enable` over SSH without first allowing port 22 locks the administrator out; the guides uniformly recommend `ufw allow ssh` (or the custom port) before enabling [secondary: https://blog.hofstede.it/linux-firewalls-how-to-actually-secure-a-cloud-server-iptables-nftables-firewalld-ufw/ retrieved 2026-09-22].
- firewalld's `--permanent` vs runtime distinction is the most common source of "rules disappeared after reboot" or "rules didn't apply" confusion; changes must be made permanent and reloaded, or applied with both flags [secondary: https://github.com/mfahad710/devops_documentations/blob/HEAD/Linux/Firewall-UFW-Firewalld.md retrieved 2026-09-22].
- On cloud servers, the host firewall is only one layer: cloud security groups act as an outer filter, and practitioners commonly run a default-deny host firewall even behind security groups for defense in depth [secondary: https://blog.hofstede.it/linux-firewalls-how-to-actually-secure-a-cloud-server-iptables-nftables-firewalld-ufw/ retrieved 2026-09-22].

---

