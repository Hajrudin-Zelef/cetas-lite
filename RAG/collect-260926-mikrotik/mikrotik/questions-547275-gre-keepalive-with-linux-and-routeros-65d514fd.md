---
id: collect-260926-mikrotik/mikrotik/questions-547275-gre-keepalive-with-linux-and-routeros-65d514fd
title: "questions-547275-gre-keepalive-with-linux-and-routeros-65d514fd"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-547275-gre-keepalive-with-linux-and-routeros-65d514fd.md
source_anchor: ""
source_lines: [1, 21]
sha256: 122a11981eea4ae43c4f0f6556ae212f176e340264a918c7b0138ef6f5163bb4
---

# questions-547275-gre-keepalive-with-linux-and-routeros-65d514fd

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
0
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a Linux host and couple of routerboadrs. I created a GRE tunnel, but Linux does not answer keepalive packages. Then router mark gre connection as unreachable, so I cant send to Linux host from router subnet. If linux sends something into tunnel (ping, etc.) - RouterOS mark connection as reacheble. Second and next packages routed nicely until one minute idle (no traffic).
Tunnel in linux a make in this way:
remote=x.x.x.x
dev=gre21
network=10.21.0.0/16
ip tunnel add ${dev} mode gre remote ${remote} ttl 255
ip addr add 172.16.1.1/24 peer 172.16.1.21 dev ${dev}
ip linkset${dev} up
ip route add ${network} dev ${dev}
And ip l:
14: gre21: <POINTOPOINT,NOARP,UP,LOWER_UP> mtu 1476 qdisc noqueue state UNKNOWN
link/gre 0.0.0.0 peer 109.60.170.15
A GRE Keepalive is a "host to router" GRE packet encapsulated inside a "router to host" GRE packet. The idea being the host (in this case Linux) receives the packet, sees the packet is actually a GRE packet for the router, and sends it back out. The router receives this packet and knows the remote end is still responding.
The Linux FIB code is such that if it receives traffic where the source is a local unicast address, the traffic is considered invalid.
This is not exactly a direct answer to the GRE tunnel keepalive, however you may find it easier to use the MikroTik EoIP option to a linux server as this has been ported and allows tunnel ID and keepalive in a similar fashion (also handles fragmentation for any size link).
