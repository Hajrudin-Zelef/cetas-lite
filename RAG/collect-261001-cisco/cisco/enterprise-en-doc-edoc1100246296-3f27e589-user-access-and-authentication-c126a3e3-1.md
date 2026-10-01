---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-1
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [1, 155]
sha256: 5be1f97c853c8268efa307e1ae7613b380b2a4719baf21b13f174c92a1f557ae
---

# Configure DeviceA to generate a local key pair.

Enterprise
In Figure 3-103, the enterprise requires that the administrator use AAA local authentication to log in to the device through STelnet. The specific requirements are as follows:
<HUAWEI> system-view 
[HUAWEI] sysname DeviceA 
[DeviceA] vlan batch 10
[DeviceA] interface 10ge 1/0/1
[DeviceA-10GE1/0/1] portswitch
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
[DeviceA-10GE1/0/1] quit
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 192.168.10.1 24
[DeviceA-Vlanif10] quit
# Configure DeviceA to generate a local key pair.
[DeviceA] rsa local-key-pair create
The key name will be:Host
The range of public key size is (2048, 4096).
NOTE: Key pair generation will take a short while.
Please input the modulus [default = 3072]:3072 
# Set the authentication mode and protocol for accessing VTY user interfaces 0 to 4 to AAA and SSH, respectively.
[DeviceA] user-interface vty 0 4 
[DeviceA-ui-vty0-4] authentication-mode aaa 
[DeviceA-ui-vty0-4] protocol inbound ssh
[DeviceA-ui-vty0-4] quit
# Enable the SSH server function on DeviceA.
[DeviceA] stelnet server enable 
[DeviceA] ssh server-source -i vlanif 10
# Set the authentication mode of all SSH users to password authentication.
[DeviceA] ssh authentication-type default password
If only the authentication mode and service type of a few SSH users need to be set to password authentication and STelnet, you can specify the SSH user name to set the authentication mode and service type of a single SSH user. For example, set the authentication mode and service type of an SSH user named admin to password authentication and STelnet, respectively.
[DeviceA] ssh user admin authentication-type password
[DeviceA] ssh user admin service-type stelnet
[DeviceA] aaa
[DeviceA-aaa] local-user user1-huawei password irreversible-cipher YsHsjx_202206
[DeviceA-aaa] local-user user1-huawei service-type ssh
Warning: If the service type of local users is configured multiple times, the new configuration overwrites the previous one. The new configuration may cause local user authentication to fail. Continue? [Y/N]:y 
[DeviceA-aaa] local-user user1-huawei privilege level 3
Warning: This operation may affect online users and will change the user privilege level, Continue? [Y/N]:y
[DeviceA-aaa] quit
The administrator can log in to DeviceA through the STelnet client after entering the correct user name and password.
#
sysname DeviceA
#
aaa
 local-user user1-huawei password irreversible-cipher $1d$OwseVRh@LH}ZeTBm$1nH4$ab>d(N{-%0!ab48y=Ic*xEUR4pVhR2"9-~,$
 local-user user1-huawei privilege level 3
 local-user user1-huawei service-type ssh   
# 
vlan batch 10
#
interface Vlanif10
 ip address 192.168.10.1 255.255.255.0
#
interface 10GE1/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
#
stelnet server enable 
ssh server-source -i Vlanif 10
#
user-interface vty 0 4           
 authentication-mode aaa           
 protocol inbound ssh
# 
return 
In Figure 3-104, a RADIUS server is deployed on an enterprise network. The enterprise requires that the administrator use RADIUS authentication to log in to DeviceA through STelnet. The specific requirements are as follows:
Ensure that the shared key in the RADIUS server template is the same as that configured on the RADIUS server.
If the RADIUS server does not support the user name containing a domain name, run the undo radius-server user-name domain-included command in the RADIUS server template view to configure the device to send packets that do not contain a domain name to the RADIUS server.
<HUAWEI> system-view
[HUAWEI] sysname DeviceA
[DeviceA] vlan batch 10 20
[DeviceA] interface 10ge 1/0/1 
[DeviceA-10GE1/0/1] portswitch
[DeviceA-10GE1/0/1] port link-type trunk
[DeviceA-10GE1/0/1] port trunk allow-pass vlan 10
[DeviceA-10GE1/0/1] quit
[DeviceA] interface 10ge 1/0/2 
[DeviceA-10GE1/0/2] portswitch
[DeviceA-10GE1/0/2] port link-type trunk
[DeviceA-10GE1/0/2] port trunk allow-pass vlan 20
[DeviceA-10GE1/0/2] quit
[DeviceA] interface vlanif 10
[DeviceA-Vlanif10] ip address 10.1.1.2 255.255.255.0
[DeviceA-Vlanif10] quit
[DeviceA] interface vlanif 20
[DeviceA-Vlanif20] ip address 10.1.6.2 255.255.255.0
[DeviceA-Vlanif20] quit
# Configure a RADIUS server template for communication between DeviceA and RADIUS server.
[DeviceA] radius-server template 1
[DeviceA-radius-1] radius-server authentication 10.1.6.6 1812
[DeviceA-radius-1] radius-server accounting 10.1.6.6 1813
[DeviceA-radius-1] radius-server shared-key cipher YsHsjx_202206139
[DeviceA-radius-1] quit
# Configure an AAA authentication scheme and set the authentication mode to RADIUS and local authentication.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme auth1
[DeviceA-aaa-authen-auth1] authentication-mode radius local
[DeviceA-aaa-authen-auth1] quit
# Configure an AAA accounting scheme named acc1 and set the accounting mode to RADIUS accounting.
[DeviceA-aaa] accounting-scheme acc1 
[DeviceA-aaa-accounting-acc1] accounting-mode radius 
[DeviceA-aaa-accounting-acc1] quit
# Apply the AAA schemes and RADIUS server template to a domain.
[DeviceA-aaa] domain huawei.com
[DeviceA-aaa-domain-huawei.com] authentication-scheme auth1
[DeviceA-aaa-domain-huawei.com] accounting-scheme acc1
[DeviceA-aaa-domain-huawei.com] radius-server 1
[DeviceA-aaa-domain-huawei.com] quit
[DeviceA-aaa] quit
[DeviceA] domain huawei.com admin
The configuration includes adding a device, adding a user, and setting the user privilege level to 3.
In the command output, the values of User access type, User Privilege, User authentication type, Current authentication method, Current authorization method, and Current accounting method indicate that the user login mode is SSH, the privilege level is 3, the authentication type is administrator authentication, and the authentication, authorization, as well as accounting modes are RADIUS.
<DeviceA> display access-user username user1-huawei detail
 ------------------------------------------------------------------------------
Basic:
  User ID                         : 16414
  User name                       : user1-huawei
  Domain-name                     : huawei.com
  User MAC                        : -
  User IP address                 : 10.1.1.10
  User IPv6 address               : -
  User access time                : 2024/07/01 10:49:22
  User accounting session ID      : example010000000000006d****010001e
  User access type                : SSH
  User Privilege                  : 3
  User Group                      : -
AAA:
  User authentication type        : Administrator authentication
  Current authentication method   : RADIUS
  Current authorization method    : RADIUS
  Current accounting method       : RADIUS
In the command output, the values of User access type, User Privilege, User authentication type, Current authentication method, Current authorization method, and Current accounting method indicate that the user login mode is SSH, the privilege level is 3, the authentication type is administrator authentication, the authentication mode is local, and the accounting mode is RADIUS.
<DeviceA> display access-user username user1-huawei detail
 ------------------------------------------------------------------------------
Basic:
  User ID                         : 16414
  User name                       : user1-huawei
  Domain-name                     : huawei.com
  User MAC                        : -
  User IP address                 : 10.1.1.10
  User IPv6 address               : -
  User access time                : 2024/07/01 10:49:22
  User accounting session ID      : example010000000000006d****010001e
  User access type                : SSH
  User Privilege                  : 3
  User Group                      : -
AAA:
  User authentication type        : Administrator authentication
  Current authentication method   : Local
  Current authorization method    : -
  Current accounting method       : RADIUS
DeviceA
#
sysname DeviceA
#
