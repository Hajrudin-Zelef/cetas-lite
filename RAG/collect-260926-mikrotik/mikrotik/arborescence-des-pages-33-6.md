---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-33-6
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-33.md
source_anchor: ""
source_lines: [810, 834]
sha256: 9b2272409a660d53a250c587652bf9533a99af6923083ddb0ed49eea0da731b4
---

# Description

```
[admin@CCR2004_2XS_111] /routing/route> print detail 
...
 Ab   afi=vpnv4 contribution=active dst-address=111.12.0.0/24&1.1.1.1:1 routing-table=main label=17 gateway=111.11.0.2 
       immediate-gw=111.11.0.2%sfp-sfpplus1 distance=200 scope=40 target-scope=30 belongs-to="bgp-VPNv4-111.11.0.2" 
       bgp.session=to-tested-1 .ext-communities=rt:1:1 .local-pref=100 .origin=igp 
       debug.fwp-ptr=0x203421E0 
 Ay   afi=vpnv4 contribution=active dst-address=111.12.0.0/24&1.1.1.2:1 routing-table=main label=33 gateway=vrf-dummy@vrfTest 
       immediate-gw=vrf-dummy distance=200 scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-1-connected-export" 
       bgp.ext-communities=rt:1:1 .origin=incomplete 
       debug.fwp-ptr=0x20342240 
[admin@CCR2004_2XS_111] /ip/route> print 
...
 DAc 111.12.0.0/24       vrf-dummy@vrfTest  vrfTest               0
 D y 111.12.0.0/24       111.11.0.2         vrfTest             200
```
 In this case compared to previous setup, when local route from any of sites becomes inactive, switch to alternative route happens momentarily.



# References

RFC 4364: BGP/MPLS IP Virtual Private Networks (VPNs)

MPLS Fundamentals, chapter 7, *Luc De Ghein*, Cisco Press 2006
