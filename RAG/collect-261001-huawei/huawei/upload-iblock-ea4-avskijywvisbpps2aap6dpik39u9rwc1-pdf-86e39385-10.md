---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-10
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [1072, 1228]
sha256: 7e1fb936a226ed89453f61306d16982864f05856011e26e1aa7daa5b237c1bbe
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

application may use more than one port at a time. Therefore, packet filtering 
firewall prevents only single-channel transmission of applications and blocks 
the applications using fixed ports, which brings about many security risks. 
ASPF listens to the port used by each connection of an application, opens an 
appropriate path to permit data of a session, and closes this path at the end of 
the session. In this way, the USG6000 series effectively implements access 
control over the applications using dynamic ports. 
When a packet reaches the USG6000 series, ASPF matches the packet with 
access rules. If a match is found, the packet can pass through the USG6000 
series. Otherwise, the packet is discarded. If a packet is used to open a control 
or data connection, ASPF dynamically modifies access rules. The returned 
packets can pass through the USG6000 series only after matching an access 
rule. When processing the returned packets, ASPF also updates the status 
information table. After a connection is closed or timed out, ASPF deletes the 
status information table of the connection, preventing unauthorized packets 
from passing through the USG6000 series. 
3.5 ACTUAL Awareness 
Networks are evolving into next-generation networks that feature explosive 
information growth, borderless network, mobile Internet, and Web2.0. 
Cybercriminals can easily penetrate a traditional firewall that uses quintuple 
ACLs by spoofing or using Trojan horses, malware, or botnets. Under this 
background, the USG6000 series of Huawei provides an "ACTUAL" 
(Application, Content, Time, User, Attack, and Location) awareness technology 
to accurately control network traffic in a refined manner, defend against 
security threats, and ensure intranet security. 
Definition 
ACTUAL awareness is the capability of identifying network traffic by 
application, content, time, user, attack, and location. Based on the ACTUAL 
awareness results, you can configure security policies such as the filtering, 
route selection, traffic control, and NAT policies. 
A pplication
L ocation
T ime
A ttack
C ontent
User
Network environment ACTUAL awareness system
Cloud service
OA service
Mobile user
Service environment
Office userApplication
Content Threat
10101100
01010100
00101011
10101001
11101101
10101101
11100100
10101010
00011101
00111110
Apps

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  27 
   
 
As shown in the previous figure, network traffic is complex. The administrator 
of a traditional firewall cannot accurately analyze or obtain real service traffic 
types, and cannot apply security policies to control network traffic. ACTUAL 
awareness of the USG6000 series analyzes the traffic of complex network 
environments, provides the administrator visibility into statistics on traffic by 
application, content, time, user, attack, and location, and helps the 
administrator configure security policies in a refined manner. 
Application Awareness 
The application awareness module identifies unknown traffic and packet 
formats, extracts the signature, payload length, content or length change rule, 
IP address, and port of a packet, and incorporates statistics on and relationship 
of packets to accurately categorize applications of the traffic. 
Huawei cloud security competence center, by virtue of its experience and 
expertise, provides an application signature database that covers more than 
6000 applications. The USG6000 series can use the application identification 
engine and online update of the signature database to identify and track the 
latest applications. 
 Application-specific security control 
The USG6000 series implements application-specific security control to 
categorize traffic in a fine-granular manner and accurately control the traffic. 
For example, the USG6000 series permits HTTP traffic and denies traffic of 
WebThunder. 
Based on unified policies, application-specific security policies has integrated 
the application dimension. The policy meaning and configuration mode remain 
unchanged. 
 Application-specific traffic management and control 
The USG6000 series implements application-specific traffic management and 
control to categorize traffic in a fine-granular manner and accurately control 
the traffic. For example, the USG6000 series limits the bandwidth of P2P 
applications to guarantee bandwidth for internal applications. 
Based on bandwidth policies, application-specific security policies has 
integrated the application dimension. The policy meaning and configuration 
mode remain unchanged. 
Traffic management and control of the USG6000 series support the guaranteed 
bandwidth and maximum bandwidth. The guaranteed bandwidth specifies the 
minimum bandwidth resources for key services to prevent other services from 
occupying too much bandwidth. The maximum bandwidth specifies the 
maximum bandwidth resources of some services to prevent the impact on other 
services. 
 Application-specific PBR 
The USG6000 series implements application-specific PBR to apply different 
route selection policies by application. For example, the USG6000 series 
selects a reliable and low-delay link for key information system applications of 
enterprises and other links for P2P applications.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  28 
   
 
Based on PBR policies, application-specific security policies has integrated the 
application dimension. The policy meaning and configuration mode remain 
unchanged. 
Content Awareness 
The USG6000 series analyzes application protocols to obtain the content 
transmitted by the application protocols and applies security policies by 
content. 
The content awareness module consists of a protocol decoding module and a 
content matching module. The protocol decoding module categorizes incoming 
packets by protocol, obtains information based on the category, decompresses 
and unpacks the obtained files, identifies real file types, sorts the obtained 
URLs, and sends them to the content matching module. The content matching 
module matches traffic information with virus signatures, intrusion rule 
signatures, sensitive information, and email contents and determines whether 
the traffic triggers security policies based on the matching results. The 
signatures can be updated on the cloud, or traffic information can be sent to the 
cloud for detection, which ensures the up-to-date and effectiveness of 
signatures. 
Time Awareness 
The USG6000 series implements time awareness based on the following 
technologies: 
 Automatic clock synchronization 
The USG6000 series uses Network Time Protocol (NTP) to obtain standard 
network time and adjust the local clock. 
 Automatic conversion of DST 
Some countries and regions use the DST system. The USG6000 series sets the 
DST clock based on the VRP , and the device clock is automatically switched 
with the DST clock. 
 Time-specific security policy 
The USG6000 series has integrated the time or time segment into the security, 
traffic control, and authentication policies as a matching condition and updates 
the policies by time or time segment to implement time-specific control. You 
can configure the USG6000 series to apply traffic control policies on network 
traffic by time segment. 
User Awareness 
Enterprise networks become borderless with increasing mobile office 
employees whose IP addresses are dynamically changed. The security policies 
of traditional firewalls are based on IP configurations, which cannot meet 
requirements on security management and control. How to accurately identify 
users and effectively manage and control user behaviors has become a top issue 
of network security.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  29 
   
 
