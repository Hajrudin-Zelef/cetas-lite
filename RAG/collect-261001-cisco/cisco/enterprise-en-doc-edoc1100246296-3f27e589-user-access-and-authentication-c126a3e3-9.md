---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-9
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [1242, 1452]
sha256: 9387811849f730ad0c32df1cf6b00a3681308810dfd3f62ff3dadbcf04bf7353
---

# Configure DeviceA to generate a local key pair.

[DeviceA-aaa] smart-access-account mac-address 00e0-fcd4-8828 service-scheme asd description smart-access-account state active
[DeviceA-aaa] quit
[DeviceA] time-range t1 from 08:30 2025/10/1 to 18:00 2025/10/1
[DeviceA] aaa
[DeviceA-aaa] service-scheme asd
[DeviceA-aaa-service-asd] quit
[DeviceA-aaa] smart-access-policy name p1
[DeviceA-aaa-smart-access-policy-p1] rule 1 match time-range t1 apply service-scheme asd
[DeviceA-aaa-smart-access-policy-p1] quit
[DeviceA-aaa] quit
# Enable MAC address authentication on 10GE 1/0/1 of DeviceA and configure smart access authentication.
[DeviceA] mac-access-profile name d1
[DeviceA-mac-access-profile-d1] smart-access enable
[DeviceA-mac-access-profile-d1] quit
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] mac-access-profile d1
[DeviceA-authen-profile-p1] quit
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] authentication-profile p1
[DeviceA-10GE1/0/1] quit
# Configure an authentication-free rule to allow packets from the CAPWAP tunnel management VLAN.
[DeviceA] free-rule-template name default_free_rule
[DeviceA-free-rule-default_free_rule] free-rule 1 source vlan 20
[DeviceA-free-rule-default_free_rule] quit
Administrators can run the display access-user access-type mac-authen smart-access command on the device to check information about online smart access authentication users.
Device_1
#
sysname Device_1
#
vlan batch 10 20
#
as-mode enable
#
authentication-profile name p1
 mac-access-profile m1
#
as access interface vlanif 20
#
interface Vlanif20
 ip address dhcp-alloc
#
interface 10GE1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10 20
#
interface 10GE1/0/2
 port link-type access
 port default vlan 10
 authentication-profile p1
 authentication access-point
#
mac-access-profile name m1
# 
return  
#
sysname DeviceA
#
vlan batch 10 20
#
time-range t1 from 08:30 2025/10/1 to 18:00 2025/10/1
#
authentication-profile name p1
 mac-access-profile m1
#
dhcp enable
# 
free-rule-template name default_free_rule
 free-rule 1 source vlan 20
# 
aaa
 authentication-scheme a1
  authentication-mode local
 service-scheme asd
 domain isp1
  authentication-scheme a1
 smart-access-account mac-address 00e0-fcd4-8828 service-scheme asd description smart-access-account state active
 smart-access-policy name p1
  rule 1 match time-range t1 apply service-scheme asd
#
domain isp1 access
#
smart-access enable
#
interface Vlanif10
 ip address 192.168.2.1 255.255.255.0
 dhcp select interface
#
interface Vlanif20
 ip address 192.168.3.1 255.255.255.0
 dhcp select interface
#
interface 10GE1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10 20
 authentication-profile p1
 authentication control-point
#
capwap source interface vlanif 20
#
as-auth
 auth-mode none
#
mac-access-profile name m1
 smart-access enable
