---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-15
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: ["2016-03-30"]
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [3274, 3491]
sha256: abade9df97f1d981195bfb77d2b7d4baa2b534516ae7a6130a1492f005481c78
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

     172.16.0.0/24         Direct   0   0          D   172.16.0.1   LoopBack2
     172.16.0.1/32         Direct   0   0          D   127.0.0.1 LoopBack2
     172.16.0.255/32       Direct   0   0          D   127.0.0.1 LoopBack2
255.255.255.255/32         Direct   0   0          D   127.0.0.1 InLoopBack0



Test network connectivity between R2 and R1 at 10.0.1.1 and between R2 and R3 at
10.0.3.3.
<R2>ping 10.0.1.1
  PING 10.0.1.1: 56 data bytes, press CTRL_C to break
    Reply from 10.0.1.1: bytes=56 Sequence=1 ttl=255 time=37 ms
    Reply from 10.0.1.1: bytes=56 Sequence=2 ttl=255 time=42 ms
    Reply from 10.0.1.1: bytes=56 Sequence=3 ttl=255 time=42 ms
    Reply from 10.0.1.1: bytes=56 Sequence=4 ttl=255 time=45 ms
    Reply from 10.0.1.1: bytes=56 Sequence=5 ttl=255 time=42 ms


--- 10.0.1.1 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 37/41/45 ms


<R2>ping 10.0.3.3
  PING 10.0.3.3: 56 data bytes, press CTRL_C to break
    Reply from 10.0.3.3: bytes=56 Sequence=1 ttl=254 time=37 ms
    Reply from 10.0.3.3: bytes=56 Sequence=2 ttl=254 time=42 ms
    Reply from 10.0.3.3: bytes=56 Sequence=3 ttl=254 time=42 ms
    Reply from 10.0.3.3: bytes=56 Sequence=4 ttl=254 time=42 ms
    Reply from 10.0.3.3: bytes=56 Sequence=5 ttl=254 time=42 ms


 --- 10.0.3.3 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
    round-trip min/avg/max = 37/41/42 ms




Run the display ospf peer command to view the OSPF neighbor status.
<R1>display ospf peer


          OSPF Process 1 with Router ID 10.0.1.1



                                                   HUAWEI TECHNOLOGIES          Page73
                   Neighbors


 Area 0.0.0.0 interface 10.0.12.1(GigabitEthernet0/0/1)'s neighbors
 Router ID: 10.0.2.2            Address: 10.0.12.2
   State: Full Mode:Nbr is Master         Priority: 1
   DR: 10.0.12.1 BDR: 10.0.12.2       MTU: 0
   Dead timer due in 32 sec
   Retrans timer interval: 5
   Neighbor is up for 00:47:59
   Authentication Sequence: [ 0 ]


                   Neighbors


 Area 0.0.0.0 interface 10.0.13.1(GigabitEthernet0/0/0)'s neighbors
 Router ID: 10.0.3.3            Address: 10.0.13.3
   State: Full Mode:Nbr is Master         Priority: 1
   DR: 10.0.13.1 BDR: 10.0.13.3       MTU: 0
   Dead timer due in 34 sec
   Retrans timer interval: 5
   Neighbor is up for 00:41:44
   Authentication Sequence: [ 0 ]


The display ospf peer command displays detailed information about any peering
neighbors. In the example given, the link 10.0.13.1 of R1 shows to be the DR. The DR
election is non pre-emptive, meaning that the link of R3 will not take over the role of
DR from R1 unless the OSPF process is reset.

The display ospf peer brief command can also be used to display a condensed
version of the OSPF peer information.
<R1>display ospf peer brief


           OSPF Process 1 with Router ID 10.0.1.1
                    Peer Statistic Information
 ----------------------------------------------------------------------------
 Area Id            Interface                             Neighbor id            State
 0.0.0.0           GigabitEthernet0/0/0                   10.0.3.3              Full
 0.0.0.0           GigabitEthernet0/0/1                   10.0.2.2              Full
----------------------------------------------------------------------------
<R2>display ospf peer brief


           OSPF Process 1 with Router ID 10.0.2.2
                    Peer Statistic Information



                                                        HUAWEI TECHNOLOGIES              Page74
 ----------------------------------------------------------------------------
 Area Id            Interface                            Neighbor id             State
 0.0.0.0          GigabitEthernet0/0/1                    10.0.1.1              Full
 ----------------------------------------------------------------------------


