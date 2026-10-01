---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-1
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["training", "ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [1, 289]
sha256: 688198057f750f826660566bba8f66b6c44e7426086c090d39933fa95f9af7b6
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

Huawei Certification 
 
HCDP-IESN 
 
Implementing Enterprise Switching Networks 
Lab Guide 
 
 
 
 
 
 
 
 
 
 
 
 
 
 
Huawei Technologies Co.,Ltd

)UV_XOMNZj.[G]KO:KINTURUMOKY)U2ZJ'RRXOMNZYXKYKX\KJ
No part of this document may be repr oduced or transmitted in any form or 
by any means without prior written consent of Huawei Technologies Co., Ltd. 
:XGJKSGXQYGTJ6KXSOYYOUTY
 and other Huawei trademarks are trademarks of Huawei Technologies 
Co., Ltd. All other trademarks and trade names mentioned in this document 
the property of their respective holders.
4UZOIK
The information in this document is subject to change without notice. Every 
effort has been made in the preparation of this document to ensure accuracy of 
the contents, but all statements, information, and recommendations in this 
document do not constitute the warranty of any kind, expressed or implied. 
 
 
 
.[G]KO)KXZOLOIGZOUT
.)*6/+94/SVRKSKTZOTM+TZKXVXOYK9]OZINOTM
2GH-[OJK

+JOZOUT


Huawei Certification System 
Relying on its strong technical and professional training system, in accordance  
with different customers at differ ent levels of ICT technology, Huawei 
certification is committed to provide customs with authentic, professional 
certification.
Based on characteristics of ICT technologies and customersÿneeds at different 
levels, Huawei certification provides customers with certification system of four 
levels.  
HCDA (Huawei Certification Datacom Associate) is primary for IP network 
maintenance engineers, and any others who want to build an understanding of 
the IP network. HCDA certification covers the TCP/IP basics, routing, switching 
and other common foundational knowledge of IP networks, together with 
Huawei communications products, versatile routing platform VRP 
characteristics and basic maintenance. 
HCDP-Enterprise (Huawei Certification Datacom Professional-Enterprise) is 
aimed at enterprise-class network ma intenance engineers, network design 
engineers, and any others who want to   grasp in depth routing, switching, 
network adjustment and optimization technologies. HCDP-Enterprise consists 
of IESN (Implementing Enterprise Switch Networks), IERN (Implementing 
Enterprise Routing Networks), and IENP (Improving Enterprise Network 
performance), which includes advanced IPv4 routing and switching technology 
principles,  network security, high av ailability and QoS, as well as the 
configuration of  Huawei products. 
HCIE-Enterprise (Huawei Certified Internetwork Expert-Enterprise) is designed 
to endue engineers with a variety of IP  technologies and proficiency in the 
maintenance, diagnostics and troubl eshooting of Huawei products, which 
equips  engineers with competence in planning, design and optimization of 
large-scale IP networks.

HCIE- 
R&S 
 
UC&C VC Cloud Storage Wireless Transmission Security 
ICT Career Certification 
Expert 
HCNA- 
Design 
HCNP- 
Design 
HCNA(HCDA) 
HCAr 
HCNA-
WLAN 
HCNA- 
UC 
HCNA- 
VC 
HCNA-
Cloud 
HCNA- 
LTE 
HCNA- 
Transmission 
HCNA-
Security 
HCNA- 
CC 
HCNP-Carrier 
(HCDP-Carrier) 
HCNP-
WLAN 
HCNP- 
UC 
HCNP- 
VC 
HCNP- 
Cloud 
HCNP- 
LTE 
HCNP-
Transmission 
HCNP-R&S 
(HCDP) 
HCNP- 
Security 
HCNP-
Storage 
HCNP- 
CC 
HCNA-
Storage 
H
Associate Professional 
HCIE- 
Design 
Proposed Advanced 
relationship 
Necessary advanced 
relationship 
 
HCIE- 
Carrier 
 
HCIE- 
LTE 
HCIE- 
WLAN 
HCIE-
Security 
HCIE-
Transmissio
n 
HCIE- 
CC 
HCIE- 
UC 
HCIE- 
Cloud 
HCIE- 
VC 
HCIE- 
Storage 
Architect 
Routing & Switching WLAN ICT Convergence 
Design

Referenced icon 


  
8U[ZKX 29]OZIN 29]OZIN ,OXK]GRR 4KZIRU[J
+ZNKXTKZROTK  9KXOGRROTK

Lab environment specification 
:NK2GHKT\OXUTSKTZOYY[MMKYZKJHKRU] 

/JKTZOLOKX *K\OIK 59\KXYOUT
8 '8 <KXYOUT<8)96) 
8 '8 <KXYOUT<8)96) 
8 '8 <KXYOUT<8)96) 
8 '8 <KXYOUT<8)96) 
8 '8 <KXYOUT<8)96) 
9 9)+/9 <KXYOUT<8)96)
9 9)+/9 <KXYOUT<8)96)
9 9:6+/') <KXYOUT<8)96)
9 9:6+/') <KXYOUT<8)96)
,= ;9- <KXYOUT<8)96) 
,= ;9- <KXYOUT<8)96) 

