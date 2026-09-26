---
id: collect-260926-mikrotik/mikrotik/questions-832887-mikrotik-routing-and-vlans-e085d702
title: "questions-832887-mikrotik-routing-and-vlans-e085d702"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/questions-832887-mikrotik-routing-and-vlans-e085d702.md
source_anchor: ""
source_lines: [1, 21]
sha256: 1c1f508e68ea7b7c4e6f11102a87eb124766f7367359009ec9ec6ca65bacf5ed
---

# questions-832887-mikrotik-routing-and-vlans-e085d702

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have setup two VLANs on one Mirkotik router. Both VLAN0 and VLAN1 have their WAN ports. WAN1 (VLAN0) is connected to network 1 and WAN2 (VLAN1) is connected to network 2.
I want to:
route the traffic from VLAN0 via WAN1 to net1
and from VLAN1 via WAN2 to net2
...but there is a routing problem. Only one routing table exist, and only one default route for both VLAN0 and 1 can be set up.
Can I somehow create different routes for both VLANs?
You can achieve this by using routing marks in the mikrotik.
In the firewall you should assign a rule for each VLAN on the prerouting chain setting the action to the mark routing, but prior to this action, you should mark packets which are coming through the vlan interfaces.
so First you mark the packets like this:
Now you can do routing marks:
At the moment I have created a routing mark ('lookup table') named "vlan1-routing".
So I can set a rule in route with the configured marks, like this:
Any router isn't needed to assign static routes for directly connected networks. You have two ways.
You assign WAN1 ip addres as default gateway for all hosts from net1. Then you assign WAN2 ip addres as default gateway for all hosts from net2.
Add static route option to DHCP servers for both networks net1 and net2. But this method is difficult and not all operational systems can understand this dhcp options.
