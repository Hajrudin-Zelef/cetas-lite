---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-16
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [3492, 3664]
sha256: 80cc80f52f3778b3989bf421685452e4b7665b1f2505023ae3606f75db9659b7
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

Check the OSPF neighbor status on R1 again.
<R1>display ospf peer brief
           OSPF Process 1 with Router ID 10.0.1.1
                     Peer Statistic Information
 ----------------------------------------------------------------------------
 Area Id             Interface                               Neighbor id         State
 0.0.0.0           GigabitEthernet0/0/0                      10.0.3.3           Full
 0.0.0.0           GigabitEthernet0/0/1                      10.0.2.2           Full
----------------------------------------------------------------------------



Step 5 Advertise default routes in OSPF.

Configure OSPF to advertise default routes on R3.
[R3]ip route-static 0.0.0.0 0.0.0.0 LoopBack 2
[R3]ospf 1
[R3-ospf-1]default-route-advertise



View routing tables of R1 and R2. You can see that R1 and R2 have learned the
default routes advertised by R3.
<R1>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
           Destinations : 16            Routes : 16


Destination/Mask        Proto         Pre Cost Flags     NextHop        Interface


     0.0.0.0/0               O_ASE       150 1           D    10.0.13.3 GigabitEthernet0/0/0
     10.0.1.0/24             Direct      0    0          D    10.0.1.1 LoopBack0
     10.0.1.1/32             Direct      0    0          D    127.0.0.1 LoopBack0
     10.0.1.255/32           Direct      0    0          D    127.0.0.1 LoopBack0
     10.0.2.2/32             OSPF        10   1          D    10.0.12.2 GigabitEthernet0/0/1


                                                         HUAWEI TECHNOLOGIES                   Page77
     10.0.3.3/32           OSPF        10   1           D   10.0.13.3 GigabitEthernet0/0/0
     10.0.12.0/24          Direct      0    0           D   10.0.12.1 GigabitEthernet0/0/1
     10.0.12.1/32          Direct      0    0           D   127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct      0    0           D   127.0.0.1 GigabitEthernet0/0/1
     10.0.13.0/24          Direct      0    0           D   10.0.13.1 GigabitEthernet0/0/0
     10.0.13.1/32          Direct      0    0           D   127.0.0.1 GigabitEthernet0/0/0
     10.0.13.255/32        Direct      0    0           D   127.0.0.1 GigabitEthernet0/0/0
     127.0.0.0/8           Direct      0    0           D   127.0.0.1 InLoopBack0
     127.0.0.1/32               Direct      0       0       D     127.0.0.1 InLoopBack0
127.255.255.255/32         Direct      0    0           D   127.0.0.1 InLoopBack0
255.255.255.255/32         Direct      0    0           D   127.0.0.1 InLoopBack0




<R2>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 14           Routes : 14


Destination/Mask        Proto       Pre Cost Flags      NextHop       Interface


     0.0.0.0/0             O_ASE       150 1            D   10.0.12.1 GigabitEthernet0/0/1
     10.0.1.1/32           OSPF1       0    1           D   10.0.12.1 GigabitEthernet0/0/1
     10.0.2.0/24           Direct      0    0           D   10.0.2.2 LoopBack0
     10.0.2.2/32           Direct      0    0           D   127.0.0.1 LoopBack0
     10.0.2.255/32         Direct      0    0           D   127.0.0.1 LoopBack0
     10.0.3.3/32           OSPF        10   2           D   10.0.12.1 GigabitEthernet0/0/1
     10.0.12.0/24          Direct      0    0           D   10.0.12.2 GigabitEthernet0/0/1
     10.0.12.2/32          Direct      0    0           D   127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct      0    0           D   127.0.0.1 GigabitEthernet0/0/1
     10.0.13.0/24          OSPF        10   2           D   10.0.12.1 GigabitEthernet0/0/1
     127.0.0.0/8           Direct      0    0           D   127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct      0    0           D   127.0.0.1 InLoopBack0
