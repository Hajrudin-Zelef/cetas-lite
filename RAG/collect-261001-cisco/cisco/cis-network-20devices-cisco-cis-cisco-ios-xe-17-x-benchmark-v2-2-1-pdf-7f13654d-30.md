---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-30
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["benchmark", "agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [5220, 5430]
sha256: def36522db925f8bf1fdaeeaefbb87992a7f92495a00b9f97563a479b200dc30
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 196 
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

Page 197 
Internal Only - General 
3.3.3 Require BGP Authentication if Protocol is Used 
Border Gateway Protocol (BGP)is a path vector protocol used for interior and exterior 
gateway routing on some networks. 
BGP is a complex protocol, with many configuration options which may have effects 
which are not immediately obvious. 
Verify Border Gateway Protocol (BGP) authentication is enabled, if routing protocol is 
used, where feasible.

Page 198 
Internal Only - General 
3.3.3.1 Set 'neighbor password' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Enable message digest5 (MD5) authentication on a TCP connection between two BGP 
peers 
Rationale: 
Enforcing routing authentication reduces the likelihood of routing poisoning and 
unauthorized routers from joining BGP routing. 
Impact: 
Organizations should plan and implement enterprise security policies that require 
rigorous authentication methods for routing protocols. Using the 'neighbor password' for 
BGP enforces these policies by restricting the type of authentication between network 
devices. 
Audit: 
Verify you see the appropriate neighbor password is defined: 
 
hostname#sh run | sec router bgp  
Remediation: 
Configure BGP neighbor authentication where feasible. 
 
hostname(config)#router bgp <<em>bgp_as-number</em>> 
hostname(config-router)#neighbor <<em>bgp_neighbor-ip</em> | <em>peer-group-
name</em>> password <<em>password</em>>  
Default Value: 
Not set 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_bgp/command/bgp-
n1.html#GUID-A8900842-ECF3-42D3-B188-921BE0EC060B 
2. http://www.cisco.com/en/US/docs/ios-xml/ios/iproute_bgp/command/bgp-
m1.html#GUID-159A8006-F0DF-4B82-BB71-C39D2C134205

Page 199 
Internal Only - General 
Additional Information: 
MD5 authentication between two BGP peers, meaning that each segment sent on the 
TCP connection between the peers is verified. MD5 authentication must be configured 
with the same password on both BGP peers. 
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

Page 200 
Internal Only - General 
Appendix: Summary Table 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
1 Management Plane 
1.1 Local Authentication, Authorization and Accounting (AAA) Rules 
1.1.1 Enable 'aaa new-model' (Automated)   
1.1.2 Enable 'aaa authentication login' (Automated)   
1.1.3 Enable 'aaa authentication enable default' (Automated)   
1.1.4 Set 'login authentication for 'line vty' (Automated)   
1.1.5 Set 'login authentication for 'ip http' (Automated)   
1.1.6 Set 'aaa accounting' to log all privileged use commands 
using 'commands 15' (Automated) 
  
1.1.7 Set 'aaa accounting connection' (Automated)   
1.1.8 Set 'aaa accounting exec' (Automated)   
1.1.9 Set 'aaa accounting network' (Automated)   
1.1.10 Set 'aaa accounting system' (Automated)   
1.2 Access Rules 
1.2.1 Set 'privilege 1' for local users (Manual)   
1.2.2 Set 'transport input ssh' for 'line vty' connections 
(Automated) 
  
1.2.3 Set 'no exec' for 'line aux 0' (Automated)   
1.2.4 Create 'access-list' for use with 'line vty' (Manual)   
1.2.5 Set 'access-class' for 'line vty' (Automated)  

Page 201 
Internal Only - General 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
1.2.6 Set 'exec-timeout' to less than or equal to 10 minutes for 
'line aux 0' (Automated) 
  
1.2.7 Set 'exec-timeout' to less than or equal to 10 minutes 
'line console 0' (Automated) 
  
1.2.8 Set 'exec-timeout' to less than or equal to 10 minutes 
'line vty' (Automated) 
  
1.2.9 Set 'http Secure-server' limit (Automated)   
1.2.10 Set 'exec-timeout' to less than or equal to 10 min on 'ip 
http' (Automated) 
  
1.3 Banner Rules 
1.3.1 Set the 'banner-text' for 'banner exec' (Automated)   
1.3.2 Set the 'banner-text' for 'banner login' (Automated)   
1.3.3 Set the 'banner-text' for 'banner motd' (Automated)   
1.3.4 Set the 'banner-text' for 'webauth banner' (Automated)   
1.4 Password Rules 
1.4.1 Set 'password' for 'enable secret' (Automated)   
1.4.2 Enable 'service password-encryption' (Automated)   
1.4.3 Set 'username secret' for all local users (Automated)   
1.5 SNMP Rules 
1.5.1 Set 'no snmp-server' to disable SNMP when unused 
(Manual) 
  
1.5.2 Unset 'private' for 'snmp-server community' (Automated)   
1.5.3 Unset 'public' for 'snmp-server community' (Automated)  

Page 202 
Internal Only - General 
CIS Benchmark Recommendation Set 
Correctly 
Yes No 
1.5.4 Do not set 'RW' for any 'snmp-server community' 
(Manual) 
  
1.5.5 Set the ACL for each 'snmp-server community' (Manual)   
1.5.6 Create an 'access-list' for use with SNMP (Manual)   
1.5.7 Set 'snmp-server host' when using SNMP (Automated)   
1.5.8 Set 'snmp-server enable traps snmp' (Automated)   
1.5.9 Set 'priv' for each 'snmp-server group' using SNMPv3 
(Automated) 
  
1.5.10 Require 'aes 128' as minimum for 'snmp-server user' 
when using SNMPv3 (Manual) 
  
2 Control Plane 
2.1 Global Service Rules 
2.1.1 Setup SSH 
2.1.1.1 Configure Prerequisites for the SSH Service 
2.1.1.1.1 Set the 'hostname' (Automated)   
2.1.1.1.2 Set the 'ip domain-name' (Automated)   
2.1.1.1.3 Set 'modulus' to greater than or equal to 2048 for 'crypto 
key generate rsa' (Manual) 
  
2.1.1.1.4 Set 'seconds' for 'ip ssh timeout' for 60 seconds or less 
(Automated) 
  
2.1.1.1.5 Set maximum value for 'ip ssh authentication-retries' 
(Automated) 
  
2.1.1.2 Set version 2 for 'ip ssh version' (Manual)   
2.1.2 Set 'no cdp run' (Manual)  

