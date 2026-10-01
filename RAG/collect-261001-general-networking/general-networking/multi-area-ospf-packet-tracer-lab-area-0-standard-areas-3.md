---
id: collect-261001-general-networking/general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas-3
title: "multi-area-ospf-packet-tracer-lab-area-0-standard-areas"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas.md
source_anchor: ""
source_lines: [388, 538]
sha256: d69906b2f2d18cedb8547a75c50d195127b8f93defaaff209dc4c0d26ad3f0f4
---

# multi-area-ospf-packet-tracer-lab-area-0-standard-areas

```
Router3#
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
     10.0.0.0/8 is variably subnetted, 6 subnets, 2 masks
O IA    10.1.0.0/24 [110/2] via 10.2.0.1, 02:47:08, GigabitEthernet0/1
C       10.2.0.0/24 is directly connected, GigabitEthernet0/1
L       10.2.0.2/32 is directly connected, GigabitEthernet0/1
O IA    10.3.0.0/24 [110/3] via 10.2.0.1, 02:47:08, GigabitEthernet0/1
O IA    10.4.0.0/24 [110/4] via 10.2.0.1, 02:17:52, GigabitEthernet0/1
O IA    10.5.0.0/24 [110/5] via 10.2.0.1, 02:11:08, GigabitEthernet0/1
```
Router4#
```
**show ip ospf neighbor** 
Neighbor ID     Pri   State           Dead Time   Address         Interface
5.5.5.5           1   FULL/DR         00:00:36    10.4.0.2        GigabitEthernet0/0
2.2.2.2           1   FULL/BDR        00:00:38    10.3.0.1        GigabitEthernet0/1
```
Router4#
```
**show ip ospf database**
            OSPF Router with ID (3.3.3.3) (Process ID 1)
                Router Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum Link count
3.3.3.3         3.3.3.3         429         0x8000000d 0x001fc4 2
5.5.5.5         5.5.5.5         429         0x8000000a 0x00946e 1
2.2.2.2         2.2.2.2         425         0x8000000a 0x0057c6 1
                Net Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum
10.3.0.2        3.3.3.3         429         0x80000006 0x003b8b
10.4.0.2        5.5.5.5         1021        0x80000005 0x001739
                Summary Net Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.0        2.2.2.2         1119        0x8000000b 0x00b48c
10.2.0.0        2.2.2.2         1119        0x8000000c 0x00b08d
10.5.0.0        5.5.5.5         684         0x80000013 0x001a0f
10.1.0.0        5.5.5.5         415         0x80000014 0x005ccd
                Router Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum Link count
3.3.3.3         3.3.3.3         1113        0x80000007 0x00ac9a 0
                Net Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum
10.4.0.1        3.3.3.3         764         0x80000006 0x003b1d

```
Router4#
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
O IA    10.1.0.0/24 [110/2] via 10.3.0.1, 02:46:14, GigabitEthernet0/1
O IA    10.2.0.0/24 [110/3] via 10.3.0.1, 02:46:14, GigabitEthernet0/1
C       10.3.0.0/24 is directly connected, GigabitEthernet0/1
L       10.3.0.2/32 is directly connected, GigabitEthernet0/1
C       10.4.0.0/24 is directly connected, GigabitEthernet0/0
L       10.4.0.1/32 is directly connected, GigabitEthernet0/0
O IA    10.5.0.0/24 [110/2] via 10.4.0.2, 02:11:46, GigabitEthernet0/0

```
Router5#
```
**show ip ospf neighbor** 
Neighbor ID     Pri   State           Dead Time   Address         Interface
6.6.6.6           1   FULL/DR         00:00:31    10.5.0.2        GigabitEthernet0/1
2.2.2.2           0   FULL/  -        00:00:31    10.3.0.1        OSPF_VL0
3.3.3.3           1   FULL/BDR        00:00:31    10.4.0.1        GigabitEthernet0/0

```
Router5#
```
**show ip ospf database**
            OSPF Router with ID (5.5.5.5) (Process ID 1)
                Router Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum Link count
5.5.5.5         5.5.5.5         1059        0x80000005 0x009874 1
2.2.2.2         2.2.2.2         487         0x8000000b 0x00e604 2
1.1.1.1         1.1.1.1         481         0x8000000b 0x0079af 1
                Net Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.2        2.2.2.2         487         0x80000005 0x001b1b
                Summary Net Link States (Area 0)
Link ID         ADV Router      Age         Seq#       Checksum
10.4.0.0        5.5.5.5         1069        0x8000000e 0x0030fe
10.3.0.0        5.5.5.5         1059        0x8000000f 0x0044e9
10.5.0.0        5.5.5.5         739         0x80000010 0x00200c
10.2.0.0        1.1.1.1         1268        0x80000006 0x00d078
10.3.0.0        2.2.2.2         1168        0x8000000a 0x009ea1
10.4.0.0        2.2.2.2         1142        0x8000000b 0x009aa2
                Router Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum Link count
5.5.5.5         5.5.5.5         485         0x8000000a 0x00946e 1
3.3.3.3         3.3.3.3         486         0x8000000d 0x001fc4 2
2.2.2.2         2.2.2.2         481         0x8000000a 0x0057c6 1
                Net Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum
10.4.0.2        5.5.5.5         1076        0x80000005 0x001739
10.3.0.2        3.3.3.3         486         0x80000006 0x003b8b
                Summary Net Link States (Area 2)
Link ID         ADV Router      Age         Seq#       Checksum
10.5.0.0        5.5.5.5         739         0x80000013 0x001a0f
10.1.0.0        2.2.2.2         1175        0x8000000b 0x00b48c
10.2.0.0        2.2.2.2         1175        0x8000000c 0x00b08d
10.1.0.0        5.5.5.5         471         0x80000014 0x005ccd
                Router Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum Link count
5.5.5.5         5.5.5.5         486         0x8000000c 0x009768 1
6.6.6.6         6.6.6.6         485         0x80000008 0x005e9d 1
                Net Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum
10.5.0.2        6.6.6.6         679         0x80000005 0x00330c
                Summary Net Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.0        5.5.5.5         1040        0x80000010 0x0064c9

```
Router5#
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
O       10.1.0.0/24 [110/3] via 10.4.0.1, 01:38:19, GigabitEthernet0/0
O IA    10.2.0.0/24 [110/4] via 10.4.0.1, 01:38:09, GigabitEthernet0/0
O       10.3.0.0/24 [110/2] via 10.4.0.1, 02:08:15, GigabitEthernet0/0
C       10.4.0.0/24 is directly connected, GigabitEthernet0/0
L       10.4.0.2/32 is directly connected, GigabitEthernet0/0
C       10.5.0.0/24 is directly connected, GigabitEthernet0/1
L       10.5.0.1/32 is directly connected, GigabitEthernet0/1

