---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-14
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [1849, 2012]
sha256: 21ddd45c059894201bc4eadacbb48370d1e83c66658f0d5266b71fd6f28d0506
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 69 
Internal Only - General 
1.4.1 Set 'password' for 'enable secret' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Use the enable secret command to provide an additional layer of security over the 
enable password. The enable secret command provides better security by storing the 
enable secret password using a nonreversible cryptographic function. The added layer 
of security encryption provides is useful in environments where the password crosses 
the network or is stored on a TFTP server. 
Rationale: 
In Cisco IOS XE, password types 8 and 9 are considered more secure than older 
password types because they use strong hashing algorithms. 
• Type 8 passwords use PBKDF2 with SHA-256, which is an improvement over 
older MD5-based hashing methods. 
• Type 9 passwords use SCRYPT, which is designed to be memory-intensive, 
making it harder for attackers to brute-force. 
Cisco generally recommends using Type 8 or Type 9 for better security, but Type 9 is 
the default in IOS XE. However, some auditors may prefer Type 8 because Type 9 is 
not NIST-approved, especially in U.S. Defense and Public Sector environments. 
Impact: 
While enable secret in Cisco IOS XE enhances security, it can have some potential 
drawbacks: 
• Password Recovery Complexity: If the password is lost, recovering access can 
be difficult, 
• Performance Overhead: Stronger hashing algorithms (like SCRYPT) can slightly 
increase CPU usage, especially on older hardware. 
• Configuration Management Issues: If not properly documented, administrators 
may struggle with access control changes. 
• Potential Lockout Risks: If centralized authentication fails and the enable secret 
is unknown, network administrators could be locked out. 
• Security Misconfiguration: If weak passwords are used, even with encryption, 
attackers could still brute-force them. 
Audit: 
Perform the following to determine enable secret is set: 
If the command does not return a result, the enable password is not set.

Page 70 
Internal Only - General 
hostname#sh run | incl enable secret  
Remediation: 
Configure a strong, enable secret password. 
hostname(config)#enable secret 9 {ENABLE_SECRET_PASSWORD}  
Default Value: 
No enable secret password setup by default 
References: 
1. https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst9600/software/releas
e/16-
12/configuration_guide/sec/b_1612_sec_9600_cg/controlling_switch_access_wit
h_passwords_and_privilege_levels.html 
Additional Information: 
Note: You cannot recover a lost encrypted password. You must clear NVRAM and set a 
new password. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
5.4 Restrict Administrator Privileges to Dedicated 
Administrator Accounts 
 Restrict administrator privileges to dedicated administrator accounts on 
enterprise assets. Conduct general computing activities, such as internet 
browsing, email, and productivity suite use, from the user’s primary, non-privileged 
account. 
● ● ● 
v7 
4.3 Ensure the Use of Dedicated Administrative Accounts 
 Ensure that all users with administrative account access use a dedicated or 
secondary account for elevated activities. This account should only be used for 
administrative activities and not internet browsing, email, or similar activities. 
● ● ●

Page 71 
Internal Only - General 
1.4.2 Enable 'service password-encryption' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
When password encryption is enabled, the encrypted form of the passwords is 
displayed when a more system:running-config command is entered. 
Rationale: 
This requires passwords to be encrypted in the configuration file to prevent 
unauthorized users from learning the passwords just by reading the configuration. When 
not enabled, many of the device's passwords will be rendered in plain text in the 
configuration file. This service ensures passwords are rendered as encrypted strings 
preventing an attacker from easily determining the configured value. 
Impact: 
Organizations implementing 'service password-encryption' reduce the risk of 
unauthorized users learning clear text passwords to Cisco IOS configuration files. 
However, the algorithm used is not designed to withstand serious analysis and should 
be treated like clear-text. 
Audit: 
Perform the following to determine if a user with an encrypted password is enabled: 
Ensure a result that matches the command return 
 
hostname#sh run | incl service password-encryption 
Remediation: 
Enable password encryption service to protect sensitive access passwords in the device 
configuration. 
 
hostname(config)#service password-encryption 
Default Value: 
Service password encryption is not set by default 
References: 
1. http://www.cisco.com/en/US/docs/ios-xml/ios/security/s1/sec-cr-s1.html#GUID-
CC0E305A-604E-4A74-8A1A-975556CE5871

Page 72 
Internal Only - General 
Additional Information: 
Caution: This command does not provide a high level of network security. If you use this 
command, you should also take additional network security measures. 
Note: You cannot recover a lost encrypted password. You must clear NVRAM and set a 
new password. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
3.11 Encrypt Sensitive Data at Rest 
 Encrypt sensitive data at rest on servers, applications, and databases containing 
sensitive data. Storage-layer encryption, also known as server-side encryption, 
meets the minimum requirement of this Safeguard. Additional encryption methods 
may include application-layer encryption, also known as client-side encryption, 
where access to the data storage device(s) does not permit access to the plain-text 
data.  
 ● ● 
v7 16.4 Encrypt or Hash all Authentication Credentials 
 Encrypt or hash with a salt all authentication credentials when stored.  ● ●

Page 73 
Internal Only - General 
1.4.3 Set 'username secret' for all local users (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Username secret password type 5 and enable secret password type 5 must be migrated 
to the stronger password type 8 or 9. IF a device is upgraded from IOS XE 16.9 or later 
the type 5 is auto converted to type 9. 
The username secret command provides an additional layer of security over the 
username password. 
Rationale: 
Default device configuration does not require strong user authentication potentially 
enabling unfettered access to an attacker that is able to reach the device. Creating a 
local account with an encrypted password enforces login authentication and provides a 
fallback authentication mechanism for configuration in a named method list in a situation 
where centralized authentication, authorization, and accounting services are 
unavailable. 
Impact: 
Organizations implementing 'username secret' across their enterprise reduce the risk of 
unauthorized users gaining access to Cisco IOS devices by applying a MD5 hash and 
encrypting user passwords. 
Audit: 
Perform the following to determine if a user with an encrypted password is enabled: 
If a result does not return with secret, the feature is not enabled 
 
 hostname#show run | incl username 
Remediation: 
Create a local user with an encrypted, complex (not easily guessed) password. 
 
hostname(config)#username {{em}LOCAL_USERNAME{/em}} secret 
{{em}LOCAL_PASSWORD{/em}} 
Default Value: 
No passwords are set by default

