---
id: collect-260926-mikrotik/mikrotik/questions-744752-route-all-ipv4-traffic-over-a-6in4-tunnel-using-routeros-6-x-f46dbc7d
title: "questions-744752-route-all-ipv4-traffic-over-a-6in4-tunnel-using-routeros-6-x-f46dbc7d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "research"]
source: docs/RAG/lot-mikrotik/RouterOS/questions-744752-route-all-ipv4-traffic-over-a-6in4-tunnel-using-routeros-6-x-f46dbc7d.md
source_anchor: ""
source_lines: [1, 11]
sha256: 339d0abffdf72b6fe7576c90fc96d8b503e57430405318ba27740259f19a8dd7
---

# questions-744752-route-all-ipv4-traffic-over-a-6in4-tunnel-using-routeros-6-x-f46dbc7d

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
Would it be possible to route all IPv4 traffic over an IPv6 tunnel?
Let me elaborate, let's say I establish a tunnelbroker (HE) 6in4 tunnel, using IPv4 (duh!), but then I want to relay all my IPv4 traffic over the IPv6 tunnel, so, in other words make this IPv6 link my default gateway for all internet traffic (IPv6 and IPv4).
A 6in4 tunnel transports IPv6 packets, so it is not possible to transport IPv4 packets in such a tunnel.
What you can do if you want all your traffic to go through the tunnel is to change your local network to IPv6-only. Usually this limits you to reaching IPv6-capable sites, but when you combine it with NAT64/DNS64 then you would be able to reach IPv4 sites through NAT64. There is a list of public NAT64 gateways here.
I wouldn't recommend such a setup though, you'll end up with horrible routing. This will not be good for your latency and bandwidth.
