---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-8
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [879, 971]
sha256: 7b0f9081afd86b2dde94d8fd36881eacaba91881af14a7ed976a3c75f90abffa
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

Manageable Security Zones 
Many firewalls in the industry provide independent Trust zones, Untrust zones, 
and DMZs. Such protection model meets most networking requirements. In 
some scenarios that have demanding requirements on security policies, this 
protection model cannot meet the requirements. 
The USG6000 series provides four default security zones: Trust, Untrust, DMZ, 
and Local. It has added the Local zone that defines packets destined for the 
firewall itself, which enhances security protection for the firewall itself. For 
example, by controlling the packets in the Local zone, the USG6000 series 
easily prevents the access initiated from insecure zones using Telnet or FTP . 
The USG6000 series also supports user-defined security zones. Independent 
interfaces can be added to each security zone. 
Policy Control by Security Zone 
The USG6000 series supports the design of security policy groups for the 
access between security zones. Each security policy group supports several 
independent rules. Such a rule system enables easy management of firewall 
policies and facilitates independent management over logical security zones. 
The policy control model based on security zones can clearly define the access 
from the Trust zone to the Untrust zone and from the DMZ to the Untrust zone. 
The model enables the network isolation function of the USG6000 series to 
provide excellent management capabilities. 
Comprehensive Service Capability 
Security zone management of a firewall covers all physical interfaces, 
subinterfaces, Loop Back interfaces, tunnel interfaces, dial-up logic interfaces, 
and virtual-template interfaces. The policies of security zone management 
support all types of services on the firewall. Independent security zone 
management of the USG6000 series isolates network areas accessing through 
VLANs. 
The USG6000 series supports management over the Local zone. You can easily 
define policies to allow external users' access to the USG6000 series itself. By 
defining these policies, you can flexibly set the management rules of the 
USG6000 series. For example, you can permit users in a security zone to log in 
to a firewall and interfaces in a security zone to communicate with the firewall. 
Such an operation manages the firewall itself and distinguishes firewall 
management policies from service flow management policies, helping you 
define clear security policies. 
The security policies of the USG6000 series can be defined on the basis of 
security zones in a centralized manner. For example, the levels of defense 
against DoS attacks may vary with security areas. Through the support of 
services, the policies and control modes of the USG6000 series can cooperate 
well with the security zones. In this way, the USG6000 series provides security 
defense and policy management at the system level, therefore facilitating 
management and implementation of services and policies, and the security 
defense system becomes clearer.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  23 
   
 
3.3 Security Policy Control 
Flexible Rule Setting 
The USG6000 series supports flexible rule settings based on packet 
characteristics. It provides the following functions: 
 Sets rules based on the protocol number of packets. 
 Sets rules based on the source and destination addresses of packets. 
 Uses a wildcard character to define an address range to specify hosts of 
the address range. 
 Sets a source or destination port for UDP or TCP . 
 Sets a port range for the source and destination ports using the methods 
such as greater than, equal to, between, or not equal to. 
 Defines the type and code of ICMP packets and configures a rule for each 
type of ICMP packets. 
 Sets flexible rules based on the ToS field of IP packets. 
 Sets filtering rules based on the user groups and names of Internet access 
users. 
 Sets filtering rules based on application categories and protocols. 
 Sets filtering rules based on locations. 
Rule Management by Time Segment 
ACL policies of the USG6000 series can be managed by time segment. You 
can configure absolute time segments or periodic time segments. You can 
easily configure time-specific policies on the USG6000 series using time 
segments. For example, forbid the use of Skype in working hours and allow the 
use of them in non-working hours. 
ACL-based policies can be configured on the basis of time segments. For 
example, NAT services define policies based on ACLs. Time segments can be 
used to provide more flexible NAT services. QoS defines data flows based on 
ACLs. Time segments can also be used to configure time-specific QoS policies. 
High-Speed Policy Matching 
Policy matching may affect firewall efficiency because each policy consists of 
many rules. 
The USG6000 series uses Huawei-proprietary ACL quick search and matching 
algorithm that enables the USG6000 series to maintain highly efficient 
forwarding when a large number rules exist. When searching thousands of ACL 
rules, the system performance is almost not affected and the processing speed 
remains unchanged. Therefore, high-speed policy matching of the USG6000 
series improves the overall system performance.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  24 
   
 
