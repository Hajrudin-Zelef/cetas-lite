---
id: collect-260926-mikrotik/mikrotik/default-route-from-bgp-to-ospf
title: "default-route-from-bgp-to-ospf"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/default-route-from-bgp-to-ospf.md
source_anchor: ""
source_lines: [1, 71]
sha256: d8184ae104cefaf055ca7bb59d7990a9ae7205c2dfb81452a67f777c5423b84f
---

# default-route-from-bgp-to-ospf

You can definitely redistribute the default route from BGP into OSPF.  On the router that is doing the redistribution from BGP into OSPF, use the following config and the default will be advertised in OSPF only as long as it exists in BGP.

```
/routing ospf instance
set [ find default=yes ] distribute-default=if-installed-as-type-1
```

**Here is an example from our lab**

```
[admin@IPA-LAB-OSPF-Customer] > ip route print detail
Flags: X - disabled, A - active, D - dynamic, 
C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, 
B - blackhole, U - unreachable, P - prohibit 
 0 ADo  dst-address=0.0.0.0/0 gateway=100.65.0.1 
        gateway-status=100.65.0.1 reachable via  ether5 distance=110 scope=20 
        target-scope=10 ospf-metric=11 ospf-type=external-type-1 
 1 ADC  dst-address=100.65.0.0/30 pref-src=100.65.0.2 gateway=ether5 
        gateway-status=ether5 reachable distance=0 scope=10 
[admin@IPA-LAB-OSPF-Customer] >
```

**Logs from  the BGP/OSPF router when the BGP default is withdrawn in the upstream BGP Internet peer**

```
23:51:57 route,debug Target withdraw  
23:51:57 route,debug     prefix=0.0.0.0/0 
23:51:57 route,debug     previous attributes 
23:51:57 route,debug         protocol=BGP 
23:51:57 route,debug         scope=40 
23:51:57 route,debug         immediate-next-hop= address=100.64.0.1 interface=ether5 
23:51:57 route,debug         next-hop= address=100.64.0.1 
23:51:57 route,debug         origin-type=BGP 
23:51:57 route,debug         origin-instance-id=0 
23:51:57 route,debug         bgp-peer-router-id=1.1.1.2 
23:51:57 route,debug         bgp-peer-flags=0 
23:51:57 route,debug         bgp-router-id=2.2.2.2 
23:51:57 route,debug         bgp-origin=IGP 
23:51:57 route,debug         bgp-as-path=7018 
23:51:57 route,debug         bgp-as-path-len=1 
23:51:57 route,debug         bgp-nexthop=100.64.0.1 
23:51:57 route,debug         use-te-nexthop=yes
```

**Same as above but as viewed from the OSPF only router**

```
23:40:57 route,debug Target withdraw  
23:40:57 route,debug     prefix=0.0.0.0/0 
23:40:57 route,debug     previous attributes 
23:40:57 route,debug         protocol=OSPF 
23:40:57 route,debug         scope=20 
23:40:57 route,debug         immediate-next-hop= address=100.65.0.1 interface=ethe
r5 
23:40:57 route,debug         next-hop= address=100.65.0.1 
23:40:57 route,debug         origin-type=OSPF 
23:40:57 route,debug         origin-instance-id=0 
23:40:57 route,debug         ospf-metric=11 
23:40:57 route,debug         ospf-type=external-type-1
```

**Finally, the routing table on the OSPF router after disabling it in the upstream BGP router.**

```
[admin@IPA-LAB-OSPF-Customer] > ip route print detail
Flags: X - disabled, A - active, D - dynamic, 
C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme, 
B - blackhole, U - unreachable, P - prohibit 
 0 ADC  dst-address=100.65.0.0/30 pref-src=100.65.0.2 gateway=ether5 
        gateway-status=ether5 reachable distance=0 scope=10
```
