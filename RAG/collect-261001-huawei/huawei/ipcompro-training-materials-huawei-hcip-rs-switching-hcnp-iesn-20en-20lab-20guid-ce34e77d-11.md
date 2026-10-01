---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-11
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "preemption"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [2228, 2427]
sha256: 39ab63fe58a03b2e65e0f60d014c67e8f3abe893f96b07bb89371ef8f2682567
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[S2]interface GigabitEthernet 0/0/23 
[S2-GigabitEthernet0/0/23]stp disable 
[S2-GigabitEthernet0/0/23]sep segment 2 edge secondary 
[S2-GigabitEthernet0/0/23]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]stp disable 
[S2-GigabitEthernet0/0/9]sep segment 1 
[S2-GigabitEthernet0/0/9]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]stp disable 
[S2-GigabitEthernet0/0/10]sep segment 1 
 
[S3]interface Ethernet 0/0/13 
[S3-Ethernet0/0/13]stp disable 
[S3-Ethernet0/0/13]sep segment 2 
[S3-Ethernet0/0/13]interface ethernet 0/0/23 
[S3-Ethernet0/0/23]stp disable 
[S3-Ethernet0/0/23]sep segment 2 
 
Configure the blocking and preemption modes. 
[S1]sep segment 1 
[S1-sep-segment1]block port optimal 
[S1-sep-segment1]quit 
[S1]sep segment 2 
[S1-sep-segment2]block port 
[S1-sep-segment2]block port optimal 
 
[S2]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]sep segment 1 priority 128 
 
[S3]inter Ethernet 0/0/23 
[S3-Ethernet0/0/23]sep segment 2 priority 128 
 
[S1]sep segment 1 
[S1-sep-segment1]preempt delay 30 
[S1-sep-segment1]quit 
[S1]sep segment 2 
[S1-sep-segment2]preempt delay 30 
 
Enable the function of advertising topology changes. 
[S1]sep segment 2 
[S1-sep-segment2]tc-notify segment 1 
 
[S2]sep segment 2

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page50 HUAWEI TECHNOLOGIES HC Series 
 
[S2-sep-segment2]tc-notify segment 1 
 
Run the display sep topology command to view SEP information. 
[S1]display sep topology 
SEP segment 1 
----------------------------------------------------------------- 
System Name          Port Name        Port Role       Port Status 
----------------------------------------------------------------- 
S1                   GE0/0/9          primary         forwarding 
S2                   GE0/0/9          common          forwarding 
S2                   GE0/0/10         common          discarding 
S1                   GE0/0/10         secondary       forwarding 
 
SEP segment 2 
----------------------------------------------------------------- 
System Name          Port Name        Port Role       Port Status 
----------------------------------------------------------------- 
S1                   GE0/0/13         primary         forwarding 
S3                   Eth0/0/13        common          forwarding 
S3                   Eth0/0/23        common          forwarding 
S2                   GE0/0/23         secondary       discarding 
 
Disable the G0/0/9 interface of S2 to check the running of SEP. 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]shutdown 
[S2-GigabitEthernet0/0/9]quit 
[S2]display sep topology 
SEP segment 1 
SEP detects a segment failure that may be caused by an incomplete topology 
----------------------------------------------------------------- 
System Name          Port Name        Port Role       Port Status 
----------------------------------------------------------------- 
S1                   GE0/0/9          secondary       discarding 
S1                   GE0/0/10         secondary       forwarding 
S2                   GE0/0/10         common          forwarding 
S2                   GE0/0/9          common          discarding 
 
SEP segment 2 
----------------------------------------------------------------- 
System Name          Port Name        Port Role       Port Status 
----------------------------------------------------------------- 
S1                   GE0/0/13         primary         forwarding

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page51 
 
S3                   Eth0/0/13        common          forwarding 
S3                   Eth0/0/23        common          forwarding 
S2                   GE0/0/23         secondary       discarding 
 
Enable the G0/0/9 interface of S2 and disable the E0/0/13 interface of S3 
to check the running of SEP. 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]undo shutdown 
 
