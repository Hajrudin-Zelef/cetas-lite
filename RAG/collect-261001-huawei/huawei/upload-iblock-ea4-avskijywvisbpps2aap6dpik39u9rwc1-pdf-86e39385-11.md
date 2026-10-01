---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-11
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [1229, 1333]
sha256: 3c1357995ac5a01439da1b2da1658ed5987982fc5cdd08f6df27c79e1a7a5ee7
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

User awareness of the USG6000 series identifies users of network traffic and 
implements security management and control by user. 
 User authentication and identification 
User identification is the prerequisite of applying differentiated policies on 
users. The user management and control module provides multiple 
authentication modes to meet the requirements of different user types and 
scenarios. 
− Authentication exemption 
Upper executives require high efficiency and authentication exemption. 
However, their activities must be highly secure. You can bind their 
accounts to IP or MAC addresses and configure authentication exemption 
for them. The USG6000 series then exempts upper executives from 
authentication and allows the login only from the bound IP or MAC 
address. 
Some enterprises have guests who may need to access the enterprise 
networks. The guests do not have dedicated accounts and cannot be 
authenticated. Therefore, their network access permissions must be 
controlled. To accommodate this situation, the user management and 
control module automatically creates temporary accounts for the guests 
with their IP addresses as user names. 
− Password authentication 
For common employees, password authentication is applied. 
Users can access the URL of an authentication page before starting service 
access. The USG6000 series supports HTTP and HTTPS authentication. 
You are advised to choose HTTPS authentication to meet high security 
requirements. 
The USG6000 series supports authentication based on user names and 
passwords. It can also interwork with the LDAP , RADIUS, and AD 
authentication servers and send user information to the authentication 
servers. 
In addition, the USG6000 series supports redirected web authentication. 
When an unauthenticated employee accesses HTTP services, the 
USG6000 series redirects the user to an authentication page and prompts 
the user to get authenticated. 
− Single Sign-On (SSO) 
If an AD server with an identity authentication system has been deployed 
on a network, the USG6000 series can interwork with the AD server to 
implement SSO. After identifying that a user is authenticated by the AD 
server, the USG6000 series permits the user without requesting the user 
name and password. 
If a user has used a VPN (such as an L2TP or SSL VPN) for access and 
the USG6000 series has authenticated the user, the USG6000 series 
normalizes the access user and the user whose online behaviors are 
managed to implement SSO and avoid re-authentication. 
− User-initiated authentication and redirected authentication 
User-initiated authentication is an authentication mode that a user logs in 
to the authentication portal page of the USG6000 series for authentication

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  30 
   
 
before accessing network resources. User-initiated authentication supports 
all access methods. 
Redirected authentication is an authentication mode that an 
unauthenticated user accesses network resources and the USG6000 series 
identifies that the user is not authenticated and pushes an authentication 
page to the user. Redirected authentication supports only HTTP access. 
 User-specific management and control policy 
− Online user management and control 
To restrict all online behaviors of some users within a time segment, you 
can lock out the online users. 
You can also force some untrustworthy online users to log out. 
− Policy management and control 
The USG6000 series supports user-specific online behavior management 
that includes application-layer management and control functions, such as 
user-specific and quintuple-based behavior control, user-specific 
application-layer protocol control, user-specific URL access control, mail 
filtering, and file filtering by keyword or type. For example, you can 
forbid instant messaging tools such as Skype during working hours and 
forbid the access to certain game or forum URLs to ensure working 
efficiency. 
The USG6000 series provides user-specific traffic management and 
control and limits the number of concurrent connections by user to 
effectively allocate and manage bandwidth resources. The USG6000 
series can audit and analyze the traffic statistics of users and user groups 
for follow-up optimization. 
The USG6000 series provides reports, such as user-specific traffic 
rankings by category and time 
Users can inherit management and control policies from user groups, and 
the user groups can inherit the policies from parent user groups. 
Attack Awareness 
Attack awareness of the USG6000 series identifies network security events and 
content security events and incorporates the awareness results of attack events, 
attack behaviors, and abnormal traffic into the reports of unified security 
policies and security postures. Attack awareness enables the USG6000 series to 
defend against attack behaviors and provides administrators and CIOs visibility 
into security postures for accurate understanding. 
Huawei security R&D team has sustained accumulation of attack awareness 
technologies as follows: 
 DoS/DDoS detection and defense 
The USG6000 series provides powerful DDoS detection capabilities based on 
the behavior analysis, legitimate traffic identification, feature identification and 
filtering, abnormal traffic baseline learning, dynamic fingerprint identification, 
reverse source detection technologies to detect malformed packet attacks (such 
as Winnuke and Teardrop), scanning and sniffing attacks (such as the IP sweep, 
port scanning, and IP source routing option attacks), and flood or traffic attacks. 
The USG6000 series also incorporates the Netstream and route-based traffic

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  31 
   
 
