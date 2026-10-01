---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107-1
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107.md
source_anchor: ""
source_lines: [1, 92]
sha256: 5c010794a44eb104e05c0946cd9fe8f02f44bb34eca16514d963821535071f21
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107

vPCs
vPCs — concept overview.
vPCs
A virtual port channel (vPC) allows links that are physically connected to two Cisco Nexus 9000 Series devices to appear as a single port channel by a third device (see figure). The third device can be a switch, server, or any other networking device that supports port channels. A vPC can provide Layer 2 multipathing, which allows you to create redundancy and increase the bisectional bandwidth by enabling multiple parallel paths between nodes and allowing load balancing traffic.
- 
                                    
                                    Allows a single device to use a port channel across two upstream devices
- 
                                    
                                    Eliminates Spanning Tree Protocol (STP) blocked ports
- 
                                    
                                    Provides a loop-free topology
- 
                                    
                                    Uses all available uplink bandwidth
- 
                                    
                                    Provides fast convergence if either the link or a device fails
- 
                                    
                                    Provides link-level resiliency
- 
                                    
                                    Assures high availability
The virtual port channel (vPC) is a technology that allows a single downstream device to connect to two upstream devices as though they were one logical device.
- 
                                    
                                    Layer 2 port channel support
- 
                                    
                                    Link Aggregation Control Protocol (LACP) optional
- 
                                    
                                    Enables redundancy and load balancing
vPC supports trunk mode port channels with or without LACP, and improves network stability and convergence.
Protocol Details and Recommendations
You can use only Layer 2 port channels in the vPC. You configure the port channels by using one of the following:
- 
                                    
                                    No protocol
- 
                                    
                                    Link Aggregation Control Protocol (LACP)
When you configure the port channels in a vPC—including the vPC Peer-Link channel—without using LACP, each device can have up to 32 active links in a single port channel. When using LACP, each device can have 32 active links and eight standby links.
| Note | You must enable the vPC feature before you can configure or run the vPC functionality. | 
The system automatically takes a checkpoint prior to disabling the feature, and you can roll back to this checkpoint.
After you enable the vPC functionality, you create the peer-keepalive link, which sends heartbeat messages between the two vPC peer devices.
To ensure that you have the correct hardware to enable and run a vPC, enter the show hardware feature-capability command. If you see an X across from the vPC in your command output, your hardware cannot enable the vPC feature.
| Note | Devices attached to a vPC domain using port channels should be connected to both of vPC peers. | 
Peer-Link Creation Example
You can create a vPC Peer-Link by configuring a port channel on one Cisco Nexus 9000 Series chassis by using two or more Ethernet ports higher speed than 1-Gigabit Ethernet.
We recommend that you configure the vPC Peer-Link Layer 2 port channels as trunks. On another Cisco Nexus 9000 Series chassis, you configure another port channel again using two or more Ethernet ports with speed higher than 1-Gigabit in the dedicated port mode.
Connecting these two port channels creates a vPC Peer-Link in which the two linked Cisco Nexus devices appear as one device to a third device.
Incorrect Hardware or Module Usage
If you are not using the correct module, the system displays an error message.
Once you configure this feature and if the primary vPC peer device fails, the system automatically suspends all the vPC links on the primary vPC peer device.
Track Object Recommendation
You can create a track object and apply that object to all links on the primary vPC peer device that connect to the core and to the vPC Peer-Link.
If you must configure all the vPC Peer-Links and core-facing interfaces on a single module, you should configure a track object.
Hitless vPC Role Changes
A virtual port channel (vPC) allows links that are physically connected to two different Cisco Nexus 9000 Series devices to appear as a single port channel. The vPC role change feature enables you switch vPC roles between vPC peers without impacting traffic flow. The vPC role switching is done based on the role priority value of the device under the vPC domain. A vPC peer device with lower role priority is selected as the primary vPC device during the vPC Role switch. You can use the vpc role preempt command to switch vPC role between peers.
Additional Information
For information about how to configure Hitless vPC Role Change, see Configure the Hitless vPC Role Change.
vPC Terminology
The terminology used in vPCs is as follows:
- 
                                       						
                                       vPC—The combined port channel between the vPC peer devices and the downstream device.
- 
                                       						
                                       vPC peer device—One of a pair of devices that are connected with the special port channel known as the vPC Peer-Link.
- 
                                       						
                                       vPC Peer-Link—The link used to synchronize state between the vPC peer devices. This link must use a 10-Gigabit Ethernet interface at a minimum. Higher-bandwidth interfaces (such as 25-Gigabit Ethernet, 40-Gigabit Ethernet, 100-Gigabit Ethernet, and so on) may also be used.
- 
                                       						
                                       vPC member port—An interface that belongs to a vPC.
- 
                                       						
                                       Host vPC port—A Fabric Extender host interface that belongs to a vPC.
- 
                                       						
                                       vPC domain—This domain includes both vPC peer devices, the vPC peer-keepalive link, and all of the port channels in the vPC connected to the downstream devices. It is also associated to the configuration mode that you must use to assign vPC global parameters.
- 
                                       						
                                       vPC peer-keepalive link—The peer-keepalive link monitors the vitality of a vPC peer Cisco Nexus 9000 Series device. The peer-keepalive link sends configurable, periodic keepalive messages between vPC peer devices. We recommend that you associate a peer-keepalive link to a separate virtual routing and forwarding (VRF) instance that is mapped to a Layer 3 interface in each vPC peer device. If you do not configure a separate VRF, the system uses the management VRF by default. However, if you use the management interfaces for the peer-keepalive link, you must put a management switch connected to both the active and standby management ports on each vPC peer device (see figure). No data or synchronization traffic moves over the vPC peer-keepalive link; the only traffic on this link is a message that indicates that the originating switch is operating and running a vPC.
- 
                                       						
                                       Dual-active—Both vPC peers act as primary. This situation occurs when the peer-keepalive and vPC Peer-Link go down while both peers are still active. In this case, the secondary vPC assumes that the primary vPC is inactive and acts as the primary vPC.
- 
                                       						
