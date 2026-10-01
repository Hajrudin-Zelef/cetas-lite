---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ospfv3-0055-html-d609fc1e
title: "Configure IPv6 addresses for interfaces on RouterA."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ospfv3-0055-html-d609fc1e.md
source_anchor: ""
source_lines: [1, 215]
sha256: 6a9279b150d4c5f0c427d1dc0a71eedcbca06cfddf4c969b10a4ee4420762187
---

# Configure IPv6 addresses for interfaces on RouterA.

As shown in Figure 1, all routers run OSPFv3. The entire autonomous system is divided into three areas. RouterB and RouterC serve as ABRs to forward the inter-area routes.
It is required that Area 2 be configured to decrease the LSAs advertised to this area, without affecting route reachability.
The configuration roadmap is as follows:
Enable basic OSPFv3 function on each router.
Configure Area 2 as a stub area to decrease the LSAs advertised to this area, without affecting route reachability.
# Configure IPv6 addresses for interfaces on RouterA.
<Huawei> system-view
[Huawei] sysname RouterA
[RouterA] ipv6
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ipv6 enable
[RouterA-GigabitEthernet1/0/0] ipv6 address 2000::1/64
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 2/0/0
[RouterA-GigabitEthernet2/0/0] ipv6 enable
[RouterA-GigabitEthernet2/0/0] ipv6 address 1001::2/64
The configurations of RouterB, RouterC and RouterD are similar to the configuration of RouterA, and are not mentioned here.
# Configure RouterA.
[RouterA] ipv6
[RouterA] ospfv3
[RouterA-ospfv3-1] router-id 1.1.1.1
[RouterA-ospfv3-1] quit
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ospfv3 1 area 1
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 2/0/0
[RouterA-GigabitEthernet2/0/0] ospfv3 1 area 1
[RouterA-GigabitEthernet2/0/0] quit
# Configure RouterB.
[RouterB] ipv6
[RouterB] ospfv3
[RouterB-ospfv3-1] router-id 2.2.2.2
[RouterB-ospfv3-1] quit
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] ospfv3 1 area 0
[RouterB-GigabitEthernet1/0/0] quit
[RouterB] interface gigabitethernet 2/0/0
[RouterB-GigabitEthernet2/0/0] ospfv3 1 area 1
[RouterB-GigabitEthernet2/0/0] quit
# Configure RouterC.
[RouterC] ipv6
[RouterC] ospfv3
[RouterC-ospfv3-1] router-id 3.3.3.3
[RouterC-ospfv3-1] quit
[RouterC] interface gigabitethernet 1/0/0
[RouterC-GigabitEthernet1/0/0] ospfv3 1 area 0
[RouterC-GigabitEthernet1/0/0] quit
[RouterC] interface gigabitethernet 2/0/0
[RouterC-GigabitEthernet2/0/0] ospfv3 1 area 2
[RouterC-GigabitEthernet2/0/0] quit
# Configure RouterD.
[RouterD] ipv6
[RouterD] ospfv3
[RouterD-ospfv3-1] router-id 4.4.4.4
[RouterD-ospfv3-1] quit
[RouterD] interface gigabitethernet 1/0/0
[RouterD-GigabitEthernet1/0/0] ospfv3 1 area 2
[RouterD-GigabitEthernet1/0/0] quit
# Display the OSPFv3 neighbors of RouterB.
[RouterB] display ospfv3 peer
OSPFv3 Process (1)
OSPFv3 Area (0.0.0.1)
Neighbor ID     Pri  State            Dead Time Interface            Instance ID
1.1.1.1           1  Full/ -          00:00:34  GE2/0/0                        0
OSPFv3 Area (0.0.0.0)
Neighbor ID     Pri  State            Dead Time Interface            Instance ID
3.3.3.3           1  Full/ -          00:00:32  GE1/0/0                        0
# Display OSPFv3 neighbors of RouterC.
[RouterC] display ospfv3 peer
OSPFv3 Process (1)
OSPFv3 Area (0.0.0.0)
Neighbor ID     Pri  State            Dead Time  Interface            Instance ID
2.2.2.2           1  Full/ -          00:00:37   GE1/0/0                        0
OSPFv3 Area (0.0.0.2)
Neighbor ID     Pri  State            Dead Time  Interface            Instance ID
4.4.4.4           1  Full/ -          00:00:33   GE2/0/0                        0
# Display the OSPFv3 routing table of RouterD.
[RouterD] display ospfv3 routing
Codes : E2 - Type 2 External, E1 - Type 1 External, IA - Inter-Area,
N - NSSA, U - Uninstalled
OSPFv3 Process (1)
   Destination                                   Metric
     Next-hop
  IA 1000::/64                                     2
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
  IA 1001::/64                                     3
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
     1002::/64                                     1
           directly-connected, GigabitEthernet1/0/0
  IA 2000::/64                                     4
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
# Configure the stub area of RouterD.
[RouterD] ospfv3
[RouterD-ospfv3-1] area 2
[RouterD-ospfv3-1-area-0.0.0.2] stub
[RouterD-ospfv3-1-area-0.0.0.2] quit
# Configure the stub area of RouterC, and set the cost of the default route advertised to the stub area to 10.
[RouterC] ospfv3
[RouterC-ospfv3-1] area 2
[RouterC-ospfv3-1-area-0.0.0.2] stub
[RouterC-ospfv3-1-area-0.0.0.2] default-cost 10
[RouterC-ospfv3-1-area-0.0.0.2] quit
# Display the OSPFv3 routing table of RouterD, and you can view a new default route in the routing table. Its cost is the sum of the cost of the directly connected routes and the configured cost.
Codes : E2 - Type 2 External, E1 - Type 1 External, IA - Inter-Area,
N - NSSA, U - Uninstalled
OSPFv3 Process (1)
OSPFv3 Process (1)
   Destination                                   Metric
     Next-hop
  IA ::/0                                          11
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
  IA 1000::/64                                     2
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
  IA 1001::/64                                     3
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
     1002::/64                                     1
           directly-connected, GigabitEthernet1/0/0
  IA 2000::/64                                     4
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
# Configure RouterC and configure Area 2 as a totally stub area.
[RouterC] ospfv3
[RouterC-ospfv3-1] area 2
[RouterC-ospfv3-1-area-0.0.0.2] stub no-summary
[RouterC-ospfv3-1-area-0.0.0.2] quit
# Display the OSPFv3 routing table of RouterD, and you can view that the entries in the routing table decrease; other non-directly connected routes are suppressed; only the default route is reserved.
Codes : E2 - Type 2 External, E1 - Type 1 External, IA - Inter-Area,
N - NSSA, U - Uninstalled
OSPFv3 Process (1)
OSPFv3 Process (1)
   Destination                                   Metric
     Next-hop
  IA ::/0                                          11
           via FE80::1572:0:5EF4:1, GigabitEthernet1/0/0
     1002::/64                                     1
           directly-connected, GigabitEthernet1/0/0
