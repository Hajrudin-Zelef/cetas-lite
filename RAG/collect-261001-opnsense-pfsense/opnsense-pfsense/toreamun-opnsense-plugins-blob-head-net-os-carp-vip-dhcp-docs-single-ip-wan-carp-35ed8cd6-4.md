---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-4
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [248, 338]
sha256: 71b670f6a2b29b6fe99917254fefd2db8321971a0630f5192c1da85e4c5c61ab
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

Set these up under Firewall ‣ Aliases before writing any rule. Referring to names instead of raw addresses keeps the ruleset readable, and for the VIP it lets the public address change without touching a single rule.
| Alias | Type | Content | Used by | 
|---|---|---|---|
| wan_carp_vip | Host | (plugin-managed, the live public VIP) | Source-NAT target; any rule that must follow the WAN address | 
| snat_internal | Network | your LAN/VLAN subnets (or the RFC 1918 ranges) | The single "internal to VIP" source-NAT rule | 
wan_carp_vip is created and owned by the plugin: give the keeper a Sync firewall alias name
and it ensures a Host alias of that name exists and keeps its content equal to the live
VIP, updated on every lease change. You may pre-create the Host alias yourself (handy if
you want to reference it in rules before the keeper runs): on start the plugin adopts a
matching pre-existing alias - it must be type Host - stamps its marker, and drives the
content from then on. Point Source NAT (and anything address-dependent)
at it and those rules follow the address with no ruleset reload. Do not hand-edit
it. Editing the alias out-of-band does not reach the live pf table until a full
filter reload (a plain refresh_aliases leaves the table stale), and the plugin
reconciles it back anyway. Let the plugin drive it. (Likewise, disabling the keeper's
service does not release the alias; the reconcile keys on the keeper's Enabled
box, so un-tick Enabled to hand the alias back.)
Work through these on node A first; confirm it holds the lease and reaches the internet, then repeat the per-node parts on node B. Config-synced items (aliases, the keeper) only need doing once.
You can stop after node A and add HA later. One node is already a complete single-IP firewall: the CARP VIP simply stays master, the plugin keeps the lease, and NAT works. Steps 1-4 and 6-8 stand alone; the SYNC bits (steps 5 and 9, plus config-sync) and node B's per-node setup come in only when you add the peer. Standing node A up alone first lets you validate the DHCP-on-a-virtual-MAC part before layering on HA.
- 
WAN-front switch physically between the ISP hand-off and both nodes' WAN ports.
- 
WAN interface per node: static link-local IP ( 169.254.255.1/29 on A,.2/29 on B; an RFC 1918 range works too but needs the no-NAT-for-CARP rule, section 6.1).
Because this interface is static (the public address lives on the VIP, not here),
OPNsense does not auto-create a gateway the way it does for a DHCP WAN - you addWAN_ISP by hand in step 6. The ISP gateway (123.123.123.1 ) is not in this
subnet, so mark it Far Gateway, otherwise OPNsense rejects the off-subnet gateway.
On-link reachability to.1 comes from the VIP's public/24 on the same interface.
- 
CARP VIP 123.123.123.123/24 , vhid 9,pass , advskew 0/100, under
Interfaces ‣ Virtual IPs (OPNsense how-to:
Setup Virtual IPs).
You don't need to know the final public address: seed the VIP with your current
WAN's public IP (the Interfaces overview shows it), or any routable placeholder if the
line is fresh. With Follow dynamic DHCP address on (step 4) the keeper DHCPs on the
virtual MAC and rewrites the VIP - address, prefix and gateway - to whatever the ISP
actually hands out (section 6.1).
- 
Plugin os-carp-vip-dhcp: a keeper on the VIP with Follow dynamic DHCP address on, and set its Sync firewall alias to wan_carp_vip (section 6.2). Only the CARP master
sources the virtual MAC: the backup stays passive (it holds no lease of its own and
transmits no DHCP) until it is promoted, so the two nodes never flap the virtual MAC on
the shared WAN switch. On failover the new master acquires the lease at once and nudges
the gateway, so failover stays CARP-speed. (This holds once both nodes run this
version or later; during a rolling upgrade a not-yet-upgraded backup still sources the
MAC, so upgrade the node that is currently CARP backup first.) Leave Client MAC
override (chaddr) blank: the keeper uses the CARP virtual MAC
automatically, which is what the lease, the ISP binding and the VIP's traffic all need.
Only set it for an ISP reservation on a fixed different MAC, and never to this node's own
NIC MAC (that binds the lease to one MAC while the VIP still sends from the CARP MAC, so an
IP-source-guard ISP blackholes the gateway).
 The keeper, advanced mode on: Follow on, aliaswan_carp_vip , Client MAC override blank. Own default route by CARP role stays off in the baseline design (section 7).
