---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-25
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1307, 1386]
sha256: 407dfb71a56465494019f97960333b2e5709b532c5db161f1e6620f0c7e1a75f
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 5 | isis hello-multiplier  				multiplier [level-1 \|  				level-2] Example:  Device(config-if)# isis hello-multiplier 5  | (Optional) Specifies the number of IS-IS hello packets a neighbor must miss before the router should declare the adjacency as down. The range is from 3 to 1000. The default is 3. Using a smaller hello-multiplier causes fast convergence, but can result in more routing instability. | 
| Step 6 | isis csnp-interval  				seconds [level-1 \|  				level-2] Example:  Device(config-if)# isis csnp-interval 15  | (Optional) Configures the IS-IS complete sequence number PDU (CSNP) interval for the interface. The range is from 0 to 65535. The default is 10 seconds. | 
| Step 7 | isis 				  retransmit-interval  				seconds Example:  Device(config-if)# isis retransmit-interval 7  | (Optional) Configures the number of seconds between retransmission of IS-IS LSPs for point-to-point links. The value you specify should be an integer greater than the expected round-trip delay between any two routers on the network. The range is from 0 to 65535. The default is 5 seconds. | 
| Step 8 | isis 				  retransmit-throttle-interval  				milliseconds Example:  Device(config-if)# isis retransmit-throttle-interval 4000  | (Optional) Configures the IS-IS LSP retransmission throttle interval, which is the maximum rate (number of milliseconds between packets) at which IS-IS LSPs will be re-sent on point-to-point links. The range is from 0 to 65535. The default is determined by the isis lsp-interval command. | 
| Step 9 | isis priority  				value [level-1 \|  				level-2] Example:  Device(config-if)# isis priority 50  | (Optional) Configures the priority to use for designated router election. The range is from 0 to 127. The default is 64. | 
| Step 10 | isis circuit-type {level-1 \|  				level-1-2 \|  				level-2-only} Example:  Device(config-if)# isis circuit-type level-1-2  | (Optional) Configures the type of adjacency desired for neighbors on the specified interface (specify the interface circuit type).  | 
| Step 11 | isis password  				password [level-1 \|  				level-2] Example:  Device(config-if)# isis password secret  | (Optional) Configures the authentication password for an interface. By default, authentication is disabled. Specifying Level 1 or Level 2 enables the password only for Level 1 or Level 2 routing, respectively. If you do not specify a level, the default is Level 1 and Level 2. | 
| Step 12 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 13 | show clns 				  interface  				interface-id Example:  Device# show clns interface gigabitethernet 1/0/1  | Verifies your entries. | 
| Step 14 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Monitoring and Maintaining ISO IGRP and IS-IS
You can remove all contents of a CLNS cache or remove information for a particular neighbor or route. You can display specific CLNS or IS-IS statistics, such as the contents of routing tables, caches, and databases. You can also display information about specific interfaces, filters, or neighbors.
The following table lists the privileged EXEC commands for clearing and displaying ISO CLNS and IS-IS routing. For explanations of the display fields, see the Cisco IOS Apollo Domain, Banyan VINES, DECnet, ISO CLNS and XNS Command Reference ,use the Cisco IOS command reference master index, or search online.
| Table 13 ISO CLNS and IS-IS Clear and 		  Show Commands |  | 
|---|---|
| Command | Purpose | 
|---|---|
| clear clns cache | Clears and reinitializes the CLNS routing cache. | 
| clear clns es-neighbors | Removes end system (ES) neighbor information from the adjacency database. | 
| clear clns is-neighbors | Removes intermediate system (IS) neighbor information from the adjacency database. | 
| clear clns neighbors | Removes CLNS neighbor information from the adjacency database. | 
| clear clns route | Removes dynamically derived CLNS routing information. | 
| show clns | Displays information about the CLNS network. | 
| show clns cache | Displays the entries in the CLNS routing cache. | 
| show clns es-neighbors | Displays ES neighbor entries, including the associated areas. | 
| show clns filter-expr | Displays filter expressions. | 
| show clns filter-set | Displays filter sets. | 
| show clns interface [interface-id] | Displays the CLNS-specific or ES-IS information about each interface. | 
| show clns neighbor | Displays information about IS-IS neighbors. | 
| show clns protocol | List the protocol-specific information for each IS-IS or ISO IGRP routing process in this router. | 
| show clns route | Displays all the destinations to which this router knows how to route CLNS packets. | 
| show clns traffic | Displays information about the CLNS packets this router has seen. | 
| show ip route isis | Displays the current state of the ISIS IP routing table. | 
| show isis database | Displays the IS-IS link-state database. | 
| show isis routes | Displays the IS-IS Level 1 routing table. | 
| show isis spf-log | Displays a history of the shortest path first (SPF) calculations for IS-IS. | 
| show isis topology | Displays a list of all connected routers in all areas. | 
| show route-map | Displays all route maps configured or only the one specified. | 
| trace clns destination | Discover the paths taken to a specified destination by packets in the network. | 
| which-route {nsap-address \| clns-name} | Displays the routing table in which the specified CLNS destination is found. | 
Configuration Examples for ISO CLNS Routing
Example: Configuring IS-IS Routing
This example shows how to configure three routers to run conventional IS-IS as an IP routing protocol. In conventional IS-IS, all routers act as Level 1 and Level 2 routers (by default).
Device(config)# clns routing
Device(config)# router isis
Device(config-router)# net 49.0001.0000.0000.000a.00
Device(config-router)# exit
Device(config)# interface gigabitethernet1/0/1
Device(config-if)# ip router isis
Device(config-if)# clns router isis
Device(config)# interface gigabitethernet1/0/2
Device(config-if)# ip router isis
Device(config-if)# clns router isis
Device(config-router)# exit
Device(config)# clns routing
Device(config)# router isis
Device(config-router)# net 49.0001.0000.0000.000b.00
Device(config-router)# exit
Device(config)# interface gigabitethernet1/0/1
Device(config-if)# ip router isis
Device(config-if)# clns router isis
Device(config)# interface gigabitethernet1/0/2
Device(config-if)# ip router isis
Device(config-if)# clns router isis
Device(config-router)# exit
Device(config)# clns routing
Device(config)# router isis
Device(config-router)# net 49.0001.0000.0000.000c.00
Device(config-router)# exit
Device(config)# interface gigabitethernet1/0/1
Device(config-if)# ip router isis
Device(config-if)# clns router isis
Device(config)# interface gigabitethernet1/0/2
Device(config-if)# ip router isis
Device(config-if)# clns router isis
Device(config-router)# exit
Information About Multi-VRF CE
Virtual Private Networks (VPNs) provide a secure way for customers to share bandwidth over an ISP backbone network. A VPN is a collection of sites sharing a common routing table. A customer site is connected to the service-provider network by one or more interfaces, and the service provider associates each interface with a VPN routing table, called a VPN routing/forwarding (VRF) table.
The switch supports multiple VPN routing/forwarding (multi-VRF) instances in customer edge (CE) devices (multi-VRF CE) when the it is running the IP services or advanced IP Services feature set. Multi-VRF CE allows a service provider to support two or more VPNs with overlapping IP addresses.
Understanding Multi-VRF CE
