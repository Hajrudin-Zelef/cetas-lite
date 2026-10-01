---
id: collect-261001-huawei/huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4-11
title: "hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4.md
source_anchor: ""
source_lines: [2487, 2662]
sha256: a376f6a5859fb614100827d69a03d72d6cda272ad9254194a36c7dc6b69dbc83
---

# hcia-routing-and-switching-v2-5-entry-lab-guide-e9f38af4

Run the display ip routing-table command to view the routing table of R2. The
routing table does not contain the routes of the two networks.
<R2>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 13         Routes : 13
Destination/Mask        Proto    Pre Cost Flags NextHop               Interface
     10.0.2.0/24           Direct 0      0            D    10.0.2.2 LoopBack0
     10.0.2.2/32           Direct 0      0            D    127.0.0.1 LoopBack0
     10.0.2.255/32         Direct 0      0            D    127.0.0.1 LoopBack0
     10.0.12.0/24          Direct 0      0            D    10.0.12.2 GigabitEthernet0/0/1
     10.0.12.2/32          Direct 0      0            D    127.0.0.1 GigabitEthernet0/0/1
     10.0.12.255/32        Direct 0      0            D    127.0.0.1 GigabitEthernet0/0/1
     10.0.23.0/24          Direct 0      0            D    10.0.23.2 GigabitEthernet0/0/2
     10.0.23.2/32          Direct 0      0            D    127.0.0.1 GigabitEthernet0/0/2
     10.0.23.255/32        Direct 0      0            D    127.0.0.1 GigabitEthernet0/0/2
     127.0.0.0/8           Direct 0      0            D    127.0.0.1 InLoopBack0
     127.0.0.1/32          Direct 0      0            D    127.0.0.1 InLoopBack0
127.255.255.255/32         Direct 0      0            D    127.0.0.1 InLoopBack0
255.255.255.255/32         Direct 0      0            D    127.0.0.1 InLoopBack0



Step 3 Configure static routes on R2.

Configure a static route for destination networks 10.0.13.0/24 and 10.0.3.0/24, with
the next hop set as the IP address 10.0.23.3 of R3, a preference value of 60 is the
default and need not be set.
[R2]ip route-static 10.0.13.0 24 10.0.23.3
[R2]ip route-static 10.0.3.0 24 10.0.23.3


Note: In the ip route-static command, 24 indicates the subnet mask length, which
can also be expressed using the decimal format 255.255.255.0.
<R2>display ip routing-table

Route Flags: R - relay, D - download to fib




                                                      HUAWEI TECHNOLOGIES                   Page56
Destination/Mask        Proto       Pre Cost Flags NextHop           Interface

   10.0.3.0/24             Static      60   0       RD 10.0.23.3           GigabitEthernet0/0/2
   10.0.12.0/24            Direct      0    0       D    10.0.12.2         GigabitEthernet0/0/1
   10.0.12.2/32            Direct      0    0       D    127.0.0.1         GigabitEthernet0/0/1
   10.0.12.255/32          Direct      0    0       D    127.0.0.1         GigabitEthernet0/0/1
   10.0.13.0/24            Static      60   0       RD 10.0.23.3           GigabitEthernet0/0/2
   10.0.23.0/24            Direct      0    0       D    10.0.23.2         GigabitEthernet0/0/2
   10.0.23.2/32            Direct      0    0       D    127.0.0.1         GigabitEthernet0/0/2



Step 4 Configure backup static routes.

The data exchanged between R2 and 10.0.13.3 and 10.0.3.3 is transmitted through
the link between R2 and R3. R2 fails to communicate with 10.0.13.3 and 10.0.3.3 if
the link between R2 and R3 is faulty.
According to the topology, R2 can communicate with R3 through R1 if the link
between R2 and R3 fails. A backup static route can be configured to enable this
redundancy. Backup static routes do not take effect in normal cases. If the link
between R2 and R3 fails, backup static routes are used to transfer data.
Amend th preferences for on the backup static routes to ensure that the routes are
used only when the primary link fails. In this example, the preference of the backup
static route is set to 80.
[R1]ip route-static 10.0.3.0 24 10.0.13.3


[R2]ip route-static 10.0.13.0 255.255.255.0 10.0.12.1 preference 80
[R2]ip route-static 10.0.3.0 24 10.0.12.1 preference 80


[R3]ip route-static 10.0.12.0 24 10.0.13.1




Step 5 Test the static routes.

