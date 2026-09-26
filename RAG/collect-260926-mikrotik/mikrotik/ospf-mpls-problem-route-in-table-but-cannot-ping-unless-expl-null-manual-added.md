---
id: collect-260926-mikrotik/mikrotik/ospf-mpls-problem-route-in-table-but-cannot-ping-unless-expl-null-manual-added
title: "ospf-mpls-problem-route-in-table-but-cannot-ping-unless-expl-null-manual-added"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/ospf-mpls-problem-route-in-table-but-cannot-ping-unless-expl-null-manual-added.md
source_anchor: ""
source_lines: [1, 140]
sha256: 834569e9e87db69f43f1086516ec027a4fb98145b47866b69432bca19019a621
---

# ospf-mpls-problem-route-in-table-but-cannot-ping-unless-expl-null-manual-added

Hello I’ve a problem with OSPF with the following configuration.

ROUTER 2 (have internet) ↔ ROUTER 1 (PPPoE server)

In the PPPoE server there is a client pppoe-test1 that has remote address 100.80.0.12 and I route him a public /29 through it.

Routerboard at the client location has the first IP of /29.

The problem is that I cannot ping the 37.x.x.x/29 from the router 2 so also from internet it’s not reachable.




If in PPPoE router I add the following all works, but I would not do this for all the pppoe routed subnets.

```
/mpls local-bindings
add dst-address=37.x.x.x/29 label=expl-nul
```

ROUTER 1 - PPPOE SERVER CONFIG

```
/interface bridge
add name=BRIDGE_PPPOE
add name=loopback
/routing ospf area
add area-id=0.0.0.1 default-cost=1 inject-summary-lsas=no name=pppoe-area type=stub
/routing ospf instance
set [ find default=yes ] redistribute-static=as-type-1 router-id=10.100.0.1
/interface bridge port
add bridge=BRIDGE_PPPOE interface=sfp3
/interface pppoe-server server
add authentication=chap,mschap2 default-profile=default disabled=no interface=BRIDGE_PPPOE keepalive-timeout=30 max-mru=1500 max-mtu=1500 service-name=PPPOE-SERVER
/ip address
add address=10.0.0.1/29 comment="PTP1" interface=sfp1 network=10.0.0.0
add address=10.100.0.1 comment=LOOPBACK interface=loopback network=10.100.0.1
add address=10.0.0.49/29 comment=PTP2 interface=sfp2 network=10.0.0.48
/mpls interface
set [ find default=yes ] mpls-mtu=1530
/mpls ldp
set enabled=yes lsr-id=10.100.0.1 transport-address=10.100.0.1
/mpls ldp interface
add interface=sfp1
add interface=sfp2
/routing ospf area range
add area=pppoe-area range=37.y.y.y/27
add area=pppoe-area range=100.80.0.0/16
/routing ospf interface
add interface=loopback network-type=broadcast passive=yes
add authentication=md5 authentication-key=xxx comment=PT1 interface=sfp1 network-type=nbma priority=100
add authentication=md5 authentication-key=xxx comment=PTP2 cost=150 interface=sfp2 network-type=nbma priority=100
add comment="ALL PASSIVE" network-type=broadcast passive=yes
/routing ospf nbma-neighbor
add address=10.0.0.6 poll-interval=30s priority=100
add address=10.0.0.54 poll-interval=30s priority=100
/routing ospf network
add area=backbone network=10.0.0.0/29
add area=backbone network=10.100.0.1/32
add area=backbone network=10.0.0.48/29
add area=pppoe-area network=37.y.y.y/27
add area=pppoe-area network=100.80.0.0/16
```

ROUTER 2

```
/interface bridge
add name=loopback
/ip address
add address=10.100.0.2 comment=LOOPBACK interface=loopback network=10.100.0.2
add address=10.0.0.6/29 comment="PTP 1" interface=ether1 network=10.0.0.0
add address=94.x.x.x/30 comment="WAN" interface=ether5
add address=10.0.0.54/29 comment="PTP 2" interface=ether6 network=10.0.0.48
/ip route
add distance=1 gateway=94.x.x.x
/mpls interface
set [ find default=yes ] mpls-mtu=1530
/mpls ldp
set enabled=yes lsr-id=10.100.0.2 transport-address=10.100.0.2
/mpls ldp interface
add interface=ether1
add interface=ether6
/routing ospf interface
add interface=loopback network-type=broadcast passive=yes
add authentication=md5 authentication-key=xxx cost=150 interface=ether1 network-type=nbma
add authentication=md5 authentication-key=xxx cost=100 interface=ether6 network-type=nbma
/routing ospf nbma-neighbor
add address=10.0.0.1 poll-interval=30s priority=1
add address=10.0.0.49 poll-interval=30s priority=1
/routing ospf network
add area=backbone network=10.100.0.2/32
add area=backbone network=10.0.0.0/29
add area=backbone network=10.0.0.48/29
```

             
            
           
          
            
            
              What version of RouterOS is on these two routers and can you post the mpls forwarding table as well as the routing table on each?

             
            
           
          
            
            
              Instead of EXPL-NULL I put the subnet routed through PPPoE in the filter and so the PPPoE disappeared from the MPLS forwarding table.

```
/mpls ldp advertise-filter
add advertise=no prefix=37.x.x.x/27
```

At the moment I cant remove the filter because it’s in production but these are the tables (before I have as I said the PPPoE interface in the forwarding table of PPPoE server and a label assigned to the /29 route in local binding).

ROUTER 1 PPPoE Server

ROUTER 2

 
            
           
          
            
            
              Versions are 6.32.3 and 6.32.4

             
            
           
          
            
            
              Anyone solved this problem? I have the same issue with 6.42.5

UPD: Sorry, I forgot to apply advertise filter, also needed to disable/enable LDP on all routers.
