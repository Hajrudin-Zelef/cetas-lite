---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-18
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [2925, 3190]
sha256: b5c53c5abd019dcea6e7d1657104d27184109fcc21edf3d87bb72de2ba6b81c4
---

# Configure DeviceA to generate a local key pair.

[DeviceA-authentication-profile-p1] portal-access-profile portal1
[DeviceA-authentication-profile-p1] access-domain example.com force
[DeviceA-authentication-profile-p1] quit
#
sysname DeviceA
#
vlan batch 10 20
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
#
return 
[DeviceA] url-template name url1 
[DeviceA-url-template-url1] url http://192.168.2.30:19008/portal 
[DeviceA-url-template-url1] url-parameter set device-ip 192.168.2.10
[DeviceA-url-template-url1] url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac 
[DeviceA-url-template-url1] quit
[DeviceA] web-auth-server server-source all-interface
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-ip 192.168.2.30
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url-template url1
[DeviceA-web-auth-server-abc] quit
#
sysname DeviceA
#
vlan batch 10 20
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
url-template name url1 
 url http://192.168.2.30:19008/portal 
 url-parameter set device-ip 192.168.2.10
 url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac 
#
web-auth-server server-source all-interface
web-auth-server abc
 server-ip 192.168.2.30
 port 50200
 shared-key cipher %^%#4~ZXE3]6@BXu;2;aw}hA{rSb,@"L@T#e{%6G1AiD%^%#
 url-template url1
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
#
return 
[DeviceA] url-template name url1 
[DeviceA-url-template-url1] url https://192.168.2.30:19008/portal 
[DeviceA-url-template-url1] url-parameter set device-ip 192.168.2.10
[DeviceA-url-template-url1] url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac
[DeviceA-url-template-url1] quit
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-ip 192.168.2.30
[DeviceA-web-auth-server-abc] protocol http
[DeviceA-web-auth-server-abc] http-method post cmd-key cmd1
[DeviceA-web-auth-server-abc] url-template url1
[DeviceA-web-auth-server-abc] quit
[DeviceA] free-rule-template name default_free_rule
[DeviceA-free-rule-default_free_rule] free-rule 1 destination ip 192.168.1.10 mask 24
[DeviceA-free-rule-default_free_rule] quit
#
pki realm abcd 
#
sysname DeviceA
#
vlan batch 10 20
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
portal web-authen-server https ssl-policy huawei 
portal web-authen-server server-source all-interface
#   
radius-server template rd1 
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 192.168.2.30 1812 weight 80
 radius-server accounting 192.168.2.30 1813 weight 80
#
ssl policy huawei
 pki-domain abcd
#
free-rule-template name default_free_rule
 free-rule 1 destination ip 192.168.1.10 mask 255.255.255.0 
# 
url-template name url1 
 url https://192.168.2.30:19008/portal 
 url-parameter set device-ip 192.168.2.10
 url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac
#
web-auth-server abc
 server-ip 192.168.2.30
 protocol http  
 http-method post cmd-key cmd1
 url-template url1
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
#
return 
HTTPS is recommended, as it is more secure than HTTP. For related examples, see Example for Configuring Wired Portal Authentication (Using the HTTPS Protocol).
[DeviceA] portal web-authen-server http
[DeviceA] portal web-authen-server server-source ip-address 192.168.2.10
#
sysname DeviceA
#
vlan batch 10 20
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
portal web-authen-server http 
portal web-authen-server server-source ip-address 192.168.2.10
#   
radius-server template rd1 
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 192.168.2.30 1812 weight 80
 radius-server accounting 192.168.2.30 1813 weight 80
#
free-rule-template name default_free_rule
 free-rule 1 destination ip 192.168.1.10 mask 255.255.255.0 
# 
url-template name url1 
 url http://192.168.2.30:19008/portal 
 url-parameter set device-ip 192.168.2.10
 url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac 
#
web-auth-server abc
 server-ip 192.168.2.30
 protocol http  
 http-method post cmd-key cmd1
 url-template url1
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
#
return 
