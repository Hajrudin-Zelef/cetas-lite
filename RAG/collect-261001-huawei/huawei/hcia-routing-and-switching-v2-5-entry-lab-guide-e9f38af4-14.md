---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-14
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [3091, 3273]
sha256: bde7f500546040d7e6c2db7b73953053bcaea436603d3635f62e9d724f3b796c
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

If you are starting this section with a non-configured device, begin here and then
move to step 3. For those continuing from previous labs, begin at step 2.

Establish the basic system configuration and addressing for the lab.


<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R1
[R1]interface GigabitEthernet 0/0/1
[R1-GigabitEthernet 0/0/1]ip address 10.0.12.1 24
[R1-GigabitEthernet 0/0/1]quit
[R1]interface GigabitEthernet 0/0/0
[R1-GigabitEthernet0/0/0]ip address 10.0.13.1 24
[R1-GigabitEthernet0/0/0]quit
[R1]interface LoopBack 0
[R1-LoopBack0]ip address 10.0.1.1 24


<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R2
[R2]interface GigabitEthernet 0/0/1
[R2-GigabitEthernet 0/0/1]ip address 10.0.12.2 24
[R2-GigabitEthernet 0/0/1]quit
[R2]interface LoopBack 0
[R2-LoopBack0]ip address 10.0.2.2 24




                                                   HUAWEI TECHNOLOGIES       Page69
<Huawei>system-view
Enter system view, return user view with Ctrl+Z.
[Huawei]sysname R3
[R3]interface GigabitEthernet 0/0/0
[R3-GigabitEthernet0/0/0]ip address 10.0.13.3 24
[R3-GigabitEthernet0/0/0]quit
[R3]interface LoopBack 0
[R3-LoopBack0]ip address 10.0.3.3 24
[R3-LoopBack0]quit
[R3]interface LoopBack 2
[R3-LoopBack2]ip address 172.16.0.1 24



Step 2 Configure OSPF.

Assign the value 10.0.1.1 (as used on logical interface loopback 0 for simplicity) as
the router ID. Use OSPF process 1 (the default process), and specify network
segments 10.0.1.0/24, 10.0.12.0/24, and 10.0.13.0/24 as part of OSPF area 0.
[R1]ospf 1 router-id 10.0.1.1
[R1-ospf-1]area 0
[R1-ospf-1-area-0.0.0.0]network 10.0.1.0 0.0.0.255
[R1-ospf-1-area-0.0.0.0]network 10.0.13.0 0.0.0.255
[R1-ospf-1-area-0.0.0.0]network 10.0.12.0 0.0.0.255


Different process ID's will generate multiple link state databases, therefore ensure
that all routers use the same OSPF process ID. The wildcard mask must be specified
as part of the network command.

Manually assign the value 10.0.2.2 as the router ID. Use OSPF process 1, and
advertise network segments 10.0.12.0/24 and 10.0.2.0/24 into OSPF area 0.
[R2]ospf 1 router-id 10.0.2.2
[R2-ospf-1]area 0
[R2-ospf-1-area-0.0.0.0]network 10.0.2.0 0.0.0.255
[R2-ospf-1-area-0.0.0.0]network 10.0.12.0 0.0.0.255




…output omitted…
Mar 30 2016 09:41:39+00:00 R2 %%01OSPF/4/NBR_CHANGE_E(l)[5]:Neighbor changes event: neighbor status changed.
(ProcessId=1, NeighborAddress=10.0.12.1, NeighborEvent=LoadingDone, NeighborPreviousState=Loading,
NeighborCurrentState=Full)



                                                   HUAWEI TECHNOLOGIES                               Page70
Adjacency is attained when “NeighborCurrentState=Full”. For R3, Manually assign
the value 10.0.3.3 as the router ID. Use OSPF process 1, and advertise network
segments 10.0.3.0/24 and 10.0.13.0/24 into OSPF area 0.
[R3]ospf 1 router-id 10.0.3.3
[R3-ospf-1]area 0
[R3-ospf-1-area-0.0.0.0]network 10.0.3.0 0.0.0.255
[R3-ospf-1-area-0.0.0.0]network 10.0.13.0 0.0.0.255
…output omitted…
Mar 30 2016 16:05:34+00:00 R3 %%01OSPF/4/NBR_CHANGE_E(l)[5]:Neighbor changes event: neighbor status changed.
(ProcessId=1, NeighborAddress=10.0.13.1, NeighborEvent=LoadingDone, NeighborPreviousState=Loading,
NeighborCurrentState=Full)




