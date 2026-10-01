---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908-2
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908.md
source_anchor: ""
source_lines: [58, 125]
sha256: a2dab8c541f2d1834bdfb4b87e9ca89a293a84a68522bc8a2f1ecd8bd43c3890
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908

                                 If the interface that the router ID is based on changes, that new IP address becomes the router ID. If any other interface changes its IP address, there is no router ID change.
BGP path selection
BGP best-path algorithm is a process used by Cisco NX-OS to select the optimal path for routing a given network prefix when multiple valid paths are available. For information on configuring additional BGP paths, see Configuring Advance BGP.
The best-path algorithm runs each time that a path is added or withdrawn for a given network. The best-path algorithm also runs if you change the BGP configuration. BGP selects the best path from the set of valid paths available for a given network.
Cisco NX-OS implements the BGP best-path algorithm in these steps:
- 
                                 				
                                 Compares two paths to determine which is better. See the Step 1 Comparing Pairs of Paths section).
- 
                                 				
                                 Explores all paths and determines in which order to compare the paths to select the overall best path. See the Step 2 Determining the Order of Comparisons section.
- 
                                 				
                                 Determines whether the old and new best paths differ enough so that the new best path should be used. See the Step 3 Determining the Best-Path Change Suppressionsection.
| Note | The order of comparison determined in Part 2 is important. Consider the case where you have three paths, A, B, and C. When Cisco NX-OS compares A and B, it chooses A. When Cisco NX-OS compares B and C, it chooses B. But when Cisco NX-OS compares A and C, it might not choose A because some BGP metrics apply only among paths from the same neighboring autonomous system and not among all paths. | 
The path selection uses the BGP AS-path attribute. The AS-path attribute includes the list of autonomous system numbers (AS numbers) traversed in the advertised path. If you subdivide your BGP autonomous system into a collection or confederation of autonomous systems, the AS-path contains confederation segments that list these locally defined autonomous systems.
| Note | VXLAN deployments use a BGP path selection process that differs from the normal selection of local over remote paths. For the EVPN address family, BGP compares the sequence number in the MAC Mobility attribute (if present) and selects the path with the higher sequence number. If both paths being compared have the attribute and the sequence numbers are the same, BGP prefers the path that is learned from the remote peer over a locally originated path. For more information, see the Cisco Nexus 9000 Series NX-OS VXLAN Configuration Guide. | 
BGP path selection - comparing pairs of paths
This first step in the BGP best-path algorithm compares two paths to determine which path is better. These following sequences describe the basic steps that Cisco NX-OS uses to compare two paths to determine the better path:
- 
                                    
                                    Cisco NX-OS chooses a valid path for comparison. For example, a path that has an unreachable next hop is not valid.
- 
                                    
                                    Cisco NX-OS chooses the path with the highest weight.
- 
                                    
                                    Cisco NX-OS chooses the path with the highest local preference.
- 
                                    
                                    If one of the paths is locally originated, Cisco NX-OS chooses that path.
- 
                                    
                                    Cisco NX-OS chooses the path with the shorter AS path. Note 
 When calculating the length of the AS-path, Cisco NX-OS ignores confederation segments and counts AS sets as 1. See the AS confederations section for more information. 
- 
                                    
                                    Cisco NX-OS chooses the path with the lower origin. Interior Gateway Protocol (IGP) is considered lower than EGP.
- 
                                    
                                    Cisco NX-OS chooses the path with the lower multiexit discriminator (MED). You can configure Cisco NX-OS to always perform the best-path algorithm MED comparison, regardless of the peer autonomous system in the paths. See the Tune the best-path algorithm section for more information. Otherwise, Cisco NX-OS performs a MED comparison that depends on the AS-path attributes of the two paths being compared. You can configure Cisco NX-OS to always perform the best-path algorithm MED comparison, regardless of the peer autonomous system in the paths. Otherwise, Cisco NX-OS performs a MED comparison that depends on the AS-path attributes of the two paths being compared these ways: 
  - 
                                          
                                          If a path has no AS-path or the AS-path starts with an AS_SET, the path is internal and Cisco NX-OS compares the MED to other internal paths.
  - 
                                          
                                          If the AS-path starts with an AS_SEQUENCE, the peer autonomous system is the first AS number in the sequence and Cisco NX-OS compares the MED to other paths that have the same peer autonomous system.
  - 
                                          
                                          If the AS-path contains only confederation segments or starts with confederation segments followed by an AS_SET, the path is internal and Cisco NX-OS compares the MED to other internal paths.
  - 
                                          
                                          If the AS-path starts with confederation segments that are followed by an AS_SEQUENCE, the peer autonomous system is the first AS number in the AS_SEQUENCE and Cisco NX-OS compares the MED to other paths that have the same peer autonomous system. Note 
 If Cisco NX-OS receives no MED attribute with the path, Cisco NX-OS considers the MED to be 0 unless you configure the best-path algorithm to set a missing MED to the highest possible value. See the Tune the best-path algorithm for more information. 
  - 
                                          
                                          If the non-deterministic MED comparison feature is enabled, the best-path algorithm uses the Cisco IOS style of MED comparison.
- 
                                          
                                          
- 
                                    
                                    If one path is from an internal peer and the other path is from an external peer, Cisco NX-OS chooses the path from the external peer.
- 
                                    
                                    If the paths have different IGP metrics to their next-hop addresses, Cisco NX-OS chooses the path with the lower IGP metric.
- 
                                    
