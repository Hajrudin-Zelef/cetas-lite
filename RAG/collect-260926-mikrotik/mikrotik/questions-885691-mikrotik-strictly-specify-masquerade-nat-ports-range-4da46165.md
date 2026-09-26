---
id: collect-260926-mikrotik/mikrotik/questions-885691-mikrotik-strictly-specify-masquerade-nat-ports-range-4da46165
title: "questions-885691-mikrotik-strictly-specify-masquerade-nat-ports-range-4da46165"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/questions-885691-mikrotik-strictly-specify-masquerade-nat-ports-range-4da46165.md
source_anchor: ""
source_lines: [1, 12]
sha256: b43ed74b87542bdcc2e90cba0db4e7730a7ccce958c411734c489fd2787d2532
---

# questions-885691-mikrotik-strictly-specify-masquerade-nat-ports-range-4da46165

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I want to limit TCP/UDP ports which can be used for masquerade NAT on my Mikrotik router. For example, I want to use only 40000-65535 TCP ports range and 20000-65535 UDP ports range for masquerade. Other ports will be used for DNAT from WAN.
In general I have 2 questions:
How can I strictly specify ports range which can be used for masquerade NAT firewall rule?
If no restriction specified does Mikrotik can skip using in masquerade ports which router opened itself (SSH, Telnet, WinBox, etc.)?
You can create 2 additional NAT rules for TCP/UDP. Set the first to match TCP packets going out to WAN and set the action to src-nat; Then specify the correct public address and port range. Do the same for UDP, then have the standard masquerade rule underneath to catch anything else.
It will not use open ports for NAT, and while I don't know the exact port range it uses, it will certainly not use low port numbers like 22/23. Not sure if it's clever enough to automatically avoid ports you have dst-nat rules set up for though.
