---
id: collect-260926-mikrotik/mikrotik/questions-846922-forward-between-interfaces-on-mikrotik-1e8bbc53
title: "questions-846922-forward-between-interfaces-on-mikrotik-1e8bbc53"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-846922-forward-between-interfaces-on-mikrotik-1e8bbc53.md
source_anchor: ""
source_lines: [1, 28]
sha256: 58a287e21053430a3c5f8bcb67d079739eb0d49b0cc18c721afb83c078ab1130
---

# questions-846922-forward-between-interfaces-on-mikrotik-1e8bbc53

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
im having a trouble
ive a mikrotik router with 2 interfaces up ( let's call lan1 and lan2 )
Lan 1 has the IP 192.168.100.1
lan 2 192.168.0.32
Lan1 the other side of the wire goes to a cisco wich IP is 192.168.100.20 , and beside that cisco its another network with IP 10.94/16
if i test over the mikrotik with winbox y can reach the cisco AND the other network itself,
now in my network we have the range 192.168.0.0/16
i can ping the lan2 of the mikrotik, but cant reach lan1 or cisco or 10.94 network,
could anyone help me wich filter rules and nat rules should i create to forward the requestest from 192.168.0.0/23 and reach 10.94.0.0/16 ? or the whole traffic coming for LAN2 forward to LAN1?
The problem is that Lan1 or Cisco or 10.94.x.x Don't have a route to the Lan2 network.
So you either have to write a route in cisco/10.94.x.x for the lan1 (192.168.100.0/24) or write a NAT rule in mikrotik to obtain access.
This solution only works if mikrotik and cisco are gateways for the networks
For instance in Cisco router(config mode) you can do:
ip route 192.168.0.0 255.255.252.0 192.168.100.1
and a route in the mikrotik for 10.94 network:
ip route add dst-address=10.94.0.0/16 gateway=192.168.0.1
(I assume mikrotik IP is 192.168.0.1)
or Do a NAT in mikrotik which seems a better way like this:
ip firewall nat add action=masquerade chain=srcnat src-address=192.168.0.0/23
ip firewall nat add action=masquerade chain=srcnat dst-address=192.168.0.0/23
or even a masquerade for all to make every packet passing Mikrotik route-able :
ip firewall nat add action=masquerade chain=srcnat
