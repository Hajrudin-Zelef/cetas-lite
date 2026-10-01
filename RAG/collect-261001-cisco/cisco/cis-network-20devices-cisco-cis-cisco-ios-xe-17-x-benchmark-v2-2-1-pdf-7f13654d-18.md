---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-18
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [2645, 2867]
sha256: 37d32f142cfd8d938d340163dc4b72409bfb8ed25ffb84b0e7ccb9d6b74587c0
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 99 
Internal Only - General 
2.1.1.1.2 Set the 'ip domain-name' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Define a default domain name that the Cisco IOS software uses to complete unqualified 
hostnames 
Rationale: 
The domain name is a prerequisite for setting up SSH. 
Impact: 
Organizations should plan the enterprise network and identify an appropriate domain 
name for the router. 
Audit: 
Perform the following to determine if the domain name is configured: 
Verify the domain name is configured properly. 
 
hostname#sh run | incl domain-name 
Remediation: 
Configure an appropriate domain name for the router. 
 
hostname (config)#ip domain-name {<em>domain-name</em>} 
Default Value: 
No domain is set. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/ipaddr/command/ipaddr-
i3.html#GUID-A706D62B-9170-45CE-A2C2-7B2052BE2CAB 
CIS Controls: 
Controls Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 100 
Internal Only - General 
Controls Version Control IG 1 IG 2 IG 3 
v7 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 101 
Internal Only - General 
2.1.1.1.3 Set 'modulus' to greater than or equal to 2048 for 'crypto 
key generate rsa' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Use this command to generate RSA key pairs for your Cisco device. 
RSA keys are generated in pairs--one public RSA key and one private RSA key. 
Rationale: 
An RSA key pair is a prerequisite for setting up SSH and should be at least 2048 bits. 
NOTE: IOS does NOT display the modulus bit value in the Audit Procedure. 
Impact: 
Organizations should plan and implement enterprise network cryptography and 
generate an appropriate RSA key pairs, such as 'modulus', greater than or equal to 
2048. 
Audit: 
Perform the following to determine if the RSA key pair is configured: 
 
 hostname#sh crypto key mypubkey rsa 
Remediation: 
Generate an RSA key pair for the router. 
 
hostname(config)#crypto key generate rsa general-keys modulus <em>2048</em> 
Default Value: 
RSA key pairs do not exist. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-c4.html#GUID-
2AECF701-D54A-404E-9614-D3AAB049BC13

Page 102 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
16.11 Leverage Vetted Modules or Services for Application 
Security Components 
 Leverage vetted modules or services for application security components, such 
as identity management, encryption, and auditing and logging. Using platform 
features in critical security functions will reduce developers’ workload and minimize 
the likelihood of design or implementation errors. Modern operating systems provide 
effective mechanisms for identification, authentication, and authorization and make 
those mechanisms available to applications. Use only standardized, currently 
accepted, and extensively reviewed encryption algorithms. Operating systems also 
provide mechanisms to create and maintain secure audit logs. 
 ● ● 
v7 
18.5 Use Only Standardized and Extensively Reviewed 
Encryption Algorithms 
 Use only standardized and extensively reviewed encryption algorithms. 
 ● ●

Page 103 
Internal Only - General 
2.1.1.1.4 Set 'seconds' for 'ip ssh timeout' for 60 seconds or less 
(Automated) 
Profile Applicability: 
•  Level 1 
Description: 
The time interval that the router waits for the SSH client to respond before 
disconnecting an uncompleted login attempt. 
Rationale: 
This reduces the risk of an administrator leaving an authenticated session logged in for 
an extended period of time. 
Impact: 
Organizations should implement a security policy requiring minimum timeout settings for 
all network administrators and enforce the policy through the 'ip ssh timeout' command. 
Audit: 
Perform the following to determine if the SSH timeout is configured: 
Verify the timeout is configured properly. 
 
hostname#sh ip ssh 
Remediation: 
Configure the SSH timeout 
 
hostname(config)#ip ssh time-out [<em>60</em>]  
Default Value: 
SSH in not enabled by default. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i3.html#GUID-
5BAC7A2B-0A25-400F-AEE9-C22AE08513C6 
Additional Information: 
This cannot exceed 120 seconds. 
Adjusting Artifact's title and regex from 'timeout' to 'time-out'

Page 104 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.3 Configure Automatic Session Locking on Enterprise 
Assets 
 Configure automatic session locking on enterprise assets after a defined period 
of inactivity. For general purpose operating systems, the period must not exceed 
15 minutes. For mobile end-user devices, the period must not exceed 2 minutes. 
● ● ● 
v7 16.11 Lock Workstation Sessions After Inactivity 
 Automatically lock workstation sessions after a standard period of inactivity. ● ● ●

Page 105 
Internal Only - General 
2.1.1.1.5 Set maximum value for 'ip ssh authentication-retries' 
(Automated) 
Profile Applicability: 
•  Level 1 
Description: 
The number of retries before the SSH login session disconnects. 
Rationale: 
This limits the number of times an unauthorized user can attempt a password without 
having to establish a new SSH login attempt. This reduces the potential for success 
during online brute force attacks by limiting the number of login attempts per SSH 
connection. 
Impact: 
Organizations should implement a security policy limiting the number of authentication 
attempts for network administrators and enforce the policy through the 'ip ssh 
authentication-retries' command. 
Audit: 
Perform the following to determine if SSH authentication retries is configured: 
Verify the authentication retries is configured properly. 
 
hostname#sh ip ssh 
Remediation: 
Configure the SSH timeout: 3 or less 
 
hostname(config)#ip ssh authentication-retries [<em>3</em>] 
Default Value: 
SSH is not enabled by default. When set, the default value is 3. When set using the 
default value it will not display under a show running-configuration. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i3.html#GUID-
5BAC7A2B-0A25-400F-AEE9-C22AE08513C6

Page 106 
Internal Only - General 
CIS Controls: 
Controls Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v7 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 107 
Internal Only - General 
2.1.1.2 Set version 2 for 'ip ssh version' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Specify the version of Secure Shell (SSH) to be run on a router 
Rationale: 
SSH Version 1 has been subject to a number of serious vulnerabilities and is no longer 
considered to be a secure protocol, resulting in the adoption of SSH Version 2 as an 
Internet Standard in 2006. 
Cisco routers support both versions, but due to the weakness of SSH Version 1 only the 
later standard should be used. 
Impact: 
To reduce the risk of unauthorized access, organizations should implement a security 
policy to review their current protocols to ensure the most secure protocol versions are 
in use. 
Audit: 
Perform the following to determine if SSH version 2 is configured: 
Verify that SSH version 2 is configured properly. 
 
hostname#sh ip ssh 
Remediation: 
Configure the router to use SSH version 2 
 
hostname(config)#ip ssh version 2 
Default Value: 
SSH is not enabled by default. When enabled, SSH operates in compatibility mode 
(versions 1 and 2 supported). 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i3.html#GUID-
170AECF1-4B5B-462A-8CC8-999DEDC45C21

Page 108 
Internal Only - General 
CIS Controls: 
Controls Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v7 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

