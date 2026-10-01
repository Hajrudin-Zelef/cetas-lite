---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-7
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [654, 825]
sha256: 8a8cf5aea9dff2b99fd9967991b38cd962f8c2d240a6166e7c6d65cb37e9fefc
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 25 
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

Page 26 
Internal Only - General 
1.1.5 Set 'login authentication for 'ip http' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
If account management functions are not automatically enforced, an attacker could gain 
privileged access to a vital element of the network security architecture 
Rationale: 
Configure the device to enforce authentication for the HTTP server (ip http) by defining 
a valid login method. This ensures that only authorized users can access the device's 
web-based management interface. Use AAA (Authentication, Authorization, and 
Accounting) or a predefined local username and password database to enforce login 
authentication 
Impact: 
Enabling Cisco AAA 'line login' is significantly disruptive as former access methods are 
immediately disabled. Therefore, before enabling Cisco AAA 'line login', the 
organization should plan and implement authentication logins and passwords, 
challenges and responses, and token technologies. 
Audit: 
Perform the following to determine if AAA authentication for line login is enabled: 
If the command does not return a result for each management access method, the 
feature is not enabled 
hostname#show running-config | inc ip http authentication 
Remediation: 
Configure management lines to require login using the default or a named AAA 
authentication list. This configuration must be set individually for all line types. 
hostname#(config)ip http secure-server 
hostname#(config)ip http authentication {default | _aaa\_list\_name_} 
Default Value: 
Login authentication is not enabled. 
Uses the default set with aaa authentication login.

Page 27 
Internal Only - General 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-k1.html#GUID-
297BDF33-4841-441C-83F3-4DA51C3C7284 
Additional Information: 
This control is usefull if http or https server is enabled, which involves global config 
contains a line starting with "ip http server" or "ip http secure-server". Then, look for ('no 
ip http server' and 'no ip http secure-server') or 'ip http authentication ' config lines 
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

Page 28 
Internal Only - General 
1.1.6 Set 'aaa accounting' to log all privileged use commands 
using 'commands 15' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Runs accounting for all commands at the specified privilege level. 
Rationale: 
Authentication, authorization and accounting (AAA) systems provide an authoritative 
source for managing and monitoring access for devices. Centralizing control improves 
consistency of access control, the services that may be accessed once authenticated 
and accountability by tracking services accessed. Additionally, centralizing access 
control simplifies and reduces administrative costs of account provisioning and de-
provisioning, especially when managing a large number of devices. AAA Accounting 
provides a management and audit trail for user and administrative sessions through 
TACACS+. 
Impact: 
Enabling 'aaa accounting' for privileged commands records and sends activity to the 
accounting servers and enables organizations to monitor and analyze privileged activity. 
Audit: 
Perform the following to determine if aaa accounting for commands is required: 
Verify a command string result returns 
hostname#show running-config | incl aaa accounting commands 
Remediation: 
Configure AAA accounting for commands. 
hostname(config)#aaa accounting commands 15 {default | list-name | guarantee-
first} 
{start-stop | stop-only | none} {radius | group group-name} 
Default Value: 
AAA accounting is disabled. 
Additional Information: 
Valid privilege level entries are integers from 0 through 15.

Page 29 
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

Page 30 
Internal Only - General 
1.1.7 Set 'aaa accounting connection' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Provides information about all outbound connections made from the network access 
server. 
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
Implementing aaa accounting connection creates accounting records about connections 
from the network access server. Organizations should regular monitor these connection 
records for exceptions, remediate issues, and report findings regularly. 
Audit: 
Perform the following to determine if aaa accounting for connection is required: 
Verify a command string result returns 
hostname#show running-config | incl aaa accounting connection 
Remediation: 
Configure AAA accounting for connections. 
hostname(config)#aaa accounting connection {default | list-name | guarantee-
first}  
{start-stop | stop-only | none} {radius | group group-name} 
Default Value: 
AAA accounting is not enabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a1.html#GUID-
0520BCEF-89FB-4505-A5DF-D7F1389F1BBA

Page 31 
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

