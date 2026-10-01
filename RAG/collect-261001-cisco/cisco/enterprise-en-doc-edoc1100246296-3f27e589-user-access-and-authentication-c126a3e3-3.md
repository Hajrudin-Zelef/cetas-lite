---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3-3
title: "Configure DeviceA to generate a local key pair."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100246296-3f27e589-user-access-and-authentication-c126a3e3.md
source_anchor: ""
source_lines: [329, 512]
sha256: 89847fd71001be809b5c400c54bab284986b028de97b709c58a05ce5d9a29291
---

# Configure DeviceA to generate a local key pair.

[DeviceA-aaa-recording-sch0] quit 
[DeviceA-aaa] cmd recording-scheme sch0
The configuration includes the following: add a device, add a user, set the user privilege level to 3, and configure command authorization. Note that the reset hwtacacs-server statistics all command cannot be configured.
You can check logs recording command execution successes and failures of all users including non-HWTACACS authentication users under Reports and Activity > TACACS+ Administration.
[DeviceA] quit
<DeviceA> reset hwtacacs-server statistics all
 Error: Failed to pass the authorization.
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
  authorization-cmd 3 hwtacacs local
 accounting-scheme sch3                                                          
  accounting-mode hwtacacs  
 recording-scheme sch0                                                           
  recording-mode hwtacacs template1                                                    
 cmd recording-scheme sch0 
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
In Figure 3-107, an LDAP server is deployed on the network, and user information as well as group information are stored on the LDAP server. The enterprise requires that the administrator can log in to the device through STelnet only after passing LDAP authentication.
# Configure an LDAP server template, and configure the IP address, port number, Base DN, administrator DN, and administrator password in the template.
[DeviceA] ldap-server template template1
[DeviceA-ldap-template1] ldap-server authentication 10.1.6.6 389 no-ssl
[DeviceA-ldap-template1] ldap-server authentication base-dn dc=esaptest,dc=com
[DeviceA-ldap-template1] ldap-server authentication manager cn=Administrator,cn=users Admin@123
[DeviceA-ldap-template1] quit
# Configure an authentication scheme named auth1 and set the authentication mode to LDAP authentication.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme auth1
[DeviceA-aaa-authen-auth1] authentication-mode ldap
[DeviceA-aaa-authen-auth1] quit
# Configure an authorization scheme named acc1 and set the authorization mode to LDAP authorization.
[DeviceA-aaa] authorization-scheme acc1
[DeviceA-aaa-author-acc1] authorization-mode ldap
[DeviceA-aaa-author-acc1] quit
# Configure the service scheme s1 and set the administrator privilege level to 3 in the service scheme.
[DeviceA-aaa] service-scheme s1
[DeviceA-aaa-service-s1] admin-user privilege level 3
[DeviceA-aaa-service-s1] quit
# Apply the created AAA schemes and LDAP server template to a specific domain.
[DeviceA-aaa] domain huawei.com
[DeviceA-aaa-domain-huawei.com] authentication-scheme auth1
[DeviceA-aaa-domain-huawei.com] authorization-scheme acc1
[DeviceA-aaa-domain-huawei.com] ldap-server template1
[DeviceA-aaa-domain-huawei.com] service-scheme s1
[DeviceA-aaa-domain-huawei.com] quit
[DeviceA-aaa] quit
#
sysname DeviceA
#
ldap-server template template1
 ldap-server authentication 10.1.6.6 389 no-ssl
 ldap-server authentication manager cn=Administrator,cn=users %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.nt[%K)u50(_zvw*/=MzISB3'K!!!!!2jp5!!!!!!:!!!!h$%%+PY@9GDRjBP8bVj/vDf{Cr}&[Sc`Kf)!!!!!%+%#
 ldap-server authentication base-dn dc=esaptest,dc=com
#
aaa
 authentication-scheme auth1    
  authentication-mode ldap  
 authorization-scheme acc1    
  authorization-mode ldap  
 service-scheme s1
  admin-user privilege level 3 
 domain huawei.com            
  authentication-scheme auth1     
  authorization-scheme acc1
  service-scheme s1 
  ldap-server template1 
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
In Figure 3-107, an AD server is deployed on the network, and user information as well as group information are stored on the AD server. The enterprise requires that the administrator can log in to the device through STelnet only after passing AD authentication.
In this example, interface 1 and interface 2 represent 10GE 1/0/1 and 10GE 1/0/2, respectively.
# Configure an AD server template, and configure the IP address, port number, Base DN, administrator DN, and administrator password in the template.
[DeviceA] ad-server template template1
[DeviceA-ad-template1] ad-server authentication 10.1.6.6 88
[DeviceA-ad-template1] ad-server authentication base-dn dc=test1,dc=com
[DeviceA-ad-template1] ad-server authentication manager cn=Administrator,cn=users Admin@123
[DeviceA-ad-template1] ad-server authentication host-name win.aa
[DeviceA-ad-template1] quit
# Configure an authentication scheme named auth1 and set the authentication mode to AD authentication.
[DeviceA] aaa
[DeviceA-aaa] authentication-scheme auth1
[DeviceA-aaa-authen-auth1] authentication-mode ad
[DeviceA-aaa-authen-auth1] quit
# Configure an authorization scheme named acc1 and set the authorization mode to local authorization.
[DeviceA-aaa] authorization-scheme acc1
[DeviceA-aaa-author-acc1] authorization-mode local
[DeviceA-aaa-author-acc1] quit
# Configure a local user and access type.
# Configure local authorization. Select one or more configuration methods as required.
[DeviceA-aaa] domain huawei.com
[DeviceA-aaa-domain-huawei.com] authentication-scheme auth1
[DeviceA-aaa-domain-huawei.com] authorization-scheme acc1
[DeviceA-aaa-domain-huawei.com] ad-server template1
[DeviceA-aaa-domain-huawei.com] service-scheme s1
[DeviceA-aaa-domain-huawei.com] quit
[DeviceA-aaa] quit
#
sysname DeviceA
#
ad-server template template1
 ad-server authentication 10.1.6.6 88 ldap-over-ssl
 ad-server authentication base-dn dc=test1,dc=com
 ad-server authentication manager cn=Administrator,cn=users %+%##!!!!!!!!!"!!!!"!!!!*!!!!SKvr${[Fs.|#*S#kB\=3@-+#@CYd5E^@|E*!!!!!2jp5!!!!!!:!!!!F4_]L0O\."9h_o8tIO/(`'>p";xS#/Fns\*!!!!!%+%#
 ad-server authentication host-name win.aa
#
aaa
 authentication-scheme auth1    
  authentication-mode ad  
 authorization-scheme acc1    
  authorization-mode local  
 service-scheme s1
  admin-user privilege level 3 
 domain huawei.com            
  authentication-scheme auth1     
  authorization-scheme acc1
  service-scheme s1
