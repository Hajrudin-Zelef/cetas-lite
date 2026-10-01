---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9.md
source_anchor: ""
source_lines: [1, 68]
sha256: 7ea28f93e4969dda965d1f0b3907c009a8e75c9245e1187d3ebe8551c3774d43
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-16-configuration-g-9862f4f9

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
All ports in an EtherChannel must be assigned to the same VLAN or they must be configured as trunk port.
The LACP 1:1 redundancy feature is supported on port channel interfaces only.
On
EtherChannel member port selection is software based in Layer 2 and Layer 3 multicast route. This means that all multicast
traffic under a group will be routed via the same physical port of the EtherChannel. As a result, the distribution of multicast
traffic load balance over etherchannels might not evenly spread across member ports.
Information About EtherChannels
The following sections provide information about EtherChannels and the various modes to configure EtherChannels.
EtherChannel provides fault-tolerant high-speed links between switches, routers, and servers. You can use the EtherChannel
to increase the bandwidth between the wiring closets and the data center, and you can deploy it anywhere in the network where
bottlenecks are likely to occur. EtherChannel provides automatic recovery for the loss of a link by redistributing the load
across the remaining links. If a link fails, EtherChannel redirects traffic from the failed link to the remaining links in
the channel without intervention.
An EtherChannel consists of individual Ethernet links bundled into a single logical link.
Figure 1. Typical EtherChannel Configuration
Each EtherChannel can consist of up to eight compatibly configured Ethernet ports.
Channel Groups and Port-Channel Interfaces
An EtherChannel comprises a channel group and a port-channel interface. The channel group binds physical ports to the port-channel
interface. Configuration changes applied to the port-channel interface apply to all the physical ports bound together in the
channel group.
Figure 2. Relationship Between Physical Ports, a Channel Group, and a Port-Channel Interface
The channel-group command binds the physical port and the port-channel interface together. Each EtherChannel has a port-channel logical interface
numbered from 1 to 128. This port-channel interface number corresponds to the one specified with the channel-group interface configuration command.
With Layer 2 ports, use the channel-group interface configuration command to dynamically create the port-channel interface.
You also can use the interface port-channelport-channel-number global configuration command to manually create the port-channel interface, but then you must use the channel-groupchannel-group-numbercommand to bind the logical interface to a physical port. The channel-group-number can be the same as the port-channel-number, or you can use a new number. If you use a new number, the channel-group command dynamically creates a new port channel.
With Layer 3 ports, you should manually create the logical interface by using the interface port-channel global configuration command followed by the no switchport interface configuration command. You then manually assign an interface to the EtherChannel by using the channel-group interface configuration command.
With Layer 3 ports, use the no switchport interface command to configure the interface as a Layer 3 interface, and then use the channel-group interface configuration command to dynamically create the port-channel interface.
Port Aggregation Protocol
The Port Aggregation Protocol (PAgP) is a Cisco-proprietary protocol that can be run only on Cisco devices and on those devices
that are licensed by vendors to support PAgP. PAgP facilitates the automatic creation of EtherChannels by exchanging PAgP
packets between Ethernet ports. PAgP can be enabled on cross-stack EtherChannels.
By using PAgP, the switch or switch stack learns the identity of partners capable of supporting PAgP and the capabilities
of each port. It then dynamically groups similarly configured ports (on a single device in the stack) into a single logical
link (channel or aggregate port). Similarly configured ports are grouped based on hardware, administrative, and port parameter
constraints. For example, PAgP groups the ports with the same speed, duplex mode, native VLAN, VLAN range, and trunking status
and type. After grouping the links into an EtherChannel, PAgP adds the group to the spanning tree as a single device port.
PAgP modes specify whether a port can send PAgP packets, which start PAgP negotiations, or only respond to PAgP packets received.
Table 1. EtherChannel PAgP Modes
Mode
Description
auto
Places a port into a passive negotiating state, in which the port responds to PAgP packets it receives but does not start
PAgP packet negotiation. This setting minimizes the transmission of PAgP packets.
desirable
Places a port into an active negotiating state, in which the port starts negotiations with other ports by sending PAgP packets.
Switch ports exchange PAgP packets only with partner ports that are configured in the auto or desirable modes. Ports that are configured in the on mode do not exchange PAgP packets.
Both the auto and desirable modes enable ports to negotiate with partner ports to form an EtherChannel based on criteria such as port speed. and for
Layer 2 EtherChannels, based on trunk state and VLAN numbers.
Ports can form an EtherChannel when they are in different PAgP modes as long as the modes are compatible. For example:
A port in the desirable mode can form an EtherChannel with another port that is in the desirable or auto mode.
A port in the auto mode can form an EtherChannel with another port in the desirable mode.
A port in the auto mode cannot form an EtherChannel with another port that is also in the auto mode because neither port starts PAgP negotiation.
If your switch is connected to a partner that is PAgP-capable, you can configure the switch port for nonsilent operation by
using the non-silent keyword. If you do not specify non-silent with the auto or desirable mode, silent mode is assumed.
Use the silent mode when the switch is connected to a device that is not PAgP-capable and seldom, if ever, sends packets.
An example of a silent partner is a file server or a packet analyzer that is not generating traffic. In this case, running
PAgP on a physical port that is connected to a silent partner prevents that switch port from ever becoming operational. However,
the silent setting allows PAgP to operate, to attach the port to a channel group, and to use the port for transmission.
Port Aggregation Protocol Learn Method and Priority
Network devices are classified as PAgP physical learners or aggregate-port learners. A device is a physical learner if it
learns addresses by physical ports and directs transmissions based on that knowledge. A device is an aggregate-port learner
if it learns addresses by aggregate (logical) ports. The learn method must be configured the same at both ends of the link.
When a device and its partner are both aggregate-port learners, they learn the address on the logical port-channel. The device
sends packets to the source by using any of the ports in the EtherChannel. With aggregate-port learning, it is not important
on which physical port the packet arrives.
PAgP cannot automatically detect when the partner device is a physical learner and when the local device is an aggregate-port
learner. Therefore, you must manually set the learning method on the local device to learn addresses by physical ports. You
