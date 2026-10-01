---
id: collect-261001-huawei/huawei/questions-43676-2e60b88f
title: "questions-43676-2e60b88f"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-huawei/questions-43676-2e60b88f.md
source_anchor: ""
source_lines: [1, 15]
sha256: db346b2cc925a21f85b192e3942fd3189de6c7ab7c84c24f367c9aabd93cf975
---

# questions-43676-2e60b88f

fortigate - how to discover devices in a large heterogenous network with huawei, cisco, alcatel and more different devices without using snmp? - Network Engineering Stack Exchange
Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
i have a huge task to do and i need your help.I am new to the whole networking world and I apologize if this is a re-post of a question.
so what i wanted is to discover and mapping all my enterprise network on a web app using network standard commands and protocol but not using snmp. I have a huge network with devices from many vendor such as huawei NE40, Cicco, Juniper, fortinet ...
I directed this question to you experts cause i do not have the knowledge or know the keywords to search online on this subject matter.
It depends on what layer you want to/can discover devices. At the link layer, LLDP comes to mind but it requires you to directly interface with the device and only gives just a few items.
In a layer 2 network, you can scan for MAC/IP addresses by ARP. The vendor specific part (OUI) of the MAC tells you the vendor. Reverse DNS possibly tells you the device name.
When it comes to applications, it gets hard. You'd need to do a port scan which takes quite a while and only tells you on which ports a device accepts connections. Port probing might reveal what is behind a port but requires to know what you're looking for.
SNMP is a standardized protocol for exactly this purpose and it allows you to efficiently collect a wide range of information. If you want to collect detailed info without SNMP you'd probably need to do it manually (ie. look at each config).
?). If you are looking for a product or resource recommendation, those are explicitly off-topic here, as they are on most SE sites, except Software Recommendations and Hardware Recommendations.
