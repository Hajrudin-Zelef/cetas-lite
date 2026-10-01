---
id: collect-261001-huawei/huawei/info-finder-encyclopedia-en-router-html-8795e4b0-2
title: "info-finder-encyclopedia-en-router-html-8795e4b0"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-07-11"]
keywords: ["dci", "throughput"]
source: docs/RAG/collect-261001-huawei/info-finder-encyclopedia-en-router-html-8795e4b0.md
source_anchor: ""
source_lines: [52, 79]
sha256: 0830ac8a8ecc0cd2aa3d4b8d077659bbbded729c6f029cd3a2006f9bf3de3b34
---

# info-finder-encyclopedia-en-router-html-8795e4b0

In real-world situations, a router can have both static and dynamic routes configured. By default, static routes take precedence over dynamic ones. If there are two routes available to the same destination — one dynamic route and another static route — the static route is preferred.
Addressing
After receiving a packet, a router parses the destination IP address of the packet, searches the routing table for the best path to the destination network, and processes the packet based on the search result.
As shown in the following figure, PC1 needs to forward a packet to PC2 through two routers. The addressing process is as follows:
Router's addressing process
- PC1 sends a packet to RouterA. After receiving the packet, Router A parses the destination IP address of the packet.
- Router A searches its routing table based on the destination IP address and processes the packet based on the search result. 
   
  - If the destination IP address is found in the routing table, RouterA forwards the packet to RouterB through the corresponding outbound interface.
  - If the destination IP address is not found in the routing table, RouterA checks whether a default route exists. If a default route exists, RouterA forwards the packet to RouterB through the outbound interface for the default route.
  - If the destination IP address is not found in the routing table and no default route exists, RouterA sends an error ICMP message to the source IP address, indicating that the packet cannot be transmitted. At the same time, RouterA discards the packet.
- After receiving the packet, RouterB processes the packet in the same way as RouterA and forwards the packet to PC2, the destination host.
What Are the Different Types of Routers?
In 1984, Leonard Bosack and Sandy Lerner, both from Stanford University, designed the first multi-protocol router, a groundbreaking innovation in networking technology. As technologies continued to evolve, many new routers emerged.
- Routers can be classified into the following types based on application scenarios: 
  
  - Enterprise routers used for enterprise intranets
  - Carrier routers used for carrier networks
  - Home routers used for Internet access from residential locations
- Routers can be classified into the following types based on network layers and functions: 
  
  - Backbone routers: These routers feature large data throughput and play a pivotal role in carrier network interconnection. Backbone routers require high performance and reliability. For example, Huawei NetEngine 5000E series backbone routers, with large capacity, high reliability, green design, and strong intelligence, can be used on enterprise backbone networks or function as MAN core nodes, DCI nodes, and IGWs. They can work as either standalone routers or back-to-back and multi-chassis clusters, facilitating on-demand capacity expansion. These capabilities enable enterprise users to effectively manage the rapid growth of Internet traffic and accommodate future service development.
  - Metro routers: These routers are deployed at the edge of a carrier network to connect enterprise networks to the carrier backbone network. For example, Huawei NetEngine 8000 series metro routers are an example of all-scenario access and aggregation routers. In addition to compact design, large capacity, and powerful switching capabilities, they provide a variety of interfaces for flexible service access. Their redundancy design for key components achieves high reliability for different types of services. In addition, they leverage new technologies such as SRv6, FlexE, and IFIT to support smooth cloud-based evolution of customer networks.
  - Access routers: These routers connect to many terminal systems. They connect enterprise networks to carrier networks. They provide multiple types of wired and wireless interfaces and a variety of service capabilities, such as VPN, security, QoS, and authentication. For example, Huawei NetEngine AR series access routers are an example of next-generation service router gateways that integrate SD-WAN, routing, switching, VPN, and security functions. These routers include the AR5700, AR6700, and AR8000 series.
- Author： Tang Dandan
- Updated on： 2025-07-11
- Views： 20870
- Average rating：
