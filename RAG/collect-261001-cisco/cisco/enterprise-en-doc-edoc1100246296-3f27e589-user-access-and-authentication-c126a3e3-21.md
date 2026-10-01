---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-21
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [3519, 3749]
sha256: d0eb1eca3f259b2f0a36b536cf3246a217c517bc5b74f861c4aa0a0b3f2e5589
---

# Configure DeviceA as a DHCP server to assign IP addresses to APs from the IP address pool on VLANIF 100, and assign IP addresses to STAs from the IP address pool on VLANIF 101.
[DeviceA] interface vlanif 200
[DeviceA-Vlanif200] ip address 10.23.200.3 24
[DeviceA-Vlanif200] quit
[DeviceA] radius-server template radius_test
[DeviceA-radius-radius_test] radius-server authentication 10.23.200.1 1812
[DeviceA-radius-radius_test] radius-server accounting 10.23.200.1 1813
[DeviceA-radius-radius_test] radius-server shared-key cipher YsHsjx_202206mc@1
[DeviceA-radius-radius_test] quit
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme radius_test
[DeviceA-aaa-authen-radius_test] authentication-mode radius
[DeviceA-aaa-authen-radius_test] quit
[DeviceA-aaa] accounting-scheme scheme1
[DeviceA-aaa-accounting-scheme1] accounting-mode radius
[DeviceA-aaa-accounting-scheme1] accounting realtime 15
[DeviceA-aaa-accounting-scheme1] quit
[DeviceA-aaa] quit
# Configure an authentication domain.
[DeviceA-aaa] domain test
[DeviceA-aaa-domain-test] authentication-scheme radius_test
[DeviceA-aaa-domain-test] accounting-scheme scheme1
[DeviceA-aaa-domain-test] radius-server radius_test
[DeviceA-aaa-domain-test] quit
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-source ip-address 10.23.200.3 
[DeviceA-web-auth-server-abc] server-ip 10.23.200.1
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url http://10.23.200.1:8080/portal
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] quit
[DeviceA] free-rule-template name default_free_rule 
[DeviceA-free-rule-default_free_rule] free-rule 1 destination ip 10.23.202.1 mask 32
[DeviceA-free-rule-default_free_rule] quit
[DeviceA] authentication-profile name p1
[DeviceA-authentication-profile-p1] portal-access-profile web1
[DeviceA-authentication-profile-p1] free-rule-template default_free_rule
[DeviceA-authentication-profile-p1] mac-access-profile m1
[DeviceA-authentication-profile-p1] access-domain test
[DeviceA-authentication-profile-p1] quit
#
sysname DeviceA
#
vlan batch 100 to 101 200
#
authentication-profile name p1
 mac-access-profile m1    
 free-rule-template default_free_rule
 portal-access-profile web1
 access-domain test
# 
dhcp enable
#
radius-server template radius_test
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%#
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80
#
free-rule-template name default_free_rule 
 free-rule 1 destination ip 10.23.202.1 mask 255.255.255.255
#
web-auth-server abc
 server-source ip-address 10.23.200.3
 server-ip 10.23.200.1
 port 50200
 url http://10.23.200.1:8080/portal
 shared-key cipher %^%#-;k-</gjVJ=ZK&Ea)<WB(j1FD8HJOGq^@$Ly=\0Y%^%#
#
portal-access-profile name web1
 web-auth-server abc 
#
aaa
 authentication-scheme radius_test
  authentication-mode radius
 accounting-scheme scheme1
  accounting-mode radius
  accounting realtime 15
 domain test
  authentication-scheme radius_test
  accounting-scheme scheme1
  radius-server radius_test
#
interface Vlanif100
 ip address 10.23.100.1 255.255.255.0
 dhcp select interface
#
interface Vlanif101
 ip address 10.23.101.1 255.255.255.0
 dhcp select interface
#
interface Vlanif200
 ip address 10.23.200.3 255.255.255.0
#
interface 10GE 1/0/1
 port link-type trunk
 port trunk allow-pass vlan 100
#
interface 10GE 1/0/2
 port link-type trunk
 port trunk allow-pass vlan 200
