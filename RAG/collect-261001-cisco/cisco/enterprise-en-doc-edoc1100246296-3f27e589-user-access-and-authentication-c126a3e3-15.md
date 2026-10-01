---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-15
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [2313, 2499]
sha256: fcc79fae6a11d1a0bc2b974de36761fb6648863aa093b58e79103e4878944588
---

# Configure the domain example, and apply the authentication scheme auth, accounting scheme acc, and RADIUS server template controller to the domain.
[DeviceA-aaa] domain example
[DeviceA-aaa-domain-example] authentication-scheme auth
[DeviceA-aaa-domain-example] accounting-scheme acc
[DeviceA-aaa-domain-example] radius-server controller
[DeviceA-aaa-domain-example] quit
[DeviceA-aaa] quit
# Bind the 802.1X access profile d1 to the authentication profile, and configure the forcible authentication domain example for users using the authentication profile.
After a forcible domain is configured in an authentication profile, users using this authentication profile are authenticated in the domain regardless of whether the user names carry domain names or carry what kind of domain names.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] dot1x-access-profile d1
[DeviceA-authen-profile-p1] access-domain example force
[DeviceA-authen-profile-p1] quit
# Configure the rights granted to users when RADIUS servers are faulty, and enable the re-authentication function when the RADIUS servers recover. In this example, service scheme-based authorization is used during user authentication bypass. For other authorization modes, see (Optional) Configuring the Bypass Function.
[DeviceA] acl 3001
[DeviceA-acl-adv-3001] rule 1 permit ip source 192.168.2.0 0.0.0.255
[DeviceA-acl-adv-3001] quit
[DeviceA] aaa
[DeviceA-aaa] service-scheme s1
[DeviceA-aaa-service-s1] acl-id 3001
[DeviceA-aaa-service-s1] quit
[DeviceA-aaa] quit
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] authentication event authen-server-down action authorize service-scheme s1
[DeviceA-authen-profile-p1] authentication event authen-server-up action re-authen
[DeviceA-authen-profile-p1] quit
# Disable the pre-connection function.
[DeviceA] undo authentication pre-authen-access enable
# Bind the authentication profile p1 to 10GE 1/0/2 and enable 802.1X authentication on 10GE 1/0/2. The configurations of other interfaces connected to terminals are similar.
#
sysname DeviceA
#
undo authentication pre-authen-access enable 
authentication-profile name p1
 dot1x-access-profile d1
 access-domain example force
 authentication event authen-server-down action authorize service-scheme s1
 authentication event authen-server-up action re-authen
#
vlan batch 10 20
#
acl 3001 
 rule 1 permit ip source 192.168.2.0 0.0.0.255
#
aaa
 authentication-scheme auth    
  authentication-mode radius 
 accounting-scheme acc
  accounting-mode radius   
 service-scheme s1 
  acl-id 3001
 domain example            
  authentication-scheme auth
  accounting-scheme acc
  radius-server controller        
# 
radius-server template controller
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server accounting 10.7.66.66 1813 weight 80   
 radius-server accounting 10.7.66.67 1813 weight 40 
 radius-server authentication 10.7.66.66 1812 weight 80 
 radius-server authentication 10.7.66.67 1812 weight 40   
 radius-server testuser username test1 password cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
#
dot1x-access-profile name d1
 dot1x timer client-timeout 30 
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
#
interface 10GE1/0/2
 port link-type hybrid
 port hybrid pvid vlan 20
 port hybrid untagged vlan 20
 authentication-profile p1
#
ip route-static 10.7.66.0 255.255.255.0 192.168.1.1
#
return 
In Figure 3-124, DeviceA functioning as a WAC in an enterprise is directly connected to an AP. The enterprise needs to deploy the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
To reduce network security risks, you can deploy Portal authentication on the WAC. The WAC works with the RADIUS server (integrated with the Portal server) to implement access control on STAs that attempt to access the enterprise network, meeting the enterprise's security requirements.
| Item | Data | 
|---|---|
| RADIUS authentication parameters |  | 
| Portal server template |  | 
| Portal access profile |  | 
| Authentication-free rule profile |  | 
| Authentication domain | The authentication domain is named example.com and bound to:  | 
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
# Configure a URL template, configure the redirect URL for the Portal server, and specify parameters carried in the URL as the user IP address and user MAC address.
[DeviceA] url-template name url1 
[DeviceA-url-template-url1] url http://10.23.200.1:19008/portal 
[DeviceA-url-template-url1] url-parameter set device-ip 10.23.100.1
[DeviceA-url-template-url1] url-parameter device-ip ac-ip user-ipaddress uaddress user-mac umac 
[DeviceA-url-template-url1] quit
# Configure a Portal server template and bind it to the URL template.
[DeviceA] web-auth-server server-source all-interface
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-ip 10.23.200.1
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url-template url1
[DeviceA-web-auth-server-abc] quit
[DeviceA] portal-access-profile name portal1
[DeviceA-portal-access-profile-portal1] web-auth-server abc
[DeviceA-portal-access-profile-portal1] quit
[DeviceA] free-rule-template name default_free_rule
[DeviceA-free-rule-default_free_rule] free-rule 1 destination ip 10.23.200.2 mask 24
[DeviceA-free-rule-default_free_rule] quit
[DeviceA] authentication-profile name p1
[DeviceA-authentication-profile-p1] portal-access-profile portal1
[DeviceA-authentication-profile-p1] free-rule-template default_free_rule
[DeviceA-authentication-profile-p1] access-domain example.com force
[DeviceA-authentication-profile-p1] quit
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