HCDP-IESN Content   
 
HC Series HUAWEI TECHNOLOGIES Page1 
 
)54:+4:9
Chapter 1 Implementing VLAN features ............................................................................................. 1 
Lab 1-1 VLAN Configuration ........................................................................................................... 1 
Lab 1-2 MUX VLAN Configuration and GVRP Configuration (Optional) ......................................... 17 
Lab 1-3 Inter-VLAN Communication ............................................................................................. 32 
Lab 1-4 SEP and Smart Link .......................................................................................................... 46 
Chapter 2 STP and SEP ..................................................................................................................... 59 
Lab 2-1 STP, RSTP, and MSTP Configuration ................................................................................. 59 
Lab 2-2 Compatibility Between Multi-Region MSTP and STP (Optional) ....................................... 82 
Chapter 3 Implementing MPLS technologies .................................................................................. 104 
Lab 3-1 MPLS LDP Configuration ................................................................................................ 104

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page1 
 
Chapter 1 Implementing VLAN features 
Lab 1-1 VLAN Configuration 
Learning Objectives 
The objectives of this lab are to learn and understand how to perform the 
following operations: 
/g120 Configure VLANs. 
/g120 Configure Eth-trunks. 
/g120 Use hybrid interfaces. 
Topology 
 
Figure 1-1 VLAN configuration 
Scenario 
Assume that you are a network administrator of a company. The company 
uses an Ethernet that has two switch es. In the preceding figure, R1 and R2 
represent computers on the network and R3 is a server. To optimize the 
network, the transmission speed and reliab ility must be improved for the links 
between S1 and S2. Two VLANs are configured to isolate broadcast domains.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page2 HUAWEI TECHNOLOGIES HC Series 
 
R2 and R3 are in the same VLAN. R1 needs to communicate with R3. 
Tasks 
Step 1 Perform basic configurations and IP addressing. 
Configure IP addresses and subnet masks for all routers. 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R1 
[R1]interface GigabitEthernet 0/0/1 
[R1-GigabitEthernet0/0/1]ip address 10.0.10.1 24 
 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R2 
[R2]interface GigabitEthernet 0/0/1 
[R2-GigabitEthernet0/0/1]ip address 10.0.10.2 24 
 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R3 
[R3]interface GigabitEthernet 0/0/2 
[R3-GigabitEthernet0/0/2]ip address 10.0.10.3 24 
 
Configure names for the two switches,shutdown some unused interfaces. 
<Quidway>system-view  
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S1 
[S1]interface GigabitEthernet 0/0/13 
[S1-GigabitEthernet0/0/13]shutdown 
[S1-GigabitEthernet0/0/13]interface GigabitEthernet 0/0/14 
[S1-GigabitEthernet0/0/14]shutdown 
 
<Quidway>system-view  
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S2

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page3 
 
