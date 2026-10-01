---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-33
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [5812, 6099]
sha256: d356b5a0ef10094efe5e1a4c4599940561422235a8e6b29106c15e5f02f430f5
---

# Configure DeviceA to generate a local key pair.

 authentication-mode hmac-sha256 password %+%##!!!!!!!!!"!!!!"!!!!*!!!!C+tR0CW9x*eB&pWp`t),Azgw-h\o8#4LZPD!!!!!!!!!!!!!!!9!!!!>fwJ)I0E{=:%,*,XRhbH&t0MCy_8=7!!!!!!!!!!%+%#
 dfs-group state switchover disable
#
vlan batch 10 to 11
#
access-user m-lag enable
#
radius-server template rd1
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
 radius-server authentication 192.168.10.1 1812 source ip-address 192.168.1.1 weight 80
 radius-server accounting 192.168.10.1 1813 source ip-address 192.168.1.1 weight 80
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme scheme2
  accounting-mode radius
  accounting realtime 15
 domain example.com
  authentication-scheme abc
  accounting-scheme scheme2
  radius-server rd1
#
dot1x-access-profile name d1
 dot1x timer client-timeout 30
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example.com force
 authentication mode multi-authen max-user 100
#
stp bridge-address 00e0-fc12-3458
stp instance 0 root primary
#
interface MEth0/0/0
 ip address 10.1.1.1 255.255.255.0
#
interface Vlanif10
 ip address 192.168.1.1 255.255.255.0
 mac-address 0000-0000-0011
#
interface Eth-Trunk0
 mode lacp-static
 stp disable
 peer-link 1
#
interface Eth-Trunk1
 port link-type access
 port default vlan 11
 stp edged-port enable
 mode lacp-static
 dfs-group 1 m-lag 1
 authentication-profile p1
#
interface Eth-Trunk2
 port link-type trunk
 port trunk allow-pass vlan 10 11
 mode lacp-static 
#
interface 10GE1/0/1
 eth-trunk 2
#
interface 10GE1/0/2
 eth-trunk 1
#
interface 10GE1/0/3
 eth-trunk 0
#
interface 10GE1/0/4
 eth-trunk 0
#
interface 10GE1/0/5
 eth-trunk 1
#
interface 10GE1/0/6
 eth-trunk 2
#
return
DeviceB
#
sysname DeviceB
#
dfs-group 1
 priority 120
 dual-active detection source ip 10.1.1.2 peer 10.1.1.1
 authentication-mode hmac-sha256 password %+%##!!!!!!!!!"!!!!"!!!!*!!!!=I9f8>C{!P_bhB31@7r-=jrS8c|_"(Bn~#=!!!!!!!!!!!!!!!9!!!!kx-6@.tGA(wAt/IQXl6>[g{6YlOi9$!!!!!!!!!!%+%#
 dfs-group state switchover disable
#
vlan batch 10 to 11
#
access-user m-lag enable
#
radius-server template rd1
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
 radius-server authentication 192.168.10.1 1812 source ip-address 192.168.1.1 weight 80
 radius-server accounting 192.168.10.1 1813 source ip-address 192.168.1.1 weight 80
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme scheme2
  accounting-mode radius
  accounting realtime 15
 domain example.com
  authentication-scheme abc
  accounting-scheme scheme2
  radius-server rd1
#
dot1x-access-profile name d1
 dot1x timer client-timeout 30
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example.com force
 authentication mode multi-authen max-user 100
#
stp bridge-address 00e0-fc12-3458
stp instance 0 root primary
#
interface MEth0/0/0
 ip address 10.1.1.2 255.255.255.0
#
interface Vlanif10
 ip address 192.168.1.1.1 255.255.255.0
 mac-address 0000-0000-0011
#
interface Eth-Trunk0
 mode lacp-static
 stp disable
 peer-link 1
#
interface Eth-Trunk1
 port link-type access
 port default vlan 11
 stp edged-port enable
 mode lacp-static
 dfs-group 1 m-lag 1
 authentication-profile p1
#
interface Eth-Trunk2
 port link-type trunk
 port trunk allow-pass vlan 10 11
 mode lacp-static
#
interface 10GE1/0/1
 eth-trunk 2
#
interface 10GE1/0/2
 eth-trunk 1
#
interface 10GE1/0/3
 eth-trunk 0
#
interface 10GE1/0/4
 eth-trunk 0
#
interface 10GE1/0/5
 eth-trunk 1
#
interface 10GE1/0/6
 eth-trunk 2
#
return
DeviceC
#
sysname DeviceC
#
vlan batch 10 to 11
#
interface Eth-Trunk2
 port link-type trunk
 port trunk allow-pass vlan 10 11
 mode lacp-static
#
interface 10GE1/0/1
 eth-trunk 2
#
interface 10GE1/0/2
 eth-trunk 2
#
return
DeviceD
#
sysname DeviceD
#
vlan batch 10 to 11
#
interface Eth-Trunk2
 port link-type trunk
 port trunk allow-pass vlan 10 11
 mode lacp-static
#
interface 10GE1/0/1
 eth-trunk 2
#
interface 10GE1/0/2
 eth-trunk 2
