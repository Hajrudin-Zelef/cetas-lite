---
id: collect-261001-cisco/cisco/questions-6359-ospf-on-asa-8-2-not-advertising-32-a31c1cd1
title: "questions-6359-ospf-on-asa-8-2-not-advertising-32-a31c1cd1"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-6359-ospf-on-asa-8-2-not-advertising-32-a31c1cd1.md
source_anchor: ""
source_lines: [1, 15]
sha256: 45a73033a3f828ca169279097ef84ec2dcb1f96f26e8ab6d3af800e58d229f67
---

# questions-6359-ospf-on-asa-8-2-not-advertising-32-a31c1cd1

We have set up OSPF between our 2 ASA 5540 running in Active/Passive and our Internet provider's Juniper routers. We would like to advertise routes based on the content of an ACL through a route-map :
router ospf 1
router-id X.X.X.A
network X.X.X.B 255.255.255.248 area 0
area 0 authentication message-digest
log-adj-changes
redistribute static metric 10 subnets route-map annonce_ospf_isp
route-map annonce_ospf_isp permit 1
match ip address redistribute_isp
access-list redistribute_isp standard permit Y.Y.Y.C 255.255.255.252
If I advertise a /30, no problem : it shows up in the ospf database and the route is inserted in other routers.
However if I advertise a /32 then nothing happens: it's not in the database, it's not advertised :
access-list redistribute_neo standard permit host Y.Y.Y.D
What is the problem?
prefix-list PF-test-ospf seq 1 permit Y.Y.Y.D/32. However I cannot use it in my route-map :(config-route-map)# match ip address ? route-map mode commands/options: WORD Access-list name
