---
id: collect-260926-mikrotik/mikrotik/questions-1471612-mikrotik-hotspot-performance-degradation-so-slow-hotspot-unusa-b745e3d8
title: "Mikrotik Hotspot Performance Degradation: So Slow Hotspot Unusable"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-1471612-mikrotik-hotspot-performance-degradation-so-slow-hotspot-unusa-b745e3d8.md
source_anchor: ""
source_lines: [1, 23]
sha256: 32275fa7a205a86fde36978da786841f3e8479681ac219e84278318cc90d14da
---

# Mikrotik Hotspot Performance Degradation: So Slow Hotspot Unusable

*Source : https://superuser.com/questions/1471612/mikrotik-hotspot-performance-degradation-so-slow-hotspot-unusable | Site : superuser.com | Score : 1*

FAULT:

Hotspot created with Mikrotik's "ip > Hotspot" wizard is so painfully slow that it is unusable

Googling "mikrotik hotspot slow" returned a gazillion results, none of which resolved this fault.

I checked queues, DNS, proxy and firewall rules created by the Mikrotik "wizard"- which appeared to provide the correct connectivity.

---

## Reponse (ACCEPTEE) — score 2

Mikrotik's automated Hotspot "wizard" created the fault: it inserts its rules ABOVE the one allowing related,established connections in the FORWARD chain.

RESOLUTION:

Moving the related,established rule to the top of the FORWARD chain successfully resolves this fault, restoring normal web browsing performance.

Hopefully long-suffering Mikrotik users will find this answer and save themselves reading numerous posts which lead them in the wrong direction...
