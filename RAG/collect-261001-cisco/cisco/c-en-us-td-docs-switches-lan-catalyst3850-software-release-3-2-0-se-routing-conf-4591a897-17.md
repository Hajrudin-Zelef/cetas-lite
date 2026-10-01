---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-17
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [875, 931]
sha256: b46537e779b3df2464de78e5cc6f883759f432cff4463153136bb2c2fe538165
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

Based on the community, you can control which routing information to accept, prefer, or distribute to other neighbors. A BGP speaker can set, append, or modify the community of a route when learning, advertising, or redistributing routes. When routes are aggregated, the resulting aggregate has a COMMUNITIES attribute that contains all communities from all the initial routes.
You can use community lists to create groups of communities to use in a match clause of a route map. As with an access list, a series of community lists can be created. Statements are checked until a match is found. As soon as one statement is satisfied, the test is concluded.
To set the COMMUNITIES attribute and match clauses based on communities, see the match community-list and set community route-map configuration commands in the “Using Route Maps to Redistribute Routing Information” section.
BGP Neighbors and Peer Groups
Often many BGP neighbors are configured with the same update policies (that is, the same outbound route maps, distribute lists, filter lists, update source, and so on). Neighbors with the same update policies can be grouped into peer groups to simplify configuration and to make updating more efficient. When you have configured many peers, we recommend this approach.
To configure a BGP peer group, you create the peer group, assign options to the peer group, and add neighbors as peer group members. You configure the peer group by using the neighbor router configuration commands. By default, peer group members inherit all the configuration options of the peer group, including the remote-as (if configured), version, update-source, out-route-map, out-filter-list, out-dist-list, minimum-advertisement-interval, and next-hop-self. All peer group members also inherit changes made to the peer group. Members can also be configured to override the options that do not affect outbound updates.
Aggregate Routes
Classless interdomain routing (CIDR) enables you to create aggregate routes (or supernets) to minimize the size of routing tables. You can configure aggregate routes in BGP either by redistributing an aggregate route into BGP or by creating an aggregate entry in the BGP routing table. An aggregate address is added to the BGP table when there is at least one more specific entry in the BGP table.
Routing Domain Confederations
One way to reduce the IBGP mesh is to divide an autonomous system into multiple subautonomous systems and to group them into a single confederation that appears as a single autonomous system. Each autonomous system is fully meshed within itself and has a few connections to other autonomous systems in the same confederation. Even though the peers in different autonomous systems have EBGP sessions, they exchange routing information as if they were IBGP peers. Specifically, the next hop, MED, and local preference information is preserved. You can then use a single IGP for all of the autonomous systems.
BGP Route Reflectors
BGP requires that all of the IBGP speakers be fully meshed. When a router receives a route from an external neighbor, it must advertise it to all internal neighbors. To prevent a routing information loop, all IBPG speakers must be connected. The internal neighbors do not send routes learned from internal neighbors to other internal neighbors.
With route reflectors, all IBGP speakers need not be fully meshed because another method is used to pass learned routes to neighbors. When you configure an internal BGP peer to be a route reflector, it is responsible for passing IBGP learned routes to a set of IBGP neighbors. The internal peers of the route reflector are divided into two groups: client peers and nonclient peers (all the other routers in the autonomous system). A route reflector reflects routes between these two groups. The route reflector and its client peers form a cluster. The nonclient peers must be fully meshed with each other, but the client peers need not be fully meshed. The clients in the cluster do not communicate with IBGP speakers outside their cluster.
When the route reflector receives an advertised route, it takes one of these actions, depending on the neighbor:
- A route from an external BGP speaker is advertised to all clients and nonclient peers.
- A route from a nonclient peer is advertised to all clients.
- A route from a client is advertised to all clients and nonclient peers. Hence, the clients need not be fully meshed.
Usually a cluster of clients have a single route reflector, and the cluster is identified by the route reflector router ID. To increase redundancy and to avoid a single point of failure, a cluster might have more than one route reflector. In this case, all route reflectors in the cluster must be configured with the same 4-byte cluster ID so that a route reflector can recognize updates from route reflectors in the same cluster. All the route reflectors serving a cluster should be fully meshed and should have identical sets of client and nonclient peers.
Route Dampening
Route flap dampening is a BGP feature designed to minimize the propagation of flapping routes across an internetwork. A route is considered to be flapping when it is repeatedly available, then unavailable, then available, then unavailable, and so on. When route dampening is enabled, a numeric penalty value is assigned to a route when it flaps. When a route’s accumulated penalties reach a configurable limit, BGP suppresses advertisements of the route, even if the route is running. The reuse limit is a configurable value that is compared with the penalty. If the penalty is less than the reuse limit, a suppressed route that is up is advertised again.
Dampening is not applied to routes that are learned by IBGP. This policy prevents the IBGP peers from having a higher penalty for routes external to the AS.
More BGP Information
For detailed descriptions of BGP configuration, see the “Configuring BGP” chapter in the “IP Routing Protocols” part of the Cisco IOS IP Configuration Guide, Release 12.4. For details about specific commands, see the Cisco IOS IP Command Reference, Volume 2 of 3: Routing Protocols, Release 12.4.
How to Configure BGP
Default BGP Configuration
The table given below shows the basic default BGP configuration. For the defaults for all characteristics, see the specific commands in the Cisco IOS IP Command Reference, Volume 2 of 3: Routing Protocols, Release 12.4.
| Table 10 Default BGP 		  Configuration |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Aggregate address | Disabled: None defined. | 
| AS path access list | None defined. | 
| Auto summary | Disabled. | 
| Best path |  | 
| BGP community list |  | 
| BGP confederation identifier/peers |  | 
| BGP Fast external fallover | Enabled. | 
| BGP local preference | 100. The range is 0 to 4294967295 with the higher value preferred. | 
| BGP network | None specified; no backdoor route advertised. | 
| BGP route dampening | Disabled by default. When enabled: | 
| BGP router ID | The IP address of a loopback interface if one is configured or the highest IP address configured for a physical interface on the router. | 
| Default information originate (protocol or network redistribution) | Disabled. | 
| Default metric | Built-in, automatic metric translations. | 
| Distance |  | 
| Distribute list |  | 
| Internal route redistribution | Disabled. | 
| IP prefix list | None defined. | 
| Multi exit discriminator (MED) |  | 
| Neighbor |  | 
| NSF1 Awareness | Disabled2. If enabled, allows Layer 3 switches to continue forwarding packets from a neighboring NSF-capable router during hardware or software changes. | 
| Route reflector | None configured. | 
| Synchronization (BGP and IGP) | Disabled. | 
| Table map update | Disabled. | 
| Timers | Keepalive: 60 seconds; holdtime: 180 seconds. | 
Enabling BGP Routing
|  | Command or Action | Purpose | 
|---|---|---|
