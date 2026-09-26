---
id: collect-260926-mikrotik/mikrotik/presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62-2
title: "presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/lot-mikrotik/RouterOS/presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62.md
source_anchor: ""
source_lines: [293, 482]
sha256: 36997a70d5e2571c78e0736ba921c7da5421e1791fbf85334994b803d9b0ad37
---

# presentations-cy15-easy-setup-of-ip-based-capsman-link-failover-cap-monitoring-p-397eee62

CAPsMAN Auto Certificate 
• Accept connections only from CAPs with 
valid certificate 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 35

CAPsMAN Antenna Gain  
 (Country Regulations) 
• Antenna-gain value  
is taken from the  
CAP interface 
• Must be configured on 
AP before you enable 
radio in CAP mode 
Example  
 Antenna -gain: 6dBi 
 EIRP: 30dB 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 36

Simple Routed CAPs Network  
(with redundancy) 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 37

Simple Routed CAPs Network  
(with redundancy) 
CAPsMAN 
 
 
 
 
 
 
 
• Ethernet1: Internet Connection 
• Ethernet2: Connection with CAP1 
• Ethernet3: Connection with CAP2 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 38

Simple Routed CAPs Network  
(with redundancy) 
CAPsMAN 
 
 
 
 
 
 
• Add IP addresses for CAPsMAN <-> CAP1  & CAP3 communication 
– CAPsMAN Ethernet2(to CAP1 Ethernet1): 192.168.100.1/30 
– CAPsMAN Ethernet3(to CAP3 Ethernet2): 192.168.100.14/30 
• Create a Bridge interface and add IP address 192.168.200.1/32 on it 
– From now on known as loopback IP address of CAPsMAN 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 39

Simple Routed CAPs Network  
(with redundancy) 
CAPsMAN 
 
 
 
 
 
 
 
• Enable OSPF routing protocol (add networks) 
– 192.168.100.0/30 
– 192.168.100.12/30 
– 192.168.200.1/32 
 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 40

Simple Routed CAPs Network  
(with redundancy) 
Similarly do on CAP1 
• Add IP addresses  
– CAP1 Ethernet1(to CAPsMAN Ethernet2): 192.168.100.2/30 
– CAP1 Ethernet2(to CAP2 Ethernet1): 192.168.100.5/30 
• Add loopback interface(new bridge) and IP address 
192.168.101.1/32 on it 
• Enable OSPF routing protocol (add networks) 
– 192.168.100.0/30 
– 192.168.100.4/30 
– 192.168.101.1/32 
 
 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 41

Simple Routed CAPs Network  
(with redundancy) 
Similarly do on CAP2 
• Add IP addresses  
– CAP2 Ethernet1(to CAP1 Ethernet2): 192.168.100.6/30 
– CAP2 Ethernet2(to CAP3 Ethernet1): 192.168.100.9/30 
• Add loopback interface(new bridge) and IP address 
192.168.101.2/32 on it 
• Enable OSPF routing protocol (add networks) 
– 192.168.100.4/30 
– 192.168.100.8/30 
– 192.168.101.2/32 
 
 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 42

Simple Routed CAPs Network  
(with redundancy) 
Similarly do on CAP3 
• Add IP addresses  
– CAP3 Ethernet1(to CAP2 Ethernet2): 192.168.100.10/30 
– CAP3 Ethernet2(to CAPsMAN Ethernet3): 192.168.100.13/30 
• Add loopback interface(new bridge) and IP address 
192.168.101.3/32 on it 
• Enable OSPF routing protocol (add networks) 
– 192.168.100.8/30 
– 192.168.100.12/30 
– 192.168.101.3/32 
 
 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 43

Monitor CAPs 
• Get notified when any of your Controlled Access Points goes down 
– Power Supply Failure? 
– Board Failure? 
– Any other reason 
 
 Just use loopback  
address of each CAP  
as “host”  
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 44

Monitor CAPs 
 You can use sms or email tool 
 
Suggestion/idea for  
MikroTik Development Team 
(maybe in CAPsMAN v2) 
 Option to Enable Monitoring  
for IP Managed CAPs(one/a group/all) 
 Every time “Monitoring” is enabled for  
an IP CAP ,a Dynamic rule could be  
created on Netwatch 
• Dynamic Rule will be remove if monitoring option is disabled 
• The same option/tab can configure the Up/Down Scripts 
 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 45

Comments? Questions? 
 
Thank You! 
Enjoy the Rest of the MUM 
 
 
Need Help? A reliable partner? Contact me for: 
o MikroTik Training Enquiries (Greek & English) 
o Consultancy & Solutions for New or Existing ISP/WISP 
o Telephony-VoIP Solutions (Wholesale or Retail/CallingCard) 
o A custom Network/Telecom service or solution  
 Georgios Argyrides 
 +357-22-030212  
 +1-561-853-0199 
 george@argyrides.gr 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 46

Live Demo!  
MUM CY2015 Larnaca Georgios Argyrides - george@argyrides.gr 47

More Comments? Questions? 
Thank You! 
Enjoy the Rest of the MUM 
Do you like MikroTik? Need to know more?  
Upcoming Public Trainings:                 In cooperation with 
      Mikrotik Greece Distributor 
o MTCNA & Introduction : June 19-21, Greece, Athens 
o MTCNA & Introduction : October 19-21, Greece, Athens 
o MTCTCE : October 23-25, Greece, Athens 
o MTCWE & SXT Workshop : November 02-04, Greece, Athens 
o MTCRE : November 06-07, Greece, Athens 
 
Need Different Place? Different Dates? Private Trainings for your company?  
 
Just Contact me  
 Georgios Argyrides 
 +357-22-030212  
 +1-561-853-0199 
 george@argyrides.gr 
Georgios Argyrides - george@argyrides.gr MUM CY2015 Larnaca 48
