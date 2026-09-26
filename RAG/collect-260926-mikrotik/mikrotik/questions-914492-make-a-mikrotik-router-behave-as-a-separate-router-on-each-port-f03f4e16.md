---
id: collect-260926-mikrotik/mikrotik/questions-914492-make-a-mikrotik-router-behave-as-a-separate-router-on-each-port-f03f4e16
title: "questions-914492-make-a-mikrotik-router-behave-as-a-separate-router-on-each-port-f03f4e16"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-914492-make-a-mikrotik-router-behave-as-a-separate-router-on-each-port-f03f4e16.md
source_anchor: ""
source_lines: [1, 14]
sha256: 84f50103e6973f866a796d22e2e8baef8a52d68a9a164099e5c9245fcb8c3c54
---

# questions-914492-make-a-mikrotik-router-behave-as-a-separate-router-on-each-port-f03f4e16

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
We use Mikrotik routers for many points in our system, but we have an odd need for one location.
We would like to use one router (example: RB960PGS-PB), where the main Internet connection comes in on port 1, and have ports 2-5 each act as a separate public-facing router.
IE, each port will be assigned a custom MAC, will have a public IP bound (we have a large pool of public IPs), and on the "inside" it will act like a normal router (firewall, dhcp for 192.168.x.x or 10.x.x.x, etc..).
Is this even possible? We are using Ubiquiti Lite APs on each of the ports (2-5) but we would like each AP to act like a separate router/network with its own public IP.
There are site restrictions that prevent us from just putting an AirCube, Linksys, or other cheap router between the Mikrotik site router and the AP.
Yes. That's literally what a router is. The term "router" has become diluted over the past 15 years to mean almost anything. It often means "Router/Switch/WiFi Access Point/DHCP Server/Firewall". But the "Routing" part of it is the part you need.
In your MikroTik, remove the ports from "bridge" interface. Bam. You no longer have a switch, you have all the individual ports acting independently.
Note that, being a router, you will still be able to access the other subnets as the router will say "Oh I know how to get to that network" and will just forward the packets. If you don't want this, if you want each port to be isolated, then you'll need to configure the firewall on your Mikrotik to deny the cross-interface traffic.
