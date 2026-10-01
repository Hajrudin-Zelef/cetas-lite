---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-30
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1706, 1746]
sha256: 88d8c38370348a4bb04a9deaf717ed4363ad203ac40cb11eea7405b878f79d6c
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

Cisco Express Forwarding (CEF) is a Layer 3 IP switching technology used to optimize network performance. CEF implements an advanced IP look-up and forwarding algorithm to deliver maximum Layer 3 switching performance. CEF is less CPU-intensive than fast switching route caching, allowing more CPU processing power to be dedicated to packet forwarding. In a switch stack, the hardware uses distributed CEF (dCEF) in the stack. In dynamic networks, fast switching cache entries are frequently invalidated because of routing changes, which can cause traffic to be process switched using the routing table, instead of fast switched using the route cache. CEF and dCEF use the Forwarding Information Base (FIB) lookup table to perform destination-based switching of IP packets.
The two main components in CEF and dCEF are the distributed FIB and the distributed adjacency tables.
- The FIB is similar to a routing table or information base and maintains a mirror image of the forwarding information in the IP routing table. When routing or topology changes occur in the network, the IP routing table is updated, and those changes are reflected in the FIB. The FIB maintains next-hop address information based on the information in the IP routing table. Because the FIB contains all known routes that exist in the routing table, CEF eliminates route cache maintenance, is more efficient for switching traffic, and is not affected by traffic patterns.
- Nodes in the network are said to be adjacent if they can reach each other with a single hop across a link layer. CEF uses adjacency tables to prepend Layer 2 addressing information. The adjacency table maintains Layer 2 next-hop addresses for all FIB entries.
Because the switch or switch stack uses Application Specific Integrated Circuits (ASICs) to achieve Gigabit-speed line rate IP traffic, CEF or dCEF forwarding applies only to the software-forwarding path, that is, traffic that is forwarded by the CPU.
How to Configure Cisco Express Forwarding
CEF or distributed CEF is enabled globally by default. If for some reason it is disabled, you can re-enable it by using the ip cef or ip cef distributed global configuration command.
The default configuration is CEF or dCEF enabled on all Layer 3 interfaces. Entering the no ip route-cache cef interface configuration command disables CEF for traffic that is being forwarded by software. This command does not affect the hardware forwarding path. Disabling CEF and using the debug ip packet detail privileged EXEC command can be useful to debug software-forwarded traffic. To enable CEF on an interface for the software-forwarding path, use the ip route-cache cef interface configuration command.
| Caution | Although the no ip route-cache cef interface configuration command to disable CEF on an interface is visible in the CLI, we strongly recommend that you do not disable CEF or dCEF on interfaces except for debugging purposes. | 
To enable CEF or dCEF globally and on an interface for software-forwarded traffic if it has been disabled:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | ip cef Example:  Device(config)# ip cef  | Enables CEF operation on a non-stacking switch. Go to Step 4. | 
| Step 3 | ip cef 							distributed Example:  Device(config)# ip cef distributed  | Enables CEF operation on a active switch. | 
| Step 4 | interface 						interface-id Example:  Device(config)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 5 | ip route-cache 							cef Example:  Device(config-if)# ip route-cache cef  | Enables CEF on the interface for software-forwarded traffic. | 
| Step 6 | end Example:  Device(config-if)# end  | Returns to privileged EXEC mode. | 
| Step 7 | show ip 						cef Example:  Device# show ip cef  | Displays the CEF status on all interfaces. | 
| Step 8 | show cef 							linecard [detail] Example:  Device# show cef linecard detail  | (Optional) Displays CEF-related interface information on a non-stacking switch. | 
| Step 9 | show cef linecard [slot-number] [detail] Example:  Device# show cef linecard 5 detail  | (Optional) Displays CEF-related interface information on a switch by stack member for all switches in the stack or for the specified switch. (Optional) For slot-number, enter the stack member switch number. | 
| Step 10 | show cef interface [interface-id] Example:  Device# show cef interface gigabitethernet 1/0/1  | Displays detailed CEF information for all interfaces or the specified interface. | 
| Step 11 | show adjacency Example:  Device# show adjacency  | Displays CEF adjacency table information. | 
| Step 12 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Number of Equal-Cost Routing Paths
Information About Equal-Cost Routing Paths
When a router has two or more routes to the same network with the same metrics, these routes can be thought of as having an equal cost. The term parallel path is another way to see occurrences of equal-cost routes in a routing table. If a router has two or more equal-cost paths to a network, it can use them concurrently. Parallel paths provide redundancy in case of a circuit failure and also enable a router to load balance packets over the available paths for more efficient use of available bandwidth. Equal-cost routes are supported across switches in a stack.
Even though the router automatically learns about and configures equal-cost routes, you can control the maximum number of parallel paths supported by an IP routing protocol in its routing table. Although the switch software allows a maximum of 32 equal-cost routes, the switch hardware will never use more than 16 paths per route.
How to Configure Equal-Cost Routing Paths
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | router {bgp \| rip \| ospf \| eigrp} Example:  Device(config)# router eigrp  | Enters router configuration mode. | 
| Step 3 | maximum-paths maximum Example:  Device(config-router)# maximum-paths 2  | Sets the maximum number of parallel paths for the protocol routing table. The range is from 1 to 16; the default is 4 for most IP routing protocols, but only 1 for BGP. | 
| Step 4 | end Example:  Device(config-router)# end  | Returns to privileged EXEC mode. | 
| Step 5 | show ip protocols Example:  Device# show ip protocols  | Verifies the setting in the Maximum path field. | 
| Step 6 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Static Unicast Routes
Information About Static Unicast Routes
Static unicast routes are user-defined routes that cause packets moving between a source and a destination to take a specified path. Static routes can be important if the router cannot build a route to a particular destination and are useful for specifying a gateway of last resort to which all unroutable packets are sent.
The switch retains static routes until you remove them. However, you can override static routes with dynamic routing information by assigning administrative distance values. Each dynamic routing protocol has a default administrative distance, as listed in Table 41-16. If you want a static route to be overridden by information from a dynamic routing protocol, set the administrative distance of the static route higher than that of the dynamic protocol.
