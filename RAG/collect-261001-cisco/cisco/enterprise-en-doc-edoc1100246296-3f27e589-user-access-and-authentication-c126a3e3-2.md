---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-2
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [156, 328]
sha256: 6192b600a72483ebc053819283291b2a2a91ef58c0353a705f5993394a649f4a
---

# Configure DeviceA to generate a local key pair.

radius-server template 1                                                        
 radius-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.3t@/5k|BENhEu>W(3\~XG!!D;!!!!!2jp5!!!!!!A!!!!3"pK8qv!}9M#(4$jGWvQF/R[CNe/+:W^jk8HUe&W%+%#
 radius-server authentication 10.1.6.6 1812 weight 80                           
 radius-server accounting 10.1.6.6 1813 weight 80
#
aaa
 authentication-scheme auth1    
  authentication-mode radius local 
 accounting-scheme acc1    
  accounting-mode radius  
 domain huawei.com            
  authentication-scheme auth1     
  accounting-scheme acc1
  radius-server 1      
 local-user user1-huawei password irreversible-cipher $1d$OwseVRh@LH}ZeTBm$1nH4$ab>d(N{-%0!ab48y=Ic*xEUR4pVhR2"9-~,$
 local-user user1-huawei privilege level 3
 local-user user1-huawei service-type ssh   
#  
domain huawei.com admin 
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
stelnet server enable 
ssh server-source -i Vlanif10
#
user-interface vty 0 4           
 authentication-mode aaa           
 protocol inbound ssh
# 
return 
# Create an HWTACACS server template named template1 for communication between DeviceA and the HWTACACS server.
[DeviceA] hwtacacs-server template template1 
[DeviceA-hwtacacs-template1] hwtacacs-server authentication 10.1.6.6 49 
[DeviceA-hwtacacs-template1] hwtacacs-server authorization 10.1.6.6 49 
[DeviceA-hwtacacs-template1] hwtacacs-server accounting 10.1.6.6 49 
[DeviceA-hwtacacs-template1] hwtacacs-server shared-key cipher YsHsjx_202206139 
[DeviceA-hwtacacs-template1] quit
# Create an authentication scheme named sch1 and set the authentication mode to HWTACACS and local authentication.
[DeviceA] aaa 
[DeviceA-aaa] authentication-scheme sch1
[DeviceA-aaa-authen-sch1] authentication-mode hwtacacs local
[DeviceA-aaa-authen-sch1] quit
# Configure an authorization scheme named sch2 and set the authorization mode to HWTACACS and local authorization.
[DeviceA-aaa] authorization-scheme sch2 
[DeviceA-aaa-author-sch2] authorization-mode hwtacacs local
[DeviceA-aaa-author-sch2] quit
# Create an accounting scheme named sch3 and set the accounting mode to HWTACACS accounting.
[DeviceA-aaa] accounting-scheme sch3 
[DeviceA-aaa-accounting-sch3] accounting-mode hwtacacs 
[DeviceA-aaa-accounting-sch3] quit
# Apply the HWTACACS server template and AAA schemes to the domain huawei.com.
[DeviceA-aaa] domain huawei.com 
[DeviceA-aaa-domain-huawei.com] hwtacacs-server template1 
[DeviceA-aaa-domain-huawei.com] authentication-scheme sch1 
[DeviceA-aaa-domain-huawei.com] authorization-scheme sch2 
[DeviceA-aaa-domain-huawei.com] accounting-scheme sch3 
[DeviceA-aaa-domain-huawei.com] quit 
[DeviceA-aaa] quit
# Set the domain huawei.com as the global default administrative domain.
In the command output, the values of User access type, User Privilege, User authentication type, Current authentication method, Current authorization method, and Current accounting method indicate that the user login mode is SSH, the privilege level is 3, the authentication type is administrator authentication, and the authentication, authorization, as well as accounting modes are HWTACACS.
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
  Current authentication method   : HWTACACS
  Current authorization method    : HWTACACS
  Current accounting method       : HWTACACS
In the command output, the values of User access type, User Privilege, User authentication type, Current authentication method, Current authorization method, and Current accounting method indicate that the login mode is SSH, the privilege level is 3, the authentication type is administrator authentication, the authentication and authorization modes are local, and the accounting mode is HWTACACS.
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
  Current authorization method    : Local
  Current accounting method       : HWTACACS
# 
sysname DeviceA 
# 
hwtacacs-server template template1                                           
 hwtacacs-server authentication 10.1.6.6                                    
 hwtacacs-server authorization 10.1.6.6                                   
 hwtacacs-server accounting 10.1.6.6  
 hwtacacs-server shared-key cipher %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs."u,S-6a-X1'[X=L"cpF!5Oz`1!!!!!2jp5!!!!!!A!!!!Ix>cM8i{y6!);(8Dr9:dK`&BHfE(H2=.:SH{@pT%+%#  
# 
aaa 
 authentication-scheme sch1     
  authentication-mode hwtacacs local
 authorization-scheme sch2                                                       
  authorization-mode hwtacacs local
 accounting-scheme sch3                                                          
  accounting-mode hwtacacs  
 domain huawei.com                                                               
  authentication-scheme sch1                                                     
  accounting-scheme sch3                                                         
  authorization-scheme sch2                                                      
  hwtacacs-server template1  
 local-user user1-huawei password irreversible-cipher $1d$OwseVRh@LH}ZeTBm$1nH4$ab>d(N{-%0!ab48y=Ic*xEUR4pVhR2"9-~,$
 local-user user1-huawei privilege level 3
 local-user user1-huawei service-type ssh   
#  
domain huawei.com admin 
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
stelnet server enable 
ssh server-source -i Vlanif10
# 
user-interface vty 0 4           
 authentication-mode aaa           
 protocol inbound ssh
# 
return
# Configure an authentication scheme named sch1 and set the authentication mode to HWTACACS and local authentication.
# Create an authorization scheme named sch2, set the authorization mode to HWTACACS authorization, and enable command authorization for the administrator with the privilege level 3.
[DeviceA-aaa] authorization-scheme sch2 
[DeviceA-aaa-author-sch2] authorization-mode hwtacacs local
[DeviceA-aaa-author-sch2] authorization-cmd 3 hwtacacs local
[DeviceA-aaa-author-sch2] quit
# Create a recording scheme named sch0 to record the commands that the administrator has executed.
[DeviceA-aaa] recording-scheme sch0 
[DeviceA-aaa-recording-sch0] recording-mode hwtacacs template1 
