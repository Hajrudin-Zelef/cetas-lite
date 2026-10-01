---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0-8
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0.md
source_anchor: ""
source_lines: [516, 571]
sha256: 7ceb743286a6a9d4c3db33e53ae769514f07110f01398c7a602901aa6d59aea8
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-102x-configuration-interfaces-cisco-nexus-90-8f5ebdf0

                                       VLANs (PVRST+)
| Note | Display the configuration on both sides of the vPC Peer-Link to ensure that the settings are identical. | 
You can use the show spanning-tree command to display information about the vPC when that feature is enabled. See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for an example.
| Note | We recommend that you configure the ports on the downstream devices as STP edge ports. You should configure all host ports connected to a switch as STP edge ports. See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for more information about STP port types. | 
vPC Peer Switches
The vPC peer switch feature was added to Cisco NX-OS to address performance concerns around STP convergence. This feature allows a pair of Cisco Nexus 9000 Series devices to appear as a single STP root in the Layer 2 topology. This feature eliminates the need to pin the STP root to the vPC primary switch and improves vPC convergence if the vPC primary switch fails.
To avoid loops, the vPC Peer-Link is excluded from the STP computation. In vPC peer switch mode, STP BPDUs are sent from both vPC peer devices to avoid issues related to STP BPDU timeout on the downstream switches, which can cause traffic disruption.
This feature can be used with the pure peer switch topology in which the devices all belong to the vPC.
Additional Information
| Note | Peer-switch feature is supported on networks that use vPC and STP-based redundancy is not supported. If the vPC Peer-Link fail in a hybrid peer-switch configuration, you can lose traffic. In this scenario, the vPC peers use the same STP root ID as well as the same bridge ID. The access switch traffic is split in two with half going to the first vPC peer and the other half to the second vPC peer. With vPC Peer-Link failure, there is no impact to the north/south traffic but the east/west traffic is lost. | 
See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for information about STP enhancement features and Rapid PVST+.
vPC Peer-Gateways
You can configure vPC peer devices to act as the gateway even for packets that are destined to the vPC peer device’s MAC address.
Use the peer-gateway command to configure this feature.
| Note |  | 
Some network-attached storage (NAS) devices or load balancers might have features that help to optimize the performances of particular applications. These features enable the device to avoid a routing-table lookup when responding to a request that originated from a host that is not locally attached to the same subnet. Such devices might reply to traffic using the MAC address of the sender Cisco Nexus 9000 Series device rather than the common HSRP gateway. This behavior is noncomplaint with some basic Ethernet RFC standards. Packets that reach a vPC device for the nonlocal router MAC address are sent across the vPC Peer-Link and could be dropped by the built in vPC loop avoidance mechanism if the final destination is behind another vPC.
The vPC peer-gateway capability allows a vPC switch to act as the active gateway for packets that are addressed to the router MAC address of the vPC peer. This feature enables local forwarding of packets without the need to cross the vPC Peer-Link. In this scenario, the feature optimizes use of the vPC Peer-Link and avoids potential traffic loss.
Configuring the peer-gateway feature must be done on both primary and secondary vPC peers and is nondisruptive to the operations of the device or to the vPC traffic. The vPC peer-gateway feature can be configured globally under the vPC domain submode.
When you enable this feature, Cisco NX-OS automatically disables IP redirects on all interface VLANs mapped over a vPC VLAN to avoid generation of IP redirect messages for packets switched through the peer gateway router.
Packets that arrive at the peer-gateway vPC device have their Time to Live (TTL) decremented, so that packets carrying a TTL of 1 might get dropped in transit due to TTL expiration. You should take this situation into account when the peer-gateway feature is enabled and particular network protocols that source packets with a TTL of 1 operate on a vPC VLAN.
vPC and ARP or ND
A feature was added to Cisco NX-OS to address table synchronization across vPC peers using the reliable transport mechanism of the Cisco Fabric Service over Ethernet (CFSoE) protocol. You must enable the ip arp synchronize and ipv6 nd synchronize commands to support faster convergence of address tables between the vPC peers. This convergence overcomes the delay that occurs in ARP table restoration for IPv4 or ND table restoration for IPv6 when the vPC Peer-Link port channel flaps or when a vPC peer comes back online.
vPC Multicast—PIM, IGMP, and IGMP Snooping
The Cisco NX-OS software for the Nexus 9000 Series devices supports the following on a vPC:
- 
                                       					
                                       PIM Any Source Multicast (ASM).
- 
                                       					
                                       					
                                       PIM Source-Specific Multicast (SSM) .
Multicast Synchronization and Traffic Flow
The software keeps the multicast forwarding state synchronized on both of the vPC peer devices. The IGMP snooping process on a vPC peer device shares the learned group information with the other vPC peer device through the vPC Peer-Link; the multicast states are always synchronized on both vPC peer devices. The PIM process in vPC mode ensures that only one of the vPC peer devices forwards the multicast traffic to the receivers.
Each vPC peer is a Layer 2 or Layer 3 device. Multicast traffic flows from only one of the vPC peer devices. You might see duplicate packets in the following scenarios:
- 
                                       					
                                       Orphan hosts
- 
                                       					
                                       When the source and receivers are in the Layer 2 vPC cloud in different VLANs with multicast routing enabled and a vPC member link goes down.
You might see negligible traffic loss in the following scenarios:
- 
                                       					
                                       When you reload the vPC peer device that is forwarding the traffic.
- 
                                       					
                                       When you restart PIM on the vPC peer device that is forwarding the traffic.
| Note | The Cisco NX-OS software does not support Bidirectional (BIDR) on a vPC. | 
Overall multicast convergence times are scale and vPC role change / PIM restart duration dependent.
Ensure that you dual-attach all Layer 3 devices to both vPC peer devices. If one vPC peer device goes down, the other vPC peer device continues to forward all multicast traffic normally.
The following outlines vPC PIM and vPC IGMP/IGMP snooping:
- 
                                       					
                                       vPC PIM—The PIM process in vPC mode ensures that only one vPC peer device forwards multicast traffic. The PIM process in vPC mode synchronizes the source state with both vPC peer devices and elects which vPC peer device forwards the traffic.
- 
                                       					
