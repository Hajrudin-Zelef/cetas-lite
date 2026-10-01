---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-22
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost", "distribution"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [4592, 4810]
sha256: b5708498a2cb8331cec89d8cd759e35073268d3db5b1c50b60d11053e22cf8ba
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

Figure 3-1 MPLS LDP topology 
Scenario 
Assume that you are a network administrator of an enterprise. Your 
enterprise uses an IP network with poor forwarding performance. You need to 
use MPLS to improve the forwarding rate of routers. Static LSPs are 
configured manually, while LDP is a protocol developed for label distribution. 
To perform flexible configuration, use LDP to set up MPLS LSPs.

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
HC Series HUAWEI TECHNOLOGIES     
Page105 
 
Tasks 
Step 1 Configure IP addresses. 
Configure IP addresses and masks for all routers. 
<Huawei>system-view  
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname S1 
[S1]interface Vlanif 1 
[S1-Vlanif1]ip address 10.0.1.2 24 
 
<Huawei>system-view  
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R1 
[R1]interface GigabitEthernet 0/0/1 
[R1-GigabitEthernet0/0/1]ip address 10.0.1.1 24 
[R1-GigabitEthernet0/0/1]quit 
[R1]interface s1/0/0 
[R1-Serial1/0/0]ip address 10.0.12.1 24 
[R1-Serial1/0/0]quit 
[R1]interface loopback 0 
[R1-LoopBack0]ip address 2.2.2.2 24 
 
<Huawei>system-view  
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R2 
[R2]interface s1/0/0 
[R2-Serial1/0/0]ip address 10.0.12.2 24 
[R2-Serial1/0/0]quit 
[R2]interface s2/0/0 
[R2-Serial2/0/0]ip address 10.0.23.2 24 
[R2-Serial2/0/0]quit 
[R2]interface loopback 0 
[R2-LoopBack0]ip address 3.3.3.3 24 
 
<Huawei>system-view  
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname R3 
[R3]interface GigabitEthernet 0/0/2 
[R3-GigabitEthernet0/0/2]ip address 10.0.2.1 24

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page106 HUAWEI TECHNOLOGIES HC Series 
 
[R3-GigabitEthernet0/0/2]quit 
[R3]interface s2/0/0 
[R3-Serial2/0/0]ip address 10.0.23.3 24 
[R3-Serial2/0/0]quit 
[R3]interface loopback 0 
[R3-LoopBack0]ip address 4.4.4.4 24 
 
<Huawei>system-view  
Enter system view, return user view with Ctrl+Z. 
[Huawei]sysname S2 
[S2]interface Vlanif 1 
[S2-Vlanif1]ip address 10.0.2.2 24 
 
After the configurations are complete, test the connectivity of direct links. 
Step 2 Configure single-area OSPF. 
Configure network segments including 10.0.12.0/24, 10.0.23.0/24, 
10.0.1.0/24, and 10.0.2.0/24 to belong to OSPF area 0.
[S1]ospf 1 router-id 1.1.1.1 
[S1-ospf-1]area 0  
[S1-ospf-1-area-0.0.0.0]network 10.0.1.0 0.0.0.255 
 
[R1]ospf 1 router-id 2.2.2.2 
[R1-ospf-1]area 0  
[R1-ospf-1-area-0.0.0.0]network 10.0.1.0 0.0.0.255 
[R1-ospf-1-area-0.0.0.0]network 10.0.12.0 0.0.0.255 
[R1-ospf-1-area-0.0.0.0]network 2.2.2.0 0.0.0.255 
 
[R2]ospf 1 router-id 3.3.3.3 
[R2-ospf-1]area 0  
[R2-ospf-1-area-0.0.0.0]network 10.0.12.0 0.0.0.255 
[R2-ospf-1-area-0.0.0.0]network 10.0.23.0 0.0.0.255 
[R2-ospf-1-area-0.0.0.0]network 3.3.3.0 0.0.0.255   
 
[R3]ospf 1 router-id 4.4.4.4 
[R3-ospf-1]area 0 
[R3-ospf-1-area-0.0.0.0]net 
[R3-ospf-1-area-0.0.0.0]network 10.0.23.0 0.0.0.255 
[R3-ospf-1-area-0.0.0.0]network 10.0.2.0 0.0.0.255 
[R3-ospf-1-area-0.0.0.0]network 4.4.4.0 0.0.0.255

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
HC Series HUAWEI TECHNOLOGIES     
Page107 
 
 
[S2]ospf 1 router-id 5.5.5.5 
[S2-ospf-1]area 0 
[S2-ospf-1-area-0.0.0.0]network 10.0.2.0 0.0.0.255 
 
