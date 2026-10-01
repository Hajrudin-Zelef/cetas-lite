---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-3
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [148, 247]
sha256: 782ff6b44d7d6afa864371b4fe2e59feefb254a64e7c32eeaaff83aa89c8bc05
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

The failover speed is something we set, not the ISP. CARP declares a master
dead after ~3 missed advertisements, so with advbase 1 the switch is ~1-3 s
(lower advbase = faster, at the cost of more advertisement chatter). The ISP
only has to relearn the VIP's MAC on the new port (gratuitous ARP), which is
near-instant. This covers the failover relearn: the virtual MAC is identical on
both nodes, so the gateway's ARP entry stays valid and only the switch relearns the
port. A separate, steady-state hazard - the gateway letting the VIP's ARP entry
expire and never re-querying it - is independent of failover and is covered in
section 9.
Lab-validated, and exercised on real ISPs. The failover itself has run on real
WANs: the VIP followed the master on a CGNAT WAN, and on a live single-IP WAN a real
CARP failover promoted the backup, which served immediately (section 10). The
connection-level details here are the lab findings - a client's TCP connection carried
data both before and after a mid-connection cable-pull (pfsync had synced the
state; outbound NAT to the VIP keeps it portable across nodes, so the promoted node
continued the same connection), and the switch relearns the virtual MAC from the new
master's gratuitous ARP with no special switch config (section 9 has the lab caveat
on faithful failure injection).
Addresses in this section (169.254.255.x node WAN, 10.2.2.x SYNC, 123.123.123.x public)
are examples, substitute your own.
At a glance, the whole build (details in 6.3):
| # | Do | Where | 
|---|---|---|
| 0 | Confirm the ISP leases to the CARP virtual MAC (pre-flight below) | Interfaces ‣ [WAN] or a DISCOVER probe | 
| A | Define the aliases wan_carp_vip andsnat_internal (6.2) | Firewall ‣ Aliases | 
| 1 | Put a switch between the ISP hand-off and both nodes' WAN ports | hardware | 
| 2 | Give each node's WAN a static link-local IP ( 169.254.255.1/29 ,.2/29 ) | Interfaces ‣ [WAN] | 
| 3 | Create the CARP VIP (vhid, password, advskew 0 / 100) | Interfaces ‣ Virtual IPs | 
| 4 | Add a plugin keeper on that VIP, Follow on, alias wan_carp_vip | Interfaces ‣ Virtual IPs DHCP | 
| 5 | SYNC interface, pfsync + config sync incl. carpvipdhcp | System ‣ High Availability | 
| 6 | Add gateway WAN_ISP (Far Gateway, Upstream), delete the old<WANIF>_DHCP one | System ‣ Gateways | 
| 7 | No gateway group for the default route | System ‣ Gateways ‣ Group | 
| 8 | Source NAT: internal nets to wan_carp_vip (translation target = the alias) | Firewall ‣ NAT ‣ Source NAT | 
| 9 | Allow SYNC net -> any | Firewall ‣ Rules ‣ SYNC | 
| 10 | Verify ( pfctl ,tcpdump -T carp , a failover) | shell | 
NAT menu: the Firewall ‣ NAT menu carries both a Source NAT page and a legacy Outbound page (both are present on 26.1 and 26.7). Either can host the source-NAT rule this guide needs; the steps below use Source NAT (Firewall ‣ NAT ‣ Source NAT). See also OPNsense's CARP how-to, Setup outbound NAT.
Pre-flight: confirm the ISP serves the virtual MAC (do this first). The whole
design hinges on the ISP leasing the public address to the CARP virtual MAC
(00:00:5e:00:01:{vhid}), not only to your interface's burned-in MAC. Some ISPs bind
the single lease to the first MAC they see and NAK a REQUEST from any other MAC and
stay silent to its DISCOVER - then this design cannot work as-is and you need a
fixed-MAC-plus-CARP-gated approach instead. Check it while the WAN is still on DHCP; it
needs no static cutover, so the reboot in section 9 First cutover (a
DHCP-to-static issue) does not apply here. Two ways:
- Simplest (all GUI): temporarily set the WAN interface's MAC address to the virtual MAC under Interfaces ‣ [WAN], Apply, and see whether it pulls the public lease. The interface stays DHCP (no reboot trap), and this runs the full exchange on the virtual MAC - conclusive, but it takes a lease, so on a MAC-binding line expect a brief WAN blip and a possible cooldown (section 9 First cutover) on revert (clear the MAC field + Apply).
- Non-disturbing: send a
DISCOVER-only probe from a throwaway locally-administered
MAC (not the real virtual MAC, so a lease-binding ISP cannot associate the probe with
your VIP; the question it answers is "does this line lease to a second MAC at all") and
watch the reply in Interfaces ‣ Diagnostics ‣ Packet Capture (or
tcpdump -ni <wan> 'udp port 67 or udp port 68 or (vlan 0 and (udp port 67 or udp port 68))'):
an OFFER to the probe MAC = good, silence = MAC-bound. Keep the vlan 0 half of
the filter: some upstreams send replies 802.1Q priority-tagged, and a bare port
filter silently misses those. A DISCOVER takes no lease, so it leaves the live line
untouched. (OPNsense's dhclient has no -r/release flag; for a maximally clean
test free the current lease with a DHCPRELEASE sent another way.) See section 9
DHCP behaviour and section 10.
| Element | Value | Synced? | Note | 
|---|---|---|---|
| WAN public VIP | 123.123.123.123/24 (vhid 9) | Yes (VIP def) | Obtained via DHCP on virtual MAC 00:00:5e:00:01:09 | 
| WAN gateway (ISP) | 123.123.123.1 | - | On-link via the VIP's /24 | 
| Node A WAN (link-local) | 169.254.255.1/29 | No (per-node) | CARP advertisement source only; never egresses | 
| Node B WAN (link-local) | 169.254.255.2/29 | No (per-node) | - | 
| SYNC subnet | 10.2.2.0/30 | - | pfsync + config-sync + transit | 
| Node A SYNC | 10.2.2.1 | No (per-node) | - | 
| Node B SYNC | 10.2.2.2 | No (per-node) | - | 
| advskew A / B | 0 /100 | No (per-node) | A is intended master, preempt=1 | 
| pass | shared secret | Yes | Authenticates advertisements on the shared segment | 
vhid 9 gives virtual MAC 00:00:5e:00:01:09 (last MAC byte = vhid in hex). In
production pick an unusual vhid - see vhid collision
about shared ISP L2.
The VIP is not a network you choose: it is whatever address the ISP leases on the
virtual MAC, so it always sits in the ISP gateway's subnet. The /24 here is
illustrative; the real mask is the ISP's. With Follow dynamic DHCP address on, the plugin follows a
changed lease, including a cross-subnet renumber: it adopts the new address and, from
the DHCP ACK's subnet mask and gateway, updates the VIP prefix and the WAN gateway too,
so outbound keeps working (the same as a plain DHCP interface). If the ACK carries no
subnet mask it can only move the address, and logs a warning to fix the prefix and
gateway by hand.
Why link-local for the node IPs. Link-local (169.254.0.0/16,
RFC 3927) is valid only on the local
wire and is never routed off it - ideal for node IPs that only carry CARP on the WAN
segment. Assign any address in that range statically; a block inside 169.254.255.0/24
stays clear of 169.254.1.0-169.254.254.255 (the range OS auto-configuration draws
from), so nothing ever auto-assigns into yours - and avoid the block's network and
broadcast (.0/.7 in a /29). Because these IPs only source CARP advertisements and
never egress, keeping them link-local means the outbound catch-all NAT
(section 6.3 step 8) never matches them, so this setup needs no
no-NAT-for-CARP rule and no extra alias. Verified on OPNsense 26.7: Block bogon
networks (blockbogons) does not block 169.254.0.0/16 (negated in the bogons
table), and Block private networks (blockpriv) was off on this interface, so the
GUI accepts a static link-local WAN
IP and CARP elects and fails over normally. (Caveat: since the catch-all matches RFC
1918 and not link-local, traffic the backup sources from its own link-local WAN IP
is not NAT'd - in practice only the WAN gateway probe, which is meant to fail on the
backup anyway.)
Alternative - RFC 1918 node IPs (e.g. 10.1.1.0/30) work too, at the cost of extra
configuration: being RFC 1918 they are caught by the outbound catch-all, which would
rewrite the source of your CARP advertisements and break the election. You then need a
wan_carp_nodes alias and a no-NAT-for-CARP rule (section 6.3 step 8). Link-local avoids both.
