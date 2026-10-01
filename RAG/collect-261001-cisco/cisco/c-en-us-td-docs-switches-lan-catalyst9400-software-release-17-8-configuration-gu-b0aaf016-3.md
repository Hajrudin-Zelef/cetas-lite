---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016-3
title: "c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016.md
source_anchor: ""
source_lines: [124, 198]
sha256: 000307d535d806aa6108aba262072680b4afe99d8db23477063c3f7cbff937a9
---

# c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016

                                 The response packet from host B would be destined to the inside global address and the NAT module intercepts this packet and translates it back to the corresponding inside local address with the help of the flow entry that has been setup in the translation database.
Host 10.1.1.1 receives the packet and continues the conversation. The device performs Steps 2 to 5 for each packet that it receives.
Outside Source Address Translation
You can translate the source address of the IP packets that travel from outside of the network to inside the network. This type of translation is usually employed in conjunction with inside source address translation to interconnect overlapping networks.
This process is explained in the section on Configuring Translation of Overlapping Networks
Port Address Translation (PAT)
You can conserve addresses in the inside global address pool by allowing a device to use one global address for many local addresses and this type of NAT configuration is called overloading or port address translation. When overloading is configured, the device maintains enough information from higher-level protocols (for example, TCP or UDP port numbers) to translate the global address back to the correct local address. When multiple local addresses map to one global address, the TCP or UDP port numbers of each inside host distinguish between the local addresses.
The figure below illustrates a NAT operation when an inside global address represents multiple inside local addresses. The TCP port numbers act as differentiators.
The device performs the following process in the overloading of inside global addresses, as shown in the figure above. Both Host B and Host C believe that they are communicating with a single host at address 203.0.113.2. Whereas, they are actually communicating with different hosts; the port number is the differentiator. In fact, many inside hosts can share the inside global IP address by using many port numbers.
- 
                                 
                                 The user at Host Y opens a connection to Host B and the user at Host X opens a connection to Host C.
- 
                                 
                                 NAT module intercepts the corresponding packets and attempts to translate the packets. Based on the presence or absence of a matching NAT rule the following scenarios are possible: 
  - 
                                       
                                       If a matching static translation rule exists, then it takes precedence and the packets are translated to the corresponding global address. Otherwise, the packets are matched against dynamic translation rule and in the event of a successful match, they are translated to the corresponding global address. NAT module inserts a fully qualified flow entry corresponding to the translated packets, into its translation database, to facilitate fast translation and forwarding of the packets corresponding to this flow, in either direction.
  - 
                                       
                                       The packets get forwarded without any address translation in the absence of a successful rule match.
  - 
                                       
                                       The packets get dropped in the event of failure to obtain a valid inside global address even though we have a successful rule match.
  - 
                                       
                                       As this is a PAT configuration, transport ports help translate multiple flows to a single global address. (In addition to source address, the source port is also subjected to translation and the associated flow entry maintains the corresponding translation mappings.)
- 
                                       
                                       
- 
                                 
                                 The device replaces inside local source address/port 10.1.1.1/1723 and 10.1.1.2/1723 with the corresponding selected global address/port 203.0.113.2/1024 and 203.0.113.2/1723 respectively and forwards the packets.
- 
                                 
                                 Host B receives the packet and responds to Host Y by using the inside global IP address 203.0.113.2, on port 1024. Host C receives the packet and responds to Host X using the inside global IP adress 203.0.113.2, on port 1723.
- 
                                 
                                 When the device receives the packets with the inside global IP address, it performs a NAT table lookup; the inside global address and port, and the outside address and port as keys; translates the addresses to the inside local addresses 10.1.1.1:1723 / 10.1.1.2:1723 and forwards the packets to Host Y and Host X respectively.
Host Y and Host X receive the packet and continue the conversation. The device performs Steps 2 to 5 for each packet it receives.
Overlapping Networks
Use NAT to translate IP addresses if the IP addresses that you use are neither legal nor officially assigned. Overlapping networks result when you assign an IP address to a device on your network that is already legally owned and assigned to a different device on the Internet or outside the network.
The following figure depicts overlapping networks: the inside network and outside network both have the same local IP addresses (10.1.1.x). You need network connectivity between such overlapping address spaces with one NAT device to translate the address of a remote peer (10.1.1.3) to a different address from the perspective of the inside.
Notice that the inside local address (10.1.1.1) and the outside global address ( 10.1.1.3) are in the same subnet. To translate the overlapping address, first, the inside source address translation happens with the inside local address getting translated to 203.0.113.2 and a half entry is created in the NAT table. On the Receiving side, the outside source address is translated to 172.16.0.3 and another half entry is created. The NAT table is then updated with a full entry of the complete translation.
The following steps describe how a device translates overlapping addresses:
- 
                                 
                                 Host 10.1.1.1 opens a connection to 172.16.0.3.
- 
                                 
                                 The NAT module sets up the translation mapping of the inside local and global addresses to each other and the outside global and local addresses to each other
- 
                                 
                                 The Source Address (SA) is replaced with inside global address and the Destination Address (DA) is replaced with outside global address.
- 
                                 
                                 Host C receives the packet and continues the conversation.
- 
                                 
                                 The device does a NAT table lookup, replaces the DA with inside local address, and replaces the SA with outside local address.
- 
                                 
                                 Host 10.1.1.1 receives the packet and the conversation continues using this translation process.
Limitations of Network Address Translation
- 
                                    
                                    There are certain NAT operations that are currently not supported in the hardware data plane. The following are such operations that are carried out in the relatively slower Software data plane: 
  - 
                                          
                                          Translation of Internet Control Message Protocol (ICMP) packets.
  - 
                                          
                                          Translation of packets that require application layer gateway (ALG) processing.
  - 
                                          
