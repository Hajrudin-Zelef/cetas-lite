---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-68653-vpn-traffic-routing-on-ubiquiti-vyatta-e78a9067
title: "questions-68653-vpn-traffic-routing-on-ubiquiti-vyatta-e78a9067"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-68653-vpn-traffic-routing-on-ubiquiti-vyatta-e78a9067.md
source_anchor: ""
source_lines: [1, 11]
sha256: fa812c705c67d7f761353e91baa1bbd5e656a0b37ace6ce6027021a22ec29ec5
---

# questions-68653-vpn-traffic-routing-on-ubiquiti-vyatta-e78a9067

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I'm using Ubiquiti to route certain outgoing ip-ranges through a VPN. I would however like to route all HTTP/HTTPS traffic via the VPN except large downloads.
I realize that this is a tall order, but is it possible? Maybe after detecting a certain amount of bytes it could just reroute the connection and resume it without going over the VPN?
Normally, routing is based on a packet's destination address. Some routers feature policy-based routing which you could use to route by application protocol, e.g. HTTP/S. However, you cannot route based on a (yet unknown) size of an application-layer reply.
Even if the router could detect the reply size, you cannot reroute a TCP connection once it has been established. Doing so would change the client address visible to the public server and break the TCP socket.
What you ask could be possible using an application-level proxy, but those are off-topic here for operating above the transport layer.
