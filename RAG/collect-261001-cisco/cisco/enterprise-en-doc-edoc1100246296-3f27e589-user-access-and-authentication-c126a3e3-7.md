---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-7
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [940, 1078]
sha256: d39f38b7dc9921001b9dcc08b6c755cd520f068dd6cf4ab82d38a9b4bcb919d1
---

# Configure DeviceA to generate a local key pair.

 local-user user1-huawei password irreversible-cipher $1d$OwseVRh@LH}ZeTBm$1nH4$ab>d(N{-%0!ab48y=Ic*xEUR4pVhR2"9-~,$
 local-user user1-huawei privilege level 3
 local-user user1-huawei service-type ssh   
#
domain huawei admin
# 
vlan batch 10 20
#
interface Vlanif10
 ip address 10.1.1.2 255.255.255.0
#
interface Vlanif20
 ip address 10.1.6.2 255.255.255.0
#
interface 10GE1/0/1 
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface 10GE1/0/2 
 port link-type trunk
 port trunk allow-pass vlan 20
#
return
In Figure 3-112, users connect to the network through DeviceA. The user names do not contain any domain name. It is required that common users connect to the network and obtain corresponding rights only after passing RADIUS authentication and that the administrator logs in to DeviceA only after passing local authentication.
Ensure that users have been configured on the RADIUS server. In this example, a user with the name test1 and password YsHsjx_202206 has been configured on the RADIUS server.
This example provides only the configuration of DeviceA. The RADIUS server configuration is not provided here.
# Create VLAN 11 on DeviceA.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 11
# On DeviceA, configure 10GE 1/0/2 connecting to the RADIUS server as an access interface, and add 10GE 1/0/1 to VLAN 11.
[DeviceA] interface 10ge 1/0/2
[DeviceA-10GE1/0/2] port link-type access
[DeviceA-10GE1/0/2] port default vlan 11
[DeviceA-10GE1/0/2] quit
# Create VLANIF 11 and configure the IP address 192.168.2.29/24 for it.
[DeviceA] interface vlanif 11
[DeviceA-Vlanif11] ip address 192.168.2.29 24
[DeviceA-Vlanif11] quit
# Create and configure the RADIUS server template rd1.
[DeviceA] radius-server template rd1
[DeviceA-radius-rd1] radius-server authentication 192.168.2.30 1812
[DeviceA-radius-rd1] radius-server accounting 192.168.2.30 1813
[DeviceA-radius-rd1] radius-server shared-key cipher YsHsjx_202206139
[DeviceA-radius-rd1] radius-server retransmit 2
[DeviceA-radius-rd1] quit
# Create the AAA authentication scheme abc and accounting scheme abc, and set the authentication mode and accounting mode to RADIUS.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme abc
[DeviceA-aaa-authen-abc] authentication-mode radius
[DeviceA-aaa-authen-abc] quit
[DeviceA-aaa] accounting-scheme abc
[DeviceA-aaa-accounting-abc] accounting-mode radius
[DeviceA-aaa-accounting-abc] quit
# Test the connection between DeviceA and the RADIUS server. (The following assumes that the test user test1 and password YsHsjx_202206 have been configured on the RADIUS server.)
[DeviceA-aaa] test-aaa test1 YsHsjx_202206 radius-template rd1
# Bind the AAA authentication scheme abc, AAA accounting scheme abc, and RADIUS server template rd1 to the domain default.
[DeviceA-aaa] domain default
[DeviceA-aaa-domain-default] authentication-scheme abc
[DeviceA-aaa-domain-default] accounting-scheme abc
[DeviceA-aaa-domain-default] radius-server rd1
[DeviceA-aaa-domain-default] quit
[DeviceA-aaa] quit
# Enable 802.1X authentication on an interface.
[DeviceA] dot1x-access-profile name d1
[DeviceA-dot1x-access-profile-d1] quit
[DeviceA] authentication-profile name p1
[DeviceA-authen-profile-p1] dot1x-access-profile d1
[DeviceA-authen-profile-p1] authentication mode multi-authen max-user 100
[DeviceA-authen-profile-p1] quit
[DeviceA] vlan batch 10
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] port link-type access
[DeviceA-10GE1/0/1] port default vlan 10 
[DeviceA-10GE1/0/1] authentication-profile p1
[DeviceA-10GE1/0/1] quit
# Set the domain default as the global default common domain. After common users enter their user names in the format of user@default, the device performs AAA authentication for these users in the domain default. If a user name does not contain a domain name or contains a nonexistent domain name, the device authenticates the user in the default common domain.
[DeviceA] domain default
[DeviceA] rsa local-key-pair create
The key name will be:Host
The range of public key size is (2048, 4098).
NOTE: Key pair generation will take a short while.
Please input the modulus [default = 3072]:3072 
[DeviceA] stelnet server enable 
[DeviceA] ssh server-source all-interface
# Create a local user named test, set the password to YsHsjx_202206, access type to SSH, and privilege level to 3.
[DeviceA] aaa
[DeviceA-aaa] local-user user1-huawei password irreversible-cipher YsHsjx_202206
[DeviceA-aaa] local-user user1-huawei service-type ssh
[DeviceA-aaa] local-user user1-huawei privilege level 3
[DeviceA-aaa] quit
# Enable the local account locking function, and set the authentication retry interval to 5 minutes, maximum number of consecutive failed password attempts to 3, and account lockout duration to 5 minutes.
[DeviceA-aaa] local-aaa-user wrong-password retry-interval 5 retry-time 3 block-time 5
# Configure an authentication scheme named auth and set the authentication mode to local authentication.
[DeviceA-aaa] authentication-scheme auth
[DeviceA-aaa-authen-auth] authentication-mode local
[DeviceA-aaa-authen-auth] quit
# Configure an authorization scheme named autho and set the authorization mode to local authorization.
[DeviceA-aaa] authorization-scheme autho
[DeviceA-aaa-author-autho] authorization-mode local
[DeviceA-aaa-author-autho] quit
# Configure the domain default_admin and apply the authentication scheme auth and authorization scheme autho to the domain.
[DeviceA-aaa] domain default_admin
[DeviceA-aaa-domain-default_admin] authentication-scheme auth
[DeviceA-aaa-domain-default_admin] authorization-scheme autho
[DeviceA-aaa-domain-default_admin] quit
[DeviceA-aaa] quit
# Set the domain default_admin as the global default administrative domain. After the administrator enters the user name in the format of user@default_admin, the device performs AAA authentication for the administrator in the domain default_admin. If the user name does not contain a domain name or contains a nonexistent domain name, the device authenticates the administrator in the default administrative domain.
[DeviceA] domain default_admin admin
[DeviceA] quit
# Run the display dot1x interface command on DeviceA to check 802.1X authentication information.
# After a common user enters the user name test1 and password YsHsjx_202206 on the 802.1X client to log in to the device, you can run the display access-user domain and display access-user user-id commands on the device to check the domain and access type of the user.
<DeviceA> display access-user domain default
 ------------------------------------------------------------------------------
 UserID Username             IP address             MAC               Status
 ------------------------------------------------------------------------------
 16040  test1                -                      00e0-fc01-31f6    Success
 ------------------------------------------------------------------------------
 Total: 1, printed: 1
<DeviceA> display access-user user-id 16040
 Total: 1                                                               
 ------------------------------------------------------------------------------ 
Basic:
  User id                         : 16040
  User name                       : test1
  Domain-name                     : default
  User MAC                        : 00e0-4c97-31f6
  User IP address                 : -
  User IPv6 address               : -
  User access time                : 2024/02/15 19:10:52
  User accounting session ID      : example255255000000000f****2016040
  Option82 information            : -
  User access type                : 802.1x
AAA:
  User authentication type        : 802.1x authentication
  Current authentication method   : RADIUS
  Current authorization method    : -
  Current accounting method       : RADIUS
 ------------------------------------------------------------------------------
