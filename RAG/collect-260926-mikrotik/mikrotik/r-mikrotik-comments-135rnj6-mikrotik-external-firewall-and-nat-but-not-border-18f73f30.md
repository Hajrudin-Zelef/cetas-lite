---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-135rnj6-mikrotik-external-firewall-and-nat-but-not-border-18f73f30
title: "r-mikrotik-comments-135rnj6-mikrotik-external-firewall-and-nat-but-not-border-18f73f30"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/firewall-nat/r-mikrotik-comments-135rnj6-mikrotik-external-firewall-and-nat-but-not-border-18f73f30.md
source_anchor: ""
source_lines: [1, 24]
sha256: 2d7ec97e91faff07d60d551151fdfc42b407e31481037fae2aee6992bd122416
---

# r-mikrotik-comments-135rnj6-mikrotik-external-firewall-and-nat-but-not-border-18f73f30

Mikrotik External Firewall and Nat but not Border Router 
        
    Hello Please help me Understand how to accomplish this. Basically, I need LAN devices or all devices after the border/core router to be behind NAT/Firewall. then the core will handle the data again to load balance using ECMP before going out to the internet
Im assuming firewall devices use ipv4 ipv6 to communicate with other devices
My question are:
- 
      how will CCR2216(core) communicate to CCR2116(Firewall) using ipv4 address?
- 
      How will CCR2116(Firewall) communicate to internet if its behind CCR2216(core)? The defaultt route of CCR2116(firewall) will be CCR2216(core) ?
- 
      What device will the LAN default route point to? To CCR2116(Firewall) or CCR2116(firewall)
?
4. if the default route is CCR2216(core). How will CCR2216(core) redirect the traffic going to the internet to CR2116(Firewall)? Do I need to use policy base routing to redirect the traffic?
5. From the point of view of CR2116(Firewall) when it receives and Applies NAT to the traffic from LAN going to the internet. what IP address will it use to represent the LAN IP addresses to the public internet?
6. What will be the default route of CR2116(Firewall) ?
Section des commentaires
Core would have to be configured with the IP addresses that are to be routed to Firewall.
Yes. If there are no other routes on the routing table of Firewall, it will send traffic to the default route, which you would point at Core.
This doesn't make sense.
Why would Core redirect Internet traffic to Firewall, which has no direct Internet connection? This question again makes no sense.
It will use whatever public IP you configured it to use. This is a problem because you have have fixed IP addresses and no BGP on two of your three connections, which means in order to use ISP1 your address must be 120.1.1.2, which you can't use because that address is on Core and not on Firewall which is providing NAT.
Per your own diagram, Core.
In summary, your diagram along with the implications from your questions suggest an unworkable design that you've concocted. It would appear that you do not understand the basics of IP routing and networking, let alone trying to set up something complicated involving ECMP, BGP and PPPoE.
In other words, you need to hire a consultant to design and stop shopping this around various networking vendor subreddits.
