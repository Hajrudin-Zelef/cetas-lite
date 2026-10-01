---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-8
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [1505, 1740]
sha256: 630890266c330754d8244219bf27621069e23dc2cb71facb38737a1a527c5cbf
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

HC Series HUAWEI TECHNOLOGIES     Page31 
 
interface GigabitEthernet0/0/4 
 port link-type access 
 port default vlan 20 
 port mux-vlan enable 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/5 
 port link-type access 
 port default vlan 100 
 port mux-vlan enable 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/9 
 port link-type trunk 
 port trunk allow-pass vlan 2 to 4094 
 undo ntdp enable 
 undo ndp enable 
 gvrp 
 gvrp registration fixed 
# 
interface GigabitEthernet0/0/10 
 shutdown 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
interface GigabitEthernet0/0/24 
voice-vlan 200 enable 
 port hybrid pvid vlan 30 
 port hybrid untagged vlan 30 
 undo ntdp enable 
 undo ndp enable 
 bpdu disable 
# 
Return

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page32 HUAWEI TECHNOLOGIES HC Series 
 
Lab 1-3 Inter-VLAN Communication 
Learning Objectives 
The objectives of this lab are to learn and understand how to perform the 
following operations: 
/g120 Configure multi-arm routing. 
/g120 Configure single-arm routing. 
/g120 Configure inter-VLAN communication. 
/g120 Configure VLAN aggregation. 
Topology 
 
Figure 1-3 Inter-VLAN communication 
Scenario 
Assume that you are a network administrator of a company. The company 
uses an Ethernet that has two switches and one router. In the preceding figure, 
R1 and R2 represent two computers in different departments and are added to 
two VLANs. As required by the compan y, R1 needs to communicate with R2.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page33 
 
Multi-arm routing is used first and later replaced with single-arm routing to 
reduce the costs. 
Then multi-layer switching is used because traffic is mostly transmitted 
between VLANs as the network architecture changes. Finally the VLAN 
aggregation technology is used to facilitate network management. 
Tasks 
Step 1 Perform basic configurations and IP addressing. 
Disable G0/0/9 of S1 to avoid affecting the experiment. 
Configure IP addresses and subnet masks for all routers. 
<huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[huawei]sysname R1 
[R1]interface GigabitEthernet 0/0/1 
[R1-GigabitEthernet0/0/1]ip address 10.0.2.2 24 
 
<huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[huawei]sysname R2 
[R2]interface GigabitEthernet 0/0/2 
[R2-GigabitEthernet0/0/2]ip address 10.0.3.2 24 
 
<Quidway>system-view 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S1 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/10]undo shutdown 
 
<Quidway>system-view 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S2 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/10]undo shutdown 
 
<huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[huawei]sysname R4 
[R4]interface GigabitEthernet 0/0/1

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page34 HUAWEI TECHNOLOGIES HC Series 
 
[R4-GigabitEthernet0/0/1]ip address 10.0.2.1 24 
[R4-GigabitEthernet0/0/1]interface GigabitEthernet 0/0/2 
[R4-GigabitEthernet0/0/2]ip address 10.0.3.1 24 
 
Run the ping command to test whether the r oute from R1 to the G0/0/1 
interface of R4 is reachable. 
[R1]ping -c 1 10.0.2.1 
  PING 10.0.2.1: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.2.1: bytes=56 Sequence=1 ttl=255 time=4 ms 
 
  --- 10.0.2.1 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 4/4/4 ms 
 
Run the ping command to test whether the r oute from R2 to the G0/0/2 
interface of R4 is reachable. 
[R2]ping -c 1 10.0.3.1 
  PING 10.0.3.1: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.3.1: bytes=56 Sequence=1 ttl=255 time=3 ms 
 
  --- 10.0.3.1 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 3/3/3 ms 
 
Step 2 Configure multi-arm routing. 
R1 and R2 are in different VLANs. 
The gateway of R1 uses the IP address of the G0/0/1 interface of R4, and 
that of R2 uses the IP address of the G0/0/2 interface of R4. 
Multiple interfaces of R4 support communication between VLANs. This is 
multi-arm routing. 
Create VLAN 2 and VLAN 3 on S1 and S2. 
[S1]vlan batch 2 3 
Info: This operation may take a few seconds. Please wait for a moment...done.

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page35 
 
[S2]vlan batch 2 3 
Info: This operation may take a few seconds. Please wait for a moment...done. 
 
Add R1 and the G0/0/1 interface of R4 to VLAN 2, and R2 and the G0/0/2 
interface of R4 to VLAN 3. 
[S1]interface GigabitEthernet 0/0/1 
[S1-GigabitEthernet0/0/1]port link-type access 
[S1-GigabitEthernet0/0/1]port default vlan 2 
[S1-GigabitEthernet0/0/1]inter g0/0/4 
[S1-GigabitEthernet0/0/4]port link-type access 
[S1-GigabitEthernet0/0/4]port default vlan 2 
 
[S2]interface GigabitEthernet 0/0/2 
[S2-GigabitEthernet0/0/2]port link-type access 
[S2-GigabitEthernet0/0/2]port default vlan 3 
[S2-GigabitEthernet0/0/2]inter g0/0/4 
[S2-GigabitEthernet0/0/4]port link-type access 
[S2-GigabitEthernet0/0/4]port default vlan 3 
 
Configure a gateway on R1 and R2. The gateway of R1 uses the IP 
address of the G0/0/1 interface of R4, and  that of R2 uses the IP address of 
the G0/0/2 interface of R4. 
[R1]ip route-static 0.0.0.0 0 10.0.2.1 
 
[R2]ip route-static 0.0.0.0 0 10.0.3.1 
 
Run the display vlan command to check the configurations. 
[S1]display vlan 2 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports 
---------------------------------------------------------------------------- 
2    common  UT:GE0/0/1(U)      GE0/0/4(U) 
VID  Status  Property      MAC-LRN Statistics Description 
---------------------------------------------------------------------------- 
2    enable  default       enable  disable    VLAN 0002 
 
[S2]display vlan 3

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page36 HUAWEI TECHNOLOGIES HC Series 
 
---------------------------------------------------------------------------- 
U: Up;         D: Down;         TG: Tagged;         UT: Untagged; 
MP: Vlan-mapping;               ST: Vlan-stacking; 
#: ProtocolTransparent-vlan;    *: Management-vlan; 
---------------------------------------------------------------------------- 
VID  Type    Ports 
---------------------------------------------------------------------------- 
3    common  UT:GE0/0/2(U)      GE0/0/4(U) 
VID  Status  Property      MAC-LRN Statistics Description 
---------------------------------------------------------------------------- 
3    enable  default       enable  disable    VLAN 0003 
 
Test whether the route from R1 to R2 is reachable. 
[R1]ping -c 1 10.0.3.2 
  PING 10.0.3.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.3.2: bytes=56 Sequence=1 ttl=254 time=3 ms 
 
  --- 10.0.3.2 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 3/3/3 ms 
 
[R2]ping -c 1 10.0.2.2 
  PING 10.0.2.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.2.2: bytes=56 Sequence=1 ttl=254 time=3 ms 
 
  --- 10.0.2.2 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 3/3/3 ms 
Step 3 Configure single-arm routing. 
Create two subinterfaces on a physical interface of R4. Inter-VLAN 
communication is implemented on relevant subinterfaces. 
This is single-alarm routing. 
Disable the G0/0/4 interface of S2. 
[S2]interface GigabitEthernet 0/0/4

