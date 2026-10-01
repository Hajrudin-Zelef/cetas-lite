---
id: collect-261001-huawei/huawei/file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0-3
title: "file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0"
domain: huawei
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["cyber"]
source: docs/RAG/collect-261001-huawei/file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0.md
source_anchor: ""
source_lines: [195, 282]
sha256: 2d2b3731aeb6eabe99c2be00602d5ef2049f6eae22304d9c470cee4c56958697
---

# file-repository-cyber-alerts-enhanced-visibility-and-hardening-guidance-for-comm-4e3353b0

TLP:CLEAR 
CISA | FBI | NSA | ASD’s ACSC | CCCS | NCSC-NZ 
TLP:CLEAR 
6 
o Ensure no passwords are reset back to the default. 
 Confirm the integrity of the software image in use by using a trusted hashing calculation 
utility, if available. 
o If a utility is unavailable, calculate a hash of the software image on a trusted 
administration workstation and compare against the vendor’s published hashes on an 
authenticated site as a trusted source of truth. This may require engaging the device’s 
maintenance contract to access source of truth hash values. For additional security, copy 
the image to a forensic workstation and calculate the hash value to compare against the 
vendor’s published hashes. 
Network Defenders 
 Disable any unnecessary, unused, exploitable, or plaintext services and protocols, such as 
Telnet, File Transfer Protocol (FTP), Trivial FTP (TFTP), SSH v1, Hypertext Transfer Protocol 
(HTTP) servers, and SNMP v1/v2c. Ensure any required internet-exposed services are 
adequately protected by ACLs and are fully patched. 
 Conduct port-scanning and scanning of known internet-facing infrastructure to ensure no 
additional services are accessible across the network or from the internet. Remove 
unnecessary internet-facing infrastructure, monitor necessary internet-facing infrastructure, 
and continuously validate the architecture. 
o Routers with an active shell environment—even if they have not been tampered with—
have significantly more listeners running at the operating system (OS) level compared to 
the software level. 
Network defenders and network engineers should ensure close collaboration and open 
communication to accomplish the following: 
 Ensure all networking configurations are stored, tracked, and regularly audited for 
compliance with security policies and best practices. 
o Whenever networking configurations are transmitted for storage, tracking, and 
troubleshooting, confirm that they are sent using encrypted protocols. Additionally, be 
sure they are not attached to plaintext emails or sent via FTP or TFTP. 
 Monitor for vendor end-of-life (EOL) announcements for hardware devices, operating system 
versions, and software, and upgrade as soon as possible. 
 Implement a change management system that anticipates both routine and emergency 
patching. Continuously monitor for vendor vulnerability and patch announcements and 
ensure patches are applied in a timely manner. Ensure use of vendor recommended version 
of the operating system for the features and capabilities required. 
o Test and validate patches as part of the change and patch management processes. 
 As part of a broader password policy, store passwords with secure hashing algorithms. 
Passwords should meet complexity requirements and should be stored using one-way 
hashing algorithms or, if available, unique keys. Follow National Institute of Standards and 
Technologies guidelines when creating password policies.

TLP:CLEAR 
CISA | FBI | NSA | ASD’s ACSC | CCCS | NCSC-NZ 
TLP:CLEAR 
7 
 Require phishing-resistant multi-factor authentication (MFA) for all accounts that access 
company systems, networks, and applications, including sensitive administrative access to 
routers. MFA should use a combination of credentials and a phishing-resistant secondary 
verification method, such as hardware-based PKI or FIDO authentication, to ensure secure 
access and prevent unauthorized entry. 
 As part of a broader identity and access management policy, use local accounts only for 
emergencies and change the passwords after each use. Verify that each use was authorized 
and expected. For everyday management of network infrastructure, use a centralized AAA 
server that supports multi-factor authentication requirements; however, ensure the AAA 
server is not linked to the primary corporate identity store. 
 Limit session token durations and require users to reauthenticate when the session expires. 
Conduct audits to determine the standard session duration for each role to implement 
session expirations. 
 Implement a Role-Based Access Control (RBAC) strategy that assigns users to a specific role 
with defined and inherited permissions to better control and manage what users can do. 
 Remove any unnecessary accounts and periodically review accounts to verify that they 
continue to be needed. Apply the principle of least privilege to make sure accounts only have 
the minimum permissions necessary to  complete their tasks. Additionally, continuously 
monitor accounts in use. 
Cisco-Specific Guidance 
Organizations in the communications sector should be aware that the authoring agencies have 
observed Cisco-specific features often being targeted by, and associated with, these PRC cyber 
threat actors’ activity. To address the risk of exploitation by these specific threat actors, the 
authoring agencies urge organizations to apply the following hardening best practices to all Cisco 
operating systems. For additional information, see Cisco’s IOS XE Hardening Guide and Guide to 
Securing NX-OS Software Devices. 
 Disable Cisco’s Smart Install service using no vstack. 
 If not required, disable the guestshell access using guestshell disable for those 
versions which support the guestshell service. 
 Disable all non-encrypted web management capabilities. If web management is required, 
configure servers in compliance with vendor recommended security settings and software 
images. 
o Always disable the underlying non-encrypted web server using no ip http server. If 
web management is not required, disable all of the underlying web servers using no ip 
http server and no ip http secure-server. 
 Disable telnet and ensure it is not available on any of the VTY lines by configuring all VTY 
stanzas with transport input ssh and transport output none. 
 To securely store passwords on Cisco devices, organizations should: 
o Use Type-8 passwords when possible.

