---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-8952629b
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-8952629b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-8952629b.md
source_anchor: ""
source_lines: [1, 66]
sha256: 9ac66730343f7e5bf98e2bcd59501b3e85aaefc284cb42f8604f9cd915f31551
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-interfaces-cisco-nexus-90-8952629b

About Layer 3 Interfaces
Layer 3 interfaces forward IPv4 and IPv6 packets to another device using static or dynamic routing protocols. You can use Layer 3 interfaces for IP routing and inter-VLAN routing of Layer 2 traffic.
Routed Interfaces
You can configure a port as a Layer 2 interface or a Layer 3 interface. A routed interface is a physical port that can route IP traffic to another device. A routed interface is a Layer 3 interface only and does not support Layer 2 protocols, such as the Spanning Tree Protocol (STP).
All Ethernet ports are routed interfaces by default. You can change this default behavior with the CLI setup script.
| Note | The default behavior varies based on the type of switch (Cisco Nexus 9300, Cisco Nexus 9500, or Cisco Nexus 3164). | 
| Note | Cisco Nexus 9300 Series switches (except Cisco Nexus 9332 switch) have a Layer 2 default mode. | 
You can assign an IP address to the port, enable routing, and assign routing protocol characteristics to this routed interface.
You can also create a Layer 3 port channel from routed interfaces. For more information about port channels, see the "Configuring Port Channels" section.
Routed interfaces support exponentially decayed rate counters. Cisco NX-OS tracks the following statistics with these averaging counters:
- 
                                    					
                                    Input packets/sec
- 
                                    					
                                    Output packets/sec
- 
                                    					
                                    Input bytes/sec
- 
                                    					
                                    Output bytes/sec
Subinterfaces
You can create virtual subinterfaces on a parent interface configured as a Layer 3 interface. A parent interface can be a physical port.
Subinterfaces divide the parent interface into two or more virtual interfaces on which you can assign unique Layer 3 parameters such as IP addresses and dynamic routing protocols. The IP address for each subinterface should be in a different subnet from any other subinterface on the parent interface.
You create a subinterface with a name that consists of the parent interface name (for example, Ethernet 2/1) followed by a period and then by a number that is unique for that subinterface. For example, you could create a subinterface for Ethernet interface 2/1 named Ethernet 2/1.1 where .1 indicates the subinterface.
Cisco NX-OS enables subinterfaces when the parent interface is enabled. You can shut down a subinterface independent of shutting down the parent interface. If you shut down the parent interface, Cisco NX-OS shuts down all associated subinterfaces as well.
One use of subinterfaces is to provide unique Layer 3 interfaces to each virtual local area network (VLAN) supported by the parent interface. In this scenario, the parent interface connects to a Layer 2 trunking port on another device. You configure a subinterface and associate the subinterface to a VLAN ID using 802.1Q trunking.
The following figure shows a trunking port from a switch that connects to router B on interface E 2/1. This interface contains three subinterfaces that are associated with each of the three VLANs carried by the trunking port.
For more information about VLANs, see the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide.
VLAN Interfaces
A VLAN interface, or switch virtual interface (SVI), is a virtual routed interface that connects a VLAN on the device to the Layer 3 router engine on the same device. Only one VLAN interface can be associated with a VLAN.
However, you need to configure a VLAN interface for a VLAN only when you want to route between VLANs or to provide IP host connectivity to the device through a virtual routing and forwarding (VRF) instance that is not the management VRF. When you enable VLAN interface creation, Cisco NX-OS creates a VLAN interface for the default VLAN (VLAN 1) to permit remote switch administration.
Enable the VLAN network interface feature using the feature interface-vlan configuration. The system automatically takes a checkpoint prior to disabling the feature, and you can roll back to this checkpoint. See the Cisco Nexus 9000 Series NX-OS System Management Configuration Guide for information on rollbacks and checkpoints.
| Note | The feature interface-vlan configuration is not available on the Nexus 9800 switches. | 
Layer 3 inter-VLAN Routing
You can route traffic across VLAN interfaces to provide Layer 3 inter-VLAN routing by configuring a VLAN interface for each VLAN, and assigning an IP address on the VLAN interface.
For more information about IP addresses and IP routing, see the Cisco Nexus 9000 Series NX-OS Unicast Routing Configuration Guide.
Connecting Two VLAN Interfaces
You can configure VLAN interfaces for each VLAN that allows Host 1 to communicate with Host 2 using IP routing between the VLANs. VLAN 1 communicates at Layer 3 over VLAN interface 1 and VLAN 10 communicates at Layer 3 over VLAN interface 10.
The following figure shows two hosts connected to two VLANs on a device.
| Note | You cannot delete the VLAN interface for VLAN 1. | 
Loopback Interfaces
A loopback interface is a virtual interface with a single endpoint that is always up. Any packet transmitted over a loopback interface is immediately received by this interface. Loopback interfaces emulate a physical interface. You can configure up to 1024 loopback interfaces, numbered 0 to 1023.
You can use loopback interfaces for performance analysis, testing, and local communications. Loopback interfaces can act as a termination address for routing protocol sessions. This loopback configuration allows routing protocol sessions to stay up even if some of the outbound interfaces are down.
High Availability
Layer 3 interfaces support stateful and stateless restarts. After the switchover, Cisco NX-OS applies the runtime configuration after the switchover.
See the Cisco Nexus 9000 Series NX-OS High Availability and Redundancy Guide for complete information about high availability.
Virtualization Support
Layer 3 interfaces support Virtual Routing and Forwarding instances (VRFs). VRFs exist within virtual device contexts (VDCs). By default, Cisco NX-OS places you in the default VDC and default VRF.
| Note | You must assign an interface to a VRF before you configure the IP address for that interface. | 
Layer 3 Static MAC Addresses
You can configure a static MAC address for the following Layer 3 interfaces:
- 
                                    					
                                    Layer 3 interfaces
- 
                                    					
                                    Layer 3 subinterfaces
- 
                                    					
                                    Layer 3 port channels
- 
                                    					
                                    VLAN network interface
| Note | You cannot configure static MAC address on tunnel interfaces. |
