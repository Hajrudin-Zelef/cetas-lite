---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-15
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [1624, 1721]
sha256: 7ccf188914a99b052c296bd850bac0bbb10df3cdb1832784b7cb023a26d4d8d8
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

Static mapping has the following shortcomings: 
1. Static mapping severely wastes public IP addresses, even if it resolves the 
reverse access issue. The NAT technology saves public IP addresses. 
However, public IP addresses cannot be fully used in static mapping 
mode. 
2. Big security problems may occur. An internal server serves only a single 
purpose. For example, the web server provides only HTTP services. This 
server needs to provide access only to port 80. However, the web server 
deployed in static mapping mode enables Internet users to access port 80 
and other ports, which bring about security risks. If a server can be 
maintained only at an intranet host using Telnet, static mapping may 
enable Internet hosts to telnet the server. 
3. Servers with non-standard ports are difficult to deploy. For example, static 
mapping cannot be used to deploy two web servers, one using port 80 and 
the other using port 8080. 
The NAT function of the USG6000 series supports port-level internal servers. 
You can configure internal servers in terms of ports and protocols for internal 
use and that for external use. In the previous example, if the NA T function of 
the USG6000 series is used, 202.38.160.1 can be used as the addresses of the 
web and FTP servers, and URL http://202.38.160.1:8080 can be used to deploy 
the second web server and users can use 202.38.160.1 to access the Internet. 
The USG6000 series provides port-based mapping of internal servers. It can 
provide port-specific services and implement one-to-one mapping of addresses. 
Moreover, each USG6000 series provides the mapping of up to 4096 internal 
servers without affecting access efficiency. 
Server Load Balancing 
Server load balancing can be considered as a static address mapping extension. 
It provides the mapping of one public IP address and multiple private IP 
addresses. Traffic destined for the public IP address is balanced to servers at the 
private IP addresses on the server. Such a function improves server 
performance and reliability. 
Server load balancing also provides port-level mapping to set up the mapping 
between a port of one public IP address and the ports of multiple private IP 
addresses. Based on the mapping relationship, the USG6000 series translates 
destination addresses and ports and forwards packets only when the packets are 
destined for the specified port of the public IP address. 
Server load balancing implements heartbeat detection to check the server 
health. If a server becomes abnormal, subsequent requests are not forwarded to 
the server. 
Perfect Service Support 
NA T has difficulties in processing the packet whose payload contains address 
information. FTP packets are typical examples. The NAT function of the 
USG6000 series supports ICMP redirect, ICMP unreachable, FTP (in passive 
and active modes), H.323, NetMeeting, PPTP , L2TP , DNS, NetBIOS, SIP , 
MGCP , and Skype. Based on available services, the USG6000 series can

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  38 
   
 
provide powerful service support to meet the requirements of most Internet 
services and prevent NAT from becoming a bottleneck in network services. 
To better adapt to the development of network services, the USG6000 series 
provides a customized ALG function. The ALG of some service applications 
can be configured using command lines. Such a function strengthens the 
USG6000 series in its service support and response speed. 
Furthermore, the architecture of the USG6000 series takes into account packet 
encryption and the rapid support for special protocols during address 
translation. Therefore, in terms of the application program gateway, the 
USG6000 series fully considers the program design and architecture. It can 
more quickly respond to user requirements and better support the ever 
changing network services. 
Limitless PAT 
The USG6000 series provides powerful PA T. PAT uses the port information of 
TCP or UDP and applies the "Address+Port" mode to identify connections 
initiated by hosts from the intranet to the Internet during NA T. In this manner, 
PAT enables users on the intranet to share one IP address to access the Internet. 
The TCP or UDP port ranges from 1 to 65535. Ports 1 to 1024 are reserved by 
the system. Theoretically, a public IP address in PAT mode can support about 
60,000 concurrent connections. The USG6000 series provides a 
Huawei-proprietary "unrestricted port" connection algorithm, which ensures 
that one public IP address can support infinite concurrent connections. This 
technology breaks through the upper limit of 65535 ports for Internet access in 
PAT mode, better meets requirements on address translation, and optimizes 
public IP addresses. 
Multi-Interface Load Balancing 
The NAT function of the USG6000 series supports Internet access using 
multiple interfaces in load balancing mode. In actual scenarios, intranet users 
may access the Internet using different interfaces or ISP networks. If address 
translation is not needed, you can configure two default routes on the USG6000 
series to implement load balancing. 
The address translation function of the USG6000 series supports the previous 
load balancing and Internet access using multiple interfaces. The function has 
an excellent effect in the Internet access scenario of a large intranet. 
3.10 Diversified Attack Defense Methods 
Excellent Necessary Capabilities for Defending Against DoS Attacks 
DoS attacks are prevailing on the Internet. A DoS attack is to congest networks 
and interrupt services by sending various junk packets to the target. IP 
communication is connectionless. Attackers take advantage of this feature to 
invent various attack means. Launching a DoS attack is simple, even only a PC 
and a packet sending tool are used. Consequently, DoS attacks prevail on the

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  39 
   
 
