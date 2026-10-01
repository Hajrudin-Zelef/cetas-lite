---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-15
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [3125, 3307]
sha256: 83fdc9a816a20041905763f30fdbf735e0289a08a8f54a05be5833aed59d268a
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

for 20 times. 
Note: Disable the GigabitEthernet 0/0/10 interface of S2 immediately after 
running the ping command on S1. 
[S1]ping -c 20 10.0.1.2 
  PING 10.0.1.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.1.2: bytes=56 Sequence=1 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=2 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=3 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=4 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=5 ttl=254 time=1 ms 
Dec 21 2011 16:37:10-05:13 S1 %%01IFNET/4/IF_STATE(l)[7]:Interface 
GigabitEthernet0/0/10 has turned into DOWN state. 
    Request time out 
    Reply from 10.0.1.2: bytes=56 Sequence=7 ttl=255 time=10 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=8 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=9 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=10 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=11 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=12 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=13 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=14 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=15 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=16 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=17 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=18 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=19 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=20 ttl=254 time=1 ms 
 
  --- 10.0.1.2 ping statistics --- 
    20 packet(s) transmitted 
    19 packet(s) received 
    5.00% packet loss 
round-trip min/avg/max = 1/1/10 ms 
 
[S2]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]shutdown 
 
View role information about interfaces of S1. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ROOT  FORWARDING      NONE

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page71 
 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/9 interface of  S1 becomes the root interface and 
enters the FORWARDING state. There is one expired packet and the network 
convergence takes 2 seconds. 
Enable the GigabitEthernet 0/0/10 interface of S2. 
[S2]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]undo shutdown 
 
Step 6 Perform compatibility configuration between RSTP and 
STP . 
Enable STP on S1 and retain other configurations. 
[S1]stp mode stp 
 
View role information about interfaces of S1. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ALTE  DISCARDING      NONE 
   0    GigabitEthernet0/0/10       ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/10 interface of  S1 is the root interface. Run the 
ping command to test whether the route from S1 to S2 is reachable for 20 
times. 
Note: Disable the GigabitEthernet 0/0/10 interface of S2 immediately after 
running the ping command on S1. 
 [S1]ping -c 20 10.0.1.2 
  PING 10.0.1.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.1.2: bytes=56 Sequence=1 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=2 ttl=254 time=1 ms 
Dec 21 2011 16:20:44-05:13 S1 %%01IFNET/4/IF_STATE(l)[5]:Interface 
GigabitEthernet0/0/10 has turned into DOWN state. 
    Request time out 
    Request time out

HCDP-IESN  Chapter 2 STP and SEP 
 
Page72 HUAWEI TECHNOLOGIES HC Series 
 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Request time out 
    Reply from 10.0.1.2: bytes=56 Sequence=18 ttl=255 time=15 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=19 ttl=254 time=1 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=20 ttl=254 time=1 ms 
 
  --- 10.0.1.2 ping statistics --- 
    20 packet(s) transmitted 
    5 packet(s) received 
    75.00% packet loss 
    round-trip min/avg/max = 1/3/15 ms 
 
[S2]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]shutdown 
 
View role information about interfaces of S1. 
[S1]display stp brief 
 MSTID  Port                        Role  STP State     Protection 
   0    GigabitEthernet0/0/9        ROOT  FORWARDING      NONE 
   0    GigabitEthernet0/0/13       DESI  FORWARDING      NONE 
   0    GigabitEthernet0/0/14       DESI  FORWARDING      NONE 
 
The GigabitEthernet 0/0/9 interface of  S1 becomes the root interface and 
enters the FORWARDING state. There are 15 expired packets and the 
network convergence takes 30 seconds. 
RSTP is compatible with STP but it still uses the convergence mechanism 
of STP. 
Enable the GigabitEthernet 0/0/10 interface of S2. 
[S2]interface GigabitEthernet 0/0/10

HCDP-IESN  Chapter 2 STP and SEP 
 
HC Series HUAWEI TECHNOLOGIES     Page73 
 
[S2-GigabitEthernet0/0/10]undo shutdown 
 
Step 7 Configure MSTP and verify the configuration. 
Create VLANs 2 to 20 and add interfaces to relevant VLANs. 
[S1]vlan batch 2 to 20 
Info: This operation may take a few seconds. Please wait for a moment...done. 
[S1]interface GigabitEthernet 0/0/9 
[S1-GigabitEthernet0/0/9]port link-type trunk 
[S1-GigabitEthernet0/0/9]port trunk allow-pass vlan 1 TO 20 
[S1-GigabitEthernet0/0/9]interface GigabitEthernet 0/0/10 
[S1-GigabitEthernet0/0/10]port link-type trunk 
[S1-GigabitEthernet0/0/10]port trunk allow-pass vlan 1 TO 20 
[S1-GigabitEthernet0/0/10]interface GigabitEthernet 0/0/13 
[S1-GigabitEthernet0/0/13]port link-type trunk 
[S1-GigabitEthernet0/0/13]port trunk allow-pass vlan 1 TO 20 
[S1-GigabitEthernet0/0/13]interface GigabitEthernet 0/0/14 
[S1-GigabitEthernet0/0/14]port link-type trunk 
[S1-GigabitEthernet0/0/14]port trunk allow-pass vlan 1 TO 20 
 
[S2]vlan batch 1 to 20 
Info: This operation may take a few seconds. Please wait for a moment...done. 
[S2]interface GigabitEthernet 0/0/9 
[S2-GigabitEthernet0/0/9]port link-type trunk 
[S2-GigabitEthernet0/0/9]port trunk allow-pass vlan 1 TO 20 
[S2-GigabitEthernet0/0/9]interface GigabitEthernet 0/0/10 
[S2-GigabitEthernet0/0/10]port link-type trunk 
[S2-GigabitEthernet0/0/10]port trunk allow-pass vlan 1 TO 20 
[S2-GigabitEthernet0/0/10]interface GigabitEthernet 0/0/23 
[S2-GigabitEthernet0/0/23]port link-type trunk 
[S2-GigabitEthernet0/0/23]port trunk allow-pass vlan 1 TO 20 
[S2-GigabitEthernet0/0/23]interface GigabitEthernet 0/0/24 
[S2-GigabitEthernet0/0/24]port link-type trunk 
[S2-GigabitEthernet0/0/24]port trunk allow-pass vlan 1 TO 20 
 
[S3]vlan batch 1 to 20 
Info: This operation may take a few seconds. Please wait for a moment...done. 
[S3]interface Ethernet0/0/1 
[S3-Ethernet0/0/1]port link-type trunk 
[S3-Ethernet0/0/1]port trunk allow-pass vlan 1 TO 20 
[S3-Ethernet0/0/1]interface Ethernet0/0/13

HCDP-IESN  Chapter 2 STP and SEP 
 
Page74 HUAWEI TECHNOLOGIES HC Series 
 
[S3-Ethernet0/0/13]port link-type trunk 
[S3-Ethernet0/0/13]port trunk allow-pass vlan 1 TO 20 
[S3-Ethernet0/0/13]interface Ethernet0/0/23 
[S3-Ethernet0/0/23]port link-type trunk 
[S3-Ethernet0/0/23]port trunk allow-pass vlan 1 TO 20 
 
