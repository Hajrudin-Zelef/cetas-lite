---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-22
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [2505, 2609]
sha256: 016253290740d5ffe4bd1f002b18c9c6a954895631c1cb470a3c5609e8d87cba
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  57 
   
 
PC
PC
PC
PC
LAN
LAN
Internet
Headquarters
LAN Switch
LAN Switch
USG 
(Active)
USG 
(Standby)
Router
Router
 
 
 Two USG6000s in the headquarters implement hot standby. One 
USG6000 functions as the active firewall and the other as the standby 
firewall to provide ACL, ASPF, traffic policing, NAT, and other security 
defense functions. 
 The USG6000s are connected by the heartbeat link. 
 The USG6000s are connected to intranet users through LAN switches. 
 The USG6000s are connected to the Internet through routers. 
4.4 VPN Applications Protected by IPSec 
As VPN gateways, the USG6000 series supports the tunneling technologies 
such as L2TP and GRE. The IPSec, unified security gateway, and QoS 
technologies ensure the high quality and security of network information 
transmission. The following figure shows the networking diagram for the 
security VPN application.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  58 
   
 
PC
PC
LAN
Branch office/Partner
Headquarters
LAN
 PC
File server
Carrier 
network
NAS
Home 
office
Mobile office
PSTN/ISDN
Access VPN Internet
VPN tunnel
IPSec
tunnel
 
 
 The Access VPN sets up a secure path for the SOHO and mobile office 
users to access the headquarters resources through PSTN/ISDN networks. 
 Intranet VPN enables branch offices and representative offices to access 
resources at the headquarters. The IPSec and IKE technologies ensures the 
secure transmission of data on the Internet to prevent interception and 
tampering. 
 The extranet VPN enables partners and customers to access the intranet 
and ensures the security of their own networks.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  59 
   
 
4.5 SSL VPN Application 
Figure 4-2 SSL VPN solution 
Partner
Mobile office
Branch office
Remote 
maintenance
Customer
Intranet
Server cluster
Encrypted intranet or Internet 
connection
Standard intranet connection
 
 
 Comprehensive user identity authentication, access authorization, and 
behavior audit to ensure legitimate user identities and implement refined 
access control policies 
 Forcible encryption for the data between remote users and the intranet to 
protect sensitive data and prevent information leaks 
 Support of a wide range of remote access services, including access to 
web resources, file systems, multiple types of C/S applications, and all-IP 
layer services that are irrelevant to applications 
 Access through standard browsers, which means no need to install, 
configure, or maintain clients for users and effectively improves the 
productivity of mobile office personnel, such as employees on the move 
 Excellent log functions, which facilitate the real-time audit and 
management of operation behaviors of administrators and other users
