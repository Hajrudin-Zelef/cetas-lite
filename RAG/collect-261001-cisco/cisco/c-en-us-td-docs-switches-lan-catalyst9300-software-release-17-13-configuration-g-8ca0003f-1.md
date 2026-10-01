---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "license"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f.md
source_anchor: ""
source_lines: [1, 64]
sha256: 8fceacd59913aa27f1a1c5156981fca1446e987682b5c9bd70d2ba495f9b6148
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-8ca0003f

Information About IPv6 Unicast Routing
This chapter describes how to configure IPv6 unicast routing on the switch.
| Note | To use all IPv6 features in this chapter, the switch or active switch must be running the Network Advantage license. Switches running the Network Essentials license support IPv6 static routing and RIP for IPv6. Switches running the Network Advantage license support OSPF, EIGRP, and BGP for IPv6. | 
Understanding IPv6
IPv4 users can move to IPv6 and receive services such as end-to-end security, quality of service (QoS), and globally unique addresses. The IPv6 address space reduces the need for private addresses and Network Address Translation (NAT) processing by border routers at network edges.
For information about how Cisco Systems implements IPv6, go to Networking Software (IOS & NX-OS)
For information about IPv6 and other features in this chapter
- 
                                 
                                 See the Cisco IOS IPv6 Configuration Library.
- 
                                 
                                 Use the Search field on Cisco.com to locate the Cisco IOS software documentation. For example, if you want information about static routes, you can enter Implementing Static Routes for IPv6 in the search field to learn about static routes.
Static Routes for IPv6
Static routes are manually configured and define an explicit route between two networking devices. Static routes are useful for smaller networks with only one path to an outside network or to provide security for certain types of traffic in a larger network.
Configuring Static Routing for IPv6 (CLI)
For configuring static routes for IPv6, see the Configuring Static Routing for IPv6 section.
For more information about static routes, see the “Implementing Static Routes for IPv6” chapter in the Cisco IOS IPv6 Configuration Library on Cisco.com.
Path MTU Discovery for IPv6 Unicast
The switch supports advertising the system maximum transmission unit (MTU) to IPv6 nodes and path MTU discovery. Path MTU discovery allows a host to dynamically discover and adjust to differences in the MTU size of every link along a given data path. In IPv6, if a link along the path is not large enough to accommodate the packet size, the source of the packet handles the fragmentation.
ICMPv6
The Internet Control Message Protocol (ICMP) in IPv6 generates error messages, such as ICMP destination unreachable messages, to report errors during processing and other diagnostic functions. In IPv6, ICMP packets are also used in the neighbor discovery protocol and path MTU discovery.
Neighbor Discovery
The switch supports NDP for IPv6, a protocol running on top of ICMPv6, and static neighbor entries for IPv6 stations that do not support NDP. The IPv6 neighbor discovery process uses ICMP messages and solicited-node multicast addresses to determine the link-layer address of a neighbor on the same network (local link), to verify the reachability of the neighbor, and to keep track of neighboring routers.
The switch supports ICMPv6 redirect for routes with mask lengths less than 64 bits. ICMP redirect is not supported for host routes or for summarized routes with mask lengths greater than 64 bits.
Neighbor discovery throttling ensures that the switch CPU is not unnecessarily burdened while it is in the process of obtaining the next hop forwarding information to route an IPv6 packet. The switch drops any additional IPv6 packets whose next hop is the same neighbor that the switch is actively trying to resolve. This drop avoids further load on the CPU.
IPv6 Router Advertisement Options for DNS Configuration
Most of the internet services are identified by a Domain Name Server (DNS) name. IPv6 Router Advertisement (RA) provides the following two options to allow IPv6 hosts to perform automatic DNS configuration:
- 
                                    
                                    Recursive DNS Server (RDNSS)
- 
                                    
                                    DNS Search List (DNSSL)
RDNSS contains the address of recursive DNS servers that help in DNS name resolution in IPv6 hosts. DNS Search List is a list of DNS suffix domain names that are used by IPv6 hosts when they perform DNS query searches.
For more information on RA options for DNS configuration, refer IETF RFC 6106.
For configuring DNSSL, see the Configuring DNS Search List Using IPv6 Router Advertisement Options section of the IP Addressing Services Configuration Guide.
Default Router Preference
The switch supports IPv6 default router preference (DRP), an extension in router advertisement messages. DRP improves the ability of a host to select an appropriate router, especially when the host is multihomed and the routers are on different links. The switch does not support the Route Information Option in RFC 4191.
An IPv6 host maintains a default router list from which it selects a router for traffic to offlink destinations. The selected router for a destination is then cached in the destination cache. NDP for IPv6 specifies that routers that are reachable or probably reachable are preferred over routers whose reachability is unknown or suspect. For reachable or probably reachable routers, NDP can either select the same router every time or cycle through the router list. By using DRP, you can configure an IPv6 host to prefer one router over another, provided both are reachable or probably reachable.
For configuring DRP for IPv6, see the Configuring Default Router Preference section.
For more information about DRP for IPv6, see the Cisco IOS IPv6 Configuration Library on Cisco.com.
Policy-Based Routing for IPv6
Policy-based routing (PBR) gives you a flexible means of routing packets by allowing you to configure a defined policy for traffic flows, which lessens reliance on routes that are derived from routing protocols. Therefore, PBR gives you more control over routing by extending and complementing the existing mechanisms that are provided by routing protocols. PBR allows you to set the IPv6 precedence. For a simple policy, you can use any one of these tasks; for a complex policy, you can use all of them. It also allows you to specify a path for certain traffic, such as priority traffic over a high-cost link.
PBR for IPv6 may be applied to both forwarded and originated IPv6 packets. For forwarded packets, PBR for IPv6 will be implemented as an IPv6 input interface feature, supported in the following forwarding paths:
- 
                                    				
                                    Process
- 
                                    				
                                    Cisco Express Forwarding (formerly known as CEF)
- 
                                    				
                                    Distributed Cisco Express Forwarding
Policies can be based on the IPv6 address, port numbers, protocols, or packet size.
PBR allows you to perform the following tasks:
- 
                                    				
                                    Classify traffic based on extended access list criteria. Access lists, then, establish the match criteria.
- 
                                    				
                                    Set IPv6 precedence bits, giving the network the ability to enable differentiated classes of service.
- 
                                    				
