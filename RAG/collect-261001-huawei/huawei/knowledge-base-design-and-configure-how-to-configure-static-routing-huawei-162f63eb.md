---
id: collect-261001-huawei/huawei/knowledge-base-design-and-configure-how-to-configure-static-routing-huawei-162f63eb
title: "knowledge-base-design-and-configure-how-to-configure-static-routing-huawei-162f63eb"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "United States"]
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/knowledge-base-design-and-configure-how-to-configure-static-routing-huawei-162f63eb.md
source_anchor: ""
source_lines: [1, 43]
sha256: d656dc09bb96321406bb0d1ecaf6714550b6dcec25e494f5ad4621b78f73f4d7
---

# knowledge-base-design-and-configure-how-to-configure-static-routing-huawei-162f63eb

Poland
GRANDMETRIC Sp. z o.o.
ul. Metalowa 5, 60-118 Poznań, Poland
NIP 7792433527
+48 61 271 04 43
info@grandmetric.com
UK
Grandmetric LTD
Office 584b
182-184 High Street North
London
E6 2JA
+44 20 3321 5276
info@grandmetric.com
US Region
Grandmetric LLC
Lewes DE 19958
16192 Coastal Hwy USA
EIN: 98-1615498
+1 302 691 94 10 
info@grandmetric.com
Technology: Network Services
Area: Static Routing
Vendor: Huawei
Software: eNSP, Quidway software
Platform: Huawei routers
Static routing is a form of routing that occurs when a router uses a manually-configured routing entry, rather than information from a dynamic routing protocols. Static routing can also be used in stub networks, or to provide a gateway of last resort..
To configure IP Static Route, use the following command:
[R1] ip route-static 10.0.3.0 24 10.0.13.3
To verify the routing table:
<HUAWEI> display ip routing-table
Route Flags: R - relay, D - download for forwarding
--------------------------------------------------------------
Routing Tables: Public
Destinations : 8 Routes : 8
Destination/Mask Proto  Pre            Cost                Flags      NextHop            Interface
0.0.0.0/0             Static 60                      0                RD            10.0.3.0                GigabitEthernet0/0/1
1.1.1.0/24           Direct 0                       0                      D            1.1.1.1                   GigabitEthernet0/0/0
1.1.1.1/32           Direct 0                       0                     D             127.0.0.1                InLoopBack0
1.1.4.0/30           Direct 0                       0                     D             10.0.10.1                Inloopback0
1.1.4.1/32           Direct 0                       0                     D             127.0.0.1               InLoopBack0
127.0.0.0/8         Direct 0                       0                    D             127.0.0.1              InLoopBack0
127.0.0.1/32 Direct 0                             0                    D             127.0.0.1              InLoopBack0