127.255.255.255/32         Direct      0    0           D   127.0.0.1 InLoopBack0
255.255.255.255/32         Direct      0    0           D   127.0.0.1 InLoopBack0




<R3>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------



                                                        HUAWEI TECHNOLOGIES                  Page78
Routing Tables: Public
          Destinations : 17           Routes : 17


Destination/Mask       Proto        Pre Cost Flags   NextHop        Interface


     0.0.0.0/0             Static      60   0        D   172.16.0.1     LoopBack2
     10.0.1.1/32           OSPF        10   1        D   10.0.13.1 GigabitEthernet0/0/0
     10.0.2.2/32           OSPF        10   2        D   10.0.13.1 GigabitEthernet0/0/0
     10.0.3.0/24           Direct      0    0        D   10.0.3.3       LoopBack0
     10.0.3.3/32           Direct      0    0        D   127.0.0.1 LoopBack0
     10.0.3.255/32         Direct      0    0        D   127.0.0.1 LoopBack0
     10.0.12.0/24          OSPF        10   2        D   10.0.13.1 GigabitEthernet0/0/0
     10.0.13.0/24          Direct      0    0        D   10.0.13.3 GigabitEthernet0/0/0
     10.0.13.3/32          Direct      0    0        D   127.0.0.1 GigabitEthernet0/0/0
     10.0.13.255/32        Direct      0    0        D   127.0.0.1 GigabitEthernet0/0/0
     127.0.0.0/8           Direct      0    0        D   127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct      0    0        D   127.0.0.1 InLoopBack0
127.255.255.255/32         Direct      0    0        D   127.0.0.1 InLoopBack0
     172.16.0.0/24         Direct      0    0        D   172.16.0.1     LoopBack2
     172.16.0.1/32         Direct      0    0        D   127.0.0.1 LoopBack2
     172.16.0.255/32       Direct      0    0        D   127.0.0.1 LoopBack2
255.255.255.255/32         Direct      0    0        D   127.0.0.1 InLoopBack0


Run the ping command to test connectivity between R2 and Loopback2 at
172.16.0.1.
<R2>ping 172.16.0.1
  PING 172.16.0.1: 56 data bytes, press CTRL_C to break
    Reply from 172.16.0.1: bytes=56 Sequence=1 ttl=254 time=47 ms
    Reply from 172.16.0.1: bytes=56 Sequence=2 ttl=254 time=37 ms
    Reply from 172.16.0.1: bytes=56 Sequence=3 ttl=254 time=37 ms
    Reply from 172.16.0.1: bytes=56 Sequence=4 ttl=254 time=37 ms
    Reply from 172.16.0.1: bytes=56 Sequence=5 ttl=254 time=37 ms


 --- 172.16.0.1 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 37/39/47 ms




                                                     HUAWEI TECHNOLOGIES                  Page79
Step 6 Control OSPF DR or BDR election.

Run the display ospf peer command to view the DR and BDR of R1 and R3.
<R1>display ospf peer 10.0.3.3


         OSPF Process 1 with Router ID 10.0.1.1
                  Neighbors


 Area 0.0.0.0 interface 10.0.13.1(GigabitEthernet0/0/0)'s neighbors
 Router ID: 10.0.3.3           Address: 10.0.13.3
   State: Full Mode:Nbr is Master       Priority: 1
   DR: 10.0.13.3 BDR: 10.0.13.1 MTU: 0
   Dead timer due in 49 sec
   Retrans timer interval: 5
   Neighbor is up for 00:17:40
   Authentication Sequence: [ 0 ]


The preceding information shows that R3 is the DR and R1 is the BDR. This is
because R3's router ID 10.0.3.3 is greater than R1's router ID 10.0.1.1. R1 and R3 use
the default priority of 1, so their router IDs are used for DR or BDR election.
Run the ospf dr-priority command to change DR priorities of R1 and R3.
[R1]interface GigabitEthernet 0/0/0
[R1-GigabitEthernet0/0/0]ospf dr-priority 200


[R3]interface GigabitEthernet 0/0/0
[R3-GigabitEthernet0/0/0]ospf dr-priority 100


