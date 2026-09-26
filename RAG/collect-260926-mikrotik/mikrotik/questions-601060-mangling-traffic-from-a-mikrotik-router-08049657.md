---
id: collect-260926-mikrotik/mikrotik/questions-601060-mangling-traffic-from-a-mikrotik-router-08049657
title: "questions-601060-mangling-traffic-from-a-mikrotik-router-08049657"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-601060-mangling-traffic-from-a-mikrotik-router-08049657.md
source_anchor: ""
source_lines: [1, 11]
sha256: 81578163e683070f71770098d2f37c13404987bd289fdd53ba0521306dfd1b39
---

# questions-601060-mangling-traffic-from-a-mikrotik-router-08049657

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a MikroTik powered Router in the house with a couple of internet connections (2 200/10Mb Cable modems and a 100/20Mb VDSL Line). I am using Mangle rules to set routing marks and NAT rules to do some load balancing, and everything seems to be going grand... But it only works for traffic from outside the router... Let me explain:
I have 4 GigE ports on the machine, WAN1,2 and 3, and a LAN port named LAN1. All traffic from LAN1 is getting mangled (as it should be) but traffic from the load router itself (proxy traffic, IPv6 tunnels, VPN connections) are not being mangled. They get the first route to 0.0.0.0/0, which in my case is WAN2, and stick with it.
So, how do I get traffic from the local router to be mangled? Originally it was proxy traffic that caused the problem, but now with IPv6 and VPN, they are more important to be mangled... last time i enabled IPv6 traffic, all traffic only went though WAN2, and the rest where unused... Any ideas?
You will need to mangle your connections at the output chain level so your router can take advantage of load balancing for his own outgoing connections, but... You must be extremely careful when doing so; remember that local addresses by default belong to the "main" routing table, which I assume it is the one you are using for management from your LAN workstations, and once you route-tag a connection, it may be impossible for the router to return packets to your admin workstation.
Before you apply routing marks at the output chain, verify that all of your tagged route tables allow traffic from your admin interfaces or you may get yourself isolated from your router admin.
