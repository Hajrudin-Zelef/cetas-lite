---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-4
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [486, 655]
sha256: cb1e060717b7adefc9dfc935e18f4ea0f52da5c26eab58aebc62e11edb48676f
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Networking in Transparent Mode Packet Forwarding
ARP Table
In Transparent mode, the Address Resolution Protocol (ARP) table is used in the following situations:
l For IP traffic received or originated by the FortiGate itself, and in destination of the management device or next-
hop.
l When IPsec is used, the FortiGate uses its ARP table to forward the traffic from the IPsec tunnel to the local
destination host(s).
All other forwarding decision is based on the Forwarding Database (FDB) table or optional settings.
Verifying the Forwarding Database
To view all instances of the forwarding database (FDB), use the following CLI command:
diagnose netlink brctl list
Example
FGT # diagnose netlink brctl list
list bridge information
1. root.b fdb: size=256 used=6 num=7 depth=2 simple=no
2. mgmt.b fdb: size=256 used=5 num=4 depth=2 simple=no
Total 2 bridges
Here above we can see two bridge instances for 2 VDOMs in Transparent mode: root andmgmt .
l This command will dump the L2 forwarding table for each VDOM bridge instance:
diagnose netlink brctl name host <VDOM_name>.b
Example for the root VDOM:
FGT# diag netlink brctl name host root.b
show bridge control interface root.b host.
fdb: size=256, used=6, num=7, depth=2, simple=no
Bridge root.b host table
port no device devname mac addr ttl atributes
2 7 wan2 02:09:0f:78:69:00 0 Local Static
5 6 trunk_1 02:09:0f:78:69:01 0 Local Static
3 8 dmz 02:09:0f:78:69:01 0 Local Static
4 9 internal 02:09:0f:78:69:02 0 Local Static
3 8 dmz 00:80:c8:39:87:5a 194
4 9 internal 02:09:0f:78:67:68 8
1 3 wan1 00:09:0f:78:69:fe 0 Local Static
Spanning Tree BPDUs Forwarding
Spanning tree Bridge Protocol Data Units (BPDUs) are not forwarded by default in Transparent Mode. To forward
spanning tree BPDUs, the following setting can be applied on each interface where this is required:
config system interface
edit port1
set stpforward enable
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
19

Network Address Translation (NAT) Networking in Transparent Mode
next
end
Non-IPv4 Ethernet frames forwarding
In the situation where non IP frames (or non Ethernet II) frames need to be accepted on a port, the parameter
l2forward can be enabled (disabled by default). This can be used to forward frames such as PPPoE PADI,
Appletalk, on other ports belonging to the same forwarding domain.
The procedure is the following:
config system interface
edit port1
set l2forward enable
next
edit port2
set l2forward enable
next
end
Network Address Translation (NAT)
While NAT is generally not used by a FortiGate in Transparent mode, both source network address translation
(SNAT) and destination network address translation (DNAT) can be configured.
Configuring SNAT
Source Network Address Translation (SNAT) is an option available in Transparent mode and configurable in CLI
only, using the following commands:
config firewall ippool
edit "nat-out"
set endip 192.168.183.48
set startip 192.168.183.48
set interface vlan18_p3
next
end
config firewall policy
edit 3
set srcintf "vlan160_p2"
set dstintf "vlan18_p3"
set srcaddr "all"
set dstaddr "all"
set action accept
set ippool enable
set poolname "nat-out"
set schedule "always"
set service "ANY"
set nat enable
next
end
The sniffer trace below shows the source IP 192.168.182.93 being source translated to 192.168.183.48:
20 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

Networking in Transparent Mode Network Address Translation (NAT)
fgt300 (TP) # diagnose sniffer packet any "host 10.2.2.1" 4
interfaces=[any]
filters=[host 10.2.2.1]
4.891970 vlan160_p2 in 192.168.182.93 -> 10.2.2.1: icmp: echo request
4.892003 vlan18_p3 out 192.168.183.48 -> 10.2.2.1: icmp: echo request
4.892007 port3 out 192.168.183.48 -> 10.2.2.1: icmp: echo request
4.933216 vlan18_p3 in 10.2.2.1 -> 192.168.183.48: icmp: echo reply
4.933249 vlan160_p2 out 10.2.2.1 -> 192.168.182.93: icmp: echo reply
4.933253 port2 out 10.2.2.1 -> 192.168.182.93: icmp: echo reply
Configuring DNAT
The following example shows how to configure Destination Network Address Translation (DNAT) using a virtual
IP on a FortiGatein Transparent Mode:
config firewall vip
edit "vip1"
set extip 192.168.183.48
set extintf "vlan160_p2"
set mappedip 192.168.182.78
next
end
config firewall policy
edit 4
set srcintf "vlan160_p2"
set dstintf "vlan18_p3"
set srcaddr "all"
set dstaddr "vip1"
set action accept
set schedule "always"
set service "ANY"
next
end
If the mappedip is on a different subnet than the management IP, the FortiGate
must have a valid routeto this destination
The sniffer trace below shows the destination IP 192.168.183.48 being translated to 192.168.182.78:
fgt300 (TP) # diagnose sniffer packet any "icmp" 4
interfaces=[any]
filters=[icmp]
4.126138 vlan160_p2 in 192.168.182.93 ->192.168.183.48: icmp: echo request
4.126190 vlan18_p3 out 192.168.182.93 ->192.168.182.78: icmp: echo request
4.126196 port3 out 192.168.182.93 -> 192.168.182.78: icmp: echo request
4.126628 vlan18_p3 in 192.168.182.78 -> 192.168.182.93: icmp: echo reply
4.126661 vlan160_p2 out 192.168.183.48 -> 192.168.182.93: icmp: echo reply
4.126667 port2 out 192.168.183.48 -> 192.168.182.93: icmp: echo reply
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
21

VLANs and Forwarding Domains Networking in Transparent Mode
VLANs and Forwarding Domains
The following sections include information about configuring virtual local area networks (VLANs) and forwarding
domains in Transparent mode:
l VLANs in Transparent Mode
l Forwarding Domains in Transparent Mode
l VLANs vs Forwarding Domains
l VLAN Forwarding
l Unknown VLANs and VLAN Forwarding
l VLAN Trunking and MAC Address Learning
l VLAN Translation
VLANs in Transparent Mode
A VLAN configured on a physical port is used to classify a packet in a broadcast domain in ingress and to tag
packet in egress. A VLAN on the FortiGate conforms to the standard 802.1q. The following rules apply to VLAN
configuration:
l a VLAN ID can be used only once on the same physical port
l the same VLAN ID can be used on a different port
l the VLAN ID range is from 1 to 4094
Forwarding Domains in Transparent Mode
A forwarding domain is used to create separate broadcast domains and confine traffic across two or more ports. It
also allows learning the same MAC in different VLANs (IVL).
A forwarding domain and its associated ID number are unique across one VDOM, or a FortiGate with VDOMs
disabled. Each new VDOM will create a new bridge instance in the FortiGate.
Even though the forwarding domain ID is not in relation with the actual
VLAN numbers, it is recommended, for maintenance and troubleshooting
purposes, to configure one forwarding domain per VLAN and use the same
forwarding domain ID as the VLANs ID.
Once forwarding domains are configured, it is possible to configure firewall
policies only between ports or VLAN belonging to the same forwarding
domain.
22 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

