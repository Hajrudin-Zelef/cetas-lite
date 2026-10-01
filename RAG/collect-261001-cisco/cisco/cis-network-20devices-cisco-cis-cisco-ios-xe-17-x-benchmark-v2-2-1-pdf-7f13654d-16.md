---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-16
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [2235, 2436]
sha256: de1bede4f6efbf3e80c461b17c17bad2eb1e4d6343014d58e2ee4fb54ac9e91e
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

hostname(config)#no snmp-server community {<em>write_community_string</em>}  
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s2.html#GUID-2F3F13E4-EE81-4590-871D-6AE1043473DE 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 83 
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

Page 84 
Internal Only - General 
1.5.5 Set the ACL for each 'snmp-server community' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
This feature specifies a list of IP addresses that are allowed to use the community string 
to gain access to the SNMP agent. 
Rationale: 
If ACLs are not applied, then anyone with a valid SNMP community string can 
potentially monitor and manage the router. An ACL should be defined and applied for all 
SNMP access to limit access to a small number of authorized management stations 
segmented in a trusted management zone. If possible, use SNMPv3 which uses 
authentication, authorization, and data privatization (encryption). 
Impact: 
To reduce the risk of unauthorized access, Organizations should enable access control 
lists for all snmp-server communities and restrict the access to appropriate trusted 
management zones. If possible, implement SNMPv3 to apply authentication, 
authorization, and data privatization (encryption) for additional benefits to the 
organization. 
Audit: 
Perform the following to determine if an ACL is enabled: 
Verify the result shows a number after the community string 
 
hostname#show run | incl snmp-server community 
Remediation: 
Configure authorized SNMP community string and restrict access to authorized 
management systems. 
 
hostname(config)#snmp-server community <<em>community_string</em>> ro 
{<em>snmp_access-list_number |  
<span>snmp_access-list_name</span></em><span>}</span> 
Default Value: 
No ACL is set for SNMP

Page 85 
Internal Only - General 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s2.html#GUID-2F3F13E4-EE81-4590-871D-6AE1043473DE 
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

Page 86 
Internal Only - General 
1.5.6 Create an 'access-list' for use with SNMP (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
You can use access lists to control the transmission of packets on an interface, control 
Simple Network Management Protocol (SNMP) access, and restrict the contents of 
routing updates. The Cisco IOS software stops checking the extended access list after a 
match occurs. 
Rationale: 
SNMP ACLs control what addresses are authorized to manage and monitor the device 
via SNMP. If ACLs are not applied, then anyone with a valid SNMP community string 
may monitor and manage the router. An ACL should be defined and applied for all 
SNMP community strings to limit access to a small number of authorized management 
stations segmented in a trusted management zone. 
Audit: 
Perform the following to determine if the ACL is created: 
Verify you the appropriate access-list definitions 
 
hostname#sh ip access-list <<em>snmp_acl_number</em>>  
Remediation: 
Configure SNMP ACL for restricting access to the device from authorized management 
stations segmented in a trusted management zone. 
 
hostname(config)#access-list <<em>snmp_acl_number</em>> permit 
<<em>snmp_access-list</em>> 
hostname(config)#access-list deny any log  
Default Value: 
SNMP does not use an access list. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a2.html#GUID-
9EA733A3-1788-4882-B8C3-AB0A2949120C

Page 87 
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

Page 88 
Internal Only - General 
1.5.7 Set 'snmp-server host' when using SNMP (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
SNMP notifications can be sent as traps to authorized management systems. 
Rationale: 
If SNMP is enabled for device management and device alerts are required, then ensure 
the device is configured to submit traps only to authorize management systems. 
Impact: 
Organizations using SNMP should restrict sending SNMP messages only to explicitly 
named systems to reduce unauthorized access. 
Audit: 
Perform the following to determine if SNMP traps are enabled: 
If the command returns configuration values, then SNMP is enabled. 
 
hostname#show run | incl snmp-server 
Remediation: 
Configure authorized SNMP trap community string and restrict sending messages to 
authorized management systems. 
 
hostname(config)#snmp-server host {ip_address} {trap_community_string} 
{notification-type}  
Default Value: 
A recipient is not specified to receive notifications. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/snmp/command/nm-snmp-cr-
s5.html#GUID-D84B2AB5-6485-4A23-8C26-73E50F73EE61

Page 89 
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

