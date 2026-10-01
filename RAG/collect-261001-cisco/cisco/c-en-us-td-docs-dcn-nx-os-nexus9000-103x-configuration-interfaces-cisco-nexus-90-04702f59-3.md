---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-3
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [210, 302]
sha256: 5590fb6d9abdaf9687b413b1a91c0537f397bb9c47b30298e103da03781a0dfc
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

as the core router. When two routers carrying GTP traffic are connected with link bundling, the traffic is required to be
distributed evenly between all bundle members.
Different Mechanisms for GTP Load Balancing
Two different kinds of mechanisms are used to achieve GTP load balancing.
From Cisco Nexus Release 10.5(2), the inner IP header fields source, destination IP address and IP protocol is used to maintain
load balancing.
Prior to Cisco Nexus Release 10.5(2), the 5-tuple load balancing mechanism is used. The load balancing mechanism takes into
account the source IP, destination IP, protocol, Layer 4 resource and destination port (if traffic is TCP or UDP) fields from
the packet. In the case of GTP traffic, a limited number of unique values for these fields restrict the equal distribution
of traffic load on the tunnel.
Inner IP Header GTP Load Balancing Mechanism
Using inner IP header fields source-ip, dest-ip and ip-protocol the load-balancing is done. Symmetric load-balancing is supported
to maintain stickiness for forward and reverse traffic of same flow.
GTP inner header based hashing works for both IPv4 and IPv6 on all interfaces. The inner IP header for both IPv4 and IPv6
uses all the 16 UDF for all cloudscale switches. Inner IP headers are used for two switch or three switches bundling.
5-Tuple GTP Load Balancing Mechanism
In order to avoid polarization for GTP traffic in load balancing, a tunnel endpoint identifier (TEID) in the GTP header is
used instead of a UDP port number. Since the TEID is unique per tunnel, traffic can be evenly load balanced across multiple
links in the bundle.
This feature overrides the source and destination port information with the 32-bit TEID value that is present in GTPU packets.
GTP tunnel load balancing feature adds support for:
GTP with IPv4/IPv6 transport header on physical interface
GTP traffic over TE tunnel
GTPU with UDP port 2152
The ip load-sharing address source-destination gtpu command enables the GTP tunnel load balancing.
To know the egress interface for GTP traffic after load balancing, use show cef {ipv4 | ipv6} exact-route command with TEID in place of L4 protocol source and destination port number. Use 16MSBist of TEID in source port and 16LSBits
of TEID in destination port.
The port-channel load-balance src-dst gtpu command enables GTP packets with UDP destination port number 2152 to load balance based on the GTP TEID value. This command
enables the switch to load balance for GTP packets even if the outer five tuples (src-ip, dst-ip, ip proto, L4 sport, L4 dport)
are same. Because the hardware controls for port channel and ECMP are same, enabling either port-channel load-balance or ip load-sharing with GTP option enables GTP TEID based load balancing.
The port-channel load-balance src-dst gtpu command is applicable for both GTP packets, with or without VXLAN encapsulation
When GTP header is a part of the outer layer, the port-channel load-balance src-dst gtpu command picks up GTP TEID from outer layer for hashing.
When GTP header is part of inner layer, the port-channel load-balance src-dst gtpu command picks up GTP TEID from inner layer for hashing.
You need to set the protocol field to 17 and set the value for other parameters when you use the show port-channel load-balance forwarding-path command. An example is listed below.
Beginning Cisco Nexus Release 9.3(3) GTP Tunnel Load Balancing is supported on Cisco Nexus 9500 platform switches with 9700-EX
and 9700-FX line cards. However, GTP Tunnel Load Balancing for IPv6 flow is supported only on Cisco Nexus 9500 platform switches
with FM-E2 fabric modules. It is not supported on Cisco Nexus 9500 platform switches with FM-E fabric modules. Because the
hardware control is same for both Port-channel and ECMP, enabling either port-channel load-balance or ip load-sharing with GTP option enables GTP TEID based load balancing for both the cases. In multi encapsulated packets, if the GTP header
is a part of outer header, it picks up GTP TEIF from outer layer for hashing. If the GTP header is a part of inner header,
it picks up GTP TEIF from inner layer for hashing.
GTP Tunnel Load Balancing is supported on Cisco Nexus 9300-EX, 9300-FX, 9300-FX2, 9364C, and 9300-GX platform switches.
Inner IP header GTP load balancing mechanism is supported on:
Cisco Nexus 9300-EX platform switches
Cisco Nexus 9300-FX and 9364C platform switches
Cisco Nexus 9500 platform switches with 9700-EX and 9700-FX line cards
Cisco Nexus 9300-EX, 9300-FX, 9300-FX2, 9364C, and 9300-GX platform switches
Cisco Nexus 9364C-H1 Switch
Note
Cisco Nexus 9364C-H1 switch can natively support inner-header based hashing for packets with GTP header of size 8 or 12 bytes
LACP
LACP allows you to configure up to 16 interfaces into a port channel.
The Link Aggregation Control Protocol (LACP) for Ethernet is defined in IEEE 802.1AX and IEEE 802.3ad. This protocol controls
how physical ports are bundled together to form one logical channel.
Note
You must enable LACP before you can use LACP. By default, LACP is disabled. See the “Enabling LACP” section for information
about enabling LACP.
The following figure shows how individual links can be combined into LACP port channels and channel groups as well as function
as individual links.
Figure 2. Individual Links Combined into a Port Channel
With LACP, you can bundle up to 32 interfaces in a channel group.
Note
When you delete the port channel, the software automatically deletes the associated channel group. All member interfaces revert
to their original configuration.
Note
If you downgrade a Cisco Nexus 9500 series switch that is configured to use LACP vPC convergence feature, that runs Cisco
NX-OS Release 7.0(3)I7(5) to a lower release, the configuration is removed. You must configure the LACP vPC convergence feature
again when you upgrade the switch.
You cannot disable LACP while any LACP configurations are present.
Port-Channel Modes
Individual interfaces in port channels are configured with channel modes. When you run static port channels with no aggregation
protocol, the channel mode is always set to on. After you enable LACP globally on the device, you enable LACP for each channel by setting the channel mode for each interface
to either active or passive. You can configure channel mode for individual links in the LACP channel group when you are adding the links to the channel
group
Note
You must enable LACP globally before you can configure an interface in either the active or passive channel mode.
The following table describes the channel modes.
Table 1. Channel Modes for Individual Links in a Port Channel
Channel Mode
Description
passive
The LACP is enabled on this port channel and the ports are in a passive negotiating state. Ports responds to LACP packets
that it receives but does not initiate LACP negotiation.
active
The LACP is enabled on this port channel and the ports are in an active negotiating state. Ports initiate negotiations with
other ports by sending LACP packets.
on
The LACP is disabled on this port channel and the ports are in a non-negotiating state. The on state of the port channel represents the static mode.
The port will not verify or negotiate port channel memberships. If you attempt to change the channel mode to active or passive
before enabling LACP, the device displays an error message. When an LACP attempts to negotiate with an interface in the on state, it does not receive any LACP packets and becomes an individual link with that interface, it does not join the LACP
channel group. The on state is the default port-channel mode
Both the passive and active modes allow LACP to negotiate between ports to determine if they can form a port channel based
on criteria such as the port speed and the trunking state.The passive mode is useful when you do not know whether the remote
system, or partner, supports LACP.
