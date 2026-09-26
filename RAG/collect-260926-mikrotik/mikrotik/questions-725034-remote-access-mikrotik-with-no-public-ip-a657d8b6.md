---
id: collect-260926-mikrotik/mikrotik/questions-725034-remote-access-mikrotik-with-no-public-ip-a657d8b6
title: "questions-725034-remote-access-mikrotik-with-no-public-ip-a657d8b6"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-725034-remote-access-mikrotik-with-no-public-ip-a657d8b6.md
source_anchor: ""
source_lines: [1, 7]
sha256: 16bac06f3f0c9ccea8dede3166d1d08b12b7a241cea801c6291843597cf8bf90
---

# questions-725034-remote-access-mikrotik-with-no-public-ip-a657d8b6

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
You can not, simple like that. It is likely that even IF you manage - which I am not sure you could - the result would be you getting fired because of you going around the business internet policy.
The problem is that you have no public IP at all and VPN is not allowed, and NAT is not - that rules out all possible ways to communicate via IP, which the internet needs.
This smells brutally like a business or internet provider telling you that internet is NOT for publishing anything - and bypassing this is not a technical question. Besides being off topic, you open yourself to criminal prosecution and immediate job loss. Do not do it.
You can create an outgoing SSTP tunnel which works using SSL/TLS over port 443 (so it shouldn't be blocked) to connect to another SSTP server outside the network.
Then you can access the router with the IP assigned on the SSTP interface.
Create NAT rules from the broadmabd router (if you have access) to forward incoming connections on port 8291 (for example) to your Mikrotik's LAN address.
