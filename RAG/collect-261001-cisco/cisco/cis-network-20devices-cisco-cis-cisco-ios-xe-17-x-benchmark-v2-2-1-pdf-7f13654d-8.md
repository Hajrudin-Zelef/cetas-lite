---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-8
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [826, 1008]
sha256: a8535e764ae4d64dea254ab4ac91a868ef13b92e7b601d814d129a25003171e0
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 32 
Internal Only - General 
1.1.8 Set 'aaa accounting exec' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Runs accounting for the EXEC shell session. 
Rationale: 
Authentication, authorization and accounting (AAA) systems provide an authoritative 
source for managing and monitoring access for devices. Centralizing control improves 
consistency of access control, the services that may be accessed once authenticated 
and accountability by tracking services accessed. Additionally, centralizing access 
control simplifies and reduces administrative costs of account provisioning and de-
provisioning, especially when managing a large number of devices. AAA Accounting 
provides a management and audit trail for user and administrative sessions through 
RADIUS and TACACS+. 
Impact: 
Enabling aaa accounting exec creates accounting records for the EXEC terminal 
sessions on the network access server. These records include start and stop times, 
usernames, and date information. Organizations should regularly monitor these records 
for exceptions, remediate issues, and report findings. 
Audit: 
Perform the following to determine if aaa accounting for EXEC shell session is required: 
Verify a command string result returns 
hostname#show running-config | incl aaa accounting exec 
Remediation: 
Configure AAA accounting for EXEC shell session. 
hostname(config)#aaa accounting exec {default | list-name | guarantee-first}  
{start-stop | stop-only | none} {radius | group group-name} 
Default Value: 
AAA accounting is not enabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a1.html#GUID-
0520BCEF-89FB-4505-A5DF-D7F1389F1BBA

Page 33 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.2 Collect Audit Logs 
 Collect audit logs. Ensure that logging, per the enterprise’s audit log 
management process, has been enabled across enterprise assets. 
● ● ● 
v8 
8.5 Collect Detailed Audit Logs 
 Configure detailed audit logging for enterprise assets containing sensitive data. 
Include event source, date, username, timestamp, source addresses, destination 
addresses, and other useful elements that could assist in a forensic investigation. 
 ● ● 
v7 
6.3 Enable Detailed Logging 
 Enable system logging to include detailed information such as an event source, 
date, user, timestamp, source addresses, destination addresses, and other useful 
elements. 
 ● ●

Page 34 
Internal Only - General 
1.1.9 Set 'aaa accounting network' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Runs accounting for all network-related service requests. 
Rationale: 
Authentication, authorization and accounting (AAA) systems provide an authoritative 
source for managing and monitoring access for devices. Centralizing control improves 
consistency of access control, the services that may be accessed once authenticated 
and accountability by tracking services accessed. Additionally, centralizing access 
control simplifies and reduces administrative costs of account provisioning and de-
provisioning, especially when managing a large number of devices. AAA Accounting 
provides a management and audit trail for user and administrative sessions through 
RADIUS and TACACS+. 
Impact: 
Implementing aaa accounting network creates accounting records for a method list 
including ARA, PPP, SLIP, and NCPs sessions. Organizations should regular monitor 
these records for exceptions, remediate issues, and report findings. 
Audit: 
Perform the following to determine if aaa accounting for connection is required: 
Verify a command string result returns 
hostname#show running-config | incl aaa accounting network 
Remediation: 
Configure AAA accounting for connections. 
hostname(config)#aaa accounting network {default | list-name | guarantee-
first}  
{start-stop | stop-only | none} {radius | group group-name} 
Default Value: 
AAA accounting is not enabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a1.html#GUID-
0520BCEF-89FB-4505-A5DF-D7F1389F1BBA

Page 35 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.2 Collect Audit Logs 
 Collect audit logs. Ensure that logging, per the enterprise’s audit log 
management process, has been enabled across enterprise assets. 
● ● ● 
v8 
8.5 Collect Detailed Audit Logs 
 Configure detailed audit logging for enterprise assets containing sensitive data. 
Include event source, date, username, timestamp, source addresses, destination 
addresses, and other useful elements that could assist in a forensic investigation. 
 ● ● 
v7 
6.3 Enable Detailed Logging 
 Enable system logging to include detailed information such as an event source, 
date, user, timestamp, source addresses, destination addresses, and other useful 
elements. 
 ● ●

Page 36 
Internal Only - General 
1.1.10 Set 'aaa accounting system' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Performs accounting for all system-level events not associated with users, such as 
reloads. 
Rationale: 
Authentication, authorization and accounting (AAA) systems provide an authoritative 
source for managing and monitoring access for devices. Centralizing control improves 
consistency of access control, the services that may be accessed once authenticated 
and accountability by tracking services accessed. Additionally, centralizing access 
control simplifies and reduces administrative costs of account provisioning and de-
provisioning, especially when managing a large number of devices. AAA Accounting 
provides a management and audit trail for user and administrative sessions through 
RADIUS and TACACS+. 
Impact: 
Enabling aaa accounting system creates accounting records for all system-level events. 
Organizations should regular monitor these records for exceptions, remediate issues, 
and report findings regularly. 
Audit: 
Perform the following to determine if aaa accounting system is required: 
Verify a command string result returns 
hostname#show running-config | incl aaa accounting system 
Remediation: 
Configure AAA accounting system. 
hostname(config)#aaa accounting system {default | list-name | guarantee-
first}  
{start-stop | stop-only | none} {radius | group group-name} 
Default Value: 
AAA accounting is not enabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a1.html#GUID-
0520BCEF-89FB-4505-A5DF-D7F1389F1BBA

Page 37 
Internal Only - General 
Additional Information: 
When system accounting is used and the accounting server is unreachable at system 
startup time, the system will not be accessible for approximately two minutes. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.2 Collect Audit Logs 
 Collect audit logs. Ensure that logging, per the enterprise’s audit log 
management process, has been enabled across enterprise assets. 
● ● ● 
v8 
8.5 Collect Detailed Audit Logs 
 Configure detailed audit logging for enterprise assets containing sensitive data. 
Include event source, date, username, timestamp, source addresses, destination 
addresses, and other useful elements that could assist in a forensic investigation. 
 ● ● 
v7 
6.3 Enable Detailed Logging 
 Enable system logging to include detailed information such as an event source, 
date, user, timestamp, source addresses, destination addresses, and other useful 
elements. 
 ● ●

Page 38 
Internal Only - General 
1.2 Access Rules 
Rules in the access class enforce controls for device administrative connections.

