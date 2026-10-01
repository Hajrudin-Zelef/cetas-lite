---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ospfv3-0096-html-695025ce-1
title: "Configure RouterA."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ospfv3-0096-html-695025ce.md
source_anchor: ""
source_lines: [1, 181]
sha256: b40b2986ea263cf5d468076312ddf346b56235b033a382fa7ccb67e54a936e96
---

# Configure RouterA.

This part provides an example for configuring BFD for OSPFv3. After BFD for OSPFv3 is configured, BFD can fast detect link faults and report them to OSPFv3 so that service traffic can be transmitted through the backup link.
As shown in Figure Networking diagram for configuring BFD for OSPFv3, it is required as follows:
Run OSPFv3 between RouterA, RouterB, and RouterC.
Enable BFD of the OSPFv3 process on RouterA, RouterB, and RouterC.
Traffic is transmitted on the active link RouterA â RouterB. The link RouterA â RouterC â RouterB acts as the standby link.
When a fault occurs on the link, BFD can quickly detect the fault and notify OSPFv3 of the fault; therefore, the traffic is transmitted on the standby link.
The configuration roadmap is as follows:
Enable the basic OSPFv3 functions on each Router.
Configuring BFD for OSPFv3.
The detailed configuration is not mentioned here.
# Configure RouterA.
[RouterA] ipv6
[RouterA] ospfv3
[RouterA-ospfv3-1] router-id 1.1.1.1
[RouterA-ospfv3-1] quit
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ipv6 enable
[RouterA-GigabitEthernet1/0/0] ospfv3 1 area 0
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 1/0/1
[RouterA-GigabitEthernet1/0/1] ipv6 enable
[RouterA-GigabitEthernet1/0/1] ospfv3 1 area 0.0.0.0
[RouterA-GigabitEthernet1/0/1] quit
# Configure RouterB.
[RouterB] ipv6 enable
[RouterB] ospfv3 1
[RouterB-ospfv3-1] router-id 2.2.2.2
[RouterB-ospfv3-1] quit
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] ipv6 enable
[RouterB-GigabitEthernet1/0/0] ospfv3 1 area 0.0.0.0
[RouterB-GigabitEthernet1/0/0] quit
[RouterB] interface gigabitethernet 1/0/1
[RouterB-GigabitEthernet1/0/1] ipv6 enable
[RouterB-GigabitEthernet1/0/1] ospfv3 1 area 0.0.0.0
[RouterB-GigabitEthernet1/0/1] quit
[RouterB] interface gigabitethernet 1/0/2
[RouterB-GigabitEthernet1/0/2] ipv6 enable
[RouterB-GigabitEthernet1/0/2] ospfv3 1 area 0.0.0.0
# Configure HuaweiC.
[RouterC] ospfv3 1
[RouterC-ospfv3-1] router-id 3.3.3.3
[RouterC-ospfv3-1] quit
[RouterC] interface gigabitethernet 1/0/0
[RouterC-GigabitEthernet1/0/0] ipv6 enable
[RouterC-GigabitEthernet1/0/0] ospfv3 1 area 0.0.0.0
[RouterC-GigabitEthernet1/0/0] quit
[RouterC] interface gigabitethernet 1/0/1
[RouterC-GigabitEthernet1/0/1] ipv6 enable
[RouterC-GigabitEthernet1/0/1] ospfv3 1 area 0.0.0.0
# After the preceding configurations are complete, run the display ospfv3 peer command. You can view that the neighboring relationship is set up between RouterA and RouterB, and that between RouterB and RouterC. Take the display of RouterA as an example:
[RouterA] display ospfv3 peer verbose
OSPFv3 Process (1)
Neighbor 2.2.2.2 is Full, interface address FE80::E0:CE19:8142:1
    In the area 0.0.0.0 via interface GE1/0/0
    DR Priority is 1 DR is 2.2.2.2 BDR is 1.1.1.1
    Options is 0x000013 (-|R|-|-|E|V6)
    Dead timer due in 00:00:34
    Neighbour is up for 01:30:52
    Database Summary Packets List 0
    Link State Request List 0
    Link State Retransmission List 0
    Neighbour Event: 6
    Neighbour If Id : 0xe
Neighbor 3.3.3.3 is Full, interface address FE80::E0:9C69:8142:2
    In the area 0.0.0.0 via interface GE1/0/1
    DR Priority is 1 DR is 3.3.3.3 BDR is 1.1.1.1
    Options is 0x000013 (-|R|-|-|E|V6)
    Dead timer due in 00:00:37
    Neighbour is up for 01:31:18
    Database Summary Packets List 0
    Link State Request List 0
    Link State Retransmission List 0
    Neighbour Event: 6
    Neighbour If Id : 0x9
# Display the information in the OSPFv3 routing table on RouterA. You can view the routing entries to RouterB and RouterC.
[RouterA] display ospfv3 routing
Codes : E2 - Type 2 External, E1 - Type 1 External, IA - Inter-Area,
N - NSSA, U - Uninstalled, D - Denied by Import Policy
OSPFv3 Process (1)
Destination                                                 Metric
  Next-hop
