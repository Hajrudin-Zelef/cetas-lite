---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-8
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [1079, 1241]
sha256: 0edc30db5fe9ccaf7f628c2fb0dbbf74119bc67cd9d6c0f7943f1e778991536a
---

# After a user enters the user name test and password YsHsjx_202206 to log in to the device through Telnet, you can run the display access-user domain and display access-user user-id commands on the device to check the domain and access type of the user.
<DeviceA> display access-user domain default_admin
 ------------------------------------------------------------------------------
 UserID Username             IP address             MAC               Status
 ------------------------------------------------------------------------------
 16009  test                 10.135.18.217          -                 Success
 ------------------------------------------------------------------------------
 Total: 1, printed: 1
<DeviceA> display access-user user-id 16009
 Total: 1                                                               
 ------------------------------------------------------------------------------ 
Basic:
  User id                         : 16009
  User name                       : test
  Domain-name                     : default_admin
  User MAC                        : -
  User IP address                 : 10.135.18.217
  User IPv6 address               : -
  User access time                : 2024/02/15 05:10:52
  User accounting session ID      : example255255000000000f****2016009
  Option82 information            : -
  User access type                : SSH
AAA:
  User authentication type        : Administrator authentication
  Current authentication method   : Local
  Current authorization method    : Local
  Current accounting method       : None
 ------------------------------------------------------------------------------
#
sysname DeviceA
#
vlan batch 10 to 11
#
authentication-profile name p1
 dot1x-access-profile d1
 authentication mode multi-authen max-user 100
#
radius-server template rd1
 radius-server shared-key cipher %^%#Q75cNQ6IF(e#L4WMxP~%^7'u17,]D87GO{"[o]`D%^%#
 radius-server authentication 192.168.2.30 1812 weight 80
 radius-server accounting 192.168.2.30 1813 weight 80
 radius-server retransmit 2
