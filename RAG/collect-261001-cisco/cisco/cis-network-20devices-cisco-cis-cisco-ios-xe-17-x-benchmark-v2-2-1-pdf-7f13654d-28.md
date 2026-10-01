---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-28
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [4796, 5006]
sha256: 38cfae536899a2c9839d2e9fe35b58fa62585455be90bcd7e65308fa860ae0d8
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

hostname#sh run | sec router eigrp  
Remediation: 
Configure the EIGRP address family. 
 
hostname(config)#router eigrp <<em>virtual-instance-name</em>> 
hostname(config-router)#address-family ipv4 autonomous-system {<em>eigrp_as-
number</em>} 
hostname(config-router-af)#af-interface default  
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
i1.html#GUID-67388D6C-AE9C-47CA-8C35-2A2CF9FA668E 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
a1.html#GUID-C03CFC8A-3CE3-4CF9-9D65-52990DBD3377 
3. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
a1.html#GUID-DC0EF1D3-DFD4-45DF-A553-FA432A3E7233

Page 183 
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

Page 184 
Internal Only - General 
3.3.1.6 Set 'authentication key-chain' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Configure the EIGRP address family key chain. 
Rationale: 
This is part of the EIGRP authentication configuration 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using the address-family 'key 
chain' for EIGRP enforces these policies by restricting the exchanges between 
predefined network devices. 
Audit: 
Verify the appropriate key chain is set 
 
hostname#sh run | sec router eigrp  
Remediation: 
Configure the EIGRP address family key chain. 
 
hostname(config)#router eigrp <virtual-instance-name> 
hostname(config-router)#address-family ipv4 autonomous-system {eigrp_as-
number} 
hostname(config-router-af)#af-interface {interface-name} 
hostname(config-router-af-interface)#authentication key-chain {eigrp_key-
chain_name} 
Default Value: 
No key chains are specified for EIGRP 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
i1.html#GUID-67388D6C-AE9C-47CA-8C35-2A2CF9FA668E 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
a1.html#GUID-C03CFC8A-3CE3-4CF9-9D65-52990DBD3377 
3. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
a1.html#GUID-6B6ED6A3-1AAA-4EFA-B6B8-9BF11EEC37A0

Page 185 
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

Page 186 
Internal Only - General 
3.3.1.7 Set 'authentication mode md5' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Configure authentication to prevent unapproved sources from introducing unauthorized 
or false service messages. 
Rationale: 
This is part of the EIGRP authentication configuration 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using the 'authentication mode' 
for EIGRP address-family or service-family packets enforces these policies by 
restricting the type of authentication between network devices. 
Audit: 
Verify the appropriate address family authentication mode is set 
 
hostname#sh run | sec router eigrp  
Remediation: 
Configure the EIGRP address family authentication mode. 
 
hostname(config)#router eigrp <virtual-instance-name> 
hostname(config-router)#address-family ipv4 autonomous-system {eigrp_as-
number} 
hostname(config-router-af)#af-interface {interface-name} 
hostname(config-router-af-interface)#authentication mode md5 
Default Value: 
Not defined 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
i1.html#GUID-67388D6C-AE9C-47CA-8C35-2A2CF9FA668E 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
a1.html#GUID-C03CFC8A-3CE3-4CF9-9D65-52990DBD3377 
3. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
a1.html#GUID-A29E0EF6-4CEF-40A7-9824-367939001B73

Page 187 
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

Page 188 
Internal Only - General 
3.3.1.8 Set 'ip authentication key-chain eigrp' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Specify the type of authentication used in Enhanced Interior Gateway Routing Protocol 
(EIGRP) packets per interface. 
Rationale: 
Configuring EIGRP authentication key-chain number and name to restrict packet 
exchanges between network devices. 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Configuring the interface with 'ip 
authentication key chain' for EIGRP by name and number enforces these policies by 
restricting the exchanges between network devices. 
Audit: 
Verify the appropriate key chain is set on the appropriate interface(s) 
hostname#sh ip eigrp int 
hostname#sh run int {<em>interface_name</em>} | incl key-chain  
Remediation: 
Configure the interface with the EIGRP key chain. 
 
