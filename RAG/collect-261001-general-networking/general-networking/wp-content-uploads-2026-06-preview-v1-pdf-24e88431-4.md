---
id: collect-261001-general-networking/general-networking/wp-content-uploads-2026-06-preview-v1-pdf-24e88431-4
title: "wp-content-uploads-2026-06-preview-v1-pdf-24e88431"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-general-networking/wp-content-uploads-2026-06-preview-v1-pdf-24e88431.md
source_anchor: ""
source_lines: [180, 383]
sha256: e7a8cb19258d432194dfd4a93fb6eac881a70b4672e66403946bf7ed5fdfa0a3
---

# wp-content-uploads-2026-06-preview-v1-pdf-24e88431

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 1-6 
 
The VPN Lab is based on the FortiGate Lab topology. The only difference is the addition of a small branch 
network and a single IP address change, introduced at the beginning of the VPN Lab section in IP addressing 
table. 
The HA (High Availability) and SD-WAN labs have different structures and are implemented in different labs. 
Each lab uses its own topology, and the IP addressing is configured separately for each lab. In the last section, 
the VDOM (Virtual Domain) feature is explained. Since, VDOM cannot be implemented in the lab environment, 
only the related CLI commands are mentioned. This topic is not included in the FortiGate Administrator 
syllabus and is added to this document as a bonus section.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-7 
 
 
In FortiGate LAB, main topics are described, configured, and tested step-by-step in the E VE-NG environment. 
The LAB begins with the definition of the IP addressing scheme, followed by a detai led explanation of the lab 
structure and setup process. 
After the lab environment is fully built, the FortiGate configuration is performed from factory default settings. The 
lab then proceeds through the remaining topics in a structured order. Each se ction includes verification steps 
a n d  s c r e e n s h o t s  t o  c l e a r l y  i l l u s t r a t e  t h e  c o n f i g u r a t i o n  a n d  r e s u l t s .

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-8 
 
2.1 Topology 
 
2.2 IP Addressing 
CI OS Image Interface Zone 
IPv4 
Address/Prefix 
Sw-core 
x86_64_crb_linux_l2-adventerprisek9-
ms.bin 
E0/0 LAN 192.168.22.0/30 
E0/1 LAN 192.168.33.0/24 
E0/2 LAN 192.168.44.0/24 
Sw1 E0/0 LAN Trunk

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-9 
 
x86_64_crb_linux_l2-adventerprisek9-
ms.bin 
E0/1 LAN 192.168.33.0/24 
Sw2 
x86_64_crb_linux_l2-adventerprisek9-
ms.bin 
E0/0 LAN Trunk 
E0/1 LAN 192.168.44.0/24 
Sw-dmz 
x86_64_crb_linux_l2-adventerprisek9-
ms.bin 
E0/0 DMZ 172.16.1.0/24 
E0/1 DMZ 172.16.2.0/24 
po/1 DMZ Trunk 
Sw-dc 
x86_64_crb_linux_l2-adventerprisek9-
ms.bin 
E0/0 
DATA 
CENTER 
10.10.11.0/24 
E0/1 
DATA 
CENTER 
10.10.10.0/24 
po/1 
DATA 
CENTER 
trunk 
2.3 LAB Topology Setup 
In this part, a multi-zone enterprise environment is built with VLAN segmentation for security and efficient traffic 
management. In the following sections, lab setup explained, starting with the DMZ and continu ing through the 
Data Center, LAN zone, Out-of-Band (OOB) management environment and Core Infrastr ucture. Each 
component is configured step by step, and all zones are interconnected through switches, with a Fort iGate 
firewall acting as the central security gateway for internal traffic control and external connectivity. 
2.3.1 DMZ Zone 
This zone, hosts web servers that are exposed to external traffic. It consists of Web-Server1 and Web-Server2 
and a switch that is connecting these servers to the main firewall. 
DMZ Subnets PC IP Address Gateway IP Address Servers 
VLAN 161 172.16.1.0/24 172.16.1.100 172.16.1.1 Blue-Web Server 
VLAN 162 172.16.2.0/24 172.16.2.100 172.16.2.1 Green-Web Server 
Copy the initial configuration and paste it into global configuration mode on SW-DMZ switch.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-10 
 
