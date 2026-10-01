---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-9
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [972, 1071]
sha256: 33af2d02c3b577ee48cab32d145bbca3824465c60cd3d7aaa0b46820491323b2
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

IP-MAC Binding 
According to user configurations, you can bind MAC addresses to IP addresses 
on the USG6000 series. If packets from an IP address do not match the bound 
MAC address, the USG6000 series discards the packets. The USG6000 series 
sends packets that are destined for an IP address to the bound MAC address to 
prevent IP spoofing attacks. 
Dynamic Policy Management — Blacklist 
The USG6000 series blacklists the source IP addresses of untrusted packets and 
discards all the packets of the blacklisted users, therefore effectively preventing 
the attacks from malicious hosts. 
The USG6000 series provides the following blacklist maintenance methods: 
 Manually adding entries to the blacklist to implement proactive defense 
 Automatically adding blacklist entries through attack defense to 
implement intelligent protection 
 Interworking with the whitelist to allow the blacklisted host to access 
some network resources. For example, users are allowed to access the 
Internet using a host even if the host is blacklisted. 
Blacklist is a dynamic policy technology and belongs to the response system. 
The USG6000 series can identify some attack behaviors during dynamic 
running. It controls the traffic of these illegitimate users through the blacklist 
dynamic response system to protect the entire system. 
3.4 Stateful Inspection Based on Flow Sessions 
Kernel Technology Based on Session Management 
The USG6000 series is an advanced stateful firewall based on flow sessions 
and has integrated powerful kernel technology based on session management. 
It provides two core processing units: first-packet processing unit and session 
management unit. They rely on independent acceleration systems in terms of 
management. Such processing has the following advantages: 
 The first-packet processing unit avoids bottleneck of the USG6000 series 
in the processing of the first packet. It enables the USG6000 series to 
provide outstanding performance of new connections per second and 
maintain excellent processing performance on the live network. 
 The session management unit equips the USG6000 series with an 
extraordinary forwarding acceleration system, delivering high forwarding 
performance. The forwarding performance of the USG6000 series for 
subsequent packets relies on the independent acceleration system to 
achieve accelerated packet forwarding, so that the USG6000 series 
delivers high forwarding performance besides the brilliant processing of 
new connections per second. 
 The connection management of the USG6000 series can be at a fine 
granularity. On most firewalls, you can configure policies only over TCP

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  25 
   
 
or UDP in terms of connection management. On the USG6000 series, you 
can configure management policies by service type. For example, you can 
configure management policies for Telnet and HTTP . 
 The service processing of a firewall is based on session management, so 
that the firewall can support abundant services. For example, the 
USG6000 series supports service features such as PBR and QoS. These 
features can be managed on the flow basis. With the flow-based 
forwarding and stateful inspection technologies, the USG6000 series 
provides diversified flow-based services to meet requirements in various 
operating environments. 
In-Depth Inspection 
The USG6000 series provides ASPF that is an advanced communication 
filtering technology to check application-layer protocol information and 
monitor connection-based application-layer protocol status. The USG6000 
series, relying on access control based on packet content, detects and defends 
against some application-layer attacks. It also detects FTP commands, SMTP 
commands, HTTP Java, and ActiveX controls. 
ASPF provides in-depth inspection based on session management. The ASPF 
technology uses information in the session management module to maintain 
session access rules. It saves session status information that cannot be saved by 
static access list rules in the session management module. The session status 
information can be used to intelligently permit or deny packets. When a session 
terminates, the ASPF session management module removes the session 
information from the session table and closes the session on the firewall. 
ASPF intelligently detects the TCP three-way handshake and the connection 
removal handshake. Stateful inspection on the handshake and connection 
removal ensures that a TCP access can normally proceed and the packets of 
incomplete TCP handshake connections are denied directly. 
Advantages of Stateful Inspection 
ACL-based IP packet filtering is applied in common scenarios. This technology 
is simple and inflexible. In many complex scenarios, common packet filtering 
is unable to protect networks. For example, configuring packet filtering rules is 
difficult for multi-channel protocols such as FTP . FTP includes a TCP control 
channel with predefined ports and a TCP data channel that is dynamically 
negotiated. You cannot obtain the port number of the data channel when 
configuring security policies on a packet filtering firewall. Therefore, the 
ingress of the data channel cannot be determined, and security policies cannot 
be accurately configured. The ASPF technology resolves this problem. It 
detects application-layer packet information and dynamically creates and 
deletes temporary rules based on packet content to permit certain packets. 
ASPF enables the USG6000 series to support multiple data connections over 
one control channel. It facilitates security policy configuration in complex 
application scenarios. Many application protocols, such as Telnet and SMTP , 
use standard or well-known ports for communications, but most multimedia 
application protocols such as H.323 and SIP , and other protocols such as FTP 
and NetMeeting use designated ports to initialize a control connection and 
dynamically select ports to transmit data. Port selection is unpredictable. An

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  26 
   
 
