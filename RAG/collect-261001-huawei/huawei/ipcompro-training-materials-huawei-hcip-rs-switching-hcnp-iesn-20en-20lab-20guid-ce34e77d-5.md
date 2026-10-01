---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-5
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [774, 1071]
sha256: 855c1c909159b4f076ba1a1a5bcf7dfea2a974132de14c9a5d7901a04dbe9756
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page14 HUAWEI TECHNOLOGIES HC Series 
 
# 
 vlan batch 10 20 
# 
 stp mode rstp 
 stp enable 
# 
 undo http server enable 
# 
interface Eth-Trunk1 
 port hybrid untagged vlan 10 
 mode lacp-static 
 max active-linknumber 1 
 bpdu enable 
# 
interface GigabitEthernet0/0/1 
 port hybrid pvid vlan 10 
 port hybrid tagged vlan 10 
# 
interface GigabitEthernet0/0/2 
 port link-type access 
 port default vlan 20 
# 
interface GigabitEthernet0/0/3 
 shutdown 
# 
interface GigabitEthernet0/0/4 
 shutdown 
# 
interface GigabitEthernet0/0/9 
 shutdown 
 eth-trunk 1 
# 
interface GigabitEthernet0/0/10 
 eth-trunk 1 
# 
interface GigabitEthernet0/0/13 
 shutdown 
# 
interface GigabitEthernet0/0/14 
 shutdown 
# 
interface GigabitEthernet0/0/21 
 shutdown

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page15 
 
# 
interface GigabitEthernet0/0/22 
 shutdown 
# 
interface GigabitEthernet0/0/23 
 shutdown 
# 
Return 
 
[S2]display current-configuration 
# 
!Software Version V100R005C01SPC100 
 sysname S2 
# 
 vlan batch 10 20 
# 
 stp mode rstp 
 stp enable 
# 
interface Eth-Trunk1 
 port hybrid untagged vlan 20 
 mode lacp-static 
 max active-linknumber 1 
 bpdu enable 
# 
interface GigabitEthernet0/0/1 
 shutdown 
# 
interface GigabitEthernet0/0/2 
 shutdown 
# 
interface GigabitEthernet0/0/3 
 port hybrid pvid vlan 20 
 port hybrid tagged vlan 20 
# 
interface GigabitEthernet0/0/4 
 shutdown 
# 
interface GigabitEthernet0/0/5 
 shutdown 
# 
interface GigabitEthernet0/0/9 
 eth-trunk 1

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page16 HUAWEI TECHNOLOGIES HC Series 
 
# 
interface GigabitEthernet0/0/10 
 eth-trunk 1 
# 
interface GigabitEthernet0/0/11 
 shutdown 
# 
interface GigabitEthernet0/0/12 
 shutdown 
# 
interface GigabitEthernet0/0/13 
 shutdown 
# 
interface GigabitEthernet0/0/23 
 shutdown 
# 
interface GigabitEthernet0/0/24 
 shutdown 
# 
return

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page17 
 
Lab 1-2 MUX VLAN Configuration and GVRP Configuration 
(Optional) 
Learning Objectives 
The objectives of this lab are to learn and understand how to perform the 
following operations: 
/g120 Configure GARP VLAN Registration Protocol (GVRP). 
/g120 Configure multiplexer (MUX) VLANs. 
/g120 Configure voice VLANs in automatic mode. 
Topology 
 
Figure 1-2 MUX VLAN configuration and GVRP configuration 
Scenario 
Assume that you are a network administrator of a company. The company 
uses an Ethernet that has two switche s. In the preceding figure, the routers 
represent computers on the network. To optimize the network, broadcast 
domains must be isolated from each other. R1 and R2 are in the same VLAN, 
and R3 and R4 are in the same VLAN. As required by the company, all

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page18 HUAWEI TECHNOLOGIES HC Series 
 
computers can access R4, but R3 and R4 cannot communicate with each 
other, with neither R1 nor R2. Voice devices will be used on the G0/0/24 
interface of S2. Therefore, plan a voic e VLAN and related configurations in 
advance. 
Tasks 
Step 1 Perform basic configurations and IP addressing. 
Configure IP addresses and subnet masks for all routers. 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R1 
[R1]interface g0/0/1 
[R1-GigabitEthernet0/0/1]ip address 10.0.10.1 24 
 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R2 
