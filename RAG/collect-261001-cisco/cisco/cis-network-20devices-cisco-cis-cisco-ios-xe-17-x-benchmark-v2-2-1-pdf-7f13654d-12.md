---
id: collect-261001-cisco/cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d-12
title: "cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-cisco/cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d.md
source_anchor: ""
source_lines: [1511, 1689]
sha256: 4f08b221e384923b660164767e17f7ef114276923101b77184e74324df141fa0
---

# cis-network-20devices-cisco-cis-cisco-ios-xe-17-x-benchmark-v2-2-1-pdf-7f13654d

Page 58 
Internal Only - General 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/D_through_E.html#GUID-76805E6F-9E89-4457-
A9DC-5944C8FE5419 
Additional Information: 
limiting http exec-timeout to 10 minutes is useful only if http server or http secure-server, 
at least one, is enable. if HTTP server is disabled (Artifact #2 "HTTP server is disabled") 
and HTTPS server is disabled (Artifact #3 "HTTPS server is disabled"), then this control 
is passed, otherwise it needs to limit exec-timeout to a maximum of 10 minutes (600 
seconds) (Artifact #1 "Max http secure-server limit"). 
CIS Controls: 
Controls Version Control IG 1 IG 2 IG 3 
v8 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped    
v7 0.0 Explicitly Not Mapped 
 Explicitly Not Mapped

Page 59 
Internal Only - General 
1.3 Banner Rules 
Rules in the banner class communicate legal rights to users.

Page 60 
Internal Only - General 
1.3.1 Set the 'banner-text' for 'banner exec' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
This command specifies a message to be displayed when an EXEC process is created 
(a line is activated, or an incoming connection is made to a vty). Follow this command 
with one or more blank spaces and a delimiting character of your choice. Then enter 
one or more lines of text, terminating the message with the second occurrence of the 
delimiting character. 
When a user connects to a router, the message-of-the-day (MOTD) banner appears 
first, followed by the login banner and prompts. After the user logs in to the router, the 
EXEC banner or incoming banner will be displayed, depending on the type of 
connection. For a reverse Telnet login, the incoming banner will be displayed. For all 
other connections, the router will display the EXEC banner. 
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
their networks by using a 'banner-text' for the banner exec command. 
Audit: 
Perform the following to determine if the exec banner is set:

Page 61 
Internal Only - General 
 
hostname#sh running-config | beg banner exec 
If the command does not return a result, the banner is not enabled 
Remediation: 
Configure the EXEC banner presented to a user when accessing the devices enable 
prompt. 
 
hostname(config)#banner exec c 
Enter TEXT message. End with the character 'c'. 
<banner-text> 
c 
Default Value: 
No banner is set by default 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/A_through_B.html#GUID-0DEF5B57-A7D9-
4912-861F-E837C82A3881 
Additional Information: 
The default is no banner. 
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

Page 62 
Internal Only - General 
1.3.2 Set the 'banner-text' for 'banner login' (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
Follow the banner login command with one or more blank spaces and a delimiting 
character of your choice. Then enter one or more lines of text, terminating the message 
with the second occurrence of the delimiting character. 
When a user connects to the router, the message-of-the-day (MOTD) banner (if 
configured) appears first, followed by the login banner and prompts. After the user 
successfully logs in to the router, the EXEC banner or incoming banner will be 
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
their networks by using a 'banner-text' for the banner login command. 
Audit: 
Perform the following to determine if the login banner is set: 
 
hostname#show running-config | beg banner login 
If the command does not return a result, the banner is not enabled.

Page 63 
Internal Only - General 
Remediation: 
Configure the device so a login banner presented to a user attempting to access the 
device. 
 
hostname(config)#banner login c 
Enter TEXT message. End with the character 'c'. 
<banner-text> 
c 
Default Value: 
No banner is set by default 
References: 
1. http://www.cisco.com/en/US/docs/ios-
xml/ios/fundamentals/command/A_through_B.html#GUID-FF0B6890-85B8-
4B6A-90DD-1B7140C5D22F 
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

