---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac.md
source_anchor: ""
source_lines: [55, 129]
sha256: 58213c60e3e8854dd558a5384dd7ac7d9a75784b668bac851da6ce56f288e3b4
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac

                                 Universal algorithm—The universal load-balancing algorithm allows each device on the network to make a different load sharing decision for each source-destination address pair, which resolves load-sharing imbalances. The device is set to perform universal load sharing by default.
How to Configure a Load-Balancing for CEF Traffic
The following sections provide information on configuring load-balancing for CEF traffic.
Enabling or Disabling CEF Per-Destination Load Balancing
To enable or disable CEF per-destination load balancing, perform the following procedure:
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device# enable   | Enters global configuration mode. | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | interface interface-id Example:  Device(config-if)# interface gigabitethernet 1/0/1  | Enters interface configuration mode, and specifies the Layer 3 interface to configure. | 
| Step 4 | [no] ip load-sharing per-destination Example:  Device(config-if)# ip load-sharing per-destination  | Enables per-destination load balancing for CEF on the interface. The no ip load-sharing per-destination command disables per-destination load balancing for CEF on the interface. | 
| Step 5 | end Example:  Device(config-if)# end | Exits interface configuration mode and returns to privileged EXEC mode. | 
Selecting a Tunnel Load-Balancing Algorithm for CEF Traffic
Select the tunnel algorithm when your network environment contains only a few source and destination pairs. The device is set to perform universal load sharing by default.
To select a tunnel load-balancing algorithm for CEF traffic, perform the following procedure:
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device# enable   | Enters global configuration mode. | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | ip cef load-sharing algorithm {original \| universal [id] } Example:  Device(config)# ip cef load-sharing algorithm universal | Selects a CEF load-balancing algorithm.  | 
| Step 4 | end Example:  Device(config)# end | Returns to privileged EXEC mode. | 
Configuration Examples for CEF Traffic Load-Balancing
The following sections provide configuration examples for CEF traffic load-balancing.
Example: Enabling or Disabling CEF Per-Destination Load Balancing
Per-destination load balancing is enabled by default when you enable CEF. The following example shows how to disable per-destination load balancing:
Device> enable
Device# configure terminal
Device(config)# interface Ethernet1/0/1
Device(config-if)# no ip load-sharing per-destination
Device(config-if)# end
Number of Equal-Cost Routing Paths
Information About Equal-Cost Routing Paths
When a router has two or more routes to the same network with the same metrics, these routes can be thought of as having an equal cost. The term parallel path is another way to see occurrences of equal-cost routes in a routing table. If a router has two or more equal-cost paths to a network, it can use them concurrently. Parallel paths provide redundancy in case of a circuit failure and also enable a router to load balance packets over the available paths for more efficient use of available bandwidth. Equal-cost routes are supported across switches in a stack.
Even though the router automatically learns about and configures equal-cost routes, you can control the maximum number of parallel paths supported by an IP routing protocol in its routing table. Although the switch software allows a maximum of 32 equal-cost routes, the switch hardware will never use more than 16 paths per route.
How to Configure Equal-Cost Routing Paths
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | router {rip \| ospf \| eigrp} Example:  Device(config)# router eigrp  | Enters router configuration mode. | 
| Step 3 | maximum-paths maximum Example:  Device(config-router)# maximum-paths 2  | Sets the maximum number of parallel paths for the protocol routing table. The range is from 1 to 16; the default is 4 for most IP routing protocols, but only 1 for BGP. | 
| Step 4 | end Example:  Device(config-router)# end  | Returns to privileged EXEC mode. | 
| Step 5 | show ip protocols Example:  Device# show ip protocols  | Verifies the setting in the Maximum path field. | 
| Step 6 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Static Unicast Routes
Information About Static Unicast Routes
Static unicast routes are user-defined routes that cause packets moving between a source and a destination to take a specified path. Static routes can be important if the router cannot build a route to a particular destination and are useful for specifying a gateway of last resort to which all unroutable packets are sent.
The switch retains static routes until you remove them. However, you can override static routes with dynamic routing information by assigning administrative distance values. Each dynamic routing protocol has a default administrative distance, as listed in Table 41-16. If you want a static route to be overridden by information from a dynamic routing protocol, set the administrative distance of the static route higher than that of the dynamic protocol.
| Table 1.  Dynamic Routing Protocol                                     		Default Administrative Distances |  | 
|---|---|
| Route Source | Default Distance | 
|---|---|
| Connected interface | 0 | 
| Static route | 1 | 
| Enhanced IRGP summary route | 5 | 
| Internal Enhanced IGRP | 90 | 
| IGRP | 100 | 
| OSPF | 110 | 
| Internal BGP | 200 | 
| Unknown | 225 | 
Static routes that point to an interface are advertised through RIP, IGRP, and other dynamic routing protocols, whether or not static redistribute router configuration commands were specified for those routing protocols. These static routes are advertised because static routes that point to an interface are considered in the routing table to be connected and hence lose their static nature. However, if you define a static route to an interface that is not one of the networks defined in a network command, no dynamic routing protocols advertise the route unless a redistribute static command is specified for these protocols.
When an interface goes down, all static routes through that interface are removed from the IP routing table. When the software can no longer find a valid next hop for the address specified as the forwarding router's address in a static route, the static route is also removed from the IP routing table.
Configuring Static Unicast Routes
Static unicast routes are user-defined routes that cause packets moving between a source and a destination to take a specified path. Static routes can be important if the router cannot build a route to a particular destination and are useful for specifying a gateway of last resort to which all unroutable packets are sent.
Follow these steps to configure a static route:
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example:  Device> enable   | Enables privileged EXEC mode.  | 
| Step 2 | configure terminal Example:  Device# configure terminal   | Enters global configuration mode. | 
| Step 3 | ip route prefix mask {address \| interface} [distance] Example:  Device(config)# ip route prefix mask gigabitethernet 1/0/4 | Establish a static route. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show ip route Example:  Device# show ip route  | Displays the current state of the routing table to verify the configuration. | 
