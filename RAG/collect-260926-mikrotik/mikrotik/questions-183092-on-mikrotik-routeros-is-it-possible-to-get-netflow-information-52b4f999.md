---
id: collect-260926-mikrotik/mikrotik/questions-183092-on-mikrotik-routeros-is-it-possible-to-get-netflow-information-52b4f999
title: "questions-183092-on-mikrotik-routeros-is-it-possible-to-get-netflow-information--52b4f999"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "open source", "research"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-183092-on-mikrotik-routeros-is-it-possible-to-get-netflow-information--52b4f999.md
source_anchor: ""
source_lines: [1, 11]
sha256: b84c26833096dd2955e558f58ca335d77673d1261e9e6f011a90d64c597b3a66
---

# questions-183092-on-mikrotik-routeros-is-it-possible-to-get-netflow-information--52b4f999

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a RouterOS box set up to bridge two ethernet connections. I have use-ip-firewall=yes in the bridge configuration, so that the ports go through the firewall.
I've enabled netflow reporting via ip/traffic-flow, but the only packets I see reported are broadcast and multicast packets, not the packets that are flowing through the bridge. The documentation indicates that traffic flow logging happens after firewall processing and that it won't work with bridged connections by default, but I would have thought that use-ip-firewall=yes ought to address this.
I think the basis for their NetFlow support is centered on open source technologies. We tried to configure Mikrotik NetFlow with Scrutinizer NetFlow & sFlow Analyzer and it wasn't possible at that time, which was about a year ago. If things have changed with their NetFlow support since then, we would certainly be willing to run more tests. Please send us a packet capture, if we can be of any help.
IP firewalled bridges will only monitor flows for packets crossing your bridge, meaning that any connection between machines located in the same side of your bridge won't be shown in NetFlow as packets never reach your router.
If you ping a node located in the remote side of the bridge, you should be able to monitor in "/ip firewall connections" both broadcast ARP packets requesting the remote node IP, and then the established ICMP connections, this time crossing your router.
