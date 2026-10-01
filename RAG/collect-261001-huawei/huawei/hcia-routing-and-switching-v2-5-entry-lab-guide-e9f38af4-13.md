---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-13
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [2856, 3090]
sha256: ace9cd0d9fcbc48115b7a0b7f6f66a204f189119eb5e00644cf7279d8ef31bef
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

View the routes of R1 when the link between R1 and R3 is operational.
<R1>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 15           Routes : 15
Destination/Mask        Proto    Pre Cost Flags NextHop               Interface
     0.0.0.0/0             Static 60      0           RD 10.0.13.3 GigabitEthernet0/0/0
     10.0.1.0/24           Direct 0       0           D    10.0.1.1 LoopBack0
     10.0.1.1/32           Direct 0       0           D    127.0.0.1 LoopBack0
     10.0.1.255/32         Direct 0       0           D    127.0.0.1 LoopBack0
     10.0.3.0/24           Static 60      0           RD 10.0.13.3 GigabitEthernet0/0/0
     10.0.12.0/24          Direct 0       0           D    10.0.12.1 GigabitEthernet0/0/1
     10.0.12.1/32          Direct 0       0           D    127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct 0       0           D    127.0.0.1 GigabitEthernet0/0/1
     10.0.13.0/24          Direct 0       0           D    10.0.13.1 GigabitEthernet0/0/0
     10.0.13.1/32          Direct 0       0           D    127.0.0.1 GigabitEthernet0/0/0
     10.0.13.255/32        Direct 0       0           D    127.0.0.1 GigabitEthernet0/0/0



                                                      HUAWEI TECHNOLOGIES                   Page63
     127.0.0.0/8           Direct 0      0            D    127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct 0      0            D    127.0.0.1 InLoopBack0
127.255.255.255/32         Direct 0      0            D    127.0.0.1 InLoopBack0
255.255.255.255/32         Direct 0      0            D    127.0.0.1 InLoopBack0



Disable Gigabit Ethernet 0/0/0 on R1 and disable interface Gigabit Ethernet 0/0/0 on
R3 to simulate a link failure, and then view the routes of R1. Compare the current
routes with the routes before Gigabit Ethernet 0/0/0 was disabled.
[R1]interface GigabitEthernet0/0/0
[R1-GigabitEthernet0/0/0]shutdown
[R1-GigabitEthernet0/0/0]quit


[R3]interface GigabitEthernet0/0/0
[R3-GigabitEthernet0/0/0]shutdown
[R3-GigabitEthernet0/0/0]quit


<R1>display ip routing-table
Route Flags: R - relay, D - download to fib
-------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 11         Routes : 11
Destination/Mask           Proto    Pre Cost Flags NextHop            Interface
     0.0.0.0/0             Static 80     0      RD 10.0.12.2 GigabitEthernet0/0/1
     10.0.1.0/24           Direct 0      0            D    10.0.1.1 LoopBack0
     10.0.1.1/32           Direct 0      0            D    127.0.0.1 LoopBack0
     10.0.1.255/32         Direct 0      0            D    127.0.0.1 LoopBack0
     10.0.12.0/24          Direct 0      0            D    10.0.12.1 GigabitEthernet0/0/1
     10.0.12.1/32          Direct 0      0            D    127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct 0      0            D    127.0.0.1 GigabitEthernet0/0/1
     127.0.0.0/8           Direct 0      0            D    127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct 0      0            D    127.0.0.1 InLoopBack0
127.255.255.255/32         Direct 0      0            D    127.0.0.1 InLoopBack0
255.255.255.255/32         Direct 0      0            D    127.0.0.1 InLoopBack0



According to the preceding routing table, the value of 80 in the Preference column
indicates that the backup default route 0.0.0.0 is actively forwarding traffic to the
next hop of 10.0.23.3.
Test network connectivity on R1.
<R1>ping 10.0.23.3



                                                      HUAWEI TECHNOLOGIES                   Page64
    PING 10.0.23.3: 56 data bytes, press CTRL_C to break
      Reply from 10.0.23.3: bytes=56 Sequence=1 ttl=254 time=76 ms
      Reply from 10.0.23.3: bytes=56 Sequence=2 ttl=254 time=250 ms
      Reply from 10.0.23.3: bytes=56 Sequence=3 ttl=254 time=76 ms
      Reply from 10.0.23.3: bytes=56 Sequence=4 ttl=254 time=76 ms
      Reply from 10.0.23.3: bytes=56 Sequence=5 ttl=254 time=76 ms
    --- 10.0.23.3 ping statistics ---
      5 packet(s) transmitted
      5 packet(s) received
      0.00% packet loss
      round-trip min/avg/max = 76/110/250 ms
