---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9-2
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9.md
source_anchor: ""
source_lines: [119, 160]
sha256: 9a40b6ef860070e560a3060d8dc1482ca2084d394b43ecc37690841a16d26f0d
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-17-configuration-g-6f5fccb9

                                 For an IPv6 flow monitor, Source Group Tag (SGT) and Destination Group Tag (DGT) fields cannot co-exist with MAC address fields.
- 
                                 					
                                 					
                                 NetFlow records do not support MultiProtocol Label Switching-enabled (MPLS-enabled) interfaces. 
- 
                                 					
                                 					
                                 Data capture based on MPLS label inside the MPLS network is not supported. Capture of IP header fields of an MPLS tagged packet
                                    is not supported.
                                 
- 
                                 					
                                 Egress flow monitors do not capture flows that are egressing out in EoMPLS mode or in L3VPN Per-Prefix mode.
- 
                                 					
                                 					
                                 Flow exporter exports the flow data only after the template data time out period ends. Configuration changes such as VPN ID
                                    modification or VRF deletion will take effect after the time out period ends.
                                 
- 
                                 					
                                 					
                                 A flow monitor cannot be shared across Layer 3 physical interfaces and logical interfaces (such as, Layer 3 port-channel interface,
                                    Layer 3 port-channel member, and switch virtual interface [SVI]), but a flow monitor can be shared across logical interfaces
                                    or Layer 3 physical interfaces.
                                 
- 
                                 					
                                 					
                                 When Flexible NetFlow and Network Address Translation (NAT) are configured on an interface,  
                                    						
                                    
  - 
                                       							
                                       Flexible NetFlow will display and export the actual flow details; but not the translated flow details. Application-level gateway
                                          (ALG) flow details are not part of the actual flow details that are exported.
                                       
  - 
                                       							
                                       If the ALG traffic gets translated through the CPU, Flexible NetFlow will display and export the translated flow details for
                                          the ALG traffic.
