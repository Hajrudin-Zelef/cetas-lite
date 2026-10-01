---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-13
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [1690, 1848]
sha256: 2de581b7f781e162b5e4acd9570d2e14b297b8ba7c9cce3d013b42a1722952c5
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 64 
Internal Only - General 
1.3.3 Set the 'banner-text' for 'banner motd' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
This MOTD banner is displayed to all terminals connected and is useful for sending 
messages that affect all users (such as impending system shutdowns). Use the no 
exec-banner or no motd-banner command to disable the MOTD banner on a line. The 
no exec-banner command also disables the EXEC banner on the line. 
When a user connects to the router, the MOTD banner appears before the login prompt. 
After the user logs in to the router, the EXEC banner or incoming banner will be 
displayed, depending on the type of connection. For a reverse Telnet login, the 
incoming banner will be displayed. For all other connections, the router will display the 
EXEC banner. 
Rationale: 
"Network banners are electronic messages that provide notice of legal rights to users of 
computer networks. From a legal standpoint, banners have four primary functions. 
• First, banners may be used to generate consent to real-time monitoring under 
Title III. 
• Second, banners may be used to generate consent to the retrieval of stored files 
and records pursuant to ECPA. 
• Third, in the case of government networks, banners may eliminate any Fourth 
Amendment "reasonable expectation of privacy" that government employees or 
other users might otherwise retain in their use of the government's network under 
O'Connor v. Ortega, 480 U.S. 709 (1987). 
• Fourth, in the case of a non-government network, banners may establish a 
system administrator's "common authority" to consent to a law enforcement 
search pursuant to United States v. Matlock, 415 U.S. 164 (1974)." (US 
Department of Justice APPENDIX A: Sample Network Banner Language) 
Impact: 
Organizations provide appropriate legal notice(s) and warning(s) to persons accessing 
their networks by using a 'banner-text' for the banner motd command. 
Audit: 
Perform the following to determine if the login banner is set: 
 
hostname#sh running-config | beg banner motd 
If the command does not return a result, the banner is not enabled.

Page 65 
Internal Only - General 
Remediation: 
Configure the message of the day (MOTD) banner presented when a user first connects 
to the device. 
 
hostname(config)#banner motd c 
Enter TEXT message. End with the character 'c'. 
<banner-text> 
c 
Default Value: 
No banner is set by default 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/A_through_B.html#GUID-7416C789-9561-
44FC-BB2A-D8D8AFFB77DD 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
14.1 Establish and Maintain a Security Awareness 
Program 
 Establish and maintain a security awareness program. The purpose of a security 
awareness program is to educate the enterprise’s workforce on how to interact with 
enterprise assets and data in a secure manner. Conduct training at hire and, at a 
minimum, annually. Review and update content annually, or when significant 
enterprise changes occur that could impact this Safeguard. 
● ● ● 
v7 
17.3 Implement a Security Awareness Program 
 Create a security awareness program for all workforce members to complete on 
a regular basis to ensure they understand and exhibit the necessary behaviors and 
skills to help ensure the security of the organization. The organization's security 
awareness program should be communicated in a continuous and engaging 
manner. 
● ● ●

Page 66 
Internal Only - General 
1.3.4 Set the 'banner-text' for 'webauth banner' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
This banner is displayed to all terminals connected and is useful for sending messages 
that affect all users (such as impending system shutdowns). Use the no exec-banner or 
no motd-banner command to disable the banner on a line. The no exec-banner 
command also disables the EXEC banner on the line. 
When a user connects to the router, the MOTD banner appears before the login prompt. 
After the user logs in to the router, the EXEC banner or incoming banner will be 
displayed, depending on the type of connection. For a reverse Telnet login, the 
incoming banner will be displayed. For all other connections, the router will display the 
EXEC banner. 
Rationale: 
"Network banners are electronic messages that provide notice of legal rights to users of 
computer networks. From a legal standpoint, banners have four primary functions. 
• First, banners may be used to generate consent to real-time monitoring under 
Title III. 
• Second, banners may be used to generate consent to the retrieval of stored files 
and records pursuant to ECPA. 
• Third, in the case of government networks, banners may eliminate any Fourth 
Amendment "reasonable expectation of privacy" that government employees or 
other users might otherwise retain in their use of the government's network under 
O'Connor v. Ortega, 480 U.S. 709 (1987). 
• Fourth, in the case of a non-government network, banners may establish a 
system administrator's "common authority" to consent to a law enforcement 
search pursuant to United States v. Matlock, 415 U.S. 164 (1974)." (US 
Department of Justice APPENDIX A: Sample Network Banner Language) 
Impact: 
Organizations provide appropriate legal notice(s) and warning(s) to persons accessing 
their networks by using a 'banner-text' for the banner motd command. 
Audit: 
Perform the following to determine if the login banner is set: 
 
hostname#show ip admission auth-proxy-banner http 
If the command does not return a result, the banner is not enabled.

Page 67 
Internal Only - General 
Remediation: 
Configure the webauth banner presented when a user connects to the device. 
hostname(config)#ip admission auth-proxy-banner http {banner-text | filepath} 
Default Value: 
No banner is set by default 
References: 
1. https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst9500/software/releas
e/16-
9/configuration_guide/sec/b_169_sec_9500_cg/configuring_web_based_authenti
cation.html 
Additional Information: 
Set the 'banner-text' for 'webauth banner' is useful only if http server or http secure-
server, at least one, is enabled. if HTTP server is disabled (Artifact #1 "HTTP server is 
disabled") and HTTPS server is disabled (Artifact #2 "HTTPS server is disabled"), then 
this control is passed, otherwise 'ip admission auth-proxy-banner http' must exist 
(Artifact #3 "check for ip admission auth-proxy-banner http"). 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v8 
14.1 Establish and Maintain a Security Awareness 
Program 
 Establish and maintain a security awareness program. The purpose of a security 
awareness program is to educate the enterprise’s workforce on how to interact with 
enterprise assets and data in a secure manner. Conduct training at hire and, at a 
minimum, annually. Review and update content annually, or when significant 
enterprise changes occur that could impact this Safeguard. 
● ● ● 
v7 
17.3 Implement a Security Awareness Program 
 Create a security awareness program for all workforce members to complete on 
a regular basis to ensure they understand and exhibit the necessary behaviors and 
skills to help ensure the security of the organization. The organization's security 
awareness program should be communicated in a continuous and engaging 
manner. 
● ● ●

Page 68 
Internal Only - General 
1.4 Password Rules 
Rules in the password class enforce secure, local device authentication credentials.

