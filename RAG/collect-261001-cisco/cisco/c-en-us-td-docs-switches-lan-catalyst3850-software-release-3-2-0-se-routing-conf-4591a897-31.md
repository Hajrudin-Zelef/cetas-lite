---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-31
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1747, 1788]
sha256: 124104dbddee71e687e9db9dfc6af140536387600b21e7cdc1d2ed2d6ff57c95
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

Static routes that point to an interface are advertised through RIP, IGRP, and other dynamic routing protocols, whether or not static redistribute router configuration commands were specified for those routing protocols. These static routes are advertised because static routes that point to an interface are considered in the routing table to be connected and hence lose their static nature. However, if you define a static route to an interface that is not one of the networks defined in a network command, no dynamic routing protocols advertise the route unless a redistribute static command is specified for these protocols.
When an interface goes down, all static routes through that interface are removed from the IP routing table. When the software can no longer find a valid next hop for the address specified as the forwarding router's address in a static route, the static route is also removed from the IP routing table.
Configuring Static Unicast Routes
Static unicast routes are user-defined routes that cause packets moving between a source and a destination to take a specified path. Static routes can be important if the router cannot build a route to a particular destination and are useful for specifying a gateway of last resort to which all unroutable packets are sent.
Follow these steps to configure a static route:
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable   | Enables privileged EXEC mode. Enter your password if prompted.  | 
| Step 2 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 3 | ip route  			 prefix mask {address \|  				interface} [distance] Example:  Device(config)# ip route prefix mask gigabitethernet 1/0/4  | Establish a static route. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show ip 				  route Example:  Device# show ip route  | Displays the current state of the routing table to verify the configuration. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Use the no ip route prefix mask {address| interface} global configuration command to remove a static route. The device retains static routes until you remove them.
Default Routes and Networks
Information About Default Routes and Networks
A router might not be able to learn the routes to all other networks. To provide complete routing capability, you can use some routers as smart routers and give the remaining routers default routes to the smart router. (Smart routers have routing table information for the entire internetwork.) These default routes can be dynamically learned or can be configured in the individual routers. Most dynamic interior routing protocols include a mechanism for causing a smart router to generate dynamic default information that is then forwarded to other routers.
If a router has a directly connected interface to the specified default network, the dynamic routing protocols running on that device generate a default route. In RIP, it advertises the pseudonetwork 0.0.0.0.
A router that is generating the default for a network also might need a default of its own. One way a router can generate its own default is to specify a static route to the network 0.0.0.0 through the appropriate device.
When default information is passed through a dynamic routing protocol, no further configuration is required. The system periodically scans its routing table to choose the optimal default network as its default route. In IGRP networks, there might be several candidate networks for the system default. Cisco routers use administrative distance and metric information to set the default route or the gateway of last resort.
If dynamic default information is not being passed to the system, candidates for the default route are specified with the ip default-network global configuration command. If this network appears in the routing table from any source, it is flagged as a possible choice for the default route. If the router has no interface on the default network, but does have a path to it, the network is considered as a possible candidate, and the gateway to the best default path becomes the gateway of last resort.
How to Configure Default Routes and Networks
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure 				  terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | ip 				  default-network  				network number Example:  Device(config)# ip default-network 1  | Specifies a default network. | 
| Step 3 | end Example:  Device(config)# end  | Returns to privileged EXEC mode. | 
| Step 4 | show ip 				  route Example:  Device# show ip route  | Displays the selected default route in the gateway of last resort display. | 
| Step 5 | copy 				  running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Route Maps to Redistribute Routing Information
Information About Route Maps
The switch can run multiple routing protocols simultaneously, and it can redistribute information from one routing protocol to another. Redistributing information from one routing protocol to another applies to all supported IP-based routing protocols.
You can also conditionally control the redistribution of routes between routing domains by defining enhanced packet filters or route maps between the two domains. The match and set route-map configuration commands define the condition portion of a route map. The match command specifies that a criterion must be matched. The set command specifies an action to be taken if the routing update meets the conditions defined by the match command. Although redistribution is a protocol-independent feature, some of the match and set route-map configuration commands are specific to a particular protocol.
One or more match commands and one or more set commands follow a route-map command. If there are no match commands, everything matches. If there are no set commands, nothing is done, other than the match. Therefore, you need at least one match or set command.
| Note | A route map with no set route-map configuration commands is sent to the CPU, which causes high CPU utilization. | 
You can also identify route-map statements as permit or deny. If the statement is marked as a deny, the packets meeting the match criteria are sent back through the normal forwarding channels (destination-based routing). If the statement is marked as permit, set clauses are applied to packets meeting the match criteria. Packets that do not meet the match criteria are forwarded through the normal routing channel.
How to Configure a Route Map
Although each of Steps 3 through 14 in the following section is optional, you must enter at least one match route-map configuration command and one set route-map configuration command.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure 				  terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | route-mapmap-tag 				[permit \|  				deny] [sequence 				  number] Example:  Device(config)# route-map rip-to-ospf permit 4  | Defines any route maps used to control redistribution and enter route-map configuration mode. map-tag—A meaningful name for the route map. The redistribute router configuration command uses this name to reference this route map. Multiple route maps might share the same map tag name. (Optional) If permit is specified and the match criteria are met for this route map, the route is redistributed as controlled by the set actions. If deny is specified, the route is not redistributed. sequence number (Optional)— Number that indicates the position a new route map is to have in the list of route maps already configured with the same name. | 