2001:DB8:1::/64                                               1
 directly connected, GigabitEthernet1/0/0
2001:DB8:2::/64                                               2
 via FE80::E0:9C69:8142:2, GigabitEthernet1/0/1
 via FE80::E0:CE19:8142:1, GigabitEthernet1/0/0
2001:DB8:3::/64                                               1
 directly connected, GigabitEthernet1/0/1
2001:DB8:4::1/64                                              1
 via FE80::E0:CE19:8142:1, GigabitEthernet1/0/0                
As shown in the OSPFv3 routing table, the next hop address of the route to 2001:DB8:4::1/64 is GigabitEthernet1/0/0 and traffic is transmitted on the active link RouterA â RouterB.
# Enable global BFD on RouterA.
[RouterA] bfd
[RouterA-bfd] quit
[RouterA-ospfv3-1] bfd all-interfaces enable
[RouterA-ospfv3-1] bfd all-interfaces min-transmit-interval 100 min-receive-interval 100 detect-multiplier 4
# Enable global BFD on RouterB.
[RouterB] bfd
[RouterB-bfd] quit
[RouterB] ospfv3
[RouterB-ospfv3-1] bfd all-interfaces enable
[RouterB-ospfv3-1] bfd all-interfaces min-transmit-interval 100 min-receive-interval 100 detect-multiplier 4
# Enable global BFD on RouterC.
[RouterC] bfd
[RouterC-bfd] quit
[RouterC] ospfv3
[RouterC-ospfv3-1] bfd all-interfaces enable
[RouterC-ospfv3-1] bfd all-interfaces min-transmit-interval 100 min-receive-interval 100 detect-multiplier 4
# After the preceding configurations are complete, run the display ospfv3 bfd session command on RouterA or RouterB. You can view that the status of the BFD session is Up.
Take the display of RouterB as an example:
<RouterB> display ospfv3 bfd session verbose
* - STALE
OSPFv3 Process (1)
   Neighbor-Id: 1.1.1.1
   BFD Status: Up
   Interface: GE1/0/0
   IPv6-Local-Address: FE80::E0:CE19:8142:1
   IPv6-Remote-Address: FE80::E0:4C3A:143:1
   BFD Module preferred timer values
      Transmit-Interval(ms): 100
      Receive-Interval(ms): 100
      Detect-Multiplier: 3
   OSPFv3 Module preferred timer values
      Transmit-Interval(ms): 100
      Receive-Interval(ms): 100
      Detect-Multiplier: 3
   Configured timer values
      Transmit-Interval(ms): 100
      Receive-Interval(ms): 100
      Detect-Multiplier: 3
   Neighbor-Id: 3.3.3.3
   BFD Status: Down
   Interface: GE1/0/1
   IPv6-Local-Address: FE80::E0:CE19:8142:2
   IPv6-Remote-Address: FE80::E0:9C69:8142:1
   BFD Module preferred timer values
      Transmit-Interval(ms): 2200
      Receive-Interval(ms): 2200
      Detect-Multiplier: 0
   OSPFv3 Module preferred timer values
      Transmit-Interval(ms): 1000
      Receive-Interval(ms): 1000
      Detect-Multiplier: 3
   Configured timer values
      Transmit-Interval(ms): 1000
      Receive-Interval(ms): 1000
      Detect-Multiplier: 3
# Run the shutdown command on GE 1/0/0 of RouterB to simulate the active link failure.
[RouterB] interface gigabitethernet1/0/0
[RouterB-GigabitEthernet1/0/0] shutdown
# Display the routing table on RouterA. The standby link RouterA â RouterC â RouterB takes effect after the active link fails. The next hop address of the route to 2001:DB8:4::1/64 becomes GigabitEthernet1/0/1.
<RouterA> display ospfv3 routing
Codes : E2 - Type 2 External, E1 - Type 1 External, IA - Inter-Area,
N - NSSA, U - Uninstalled, D - Denied by Import Policy
OSPFv3 Process (1)
Destination                                                 Metric
  Next-hop
2001:DB8:1::/64                                               1
       directly connected, GigabitEthernet/0/0
2001:DB8:2::/64                                               2
       via FE80::E0:9C69:8142:2, GigabitEthernet1/0/1
2001:DB8:3::/64                                               1
       directly connected, GigabitEthernet1/0/1
2001:DB8:4::1/64                                              2
       via FE80::E0:9C69:8142:2, GigabitEthernet1/0/1          
Configuration file of RouterA
#
 sysname RouterA
#
 ipv6
#
 bfd
#
ospfv3 1
 router-id 1.1.1.1
 bfd all-interfaces enable
 bfd all-interfaces min-transmit-interval 100 min-receive-interval 100 detect-multiplier 4
#
interface gigabitethernet1/0/0 
 ipv6 enable
