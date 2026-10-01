---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide-e2ffcac9-2
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "cost", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9.md
source_anchor: ""
source_lines: [77, 170]
sha256: ab1784f9f3a551dce54cec8fe22797cec6d6a2a9786b5ba3437e3e9ff110a3df
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-layer2-configuration-guide--e2ffcac9

                                          20480
-  
                                          			 
                                          24576
-  
                                          			 
                                          28672
-  
                                          			 
                                          32768
-  
                                          			 
                                          36864
-  
                                          			 
                                          40960
-  
                                          			 
                                          45056
-  
                                          			 
                                          49152
-  
                                          			 
                                          53248
-  
                                          			 
                                          57344
-  
                                          			 
                                          61440
STP uses the extended system ID plus a MAC address to make the bridge ID unique for each VLAN.
| Note | If another bridge in the same spanning tree domain does not run the MAC address reduction feature, it could win the root bridge ownership because of the finer granularity in the selection of its bridge ID. | 
BPDUs
Network devices transmit BPDUs throughout the STP instance. Each network device sends configuration BPDUs to communicate and compute the spanning tree topology. Each configuration BPDU contains the following minimal information:
-  
                                       			 
                                       The unique bridge ID of the network device that the transmitting network device believes to be the root bridge
-  
                                       			 
                                       The STP path cost to the root
-  
                                       			 
                                       The bridge ID of the transmitting bridge
-  
                                       			 
                                       The message age
-  
                                       			 
                                       The identifier of the transmitting port
-  
                                       			 
                                       Values for the hello, forward delay, and max-age protocol timer
-  
                                       			 
                                       Additional information for STP extension protocols
When a network device transmits a Rapid PVST+ BPDU frame, all network devices connected to the VLAN on which the frame is transmitted receive the BPDU. When a network device receives a BPDU, it does not forward the frame but instead uses the information in the frame to calculate a BPDU. If the topology changes, the device initiates a BPDU exchange.
A BPDU exchange results in the following:
-  
                                       			 
                                       One network device is elected as the root bridge.
-  
                                       			 
                                       The shortest distance to the root bridge is calculated for each network device based on the path cost.
-  
                                       			 
                                       A designated bridge for each LAN segment is selected. This network device is closest to the root bridge through which frames are forwarded to the root.
-  
                                       			 
                                       A root port is elected. This port provides the best path from the bridge to the root bridge.
-  
                                       			 
                                       Ports included in the spanning tree are selected.
Election of the Root Bridge
For each VLAN, the network device with the lowest numerical ID is elected as the root bridge. If all network devices are configured with the default priority (32768), the network device with the lowest MAC address in the VLAN becomes the root bridge. The bridge priority value occupies the most significant bits of the bridge ID.
When you change the bridge priority value, you change the probability that the device will be elected as the root bridge. Configuring a lower value increases the probability; a higher value decreases the probability.
The STP root bridge is the logical center of each spanning tree topology in a Layer 2 network. All paths that are not needed to reach the root bridge from anywhere in the Layer 2 network are placed in STP blocking mode.
BPDUs contain information about the transmitting bridge and its ports, including bridge and MAC addresses, bridge priority, port priority, and path cost. STP uses this information to elect the root bridge for the STP instance, to elect the root port that leads to the root bridge, and to determine the designated port for each Layer 2 segment.
Creating the Spanning Tree Topology
By lowering the numerical value of the ideal network device so that it becomes the root bridge, you force an STP recalculation to form a new spanning tree topology with the ideal network device as the root.
When the spanning tree topology is calculated based on default parameters, the path between the source and destination end stations in a switched network might not be ideal. For instance, connecting higher-speed links to a port that has a higher number than the current root port can cause a root-port change. The goal is to make the fastest link the root port.
For example, assume that one port on switch B is a fiber-optic link, and another port on switch B (an unshielded twisted-pair [UTP] link) is the root port. Network traffic might be more efficient over the high-speed fiber-optic link. By changing the STP port priority on the fiber-optic port to a higher priority (lower numerical value) than the root port, the fiber-optic port becomes the new root port.
Rapid PVST+
Rapid PVST+ is the default spanning tree mode for the software and is enabled by default on the default VLAN and all newly created VLANs.
A single instance, or topology, of RSTP runs on each configured VLAN, and each Rapid PVST+ instance on a VLAN has a single root device. You can enable and disable STP on a per-VLAN basis when you are running Rapid PVST+.
Overview of Rapid PVST+
Rapid PVST+ is the IEEE 802.1w (RSTP) standard implemented per VLAN. A single instance of STP runs on each configured VLAN (if you do not manually disable STP). Each Rapid PVST+ instance on a VLAN has a single root switch. You can enable and disable STP on a per-VLAN basis when you are running Rapid PVST+.
| Note | Rapid PVST+ is the default STP mode for the device. | 
Rapid PVST+ uses point-to-point wiring to provide rapid convergence of the spanning tree. The spanning tree reconfiguration can occur in less than 1 second with Rapid PVST+ (in contrast to 50 seconds with the default settings in the 802.1D STP). The device automatically checks the PVID.
| Note | Rapid PVST+ supports one STP instance for each VLAN. | 
Using Rapid PVST+, STP convergence occurs rapidly. By default, each designated port in the STP sends out a BPDU every 2 seconds. On a designated port in the topology, if hello messages are missed three consecutive times, or if the maximum age expires, the port immediately flushes all protocol information in the table. A port considers that it loses connectivity to its direct neighbor designated port if it misses three BPDUs or if the maximum age expires. This rapid aging of the protocol information allows quick failure detection.
Rapid PVST+ provides for rapid recovery of connectivity following the failure of a device, a device port, or a LAN. It provides rapid convergence for edge ports, new root ports, and ports connected through point-to-point links as follows:
-  
                                       			 
