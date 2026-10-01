---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-28
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [4916, 5109]
sha256: d9c37b49d454abf370ed425cc6eac4691c7016105b786d4c0ce555e57d81a389
---

# Configure DeviceA to generate a local key pair.

 ------------------------------------------------------------------------------  
 Total: 2, printed: 2  
# 
sysname SwitchA 
# 
voice-vlan mac-address 00e0-fcc7-0000 mask ffff-ffff-0000 
voice-vlan mac-address 00e0-fc8f-0000 mask ffff-ffff-0000 
# 
vlan batch 100 200 
# 
authentication-profile name ipphone 
 mac-access-profile ipphone 
# 
dhcp enable 
# 
radius-server template ipphone 
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%# 
 radius-server authentication 192.168.100.182 1812 weight 80 
 radius-server accounting 192.168.100.182 1813 weight 80 
# 
aaa 
 authentication-scheme radius 
  authentication-mode radius 
 domain default 
  authentication-scheme radius 
  radius-server ipphone 
# 
interface Vlanif100 
 ip address 10.20.20.1 255.255.255.0 
 dhcp select relay 
 dhcp relay server-ip 10.10.20.2 
# 
interface Vlanif200 
 ip address 10.10.20.1 255.255.255.0 
# 
interface 10GE1/0/1        
 port link-type hybrid 
 voice-vlan 100 enable 
 voice-vlan remark-mode mac-address 
 port hybrid pvid vlan 100 
 port hybrid untagged vlan 100 
 authentication-profile ipphone 
# 
interface  10GE 1/0/2
 port link-type hybrid 
 voice-vlan 100 enable 
 voice-vlan remark-mode mac-address 
 port hybrid pvid vlan 100 
 port hybrid untagged vlan 100 
 authentication-profile ipphone 
# 
interface 10GE 1/0/3        
 port link-type access 
 port default vlan 200 
# 
ip route-static 0.0.0.0 0.0.0.0 10.10.20.2 
# 
mac-access-profile name ipphone 
# 
return     
# 
sysname SwitchB 
# 
vlan batch 200 
# 
dhcp enable 
# 
ip pool ip-phone 
 gateway-list 10.20.20.1  
 network 10.20.20.0 mask 255.255.255.0  
# 
interface Vlanif200 
 ip address 10.10.20.2 255.255.255.0 
 dhcp select global 
# 
interface 10GE 1/0/3 
 port link-type access 
 port default vlan 200 
# 
ip route-static 10.20.20.0 255.255.255.0 10.10.20.1 
# 
return     
In Figure 3-142, terminals in an enterprise's office area are connected to the enterprise's intranet through DeviceA. To meet high security requirements of the enterprise, the enterprise wants to deploy 802.1X authentication and use a RADIUS server to authenticate terminals.
Enterprise users often connect their terminals to the network from different locations; for example, they may move their laptops to other offices for working or presentation. By default, a user cannot immediately initiate authentication or access the network after connecting to a new interface. Instead, the user can only initiate authentication on the new interface after the user status detection period expires, or after user online entries are cleared by manually disabling and enabling the interface. To improve user experience, MAC address migration needs to be configured so that users can immediately initiate authentication and access the network after connecting to a new interface.
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 192.168.2.30 1812
[DeviceA-radius-rd1] radius-server accounting 192.168.2.30 1813
[DeviceA-radius-rd1] radius-server shared-key cipher YsHsjx_202206mc@1
[DeviceA-radius-rd1] quit
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme abc
[DeviceA-aaa-authen-abc] authentication-mode radius
[DeviceA-aaa-authen-abc] quit
# Configure the authentication profile p1, bind the 802.1X access profile d1 to the authentication profile, specify the domain example.com as the forcible authentication domain in the authentication profile, set the user access mode to multi-authen, and set the maximum number of access users to 100.
# Bind the authentication profile p1 to 10GE 1/0/1 and enable 802.1X authentication on 10GE 1/0/1.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] authentication-profile p1
[DeviceA-10GE1/0/1] quit
[DeviceA] authentication mac-move enable vlan all
[DeviceA] authentication mac-move detect enable
#
sysname DeviceA
#
vlan batch 10 20 
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example.com force
 authentication mode multi-authen max-user 100
