---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-15
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [2013, 2234]
sha256: 64b4a12cac0678ecd0656c789686c393249fbe3d6dd99c12774d1de0b4758a0b
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 74 
Internal Only - General 
References: 
1. https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst9600/software/releas
e/16-
12/configuration_guide/sec/b_1612_sec_9600_cg/controlling_switch_access_wit
h_passwords_and_privilege_levels.html 
Additional Information: 
if any local user is defined ensure that "secret 9" is used for each of them (Artifact #1 
"check all username are using secret 9 only"), so that reverting the hash to the 
password is difficult. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
3.11 Encrypt Sensitive Data at Rest 
 Encrypt sensitive data at rest on servers, applications, and databases containing 
sensitive data. Storage-layer encryption, also known as server-side encryption, 
meets the minimum requirement of this Safeguard. Additional encryption methods 
may include application-layer encryption, also known as client-side encryption, 
where access to the data storage device(s) does not permit access to the plain-text 
data.  
 ● ● 
v7 16.4 Encrypt or Hash all Authentication Credentials 
 Encrypt or hash with a salt all authentication credentials when stored.  ● ●

Page 75 
Internal Only - General 
1.5 SNMP Rules 
Simple Network Management Protocol (SNMP) provides a standards-based interface to 
manage and monitor network devices. This section provides guidance on the secure 
configuration of SNMP parameters. 
The recommendations in this Section apply to Organizations using SNMP. 
Organizations using SNMP should review and implement the recommendations in this 
section.

Page 76 
Internal Only - General 
1.5.1 Set 'no snmp-server' to disable SNMP when unused 
(Manual) 
Profile Applicability: 
•  Level 1 
Description: 
If not in use, disable simple network management protocol (SNMP), read and write 
access. 
Rationale: 
SNMP read access allows remote monitoring and management of the device. 
Impact: 
Organizations not using SNMP should require all SNMP services to be disabled by 
running the 'no snmp-server' command. 
Audit: 
Verify the result reads "SNMP agent not enabled" 
hostname#show snmp community 
Remediation: 
Disable SNMP read and write access if not in used to monitor and/or manage device. 
hostname(config)#no snmp-server  
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
book.html 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v8 
4.4 Implement and Manage a Firewall on Servers 
 Implement and manage a firewall on servers, where supported. Example 
implementations include a virtual firewall, operating system firewall, or a third-
party firewall agent. 
● ● ●

Page 77 
Internal Only - General 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.5 Implement and Manage a Firewall on End-User 
Devices 
 Implement and manage a host-based firewall or port-filtering tool on end-user 
devices, with a default-deny rule that drops all traffic except those services and 
ports that are explicitly allowed. 
● ● ● 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 78 
Internal Only - General 
1.5.2 Unset 'private' for 'snmp-server community' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
An SNMP community string permits read-only access to all objects. 
Rationale: 
The default community string "private" is well known. Using easy to guess, well known 
community string poses a threat that an attacker can effortlessly gain unauthorized 
access to the device. 
Impact: 
To reduce the risk of unauthorized access, Organizations should disable default, easy 
to guess, settings such as the 'private' setting for snmp-server community. 
Audit: 
Perform the following to determine if the public community string is enabled: 
Ensure private does not show as a result 
hostname# show snmp community 
Remediation: 
Disable the default SNMP community string private 
hostname(config)#no snmp-server community {private} 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s2.html#GUID-2F3F13E4-EE81-4590-871D-6AE1043473DE 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 79 
Internal Only - General 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.4 Implement and Manage a Firewall on Servers 
 Implement and manage a firewall on servers, where supported. Example 
implementations include a virtual firewall, operating system firewall, or a third-
party firewall agent. 
● ● ● 
v8 
4.5 Implement and Manage a Firewall on End-User 
Devices 
 Implement and manage a host-based firewall or port-filtering tool on end-user 
devices, with a default-deny rule that drops all traffic except those services and 
ports that are explicitly allowed. 
● ● ● 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 80 
Internal Only - General 
1.5.3 Unset 'public' for 'snmp-server community' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
An SNMP community string permits read-only access to all objects. 
Rationale: 
The default community string "public" is well known. Using easy to guess, well known 
community string poses a threat that an attacker can effortlessly gain unauthorized 
access to the device. 
Impact: 
To reduce the risk of unauthorized access, Organizations should disable default, easy 
to guess, settings such as the 'public' setting for snmp-server community. 
Audit: 
Perform the following to determine if the public community string is enabled: Ensure 
public does not show as a result 
 
hostname# show snmp community 
Remediation: 
Disable the default SNMP community string "public" 
 
hostname(config)#no snmp-server community {public}  
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s2.html#GUID-2F3F13E4-EE81-4590-871D-6AE1043473DE 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 81 
Internal Only - General 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.4 Implement and Manage a Firewall on Servers 
 Implement and manage a firewall on servers, where supported. Example 
implementations include a virtual firewall, operating system firewall, or a third-
party firewall agent. 
● ● ● 
v8 
4.5 Implement and Manage a Firewall on End-User 
Devices 
 Implement and manage a host-based firewall or port-filtering tool on end-user 
devices, with a default-deny rule that drops all traffic except those services and 
ports that are explicitly allowed. 
● ● ● 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 82 
Internal Only - General 
1.5.4 Do not set 'RW' for any 'snmp-server community' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Specifies read-write access. Authorized management stations can both retrieve and 
modify MIB objects. 
Rationale: 
Enabling SNMP read-write enables remote management of the device. Unless 
absolutely necessary, do not allow simple network management protocol (SNMP) write 
access. 
Impact: 
To reduce the risk of unauthorized access, Organizations should disable the SNMP 
'write' access for snmp-server community. 
Audit: 
Perform the following to determine if a read/write community string is enabled: 
Verify the result does not show a community string with a "RW" 
 
hostname#show run | incl snmp-server community 
Remediation: 
Disable SNMP write access. 
 
