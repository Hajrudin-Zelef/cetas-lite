---
id: collect-260926-mikrotik/mikrotik/questions-1312256-site-to-site-vpn-between-mikrotik-and-wathguard-7dbc0052
title: "questions-1312256-site-to-site-vpn-between-mikrotik-and-wathguard-7dbc0052"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/vpn/questions-1312256-site-to-site-vpn-between-mikrotik-and-wathguard-7dbc0052.md
source_anchor: ""
source_lines: [1, 11]
sha256: aff0907497b505f2f5a948ec38b88ccf6f6cf9a3ba8a98909540d045ff1d878e
---

# questions-1312256-site-to-site-vpn-between-mikrotik-and-wathguard-7dbc0052

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I've built an IPSEC site-to-site vpn between a Mikrotik router and a Watchguard firewall.
Everything works fine, only one thing left:
I'd like to route all the traffic from Mikrotik over the Watchguard (because on Whatchguard there's some website filtering rule and I want to accept this user restrictions on the Mikrotik site as well).
I think that is impossible using only IPSec site-to-site, but it's very easy to setup a simple GRE tunnel (no encryption) that flows over the IPSec; then you can redirect all the traffic thru the IP address which Watchguard have on the GRE interface.
172.16.22.0/24subnet? I've never done this personally like that but I would think it's possible by simply changing some routes, the default gateway routes use, or something along those lines. Since you setup the site VPN and configure the router, I thought maybe the documentation would give you some more specific pointers about this functionality.
