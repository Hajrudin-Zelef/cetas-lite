---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-19
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [3191, 3356]
sha256: 1e7bfabb4b244ba626695915868601912e1a6d5dc656c59a2caef3c24f653ae0
---

# Configure DeviceA to generate a local key pair.

In Figure 3-132, users in a company's guest area access the company's intranet through DeviceA. Unauthorized access to the intranet will damage the company's service system and cause leakage of key information. Therefore, the administrator requires that DeviceA should control users' network access rights to ensure intranet security.
# Create VLANs 10 and 20.
# Configure 10GE 1/0/1 connecting DeviceA to users as an access interface and add the interface to VLAN 10.
[DeviceA] interface 10ge 1/0/1 
[DeviceA-10GE1/0/1] port link-type access
[DeviceA-10GE1/0/1] port default vlan 10 
[DeviceA-10GE1/0/1] quit
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ipv6 enable
[DeviceA-Vlanif10] ipv6 address fc00:2::1 112
[DeviceA-Vlanif10] ipv6 nd ra halt disable
[DeviceA-Vlanif10] quit
# Configure 10GE 1/0/2 connecting DeviceA to the RADIUS server as an access interface and add the interface to VLAN 20.
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] port link-type access
[DeviceA-10GE1/0/2] port default vlan 20 
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 20
[DeviceA-Vlanif20] ip address 192.168.2.10 24
[DeviceA-Vlanif20] ipv6 enable
[DeviceA-Vlanif20] ipv6 address fc00:3::1 112
[DeviceA-Vlanif20] ipv6 nd ra halt disable
[DeviceA-Vlanif20] quit
[DeviceA] ip route-static 192.168.3.0 255.255.255.0 192.168.2.1
[DeviceA] ipv6 route-static fc00:1:: 112 fc00:3::2
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication fc00:1::1 1812
[DeviceA-radius-rd1] radius-server accounting fc00:1::1 1813
[DeviceA-radius-rd1] radius-server shared-key cipher YsHsjx_202206139
[DeviceA-radius-rd1] quit
[DeviceA-aaa] accounting-scheme scheme1 
[DeviceA-aaa-accounting-scheme1] accounting-mode radius 
[DeviceA-aaa-accounting-scheme1] accounting realtime 15 
[DeviceA-aaa-accounting-scheme1] quit 
# Create the authentication domain example.com and bind the AAA authentication scheme, AAA accounting scheme, and RADIUS server template to the authentication domain.
[DeviceA-aaa] domain example.com
[DeviceA-aaa-domain-example.com] authentication-scheme abc
[DeviceA-aaa-domain-example.com] accounting-scheme scheme1
[DeviceA-aaa-domain-example.com] radius-server rd1
[DeviceA-aaa-domain-example.com] quit
[DeviceA-aaa] quit
# Check whether a user can pass RADIUS authentication. (The test user test and password YsHsjx_2022061 have been configured on the RADIUS server.)
[DeviceA] web-auth-server abc
[DeviceA-web-auth-server-abc] server-source ip-address 192.168.2.10 
[DeviceA-web-auth-server-abc] server-ip 192.168.3.30
[DeviceA-web-auth-server-abc] server-ip ipv6 fc00:1::1
[DeviceA-web-auth-server-abc] port 50200
[DeviceA-web-auth-server-abc] url http://[192.168.3.30]:8445/portal
[DeviceA-web-auth-server-abc] url ipv6 http://[FC00:1::1]:8445/portal
[DeviceA-web-auth-server-abc] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-abc] quit
Ensure that the port number configured on the device is the same as that used by the Portal server.
[DeviceA] portal-access-profile name web1
[DeviceA-portal-access-profile-web1] web-auth-server abc
[DeviceA-portal-access-profile-web1] quit
# Configure the authentication profile p1, bind the Portal access profile web1 to the authentication profile, specify the domain example.com as the forcible authentication domain in the authentication profile, set the user access mode to multi-authen, and set the maximum number of access users to 100.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] portal-access-profile web1
[DeviceA-authen-profile-p1] access-domain example.com force
[DeviceA-authen-profile-p1] authentication mode multi-authen max-user 100
# Bind the authentication profile p1 to 10GE 1/0/1 and enable Portal authentication on 10GE 1/0/1.
#
sysname DeviceA
#
vlan batch 10 20 
#
authentication-profile name p1
 portal-access-profile web1
 access-domain example.com force
 authentication mode multi-authen max-user 100
