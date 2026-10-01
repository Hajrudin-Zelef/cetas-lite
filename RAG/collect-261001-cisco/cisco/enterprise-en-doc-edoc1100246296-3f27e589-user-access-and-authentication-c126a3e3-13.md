---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-13
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [1956, 2143]
sha256: 7fb3778754e9bd6fd3678ed7c93c98955435da39e115e7eeb2f7507728f95246
---

# Configure DeviceA to generate a local key pair.

  authentication-scheme scheme1
  accounting-scheme scheme2
  radius-server radius_huawei
#
interface Vlanif100
 ip address 10.23.100.1 255.255.255.0
 dhcp select interface
#
interface Vlanif101
 ip address 10.23.101.1 255.255.255.0
 dhcp select interface
#
interface 10GE1/0/1
 port link-type trunk
 port trunk pvid vlan 100
 port trunk allow-pass vlan 100
#
interface 10GE1/0/2
 port link-type trunk
 port trunk allow-pass vlan 101
#
ip route-static 10.23.200.0 255.255.255.0 10.23.101.2 
#  
capwap source interface vlanif 100
capwap dtls psk %+%##!!!!!!!!!"!!!!"!!!!*!!!!a4O4$.&PxRTGtHR(VaB4*(KhWT"C@K+HZ%@!!!!!2jp5!!!!!!;!!!!..AjEF}u<=!%e=F%Dz-Y[R--3yrC-*MVXf<!!!!!%+%#
#
wlan
 security-profile name wlan-security
  security wpa2 dot1x aes
 ssid-profile name wlan-ssid
  ssid wlan-net
 vap-profile name wlan-vap
  forward-mode tunnel
  service-vlan vlan-id 101
  ssid-profile wlan-ssid
  security-profile wlan-security
  authentication-profile p1
 regulatory-domain-profile name domain1
 ap-group name ap-group1
  regulatory-domain-profile domain1
  radio 0
   vap-profile wlan-vap wlan 1
  radio 1
   vap-profile wlan-vap wlan 1
 ap-id 0 type-id 1 ap-mac 00e0-fc12-3456 ap-sn 210235554710CB000042
  ap-name area_1
  ap-group ap-group1
  radio 0
   channel 20mhz 6
   eirp 127
   calibrate auto-channel-select disable  
   calibrate auto-txpower-select disable
  radio 1
   channel 80mhz 149
   eirp 127
   calibrate auto-channel-select disable 
   calibrate auto-txpower-select disable
#
return
In Figure 3-120, terminals in a company's office area are connected to the company's intranet through DeviceB. DeviceB functioning as an access switch is connected to 10GE 1/0/2 on DeviceA functioning as an aggregation switch through 10GE 1/0/1.
To meet the enterprise's high security requirements, 802.1X authentication needs to be configured to authenticate terminals in the office area through the RADIUS server. Additionally, to facilitate maintenance and reduce the number of authentication points, an authentication point needs to be deployed on 10GE 1/0/2 of DeviceA.
In the 802.1X authentication scenario, if there is a Layer 2 switch between the 802.1X-enabled device and users, Layer 2 transparent transmission must be enabled for 802.1X authentication packets on the Layer 2 switch; otherwise, users cannot be successfully authenticated. This function needs to be configured on all interfaces through which 802.1X authentication packets pass.
<HUAWEI> system-view
[HUAWEI] sysname DeviceB
[DeviceB] l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
[DeviceB] interface 10ge 1/0/1
[DeviceB-10GE1/0/1] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceB-10GE1/0/1] port link-type trunk
[DeviceB-10GE1/0/1] port trunk allow-pass vlan 20
[DeviceB-10GE1/0/1] quit
[DeviceB] interface 10ge 1/0/2
[DeviceB-10GE1/0/2] l2protocol-tunnel user-defined-protocol 802.1X enable
[DeviceB-10GE1/0/2] port link-type access
[DeviceB-10GE1/0/2] port default vlan 20
[DeviceB-10GE1/0/2] quit
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 10 20
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] port link-type trunk
[DeviceA-10GE1/0/2] port trunk allow-pass vlan 20
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 192.168.1.10 24
[DeviceA-Vlanif10] quit
[DeviceA] interface vlanif 20  
[DeviceA-Vlanif20] ip address 192.168.2.10 24
[DeviceA-Vlanif20] quit
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 192.168.1.30 1812
[DeviceA-radius-rd1] radius-server accounting 192.168.1.30 1813
[DeviceA-radius-rd1] radius-server shared-key cipher YsHsjx_202206mc@1
[DeviceA-radius-rd1] quit
Check whether a user can pass RADIUS authentication. (The test user test and password YsHsjx_2022061 have been configured on the RADIUS server.)
[DeviceA] test-aaa test YsHsjx_2022061 radius-template rd1
Info: Account test succeeded.
[DeviceA] dot1x-access-profile name d1
[DeviceA-dot1x-access-profile-d1] dot1x authentication-method eap
[DeviceA-dot1x-access-profile-d1] dot1x timer client-timeout 30
[DeviceA-dot1x-access-profile-d1] quit
Configure the authentication profile p1, bind the 802.1X access profile d1 to the authentication profile, specify the domain example.com as the forcible authentication domain in the authentication profile, set the user access mode to multi-authen, and set the maximum number of access users to 100.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] dot1x-access-profile d1
[DeviceA-authen-profile-p1] access-domain example.com force
[DeviceA-authen-profile-p1] authentication mode multi-authen max-user 100
[DeviceA-authen-profile-p1] quit
# Bind the authentication profile p1 to 10GE 1/0/2 and enable 802.1X authentication on 10GE 1/0/2.
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] authentication-profile p1
[DeviceA-10GE1/0/2] quit
#
sysname DeviceB
#
vlan batch 20 
#
l2protocol-tunnel user-defined-protocol 802.1X protocol-mac 0180-c200-0003 group-mac 0100-0000-0002
#
interface 10GE 1/0/1
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type trunk
 port trunk allow-pass vlan 20
#
interface 10GE 1/0/2
 l2protocol-tunnel user-defined-protocol 802.1X enable
 port link-type access
 port default vlan 20
#
return
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
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
 radius-server authentication 192.168.1.30 1812 weight 80
 radius-server accounting 192.168.1.30 1813 weight 80
#
interface 10GE1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface 10GE1/0/2
 port link-type trunk
 port trunk allow-pass vlan 20
 authentication-profile p1
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
In Figure 3-121, terminals in a company's office area are connected to the company's intranet through DeviceA. The downlink interfaces (for example, 10GE 1/0/2) of DeviceA are directly connected to terminals in the office area, and the uplink interface 10GE 1/0/1 of DeviceA is connected to the RADIUS server through the intranet.
To meet the company's high security requirements, 802.1X authentication needs to be configured to authenticate terminals in the office area through the RADIUS server. Additionally, authentication points need to be deployed on DeviceA's interfaces (for example, 10GE 1/0/2) that are directly connected to the terminals.
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 192.168.1.30 1812
[DeviceA-radius-rd1] radius-server accounting 192.168.1.30 1813
[DeviceA-radius-rd1] radius-server shared-key cipher Huawei@123456789
[DeviceA-radius-rd1] quit
# Create the AAA authentication scheme abc and set the authentication mode to RADIUS authentication.
# Check whether a user can pass RADIUS authentication. (The test user test and password Example2012 have been configured on the RADIUS server.)
[DeviceA] test-aaa test Example2012 radius-template rd1
Info: Account test succeeded.
