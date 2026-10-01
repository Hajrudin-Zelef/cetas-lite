---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-4
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [267, 351]
sha256: cb36af449e604f3287555433de97bcb18c6d2658b2b0316f2d9ce87b84704d8c
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

Traffic that is replicated for a multicast group and the multicast receivers are on the peer switch.
The known unicast destination MAC address is on the peer switch.
The packet is a MAC notification frame destined for a port on the peer switch.
An SVL also transports system data, such as NetFlow export data and SNMP data, from the Cisco StackWise Virtual standby switch
to the Cisco StackWise Virtual active switch.
Traffic on the SVL is load balanced with the same global hashing algorithms available for EtherChannels (the default algorithm
is source-destination IP).
Interface numbering on Cisco Catalyst 9500X Series Switches operating in SVL mode is 3 tuple for SVL links, DAD links, and
regular data traffic front-panel ports. For breakout interfaces, interface numbering is 4 tuple.
Layer 2 Protocols
The Cisco StackWise Virtual active switch runs the Layer 2 protocols (such as STP and VTP) for the switching modules on both
the switches. Protocol messages that are received on the standby switch ports must traverse SVLs to reach the active switch
where they are processed. Similarly, protocol messages that are transmitted from the standby switch ports originate on the
active switch, and traverse the SVLs to reach the standby ports.
All the Layer 2 protocols in Cisco StackWise Virtual work similarly in standalone mode. The following sections describe the
difference in behavior for some protocols in Cisco StackWise Virtual.
Spanning Tree Protocol
The Cisco StackWise Virtual active switch runs the STP. The Cisco StackWise Virtual standby switch redirects the STP BPDUs
across an SVL to the StackWise Virtual active switch.
The STP bridge ID is commonly derived from the switch MAC address. To ensure that the bridge ID does not change after a switchover,
Cisco StackWise Virtual continues to use the original switch MAC address for the STP Bridge ID.
EtherChannel Control Protocols
Link Aggregation Control Protocol (LACP) and Port Aggregation Protocol (PAgP) packets contain a device identifier. Cisco
StackWise Virtual defines a common device identifier for both the switches. Use either PAgP or LACP on Multi EtherChannels
instead of mode ON, even if all the three modes are supported.
Note
A new PAgP enhancement has been defined for assisting with dual-active scenario detection.
Switched Port Analyzer
Switched Port Analyzer (SPAN) on SVL and fast hello DAD link ports is not supported. These ports can be neither a SPAN source,
nor a SPAN destination. Cisco StackWise Virtual supports all the SPAN features for non-SVL interfaces. The number of SPAN
sessions that are available on Cisco StackWise Virtual matches that on a single switch running in standalone mode.
Private VLANs
Private VLANs on StackWise Virtual work the same way as in standalone mode. The only exception is that the native VLAN on
isolated trunk ports must be configured explicitly.
Apart from STP, EtherChannel Control Protocols, SPAN, and private VLANs, the Dynamic Trunking Protocol (DTP), Cisco Discovery
Protocol (CDP), VLAN Trunk Protocol (VTP), and Unidirectional Link Detection Protocol (UDLD) are the additional Layer 2 control-plane
protocols that run over the SVL connections.
Broadcast, Unknown Unicast and Multicast
Cisco StackWise Virtual supports local switching for Broadcast, Unknown unicast and Multicast (BUM) traffic. In uncommon deployment
scenarios, BUM traffic traverses through the StackWise Virtual Links. This section explains how BUM traffic is handled in
a Cisco StackWise Virtual setup and in local switching.
When a VLAN is created, StackWise Virtual ports are added to the VLAN flood list. The ingress BUM traffic on active or standby
switch traverses through the StackWise Virtual link to the other switch instead of a port in the VLAN. This traffic floods
the StackWise Virtual links which impacts the system and network performance.
To address this, StackWise Virtual BUM optimization feature is introduced.
A general deployment guideline for Cisco StackWise Virtual is to distribute MEC ports evenly at the uplink and downlink as
shown in the figure. In this topology, BUM traffic prefers the local link on MEC to send the traffic out instead of the StackWise
Virtual link. In a scenario where there is a standalone port on a switch or members of EtherChannel on active or standby switch
are down, BUM traffic traverses the StackWise Virtual link. When StackWise Virtual BUM optimization is enabled on VLAN, StackWise
Virtual port is not added to the VLAN flood list. This design ensures BUM traffic does not traverse StackWise Virtual link
only when MEC port channels are part of the VLAN. No optimization is done for VLANs with standalone or physical ports.
Layer 3 Protocols
The Cisco StackWise Virtual active switch runs the Layer 3 protocols and features for the StackWise Virtual. All the Layer
3 protocol packets are sent to and processed by the Cisco StackWise Virtual active switch. Both the member switches perform
hardware forwarding for ingress traffic on their interfaces. When software forwarding is required, packets are sent to the
Cisco StackWise Virtual active switch for processing.
The same router MAC address assigned by the Cisco StackWise Virtual active switch is used for all the Layer 3 interfaces
on both the Cisco StackWise Virtual member switches. After a switchover, the original router MAC address is still used. The
router MAC address is chosen based on chassis-mac and is preserved after switchover by default. Cisco Catalyst 9500 Series High Performance switches support Layer 3 subinterfaces.
The following sections describe the Layer 3 protocols for Cisco StackWise Virtual.
IPv4 Unicast
The CPU on the Cisco StackWise Virtual active switch runs the IPv4 routing protocols and performs any required software forwarding.
All the routing protocol packets received on the Cisco StackWise Virtual standby switch are redirected to the Cisco StackWise
Virtual active switch across the SVL. The Cisco StackWise Virtual active switch generates all the routing protocol packets
to be sent out over ports on either of the Cisco StackWise Virtual member switches.
Hardware forwarding is distributed across both members on Cisco StackWise Virtual. The CPU on the Cisco StackWise Virtual
active switch sends Forwarding Information Base (FIB) updates to the Cisco StackWise Virtual standby switch, which in turn
installs all the routes and adjacencies into hardware.
Packets intended for a local adjacency (reachable by local ports) are forwarded locally on the ingress switch. Packets intended
for a remote adjacency (reachable by remote ports) must traverse the SVL.
The CPU on the Cisco StackWise Virtual active switch performs all software forwarding and feature processing (such as fragmentation
and Time to Live exceed functions). If a switchover occurs, software forwarding is disrupted until the new Cisco StackWise
Virtual active switch obtains the latest Cisco Express Forwarding and other forwarding information.
In virtual switch mode, the requirements to support non-stop forwarding (NSF) match those in the standalone redundant mode
of operation.
From a routing peer perspective, Multi-Chassis EtherChannels (MEC) remain operational during a switchover, that is, only
the links to the failed switch are down, but the routing adjacencies remain valid.
Cisco StackWise Virtual achieves Layer 3 load balancing over all the paths in the Forwarding Information Base entries, be
it local or remote.
IPv6
Cisco StackWise Virtual supports IPv6 unicast and multicast because it is present in the standalone system.
IPv4 Multicast
The IPv4 multicast protocols run on the Cisco StackWise Virtual active switch. Internet Group Management Protocol (IGMP)
and Protocol Independent Multicast (PIM) protocol packets received on the Cisco StackWise Virtual standby switch are transmitted
across an SVL to the StackWise Virtual active switch. The latter generates IGMP and PIM protocol packets to be sent over ports
