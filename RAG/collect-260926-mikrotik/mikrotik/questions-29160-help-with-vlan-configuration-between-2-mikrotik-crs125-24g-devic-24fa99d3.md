---
id: collect-260926-mikrotik/mikrotik/questions-29160-help-with-vlan-configuration-between-2-mikrotik-crs125-24g-devic-24fa99d3
title: "questions-29160-help-with-vlan-configuration-between-2-mikrotik-crs125-24g-devic-24fa99d3"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-29160-help-with-vlan-configuration-between-2-mikrotik-crs125-24g-devic-24fa99d3.md
source_anchor: ""
source_lines: [1, 11]
sha256: 9c6b5c29de5181ce51a002825a293fe76d290f09f33b977f0d34992cc8708a27
---

# questions-29160-help-with-vlan-configuration-between-2-mikrotik-crs125-24g-devic-24fa99d3

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a SonicWALL with X0=10.11.10.0/24 and X2=192.168.11.0/24. The downstairs MikroTik ports 1-19 are in 192.168.11.0/24, and ports 21-24 are in 10.11.10.0/24 I am assuming port 20 will be my trunk port to the 2nd MikroTik upstairs. on the 2nd MikroTik, I want port 1 to be the trunk, with Ports 2-20 on 192.168.11.0/24, and ports 21-24 on 10.11.10.0/24.
The reason I am trying to do it this way is that this is a business in an older home, and I only have one Category-5E cable that is between my downstairs MikroTik and the upstairs MikroTik.
Is this doable, and if so, can someone help me out on the configurations, as I am new to the VLAN game?
Basically, what I would do is reset the switch and make all ports part of the same switch with ether1 being the master port. Then I would specify VLAN for separation and proper tagging for trunk ports. This way you would have full speed switching.
Generic VLAN stuff on Mikrotik is CPU bound so while it would work on CRS it would not be very fast.
