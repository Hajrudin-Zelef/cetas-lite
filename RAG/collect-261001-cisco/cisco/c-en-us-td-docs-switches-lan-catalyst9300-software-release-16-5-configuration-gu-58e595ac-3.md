---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac-3
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac.md
source_anchor: ""
source_lines: [130, 172]
sha256: e5b2cbf78aad1d9d3a1ecea7e07f5f753f7eb0b6d2b686c55c07651b16228ea4
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac

| Step 6 | copy running-config startup-config Example:  Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
What to do next
Use the no ip route prefix mask {address| interface} global configuration command to remove a static route. The device retains static routes until you remove them.
Default Routes and Networks
Information About Default Routes and Networks
A router might not be able to learn the routes to all other networks. To provide complete routing capability, you can use some routers as smart routers and give the remaining routers default routes to the smart router. (Smart routers have routing table information for the entire internetwork.) These default routes can be dynamically learned or can be configured in the individual routers. Most dynamic interior routing protocols include a mechanism for causing a smart router to generate dynamic default information that is then forwarded to other routers.
If a router has a directly connected interface to the specified default network, the dynamic routing protocols running on that device generate a default route. In RIP, it advertises the pseudonetwork 0.0.0.0.
A router that is generating the default for a network also might need a default of its own. One way a router can generate its own default is to specify a static route to the network 0.0.0.0 through the appropriate device.
When default information is passed through a dynamic routing protocol, no further configuration is required. The system periodically scans its routing table to choose the optimal default network as its default route. In IGRP networks, there might be several candidate networks for the system default. Cisco routers use administrative distance and metric information to set the default route or the gateway of last resort.
If dynamic default information is not being passed to the system, candidates for the default route are specified with the ip default-network global configuration command. If this network appears in the routing table from any source, it is flagged as a possible choice for the default route. If the router has no interface on the default network, but does have a path to it, the network is considered as a possible candidate, and the gateway to the best default path becomes the gateway of last resort.
How to Configure Default Routes and Networks
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | ip default-network network number Example:  Device(config)# ip default-network 1 | Specifies a default network. | 
| Step 3 | end Example:  Device(config)# end  | Returns to privileged EXEC mode. | 
| Step 4 | show ip route Example:  Device# show ip route  | Displays the selected default route in the gateway of last resort display. | 
| Step 5 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Route Maps to Redistribute Routing Information
Information About Route Maps
The switch can run multiple routing protocols simultaneously, and it can redistribute information from one routing protocol to another. Redistributing information from one routing protocol to another applies to all supported IP-based routing protocols.
You can also conditionally control the redistribution of routes between routing domains by defining enhanced packet filters or route maps between the two domains. The match and set route-map configuration commands define the condition portion of a route map. The match command specifies that a criterion must be matched. The set command specifies an action to be taken if the routing update meets the conditions defined by the match command. Although redistribution is a protocol-independent feature, some of the match and set route-map configuration commands are specific to a particular protocol.
One or more match commands and one or more set commands follow a route-map command. If there are no match commands, everything matches. If there are no set commands, nothing is done, other than the match. Therefore, you need at least one match or set command.
| Note | A route map with no set route-map configuration commands is sent to the CPU, which causes high CPU utilization. | 
You can also identify route-map statements as permit or deny . If the statement is marked as a deny, the packets meeting the match criteria are sent back through the normal forwarding channels (destination-based routing). If the statement is marked as permit, set clauses are applied to packets meeting the match criteria. Packets that do not meet the match criteria are forwarded through the normal routing channel.
How to Configure a Route Map
Although each of Steps 3 through 14 in the following section is optional, you must enter at least one match route-map configuration command and one set route-map configuration command.
| Note | The keywords are the same as defined in the procedure to control the route distribution. | 
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | route-mapmap-tag [permit \| deny] [sequence number] Example:  Device(config)# route-map rip-to-ospf permit 4  | Defines any route maps used to control redistribution and enter route-map configuration mode. map-tag —A meaningful name for the route map. The redistribute router configuration command uses this name to reference this route map. Multiple route maps might share the same map tag name. (Optional) If permit is specified and the match criteria are met for this route map, the route is redistributed as controlled by the set actions. If deny is specified, the route is not redistributed. sequence number (Optional)— Number that indicates the position a new route map is to have in the list of route maps already configured with the same name. | 
| Step 3 | match as-path path-list-number Example:  Device(config-route-map)#match as-path 10  | Matches a BGP AS path access list. | 
| Step 4 | match community-list community-list-number [exact] Example:  Device(config-route-map)# match community-list 150  | Matches a BGP community list. | 
| Step 5 | match ip address {access-list-number \| access-list-name} [ ...access-list-number \| ...access-list-name] Example:  Device(config-route-map)# match ip address 5 80  | Matches a standard access list by specifying the name or number. It can be an integer from 1 to 199. | 
| Step 6 | match metric metric-value Example:  Device(config-route-map)# match metric 2000  | Matches the specified route metric. The metric-value can be an EIGRP metric with a specified value from 0 to 4294967295. | 
| Step 7 | match ip next-hop {access-list-number \| access-list-name} [ ...access-list-number \| ...access-list-name] Example:  Device(config-route-map)# match ip next-hop 8 45  | Matches a next-hop router address passed by one of the access lists specified (numbered from 1 to 199). | 
| Step 8 | match tag tag value [...tag-value] Example:  Device(config-route-map)# match tag 3500  | Matches the specified tag value in a list of one or more route tag values. Each can be an integer from 0 to 4294967295. | 
| Step 9 | match interfacetype number [...type-number] Example:  Device(config-route-map)# match interface gigabitethernet 1/0/1  | Matches the specified next hop route out one of the specified interfaces. | 
| Step 10 | match ip route-source {access-list-number \| access-list-name} [ ...access-list-number \| ...access-list-name] Example:  Device(config-route-map)# match ip route-source 10 30  | Matches the address specified by the specified advertised access lists. | 
| Step 11 | match route-type {local \| internal \| external [type-1 \| type-2]} Example:  Device(config-route-map)# match route-type local  | Matches the specified route-type :  | 
