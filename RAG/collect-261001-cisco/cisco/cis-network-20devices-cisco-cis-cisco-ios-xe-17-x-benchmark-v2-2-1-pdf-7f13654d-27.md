---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-27
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [4564, 4795]
sha256: 9b07d9ceb4d1f86a809d4f80f3280c66763f7573bc8f845bf98b91c5a1cc1e5e
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 175 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.2 Establish and Maintain a Secure Configuration 
Process for Network Infrastructure 
 Establish and maintain a secure configuration process for network devices. 
Review and update documentation annually, or when significant enterprise 
changes occur that could impact this Safeguard. 
● ● ● 
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

Page 176 
Internal Only - General 
3.3.1.2 Set 'key' (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Configure an authentication key on a key chain. 
Rationale: 
This is part of the routing authentication setup 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using 'key numbers' for key 
chains for routing protocols enforces these policies. 
Audit: 
Verify the appropriate key chain is defined 
 
hostname#sh run | sec key chain  
Remediation: 
Configure the key number. 
 
hostname(config-keychain)#key {<em>key-number</em>} 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_pi/command/iri-cr-
a1.html#GUID-3F31B2E0-0E4B-4F49-A4A8-8ADA1CA0D73F 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.2 Establish and Maintain a Secure Configuration 
Process for Network Infrastructure 
 Establish and maintain a secure configuration process for network devices. 
Review and update documentation annually, or when significant enterprise 
changes occur that could impact this Safeguard. 
● ● ●

Page 177 
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

Page 178 
Internal Only - General 
3.3.1.3 Set 'key-string' (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Configure the authentication string for a key. 
Rationale: 
This is part of the routing authentication setup 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using 'key strings' for key chains 
for routing protocols enforces these policies. 
Audit: 
Verify the appropriate key chain is defined 
 
hostname#sh run | sec key chain  
Remediation: 
Configure the key string. 
 
hostname(config-keychain-key)#key-string <<em>key-string</em>>  
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_pi/command/iri-cr-
a1.html#GUID-D7A8DC18-2E16-4EA5-8762-8B68B94CC43E

Page 179 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.2 Establish and Maintain a Secure Configuration 
Process for Network Infrastructure 
 Establish and maintain a secure configuration process for network devices. 
Review and update documentation annually, or when significant enterprise 
changes occur that could impact this Safeguard. 
● ● ● 
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

Page 180 
Internal Only - General 
3.3.1.4 Set 'address-family ipv4 autonomous-system' (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Configure the EIGRP address family. 
Rationale: 
Rationale: EIGRP is a true multi-protocol routing protocol and the 'address-family' 
feature enables restriction of exchanges with specific neighbors 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using 'address-family' for EIGRP 
enforces these policies by restricting the exchanges between predefined network 
devices. 
Audit: 
Verify the appropriate address family is set 
 
hostname#sh run | sec router eigrp  
Remediation: 
Configure the EIGRP address family. 
 
hostname(config)#router eigrp <<em>virtual-instance-name</em>> 
hostname(config-router)#address-family ipv4 autonomous-system {<em>eigrp_as-
number</em>} 
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
i1.html#GUID-67388D6C-AE9C-47CA-8C35-2A2CF9FA668E 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
a1.html#GUID-C03CFC8A-3CE3-4CF9-9D65-52990DBD3377

Page 181 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
4.2 Establish and Maintain a Secure Configuration 
Process for Network Infrastructure 
 Establish and maintain a secure configuration process for network devices. 
Review and update documentation annually, or when significant enterprise 
changes occur that could impact this Safeguard. 
● ● ● 
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

Page 182 
Internal Only - General 
3.3.1.5 Set 'af-interface default' (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Defines user defaults to apply to EIGRP interfaces that belong to an address-family. 
Rationale: 
Part of the EIGRP address-family setup 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using 'af-interface default' for 
EIGRP interfaces enforces these policies by restricting the exchanges between 
predefined network devices. 
Audit: 
Verify the setting 
 