[S3]interface Ethernet 0/0/13 
[S3-Ethernet0/0/13]shutdown 
[S3]display sep topology 
SEP segment 2 
SEP detects a segment failure that may be caused by an incomplete topology 
----------------------------------------------------------------- 
System Name          Port Name        Port Role       Port Status 
----------------------------------------------------------------- 
S2                   GE0/0/23         secondary       forwarding 
S3                   Eth0/0/23        common          forwarding 
S3                   Eth0/0/13        common          discarding 
 
Enable the E0/0/13 interface of S3to check the running of SEP. 
[S3]interface Ethernet 0/0/13 
[S3-Ethernet0/0/13]undo shutdown 
 
[S1]display sep topology 
SEP segment 1 
----------------------------------------------------------------- 
System Name          Port Name        Port Role       Port Status 
----------------------------------------------------------------- 
S1                   GE0/0/9          primary         forwarding 
S2                   GE0/0/9          common          forwarding 
S2                   GE0/0/10         common          discarding 
S1                   GE0/0/10         secondary       forwarding 
 
SEP segment 2 
----------------------------------------------------------------- 
System Name          Port Name        Port Role       Port Status 
----------------------------------------------------------------- 
S1                   GE0/0/13         primary         forwarding 
S3                   Eth0/0/13        common          forwarding 
S3                   Eth0/0/23        common          discarding

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page52 HUAWEI TECHNOLOGIES HC Series 
 
S2                   GE0/0/23         secondary       forwarding 
 
The preceding information shows that SEP changes the network topology 
when a fault occurs. After the fault is rectified , SEP blo cks the specified 
interface. 
Step 3 Configure Smart Link. 
The Smart Link technology provides  redundant upstream connections for 
users. The Smart Link technology sets the E0/0/14 interface of S4 as the 
master interface and the E0/0/24 interface  as the slave interface. When the 
master interface is faulty, S4 switches to the slave interface quickly, ensuring 
service continuity for users. 
Configure a control VLAN on S4 and add the interfaces to the control 
VLAN. 
[S4]vlan 100 
[S4-vlan100]quit 
[S4]interface Ethernet 0/0/14 
[S4-Ethernet0/0/14]port link-type trunk 
[S4-Ethernet0/0/14]port trunk allow-pass vlan 100 
[S4-Ethernet0/0/14]interface ethe0/0/24 
[S4-Ethernet0/0/24]port link-type trunk 
[S4-Ethernet0/0/24]port trunk allow-pass vlan 100 
 
Configure a control VLAN on S1 and S2. Add the G0/0/14 interfaces of S1 
and S2 to this VLAN. 
[S1]vlan 100 
[S1-vlan100]quit 
[S1]interface GigabitEthernet 0/0/14 
[S1-GigabitEthernet0/0/14]port link-type trunk 
[S1-GigabitEthernet0/0/14]port trunk allow-pass vlan 100 
 
[S2]vlan 100 
[S2-vlan100]quit 
[S2]interface GigabitEthernet 0/0/24 
[S2-GigabitEthernet0/0/24]port link-type trunk 
[S2-GigabitEthernet0/0/24]port trunk allow-pass vlan 100 
 
Disable Spanning Tree Protocol (STP) for the E0/0/14 and E0/0/24 
interfaces of S4.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page53 
 
[S4]interface Ethernet 0/0/14 
[S4-Ethernet0/0/14]stp disable 
[S4-Ethernet0/0/14]interface ethe0/0/24 
[S4-Ethernet0/0/24]stp disable 
 
Add the E0/0/14 and E0/0/24 interfaces  of S4 to the Smart Link group, 
E0/0/14 as the master interface and E0/0/24 as the slave interface. 
[S4]smart-link group 1 
[S4-smlk-group1]port Ethernet 0/0/14 master 
[S4-smlk-group1]port Ethernet 0/0/24 slave 
 
Enable the switchover function on S4 and set the switchover time. 
[S4]smart-link group 1 
[S4-smlk-group1]restore enable 
[S4-smlk-group1]timer wtr 30 
 
