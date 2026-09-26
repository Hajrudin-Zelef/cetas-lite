---
id: collect-260926-mikrotik/mikrotik/questions-456814-bonding-adsl-lines-on-mikrotik-d38cb61a
title: "questions-456814-bonding-adsl-lines-on-mikrotik-d38cb61a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-456814-bonding-adsl-lines-on-mikrotik-d38cb61a.md
source_anchor: ""
source_lines: [1, 10]
sha256: d45cde8d0e342d895bd33f841a1067909105895139574bc3196f97e76b644baf
---

# questions-456814-bonding-adsl-lines-on-mikrotik-d38cb61a

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I Currently have a Cisco router which I rent from Mweb. On this router I have two incoming 4MB ADSL Lines which mweb bonds for me to acquire 8Mbps. I Don't have any access to the router they do the whole configuration for me.
I Want to know if it is possible to do this myself "bonding" on any of the Mikrotik Routerboards? I Have heard from a friend that he did this by using PPPOE and Tunnels etc. And want to know before trying if it is possible?
Any indication that it would be possible would be appreciated
You can't make a 8M internet from two 4M. This is only possible if they are from one ISP and that ISP supports mlppp, or equivalent LACP protocols (which is rare). but you can loadbalance, and manage bandwidth between devices, and destinations, and protocols, and Many things else.
