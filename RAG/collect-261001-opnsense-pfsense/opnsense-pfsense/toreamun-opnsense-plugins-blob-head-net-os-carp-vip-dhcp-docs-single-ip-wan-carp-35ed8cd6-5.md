---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-5
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [339, 431]
sha256: d0d7381fcd21f2e90c230e83101e7ff2c3005cb6feb57af2b246c24c309304a5
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

Source NAT (Firewall ‣ NAT ‣ Source NAT, mode Hybrid, or Manual if you want to own every rule). This is ordinary OPNsense outbound NAT - nothing plugin-specific; the CARP how-to covers the same ground (Setup outbound NAT). Bind the rules to the WAN interface (the rule's Interface field). If your two nodes number the WAN under different config-keys (e.g. a physical+VM pair, section 9), bind instead to a WAN interface group whose members include both keys (e.g.opt3,wan ), so the one
synced rule resolves on both. Where a rule translates to the VIP, set the
Translate Source IP field (Translation / target on the legacy Outbound page; it
otherwise defaults to Interface address) to thewan_carp_vip alias.
  - internal to VIP (the catch-all): source snat_internal (your LAN/VLAN subnets,
or simply the RFC 1918 ranges), destination any, Translate Source IPwan_carp_vip . One rule replaces per-subnet rules and follows the lease via the
alias; if its source range covers the SYNC net it also NATs the backup's borrowed
traffic (section 7.3), so no separate backup-transit rule is needed.
  - backup transit (optional): only if the catch-all above does not cover the SYNC
net - source the SYNC net to any, Translate Source IP wan_carp_vip .
 RFC 1918 node IPs only: add a no-NAT-for-CARP rule above the catch-all - source the node range ( wan_carp_nodes ), destination224.0.0.18 , Do not NAT -
else the catch-all rewrites your CARP advertisements' source and breaks the election.
Link-local node IPs (section 6.1, the default here) do not need this rule.Why the alias, not "Interface address": the WAN interface's primary address is the link-local node IP, so translating to "Interface address" would NAT to 169.254.x ,
unroutable.wan_carp_vip is the public address. (Cosmetic GUI quirk: a Do-not-NAT
rule still shows "Interface address" in the Translate column; that field is
ignored when Do-not-NAT is ticked; it renders asno nat in pf.)
 The catch-all Source NAT rule: sourcesnat_internal , Translate Source IP = the plugin-managedwan_carp_vip alias.
- internal to VIP (the catch-all): source 
- 
Firewall (SYNC): allow SYNC net -> any , kept tight (pkg/NTP/DNS).
- 
Verify: 
  - ifconfig -g WAN on both nodes: it lists that node's real WAN device, so the
group-bound NAT resolves on each (section 9).
  - pfctl -sn | grep -E 'no nat|nat on|wan_carp' : the catch-all-> <wan_carp_vip> rule is bound to the WAN interface, and (with RFC 1918 node IPs) theno nat … to 224.0.0.18 rule sits above it.
  - pfctl -t wan_carp_vip -T show : the alias table holds the live public address
(not a stale or private one).
  - tcpdump -T carp on the WAN (advertisements), a failover test (down the master
NIC), the promoted node keeping its pinnedWAN_ISP default and NATing out the VIP,
and the VIP lease following the master.
Two supported designs for the default route; both fail over seamlessly, they differ in what the backup node can do on its own:
| Design | Default route | Backup's own internet | Pick it when | 
|---|---|---|---|
| Baseline ( Own default route by CARP role = off) | pinned to WAN_ISP on both nodes, never moves | none by default; borrow the master's on demand (7.2) | you want the fewest moving parts, and the backup rarely needs to reach out itself | 
| Role-driven ( enforce +force_down , section 8) | only the CARP master holding the lease has one; fail-stop | via the master, with the optional built-in backup egress (off by default, backup-egress.md) | you want the backup online for updates/monitoring, or a dynamic router redistributes your default | 
Never use an auto-switching gateway group for the default: it lags the CARP event and blackholes a promoted node (7.1).
Giving the backup its own internet is a convenience (self-pkg/NTP), not an HA
requirement. Both nodes share one physical WAN uplink, so a backup
monitoring that same line buys nothing (if the WAN is down, it is down for both); the
backup can skip its own internet entirely and pull pkg/NTP/DNS/config from the master
over SYNC.
The reason it has none by default: the default route on both nodes points at the ISP
gateway (WAN_ISP, reached via the VIP) and never moves. Because WAN_ISP is the sole
upstream and gateway switching is off (section 6.3 step 6), OPNsense installs that default
on both nodes and never withdraws it; the backup just holds it dormant - unusable
only because the VIP, its on-link path to the gateway, is suppressed in backup state
(section 2). That is exactly what makes failover seamless: a promoted backup already
has the right default route installed, and the moment CARP brings the VIP up on it, that
route starts working - no routing reconfiguration on role change.
Own default route by CARP role (the keeper's defaultRouteMode setting): off is
this baseline, and the rest of this section applies unchanged. enforce hands ownership
of the default to the keeper per CARP role (section 8); the built-in backup-egress
feature then gives the backup seamless internet via the master, the productized version of
the manual toggle in section 7.2 (backup-egress.md). observe
only logs what either would do, with no routing-table change.
flowchart LR
    subgraph BK["Backup state"]
        B1["Default route = WAN_ISP<br/>installed, far gateway"] --> B2["VIP suppressed in backup state<br/>gateway not reachable on-link<br/>no internet of its own"]
    end
    subgraph MS["Promoted to master"]
        M1["VIP activates"] --> M2["same default route now works<br/>straight out the VIP"]
    end
    B2 -.->|"CARP promotes this node"| M1
    classDef backup fill:#f4e1ec,stroke:#cc79a7,color:#5c2547
    classDef ok fill:#cfe6f5,stroke:#0072b2,color:#00344f
    class B1,B2 backup
    class M1,M2 ok
    If you do want the backup online independently (e.g. to run pkg on it directly),
give it internet on demand or through a role-change hook - never through
automatic gateway switching. See also OPNsense's CARP how-to,
Backup node cannot reach internet.
The tempting design is a gateway group WAN_HA = [WAN_ISP tier 1, PEER_SYNC tier 2]
set as the default with Allow default gateway switching on
(System ‣ Settings ‣ General). In steady state it even looks right: the backup's
WAN_ISP monitor is DOWN (VIP suppressed, section 2) so the backup uses tier 2 (out via the
master over SYNC), while the master uses tier 1. That part works.
It is failover-unsafe. Switching is driven by dpinger monitor transitions, which
lag (a detection interval plus an up-confirmation), and OPNsense's native gateway switching
has no CARP role-change hook to re-evaluate routing on promotion (the plugin's keeper does,
by CARP role, see section 8; core gateway groups do not). When the backup is promoted, its default stays on
tier 2 - pointing over SYNC at the node that just died - until dpinger independently
notices WAN_ISP has come UP. During that window the new master, and everything behind
it, blackholes. Compounding it:
- FreeBSD keeps a single default route (no Linux-style cost-weighted multi-default with instant failover); switching replaces the one default, it does not pre-stage a backup one.
- Per the OPNsense manual, "local generated traffic will only use the current default gateway which will not change without [gateway switching]" (MultiWAN how-to).
- A gateway group cannot be the system default directly; it steers routing only through firewall policy rules, which do not cover the firewall's own traffic (same how-to).
This is not hypothetical - in testing, the site stayed dark until the monitor caught up.
To bring the backup online for a maintenance window, temporarily point its default at the peer's SYNC IP, and revert before any failover:
# ON  - backup borrows the master's internet over SYNC (run from node B)
route -nq change -inet default 10.2.2.1        # peer (master) SYNC IP
# OFF - restore the configured WAN_ISP default
configctl interface routes configure
It is transient by design: a config reload re-installs the WAN_ISP default, so run it