#
capwap source interface vlanif 100
capwap dtls psk %+%##!!!!!!!!!"!!!!"!!!!*!!!!a4O4$.&PxRTGtHR(VaB4*(KhWT"C@K+HZ%@!!!!!2jp5!!!!!!;!!!!..AjEF}u<=!%e=F%Dz-Y[R--3yrC-*MVXf<!!!!!%+%#
#
wlan
 security-profile name wlan-security
  security open
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
 ap-id 0 ap-mac 00e0-fc76-e360
  ap-name area_1
  ap-group ap-group1
  radio 0
   channel 20mhz 6
   eirp 127
  radio 1
   channel 80mhz 149
   eirp 127
 ap-id 0 type-id 1 ap-mac 00e0-fc12-3456 ap-sn 210235554710CB000042
  ap-name area_2
  ap-group ap-group1
  radio 0
   channel 20mhz 6
   eirp 127
  radio 1
   channel 80mhz 149
   eirp 127
#
mac-access-profile name m1
#
return
#
sysname DeviceB
#
vlan batch 100 
#
interface 10GE 1/0/1
 port link-type trunk
 port trunk pvid vlan 100
 port trunk allow-pass vlan 100
#
interface 10GE 1/0/2
 port link-type trunk
 port trunk pvid vlan 100
 port trunk allow-pass vlan 100
#
interface 10GE 1/0/3
 port link-type trunk
 port trunk allow-pass vlan 100
#
return
On an enterprise network, DeviceA is connected to APs through DeviceB (access switch). The enterprise plans to deploy a wireless network named wlan-net to provide wireless access for employees. To enhance network security, mobile phone users are allowed access to the network only after successful authentication of user names and passwords.
In this example, interfaces 1, 2, and 3 represent 10GE 1/0/1, 10GE 1/0/2, and 10GE 1/0/3, respectively.
| Item | Data | 
|---|---|
| Portal page to be pushed | Default page for user name and password authentication provided by iMaster NCE-Campus. | 
| Portal authentication mode | User name and password authentication:  | 
| Portal authentication-free validity period | Retain the default setting, which is 2 hours. | 
| Item | Data | 
|---|---|
| SSID profile |  | 
| IP address of the AC's source interface | VLANIF 102: 192.168.102.1/24 | 
| Security profile |  | 
| HACA server template |  | 
| Portal server configuration |  | 
| Portal access profile |  | 
| Authentication profile |  | 
| VAP profile |  | 
<HUAWEI> system-view
[HUAWEI] sysname DeviceB
[DeviceB] vlan batch 102
[DeviceB] interface 10ge 1/0/1
[DeviceB-10GE1/0/1] description to_AP1
[DeviceB-10GE1/0/1] port link-type trunk
[DeviceB-10GE1/0/1] port trunk pvid vlan 102
[DeviceB-10GE1/0/1] undo port trunk allow-pass vlan 1
[DeviceB-10GE1/0/1] port trunk allow-pass vlan 102
[DeviceB-10GE1/0/1] quit
[DeviceB] interface 10ge 1/0/2
[DeviceB-10GE1/0/2] description to_AP2
[DeviceB-10GE1/0/2] port link-type trunk
[DeviceB-10GE1/0/2] port trunk pvid vlan 102
[DeviceB-10GE1/0/2] undo port trunk allow-pass vlan 1
[DeviceB-10GE1/0/2] port trunk allow-pass vlan 102
[DeviceB-10GE1/0/2] quit
[DeviceB] interface 10ge 1/0/3
[DeviceB-10GE1/0/3] description to_WAC
[DeviceB-10GE1/0/3] port link-type trunk
[DeviceB-10GE1/0/3] undo port trunk allow-pass vlan 1
[DeviceB-10GE1/0/3] port trunk allow-pass vlan 102
[DeviceB-10GE1/0/3] quit
<HUAWEI> system-view
[HUAWEI] sysname DeviceA 
[DeviceA] vlan batch 101 102 103
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] description to_DeviceB
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] undo port trunk allow-pass vlan 1
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 102
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] description to_HACA_Server
[DeviceA-10GE1/0/2] port link-type trunk
[DeviceA-10GE1/0/2] undo port trunk allow-pass vlan 1
[DeviceA-10GE1/0/2] port trunk allow-pass vlan 103
[DeviceA-10GE1/0/2] quit
# Configure DHCP on DeviceA. Configure VLANIF 102 to assign IP addresses to APs and VLANIF 101 to assign IP addresses to wireless users.
[DeviceA] dhcp enable
[DeviceA] interface vlanif 101
[DeviceA-Vlanif101] ip address 192.168.101.1 24
[DeviceA-Vlanif101] dhcp select interface
[DeviceA-Vlanif101] quit
[DeviceA] interface vlanif 102
[DeviceA-Vlanif102] ip address 192.168.102.1 24
[DeviceA-Vlanif102] dhcp select interface 
[DeviceA-Vlanif102] quit
[DeviceA] vlan 103
[DeviceA] interface vlanif 103
