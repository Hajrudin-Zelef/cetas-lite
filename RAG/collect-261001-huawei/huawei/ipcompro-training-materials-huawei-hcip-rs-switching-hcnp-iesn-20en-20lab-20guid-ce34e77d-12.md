---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-12
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [2428, 2722]
sha256: e6661b588d8ee6e5b1d473379aa05e3b40e5d151b3671065e5228f4869f8d018
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

Enable the function of sending Flush packets on S4, and the function of 
receiving Flush packets on S1 and S2. 
[S4]smart-link group 1 
[S4-smlk-group1]flush send control-vlan 100 password simple huawei 
 
[S1]interface GigabitEthernet 0/0/14 
[S1-GigabitEthernet0/0/14]smart-link flush receive control-vlan 100 password 
simple huawei 
 
[S2]interface GigabitEthernet 0/0/24 
[S2-GigabitEthernet0/0/24]smart-link flush receive control-vlan 100 password 
simple huawei 
 
Enable the Smart Link function on S4. 
[S4]smart-link group 1 
[S4-smlk-group1]smart-link enable 
 
Run the display smart-link group command to view the information about 
the Smart Link group on S4. 
[S4]display smart-link group 1 
Smart Link group 1 information : 
  Smart Link group was enabled 
  Wtr-time is: 30 sec. 
  There is no Load-Balance

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page54 HUAWEI TECHNOLOGIES HC Series 
 
  There is no protected-vlan reference-instance 
  DeviceID: 5489-98ec-f012  Control-vlan ID: 100 
      Member              Role   State    Flush Count Last-Flush-Time 
  ---------------------------------------------------------------------- 
  Ethernet0/0/14          Master Active   1           2008/01/05 03:11:18 
UTC-05:13 
  Ethernet0/0/24          Slave  Inactive 0           0000/00/00 00:00:00 
UTC+00:00 
 
Disable the E0/0/14 interface of S4 to verify the Smart Link function. 
[S4]interface Ethernet 0/0/14 
[S4-Ethernet0/0/14]shutdown 
[S4]display smart-link group 1 
Smart Link group 1 information : 
  Smart Link group was enabled 
  Wtr-time is: 30 sec. 
  There is no Load-Balance 
  There is no protected-vlan reference-instance 
  DeviceID: 5489-98ec-f012  Control-vlan ID: 100 
      Member              Role   State    Flush Count Last-Flush-Time 
  ---------------------------------------------------------------------- 
 
  Ethernet0/0/14          Master Inactive 1           2008/01/05 03:11:18 UTC-05:13 
  Ethernet0/0/24          Slave  Active   1          2008/01/05 03:14:57 UTC-05:13 
 
The preceding information shows that the Smart Link function enables S4 
to switch over to the slave interface when the master interface is faulty. 
Step 4 Configure hybrid networking of SEP and Smart Link. 
When Smart Link is configured on the lower-layer network, SEP must learn 
the conditions of the lower-layer net work to meet network changes. Enable 
SEP and Smart Link on the network. 
On S1 and S2, enable the function of processing Smart Link Flush 
packets. 
[S1]sep segment 1 
[S1-sep-segment1]deal smart-link-flush 
 
[S2]sep segment 1 
[S2-sep-segment1]deal smart-link-flush

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page55 
 
Additional Exercises: Analyzing and Verifying 
Compare SEP and STP, and figure out their respective strong points and 
drawbacks. 
Final Configurations 
[S1]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S1 
# 
 vlan batch 10 20 100 
# 
sep segment 1 
 control-vlan 10 
 block port optimal 
 preempt delay 30 
 protected-instance 0 to 48 
 deal smart-link-flush 
sep segment 2 
 control-vlan 20 
 block port optimal 
 preempt delay 30 
 tc-notify segment 1 
 protected-instance 0 to 48 
# 
interface GigabitEthernet0/0/9 
 port hybrid tagged vlan 10 
 stp disable 
 sep segment 1 edge primary 
# 
interface GigabitEthernet0/0/10 
 port hybrid tagged vlan 10 
 stp disable 
 sep segment 1 edge secondary 
# 
interface GigabitEthernet0/0/13 
 port hybrid tagged vlan 20 
 stp disable 
 sep segment 2 edge primary 
