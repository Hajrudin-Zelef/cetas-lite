---
id: collect-261001-general-networking/general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas-2
title: "multi-area-ospf-packet-tracer-lab-area-0-standard-areas"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas.md
source_anchor: ""
source_lines: [229, 387]
sha256: 6908694a82ea1943072c36ee2c9d566032f16c08264aaa5d2095a26b34c0fc2e
---

# multi-area-ospf-packet-tracer-lab-area-0-standard-areas

We have completed our **Multi Area OSPF Packet Tracer Configuration**. Now, let’s verify the **Routing Tables, Topology Databases (OSPF Databases)** and **OSPF Neighbourhip** on all OSPF routers.

```
Router1#
```
**show ip ospf neighbor** 
Neighbor ID     Pri   State           Dead Time   Address         Interface
2.2.2.2           1   FULL/DR         00:00:38    10.1.0.2        GigabitEthernet0/0
3.3.3.3           1   FULL/DR         00:00:37    10.2.0.2        GigabitEthernet0/1

```
Router1#
```
**show ip ospf database** 
            OSPF Router with ID (1.1.1.1) (Process ID 1)
                Router Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum Link count
1.1.1.1         1.1.1.1         223         0x8000000b 0x0079af 1
5.5.5.5         5.5.5.5         802         0x80000005 0x009874 1
2.2.2.2         2.2.2.2         228         0x8000000b 0x00e604 2
                Net Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.2        2.2.2.2         228         0x80000005 0x001b1b
                Summary Net Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum
10.2.0.0        1.1.1.1         1009        0x80000006 0x00d078
10.3.0.0        2.2.2.2         909         0x8000000a 0x009ea1
10.4.0.0        2.2.2.2         884         0x8000000b 0x009aa2
10.4.0.0        5.5.5.5         812         0x8000000e 0x0030fe
10.3.0.0        5.5.5.5         802         0x8000000f 0x0044e9
10.5.0.0        5.5.5.5         482         0x80000010 0x00200c
                Router Link States (Area 1)
Link ID         ADV Router      Age         Seq#       Checksum Link count
1.1.1.1         1.1.1.1         224         0x8000000a 0x008f98 1
3.3.3.3         3.3.3.3         229         0x80000009 0x000415 1
                Net Link States (Area 1)
Link ID         ADV Router      Age         Seq#       Checksum
10.2.0.2        3.3.3.3         229         0x80000005 0x00f67e
                Summary Net Link States (Area 1)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.0        1.1.1.1         1008        0x80000013 0x00c27a
10.3.0.0        1.1.1.1         902         0x80000014 0x00b286
10.4.0.0        1.1.1.1         879         0x80000015 0x00ae87
10.5.0.0        1.1.1.1         477         0x80000016 0x00aa88

```
Router1#
```
**show ip route** 
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default, U - per-user static route, o - ODR
       P - periodic downloaded static route
Gateway of last resort is not set
     10.0.0.0/8 is variably subnetted, 7 subnets, 2 masks
C       10.1.0.0/24 is directly connected, GigabitEthernet0/0
L       10.1.0.1/32 is directly connected, GigabitEthernet0/0
C       10.2.0.0/24 is directly connected, GigabitEthernet0/1
L       10.2.0.1/32 is directly connected, GigabitEthernet0/1
O IA    10.3.0.0/24 [110/2] via 10.1.0.2, 02:04:54, GigabitEthernet0/0
O IA    10.4.0.0/24 [110/3] via 10.1.0.2, 02:04:54, GigabitEthernet0/0
O IA    10.5.0.0/24 [110/4] via 10.1.0.2, 02:04:54, GigabitEthernet0/0

```
Router2#
```
**show ip ospf neighbor** 
Neighbor ID     Pri   State           Dead Time   Address         Interface
1.1.1.1           1   FULL/BDR        00:00:39    10.1.0.1        GigabitEthernet0/0
5.5.5.5           0   FULL/  -        00:00:39    10.4.0.2        OSPF_VL0
3.3.3.3           1   FULL/DR         00:00:39    10.3.0.2        GigabitEthernet0/1

