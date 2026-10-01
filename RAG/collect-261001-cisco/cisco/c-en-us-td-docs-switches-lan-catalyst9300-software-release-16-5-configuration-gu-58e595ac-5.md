---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac-5
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac.md
source_anchor: ""
source_lines: [238, 304]
sha256: 7b37b0af04c82bb97f4d32b342051b05e2a765ae8796c940d1d33cf3848eed25
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac

                                          For example: match length A B match ip address acl1 acl2 match ip address acl3
 A packet is permitted if it is permitted by match length A B or acl1 or acl2 or acl3 
  - If the decision reached is permit, then the action specified by the set command is applied on the packet .
  - If the decision reached is deny, then the PBR action (specified in the set command) is not applied. Instead the processing logic moves forward to look at the next route-map statement in the sequence (the statement with the next higher sequence number). If no next statement exists, PBR processing terminates, and the packet is routed using the default IP routing table.
- A match command can match on
                                          				length or multiple ACLs. A route map statement can contain multiple match
                                          				commands. Logical or algorithm function is performed across all the match
                                          				commands to reach a permit or deny decision. 
                                          				
                                          
- For PBR, route-map statements marked as deny are not supported.
You can use standard IP ACLs to specify match criteria for a source address or extended IP ACLs to specify match criteria based on an application, a protocol type, or an end station. The process proceeds through the route map until a match is found. If no match is found, normal destination-based routing occurs. There is an implicit deny at the end of the list of match statements.
If match clauses are satisfied, you can use a set clause to specify the IP addresses identifying the next hop router in the path.
How to Configure PBR
- 
                                       					
                                       To use PBR, you must have the Network Essentials license enabled on the switch or active switch.
- 
                                       			 
                                       Multicast traffic is not policy-routed. PBR applies only to unicast traffic.
- 
                                       			 
                                       You can enable PBR on a routed port or an SVI.
- 
                                       			 
                                       The switch supports PBR based on match length.
- 
                                       			 
                                       You can apply a policy route map to an EtherChannel port channel in Layer 3 mode, but you cannot apply a policy route map to a physical interface that is a member of the EtherChannel. If you try to do so, the command is rejected. When a policy route map is applied to a physical interface, that interface cannot become a member of an EtherChannel.
- 
                                       			 
                                       You can define a mazimum of 128 IP policy route maps on the switch or switch stack.
- 
                                       			 
                                       You can define a maximum of 512 access control entries(ACEs) for PBR on the switch or switch stack.
- 
                                       			 
                                       When configuring match criteria in a route map, follow these guidelines: 
  - 
                                             				  
                                             Do not match ACLs that permit packets destined for a local address.
- 
                                             				  
                                             
- 
                                       			 
                                       VRF and PBR are mutually exclusive on a switch interface. You cannot enable VRF when PBR is enabled on an interface. The reverse is also true, you cannot enable PBR when VRF is enabled on an interface.
- 
                                       			 
                                       Web Cache Communication Protocol (WCCP) and PBR are mutually exclusive on a switch interface. You cannot enable WCCP when PBR is enabled on an interface. The reverse is also true, you cannot enable PBR when WCCP is enabled on an interface.
- 
                                       			 
                                       The number of hardware entries used by PBR depends on the route map itself, the ACLs used, and the order of the ACLs and route-map entries.
- 
                                       			 
                                       PBR based on TOS, DSCP and IP Precedence are not supported.
-  
                                       			 
                                       Set interface, set default next-hop and set default interface are not supported.
-  
                                       			 
                                       ip next-hop recursive and ip next-hop verify availability features are not available and the next-hop should be directly connected.
-  
                                       			 
                                       Policy-maps with no set actions are supported. Matching packets are routed normally.
-  
                                       			 