#

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page56 HUAWEI TECHNOLOGIES HC Series 
 
interface GigabitEthernet0/0/14 
 port link-type trunk 
 port trunk allow-pass vlan 100 
 smart-link flush receive control-vlan 100 password simple huawei 
# 
Return 
 
[S2]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S2 
# 
 vlan batch 10 20 100 
# 
sep segment 1 
 control-vlan 10 
 protected-instance 0 to 48 
 deal smart-link-flush 
sep segment 2 
 control-vlan 20 
 tc-notify segment 1 
 protected-instance 0 to 48 
# 
interface GigabitEthernet0/0/9 
 port hybrid tagged vlan 10 
 stp disable 
 sep segment 1 
# 
interface GigabitEthernet0/0/10 
 port hybrid tagged vlan 10 
 stp disable 
 sep segment 1 
 sep segment 1 priority 128 
# 
interface GigabitEthernet0/0/23 
 port hybrid tagged vlan 20 
 stp disable 
 sep segment 2 edge secondary 
# 
interface GigabitEthernet0/0/24 
 port link-type trunk 
 port trunk allow-pass vlan 100 
 smart-link flush receive control-vlan 100 password simple huawei

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page57 
 
# 
return 
 
[S3]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S3 
# 
 vlan batch 20 
# 
sep segment 2 
 control-vlan 20 
 protected-instance 0 to 48 
# 
# 
interface Ethernet0/0/13 
 port hybrid tagged vlan 20 
 stp disable 
 sep segment 2 
# 
interface Ethernet0/0/23 
 port hybrid tagged vlan 20 
 stp disable 
 sep segment 2 
 sep segment 2 priority 128 
# 
Return 
 
[S4]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S4 
# 
 vlan batch 20 100 
# 
sep segment 2 
 control-vlan 20 
 protected-instance 0 to 48 
# 
interface Ethernet0/0/1 
 shutdown 
# 
interface Ethernet0/0/14

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page58 HUAWEI TECHNOLOGIES HC Series 
 
 shutdown 
 port link-type trunk 
 port trunk allow-pass vlan 100 
 stp disable 
# 
interface Ethernet0/0/24 
 port link-type trunk 
 port trunk allow-pass vlan 100 
 stp disable 
# 
smart-link group 1 
 restore enable 
 smart-link enable 
 port Ethernet0/0/14 master 
 port Ethernet0/0/24 slave 
 timer wtr 30 
 flush send control-vlan 100 password simple huawei 
# 
Return

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page59 
 
Chapter 2 STP and SEP 
Lab 2-1 STP, RSTP, and MSTP Configuration 
Learning Objectives 
The objectives of this lab are to learn and understand how to perform the 
following operations: 
/g120 Learn the differences among Spanning Tree Protocol (STP), Rapid 
Spanning Tree Protocol (RSTP), and Multiple Spanning Tree Protocol 
(MSTP). 
/g120 Modify bridge priorities to control root bridge election. 
/g120 Modify interface priorities to control root interface and designated 
interface election. 
/g120 Configure RSTP and conduct compatibility configuration between STP 
and RSTP. 
/g120 Configure MSTP to implement load balancing among VLANs. 
Topology

HCDP-IESN  Chapter 2 STP and SEP 
 
Page60 HUAWEI TECHNOLOGIES HC Series 
 
Figure 2-1 STP, RSTP, and MSTP topology 
Scenario 
Assume that you are a network administrator of a company. The company 
uses a backup network. STP is used to avoid loops. By default ,STP 
convergence takes a long period of time. You can use RSTP to speed up 
network convergence. By default, all VLANs share a spanning tree. If you want 
to load balancing among VLANs,you can use MSTP. 
Tasks 
Step 1 Configure STP and verify the configuration. 
If the device default spanning tree is not turned on, use the following 
command to open. Configuration using the STP: 
<Quidway>system-view  
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S1 
[S1]stp enable 
[S1]stp mode stp 
 
<Quidway>system-view  
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S2 
[S2]stp enable 
[S2]stp mode stp 
 
<Quidway>system-view  
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S3 
[S3]stp enable 
[S3]stp mode stp 
 
<Quidway>system-view  
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S4 
[S4]stp enable 
[S4]stp mode stp

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page61 
 
