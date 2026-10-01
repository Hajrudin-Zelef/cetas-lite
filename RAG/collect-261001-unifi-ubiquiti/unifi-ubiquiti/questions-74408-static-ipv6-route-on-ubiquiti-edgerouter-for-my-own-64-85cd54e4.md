---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-74408-static-ipv6-route-on-ubiquiti-edgerouter-for-my-own-64-85cd54e4
title: "questions-74408-static-ipv6-route-on-ubiquiti-edgerouter-for-my-own-64-85cd54e4"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-74408-static-ipv6-route-on-ubiquiti-edgerouter-for-my-own-64-85cd54e4.md
source_anchor: ""
source_lines: [1, 22]
sha256: a66d68e7b58f4f50ba834ceb5a504c0af6e7812e149dff168a03704228f6d86b
---

# questions-74408-static-ipv6-route-on-ubiquiti-edgerouter-for-my-own-64-85cd54e4

I am colocating some hardware in a local datacenter, and I'm trying to get IPv6 going. My datacenter has assigned me a /64, but they have given me a gateway address within that /64.
The gateway (that I don't control) is 2606:1000:1000:1::1
I have two uplinks into the datacenter, which go to (my) two Ubiquiti Edgerouters that I control. Eventually, I'd like to get VRRP going, but for now, I just have some of my IP space on 1 router, and some on the other.
So on one of the edgerouters, I've configured the following:
- eth0 = 2606:1000:1000:1::11/64
- eth2 = 2606:1000:1000:1::111/64
- static route is set like so: set protocols static route6 '::/0' next-hop '2606:1000:1000:1::1' interface eth0
- I've assigned 1 of my hosts the following IP: 2606:1000:1000:1::210/64
My symptoms are as follows:
- The router can ping6 google.com and get responses back.
- The host can ping6 2606:1000:1000:1::111 and get responses back.
- The host cannot ping6 google.com.
- The router can NOT ping6 the host at 2606:1000:1000:1::210
I feel fairly positive that the router's set protocols static route6 is causing the issue here, but I'm having trouble figure out how to get a static route setup for the /64 to hairpin it back into eth2.
The following attempt doesn't work, when I try to configure a static route for my own /64:
cha-rtr1# set protocols static route6 '2606:1000:1000:1::0/64' next-hop '2606:1000:1000:1::111' interface eth2
[edit]
cha-rtr1# commit
[ protocols static route6 2606:1000:1000:1::0/64 next-hop 2606:1000:1000:1::111 ]
Error: Nexthop address cannot be same as own interface address
What am I missing here? Can someone help explain to me how I can get traffic flowing to my host? Or do I need to contact my datacenter, and ask them to stop being cheap, and allocate a /48 to me, or something?
/64networks. You could get a/60prefix that would give you a/64transit network and 15/64access networks. I'm guessing the data center guys really do not understand IPv6 or the fact it was designed to waste addressing, and they are trying to apply IPv4 address conservation that is detrimental to IPv6.
