---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0-2
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0.md
source_anchor: ""
source_lines: [89, 136]
sha256: 1de0a5a4d3c97dbf87288bee3a198457f2b1555063b3c864f14339b17b49d14b
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0

                                       Recovery—When the peer-keepalive and the vPC Peer-Link come up, one switch becomes the secondary vPC. On the switch that becomes the secondary vPC, the vPC links go down and come back up.
vPC Peer-Links
You can have only two devices as vPC peers; each device can serve as a vPC peer to only one other vPC peer. The vPC peer devices can also have non-vPC links to other devices.
Reference Information
See the following figure for invalid vPC peer configurations.
To make a valid configuration, you first configure a port channel on each device and then configure the vPC domain. You assign the port channel on each device as a vPC Peer-Link, using the same vPC domain ID. For redundancy, we recommend that you should configure at least two of the dedicated ports into the port channel because if one of the interfaces in the vPC Peer-Link fails, the device automatically falls back to use another interface in the vPC Peer-Link.
| Note | We recommend that you configure the Layer 2 port channels in trunk mode. | 
Compatibility and Configuration Consistency
Many operational parameters and configuration parameters must be the same in each device connected by a vPC Peer-Link (see the Compatibility Parameters for vPC Interfaces section). Because each device is completely independent on the management plane, you must ensure that the devices are compatible on the critical parameters. vPC peer devices have separate control planes. After configuring the vPC Peer-Link, you should display the configuration on each vPC peer device to ensure that the configurations are compatible.
| Note | You must ensure that the two devices connected by the vPC Peer-Link have certain identical operational and configuration parameters. For more information on required configuration consistency, see the Compatibility Parameters for vPC Interfaces section. | 
Primary and Secondary Device Roles
When you configure the vPC Peer-Link, the vPC peer devices negotiate that one of the connected devices is the primary device and the other connected device is the secondary device (see the “Configuring vPCs” section). By default, the Cisco NX-OS software uses the lowest MAC address to elect the primary device. However, if the role priority is set, then the device with the lowest priority will be elected as the primary device. The software takes different actions on each device—that is, the primary and secondary—only in certain failover conditions. If the primary device fails, the secondary device becomes the new primary device when the system recovers, and the previously primary device is now the secondary device.
You can also configure which of the vPC devices is the primary device. Changing the priority of the vPC peer devices can cause the interfaces in your network to go up and down. If you want to configure the role priority again to make one vPC device the primary device, configure the role priority on both the primary vPC device with a lower priority value and the secondary vPC device with the higher value. Then, shut down the port channel that is the vPC Peer-Link on both devices by entering the shutdown command, and finally reenable the port channel on both devices by entering the no shutdown command.
| Note | We recommend that you use two different modules for redundancy on each vPC peer device on each vPC Peer-Link. | 
Traffic Flow and Load Balancing
The software keeps all traffic that forwards across the vPC peer devices as local traffic. A packet that ingresses the port channel uses one of the local links rather than moving across the vPC Peer-Link. Unknown unicast, multicast, and broadcast traffic (including STP BPDUs) are flooded across the vPC Peer-Link. The software keeps the multicast forwarding state synchronized on both of the vPC peer devices.
You can configure any of the standard load-balancing schemes on both the vPC Peer-Link devices and the downstream device (see the Configuring Port Channels chapter for information about load balancing).
Configuration and MAC Address Synchronization
Configuration information flows across the vPC Peer-Links using the Cisco Fabric Services over Ethernet (CFSoE) protocol. (See the section for more information about CFSoE.)
All MAC addresses for those VLANs configured on both devices are synchronized between vPC peer devices. The software uses CFSoE for this synchronization. (See the section for information about CFSoE.)
vPC Peer-Link Failure and Peer-Keepalive
If the vPC Peer-Link fails, the software checks the status of the remote vPC peer device using the peer-keepalive link, which is a link between vPC peer devices that ensures that both devices are up. If the vPC peer device is up, the secondary vPC device disables all vPC ports on its device, to prevent loops and disappearing or flooding traffic. The data then forwards down the remaining active links of the port channel.
The software learns of a vPC peer device failure when the keepalive messages are not returned over the peer-keepalive link.
Use a separate link (vPC peer-keepalive link) to send configurable keepalive messages between the vPC peer devices. The keepalive messages on the vPC peer-keepalive link determines whether a failure is on the vPC Peer-Link only or on the vPC peer device. The keepalive messages are used only when all the links in the vPC Peer-Link fail. See the “Peer-Keepalive Link and Messages” section for information about the keepalive message.
Features That You Must Manually Configure on the Primary and Secondary Devices
You must manually configure the following features to conform to the primary/secondary mapping of each of the vPC peer devices.
STP Root Configuration
STP root—Configure the primary vPC peer device as the STP primary root device and configure the vPC secondary device to be the STP secondary root device. See the “vPC Peer-Links and STP” section for more information about vPCs and STP.
- 
                                       					
                                       We recommend that you configure the vPC Peer-Link interfaces as STP network ports so that Bridge Assurance is enabled on all vPC Peer-Links.
- 
                                       					
                                       We recommend that you configure Rapid per VLAN Spanning Tree plus (PVST+) so that the primary device is the root for all VLANs and configure Multiple Spanning Tree (MST) so that the primary device is the root for all instances.
Layer 3 VLAN Network Interface Configuration
Layer 3 VLAN network interface—Configure Layer 3 connectivity from each vPC peer device by configuring a VLAN network interface for the same VLAN from both devices.
HSRP Active Configuration
HSRP active—If you want to use Hot Standby Router Protocol (HSRP) and VLAN interfaces on the vPC peer devices, configure the primary vPC peer device with the HSRP active highest priority. Configure the secondary device to be the HSRP standby and ensure that you have VLAN interfaces on each vPC device that are in the same administrative and operational mode. (See the “vPC Peer-Links and Routing” section for more information on vPC and HSRP.)
UDLD Configuration Recommendations
While you configure Unidirectional Link Detection (UDLD), note the following recommendations:
- 
                                       					
                                       If LACP is used as port-channel aggregation protocol, UDLD is not required in a vPC domain.
- 
                                       					
                                       If LACP is not used as the port-channel aggregation protocol (static port-channel), use UDLD in normal mode on vPC member ports.
- 
                                       					
