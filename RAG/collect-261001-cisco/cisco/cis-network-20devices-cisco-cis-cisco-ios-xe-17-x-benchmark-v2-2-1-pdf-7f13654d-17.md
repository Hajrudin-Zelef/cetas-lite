---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-17
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [2437, 2644]
sha256: 795dcd9ed76fd5bbe6c68881536c6c7aef293aaeea52bccd126ccbec37e3b7f8
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 90 
Internal Only - General 
1.5.8 Set 'snmp-server enable traps snmp' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
SNMP notifications can be sent as traps to authorized management systems. 
Rationale: 
SNMP has the ability to submit traps . 
Impact: 
Organizations using SNMP should restrict trap types only to explicitly named traps to 
reduce unintended traffic. Enabling SNMP traps without specifying trap type will enable 
all SNMP trap types. 
Audit: 
Perform the following to determine if SNMP traps are enabled: 
If the command returns configuration values, then SNMP is enabled. 
 
hostname#show run | incl snmp-server 
Remediation: 
Enable SNMP traps. 
hostname(config)#snmp-server enable traps snmp authentication linkup linkdown 
coldstart  
Default Value: 
SNMP notifications are disabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s3.html#GUID-EB3EB677-A355-42C6-A139-85BA30810C54

Page 91 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
12.8 Establish and Maintain Dedicated Computing 
Resources for All Administrative Work 
 Establish and maintain dedicated computing resources, either physically or 
logically separated, for all administrative tasks or tasks requiring administrative 
access. The computing resources should be segmented from the enterprise's 
primary network and not be allowed internet access. 
  ● 
v7 
11.7 Manage Network Infrastructure Through a Dedicated 
Network 
 Manage the network infrastructure across network connections that are 
separated from the business use of that network, relying on separate VLANs or, 
preferably, on entirely different physical connectivity for management sessions for 
network devices. 
 ● ●

Page 92 
Internal Only - General 
1.5.9 Set 'priv' for each 'snmp-server group' using SNMPv3 
(Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Specifies authentication of a packet with encryption when using SNMPv3 
Rationale: 
SNMPv3 provides much improved security over previous versions by offering options 
for Authentication and Encryption of messages. When configuring a user for SNMPv3 
you have the option of using a range of encryption schemes, or no encryption at all, to 
protect messages in transit. AES128 is the minimum strength encryption method that 
should be deployed. 
Impact: 
Organizations using SNMP can significantly reduce the risks of unauthorized access by 
using the 'snmp-server group v3 priv' setting to encrypt messages in transit. 
Audit: 
Verify the result show the appropriate group name and security model 
 
hostname#show snmp group 
Remediation: 
For each SNMPv3 group created on your router add privacy options by issuing the 
following command... 
 
hostname(config)#snmp-server group {<em>group_name</em>} v3 priv 
Default Value: 
No SNMP server groups are configured. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s5.html#GUID-56E87D02-C56F-4E2D-A5C8-617E31740C3F

Page 93 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
6.5 Require MFA for Administrative Access 
 Require MFA for all administrative access accounts, where supported, on all 
enterprise assets, whether managed on-site or through a third-party provider. 
● ● ● 
v7 
4.5 Use Multifactor Authentication For All Administrative 
Access 
 Use multi-factor authentication and encrypted channels for all administrative 
account access. 
 ● ●

Page 94 
Internal Only - General 
1.5.10 Require 'aes 128' as minimum for 'snmp-server user' when 
using SNMPv3 (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Specify the use of a minimum of 128-bit AES algorithm for encryption when using 
SNMPv3. 
Rationale: 
SNMPv3 provides much improved security over previous versions by offering options 
for Authentication and Encryption of messages. When configuring a user for SNMPv3 
you have the option of using a range of encryption schemes, or no encryption at all, to 
protect messages in transit. AES128 is the minimum strength encryption method that 
should be deployed. 
Impact: 
Organizations using SNMP can significantly reduce the risks of unauthorized access by 
using the 'snmp-server user' setting with appropriate authentication and privacy 
protocols to encrypt messages in transit. 
Audit: 
Verify the result show the appropriate user name and security settings 
 
hostname#show snmp user 
Remediation: 
For each SNMPv3 user created on your router add privacy options by issuing the 
following command. 
 
hostname(config)#snmp-server user {user_name} {group_name} v3 auth sha 
{auth_password} priv aes 128 {priv_password} {acl_name_or_number} 
Default Value: 
SNMP username as not set by default. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s5.html#GUID-4EED4031-E723-4B84-9BBF-610C3CF60E31

Page 95 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
6.5 Require MFA for Administrative Access 
 Require MFA for all administrative access accounts, where supported, on all 
enterprise assets, whether managed on-site or through a third-party provider. 
● ● ● 
v7 
4.5 Use Multifactor Authentication For All Administrative 
Access 
 Use multi-factor authentication and encrypted channels for all administrative 
account access. 
 ● ● 
 
2 Control Plane 
The control plane covers monitoring, route table updates, and generally the dynamic 
operation of the router. Services, settings, and data streams that support and document 
the operation, traffic handling, and dynamic status of the router. Examples of control 
plane services include: logging (e.g. Syslog), routing protocols, status protocols like 
CDP and HSRP, network topology protocols like STP, and traffic security control 
protocols like IKE. Network control protocols like ICMP, NTP, ARP, and IGMP directed 
to or sent by the router itself also fall into this area. 
2.1 Global Service Rules 
Rules in the global service class enforce server and service controls that protect against 
attacks or expose the device to exploitation. 
2.1.1 Setup SSH 
Ensure use of SSH remote console sessions to Cisco routers.

Page 96 
Internal Only - General 
2.1.1.1 Configure Prerequisites for the SSH Service 
[This space intentionally left blank]

Page 97 
Internal Only - General 
2.1.1.1.1 Set the 'hostname' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
The hostname is used in prompts and default configuration filenames. 
Rationale: 
The domain name is prerequisite for setting up SSH. 
Impact: 
Organizations should plan the enterprise network and identify an appropriate host name 
for each router. 
Audit: 
Perform the following to determine if the local time zone is configured: 
Verify the result shows the summer-time recurrence is configured properly. 
 
hostname#sh run | incl hostname 
Remediation: 
Configure an appropriate host name for the router. 
 
hostname(config)#hostname {<em>router_name</em>} 
Default Value: 
The default hostname is Router. 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/F_through_K.html#GUID-F3349988-EC16-
484A-BE81-4C40110E6625 
CIS Controls: 
Controls Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 98 
Internal Only - General 
Controls Version Control IG 1 IG 2 IG 3 
v7 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

