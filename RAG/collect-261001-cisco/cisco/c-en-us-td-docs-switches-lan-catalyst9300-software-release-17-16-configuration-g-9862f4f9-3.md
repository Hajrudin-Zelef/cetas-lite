---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [145, 219]
sha256: 815731abd872a889a8146f4f3107572adb3a0a5b94585169c140ea6be6dff494
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

LACP sends and receives LACP PDUs only from ports that are up and have LACP enabled for the active or passive mode.
Link Aggregation Control Protocol Interaction with Other Features 1:1 Redundancy
The LACP 1:1 Redundancy feature supports an EtherChannel configuration with one active link, and fast switchover to a hot-standby
link. The link that is connected to the port with the lower port priority number (and therefore, of a higher priority) will
be the active link, and the other link will be in a hot-standby state. If the active link goes down, LACP performs a fast
switchover to the hot-standby link to keep the EtherChannel up. When the failed link becomes operational again, LACP performs
another fast switchover to revert to the original active link.
To allow the higher priority port to stabilize when it becomes active again after a higher-priority to lower-priority switchover,
the LACP 1:1 Hot Standby Dampening feature configures a timer that delays switchover back to the higher priority port after
higher priority port becomes active.
EtherChannel On Mode
EtherChannel on mode can be used to manually configure an EtherChannel. The on mode forces a port to join an EtherChannel without negotiations. The on mode can be useful if the remote device does not support PAgP or LACP. In the on mode, a usable EtherChannel exists only when the devices at both ends of the link are configured in the on mode.
Ports that are configured in the on mode in the same channel group must have compatible port characteristics, such as speed and duplex. Ports that are not compatible
are suspended, even though they are configured in the on mode.
Caution
You should use care when using the on mode. This is a manual configuration, and ports on both ends of the EtherChannel must have the same configuration. If the
group is misconfigured, packet loss or spanning-tree loops can occur.
Load-Balancing and Forwarding Methods
EtherChannel balances the traffic load across the links in a channel by reducing part of the binary pattern that is formed
from the addresses in the frame to a numerical value that selects one of the links in the channel. You can specify one of
several different load-balancing modes, including load distribution based on MAC addresses, IP addresses, source addresses,
destination addresses, or both source and destination addresses. The selected mode applies to all EtherChannels configured
on the device.
Note
Layer 3 Equal-cost multi path (ECMP) load balancing is based on source IP address, destination IP address, source port, destination
port, and layer 4 protocol. Fragmented packets will be treated on two different links based on the algorithm that is calculated
using these parameters. Any changes in one of these parameters result in load balancing.
With source-MAC address forwarding, when packets are forwarded to an EtherChannel, they are distributed across the ports in
the channel based on the source-MAC address of the incoming packet. Therefore, to provide load-balancing, packets from different
hosts use different ports in the channel, but packets from the same host use the same port in the channel.
With destination-MAC address forwarding, when packets are forwarded to an EtherChannel, they are distributed across the ports
in the channel based on the destination host’s MAC address of the incoming packet. Therefore, packets to the same destination
are forwarded over the same port, and packets to a different destination are sent on a different port in the channel.
With source-and-destination MAC address forwarding, when packets are forwarded to an EtherChannel, they are distributed across
the ports in the channel based on both the source and destination MAC addresses. This forwarding method, a combination source-MAC
and destination-MAC address forwarding methods of load distribution, can be used if it is not clear whether source-MAC or
destination-MAC address forwarding is better suited on a particular device. With source-and-destination MAC-address forwarding,
packets sent from host A to host B, host A to host C, and host C to host B could all use different ports in the channel.
IP Address Forwarding
With source-IP address-based forwarding, packets are distributed across the ports in the EtherChannel based on the source-IP
address of the incoming packet. To provide load balancing, packets from different IP addresses use different ports in the
channel, and packets from the same IP address use the same port in the channel.
With destination-IP address-based forwarding, packets are distributed across the ports in the EtherChannel based on the destination-IP
address of the incoming packet. To provide load balancing, packets from the same IP source address that is sent to different
IP destination addresses could be sent on different ports in the channel. Packets sent from different source IP addresses
to the same destination IP address are always sent on the same port in the channel.
With source-and-destination IP address-based forwarding, packets are distributed across the ports in the EtherChannel based
on both the source and destination IP addresses of the incoming packet. This forwarding method, a combination of source-IP
and destination-IP address-based forwarding, can be used if it is not clear whether source-IP or destination-IP address-based
forwarding is better suited on a particular device. In this method, packets sent from the IP address A to IP address B, from
IP address A to IP address C, and from IP address C to IP address B could all use different ports in the channel.
Load-Balancing Advantages
Different load-balancing methods have different advantages, and the choice of a particular load-balancing method should be
based on the position of the device in the network and the kind of traffic that needs to be load-distributed.
Figure 3. Load Distribution and Forwarding Methods. In the following figure, an EtherChannel of four workstations communicates with a router. Because the router is a single
MAC-address device, source-based forwarding on the switch EtherChannel ensures that the switch uses all available bandwidth
to the router. The router is configured for destination-based forwarding because the large number of workstations ensures
that the traffic is evenly distributed from the router EtherChannel.
Use the option that provides the greatest variety in your configuration. For example, if the traffic on a channel is going
only to a single MAC address, using the destination-MAC address always chooses the same link in the channel. Using source
addresses or IP addresses might result in better load-balancing.
EtherChannel and Switch Stacks
If a stack member that has ports participating in an EtherChannel fails or leaves the stack, the active switch removes the
failed stack member switch ports from the EtherChannel. The remaining ports of the EtherChannel, if any, continue to provide
connectivity.
When a switch is added to an existing stack, the new switch receives the running configuration from the active switch and
updates itself with the EtherChannel-related stack configuration. The stack member also receives the operational information
(the list of ports that are up and are members of a channel).
When two stacks merge that have EtherChannels configured between them, self-looped ports result. Spanning tree detects this
condition and acts accordingly. Any PAgP or LACP configuration on a winning switch stack is not affected, but the PAgP or
LACP configuration on the losing switch stack is lost after the stack reboots.
With PAgP, if the active switch fails or leaves the stack, the standby switch becomes the new active switch. The new active
switch synchronizes the configuration of the stack members to that of the active switch. The PAgP configuration is not affected
after an active switch change unless the EtherChannel has ports residing on the old active switch.
Switch Stacks and Link Aggregation Control Protocol