#
radius-server template rd1
 radius-server authentication 192.168.2.30 1812 weight 80
 radius-server accounting 192.168.2.30 1813 weight 80
 radius-server shared-key cipher  %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
#
authentication mac-move enable vlan all
authentication mac-move detect enable
#
interface 10GE 1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
 authentication-profile p1
#
interface 10GE 1/0/2
 port link-type access
 port default vlan 20
#
interface vlanif 10
 ip address 192.168.1.10 255.255.255.0
#
interface vlanif 20
 ip address 192.168.2.10 255.255.255.0
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
return
In Figure 3-143, the WAC of an enterprise is connected to an AP through Switch_A (access switch). The enterprise needs to deploy the WLAN wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
As the WLAN is open to users, there are potential security risks to enterprise information if no access control is available to the WLAN. To meet high security requirements, the enterprise needs to configure the AP as an 802.1X client before 802.1X authentication is performed on STAs. The AP can go online only after being authenticated by Switch_A. Then, the WAC performs 802.1X authentication on STAs and the RADIUS server authenticates identities of STAs.
| Item | Data | 
|---|---|
| RADIUS authentication parameters | RADIUS authentication scheme name: radius_huawei RADIUS accounting scheme name: scheme1 RADIUS server template name: radius_huawei  | 
| 802.1X access profile |  | 
| Authentication domain |  | 
| Authentication profile |  | 
| Item | Data | 
|---|---|
| RADIUS authentication parameters | RADIUS authentication scheme name: scheme1 RADIUS accounting scheme name: scheme2 RADIUS server template name: radius_huawei  | 
| 802.1X access profile |  | 
| Authentication domain |  | 
| Authentication profile |  | 
| DHCP server | The WAC functions as a DHCP server to assign IP addresses to the AP and STAs.  | 
| IP address of the WAC's source interface | VLANIF 100: 10.23.100.1/24 | 
| AP group |  | 
| Regulatory domain profile |  | 
| SSID profile |  | 
| Security profile |  | 
| VAP profile |  | 
<HUAWEI> system-view
[HUAWEI] sysname Switch_A
[Switch_A] vlan batch 100
[Switch_A] interface 10ge 1/0/1
[Switch_A-10GE1/0/1] port link-type trunk
[Switch_A-10GE1/0/1] port trunk pvid vlan 100
[Switch_A-10GE1/0/1] port trunk allow-pass vlan 100
[Switch_A-10GE1/0/1] quit
[Switch_A] interface 10ge 1/0/2
[Switch_A-10GE1/0/2] port link-type trunk
[Switch_A-10GE1/0/2] port trunk allow-pass vlan 100
[Switch_A-10GE1/0/2] quit
[Switch_A] interface vlanif 100
[Switch_A-Vlanif100] ip address 10.23.100.2 24
[Switch_A-Vlanif100] quit
[Switch_A] ip route-static 10.23.200.0 255.255.255.0 10.23.100.1
# Create and configure the RADIUS server template radius_huawei.
[Switch_A] radius-server template radius_huawei
[Switch_A-radius-radius_huawei] radius-server authentication 10.23.200.1 1812 weight 80
[Switch_A-radius-radius_huawei] radius-server accounting 10.23.200.1 1813 weight 80
[Switch_A-radius-radius_huawei] radius-server shared-key cipher YsHsjx_202206
[Switch_A-radius-radius_huawei] quit
# Create the RADIUS authentication scheme radius_huawei and RADIUS accounting scheme scheme1, and set the authentication mode and accounting mode to RADIUS.
[Switch_A] aaa  
[Switch_A-aaa] authentication-scheme radius_huawei
