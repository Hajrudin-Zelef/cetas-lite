---
id: collect-260926-mikrotik/mikrotik/bgp-and-ospf-network-configuration
title: "bgp-and-ospf-network-configuration"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/bgp-and-ospf-network-configuration.md
source_anchor: ""
source_lines: [1, 85]
sha256: 680c99f4e263c80d7edbedb41e04275d2e9c3c86aa791c5d2f35294b7540928b
---

# bgp-and-ospf-network-configuration

Hi Guys / Girls,

Looking for some general advice on moving my network to BGP routing. Currently I have two RB1100AHx2 in different locations. We have a layer 2 100mbit transit between the two routers and each router has it’s own transit to the internet (currently static routing). Each router is responsible for it’s own public /27 which is advertised (OSPF) on the cross connect link as well to route the internal traffic.

After talking with the ISP they can configure us with full feed routing on both links, so we can utilize both links if one goes down and gives us the ability to move address space around easier on our network. As we only peer with them for now we will be using a private ASN. I have minimal experience with BGP so looking for some feedback on my idea.

OSPF is currently configured on the internal link like this:

R1:

```
add area=backbone disabled=no network=10.222.0.0/24
add area=backbone disabled=no network=X.X.245.160/27
```

R2:

```
add area=backbone disabled=no network=10.222.0.0/24
add area=backbone disabled=no network=X.X.137.96/27
```

In my test network (two little RB750’s) I have configured a BGP session between the two routers and added a test eBGP peer (a RB751G) on each router:

R1:

```
/routing bgp instance
set default as=20 router-id=10.10.10.10
/routing bgp network
add disabled=no network=X.X.245.160/27 synchronize=no
/routing bgp peer
add address-families=ip as-override=no default-originate=never disabled=no \
    hold-time=30s in-filter="" instance=default multihop=no name=router2 \
    nexthop-choice=default out-filter="" passive=no remote-address=10.222.0.2 \
    remote-as=20 remove-private-as=no route-reflect=no tcp-md5-key="" ttl=\
    default use-bfd=no
add address-families=ip as-override=no default-originate=never disabled=no \
    hold-time=1m in-filter=ISP1-in instance=default multihop=no name=ISP1 \
    nexthop-choice=default out-filter=ISP1-out passive=no remote-address=\
    40.0.10.22 remote-as=444 remove-private-as=no route-reflect=no tcp-md5-key=\
    "" ttl=default use-bfd=no
/routing filter
add action=accept chain=ISP1-out disabled=no in
    X.X.245.160/27 set-bgp-prepend-path=""
add action=accept chain=ISP1-out disabled=no in
    X.X.137.96/27 set-bgp-prepend=2 set-bgp-pre
add action=discard chain=ISP1-out disabled=no i
    set-bgp-prepend-path=""
```

R2:

```
/routing bgp instance
set default as=20 router-id=20.20.20.20
/routing bgp network
add disabled=no network=X.X.137.96/27 synchronize=no
/routing bgp peer
add address-families=ip as-override=no default-originate=never disabled=no \
    hold-time=30s in-filter="" instance=default multihop=no name=router1 \
    nexthop-choice=default out-filter="" passive=no remote-address=10.222.0.1 \
    remote-as=20 remove-private-as=no route-reflect=no tcp-md5-key="" ttl=\
    default use-bfd=no
add address-families=ip as-override=no default-originate=never disabled=no \
    hold-time=1m in-filter=ISP2-in instance=default multihop=no name=ISP2 \
    nexthop-choice=default out-filter=ISP2-out passive=no remote-address=\
    40.0.20.22 remote-as=444 remove-private-as=no route-reflect=no tcp-md5-key=\
    "" ttl=default use-bfd=no
/routing filter
add action=accept chain=ISP2-out disabled=no in
    X.X.137.96/27 set-bgp-prepend-path=""
add action=accept chain=ISP2-out disabled=no in
    X.X.245.160/27 set-bgp-prepend=2 set-bgp-pre
add action=discard chain=ISP2-out disabled=no i
    set-bgp-prepend-path=""
```

This sort of works, I get prefixes from the ISP peers and I can see them recursively over the iBGP session. The prefix which I announce to the ISP is also seen but they don’t seem to announce the other prefix which it has learned over the iBGP session to the ISP. I found if I remove the network from OSPF “networks” it does announce it and the prepend path works fine.

How can I work around this behaviour, am I doing something wrong, or is there a better way of doing this? Obviously I want to announce the internal prefixes so they route over the internal link and if the BGP sessions drop I still want to have my internal routes.

Also I found if I pull one of the ISP cables the routes learned from that peer stay in the routing table and say “next hop unreachable” until the BGP session dies. When this happens is there a way to remove them instantly as I have the routes learned from iBGP which could be used but they are not as they have a lower priority.

Sorry for all the questions, but any advice would be appreciated
