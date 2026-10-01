---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-5
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "throughput"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [471, 630]
sha256: 795f7e2d0bf5386fc2950389e194bef5cdca0e47267167d9ad3757f421d26347
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

 Defends against scanning and sniffing attacks. 
 Provides comprehensive and diversified attack defense methods. DoS 
attacks are launched using different means. Therefore, the firewall must 
provide diversified methods to defend against these DoS attacks. 
 Has excellent processing capabilities. An important feature of DoS 
attacks is the sudden increase of network traffic. If a firewall does not 
have excellent processing capabilities, the firewall itself becomes a 
bottleneck when processing the traffic of DoS attacks. Defending against 
the DoS attacks is impossible. A DoS attack is to make the target network 
paralyzed. If network congestion occurs on a key device, the attack 
objective is achieved. 
 Has accurate attack identification capabilities. When processing traffic of 
DoS attacks, many firewalls only ensure that the traffic passing through 
them falls into an acceptable range, but cannot accurately identify attack 
packets. Such processing ensures the normal network traffic and server 
operating, but blocks legitimate users from accessing the Internet. The 
network plane is normal, but services of the legitimate users are denied. 
Therefore, the firewalls still fail to defend against DoS attacks. 
2.12 Networking Adaptability 
Because of the complex network deployment, the firewall must provide 
excellent networking adaptability for constructing service networks flexibly. 
Excellent networking adaptability includes the following aspects: 
 Supports abundant interfaces to meet the requirements on networking 
adaptability at the physical connection layer. 
 Supports routing protocols. Most firewalls support static routing 
protocols, but not dynamic routing protocols. However, supporting 
dynamic routing protocols can effectively improve the networking 
adaptability of a firewall. 
 Supports the transparent mode. The transparent mode helps a firewall to 
work in Layer 2 mode. Therefore, when you add the firewall to a 
network, the existing network topology is not affected. 
 Supports various virtual interfaces, such as VLAN sub-interfaces and 
tunnel interfaces. A firewall provides limited physical interfaces. To 
adapt to more complex networking schemes, the firewall must support 
various virtual interfaces. 
2.13 VPN Service 
Firewalls are deployed at the network border of an enterprise. The firewalls, 
with powerful control capabilities, can provide VPN services to ensure the 
communication between the headquarters and branch offices. 
IP VPN is a common VPN technology, including the IPSec VPN, L2TP VPN, 
and GRE VPN. The IP VPN technology, applied at network borders, enables 
remote users and mobile users to securely and efficiently access the intranet. 
The firewall provides the following VPN services:

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  14    
 
 Provides VPN services to enable the communication among branch 
offices. IPSec tunnels are used to provide secure and reliable VPN 
services. 
 Provides VPN access services for mobile office employees. The firewall 
must support Layer 2 VPN protocols. The widely applied Layer 2 VPN 
protocol is L2TP . L2TP provides VPN services and enables employees 
on the move to securely access the intranet using accounts and 
passwords. 
 Provides efficient encryption services. 
 Supports comprehensive VPN protocols, including GRE, IPSec, and 
L2TP . 
 Strictly complies with RFC and protocol standards to interwork with the 
VPN devices of other vendors. 
2.14 Management System 
The firewall management system must have the following features: 
 User-friendly man-machine interface. Users can manage a firewall 
through diversified methods. 
 Easy upgrade methods, such as online upgrade using hot patches 
 Graphical management that allows convenient configuration and policy 
management 
 Remote maintenance and monitoring 
 Secure and reliable remote logins, such as, the remote login using SSH 
2.15 Log System 
System logs enable after-the-event audit. A firewall logs various operations 
and attacks and provides the log query and filtering means to facilitate search 
and analysis.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  15 
   
 
3 Technical Features of HUAWEI 
Secospace USG6000 Series 
3.1 High Reliability Design 
HUAWEI Secospace USG6000 series uses the carrier-class hardware system 
and dedicated software system (Huawei-proprietary VRP) to provide high 
security and reliability and effectively resolve the conflicts between high 
performance and complex service processing. With the highly reliable 
hardware design, robust software system, hot standby, link backup, and hot 
backup, the USG6000 series ensures high network reliability. 
NG_Security Hardware Platform of NGFW 
NG_Security hardware platform is a next-generation high-performance 
hardware platform developed by Huawei for security products. This platform, 
with a "Multi-core MIPS+Hardware co-processor acceleration+High-speed 
Switch Fabric" architecture, uses a high-speed bus to implement the 
communications between the multi-core CPU and the service processing and 
interface expansion modules. The redundancy design of the platform improves 
hardware reliability. In addition, the NG_Security hardware platform has 
enhanced performance and functionality and expands storage to meet 
requirements on the local storage of security device logs.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  16 
   
 
Figure 3-1 Huawei NG_Security hardware architecture 
 
 
 Multi-core MIPS CPU 
Huawei NG_Security hardware platform uses a 64-bit multi-core MIPS 
architecture that has high performance and is based on the regularly encoded 
instruction set of a fixed length. The MIPS architecture provides streamlined 
instruction sets, hierarchical design of the instruction and high-speed data 
cache, concurrent multi-level flow lines, and dedicated high-speed interfaces 
and DMA capabilities for traffic throughput, and incorporates Huawei 
carrier-class embedded real-time operating system to ensure the high 
performance of the NGFW platform. 
CPU Core1
OS Layer OS Layer
VRP Forwarding 
subsystem
VPN 
subsystem
CPU
OS
Application
 Core 2   …   Core n
UTM 
subsystem ...
 
The NG_Security hardware platform can be expanded with a CPU processing 
unit to implement "1+1" CPU capabilities. Each CPU is a multi-core MIPS 
processor. Such expansion doubles hardware processing capabilities.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  17 
   
 
Figure 3-2 Huawei NG_Security software platform 
AV
IPS
DLPData
Action
Head
….
HTTP
Facebook
Single engine 
concurrent 
processing
Packet reassembly 
Single resolution
Application 
identification
Protocol decoding
Parallel matching
IAE: Intelligence Aware Engine
 
