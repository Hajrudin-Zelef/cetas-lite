---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-23
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [3752, 3948]
sha256: d309fd9c5ac25cdc85013a72f7d84adf2a2ce28302b95876ba9c1b2af54fc5fc
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

hostname#show run | include ntp trusted-key  
The above command should return any NTP server(s) configured with encryption keys. 
This value should be the same as the total number of servers configured as tested in. 
Remediation: 
Configure the NTP trusted key using the following command 
 
hostname(config)#ntp trusted-key {ntp_key_id} 
Default Value: 
Authentication of the identity of the system is disabled. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/bsm/command/bsm-cr-
n1.html#GUID-89CA798D-0F12-4AE8-B382-DE10CBD261DB

Page 144 
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

Page 145 
Internal Only - General 
2.3.1.4 Set 'key' for each 'ntp server' (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Specifies the authentication key for NTP. 
Rationale: 
This authentication feature provides protection against accidentally synchronizing the 
ntp system to another system that is not trusted, because the other system must know 
the correct authentication key. 
Impact: 
Organizations should establish three Network Time Protocol (NTP) hosts to set 
consistent time across the enterprise. Enabling the 'ntp server key' command enforces 
encrypted authentication between NTP hosts. 
Audit: 
From the command prompt, execute the following commands: 
 
hostname#show run | include ntp server 
Remediation: 
Configure each NTP Server to use a key ring using the following command. 
 
hostname(config)#ntp server {<em>ntp-server_ip_address</em>}{key 
<em>ntp_key_id</em>}  
Default Value: 
No NTP key is set by default 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.4 Standardize Time Synchronization 
 Standardize time synchronization. Configure at least two synchronized time 
sources across enterprise assets, where supported. 
 ● ●

Page 146 
Internal Only - General 
Controls 
Version Control IG 1 IG 2 IG 3 
v7 
6.1 Utilize Three Synchronized Time Sources 
 Use at least three synchronized time sources from which all servers and 
network devices retrieve time information on a regular basis so that timestamps 
in logs are consistent. 
 ● ●

Page 147 
Internal Only - General 
2.3.2 Set 'ip address' for 'ntp server' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Use this command if you want to allow the system to synchronize the system software 
clock with the specified NTP server. 
Rationale: 
To ensure that the time on your Cisco router is consistent with other devices in your 
network, at least two (and preferably at least three) NTP Server/s external to the router 
should be configured. 
Ensure you also configure consistent timezone and daylight savings time setting for all 
devices. For simplicity, the default of Coordinated Universal Time (UTC). 
Impact: 
Organizations should establish multiple Network Time Protocol (NTP) hosts to set 
consistent time across the enterprise. Enabling the 'ntp server ip address' enforces 
encrypted authentication between NTP hosts. 
Audit: 
From the command prompt, execute the following commands: 
 
hostname#sh ntp associations 
Remediation: 
Configure at least one external NTP Server using the following commands 
 
hostname(config)#ntp server {ntp-server_ip_address} 
or  
hostname(config)#ntp server {ntp server vrf [vrf name] ip address} 
Default Value: 
No servers are configured by default. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/bsm/command/bsm-cr-
n1.html#GUID-255145EB-D656-43F0-B361-D9CBCC794112 
2. https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/bsm/command/bsm-cr-
book/bsm-cr-n1.html#wp3294676008

Page 148 
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

Page 149 
Internal Only - General 
2.4 Loopback Rules 
When a router needs to initiate connections to remote hosts, for example for SYSLOG 
or NTP, it will use the nearest interface for the packets source address. This can cause 
issues due to the possible variation in source, potentially causing packets to be denied 
by intervening firewalls or handled incorrectly by the receiving host. To prevent these 
problems the router should be configured with a Loopback interface and any services 
should be bound to this address.

Page 150 
Internal Only - General 
2.4.1 Create a single 'interface loopback' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Configure a single loopback interface. 
Rationale: 
Software-only loopback interface that emulates an interface that is always up. It is a 
virtual interface supported on all platforms. 
Alternate loopback addresses create a potential for abuse, mis-configuration, and 
inconsistencies. Additional loopback interfaces must be documented and approved prior 
to use by local security personnel. 
Impact: 
Organizations should plan and establish 'loopback interfaces' for the enterprise network. 
Loopback interfaces enable critical network information such as OSPF Router IDs and 
provide termination points for routing protocol sessions. 
Audit: 
Perform the following to determine if a loopback interface is defined: 
Verify an IP address returns for the defined loopback interface 
hostname#sh ip int brief | incl Loopback 
Remediation: 
Define and configure one loopback interface. Below is an example 
hostname(config)#interface loopback 0 
hostname(config-if)#ip address 10.10.10.10 255.255.255.0 
Default Value: 
There are no loopback interfaces defined by default. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/interface/command/ir-
i1.html#GUID-0D6BDFCD-3FBB-4D26-A274-C1221F8592DF

Page 151 
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