Step 3 Verify the OSPF configuration.

After OSPF route convergence is complete, view routing tables of R1, R2, and R3.
<R1>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 15           Routes : 15


Destination/Mask        Proto       Pre Cost Flags    NextHop         Interface


     10.0.1.0/24           Direct      0    0         D    10.0.1.1 LoopBack0
     10.0.1.1/32           Direct      0    0         D    127.0.0.1 LoopBack0
     10.0.1.255/32         Direct      0    0         D    127.0.0.1 LoopBack0
     10.0.2.2/32           OSPF        10   1         D    10.0.12.2 GigabitEthernet0/0/1
     10.0.3.3/32           OSPF        10   1         D    10.0.13.3 GigabitEthernet0/0/0
     10.0.12.0/24          Direct      0    0         D    10.0.12.1 GigabitEthernet0/0/1
     10.0.12.1/32          Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/1
     10.0.13.0/24          Direct      0    0         D    10.0.13.1 GigabitEthernet0/0/0
     10.0.13.1/32          Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/0
     10.0.13.255/32        Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/0
     127.0.0.0/8           Direct      0    0         D    127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct      0    0         D    127.0.0.1 InLoopBack0
127.255.255.255/32         Direct      0    0         D    127.0.0.1 InLoopBack0
255.255.255.255/32         Direct      0    0         D    127.0.0.1 InLoopBack0




                                                      HUAWEI TECHNOLOGIES                            Page71
<R2>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 13           Routes : 13


Destination/Mask        Proto       Pre Cost        Flags NextHop              Interface


     10.0.1.1/32           OSPF        10   1         D    10.0.12.1 GigabitEthernet0/0/1
     10.0.2.0/24           Direct      0    0         D    10.0.2.2        LoopBack0
     10.0.2.2/32           Direct      0    0         D    127.0.0.1 LoopBack0
     10.0.2.255/32         Direct      0    0         D    127.0.0.1 LoopBack0
     10.0.3.3/32           OSPF        10   2         D    10.0.12.1 GigabitEthernet0/0/1
     10.0.12.0/24          Direct      0    0         D    10.0.12.2 GigabitEthernet0/0/1
     10.0.12.2/32          Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/1
     10.0.13.0/24          OSPF        10   2         D    10.0.12.1 GigabitEthernet0/0/1
     127.0.0.0/8           Direct      0    0         D    127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct      0    0         D    127.0.0.1 InLoopBack0
127.255.255.255/32         Direct      0    0         D    127.0.0.1 InLoopBack0
255.255.255.255/32         Direct      0    0         D    127.0.0.1 InLoopBack0


<R3>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 16           Routes : 16


Destination/Mask        Proto       Pre Cost Flags NextHop            Interface


     10.0.1.1/32           OSPF        10   1         D    10.0.13.1 GigabitEthernet0/0/0
     10.0.2.2/32           OSPF        10   2         D    10.0.13.1 GigabitEthernet0/0/0
     10.0.3.0/24           Direct      0    0         D    10.0.3.3 LoopBack0
     10.0.3.3/32           Direct      0    0         D    127.0.0.1 LoopBack0
     10.0.3.255/32         Direct      0    0         D    127.0.0.1 LoopBack0
     10.0.12.0/24          OSPF        10   2         D    10.0.13.1 GigabitEthernet0/0/0
     10.0.13.0/24          Direct      0    0         D    10.0.13.3 GigabitEthernet0/0/0
     10.0.13.3/32          Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/0
     10.0.13.255/32        Direct      0    0         D    127.0.0.1 GigabitEthernet0/0/0
     127.0.0.0/8           Direct      0    0         D    127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct      0    0         D    127.0.0.1 InLoopBack0
127.255.255.255/32         Direct      0    0         D    127.0.0.1 InLoopBack0



                                                      HUAWEI TECHNOLOGIES                   Page72

