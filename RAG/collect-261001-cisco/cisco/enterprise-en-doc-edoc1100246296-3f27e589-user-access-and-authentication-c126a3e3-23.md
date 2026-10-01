---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-23
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [3956, 4148]
sha256: 41a9b2f860930ddfafa9c5cd3ebb259db240a751302116a7ed6138d68c1bc50d
---

# Configure DeviceA to generate a local key pair.

   vap-profile wlan-net-vap wlan 1
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
ip route-static 192.168.100.0 255.255.255.0 192.168.103.2
#
On the network shown in Figure 3-136, the AC and AP are third-party WLAN access devices. Mobile STAs can access the Internet only after associating with a Wi-Fi network and passing Portal authentication. Specific requirements are as follows:
# Add 10GE 1/0/1 that connects DeviceA to DeviceB to VLAN 100 and VLAN 101. In this example, the AC uses the direct forwarding mode. If the AC uses the tunnel forwarding mode, you do not need to add this interface to VLAN 101.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 20 100 101
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 100 101
[DeviceA-10GE1/0/1] quit
# Add 10GE 1/0/2 that connects DeviceA to the AC to VLAN 100 and VLAN 101.
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] port link-type trunk
[DeviceA-10GE1/0/2] port trunk allow-pass vlan 100 101
[DeviceA-10GE1/0/2] quit
# Add 10GE 1/0/3 that connects DeviceA to the server to VLAN 20.
[DeviceA] interface 10ge 1/0/3
[DeviceA-10GE1/0/3] port link-type trunk
[DeviceA-10GE1/0/3] port trunk allow-pass vlan 20
[DeviceA-10GE1/0/3] quit
# Add interfaces on DeviceB to VLAN 100 and VLAN 101. In this example, the AC uses the direct forwarding mode. If the AC uses the tunnel forwarding mode, you do not need to add these interfaces to VLAN 101.
<HUAWEI> system-view
[HUAWEI] sysname DeviceB
[DeviceB] vlan batch 100 101
[DeviceB] interface 10ge 1/0/1
[DeviceB-10GE1/0/1] port link-type trunk
[DeviceB-10GE1/0/1] port trunk allow-pass vlan 100 101
[DeviceB-10GE1/0/1] quit
[DeviceB] interface 10ge 1/0/2
[DeviceB-10GE1/0/2] port link-type trunk
[DeviceB-10GE1/0/2] port trunk allow-pass vlan 100 101
[DeviceB-10GE1/0/2] port trunk pvid vlan 100
[DeviceB-10GE1/0/2] quit
# On DeviceA, configure an IP address for a VLANIF interface and configure a route destined for the server. In this example, the next-hop IP address is 10.4.1.2.
[DeviceA] interface vlanif 20
[DeviceA-Vlanif20] ip address 10.4.1.1 24
[DeviceA-Vlanif20] quit
[DeviceA] interface vlanif 101
[DeviceA-Vlanif101] ip address 10.3.1.2 24
[DeviceA-Vlanif101] quit
[DeviceA] ip route-static 10.5.1.0 255.255.255.0 10.4.1.2
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 10.5.1.3 1812
[DeviceA-radius-rd1] radius-server accounting 10.5.1.3 1813
[DeviceA-radius-rd1] radius-server shared-key cipher YsHsjx_202206139
[DeviceA-radius-rd1] quit
[DeviceA-aaa] accounting-scheme acco1
[DeviceA-aaa-accounting-acco1] accounting-mode radius
[DeviceA-aaa-accounting-acco1] accounting realtime 15  
[DeviceA-aaa-accounting-acco1] quit
[DeviceA-aaa] domain isp
[DeviceA-aaa-domain-isp] authentication-scheme abc
[DeviceA-aaa-domain-isp] accounting-scheme acco1
[DeviceA-aaa-domain-isp] radius-server rd1
[DeviceA-aaa-domain-isp] quit
[DeviceA-aaa] quit
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme noauthen
[DeviceA-aaa-authen-noauthen] authentication-mode none
[DeviceA-aaa-authen-noauthen] quit
[DeviceA-aaa] domain ap_noauthen
[DeviceA-aaa-domain-ap_noauthen] authentication-scheme noauthen
[DeviceA-aaa-domain-ap_noauthen] quit
[DeviceA-aaa] quit
# Configure non-authentication for the AP using either of the following methods:
[DeviceA] domain ap_noauthen mac-authen force mac-address 00e0-fc74-9640 mask ffff-ffff-ff00
[DeviceA] access-context profile enable 
[DeviceA] access-context profile name ap_access   
[DeviceA-access-context-ap_access] if-match vlan-id 100   
[DeviceA-access-context-ap_access] quit
[DeviceA] access-author policy name ap_noauthen   
[DeviceA-access-author-ap_noauthen] match access-context-profile ap_access action access-domain ap_noauthen force
[DeviceA-access-author-ap_noauthen] quit
[DeviceA] access-author policy ap_noauthen global  
The priorities of the forcible domain, domain carried in the user name, and default domain in different views are as follows in descending order: forcible domain with a specified authentication mode in an authentication profile > forcible domain in an authentication profile > forcible domain with a specified authentication mode based on a user context profile > forcible domain based on a user context profile > domain carried in the user name > default domain with a specified authentication mode in an authentication profile > default domain in an authentication profile > default domain with a specified authentication mode based on a user context profile > default domain based on a user context profile > global default domain.
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-source ip-address 10.4.1.1  
[DeviceA-web-auth-server-abc] server-ip 10.5.1.3
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url http://10.5.1.3:8445/portal
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] quit
# Configure the Portal access profile web1.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] mac-access-profile m1    
[DeviceA-authen-profile-p1] portal-access-profile web1 
[DeviceA-authen-profile-p1] access-domain isp  
[DeviceA-authen-profile-p1] quit
#
sysname DeviceB
#
vlan batch 100 101
#
interface 10GE 1/0/1
 port link-type trunk
 port trunk allow-pass vlan 100 101
#
interface 10GE 1/0/2
 port link-type trunk
 port trunk allow-pass vlan 100 101
 port trunk pvid vlan 100
#
return
#
sysname DeviceA
#
vlan batch 20 100 101
#
authentication-profile name p1
 mac-access-profile m1    
 portal-access-profile web1 
 access-domain isp
#
radius-server template rd1
 radius-server authentication 10.5.1.3 1812
 radius-server accounting 10.5.1.3 1813
 radius-server shared-key cipher %^%#6A6YSa1*aLmY+*4|W|<Xqr9XDeO6l$IHV{49K(s3%^%#
#
web-auth-server abc
 server-source ip-address 10.4.1.1
 server-ip 10.5.1.3
 port 50200
 url http://10.5.1.3:8445/portal
 shared-key cipher %^%#-;k-</gjVJ=ZK&Ea)<WB(j1FD8HJOGq^@$Ly=\0Y%^%#
#
portal-access-profile name web1
 web-auth-server abc
#
interface 10GE 1/0/1
 port link-type trunk
 port trunk allow-pass vlan 100 101
 authentication-profile p1
#
interface 10GE 1/0/2
 port link-type trunk
 port trunk allow-pass vlan 100 101
#
interface 10GE 1/0/3
 port link-type trunk
 port trunk allow-pass vlan 20
#
interface vlanif 20
 ip address 10.4.1.1 255.255.255.0
#
interface vlanif 101
 ip address 10.3.1.2 255.255.255.0
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
domain ap_noauthen mac-authen force mac-address 00e0-fc74-9640 mask ffff-ffff-ff00
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
ip route-static 10.5.1.0 255.255.255.0 10.4.1.2
#
return
