---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-1
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "liability"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1, 64]
sha256: 22a63c58914614813f69cba005e91747ad2d6c9727a93d1bbc7f35eb7ca7591f
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

Configuring IP Unicast Routing
Finding Feature Information
Your software release may not support all the features documented in this module. For the latest caveats and feature information, see Bug Search Tool and the release notes for your platform and software release. To find information about the features documented in this module, and to see a list of the releases in which each feature is supported, see the feature information table at the end of this module.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature Navigator, go to http://www.cisco.com/go/cfn. An account on Cisco.com is not required.
Information About Configuring IP Unicast Routing
This module describes how to configure IP Version 4 (IPv4) unicast routing on the switch.
A switch stack operates and appears as a single router to the rest of the routers in the network. Basic routing functions like static routing are available with . To use advanced routing features and other routing protocols, you must have the IP Services feature set enabled on the standalone switch or on the active switch.
Information About IP Routing
In some network environments, VLANs are associated with individual networks or subnetworks. In an IP network, each subnetwork is mapped to an individual VLAN. Configuring VLANs helps control the size of the broadcast domain and keeps local traffic local. However, network devices in different VLANs cannot communicate with one another without a Layer 3 device (router) to route traffic between the VLAN, referred to as inter-VLAN routing. You configure one or more routers to route traffic to the appropriate destination VLAN.
When Host A in VLAN 10 needs to communicate with Host B in VLAN 10, it sends a packet addressed to that host. Switch A forwards the packet directly to Host B, without sending it to the router.
When Host A sends a packet to Host C in VLAN 20, Switch A forwards the packet to the router, which receives the traffic on the VLAN 10 interface. The router checks the routing table, finds the correct outgoing interface, and forwards the packet on the VLAN 20 interface to Switch B. Switch B receives the packet and forwards it to Host C.
Types of Routing
Routers and Layer 3 switches can route packets in these ways:
- 
		  By using default routing
- 
		  By using preprogrammed static routes for the traffic
- 
		  By dynamically calculating routes by using a routing protocol
Default routing refers to sending traffic with a destination unknown to the router to a default outlet or destination.
Static unicast routing forwards packets from predetermined ports through a single path into and out of a network. Static routing is secure and uses little bandwidth, but does not automatically respond to changes in the network, such as link failures, and therefore, might result in unreachable destinations. As networks grow, static routing becomes a labor-intensive liability.
Switches running the LAN base feature set support 16 user-configured static routes, in addition to any default routes used for the management interface. The LAN base image supports static routing only on SVIs.
Dynamic routing protocols are used by routers to dynamically calculate the best route for forwarding traffic. There are two types of dynamic routing protocols:
- 
		  Routers using distance-vector protocols maintain routing tables with distance values of networked resources, and periodically pass these tables to their neighbors. Distance-vector protocols use one or a series of metrics for calculating the best routes. These protocols are easy to configure and use.
- 
		  Routers using link-state protocols maintain a complex database of network topology, based on the exchange of link-state advertisements (LSAs) between routers. LSAs are triggered by an event in the network, which speeds up the convergence time or time required to respond to these changes. Link-state protocols respond quickly to topology changes, but require greater bandwidth and more resources than distance-vector protocols.
Distance-vector protocols supported by the switch are Routing Information Protocol (RIP), which uses a single distance metric (cost) to determine the best path and Border Gateway Protocol (BGP), which adds a path vector mechanism. The switch also supports the Open Shortest Path First (OSPF) link-state protocol and Enhanced IGRP (EIGRP), which adds some link-state routing features to traditional Interior Gateway Routing Protocol (IGRP) to improve efficiency.
| Note | On a switch or switch stack, the supported protocols are determined by the software running on the active switch. If the active switch is running the IP base feature set, only default routing, static routing and RIP are supported. If the switch is running the LAN base feature set, you can configure 16 static routes on SVIs. All other routing protocols require the IP services feature set. | 
IP Routing and Switch Stacks
A switch stack appears to the network as a single switch, regardless of which switch in the stack is connected to a routing peer.
The active switch performs these functions:
-  
		  It initializes and configures the routing protocols.
-  
		  It sends routing protocol messages and updates to other routers.
-  
		  It processes routing protocol messages and updates received from peer routers.
-  
		  It generates, maintains, and distributes the distributed Cisco Express Forwarding (dCEF) database to all stack members. The routes are programmed on all switches in the stack bases on this database.
-  
		  The MAC address of the active switch is used as the router MAC address for the whole stack, and all outside devices use this address to send IP packets to the stack.
-  
		  All IP packets that require software forwarding or processing go through the CPU of the active switch.
Stack members perform these functions:
-  
		  They act as routing standby switches, ready to take over in case they are elected as the new active switch if the active switch fails.
-  
		  They program the routes into hardware.
If a active switch fails, the stack detects that the active switch is down and elects one of the stack members to be the new active switch. During this period, except for a momentary interruption, the hardware continues to forward packets with no active protocols.
However, even though the switch stack maintains the hardware identification after a failure, the routing protocols on the router neighbors might flap during the brief interruption before the active switch restarts. Routing protocols such as OSPF and EIGRP need to recognize neighbor transitions. The router uses two levels of nonstop forwarding (NSF) to detect a switchover, to continue forwarding network traffic, and to recover route information from peer devices:
-  
		  NSF-aware routers tolerate neighboring router failures. After the neighbor router restarts, an NSF-aware router supplies information about its state and route adjacencies on request.
-  
		  NSF-capable routers support NSF. When they detect a active switch change, they rebuild routing information from NSF-aware or NSF-capable neighbors and do not wait for a restart.
The switch stack supports NSF-capable routing for OSPF and EIGRP.
Upon election, the new active switch performs these functions:
-  
		  It starts generating, receiving, and processing routing updates.
-  
		  It builds routing tables, generates the CEF database, and distributes it to stack members.
-  
		  It uses its MAC address as the router MAC address. To notify its network peers of the new MAC address, it periodically (every few seconds for 5 minutes) sends a gratuitous ARP reply with the new router MAC address. 
 Note
