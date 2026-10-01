---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-12
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [2663, 2855]
sha256: 31ba372fdc15596932fbd1391c002a4fc0d4408021195a360bd86bc2f2cd9579
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

Compare the routing tables with the previous routing tables before Gigabit Ethernet
0/0/2 was disabled.
<R2>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 12          Routes : 12


Destination/Mask        Proto    Pre Cost Flags NextHop               Interface


     10.0.2.0/24           Direct 0      0           D     10.0.2.2        LoopBack0


                                                      HUAWEI TECHNOLOGIES              Page59
    10.0.2.2/32           Direct 0    0       D   127.0.0.1 LoopBack0
    10.0.2.255/32         Direct 0    0       D   127.0.0.1 LoopBack0
    10.0.3.0/24           Static 80   0      RD    10.0.12.1   GigabitEthernet0/0/1
    10.0.12.0/24          Direct 0    0       D   10.0.12.2 GigabitEthernet0/0/1
    10.0.12.2/32          Direct 0    0       D   127.0.0.1 GigabitEthernet0/0/1
    10.0.12.255/32        Direct 0    0       D   127.0.0.1 GigabitEthernet0/0/1
    10.0.13.0/24          Static 80   0      RD    10.0.12.1 GigabitEthernet0/0/1
    127.0.0.0/8           Direct 0    0       D   127.0.0.1 InLoopBack0
    127.0.0.1/32          Direct 0    0       D   127.0.0.1 InLoopBack0
127.255.255.255/32        Direct 0    0       D   127.0.0.1 InLoopBack0
255.255.255.255/32        Direct 0    0       D   127.0.0.1 InLoopBack0



The next hops and preferences of the two routes as shown in the preceding routing
table for R2 have changed.
Test connectivity between R2 and the destination addresses 10.0.13.3 and 10.0.3.3
on R2.
<R2>ping 10.0.3.3
 PING 10.0.3.3: 56 data bytes, press CTRL_C to break
   Reply from 10.0.3.3: bytes=56 Sequence=1 ttl=255 time=3 ms
   Reply from 10.0.3.3: bytes=56 Sequence=2 ttl=255 time=2 ms
   Reply from 10.0.3.3: bytes=56 Sequence=3 ttl=255 time=2 ms
   Reply from 10.0.3.3: bytes=56 Sequence=4 ttl=255 time=2 ms
   Reply from 10.0.3.3: bytes=56 Sequence=5 ttl=255 time=2 ms


 --- 10.0.3.3 ping statistics ---
   5 packet(s) transmitted
   5 packet(s) received
   0.00% packet loss
   round-trip min/avg/max = 2/2/3 ms


<R2>ping 10.0.13.3
 PING 10.0.13.3: 56 data bytes, press CTRL_C to break
   Reply from 10.0.13.3: bytes=56 Sequence=1 ttl=255 time=3 ms
   Reply from 10.0.13.3: bytes=56 Sequence=2 ttl=255 time=2 ms
   Reply from 10.0.13.3: bytes=56 Sequence=3 ttl=255 time=2 ms
   Reply from 10.0.13.3: bytes=56 Sequence=4 ttl=255 time=2 ms
   Reply from 10.0.13.3: bytes=56 Sequence=5 ttl=255 time=2 ms


 --- 10.0.13.3 ping statistics ---
   5 packet(s) transmitted




                                              HUAWEI TECHNOLOGIES                     Page60
    5 packet(s) received
    0.00% packet loss
round-trip min/avg/max = 2/2/3 ms


The network is not disconnected when the link between R2 and R3 is shut down.


The tracert command can also be run to view through over which path the data is
being forwarded.
<R2>tracert 10.0.13.3
 traceroute to 10.0.13.3(10.0.13.3), max hops: 30 ,packet length: 40,press CTRL_C to break
 1 10.0.12.1 40 ms    21 ms 21 ms
 2 10.0.13.3 30 ms    21 ms 21 ms


<R2>tracert 10.0.3.3
 traceroute to 10.0.3.3(10.0.3.3), max hops: 30 ,packet length: 40,press CTRL_C to break
 1 10.0.12.1 40 ms    21 ms 21 ms
 2 10.0.13.3 30 ms    21 ms 21 ms


