---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59-2
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59.md
source_anchor: ""
source_lines: [103, 209]
sha256: 6acd45c1f185067824fbbb0b13c93044e41a362ddc88176559c90179c985b7bb
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-04702f59

Flow-control configuration
When the interface joins a port channel, some of its individual parameters are removed and replaced with the values on the
port channel as follows:
Bandwidth
Delay
Extended Authentication Protocol over UDP
VRF
IP address
MAC address
Spanning Tree Protocol
NAC
Service policy
Access control lists (ACLs)
Many interface parameters remain unaffected when the interface joins or leaves a port channel as follows:
Beacon
Description
CDP
LACP port priority
Debounce
UDLD
MDIX
Rate mode
Shutdown
SNMP trap
Note
When you delete the port channel, the software sets all member interfaces as if they were removed from the port channel.
See the “LACP Marker Responders” section for information about port-channel modes.
Load Balancing Using Port Channels
The Cisco NX-OS software load balances traffic across all operational interfaces in a port channel by hashing the addresses
in the frame to a numerical value that selects one of the links in the channel. Port channels provide load balancing by default.
Port-channel load balancing uses MAC addresses, IP addresses, or Layer 4 port numbers to select the link. Port-channel load
balancing uses either source or destination addresses or ports, or both source and destination addresses or ports.
You can configure the load- balancing mode to apply to all port channels that are configured on the entire device. You can
configure one load-balancing mode for the entire device. You cannot configure the load-balancing method per port channel.
You can configure the type of load-balancing algorithm used. You can choose the load-balancing algorithm that determines which
member port to select for egress traffic by looking at the fields in the frame.
The default load-balancing mode for Layer 3 interfaces is the source and destination IP L4 ports, and the default load-balancing
mode for non-IP traffic is the source and destination MAC address. Use the port-channel load-balance command to set the load-balancing method among the interfaces in the channel-group bundle. The default method for Layer 2
packets is src-dst-mac. The default method for Layer 3 packets is src-dst ip-l4port.
You can configure the device to use one of the following methods to load balance across the port channel:
Destination MAC address
Source MAC address
Source and destination MAC address
Destination IP address
Source IP address
Source and destination IP address
Source TCP/UDP port number
Destination TCP/UDP port number
Source and destination TCP/UDP port number
GRE inner IP headers with source, destination and source-destination
Non-IP and Layer 3 port channels both follow the configured load-balancing method, using the source, destination, or source
and destination parameters. For example, when you configure load balancing to use the source IP address, all non-IP traffic
uses the source MAC address to load balance the traffic while the Layer 3 traffic load balances the traffic using the source
IP address. Similarly, when you configure the destination MAC address as the load-balancing method, all Layer 3 traffic uses
the destination IP address while the non-IP traffic load balances using the destination MAC address.
The unicast and multicast traffic is load-balanced across port-channel links based on configured load-balancing algorithm
displayed in show port-channel load-balancing command output.
The multicast traffic uses the following methods for load balancing with port channels:
Multicast traffic with Layer 4 information—Source IP address, source port, destination IP address, destination port
Multicast traffic without Layer 4 information—Source IP address, destination IP address
Non-IP multicast traffic—Source MAC address, destination MAC address
Note
Devices that run Cisco IOS can optimize the behavior of the member ports ASICs if a failure of a single member occurred by
running the port-channel hash-distribution command. The Cisco Nexus 9000 Series device performs this optimization by default
and does not require or support this command. Cisco NX-OS does support the customization of the load-balancing criteria on
port channels through the port-channel load-balance command for the entire device.
In a multi-tier Layer 2 or Layer 3 network, polarization can occur. To prevent this, modify the load-balancing algorithm.
One way is to set a different rotate bit for each tier in the network.
Symmetric Hashing
To be able to effectively monitor traffic on a port channel, it is essential that each interface connected to a port channel
receives both forward and reverse traffic flows. Normally, there is no guarantee that the forward and reverse traffic flows
will use the same physical interface. However, when you enable symmetric hashing on the port channel, bidirectional traffic
is forced to use the same physical interface and each physical interface in the port channel is effectively mapped to a set
of flows.
When symmetric hashing is enabled, the parameters used for hashing, such as the source and destination IP address, are normalized
before they are entered into the hashing algorithm. This process ensures that when the parameters are reversed (the source
on the forward traffic becomes the destination on the reverse traffic), the hash output is the same. Therefore, the same interface
is chosen.
Only the following load-balancing algorithms support symmetric hashing:
src-dst ip
src-dst ip-l4port
Guidelines and Limitations for ECMP
You might observe that load balancing with Layer 2/Layer 3 GW flows are not load balanced equally among all links when the
switch comes up initially after reload. There are two CLIs to change the ECMP hash configuration in the hardware. The two
CLI commands are mutually exclusive.
Enter the port-channel load-balance [src | src-dst | dst] mac command for MAC-based only hash.
For hash based on IP/Layer 4 ports, enter either the ip load-share or port-channel load-balance command.
The port-channel load-balance command can overwrite the ip load-share command. It is better to enter the port-channel load-balance command which helps to set both the IP and MAC parameters.
There are no options to force the hashing algorithm based on the IP/Layer 4 port. The default MAC configuration is always
programmed as a part of the port channel configuration.
ECMP resilient hashing is not supported for traffic flows over tunnel.
Resilient Hashing
With the exponential increase in the number of physical links used in data centers, there is also the potential for an increase
in the number of failed physical links. In static hashing systems that are used for load balancing flows across members of
port channels or Equal Cost Multipath (ECMP) groups, each flow is hashed to a link. If a link fails, all flows are rehashed
across the remaining working links. This rehashing of flows to links results in some packets being delivered out of order
even for those flows that were not hashed to the failed link.
This rehashing also occurs when a link is added to the port channel or Equal Cost Multipath (ECMP) group. All flows are rehashed
across the new number of links, which results in some packets being delivered out of order.
Resilient hashing maps flows to physical ports and it is supported for both ECMP groups and port channel interfaces.
If a physical link fails, the flows originally assigned to the failed link are redistributed uniformly among the remaining
working links. The existing flows through the working links are not rehashed and hence are not impacted.
Resilient hashing supports IPv4 and IPv6 unicast traffic, but it does not support IPv4 multicast traffic.
Resilient hashing is supported on all the Cisco Nexus 9000 Series platforms . Beginning Cisco NX-OS Release 9.3(3), resilient
hashing is supported on Cisco Nexus 92160YC-X, 92304QC, 9272Q, 9232C, 9236C, 92300YC switches.
GTP Tunnel Load Balancing
GPRS Tunneling Protocol (GTP) is used mainly to deliver mobile data on wireless networks via Cisco Nexus 9000 Series switches
