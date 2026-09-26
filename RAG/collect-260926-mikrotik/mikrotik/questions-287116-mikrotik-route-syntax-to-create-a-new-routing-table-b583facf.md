---
id: collect-260926-mikrotik/mikrotik/questions-287116-mikrotik-route-syntax-to-create-a-new-routing-table-b583facf
title: "Mikrotik route syntax to create a new routing table"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-287116-mikrotik-route-syntax-to-create-a-new-routing-table-b583facf.md
source_anchor: ""
source_lines: [1, 26]
sha256: 3318f5d36942539ca22f3c251b6b43a39800abdee4d62d03adf1499f7474021b
---

# Mikrotik route syntax to create a new routing table

*Source : https://serverfault.com/questions/287116/mikrotik-route-syntax-to-create-a-new-routing-table | Site : serverfault.com | Score : 2*

I want to replace my linux gateway box to Mikrotik. But I never found how to write this linux's iproute2 command in mikrotik's way:

ip route add default via 10.1.1.1 table browser
ip route add 10.1.2.0/24 via 10.1.1.2 table browser
ip rule add prio 100 from 10.1.6.2 lookup browser

anyone could lead me how to create a new routing table instead using the main's route?

---

## Reponse (ACCEPTEE) — score 4

The following should do the trick:

/ip route
add gateway=10.1.1.1 routing-mark=browser
/ip route
add gateway=10.1.1.2 dst-address=10.1.2.0/24 routing-mark=browser
/ip route rule
add src-address=10.1.6.2 action=lookup table=browser

Note that in RouterOS, the priority is set by the order of the routing rules.
