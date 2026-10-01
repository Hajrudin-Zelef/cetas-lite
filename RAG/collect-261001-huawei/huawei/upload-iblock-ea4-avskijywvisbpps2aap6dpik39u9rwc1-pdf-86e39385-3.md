---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-3
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "copyright", "throughput", "voice"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [239, 335]
sha256: a25f400817e1b7579b8746bcf978c3948f2aef59c2801c94e5a398ab6f404862
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

2 Technical Features of NGFWs 
2.1 Reliability Design 
A firewall is the key network device deployed at the network egress. The 
firewall requires high reliability because of its location and functions. 
The high reliability is implemented on the basis of the following technologies: 
 Reliable hardware design. Different from personal or household systems, 
network devices are required to work 24 hours without any interruptions. 
This is demanding for hardware components, such as the main board, 
CPU, fans, and cards. To ensure uninterrupted operating for a long time, 
the firewall must have an excellent hardware structure. 
 Hot standby technology. To ensure the reliable operating at a key 
location, the firewall must provide hot standby. Hot standby requires two 
independent devices of the same model to work together to provide a 
more reliable working environment. Two devices deployed in hot 
standby mode can work in either of the following modes. Only one of the 
two devices is working, and the other device takes over services if one 
device fails. Two devices are working. If one device fails, the other 
device takes over all services. 
 Link backup technology. Link backup prevents physical link faults from 
interrupting services. Link backup is implemented as follows: Two links 
are used to carry services. When both links are normal, service traffic 
may select links in load balancing mode. When one link fails, service 
traffic of that link is automatically switched to the other link. To 
implement link backup, the firewall must support various routing 
protocols and provide route management functions. The route-based link 
backup technology can well suit different scenarios and provide more 
reliable services by implementing the mutual backup of links. 
 Hot backup technology. Hot backup means that services are not affected 
during the device or link switchover when a fault occurs. If the backup 
occurs when services are interrupted due to a fault, such backup 
mechanism is called cold or warm backup. In most documents, hot 
backup, warm backup, and cold backup are not strictly distinguished. 
Many vendors advertise their hot backup concepts, but most their backup 
mechanisms are cold or warm backup. More dynamic information has 
more complex hot backup mechanism. Each firewall maintains large 
amounts of rule and connection data. The hot backup mechanism of 
firewalls is complex. Therefore, you must distinguish hot backup from 
cold backup when choosing firewall backup technologies. 
The reliability design of firewalls reflects comprehensive considerations. 
Firewalls are important network devices that have demanding requirements on 
reliability. Therefore, you must consider the reliability design during firewall 
selection.

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  9    
 
2.2 Performance Model 
This section describes the indexes that you must pay attention to when 
measuring firewall performance. 
Throughput is a key index to evaluate firewall performance in the industry. 
Throughput refers to the total traffic that a firewall can forward with the best 
effort in the case of large packets, in bit per second (bit/s). However, the 
throughput does not reflect the actual working capabilities of the firewall, and 
using the throughput as the only performance index is one-sided. 
In addition to the throughput, you must consider the following indexes: 
1. Small-packet forwarding capability 
In the industry, large packets of 1 KB to 1.5 KB are used to measure the 
processing capability of a firewall. Since network traffic mainly comprises 
200-byte packets, the capability of forwarding small packets must be assessed. 
This performance reflects the actual forwarding capacity of the firewall on the 
live network. 
2. Impacts on forwarding efficiency by rule quantity 
A firewall is generally running with a large number of rules. The 
implementation of rules and services may affect the forwarding performance. 
Therefore, you must pay attention to the forwarding efficiency of a firewall in 
the scenarios where massive rules and services exist to avoid performance 
deterioration. 
3. Number of new connections per second 
The index is the number of TCP connections that can be established on a 
firewall per second. Connections are dynamically established on the basis the 
communication status of two parties. Each session must establish a connection 
on the firewall before data exchange. If the firewall has a low connection 
setup rate, the communication delay is long on clients. The larger the 
specification, the higher the forwarding rate, the stronger the status backup 
capability, and the more powerful the attack defense capability. The number of 
new connections per second is an important index to measure firewall 
functionality. If this index is low, the firewall cannot present excellent 
performance in actual network environments and even cannot work under 
DoS attacks. 
4. Number of concurrent connections 
A firewall processes packets based on connections. The index the maximum 
number of connections supported by the firewall. Each connection is 
TCP/UDP access. 
5. Delay 
Delay is the time for transmitting data in the case of no packet loss. The delay 
must be as short as possible. The delay is critical in the scenarios that require 
high timeliness, such as voice and video services. The long delay of a firewall 
results in harmonic distortion and service interruption. Therefore, the delay is 
a key index of firewall performance.

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  10    
 
