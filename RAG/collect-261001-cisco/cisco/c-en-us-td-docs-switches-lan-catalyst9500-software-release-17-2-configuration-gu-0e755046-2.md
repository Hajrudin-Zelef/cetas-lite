---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046-2
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046.md
source_anchor: ""
source_lines: [77, 213]
sha256: 016063bfbedf3010a265ab01b0616aac4017b63bc5639b54f0ff455ddba41502
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046

blocking mode.
Bridge ID, Device Priority, and Extended System ID
The IEEE 802.1D standard requires that each device has a unique bridge identifier (bridge ID), which controls the selection of the root switch. Because each VLAN is considered as a different
logical bridge with PVST+ and Rapid PVST+, the same device must have a different bridge ID for each configured VLAN. Each VLAN on the device
has a unique 8-byte bridge ID. The 2 most-significant bytes are used for the device priority, and the remaining 6 bytes are
derived from the device MAC address.
The 2 bytes previously used for the device priority are reallocated into a 4-bit priority value and a 12-bit extended system
ID value equal to the VLAN ID.
Table 1. Device Priority Value and Extended System ID
Priority Value
Extended System ID (Set Equal to the VLAN ID)
Bit 16
Bit 15
Bit 14
Bit 13
Bit 12
Bit 11
Bit 10
Bit 9
Bit 8
Bit 7
Bit 6
Bit 5
Bit 4
Bit 3
Bit 2
Bit 1
32768
16384
8192
4096
2048
1024
512
256
128
64
32
16
8
4
2
1
Spanning tree uses the extended system ID, the device priority, and the allocated spanning-tree MAC address to make the bridge
ID unique for each VLAN.
Support for the extended system ID affects how you manually configure the root switch, the secondary root switch, and the
switch priority of a VLAN. For example, when you change the switch priority value, you change the probability that the switch
will be elected as the root switch. Configuring a higher value decreases the probability; a lower value increases the probability.
Port Priority Versus Path Cost
If a loop occurs, spanning tree uses port priority when selecting an interface to put into the forwarding state. You can assign
higher priority values (lower numerical values) to interfaces that you want selected first and lower priority values (higher
numerical values) that you want selected last. If all interfaces have the same priority value, spanning tree puts the interface
with the lowest interface number in the forwarding state and blocks the other interfaces.
The spanning-tree path cost default value is derived from the media speed of an interface. If a loop occurs, spanning tree
uses cost when selecting an interface to put in the forwarding state. You can assign lower cost values to interfaces that
you want selected first and higher cost values that you want selected last. If all interfaces have the same cost value, spanning
tree puts the interface with the lowest interface number in the forwarding state and blocks the other interfaces.
If your device is a member of a switch stack, you must assign lower cost values to interfaces that you want selected first
and higher cost values that you want selected last instead of adjusting its port priority.
Spanning-Tree Interface States
Propagation delays can occur when protocol information passes through a switched LAN. As a result, topology changes can take
place at different times and at different places in a switched network. When an interface transitions directly from nonparticipation
in the spanning-tree topology to the forwarding state, it can create temporary data loops. Interfaces must wait for new topology
information to propagate through the switched LAN before starting to forward frames. They must allow the frame lifetime to
expire for forwarded frames that have used the old topology.
Each Layer 2 interface on a device using spanning tree exists in one of these states:
Blocking—The interface does not participate in frame forwarding.
Listening—The first transitional state after the blocking state when the spanning tree decides that the interface should participate
in frame forwarding.
Learning—The interface prepares to participate in frame forwarding.
Forwarding—The interface forwards frames.
Disabled—The interface is not participating in spanning tree because of a shutdown port, no link on the port, or no spanning-tree
instance running on the port.
An interface moves through these states:
From initialization to blocking
From blocking to listening or to disabled
From listening to learning or to disabled
From learning to forwarding or to disabled
From forwarding to disabled
Figure 1. Spanning-Tree Interface States. An interface moves through the states.
When you power up the device, spanning tree is enabled by default, and every interface in the device, VLAN, or network goes
through the blocking state and the transitory states of listening and learning. Spanning tree stabilizes each interface at
the forwarding or blocking state.
When the spanning-tree algorithm places a Layer 2 interface in the forwarding state, this process occurs:
The interface is in the listening state while spanning tree waits for protocol information to move the interface to the blocking
state.
While spanning tree waits for the forward-delay timer to expire, it moves the interface to the learning state and resets the
forward-delay timer.
In the learning state, the interface continues to block frame forwarding as the device learns end-station location information
for the forwarding database.
When the forward-delay timer expires, spanning tree moves the interface to the forwarding state, where both learning and frame
forwarding are enabled.
A Layer 2 interface in the blocking state does not participate in frame forwarding. After initialization, a BPDU is sent to
each device interface. A device initially functions as the root until it exchanges BPDUs with other devices. This exchange
establishes which device in the network is the root or root device. If there is only one device in the network, no exchange
occurs, the forward-delay timer expires, and the interface moves to the listening state. An interface always enters the blocking
state after device initialization.
An interface in the blocking state performs these functions:
Discards frames received on the interface
Discards frames that are switched from another interface for forwarding
Does not learn addresses
Receives BPDUs
Listening State
The listening state is the first state a Layer 2 interface enters after the blocking state. The interface enters this state
when the spanning tree decides that the interface should participate in frame forwarding.
An interface in the listening state performs these functions:
Discards frames received on the interface
Discards frames that are switched from another interface for forwarding
Does not learn addresses
Receives BPDUs
Learning State
A Layer 2 interface in the learning state prepares to participate in frame forwarding. The interface enters the learning state
from the listening state.
An interface in the learning state performs these functions:
Discards frames received on the interface
Discards frames that are switched from another interface for forwarding
Learns addresses
Receives BPDUs
Forwarding State
A Layer 2 interface in the forwarding state forwards frames. The interface enters the forwarding state from the learning state.
An interface in the forwarding state performs these functions:
Receives and forwards frames that are received on the interface.
Forwards frames that are switched from another interface
Learns addresses
Receives BPDUs
Disabled State
A Layer 2 interface in the disabled state does not participate in frame forwarding or in the spanning tree. An interface in
the disabled state is nonoperational.
A disabled interface performs these functions:
Discards frames received on the interface
Discards frames that are switched from another interface for forwarding
Does not learn addresses
Does not receive BPDUs
How a Device or Port Becomes the Root Device or Root Port
If all devices in a network are enabled with default spanning-tree settings, the device with the lowest MAC address becomes
the root device.
Figure 2. Spanning-Tree Topology. Switch A is elected as the root device because the device priority of all the devices is set to the default (32768) and Switch A