<R1>tracert 10.0.23.3
traceroute to 10.0.23.3(10.0.23.2), max hops: 30 ,packet length: 40,press CTRL_C to break
 1 10.0.12.2 30 ms      26 ms 26 ms
 2 10.0.23.3 60 ms      53 ms 56 ms


The IP packets are reaching R3 (10.0.23.3) via the next hop 10.0.12.2 of R2.


Final Configuration

<R1>dis current-configuration
[V200R007C00SPC600]
#
 sysname R1
#
interface GigabitEthernet0/0/0
 shutdown
 ip address 10.0.13.1 255.255.255.0
#
interface GigabitEthernet0/0/1
 ip address 10.0.12.1 255.255.255.0
#
interface LoopBack0
 ip address 10.0.1.1 255.255.255.0
#
ip route-static 0.0.0.0 0.0.0.0 10.0.13.3
ip route-static 0.0.0.0 0.0.0.0 10.0.12.2 preference 80
ip route-static 10.0.3.0 255.255.255.0 10.0.13.3
#
user-interface con 0
 authentication-mode password




                                                    HUAWEI TECHNOLOGIES                     Page65
 set authentication password cipher %$%$+L'YR&IZt'4,)>-*#lH",}%K-oJ_M9+'lOU~bD (\WTqB}%N,%$%$
user-interface vty 0 4
#
return
<R2>display current-configuration
[V200R007C00SPC600]
#
 sysname R2
interface GigabitEthernet0/0/1
 ip address 10.0.12.2 255.255.255.0
#
interface GigabitEthernet0/0/2
 ip address 10.0.23.2 255.255.255.0
#
interface LoopBack0
 ip address 10.0.2.2 255.255.255.0
#
ip route-static 10.0.3.0 255.255.255.0 10.0.23.3
ip route-static 10.0.3.0 255.255.255.0 10.0.12.1 preference 80
ip route-static 10.0.13.0 255.255.255.0 10.0.23.3
ip route-static 10.0.13.0 255.255.255.0 10.0.12.1 preference 80
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$1=cd%b%/O%Id-8X:by1N,+s}'4wD6TvO<I|/pd# #44C@+s#,%$%$
user-interface vty 0 4
#
return


<R3>display current-configuration
[V200R007C00SPC600]
#
 sysname R3
#
interface GigabitEthernet0/0/0
 shutdown
 ip address 10.0.13.3 255.255.255.0
#
interface GigabitEthernet0/0/2
 ip address 10.0.23.3 255.255.255.0
#
interface LoopBack0



                                                    HUAWEI TECHNOLOGIES              Page66
 ip address 10.0.3.3 255.255.255.0
#
ip route-static 10.0.12.0 255.255.255.0 10.0.13.1
ip route-static 10.0.12.0 255.255.255.0 10.0.23.2 preference 80
#
user-interface con 0
 authentication-mode password
 set authentication password cipher %$%$ksXDMg7Ry6yUU:63:DQ),#/sQg"@*S\U#.s.bHW xQ,y%#/v,%$%$
user-interface vty 0 4
#
return




                                                    HUAWEI TECHNOLOGIES              Page67
Lab 4-2 OSPF Single-Area Configuration


Learning Objectives

As a result of this lab section, you should achieve the following tasks:

      Configuration of the Router-ID for OSPF.

      Establish OSPF on a specified interface or network.

      View OSPF operations using display commands.

      Advertisement of default routes in OSPF.

      Change of the OSPF hello interval and dead interval.

      Familiarization with DR or BDR election on multi-access networks.

      Change of the OSPF route priority to manipulate DR election.


Topology




                          Figure 4.2 OSPF single area topology




                                      HUAWEI TECHNOLOGIES                  Page68
Scenario

As the network administrator of an establishing small enterprise, it is required that a
network be implemented using OSPF. Then network is to support a single area and
with consideration for future expansion it is requested that this area be set as area 0.
OSPF is required to advertise default routes and also elect both a DR and BDR for
network resiliency.


Tasks


Step 1 Prepare the environment

