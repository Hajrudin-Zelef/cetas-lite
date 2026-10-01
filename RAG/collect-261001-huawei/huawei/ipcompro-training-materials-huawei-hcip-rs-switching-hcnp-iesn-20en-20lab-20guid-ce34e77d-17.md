---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-17
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [3531, 3760]
sha256: 0579610b4b2ca6f0abd84f297e5550328cfba36e077ad3ce72068cbdf9c0e146
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

  region-name RG1 
  instance 1 vlan 1 to 10 
  instance 2 vlan 11 to 20 
  active region-configuration 
# 
interface Vlanif1 
 ip address 10.0.1.2 255.255.255.0 
# 
interface GigabitEthernet0/0/9 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
 stp instance 0 port priority 32 
# 
interface GigabitEthernet0/0/10 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
 stp instance 0 port priority 16 
# 
interface GigabitEthernet0/0/23 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
interface GigabitEthernet0/0/24 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
Return 
 
[S3]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S3 
# 
vlan batch 2 to 20 
# 
stp region-configuration 
  region-name RG1 
  instance 1 vlan 1 to 10 
  instance 2 vlan 11 to 20 
  active region-configuration 
# 
interface Ethernet0/0/1 
 port link-type trunk

HCDP-IESN  Chapter 2 STP and SEP 
 
Page80 HUAWEI TECHNOLOGIES HC Series 
 
 port trunk allow-pass vlan 2 to 20 
# 
interface Ethernet0/0/13 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
interface Ethernet0/0/23 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
Return 
 
[S4]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S4 
# 
 vlan batch 2 to 20 
# 
 stp region-configuration 
  region-name RG1 
  instance 1 vlan 1 to 10 
  instance 2 vlan 11 to 20 
  active region-configuration 
# 
interface Ethernet0/0/1 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
interface Ethernet0/0/14 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
# 
interface Ethernet0/0/23 
# 
interface Ethernet0/0/24 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 20 
 stp instance 0 cost 2000000 
# 
Return

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page81

HCDP-IESN  Chapter 2 STP and SEP 
 
Page82 HUAWEI TECHNOLOGIES HC Series 
 
Lab 2-2 Compatibility Between Multi-Region MSTP and STP 
(Optional) 
Learning Objectives 
The objectives of this lab are to learn and understand how to perform the 
following operations: 
/g120 Configure multi-instance Multiple Spanning Tree Protocol (MSTP) and 
multi-region MSTP. 
/g120 Configure MSTP and Spanning Tree Protocol (STP) to be compatible.  
/g120 Configure MSTP edge interface protection, designated interface 
protection, loop protection, and TC-BPDU protection. 
Topology 
 
Figure 2-2 Compatibility between multi-region MSTP and STP 
Scenario 
Assume that you are a network administrator of a company. In the Layer 2 
networking architecture, MSTP is deployed to: 
Avoid sub-optimal routes. 
Resolve the problem that some virtual local area network (VLAN) paths

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page83 
 
are unreachable if a single spanning tree instance is used. 
Implement load balancing. 
Interoperate with traditional spanning trees. 
Tasks 
Step 1 Perform basic configurations. 
Disable unused interfaces before carrying out the experiment. 
<S1>system-view 
Enter system view, return user view with Ctrl+Z. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]shutdown 
 
<S3>system-view 
Enter system view, return user view with Ctrl+Z. 
[S3]interface Ethernet 0/0/23 
[S3-Ethernet0/0/23]shutdown 
 
<S4>system-view 
Enter system view, return user view with Ctrl+Z. 
[S4]interface Ethernet 0/0/14 
[S4-Ethernet0/0/14]shutdown 
 
Create VLANs 3, 4, 5, 6, 7, and 8 on all switches. 
[S1]vlan batch 3 to 8 
 
[S2]vlan batch 3 to 8 
 
[S3]vlan batch 3 to 8 
 
[S4]vlan batch 3 to 8 
 
View information about the created VLANs. 
[S1]display vlan 
* : management-vlan 
--------------------- 
The total number of vlans is : 7 
VLAN ID Type         Status   MAC Learning Broadcast/Multicast/Unicast Property

HCDP-IESN  Chapter 2 STP and SEP 
 
Page84 HUAWEI TECHNOLOGIES HC Series 
 
---------------------------------------------------------------------------- 
1       common       enable   enable       forward   forward   forward default 
3       common       enable   enable       forward   forward   forward default 
4       common       enable   enable       forward   forward   forward default 
5       common       enable   enable       forward   forward   forward default 
6       common       enable   enable       forward   forward   forward default 
7       common       enable   enable       forward   forward   forward default 
8       common       enable   enable       forward   forward   forward default 
 
[S2]display vlan 
* : management-vlan 
--------------------- 
The total number of vlans is : 7 
VLAN ID Type         Status   MAC Learning Broadcast/Multicast/Unicast Property 
---------------------------------------------------------------------------- 
1       common       enable   enable       forward   forward   forward default 
3       common       enable   enable       forward   forward   forward default 
4       common       enable   enable       forward   forward   forward default 
5       common       enable   enable       forward   forward   forward default 
6       common       enable   enable       forward   forward   forward default 
7       common       enable   enable       forward   forward   forward default 
8       common       enable   enable       forward   forward   forward default 
 
[S3]display vlan 
* : management-vlan 
--------------------- 
The total number of vlans is : 7 
VLAN ID Type         Status   MAC Learning Broadcast/Multicast/Unicast Property 
---------------------------------------------------------------------------- 
1       common       enable   enable       forward   forward   forward default 
3       common       enable   enable       forward   forward   forward default 
4       common       enable   enable       forward   forward   forward default 
5       common       enable   enable       forward   forward   forward default 
6       common       enable   enable       forward   forward   forward default 
7       common       enable   enable       forward   forward   forward default 
8       common       enable   enable       forward   forward   forward default 
 
[S4]display vlan 
* : management-vlan 
--------------------- 
The total number of vlans is : 7 
VLAN ID Type         Status   MAC Learning Broadcast/Multicast/Unicast Property 
----------------------------------------------------------------------------

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page85 
 
1       common       enable   enable       forward   forward   forward default 
3       common       enable   enable       forward   forward   forward default 
4       common       enable   enable       forward   forward   forward default 
5       common       enable   enable       forward   forward   forward default 
6       common       enable   enable       forward   forward   forward default 
7       common       enable   enable       forward   forward   forward default 
8       common       enable   enable       forward   forward   forward default 
 
Set the type to Trunk for the links between switches. Enable the links to 
receive bridge protocol data unit (BDPU) packets and to allow the access from 
all VLANs. Note: The direct link between S2 and S3 remains unchanged. 
[S1]interface GigabitEthernet 0/0/13 
[S1-GigabitEthernet0/0/13]port link-type trunk 
[S1-GigabitEthernet0/0/13]port trunk allow-pass vlan all 
[S1-GigabitEthernet0/0/13]bpdu enable 
[S1-GigabitEthernet0/0/13]interface GigabitEthernet 0/0/10 
[S1-GigabitEthernet0/0/10]port link-type trunk 
[S1-GigabitEthernet0/0/10]port trunk allow-pass vlan all 
[S1-GigabitEthernet0/0/10]bpdu enable 
 
