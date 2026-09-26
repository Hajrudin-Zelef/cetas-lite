---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-34-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-34.md
source_anchor: ""
source_lines: [15, 73]
sha256: 42a2443e6f977ff41307705583c1021a7a192c4fc92b24d3a9c791976139c2f0
---

# Introduction

| Property | Description | 
|---|---|
| **allow-fast-path** (*yes \| no* ; Default:**yes** ) | Whether to allow Fast Path processing. Fragmented and flooded packets over VXLAN are redirected via a slow path. Fast Path is disabled for VXLAN interface that uses VRF. The setting is available since RouterOS version 7.8. | 
| **arp** (*disabled \| enabled \| local-proxy-arp \| proxy-arp \| reply-only* ; Default:**enabled** ) | Address Resolution Protocol setting  | 
| **arp-timeout** (*auto \| integer* ; Default:**auto** ) | How long the ARP record is kept in the ARP table after no packets are received from IP. Value `auto` equals to the value of`arp-timeout` in IP/Settings, default is the 30s. | 
| **bridge** (*name* ; Default: ) | Name of the bridge interface to which VXLAN interface will be added as a slave port. | 
| **bridge-pvid** (*integer 1..4094;* Default:**1** ) | Used to assign PVID parameter for dynamically bridge port. This property only has an effect when bridge vlan-filtering is set to yes. | 
| **checksum**  (*yes \| no* ; Default:**no** ) | Setting controls whether a UDP checksum is calculated in the transmitted outer VXLAN packets:  If hardware offloading is used for packet transmission, this setting is ignored, and the behavior defaults to sending packets with a zero UDP checksum. | 
| **comment** (*string* ; Default: ) | Short description of the interface. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Changes whether the interface is disabled. | 
| **dont-fragment**  (*auto \| disabled \| enabled \| inherit* ; Default:**auto** ) | The Don't Fragment (DF) flag controls whether a packet can be broken into smaller packets, called fragments, before being sent over a network. When configuring VXLAN, this setting determines the presence of the DF flag on the outer IPv4 header and can control packet fragmentation if the encapsulated packet exceeds the outgoing interface MTU. This setting has three options:  The setting is available since RouterOS version 7.8. | 
| **group** (*IPv4 \| IPv6* ; Default: ) | When specified, a multicast group address can be used to forward broadcast, unknown-unicast, and multicast traffic between VTEPs. This property requires specifying the `interface` setting. The interface will use IGMP or MLD to join the specified multicast group, make sure to add the necessary PIM and IGMP/MDL configuration. When this property is set, the`vteps-ip-version` automatically gets updated to the used multicast IP version. Disables hardware offloading on compatible devices. | 
| **hw** (*yes \| no* ; Default:**yes** ) | Allows to disable hardware offloading, only applies to devices that support VXLAN offloading. | 
| **interface** (*name* ; Default: ) | Interface name used for multicast forwarding. This property requires specifying the `group` setting. Disables hardware offloading on compatible devices. | 
| **learning**  (*yes \| no* ; Default:**yes** ) | Setting controls whether inner source MAC addresses and remote VTEP IP/IPv6 addresses are learned dynamically from received packets. | 
| **local-address** (*IPv4 \| IPv6* ; Default: ) | Specifies the local source address for the VXLAN interface. If not set, one IP address of the egress interface will be selected as a source address for VXLAN packets. When the property is set, the `vteps-ip-version` automatically gets updated to the used local IP version. The setting is available since RouterOS version 7.7. | 
| **mac-address** (*MAC* ; Default: ) | Static MAC address of the interface. A randomly generated MAC address will be assigned when not specified. | 
| **max-fdb-size** (*integer: 1*..65535** ; Default:**4096** ) | Limits the maximum number of MAC addresses that VXLAN can store in the forwarding database (FDB). | 
| **mtu** (*integer* ; Default:**1500** ) | For the maximum transmission unit, the VXLAN interface will set MTU to 1500 by default. The `l2mtu` will be set automatically according to the associated`interface` (subtracting 50 bytes corresponding to the VXLAN header). If no interface is specified, the`l2mtu` value of 65535 is used. The`l2mtu` cannot be changed. | 
| **name** (*text* ; Default:**vxlan1** ) | Name of the interface. | 
| **port** (*integer: 1*..65535** ; Default:**4789** ) | Used UDP port number for listening and sending packets to remote VTEPs. | 
| **rem-csum** (*both \| none \| rx \| tx* ; Default:**none** ) | Changes the Remote Checksum Offload (RCO) settings for VXLAN interface. RCO is a technique for eliding the inner checksum of an encapsulated datagram, allowing the outer checksum to be offloaded by network driver. It does, however, involve a change to the encapsulation protocols, which the receiver must also support. For this reason, it is disabled by default and setting is available to ensure compatibility with systems that rely on this feature. RCO is detailed in the following Internet-Drafts: Remote checksum offload for VXLAN, draft-herbert-vxlan-rco-00. Remote checksum offload for encapsulation, draft-herbert-remotecsumoffload-00. If hardware offloading is used, this setting is ignored, and the behavior defaults to none. | 
| **ttl** (*auto \| integer: 0*..255** ; Default:**auto** ) | Specifies the TTL value to use in outgoing packets. By default, the TTL is set to 64 when using the `auto` option. However, if VXLAN is using a multicast underlay network, the default TTL is set to 1. If the multicast network involves routing, you will need to increase the TTL to a higher value. | 
| **vni** (*integer: 1..16777216* ; Default: ) | VXLAN Network Identifier (VNI). | 
| **vtep-vrf** (*name* ; Default:**main** ) | Set VRF for the VXLAN interface on which the VTEPs listen and make connections. VRF is not supported when using `interface` and multicast`group` settings. The same UDP`port` cannot be used in multiple routing tables at the same time. When using a VRF that is not set as the "main", hardware offloading is disabled on compatible devices. The setting is available since RouterOS version 7.7. | 
| **vteps-ip-version** (*ipv4 \| ipv6* ; Default:**ipv4** ) | Used IP protocol version for statically configured VTEPs. The RouterOS VXLAN interface does not support dual-stack, any configured remote VTEPs with the opposite IP version will be ignored. When multicast `group` or`local-address` properties are set, the`vteps-ip-version` automatically gets updated to the used IP version. Using IPv6 disables hardware offloading on compatible devices. The setting is available since RouterOS version 7.6. | 

