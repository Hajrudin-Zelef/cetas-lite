---
id: collect-261001-huawei/huawei/info-finder-encyclopedia-en-router-html-8795e4b0-1
title: "info-finder-encyclopedia-en-router-html-8795e4b0"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/info-finder-encyclopedia-en-router-html-8795e4b0.md
source_anchor: ""
source_lines: [1, 51]
sha256: 770b4d5e2d66962284031f20fb152da8c30db3f48ce6fe4f9960ccc8420fb5f3
---

# info-finder-encyclopedia-en-router-html-8795e4b0

What Is a Router?
A router is a hardware device that connects two or more networks. It works at the network layer and serves as a gateway between networks. After receiving a packet, it searches its routing table to determine the best transmission path and then forwards the packet through the corresponding interface to the destination address.
Why Do We Need a Router?
A router is an indispensable network device that mainly provides the following functions:
- Communication between networks: A router has multiple types of interfaces and supports a variety of protocols, enabling communication between different networks, such as LANs and WANs.
- High-performance data forwarding: A router can determine the best paths for received packets to travel, improving communication efficiency and saving network resources.
- Network management: A router provides functions such as configuration management, performance management, fault tolerance management, and traffic control.
How Does a Router Work?
The following figure illustrates how a router forwards a data packet on a network. This process is similar to how an express company delivers a package. Just like package delivery, the data packet does not arrive at its final destination instantly. Instead, it is sorted and relayed through different sub-stations (or routers) along the way before reaching its ultimate destination. A router is located between different networks, but it is not necessarily the final receiver of a packet. That is why a router usually maintains a routing table. A router searches for the next forwarding address based on the destination address in the data packet. This process is called addressing. Before learning the addressing process, we need to understand two basic concepts: routing table and routing protocol.
How a router works on a network
Routing Table
A routing table is a data table stored on each router. It is like a map, on which you can search for the path to a specific destination network address. A routing table has many entries, each indicating a route that corresponds to a destination on a network. The following figure shows a routing table.
When a router receives an IP packet, it checks the destination address in the packet header and searches its routing table for an entry matching the destination address. If a matching entry is found, the router forwards the packet based on the outbound interface and next hop specified in the entry. If no matching entry is found, the router checks whether a default route exists; if no default route exists, the router discards the IP packet. At the same time, the router sends an Internet Control Message Protocol (ICMP) error message to the source device of the packet, reporting that the destination address or network is unreachable.
Each router stores one local core routing table, which stores the routes discovered by various routing protocols. Among these routes, the optimal route is selected. The routing table may contain the following types of routes:
- Direct route: indicates a route to the network segment where a local interface of the router resides. Such routes are automatically discovered and added to the routing table by the router.
- Static route: indicates a route manually configured by the network administrator.
- Dynamic route: indicates a route discovered by the router through a dynamic routing protocol, such as Border Gateway Protocol (BGP), Intermediate System to Intermediate System (IS-IS), Open Shortest Path First (OSPF), or Routing Information Protocol (RIP).
Routes can also be classified into the following two types by the destination address type:
- Host route: indicates a route destined for a host address. Through this route, data packets are sent only to this host address.
- Network segment route: indicates a route destined for a network segment address. Through this route, data packets are sent to all hosts on the network segment.
On NetEngine series routers, you can run the display ip routing-table command to view information about the routing table. For example:
<HUAWEI> display ip routing-table
Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
------------------------------------------------------------------------------
Routing Table: Public
         Destinations : 8        Routes : 8
Destination/Mask     Proto  Pre  Cost   Flags  NextHop        Interface
       0.0.0.0/0      Static 60   0      D     10.1.4.2       GigabitEthernet1/0/0
       10.1.4.0/30    OSPF   10   0      D     10.1.4.1       GigabitEthernet1/0/0
       10.1.4.1/32    Direct 0    0      D     127.0.0.1      InLoopBack0
       10.1.4.2/32    OSPF   10   0      D     10.1.4.2       GigabitEthernet1/0/0
     127.0.0.0/8      Direct 0    0      D     127.0.0.1      InLoopBack0
     127.0.0.1/32     Direct 0    0      D     127.0.0.1      InLoopBack0
127.255.255.255/32    Direct 0    0      D     127.0.0.1      InLoopBack0
255.255.255.255/32    Direct 0    0      D     127.0.0.1      InLoopBack0
The routing table contains the following key information:
- Destination: indicates the destination address or destination network segment of the route.
- Mask: indicates the network mask. It is used together with the destination address (through an AND operation) to identify the address of the network segment where a destination host or router resides. For example, if the destination address is 1.1.1.1 and the mask is 255.255.255.0, the address of the network segment where the host or router resides is 1.1.1.0.
- Proto: indicates the method by which the route is obtained. For example, Direct indicates a direct route, and Static indicates a static route.
- Pre: indicates the preference of the route in the IP routing table. Multiple routes to the same destination may exist and have different next hops or outbound interfaces. These routes may be discovered by different routing protocols or manually configured static routes. Among these routes, the one with the highest priority (smallest preference value) is selected as the optimal route.
- Cost: indicates the cost of the route. When multiple routes to the same destination have the same preference, the route with the smallest cost is selected as the optimal route. The route cost is a value used to evaluate the quality of a route path. It is also called the route metric and indicates the cost of the route path from the source node to the destination node. The cost depends on multiple factors such as the network topology, bandwidth, delay, reliability, and security.
- Flags: indicates the source and destination of a data packet and how the router processes the packet.
- Nexthop: indicates the IP address of the next hop.
- Interface: indicates the outbound interface through which IP packets are forwarded.
Routing Protocols
Routing protocols are rules used by routers to maintain routing tables, such as discovering routes, generating routing entries, and guiding packet forwarding.
Routers support not only static routes, but also dynamic routes. Dynamic routing protocols have their own routing algorithms, which enable them to automatically adapt to network topology changes. They are particularly suitable for networks with a large number of Layer 3 devices. Generally, these protocols can be classified into the following types:
- Interior Gateway Protocol (IGP): routing protocol used within an AS. This protocol is irrelevant to routing protocol selection by other ASs on the Internet. Common IGPs include RIP, OSPF, and IS-IS.
- Exterior Gateway Protocol (EGP): protocol used to transmit data packets across the boundary of an AS when the source and destination devices are located in different ASs. Currently, BGP is the most widely used EGP.
The following figure shows the differences between static and dynamic routes, along with their respective application scenarios.
Differences between static and dynamic routes, along with their respective application scenarios
