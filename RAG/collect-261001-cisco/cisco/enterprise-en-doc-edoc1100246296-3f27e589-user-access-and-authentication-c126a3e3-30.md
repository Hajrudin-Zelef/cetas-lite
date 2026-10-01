---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-30
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [5237, 5432]
sha256: 443899717adf37ed4e50d43f8143e7c3d7bced1ac492a6e9819f2c81ed3ffe1b
---

# Configure DeviceA to generate a local key pair.

[WAC-radius-radius_huawei] radius-server accounting 10.23.200.1 1813
[WAC-radius-radius_huawei] radius-server shared-key cipher YsHsjx_202206mc@1
[WAC-radius-radius_huawei] quit
[WAC] aaa
[WAC-aaa] authentication-scheme scheme1
[WAC-aaa-authen-scheme1] authentication-mode radius
[WAC-aaa-authen-scheme1] quit
[WAC-aaa] accounting-scheme scheme2
[WAC-aaa-accounting-scheme2] accounting-mode radius
[WAC-aaa-accounting-scheme2] accounting realtime 15
[WAC-aaa-accounting-scheme2] quit
In this example, the device is connected to iMaster NCE-Campus. The accounting function is not provided for accounting purposes, and is only used to maintain terminal online information through accounting packets.
[WAC-aaa] domain example.com
[WAC-aaa-domain-example.com] authentication-scheme scheme1
[WAC-aaa-domain-example.com] accounting-scheme scheme2
[WAC-aaa-domain-example.com] radius-server radius_huawei
[WAC-aaa-domain-example.com] quit
[WAC-aaa] quit
[WAC] dot1x-access-profile name d1
[WAC-dot1x-access-profile-d1] quit
[WAC] authentication-profile name p1
[WAC-authentication-profile-p1] dot1x-access-profile d1
[WAC-authentication-profile-p1] access-domain example.com force
[WAC-authentication-profile-p1] quit
[WAC] wlan
[WAC-wlan] security-profile name wlan-security
[WAC-wlan-sec-prof-wlan-security] security wpa2 dot1x aes
[WAC-wlan-sec-prof-wlan-security] quit
[WAC-wlan] ssid-profile name wlan-ssid
[WAC-wlan-ssid-prof-wlan-ssid] ssid wlan-net
[WAC-wlan-ssid-prof-wlan-ssid] quit
[WAC-wlan] vap-profile name wlan-vap
[WAC-wlan-vap-prof-wlan-vap] forward-mode tunnel
[WAC-wlan-vap-prof-wlan-vap] service-vlan vlan-id 101
[WAC-wlan-vap-prof-wlan-vap] security-profile wlan-security
[WAC-wlan-vap-prof-wlan-vap] ssid-profile wlan-ssid
[WAC-wlan-vap-prof-wlan-vap] authentication-profile p1
[WAC-wlan-vap-prof-wlan-vap] quit
# Bind the VAP profile wlan-vap to the AP group and apply the profile to radios 0 and 1 of the AP.
[WAC-wlan] ap-group name ap-group1
[WAC-wlan-ap-group-ap-group1] vap-profile wlan-vap wlan 1 radio 0
[WAC-wlan-ap-group-ap-group1] vap-profile wlan-vap wlan 1 radio 1
[WAC-wlan-ap-group-ap-group1] quit
The automatic channel and power calibration functions of radios are enabled by default. The manual channel and power settings take effect only when the two functions are disabled. The channel and power settings for the AP radios in this example are for reference only. In practice, configure the channel and power of AP radios based on the actual country code of the AP and network planning.
[WAC-wlan] ap-id 0
[WAC-wlan-ap-0] radio 0
[WAC-wlan-ap-0-radio-0] calibrate auto-channel-select disable
[WAC-wlan-ap-0-radio-0] calibrate auto-txpower-select disable
[WAC-wlan-ap-0-radio-0] channel 20mhz 6
Warning: This action may cause service interruption. Continue?[Y/N]y 
[WAC-wlan-ap-0-radio-0] eirp 127
[WAC-wlan-ap-0-radio-0] quit
[WAC-wlan-ap-0] radio 1
[WAC-wlan-ap-0-radio-1] calibrate auto-channel-select disable
[WAC-wlan-ap-0-radio-1] calibrate auto-txpower-select disable
[WAC-wlan-ap-0-radio-1] channel 80mhz 149
Warning: This action may cause service interruption. Continue?[Y/N]y 
[WAC-wlan-ap-0-radio-1] eirp 127
[WAC-wlan-ap-0-radio-1] quit
[WAC-wlan-ap-0] quit
[WAC-wlan] quit
#
 sysname Switch_A    
#                                                                                                        
vlan batch 100                                                             
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain huawei force  
#
dot1x-access-profile name d1
#
radius-server template radius_huawei
 radius-server shared-key cipher %^%#ANM|Cb!>GNo=U@V~_{E1fQ>;I2#2l(3Q%1~Z.u|R%^%# 
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80                    
#
aaa        
 authentication-scheme radius_huawei
  authentication-mode radius
 accounting-scheme scheme1                                                                   
  accounting-mode radius       
  accounting realtime 15                                                                             
 domain huawei                                                                            
  authentication-scheme radius_huawei                                                                          
  accounting-scheme scheme1
  radius-server radius_huawei
#
interface 10GE1/0/1
 port link-type trunk 
 port trunk pvid vlan 100
 port trunk allow-pass vlan 100
 authentication-profile p1
#
interface 10GE1/0/2
 port link-type trunk
 port trunk allow-pass vlan 100  
#
interface vlanif 100
 ip address 10.23.100.2 255.255.255.0
#
ip route-static 10.23.200.0 255.255.255.0 10.23.100.1
#
return
#
sysname DeviceA
#
vlan batch 100 to 101 200
#
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example.com force
#
dot1x-access-profile name d1
#
dhcp enable
#
dot1x-client-profile name huawei
 eap-method eap-peap username huawei password cipher %^%#f,x[/WLW|B;vh/Nbaey$V4s17cL/R06x|d$G%!q'%^%#  
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
  authentication-scheme scheme1
  accounting-scheme scheme2
  radius-server radius_huawei
#
interface Vlanif100
 ip address 10.23.100.1 255.255.255.0
 dhcp select interface
 dhcp server excluded-ip-address 10.23.100.2
#
interface Vlanif101
 ip address 10.23.101.1 255.255.255.0
 dhcp select interface
#
interface Vlanif200 
 ip address 10.23.200.2 255.255.255.0 
#
interface 10GE1/0/1
 port link-type trunk
 port trunk pvid vlan 100
 port trunk allow-pass vlan 100
#
interface 10GE1/0/2
 port link-type trunk
 port trunk allow-pass vlan 200
#
ip route-static 10.23.200.0 255.255.255.0 10.23.101.2 
#  
capwap source interface vlanif 100
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
  wired-port-profile wired-port1 MultiGE 0
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
In Figure 3-144, DeviceA functioning as a WAC in an enterprise is directly connected to an AP. The enterprise needs to deploy the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs. To meet the enterprise's security requirements, Portal authentication needs to be deployed on the WAC. The WAC then works with the RADIUS server (integrated with the Portal server) to implement access control on STAs that attempt to access the enterprise network.