#
radius-server template rd1
 radius-server authentication fc00:1::1 1812
 radius-server accounting fc00:1::1 1813
 radius-server shared-key cipher %^%#b<4UC_J36%l@*;E]1\s6fJIY85mHu68SrhKtU%"B%^%#
#
web-auth-server abc
 server-source ip-address 192.168.2.10
 server-ip 192.168.3.30
 server-ip ipv6 fc00:1::1
 port 50200
 url http://[192.168.3.30]:8445/portal
 url ipv6 http://[FC00:1::1]:8445/portal
 shared-key cipher %^%#S/:6+:&#J0Sj|H:P2"MOcOh<!DZqX'kF.uSH'~Q'%^%#
#
portal-access-profile name web1
 web-auth-server abc
#
interface 10GE 1/0/1
 port link-type access
 port default vlan 10
 authentication-profile p1
#
interface 10GE 1/0/2
 port link-type access
 port default vlan 20
#
interface vlanif 10
 ipv6 enable
 ip address 192.168.1.10 255.255.255.0
 ipv6 address fc00:2::1 112
 ipv6 nd ra halt disable
#
interface vlanif 20
 ipv6 enable
 ip address 192.168.2.10 255.255.255.0
 ipv6 address fc00:3::1 112
 ipv6 nd ra halt disable
#
aaa
 authentication-scheme abc
  authentication-mode radius
 accounting-scheme scheme1 
  accounting-mode radius 
  accounting realtime 15
 domain example.com
  authentication-scheme abc
  accounting-scheme scheme1
  radius-server rd1
#
ip route-static 192.168.3.0 255.255.255.0 192.168.2.1
ipv6 route-static fc00:1:: 112 fc00:3::2
#
return
On the enterprise network shown in Figure 3-133, DeviceA functions as an access device, and two RADIUS/Portal servers are deployed for Portal and RADIUS authentication of users on the enterprise network. Users can access the Internet only after being successfully authenticated. The administrator has the following user authentication requirements: When the two RADIUS/Portal servers are faulty, users bypass authentication and are granted the same network access rights as they are successfully authenticated. After the RADIUS/Portal servers recover, users are re-authenticated and re-authorized by the RADIUS servers.
The Portal shared key and port number configured on the Portal server must be the same as the shared key configured on the device and the port number used by the device to listen to Portal protocol packets, respectively.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 10 20
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] port link-type access
[DeviceA-10GE1/0/2] port default vlan 20
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 20
[DeviceA-Vlanif20] ip address 192.168.2.10 24
[DeviceA-Vlanif20] quit
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
[DeviceA-10GE1/0/1] quit
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 192.168.1.10 24
[DeviceA-Vlanif10] quit
# Create the RADIUS server template rad1.
[DeviceA] radius-server template rad1
[DeviceA-radius-rad1] radius-server authentication 10.7.66.66 1812 weight 80
[DeviceA-radius-rad1] radius-server accounting 10.7.66.66 1813 weight 80
[DeviceA-radius-rad1] radius-server authentication 10.7.66.67 1812 weight 40
[DeviceA-radius-rad1] radius-server accounting 10.7.66.67 1813 weight 40
[DeviceA-radius-rad1] radius-server algorithm master-backup
[DeviceA-radius-rad1] radius-server shared-key cipher YsHsjx_202206139
[DeviceA-radius-rad1] radius-server testuser username test1 password cipher YsHsjx_2022061
[DeviceA-radius-rad1] radius-server detect-server interval 60
[DeviceA-radius-rad1] radius-server detect-server timeout 3
[DeviceA-radius-rad1] radius-server retransmit 3 timeout 5
[DeviceA-radius-rad1] quit
# Configure the domain example, and apply the authentication scheme auth, accounting scheme acc, and RADIUS server template rad1 to the domain.
[DeviceA-aaa] domain example
[DeviceA-aaa-domain-example] authentication-scheme auth
[DeviceA-aaa-domain-example] accounting-scheme acc
[DeviceA-aaa-domain-example] radius-server rad1
[DeviceA-aaa-domain-example] quit
[DeviceA-aaa] quit
[DeviceA] web-auth-server server1
[DeviceA-web-auth-server-server1] server-source ip-address 192.168.1.10  
[DeviceA-web-auth-server-server1] server-ip 10.7.66.66
