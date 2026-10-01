---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107-5
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107.md
source_anchor: ""
source_lines: [203, 340]
sha256: b9be0579e112f5617f09c3dcaef291f6197f6afadef74fca7830bd358236d7c0
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-interfaces-configuration-gu-240a5107

                                    Layer 2 port channels for vPC Peer-Link must be configured in trunk mode.
- 
                                    					
                                    Compatibility parameters must be identical across all interfaces in the vPC.
For example, the compatibility check process differs for vPCs compared to regular port channels.
Configuration and Guidelines
After enabling the vPC feature and configuring the vPC Peer-Link, Cisco Fabric Services (CFS) ensures configuration consistency between the local and remote vPC peer devices.
| Note | Enter the show vpc consistency-parameters command to display the configured values on all interfaces in the vPC. The displayed configurations are only those that would limit the vPC Peer-Link and vPC from coming up. | 
| Note | The port channel compatibility parameters must be the same for all the port channel members on the physical switch. You cannot configure shared interfaces to be part of a vPC. | 
See the “Configuring Port Channels” chapter for more details about regular port channels.
Configuration Parameters That Must Be Identical
The configuration parameters in this section must be configured identically on both devices of the vPC Peer-Link; otherwise, the vPC moves fully or partially into a suspended mode.
- 
                                          						
                                          You must ensure that all interfaces in the vPC have the identical operational and configuration parameters listed in this section.
- 
                                          						
                                          Enter the show vpc consistency-parameters command to display the configured values on all interfaces in the vPC. The displayed configurations are only those configurations that would limit the vPC Peer-Link and vPC from coming up.
The devices automatically check for compatibility for some of these parameters on the vPC interfaces. The per-interface parameters must be consistent per interface, and the global parameters must be consistent globally.
- 
                                          						
                                          Port-channel mode: on, off, or active (port-channel mode can, however, be active/passive on each side of the vPC peer)
- 
                                          						
                                          Link speed per channel
- 
                                          						
                                          Duplex mode per channel
- 
                                          						
                                          Trunk mode per channel: 
  - 
                                                								
                                                Native VLAN
  - 
                                                								
                                                VLANs allowed on trunk
  - 
                                                								
                                                Tagging of native VLAN traffic
- 
                                                								
                                                
- 
                                          						
                                          Spanning Tree Protocol (STP) mode
- 
                                          						
                                          STP region configuration for Multiple Spanning Tree
- 
                                          						
                                          Enable/disable state per VLAN
- 
                                          						
                                          STP global settings: 
  - 
                                                								
                                                Bridge Assurance setting
  - 
                                                								
                                                Port type setting
  - 
                                                								
                                                Loop Guard settings
- 
                                                								
                                                
- 
                                          						
                                          STP interface settings: 
  - 
                                                								
                                                Port type setting
  - 
                                                								
                                                Loop Guard
  - 
                                                								
                                                Root Guard
- 
                                                								
                                                
- 
                                          						
                                          Maximum Transmission Unit (MTU)
If any of these parameters are not enabled or defined on either device, the vPC consistency check ignores those parameters.
- 
                                          						
                                          To ensure that none of the vPC interfaces are in the suspend mode, enter the show vpc brief and show vpc consistency-parameters commands and check the syslog messages. In the output of show vpc or show vpc brief command, after every 50th configured vPC port-channel the following message will be displayed: Please check "show vpc consistency-parameters vpc <vpc-num>" for the consistency reason of down vpc and for type-2 consistency reasons for any vpc.
Configuration Parameters That Should Be Identical
Configure the following parameters identically on both vPC peer devices to prevent misconfiguration and undesirable traffic behavior.
- 
                                          						
                                          MAC aging timers
- 
                                          						
                                          Static MAC entries
- 
                                          						
                                          VLAN interface—Each device on the end of the vPC Peer-Link must have a VLAN interface configured for the same VLAN on both ends and they must be in the same administrative and operational mode. Those VLANs configured on only one device of the vPC Peer-Link do not pass traffic using the vPC or vPC Peer-Link. You must create all VLANs on both the primary and secondary vPC devices, or the VLAN will be suspended.
- 
                                          						
                                          All ACL configurations and parameters
- 
                                          						
                                          Quality of Service (QoS) configuration and parameters
- 
                                          						
                                          STP interface settings: 
  - 
                                                								
                                                BPDU Filter
  - 
                                                								
                                                BPDU Guard
  - 
                                                								
                                                Cost
  - 
                                                								
                                                Link type
  - 
                                                								
                                                Priority
  - 
                                                								
                                                VLANs (Rapid PVST+)
- 
                                                								
                                                
- 
                                          						
                                          Port security
- 
                                          						
                                          Cisco Trusted Security (CTS)
- 
                                          						
