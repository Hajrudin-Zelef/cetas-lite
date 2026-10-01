---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-31
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [5433, 5659]
sha256: d987ea91f21841f544acc3e9f103e2693f3e1e6c07b43e407b8919b2c7a52dfc
---

# Configure DeviceA to generate a local key pair.

Employees need to access the network using domain names and interact with third-party software or services (such as Weibo or Microsoft). In this context, access to the software or services must be permitted before employees pass the authentication. As the services are deployed in distributed mode, one domain name corresponds to multiple IP addresses that may change. As such, you need to permit access to domain names of these services on an authentication device.
| Item | Data | 
|---|---|
| RADIUS authentication parameters |  | 
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
[DeviceA] dns resolve
[DeviceA] dns server 10.23.200.2
[DeviceA] passthrough-domain name .microsoftonline.com id 1
[DeviceA] acl number 6001
[DeviceA-acl-ucl-6001] rule permit ip destination passthrough-domain .microsoftonline.com 
[DeviceA-acl-ucl-6001] rule permit ip destination ip-address 10.23.200.2 0.0.0.0 
[DeviceA-acl-ucl-6001] quit
[DeviceA] free-rule-template name default_free_rule
[DeviceA-free-rule-default_free_rule] free-rule acl 6001
[DeviceA-free-rule-default_free_rule] quit
#
sysname DeviceA
#
vlan batch 100 to 101
#
dhcp enable
#
dns server 10.23.200.2 
#
dns resolve
#
passthrough-domain name .microsoftonline.com id 1
#
acl number 6001
 rule 5 permit ip destination passthrough-domain .microsoftonline.com
 rule 15 permit ip destination ip-address 10.23.200.2 0
#
free-rule-template name default_free_rule
 free-rule acl 6001
# 
radius-server template radius_huawei
 radius-server shared-key cipher %^%#Oc6_BMCw#9gZ2@SMVtk!PAC6>Ou*eLW/"qLp+f#$%^%#
 radius-server authentication 10.23.200.1 1812 weight 80
 radius-server accounting 10.23.200.1 1813 weight 80
#
url-template name url1  
 url http://10.23.200.1:19008/portal  
 url-parameter set device-ip 10.23.100.1 
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
authentication-profile name p1
 portal-access-profile portal1
 free-rule-template default_free_rule
 access-domain example.com force
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
On the network shown in Figure 3-145, users in a company's guest area access the company's intranet through DeviceA. Unauthorized access to the intranet will damage the company's service system and cause leakage of key information. Therefore, the administrator requires that DeviceA should control users' network access rights to ensure intranet security. Because guests move frequently, Portal authentication is used and the RADIUS server authenticates them.
[DeviceA] dns resolve
[DeviceA] dns server 192.168.2.100
[DeviceA] ucl-group name group1
[DeviceA-ucl-group-group1] domain domain-name .microsoftonline.com //Configure a domain name in a static UCL group.
[DeviceA-ucl-group-group1] quit
[DeviceA] acl number 6001 
[DeviceA-acl-ucl-6001] rule 1 permit ip destination ucl-group name group1 //Configure an ACL rule to permit packets destined for the domain name in the static UCL group.
[DeviceA-acl-ucl-6001] rule permit ip destination ip-address 192.168.2.100 0.0.0.0 //Configure an ACL rule to permit traffic destined for the DNS server.
[DeviceA-acl-ucl-6001] quit
[DeviceA] free-rule-template name default_free_rule
[DeviceA-free-rule-default_free_rule] free-rule acl 6001 //Configure an authentication-free rule to permit access to domain names.
[DeviceA-free-rule-default_free_rule] quit
[DeviceA] dns snooping ttl delay-time 5760
[DeviceA] dns snooping server-ip-address 192.168.2.100
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] dns snooping enable //Enable DNS snooping on an uplink interface.
[DeviceA-10GE1/0/2] quit
#
sysname DeviceA
#
vlan batch 10 20
#
dns server 192.168.2.100 
#
dns resolve
#
dns snooping ttl delay-time 5760
dns snooping server-ip-address 192.168.2.100
#
ucl-group 1 name group1
 domain domain-name .microsoftonline.com
#
acl number 6001
 rule 5 permit ip destination ucl-group name group1
 rule 10 permit ip destination ip-address 192.168.2.100 0
#
free-rule-template name default_free_rule 
 free-rule acl 6001
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
radius-server template rd1 
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 192.168.2.30 1812 weight 80
 radius-server accounting 192.168.2.30 1813 weight 80
#
web-auth-server server-source all-interface
web-auth-server abc
 server-ip 192.168.2.30
 port 50200
 shared-key cipher %^%#4~ZXE3]6@BXu;2;aw}hA{rSb,@"L@T#e{%6G1AiD%^%#
 url http://192.168.2.30:8445/portal
#
portal-access-profile name portal1
 web-auth-server abc
#
authentication-profile name p1
 portal-access-profile portal1
 access-domain example.com force
#
interface Vlanif10
 ip address 192.168.1.10 255.255.255.0
#
interface Vlanif20
 ip address 192.168.2.10 255.255.255.0
#
interface 10GE1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
 authentication-profile p1
#
interface 10GE1/0/2
 port link-type access
 port default vlan 20
 dns snooping enable
#
return 
In Figure 3-146, user PCs access the network through DeviceE, and DeviceE connects to DeviceA and DeviceB using M-LAG. When both DeviceA and DeviceB work properly, traffic is load balanced to them. In addition, services will not be affected if any of the two devices fails, ensuring high service reliability.
