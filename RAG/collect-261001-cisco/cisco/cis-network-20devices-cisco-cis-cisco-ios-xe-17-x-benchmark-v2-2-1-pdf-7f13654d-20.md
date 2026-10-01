---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-20
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [3084, 3314]
sha256: 87b7f7621dcbb3baced1619c31f7bd4ca235718d1909fd37e4e81b253a6b0e7a
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 116 
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

Page 117 
Internal Only - General 
2.1.6 Set 'service tcp-keepalives-out' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Generate keepalive packets on idle outgoing network connections. 
Rationale: 
Stale connections use resources and could potentially be hijacked to gain illegitimate 
access. The TCP keepalives-in service generates keepalive packets on idle incoming 
network connections (initiated by remote host). This service allows the device to detect 
when the remote host fails and drop the session. If enabled, keepalives are sent once 
per minute on idle connections. The closes connection is closed within five minutes if no 
keepalives are received or immediately if the host replies with a reset packet. 
Impact: 
To reduce the risk of unauthorized access, organizations should implement a security 
policy restricting how long to allow terminated sessions and enforce this policy through 
the use of 'tcp-keepalives-out' command. 
Audit: 
Perform the following to determine if the feature is enabled: 
Verify a command string result returns 
 
hostname#show run | incl service tcp 
Remediation: 
Enable TCP keepalives-out service: 
 
hostname(config)#service tcp-keepalives-out 
Default Value: 
Disabled by default. 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/R_through_setup.html#GUID-9321ECDC-6284-
4BF6-BA4A-9CEEF5F993E5

Page 118 
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

Page 119 
Internal Only - General 
2.1.7 Set 'no service pad' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Disable X.25 Packet Assembler/Disassembler (PAD) service. 
Rationale: 
If the PAD service is not necessary, disable the service to prevent intruders from 
accessing the X.25 PAD command set on the router. 
Impact: 
To reduce the risk of unauthorized access, organizations should implement a security 
policy restricting unnecessary services such as the 'PAD' service. 
Audit: 
Perform the following to determine if the feature is disabled: 
Verify no result returns 
 
hostname#show run all| incl service pad 
Remediation: 
Disable the PAD service. 
 
hostname(config)#no service pad 
Default Value: 
Enabled by default. 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/wan/command/wan-s1.html#GUID-
C5497B77-3FD4-4D2F-AB08-1317D5F5473B 
Additional Information: 
Reverting the Artifact logic will satisfy this control and will make this control independent 
of the way the assessor tool used checks the configuration (CIS CAT Pro Assessor 
check in 'sh run' results instead of 'sh run all' results)

Page 120 
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

Page 121 
Internal Only - General 
2.2 Logging Rules 
Rules in the logging class enforce controls that provide a record of system activity and 
events.

Page 122 
Internal Only - General 
2.2.1 Set 'logging enable' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Enable logging of system messages. 
Rationale: 
Logging provides a chronological record of activities on the Cisco device and allows 
monitoring of both operational and security related events. 
Impact: 
Enabling the Cisco IOS 'logging enable' command enforces the monitoring of 
technology risks for the organizations' network devices. 
Audit: 
Perform the following to determine if the feature is enabled: 
Verify no result returns 
 
hostname#show run | i logging host 
Remediation: 
Enable system logging. 
 
hostname(config)#archive 
hostname(config-archive)#log config 
hostname(config-archive-log-cfg)#logging enable 
hostname(config-archive-log-cfg)#end 
Default Value: 
Logging is not enabled/ 
References: 
1. https://community.cisco.com/t5/networking-knowledge-base/how-to-configure-
logging-in-cisco-ios/ta-p/3132434

Page 123 
Internal Only - General 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
8.5 Collect Detailed Audit Logs 
 Configure detailed audit logging for enterprise assets containing sensitive data. 
Include event source, date, username, timestamp, source addresses, destination 
addresses, and other useful elements that could assist in a forensic investigation. 
 ● ● 
v7 
6.3 Enable Detailed Logging 
 Enable system logging to include detailed information such as an event source, 
date, user, timestamp, source addresses, destination addresses, and other useful 
elements. 
 ● ●

Page 124 
Internal Only - General 
2.2.2 Set 'buffer size' for 'logging buffered' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Enable system message logging to a local buffer. 
Rationale: 
The device can copy and store log messages to an internal memory buffer. The 
buffered data is available only from a router exec or enabled exec session. This form of 
logging is useful for debugging and monitoring when logged in to a router. 
Impact: 
Data forensics is effective for managing technology risks and an organization can 
enforce such policies by enabling the 'logging buffered' command. 
Audit: 
Perform the following to determine if the feature is enabled: 
Verify a command string result returns 
 
hostname#show run | incl logging buffered 
Remediation: 
Configure buffered logging (with minimum size). Recommended size is 64000. 
 
hostname(config)#logging buffered [<em>log_buffer_size</em>] 
Default Value: 
No logging buffer is set by default 
References: 
1. http://www.cisco.com/en/US/docs/ios/netmgmt/command/reference/nm_09.html#
wp1060051

