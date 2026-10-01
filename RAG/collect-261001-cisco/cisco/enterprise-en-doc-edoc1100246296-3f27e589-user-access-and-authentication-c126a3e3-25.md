---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-25
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [4320, 4551]
sha256: 9ce3bacf4ff05fd916f6fb390069298f738d1f61d7184f86f11a8c00e61b1a23
---

# Configure DeviceA to generate a local key pair.

 radius-server authentication 10.5.1.3 1812
 radius-server accounting 10.5.1.3 1813
 radius-server shared-key cipher %^%#x$`LC*6I3H&~})~8O[$F,,o6FN!+35|H-E3Wi}Z:%^%#
#
interface eth-trunk 10
 port link-type trunk
 port trunk allow-pass vlan 100 200
 authentication-profile p1
#
interface 10GE 1/0/1
 eth-trunk 10
#
interface 10GE 1/0/2
 eth-trunk 10
#
interface 10GE 1/0/3
 port link-type trunk
 port trunk allow-pass vlan 300
#
interface vlanif 100
 ip address 10.1.1.1 255.255.255.0 
 dhcp select interface
#
interface vlanif 200
 ip address 10.2.1.1 255.255.255.0
 dhcp select interface
#
interface vlanif 201
 ip address 10.2.2.1 255.255.255.0
 dhcp select interface
#
interface vlanif 300
 ip address 10.3.1.1 255.255.255.0
#
aaa
 authentication-scheme abc
  authentication-mode radius
 authentication-scheme noauthen
  authentication-mode none
 accounting-scheme acco1
  accounting-mode radius
  accounting realtime 15
 domain isp
  authentication-scheme abc
  accounting-scheme acco1
  radius-server rd1
 domain ap_noauthen
  authentication-scheme noauthen
#
wlan
 security-profile name wlan-net-security
  security open
 ssid-profile name wlan-net-ssid
  ssid wlan-net
 vap-profile name wlan-net-vap
  forward-mode tunnel
  service-vlan vlan-id 201
  ssid-profile wlan-net-ssid
  security-profile wlan-net-security
  authentication-profile p2
 regulatory-domain-profile name domain1
  country-code cn
 ap-group name ap-group1
  regulatory-domain-profile domain1
  radio 0
   vap-profile wlan-net-vap wlan 1
  radio 1
   vap-profile wlan-net-vap wlan 1
  radio 2
   vap-profile wlan-net-vap wlan 1
 ap-id 0 type-id 1 ap-mac 00e0-fc12-3456 ap-sn 210235554710CB000042
  ap-name area_1
  ap-group ap-group1
#
domain ap_noauthen mac-authen force mac-address 00e0-fc76-e360 mask ffff-ffff-ff00  
#
access-context profile name ap_access
 if-match vlan-id 100
access-context profile enable
#
access-author policy name ap_noauthen
 match access-context-profile ap_access action access-domain ap_noauthen force
access-author policy ap_noauthen global
#
mac-access-profile name m1
#
dot1x-access-profile name d1
#
capwap source interface vlanif 100
#
ip route-static 10.5.1.0 255.255.255.0 10.3.1.2
#
return
#
sysname DeviceB
#
vlan batch 100 200
#
l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
#
interface eth-trunk 10
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type trunk
 port trunk allow-pass vlan 100 200
#
interface 10GE 1/0/1
 eth-trunk 10
#
interface 10GE 1/0/2
 eth-trunk 10
#
interface 10GE 1/0/3
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type trunk
 port trunk allow-pass vlan 100 200
#
return
#
sysname DeviceC
#
vlan batch 100 200
#
l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
#
interface 10GE 1/0/1
 port link-type trunk
 undo port trunk allow-pass vlan 1
 port trunk pvid vlan 100
 port trunk allow-pass vlan 100
#
interface 10GE 1/0/2
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type access
 port default vlan 200
#
interface 10GE 1/0/3
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type trunk
 port trunk allow-pass vlan 100 200
