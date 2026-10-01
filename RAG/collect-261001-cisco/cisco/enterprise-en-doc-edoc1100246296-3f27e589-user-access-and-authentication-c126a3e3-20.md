---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-20
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [3357, 3518]
sha256: e852ba5b7265f9b4b9c5adaad49c575f4be1b3f78ae23709d019b5a59d591e34
---

# Configure DeviceA to generate a local key pair.

[DeviceA-web-auth-server-server1] port 50200
[DeviceA-web-auth-server-server1] url http://10.7.66.66:8080/portal
[DeviceA-web-auth-server-server1] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-server1] server-detect max-times 3 interval 60
[DeviceA-web-auth-server-server1] quit
[DeviceA] web-auth-server server2
[DeviceA-web-auth-server-server1] server-source ip-address 192.168.1.10  
[DeviceA-web-auth-server-server2] server-ip 10.7.66.67
[DeviceA-web-auth-server-server2] port 50200
[DeviceA-web-auth-server-server2] url http://10.7.66.67:8080/portal
[DeviceA-web-auth-server-server2] shared-key cipher YsHsjx_202206
[DeviceA-web-auth-server-server2] server-detect max-times 3 interval 60
[DeviceA-web-auth-server-server2] quit
[DeviceA] portal-access-profile name web1
[DeviceA-portal-access-profile-web1] web-auth-server server1 server2
[DeviceA-portal-access-profile-web1] quit
[DeviceA] acl 3001
[DeviceA-acl-adv-3001] rule 0 permit ip
[DeviceA-acl-adv-3001] quit
[DeviceA] acl 3001
[DeviceA-acl-adv-3001] rule 1 permit ip destination 192.168.10.166 0
[DeviceA-acl-adv-3001] rule 2 permit ip destination 192.168.11.31 0
[DeviceA-acl-adv-3001] rule 3 permit ip destination 192.168.12.26 0
[DeviceA-acl-adv-3001] rule 4 deny ip     //This deny rule must be configured, indicating that it is not allowed to access all network segments except the network segments allowed access by rules 1, 2, and 3.
[DeviceA-acl-adv-3001] quit
# Configure the rights granted to users when RADIUS servers are faulty. In this example, service scheme-based authorization is used during user authentication bypass. For other authorization modes, see (Optional) Configuring the Bypass Function.
[DeviceA] aaa
[DeviceA-aaa] service-scheme s1
[DeviceA-aaa-service-s1] acl-id 3001
[DeviceA-aaa-service-s1] quit
[DeviceA-aaa] quit
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] authentication event authen-server-down action authorize service-scheme s1
[DeviceA-authen-profile-p1] quit
# Configure the rights granted to users when Portal servers are faulty, and enable the re-authentication function when the Portal servers recover. In this example, service scheme-based authorization is used during user authentication bypass. For other authorization modes, see (Optional) Configuring the Portal Authentication Bypass Function.
[DeviceA] portal-access-profile name web1
[DeviceA-portal-access-profile-web1] authentication event portal-server-down action authorize service-scheme s1
[DeviceA-portal-access-profile-web1] authentication event portal-server-up action re-authen
[DeviceA-portal-access-profile-web1] quit
# Bind the Portal access profile web1 to the authentication profile, and configure the forcible authentication domain example.com for users using this authentication profile.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] portal-access-profile web1
[DeviceA-authen-profile-p1] access-domain example.com force
[DeviceA-authen-profile-p1] quit
# Bind the authentication profile p1 to 10GE 1/0/2 and enable Portal authentication on 10GE 1/0/2. The configurations of other downlink interfaces for user access are similar.
#
sysname DeviceA
#
vlan batch 10 20 
#
authentication-profile name p1
 authentication event authen-server-down action authorize service-scheme s1
 portal-access-profile web1
 access-domain example.com force
#
radius-server template rad1
 radius-server authentication 10.7.66.66 1812 weight 80
 radius-server accounting 10.7.66.66 1813 weight 80
 radius-server authentication 10.7.66.67 1812 weight 40
 radius-server accounting 10.7.66.67 1813 weight 40
 radius-server shared-key cipher %^%#b<4UC_J36%l@*;E]1\s6fJIY85mHu68SrhKtU%"B%^%#
 radius-server testuser username test1 password cipher %^%#O80v=d`~o>9:XXAS|o9'wrqD9hY+Q!El"EK:GT{G%^%#
