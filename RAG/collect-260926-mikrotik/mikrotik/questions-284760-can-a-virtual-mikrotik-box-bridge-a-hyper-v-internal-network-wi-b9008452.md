---
id: collect-260926-mikrotik/mikrotik/questions-284760-can-a-virtual-mikrotik-box-bridge-a-hyper-v-internal-network-wi-b9008452
title: "questions-284760-can-a-virtual-mikrotik-box-bridge-a-hyper-v-internal-network-wi-b9008452"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-284760-can-a-virtual-mikrotik-box-bridge-a-hyper-v-internal-network-wi-b9008452.md
source_anchor: ""
source_lines: [1, 8]
sha256: 68d51bad71a8b5f98af68277f4d0df51ad356b1ab667cba21777e9afa16a4545
---

# questions-284760-can-a-virtual-mikrotik-box-bridge-a-hyper-v-internal-network-wi-b9008452

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am trying to set up a Mikrotik router as a transparent firewall on my network. I got the machine working on a hardware MT box, but my boss wants the MT virtualized. I have been trying the set up where my virtual windows box talks to the Mikrotik via private or internal network on the Hyper-V host. I can get the two machines to talk, but as soon as I set up a bridge on the MT, all traffic ceases between the two. Is it possible to create a bridge for this purpose (having the MT silently in front of my firewalled server)?
Yes, it can. Mikrotik can run as bridge or router between different NIC's - if you run CHR (the virtual mikrotik you should use) on a VM and have multiple virtual NIC that point to the external and internal network - you got yourself a high speed connection. WAY higher speed than you can do at the moment with Mikrotik hardware, outside of pure bridging (switches now can handle 40g on mikrotik - that is hard to beat - but as router.... use that myself).
