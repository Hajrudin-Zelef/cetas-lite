---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908-4
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908.md
source_anchor: ""
source_lines: [185, 294]
sha256: 4b88713379ed204387a4b58560a5683b669398c32226d508ff6679b9384ddd48
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-101x-configuration-unicast-cisco-n9000-nx-os-4a750908

In this example, S3 advertises to S2 the prefixes Z1…Zn to reach (with itself as the next hop). With BGP PIC edge enabled, BGP on S2 installs both the best path (through S4) and the backup path (through S3 or S5) toward the AS6500 into the RIB. Then the RIB downloads both routes to the FIB.
If the S2-S4 link goes down, the FIB on S2 detects the link failure. It automatically switches from the primary path to the backup path and points to the new next hop S3. Traffic is quickly rerouted due to the local fast re-convergence in the FIB. After learning of the link failure event, BGP on S2 recomputes the best path (which is the previous backup path), removes next hop S4 from the RIB, and reinstalls S3 as the primary next hop into the RIB. BGP also computes a new backup path, if any, and notifies the RIB. With the support of the BGP PIC edge feature, the FIB can switch to the available backup route instantly upon detection of a link failure on the primary route without waiting for BGP to select the new best path and converge to achieve a fast reroute.
BGP PIC edge with multipath
The following figure shows a BGP PIC edge multipath topology.
In this topology, there are six paths for a given prefix:
- 
                                    
                                    eBGP paths: e1, e2, e3
- 
                                    
                                    iBGP paths: i1, i2, i3
The order of preference is e1 > e2 > e3 > i1 > i2 > i3.
The potential multipath situations are:
- 
                                    
                                    No multipaths configured: 
  - 
                                          
                                          bestpath = e1
  - 
                                          
                                          multipath-set = []
  - 
                                          
                                          backup path = e2
  - 
                                          
                                          PIC behavior: When e1 fails, e2 is activated.
- 
                                          
                                          
- 
                                    
                                    Two-way eBGP multipaths configured: 
  - 
                                          
                                          bestpath = e1
  - 
                                          
                                          multipath-set = [e1, e2]
  - 
                                          
                                          backup path = e3
  - 
                                          
                                          PIC behavior: Active multipaths are mutually backed up. When all multipaths fail, e3 is activated.
- 
                                          
                                          
- 
                                    
                                    Three-way eBGP multipaths configured: 
  - 
                                          
                                          bestpath = e1
  - 
                                          
                                          multipath-set = [e1, e2, e3]
  - 
                                          
                                          backup path = i1
  - 
                                          
                                          PIC behavior: Active multipaths are mutually backed up. When all multipaths fail, i1 is activated.
- 
                                          
                                          
- 
                                    
                                    Four-way eBGP multipaths configured: 
  - 
                                          
                                          – bestpath = e1
  - 
                                          
                                          – multipath-set = [e1, e2, e3, i1]
  - 
                                          
                                          – backup path = i2
  - 
                                          
                                          – PIC behavior: Active multipaths are mutually backed up. When all multipaths fail, i2 is activated.
- 
                                          
                                          
When the Equal Cost Multipath Protocol (ECMP) is enabled, none of the multipaths can be selected as the backup path.
For multipaths with the backup path scenario, faster convergence is not expected with simultaneous failure of all active multipaths.
BGP PIC core
BGP Prefix Independent Convergence (PIC) in Core is a BGP optimization technique that
- 
                                    
                                    improves BGP convergence speed after a network failure,
- 
                                    
                                    reduces the time and resources needed to update forwarding information for multiple prefixes, and
- 
                                    
                                    enables immediate leveraging of IGP convergence by hierarchical FIB programming.
When a link fails on Provider Edge (PE), the Routing Information Base (RIB) updates the Forwarding Information Base (FIB) with new next hop. FIB must update all BGP prefixes that point to the failed next hop and point to the new one. This can be time and resource consuming. With BGP PIC Core enabled, the prefix is programmed in the FIB in a hierarchical way. All prefixes point to the ECMP group instead of the recursive next hop. When the same failure happens, the FIB only needs to update the ECMP group to point to the new next hop without updating prefixes. This gives BGP immediate leveraging of IGP convergence.
BGP PIC feature support matrix
| Table 2. BGP PIC Feature Support Matrix |  |  | 
|---|---|---|
| BGP PIC | IPv4 Unicast | IPv6 Unicast | 
|---|---|---|
| Edge unipath | Yes | No | 
| Edge with multipath (multiple active ECMPs, only one backup) | Yes | No | 
| Core | Yes | Yes | 
| Note | The PIC Core (system pic-core command) and PIC Edge must be used exclusively in a Layer 3 environment and are not compatible with VXLAN environment. | 
BGP virtualization
BGP supports virtual routing and forwarding (VRF) instances.
