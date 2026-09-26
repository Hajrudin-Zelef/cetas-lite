---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-ezpuu1-dual-wan-with-failover-up-to-date-2020-routeros-9cb04b98
title: "r-mikrotik-comments-ezpuu1-dual-wan-with-failover-up-to-date-2020-routeros-9cb04b98"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/r-mikrotik-comments-ezpuu1-dual-wan-with-failover-up-to-date-2020-routeros-9cb04b98.md
source_anchor: ""
source_lines: [1, 29]
sha256: 1882c07915ffd4e8949eb4dbbfff47bc01d83fbe9881c246034bca85f6209fa7
---

# r-mikrotik-comments-ezpuu1-dual-wan-with-failover-up-to-date-2020-routeros-9cb04b98

Dual WAN with failover - Up to date 2020 routeros method\tutorial 
        
    Hi guys,
I'm in need of configuring a RB951G with 2 WANs, one dynamic + one static with simple failover, no load-balancing.
There are many youtube and written tutorials\guides out there, some work, some stopped working with recent routeros versions or didn't work at all.
I was wondering what methods are you guys using that are up to date, known to work with latest routeros.
What do you think about this one for example -https://www.prinmath.com/ham/mikrotik-failover.htm
Section des commentaires
You can use static routes on both of your WAN interfaces, but with the different route distances.
For example, add two routes with Dst. Address to 0.0.0.0/0 and gateway to your static WAN interface. With route distance, for example, 3 hops. Or 4. As you wish
I'd recommend to use static WAN interface as main, and the dynamic as the reserved one.
So. After that set the second one the same way on your dynamic WAN, but with 10 as the route distance.
And in that case, if your static WAN is going to fall, then your traffic automaticly will routed to the dynamic WAN interface.
Do you have an export of a config like that? I'll look for some unused mikrotiks and set up a demo LAN with two simulated ISPs to test these scenarios, post the configs here :), it would be awsome to have some visual examples of different types of WANs, static, dynamic, PPPoE.
/ip route
add check-gateway=arp distance=5 dst-address=0.0.0.0/0 gateway={your WAN gateway-1}
add check-gateway=arp distance=10 dst-address=0.0.0.0/0 gateway={your WAN gateway-2}
And don't forget to set addresses on that interfaces.
It's on your device, where you want to set two WAN interfaces.
And that's it.
Clear and stylish, as i see it. Saw that decision somewhere, and i think it's exactly your case.
It’ll work. There’s about 3-4 ways to do this.
However you need to explain the use case better. Does next hop fall over, does the device that connects to one of the WAN interfaces fall over?
Hi,
The usual scenario is that WAN1 link stays up but Internet is not accesible. There are rare cases when the actual links fail. Some tutorials available on the web are not good for the above scenario and unless the link gets dropped fail-over don't happen.
When the internet fails, have you performed a traceroute to see where it’s failing?
The reason why I ask this is you’ll need to check/ping a host beyond that failure.
Depending on how fancy you want to be , I usually just write a gateway check script with scheduler on boot with different set distances depending on weighting required etc.Allows me to set how sensitive I want the link checking to be as well as getting syslog and email alerts when it goes up/down.
Haven't done the recursive routing check myself but should work I think.
