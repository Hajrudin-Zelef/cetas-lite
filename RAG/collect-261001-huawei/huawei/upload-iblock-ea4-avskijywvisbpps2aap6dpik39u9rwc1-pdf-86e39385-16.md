---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-16
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [1722, 1820]
sha256: e4220cd7aa4e0c6add75ec7f8e58364d0136e7bdf9967b173d43077e3996a0b1
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

Internet and exert severe impacts on intranets and even backbone networks, 
leading to severe network accidents. Therefore, an excellent anti-DoS 
capability is indispensable to firewalls. 
Almost all firewalls advertise anti-DoS functions. Why do DoS attacks 
frequently break down networks? An excellent anti-DoS system must have the 
following features: 
 Provides comprehensive and diversified attack defense methods. The 
firewall must provide diversified methods to defend against DoS attacks, 
because they are launched using different means. 
 Has excellent processing capabilities. An important feature of DoS attacks 
is the sudden increase of network traffic. If a firewall does not have 
excellent processing capabilities, the firewall itself becomes a bottleneck 
when processing the traffic of DoS attacks. Defending against the DoS 
attacks is impossible. A DoS attack is to make the target network 
paralyzed. If network congestion occurs on a key device, the attack 
objective is achieved. Note that you must consider forwarding 
performance and service processing capabilities of a firewall. In the 
anti-DoS defense process, the number of new connections per second is a 
key index to ensure network connectivity. Attackers randomly change 
source addresses to launch DoS attacks, and all connections are new ones. 
 Has accurate attack identification capabilities. When processing traffic of 
DoS attacks, many firewalls only ensure that the traffic passing through 
them falls into an acceptable range, but cannot accurately identify attack 
packets. Such processing ensures the normal network traffic and server 
operating, but blocks legitimate users from accessing the Internet. The 
network plane is normal, but services of the legitimate users are denied. 
Therefore, the firewalls still fail to defend against DoS attacks. 
The USG6000 series has thoroughly considered all the previous aspects, so it 
has big advantages over other firewalls in anti-DoS performance and 
functionality. 
Diversified Anti-DoS Methods 
The USG6000 series can defend against DoS attacks, such as ICMP flood, 
SYN flood, and UDP flood attacks based on the characteristics of data packets 
and the attack means. The USG6000 series proactively identifies dozens of 
common attack types. Many types of attacks may result in DoS. The USG6000 
series can proactively detect and block illegal attacks to protect the intranet. 
The USG6000 series can be used to set up a secure defense system that has 
various attack defense methods to protect the network from DoS attacks. 
The USG6000 series uses some unique defense technologies according to 
attack features to ensure that it can more specifically defend against DoS 
attacks and provide a complete attack defense feature. 
In addition to careful consideration of attack means, the USG6000 series has 
fully taken into account the usage and network adaptability. The attack defense 
may protect a host or all hosts in a security zone.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  40 
   
 
Advanced TCP Proxy 
The USG6000 series can use TCP proxy to prevent DoS attacks such as SYN 
flood, which may quickly exhaust all server resources and crash the server. The 
common anti-DoS technology cannot accurately identify traffic of legitimate 
users and attack packets when attacks are launched. The USG6000 series uses 
transparent TCP proxy to defend against DoS attacks. It can accurately identify 
attack packets based on precise authentication, allow the normal packets to 
access firewall resources, and discard attack packets directly. 
Some attacks set up complete TCP connections to exhaust server resources. 
The USG6000 series implements an enhanced proxy function. It checks 
whether the client has any data packet to send after the connection with the 
client is established. If yes, the USG6000 series connects to the server. If no, 
the USG6000 series discards the packet from the client. Such a function 
ensures that the USG6000 series can identify the attacks that consume server 
resources even using the complete TCP three-way handshake. 
Defense Against Scanning and Sniffing Attacks 
Scanning and sniffing attacks use the ping sweep (ICMP and TCP) to identify 
the systems on the network, accurately locating potential targets. Alternatively, 
the scanning and sniffing attacks use the TCP and UCP port scanning to detect 
the potential services monitored by the operating system. Through scanning 
and sniffing, attackers can roughly learn about potential security vulnerabilities 
of and service types provided by the target system, preparing for further 
attacks. 
The USG6000 series can flexibly and efficiently detect such scanning and 
sniffing packets using comparative analysis and prevent the subsequent attacks. 
Such scanning and sniffing attacks include address scanning, port scanning, IP 
Source Route attacks, IP Route Record attacks, and network structure sniffing 
through Tracert. 
Malformed Packet Attack Prevention 
The USG6000 series automatically detects attack packets and defends against 
the attacks that utilize malformed packets, including Land, Smurf, Fraggle, 
WinNuke, ICMP Redirect or Unreachable packets, illegitimate TCP packet flag 
bits (such as ACK, SYN, and FIN), Ping of Death, and Teardrop. 
Defense Against Application-Layer DDoS Attacks 
The anti-DDoS function of the USG6000 series defends against IP layer, 
transport layer, and application layer DDoS attacks such as the SIP flood, 
HTTP flood, HTTPS flood, DNS Request flood, and DNS Reply flood attacks. 
The USG6000 series automatically detects DDoS attacks. When a DDoS attack 
is found, the USG6000 series enables the anti-DDoS function to block attack 
traffic and permit normal traffic. 
The anti-DDoS function of the USG6000 series supports threshold 
self-learning to provide references for the attack defense threshold setting and 
improve policy effectiveness. In normal cases, the system collects statistics on

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  41 
   
 
