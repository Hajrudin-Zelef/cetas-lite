---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-67-4
title: "Fast Path"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-67.md
source_anchor: ""
source_lines: [269, 379]
sha256: 28897d5675f87212606d821b468eb0c2eb260e1f1e378670bce7c5ae5d3b431d
---

# Fast Path

For example traffic from src:111.15.0.1 to dst:111.13.0.1

Packet will be seen in the output and postrouting chains, because now it is locally originated packet with source MAC address equal to vrfInterface:

On the other hand, for packets routed in the direction CE→PE will be seen in the "forward" as any other routed IP traffic before being sent to the MPLS:

But there can be an exception. If destination IP of reply packet is local to the router and connection tracking is performing NAT translation then connection tracking will "force" packet to move through the firewall prerouting/forward/postrouting chains.

## Logical Interfaces

So far we looked at examples when in or out interfaces are actual physical interfaces (Ethernet, wireless), but how packets will flow if the router receives tunnel encapsulated packets?

Let's assume that there is an IPIP packet coming into the router. Since it is a regular IPv4 packet it will be processed through all routing-related facilities ( until "J" in the diagram). Then the router will look if the packet needs to be decapsulated., in this case, it is an IPIP packet so "yes" send the packet to decapsulation. After that packet will go another loop through all the facilities but this time as a decapsulated IPv4 packet.

It is very important because the packet actually travels through the firewall twice, so if there is a strict firewall, then there should be "accept" rules for IPIP encapsulated packets as well as decapsulated IP packets.

Packet encapsulation and decapsulation using a bridge with enabled `vlan-filtering` do not relate to logical interfaces. See more details in the bridging section.

## IPSec Policies

Let's take a look at another tunnel type - IPSec. This type of VPN does not have logical interfaces but is processed in a similar manner.

# Fast Path

From what we learned so far, it is quite obvious that such packet processing takes a lot of CPU resources. To fast things up FastPath was introduced in the first RouterOS v6. What it does is it skips processing in the Linux kernel, basically trading some RouterOS functionality for performance. For FastPath to work, interface driver support and specific configuration conditions are required.

## How Fast Path Works

FastPath is an interface driver extension, that allows a driver to talk directly to specific RouterOS facilities and skip all others.

The packet can be forwarded by a fast path handler only if at least the source interface supports a fast path. For complete fast-forwarding, destination interface support is also required.

Currently, RouterOS has the following FastPath handlers:

- IP traffic processing
- FastTrack
- Traffic Generator
- MPLS
- Bridge

FastPath handler is used if the following conditions are met:

- firewall rules are not configured;
- simple queue or queue trees with *parent=global* are not configured;
- no mesh, metarouter interface configuration;
- sniffer or torch is not running;
- connection tracking is not active;
- IP accounting is disabled;
- VRFs are not configured (`/ip route vrf` is empty);
- A hotspot is not used (`/ip hotspot` has no interfaces);
- IPSec policies are not configured;
- `/tool mac-scan` is not actively used;
- `/tool ip-scan` is not actively used.

Packets will travel the FastPath way if FastTrack is used no matter if the above conditions are met.

Traffic Generator automatically use FastPath if the interface supports this feature.

Currently, MPLS fast-path applies to MPLS switched traffic (frames that enter router as MPLS and must leave router as MPLS) and VPLS endpoint that do VPLS encap/decap. Other MPLS ingress and egress will operate as before.

A Bridge handler is used if the following conditions are met:

- there are no bridge Calea, filter, NAT rules;
- *use-ip-firewall* is disabled;
- no mesh, MetaRouter interface configuration;
- sniffer, torch, and traffic generator are not running;
- bridge vlan-filtering is disabled (condition is removed since RouterOS 7.2 version);
- bridge dhcp-snooping is disabled.

FastPath on the vlan-filtering bridge does NOT support priority-tagged packets (packets with VLAN header but VLAN ID = 0). Those packets are redirected via a slow path.

Interfaces that support FastPath:

| RouterBoard | Interfaces | 
|---|---|
| **RB6xx series** | ether1,2 | 
| **RB800** | ether1,2 | 
| **RB1100 series** | ether1-11 | 
| **All devices**  | Ethernet interfaces | 
|  | wireless interfaces | 
|  | bridge interfaces | 
|  | VLAN, VRRP interfaces | 
|  | bonding interfaces (RX only) | 
|  | PPPoE, L2TP interfaces | 
|  | EoIP, GRE, IPIP, VXLAN interfaces. | 
|  | VPLS (starting from v7.17) | 

EoIP, Gre, IPIP, VXLAN and L2TP interfaces have per-interface setting *allow-fast-path.* Allowing a fast path on these interfaces has a side effect of bypassing firewall, connection tracking, simple queues, queue tree with parent=global, IP accounting, IPsec, hotspot universal client, vrf assignment for encapsulated packets that go through a fast-path. Also, packet fragments cannot be received in FastPath.

Whether FastPath is being used can be verified with `/interface print stats-detail`

Only interface queue that guarantees FastPath is only-hardware-queue. If you need an interface queue other than hardware then the packet will not go fully FastPath, but there is not a big impact on performance, as "interface queue" is the last step in the packet flow.

The packet may go Half-FastPath by switching from FastPath to SlowPath, but not the other way around. So, for example, if the receiving interface has FastPath support, but the out interface does not, then the router will process the packet by FastPath handlers as far as it can and then proceed with SlowPath. If the receiving interface does not support FastPath but the out interface does, the packet will be processed by SlowPath all the way through the router.

# FastTrack

Fasttrack can be decoded as Fast Path + Connection Tracking. It allows marking connections as "fast-tracked", marking packets that belong to fast-tracked connection will be sent fast-path way. The connection table entry for such a connection now will have a fast-tracked flag.

FastTrack packets bypass firewall, connection tracking, simple queues, queue tree with parent=global, ip traffic-flow, IP accounting, IPSec, hotspot universal client, VRF assignment, so it is up to the administrator to make sure FastTrack does not interfere with other configuration!

To mark a connection as fast-tracked new action was implemented "*fasttrack-connection"* for firewall filter and mangle. Currently, only TCP and UDP connections can be fast-tracked and to maintain connection tracking entries some random packets will still be sent to a slow path. This must be taken into consideration when designing firewalls with enabled "fasttrack".

FastTrack handler also supports source and destination NAT, so special exceptions for NATed connections are not required.

The easiest way to start using this feature on home routers is to enable "fasttrack" for all *established, related* connections:

Notice that the first rule marks established/related connections as fast-tracked, the second rule is still required to accept packets belonging to those connections. The reason for this is that, as was mentioned earlier, some random packets from fast-tracked connections are still sent the slow pathway and only UDP and TCP are fast-tracked, but we still want to accept packets for other protocols.

After adding the "FastTrack" rule special dummy rule appeared at the top of the list. This is not an actual rule, it is for visual information showing that some of the traffic is traveling FastPath and will not reach other firewall rules.

