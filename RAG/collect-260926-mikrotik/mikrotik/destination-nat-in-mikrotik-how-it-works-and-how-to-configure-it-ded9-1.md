---
id: collect-260926-mikrotik/mikrotik/destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9-1
title: "destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2019-08"]
keywords: ["research", "training"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9.md
source_anchor: ""
source_lines: [1, 68]
sha256: df16d9619440d17a14843a64a8c784ddb96f97249e7f18283a5180ebb2083754
---

# destination-nat-in-mikrotik-how-it-works-and-how-to-configure-it-ded9

Mikrotik is the designer of the Mikrotik operating system and also manufactures hardware products called Routerboards and products such as bandwidth controllers. One of the useful commands in Mikrotik routers is Destination NAT, which is used to change the public IP address to a private internal IP address. In this article, from the IT assistant, we will learn about NAT technology, Destination NAT, and Destination NAT training in Mikrotik.

## What is NAT technology?

NAT stands for Network Address Translation. It is a method of routing information in the Internet space that translates several local private addresses to public addresses before transferring the information. Organizations that want multiple devices to use the same IP address in their internal environment use NAT. Home routers work based on this technology. NAT technology is one of the most important technologies in computer networks.

## How does NAT work?

Suppose a laptop is connected to a home router. A person uses the laptop to search for directions to their favorite restaurant. The laptop sends this request in a data packet to the router, which forwards it to the web. But before sending the information, the router first changes the outgoing IP address from a private local address to a public address.

If the intended data packet was sent with the same private IP address, the receiving server would not know where to send the data. This issue is similar to sending physical letters and requesting return or reply services. If the IP address is sent to the receiving server, it is the same as providing an anonymous address instead of our address on the mail or leaving it blank. With NAT, information is routed back to the sender’s laptop using the router’s public address, not the laptop’s private address.

## The difference between Source NAT and Destination NAT

NAT routing technology works in the third layer of the OSI model, or the Network layer. This is why NAT works with IP addresses. As mentioned earlier, the basic job of NAT is to convert a local private IP address to a public IP address and vice versa. The same route and round-trip create two types of NAT operations. The first case is when the private IP is converted to a public IP, called the Source NAT operation. The second case is when the public IP is converted or translated into a private IP called Destination NAT.

- Destination NAT operation and technology change the destination address of the packets passing through the router. This method provides a solution for port translation in TCP/UDP headers.
- Destination NAT usually directs incoming packets with an external address or port destination to an internal IP address or port within the network.
- Different applications of Destination NAT

Destination NAT is performed, destination IP addresses are configured and translated according to Destination NAT rules, and security policies are applied. Destination NAT is commonly used to perform the following actions:

- Translate a single IP address to another: for example, to allow a device on the Internet to connect to a host on a private network
- Translate a contiguous block of IP addresses into another block of addresses of the same size: for example, to allow access to a group of servers.
- Translate the destination IP address and port to another destination IP address and port: for example, to allow access to multiple services using the same IP address but different ports.

## Destination NAT address pool definition

A NAT Pool is a set of user-defined IP addresses used for translation. Unlike static NAT, where there is one-to-one routing and IP translation, group-to-group IP translation is possible here. Meanwhile, in this case, the original destination IP address is translated to an IP address from a user-defined pool. Therefore, if the original destination IP address range is larger than the address range in the user-defined address pool, any untranslated packets are dropped. This feature is one of the applications of Destination NAT.

You can configure a NAT Pool to exist in the default routing instance. A configuration option specifies whether a NAT pool exists in the default routing instance. As a result, the NAT pool is discoverable and reachable from zones in the default routing instance and from zones in other routing instances.

Introduction of Mikrotik company

MikroTik is a network equipment manufacturer based in Latvia, Europe. It develops and sells wired and wireless network routers, switches, operating systems, and supporting software.

Mikrotik was founded in 1996 to sell network equipment in emerging markets. As of August 2019, the company’s website reported the number of employees to be around 280. In 2015, Mikrotik was the 20th largest company in Latvia, with revenues of 202 million euros. Mikrotik was originally a PC software company. In 2002, the company began producing its hardware.

## What is the Mikrotik router operating system?

MikroTik router operating system, MikroTik RouterOS, is an independent Linux operating system used in MikroTik network equipment. Of course, this is more than just an operating system for routers. This software can even be installed on regular PCs to turn them into dedicated routers.

Mikrotik’s operating system increases the efficiency of network equipment manufactured by Mikrotik. This operating system is also highly flexible, allowing it to be installed on computers and turned into routers. The Mikrotik operating system has routing, Firewall, bandwidth management, creation of wireless access points, backhaul link, spot ports, and a VPN server.

RouterOS operating system can be used in networks in two ways: in the first case, this operating system can be installed on a personal computer or virtual machine; The second mode is when this operating system is installed on the RouterBoard hardware in Mikrotik’s physical equipment. In both cases, we have created a Mikrotik router. Routerboards are Mikrotik’s exclusive hardware and RouterOS operating system with several network ports.

## Advantages of using the Mikrotik router operating system

Companies can use Mikrotik’s operating system as a trial solution during the R&D process. Practical results have shown that this operating system’s research and development phase has succeeded. Many companies have extensively implemented it to serve their customers as part of their development solution. Today, this operating system can be converted into a cloud management suite since it adds many improvements and functions. Let’s review some of the advantages of using this operating system:

- It is affordable
- It is easy to deploy
- It is powerful and practical
- Compared to other brands at the same price, it has more features
- It is widely available
- Today, it can be easily obtained
- It is very adjustable and flexible
- You can use it to write a powerful script to improve performance
- You can create a different configuration approach based on your needs

User management training with Mikrotik – click

## Training destination NAT in Mikrotik for more safety

Network address translation in the Destination NAT method is possible by changing the network address information in the IP header of packets. Let’s look at a common setup where a network administrator wants to access an administrative server from the Internet. We want to allow connections from the Internet to the administration server, whose local IP is 10.0.0.3. In this case, we need to configure a destination address translation rule on the administrative gateway router as follows:

The above rule translates to: When an incoming connection requests TCP port 22 with destination address 172.16.16.1, use the Destination NAT or DST-nat function and forward the packets to the device with local IP address 10.0.0.3 and port 22 transfer.

