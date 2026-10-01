---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-10
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [1453, 1610]
sha256: 8e6fe8c53bd31a0eaa5552b33d560b3d04ae0a4842b65e0b52b285625e918b99
---

# Configure the authentication profile p1, bind the MAC access profile m1 to the authentication profile, specify the domain example.com as the forcible authentication domain in the authentication profile, set the user access mode to multi-authen, and set the maximum number of access users to 100.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] mac-access-profile m1
[DeviceA-authen-profile-p1] access-domain example.com force
[DeviceA-authen-profile-p1] authentication mode multi-authen max-user 100
[DeviceA-authen-profile-p1] quit
#
sysname DeviceA
#
authentication-profile name p1
 mac-access-profile m1
 access-domain example.com force
 authentication mode multi-authen max-user 100
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
mac-access-profile name m1
#
interface Vlanif10
 ip address 192.168.1.10 255.255.255.0
#
interface Vlanif20
 ip address 192.168.2.10 255.255.255.0
#
interface 10GE 1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
 authentication-profile p1
#
interface 10GE 1/0/2
 port link-type access
 port default vlan 20
#
return 
In Figure 3-116, dumb terminals are connected to the company intranet through DeviceA. To ensure intranet security, the administrator requires that DeviceA control the network access rights of users.
Because dumb terminals cannot have the authentication client installed, MAC address authentication needs to be configured on DeviceA. In addition, local authentication is used to authenticate users.
# Configure the authorization scheme b1 and set the authorization mode to local authorization.
[DeviceA-aaa] authorization-scheme b1
[DeviceA-aaa-author-b1] authorization-mode local
[DeviceA-aaa-author-b1] quit
# Configure the user name, password, and access type of a local user.
Configure a terminal's MAC address as the local user name, set the password to Example@123, and set the access type to MAC address authentication (dot1x). The following assumes that the MAC address of printer 1 is 00e0-fcd4-8828.
[DeviceA-aaa] local-access-user 00e0-fcd4-8828
[DeviceA-aaa-access-user-00e0-fcd4-8828] password cipher Example@123
[DeviceA-aaa-access-user-00e0-fcd4-8828] service-type dot1x
Warning: If the service type of local access users is configured multiple times, the new configuration overwrites the previous one. The new configuration may cause local access users unable to go online. Continue? [Y/N]:y
[DeviceA-aaa-access-user-00e0-fcd4-8828] quit
# Configure the domain example.com, and apply the authentication scheme a1 and authorization scheme b1 to the domain.
[DeviceA-aaa] domain example.com
[DeviceA-aaa-domain-example.com] authentication-scheme a1
[DeviceA-aaa-domain-example.com] authorization-scheme b1
[DeviceA-aaa-domain-example.com] quit
[DeviceA-aaa] quit
When AAA local authentication and authorization are used, the user name and password used for MAC address authentication must be the same as those of the AAA local user. In this example, the user name of the local user is the terminal's MAC address containing hyphens (-), and the password is Example@123.
[DeviceA] mac-access-profile name m1
[DeviceA-mac-access-profile-m1] mac-authen username macaddress format with-hyphen password cipher Example@123
[DeviceA-mac-access-profile-m1] quit
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE 1/0/1] authentication-profile p1
[DeviceA-10GE 1/0/1] quit
#
sysname DeviceA
#
authentication-profile name p1
 mac-access-profile m1
 access-domain example.com force
#
vlan batch 10 20
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
mac-access-profile name m1
 mac-authen username macaddress format with-hyphen password cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.vW-9,<\Dz<6:AiLq@Ls2I%l3/!!!!!2jp5!!!!!!<!!!!5v2C4u&Z-8<3C6-//E:GF\uP8a9Le$7xWEF!!!!!%+%#
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
As shown in Figure 3-117, DeviceA functioning as a WAC in an enterprise is directly connected to an AP. The enterprise deploys the WLAN named wlan-net to provide wireless network access for employees. The WAC functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
As the WLAN is open to users, there are potential security risks to enterprise information if no access control is configured for the WLAN. To meet the enterprise's security requirements, MAC address authentication is used for authenticating dumb terminals such as wireless printers and wireless phones that cannot have an authentication client installed. In this case, MAC addresses of terminals are used as user information and sent to the RADIUS server for authentication. When users connect to the WLAN, they do not need to enter their user names and passwords for authentication.
| Configuration Item | Data | 
|---|---|
| RADIUS authentication parameters |  | 
| MAC access profile |  | 
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
# Configure the WAC and add 10GE 1/0/1 to VLAN 100 (management VLAN).
In this example, the service data forwarding mode is tunnel forwarding. In this mode, the management VLAN and service VLAN cannot be the same.
In direct forwarding mode, you are advised to configure port isolation on the WAC's interface 10GE 1/0/1 connected to the AP. If port isolation is not configured, unnecessary broadcast packets will be transmitted in VLANs or WLAN users on different APs will be able to communicate with each other at Layer 2.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 100 101
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] portswitch
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk pvid vlan 100
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 100
[DeviceA-10GE1/0/1] quit
# Add the WAC's uplink interface 10GE 1/0/2 to VLAN 101 (service VLAN).
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] portswitch
[DeviceA-10GE1/0/2] port link-type trunk
[DeviceA-10GE1/0/2] port trunk allow-pass vlan 101
[DeviceA-10GE1/0/2] quit
# Configure the WAC as a DHCP server to assign an IP address to the AP from the IP address pool on VLANIF 100, and assign IP addresses to STAs from the IP address pool on VLANIF 101.
[DeviceA] dhcp enable
[DeviceA] interface vlanif 100
[DeviceA-Vlanif100] ip address 10.23.100.1 24
[DeviceA-Vlanif100] dhcp select interface
[DeviceA-Vlanif100] quit
[DeviceA] interface vlanif 101