configure terminal 
hostname SW-DMZ 
interface Ethernet0/0 
switchport access vlan 161 
switchport mode access 
interface Ethernet0/1 
switchport access vlan 162 
switchport mode access 
interface ethernet0/2 
switchport trunk encapsulation dot1q 
 switchport trunk allowed vlan 161,162 
 switchport mode trunk 
interface ethernet0/3 
switchport trunk encapsulation dot1q 
 switchport trunk allowed vlan 161,162 
 switchport mode trunk 
end 
wr 
2.3.2 Data Center Zone 
This zone hosts essential services and servers within the internal network. Components are, NTP Server that 
provides time synchronization in VLAN101, a server as a domain controller in VLAN 100 and a switch that is 
connecting servers in this zone to the core infrastructure. 
Data Center Subnets PC IP Address Gateway IP Address Servers 
VLAN 101 10.10.11.0/24 10.10.11.100 10.10.11.1 NTP 
VLAN 100 10.10.10.0/24 10.10.10.100 10.10.10.1 Active Directory & CA 
Copy the initial configuration and paste it into global configuration mode on SW-DC switch. 
configure terminal 
hostname SW-DC 
interface ethernet0/0 
switchport mode access

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-11 
 
switchport access vlan 101 
interface ethernet0/1 
switchport mode access 
switchport access vlan 100 
interface ethernet0/2 
switchport trunk encapsulation dot1q 
 switchport trunk allowed vlan 100,101 
 switchport mode trunk 
interface ethernet0/3 
switchport trunk encapsulation dot1q 
 switchport trunk allowed vlan 100,101 
 switchport mode trunk 
end 
wr 
2.3.3 LAN Zone 
The LAN Zone is the local network for client access, comprising SW1, SW2 and SW-Core. 
LAN Subnets PC IP Address Gateway IP Address Clients 
VLAN 33 192.168.33.0/24 192.168.33.100 192.168.33.1 Client1 
VLAN 44 192.168.44.0/24 192.168.44.100 192.168.44.1 Client2 
Copy the initial configuration and paste it into global configuration mode on SW1. 
Configure terminal 
interface ethernet0/0 
switchport mode trunk 
switchport trunk encapsulation dot1q 
interface ethernet0/1 
switchport mode access 
switchport access vlan 33 
end 
wr 
Copy the initial configuration and paste it into global configuration mode on SW2.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-12 
 
Configure terminal 
interface ethernet0/0 
switchport mode trunk 
switchport trunk encapsulation dot1q 
interface ethernet0/1 
switchport mode access 
switchport access vlan 44 
end 
wr 
The core switch is interconnecting sw1 and sw2 switches. Copy the initial confi guration and paste it into global 
configuration mode on SW-core switch. 
configure terminal 
int ethernet0/0 
switchport mode access 
switchport access vlan 22 
int ethernet0/1 
switchport mode trunk 
switchport trunk encapsulation dot1q 
int ethernet0/2 
switchport mode trunk 
switchport trunk encapsulation dot1q 
int vlan 22 
ip address 192.168.22.2 255.255.255.252 
vlan 33 
int vlan 33 
ip address 192.168.33.1 255.255.255.0 
vlan 44 
int vlan 44 
ip address 192.168.44.1 255.255.255.252 
end 
wr 
2.3.4 OOB 
This zone is dedicated to Out-of-Band (OOB) management and includes a Wind ows system in the 
172.20.20.0/24 subnet. It is connected to the network through a standalone switch (SW-OOB). No configuration 
is required for the switch. FW-Core is connected to SW-OOB using port 1.

FortiGate Workbook 
 
COPYRIGHT SMENODE LABS. ALL RIGHTS RESERVED. 2-13 
 
