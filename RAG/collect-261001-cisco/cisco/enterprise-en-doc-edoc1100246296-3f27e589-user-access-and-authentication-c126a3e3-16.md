---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-16
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [2500, 2709]
sha256: c0a009aac73df1456bc9e65cc818be7992cae06e890c006acb9734d25034f2b6
---

# Configure DeviceA to generate a local key pair.

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
In Figure 3-125, DeviceA functioning as a WAC in an enterprise is directly connected to an AP. The enterprise needs to deploy the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
| Item | Data | 
|---|---|
| RADIUS authentication parameters |  | 
| SSL policy |  | 
| URL template |  | 
| Portal server template |  | 
| Portal access profile |  | 
| Authentication-free rule profile |  | 
| Authentication domain | The authentication domain is named example.com and is bound to:  | 
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
[DeviceA] url-template name url1
[DeviceA-url-template-url1] url https://10.23.200.1:8445/portal
[DeviceA-url-template-url1] url-parameter login-url wac_url https://10.23.100.1:8443
[DeviceA-url-template-url1] quit
The login-url parameter is used by terminals to send account information to the device. If a Portal server does not support the login URL setting, you need to run the url-parameter command with the login-url parameter specified. The IP address in the login-url parameter is the WAC's local IP address, which needs to be permitted in the authentication-free rule profile to ensure that there is a reachable route between the terminal and the IP address. The port number in the parameter is specified using the portal web-authen-server https command, and the default port number is 8443.
[DeviceA] pki realm abcd 
[DeviceA-pki-abcd] quit 
[DeviceA] pki import-certificate ca realm abcd pem filename test.cer
[DeviceA] ssl policy huawei
[DeviceA-ssl-policy-huawei] pki-domain abcd
[DeviceA-ssl-policy-huawei] quit
[DeviceA] portal web-authen-server https ssl-policy huawei
[DeviceA] portal web-authen-server server-source all-interface  
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] protocol http
[DeviceA-web-auth-server-abc] quit
Ensure that the Portal server IP address, URL, and Portal authentication port number in the URL are configured correctly and are the same as those on the Portal server.
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-ip 10.23.200.1
[DeviceA-web-auth-server-abc] url-template url1
[DeviceA-web-auth-server-abc] quit
[DeviceA] free-rule-template name default_free_rule
[DeviceA-free-rule-default_free_rule] free-rule 1 destination ip 10.23.200.2 mask 24
[DeviceA-free-rule-default_free_rule] free-rule 2 destination ip 10.23.100.1 mask 24
[DeviceA-free-rule-default_free_rule] quit
WAC
#
pki realm abcd 
#
 sysname AC
#    
vlan batch 100 to 101
#
authentication-profile name p1
 portal-access-profile portal1
 free-rule-template default_free_rule
 access-domain example.com force
#
portal web-authen-server https ssl-policy huawei 
portal web-authen-server server-source all-interface
#   
dhcp enable
#
radius-server template radius_huawei
 radius-server shared-key cipher %^%#Oc6_BMCw#9gZ2@SMVtk!PAC6>Ou*eLW/"qLp+f#$%^%#
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80
#
ssl policy huawei
 pki-domain abcd
#
free-rule-template name default_free_rule
 free-rule 1 destination ip 10.23.200.2 mask 255.255.255.0 
 free-rule 2 destination ip 10.23.100.1 mask 255.255.255.0
# 
url-template name url1
 url https://10.23.200.1:8445/portal
 url-parameter login-url wac_url https://10.23.100.1:8443
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
In Figure 3-126, DeviceA functioning as a WAC in an enterprise is directly connected to an AP. The enterprise needs to deploy the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
HTTPS is recommended, as it is more secure than HTTP. For related examples, see Example for Configuring Wireless Portal Authentication (Using the HTTPS Protocol).
| Item | Data | 
|---|---|
| RADIUS authentication parameters |  | 
| URL template |  | 
| Portal server template |  | 
| Portal access profile |  | 
| Authentication-free rule profile |  | 
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
[DeviceA] url-template name url1
[DeviceA-url-template-url1] url http://10.23.200.1:8080/portal
[DeviceA-url-template-url1] url-parameter login-url ac_url http://10.23.100.1:8000
[DeviceA-url-template-url1] quit
