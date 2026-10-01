---
id: collect-261001-general-networking/general-networking/t5-support-forum-advpn-split-tunnel-m-p-199522-5fb241bd
title: "t5-support-forum-advpn-split-tunnel-m-p-199522-5fb241bd"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-general-networking/t5-support-forum-advpn-split-tunnel-m-p-199522-5fb241bd.md
source_anchor: ""
source_lines: [1, 8]
sha256: a48db72a0d0d12959db2a02f7ba7b9ac01231bf24cc48465d2e6e79aadd3fe25
---

# t5-support-forum-advpn-split-tunnel-m-p-199522-5fb241bd

Obviously, I'm not so good with ADVPN and I also just started learning FortiGate as I got my first FortiGate and trying to build some network. I'm not sure if it's possible and how to make split tunneling on my FortiGate. Our HQ doesn't have that good internet connection and I wouldn't like it that we have more troubles when we add all planned BO. For easier configuration, we chose to use ADVPN with BGP. Btw. We have two BO with a second WAN that uses LTE (no public IP) and at the moment that BO has configuration on some Cisco routers to combine two WANs for better connectivity.
I want to let you know that from my point of view there are no dummy questions, not even dummy answers . I will try my best to not give you a dummy answer :).
ADVPN would be defined by 2 main purposes:
1. achieving full mesh between the spokes
2.taking the throughput load off the hub by establishing spoke to spoke shortcut tunnels.
Lets talk about Split tunneling
Split tunneling by its purpose, beside the P2 selectors, would also add specific routes in the routing table (to steer specific traffic)- this is rendered useful in combination with ADVPN as routes need to be dynamically changed when a shortcut tunnel goes up.
As you already chosen the best path - BGP, BGP will be your Angel in the routing decision - will install routes in the routing table and dynamically change the outgoing interface in the routing table based on the next-hop reachability (either through the hub if no shortcut is created, either through the shortcut tunnel - this is a direct tunnel to the remote side).
