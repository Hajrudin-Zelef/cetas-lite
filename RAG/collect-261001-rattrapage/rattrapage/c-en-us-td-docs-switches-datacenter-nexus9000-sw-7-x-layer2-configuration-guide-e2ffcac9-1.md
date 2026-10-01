---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide-e2ffcac9-1
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9.md
source_anchor: ""
source_lines: [1, 76]
sha256: aa8f3ecbb7699024346f239aff3c483b459aa745416d97b4b12be5ff804bcc03
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9

Information About Rapid PVST+
| Note | See the Cisco Nexus 9000 Series NX-OS Interfaces Configuration Guide, for information on creating Layer 2 interfaces. | 
The Spanning Tree Protocol (STP) was implemented to provide a loop-free network at Layer 2 of the network. Rapid PVST+ is an updated implementation of STP that allows you to create one spanning tree topology for each VLAN. Rapid PVST+ is the default STP mode on the device.
| Note | Spanning tree is used to refer to IEEE 802.1w and IEEE 802.1s. If the IEEE 802.1D Spanning Tree Protocol is discussed in this publication, then 802.1D is stated specifically. | 
| Note | Rapid PVST+ is the default STP mode. | 
The Rapid PVST+ protocol is the IEEE 802.1w standard, Rapid Spanning Tree Protocol (RSTP), implemented on a per VLAN basis. Rapid PVST+ interoperates with the IEEE 802.1Q VLAN standard, which mandates a single STP instance for all VLANs, rather than per VLAN.
Rapid PVST+ is enabled by default on the default VLAN (VLAN1) and on all newly created VLANs on the device. Rapid PVST+ interoperates with devices that run legacy IEEE 802.1D STP.
RSTP is an improvement on the original STP standard, 802.1D, which allows faster convergence.
| Note | The device supports full nondisruptive upgrades for Rapid PVST+. See the Cisco Nexus 9000 Series NX-OS High Availability and Redundancy Guide, for complete information on nondisruptive upgrades. | 
STP
STP is a Layer 2 link-management protocol that provides path redundancy while preventing loops in the network.
Overview of STP
In order for a Layer 2 Ethernet network to function properly, only one active path can exist between any two stations. STP operation is transparent to end stations, which cannot detect whether they are connected to a single LAN segment or a switched LAN of multiple segments.
When you create fault-tolerant internetworks, you must have a loop-free path between all nodes in a network. The STP algorithm calculates the best loop-free path throughout a switched Layer 2 network. Layer 2 LAN ports send and receive STP frames, which are called Bridge Protocol Data Units (BPDUs), at regular intervals. Network devices do not forward these frames but use the frames to construct a loop-free path.
Multiple active paths between end stations cause loops in the network. If a loop exists in the network, end stations might receive duplicate messages and network devices might learn end station MAC addresses on multiple Layer 2 LAN ports.
STP defines a tree with a root bridge and a loop-free path from the root to all network devices in the Layer 2 network. STP forces redundant data paths into a blocked state. If a network segment in the spanning tree fails and a redundant path exists, the STP algorithm recalculates the spanning tree topology and activates the blocked path.
When two Layer 2 LAN ports on a network device are part of a loop, the STP port priority and port path-cost setting determine which port on the device is put in the forwarding state and which port is put in the blocking state. The STP port priority value is the efficiency with which that location allows the port to pass traffic. The STP port path-cost value is derived from the media speed.
How a Topology is Created
All devices in a LAN that participate in a spanning tree gather information about other switches in the network by exchanging BPDUs. This exchange of BPDUs results in the following actions:
-  
                                       			 
                                       The system elects a unique root switch for the spanning tree network topology.
-  
                                       			 
                                       The system elects a designated switch for each LAN segment.
-  
                                       			 
                                       The system eliminates any loops in the switched network by placing redundant switch ports in a backup state; all paths that are not needed to reach the root device from anywhere in the switched network are placed in an STP-blocked state.
The topology on an active switched network is determined by the following:
-  
                                       			 
                                       The unique device identifier Media Access Control (MAC) address of the device that is associated with each device
-  
                                       			 
                                       The path cost to the root that is associated with each switch port
-  
                                       			 
                                       The port identifier that is associated with each switch port
In a switched network, the root switch is the logical center of the spanning tree topology. STP uses BPDUs to elect the root switch and root port for the switched network.
| Note | The mac-address bpdu source version 2 command enables STP to use the new Cisco MAC address (00:26:0b:xx:xx:xx) as the source address of BPDUs generated on vPC ports. To apply this command, you must have identical configurations for both vPC peer switches or peers. Cisco strongly recommends that you disable ether channel guard on the edge devices before issuing this command to minimize traffic disruption from STP inconsistencies. Re-enable the ether channel guard after updating on both peers. | 
Bridge ID
Each VLAN on each network device has a unique 64-bit bridge ID that consists of a bridge priority value, an extended system ID (IEEE 802.1t), and an STP MAC address allocation.
Bridge Priority Value
The bridge priority is a 4-bit value when the extended system ID is enabled.
You can only specify a device bridge ID (used by the spanning tree algorithm to determine the identity of the root bridge; the lowest number is preferred) as a multiple of 4096.
| Note | In this device, the extended system ID is always enabled; you cannot disable the extended system ID. | 
Extended System ID
The device always uses the 12-bit extended system ID.
This table shows how the system ID extension combined with the bridge ID functions as the unique identifier for a VLAN.
| Table 1. Bridge Priority Value and Extended System ID                                              with the Extended System ID Enabled |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Bridge Priority Value |  |  |  |  | Extended System ID (Set Equal to the VLAN ID) |  |  |  |  |  |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| Bit 16 | Bit 15 | Bit 14 | Bit 13 |  | Bit 12 | Bit 11 | Bit 10 | Bit 9 | Bit 8 | Bit 7 | Bit 6 | Bit 5 | Bit 4 | Bit 3 | Bit 2 | Bit 1 | 
| 32768 | 16384 | 8192 | 4096 |  | 2048 | 1024 | 512 | 256 | 128 | 64 | 32 | 16 | 8 | 4 | 2 | 1 | 
STP MAC Address Allocation
| Note | MAC address reduction is always enabled on the device. | 
Because MAC address reduction is always enabled on the device, you should also enable MAC address reduction on all other Layer 2 connected network devices to avoid undesirable root bridge election and spanning tree topology issues.
When MAC address reduction is enabled, the root bridge priority becomes a multiple of 4096 plus the VLAN ID. You can only specify a device bridge ID (used by the spanning tree algorithm to determine the identity of the root bridge; the lowest number is preferred) as a multiple of 4096. Only the following values are possible:
-  
                                          			 
                                          0
-  
                                          			 
                                          4096
-  
                                          			 
                                          8192
-  
                                          			 
                                          12288
-  
                                          			 
                                          16384
-  
                                          			 
