---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-configuration-guides-layer-3-switchin-e1d96fd7
title: "switching-ms-switches-design-and-configure-configuration-guides-layer-3-switchin-e1d96fd7"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "training"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-configuration-guides-layer-3-switchin-e1d96fd7.md
source_anchor: ""
source_lines: [1, 71]
sha256: a302710a52f79a8fe4aae634637bafcf4f90358344e2ae7396d61cd14c9686dc
---

# switching-ms-switches-design-and-configure-configuration-guides-layer-3-switchin-e1d96fd7

MS OSPF Overview
Learn more with this free online training course on the Meraki Learning Hub:
MS OSPF
Cisco Meraki layer-3 MS switches support the use of the OSPF routing protocol to advertise its subnets to neighboring OSPF-capable layer 3 devices. OSPF may be desirable in more complex network topologies with a layered switch distribution, where static routes are not ideal. This article outlines the OSPF implementation and configuration options available on the Cisco Meraki MS platform, and walks through an example packet capture for reference purposes.
OSPF(v2) on the MS series uses RFC 2328 with cost metric calculations using RFC 1583. MS switches also support Equal-cost Multipath (ECMP) when the routes are learned via OSPF. For the multipath routes to be used by the MS switches, the next route needs to be learned by the same type of OSPF route (Inter-area/Intra-area/External type 1, etc.). A mix of next hops spanned over these different route types won't be used as multipath.
See below on how to enable and configure OSPF on Meraki MS switches supporting L3 routing.
Supported Models: MS250, MS300 series, MS400 series. Also supported Cloud managed Catalyst, all C9300/X/L models including MS390.
OSPF and Warm Spare do not operate concurrently on a switch.
Dashboard Configuration
The following sections describe the different configuration options available in Dashboard under Switches > Configure > OSPF Routing.
Areas
Area types in OSPF are used to define what kinds of Link State Advertisements (LSAs) will be found within an area, and determine how the route table will be generated in each area.
MS switches support 3 area types:
- Normal Areas
- Stub Areas
- Not-So-Stubby Areas (NSSA).
Each area in Dashboard requires an Area ID, a descriptive name, and the type of area. If configuring an MS to be part of an existing OSPF Autonomous System and/or Area, be sure to reference the existing Area IDs:
Normal Area
Normal Areas allow for the creation of a full link state database on all routers in the area. This database allows all routers in the area to know of all routes in the Autonomous System (AS). Normal areas are generally acceptable unless the network utilizes a router that cannot run recalculations without slowing itself down.
Normal areas can contain LSA types 1,2,3,4 and 5.
Stub Area
Stub areas are ideal for branch locations where not every route needs to be advertised, so a default route to the core would suffice. Stub areas allow the L3 switch to save resources and bandwidth by cutting down recalculations and the number of LSAs going over the wire.
Stub areas can contain LSA types 1,2, and 3.
Not-So-Stubby Areas (NSSA)
Not-so-stubby areas are similar to Stub Areas, with the caveat that they allow external routes to be introduced to them from a Not-so-stubby Area Border Router (ABR). In this scenario, the MS can inject outside routes into the NSSA which will then pass them onto the ABR. As Type 5 LSAs are not allowed to be in any sort of stub networks, NSSAs use Type 7 LSAs, which are functionally similar to Type 5 LSAs. Once they hit an ABR, the ABR converts it to a Type 5 and sends it out as necessary.
Not-so-stubby Areas support LSA types 1,2 and 7.
Configuring Interfaces
OSPF configuration is handled on a per-interface basis, to determine what networks will be advertised in which areas (if at all).
Catalyst/MS switches only support broadcast/multicast-based peering. Point-to-point/multipoint network types are not currently supported.
- Area
The OSPF area to which this interface should belong.
- Cost
The path cost for this interface. Defaults to 1, but can be increased to give lower priority.
- Passive
    
  - Yes: The interface will not send hellos, form adjacency or learn routes.
  - No: The interface will send hellos, form adjacency and learn routes.
Static Routes
Dashboard provides the ability to pick and choose which static routes should be redistributed into the OSPF domain. This is done by selecting the route(s) and configuring Advertise via OSPF to “Yes,” then choosing the relative priority as needed.
Router ID
Router IDs for Cisco Meraki Switches (MS) can be configured by going to Switch > Monitor > Switches > (pick switch that is participating in OSPF) > L3 Routing Tab:
Note: You can use the pencil icon to edit the router ID.
For stacks, the router ID can be found under Switch > Monitor > Switch Stacks > (pick the switch stack that is particpating in OSPF) > Layer 3 routing
Timers
Note: The value configured for timers must be identical between all participating OSPF neighbors. If introducing an MS switch to an existing OSPF topology, be sure to reference the existing configuration.
There are two timers used in OSPF, as follows:
- Hello Timer
Denotes the frequency at which the MS switches will send hello packets out to OSPF neighbors to maintain connectivity.
- Dead Timer
The value used to determine when peers will be declared as “dead” or no longer active.
Authentication
If enabled, this allows the use of MD5 Authentication for the OSPF instance, which can be used to help secure the network by preventing attackers from learning about the topology through OSPF.
If MD5 Authentication is enabled, you will be prompted for the authentication key to be used (ID and password):
Example Packet Capture Breakdown
Below is a breakdown of a packet capture with a Cisco Router and an MS-320P forming an OSPF adjacency, which outlines how OSPF functions in practice.
This capture can be broken down into the following processes:
- Learning about neighbors from Hello packets.
- Syncing OSPF databases with LSA Updates.
- Keeping neighbors alive with Hello packets.
Example Topology
Learning about Neighbors from Hello Packets
As OSPF is a dynamic routing protocol, neighbors need to be able to dynamically learn about other devices on the network that they can create adjacencies with. With OSPF, this is done by sending OSPF Hello Packets to the OSPF Multicast Address of 224.0.0.5. This mechanism is also used to detect dead peers in an OSPF area.
In the image above, we can see 10.0.10.243 (the MS) sending hello packets every 10 seconds, as per its configured hello timer interval. Right before packet 1449, OSPF was enabled on the Cisco ISR, which in turn caused the ISR to start sending hello messages itself. Within these hello messages, there are 4 fields that need to match to ensure an adjacency can be formed: Area ID, Auth Type, Hello Interval and Dead Interval:
Syncing OSPF Databases with LSA Updates
The next step in forming an adjacency is syncing OSPF Databases and exchanging LSA Updates, by setting up a poll/response (primary/secondary) relationship between neighbors and exchanging information between the two until everything is synced. The image below shows a typical expected packet exchange between an MS and Cisco ISR:
We can look deeper into the LS Update packets to see the LSA Type being sent, as well as the data (networks) to go along with it:
Keeping Neighbors Alive with Hello Packets
Once an adjacency has established, OSPF peers will utilize OSPF Hello messages again to keep the adjacency alive, as seen below:
If a device fails to hear a hello from an adjacent for the Dead Timer interval (40 seconds, or 4 missed Hellos in the Meraki default configuration), it will mark the peer as dead.
Additional Resources
For additional information regarding MS layer-3 capabilities, please refer to the following articles:
