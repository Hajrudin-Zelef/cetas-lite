---
id: collect-261001-general-networking/general-networking/manual-dynamic-routing-html-909fa732-2
title: "manual-dynamic-routing-html-909fa732"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-general-networking/manual-dynamic-routing-html-909fa732.md
source_anchor: ""
source_lines: [76, 143]
sha256: 48069d560d382b6efe7aaf436aba5e69331029cdd111bd76ba6204dc22cd3757
---

# manual-dynamic-routing-html-909fa732

Any route received with dynamic routing protocols will only be installed if no similar local route already exists. E.g., if a neighbor advertises a default gateway route, but a directly attached default gateway route already exists, the local route will be preferred and the advertised route will be discarded.
RIP (Routing Information Protocol) - legacy
| Options | Description | 
|---|---|
| enable | This will activate the RIP service. | 
| Version | Choose your RIP version (1 or 2). 1 is classful, 2 supports CIDR. | 
| Passive Interfaces | Select the interfaces, where no RIP packets should be sent to, (e.g., WAN interface). | 
| Route Redistribution | Select other routing sources, which should be redistributed to the other nodes. A good choice is Connected Routes to automatically redistribute all locally attached routes to other routers with RIP. Otherwise use the Networks option to manually insert networks to distribute. | 
| Networks | Enter your networks in CIDR notation like 127.0.0.0/8 . | 
| Default Metric | Set the default metric to a value between 1 and 16. Routes with lower metrics will be preferred, while higher metrics indicate less preferred or distant paths. | 
The Routing Information Protocol (RIP) is a basic distance-vector routing protocol that determines the best path to a network destination based on hop count. With a maximum limit of 15 hops, RIP is suitable only for smaller networks. To prevent routing loops, RIP employs techniques like split horizon, route poisoning, and holddown timers. While easy to configure, RIP has slow convergence and limited scalability, making it less popular in modern networks compared to more efficient protocols like OSPF. It should be considered a legacy protocol.
OSPF/OSPFv3 (Open Shortest Path First)
| Options | Description | 
|---|---|
| Enable | This will activate the OSPF service. | 
| CARP demote | Register CARP status monitor. When no neighbors are found, consider this node less attractive. Requires syslog enabled with “Debugging” logging. Incompatible with “Enable CARP Failover”. | 
| Router ID | (OSPF) If you have a CARP setup, you may want to configure a router id in case of a conflict. (OSPFv3) Router ID as an IPv4 Address to uniquely identify the router. | 
| Reference Cost | (OSPF only) Adjust the reference cost in Mbps for path calculation, useful when bundling interfaces for higher bandwidth. | 
| Passive Interfaces | Select the interfaces where no OSPF packets should be sent. | 
| Redistribution Map | Route Map to set for Redistribution, can be used to send a specific network as advertisement when it is defined in a Prefix List attached to a Route Map. | 
| Log Adjacency Changes | If it should be logged when the topology of the area changes. | 
| Advertise Default Gateway | This will send the information that we have a default gateway. | 
| Always Advertise Default Gateway | Always sends default gateway information, regardless of availability. | 
| Advertise Default Gateway Metric | Allows manipulation of the metric when advertising the default gateway. | 
| Route Redistribution | Select other routing sources to redistribute to other nodes. Can be combined with a Route Map per redistribution. | 
| Options | Description | 
|---|---|
| Enable | (OSPF only) Enable / Disable | 
| Description | (OSPF only) Optional description for the neighbor. | 
| Peer-IP | (OSPF only) Specify the IP address of the OSPF neighbor. | 
| Poll-Interval | (OSPF only) The poll-interval specifies the rate for sending hello packets to neighbors that are not active. When the configured neighbor is discovered, hello packets will be sent at the rate of the hello-interval. The default poll-interval is 60 seconds. | 
| Priority | (OSPF only) The priority is used to for the Designated Router (DR) election on non-broadcast multi-access networks. | 
| Options | Description | 
|---|---|
| Enabled | (OSPF only) Enable / Disable | 
| Area ID | (OSPF only) Enter area ID in dotted (e.g. 0.0.0.1) format. You only need to define areas that are not normal. All areas defined in the network or interface tab will automatically be normal, unless explicitly overwritten here with a different area type. | 
| Area Type | (OSPF only) Select area behavior (e.g. stub no-summary) | 
| Options | Description | 
|---|---|
| Enabled | Enable / Disable | 
| Network Address | (OSPF) Specifies the network address (e.g., 192.168.1.0) to include in OSPF. (OSPFv3) Specifies the IPv6 network address (e.g., fe80::1234) to include in OSPFv3. | 
| Network Mask | (OSPF) Defines the network mask (e.g., 24) for the specified network. (OSPFv3) Defines the network prefix length (e.g., 64) for the specified IPv6 address range. | 
| Area | Assigns the network to an OSPF area using an identifier like 0.0.0.0 (Backbone Area). The Backbone Area connects other areas, supporting inter-area communication, while additional areas (e.g., 0.0.0.1, 0.0.0.255) segment the network logically to limit routing updates. | 
| Area Range | Summarizes multiple networks in the specified area, consolidating multiple networks into a single summarized route (OSPF) 192.168.0.0/23, (OSPFv3) fe80:1234::/64. | 
| Prefix-List In | Filters inbound route advertisements using a prefix list. | 
| Prefix-List Out | Filters outbound route advertisements using a prefix list. | 
Note
Using a Network configuration with 0.0.0.0 as the Backbone Area is beneficial in larger networks, where defining broad network ranges streamlines OSPF configuration. This approach avoids the need to configure OSPF individually on each interface by including all subnets within the specified range. The Backbone Area serves as the primary route aggregation point, allowing inter-area communication which is essential in hierarchical OSPF networks. Networks and Interfaces cannot have the same Area, only one of them can be defined in the Backbone Area.
| Options | Description | 
|---|---|
| Enabled | Enable / Disable | 
| Interface | Select an interface where these settings apply. | 
| Authentication Type | (OSPF only) Defines security method for OSPF exchanges (None, plain, or MD5) to prevent unauthorized updates. | 
| Authentication Key | (OSPF only) Specifies a password or key used for plain or MD5 authentication. | 
| Authentication Key ID | (OSPF only) Numeric identifier for MD5 authentication, ensuring correct key selection. | 
| Area | Assigns the network to an OSPF area using an identifier like 0.0.0.0 (Backbone Area). The Backbone Area connects other areas, supporting inter-area communication, while additional areas (e.g., 0.0.0.1, 0.0.0.255) segment the network logically to limit routing updates. | 
| Passive Interface | (OSPFv3 only) Disables OSPF Hello packets on the interface, preventing neighbor relationships (used for security or optimization). | 
| Cost | Sets the OSPF metric for path selection; lower costs are preferred paths within the area. | 
| Cost (when demoted) | Specifies metric cost when interface is in backup mode via CARP, deprioritizing paths dynamically. | 
| Depend on (carp) | Links the interface cost to a CARP VHID, adjusting costs based on primary or backup status. | 
| Hello Interval | Sets frequency (in seconds) of Hello packets to maintain OSPF neighbor relationships. | 
| Dead Interval | Defines the timeout period for OSPF neighbors; after this period, the neighbor is marked as down. | 
| Retransmission Interval | Time (seconds) to wait before resending Link-State Advertisements (LSAs) if acknowledgment is delayed. | 
| Retransmission Delay | Configures the hold time before LSAs are resent, accommodating slow or high-latency links. | 
| Priority | Determines the likelihood of becoming a Designated Router; higher values increase priority. | 
| BFD | Activates Bidirectional Forwarding Detection for rapid link failure detection; peer configuration required. | 
| Network Type |  | 
Note
