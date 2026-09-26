---
id: collect-260926-mikrotik/mikrotik/ospf-and-load-balancing-example-of-configuration
title: "ospf-and-load-balancing-example-of-configuration"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ospf-and-load-balancing-example-of-configuration.md
source_anchor: ""
source_lines: [1, 156]
sha256: 9ab1e4c0423101c84773ea7b446509cbd87662ecb23f297e8b7064324f6dc45f
---

# ospf-and-load-balancing-example-of-configuration

I want to present my OSPF configuration with load balancing (by using mangle), so as to receive some suggestions, especially if somebody has ideas about introducing multicast and PIM into my configuration.

The purpose is to use redundant links to balance unicast (and possibly multicast) traffic coming from network cameras.

At the present, the failover and balancing mechanism works with unicast traffic. I have no ideas about moving to multicast…

I have two sites, each one composed of a dual radio board and a single radio board. So there are 3 wireless links running. All units run RouterOS3.17. Boards are RB433 and RB411.

All radio units are configured in OSPF mode. The system works perfectly, so video coming from CAM1 is sent through wlan1 of RU1-RU3, video coming from CAM2 is sent though wlan2 of RU1-RU3 and video coming from CAM3 is sent through wlan1 of RU2-RU4. If some of the links goes down, the video is automatically forwarded through one of the remaining links (metrics configuration).

When I try to use PIM (so multicast package), everything is completely dynamic and based on metrics, and I cannot use static IP routes (they accept only unicast packets) and mangle. How to establish a “preferred” route for multicast packets, like I did for unicast?

Here is the main configuration of Radio Unit 1 (RU1):

Wlan1 station, wlan2 ap-bridge

/routing ospf area

add area-id=0.0.0.1 authentication=none disabled=no name=area1 type=default

/routing ospf

set distribute-default=if-installed-as-type-2 metric-bgp=20 metric-connected=

20 metric-default=1 metric-rip=20 metric-static=20 mpls-te-area=

unspecified mpls-te-router-id=unspecified redistribute-bgp=no 

redistribute-connected=as-type-1 redistribute-rip=no redistribute-static=

no router-id=0.0.0.0

/routing ospf interface

add authentication=none authentication-key=“” authentication-key-id=1 cost=11 

dead-interval=40s disabled=no hello-interval=10s interface=wlan1 

network-type=broadcast passive=no priority=1 retransmit-interval=5s 

transmit-delay=1s

add authentication=none authentication-key=“” authentication-key-id=1 cost=10 

dead-interval=40s disabled=no hello-interval=10s interface=wlan2 

network-type=broadcast passive=no priority=1 retransmit-interval=5s 

transmit-delay=1s

add authentication=none authentication-key=“” authentication-key-id=1 cost=1 

dead-interval=40s disabled=no hello-interval=10s interface=ether1 

network-type=broadcast passive=no priority=1 retransmit-interval=5s 

transmit-delay=1s

/routing ospf network

add area=area1 disabled=no network=10.0.0.0/16

add area=area1 disabled=no network=172.19.0.0/16

/ip firewall mangle

add action=mark-packet chain=prerouting comment=“” disabled=no in-interface=

ether1 new-packet-mark=pktcam1 passthrough=yes src-address=172.19.40.2

add action=mark-routing chain=prerouting comment=“” disabled=no in-interface=

ether1 new-routing-mark=mk-link1 packet-mark=pktcam1 passthrough=no

add action=mark-packet chain=prerouting comment=“” disabled=no in-interface=

ether1 new-packet-mark=pktcam2 passthrough=yes src-address=172.19.40.3

add action=mark-routing chain=prerouting comment=“” disabled=no in-interface=

ether1 new-routing-mark=mk-link2 packet-mark=pktcam2 passthrough=no

add action=mark-packet chain=prerouting comment=“” disabled=no in-interface=

ether1 new-packet-mark=pktcam3 passthrough=yes src-address=172.19.40.4

add action=mark-routing chain=prerouting comment=“” disabled=no in-interface=

ether1 new-routing-mark=mk-link3 packet-mark=pktcam3 passthrough=no

/ip route

add comment=“” disabled=no distance=1 dst-address=172.20.0.0/16 gateway=

10.0.0.1 routing-mark=mk-link1 scope=30 target-scope=10

add comment=“” disabled=no distance=1 dst-address=172.20.0.0/16 gateway=

10.0.1.1 routing-mark=mk-link2 scope=30 target-scope=10

add comment=“” disabled=no distance=1 dst-address=172.20.0.0/16 gateway=

10.0.2.1 routing-mark=mk-link3 scope=30 target-scope=10

Here is the main configuration of Radio Unit 2 (RU2):

Wlan1 station

/routing ospf area

add area-id=0.0.0.1 authentication=none disabled=no name=area1 type=default

/routing ospf

set distribute-default=never metric-bgp=20 metric-connected=20 

metric-default=1 metric-rip=20 metric-static=20 mpls-te-area=unspecified 

mpls-te-router-id=unspecified redistribute-bgp=no redistribute-connected=

as-type-1 redistribute-rip=no redistribute-static=no router-id=0.0.0.0

/routing ospf interface

add authentication=none authentication-key=“” authentication-key-id=1 cost=11 

dead-interval=40s disabled=no hello-interval=10s interface=wlan1 

network-type=broadcast passive=no priority=1 retransmit-interval=5s 

transmit-delay=1s

add authentication=none authentication-key=“” authentication-key-id=1 cost=1 

dead-interval=40s disabled=no hello-interval=10s interface=ether1 

network-type=broadcast passive=no priority=1 retransmit-interval=5s 

transmit-delay=1s

/routing ospf network

add area=area1 disabled=no network=10.0.0.0/16

add area=area1 disabled=no network=172.19.0.0/16

The configuration of RU3 and RU4 is similar (except that there is no mangle).


             
            
           
          
            
            
              Anyone has some comments?
