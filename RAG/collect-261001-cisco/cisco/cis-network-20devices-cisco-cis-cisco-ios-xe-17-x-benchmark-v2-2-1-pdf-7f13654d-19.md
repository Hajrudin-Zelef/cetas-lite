---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-19
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [2868, 3083]
sha256: f00875c1e1bd7ea8282c23549db994555b7e6197f9d4cd9c0c4e7d0ab8af4e5f
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 109 
Internal Only - General 
2.1.2 Set 'no cdp run' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Disable Cisco Discovery Protocol (CDP) service at device level. 
Rationale: 
The Cisco Discovery Protocol is a proprietary protocol that Cisco devices use to identify 
each other on a LAN segment. It is useful only in network monitoring and 
troubleshooting situations but is considered a security risk because of the amount of 
information provided from queries. In addition, there have been published denial-of-
service (DoS) attacks that use CDP. CDP should be completely disabled unless 
necessary. 
Impact: 
To reduce the risk of unauthorized access, organizations should implement a security 
policy restricting network protocols and explicitly require disabling all insecure or 
unnecessary protocols. 
Audit: 
Perform the following to determine if CDP is enabled: 
Verify the result shows "CDP is not enabled" 
 
hostname#show cdp 
Remediation: 
Disable Cisco Discovery Protocol (CDP) service globally. 
 
hostname(config)#no cdp run 
Default Value: 
Enabled on all platforms except the Cisco 10000 Series Edge Services Router 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/cdp/command/cdp-cr-
a1.html#GUID-E006FAC8-417E-4C3F-B732-4D47B0447750

Page 110 
Internal Only - General 
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

Page 111 
Internal Only - General 
2.1.3 Set 'no ip bootp server' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Disable the Bootstrap Protocol (BOOTP) service on your routing device. 
Rationale: 
BootP allows a router to issue IP addresses. This should be disabled unless there is a 
specific requirement. 
Impact: 
To reduce the risk of unauthorized access, organizations should implement a security 
policy restricting network protocols and explicitly require disabling all insecure or 
unnecessary protocols such as 'ip bootp server'. 
Audit: 
Perform the following to determine if bootp is enabled: 
Verify a "no ip bootp server" result returns 
 
hostname#show run | incl bootp 
Remediation: 
Disable the bootp server. 
 
hostname(config)#ip dhcp bootp ignore 
Default Value: 
Enabled 
References: 
1. Cisco IOS software receives Cisco Discovery Protocol information 
Additional Information: 
Adjusting the Artifact's logic reversing it: check that 'ip bootp server' does not exist in the 
global config.

Page 112 
Internal Only - General 
Doing so allows CIS CAT Pro assessor to provide the right result, because CIS CAT 
Pro Assessor relies on a 'sh run' results and as 'no ip bootp server' is the default setting 
it doesn't show up in 'sh run' results. When bootp server is enabled, the line 'ip bootp 
server' exists in global config and shows up in 'sh run' results. 
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

Page 113 
Internal Only - General 
2.1.4 Set 'no service dhcp' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Disable the Dynamic Host Configuration Protocol (DHCP) server and relay agent 
features on your router. 
Rationale: 
The DHCP server supplies automatic configuration parameters, such as dynamic IP 
address, to requesting systems. A dedicated server located in a secured management 
zone should be used to provide DHCP services instead. Attackers can potentially be 
used for denial-of-service (DoS) attacks. 
Impact: 
To reduce the risk of unauthorized access, organizations should implement a security 
policy restricting network protocols and explicitly require disabling all insecure or 
unnecessary protocols such as the Dynamic Host Configuration Protocol (DHCP). 
Audit: 
Perform the following to determine if the DHCP service is enabled: 
Verify no result returns 
 
hostname#show run | incl dhcp 
Remediation: 
Disable the DHCP server. 
 
hostname(config)#<strong>no service dhcp</strong> 
Default Value: 
Enabled by default, but also requires a DHCP pool to be set to activate the DHCP 
server. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/ipaddr/command/ipaddr-
r1.html#GUID-1516B259-AA28-4839-B968-8DDBF0B382F6

Page 114 
Internal Only - General 
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

Page 115 
Internal Only - General 
2.1.5 Set 'service tcp-keepalives-in' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Generate keepalive packets on idle incoming network connections. 
Rationale: 
Stale connections use resources and could potentially be hijacked to gain illegitimate 
access. The TCP keepalives-in service generates keepalive packets on idle incoming 
network connections (initiated by remote host). This service allows the device to detect 
when the remote host fails and drop the session. If enabled, keepalives are sent once 
per minute on idle connections. The connection is closed within five minutes if no 
keepalives are received or immediately if the host replies with a reset packet. 
Impact: 
To reduce the risk of unauthorized access, organizations should implement a security 
policy restricting how long to allow terminated sessions and enforce this policy through 
the use of 'tcp-keepalives-in' command. 
Audit: 
Perform the following to determine if the feature is enabled: 
Verify a command string result returns 
 
hostname#show run | incl service tcp 
Remediation: 
Enable TCP keepalives-in service: 
 
hostname(config)#service tcp-keepalives-in 
Default Value: 
Disabled by default. 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/R_through_setup.html#GUID-1489ABA3-2428-
4A64-B252-296A035DB85E

