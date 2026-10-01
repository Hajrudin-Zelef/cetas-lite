---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908-3
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908.md
source_anchor: ""
source_lines: [126, 184]
sha256: 72feec3a0ca56f403d64200f460e349045826bba45b6f2eb192e6e7d2bdd161c
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908

                                    Cisco NX-OS uses the path that was selected by the best-path algorithm the last time that it was run. If all path parameters in Step 1 through Step 9 are the same, you can configure the best-path algorithm to enforce comparison of the router IDs when both paths are eBGP by configuring “compare router-id”. In all other cases, the router-id comparison is done by default. See the Tune the best-path algorithm for more information. If the path includes an originator attribute, Cisco NX-OS uses that attribute as the router ID to compare to; otherwise, Cisco NX-OS uses the router ID of the peer that sent the path. If the paths have different router IDs, Cisco NX-OS chooses the path with the lower router ID. Note 
 When using the attribute originator as the router ID, it is possible that two paths have the same router ID. It is also possible to have two BGP sessions with the same peer router, so you could receive two paths with the same router ID. 
- 
                                    
                                    Cisco NX-OS selects the path with the shorter cluster length. If a path was not received with a cluster list attribute, the cluster length is 0.
- 
                                    
                                    Cisco NX-OS chooses the path received from the peer with the lower IP address. Locally generated paths (for example, redistributed paths) have a peer IP address of 0. Note 
 Paths that are equal after Step 9 can be used for multipath if you configure multipath. See the Load sharing and multipath section for more information. 
BGP path selection - determining the order of comparisons
In the second step of the BGP best-path algorithm implementation, Cisco NX-OS uses these steps to compares the paths:
- 
                                    				
                                    Cisco NX-OS partitions the paths into groups. Within each group, Cisco NX-OS compares the MED among all paths. Cisco NX-OS uses the same rules as in the Step 1—Comparing Pairs of Paths to determine whether MED can be compared between any two paths. Typically, this comparison results in one group being chosen for each neighbor autonomous system. If you configure the bgp bestpath med always command, Cisco NX-OS chooses just one group that contains all the paths.
- 
                                    				
                                    Cisco NX-OS determines the best path in each group by iterating through all paths in the group and keeping track of the best one so far. Cisco NX-OS compares each path with the temporary best path found so far and if the new path is better, it becomes the new temporary best path and Cisco NX-OS compares it with the next path in the group.
- 
                                    				
                                    Cisco NX-OS forms a set of paths that contain the best path selected from each group in Step 2. Cisco NX-OS selects the overall best path from this set of paths by going through them as in Step 2.
BGP path selection - determining the best-path change suppression
The next part of the implementation is to determine whether Cisco NX-OS uses the new best path or suppresses the new best path. The router can continue to use the existing best path if the new one is identical to the old path (if the router ID is the same). Cisco NX-OS continues to use the existing best path to avoid route changes in the network.
You can turn off the suppression feature by configuring the best-path algorithm to compare the router IDs. See the Tuning the Tune the best-path algorithm section for more information. If you configure this feature, the new best path is always preferred to the existing one.
BGP and the unicast RIB
BGP communicates with the unicast routing information base (unicast RIB) to store IPv4 and IPv6 routes in the unicast routing table. After selecting the best path, if BGP determines that the best path change needs to be reflected in the routing table, it sends a route update to the unicast RIB.
BGP receives route notifications regarding changes to its routes in the unicast RIB. It also receives route notifications about other protocol routes to support redistribution.
BGP also receives notifications from the unicast RIB regarding next-hop changes. BGP uses these notifications to keep track of the reachability and IGP metric to the next-hop addresses.
Whenever the next-hop reachability or IGP metrics in the unicast RIB change, BGP triggers a best-path recalculation for affected routes.
BGP communicates with the IPv6 unicast RIB to perform these operations for IPv6 routes.
BGP prefix independent convergence
The BGP prefix independent convergence (PIC) edge feature achieves faster convergence in the forwarding plane for BGP IP routes to a BGP backup path when there is a link failure.
The BGP PIC edge feature improves BGP convergence after a network failure. This convergence applies to edge failures in an IP network. This feature creates and stores a backup path in the routing information base (RIB) and forwarding information base (FIB) so that when the primary path fails, the backup path can immediately take over, enabling fast failover in the forwarding plane. BGP PIC edge supports only IPv4 address families.
When BGP PIC edge is configured, BGP calculates a second-best path (the backup path) along with the primary best path. BGP installs both best and backup paths for the prefixes with PIC support into the BGP RIB. BGP also downloads the backup path along with the remote next hop through APIs to the URIB, which then updates the FIB with the next hop marked as a backup. The backup path provides a fast reroute mechanism to counter a singular network failure.
This feature detects both local interface failures and remote interface or link failures and triggers the use of the backup path.
BGP PIC edge supports both unipath and multipath.
BGP PIC edge unipath
Figure GBP PIC Edge Unipath shows a BGP PIC edge unipath topology.
In this figure:
- 
                                    
                                    eBGP sessions are between S2-S4 and S3-S5.
- 
                                    
                                    The iBGP session is between S2-S3.
- 
                                    
                                    Traffic from S1 uses S2 and uses the e1 interface to reach prefixes Z1...Zn.
- 
                                    
                                    S2 has two paths to reach Z1…Zn: 
  - 
                                          
                                          A primary path through S4
  - 
                                          
                                          A backup path through S5
- 
                                          
                                          
