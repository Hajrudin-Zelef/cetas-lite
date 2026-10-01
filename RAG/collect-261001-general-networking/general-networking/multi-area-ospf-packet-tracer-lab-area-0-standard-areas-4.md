---
id: collect-261001-general-networking/general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas-4
title: "multi-area-ospf-packet-tracer-lab-area-0-standard-areas"
domain: general-networking
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/multi-area-ospf-packet-tracer-lab-area-0-standard-areas.md
source_anchor: ""
source_lines: [539, 639]
sha256: 255bddcd7d58d4e11ce6d5cf0c6789f154c405d75fd4ee5b9ad76300eb685b24
---

# multi-area-ospf-packet-tracer-lab-area-0-standard-areas

```
Router6#
```
**show ip ospf neighbor** 
Neighbor ID     Pri   State           Dead Time   Address         Interface
5.5.5.5           1   FULL/BDR        00:00:37    10.5.0.1        GigabitEthernet0/1
Router6#**show ip ospf database**
            OSPF Router with ID (6.6.6.6) (Process ID 1)
                Router Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum Link count
6.6.6.6         6.6.6.6         520         0x80000008 0x005e9d 1
5.5.5.5         5.5.5.5         522         0x8000000c 0x009768 1
                Net Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum
10.5.0.2        6.6.6.6         714         0x80000005 0x00330c
                Summary Net Link States (Area 3)
Link ID         ADV Router      Age         Seq#       Checksum
10.1.0.0        5.5.5.5         1076        0x80000010 0x0064c9

```
Router6#
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
     10.0.0.0/8 is variably subnetted, 3 subnets, 2 masks
O IA    10.1.0.0/24 [110/4] via 10.5.0.1, 02:12:20, GigabitEthernet0/1
C       10.5.0.0/24 is directly connected, GigabitEthernet0/1
L       10.5.0.2/32 is directly connected, GigabitEthernet0/1

You can also used **“show ip route ospf”** command for displaying only OSPF routes like below:

```
Router1#
```
**show ip route ospf** 
     10.0.0.0/8 is variably subnetted, 7 subnets, 2 masks
O IA    10.3.0.0 [110/2] via 10.1.0.2, 02:14:52, GigabitEthernet0/0
O IA    10.4.0.0 [110/3] via 10.1.0.2, 02:14:52, GigabitEthernet0/0
O IA    10.5.0.0 [110/4] via 10.1.0.2, 02:14:52, GigabitEthernet0/0

Lastly, let’s check the virtual-links.

```
Router2#
```
**show ip ospf virtual-links** 
Virtual Link OSPF_VL0 to router 5.5.5.5 is up
  Run as demand circuit
  Transit area 2, via interface GigabitEthernet0/1, Cost of using 2
  Transmit Delay is 1 sec, State POINT_TO_POINT,
  Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    Hello due in 00:00:09
    Adjacency State FULL
    Index 1/2, retransmission queue length 0, number of retransmission 0
        First 0x0(0)/0x0(0) Next 0x0(0)/0x0(0)
        Last retransmission scan length is 0, maximum is 0
        Last retransmission scan time is 0 msec, maximum is 0 msec

```
Router5#
```
**show ip ospf virtual-links** 
Virtual Link OSPF_VL0 to router 2.2.2.2 is up
  Run as demand circuit
  Transit area 2, via interface GigabitEthernet0/0, Cost of using 2
  Transmit Delay is 1 sec, State POINT_TO_POINT,
  Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    Hello due in 00:00:00
    Adjacency State FULL
    Index 1/2, retransmission queue length 0, number of retransmission 0
        First 0x0(0)/0x0(0) Next 0x0(0)/0x0(0)
        Last retransmission scan length is 0, maximum is 0
        Last retransmission scan time is 0 msec, maximum is 0 msec

In this post, we have talked about Standard OSPF Areas and Backbone Areas, beside Virtual-Links. We focused on configuration mostly. I hope this post will be helpful for you.

In the next post we will focus on other OSPF Area Types. Keep on, IPcisco.com ;)

You can **DOWNLOAD** the **Packet Tracer** example with **.pkt** format  **HERE**.




**You can download “Packet Tracer” in Tools section.**

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