<R3>display ospf peer brief
           OSPF Process 1 with Router ID 10.0.3.3
                    Peer Statistic Information
 ----------------------------------------------------------------------------
 Area Id            Interface                            Neighbor id             State
 0.0.0.0           GigabitEthernet0/0/0                   10.0.1.1              Full
 ----------------------------------------------------------------------------

Step 4 Change the OSPF hello interval and dead interval.

Run the display ospf interface GigabitEthernet 0/0/0 command on R1 to view the
default OSPF hello interval and dead interval.
<R1>display ospf interface GigabitEthernet 0/0/0


           OSPF Process 1 with Router ID 10.0.1.1
                   Interfaces




 Interface: 10.0.13.1 (GigabitEthernet0/0/0)
 Cost: 1        State: DR           Type: Broadcast       MTU: 1500
 Priority: 1
 Designated Router: 10.0.13.1
 Backup Designated Router: 10.0.13.3
 Timers: Hello 10 , Dead 40 , Poll 120 , Retransmit 5 , Transmit Delay 1


Run the ospf timer command to change the OSPF hello interval and dead interval
on GE0/0/0 of R1 to 15s and 60s respectively.
[R1]interface GigabitEthernet 0/0/0
[R1-GigabitEthernet0/0/0]ospf timer hello 15
[R1-GigabitEthernet0/0/0]ospf timer dead 60
Mar 30 2016 16:58:39+00:00 R1 %%01OSPF/3/NBR_DOWN_REASON(l)[1]:Neighbor state leaves full or changed
to Down. (ProcessId=1, NeighborRouterId=10.0.3.3, NeighborAreaId=0,
NeighborInterface=GigabitEthernet0/0/0,NeighborDownImmediate reason=Neighbor Down Due to Inactivity,
NeighborDownPrimeReason=Interface Parameter Mismatch, NeighborChangeTime=2016-03-30 16:58:39)


<R1>display ospf interface GigabitEthernet 0/0/0



                                                      HUAWEI TECHNOLOGIES                Page75
           OSPF Process 1 with Router ID 10.0.1.1
                   Interfaces




 Interface: 10.0.13.1 (GigabitEthernet0/0/0)
 Cost: 1        State: DR           Type: Broadcast       MTU: 1500
 Priority: 1
 Designated Router: 10.0.13.1
 Backup Designated Router: 10.0.13.3
 Timers: Hello 15 , Dead 60 , Poll 120 , Retransmit 5 , Transmit Delay 1


Check the OSPF neighbor status on R1.
<R1>display ospf peer brief


           OSPF Process 1 with Router ID 10.0.1.1
                    Peer Statistic Information
 ----------------------------------------------------------------------------
 Area Id            Interface                            Neighbor id             State
 0.0.0.0          GigabitEthernet0/0/1                    10.0.2.2              Full
 ----------------------------------------------------------------------------


The preceding information shows that R1 has only one neighbor, R2. Since the OSPF
hello intervals and dead intervals on R1 and R3 are different, R1 and R3 will fail to
establish an OSPF neighbor relationship.
Run the ospf timer command to change the OSPF hello interval and dead interval
on GE0/0/0 of R3 to 15s and 60s respectively.
[R3]interface GigabitEthernet 0/0/0
[R3-GigabitEthernet0/0/0]ospf timer hello 15
[R3-GigabitEthernet0/0/0]ospf timer dead 60
…output omitted…
Mar 30 2016 17:03:33+00:00 R3 %%01OSPF/4/NBR_CHANGE_E(l)[4]:Neighbor changes event: neighbor status
changed. (ProcessId=1, NeighborAddress=10.0.13.1, NeighborEvent=LoadingDone,
NeighborPreviousState=Loading, NeighborCurrentState=Full)


<R3>display ospf interface GigabitEthernet 0/0/0


           OSPF Process 1 with Router ID 10.0.3.3
                   Interfaces




                                                      HUAWEI TECHNOLOGIES                Page76
 Interface: 10.0.13.3 (GigabitEthernet0/0/0)
 Cost: 1         State: DR             Type: Broadcast       MTU: 1500
 Priority: 1
 Designated Router: 10.0.13.3
 Backup Designated Router: 10.0.13.1
 Timers: Hello 15 , Dead 60 , Poll 120 , Retransmit 5 , Transmit Delay 1



