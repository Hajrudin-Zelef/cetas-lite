---
id: collect-260926-mikrotik/mikrotik/questions-990453-route-internet-traffic-through-vpn-failed-mikrotik-tried-ikev2-c1b0b7f7
title: "questions-990453-route-internet-traffic-through-vpn-failed-mikrotik-tried-ikev2--c1b0b7f7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-990453-route-internet-traffic-through-vpn-failed-mikrotik-tried-ikev2--c1b0b7f7.md
source_anchor: ""
source_lines: [1, 8]
sha256: 498e5c460d809e7d09eb9d53681389872a23a96555bddc4870e938f7707278c6
---

# questions-990453-route-internet-traffic-through-vpn-failed-mikrotik-tried-ikev2--c1b0b7f7

want to route some internet traffic through vpn, from local site to remote site router(R2)
R1: LOCAL site (750Gr3): LAN / 172.16.1.1/24, gre over ipsec tunnel ip: 192.168.8.1/24
R2: Remote site (CHR, cloud vm):  WAN / 172.18.0.4/24,LAN / 172.18.1.4/24, gre over ipsec tunnel ip: 192.168.8.2/24
use 8.8.8.8 for example, have add route on R1: add distance=1 dst-address=8.8.8.8/32 gateway=192.168.89.2 (also tried use mangle to mark routing and use for routing)
(from R1 LAN side) mtr 8.8.8.8 can reach 192.168.8.2, but all ends there.
have tried nat traffic from 172.16.1.0/24 (and src-nat/src-nat), not work. (R2 can reach 172.16.1.x) logging will only show first packet, then no further packet is pass through the policy.
and have tried different vpn setup: a) ikev2 site to site vpn, with tunnel mode b) gre over previous ikev2 vpn. internet routing both not work (only LAN to LAN work).
Any ideas?