[R2]interface g0/0/1 
[R2-GigabitEthernet0/0/1]ip address 10.0.10.2 24 
 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R3 
[R3]interface g0/0/1 
[R3-GigabitEthernet0/0/1]ip address 10.0.10.3 24 
 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R4 
[R4]interface g0/0/2 
[R4-GigabitEthernet0/0/2]ip address 10.0.10.4 24 
 
<Huawei>system-view 
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R5 
[R5]interface g0/0/2 
[R5-GigabitEthernet0/0/2]ip address 10.0.10.5 24 
 
<Quidway>system-view

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page19 
 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S1 
[S1] 
 
<Quidway>system-view 
Enter system view, return user view with Ctrl+Z. 
[Quidway]sysname S2 
[S2] 
 
Test whether the routes from R1 to R2, R3, R4, and R5 are reachable. 
[R1]ping -c 1 10.0.10.2 
  PING 10.0.10.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.2: bytes=56 Sequence=1 ttl=255 time=14 ms 
 
  --- 10.0.10.2 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 14/14/14 ms 
 
[R1]ping -c 1 10.0.10.3 
  PING 10.0.10.3: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.3: bytes=56 Sequence=1 ttl=255 time=5 ms 
 
  --- 10.0.10.3 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 5/5/5 ms 
 
[R1]ping -c 1 10.0.10.4 
  PING 10.0.10.4: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.10.4: bytes=56 Sequence=1 ttl=255 time=15 ms 
 
  --- 10.0.10.4 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 15/15/15 ms 
 
[R1]ping -c 1 10.0.10.5 
  PING 10.0.10.5: 56  data bytes, press CTRL_C to break

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
Page20 HUAWEI TECHNOLOGIES HC Series 
 
    Reply from 10.0.10.5: bytes=56 Sequence=1 ttl=255 time=6 ms 
 
  --- 10.0.10.5 ping statistics --- 
    1 packet(s) transmitted 
    1 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 6/6/6 ms 
 
Step 2 Configure GVRP. 
GVRP transmits VLAN information across the network, improving the 
VLAN configuration efficiency. Configure the G0/0/9 interface of S2 to work in 
Fixed mode and that of S1 to work in Normal mode so that S1 can learn the 
VLAN configurations on S2. 
Configure the G0/0/9 interfaces of S1 and S2 to work in Trunk mode and 
allow the access from all VLANs. 
[S1]interface g0/0/9 
[S1-GigabitEthernet0/0/9]port link-type trunk 
[S1-GigabitEthernet0/0/9]port trunk allow-pass vlan all 
 
[S2]interface g0/0/9 
[S2-GigabitEthernet0/0/9]port link-type trunk 
[S2-GigabitEthernet0/0/9]port trunk allow-pass vlan all 
 
Enable GVRP on S1 and S2. 
[S1]gvrp 
 
[S2]gvrp 
 
Run the display gvrp status command to view GVRP status information. 
[S1]display gvrp status 
 GVRP is enabled 
 
[S2]display gvrp status 
 GVRP is enabled 
 
Configure the G0/0/9 interface of S1 to work in Normal mode. 
[S1]interface g0/0/9 
[S1-GigabitEthernet0/0/9]gvrp

HCDP-IESN  Chapter 1 Implementing VLAN features 
 
HC Series HUAWEI TECHNOLOGIES     Page21 
 
[S1-GigabitEthernet0/0/9]gvrp registration normal 
[S1-GigabitEthernet0/0/9]bpdu enable 
 
Configure the G0/0/9 interface of S2 to work in Fixed mode. 
[S2]interface g0/0/9 
[S2-GigabitEthernet0/0/9]gvrp 
[S2-GigabitEthernet0/0/9]gvrp registration fixed 
[S2-GigabitEthernet0/0/9]bpdu enable 
 
Run the display gvrp statistics command to view GVRP statistics. 
[S1]display gvrp statistics 
 
  GVRP statistics on port GigabitEthernet0/0/9 
    GVRP status                         : Enabled 
    GVRP registrations failed           : 0 
    GVRP last PDU origin                        : 4c1f-cc45-aacc 
GVRP registration type              : Normal 
 