#
return
In Figure 3-138, to save investment costs, users require that IP phones and PCs connect to the network through VoIP. IP phones support LLDP and can obtain the voice VLAN through LLDP. The network needs to meet the following requirements:
The priority of voice packets sent by IP phones is low and needs to be increased to ensure communication quality.
Voice and data packets are transmitted in VLAN 200 and VLAN 100, respectively.
IP phones and PCs obtain IP addresses from a DHCP server.
IP phones connect to DeviceA without authentication and PCs connect to DeviceA through 802.1X authentication.
[HUAWEI] authentication mac-move enable vlan all
[HUAWEI] authentication mac-move detect enable
# Create VLANs and enable the LLDP function.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 100 200 300
[DeviceA] lldp enable
# Add interfaces to the data VLAN.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] portswitch
[DeviceA-10GE1/0/1] port link-type hybrid
[DeviceA-10GE1/0/1] port hybrid pvid vlan 100
[DeviceA-10GE1/0/1] port hybrid untagged vlan 100
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] portswitch
[DeviceA-10GE1/0/2] port link-type hybrid
[DeviceA-10GE1/0/2] port hybrid pvid vlan 100
[DeviceA-10GE1/0/2] port hybrid untagged vlan 100
[DeviceA-10GE1/0/2] quit
[DeviceA] interface 10ge 1/0/3
[DeviceA-10GE1/0/3] portswitch
[DeviceA-10GE1/0/3] port link-type access
[DeviceA-10GE1/0/3] port default vlan 300
[DeviceA-10GE1/0/3] quit
[DeviceA] interface vlanif 100
[DeviceA-Vlanif100] ip address 10.1.1.1 24  
[DeviceA-Vlanif100] quit
[DeviceA] interface vlanif 300
[DeviceA-Vlanif300] ip address 10.3.1.1 24
[DeviceA-Vlanif300] quit
# Add interfaces to the voice VLAN.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] port hybrid tagged vlan 200
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] port hybrid tagged vlan 200
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 200
[DeviceA-Vlanif200] ip address 10.2.1.1 24  
[DeviceA-Vlanif200] quit
# Enable the voice VLAN on interfaces and configure the interfaces to authorize the voice VLAN in TLV mode.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] voice-vlan 200 enable   
[DeviceA-10GE1/0/1] stp edged-port enable   
[DeviceA-10GE1/0/1] lldp tlv-enable med-tlv network-policy voice-vlan vlan 200 cos 6 dscp 60   
[DeviceA-10GE1/0/1] lldp mdn enable txrx //For IP phones that use non-standard discovery protocols such as CDP, perform this step to ensure that the device can establish neighbor relationships with these IP phones.
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] voice-vlan 200 enable
[DeviceA-10GE1/0/2] stp edged-port enable 
[DeviceA-10GE1/0/2] lldp tlv-enable med-tlv network-policy voice-vlan vlan 200 cos 6 dscp 60  
[DeviceA-10GE1/0/2] lldp mdn enable txrx //For IP phones that use non-standard discovery protocols such as CDP, perform this step to ensure that the device can establish neighbor relationships with these IP phones.
[DeviceA-10GE1/0/2] quit
# Configure a static route from DeviceA to the server zone. In this example, the next-hop IP address is 10.3.1.2.
[DeviceA] ip route-static 0.0.0.0 0.0.0.0 10.3.1.2 
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 10.5.1.3 1812
[DeviceA-radius-rd1] radius-server accounting 10.5.1.3 1813
[DeviceA-radius-rd1] radius-server shared-key cipher Huawei@123456789
[DeviceA-radius-rd1] quit
# Configure a forcible authentication domain. If user names carry domain names, you do not need to configure a forcible authentication domain for 802.1X users.
[DeviceA-authen-profile-p1] access-domain isp dot1x force  
# Enable IP phones to go online without authentication.
[DeviceA-authen-profile-p1] authentication device-type voice authorize  
[DeviceA-authen-profile-p1] quit
# Bind the authentication profile to interfaces.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] authentication-profile p1
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] authentication-profile p1
[DeviceA-10GE1/0/2] quit
After an IP phone starts, it automatically obtains an IP address and connects to the network without authentication.
A user starts a PC and enters the user name and password, triggering 802.1X authentication. After the authentication is successful, the user can access the network.
#
sysname DeviceA
#
vlan batch 100 200 300
#
authentication-profile name p1
 dot1x-access-profile d1 
 access-domain isp dot1x force
 authentication device-type voice authorize
#
radius-server template rd1
