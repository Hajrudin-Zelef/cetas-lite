---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-1f4955e4-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-1f4955e4"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-1f4955e4.md
source_anchor: ""
source_lines: [64, 121]
sha256: db262b760ee15bd7dd524fefd1d8e17b447926e4e764c5319739b28c4ca72462
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-1f4955e4

                                    ARP entries are learned in separate VRFs. The user can display Address Resolution Protocol (ARP) entries for specific VRFs.
Multi-VRF CE Configuration Guidelines
This section provides guidelines for configuring multi-VRF CE:
| Note | To use multi-VRF CE, you must have the Network Advantage license enabled on your switch. | 
- 
                                    
                                    A switch with multi-VRF CE is shared by multiple customers, and each customer has its own routing table.
- 
                                    
                                    Because customers use different VRF tables, the same IP addresses can be reused. Overlapped IP addresses are allowed in different VPNs.
- 
                                    
                                    Multi-VRF CE lets multiple customers share the same physical link between the PE and the CE. Trunk ports with multiple VLANs separate packets among customers. Each customer has its own VLAN.
- 
                                    
                                    Multi-VRF CE does not support all MPLS-VRF functionality. It does not support label exchange, LDP adjacency, or labeled packets.
- 
                                    
                                    For the PE router, there is no difference between using multi-VRF CE or using multiple CEs. In Figure 41-6, multiple virtual Layer 3 interfaces are connected to the multi-VRF CE device.
- 
                                    
                                    The switch supports configuring VRF by using physical ports, VLAN SVIs, or a combination of both. The SVIs can be connected through an access port or a trunk port.
- 
                                    
                                    A customer can use multiple VLANs as long as they do not overlap with those of other customers. A customer’s VLANs are mapped to a specific routing table ID that is used to identify the appropriate routing tables stored on the switch.
- 
                                    
                                    The switch supports one global network and up to 256 VRFs.
- 
                                    
                                    Most routing protocols (BGP, OSPF, RIP, and static routing) can be used between the CE and the PE. However, we recommend using external BGP (EBGP) for these reasons: 
  - 
                                          
                                          BGP does not require multiple algorithms to communicate with multiple CEs.
  - 
                                          
                                          BGP is designed for passing routing information between systems run by different administrations.
  - 
                                          
                                          BGP makes it easy to pass attributes of the routes to the CE.
- 
                                          
                                          
- 
                                    
                                    Multi-VRF CE does not affect the packet switching rate.
- 
                                    
                                    VPN multicast is not supported.
- 
                                    
                                    You can enable VRF on a private VLAN, and the reverse.
- 
                                    
                                    You cannot enable VRF when policy-based routing (PBR) is enabled on an interface, and the reverse.
- 
                                    
                                    You cannot enable VRF when Web Cache Communication Protocol (WCCP) is enabled on an interface, and the reverse.