#
web-auth-server server1
 server-source ip-address 192.168.1.10  
 server-ip 10.7.66.66
 port 50200
 url http://10.7.66.66:8080/portal
 shared-key cipher %^%#-;k-</gjVJ=ZK&Ea)<WB(j1FD8HJOGq^@$Ly=\0Y%^%#
 server-detect max-times 3 interval 60
web-auth-server server2
 server-source ip-address 192.168.1.10   
 server-ip 10.7.66.67
 port 50200
 url http://10.7.66.67:8080/portal
 shared-key cipher %^%#S/:6+:&#J0Sj|H:P2"MOcOh<!DZqX'kF.uSH'~Q'%^%#
 server-detect max-times 3 interval 60
#
portal-access-profile name web1
 web-auth-server server1 server2 
 authentication event portal-server-down action authorize service-scheme s1
 authentication event portal-server-up action re-authen
#
interface 10GE 1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface 10GE 1/0/2
 port link-type access
 port default vlan 20
 authentication-profile p1
#
interface vlanif 10
 ip address 192.168.1.10 255.255.255.0
#
interface vlanif 20
 ip address 192.168.2.10 255.255.255.0
#
acl 3001
 rule 0 permit ip
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
  radius-server rad1
#
ip route-static 10.7.66.0 255.255.255.0 192.168.1.1
#
return
On the enterprise network shown in Figure 3-134, DeviceA provides the native AC function and is connected to APs through DeviceB (access switch). The enterprise plans to deploy a wireless network named wlan-net to provide wireless access for employees. DeviceA functions as a DHCP server to assign IP addresses on the network segment 10.23.101.0/24 to STAs.
MAC address-prioritized Portal authentication can be enabled to allow users to connect to the wireless network without entering their user names and passwords when they move in and out of the wireless coverage area repeatedly within a period of time (for example, 60 minutes).
| Item | Data | 
|---|---|
| RADIUS authentication parameters | RADIUS authentication scheme: radius_test RADIUS accounting scheme: scheme1 RADIUS server template: radius_test  | 
| Portal parameters |  | 
| Authentication profile |  | 
| DHCP server | DeviceA functions as a DHCP server to assign IP addresses to STAs and APs.  | 
| IP address of the AC's source interface | VLANIF 100: 10.23.100.1/24 | 
| AP group |  | 
| Regulatory domain profile |  | 
| SSID profile |  | 
| Security profile |  | 
| VAP profile |  | 
| Authentication-free rule profile |  | 
MAC address-prioritized Portal authentication:
<HUAWEI> system-view
[HUAWEI] sysname DeviceB
[DeviceB] vlan batch 100
[DeviceB] interface 10ge 1/0/1
[DeviceB-10GE1/0/1] port link-type trunk
[DeviceB-10GE1/0/1] port trunk pvid vlan 100
[DeviceB-10GE1/0/1] port trunk allow-pass vlan 100
[DeviceB-10GE1/0/1] quit
[DeviceB] 10ge 1/0/2
[DeviceB-10GE1/0/2] port link-type trunk
[DeviceB-10GE1/0/2] port trunk pvid vlan 100
[DeviceB-10GE1/0/2] port trunk allow-pass vlan 100
[DeviceB-10GE1/0/2] quit
[DeviceB] interface 10ge 1/0/3
[DeviceB-10GE1/0/3] port link-type trunk
[DeviceB-10GE1/0/3] port trunk allow-pass vlan 100
[DeviceB-10GE1/0/3] quit
# Configure DeviceA and add 10GE 1/0/1 to VLAN 100 (management VLAN).
In this example, the service data forwarding mode is tunnel forwarding. If direct forwarding is used, you are advised to configure port isolation on GE 0/0/1 that connects the AC to an AP. If port isolation is not configured, unnecessary broadcast packets will be transmitted in VLANs or WLAN users on different APs will be able to communicate with each other at Layer 2.
In tunnel forwarding mode, the management VLAN and service VLAN cannot be the same.
[DeviceA] vlan batch 100 101 200
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 100
[DeviceA-10GE1/0/1] quit
# Add the uplink interface 10GE 1/0/2 of DeviceA to VLAN 200 (VLAN for communicating with the RADIUS server).
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] port link-type trunk
[DeviceA-10GE1/0/2] port trunk allow-pass vlan 200
[DeviceA-10GE1/0/2] quit