#  
return  
In Figure 3-114, user terminals are connected to the company intranet through DeviceA. To ensure efficient and secure terminal access when no AAA server is deployed in the company, the administrator wants to configure smart access on DeviceA.
# Create VLANs, configure the allowed VLANs on interfaces, and configure IP addresses for interfaces.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 10 20
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] portswitch
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
[DeviceA-10GE1/0/1] quit
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 192.168.1.10 24
[DeviceA-Vlanif10] quit
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] portswitch
[DeviceA-10GE1/0/2] port link-type access
[DeviceA-10GE1/0/2] port default vlan 20
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 20
[DeviceA-Vlanif20] ip address 192.168.2.10 24
[DeviceA-Vlanif20] quit
# Create the domain example.com and apply the authentication scheme a1 to the domain.
[DeviceA-aaa] domain example.com
[DeviceA-aaa-domain-example.com] authentication-scheme a1
[DeviceA-aaa-domain-example.com] quit
[DeviceA-aaa] quit
# Configure the MAC access profile m1 and configure smart access authentication.
[DeviceA] mac-access-profile name m1
[DeviceA-mac-access-profile-m1] smart-access enable
[DeviceA-mac-access-profile-m1] quit
# Configure the authentication profile p1, bind the MAC access profile m1 to the authentication profile, and specify the domain example.com as the forcible authentication domain in the authentication profile.
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] mac-access-profile m1
[DeviceA-authen-profile-p1] access-domain example.com force
[DeviceA-authen-profile-p1] quit
# Bind the authentication profile p1 to 10GE 1/0/1 and enable MAC address authentication on 10GE 1/0/1.
#
sysname DeviceA
#
authentication-profile name p1
 mac-access-profile m1
 access-domain example.com force
#
vlan batch 10 20
#
time-range t1 from 08:30 2025/10/1 to 18:00 2025/10/1
#
aaa
 authentication-scheme a1    
  authentication-mode local
 service-scheme asd
 domain example.com            
  authentication-scheme a1
 smart-access-account mac-address 00e0-fcd4-8828 service-scheme asd description smart-access-account state active
 smart-access-policy name p1
  rule 1 match time-range t1 apply service-scheme asd
#
smart-access enable
#  
mac-access-profile name m1
 smart-access enable
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
In Figure 3-115, dumb terminals are connected to the company intranet through DeviceA. To ensure intranet security, the administrator requires that DeviceA control the network access rights of users.
Because dumb terminals cannot have the authentication client installed, MAC address authentication needs to be configured on DeviceA. As a result, MAC addresses of the dumb terminals are used as user information and sent to the RADIUS server for authentication.
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 192.168.2.30 1812
[DeviceA-radius-rd1] radius-server accounting 192.168.2.30 1813
[DeviceA-radius-rd1] radius-server shared-key cipher Huawei@123456789
[DeviceA-radius-rd1] quit
# Create the authentication scheme abc and set the authentication mode to RADIUS authentication.
# Create the accounting scheme scheme2 and set the accounting mode to RADIUS accounting.
[DeviceA-aaa] accounting-scheme scheme2
[DeviceA-aaa-accounting-scheme2] accounting-mode radius
[DeviceA-aaa-accounting-scheme2] accounting realtime 15
[DeviceA-aaa-accounting-scheme2] quit
# Create the authentication domain example.com, and bind the authentication scheme abc, accounting scheme scheme2, and RADIUS server template rd1 to the domain.
[DeviceA-aaa] domain example.com
[DeviceA-aaa-domain-example.com] authentication-scheme abc
[DeviceA-aaa-domain-example.com] accounting-scheme scheme2
[DeviceA-aaa-domain-example.com] radius-server rd1
[DeviceA-aaa-domain-example.com] quit
[DeviceA-aaa] quit
# Check whether a user can pass RADIUS authentication. (The test user test1 and password test1@123 have been configured on the RADIUS server.)
[DeviceA] test-aaa test1 test1@123 radius-template rd1
Info: Account test succeeded.
# Configure the MAC access profile m1.
In a MAC access profile, the user name and password for MAC address authentication are MAC addresses without hyphens (-) by default. Ensure that the user name and password used for MAC address authentication configured on the RADIUS server are the same as those configured on the access device.
[DeviceA] mac-access-profile name m1
[DeviceA-mac-access-profile-m1] quit