- 
SYNC interface: 10.2.2.1/30 /.2/30 ; pfsync + XMLRPC config-sync, under
System ‣ High Availability (OPNsense how-to:
Setup pfSync and HA sync).
Includecarpvipdhcp in the synchronized services so the keeper config replicates;
the CARP VIPs stay per-node (advskew differs). OPNsense syncs on demand, not on save:
after every later keeper change on the master, go to System ‣ High Availability ‣
Status and press the keeper's Synchronize and Restart button, which copies the
config to the backup and restarts that keeper with it. Synchronize and reconfigure
all also works but restarts every service on the backup, and it does not start a
keeper you have just added (start it from its own row after the sync). The cron job
HA update and reconfigure backup copies the config but does not restart the
keepers, so the backup keeps running its previous settings (for example a stale
client MAC override) until you restart them.
- 
Gateway: add WAN_ISP (123.123.123.1 , on the WAN interface, Far Gateway,
see step 2) under System ‣ Gateways and mark it Upstream Gateway so it is the
system default. There is no "pick the default" dropdown on 26.7: the default is the
upstream gateway with the best priority
(Gateways manual). Leave Allow
default gateway switching OFF and, for a single-uplink pair, gateway monitoring off
(section 7.1, section 9) - there is nothing to switch to. You do not need aPEER_SYNC gateway
object; the optional backup path (section 7.2) routes straight to the peer's SYNC IP.
 The WAN gateway (namedWAN_TN_GW here,WAN_ISP in the text): Upstream, Far Gateway, monitoring off. Mark Gateway as Down is ticked because this pair runs the role-driven design (section 8); leave it unticked in the baseline design.Gotcha - the leftover DHCP gateway after a DHCP-to-static WAN. This design turns the WAN interface from DHCP to static (the public address lives on the VIP, not here). Two things can then quietly leave the node with no default route, even while it is CARP master and owns the VIP: 
  - OPNsense does not remove the gateway that belonged to the old DHCP interface
(named <WANIF>_DHCP ). It can linger, still flagged as the default gateway but now
with an empty address, so no IPv4 default route installs. Delete it under
System ‣ Gateways after the cutover.
  - On OPNsense 26.7 and later, gateways live in an MVC model that the legacy config is
migrated into once; afterwards the model is authoritative. Create WAN_ISP through the GUI so it lands in the model. A gateway written only into the legacy<gateways> section afterwards (by hand, or via a restored/templated config) is not
loaded.
 Symptom: the node owns the VIP as CARP master but has no default route or egress, and configctl interface gateways list does not listWAN_ISP .
- OPNsense does not remove the gateway that belonged to the old DHCP interface
(named 
- 
No gateway group for the default route. The single upstream WAN_ISP from step 6
is the system default; keeping it pinned there is what makes failover seamless. Do
not point the default at a[WAN_ISP, PEER_SYNC] group with switching - it is
failover-unsafe (section 7.1). A gateway group still has its place for policy-routed transit
traffic, just not for the firewall's own default route.
- 
