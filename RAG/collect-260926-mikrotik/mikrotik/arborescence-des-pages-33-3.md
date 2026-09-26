---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-33-3
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-33.md
source_anchor: ""
source_lines: [303, 461]
sha256: 51fc65302d270d85ac1def137f96a697e8a2fa72018ac808ca993783d6ec3994
---

# Description

```
[admin@CCR2004_2XS] > /routing route print where dst-address in 111.0.0.0/8 && afi=ip4
...
 Ac   afi=ip4 contribution=active dst-address=111.11.0.0/24 routing-table=vrfTest1 gateway=sfp-sfpplus1@vrfTest1 immediate-gw=sfp-sfpplus1 distance=0 scope=10 
       belongs-to="connected" local-address=111.11.0.1%sfp-sfpplus1@vrfTest1 
       debug.fwp-ptr=0x202421E0 
 Ay   afi=ip4 contribution=best-candidate dst-address=111.12.0.0/24 routing-table=vrfTest1 label=17 gateway=vrfTest2@vrfTest2 immediate-gw=sfp-sfpplus2 
       distance=200 scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-1-bgp-mpls-vpn-2-connected-export-import" 
       bgp.ext-communities=rt:1:2 .atomic-aggregate=no .origin=incomplete 
       debug.fwp-ptr=0x202425A0 
 Ay   afi=ip4 contribution=best-candidate dst-address=111.11.0.0/24 routing-table=vrfTest2 label=16 gateway=vrfTest1@vrfTest1 immediate-gw=sfp-sfpplus1 
       distance=200 scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-2-bgp-mpls-vpn-1-connected-export-import" 
       bgp.ext-communities=rt:1:1 .atomic-aggregate=no .origin=incomplete 
       debug.fwp-ptr=0x202424E0 
 Ac   afi=ip4 contribution=active dst-address=111.12.0.0/24 routing-table=vrfTest2 gateway=sfp-sfpplus2@vrfTest2 immediate-gw=sfp-sfpplus2 distance=0 scope=10 
       belongs-to="connected" local-address=111.12.0.1%sfp-sfpplus2@vrfTest2 
       debug.fwp-ptr=0x20242240 
```
 And IPv6 too:

```
[admin@CCR2004_2XS] /routing/route> print detail where dst-address in 2001::/8 && afi=ip6
...
 Ac   afi=ip6 contribution=active dst-address=2001:1::/64 routing-table=vrfTest1 gateway=sfp-sfpplus1@vrfTest1 immediate-gw=sfp-sfpplus1 distance=0 scope=10 
       belongs-to="connected" local-address=2001:1::1%sfp-sfpplus1@vrfTest1 
       debug.fwp-ptr=0x20242300 
 Ay   afi=ip6 contribution=active dst-address=2001:2::/64 routing-table=vrfTest1 label=17 gateway=vrfTest2@vrfTest2 immediate-gw=sfp-sfpplus2 distance=200 
       scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-1-bgp-mpls-vpn-2-connected-export-import" 
       bgp.ext-communities=rt:1:2 .atomic-aggregate=no .origin=incomplete 
       debug.fwp-ptr=0x202425A0 
 Ay   afi=ip6 contribution=active dst-address=2001:1::/64 routing-table=vrfTest2 label=16 gateway=vrfTest1@vrfTest1 immediate-gw=sfp-sfpplus1 distance=200 
       scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-2-bgp-mpls-vpn-1-connected-export-import" 
       bgp.ext-communities=rt:1:1 .atomic-aggregate=no .origin=incomplete 
       debug.fwp-ptr=0x202424E0 
 Ac   afi=ip6 contribution=active dst-address=2001:2::/64 routing-table=vrfTest2 gateway=sfp-sfpplus2@vrfTest2 immediate-gw=sfp-sfpplus2 distance=0 scope=10 
       belongs-to="connected" local-address=2001:2::1%sfp-sfpplus2@vrfTest2 
       debug.fwp-ptr=0x20242360 
```
 ## Dynamic Vrf-Lite route leaking (old workaround)

Before ROS v7.14 there were no mechanism to leak routes from one VRF instance to another within the same router.

As a workaround, it was possible to create a tunnel between two locally configure loopback addresses and assign each tunnel endpoint to its own VRF. Then it is possible to run either dynamic routing protocols or set up static routes to leak between both VRFs.

The downside of this approach is that tunnel must be created between each VRF where routes should be leaked (create a full mesh), which significantly complicates configuration even if there are just several VRFs, not to mention more complicated setups.

For example, to leak routes between 5 VRFs it would require n * ( n – 1) / 2 connections, which will lead to the setup with 20 tunnel endpoints and 20 OSPF instances on one router.

Example config with two VRFs of this method:

