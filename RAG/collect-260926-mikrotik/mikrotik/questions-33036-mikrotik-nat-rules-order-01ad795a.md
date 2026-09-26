---
id: collect-260926-mikrotik/mikrotik/questions-33036-mikrotik-nat-rules-order-01ad795a
title: "questions-33036-mikrotik-nat-rules-order-01ad795a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-33036-mikrotik-nat-rules-order-01ad795a.md
source_anchor: ""
source_lines: [1, 11]
sha256: 7740f94a8185426d0281d07107ed44ef076aaa78f60f81c567d2266b0bfa5c19
---

# questions-33036-mikrotik-nat-rules-order-01ad795a

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have Mikrotik router, which has WAN connected at ether1. Now I need to set up NAT to send TCP and UDP on port 25565 to IP 192.168.2.42 and every other port (all protocols) to 192.168.2.41
This is my current configuration:
What I am looking for is to descrease priority of rule #1 and increase priority of rule #2 and #3.
How do I do this? If there are no priorities, what is the correct way of this implementation?
The rules are tested in order. It appears that drag-and-drop does not work with this, and it seems that you need to delete the rules and recreate them in the correct order.
