---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide-e2ffcac9-3
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "full-duplex"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9.md
source_anchor: ""
source_lines: [171, 235]
sha256: a36204a0cb7e27935b101b5d67e90ecc6494a623b24e091f5359bb447ebf99ae
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9

                                       Edge ports—When you configure a port as an edge port on an RSTP device, the edge port immediately transitions to the forwarding state. (This immediate transition was previously a Cisco-proprietary feature named PortFast.) You should only configure ports that connect to a single end station as edge ports. Edge ports do not generate topology changes when the link changes.
Enter the spanning-tree port type interface configuration command to configure a port as an STP edge port.
| Note | We recommend that you configure all ports connected to a Layer 2 host as edge ports. | 
-  
                                       			 
                                       Root port—If Rapid PVST+ selects a new root port, it blocks the old root port and immediately transitions the new root port to the forwarding state.
-  
                                       			 
                                       Point-to-point links—If you connect a port to another port through a point-to-point link and the local port becomes a designated port, it negotiates a rapid transition with the other port by using the proposal-agreement handshake to ensure a loop-free topology.
Rapid PVST+ achieves rapid transition to the forwarding state only on edge ports and point-to-point links. Although the link type is configurable, the system automatically derives the link type information from the duplex setting of the port. Full-duplex ports are assumed to be point-to-point ports, while half-duplex ports are assumed to be shared ports.
Edge ports do not generate topology changes, but all other designated and root ports generate a topology change (TC) BPDU when they either fail to receive three consecutive BPDUs from the directly connected neighbor or the maximum age times out. At this point, the designated or root port sends a BPDU with the TC flag set. The BPDUs continue to set the TC flag as long as the TC While timer runs on that port. The value of the TC While timer is the value set for the hello time plus 1 second. The initial detector of the topology change immediately floods this information throughout the entire topology.
When Rapid PVST+ detects a topology change, the protocol does the following:
-  
                                       			 
                                       Starts the TC While timer with a value equal to twice the hello time for all the nonedge root and designated ports, if necessary.
-  
                                       			 
                                       Flushes the MAC addresses associated with all these ports.
The topology change notification floods quickly across the entire topology. The system flushes dynamic entries immediately on a per-port basis when it receives a topology change.
| Note | The TCA flag is used only when the device is interacting with devices that are running legacy 802.1D STP. | 
The proposal and agreement sequence then quickly propagates toward the edge of the network and quickly restores connectivity after a topology change.
Rapid PVST+ BPDUs
Rapid PVST+ and 802.1w use all six bits of the flag byte to add the following:
-  
                                       			 
                                       The role and state of the port that originates the BPDU
-  
                                       			 
                                       The proposal and agreement handshake
Another important change is that the Rapid PVST+ BPDU is type 2, version 2, which makes it possible for the device to detect connected legacy (802.1D) bridges. The BPDU for 802.1D is type 0, version 0.
Proposal and Agreement Handshake
The switch learns the link type from the port duplex mode; a full-duplex port is considered to have a point-to-point connection and a half-duplex port is considered to have a shared connection. You can override the default setting that is controlled by the duplex setting by entering the spanning-tree link-type interface configuration command.
This proposal/agreement handshake is initiated only when a nonedge port moves from the blocking to the forwarding state. The handshaking process then proliferates step-by-step throughout the topology.
Protocol Timers
This table describes the protocol timers that affect the Rapid PVST+ performance.
| Table 2. Rapid PVST+                                           		  Protocol Timers |  | 
|---|---|
| Variable | Description | 
|---|---|
| Hello timer | Determines how often each device broadcasts BPDUs to other network devices. The default is 2 seconds, and the range is from 1 to 10. | 
| Forward delay timer | Determines how long each of the listening and learning states last before the port begins forwarding. This timer is generally not used by the protocol, but it is used when interoperating with the 802.1D spanning tree. The default is 15 seconds, and the range is from 4 to 30 seconds. | 
| Maximum age timer | Determines the amount of time that protocol information received on a port is stored by the network device. This timer is generally not used by the protocol, but it is used when interoperating with the 802.1D spanning tree. The default is 20 seconds; the range is from 6 to 40 seconds. | 
Port Roles
Rapid PVST+ provides rapid convergence of the spanning tree by assigning port roles and learning the active topology. Rapid PVST+ builds upon the 802.1D STP to select the device with the highest switch priority (lowest numerical priority value) as the root bridge. Rapid PVST+ assigns one of these port roles to individual ports:
- 
                                       					
                                       Root port—Provides the best path (lowest cost) when the device forwards packets to the root bridge.
- 
                                       					
                                       Designated port—Connects to the designated device that has the lowest path cost when forwarding packets from that LAN to the root bridge. The port through which the designated device is attached to the LAN is called the designated port.
- 
                                       					
                                       Alternate port—Offers an alternate path toward the root bridge to the path provided by the current root port. An alternate port provides a path to another device in the topology.
- 
                                       					
                                       Backup port—Acts as a backup for the path provided by a designated port toward the leaves of the spanning tree. A backup port can exist only when two ports are connected in a loopback by a point-to-point link or when a device has two or more connections to a shared LAN segment. A backup port provides another path in the topology to the device.
- 
                                       					
                                       Disabled port—Has no role within the operation of the spanning tree.
In a stable topology with consistent port roles throughout the network, Rapid PVST+ ensures that every root port and designated port immediately transition to the forwarding state while all alternate and backup ports are always in the blocking state. Designated ports start in the blocking state. The port state controls the operation of the forwarding and learning processes.
Rapid PVST+ Port State Overview
Propagation delays can occur when protocol information passes through a switched LAN. As a result, topology changes can take place at different times and at different places in a switched network. When a Layer 2 LAN port transitions directly from nonparticipation in the spanning tree topology to the forwarding state, it can create temporary data loops. Ports must wait for new topology information to propagate through the switched LAN before starting to forward frames.
Each Layer 2 LAN port on the device that uses Rapid PVST+ or MST exists in one of the following four states:
-  
                                       			 
