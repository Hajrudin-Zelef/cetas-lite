---
id: collect-260926-mikrotik/mikrotik/questions-1778589-how-to-reset-port-forwarding-without-a-factory-reset-on-a-mikr-8d78bfb2
title: "questions-1778589-how-to-reset-port-forwarding-without-a-factory-reset-on-a-mikr-8d78bfb2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/troubleshooting/questions-1778589-how-to-reset-port-forwarding-without-a-factory-reset-on-a-mikr-8d78bfb2.md
source_anchor: ""
source_lines: [1, 11]
sha256: dd5ea99e104a183fc7c177f8b42c14628e0580445d604fffd2f2218e7b25e2ea
---

# questions-1778589-how-to-reset-port-forwarding-without-a-factory-reset-on-a-mikr-8d78bfb2

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I was working on a network, that goes something like this: tp-link router -> mikrotik router -> splitter for computers. And it has wifi from both routers. I accidentally forwarded ports(the 80 port specifically) on mikrotik in such a way, that it redirects most ips(not all) to 192.168.1.1, which is tp-link settings, the mikrotik settings were on ip-address 192.168.2.1, which redirects to 192.168.1.1 too. So I cannot access the mikrotik settings. Same happens when connecting to it by wifi. Can I reset the port forwarding or fix this problem in some other way without doing a factory reset, as it will probably break a lot of things and dissrupt the workflow for a lot of people.
Sorry for being such a noob, because I used the wrong port.
Use the other configuration interfaces that RouterOS supports. If you lost HTTP access, connect using Winbox; if you can't connect using Winbox, use SSH or even Telnet (whichever is enabled).
If regular access via TCP/IP doesn't work at all, connect via MAC-Winbox (by selecting the device's MAC address in Winbox neighbor list, instead of the IP address) or via MAC-Telnet (/tool/mac-telnet from another RouterOS device).
(If you have multiple RouterOS devices, consider enabling RoMON in the future, so that you could connect through one to all of the others without relying on IP at all.)
