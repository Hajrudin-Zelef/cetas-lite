---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0-4
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0.md
source_anchor: ""
source_lines: [152, 206]
sha256: a712738fa73038a48838e5d435fde5192b23a15c6a8192bafcad269b0d528624
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0

                                    During the timeout, the vPC secondary device takes action to become the vPC primary device if no keepalive message is received by the end of the configured interval.
See the “Configuring vPC Keepalive Link and Messages” section for information about configuring the timer for the keepalive messages.
| Note | Ensure that both the source and destination IP addresses used for the peer-keepalive messages are unique in your network and these IP addresses are reachable from the VRF associated with the vPC peer-keepalive link. Peer-keepalive IP addresses must be global unicast addresses. Link-local addresses are not supported. | 
Configuring Trusted Ports for Peer-Keepalive
Use the command-line interface (CLI) to configure the interfaces you are using the vPC peer-keepalive messages as trusted ports. Leave the precedence at the default (6) or configure it higher.
vPC Domains
You can use the vPC domain ID to identify the vPC Peer-Links and the ports that are connected to the vPC downstream devices.
The vPC domain is also a configuration mode that you use to configure the keepalive messages and other vPC Peer-Link parameters rather than accept the default values. See the “Configuring vPCs” section for more information about configuring these parameters.
vPC Domain Creation and Peer-Link Configuration
To create a vPC domain, you must first create a vPC domain ID on each vPC peer device using a number from 1 to 1000. You can have only one vPC domain per vPC peer.
You must explicitly configure the port channel that you want to act as the vPC Peer-Link on each device. You associate the port channel that you made a vPC Peer-Link on each device with the same vPC domain ID to form a single vPC domain. Within this domain, the system provides a loop-free topology and Layer 2 multipathing.
You can only configure these port channels and vPC Peer-Links statically. You can configure the port channels and vPC Peer-Links either using LACP or no protocol. We recommend that you use LACP with the interfaces in active mode to configure port channels in each vPC, which ensures an optimized, graceful recovery in a port-channel failover scenario and provides configuration checks against configuration mismatches among the port channels themselves.
vPC System MAC Address Assignment
The vPC peer devices use the vPC domain ID that you configure to automatically assign a unique vPC system MAC address. Each vPC domain has a unique MAC address that is used as a unique identifier for the specific vPC-related operations, although the devices use the vPC system MAC addresses only for link-scope operations, such as LACP. We recommend that you create each vPC domain within the contiguous Layer 2 network with a unique domain ID. You can also configure a specific MAC address for the vPC domain, rather than having the Cisco NX-OS software assign the address.
See the “vPC and Orphan Ports” section for more information about displaying the vPC MAC table.
vPC Domain System Priority
After you create a vPC domain, the Cisco NX-OS software creates a system priority for the vPC domain. You can also configure a specific system priority for the vPC domain.
| Note | When manually configuring the system priority, you must ensure that you assign the same priority value on both vPC peer devices. If the vPC peer devices have different system priority values, vPC does not come up. | 
How VPC Topologies Work
In both topologies, port channels P020 and P0200 must be configured identically on the peer switches and configuration synchronization is used to synchronize the configurations of the vPC switches.
Summary
This document describes two common vPC topologies: a basic configuration with directly connected Cisco Nexus 9000 Series devices and a configuration involving Fabric Extenders (FEXs) for host vPC.
Workflow
These stages describe how vPC topologies work.
- The first topology shows a basic configuration in which the Cisco Nexus 9000 Series device ports are directly connected to another switch or host and are configured as part of a port channel that becomes part of a vPC. In this configuration, vPC 20 is configured on port channel 20, which has Eth1/10 on the first device and Eth2/1 on the second as member ports.
- The second topology illustrates how to configure a vPC from the peer devices through Fabric Extenders (FEXs). In this FEX straight-through topology, each FEX is single-homed with a Cisco Nexus 9000 Series device. The host interfaces on this FEX are configured as port channels, and those port channels are configured as vPCs. For example, Eth101/1/1 and Eth102/1/5 are configured as members of PO200, and PO200 is configured for vPC 200.
What’s next
See the Cisco Nexus 2000 Series NX-OS Fabric Extender Configuration Guide for Cisco Nexus 9000 Series Switches for more information about configuring FEX ports.
Compatibility Parameters
Many configuration and operational parameters must be identical on all interfaces in the vPC. We recommend that you configure the Layer 2 port channels that you use for the vPC Peer-Link in trunk mode.
After you enable the vPC feature and configure the vPC Peer-Link on both vPC peer devices, Cisco Fabric Services (CFS) messages provide a copy of the configuration on the local vPC peer device configuration to the remote vPC peer device. The system then determines whether any of the crucial configuration parameters differ on the two devices. (See the “vPC and Orphan Ports” section for more information about CFS.)
The vPC Peer-Link is a core component of vPC functionality, requiring consistent configuration across both peer devices.
- 
                                    					
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
                                          						