#
return
#
sysname DeviceE
#
vlan batch 11 
#
l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
#
interface Eth-Trunk2
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type access
 port default vlan 11
#
interface 10GE 1/0/2
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type access
 port default vlan 11
#
return
In Figure 3-147, user PCs are connected to the network through DeviceD. DeviceA and DeviceB set up an M-LAG, are connected to DeviceC in the uplink direction, and are connected to DeviceD in the downlink direction.
In addition, to meet high security requirements, Portal authentication is used for access control of user PCs through the RADIUS server, and the authentication point is deployed on M-LAG member devices (DeviceA and DeviceB).
In this example, interfaces 1, 2, 3, 4, 5, and 6 on DeviceA represent 10GE 1/0/1, 10GE 1/0/2, 10GE 1/0/3, 10GE 1/0/4, 10GE 1/0/5, and 10GE 1/0/6, respectively.
Interfaces 1, 2, 3, 4, 5, and 6 on DeviceB represent 10GE 1/0/1, 10GE 1/0/2, 10GE 1/0/3, 10GE 1/0/4, 10GE 1/0/5, and 10GE 1/0/6, respectively.
On DeviceA and DeviceB, Eth-Trunk 0 serves as a peer-link interface, M-LAG interface Eth-Trunk 1 connects to downlink terminals and is added to VLAN 11, and M-LAG interface Eth-Trunk 2 connects to the uplink server and is added to VLAN 10.
When configuring Portal authentication in an M-LAG scenario:
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] stp mode rstp
[DeviceA] stp v-stp enable
<HUAWEI> system-view
[HUAWEI] sysname DeviceB
[DeviceB] stp mode rstp
[DeviceB] stp v-stp enable
[DeviceA] interface eth-trunk 0
[DeviceA-Eth-Trunk0] mode lacp-static
[DeviceA-Eth-Trunk0] trunkport 10ge 1/0/3
[DeviceA-Eth-Trunk0] trunkport 10ge 1/0/4
[DeviceA-Eth-Trunk0] peer-link 1
[DeviceA-Eth-Trunk0] quit
[DeviceB] interface eth-trunk 0
[DeviceB-Eth-Trunk0] mode lacp-static
[DeviceB-Eth-Trunk0] trunkport 10ge 1/0/3
[DeviceB-Eth-Trunk0] trunkport 10ge 1/0/4
[DeviceB-Eth-Trunk0] peer-link 1
[DeviceB-Eth-Trunk0] quit
[DeviceA] vlan batch 11
[DeviceA] interface eth-trunk 1
[DeviceA-Eth-Trunk1] trunkport 10ge 1/0/2
[DeviceA-Eth-Trunk1] trunkport 10ge 1/0/5
[DeviceA-Eth-Trunk1] mode lacp-static
[DeviceA-Eth-Trunk1] port link-type access
[DeviceA-Eth-Trunk1] port default vlan 11
[DeviceA-Eth-Trunk1] dfs-group 1 m-lag 1
[DeviceA-Eth-Trunk1] quit
[DeviceB] vlan batch 11
[DeviceB] interface eth-trunk 1
[DeviceB-Eth-Trunk1] trunkport 10ge 1/0/2
[DeviceB-Eth-Trunk1] trunkport 10ge 1/0/5
[DeviceB-Eth-Trunk1] mode lacp-static
[DeviceB-Eth-Trunk1] port link-type access
[DeviceB-Eth-Trunk1] port default vlan 11
[DeviceB-Eth-Trunk1] dfs-group 1 m-lag 1
[DeviceB-Eth-Trunk1] quit
DeviceA and DeviceB must be configured with the same virtual IP address and virtual MAC address.
[DeviceA] interface vlanif 11
[DeviceA-Vlanif11] ip address 10.2.1.1 24
[DeviceA-Vlanif11] mac-address 0000-5e00-0101
[DeviceA-Vlanif11] quit
[DeviceB] interface vlanif 11
[DeviceB-Vlanif11] ip address 10.2.1.1 24
[DeviceB-Vlanif11] mac-address 0000-5e00-0101
[DeviceB-Vlanif11] quit
[DeviceA] vlan batch 10
[DeviceA] interface eth-trunk 2
[DeviceA-Eth-Trunk2] mode lacp-static
[DeviceA-Eth-Trunk2] port link-type trunk
[DeviceA-Eth-Trunk2] port trunk allow-pass vlan 10
[DeviceA-Eth-Trunk2] trunkport 10ge 1/0/1
[DeviceA-Eth-Trunk2] trunkport 10ge 1/0/6
[DeviceA-Eth-Trunk2] dfs-group 1 m-lag 2
[DeviceA-Eth-Trunk2] quit
[DeviceB] vlan batch 10
[DeviceB] interface eth-trunk 2
[DeviceB-Eth-Trunk2] mode lacp-static
[DeviceB-Eth-Trunk2] port link-type trunk
[DeviceB-Eth-Trunk2] port trunk allow-pass vlan 10
[DeviceB-Eth-Trunk2] trunkport 10ge 1/0/1
[DeviceB-Eth-Trunk2] trunkport 10ge 1/0/6
[DeviceB-Eth-Trunk2] dfs-group 1 m-lag 2
[DeviceB-Eth-Trunk2] quit
