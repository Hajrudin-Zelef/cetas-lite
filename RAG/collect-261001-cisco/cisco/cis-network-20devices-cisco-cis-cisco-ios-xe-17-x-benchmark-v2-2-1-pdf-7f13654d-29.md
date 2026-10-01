---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-29
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [5007, 5219]
sha256: bb950a5ad85adb52d42f85da547447a939c5734dc27731d97a3917eac88188fd
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

hostname(config)#interface {<em>interface_name</em>} 
hostname(config-if)#ip authentication key-chain eigrp {<em>eigrp_as-
number</em>} {<em>eigrp_key-chain_name</em>} 
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/interface/command/ir-
i1.html#GUID-0D6BDFCD-3FBB-4D26-A274-C1221F8592DF 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
i1.html#GUID-0B344B46-5E8E-4FE2-A3E0-D92410CE5E91

Page 189 
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

Page 190 
Internal Only - General 
3.3.1.9 Set 'ip authentication mode eigrp' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Configure authentication to prevent unapproved sources from introducing unauthorized 
or false routing messages. 
Rationale: 
This is part of the EIGRP authentication configuration 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Configuring the interface with 'ip 
authentication mode' for EIGRP by number and mode enforces these policies by 
restricting the exchanges between network devices. 
Audit: 
Verify the appropriate authentication mode is set on the appropriate interface(s) 
hostname#sh ip eigrp int 
hostname#sh run int {<em>interface_name</em>} | incl authentication mode 
Remediation: 
Configure the interface with the EIGRP authentication mode. 
 
hostname(config)#interface {<em>interface_name</em>} 
hostname(config-if)#ip authentication mode eigrp {<em><span>eigrp_as-
number</span></em><span>}</span> md5 
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/interface/command/ir-
i1.html#GUID-0D6BDFCD-3FBB-4D26-A274-C1221F8592DF 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_eigrp/command/ire-
i1.html#GUID-8D1B0697-8E96-4D8A-BD20-536956D68506

Page 191 
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

Page 192 
Internal Only - General 
3.3.2 Require OSPF Authentication if Protocol is Used 
Verify open shortest path first (OSPF) authentication is enabled, where feasible.

Page 193 
Internal Only - General 
3.3.2.1 Set 'authentication message-digest' for OSPF area 
(Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Enable MD5 authentication for OSPF. 
Rationale: 
This is part of the OSPF authentication setup. 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Configuring the area 
'authentication message-digest' for OSPF enforces these policies by restricting 
exchanges between network devices. 
Audit: 
Verify message digest for OSPF is defined 
hostname#sh run | sec router ospf  
Remediation: 
Configure the Message Digest option for OSPF. 
 
hostname(config)#router ospf <<em>ospf_process-id</em>> 
hostname(config-router)#area <<em>ospf_area-id</em>> authentication message-
digest  
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_ospf/command/ospf-
i1.html#GUID-3D5781A3-F8DF-4760-A551-6A3AB80A42ED 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_ospf/command/ospf-
a1.html#GUID-81D0F753-D8D5-494E-9A10-B15433CFD445

Page 194 
Internal Only - General 
Additional Information: 
The authentication type must be the same for all routers and access servers in an area. 
The authentication password for all OSPF routers on a network must be the same if 
they are to communicate with each other via OSPF 
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

Page 195 
Internal Only - General 
3.3.2.2 Set 'ip ospf message-digest-key md5' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Enable Open Shortest Path First (OSPF) Message Digest 5 (MD5) authentication. 
Rationale: 
This is part of the OSPF authentication setup 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Configuring the proper 
interface(s) for 'ip ospf message-digest-key md5' enforces these policies by restricting 
exchanges between network devices. 
Audit: 
Verify the appropriate md5 key is defined on the appropriate interface(s) 
 
hostname#sh run int {<em>interface</em>} 
Remediation: 
Configure the appropriate interface(s) for Message Digest authentication 
 
hostname(config)#interface {<em>interface_name</em>} 
hostname(config-if)#ip ospf message-digest-key {<em>ospf_md5_key-id</em>} md5 
{<em>ospf_md5_key</em>} 
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/interface/command/ir-
i1.html#GUID-0D6BDFCD-3FBB-4D26-A274-C1221F8592DF 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_ospf/command/ospf-
i1.html#GUID-939C79FF-8C09-4D5A-AEB5-DAF25038CA18

