---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066-3
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066.md
source_anchor: ""
source_lines: [182, 266]
sha256: 5a10393db920874a1ac9d696ed725dcfb5c16b042621f783d246b3d974320a96
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-13-configuration-g-1be52066

The CPU in the StackWise Virtual standby switch runs in hot standby state. StackWise Virtual uses SVL to synchronize configuration
data from the StackWise Virtual active switch to the StackWise Virtual standby switch. Also, protocols and features that support
high availability synchronize their events and state information to the StackWise Virtual standby switch.
Nonstop Forwarding
While implementing Nonstop Forwarding (NSF) technology in systems using SSO redundancy mode, network disruptions are minimized
for campus users and applications. High availability is provided even when the control-plane processing stack-member switch
is reset. During a failure of the underlying Layer 3, NSF-capable protocols perform graceful network topology resynchronization.
The preset forwarding information on the redundant stack-member switch remains intact; this switch continues to forward the
data in the network. This service availability significantly lowers the mean time to repair (MTTR) and increases the mean
time between failure (MTBF) to achieve a high level of network availability.
Multichassis EtherChannels
Multichassis EtherChannel (MEC) is an EtherChannel bundled with physical ports having common characteristics such as speed
and duplex, that are distributed across each Cisco StackWise Virtual system. A Cisco StackWise Virtual MEC can connect to
any network element that supports EtherChannel (such as a host, server, router, or switch).
Cisco StackWise Virtual support up to 128 MECs deployed in Layer 2 or Layer 3 modes. EtherChannel 127 and 128 are reserved
for SVL connections. Hence, the maximum available MEC count is 126.On the Cisco Catalyst 9500X Series Switches, Cisco StackWise Virtual support up to 240 MECs deployed in Layer 2 or Layer 3 modes, and EtherChannel 241 is reserved for
internal SVL link EtherChannel bundling.
In a Cisco StackWise Virtual system, an MEC is an EtherChannel with additional capability. A multichassis EtherChannel link
reduces the amount of traffic that requires transmission across the SVL by populating the index port only with the ports local
to the physical switch. This allows the switch to give precedence to the local ports of the multichassis EtherChannel link
over those on the remote switch.
Each MEC can optionally be configured to support either Cisco PAgP, IEEE LACP, or Static ON mode. We recommend that you implement
EtherChannel using Cisco PAgP or LACP with a compatible neighbor. If a remotely connected neighbor such as Cisco Wireless
LAN Controller (WLC) does not support this link-bundling protocol, then a Static ON mode can be deployed. These protocols
run only on the Cisco StackWise Virtual active switch.
Note
On an SVL system, a maximum of 8 ports is supported in an LACP EtherChannel configuration.
An MEC can support up to eight physical links that can be distributed in any proportion between the Cisco StackWise Virtual
active switch and the Cisco StackWise Virtual standby switch. We recommend that you distribute the MEC ports across both switches
evenly.
MEC Minimum Latency Load Balancing
The StackWise Virtual environment is designed such that data forwarding always remains within the switch. The Virtual Stack
always tries to forward traffic on the locally available links. This is true for both Layer 2 and Layer3 links. The primary
motivation for local forwarding is to avoid unnecessarily sending data traffic over the SVL and thus reduce the latency (extra
hop over the SVL) and congestion.The bidirectional traffic is load-shared between the two StackWise Virtual members. However,
for each StackWise Virtual member, ingress and egress traffic forwarding is based on locally-attached links that are part
of MEC. This local forwarding is a key concept in understanding convergence and fault conditions in a StackWise Virtual enabled
campus network.
The active and standby switches support local forwarding that will individually perform the desired lookups and forward the
traffic on local links to uplink neighbors. If the destination is a remote switch in the StackWise Virtual domain, ingress
processing is performed on the ingress switch and then traffic is forwarded over the SVL to the egress switch where only egress
processing is performed.
MEC Failure Scenarios
The following sections describe issues that may arise and the resulting impact:
Single MEC Link Failure
If a link within a MEC fails (and other links in the MEC are still operational), the MEC redistributes the load among the
operational links, as in a regular port.
All MEC Links to the Cisco StackWise Virtual Active Switch Fail
If all the links to the Cisco StackWise Virtual active switch fail, a MEC becomes a regular EtherChannel with operational
links to the Cisco StackWise Virtual standby switch.
Data traffic that terminates on the Cisco StackWise Virtual active switch reaches the MEC by crossing the SVL to the Cisco
StackWise Virtual standby switch. Control protocols continue to run in the Cisco StackWise Virtual active switch. Protocol
messages reach the MEC by crossing the SVL.
All MEC Links Fail
If all the links in an MEC fail, the logical interface for the EtherChannel is set to Unavailable. Layer 2 control protocols
perform the same corrective action as for a link-down event on a regular EtherChannel.
On adjacent switches, routing protocols and the Spanning Tree Protocol (STP) perform the same corrective action as for a
regular EtherChannel.
Cisco StackWise Virtual Standby Switch Failure
If the Cisco StackWise Virtual standby switch fails, a MEC becomes a regular EtherChannel with operational links on the Cisco
StackWise Virtual active switch. Connected peer switches detect the link failures, and adjust their load-balancing algorithms
to use only the links to the StackWise Virtual active switch.
Cisco StackWise Virtual Active Switch Failure
Cisco StackWise Virtual active switch failure results in a stateful switchover (SSO). After the switchover, a MEC is operational
on the new Cisco StackWise Virtual active switch. Connected peer switches detect the link failures (to the failed switch),
and adjust their load-balancing algorithms to use only the links to the new Cisco StackWise Virtual active switch.
Cisco StackWise Virtual Packet Handling
In Cisco StackWise Virtual, the Cisco StackWise Virtual active switch runs the Layer 2 and Layer 3 protocols and features
and manages the ports on both the switches. Cisco StackWise Virtual uses SVL to communicate system and protocol information
between the peer switches and to carry data traffic between the two switches.
The following sections describe packet handling in Cisco StackWise Virtual.
Traffic on StackWise Virtual Link
SVL carries data traffic and in-band control traffic between two switches. All the frames that are forwarded over the SVL
are encapsulated with a special StackWise Virtual Header (SVH). The SVH adds an overhead of 64 bytes for control and data
traffic, which provides information for Cisco StackWise Virtual to forward the packet on the peer switch.
An SVL transports control messages between two switches. Messages include protocol messages that are processed by the Cisco
StackWise Virtual active switch, but received or transmitted by interfaces on the Cisco StackWise Virtual standby switch.
Control traffic also includes module programming between the Cisco StackWise Virtual active switch and the switching modules
on the Cisco StackWise Virtual standby switch.
Cisco StackWise Virtual transmits data traffic over an SVL under the following circumstances:
Layer 2 traffic flooded over a VLAN (even for dual-homed links).
Packets processed by software on the Cisco StackWise Virtual active switch where the ingress interface is on the Cisco StackWise
Virtual standby switch.
The packet destination is on the peer switch, as described in the following examples:
Traffic within a VLAN where the known destination interface is on the peer switch.