The command output shows that the data sent by R2 reaches R3 via the 10.0.12.0
and 10.0.13.0 networks connected to R1.

Step 7 Using default routes to implement network connectivity.

Enable the interface that was disabled in step 6 on R2.
[R2]interface GigabitEthernet 0/0/2
[R2-GigabitEthernet0/0/2]undo shutdown



Verify connectivity to the network 10.0.23.0 from R1.
[R1]ping 10.0.23.3
  PING 10.0.23.3: 56 data bytes, press CTRL_C to break
    Request time out
    Request time out
    Request time out
    Request time out
    Request time out


  --- 10.0.23.3 ping statistics ---
    5 packet(s) transmitted




                                                 HUAWEI TECHNOLOGIES                         Page61
    0 packet(s) received
100.00% packet loss


R3 cannot be reached because the route destined for 10.0.23.3 is not configured on
R1.


<R1>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 14           Routes : 14
Destination/Mask        Proto       Pre Cost Flags NextHop            Interface


     10.0.1.0/24           Direct     0     0         D    10.0.1.1 LoopBack0
     10.0.1.1/32           Direct     0     0         D    127.0.0.1 LoopBack0
     10.0.1.255/32         Direct     0     0         D    127.0.0.1 LoopBack0
     10.0.3.0/24           Static     60    0         RD 10.0.13.3 GigabitEthernet0/0/0
     10.0.12.0/24          Direct     0     0         D    10.0.12.1 GigabitEthernet0/0/1
     10.0.12.1/32          Direct     0     0         D    127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct     0     0         D    127.0.0.1 GigabitEthernet0/0/1
     10.0.13.0/24          Direct     0     0         D    10.0.13.1 GigabitEthernet0/0/0
     10.0.13.1/32          Direct     0     0         D    127.0.0.1 GigabitEthernet0/0/0
     10.0.13.255/32        Direct     0     0         D    127.0.0.1 GigabitEthernet0/0/0
     127.0.0.0/8           Direct     0     0         D    127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct     0     0         D    127.0.0.1 InLoopBack0
127.255.255.255/32         Direct     0     0         D    127.0.0.1 InLoopBack0
255.255.255.255/32         Direct     0     0         D    127.0.0.1 InLoopBack0


A default route can be configured on R1 to implement network connectivity via a
next hop of 10.0.13.3.
[R1]ip route-static 0.0.0.0 0.0.0.0 10.0.13.3


After the configuration is complete, test connectivity between R1 and 10.0.23.3.
<R1>ping 10.0.23.3
  PING 10.0.23.3: 56 data bytes, press CTRL_C to break
    Reply from 10.0.23.3: bytes=56 Sequence=1 ttl=255 time=3 ms
    Reply from 10.0.23.3: bytes=56 Sequence=2 ttl=255 time=2 ms
    Reply from 10.0.23.3: bytes=56 Sequence=3 ttl=255 time=2 ms
    Reply from 10.0.23.3: bytes=56 Sequence=4 ttl=255 time=2 ms
    Reply from 10.0.23.3: bytes=56 Sequence=5 ttl=255 time=2 ms




                                                      HUAWEI TECHNOLOGIES                   Page62
  --- 10.0.23.3 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
round-trip min/avg/max = 2/2/3 ms
The default route forwards traffic destined for 10.0.23.3 to the next hop of 10.0.13.3
on R3. R3 is directly connected to the 10.0.23.0 network.

Step 8 Configure a backup default route.

If the link between R1 and R3 fails, a backup default route can be used to
communicate with 10.0.23.3 and 10.0.3.3 via the 10.0.12.0 network.
However, R1 is not directly connected to these networks and therefore a backup
route (in both directions) must be configured to provide a forwarding path.
[R1]ip route-static 0.0.0.0 0.0.0.0 10.0.12.2 preference 80


[R3]ip route-static 10.0.12.0 24 10.0.23.2 preference 80



Step 9 Test the backup default route.