After the configurations are complete , check routing tables and test the 
network connectivity. 
[R2]ping 10.0.1.2 
  PING 10.0.1.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.1.2: bytes=56 Sequence=1 ttl=253 time=36 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=2 ttl=253 time=31 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=3 ttl=253 time=31 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=4 ttl=253 time=31 ms 
    Reply from 10.0.1.2: bytes=56 Sequence=5 ttl=253 time=31 ms 
 
  --- 10.0.1.2 ping statistics --- 
    5 packet(s) transmitted 
    5 packet(s) received 
    0.00% packet loss 
    round-trip min/avg/max = 31/32/36 ms 
 
[R2]ping 10.0.2.2 
  PING 10.0.2.2: 56  data bytes, press CTRL_C to break 
    Reply from 10.0.2.2: bytes=56 Sequence=1 ttl=253 time=38 ms 
    Reply from 10.0.2.2: bytes=56 Sequence=2 ttl=253 time=33 ms 
    Reply from 10.0.2.2: bytes=56 Sequence=3 ttl=253 time=33 ms 
    Reply from 10.0.2.2: bytes=56 Sequence=4 ttl=253 time=33 ms 
    Reply from 10.0.2.2: bytes=56 Sequence=5 ttl=253 time=33 ms 
 
  --- 10.0.2.2 ping statistics --- 
    5 packet(s) transmitted 
    5 packet(s) received 
    0.00% packet loss 
round-trip min/avg/max = 33/34/38 ms 
 
Run the display ip routing-table  command to view the OSPF routing 
tables of each router.  
[R2]display ip routing-table  
Route Flags: R - relay, D - download to fib 
---------------------------------------------------------------------------- 
Routing Tables: Public

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page108 HUAWEI TECHNOLOGIES HC Series 
 
         Destinations : 14       Routes : 14        
 
Destination/Mask    Proto   Pre  Cost      Flags NextHop         Interface 
 
       10.0.1.0/24   OSPF    10   1563        D   10.0.12.1       Serial1/0/0 
       10.0.2.0/24   OSPF    10   1563        D   10.0.23.3       Serial2/0/0 
      10.0.12.0/24   Direct  0    0           D   10.0.12.2       Serial1/0/0 
      10.0.12.1/32   Direct  0    0           D   10.0.12.1       Serial1/0/0 
      10.0.12.2/32   Direct  0    0           D   127.0.0.1       InLoopBack0 
    10.0.12.255/32   Direct  0    0           D   127.0.0.1       InLoopBack0 
      10.0.23.0/24   Direct  0    0           D   10.0.23.2       Serial2/0/0 
      10.0.23.2/32   Direct  0    0           D   127.0.0.1       InLoopBack0 
      10.0.23.3/32   Direct  0    0           D   10.0.23.3       Serial2/0/0 
    10.0.23.255/32   Direct  0   0           D   127.0.0.1       InLoopBack0 
      127.0.0.0/8    Direct  0    0           D   127.0.0.1       InLoopBack0 
      127.0.0.1/32   Direct  0    0           D   127.0.0.1       InLoopBack0 
127.255.255.255/32  Direct  0   0           D   127.0.0.1       InLoopBack0 
255.255.255.255/32  Direct  0   0           D   127.0.0.1       InLoopBack0 
 
Step 3 Configure MPLS LDP. 
Configure global MPLS and LDP on each LSR. 
[R1]mpls lsr-id 2.2.2.2 
[R1]mpls 
Info: Mpls starting, please wait... OK! 
[R1-mpls]mpls ldp 
 
[R2]mpls lsr-id 3.3.3.3 
[R2]mpls 
Info: Mpls starting, please wait... OK! 
[R2-mpls]mpls ldp 
 
[R3]mpls lsr-id 4.4.4.4 
[R3]mpls 
Info: Mpls starting, please wait... OK! 
[R3-mpls]mpls ldp 
 
Configure MPLS and LDP on interfaces of each LSR. 
[R1]interface Serial 1/0/0 
[R1-Serial1/0/0]mpls

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
HC Series HUAWEI TECHNOLOGIES     
Page109 
 
[R1-Serial1/0/0]mpls ldp 
 
[R2]interface Serial 1/0/0 
[R2-Serial1/0/0]mpls 
[R2-Serial1/0/0]mpls ldp 
[R2-Serial1/0/0]interface Serial 2/0/0 
[R2-Serial2/0/0]mpls 
[R2-Serial2/0/0]mpls ldp 
 
[R3]interface Serial 2/0/0  
[R3-Serial2/0/0]mpls 
[R3-Serial2/0/0]mpls ldp 
 
After the configurations are complete, run the display mpls ldp session  
command on R1, R2, and R3. The status of the local LDP session between R1, 
R2 and R3 is Operational. 
[R1]display mpls ldp session 
 LDP Session(s) in Public Network  
 Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM) 
 A '*' before a session means the session is being deleted. 
 ---------------------------------------------------------------------------- 
 PeerID             Status      LAM  SsnRole  SsnAge      KASent/Rcv 
 ---------------------------------------------------------------------------- 
 3.3.3.3:0          Operational DU   Passive  0000:00:10  41/41 
 ---------------------------------------------------------------------------- 
TOTAL: 1 session(s) Found. 
 
