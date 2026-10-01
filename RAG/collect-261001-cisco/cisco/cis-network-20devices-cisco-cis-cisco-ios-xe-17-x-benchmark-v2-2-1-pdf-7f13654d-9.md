---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-9
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [1009, 1192]
sha256: cac71d973974860d5ca9bbcfe0b0c2bb378adb763cd10973ed41b15e41727e41
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 39 
Internal Only - General 
1.2.1 Set 'privilege 1' for local users (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Sets the privilege level for the user. 
Rationale: 
Default device configuration does not require strong user authentication potentially 
enabling unfettered access to an attacker that is able to reach the device. Creating a 
local account with privilege level 1 permissions only allows the local user to access the 
device with EXEC-level permissions and will be unable to modify the device without 
using the enable password. In addition, require the use of an encrypted password as 
well (see Section 1.1.4.4 - Require Encrypted User Passwords). 
Impact: 
Organizations should create policies requiring all local accounts with 'privilege level 1' 
with encrypted passwords to reduce the risk of unauthorized access. Default 
configuration settings do not provide strong user authentication to the device. 
Audit: 
Perform the following to determine if a user with an encrypted password is enabled: 
Verify all username results return "privilege 1" 
hostname#show running-config | incl privilege 
Remediation: 
Set the local user to privilege level 1. 
hostname(config)#username <LOCAL_USERNAME> privilege 1  
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/s1/sec-cr-t2-z.html#GUID-
34B3E43E-0F79-40E8-82B6-A4B5F1AFF1AD 
CIS Controls: 
Controls Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 40 
Internal Only - General 
Controls Version Control IG 1 IG 2 IG 3 
v7 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 41 
Internal Only - General 
1.2.2 Set 'transport input ssh' for 'line vty' connections 
(Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Selects the Secure Shell (SSH) protocol. 
Rationale: 
Configuring VTY access control restricts remote access to only those authorized to 
manage the device and prevents unauthorized users from accessing the system. 
Impact: 
To reduce risk of unauthorized access, organizations should require all VTY 
management line protocols to be limited to ssh. 
Audit: 
Perform the following to determine if SSH is the only transport method for incoming VTY 
logins: 
The result should show only "ssh" for "transport input" 
hostname#show running-config | sec vty    
Remediation: 
Apply SSH to transport input on all VTY management lines 
hostname(config)#line vty <line-number> <ending-line-number> 
hostname(config-line)#transport input ssh  
References: 
1. http://www.cisco.com/en/US/docs/ios/termserv/command/reference/tsv_s1.html#
wp1069219 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
6.5 Require MFA for Administrative Access 
 Require MFA for all administrative access accounts, where supported, on all 
enterprise assets, whether managed on-site or through a third-party provider. 
● ● ●

Page 42 
Internal Only - General 
Controls 
Version Control IG 1 IG 2 IG 3 
v7 
4.5 Use Multifactor Authentication For All Administrative 
Access 
 Use multi-factor authentication and encrypted channels for all administrative 
account access. 
 ● ●

Page 43 
Internal Only - General 
1.2.3 Set 'no exec' for 'line aux 0' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
The 'no exec' command restricts a line to outgoing connections only. 
Rationale: 
Unused ports should be disabled, if not required, since they provide a potential access 
path for attackers. Some devices include both an auxiliary and console port that can be 
used to locally connect to and configure the device. The console port is normally the 
primary port used to configure the device; even when remote, backup administration is 
required via console server or Keyboard, Video, Mouse (KVM) hardware. The auxiliary 
port is primarily used for dial-up administration via an external modem; instead, use 
other available methods. 
Impact: 
Organizations can reduce the risk of unauthorized access by disabling the 'aux' port 
with the 'no exec' command. Conversely, not restricting access through the 'aux' port 
increases the risk of remote unauthorized access. 
Audit: 
Perform the following to determine if the EXEC process for the aux port is disabled: 
Verify no exec 
hostname#show running-config all | sec aux 
Verify you see the following "no exec" 
Remediation: 
Disable the EXEC process on the auxiliary port. 
hostname(config)#line aux 0 
hostname(config-line)#no exec 
References: 
1. https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst9200/software/releas
e/17-6/command_reference/b_176_9200_cr.html

Page 44 
Internal Only - General 
Additional Information: 
Some Cisco devices don't even have an auxiliary port, therefore this control will fail. 
Adjusting this control logic: check for existence of auxiliary port (ie check for 'none_exist' 
of a section with 'aux') combined with a logical OR with the current control. 
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

Page 45 
Internal Only - General 
1.2.4 Create 'access-list' for use with 'line vty' (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Access lists control the transmission of packets on an interface, control Virtual Terminal 
Line (VTY) access, and restrict the contents of routing updates. The Cisco IOS software 
stops checking the extended access list after a match occurs. 
Rationale: 
VTY ACLs control what addresses may attempt to log in to the router. Configuring VTY 
lines to use an ACL, restricts the sources where a user can manage the device. You 
should limit the specific host(s) and or network(s) authorized to connect to and configure 
the device, via an approved protocol, to those individuals or systems authorized to 
administer the device. For example, you could limit access to specific hosts, so that only 
network managers can configure the devices only by using specific network 
management workstations. Make sure you configure all VTY lines to use the same ACL. 
Impact: 
Organizations can reduce the risk of unauthorized access by implementing access-lists 
for all VTY lines. Conversely, using VTY lines without access-lists increases the risk of 
unauthorized access. 
Audit: 
Perform the following to determine if the ACL is created: 
Verify the appropriate access-list definitions 
hostname#sh ip access-list <vty_acl_number> 
Remediation: 
Configure the VTY ACL that will be used to restrict management access to the device. 
hostname(config)#access-list <vty_acl_number> permit tcp 
<vty_acl_block_with_mask> any 
hostname(config)#access-list <vty_acl_number> permit tcp host <vty_acl_host> 
any 
hostname(config)#deny ip any any log 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/a1/sec-cr-a2.html#GUID-
9EA733A3-1788-4882-B8C3-AB0A2949120C

