---
id: collect-260926-mikrotik/mikrotik/questions-530901-routing-between-multiple-mikrotik-firewalls-0d073b49
title: "questions-530901-routing-between-multiple-mikrotik-firewalls-0d073b49"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-530901-routing-between-multiple-mikrotik-firewalls-0d073b49.md
source_anchor: ""
source_lines: [1, 26]
sha256: af090669b25f46432bc2f74843ed5128995c2cbc1442f1710ea612db000d546a
---

# questions-530901-routing-between-multiple-mikrotik-firewalls-0d073b49

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have 2 Miktorik RouterBoards, an RB1100 and a RB951G. The 951G is acting as my Wireless box, and has Guest, Internal and Internet Only Wireless networks. The RB1100 has 3 WAN connections (2x150/10Mb Cable modems and a 70/20Mb VDSL modem) and does Load Balancing, firewalling, etc, for the whole network.
The RB1100 is on network 192.168.0.0/24 and the 951 has 3 address ranges:
Guest -> 192.168.87.0/24
Internal -> 192.168.88.0/24
Internet Only -> 192.168.89.0/24
The idea is that guest is firewalled big time (limited bandwidth, limited sites, etc) which i have working with the help of the hotspot.
Internet Only should only be routed to the internet, possibly limiting some ports, and should not see anything on the 192.168.0.0/24 network.
Internal should have access to both the internet and also the 192.168.0.0/24 network, and anything on the 192.168.0.0/24 network should be able to see the 192.168.88.0/24 network also...
I Had the Internet Only part working to an extent, but accidentally cleared my router config (doh) but i never managed to setup the Internal network correctly...
Currently i have NAT enabled and that allows me to see all machines on the 192.168.0.0/24 network from the 88.0/24 network, but 0.0/24 cannot see 88.0/24 network...
I know i need to do something with routes, but even when i had that, something was not allowing me to see machines (laptop on wifi could not see desktop on wired).
So, Where am i going wrong?
Again, sorry i cant post the exact config... lost it in a firewall rule screw up...
Ok, thanks to DKNUKLES for his answer, but the problem was more a configuration problem on my end... I ended up adding 2 routing items to the routers, one on the 1100 and one on the 951. the 951 said to route all traffic (0.0.0.0/0) to the RB1100. the RB1100 had a route to point all 192.168.88.0/24 traffic to the IP of the 951. but still no joy...
the problem was my pre-routing. since i have multiple WAN connections and since my pre-routing did round robin, all traffic on my network was being routed to one of the WAN ports, if it was not for the 192.168.0.0/24 network. So, at the end of my pre-routing block, i added a rule, any traffic going to 192.168.88.0/24 with a connection-mark was cleared of its connection mark. this by-passed the loadbalancing rules and allows the traffic to go in the correct area... I realized it could be something like that when running traceroute. from one of the non WiFi connected machines to a WiFi connected machine, i was being routed out to one of my modems... Anyway, all fixed... I can now get to machines from the 192.168.88.0 network and they can see machines in the 192.168.0.0 network. happy days!
For your internet only group you would create a firewall rule that states that all traffic from that "Guest" network that is destined to your internal networks should be dropped. Make sure to create a rule to drop all input traffic from that subnet as well, as you don't want people fudging with your router (or at least attempting to...)
As for your wireless network there are two ways you can do this.
You can set up a route that connects two networks (wireless and trusted wired clients) - similar to a site-to-site VPN route. The disadvantage to this is that you'll lose your source address and all traffic will appear coming from the "route" network.
You can set up your wireless router to receive DHCP addresses from your RB1100 as a DHCP relay. This will allow you to preserve your source addresses and give access to your wireless network. You can then manipulate traffic as you see fit.
If I understand your problem correctly, you should not need NAT to accomplish what you're hoping to do as MikroTik will take care of the routes for you.
