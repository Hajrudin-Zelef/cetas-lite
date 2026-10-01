---
id: collect-261001-general-networking/general-networking/questions-1164123-routing-traffic-between-2-routers-e6ff57d7
title: "Routing traffic between 2 routers"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/questions-1164123-routing-traffic-between-2-routers-e6ff57d7.md
source_anchor: ""
source_lines: [1, 41]
sha256: 152e13629acd33f2f24961d12614f5f1673474863de46063cc5dd240723d6950
---

# Routing traffic between 2 routers

*Score : -1 | Source : https://serverfault.com/questions/1164123/routing-traffic-between-2-routers*

This is roughly the network infrastructure as of now
As you can see I have two routers/firewalls in Nethserver and OpnSense (This is newly installed with the new ISP).
Nethserver LAN's subnet is 192.168.0.0/24 (called 0 subnet) OpnSense LAN's subnet is 192.168.2.0/24 (called 2 subnet)
Most of the devices connected to the network (access point or via the switch) are of the subnet 192.168.0.0/24 with the DG of 192.168.0.1
I understand that devices on the 0 subnet aren't able to directly connect/communicate to devices of 2 subnet since they are of different subnets (ARP won't work) and any request is basically sent eventually to the DG 192.168.0.1
Once here, the Nethserver router is sending it out of the network to find the device(s)
I've tried adding a static route on Nethserver using the command below
ip route add 192.168.2.0/24 via 192.168.2.1 dev eth0
But this just gives the output
RTNETLINK answers: Network is unreachable
I'm missing something here. I know it. But can't point a finger on it.

---

### Reponse (acceptee) — score 0

I understand that devices on the 0 subnet aren't able to directly connect/communicate to devices of 2 subnet since they are of different subnets (ARP won't work)
If I understood your diagram correctly, both IP subnets are on the same Ethernet (no VLANs or any other L2 isolation). In this setup, ARP can work – the hosts just don't know about the possibility. This is directly related to your error message.
answers: Network is unreachable
The "next hop" must be an address that your system already has a direct route for. You can't route 192.168.2.0/24 via 192.168.2.1 because the routing table doesn't know how to reach 192.168.2.1 yet.
If you're sure that 192.168.2.1 is on the same Ethernet network (i.e. it could be resolved via ARP), there are two ways to tell Linux about this:
Use the onlink flag to override the check.
ip route add 192.168.2.0/24 via 192.168.2.1 dev eth0 onlink
Or, manually create a local (non-gateway) route for the nexthop.
ip route add 192.168.2.1/32 dev eth0
ip route add 192.168.2.0/24 via 192.168.2.1
But... because the entire subnet is equally on the same Ethernet, you can just declare the entire subnet as on-link instead:
ip route add 192.168.2.0/24 dev eth0
...and then the system will directly use ARP for all hosts. (You can also deploy this route to client devices via DHCP for better performance.)
Hopefully, all of this is temporary – although having two subnets and two routers overlapping like this isn't the worst, it certainly isn't good practice either. (For example, you will get into situations where traffic goes through a router in one direction but directly in the other, and that confuses its firewall.) Ideally the two subnets would be different VLANs, and the routers would have a dedicated /30 connection between them (or of course one router to handle both subnets).
Alternatively you could have one IP subnet with two gateways, DHCP providing one gateway address and certain servers being manually configured to use the other. Still not great, but seems a little better than the current setup.

---

### Reponse — score 0

It's also worth to note, that for the full connectivity the route "FROM" should also be present - the "other end" of communication needs to know where to send the reply. Meaning: you should also create a similar route on the OpnSense.
