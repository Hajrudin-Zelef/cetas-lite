---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-17
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [2710, 2924]
sha256: e41d1c510aa240c1d13609603a500ee7385c628356f274aef228fcdc41c7553c
---

# Configure DeviceA to generate a local key pair.

The login-url parameter is used by terminals to send account information to the device. Some Portal servers do not support the login-url parameter setting. Therefore, you need to configure the URL to carry the login-url parameter so that the URL can carry the login-url parameter to the Portal server when a terminal accesses the Portal page. The IP address in the login-url parameter is the WAC's local IP address, which needs to be permitted in the authentication-free rule profile to ensure that there is a reachable route between the terminal and the IP address. The port number in the parameter is specified using the portal web-authen-server http command, and the default port number is 8000.
[DeviceA] portal web-authen-server http  
[DeviceA] portal web-authen-server server-source all-interface
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] protocol http
[DeviceA-web-auth-server-abc] quit
#
 sysname DeviceA
#    
vlan batch 100 to 101
#
authentication-profile name p1
 portal-access-profile portal1
 free-rule-template default_free_rule
 access-domain example.com force
#
portal web-authen-server http   
portal web-authen-server server-source all-interface                
#   
dhcp enable
#
radius-server template radius_huawei
 radius-server shared-key cipher %^%#Oc6_BMCw#9gZ2@SMVtk!PAC6>Ou*eLW/"qLp+f#$%^%#
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80                                                              
#
free-rule-template name default_free_rule                                                                                           
 free-rule 1 destination ip 10.23.200.2 mask 255.255.255.0
 free-rule 2 destination ip 10.23.100.1 mask 255.255.255.0                                                               
#
url-template name url1
 url http://10.23.200.1:8080/portal
 url-parameter login-url ac_url http://10.23.100.1:8000
# 
web-auth-server abc
 server-ip 10.23.200.1
 url-template url1
 protocol http
#
portal-access-profile name portal1
 web-auth-server abc
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
#
return
In Figure 3-127, DeviceA of an enterprise functions as a WAC and is directly connected to an AP. The enterprise deploys the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
To reduce network security risks, you can deploy Portal authentication on the WAC. The WAC works with the RADIUS server (integrated with the Portal server) to implement access control on STAs who attempt to access the enterprise network, meeting the enterprise's security requirements.
Portal authentication allows you to specify parameters carried in a redirect URL so that the Portal server can obtain user terminal information based on the parameters in the URL and provide different web authentication pages for different users based on the user terminal information.
[DeviceA] url-template name url1 
[DeviceA-url-template-url1] url http://10.23.200.1:19008/portal 
[DeviceA-url-template-url1] url-parameter set device-ip 10.23.101.1
[DeviceA-url-template-url1] url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac
[DeviceA-url-template-url1] quit
#
sysname DeviceA
#
vlan batch 100 to 101
#
authentication-profile name p1
 portal-access-profile portal1
 free-rule-template default_free_rule
 access-domain example.com force
#
dhcp enable
#
radius-server template radius_huawei
 radius-server shared-key cipher %^%#Oc6_BMCw#9gZ2@SMVtk!PAC6>Ou*eLW/"qLp+f#$%^%#
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80
#
free-rule-template name default_free_rule
 free-rule 1 destination ip 10.23.200.2 mask 255.255.255.0
# 
url-template name url1
 url https://10.23.200.1:19008/portal
 url-parameter set device-ip 10.23.101.1
 url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac
#
web-auth-server server-source all-interface
web-auth-server abc
 server-ip 10.23.200.1
 port 50200
 shared-key cipher %^%#4~ZXE3]6@BXu;2;aw}hA{rSb,@"L@T#e{%6G1AiD%^%#
 url-template url1
#
portal-access-profile name portal1
 web-auth-server abc
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
On the network shown in Figure 3-128, users in a company's guest area access the company's intranet through DeviceA. Unauthorized access to the intranet will damage the company's service system and cause leakage of key information. Therefore, the administrator requires that DeviceA should control users' network access rights to ensure intranet security. Because guest users move frequently, Portal authentication is used and the RADIUS server authenticates the users.
[DeviceA] web-auth-server server-source all-interface
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-ip 192.168.2.30
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url http://192.168.2.30:8080/portal
[DeviceA-web-auth-server-abc] quit
[DeviceA] authentication-profile name p1
