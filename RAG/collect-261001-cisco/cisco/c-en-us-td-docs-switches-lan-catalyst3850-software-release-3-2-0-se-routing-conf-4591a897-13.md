---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-13
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [688, 750]
sha256: 67447f9cb641ee13345061be7873f0fcdf86e5e993631fa8e3ea4017502bc68b
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

If at least one of the stack peer neighbors is NSF-aware, the stack master receives updates and rebuilds its database. Each NSF-aware neighbor sends an end of table (EOT) marker in the last update packet to mark the end of the table content. The stack master recognizes the convergence when it receives the EOT marker, and it then begins sending updates. When the stack master has received all EOT markers from its neighbors or when the NSF converge timer expires, EIGRP notifies the routing information database (RIB) of convergence and floods its topology table to all NSF-aware peers.
EIGRP Stub Routing
The EIGRP stub routing feature, available in all feature sets, reduces resource utilization by moving routed traffic closer to the end user.
| Note | The IP Base feature set contains EIGRP stub routing capability, which only advertises connected or summary routes from the routing tables to other Devicees in the network. The Device uses EIGRP stub routing at the access layer to eliminate the need for other types of routing advertisements. For enhanced capability and complete EIGRP routing, the Device must be running the IP Base feature set. On a Device running the IP base feature set, if you try to configure multi-VRF-CE and EIGRP stub routing at the same time, the configuration is not allowed. IPv6 EIGRP stub routing is not supported with the IP base feature set. | 
In a network using EIGRP stub routing, the only allowable route for IP traffic to the user is through a Device that is configured with EIGRP stub routing. The Device sends the routed traffic to interfaces that are configured as user interfaces or are connected to other devices.
When using EIGRP stub routing, you need to configure the distribution and remote routers to use EIGRP and to configure only the Device as a stub. Only specified routes are propagated from the Device. The Device responds to all queries for summaries, connected routes, and routing updates.
Any neighbor that receives a packet informing it of the stub status does not query the stub router for any routes, and a router that has a stub peer does not query that peer. The stub router depends on the distribution router to send the proper updates to all peers.
In the figure given below, Device B is configured as an EIGRP stub router. Devicees A and C are connected to the rest of the WAN. Device B advertises connected, static, redistribution, and summary routes to Device A and C. Device B does not advertise any routes learned from Device A (and the reverse).
For more information about EIGRP stub routing, see “Configuring EIGRP Stub Routing” section of the Cisco IOS IP Configuration Guide, Volume 2 of 3: Routing Protocols.
How to Configure EIGRP
To create an EIGRP routing process, you must enable EIGRP and associate networks. EIGRP sends updates to the interfaces in the specified networks. If you do not specify an interface network, it is not advertised in any EIGRP update.
| Note | If you have routers on your network that are configured for IGRP, and you want to change to EIGRP, you must designate transition routers that have both IGRP and EIGRP configured. In these cases, perform Steps 1 through 3 in the next section and also see the “Configuring Split Horizon” section. You must use the same AS number for routes to be automatically redistributed. | 
- Default EIGRP Configuration
- Configuring Basic EIGRP Parameters
- Configuring EIGRP Interfaces
- Configuring EIGRP Route Authentication
Default EIGRP Configuration
| Table 7 Default EIGRP 		Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Auto summary | Disabled. | 
| Default-information | Exterior routes are accepted and default information is passed between EIGRP processes when doing redistribution. | 
| Default metric | Only connected routes and interface static routes can be redistributed without a default metric. The metric includes:  | 
| Distance | Internal distance: 90. External distance: 170. | 
| EIGRP log-neighbor changes | Disabled. No adjacency changes logged. | 
| IP authentication key-chain | No authentication provided. | 
| IP authentication mode | No authentication provided. | 
| IP bandwidth-percent | 50 percent. | 
| IP hello interval | For low-speed nonbroadcast multiaccess (NBMA) networks: 60 seconds; all other networks: 5 seconds. | 
| IP hold-time | For low-speed NBMA networks: 180 seconds; all other networks: 15 seconds. | 
| IP split-horizon | Enabled. | 
| IP summary address | No summary aggregate addresses are predefined. | 
| Metric weights | tos: 0; k1 and k3: 1; k2, k4, and k5: 0 | 
| Network | None specified. | 
| Nonstop Forwarding (NSF) Awareness | Enabled for IPv4 on switches running the IP services feature set. Allows Layer 3 switches to continue forwarding packets from a neighboring NSF-capable router during hardware or software changes. | 
| NSF capability | Disabled. | 
| Offset-list | Disabled. | 
| Router EIGRP | Disabled. | 
| Set metric | No metric set in the route map. | 
| Traffic-share | Distributed proportionately to the ratios of the metrics. | 
| Variance | 1 (equal-cost load-balancing). | 
| Note | The Device supports EIGRP NSF-capable routing for IPv4. | 
Configuring Basic EIGRP Parameters
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router 				  eigrp autonomous-system Example:  Device(config)# router eigrp 10  | Enables an EIGRP routing process, and enter router configuration mode. The AS number identifies the routes to other EIGRP routers and is used to tag routing information. | 
| Step 3 | nsf Example:  Device(config)# nsf  | (Optional) Enables EIGRP NSF. Enter this command on the stack master and on all of its peers. | 
| Step 4 | network network-number Example:  Device(config)# network 192.168.0.0  | Associate networks with an EIGRP routing process. EIGRP sends updates to the interfaces in the specified networks. | 
| Step 5 | eigrp 				  log-neighbor-changes Example:  Device(config)# eigrp log-neighbor-changes  | (Optional) Enables logging of EIGRP neighbor changes to monitor routing system stability. | 
| Step 6 | metric 				  weights  				tos k1 k2 k3 k4 				  k5 Example:  Device(config)# metric weights 0 2 0 2 0 0  | (Optional) Adjust the EIGRP metric. Although the defaults have been carefully set to provide excellent operation in most networks, you can adjust them. | 
| Step 7 | offset-list [access-list 				  number \|  				name] {in \|  				out}  				offset [type number] Example:  Device(config)# offset-list 21 out 10  | (Optional) Applies an offset list to routing metrics to increase incoming and outgoing metrics to routes learned through EIGRP. You can limit the offset list with an access list or an interface. | 
| Step 8 | auto-summary Example:  Device(config)# auto-summary  | (Optional) Enables automatic summarization of subnet routes into network-level routes. | 
| Step 9 | ip 				  summary-address eigrp  				autonomous-system-number 				  address mask Example:  Device(config)# ip summary-address eigrp 1 192.168.0.0 255.255.0.0  | (Optional) Configures a summary aggregate. | 
| Step 10 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 11 | show ip 				  protocols Example:  Device# show ip protocols  | Verifies your entries. For NSF awareness, the output shows: *** IP Routing is NSF aware *** EIGRP NSF enabled | 
| Step 12 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Caution | Setting metrics is complex and is not recommended without guidance from an experienced network designer. | 
Configuring EIGRP Interfaces
Other optional EIGRP parameters can be configured on an interface basis.
|  | Command or Action | Purpose | 
|---|---|---|
