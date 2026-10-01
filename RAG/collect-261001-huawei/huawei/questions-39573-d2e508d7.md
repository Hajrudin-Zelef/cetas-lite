---
id: collect-261001-huawei/huawei/questions-39573-d2e508d7
title: "questions-39573-d2e508d7"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-39573-d2e508d7.md
source_anchor: ""
source_lines: [1, 16]
sha256: c0a513aa27b0eb189a31f40b1eb679cae301529ddf49500ebec307773445fc70
---

# questions-39573-d2e508d7

three routers are connected where
  Router A connected to router B
A (gig0/0/0  ip address 10.10.1.1)====> B(gig0/0/0 ip address 10.10.1.2    
A(gig0/0/0=> ip address 10.11.1.1===>B(gig0/0/1  ip address 10.11.1.2))
Router A connected toRouter C
(A (gig0/0/2=> ip address 10.12.1.1 )====>  C (gig0/0/0  ip address 10.12.1.2)
An IS-IS process is running among nodes. At Router A, Can apply a route policy so that link between Router A to Router B only one link active for certain routes based on next hope Address match condition. for example
if routes 157.1.1.1/32, 57.1.1.1/32, 47.1.1.1/32 has two next hop 10.10.1.1 and 10.11.1.1 for Same destination Router B.how can i apply effectively in ISIS process the below policy:-
ip ip-prefix Next_add index 10 permit 10.11.1.0 30 
ip ip-prefix rouisis index 10 permit 157.1.1.1 32
ip ip-prefix rouisis index 20 permit 57.1.1.1 32  
ip ip-prefix rouisis index 30 permit 47.1.1.1 32 
route-policy isispref permit node 20  
if-match ip-prefix rouisis 
if-match ip next-hop ip-prefix Next_add
apply preference 200
