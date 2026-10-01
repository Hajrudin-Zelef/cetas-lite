---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6-2
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6.md
source_anchor: ""
source_lines: [19, 62]
sha256: 2f73f6477916b5453083dfd4518f955da52e861493d1a3090a5aeba372732d67
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-93x-interfaces-configuration-gu-d98dd8d6

                                    A trunk port carries traffic for two or more VLANs simultaneously.
Access and Trunk Port Details
By default, all ports on Cisco Nexus 9300-EX switches are Layer 3 ports, while all ports on Cisco Nexus 9300 switches are Layer 2 ports. You can configure ports as Layer 2 using the setup script, the system default switchport command, or the switchport command.
Trunk ports require IEEE 802.1Q encapsulation to deliver traffic for multiple VLANs. All ports in a trunk must reside in the same VDC, as trunk ports cannot carry VLANs from different VDCs. If an access port receives a packet with an 802.1Q tag other than the access VLAN, the port drops the packet.
To optimize performance, you can configure an end station port as a host port, which automatically sets it as an access port and disables channel grouping to decrease packet forwarding time. A Layer 2 interface cannot function as both an access and trunk port simultaneously. Changing a Layer 2 interface to a Layer 3 interface removes all Layer 2 configurations and restores default VLAN settings.
| Note | See the Cisco Nexus 9000 Series NX-OS Fundamentals Configuration Guide for setup script details, the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for VLAN information, and the Cisco Nexus 9000 Series NX-OS Unicast Routing Configuration Guide for subinterface information. | 
Only an end station can be set as a host port; you will receive an error message if you attempt to configure other ports as hosts.
IEEE 802.1Q Encapsulation
IEEE 802.1Q encapsulation is a tagging method used on trunk ports to identify which VLAN a frame belongs to, allowing multiple VLANs to traverse the same physical link while maintaining traffic separation.
IEEE 802.1Q Encapsulation Details
| Note | For information about VLANs, see the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide. | 
A trunk is a point-to-point link between the switch and another networking device. Trunks carry the traffic of multiple VLANs over a single link and allow you to extend VLANs across an entire network.
To correctly deliver the traffic on a trunk port with several VLANs, the device uses the IEEE 802.1Q encapsulation, or tagging, method that uses a tag that is inserted into the frame header. This tag carries information about the specific VLAN to which the frame and packet belong. This method allows packets that are encapsulated for several different VLANs to traverse the same port and maintain traffic separation between the VLANs. Also, the encapsulated VLAN tag allows the trunk to move traffic end-to-end through the network on the same VLAN.
Access VLANs
An access VLAN is the VLAN assigned to an access port that determines which VLAN traffic the port carries.
Access VLAN Behavior
When you configure a port in access mode, you can specify which VLAN will carry the traffic for that interface. If you do not configure the VLAN for a port in access mode, or an access port, the interface carries traffic for the default VLAN (VLAN1).
You can change the access port membership in a VLAN by specifying the new VLAN. You must create the VLAN before you can assign it as an access VLAN for an access port. If you change the access VLAN on an access port to a VLAN that is not yet created, the system shuts that access port down.
If an access port receives a packet with an 802.1Q tag in the header other than the access VLAN value, that port drops the packet without learning its MAC source address.
Native VLAN IDs for Trunk Ports
A trunk port can carry nontagged packets simultaneously with the 802.1Q tagged packets. When you assign a default port VLAN ID to the trunk port, all untagged traffic travels on the default port VLAN ID for the trunk port, and all untagged traffic is assumed to belong to this VLAN. This VLAN is referred to as the native VLAN ID for a trunk port. That is, the native VLAN ID is the VLAN that carries untagged traffic on trunk ports.
| Note | Native VLAN ID numbers must match on both ends of the trunk. | 
The trunk port sends an egressing packet with a VLAN that is equal to the default port VLAN ID as untagged; all the other egressing packets are tagged by the trunk port. If you do not configure a native VLAN ID, the trunk port uses the default VLAN.
| Note | You cannot use a Fibre Channel over Ethernet (FCoE) VLAN as a native VLAN for an Ethernet trunk switchport. | 
Tagging Native VLAN Traffic
Tagging native VLAN traffic is a feature that controls whether packets on the native VLAN of an 802.1Q trunk port retain or lose their VLAN tags as they traverse the trunk.
Native VLAN Tagging Behavior
The Cisco software supports the IEEE 802.1Q standard on trunk ports. In order to pass untagged traffic through the trunk ports, you must create a VLAN that does not tag any packets (or you can use the default VLAN). Untagged packets can pass through trunk ports and access ports.
However, all packets that enter the device with an 802.1Q tag that matches the value of the native VLAN on the trunk are stripped of any tagging and egress the trunk port as untagged packets. This situation can cause problems because you may want to retain the tagging on packets on the native VLAN for the trunk port.
You can configure the device to drop all untagged packets on the trunk ports and to retain the tagging of packets entering the device with 802.1Q values that are equal to that of the native VLAN ID. All control traffic still passes on the native VLAN. This configuration is global; trunk ports on the device either do or do not retain the tagging for the native VLAN.
Allowed VLANs
Allowed VLANs are the set of VLANs permitted to send and receive traffic over a trunk port, controlling which VLAN traffic traverses the trunk link.
Managing Allowed VLANs on Trunk Ports
By default, a trunk port sends traffic to and receives traffic from all VLANs. All VLAN IDs are allowed on each trunk. However, you can remove VLANs from this inclusive list to prevent traffic from the specified VLANs from passing over the trunk. Later, you can add any specific VLANs that you may want the trunk to carry traffic for back to the list.
To partition the Spanning Tree Protocol (STP) topology for the default VLAN, you can remove VLAN1 from the list of allowed VLANs. Otherwise, VLAN1, which is enabled on all ports by default, will have a very big STP topology, which can result in problems during STP convergence. When you remove VLAN1, all data traffic for VLAN1 on this port is blocked, but the control traffic continues to move on the port.
| Note | See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for more information about STP. | 
| Note | You can change the block of VLANs reserved for internal use. See the Cisco Nexus 9000 Series NX-OS Layer 2 Switching Configuration Guide for more information about changing the reserved VLANs. | 
Default Interfaces
You can use the default interface feature to clear the configured parameters for both physical and logical interfaces such as the Ethernet, loopback, VLAN network, tunnel, and the port-channel interface.
| Note | A maximum of eight ports can be selected for the default interface. The default interfaces feature is not supported for management interfaces because the device could go to an unreachable state. | 
Switch Virtual Interface and Autostate Behavior
A switch virtual interface (SVI) represents a logical interface between the bridging function and the routing function of a VLAN in the device, with its operational state governed by the STP forwarding state of ports in its corresponding VLAN.
SVI Autostate Behavior
In Cisco NX-OS, a switch virtual interface (SVI) represents a logical interface between the bridging function and the routing function of a VLAN in the device.
