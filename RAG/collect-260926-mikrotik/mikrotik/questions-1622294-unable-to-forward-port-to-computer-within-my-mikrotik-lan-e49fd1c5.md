---
id: collect-260926-mikrotik/mikrotik/questions-1622294-unable-to-forward-port-to-computer-within-my-mikrotik-lan-e49fd1c5
title: "questions-1622294-unable-to-forward-port-to-computer-within-my-mikrotik-lan-e49fd1c5"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1622294-unable-to-forward-port-to-computer-within-my-mikrotik-lan-e49fd1c5.md
source_anchor: ""
source_lines: [1, 3]
sha256: 11497e17d8036385e192c565154b69075af79d3c0bb8121bd3612a697e418ec9
---

# questions-1622294-unable-to-forward-port-to-computer-within-my-mikrotik-lan-e49fd1c5

I do not have extensive background in networking and I want to achieve something simple within my LAN. I have a server running on port 8080 on my computer and I want to make it accessible through the router's public IP. I have done that with regular residential gateways many times, but recently I switched to Mikrotik hAP ac² and cannot seem to succeed. I tried adding the following rule:
/ip firewall nat add chain=dstnat dst-port=8080 action=dst-nat protocol=tcp to-address=192.168.*.* to-port=8080
but when I type <public-ip>:8080 in the browser it gets stuck for some time and then returns This site can't be reached. I tried to configure it through WebFig, but without success. The only other rule I have is a pre-existing masquerade. What am I doing wrong?
