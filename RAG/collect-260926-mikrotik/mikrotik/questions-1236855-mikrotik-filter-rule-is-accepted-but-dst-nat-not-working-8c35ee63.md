---
id: collect-260926-mikrotik/mikrotik/questions-1236855-mikrotik-filter-rule-is-accepted-but-dst-nat-not-working-8c35ee63
title: "questions-1236855-mikrotik-filter-rule-is-accepted-but-dst-nat-not-working-8c35ee63"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-1236855-mikrotik-filter-rule-is-accepted-but-dst-nat-not-working-8c35ee63.md
source_anchor: ""
source_lines: [1, 8]
sha256: e78f93d934b73c0a633d0066f1db41746945d9447c9fa80d79a09d81b48ae12f
---

# questions-1236855-mikrotik-filter-rule-is-accepted-but-dst-nat-not-working-8c35ee63

I have an NAT issue with mikrotik CCR1016-12G - namely I can't see any packets going to my dst-nat rules and dst-nat rules are not working.
I have tried disabling filtering rules that might drop any packets, and I've also tried to make a specific rule to accept a specific port with no luck:
chain=input action=accept connection-state=new protocol=tcp 
src-address=0.0.0.0/0 in-interface=ether01-gateway dst-port=8088 
I can see the packets coming in to the filter rules and assume that they get accepted(as I have the previous rule in place), but no packets are going to the dst-nat part:
chain=dstnat action=dst-nat to-addresses=192.168.1.88 to-ports=80 
protocol=tcp src-address=0.0.0.0 dst-address="Outer-IP" dst port=8088   
What could be the issue that I don't see any packets going to NAT?
