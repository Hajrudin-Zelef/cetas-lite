---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046-1
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "cost", "distribution", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046.md
source_anchor: ""
source_lines: [1, 76]
sha256: ac626d3294a2b1c28bae0c0ed7715035d24dab87b14f1b9859046fac8d616420
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
This chapter describes how to configure the Spanning Tree Protocol (STP) on port-based VLANs on the Catalyst devices. The
device can use either the per-VLAN spanning-tree plus (PVST+) protocol based on the IEEE 802.1D standard and Cisco proprietary
extensions, or the rapid per-VLAN spanning-tree plus (rapid-PVST+) protocol based on the IEEE 802.1w standard. A device stack
appears as a single spanning-tree node to the rest of the network, and all stack members use the same bridge ID.
An attempt to configure a device as the root device fails if the value necessary to be the root device is less than 1.
If your network consists of devices that support and do not support the extended system ID, it is unlikely that the device
with the extended system ID support will become the root device. The extended system ID increases the device priority value
every time the VLAN number is greater than the priority of the connected devices running older software.
The root device for each spanning tree instance should be a backbone or distribution device. Do not configure an access device
as the spanning tree primary root.
Information About Spanning Tree Protocol
The following sections provide information about spanning tree protocol:
Spanning Tree Protocol (STP) is a Layer 2 link management protocol that provides path redundancy while preventing loops in
the network. For a Layer 2 Ethernet network to function properly, only one active path can exist between any two stations.
Multiple active paths among end stations cause loops in the network. If a loop exists in the network, end stations might receive
duplicate messages. Devices might also learn end-station MAC addresses on multiple Layer 2 interfaces. These conditions result
in an unstable network. Spanning-tree operation is transparent to end stations, which cannot detect whether they are connected
to a single LAN segment or a switched LAN of multiple segments.
The STP uses a spanning-tree algorithm to select one device of a redundantly connected network as the root of the spanning
tree. The algorithm calculates the best loop-free path through a switched Layer 2 network by assigning a role to each port
based on the role of the port in the active topology:
Root—A forwarding port elected for the spanning-tree topology
Designated—A forwarding port elected for every switched LAN segment
Alternate—A blocked port providing an alternate path to the root bridge in the spanning tree
Backup—A blocked port in a loopback configuration
The device that has all of its ports as the designated role or as the backup role is the root device. The device that has at least one of its ports in the designated role is called the designated device.
Spanning tree forces redundant data paths into a standby (blocked) state. If a network segment in the spanning tree fails
and a redundant path exists, the spanning-tree algorithm recalculates the spanning-tree topology and activates the standby
path. Devices send and receive spanning-tree frames, called bridge protocol data units (BPDUs), at regular intervals. The devices do not forward these frames but use them to construct
a loop-free path. BPDUs contain information about the sending device and its ports, including device and MAC addresses, device priority, port
priority, and path cost. Spanning tree uses this information to elect the root device and root port for the switched network
and the root port and designated port for each switched segment.
When two ports on a device are part of a loop, the spanning-tree and path cost settings control which port is put in the forwarding state and which is put in the blocking state. The spanning-tree
port priority value represents the location of a port in the network topology and how well it is located to pass traffic.
The path cost value represents the media speed.
Note
On the C9500-32C, C9500-32QC, C9500-48Y4C, and C9500-24Y4C models of the Cisco
Catalyst 9500 Series Switches, the long path cost method is the default STP path
cost method.
On the C9500-12Q, C9500-16X, C9500-24Q, C9500-40X models of the Cisco Catalyst 9500
Series Switches, the short path cost method is the default STP path cost method.
Note
In addition to STP, the device uses keepalive messages to detect loops. By default, keepalive is enabled on Layer 2 ports.
To disable keepalive, use the no keepalive command in interface configuration mode.
Spanning-Tree Topology and Bridge Protocol Data Units
The stable, active spanning-tree topology of a switched network is controlled by these elements:
The unique bridge ID (device priority and MAC address) associated with each VLAN on each device. In a switch stack, all switches use the same bridge ID for a given spanning-tree instance.
The spanning-tree path cost to the root device.
The port identifier (port priority and MAC address) associated with each Layer 2 interface.
When the devices in a network are powered up, each functions as the root device. Each device sends a configuration BPDU through
all its ports. The BPDUs communicate and compute the spanning-tree topology. Each configuration BPDU contains this information:
The unique bridge ID of the device that the sending device identifies as the root device.
The spanning-tree path cost to the root
The bridge ID of the sending device
Message age
The identifier of the sending interface
Values for the hello, forward delay, and max-age protocol timers
When a device receives a configuration BPDU that contains superior information (lower bridge ID, lower path cost, and so forth), it stores the information for that port. If this BPDU is received
on the root port of the device, the device also forwards it with an updated message to all attached LANs for which it is the
designated device.
If a device receives a configuration BPDU that contains inferior information to that currently stored for that port, it discards the BPDU. If the device is a designated device for the LAN
from which the inferior BPDU was received, it sends that LAN a BPDU containing the up-to-date information stored for that
port. In this way, inferior information is discarded, and superior information is propagated on the network.
A BPDU exchange results in these actions:
One device in the network is elected as the root switch (the logical center of the spanning-tree topology in a switched network). See the figure following the bullets.
For each VLAN, the device with the highest device priority (the lowest numerical priority value) is elected as the root switch.
If all devices are configured with the default priority (32768), the devices with the lowest MAC address in the VLAN becomes
the root device. The device priority value occupies the most significant bits of the bridge ID, .
A root port is selected for each device (except the root switch). This port provides the best path (lowest cost) when the
device forwards packets to the root switch.
The shortest distance to the root switch is calculated for each device based on the path cost.
A designated device for each LAN segment is selected. The designated device incurs the lowest path cost when forwarding packets
from that LAN to the root switch. The port through which the designated device is attached to the LAN is called the designated
port.
All paths that are not needed to reach the root switch from anywhere in the switched network are placed in the spanning-tree