```
Router2#
```
**show ip ospf database**
            OSPF Router with ID (2.2.2.2) (Process ID 1)
                Router Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum Link count
2.2.2.2         2.2.2.2         349         0x8000000b 0x00e604 2
5.5.5.5         5.5.5.5         922         0x80000005 0x009874 1
1.1.1.1         1.1.1.1         344         0x8000000b 0x0079af 1
                Net Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.2        2.2.2.2         349         0x80000005 0x001b1b
                Summary Net Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum
10.3.0.0        2.2.2.2         1030        0x8000000a 0x009ea1
10.4.0.0        2.2.2.2         1005        0x8000000b 0x009aa2
10.2.0.0        1.1.1.1         1131        0x80000006 0x00d078
10.4.0.0        5.5.5.5         933         0x8000000e 0x0030fe
10.3.0.0        5.5.5.5         922         0x8000000f 0x0044e9
10.5.0.0        5.5.5.5         603         0x80000010 0x00200c
                Router Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum Link count
2.2.2.2         2.2.2.2         344         0x8000000a 0x0057c6 1
3.3.3.3         3.3.3.3         349         0x8000000d 0x001fc4 2
5.5.5.5         5.5.5.5         349         0x8000000a 0x00946e 1
                Net Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum
10.4.0.2        5.5.5.5         941         0x80000005 0x001739
10.3.0.2        3.3.3.3         349         0x80000006 0x003b8b
                Summary Net Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.0        2.2.2.2         1039        0x8000000b 0x00b48c
10.2.0.0        2.2.2.2         1039        0x8000000c 0x00b08d
10.5.0.0        5.5.5.5         603         0x80000013 0x001a0f
10.1.0.0        5.5.5.5         335         0x80000014 0x005ccd

```
Router2#
```
**show ip route**
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2, E - EGP
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
       * - candidate default, U - per-user static route, o - ODR
       P - periodic downloaded static route
Gateway of last resort is not set
     10.0.0.0/8 is variably subnetted, 7 subnets, 2 masks
C       10.1.0.0/24 is directly connected, GigabitEthernet0/0
L       10.1.0.2/32 is directly connected, GigabitEthernet0/0
O IA    10.2.0.0/24 [110/2] via 10.1.0.1, 02:06:09, GigabitEthernet0/0
C       10.3.0.0/24 is directly connected, GigabitEthernet0/1
L       10.3.0.1/32 is directly connected, GigabitEthernet0/1
O       10.4.0.0/24 [110/2] via 10.3.0.2, 02:15:59, GigabitEthernet0/1
O IA    10.5.0.0/24 [110/3] via 10.3.0.2, 02:10:17, GigabitEthernet0/1

```
Router3#
```
**show ip ospf neighbor** 
Neighbor ID     Pri   State           Dead Time   Address         Interface
1.1.1.1           1   FULL/BDR        00:00:34    10.2.0.1        GigabitEthernet0/1

```
Router3#
```
**show ip ospf database**
            OSPF Router with ID (3.3.3.3) (Process ID 1)
                Router Link States (Area 1)
Link ID         ADV Router      Age         Seq#       Checksum Link count
3.3.3.3         3.3.3.3         394         0x80000009 0x000415 1
1.1.1.1         1.1.1.1         389         0x8000000a 0x008f98 1
                Net Link States (Area 1)
Link ID         ADV Router      Age         Seq#       Checksum
10.2.0.2        3.3.3.3         394         0x80000005 0x00f67e
                Summary Net Link States (Area 1)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.0        1.1.1.1         1174        0x80000013 0x00c27a
10.3.0.0        1.1.1.1         1067        0x80000014 0x00b286
10.4.0.0        1.1.1.1         1045        0x80000015 0x00ae87
10.5.0.0        1.1.1.1         642         0x80000016 0x00aa88

