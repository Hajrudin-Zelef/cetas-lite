---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-28
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1497, 1555]
sha256: ee642d3ec5e59ff27a03f0b006f94ede58455c37dcaa3acfca12d65dc92bbbe2
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

Configuring Multicast VRFs
For complete syntax and usage information for the commands, see the switch command reference for this release and the Cisco IOS IP Multicast Command Reference.
For more information about configuring a multicast within a Multi-VRF CE, see the IP Routing: Protocol-Independent Configuration Guide, Cisco IOS Release 15S.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | ip 				  routing Example:  Device(config)# ip routing  | Enables IP routing mode. | 
| Step 3 | ip vrf  				vrf-name Example:  Device(config)# ip vrf vpn1  | Names the VRF, and enter VRF configuration mode. | 
| Step 4 | rd  				route-distinguisher Example:  Device(config-vrf)# rd 100:2  | Creates a VRF table by specifying a route distinguisher. Enter either an AS number and an arbitrary number (xxx:y) or an IP address and an arbitrary number (A.B.C.D:y) | 
| Step 5 | route-target {export \|  				import \|  				both}  				route-target-ext-community Example:  Device(config-vrf)# route-target import 100:2  | Creates a list of import, export, or import and export route target communities for the specified VRF. Enter either an AS system number and an arbitrary number (xxx:y) or an IP address and an arbitrary number (A.B.C.D:y). The route-target-ext-community should be the same as the route-distinguisher entered in Step 4. | 
| Step 6 | import map  				route-map Example:  Device(config-vrf)# import map importmap1  | (Optional) Associates a route map with the VRF. | 
| Step 7 | ip 				  multicast-routing vrf  				vrf-name  				distributed Example:  Device(config-vrf)# ip multicast-routing vrf vpn1 distributed  | (Optional) Enables global multicast routing for VRF table. | 
| Step 8 | interface  				interface-id Example:  Device(config-vrf)# interface gigabitethernet 1/0/2  | Specifies the Layer 3 interface to be associated with the VRF, and enter interface configuration mode. The interface can be a routed port or an SVI. | 
| Step 9 | ip vrf 				  forwarding  				vrf-name Example:  Device(config-if)# ip vrf forwarding vpn1  | Associates the VRF with the Layer 3 interface. | 
| Step 10 | ip address  				ip-address  				mask Example:  Device(config-if)# ip address 10.1.5.1 255.255.255.0  | Configures IP address for the Layer 3 interface. | 
| Step 11 | ip pim 				  sparse-dense mode Example:  Device(config-if)# ip pim sparse-dense mode  | Enables PIM on the VRF-associated Layer 3 interface. | 
| Step 12 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 13 | show ip vrf [brief \|  				detail \|  				interfaces] [vrf-name] Example:  Device# show ip vrf detail vpn1  | Verifies the configuration. Displays information about the configured VRFs. | 
| Step 14 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring a VPN Routing Session
Routing within the VPN can be configured with any supported routing protocol (RIP, OSPF, EIGRP, or BGP) or with static routing. The configuration shown here is for OSPF, but the process is the same for other protocols.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure 				  terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | router ospf 				 				process-id  				vrf  				vrf-name Example:  Device(config)# router ospf 1 vrf vpn1  | Enables OSPF routing, specifies a VPN forwarding table, and enter router configuration mode. | 
| Step 3 | log-adjacency-changes Example:  Device(config-router)# log-adjacency-changes  | (Optional) Logs changes in the adjacency state. This is the default state. | 
| Step 4 | redistribute 				  bgp  				autonomous-system-number  				subnets Example:  Device(config-router)# redistribute bgp 10 subnets  | Sets the switch to redistribute information from the BGP network to the OSPF network. | 
| Step 5 | network  				network-number  				area  				area-id Example:  Device(config-router)# network 1 area 2  | Defines a network address and mask on which OSPF runs and the area ID for that network address. | 
| Step 6 | end Example:  Device(config-router)# end  | Returns to privileged EXEC mode. | 
| Step 7 | show ip 				  ospf  				process-id Example:  Device# show ip ospf 1  | Verifies the configuration of the OSPF network. | 
| Step 8 | copy 				  running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Configuring BGP PE to CE Routing Sessions
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure 				  terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | router bgp  				autonomous-system-number Example:  Device(config)# router bgp 2  | Configures the BGP routing process with the AS number passed to other BGP routers, and enter router configuration mode. | 
| Step 3 | network  				network-number  				mask  				network-mask Example:  Device(config-router)# network 5 mask 255.255.255.0  | Specifies a network and mask to announce using BGP. | 
| Step 4 | redistribute 				  ospf  				process-id  				match internal Example:  Device(config-router)# redistribute ospf 1 match internal  | Sets the switch to redistribute OSPF internal routes. | 
| Step 5 | network  				network-number  				area  				area-id Example:  Device(config-router)# network 5 area 2  | Defines a network address and mask on which OSPF runs and the area ID for that network address. | 
| Step 6 | address-family 				  ipv4 vrf  				vrf-name Example:  Device(config-router)# address-family ipv4 vrf vpn1  | Defines BGP parameters for PE to CE routing sessions, and enter VRF address-family mode. | 
| Step 7 | neighbor  				address  				remote-as  				as-number Example:  Device(config-router)# neighbor 10.1.1.2 remote-as 2  | Defines a BGP session between PE and CE routers. | 
| Step 8 | neighbor  				address  				activate Example:  Device(config-router)# neighbor 10.2.1.1 activate  | Activates the advertisement of the IPv4 address family. | 
| Step 9 | end Example:  Device(config-router)# end  | Returns to privileged EXEC mode. | 
| Step 10 | show ip bgp [ipv4] [neighbors] Example:  Device# show ip bgp ipv4 neighbors  | Verifies BGP configuration. | 
| Step 11 | copy 				  running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Monitoring Multi-VRF CE
| Table 15 Commands for Displaying Multi-VRF CE Information |  | 
|---|---|
| show ip protocols vrf vrf-name | Displays routing protocol information associated with a VRF. | 
| show ip route vrf vrf-name [connected] [protocol [as-number]] [list] [mobile] [odr] [profile] [static] [summary] [supernets-only] | Displays IP routing table information associated with a VRF. | 
| show ip vrf [brief \| detail \| interfaces] [vrf-name] | Displays information about the defined VRF instances. | 
For more information about the information in the displays, see the Cisco IOS Switching Services Command Reference, Release 12.4.
Configuration Examples for Multi-VRF CE
Multi-VRF CE Configuration Example
OSPF is the protocol used in VPN1, VPN2, and the global network. BGP is used in the CE to PE connections. The examples following the illustration show how to configure a switch as CE Switch A, and the VRF configuration for customer switches D and F. Commands for configuring CE Switch C and the other customer switches are not included but would be similar. The example also includes commands for configuring traffic to Switch A for a Catalyst 6000 or Catalyst 6500 switch acting as a PE router.
On Switch A, enable routing and configure VRF.
Device# configure terminal 
Enter configuration commands, one per line.  End with CNTL/Z.
Device(config)# ip routing
