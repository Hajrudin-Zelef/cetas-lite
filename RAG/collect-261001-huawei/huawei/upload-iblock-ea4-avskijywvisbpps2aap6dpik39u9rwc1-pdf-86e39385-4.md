---
id: collect-261001-huawei/huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385-4
title: "upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["copyright", "voice"]
source: docs/RAG/collect-261001-huawei/upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385.md
source_anchor: ""
source_lines: [336, 470]
sha256: bddf3d500ed8936a40504cbb6911a784e01b0f8b035f13bd177967e2fe3d5b39
---

# upload-iblock-ea4-avskijywvisbpps2aap6dpik39u9rwc1-pdf-86e39385

During firewall selection, you may consider other indexes based on actual 
requirements. You must note that a firewall is a data communication device to 
process complex services and the performance indexes are much more in 
amount than any traditional data communication device. The performance 
indexes of a firewall also reflect the comprehensive indexes of the firewall, 
such as the software design and hardware design. Therefore, the performance 
indexes are important during firewall selection. 
2.3 Network Isolation 
The essential function of a firewall is to isolate network areas. The firewall 
isolates the logical networks of common areas and key areas to avoid the 
spread of insecure factors. Network isolation is an important feature in the 
firewall technology system. Security policies can be effectively implemented 
only after you have divided network areas properly. To check whether 
network isolation is correct, examine the following aspects: 
The network isolation system of a firewall has a clear logic structure to make 
the firewall meet requirements in different scenarios. For example, a firewall 
must have a Demilitarized Zone (DMZ). 
Network areas must interwork with physical interfaces during network 
isolation, and the division of network areas cannot rely on physical interfaces 
only. If only physical interfaces are used for network isolation, requirements 
on flexible implementation cannot be met. Network isolation is a logic 
concept and must be flexibly implemented to meet service requirements. 
When you isolate networks, you must consider the implementation of virtual 
interfaces, such as the tunnel, VPN, and VLAN interfaces. Network services 
are ever-changing. VPN isolation and VLAN isolation are widely applied on 
networks. Domain isolation must apply different virtual interfaces with 
services such as VPNs and VLANs. 
You must consider the security of a firewall itself. The firewall is a control 
point of network isolation and must be secure. The firewall security is the 
basis of network security. You must also consider the access to a firewall from 
the network areas that are isolated by the firewall. 
2.4 Access Control 
The access control function of a firewall is important and applies Access 
Control Lists (ACLs). Each ACL defines a series of rules based on packet 
characteristics to control the packets that pass through the firewall. In some 
scenarios, a large number of rules are specified on the firewall. Therefore, the 
rule capacity is a key index of evaluating firewall performance and 
functionality.

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  11    
 
2.5 Flow-based Stateful Inspection Technology 
The ACL-based IP packet filtering technology is widely used in access control. 
This technology is simple and reliable, but lacks flexibility. For the 
communication using multi-channel protocols such FTP , the firewall is 
difficult to configure. FTP includes a TCP control channel with predefined 
ports and a TCP data channel that is dynamically negotiated. You cannot 
obtain the port number of the data channel when configuring security policies 
on a common firewall. Therefore, the ingress of the data channel cannot be 
determined. The stateful inspection technology can resolve such an issue. By 
detecting the status of data packets, the firewall dynamically discovers ports 
to be opened to determine the packets that are allowed to pass through the 
firewall during the communication process. 
The flow-based stateful inspection technology provides high forwarding 
performance. ACL-based packet filtering detects packets one by one. As a 
result, the firewall performance is degraded when there are massive filtering 
rules. Flow-based stateful inspection, however, determines whether a packet is 
allowed to pass through the firewall based on flow information. Such 
processing improves the forwarding performance. 
Mainstream firewalls mainly use the stateful inspection technology. You are 
advised to consider a stateful firewall first. 
2.6 User-based Management and Control 
An NGFW performs security policy control by IP address and user identity. 
The NGFW must monitor the user logins and logouts, and control user 
permissions and assign bandwidths by user or user group. 
2.7 Application-based Management and Control 
An NGFW performs security policy control by port as well as in-depth 
application identification by protocol and carries out application-based 
management and control according to the identification results. 
The NGFW must support continuous updates of pattern files (used for 
identifying applications) to prevent employees from evading firewall 
monitoring by updating applications or using new applications. 
2.8 Application-Layer Intrusion Prevention 
An NGFW defends against traditional network-layer attacks as well as 
application-layer threats. The NGFW must integrate application identification 
and decoding capabilities, identify worms, botnets, and other 
application-based attacks, detect the contents transmitted by applications, and 
perform application-layer content filtering to prevent information leaks and 
illegitimate transmission.

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  12    
 
2.9 Service Support 
A firewall is deployed at the control point of network services. An important 
measure of the network security solution is to find a balance between 
openness and security. Because of technical features, the firewall may affect 
some services when being deployed on a network. To meet the requirements 
on service expansion, you must consider the service support capabilities of a 
firewall as follows: 
 Supports diversified services using flow-based stateful inspection. With 
the growth of network resources and bandwidths, more and more 
services based on broadband applications come into being. You must 
ensure that the flow-based stateful inspection technology supports 
various services. 
 Supports all multimedia services, such as voice and video services based 
on H.323, SIP , and RTSP , that account for a large proportion of 
broadband services. 
 Supports powerful Network Address Translation (NAT) functions. Public 
IPv4 addresses are in great shortage. Therefore, NAT is a must to provide 
services. Because a firewall is deployed at a key position, configuring 
NA T on the firewall is one of the most common services. In addition, 
NA T hides the intranet structure, which effectively ensures intranet 
security. 
 Supports necessary multicast services. 
 Supports various Quality of Service (QoS) measures. 
2.10 NAT 
With the rapid development of the Internet, public IPv4 addresses are 
exhausted. Before IPv6 is applied, NAT is a major technology that resolves 
this problem. 
NA T is proposed to resolve public IP address shortage to enable intranet users 
to access the Internet. NAT protects privacy of the intranet and provides 
Internet users with services such as WWW, FTP , Telnet, SMTP , and POP3. 
NA T functions include forward NAT and reverse NA T. The forward NAT has 
two forms: NAT and Port Address Translation (PAT). 
Because of the deployment position and technical features of a firewall, NAT 
services provided by the firewall are suitable. Therefore, providing 
comprehensive NAT services is a necessary feature of the firewall. 
2.11 Attack Defense 
Attack defense is a key firewall function. The firewall must have the 
following attack defense capabilities: 
 Defends against DoS attacks. 
 Defends against malformed packet attacks and intelligently identifies 
attack packets.

HUAWEI Secospace USG6000 Series Technical White 
Paper  
 
Copyright ©2013  Huawei T echnologies Co., Ltd.  All rights reserved.  13    
 
