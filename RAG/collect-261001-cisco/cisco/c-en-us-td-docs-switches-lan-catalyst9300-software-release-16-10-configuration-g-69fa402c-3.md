---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "throughput"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c.md
source_anchor: ""
source_lines: [165, 272]
sha256: a99b9f0dfa647fc3263c90f6f12109cce53a7bf20d41b245c329f1dae5cfa4a2
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c

When the device receives the packets with the inside global IP address, it performs a NAT table lookup; the inside global
address and port, and the outside address and port as keys; translates the addresses to the inside local addresses 10.1.1.1:1723
/ 10.1.1.2:1723 and forwards the packets to host 10.1.1.1. and 10.1.1.2 respectively.
Host 10.1.1.1 and Host 10.1.1.2 receive the packet and continue the conversation. The device performs Steps 2 to 5 for each
packet it receives.
Overlapping Networks
Use NAT to translate IP addresses if the IP addresses that you use are neither legal nor officially assigned. Overlapping
networks result when you assign an IP address to a device on your network that is already legally owned and assigned to a
different device on the Internet or outside the network.
The following figure depicts overlapping networks: the inside network and outside network both have the same local IP addresses
(10.1.1.x). You need network connectivity between such overlapping address spaces with one NAT device to translate the address
of a remote peer (10.1.1.3) to a different address from the perspective of the inside.
Notice that the inside local address (10.1.1.1) and the outside global address ( 10.1.1.3) are in the same subnet. To translate
the overlapping address, first, the inside source address translation happens with the inside local address getting translated
to 203.0.113.2 and a half entry is created in the NAT table. On the Receiving side, the outside source address is translated
to 172.16.0.3 and another half entry is created. The NAT table is then updated with a full entry of the complete translation.
The following steps describe how a device translates overlapping addresses:
Host 10.1.1.1 opens a connection to 172.16.0.3.
The NAT module sets up the translation mapping of the inside local and global addresses to each other and the outside global
and local addresses to each other
The Source Address (SA) is replaced with inside global address and the Destination Address (DA) is replaced with outside
global address.
Host C receives the packet and continues the conversation.
The device does a NAT table lookup, replaces the DA with inside local address, and replaces the SA with outside local address.
Host 10.1.1.1 receives the packet and the conversation continues using this translation process.
Limitations of NAT
There are certain NAT operations that are currently not supported in the Hardware data plane. The following are such operations
that are carried out in the relatively slower Software data plane:
Translation of Internet Control Message Protocol (ICMP) packets.
Translation of packets that require application layer gateway (ALG) processing.
Packets that require both inside and outside translation.
The maximum number of sessions that can be translated and forwarded in the hardware in an ideal setting is limited to 2500.
Additional flows that require translation are handled in the software data plane at a reduced throughput.
Note
Each translation consumes two entries in TCAM.
A configured NAT rule might fail to get programmed into the hardware owing to resource constraint. This could result in packets
that correspond to the given rule to get forwarded without translation.
ALG support is currently limited to FTP, TFTP and ICMP protocols. Also, although TCP SYN, TCP FIN and TCP RST are not part
of ALG traffic, they are processed as part of ALG traffic.
Dynamically created NAT flows age out after a period of inactivity. The number of NAT flows whose activity can be tracked
is limited to 4000.
Port channel is not supported in NAT configuration.
NAT does not support translation of fragmented packets.
Bidirectional Forwarding Detection (BFD) is not supported with NAT configuration.
NAT does not support Stateful Switchover (SSO). Dynamically created NAT states are not synchronized between the Active and
Standby devices.
NAT configuration must be done without using route-maps, as route-mapped NAT is not supported.
Explicit deny access control entry (ACE) in NAT ACL is not supported. Only explicit permit ACE is supported.
The maximum number of TCAM flows that are available in the hardware is 5000.
Note
Using Address Only Translation optimizes the handling of flows and enhances the scale of the NAT feature.
Address Only Translation
Address only Translation (AOT) functionality can be employed in situations that require only the address fields to be translated
and not the transport ports. In such settings, enabling AOT functionality significantly increases the number of flows that
can be translated and forwarded in the hardware at line-rate. This improvement is brought about by optimizing the usage of
various hardware resources associated with translation and forwarding. A typical NAT focused resource allocation scheme sets
aside 5000 TCAM entries for performing hardware translation. This places a strict upper limit on the number of flows that can be translated
and forwarded at line-rate. Under AOT scheme, the usage of TCAM resource is highly optimized thereby enabling the accommodation
of more number of flows in the TCAM tables and this provides a significant improvement in the hardware translation and forwarding
scale. AOT can be very effective in situations where majority of the flows are destined to a single or a small set of destinations.
Under such favourable conditions, AOT can potentially enable line-rate translation and forwarding of all the flows originating
from the given end-point(s). AOT functionality is disabled by default. It can be enabled using the no ip nat create flow-entries command. The existing dynamic flow can be cleared using the clear ip nat translation command. The AOT feature can be disabled using the ip nat create flow-entries command.
Restrictions for Address Only Translation
AOT feature is expected to function correctly only in translation scenarios corresponding to simple inside static and inside
dynamic rules
When AOT is enabled, the show ip nat translation command will not give visibility into all the NAT flows being translated and forwarded.
Configuring NAT
The tasks described in this section will help you configure NAT. Based on the desired configuration, you may need to configure
more than one task.
Configuring Static Translation of Inside Source Addresses
Configure static translation of inside source address to allow one-to-one mapping between an inside local address and an inside
global address. Static translation is useful when a host on the inside must be accessible by a fixed address from the outside.
SUMMARY STEPS
enable
configure terminal
Use any of the following three commands depending on the requirement:
ip nat inside source static local-ip global-ip
Switch(config)# ip nat inside source static 10.10.10.1 172.16.131.
ip nat inside source static protocol local-ip port global-ip port
Establishes static translation between an inside local address and an inside global address.
Establishes a static port translation between an inside local address and an inside global address.
Establishes a static translation between an inside local address and an inside global address. You can specify a range of
subnets to be translated to the inside global address, wherein the host protion of the IP address gets translated and the
network protion of the IP remains the same.
Step 4
interface type number
Example:
Switch(config)# interface ethernet 1
Specifies an interface and enters interface configuration mode.
Step 5
ip addressip-address mask[secondary]
Example:
Switch(config-if)# ip address 10.114.11.39 255.255.255.0
Sets a primary IP address for an interface.
Step 6
ip nat inside
Example:
Switch(config-if)# ip nat inside
Connects the interface to the inside network, which is subject to NAT.
Step 7
exit
Example:
Switch(config-if)# exit
Exits interface configuration mode and returns to global configuration mode.
Step 8
interface type number
Example:
Switch(config)# interface gigabitethernet 0/0/0
