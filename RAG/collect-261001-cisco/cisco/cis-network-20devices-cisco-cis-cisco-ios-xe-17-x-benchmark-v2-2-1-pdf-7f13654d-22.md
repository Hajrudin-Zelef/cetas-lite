---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-22
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [3526, 3751]
sha256: f07a6fba5f5791889df587801d68149dd4c3fc4315833778ff1fc4b28d232016
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 133 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
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

Page 134 
Internal Only - General 
2.2.7 Set 'logging source interface' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Specify the source IPv4 or IPv6 address of system logging packets 
Rationale: 
This is required so that the router sends log messages to the logging server from a 
consistent IP address. 
Impact: 
Logging is an important process for an organization managing technology risk and 
establishing a consistent source of messages for the logging host is critical. The 'logging 
source interface loopback' command sets a consistent IP address to send messages to 
the logging host and enforces the logging process. 
Audit: 
Perform the following to determine if logging services are bound to a source interface: 
Verify a command string result returns 
 
hostname#sh run | incl logging source 
Remediation: 
Bind logging to the loopback interface. 
 
hostname(config)#logging source-interface loopback 
{<em>loopback_interface_number</em>} 
Default Value: 
The wildcard interface address is used. 
References: 
1. http://www.cisco.com/en/US/docs/ios/netmgmt/command/reference/nm_09.html#
wp1095099

Page 135 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
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

Page 136 
Internal Only - General 
2.2.8 Set 'login success/failure logging' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Without generating audit records that are specific to the security and mission needs of 
the organization, it would be difficult to establish, correlate, and investigate the events 
relating to an incident or identify those responsible for one. 
Rationale: 
Audit records can be generated from various components within the information system 
(e.g., module or policy filter). 
Audit: 
hostname(config)#sho running-config | inc login on- 
Remediation: 
hostname(config)#login on-failure log 
hostname(config)#login on-success log 
hostname(config)#end 
References: 
1. https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/config-mgmt/configuration/xe-
16-6/config-mgmt-xe-16-6-book/cm-config-logger.pdf 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
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

Page 137 
Internal Only - General 
2.3 NTP Rules 
Network Time Protocol allows administrators to set the system time on all of their 
compatible systems from a single source, ensuring a consistent time stamp for logging 
and authentication protocols. NTP is an internet standard, defined in RFC1305.

Page 138 
Internal Only - General 
2.3.1 Require Encryption Keys for NTP 
Encryption keys should be set for NTP Servers.

Page 139 
Internal Only - General 
2.3.1.1 Set 'ntp authenticate' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Enable NTP authentication. 
Rationale: 
Using authenticated NTP ensures the Cisco device only permits time updates from 
authorized NTP servers. 
Impact: 
Organizations should establish three Network Time Protocol (NTP) hosts to set 
consistent time across the enterprise. Enabling the 'ntp authenticate' command enforces 
authentication between NTP hosts. 
Audit: 
From the command prompt, execute the following commands: 
hostname#show run | include ntp  
Remediation: 
Configure NTP authentication: 
 
hostname(config)#ntp authenticate 
Default Value: 
NTP authentication is not enabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/bsm/command/bsm-cr-
n1.html#GUID-8BEBDAF4-6D03-4C3E-B8D6-6BCBC7D0F324 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.4 Standardize Time Synchronization 
 Standardize time synchronization. Configure at least two synchronized time 
sources across enterprise assets, where supported. 
 ● ●

Page 140 
Internal Only - General 
Controls 
Version Control IG 1 IG 2 IG 3 
v7 
6.1 Utilize Three Synchronized Time Sources 
 Use at least three synchronized time sources from which all servers and 
network devices retrieve time information on a regular basis so that timestamps 
in logs are consistent. 
 ● ●

Page 141 
Internal Only - General 
2.3.1.2 Set 'ntp authentication-key' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Define an authentication key for Network Time Protocol (NTP). 
Rationale: 
Using an authentication key provides a higher degree of security as only authenticated 
NTP servers will be able to update time for the Cisco device. 
Impact: 
Organizations should establish three Network Time Protocol (NTP) hosts to set 
consistent time across the enterprise. Enabling the 'ntp authentication-key' command 
enforces encrypted authentication between NTP hosts. 
Audit: 
From the command prompt, execute the following commands: 
 
hostname#show run | include ntp authentication-key 
Remediation: 
Configure at the NTP key ring and encryption key using the following command 
 
hostname(config)#ntp authentication-key {ntp_key_id} md5 {ntp_key_hash} 
Default Value: 
No authentication key is defined for NTP. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/bsm/command/bsm-cr-
n1.html#GUID-0435BFD1-D7D7-41D4-97AC-7731C11226BC

Page 142 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.4 Standardize Time Synchronization 
 Standardize time synchronization. Configure at least two synchronized time 
sources across enterprise assets, where supported. 
 ● ● 
v7 
6.1 Utilize Three Synchronized Time Sources 
 Use at least three synchronized time sources from which all servers and 
network devices retrieve time information on a regular basis so that timestamps 
in logs are consistent. 
 ● ●

Page 143 
Internal Only - General 
2.3.1.3 Set the 'ntp trusted-key' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Ensure you authenticate the identity of a system to which Network Time Protocol (NTP) 
will synchronize 
Rationale: 
This authentication function provides protection against accidentally synchronizing the 
system to another system that is not trusted, because the other system must know the 
correct authentication key. 
Impact: 
Organizations should establish three Network Time Protocol (NTP) hosts to set 
consistent time across the enterprise. Enabling the 'ntp trusted-key' command enforces 
encrypted authentication between NTP hosts. 
Audit: 
From the command prompt, execute the following commands: 
 
