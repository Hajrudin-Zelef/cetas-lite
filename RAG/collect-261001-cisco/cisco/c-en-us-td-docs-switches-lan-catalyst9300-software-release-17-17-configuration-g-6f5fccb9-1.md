---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["asic", "ethernet", "memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9.md
source_anchor: ""
source_lines: [1, 118]
sha256: 41c3c8a2b5687a7a377b3e36cf44e562ea91f6a77ffcd81a034bff08718307f7
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9

- 
                                 					
                                 Flexible NetFlow is not supported on the Layer 2 port-channel interface, but is supported on the Layer 2 port-channel member
                                    ports.
                                 
- 
                                 					
                                 					
                                 Flexible NetFlow is supported on the Layer 3 port-channel interfaces and member ports but not on both at the same time for
                                    the same traffic type and direction.
                                 
- 
                                 					
                                 Flexible NetFlow Version 9 and Version 10 (IPFIX) export formats are supported. Flexible Netflow Version 5 Export Protocol is also supported. However, if you have not configured the export protocol, Version 9 export format is applied by default. 
                                 
- 
                                 					
                                 					
                                 Flexible Netflow without AVC and other NBAR or AVC enabled features cannot be configured at the same time on the same interface.
                                    Use AVC Flexible Netflow with these features.
                                  For information on how to configure AVC Flexible Netflow, refer "Configuring Application Visibility and Control in a Wired
                                    Network" section of the System Management Configuration guide.
                                 
- 
                                 					
                                 Layer 2, IPv4, and IPv6 traffic types are supported. A unique flow monitor for each traffic type can be applied for a given
                                    interface and direction. Multiple flow monitors for the same traffic type cannot be applied to an interface in the same direction.
                                 
- 
                                 					
                                 					
                                  Layer 2, VLAN, Layer 3 and SVI interfaces are supported, but the device does not support tunnels. 
- 
                                 					
                                 The following NetFlow table sizes are supported:  
                                    
                                    
                                       
                                       
                                       
                                    
                                    
                                       								
                                       
                                          									
                                          | Trim Level  | Ingress NetFlow Table  | Egress NetFlow Table  |  
                                          									
                                          | Network Essentials  | 32 K  | 32 K  |  
                                          									
                                          | Network Advantage  |  32 K  | 32 K  | 
- 
                                 					
                                 Depending on the switch type, a switch will have one or more forwarding ASICs, with multiple cores per ASIC. The capacities
                                    listed in the above table are on a per-Core/per-ASIC basis. 
                                 
- 
                                 					
                                 The hardware NetFlow tables which exist on separate ASIC cores cannot be combined. Flows will be created in the table on the
                                    core which processed that packet.
                                 
- 
                                 					
                                 NetFlow hardware implementation supports four hardware samplers. You can select a sampler rate from 1 out of 2 to 1 out of
                                    1024. Both — random and deterministic — sampling modes are supported. 
                                 
- 
                                 					
                                 NetFlow hardware uses hash tables internally. Hash collisions can occur in the hardware. Therefore, in spite of the internal
                                    overflow Content Addressable Memory (CAM), the actual NetFlow table utilization could be about 80 percent. 
                                 
- 
                                 					
                                 Depending on the fields that are used for the flow, a single flow could take two consecutive entries. IPv6 and datalink flows
                                    also take two entries. In these situations, the effective usage of NetFlow entries is half the table size, which is separate
                                    from the above hash collision limitation. 
                                 
- 
                                 					
                                 The device supports up to 15 flow monitors. 
- 
                                 					
                                 Ingress flows are stored in the ingress table for the corresponding ASIC and core that received the packets for that flow.
                                    Egress flows are stored in the egress table for the corresponding ASIC and core that sent packets for that flow.
                                 
- 
                                 					
                                 The reported value for the bytes count field (called “bytes long”) is Layer-2-packet-size—18 bytes. For classic Ethernet traffic
                                    (802.3), this will be accurate. For all other Ethernet types, this field will not be accurate. Use the "bytes layer2” field,
                                    which always reports the accurate Layer 2 packet size. For information about supported Flexible NetFlow fields, see 'Supported
                                    Flexible NetFlow Fields' topic.
                                 
- 
                                 					
                                 Flexible NetFlow export is not supported on the Ethernet management port, GigabitEthernet 0/0. 
- 
                                 					
                                  When a flow record has only Source Group Tag (SGT) and Destination Group Tag (DGT) fields (or only either of the two) and
                                    if both the values are not applicable, then a flow will still be created with zero values for SGT and DGT. The flow records
                                    are expected to include source and destination IP addresses, along with SGT and DGT fields. 
                                 
- 
                                 					
                                 					
                                 On non-Cisco TrustSec interfaces, an SGT value of zero implies that there is no command header. On Cisco TrustSec interfaces,
                                    an SGT value of zero implies an unknown tag.
                                 
- 
                                 					
                                 					
                                 					
                                 When a quality of service (QoS) marked packet is received on an interface which has NetFlow configured in the ingress direction,
                                    the QoS value of the packet is captured by the NetFlow collector. However, when the packet is received on an interface which
                                    has NetFlow configured in the egress direction and the QoS value has been rewritten on ingress by the switch, the new QoS
                                    value of the packet is not captured by the collector. 
                                 
- 
                                 					
                                 					
