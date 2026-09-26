---
id: collect-260926-mikrotik/mikrotik/questions-661493-mikrotik-routing-a-single-address-part-of-a-direct-accessed-sub-f65eefdd
title: "jan/22/2015 13:38:30 by RouterOS 6.25"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-661493-mikrotik-routing-a-single-address-part-of-a-direct-accessed-sub-f65eefdd.md
source_anchor: ""
source_lines: [1, 48]
sha256: 08f139a8c104065a727706918a06554682fcce45d32ecfde98ad94f73b83a1a1
---

# jan/22/2015 13:38:30 by RouterOS 6.25

I have a Mikrotik RB2011 and several TP-Links - WR740N, located at different geo locations, part of my ISP MAN network.
My ISP provides me with an (static) address/mask and a gateway for each device. Ie:
- routerboard - 192.168.5.10/24 - gate 192.168.5.1
- TPLink1 - 192.168.10.5/24 - gate 192.168.10.1
- TPLink2 - 192.168.20.10/24 - gate 192.168.20.1
- TPLink3 - 192.168.30.15/24 - gate 192.168.30.1
- etc...
Because the routerboard has more than one WAN address, I configured the routes to the networks from above this way:
- accessing TPLink1 - 192.168.10.0/24 -> 192.168.5.1
- accessing TPLink2 - 192.168.20.0/24 -> 192.168.5.1
- accessing TPLink3 - 192.168.30.0/24 -> 192.168.5.1
- and so on...
Everything goes fine. I'm able to access each TPLink from the routerboard.
But I have a TPLinkX with an assigned address 192.168.5.6/29 and its network (5.0/29) is physically different from the routerboards one (5.0/24).
So, I added a new route (routerboard site) - 192.168.5.6/32 -> 192.168.5.1 and everything works, but after some time (5-10-15-20 minutes) this route becomes ignored. If I disable the route and enable it again - it becomes to work again (again for a short period of time). By the way, I'm surprised that it even works (although for a brief), because by default I have a dynamic route - 192.168.5.0/24 -> interface with a distance of 0 (generated because the static WAN address).
Is there any way to "bypass" the default route just for one host (or another approach) ?
Thanks in advance
EDIT
/ip routes
# jan/22/2015 13:38:30 by RouterOS 6.25
# software id = 8IZ2-4V85
 0 A S  dst-address=192.168.5.6/32 gateway=192.168.5.1
        gateway-status=192.168.5.1 reachable via  ether1-gateway distance=1 
        scope=30 target-scope=10
 1 ADS  dst-address=0.0.0.0/0 gateway=XXX.XXX.XXX.XXX 
        gateway-status=XXX.XXX.XXX.XXX reachable via  pppoe distance=1 
        scope=30 target-scope=10 
 3 ADC  dst-address=192.168.1.0/24 pref-src=192.168.1.1 gateway=ether2 
        gateway-status=ether2 reachable distance=0 scope=10 
 4 A S  dst-address=192.168.8.0/24 gateway=192.168.5.1 
        gateway-status=192.168.5.1 reachable via  ether1-gateway distance=1 
        scope=30 target-scope=10 
14 A S  dst-address=192.168.12.0/24 gateway=192.168.5.1 
        gateway-status=192.168.5.1 reachable via  ether1-gateway distance=1 
        scope=30 target-scope=10 
15 A S  dst-address=192.168.20.0/24 gateway=192.168.5.1 
        gateway-status=192.168.5.1 reachable via  ether1-gateway distance=1 
        scope=30 target-scope=10 
16 A S  dst-address=192.168.24.0/24 gateway=192.168.5.1 
        gateway-status=192.168.5.1 reachable via  ether1-gateway distance=1 
        scope=30 target-scope=10 
17 ADC  dst-address=192.168.5.0/22 pref-src=192.168.5.11 
        gateway=ether1-gateway gateway-status=ether1-gateway reachable 
        distance=0 scope=10 
23 ADC  dst-address=XXX.XXX.XXX.XXX/32 pref-src=XXX.XXX.XXX.XXX gateway=pppoe client
        gateway-status=pppoe reachable distance=0 scope=10 
The problematic one is the first route. It works for a while, but then suddenly becomes ignored.
/ip route print detail
