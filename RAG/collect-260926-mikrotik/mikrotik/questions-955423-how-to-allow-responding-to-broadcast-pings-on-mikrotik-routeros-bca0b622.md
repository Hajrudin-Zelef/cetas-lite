---
id: collect-260926-mikrotik/mikrotik/questions-955423-how-to-allow-responding-to-broadcast-pings-on-mikrotik-routeros-bca0b622
title: "How to allow responding to broadcast pings on Mikrotik RouterOS?"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-955423-how-to-allow-responding-to-broadcast-pings-on-mikrotik-routeros-bca0b622.md
source_anchor: ""
source_lines: [1, 9]
sha256: 7e3e269a17d9584900113090729369eaab714ff449272cbe7807c74e7d112333
---

# How to allow responding to broadcast pings on Mikrotik RouterOS?

*Source : https://serverfault.com/questions/955423/how-to-allow-responding-to-broadcast-pings-on-mikrotik-routeros | Site : serverfault.com | Score : 1*

By default, most systems don't respond to broadcast ping requests. However, there exist ways to explicitly allow responding to broadcast ICMP requests on multiple systems, such as Linux.

How can I enable responding to broadcast ping requests on Mikrotik?

Using /tool sniffer start and /tool sniffer packet print I was able to verify that the broadcast packets do indeed reach the Mikrotik routers, but a response is never sent.
