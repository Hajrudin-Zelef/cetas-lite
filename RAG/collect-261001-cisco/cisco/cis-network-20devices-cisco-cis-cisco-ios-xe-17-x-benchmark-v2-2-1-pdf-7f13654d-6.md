---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-6
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [472, 653]
sha256: 9c5754b836e9b75883846e351027135d273d1b0feaf4c9aa7845b7279e6a193f
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 18 
Internal Only - General 
1.1.1 Enable 'aaa new-model' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
This command enables the AAA access control system. 
Rationale: 
Authentication, authorization and accounting (AAA) services provide an authoritative 
source for managing and monitoring access for devices. Centralizing control improves 
consistency of access control, the services that may be accessed once authenticated 
and accountability by tracking services accessed. Additionally, centralizing access 
control simplifies and reduces administrative costs of account provisioning and de-
provisioning, especially when managing a large number of devices. 
Impact: 
Implementing Cisco AAA is significantly disruptive as former access methods are 
immediately disabled. Therefore, before implementing Cisco AAA, the organization 
should carefully review and plan their authentication criteria (logins & passwords, 
challenges & responses, and token technologies), authorization methods, and 
accounting requirements. 
Audit: 
Perform the following to determine if AAA services are enabled: 
 hostname#show running-config | inc aaa new-model 
If the result includes a "no", the feature is not enabled. 
Remediation: 
Globally enable authentication, authorization and accounting (AAA) using the new-
model command. 
hostname(config)#aaa new-model 
Default Value: 
AAA is not enabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a2.html#GUID-
E05C2E00-C01E-4053-9D12-EC37C7E8EEC5

Page 19 
Internal Only - General 
Additional Information: 
this is a cosmetic change: adjusting regex to match 'aaa new-model' and nothing from 
line which follows next. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 5.6 Centralize Account Management 
 Centralize account management through a directory or identity service.  ● ● 
v7 
16.2 Configure Centralized Point of Authentication 
 Configure access for all accounts through as few centralized points of 
authentication as possible, including network, security, and cloud systems. 
 ● ●

Page 20 
Internal Only - General 
1.1.2 Enable 'aaa authentication login' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Sets authentication, authorization and accounting (AAA) authentication at login. 
Rationale: 
Using AAA authentication for interactive management access to the device provides 
consistent, centralized control of your network. The default under AAA (local or network) 
is to require users to log in using a valid user name and password. This rule applies for 
both local and network AAA. Fallback mode should also be enabled to allow emergency 
access to the router or switch in the event that the AAA server was unreachable, by 
utilizing the LOCAL keyword after the AAA server-tag. 
Impact: 
Implementing Cisco AAA is significantly disruptive as former access methods are 
immediately disabled. Therefore, before implementing Cisco AAA, the organization 
should carefully review and plan their authentication methods such as logins and 
passwords, challenges and responses, and which token technologies will be used. 
Audit: 
Perform the following to determine if AAA authentication for login is enabled: 
hostname#show running-config | incl aaa authentication login 
If a result does not return, the feature is not enabled. 
Remediation: 
Configure AAA authentication method(s) for login authentication. 
hostname(config)#aaa authentication login {default | aaa_list_name} [passwd-
expiry] 
[method1] [method2] 
  
Default Value: 
AAA authentication at login is disabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a1.html#GUID-
3DB1CC8A-4A98-400B-A906-C42F265C7EA2

Page 21 
Internal Only - General 
Additional Information: 
Only “the default method list is automatically applied to all interfaces except those that 
have a named method list explicitly defined. A defined method list overrides the default 
method list.” (1) 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 5.6 Centralize Account Management 
 Centralize account management through a directory or identity service.  ● ● 
v7 
16.2 Configure Centralized Point of Authentication 
 Configure access for all accounts through as few centralized points of 
authentication as possible, including network, security, and cloud systems. 
 ● ●

Page 22 
Internal Only - General 
1.1.3 Enable 'aaa authentication enable default' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Authenticates users who access privileged EXEC mode when they use the enable 
command. 
Rationale: 
Using AAA authentication for interactive management access to the device provides 
consistent, centralized control of your network. The default under AAA (local or network) 
is to require users to log in using a valid user name and password. This rule applies for 
both local and network AAA. 
Impact: 
Enabling Cisco AAA 'authentication enable' mode is significantly disruptive as former 
access methods are immediately disabled. Therefore, before enabling 'aaa 
authentication enable default' mode, the organization should plan and implement 
authentication logins and passwords, challenges and responses, and token 
technologies. 
Audit: 
Perform the following to determine if AAA authentication enable mode is enabled: 
hostname#show running-config | incl aaa authentication enable 
If a result does not return, the feature is not enabled 
Remediation: 
Configure AAA authentication method(s) for enable authentication. 
hostname(config)#aaa authentication enable default {method1} enable  
Default Value: 
By default, fallback to the local database is disabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a1.html#GUID-
4171D649-2973-4707-95F3-9D96971893D0

Page 23 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 5.6 Centralize Account Management 
 Centralize account management through a directory or identity service.  ● ● 
v7 
16.2 Configure Centralized Point of Authentication 
 Configure access for all accounts through as few centralized points of 
authentication as possible, including network, security, and cloud systems. 
 ● ●

Page 24 
Internal Only - General 
1.1.4 Set 'login authentication for 'line vty' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Authenticates users who access the router or switch remotely through the VTY port. 
Rationale: 
Using AAA authentication for interactive management access to the device provides 
consistent, centralized control of your network. The default under AAA (local or network) 
is to require users to log in using a valid user name and password. This rule applies for 
both local and network AAA. 
Impact: 
Enabling Cisco AAA 'login authentication for line VTY' is significantly disruptive as 
former access methods are immediately disabled. Therefore, before enabling Cisco 
AAA 'login authentication for line VTY', the organization should plan and implement 
authentication logins and passwords, challenges and responses, and token 
technologies. 
Audit: 
Perform the following to determine if AAA authentication for line login is enabled: 
If the command does not return a result for each management access method, the 
feature is not enabled 
hostname#show running-config | sec line | incl login authentication 
Remediation: 
Configure management lines to require login using the default or a named AAA 
authentication list. This configuration must be set individually for all line types. 
hostname(config)#line vty {line-number} [<em>ending-line-number] 
hostname(config-line)#login authentication {default | aaa_list_name} 
Default Value: 
Login authentication is not enabled. 
Uses the default set with aaa authentication login. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-k1.html#GUID-
297BDF33-4841-441C-83F3-4DA51C3C7284

