---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9.md
source_anchor: ""
source_lines: [1, 34]
sha256: 36658b3b7e60c26c9a395929721a87d0c7008194d15fbba425f307eb41bfe40d
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-8-configuration-gu-977251c9

Distributed Cisco Express Forwarding and Load-Balancing Scheme for CEF Traffic
The following sections provide information about distributed Cisco express forwarding (CEF) and load-balancing scheme for CEF traffic.
Restrictions for Configuring a Load-Balancing Scheme for CEF Traffic
- 
                                 
                                 You must globally configure load balancing on device or device stack members in the same way.
- 
                                 
                                 Per-packet load balancing for CEF traffic is not supported.
Information About Cisco Express Forwarding
Cisco Express Forwarding (CEF) is a Layer 3 IP switching technology used to optimize network performance. CEF implements an advanced IP look-up and forwarding algorithm to deliver maximum Layer 3 switching performance. CEF is less CPU-intensive than fast switching route caching, allowing more CPU processing power to be dedicated to packet forwarding. In a switch stack, the hardware uses distributed CEF (dCEF) in the stack. In dynamic networks, fast switching cache entries are frequently invalidated because of routing changes, which can cause traffic to be process switched using the routing table, instead of fast switched using the route cache. CEF and dCEF use the Forwarding Information Base (FIB) lookup table to perform destination-based switching of IP packets.
The two main components in CEF and dCEF are the distributed FIB and the distributed adjacency tables.
- 
                                 
                                 The FIB is similar to a routing table or information base and maintains a mirror image of the forwarding information in the IP routing table. When routing or topology changes occur in the network, the IP routing table is updated, and those changes are reflected in the FIB. The FIB maintains next-hop address information based on the information in the IP routing table. Because the FIB contains all known routes that exist in the routing table, CEF eliminates route cache maintenance, is more efficient for switching traffic, and is not affected by traffic patterns.
- 
                                 
                                 Nodes in the network are said to be adjacent if they can reach each other with a single hop across a link layer. CEF uses adjacency tables to prepend Layer 2 addressing information. The adjacency table maintains Layer 2 next-hop addresses for all FIB entries.
Because the switch or switch stack uses Application Specific Integrated Circuits (ASICs) to achieve Gigabit-speed line rate IP traffic, CEF or dCEF forwarding applies only to the software-forwarding path, that is, traffic that is forwarded by the CPU.
CEF Load-Balancing Overview
CEF load balancing allows you to optimize resources by distributing traffic over multiple paths. CEF load balancing works based on a combination of source and destination packet information.
You can configure load balancing on a per-destination. Because load-balancing decisions are made on the outbound interface, load balancing must be configured on the outbound interface.
Per-Destination Load Balancing for CEF Traffic
Per-destination load balancing allows the device to use multiple paths to achieve load sharing across multiple source-destination host pairs. Packets for a given source-destination host pair are guaranteed to take the same path, even if multiple paths are available. Traffic streams destined for different pairs tend to take different paths.
Per-destination load balancing is enabled by default when you enable CEF. To use per-destination load balancing, you do not perform any additional tasks once CEF is enabled. Per-destination is the load-balancing method of choice for most situations.
Because per-destination load balancing depends on the statistical distribution of traffic, load sharing becomes more effective as the number of source-destination host pairs increases.
You can use per-destination load balancing to ensure that packets for a given host pair arrive in order. All packets intended for a certain host pair are routed over the same link (or links).
Load-Balancing Algorithms for CEF Traffic
The following load-balancing algorithms are provided for use with CEF traffic. Select a load-balancing algorithm with the ip cef load-sharing algorithm command.
- 
                                    
                                    Original algorithm—The original load-balancing algorithm produces distortions in load sharing across multiple devices because the same algorithm was used on every device. Depending on your network environment, you should select the algorithm.
- 
                                    
