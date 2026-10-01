---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-10
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [1947, 2227]
sha256: de31378ff434aa2e9c46ebfe26858ee37dd15f641fa2c7a141016cdb0fedb269
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[R2]interface GigabitEthernet 0/0/2 
[R2-GigabitEthernet0/0/2]ip address 10.0.100.3 24 
[R2-GigabitEthernet0/0/2]quit 
[R2]ip route-static 0.0.0.0 0 10.0.100.1 
 
Test whether the routes from R1 and R2 to the VLANIF 100 interface of S1 
are reachable. 
[R1]ping -c 1 10.0.100.1 
  PING 10.0.100.1: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.100.1: bytes=56 Sequence=1 ttl=254 time=3 ms 
 
  --- 10.0.100.1 ping statistics ---

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page42 HUAWEI TECHNOLOGIES HC Series 
 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 3/3/3 ms 
 
[R1]ping -c 1 10.0.100.3 
  PING 10.0.100.3: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.100.3: bytes=56 Sequence=1 ttl=254 time=2 ms 
 
  --- 10.0.100.3 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 2/2/2 ms 
 
[R2]pin -c 1 10.0.100.1 
  PING 10.0.100.1: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.100.1: bytes=56 Sequence=1 ttl=254 time=3 ms 
 
  --- 10.0.100.1 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 3/3/3 ms 
 
The preceding information shows that R1 and R2 can communicate with 
the VLANIF 100 interface of S1. Compared with Layer 3 switching, VLAN 
aggregation allows different VLANs to use the same gateway for 
communication. This reduces the number of required IP addresses and 
improves management efficiency. Howev er, computers on the same network 
segment communicate with each other over  the same VLANIF interface, this 
challenges the robustness of the interface. 
Additional Exercises: Analyzing and Verifying 
Figure out the features, strong points, drawbacks, and application 
scenarios of multi-arm routing, single-arm routing, inter-VLAN communication, 
and VLAN aggregation.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page43 
 
Final Configurations 
[S1]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S1 
# 
 vlan batch 2 to 3 10 20 100 
# 
vlan 100 
 aggregate-vlan 
 access-vlan 10 20 
# 
interface Vlanif2 
 ip address 10.0.20.1 255.255.255.0 
# 
interface Vlanif3 
 ip address 10.0.30.1 255.255.255.0 
# 
interface Vlanif100 
 ip address 10.0.100.1 255.255.255.0 
 arp-proxy inter-sub-vlan-proxy enable 
# 
interface GigabitEthernet0/0/1 
 port link-type access 
 port default vlan 10 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/4 
 shutdown 
port link-type trunk 
 port trunk allow-pass vlan 2 to 3 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/9 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 3 10 20 
 undo ntdp enable

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page44 HUAWEI TECHNOLOGIES HC Series 
 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/10 
 shutdown 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
Return 
 
[S2]display current-configuration 
# 
!Software Version V100R006C00SPC800 
 sysname S2 
# 
 vlan batch 2 to 3 10 20 100 
# 
interface GigabitEthernet0/0/2 
 port link-type access 
 port default vlan 20 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/4 
 shutdown 
 port link-type access 
 port default vlan 3 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/9 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 3 10 20 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
return 
 
[R4]display current-configuration

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page45 
 
[V200R001C00SPC500] 
# 
 sysname R4 
# 
interface GigabitEthernet0/0/1 
 ip address 10.0.2.1 255.255.255.0 
# 
interface GigabitEthernet0/0/1.2 
 control-vid 20 dot1q-termination 
 dot1q termination vid 2 
 ip address 10.0.20.1 255.255.255.0 
 arp broadcast enable 
# 
interface GigabitEthernet0/0/1.3 
 control-vid 30 dot1q-termination 
 dot1q termination vid 3 
 ip address 10.0.30.1 255.255.255.0 
 arp broadcast enable 
# 
interface GigabitEthernet0/0/0 
ip address 10.0.3.1 255.255.255.0 
# 
Return

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page46 HUAWEI TECHNOLOGIES HC Series 
 
Lab 1-4 SEP and Smart Link 
Learning Objectives 
The objectives of this lab are to learn and understand how to perform the 
following operations: 
/g120 Configure SEP. 
/g120 Configure Smart Link. 
/g120 Configure hybrid networking of SEP and Smart Lin k.
Topology 
 
Figure 1-4 SEP configuration 
Scenario 
Assume that you are a network administrator of a company. The company

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page47 
 
uses an Ethernet that has four switches. S1, S2, and S3 comprise the core 
network. The backup design is used to improve the robustness of the network. 
That is, SEP is used for loop protection. Configure two upstream links on the 
access layer switch (S4). The Smart Link technology provides master and 
slave links for a network, ensuring network reliability. 
Tasks 
Step 1 Perform basic configurations. 
Configure names for all devices and disa ble the E0/0/1 interface of S4 to 
avoid affecting the experiment. 
<Quidway>system-view 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S1 
[S1] 
 
<Quidway>system-view 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S2 
[S2] 
 
<Quidway>system-view 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S3 
[S3] 
 
<Quidway>system-view 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S4 
[S4]interface Ethernet 0/0/1 
[S4-Ethernet0/0/1]shutdown 
 
Step 2 Configure SEP. 
Redundant connections are created between S1, S2, and S3 to improve 
the robustness of the network. Two network loops exist. 
The G0/0/9 and G0/0/10 interfaces of S1 and S2 form a closed loop. The

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page48 HUAWEI TECHNOLOGIES HC Series 
 
G0/0/13 interface of S1, G0/0/13 and G0 /0/23 of S3, and G0/0/23 of S2 form 
an open loop. SEP is used to  provide redundancy protection for the two 
network loops. 
Create a SEP network segment, configure a control VLAN, and specify 
protection instances. 
[S1]sep segment 1 
[S1-sep-segment1]control-vlan 10 
[S1-sep-segment1]protected-instance all 
[S1-sep-segment1]quit 
[S1]sep segment 2 
[S1-sep-segment2]control-vlan 20 
[S1-sep-segment2]protected-instance all 
 
[S2]sep segment 1 
[S2-sep-segment1]control-vlan 10 
[S2-sep-segment1]protected-instance all 
[S2-sep-segment1]quit 
[S2]sep segment 2 
[S2-sep-segment2]control-vlan 20 
[S2-sep-segment2]protected-instance all 
 
[S3]sep segment 2 
[S3-sep-segment2]control-vlan 20 
[S3-sep-segment2]protected-instance all 
 
[S4]sep segment 2 
[S4-sep-segment2]control-vlan 20 
[S4-sep-segment2]protected-instance all 
 
Add interfaces to the SEP network se gment and configure roles for the 
interfaces. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]stp disable 
[S1-GigabitEthernet0/0/9]sep segment 1 edge primary 
[S1-GigabitEthernet0/0/9]interface GigabitEthernet 0/0/10 
[S1-GigabitEthernet0/0/10]stp disable 
[S1-GigabitEthernet0/0/10]sep segment 1 edge secondary 
[S1-GigabitEthernet0/0/10]interface GigabitEthernet 0/0/13 
[S1-GigabitEthernet0/0/13]stp disable 
[S1-GigabitEthernet0/0/13]sep segment 2 edge primary

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page49 
 
