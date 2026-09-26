---
id: collect-260926-mikrotik/mikrotik/destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9-2
title: "destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "memory", "training"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9.md
source_anchor: ""
source_lines: [69, 165]
sha256: b64295e69f66dc31da2e125a6a13e1a358962764f971b9887c6a87d867295175
---

# destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9

Note: To allow access, but only from the home PC, we can improve our DST-nat rule with “src-address=192.168.88.1,” which is a public IP address of the home PC for added security. You can choose and participate in this training course to fully familiarize yourself with topics related to network security:

Training destination NAT in Mikrotik for port forwarding operations

As we explained at the beginning of the article, Destination NAT is used to change the IP address and port. Port forwarding is a daily need for internet users, website administrators, and server and host management. In the following, we will learn about this concept, the use of Destination NAT, and how to set it up in Mikrotik.

## What is port forwarding?

Port forwarding intercepts and routes data traffic to a computer’s IP/port combination and redirects it to another IP or port. This process can be easily done using a MikroTik router or any system running RouterOS. It also creates a better user experience and optimal user management when using the Mikrotik router.

### How to use and configure port forwarding in Mikrotik

To understand the necessity and application of the port forwarding process, suppose you are an IT manager. You have created a large network, and someone wants to remotely connect to your VPS or dedicated server to work remotely. You cannot share the server IP with that person for security reasons. What should you do?

In this situation, you must use port forwarding in the Mikrotik router to handle all requests, and this operation is performed based on the Destination NAT feature. To configure port forwarding in MikroTik, you must first make sure that you have installed the latest version of MikroTik RouterOS, and then do these things step by step:

- Step 1: Log in to your MikroTik server with admin privileges
- Step 2: Click on IP from the left panel
- Step 3: In the submenu that opens, click on Firewall
- Step 4: Go to the NAT tab in the Firewall window
- Step 5: Click the + button to create a new rule. Note that in this scenario, the router is assumed to be connected to the IP (10.10.10.10), and we want to forward all requests from (10.10.10.10:5847) to (20.20.20.20:4324).
- Step 6: Click the General tab and select distant from the drop-down list
- Step 7: In the Dst field. Address: Type the IP from which you want to forward all requests
- Step 8: From the protocol list, select the connection protocol, such as TCP
- Step 9: In the Dst field. Port Type the port you want to forward requests to
- Step 10: Now, go to the Action tab
- Step 11: From the Action drop-down list, select DST-nat
- Step 12: In the To Addresses field, type the desired IP to send all requests to
- Step 13: In the To Ports field, type the port code to which you want to forward requests.
- Step 14: Click Apply and then OK to save and add the new rule

This way, you have successfully configured your first port forwarding rule in MikroTik. To add new port forwarding rules, follow the same steps for new ports or IPs.

## Source NAT training in Mikrotik

Suppose you want to hide your local devices behind the public IP address received from your Internet Service Provider (ISP). In that case, you must configure your MikroTik router’s source network address translation or masquerading feature. Suppose you want to hide the office computer and server behind the public IP 172.16.16.1. The rule required for this work in the operating system will be as follows:

By setting this rule, your ISP will see all requests sent to the 172.16.16.1 IP address, not your LAN IP addresses.

MikroTik routers and switches are affordable and reliable products. Due to their low cost, they are very suitable for small and medium-sized businesses. In some cases, even large companies use these products because they offer many features at a very low cost. However, they are unsuitable for supporting large volumes of data traffic like the main tier of ISPs. If you are interested in learning about the Mikrotik routing system, the following course is the best option:

Managers of one of the companies that have used MikroTik routers and switches in their development process, especially the CCR-1036 and CCR-1072 models, said about their experience: “What we like is the ease of use and how they work. It is data processing. Likewise, we have deployed RB devices as the customer’s CPE, and the results have been satisfactory.” Let’s review some features of Mikrotik routers:

- Support for Destination NAT to access servers inside the network from outside
- Support MPLS, BGP, and OSPF routing protocols
- The possibility of setting up a hotspot
- The possibility of network accounting
- An IP address supports version 6
- High boot speed on the router
- Support for the routing system based on PBR or Policy-Based Routing
- The possibility of remote access to a remote network through various VPN protocols
- VRF support
- Support for multiple virtual routers
- Ability to perform network traffic control operations
- Apply a speed limit for users
- The possibility of connecting to several Internet service providers
- The possibility of distributing users’ Internet traffic over several Internet links
- Ability to set up DHCP on the network to configure network clients
- The possibility of implementing quality of service on network packets
- The ability to use the MAC address for the initial configuration of the device without the need for an IP address at first

### Mikrotik router category and some of its most commonly used types

Mikrotik routers can be divided into four basic groups. This classification is based on several factors, including the connection technologies, the number of ports, the size, and the field of work of the devices. The four groups of these routers are:

- 3DIGIT routers
- 4DIGIT routers
- Naming routers
- Cloud Core routers

### Some of the famous models of these routers are

- RB133C: Simple Mikrotik router for subscriber-side wireless service
- RB450: from the Router Banel series with 300MHz processor, 32MB memory, and 5 Ethernet ports
- RB411: This router is used in links with medium bandwidth and high distances
- RB493: suitable for routing and wireless servicing of offices and centers connected to multiple Sowers
- RB433: suitable for wireless links with high bandwidth and long distances
- 433AH: Very powerful router with many features
- RB600A: It has high processing power and is similar to the 433AH router
- RB1000: It has a powerful 13333MHZ processor and 512MB memory

To fully familiarize yourself with Mikrotik company’s routers and their application and features, we suggest that all those interested participate in this course:

Learning and teaching destination NAT in Mikrotik: A factor for the development of companies

Mikrotik is one of the top companies producing network equipment and software required for it. Among its most important products are Mikrotik routers, as well as the company’s proprietary operating system and hardware called RouterOS. One useful command in this operating system is the destination NAT, which works based on the network address translation method.

This command and other functional commands in this operating system allow you to increase the efficiency of your equipment based on individual or corporate needs. Therefore, we suggest learning the skills of working with the Mikrotik operating system and those related to network security to all people who intend to learn the optimal and practical management of the network, server, and related equipment, or to teach others how to work with them.

## FAQ

### What is Destination NAT in MikroTik?

Destination NAT (dst-nat) changes the destination IP and/or port of incoming packets to forward them from a public address to an internal private address.

### Why use Destination NAT on a MikroTik router?

