---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-3
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [327, 485]
sha256: 8578f64232dfb3c853f62c1e047a2bb2bee75eb5dfe7fad077759e22808e2bf2
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Networking in Transparent Mode
This section contains information about networking concepts in Transparent mode. It contains the following
topics:
l Packet Forwarding
l Network Address Translation (NAT)
l VLANs and Forwarding Domains
l Inter-VDOM links between NAT/Route and Transparent VDOMs
l Packet Forwarding using Cisco Protocols
l Configuration Example
Packet Forwarding
The following sections include information aboutconfiguring packet forwarding in Transparent mode:
l MAC learning and L2 Forwarding Table
l Broadcast, Multicast, and Unicast Forwarding
l Multicast Processing
l Source MAC Addresses
l ARP Table
l Verifying the Forwarding Database
l Spanning Tree BPDUs Forwarding
l Non-IPv4 Ethernet frames forwarding
MAC learning and L2 Forwarding Table
When operating in Transparent mode, a FortiGate behaves like an L2 switch in accordance with 802.1d
principles:
l The forwarding database (FDB) is populated with the network devices MAC addresses during a MAC learning
process, based on the source addresses seen in the Ethernet frames ingressing a FortiGate port. Static MAC
entries can also be configured using the following CLI command:
config system mac-address-table
edit 00:01:02:03:04:05
set interface "port3"
next
end
The FDB table can be verified with the following command: diagnose netlink
brctl name host TP.b
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
15

Packet Forwarding Networking in Transparent Mode
l Ethernet IP frames forwarding is based on known MAC address on each port.
l As Spanning Tree in not running on the FortiGate, a port that comes up goes immediately into forwarding or
flooding state. This last state will not occur once unicast MAC addresses are present in the FDB.
If the FortiGate in Transparent mode bridges traffic to a router or host using a virtual
MAC for one direction and a different physical MAC for the other direction (for example
when VRRP or HSRP protocols are used), it is highly recommended to create a static
MAC entry for the virtual MAC. This is to make sure that the virtual MAC address is
present in the FDB.
Broadcast, Multicast, and Unicast Forwarding
In Transparent mode, IPv4 packets are typically only forwarded by the FortiGate from a port to another port when
a firewall policy is matched with action ACCEPT.
Below are exceptions.
l L2 (IP) Broadcast frames forwarding:
L2 (IP) means a L2 frame type 0x0800 (IP) or 0x0806 (ARP)
l ARP: by default, ARP broadcasts and ARP reply packets are flooded/forwarded on all ports or VLANs
belonging to the same forwarding domain, without the need of firewall policies between the ports. This default
behavior is necessary to allow the population of the FDB and allow further firewall policy lookup (see section
Transparent mode Firewall processing for more details). This option is configurable at the interface settings
level with the parameter arpforward (enabled by default).
l Non-ARP: To forward non-ARP broadcasts, the following CLI command is used:
config system interface
edit "port2"
set broadcast-forward enable
next
end
l L2 (IP) Multicast frames forwarding: the FortiGate does not forward frames with multicast destination MAC
addresses by default. Multicast traffic such as one used by routing protocols or streaming media may need to
traverse the FortiGate which should not interfere this communication.
Fortinet recommends that the FortiGate is set up using Multicast policies. This allows for greater control and
predictability on traffic behavior. However Multicast traffic may be forwarded through a Transparent mode device
using the multicast-skip-policy setting. This is detailed in the section "Multicast Processing" on page 17
l L2 (IP) Unicast frames forwarding: a frame with a unicast destination MAC address is subject to firewall
processing before being forwarded (see "Firewall Policy Look Up" on page 29 for more details). This does not apply
to ARP replies.
16 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

Networking in Transparent Mode Packet Forwarding
Multicast Processing
In Transparent mode, a FortiGate does not forward frames with multicast destination MAC addresses by default.
If multicast traffic is required, multicast policies are recommended to allow finer control of this traffic.
Forwarding all multicast traffic with policy
Multicast traffic may have to be forwarded through a Transparent mode device using the multicast-skip-
policy sytem setting. This is the configuration for this solution:
config system settings
set multicast-skip-policy enable
end
In that case, no check is performed on sources/destinations/interfaces. A multicast packet received on an
interface is flooded unconditionally to all interfaces (except the incoming interface) belonging to the same
forwarding domain.
Configuring firewall multicast-policy
The use of firewall multicast-policy allows a finer control over the multicast packets. Hereafter are
some commented examples. Note that the parameter multicast-skip-policy mentioned above must be left to
disabled.
Those policies can only be configured from the CLI.
1- Simple policy
config firewall multicast-policy
edit 1
set action accept
next
end
In that case, no check is performed on sources/destinations/interfaces. A multicast packet received on an
interface is flooded unconditionally to all interfaces (except the incoming interface) belonging to the same
forwarding domain.
2- To restrict incoming and outgoing interfaces:
config firewall multicast-policy
edit 1
set srcintf "port1"
set dstintf "port2"
set action accept
next
end
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
17

Packet Forwarding Networking in Transparent Mode
3- To be more restrictive (example to allow RIP2 packets from port1 to port2 and sourced by 10.10.0.10):
config firewall multicast-policy
edit 1
set srcintf "port1"
set srcaddr 10.10.0.10 255.255.255.255
set dstintf "port2"
set dstaddr 224.0.0.9 255.255.255.255
set action accept
next
end
4- This policy will allow all 224.0.0.0/255 range (OSPF, RIPv2, DVMPR…) from port1 to port2
config firewall multicast-policy
edit 1
set srcintf "port1"
set dstintf "port2"
set dstaddr 224.0.0.0 255.255.255
set action accept
next
end
Source MAC Addresses
When a FortiGate is in Transparent Mode, it does not typically alter the original source and destination address of
packets that flow through the unit. Because of this, end devices do not “see” the MAC address of the FortiGate.
However, if network address translation (NAT) is enabled by a firewall policy, the source MAC address will be the
MAC address of the FortiGate's management interface.
IP packets that are initiated by the FortiGate (remote management, access to FortiGuard server…) are sent in L2
Ethernet frames that have a source MAC address of the interface in the virtual domain (VDOM) with the lowest
MAC address. Below is an example with port2 and port3 in the same VDOM, remote access done via port2, but
the sniffer trace showing MAC address of port2. The address of port2 is shown in bold.
diagnose hardware deviceinfo nic port2
[…]
Current_HWaddr          00:09:0F:85:3F:C4
Permanent_HWaddr        00:09:0F:85:3F:C4
fgt300 (global) # diagnose hardware deviceinfo nic port3
[…]
Current_HWaddr          00:09:0F:85:3F:C5
Permanent_HWaddr        00:09:0F:85:3F:C5
diagnose sniffer packet port3 "port 80" 6
3.774236 port3 -- 192.168.171.165.2619 -> 192.168.182.136.80: syn 3961770249
0x0000   0009 0f85 3fc4 0009 0f09 3204 0800 4500      ....? .... 2...E.
0x0010   0030 8071 4000 7e06 98d7 c0a8 aba5 c0a8      .0.q@.~ ........
0x0020   b688 0a3b 0050 ec23 d109 0000 0000 7002      ...;.P.# ..... p.
0x0030   ffff d7e7 0000 0204 05b4 0101 0402
18 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

