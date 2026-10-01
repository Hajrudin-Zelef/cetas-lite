---
id: collect-261001-huawei/huawei/info-finder-encyclopedia-en-network-gateway-html-f4e09514
title: "info-finder-encyclopedia-en-network-gateway-html-f4e09514"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2026-01-08"]
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/info-finder-encyclopedia-en-network-gateway-html-f4e09514.md
source_anchor: ""
source_lines: [1, 48]
sha256: 621174e63c322f03632efdc49bd902dbed6623b4f3ee4c720096e87a6f790855
---

# info-finder-encyclopedia-en-network-gateway-html-f4e09514

What Is a Network Gateway?
A network gateway serves as an exit and entry point for a local network since all traffic traveling to or from external networks must pass through the gateway. When a user attempts to access a resource not available in the local network, the user-originated data packet is passed to the gateway and then analyzed, translated, and routed to the destination. By doing so, the gateway ensures seamless communication between networks.
Why Do We Need a Network Gateway?
A network gateway bridges local and external networks, responsible for translating data packets to ensure smooth communication between dissimilar networks. It is essential for network interworking, offering functions such as:
- Network connection: A network gateway can connect wide area networks (WANs) and local area networks (LANs). By translating data between distinct protocols, it connects dissimilar networks that are unable to communicate directly.
- Security: By implementing access control and security policies, a network gateway can defend against invalid packet attacks, enhancing network security.
- Network management: A network gateway can manage devices on networks by integrating and processing network data. For example, it can detect the device operating status, manage configurations, and diagnose faults.
How Does a Network Gateway Work?
As illustrated in the figure below, a network gateway is deployed at the edge of a network and manages all data directed internally or externally from that network. When one network wants to communicate with another, the data packet is passed to the network gateway prior to being routed to the destination through the most optimal path. A network gateway can be compared to an express station. When sending a parcel, you only need to provide the express station with the destination address, without considering the transportation process or how many times the parcel needs to be transferred.
Working principle of a network gateway
Forwarding a data packet through a network gateway involves the following steps:
- Receiving a packet from the source network: The gateway receives a request packet from a PC on the source network.
- Sending the packet from the source network: The gateway converts the received packet to a format that can be parsed by the destination network and then sends it to the destination network.
- Receiving a packet from the destination network: The gateway receives a response packet from the destination network.
- Sending the packet from the destination network: The gateway converts the received packet to a format that can be parsed by the PC and then sends it back to the source network.
In general, a network gateway receives, parses, converts, and forwards packets to connect devices across different networks.
How Is a Network Gateway Different from a Router?
A network gateway connects networks by translating protocols, while a router primarily forwards data packets. Their main differences can be summarized as follows:
- Network location 
  
  - Network gateway: operates at the application layer (Layer 7) of the Open Systems Interconnection (OSI) model and is typically located at the edge of a network. It connects dissimilar networks by translating packets between their protocols and supported data formats.
  - Router: operates at the network layer (Layer 3) of the OSI model and is usually located inside a network. It transmits data between networks in the form of packets.
- Function 
  
  - Network gateway: connects networks using different protocols by converting the data transmitted between them. For example, a gateway can link a LAN to the Internet.
  - Router: is mainly responsible for forwarding data packets based on their destination IP addresses along the most optimal routes.
- Deployment scope 
  
  - Network gateway: Generally, a network has only one gateway for connecting the local network to external networks.
  - Router: A network may have multiple routers for routing packets within that network.
What Are the Application Scenarios of Network Gateways?
A network gateway is a network device typically situated at the edge of a network. It facilitates communication between distinct networks by forwarding data between them, and plays a crucial role in many scenarios. Currently, AR routers can function as network gateways in the following typical scenarios:
- Internet access gateway: As shown in the following figure, RouterA, RouterB, and RouterC are egress gateways of the enterprise intranet and provide cost-effective and reliable Internet access through wireless (3G, LTE, and 5G) and wired (GE and FE) interfaces.ARs offer the following functions to meet a range of network requirements: 
  - Provide wired (GE and FE) and wireless (3G, LTE, and 5G) interfaces for connecting enterprise branches to the Internet. In wired access mode, ARs can obtain IP addresses for Internet access through static IP address configuration, DHCP (as DHCP clients), or PPPoE (as PPPoE clients).
  - Support Network Address Translation (NAT) that enables internal users to use private IP addresses for accessing external resources. This overcomes the public address shortage at enterprise branches.
  - Allow internal users' access to network resources through domain names. ARs can function as DNS clients to dynamically obtain IP addresses that map domain names from the DNS server, facilitating network resource access.
 
ARs as enterprise egress gateways for Internet access
- Internet of Things (IoT) gateway: In smart home, smart city, and industrial fields, an IoT gateway functions as a pipe to connect IoT devices and cloud networks, enabling communication between common objects with independent functions. An Edge Computing IoT (EC-IoT) gateway offers diverse IoT interfaces, for example, IP-based Power-line Communication (PLC), radio frequency (RF), RS485, and RS232 interfaces, for connecting to sensors and terminals, enabling massive numbers of terminals to connect to IoT networks.The following figure shows the EC-IoT architecture, which consists of the platform and application layer, network layer, and sensing layer. 
  - Platform and application layer: This layer includes the industry IoT platform and controller. The industry IoT platform performs analysis and computing, generates policies, and displays results based on the received IoT data. The controller uniformly manages a large number of terminals, networks, and data, enabling upper-layer IoT applications.
  - Network layer: This layer provides an extensive range of functions such as network connection establishment and maintenance; data collection and front-end processing; on-demand local analysis and decision-making; and backhaul of required data to the cloud.
  - Sensing layer: This layer is responsible for connections to huge numbers of sensors and smart terminals.
 
ARs as EC-IoT gateways
- Author： Meng Hongyan
- Updated on： 2026-01-08
- Views： 12402
- Average rating： 0
