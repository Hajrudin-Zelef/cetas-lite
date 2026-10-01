---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-2
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet", "license"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [103, 181]
sha256: 489f239faf14cc144c63d25232aa3e9a47ee1035a71d5c131d30876fb07b1066
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

logical virtual switch using an Ethernet connection.
Cisco StackWise Virtual Topology
A typical network design consists of core, distribution, and access layers. The default mode of a switch is standalone. When
two redundant switches are deployed in the distribution layer, the following network challenges arise:
If VLAN IDs are reused between access layers then, it will introduce a spanning tree loop that will impact the overall performance
of the network.
Spanning tree protocols and configuration are required to protect Layer 2 network against spanning tree protocol loop, and
root and bridge protocol data unit management.
Additional protocols such as first hop redundancy protocol are required to virtualize the IP gateway function. This should
align with STP root priorities for each VLAN.
The Protocol independent multicast designated router (PIM DR) configuration should be fine-tuned to selectively build a multicast
forwarding topology on a VLAN.
The standalone distribution layer system provides protocol-driven remote failure and detection, which results in slower convergence
time. Fine-tune First Hop Redundancy Protocol (FHRP) and PIM timers for rapid fault detection and recovery process.
We recommend Cisco StackWise Virtual model for aggregation layers and collapsed aggregation and core layers. The stack can
be formed over a 100G, 10G, 25G, 40G, and 100G link to ensure that the distribution or the aggregation switches can be deployed over a large distance.
Additionally, on the Cisco Catalyst 9500X Series Switches, the stack can be formed over a 400-G link.
Note that STP keeps one of the ports connected to the distribution switches blocked on the access switches. As a result of
this, an active link failure causes STP convergence and the network suffers from traffic loss, flooding, and a possible transient
loop in the network. On the other hand, if the switches are logically merged into one switch, all the access switches might
form an EtherChannel bundle with distribution switches, and a link failure within an EtherChannel would not have any impact
as long as at least one member within the EtherChannel is active.
Etherchannel in StackWise Virtual is capable of implementing Multi-chassis EtherChannel (MEC) across the stack members. When
access layer and aggregation layer are collapsed into a single StackWise Virtual system, MEC across the different access layer
domain members and across distribution and access layer switches will not be supported. MEC is designed to forward the traffic
over the local link irrespective of the hash result.
Since the control plane, management plane, and data plane are integrated, the system behaves as a single switch.
The virtualization of multiple physical switches into a single logical switch is from a control and management plane perspective
only. Because of the control plane being common, it may look like a single logical entity to peer switches. The data plane
of the switches is distributed. Each switch is capable of forwarding over its local interfaces without involving other members.
However, when a packet coming into a switch has to be forwarded over a different member’s port, the forwarding context of
the packet is carried over to the destination switch after ingress processing is performed in the ingress switch. Egress processing
is done only in the egress switch. This provides a uniform data plane behavior to the entire switch irrespective whether of
the destination port is in a local switch or in a remote switch. However, the common control plane ensures that all the switches
have equivalent data plane entry for each forwarding entity.
An election mechanism elects one of the switches to be Cisco StackWise Virtual active and the other switch to be Cisco StackWise
Virtual standby in terms of Control Plane functions. The active switch is responsible for all the management, bridging and
routing protocols, and software data path. The standby switch is in hot standby state ready to take over the role of active,
if the active switch fails over.
The following are the components of the Cisco StackWise Virtual solution:
Stack members
SVL: 400G, 100G, 50G, 10G, 25G, 40G, and 100G Ethernet connections. SVL is established using the 400G, 100G, 50G, 10G, 25G, 40G, and 100G interfaces depending on the switch models. However, a combination of two different speeds is not supported.
SVL is the link that connects the switches over Ethernet. Typically, Cisco StackWise Virtual consists of multiple 400G, 100G, 50G, 10G, 25G, 40G, and 100G physical links. It carries all the control and data traffic between the switching units. You can configure
SVL on a supported port. When a switch is powered up and the hardware is initialized, it looks for a configured SVL before
the initialization of the control plane.
The Link Management Protocol (LMP) is activated on each link of the SVL as soon as the links are established. LMP ensure the
integrity of the links and monitors and maintains the health of the links. The redundancy role of each switch is resolved
by the StackWise Discovery Protocol (SDP). It ensures that the hardware and software versions are compatible to form the SVL
and determines which switch becomes active or standby from a control plane perspective.
Note
On the Cisco Catalyst 9500X Series Switches, Link Aggregation Control Protocol (LACP) replaces LMP, and Intermediate System to Intermediate System (ISIS) replaces SDP.
Cisco StackWise Virtual Header (SVH) is 64-byte frame header that is prepended over all control, data, and management plane
traffic that traverse over each SVL between the two stack members of the Cisco StackWise Virtual domain. The SVH-encapsulated
traffic operates at OSI Layer 2 and can be recognized and processed only by Cisco StackWise Virtual-enabled switches. SVL
interfaces are non-bridgeable and non-routeable, and allows non-routeable traffic over L2 or L3 network.
Cisco StackWise Virtual Redundancy
Cisco StackWise Virtual operates stateful switchover (SSO) between the active and standby switches. The following are the
ways in which Cisco StackWise Virtual's redundancy model differs from that of the standalone mode:
The Cisco StackWise Virtual active and standby switches are hosted in separate switches and use a StackWise Virtual link to
exchange information.
The active switch controls both the switches of Cisco StackWise Virtual. The active switch runs the Layer 2 and Layer 3 control
protocols and manages the switching modules of both the switches.
The Cisco StackWise Virtual active and standby switches perform data traffic forwarding.
Note
If the Cisco StackWise Virtual active switch fails, the standby switch initiates a switchover and assumes the Cisco StackWise
Virtual active switch role.
SSO Redundancy
A StackWise Virtual system operates with SSO redundancy if it meets the following requirements:
Both the switches must be running the same software version, unless they are in the process of software upgrade.
SVL-related configuration in the two switches must match.
License type must be same on both the switch models.
Both the switch models must be in the same StackWise Virtual domain.
With SSO redundancy, the StackWise Virtual standby switch is always ready to assume control if a fault occurs on the StackWise
Virtual active switch. Configuration, forwarding, and state information are synchronized from the StackWise Virtual active
switch to the redundant switch at startup, and whenever changes to the StackWise Virtual active switch configuration occur.
If a switchover occurs, traffic disruption is minimized.
If StackWise Virtual does not meet the requirements for SSO redundancy, it will be incapable of establishing a relationship
with the peer switch. StackWise Virtual runs stateful switchover (SSO) between the StackWise Virtual active and standby switches.
The StackWise Virtual determines the role of each switch during initialization.
