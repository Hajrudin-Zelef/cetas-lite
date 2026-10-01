---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-1
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "exploit"]
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [1, 83]
sha256: 2f4e77e8b7454f8204e27c37ae30d16574171c0a59c130d7c4d8fc33b2e3b1ee
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

This guide builds a two-node OPNsense HA pair on a WAN where the ISP hands out one DHCP address, the classic "single-IP" obstacle to CARP on the WAN side. The os-carp-vip-dhcp plugin keeps that single public lease alive on a virtual CARP MAC, so the address floats between the nodes on failover.
Status: field-validated. The full topology runs on a live single-IP DHCP WAN, including a real CARP failover, on top of an isolated two-node lab and a real CGNAT WAN. It is one deployment so far, not yet exercised across many ISP/CPE combinations. Section 10 lists what is validated piece by piece. One design was rejected on the way: an automatic gateway-group tier flip for the backup's internet is failover-unsafe (section 7.1); use the built-in backup-egress feature or the on-demand toggle instead.
All addresses below are examples, substitute your own. The per-node WAN IPs are
link-local (169.254.0.0/16); the other private ranges (SYNC, internal) are
RFC 1918; the public side uses an
arbitrary, good-looking address (123.123.123.123) purely for illustration.
In a hurry? Jump to Implementation steps for the click-by-click setup; the sections before it explain why it works.
CARP vocabulary used in this guide (expand if advskew, pfsync or demotion counter are new to you)
| Term | Meaning | 
|---|---|
| CARP | Common Address Redundancy Protocol - floats a virtual IP between two nodes on a shared segment ( carp(4) ) | 
| VIP | Virtual IP - the floating address a CARP group owns and answers on | 
| vhid | Virtual Host ID - identifies a CARP group on the segment and sets its virtual MAC 00:00:5e:00:01:{vhid} | 
| virtual MAC | The CARP-derived MAC 00:00:5e:00:01:{vhid} the master answers the VIP's ARP with; identical on both nodes, so the address floats on failover without the gateway relearning it | 
| advbase | CARP advertisement base interval - the seconds between a node's CARP advertisements | 
| advskew | Skew added to advbase to set a node's CARP advertisement interval; the lowest wins the election (preferred master) | 
| pass | The CARP passphrase that authenticates advertisements. carp(4) /ifconfig call itpass ; the OPNsense GUI labels the same field "Password" | 
| preempt | The FreeBSD CARP sysctl net.inet.carp.preempt (1 = on, the OPNsense default) - not a per-VIP GUI field - that lets a recovered node with the loweradvskew take the master role back | 
| demotion counter | FreeBSD CARP counter ( net.inet.carp.demotion ) added toadvskew ; raised when an interface goes down (bynet.inet.carp.ifdown_demotion_factor , default 240) orpfsync is mid-sync, to hand the role to the peer | 
| pfsync | Kernel protocol that replicates firewall connection state between the two nodes | 
| dpinger | OPNsense's gateway-monitoring daemon | 
| SYNC | The dedicated inter-node link carrying pfsync + config-sync | 
| link-local | 169.254.0.0/16 addresses (RFC 3927) - valid only on the local segment, never routed (see IP and CARP plan) | 
| L2 (Layer 2) | The switched Ethernet segment (data-link layer, one broadcast domain) the two nodes share; CARP elects and fails over within it, with no router in between | 
Goal: seamless firewall failover (hot-warm HA) without depending on the ISP handing out more than one IP.
The problem: classic CARP on a WAN wants three IPs on the WAN segment:
| IP | Role | 
|---|---|
| Node A's own | Source for CARP advertisements + node A's own WAN access | 
| Node B's own | Same, for node B | 
| Floating VIP | The address services answer on / NAT out of | 
Most ISPs give you one DHCP address (here: gateway 123.123.123.1, one leased
public IP). That leaves you two short. This document works around that.
A small static block (e.g. a /30) has the same shortage. A /30 gives two
usable IPs - still one short of three. The topology below applies unchanged, with
one simplification: a static public IP needs no lease-keeping, so you skip the
DHCP/plugin part (section 3 step 2 / section 6.3 step 4) and simply assign the public address to
the CARP VIP (or bind it as an IP-alias VIP to the CARP VIP). Everything else -
the per-node WAN IPs for CARP, and the backup's own-internet handling (section 7) - is
identical. The plugin is only needed when that single public address is
handed out by DHCP.
2 CARP mechanics (from carp(4), FreeBSD + OpenBSD)
The facts that drive the design:
- Failover needs only a shared Layer 2 (L2) segment. The master is elected via advertisements -
link-local IP multicast (224.0.0.18 , proto 112) that never leaves the segment -
carryingvhid ,advbase ,advskew and a crypto checksum over the VIP prefixes +pass . Nodes need no routing between them, just a shared segment and each other's
presence.
- advskew /preempt : lowestadvskew becomes master;preempt=1 lets the
intended master take the role back.
- Auto-demotion: FreeBSD raises the CARP demotion counter (added to advskew when computing the advertisement interval) when a vhid's interface goes down
(net.inet.carp.ifdown_demotion_factor=240 ) orpfsync is mid-sync, so the master demotes itself and
the backup takes over.
- Virtual MAC 00:00:5e:00:01:{vhid} : the master answers ARP for the VIP with
this address.
- Backup state suppresses the VIP. A vhid address in BACKUP state is not active
on the interface, so the backup has no address in the VIP's subnet and no
connected route to the ISP gateway. This, not source-address selection, is the
real reason the backup's gateway monitor fails and why the backup has no internet of
its own (section 7). It is not a safe trigger for automatic gateway switching (section 7.1).
- Important: OpenBSD's warning that the carp device must share a subnet with
the CARP VIP applies only to balancing mode. In ordinary master/backup,
a non-routable node IP (link-local or RFC 1918) plus a public VIP in a different subnet works fine -
that is exactly what we exploit.
- A WAN-front switch gives both nodes access to the same WAN segment - the physical prerequisite everything else builds on. Any ordinary switch works - it just has to keep both nodes and the ISP hand-off in one L2 broadcast domain; managed or unmanaged, L2-only or L3-capable (as long as you don't route this segment).
- The CARP VIP owns the single public IP, obtained over DHCP on the virtual
CARP MAC (00:00:5e:00:01:{vhid} ) via the
os-carp-vip-dhcp plugin. The lease follows the master on
failover.
- Each node is assigned a small link-local static IP (169.254.x ) on the WAN
interface - set by hand, not via DHCP - used only for CARP advertisements + node
identity. The ISP
never routes it; the CARP advertisements are link-local multicast (224.0.0.18 )
that stays on the WAN segment. (The ISP's on-segment access gear does still see the
master's physical MAC as source - CARP advertisements and egress both use it - plus
the virtual MAC; the backup stays silent on the WAN until it takes over, when that
physical MAC flips to it. See section 9 for the strict one-MAC-per-port case.)
Because that range is not RFC 1918, the outboundinternal to VIP SNAT catch-all never
matches the node IP as a source, so it needs no no-NAT-for-CARP carve-out - one rule
fewer (IP and CARP plan). (Egress still leaves via the VIP: that SNAT's target is set
explicitly to the VIP, not the interface's link-local primary - which is unroutable -
see section 6.3 step 8.)
- The default route stays pinned to the ISP gateway (via the VIP) on both nodes, so a promoted backup routes out the instant it owns the VIP - failover needs no routing change. (This is the baseline; section 7 has the role-driven alternative, where the master owns the default and the backup gets internet via the master.) In backup state a node has no internet of its own; if it must self-update, it borrows the master's path over SYNC on demand (section 7), never via automatic gateway switching (failover-unsafe, section 7.1).
Direct VIP vs. IP-alias - why the public address sits straight on the CARP VIP
