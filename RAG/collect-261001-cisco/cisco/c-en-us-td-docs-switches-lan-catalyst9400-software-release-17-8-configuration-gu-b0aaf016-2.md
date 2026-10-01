---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016-2
title: "c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016.md
source_anchor: ""
source_lines: [64, 123]
sha256: 1ddc6c8c2012d99e6f14b7f9c3d7f37d022de16e3f2e6f08ffd83bd7777536de
---

# c-en-us-td-docs-switches-lan-catalyst9400-software-release-17-8-configuration-gu-b0aaf016

                                 Half Entry—represents a mapping between the local and global address/ports and is maintained in the translation database of NAT module. A half entry may be created statically or dynamically based on the configured NAT rule.
- 
                                 
                                 Full Entry/Flow entry—represents a unique flow corresponding to a given session. In addition to the local to global mapping, it also maintains the destination information which fully qualifies the given flow. A Full entry is always created dynamically and maintained in the translation database of NAT module.
VRF-Aware Network Address Translation
NAT is typically configured to operate across the default or global routing domain. As per this feature, the inside and outside NAT domains are associated with the default VRF space and the translations are effected accordingly. However, there are certain scenarios where NAT is required to operate in a VRF setting. One common scenario involves enabling shared service access for private networks that have overlapping address space. In such cases, the given private networks can be placed in different VRFs and global service access can be achieved by configuring VRF-aware NAT rules that map overlapping private address to unique global address. VRF-awareness enables NAT to carry out address and port translation by taking the VRF of the private networks into consideration.
VRF-aware NAT supports only VRF to Global translation of IP addresses. VRF to Global translation is between a NAT inside interface that is associated with a specific VRF and a NAT outside interface that is associated with the global VRF. Intra-VRF NAT translation (which involves the NAT-inside and NAT-outside interfaces of the same specific VRF) and Inter-VRF NAT translation (which involves NAT-inside and NAT-outside interfaces that are associated with different VRFs) are not supported. NAT behavior is undefined in such unsupported scenarios. We recommend that you deploy only the VRF to Global NAT translation in your network.
Types of NAT
You can configure NAT such that it will advertise only a single address for your entire network to the outside world. Doing this effectively hides the internal network from the world, giving you some additional security.
The types of NAT include:
- 
                                 
                                 Static address translation (static NAT)—Allows one-to-one mapping between local and global addresses.
- 
                                 
                                 Dynamic address translation (dynamic NAT)—Maps unregistered IP addresses to registered IP addresses from a pool of registered IP addresses.
- 
                                 
                                 Overloading / PAT—Maps multiple unregistered IP addresses to a single registered IP address (many to one) using different Layer 4 ports. This method is also known as Port Address Translation (PAT). By using overloading, thousands of users can be connected to the Internet by using only one real global IP address.
Using NAT to Route Packets to the Outside Network (Inside Source Address Translation)
You can translate unregistered IP addresses into globally unique IP addresses when communicating outside your network.
You can configure static or dynamic inside source address translation as follows:
- 
                                 
                                 Static translation establishes a one-to-one mapping between the inside local address and an inside global address. Static translation is useful when a host on the inside must be accessible by a fixed address from the outside. Static translation can be enabled by configuring a static NAT rule as explained in the Configuring Static Translation of Inside Source Addresses section.
- 
                                 
                                 Dynamic translation establishes a mapping between an inside local address and a pool of global addresses dynamically. Dynamic translation can be enabled by configuring a dynamic NAT rule and the mapping is established based on the result of the evaluation of the configured rule at run-time. You can employ an Access Control List (ACL), both Standarad and Extended ACLs, to specify the inside local address. The inside global address can be specified through an address pool or an interface. Dynamic translation is enabled by configuring a dynamic rule as explained in the Configuring Dynamic Translation of Inside Source Addresses section.
The following figure illustrates a device that is translating a source address inside a network to a source address outside the network.
The following process describes the inside source address translation, as shown in the figure above:
- 
                                 
                                 The user at host 10.1.1.1 opens a connection to Host B in the outside network.
- 
                                 
                                 NAT module intercepts the corresponding packet and attempts to translate the packet. The following scenarios are possible based on the presence or absence of a matching NAT rule: 
  - 
                                       
                                       If a matching static translation rule exists, the packet gets translated to the corresponding inside global address. Otherwise, the packet is matched against the dynamic translation rule and in the event of a successful match, it gets translated to the corresponding inside global address. The NAT module inserts a fully qualified flow entry corresponding to the translated packet, into its translation database. This facilitates fast translation and forwarding of the packets corresponding to this flow, in either direction.
  - 
                                       
                                       The packet gets forwarded without any address translation in the absence of a successful rule match.
  - 
                                       
                                       The packet gets dropped in the event of failure to obtain a valid inside global address even-though we have a successful rule match. Note 
 If an ACL is employed for dynamic translation, NAT evaluates the ACL and ensures that only the packets that are permitted by the given ACL are considered for translation. 
- 
                                       
                                       
- 
                                 
                                 The device replaces the inside local source address of host 10.1.1.1 with the inside global address of the translation, 203.0.113.2, (only the packet-relevant checksums get updated and all other fields in the packet remain unchanged) and forwards the packet.
- 
                                 
                                 The NAT module inserts a fully qualified flow entry corresponding to the translated packet flow, into its translation database. This facilitates fast translation and forwarding of packets corresponding to the flow in either direction.
- 
                                 
                                 Host B receives the packet and responds to host 10.1.1.1 by using the inside global IP destination address (DA) 203.0.113.2
- 
                                 
