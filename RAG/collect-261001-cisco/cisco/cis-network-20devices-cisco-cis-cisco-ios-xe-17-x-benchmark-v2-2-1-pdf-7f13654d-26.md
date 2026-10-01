---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-26
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [4366, 4563]
sha256: d7d1860b6a80363881d50353d4327b49a60f1d438d03600e9bbc999d3fbcf1eb
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 168 
Internal Only - General 
3.2.1 Set 'ip access-list extended' to Forbid Private Source 
Addresses from External Networks (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
This command places the router in access-list configuration mode, where you must 
define the denied or permitted access conditions by using the deny and permit 
commands. 
Rationale: 
Configuring access controls can help prevent spoofing attacks. To reduce the 
effectiveness of IP spoofing, configure access control to deny any traffic from the 
external network that has a source address that should reside on the internal network. 
Include local host address or any reserved private addresses (RFC 1918). 
Ensure the permit rule(s) above the final deny rule only allow traffic according to your 
organization's least privilege policy. 
Impact: 
Organizations should plan and implement enterprise security policies that explicitly 
separate internal from external networks. Adding 'ip access-list' explicitly permitting and 
denying internal and external networks enforces these policies. 
Audit: 
Verify you have the appropriate access-list definitions 
 
hostname#sh ip access-list {<em>name | number</em>} 
Remediation: 
Configure ACL for private source address restrictions from external networks.

Page 169 
Internal Only - General 
 
hostname(config)#ip access-list extended {<span><em>name | number</em>}  
</span><span>hostname(config-nacl)#deny ip 
{</span><em>internal_networks</em>} any log 
hostname(config<span>-nacl</span>)#deny ip 127.0.0.0 0.255.255.255 any log 
hostname(config<span>-nacl</span>)#deny ip 10.0.0.0 0.255.255.255 any log 
hostname(config<span>-nacl</span>)#deny ip 0.0.0.0 0.255.255.255 any log 
hostname(config<span>-nacl</span>)#deny ip 172.16.0.0 0.15.255.255 any log 
hostname(config<span>-nacl</span>)#deny ip 192.168.0.0 0.0.255.255 any log 
hostname(config<span>-nacl</span>)#deny ip 192.0.2.0 0.0.0.255 any log 
hostname(config<span>-nacl</span>)#deny ip 169.254.0.0 0.0.255.255 any log 
hostname(config<span>-nacl</span>)#deny ip 224.0.0.0 31.255.255.255 any log 
hostname(config<span>-nacl</span>)#deny ip host 255.255.255.255 any log 
hostname(config<span>-nacl</span>)#permit {protocol} {source_ip} 
{source_mask} {destination} {destination_mask} log 
hostname(config<span>-nacl</span>)#deny any any log 
hostname(config)#interface <external_<em>interface</em>> 
hostname(config-if)#access-group <<em>access-list</em>> in 
  
Default Value: 
No access list defined 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i1.html#GUID-
BD76E065-8EAC-4B32-AF25-04BA94DD2B11 
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

Page 170 
Internal Only - General 
Controls 
Version Control IG 1 IG 2 IG 3 
v7 
9.2 Ensure Only Approved Ports, Protocols and Services 
Are Running 
 Ensure that only network ports, protocols, and services listening on a system 
with validated business needs, are running on each system. 
 ● ●

Page 171 
Internal Only - General 
3.2.2 Set inbound 'ip access-group' on the External Interface 
(Manual) 
Profile Applicability: 
•  Level 2 
Description: 
This command places the router in access-list configuration mode, where you must 
define the denied or permitted access conditions by using the deny and permit 
commands. 
Rationale: 
Configuring access controls can help prevent spoofing attacks. To reduce the 
effectiveness of IP spoofing, configure access control to deny any traffic from the 
external network that has a source address that should reside on the internal network. 
Include local host address or any reserved private addresses (RFC 1918). 
Ensure the permit rule(s) above the final deny rule only allow traffic according to your 
organization's least privilege policy. 
Impact: 
Organizations should plan and implement enterprise security policies explicitly 
permitting and denying access based upon access lists. Using the 'ip access-group' 
command enforces these policies by explicitly identifying groups permitted access. 
Audit: 
Verify the access-group is applied to the appropriate interface 
 
hostname#sh run | sec interface {<em>external_interface</em>} 
Remediation: 
Apply the access-group for the external (untrusted) interface 
 
hostname(config)#interface {external_interface} 
hostname(config-if)#ip access-group {name | number} in 
Default Value: 
No access-group defined

Page 172 
Internal Only - General 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/interface/command/ir-
i1.html#GUID-0D6BDFCD-3FBB-4D26-A274-C1221F8592DF 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/security/d1/sec-cr-i1.html#GUID-
D9FE7E44-7831-4C64-ACB8-840811A0C993 
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
 
3.3 Neighbor Authentication 
Enable routing authentication.

Page 173 
Internal Only - General 
3.3.1 Require EIGRP Authentication if Protocol is Used 
Verify enhanced interior gateway routing protocol (EIGRP) authentication is enabled, if 
routing protocol is used, where feasible.

Page 174 
Internal Only - General 
3.3.1.1 Set 'key chain' (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Define an authentication key chain to enable authentication for routing protocols. A key 
chain must have at least one key and can have up to 2,147,483,647 keys. 
NOTE: Only DRP Agent, EIGRP, and RIPv2 use key chains. 
Rationale: 
Routing protocols such as DRP Agent, EIGRP, and RIPv2 use key chains for 
authentication. 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using 'key chains' for routing 
protocols enforces these policies. 
Audit: 
Verify the appropriate key chain is defined 
 
hostname#sh run | sec key chain  
Remediation: 
Establish the key chain. 
 
hostname(config)#key chain {<em>key-chain_name</em>} 
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_pi/command/iri-cr-
a1.html#GUID-A62E89F5-0B8B-4CF0-B4EB-08F2762D88BB

