---
id: collect-260926-mikrotik/mikrotik/questions-1029386-mikrotik-join-2-different-ip-networks-085451b4
title: "MIKROTIK - Join 2 different IP networks"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-1029386-mikrotik-join-2-different-ip-networks-085451b4.md
source_anchor: ""
source_lines: [1, 21]
sha256: 97381babc33a0ce8bbafed9ce54b385301ae46bb44ecc033ad2da543cc248259
---

# MIKROTIK - Join 2 different IP networks

*Source : https://serverfault.com/questions/1029386/mikrotik-join-2-different-ip-networks | Site : serverfault.com | Score : 0*

Scenario:

I have 1 ISP.
The Mikrotik router is connected to that ISP in ether1 (192.168.0.3).
Mikrotik gives out DHCP addresses 10.0.0.0/8. My PC is connected to it (10.0.0.6).
The other network has a router(ddwrt) that connects to the mikrotik router and gives out 10.1.1.1/8 addresses

I want network 10.0.0.x to be able to connect with network 10.1.1.x, specifically my PC to monitor other computers on the other network.
How can I proceed of this? Thank you so much!

---

## Reponse — score 1

Adding a static route on the mikrotik router should be enough (https://wiki.mikrotik.com/wiki/Simple_Static_Routes_Example)

/ip route add dst-address=10.1.1.0/24 gateway=10.0.0.7
