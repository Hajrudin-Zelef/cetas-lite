---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-7
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "parameters", "preemption"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [782, 878]
sha256: 32ebda494b83e6c538d6b05fc4b7dce995adbee05767734b012bd9c32d8f0686
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

expanded for the further development of security technologies. All these factors 
endow the technology advance of the USG6000 series. 
Hot Standby 
Hot standby of the USG6000 series means that two independent devices of the 
same model work simultaneously to provide a more reliable operating 
environment. The USG6000 series can work in either of the following modes: 
 Only one of the two devices is working. If one device fails, the other 
device takes over its services. 
 Both devices are working to implement load balancing. If one device fails, 
the other device automatically takes over all tasks. 
Hot Backup 
Hot backup means that services are not affected during the device or link 
switchover when a fault occurs. If the backup occurs when services are 
interrupted due to a fault, such backup mechanism is called cold backup. The 
USG6000 series implements hot backup on firewall configuration and dynamic 
traffic, including filtering rules, connections, dynamic routing information, and 
state machines of application-layer protocols in status check. The more 
dynamic information is, the more complex the hot backup mechanism is. 
Link Backup 
Link backup prevents physical link faults from interrupting services. The 
USG6000 series provides two links to carry services. When the two links are 
normal, traffic may select both links in load balancing. When one link fails, 
traffic of that link automatically fails over to the other link. The USG6000 
series dynamically adjusts routing protocols during the switchover. Therefore, 
the route-based link backup technology of the USG6000 series can well suit 
different scenarios and provide more reliable services based on the mutual 
backup of links. 
BFD 
Bidirectional Forwarding Detection (BFD) quickly identifies communications 
faults between systems and reports the faults to upper-level applications. 
As an independent hello protocol, BFD implements low-overhead and rapid 
fault detection. By interworking with upper-layer protocols, BFD enables them 
to rapidly identify and recover from faults. BFD can interwork with OSPF, 
static routing, Fast ReRoute (FRR), policy-based routing (PBR), and DHCP to 
rapidly identify link faults. 
Advantages of Huawei Firewalls in Reliability 
The hot standby mechanism of HUAWEI Secospace USG6000 series has the 
following advantages: 
 Since Huawei USG6000 series has expanded the Virtual Router 
Redundancy Protocol (VRRP) to the VRRP Group Management Protocol

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  21 
   
 
(VGMP) to control and guarantee the consistency of VRRP , it has 
abundant advantages in LAN application. Because VRRP reliability 
technology is proved to be stable and reliable in LANs and can be 
transparent to users in LANs, the hot standby solution of the USG6000 
series has distinct advantages in LANs or the access points of intranets. 
 The hot standby technology of the USG6000 series is based on HRP , 
which is a quick and efficient hot standby technology developed by 
Huawei. Through the hot standby technology, multiple HRP backup 
channels with different priorities can be configured according to 
live-network traffic. Due to the quick backup of the session table, users' 
applications are not interrupted during the active/standby switchover 
caused by firewall faults. 
 The hot standby technology of the USG6000 series supports preemption, 
which is important for the networking in which devices back up each 
other to share traffic. Since all the traffic is switched over to one firewall 
once the other is faulty, a practical mechanism is needed to ensure that the 
traffic can smoothly switch back to the original faulty firewall when it 
recovers. The hot standby technology supporting preemption guarantees 
the smooth switchover and therefore ensures the reliable operating of the 
devices in mutual backup networking. 
 The USG6000 series supports OSPF +VRRP hybrid networking. When a 
fault occurs, the firewalls dynamically adjust OSPF parameters so that the 
traffic can be quickly switched over to the other device. In this way, traffic 
can be smoothly switched back in the event of failure recovery and 
reliable operating of the backup networking is guaranteed. 
 The USG6000 series supports the hot standby solution in hybrid mode, 
ensuring that the service interfaces of the firewalls can back up traffic and 
work in transparent mode without any influence on the existing network 
topology, so that users' services are not interrupted during the switchover 
caused by firewall faults. 
 The USG6000 series supports diversified networking modes, and each 
mode can provide the full redundancy of devices and links, which ensures 
the stable operating of the high reliability network. 
3.2 Flexible Security Zone Management 
Isolation by Security Zone 
Based on security zones, the security isolation design of the USG6000 series 
provides an excellent management model for users in the actual application of 
firewalls. 
The core function of a firewall is network isolation, and the network isolation 
technology does not rely only on interfaces in network division. Network 
topologies vary with actual conditions. Network isolation based on fixed 
interfaces cannot meet requirements on the live network. 
The USG6000 series provides an isolation model based on security zones. Each 
security zone can be added to any interface according to actual conditions, not 
affected by the network topology.

HUAWEI Secospace USG6000 Series Technical White Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  22 
   
 
