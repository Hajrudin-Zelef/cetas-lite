---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-bgp-cfg-4094-htm-6614cb13
title: "hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-bgp-cfg-4094-htm-6614cb13"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-bgp-cfg-4094-htm-6614cb13.md
source_anchor: ""
source_lines: [1, 57]
sha256: 0b330045283509a02519a4377e2708b8898dc481451d7ce8ac4415b539eb9ed0
---

# hedex-api-pages-edoc1100277644-aem10221-04-resources-vrp-dc-vrp-bgp-cfg-4094-htm-6614cb13

Configuring BGP load balancing better utilizes network resources and reduces network congestion.
On large networks, there may be multiple valid routes to the same destination. BGP, however, advertises only the optimal route to its peers, which may result in load imbalance.
Either of the following methods can be used to resolve the load imbalance:
Use BGP routing policies to allow traffic to be balanced. For example, use a route-policy to modify the Local_Pref, AS_Path, Origin, or MED attribute of BGP routes to control traffic forwarding paths, helping implement load balancing.
Use multiple paths to implement traffic load balancing. This method requires that multiple equal-cost routes exist and the number of routes allowed to participate in load balancing be set. Load balancing can be implemented globally or for a specified peer or peer group.
You can change load balancing rules through configurations. For example, you can prevent the device from comparing AS_Path attributes or IGP costs. When performing these configurations, ensure that no routing loops will occur.
Locally leaked routes and routes imported between public network and VPN instances do not support load balancing.
Run system-view
The system view is displayed.
Run bgp as-number
The BGP view is displayed.
The BGP-IPv4 unicast address family view is displayed.
Run peer { ipv4-address | group-name } load-balancing [ as-path-ignore | as-path-relax ]
BGP peer or peer group-based load balancing is enabled.
Return to the BGP view.
Return to the system view.
The VPN instance view is displayed.
The VPN instance IPv4 address family view is displayed.
An RD is configured for the VPN instance IPv4 address family.
Return to the VPN instance view.
A VPN instance is specified.
A peer relationship is established with the peer with the specified AS number.
The BGP labeled VPN instance IPv4 address family view is displayed.
The device is enabled to exchange routing information with the specified peer.
Peer or peer group-based load balancing among VPN routes is enabled.
(Optional) Change load balancing rules.
The address family views supported by the preceding commands are different. When running any of the commands, ensure that the command is run in the correct address family view.
Change load balancing rules based on networking. Exercise caution when running the preceding commands.
Run commit
The configuration is committed.
Run ipv4-family unicast
The IPv4 unicast address family view is displayed.
Run maximum load-balancing [ ebgp | ibgp ] number [ ecmp-nexthop-changed ]
The maximum number of BGP equal-cost routes for load balancing is set.
ebgp indicates that load balancing is implemented only among EBGP routes.
ibgp indicates that load balancing is implemented only among IBGP routes.
If neither ebgp nor ibgp is specified, both EBGP and IBGP routes can balance traffic, and the number of EBGP routes for load balancing is the same as the number of IBGP routes for load balancing.
If multiple routes with the same destination address exist on the public network, the system selects the optimal route first. If IBGP routes are optimal, only IBGP routes carry out load balancing. If EBGP routes are optimal, only EBGP routes carry out load balancing. This means that load balancing cannot be implemented using both IBGP and EBGP routes with the same destination address.
Set the maximum number of EBGP and IBGP routes for load balancing.
This configuration is used in a VPN where a CE is dual-homed to two PEs. When the CE resides in the same AS as only one of the PEs, you can set the maximum number of EBGP and IBGP routes for load balancing so that VPN traffic can be balanced among EBGP and IBGP routes.
Run ipv4-family vpn-instance vpn-instance-name
The BGP-VPN instance IPv4 address family view is displayed.
Run maximum load-balancing eibgp number [ ecmp-nexthop-changed ]
The maximum number of EBGP and IBGP routes for load balancing is set.
After the maximum load-balancing eibgp number command is run on a device, the device, by default, changes the next hop of each route to itself before advertising the route to a peer, regardless of whether the route is to be used for load balancing. However, in RR or BGP confederation scenarios, the device does not change the next hop addresses of non-local routes to be advertised to a local address. As a result, besides the routes for load-balancing, those routes that are not supposed to participate in load balancing divert traffic to the device, which overburdens the device. To address this problem, you can set ecmp-nexthop-changed so that the device changes the next hop of only routes that are to be used for load balancing to itself before advertising them to peers.
This configuration is mainly used in an EIBGP load balancing scenario where a CE that is single-homed to a PE accesses a CE that is dual-homed to PEs. As shown in Figure 1, CE2 is dual-homed to PE1 and PE2. CE1 and CE2 are connected to PE1, and CE2 and CE3 are connected to PE2. When CE1 or CE3 needs to access CE2, you can configure load balancing among VPN unicast routes and leaked routes on PE1 and PE2. In this manner, when CE1 or CE3 accesses CE2, load balancing can be implemented through PE1 and PE2, that is, load balancing among VPN routes and leaked routes.
A VPN instance is created, and its view is displayed.
An RD is configured for the VPN instance.
A label distribution mode is configured for the current VPN.
Run load-balancing local-learning cross
Load balancing among VPN unicast routes and leaked routes is configured.
The load-balancing local-learning cross command can be run only after the apply-label { per-nexthop | per-route } pop-go command is run in the corresponding address family view of a VPN instance. If the apply-label { per-nexthop | per-route } pop-go command is not run, routing loops may occur between PEs.
The load-balancing local-learning cross command is mutually exclusive with the segment-routing ipv6 locator evpn, segment-routing ipv6 locator, vxlan vni, and evpn mpls routing-enable commands.
When BGP routes carrying the link bandwidth extended community attribute are available for load balancing and they all recurse to IP routes or tunnels, you can run the load-balancing ucmp command in the BGP-IPv4 unicast address family view or BGP view to implement unequal-cost load balancing among BGP routes based on the link bandwidth extended community attribute. With this function, when there are multiple egress devices to the destination, unequal-cost load balancing can be implemented based on the actual bandwidth capability of each egress device. The methods of configuring the link bandwidth extended community attribute are as follows:
After the configuration is complete, verify it.
Run the display bgp routing-table [ network ] [ mask | mask-length ] [ longer-prefixes ] command to check information about the BGP routing table.
Run the display ip routing-table [ verbose ] command to check information about the IP routing table.
