---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-12
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [1744, 1955]
sha256: d493bc076c7e738ce58e481a8137a9124b97d622e72ddb1f5f74ab44c63bc38b
---

# Configure DeviceA to generate a local key pair.

  accounting-mode radius
  accounting realtime 15
 domain example.com
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
In Figure 3-118, DeviceA functioning as a WAC in an enterprise is directly connected to an AP. The enterprise deploys the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
As the WLAN is open to users, there are potential security risks to enterprise information if no access control is configured for the WLAN. Dumb terminals (such as printers) in the physical access control department cannot have an authentication client installed. To meet the enterprise's security requirements, MAC address authentication can be configured on the WAC to authenticate dumb terminals in local authentication mode.
| Item | Data | 
|---|---|
| Local authentication parameters | Local authentication scheme name: a1 User name, password, and access type of the local user (STA1 is used as an example):  | 
| MAC access profile |  | 
| Authentication domain | The authentication domain is named example.com and is bound to:  | 
| Authentication profile |  | 
| DHCP server | The WAC functions as a DHCP server to assign IP addresses to the AP and STAs. | 
| IP address pool for the AP | 10.23.100.2/24 to 10.23.100.254/24 | 
| IP address pool for STAs | 10.23.101.2/24 to 10.23.101.254/24 | 
| IP address of the WAC's source interface | VLANIF 100: 10.23.100.1/24 | 
| AP group |  | 
| Regulatory domain profile |  | 
| SSID profile |  | 
| Security profile |  | 
| VAP profile |  | 
Configure a terminal's MAC address as the local user name, set the password to Example@123, and set the access type to MAC address authentication (dot1x). In this example, the MAC address of printer 1 is 00e0-fcd4-8828.
#
sysname DeviceA
#
vlan batch 100 to 101
#
authentication-profile name p1
 mac-access-profile m1
 access-domain example.com force
#
mac-access-profile name m1
 mac-authen username macaddress format with-hyphen password cipher %^%#PW~_5m;sAFFI.cEB"%^@6@4$96ds_5+O'28+d3:A%^%# 
#
dhcp enable
#
aaa
 authentication-scheme a1    
  authentication-mode local
 authorization-scheme b1
  authorization-mode local
 domain example.com            
  authentication-scheme a1
  authorization-scheme b1
 local-access-user 00e0-fcd4-8828
  password cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.<FvBB,.w;M75IN5Z>'!L8G:n-!!!!!2jp5!!!!!!<!!!!k9&fPO<BSRW}jPT(,ewKyfIL"zVtM1~=>e.!!!!!%+%# 
  service-type dot1x 
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
In Figure 3-119, DeviceA functioning as a WAC in an enterprise is directly connected to an AP. The enterprise deploys the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
As the WLAN is open to users, there are potential security risks to enterprise information if no access control is available to the WLAN. To meet the enterprise's high security requirements, 802.1X authentication is used for authenticating wireless users through a RADIUS server.
| Item | Data | 
|---|---|
| RADIUS authentication parameters | Authentication scheme name: scheme1 Accounting scheme name: scheme2 RADIUS server template name: radius_huawei  | 
| Authentication domain | The authentication domain is named example.com and is bound to:  | 
| 802.1X access profile |  | 
| Authentication profile |  | 
| DHCP server | The WAC functions as a DHCP server to assign IP addresses to the AP and STAs. | 
| IP address pool for the AP | 10.23.100.2 to 10.23.100.254/24 | 
| IP address pool for STAs | 10.23.101.2 to 10.23.101.254/24 | 
| IP address of the WAC's source interface | VLANIF 100: 10.23.100.1/24 | 
| AP group |  | 
| Regulatory domain profile |  | 
| SSID profile |  | 
| Security profile |  | 
| VAP profile |  | 
By default, an 802.1X access profile uses EAP authentication. Ensure that the RADIUS server supports the EAP protocol. Otherwise, the RADIUS server cannot process 802.1X authentication requests.
[DeviceA] authentication-profile name p1
[DeviceA-authentication-profile-p1] dot1x-access-profile d1
[DeviceA-authentication-profile-p1] access-domain example.com force
[DeviceA-authentication-profile-p1] quit
[DeviceA] wlan
[DeviceA-wlan] security-profile name wlan-security
[DeviceA-wlan-sec-prof-wlan-security] security wpa2 dot1x aes
[DeviceA-wlan-sec-prof-wlan-security] quit
Configuration in the Windows XP operating system:
Configuration in the Windows 7 operating system:
#
sysname DeviceA
#
vlan batch 100 to 101
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example.com force
#
dot1x-access-profile name d1
#
dhcp enable
#
radius-server template radius_huawei
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!Cd/`W03KjAwAqn64E<\TxGC_SOri<2BP\A+!!!!!2jp5!!!!!!B!!!!Oe2HMc->XMa#TLDZUaJBFJtm#XVj*E:S*|(N7`J1B:3QY!!!!!!!!!!!!!!!%+%# 
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80
#
aaa
 authentication-scheme scheme1
  authentication-mode radius
 accounting-scheme scheme2
  accounting-mode radius
  accounting realtime 15
 domain example.com