/interface bridge
add name=dummy_custC
add name=dummy_custB
add name=lo1
add name=lo2
/ip address
add address=111.255.255.1 interface=lo1 network=111.255.255.1
add address=111.255.255.2 interface=lo2 network=111.255.255.2
add address=172.16.1.0/24 interface=dummy_custC network=172.16.1.0
add address=172.16.2.0/24 interface=dummy_custB network=172.16.2.0
/interface ipip
add local-address=111.255.255.1 name=ipip-tunnel1 remote-address=111.255.255.2
add local-address=111.255.255.2 name=ipip-tunnel2 remote-address=111.255.255.1
/ip address
add address=192.168.1.1/24 interface=ipip-tunnel1 network=192.168.1.0
add address=192.168.1.2/24 interface=ipip-tunnel2 network=192.168.1.0
/ip vrf
add interfaces=ipip-tunnel1,dummy_custC name=custC
add interfaces=ipip-tunnel2,dummy_custB name=custB
/routing ospf instance
add disabled=no name=i2_custB redistribute=connected,static,copy router-id=192.168.1.1 routing-table=custB vrf=custB
add disabled=no name=i2_custC redistribute=connected router-id=192.168.1.2 routing-table=custC vrf=custC
/routing ospf area
add disabled=no instance=i2_custB name=custB_bb
add disabled=no instance=i2_custC name=custC_bb
/routing ospf interface-template
add area=custB_bb disabled=no networks=192.168.1.0/24
add area=custC_bb disabled=no networks=192.168.1.0/24

 Result:

```
[admin@rack1_b36_CCR1009] /routing/ospf/neighbor> print 
Flags: V - virtual; D - dynamic 
 0  D instance=i2_custB area=custB_bb address=192.168.1.1 priority=128 router-id=192.168.1.2 dr=192.168.1.1 bdr=192.168.1.2 
      state="Full" state-changes=6 adjacency=41m28s timeout=33s 
 1  D instance=i2_custC area=custC_bb address=192.168.1.2 priority=128 router-id=192.168.1.1 dr=192.168.1.1 bdr=192.168.1.2 
      state="Full" state-changes=6 adjacency=41m28s timeout=33s 
[admin@rack1_b36_CCR1009] /ip/route> print where routing-table=custB
Flags: D - DYNAMIC; A - ACTIVE; c, s, o, y - COPY
Columns: DST-ADDRESS, GATEWAY, DISTANCE
     DST-ADDRESS       GATEWAY                         DISTANCE
  DAo 172.16.1.0/24     192.168.1.1%ipip-tunnel2@custB       110
  DAc 172.16.2.0/24     dummy_custB@custB                      0
  DAc 192.168.1.0/24    ipip-tunnel2@custB                     0
[admin@rack1_b36_CCR1009] > /ip route/print where routing-table=custC
Flags: D - DYNAMIC; A - ACTIVE; c, o, y - COPY
Columns: DST-ADDRESS, GATEWAY, DISTANCE
    DST-ADDRESS       GATEWAY                         DISTANCE
  DAc 172.16.1.0/24     dummy_custC@custC                      0
  DAo 172.16.2.0/24     192.168.1.2%ipip-tunnel1@custC       110 
  DAc 192.168.1.0/24    ipip-tunnel1@custC                     0
```
 


## The simplest MPLS VPN setup

In this example, a rudimentary MPLS backbone (consisting of two Provider Edge (PE) routers PE1 and PE2) is created and configured to forward traffic between Customer Edge (CE) routers CE1 and CE2 routers that belong to *cust-one* VPN.

### CE1 Router

/ip address add address=10.1.1.1/24 interface=ether1 
# use static routing 
/ip route add dst-address=10.3.3.0/24 gateway=10.1.1.2

 


### CE2 Router

/ip address add address=10.3.3.4/24 interface=ether1 
/ip route add dst-address=10.1.1.0/24 gateway=10.3.3.3

 


### PE1 Router

/interface bridge add name=lobridge 
/ip address add address=10.1.1.2/24 interface=ether1 
/ip address add address=10.2.2.2/24 interface=ether2 
/ip address add address=10.5.5.2/32 interface=lobridge 
/ip vrf add name=cust-one interfaces=ether1 
/mpls ldp add enabled=yes transport-address=10.5.5.2 lsr-id=10.5.5.2
/mpls ldp interface add interface=ether2 
/routing bgp template set default as=65000 
/routing bgp vpn 
add vrf=cust-one \
  route-distinguisher=1.1.1.1:111 \
  import.route-targets=1.1.1.1:111 \
  import.router-id=cust-one \
  export.redistribute=connected \
  export.route-targets=1.1.1.1:111 \
  label-allocation-policy=per-vrf
/routing bgp connection 
add template=default remote.address=10.5.5.3 address-families=vpnv4 local.address=10.5.5.2
# add route to the remote BGP peer's loopback address 
/ip route add dst-address=10.5.5.3/32 gateway=10.2.2.3

 




### PE2 Router (Cisco)

 

