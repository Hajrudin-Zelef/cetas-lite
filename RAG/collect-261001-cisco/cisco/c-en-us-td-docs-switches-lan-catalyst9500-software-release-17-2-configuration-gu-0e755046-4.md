---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046-4
title: "c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046.md
source_anchor: ""
source_lines: [302, 477]
sha256: b51c93491b3254d14ddb263db0a5244b096392dd1d0f0e853636b5803240a316
---

# c-en-us-td-docs-switches-lan-catalyst9500-software-release-17-2-configuration-gu-0e755046

requires only one spanning-tree instance for all VLANs allowed on the trunks. However, in a network of Cisco devices that are connected through IEEE 802.1Q trunks, the devices
maintain one spanning-tree instance for each VLAN allowed on the trunks.
When you connect a Cisco device to a non-Cisco device through an IEEE 802.1Q trunk, the Cisco device uses PVST+ to provide
spanning-tree interoperability. If Rapid PVST+ is enabled, the device uses it instead of PVST+. The device combines the spanning-tree
instance of the IEEE 802.1Q VLAN of the trunk with the spanning-tree instance of the non-Cisco IEEE 802.1Q device.
However, all PVST+ or Rapid PVST+ information is maintained by Cisco devices that are separated by a cloud of non-Cisco IEEE
802.1Q devices. The non-Cisco IEEE 802.1Q cloud separating the Cisco devices is treated as a single trunk link between the
devices.
Rapid PVST+ is automatically enabled on IEEE 802.1Q trunks, and no user configuration is required. The external spanning-tree
behavior on access ports and Inter-Switch Link (ISL) trunk ports is not affected by PVST+.
Spanning Tree and Switch Stacks
When the switch stack is operating in PVST+ or Rapid PVST+ mode:
A switch stack appears as a single spanning-tree node to the rest of the network, and all stack members use the same bridge
ID for a given spanning tree. The bridge ID is derived from the MAC address of the active switch.
When a new device joins the stack, it sets its bridge ID to the active switch bridge ID. If the newly added device has the
lowest ID and if the root path cost is the same among all stack members, the newly added device becomes the stack root.
When a stack member leaves the stack, spanning-tree reconvergence occurs within the stack (and possibly outside the stack).
The remaining stack member with the lowest stack port ID becomes the stack root.
If the switch stack is the spanning-tree root and the active switch fails or leaves the stack, the standby switch becomes
the new active switch, bridge IDs remain the same, and a spanning-tree reconvergence might occur.
If a neighboring device external to the switch stack fails or is powered down, normal spanning-tree processing occurs. Spanning-tree
reconvergence might occur as a result of losing a device in the active topology.
If a new device external to the switch stack is added to the network, normal spanning-tree processing occurs. Spanning-tree
reconvergence might occur as a result of adding a device in the network.
Default Spanning-Tree Configuration
Table 3. Default Spanning-Tree Configuration
Feature
Default Setting
Enable state
Enabled on VLAN 1.
Spanning-tree mode
Rapid PVST+ ( PVST+ and MSTP are disabled.)
Device priority
32768
Spanning-tree port priority (configurable on a per-interface basis)
128
Spanning-tree port cost (configurable on a per-interface basis)
Note
These values are supported on the C9500-12Q, C9500-24Q, C9500-16X and C9500-40X models of the Cisco Catalyst 9500 Series Switches.
10 Mbps: 100
100 Mbps: 19
1 Gbps: 4
10 Gbps: 2
25 Gbps: 1
40 Gbps: 1
Spanning-tree port cost (configurable on a per-interface basis)
Note
These values are supported on the C9500-32C, C9500-32QC, C9500-48Y4C, and C9500-24Y4C models of the Cisco Catalyst 9500 Series
Switches.
10 Mbps: 2000000
100 Mbps: 200000
1 Gbps: 20000
10 Gbps: 2000
25 Gbps: 800
40 Gbps: 500
100 Gbps: 200
1 Tbps: 20
10 Tbps: 2
Spanning-tree VLAN port priority (configurable on a per-VLAN basis)
128
Spanning-tree VLAN port cost (configurable on a per-VLAN basis)
Note
These values are supported on the C9500-12Q, C9500-24Q, C9500-16X and C9500-40X models of the Cisco Catalyst 9500 Series Switches
10 Mbps: 100
100 Mbps: 19
1 Gbps: 4
10 Gbps: 2
25 Gbps: 1
40 Gbps: 1
Spanning-tree VLAN port cost (configurable on a per-VLAN basis)
Note
These values are supported on the C9500-32C, C9500-32QC, C9500-48Y4C, and C9500-24Y4C models of the Cisco Catalyst 9500 Series
Switches.
10 Mbps: 2000000
100 Mbps: 200000
1 Gbps: 20000
10 Gbps: 2000
25 Gbps: 800
40 Gbps: 500
100 Gbps: 200
1 Tbps: 20
10 Tbps: 2
Spanning-tree timers
Hello time: 2 seconds
Forward-delay time: 15 seconds
Maximum-aging time: 20 seconds
Transmit hold count: 6 BPDUs
Note
Beginning in Cisco IOS Release 15.2(4)E, the default STP mode is Rapid PVST+.
How to Configure Spanning Tree Protocol
The following sections provide information about configuring spanning tree protocol:
The switch supports three spanning-tree modes: per-VLAN spanning tree plus (PVST+), Rapid PVST+, or Multiple Spanning Tree
Protocol (MSTP). By default, the device runs the Rapid PVST+ protocol.
If you want to enable a mode that is different from the default mode, this procedure is required.
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
spanning-tree mode {pvst | mst | rapid-pvst}
Example:
Device(config)# spanning-tree mode pvst
Configures a spanning-tree mode.
All stack members run the same version of spanning tree.
Select pvst to enable PVST+.
Select mst to enable MSTP.
Select rapid-pvst to enable rapid PVST+.
Step 4
interfaceinterface-id
Example:
Device(config)# interface GigabitEthernet1/0/1
Specifies an interface to configure, and enters interface configuration mode. Valid interfaces include physical ports, VLANs,
and port channels. The VLAN ID range is 1 to 4094. The port-channel range is 1 to 128.
Specifies that the link type for this port is point-to-point.
If you connect this port (local port) to a remote port through a point-to-point link and the local port becomes a designated
port, the device negotiates with the remote port and rapidly changes the local port to the forwarding state.
Step 6
end
Example:
Device(config-if)# end
Returns to privileged EXEC mode.
Step 7
clear spanning-tree detected-protocols
Example:
Device# clear spanning-tree detected-protocols
If any port on the device is connected to a port on a legacy IEEE 802.1D device, this command restarts the protocol migration
process on the entire device.
This step is optional if the designated device detects that this device is running rapid PVST+.
(Optional) Disabling Spanning Tree
Spanning tree is enabled by default on VLAN 1 and on all newly created VLANs up to the spanning-tree limit. Disable spanning
tree only if you are sure that there are no loops in the network topology.
Caution
When spanning tree is disabled and loops are present in the topology, excessive traffic and indefinite packet duplication
can drastically reduce network performance.
To disable spanning tree, perform this procedure:
Procedure
Command or Action
Purpose
Step 1
enable
Example:
Device> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Device# configure terminal
Enters global configuration mode.
Step 3
no spanning-tree vlanvlan-id
Example:
Device(config)# no spanning-tree vlan 300
For vlan-id, the range is 1 to 4094.
Step 4
end
Example:
Device(config)# end
Returns to privileged EXEC mode.
(Optional) Configuring the Root Device
To configure a device as the root for the specified VLAN, use the spanning-tree vlanvlan-idroot global configuration command to modify the device priority from the default value (32768) to a significantly lower value.
When you enter this command, the software checks the switch priority of the root switches for each VLAN. Because of the extended
system ID support, the switch sets its own priority for the specified VLAN to 24576 if this value causes this switch to become
the root for the specified VLAN.
Use the diameter keyword to specify the Layer 2 network diameter (that is, the maximum number of device hops between any two end stations
in the Layer 2 network). When you specify the network diameter, the device automatically sets an optimal hello time, forward-delay
