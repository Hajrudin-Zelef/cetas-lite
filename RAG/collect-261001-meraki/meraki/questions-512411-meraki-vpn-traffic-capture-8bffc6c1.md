---
id: collect-261001-meraki/meraki/questions-512411-meraki-vpn-traffic-capture-8bffc6c1
title: "questions-512411-meraki-vpn-traffic-capture-8bffc6c1"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-meraki/questions-512411-meraki-vpn-traffic-capture-8bffc6c1.md
source_anchor: ""
source_lines: [1, 9]
sha256: f887c635c21cd8777d63cc88ce056623bd0bd1a26076f7e5553aa1d3f3ac6ada
---

# questions-512411-meraki-vpn-traffic-capture-8bffc6c1

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We have multiple offices interconnected via VPN using Meraki switches. I've been looking around and can't seem to find a way, but was thinking it's possible. Is it possible for me to put my desktop on one of the remote vpn'd subnets so I can monitor the traffic? We have one of our offices experiencing high traffic volume and we're looking for a comparison without having to remote into a local computer.
If you only have Meraki switches you can enable port mirroring (Configure -> Switch Settings) on the ports sending the vpn traffic and view it through your desktop and use something like Wireshark.
If you have Meraki APs or Appliances the monitoring is easier as packet capture and traffic analysis is all built in.
