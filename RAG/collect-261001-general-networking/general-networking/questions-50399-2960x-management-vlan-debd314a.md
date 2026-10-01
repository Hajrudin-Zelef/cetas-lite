---
id: collect-261001-general-networking/general-networking/questions-50399-2960x-management-vlan-debd314a
title: "questions-50399-2960x-management-vlan-debd314a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-50399-2960x-management-vlan-debd314a.md
source_anchor: ""
source_lines: [1, 11]
sha256: 5101d177c5cafbcb1e3dfd35381e5f6dabd78e0e64f42f4520be6b48d7c26106
---

# questions-50399-2960x-management-vlan-debd314a

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
6
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Cisco ws-c2960x-24td-l running 15.2(2)E7
I have created a management vlan (500) and want to put the management interface (FE0) into it, however the "switchport" command is not available for this interface.
Can somebody let me know what I am doing wrong please?
You can disregard the dedicated management interface and simply assign an IP address to the VLAN 500 interface on the switch and use this IP for management purpose, with appropriate access lists.
FE0 is useful for out of band management. Nothing prevent you to perform in-band management without it.