Configuration file of RouterA
#
 sysname RouterA
#
ipv6
#
interface GigabitEthernet1/0/0
 ipv6 enable
 ipv6 address 2000::1/64
 ospfv3 1 area 0.0.0.1
#
interface GigabitEthernet2/0/0
 ipv6 enable
 ipv6 address 1001::2/64
 ospfv3 1 area 0.0.0.1
#
ospfv3 1
 router-id 1.1.1.1
#
return
Configuration file of RouterB
#
 sysname RouterB
#
ipv6
#
interface GigabitEthernet1/0/0
 ipv6 enable
 ipv6 address 1000::1/64
 ospfv3 1 area 0.0.0.0
#
interface GigabitEthernet2/0/0
 ipv6 enable
 ipv6 address 1001::1/64
 ospfv3 1 area 0.0.0.1
#
ospfv3 1
 router-id 2.2.2.2
#
return
Configuration file of RouterC
#
 sysname RouterC
#
ipv6
#
interface GigabitEthernet1/0/0
 ipv6 enable
 ipv6 address 1000::2/64
 ospfv3 1 area 0.0.0.0
#
interface GigabitEthernet2/0/0
 ipv6 enable
 ipv6 address 1002::1/64
 ospfv3 1 area 0.0.0.2
#
ospfv3 1
 router-id 3.3.3.3
 area 0.0.0.2
  stub no-summary
  default-cost 10
#
return
Configuration file of RouterD
#
 sysname RouterD
#
ipv6
#
interface GigabitEthernet1/0/0
 ipv6 enable
 ipv6 address 1002::2/64
 ospfv3 1 area 0.0.0.2
#
ospfv3 1
 router-id 4.4.4.4
 area 0.0.0.2
  stub
#
return