View the current static route configuration in the routing table of R2.
<R2>display ip routing-table
Route Flags: R - relay, D - download to fib
----------------------------------------------------------------------------
Routing Tables: Public
          Destinations : 15           Routes : 15
  Destination/Mask      Proto       Pre Cost    Flags NextHop        Interface



                                                        HUAWEI TECHNOLOGIES                       Page57
       10.0.2.0/24        Direct 0    0       D   10.0.2.2 LoopBack0
       10.0.2.2/32        Direct 0    0       D   127.0.0.1 LoopBack0
     10.0.2.255/32        Direct 0    0       D   127.0.0.1 LoopBack0
       10.0.3.0/24        Static 60   0   RD 10.0.23.3 GigabitEthernet0/0/2
      10.0.12.0/24        Direct 0    0       D   10.0.12.2 GigabitEthernet0/0/1
      10.0.12.2/32        Direct 0    0       D   127.0.0.1 GigabitEthernet0/0/1
    10.0.12.255/32        Direct 0    0       D   127.0.0.1 GigabitEthernet0/0/1
      10.0.13.0/24        Static 60   0   RD 10.0.23.3 GigabitEthernet0/0/2
      10.0.23.0/24        Direct 0    0       D   10.0.23.2 GigabitEthernet0/0/2
      10.0.23.2/32        Direct 0    0       D   127.0.0.1 GigabitEthernet0/0/2
    10.0.23.255/32        Direct 0    0       D   127.0.0.1 GigabitEthernet0/0/2
       127.0.0.0/8        Direct 0    0       D   127.0.0.1 InLoopBack0
      127.0.0.1/32        Direct 0    0       D   127.0.0.1 InLoopBack0
127.255.255.255/32        Direct 0    0       D   127.0.0.1 InLoopBack0
255.255.255.255/32        Direct 0    0       D   127.0.0.1 InLoopBack0


The routing table contains two static routes that were configured in step 3. The value
of the Protocol field is Static, indicating a static route. The value of the Preference
field is 60, indicating the default preference is used for the route.
Test network connectivity to ensure the route between R2 and R3 exists.
<R2>ping 10.0.13.3
 PING 10.0.13.3: 56 data bytes, press CTRL_C to break
   Reply from 10.0.13.3: bytes=56 Sequence=1 ttl=255 time=34 ms
   Reply from 10.0.13.3: bytes=56 Sequence=2 ttl=255 time=34 ms
   Reply from 10.0.13.3: bytes=56 Sequence=3 ttl=255 time=34 ms
   Reply from 10.0.13.3: bytes=56 Sequence=4 ttl=255 time=34 ms
   Reply from 10.0.13.3: bytes=56 Sequence=5 ttl=255 time=34 ms


 --- 10.0.13.3 ping statistics ---
   5 packet(s) transmitted
   5 packet(s) received
   0.00% packet loss
   round-trip min/avg/max = 34/34/34 ms


<R2>ping 10.0.3.3
 PING 10.0.3.3: 56 data bytes, press CTRL_C to break
   Reply from 10.0.3.3: bytes=56 Sequence=1 ttl=255 time=41 ms
   Reply from 10.0.3.3: bytes=56 Sequence=2 ttl=255 time=41 ms
   Reply from 10.0.3.3: bytes=56 Sequence=3 ttl=255 time=41 ms
   Reply from 10.0.3.3: bytes=56 Sequence=4 ttl=255 time=41 ms
   Reply from 10.0.3.3: bytes=56 Sequence=5 ttl=255 time=41 ms



                                               HUAWEI TECHNOLOGIES                 Page58
  --- 10.0.3.3 ping statistics ---
    5 packet(s) transmitted
    5 packet(s) received
    0.00% packet loss
round-trip min/avg/max = 41/41/41 ms


The command output shows that the route is functioning normally. The tracert
command can also be run to view the path over which the data is transferred.
<R2>tracert 10.0.13.3
 traceroute to 10.0.13.3(10.0.13.3), max hops: 30 ,packet length: 40,
 press CTRL_C to break
 1 10.0.23.3 40 ms    31 ms 30 ms


<R2>tracert 10.0.3.3
 traceroute to 10.0.3.3(10.0.3.3), max hops: 30 ,packet length: 40,
 press CTRL_C to break
 1 10.0.23.3 40 ms    30 ms 30 ms



The command output verifies that R2 directly sends data to R3.

Step 6 Test the backup static routes.

Disable the path to 10.0.23.3 via GigabitEthernet0/0/2 on R2 and observe the
changes in the IP routing tables.
[R2]interface GigabitEthernet0/0/2
[R2-GigabitEthernet0/0/2]shutdown
[R2-GigabitEthernet0/0/2]quit