**Sub-menu:** `/interface vxlan vteps`

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Short description of the configured VTEP. | 
| **interface** (*name* ; Default: ) | Name of the VXLAN interface. | 
| **remote-ip** (*IPv4 \| IPv6* ; Default: ) | Defines the VTEP endpoint IPv4 or IPv6 address which is used when VXLAN interface needs to send BUM (broadcast, unknown-unicast, multicast) traffic. It is not used as access control. | 

# Forwarding table

Since RouterOS version 7.9, it is possible to monitor the learned MAC addresses from remote VTEPs.

**Sub-menu:** `/interface vxlan fdb`

| Property | Description | 
|---|---|
| **interface** (*read-only: *name** ) | Name of the VXLAN interface. | 
| **mac-address** (*read-only: MAC address* ) | MAC address. | 
| **remote-ip** (*read-only: IPv4 \| IPv6 address* ) | The IPv4 or IPv6 destination address of remote VTEP. | 

# Configuration example

This configuration example creates a single VXLAN tunnel between two statically configured VTEP endpoints.

First, create VXLAN interfaces on both routers.

Then configure VTEPs on both routers with respective IPv4 destination addresses. Both devices should have an active route toward the destination address.

The configuration is complete. It is possible to include the VXLAN interface into a bridge with other Ethernet interfaces.

# Hardware offloaded VXLAN

