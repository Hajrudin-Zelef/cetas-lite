---
id: collect-261001-general-networking/general-networking/questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing-60a06b3c-2
title: "questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing--60a06b3c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing--60a06b3c.md
source_anchor: ""
source_lines: [164, 198]
sha256: cde59ec022f99467d3db4a51bd1fac663e485160de9af05c59b69c9b6fd4228a
---

# questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing--60a06b3c

Loopback20                 192.168.20.1    YES NVRAM  up                    up
Loopback25                 192.168.25.1    YES NVRAM  up                    up
Loopback30                 192.168.30.1    YES NVRAM  up                    up
Loopback35                 192.168.35.1    YES NVRAM  up                    up
Loopback40                 192.168.40.1    YES NVRAM  up                    up
R3#show ip route eigrp
R3#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is not set
C    192.168.30.0/24 is directly connected, Loopback30
C    192.168.8.0/24 is directly connected, Loopback8
C    192.168.25.0/24 is directly connected, Loopback25
C    192.168.9.0/24 is directly connected, Loopback9
C    192.168.10.0/24 is directly connected, Loopback10
C    192.168.40.0/24 is directly connected, Loopback40
     172.16.0.0/24 is subnetted, 5 subnets
C       172.16.23.0 is directly connected, Serial0/0
O E2    172.16.12.0 [110/20] via 172.16.23.2, 00:35:47, Serial0/0
O E2    172.16.1.0 [110/20] via 172.16.23.2, 00:35:49, Serial0/0
O E2    172.16.2.0 [110/20] via 172.16.23.2, 00:35:49, Serial0/0
C       172.16.3.0 is directly connected, Loopback0
C    192.168.11.0/24 is directly connected, Loopback11
C    192.168.20.0/24 is directly connected, Loopback20
O E2 192.168.51.0/24 [110/20] via 172.16.23.2, 00:35:50, Serial0/0
O E2 192.168.50.0/24 [110/20] via 172.16.23.2, 00:35:50, Serial0/0
C    192.168.35.0/24 is directly connected, Loopback35
O E2 192.168.70.0/24 [110/20] via 172.16.23.2, 00:35:50, Serial0/0
O    192.168.8.0/22 is a summary, 00:55:56, Null0
O E2 192.168.48.0/23 [110/20] via 172.16.23.2, 00:35:50, Serial0/0
