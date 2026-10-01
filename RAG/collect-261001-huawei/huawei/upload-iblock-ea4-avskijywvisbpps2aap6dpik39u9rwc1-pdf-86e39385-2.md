---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-2
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [112, 238]
sha256: 1ed25fd57a3f112919a6dd3dfe812e2b4de7ff6158aa1b6a0ec1bda3a8ba86b5
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

HUAWEI Secospace USG6000 Series Technical 
White Paper 
Keywords: NGFW, HUAWEI Secospace USG6000 series, network security, VPN, tunneling 
technology, L2TP , IPSec, IKE  
Abstract: This document describes the technical features and working mechanisms of the 
Secospace USG6000 series and analyzes technical issues that you should pay attention to during 
firewall selection. 
Acronym and Abbreviation Full Spelling 
VPN Virtual Private Network 
AAA Authentication, Authorization, and 
Accounting 
ASPF Application Specific Packet Filter 
DoS Denial of Service 
L2TP Layer 2 Tunneling Protocol 
IPSec Internet Protocol Security 
IKE Internet Key Exchange 
NGFW Next Generation Firewall

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  5    
 
1 Overview 
1.1 Network Threats and Emergence of 
Next-Generation Firewalls (NGFWs) 
With the rapid Internet development, increasing application quantity, and 
Web2.0 popularity, more bandwidth requirements and new application 
architectures (such as Web2.0) are changing the protocol use and data 
transmission methods. More and more applications use a few ports for 
transmission. New threats (such as worms, botnets, and other 
application-based attacks) continuously come into being. Security threats 
focus on seducing users into installing malicious programs that can evade the 
detection of security devices and software. 
Traditional firewalls identify applications based on ports and protocols and 
perform attack detection and defense based on transport-layer characteristics. 
Security policies based on ports and protocols cannot provide sufficient 
defense capabilities in the scenarios where a large number of applications use 
a few ports or some applications use non-standard ports. The traditional 
firewalls cannot defend against application-based threats, such as worms and 
botnets. 
Ever-changing business processes, enterprise deployment technologies, and 
threats bring about new requirements on network security. New security 
requirements have promoted the emergence of NGFWs. 
1.2 NGFW Definition 
Gartner defines a network firewall as an in-line security control that 
implements network security policy between networks of different trust levels 
in real time. Gartner has used the NGFW term to indicate the necessary 
evolution of a firewall to deal with changes in business processes, IT 
technologies, and network threats. 
An NGFW has at least the following attributes: 
1. Supports online "bump-in-the-wire" configuration without interrupt any 
network connection.

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  6    
 
2. Checks network flows during transmission and implements security 
policies. The features are as follows: 
− Standard first-generation firewall capabilities, including packet 
filtering, Network Address Translation (NA T), stateful protocol 
inspection, and VPN 
− Integrated network intrusion detection: The NGFW supports 
vulnerability-specific and threat-specific feature codes. The 
interaction effect of the IPS and firewall is greater than the sum of 
two separate parts. For example, an NGFW automatically binds 
conditions to apply firewall rules to prevent an address from loading 
malicious traffic to the IPS, but does not require administrators to 
deploy the solution cross consoles. The NGFW has integrated 
powerful IPS engines and feature codes. 
− Application awareness and full-stack visibility: The NGFW identifies 
applications and implements network security policies that are 
independent from ports, protocols, and services. For example, the 
NGFW allows the use of Skype, but disables file sharing in Skype or 
always blocks the GoToMyPC function. 
− Excellent firewall intelligence: The NGFW collects incoming 
information, helps administrators make better decisions, and 
optimizes the deny rule database. For example, the NGFW binds the 
deny action to user identities or set up the address blacklist and 
whitelist. 
1.3 Instructions on Using Firewalls 
Firewalls are deployed at convergence points on the entire network. If traffic 
of the protected network has bypassed a firewall, the security defense function 
of the firewall does not take effect. Therefore, when you use a firewall, ensure 
that all traffic of the protected network passes through the firewall. 
By default, firewall rules deny all access requests. After connecting a firewall 
to the network, you must configure security policies based on actual 
conditions. Policy effectiveness, diversity, and flexibility are important 
indexes to evaluate a firewall. In the complex network environment that has 
massive rules, you must consider the rule capacity and forwarding 
performance of the firewall. 
Firewall security is also an important criterion. Security performance of a 
firewall depends on the operating system and hardware platform. The secure 
operating system guarantees the software security of the firewall, and the 
dedicated hardware platform enables the firewall to run stably for a long time. 
Firewall is a basic network device and must run for a long time without any 
interruption. Therefore, hardware reliability is critical. 
Before using a firewall, you must determine issues to be resolved on the live 
network and select the firewall that meets requirements on performance and 
functionality. The firewall performance and functionality must be balanced. 
However, you must pay special attention to performance indexes, because 
they are critical to the actual operating. If the firewall is poor in performance, 
network congestions and failures frequently occur. Such a network is insecure. 
Performance indexes reflect the available firewall capabilities and the costs

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  7    
 
that an enterprise uses the firewall. The enterprise cannot afford too high costs. 
If a firewall causes long delays, users also experience great losses. 
Mainstream firewalls are the stateful inspection ones that are sensitive to 
service applications. Protocols of multimedia services (such as audio and 
video services) are complex. If the protocol status is improperly processed, 
services may be interrupted or unnecessary ports must be opened to ensure 
service continuity. However, the opening of unnecessary ports degrades 
security. Therefore, you must consider the service adaptability of stateful 
firewalls, so that the firewalls do not adversely affect network services.

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  8    
 
