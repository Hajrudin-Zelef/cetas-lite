---
id: collect-260926-mikrotik/mikrotik/questions-28494-nat-network-setup-with-dsl-modem-and-mikrotik-router-a5ba91a0
title: "questions-28494-nat-network-setup-with-dsl-modem-and-mikrotik-router-a5ba91a0"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-28494-nat-network-setup-with-dsl-modem-and-mikrotik-router-a5ba91a0.md
source_anchor: ""
source_lines: [1, 24]
sha256: c4d647347e2df3873b170f0a8d8a00279b212dad68caa7182b0e14d95701dce0
---

# questions-28494-nat-network-setup-with-dsl-modem-and-mikrotik-router-a5ba91a0

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
2
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have littel experience in this kind of setups, so please bear with me. I think the question is kind of basic.
I have to install a router in a LAN since I'll have to have it mandle a VPN tonnel later. In short I have to go from this:
< PUBLIC IP >[DSL Modem]<192.168.1.1> ------ <192.168.1.0>[A BUNCH OF SERVERS]
to this:
< PUBLIC IP >[DSL Modem]<192.168.0.1> --- <192.168.0.2>[Mikrotik Routerboard]<192.168.1.1> --- <192.168.1.0>[A BUNCH OF SERVERS]
I NATted the services I need to forward to the servers using the DSL modem interface (mail, rdp, ssh, etc). Now with the new setup I'm really confused on how should I procede.
Is it sensible to have 2 NAT "hops"? DSL modem forwarding all services requests from the WAN to 192.168.0.2 (again by using the NAT interface) and then forwarding them again to the server (from the router)? Does NAT work this way? Is there another smarter way to solve this?
Don't know what DSL modem you're using, but you're probably better to make that DSL modem work in bridge mode and have your MikroTik router working as a normal WAN router facing WAN and LAN sides. That way you'll have no issues with VPN, be able to forward traffic as you like and set up a good firewall. This is very typical and good setup. And this is how I'd do it.
If you haven't got enough control over your DSL modem, then sure, you can do chained NAT, forward everything to MT router or, as Benoit suggested, set up a DMZ feature. That way you'll still be able to bring the VPN tunnel up, but only as a client.
The scenario you describe works, you can "chain NATs" without problem, for example
[Public IP]:58000 ->NATed by DSL modem on 192.168.0.2:58000 ->NATed by mikrotik on 192.168.1.55:80
Or you can use the "DMZ" feature on your modem to redirect all incoming ports to 192.168.0.2 and do the NAT on Mikrotik only.
Or you can set up your DSL modem in bridge mode (if it can do that) so all incoming connections are directly sent (bridged) to Mikrotik router, which will do the NAT.
You can have double NAT if you want but you will find it easier to purchase a router that has a DSL modem built in.
You could also have your current modem act only as a modem and bridge the connection to the new router using something like PPPoE.
If you absolutely cannot change your set up then it will work but it will be a hassle if you have to port forward anything . If you port forward port 25 (SMTP) on the outside interface of your modem to your mikrotik routers outside interface (192.168.0.2), you will also need to port forward port 25 on your mikrotik routers outside interface to the internal resource on the LAN. You'll be doing everything twice basically.
Assuming your public ip is 1.1.1.1 and you want to forward SMTP traffic from the DSL modem to an email server on your LAN at 192.168.1.5 it would look like the following:
A VPN tunnel should be fine as long as you have NAT-T on both ends of the tunnel and the other side has a static IP without NAT. You can set your side to be dynamic and only have it initiate the connection
