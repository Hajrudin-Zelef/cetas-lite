---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0-5
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0.md
source_anchor: ""
source_lines: [207, 347]
sha256: 0e587c574fdc950f22bbaa222721c1dd1cb9266aac7f2f720458a61d352522f5
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0

                                          Port-channel mode: on, off, or active (port-channel mode can, however, be active/passive on each side of the vPC peer)
- 
                                          						
                                          The [no] lacp suspend-individual pxe configuration must be the same on both the sides of the vPC.
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
                                          						
                                          Dynamic Host Configuration Protocol (DHCP) snooping
- 
                                          						
                                          Network Access Control (NAC)
- 
                                          						
                                          Dynamic ARP Inspection (DAI)
- 
                                          						
                                          IP source guard (IPSG)
- 
                                          						
                                          Internet Group Management Protocol (IGMP) snooping
- 
                                          						
                                          Hot Standby Routing Protocol (HSRP)
- 
                                          						
                                          Protocol Independent Multicast (PIM)
- 
                                          						