#
aaa
 authentication-scheme abc
  authentication-mode radius
 authentication-scheme auth
 authorization-scheme autho
 accounting-scheme abc
  accounting-mode radius
 domain default
  authentication-scheme abc
  accounting-scheme abc
  radius-server rd1
 domain default_admin
  authentication-scheme auth
  authorization-scheme autho
 local-user user1-huawei password irreversible-cipher $1d$OwseVRh@LH}ZeTBm$1nH4$ab>d(N{-%0!ab48y=Ic*xEUR4pVhR2"9-~,$
 local-user user1-huawei privilege level 3
 local-user user1-huawei service-type ssh   
#
interface Vlanif11
 ip address 192.168.2.29 255.255.255.0
#
interface 10GE1/0/1
 port link-type access
 port default vlan 10
 authentication-profile p1
# 
interface 10GE1/0/2
 port link-type access
 port default vlan 11
#
stelnet server enable 
ssh server-source all-interface
#
user-interface vty 0 4           
 authentication-mode aaa           
 protocol inbound ssh
# 
dot1x-access-profile name d1
#  
return
In Figure 3-113, the gateway DeviceA functions as an authentication control device and Device_1 to Device_N function as authentication access devices. Smart access authentication needs to be performed for users on the authentication control device and their access policies need to be implemented on the authentication access devices. VLAN 10 is a user VLAN, and VLAN 20 is the management VLAN of CAPWAP tunnels. MAC address authentication is used for user access authentication.
In the policy association scenario, smart access needs to be configured on an authentication control device. For details about the authentication access device and authentication control device models, see Support for Policy Association.
# Create VLAN 10 and VLAN 20 on Device_1.
<HUAWEI> system-view
[HUAWEI] sysname Device_1
[Device_1] vlan batch 10 20
# On Device_1, configure 10GE 1/0/2 connected to users as an access interface and add it to VLAN 10.
[Device_1] interface 10ge 1/0/2
[Device_1-10GE1/0/2] port link-type access
[Device_1-10GE1/0/2] port default vlan 10 
[Device_1-10GE1/0/2] quit
# On Device_1, configure 10GE 1/0/1 connected to DeviceA as a trunk interface and configure the interface to permit packets from VLAN 10 and VLAN 20.
[Device_1] interface 10ge 1/0/1
[Device_1-10GE1/0/1] port link-type trunk
[Device_1-10GE1/0/1] port trunk allow-pass vlan 10 20
[Device_1-10GE1/0/1] quit
# Create VLAN 10 and VLAN 20 on DeviceA.
# On DeviceA, configure 10GE 1/0/1 connected to Device_1 as a trunk interface and configure the interface to permit packets from VLAN 10 and VLAN 20.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 10 20
[DeviceA-10GE1/0/1] quit
[DeviceA] dhcp enable
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 192.168.2.1 255.255.255.0
[DeviceA-Vlanif10] dhcp select interface
[DeviceA-Vlanif10] quit
[Device_1] as-mode enable
Warning: If the device switches to the AS mode, UC configurations will be deleted. Continue? [Y/N]:y 
# Create VLANIF 20 on Device_1, configure VLANIF 20 to obtain an IP address through DHCP, and specify VLANIF 20 as the source interface of a CAPWAP tunnel.
[Device_1] interface vlanif 20
[Device_1-Vlanif20] ip address dhcp-alloc
[Device_1-Vlanif20] quit
[Device_1] as access interface vlanif 20
# Create VLANIF 20 on DeviceA, configure an address pool on VLANIF 20, and specify VLANIF 20 as the source interface of a CAPWAP tunnel.
[DeviceA] interface vlanif 20
[DeviceA-Vlanif20] ip address 192.168.3.1 255.255.255.0
[DeviceA-Vlanif20] dhcp select interface
[DeviceA-Vlanif20] quit
[DeviceA] capwap source interface vlanif 20 
//When the source interface of a CAPWAP tunnel is configured for the first time, the system displays a message, asking you to check whether security-related configurations exist, including the PSK for DTLS encryption. If DTLS encryption is disabled, the message does not affect policy association configurations.
[DeviceA] as-auth
[DeviceA-as-auth] auth-mode none
Warning: None authentication is configured, which has security risks. Continue? [Y/N]: y
[DeviceA-as-auth] quit
# Configure the authentication scheme a1 and set the authentication mode to local authentication.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme a1
[DeviceA-aaa-authen-a1] authentication-mode local
[DeviceA-aaa-authen-a1] quit
# Create the authentication domain isp1 and bind the AAA authentication scheme a1 and authorization scheme b1 to the domain.
[DeviceA-aaa] domain isp1
[DeviceA-aaa-domain-isp1] authentication-scheme a1
[DeviceA-aaa-domain-isp1] quit
[DeviceA-aaa] quit
# Configure the global default domain isp1.
[DeviceA] domain isp1 access
When a user enters the user name in the format of user@isp1 for access authentication, the user is authenticated in the authentication domain isp1. If the user name does not contain a domain name or contains an invalid domain name, the user is authenticated in the default domain.
# Configure 10GE 1/0/2 on Device_1 as an access point.
[Device_1] interface 10ge 1/0/2
[Device_1-10GE1/0/2] authentication access-point
[Device_1-10GE1/0/2] quit
# Configure 10GE 1/0/1 on DeviceA as the control point.
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] authentication control-point
[DeviceA-10GE1/0/1] quit
[Device_1] mac-access-profile name m1
[Device_1-mac-access-profile-d1] quit
[Device_1] authentication-profile name p1
[Device_1-authen-profile-p1] mac-access-profile m1
[Device_1-authen-profile-p1] quit
[Device_1] interface 10ge 1/0/2
[Device_1-10GE1/0/2] authentication-profile p1
[Device_1-10GE1/0/2] quit
# Enable smart access globally.
[DeviceA] smart-access enable
# Configure a smart access account and authorize a service scheme. You can select one or more of the following methods:
[DeviceA] aaa
[DeviceA-aaa] service-scheme asd
[DeviceA-aaa-service-asd] quit
