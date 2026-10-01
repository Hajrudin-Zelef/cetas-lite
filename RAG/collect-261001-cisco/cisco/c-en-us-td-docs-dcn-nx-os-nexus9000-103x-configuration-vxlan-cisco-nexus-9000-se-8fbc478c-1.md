---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c-1
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c.md
source_anchor: ""
source_lines: [1, 158]
sha256: 53a013da103d3187df3ac8d0b84b93927a71208fdd69219ae5b07a7d9f04711a
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-8fbc478c

Configuring Features
                           
                           Ensure the vPC Fabric Peering DSCP value is consistent on both vPC member switches. Ensure that the corresponding QoS policy matches the vPC Fabric Peering DSCP marking.
                           
                           
                           All VLANs that require communication traversing the vPC Fabric Peering must have a VXLAN enabled (vn-segment); this includes the native VLAN. 
                           
                           
                           
                              
                                 |  Note |  For MSTP, VLAN 1 must be extended across vPC Fabric Peering if the peer-link and vPC legs have the default native VLAN configuration.                                           This behavior can be achieved by extending VLAN 1 over VXLAN (vn-segment). If the peer-link and vPC legs have non-default                                           native VLANs, those VLANs must be extended across vPC Fabric Peering by associating the VLANs with VXLAN (vn-segment).                                          | 
                           
Use the show vpc virtual-peerlink vlan consistency  command for verification of the existing VLAN-to-VXLAN mapping used for vPC Fabric Peering.
                           
                           
                           peer-keepalive  command for vPC Fabric Peering is supported with one of the following configurations:
                           
                           
                           
                           
                           
                           Example uses OSPF as the underlay routing protocol. 
                           configure terminal
nv overlay evpn
feature ospf
feature bgp
feature pim
feature interface-vlan
feature vn-segment-vlan-based
feature vpc
feature nv overlay
                           
                        
                           vPC Configuration
                           
                           
                           
                              
                                 |  Note |  To change the vPC Fabric Peering source or destination IP, the vPC domain must be shutdown prior to modification. The vPC domain can be returned to operation                                           after the modifying by using the no shutdown  command.                                           | 
                           
Configuring TCAM Carving
                           hardware access-list tcam region ing-racl 0
hardware access-list tcam region ing-sup 768
hardware access-list tcam region ing-flow-redirect 512
                           
                              
                                 |  Note |                                                                                                                                                                                    When configuring fabric vPC peering, the minimum size for Ingress-Flow-redirect TCAM region size is 512. Also ensure that                                                 the TCAM region size is always configured  in multiples of 512.                                                                                                                                                                                         TCAM carving for ing-flow-redirect region is only required on Cisco Nexus 9300-EX, 9300-FX, 9300-FX2, 9300-FX3, and 9364C platform switches.                                                                                                                                           Switch reload is required for the TCAM carving to take effect.  | 
                           
Configuring the vPC Domain
                           
                           For IPv4
                           vpc domain 100
peer-keepalive destination 192.0.2.1 
virtual peer-link destination 192.0.2.100 source 192.0.2.20/32 [dscp <dscp-value>] 
Warning: Appropriate TCAM carving must be configured for virtual peer-link vPC
peer-switch
peer-gateway
ip arp synchronize
ipv6 nd synchronize
exit
                           For IPv6
                           vpc domain 100
peer-keepalive destination 192:0:2::1 
virtual peer-link destination 192:0:2::100 source 192:0:2::20/32 [dscp <dscp-value>] 
Warning: Appropriate TCAM carving must be configured for virtual peer-link vPC
peer-switch
peer-gateway
ipv6 arp synchronize
ipv6 nd synchronize
exit
                           
                              
                                 |  Note |  The dscp  keyword in optional. Range is 1 to 63. The default value is 56.                                          | 
                           
Configuring vPC Fabric Peering Port Channel
                           
                           No need to configure members for the following port channel.
                           interface port-channel 10
switchport
switchport mode trunk
vpc peer-link 
                           interface loopback0
                           
                              
                                 |  Note |  This loopback is not the NVE source-interface loopback (interface used for the VTEP IP address).   | 
                           
For IPv4
                           interface loopback 0
ip address 192.0.2.20/32
ip router ospf 1 area 0.0.0.0
                           For IPv6
                           interface loopback 0
ipv6 address 192:0:2::20/32
ipv6 router ospfv3 1 area 0.0.0.0
                           
                              
                                 |  Note |  You can use the loopback for BGP peering or a dedicated loopback. This lookback must be different that the loopback for peer                                           keep alive.                                           | 
                           
Configuring the Underlay Interfaces
                           
                           
                           Both L3 physical and L3 port channels are supported. SVI and sub-interfaces are not supported. 
                           
                           For IPv4
                           router ospf 1
interface Ethernet1/16
port-type fabric 
ip address 192.0.2.2/24
ip router ospf 1 area 0.0.0.0
no shutdown
interface Ethernet1/17
port-type fabric 
ip address 192.0.2.3/24
ip router ospf 1 area 0.0.0.0
no shutdown
interface Ethernet1/40
port-type fabric 
ip address 192.0.2.4/24
ip router ospf 1 area 0.0.0.0
no shutdown
interface Ethernet1/41
port-type fabric 
ip address 192.0.2.5/24
ip router ospf 1 area 0.0.0.0
no shutdown
                           For IPv6
                           router ospfv3 1
interface Ethernet1/16
port-type fabric
ipv6 address 192:0:2::2/24
ipv6 router ospfv3 1 area 0.0.0.0
no shutdown
interface Ethernet1/17
port-type fabric 
ipv6 address 192:0:2::3/24
ipv6 router ospfv3 1 area 0.0.0.0
no shutdown
interface Ethernet1/40
port-type fabric 
ipv6 address 192:0:2::4/24
ipv6 router ospfv3 1 area 0.0.0.0
no shutdown
interface Ethernet1/41
port-type fabric 
ipv6 address 192:0:2::5/24
ipv6 router ospfv3 1 area 0.0.0.0
no shutdown
                           
                              
                                 |  Note |  All ports connected to spines must be port-type fabric.  | 
                           
VXLAN Configuration
                           
                           
                           
                              
                                 |  Note |  Configuring advertise virtual-rmac  (NVE) and advertise-pip  (BGP) are required steps. For more information, see the Configuring vPC Multi-Homing chapter.   | 
                           
