---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-33-5
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-33.md
source_anchor: ""
source_lines: [645, 809]
sha256: fc02bca2de94bffcc340bdf622483d4bea3b10ec82a66de7892c88bc69d80e86
---

# Description

ip vrf cust-one 
rd 1.1.1.1:111 
route-target export 1.1.1.1:111 
route-target import 1.1.1.1:111 
route-target import 2.2.2.2:222 
exit 
ip vrf cust-two 
rd 2.2.2.2:222 
route-target export 2.2.2.2:222 
route-target import 1.1.1.1:111 
route-target import 2.2.2.2:222 
exit 
interface FastEthernet2/0 
ip vrf forwarding cust-two 
ip address 10.4.4.3 255.255.255.0 
router bgp 65000 
address-family ipv4 vrf cust-two 
redistribute connected 
exit-address-family

 


## Variation: replace the Cisco with another MT

### PE2 Mikrotik config

 

/interface bridge add name=lobridge
/ip address
add address=10.2.2.3/24 interface=ether1
add address=10.3.3.3/24 interface=ether2
add address=10.4.4.3/24 interface=ether3
add address=10.5.5.3/32 interface=lobridge
/ip vrf
add name=cust-one interfaces=ether2
add name=cust-two interfaces=ether3
/mpls ldp add enabled=yes transport-address=10.5.5.3
/mpls ldp interface add interface=ether1
/routing bgp template set default as=65000 
/routing bgp vpn 
add vrf=cust-one \
  export.redistribute=connected \
  route-distinguisher=1.1.1.1:111 \
  import.route-targets=1.1.1.1:111,2.2.2.2:222 \
  export.route-targets=1.1.1.1:111 \
add vrf=cust-two \
  export.redistribute=connected \
  route-distinguisher=2.2.2.2:222 \
  import.route-targets=1.1.1.1:111,2.2.2.2:222 \
  export.route-targets=2.2.2.2:222 \
/routing bgp connection 
add template=default remote.address=10.5.5.2 address-families=vpnv4 local.address=10.5.5.3
# add route to the remote BGP peer's loopback address
/ip route add dst-address=10.5.5.2/32 gateway=10.2.2.2

 Results

The output of **/ip route print** now is interesting enough to deserve detailed observation.

 

[admin@PE2] /ip route> print
Flags: X - disabled, A - active, D - dynamic,
C - connect, S - static, r - rip, b - bgp, o - ospf, m - mme,
B - blackhole, U - unreachable, P - prohibit
# DST-ADDRESS PREF-SRC GATEWAY DISTANCE
0 ADb 10.1.1.0/24 10.5.5.2 recurs... 20
1 ADC 10.3.3.0/24 10.3.3.3 ether2 0
2 ADb 10.4.4.0/24 20
3 ADb 10.1.1.0/24 10.5.5.2 recurs... 20
4 ADb 10.3.3.0/24 20
5 ADC 10.4.4.0/24 10.4.4.3 ether3 0
6 ADC 10.2.2.0/24 10.2.2.3 ether1 0
7 A S 10.5.5.2/32 10.2.2.2 reacha... 1
8 ADC 10.5.5.3/32 10.5.5.3 lobridge 0

 The route 10.1.1.0/24 was received from a remote BGP peer and is installed in both VRF routing tables.

The routes 10.3.3.0/24 and 10.4.4.0/24 are also installed in both VRF routing tables. Each is a connected route in one table and a BGP route in another table. This has nothing to do with their being advertised via BGP. They are simply being "advertised" to the local VPNv4 route table and locally reimported after that. Import and export **route-targets** determine in which tables they will end up.

This can be deduced from its attributes - they don't have the usual BGP properties. (Route 10.4.4.0/24.)

 

[admin@PE2] /routing/route> print detail where routing-table=cust-one
...

 


## Unique RD per-site vs unique RD per-customer

Let's consider BGP VPN setup where two CUST_A sites announce the same network (111.12.0.0/24).

```
                         +----------+               +----------+ 
                         |+-vrf-+   |               |   +-vrf-+|
CUST_A(10.12.0.0/24)-----|+-----+   |---(BGP VPN)---|   +-----+|-------CUST_A(10.12.0.0/24)
                         +----------+               +----------+ 
```
 There are two ways to handle setups like this:

- per-customer (CUST_A VPN on Router 1 have the same route distinguisher (lets assume RD=1:1) as CUST_A VPN on Router 2)
- per-site (each CUST_A site  have unique Route Distinguisher)

In first case CUST_A sites on Router1  have exported the VPNv4 route and advertise it to remote PE.  

```
 Ay   afi=vpnv4 contribution=active dst-address=111.12.0.0/24&1:1 routing-table=main label=16 gateway=vrf-dummy@vrfTest 
       immediate-gw=vrf-dummy distance=200 scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-1-connected-export" 
       bgp.ext-communities=rt:1:1 .origin=incomplete 
       debug.fwp-ptr=0x20302600 
```
 


CUST_A site on Router2 also exports VPNv4 route and has received VPNv4 route from another site as well:

```
 Ab + afi=vpnv4 contribution=active dst-address=111.12.0.0/24&1:1 routing-table=main label=16 gateway=111.11.0.2 
       immediate-gw=111.11.0.2%sfp-sfpplus1 distance=200 scope=40 target-scope=30 belongs-to="bgp-VPNv4-111.11.0.2" 
       bgp.session=to-tested-1 .ext-communities=rt:1:1 .local-pref=100 .origin=igp 
       debug.fwp-ptr=0x203421E0 
 Ay + afi=vpnv4 contribution=active dst-address=111.12.0.0/24&1:1 routing-table=main label=32 gateway=vrf-dummy@vrfTest 
       immediate-gw=vrf-dummy distance=200 scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-1-connected-export" 
       bgp.ext-communities=rt:1:1 .origin=incomplete 
       debug.fwp-ptr=0x20342540 
```
 Currently RouterOS advertises only one best route. In case of ECMP by default it picks the first one from the list which happens to be BGP VPNv4 route received from remote site, and of course remote site will not get the second route. This leads to situation that one site has two redundant routes in the VRF but other site does not. On that site where VRF does not have installed VPN route and local route becomes inactive, BGP needs to send withdraw, recalculate main table, receive update from remote site and import new best route into CUST_A VRF, leading to slower convergence.

 This behavior of course could be changed with route selection filters, but it is outside the scope of this example.

Similar situation happens if exported to VPNV4 route is also BGP route received from customers CE-PE session, except that now instead of ECMP, BGP best-path selection process is applied to VPNv4 routes and only one best is selected.



Now if we assign unique route-distinguisher per site, lets say (1.1.1.1:1 on first site and 1.1.1.2:1 on second site), there is no longer selection on VPNv4 routes because these are now considered unique destinations and both destinations are imported into VRF.

CUST_A on Router1

```
[admin@CCR2004_2XS_111] /routing/route> print detail 
...
 Ay   afi=vpnv4 contribution=active dst-address=111.12.0.0/24&1.1.1.1:1 routing-table=main label=17 gateway=vrf-dummy@vrfTest 
       immediate-gw=vrf-dummy distance=200 scope=40 target-scope=10 belongs-to="bgp-mpls-vpn-1-connected-export" 
       bgp.ext-communities=rt:1:1 .origin=incomplete 
       debug.fwp-ptr=0x20302240 
 Ab   afi=vpnv4 contribution=active dst-address=111.12.0.0/24&1.1.1.2:1 routing-table=main label=33 gateway=111.11.0.1 
       immediate-gw=111.11.0.1%sfp-sfpplus1 distance=200 scope=40 target-scope=30 belongs-to="bgp-VPNv4-111.11.0.1" 
       bgp.session=to-tester-1 .ext-communities=rt:1:1 .local-pref=100 .origin=igp 
       debug.fwp-ptr=0x20302480 
[admin@CCR2004_2XS_111] /ip/route> print 
...
  DAc  111.12.0.0/24       vrf-dummy@vrfTest  vrfTest               0
  D y  111.12.0.0/24       111.11.0.1         vrfTest             200
```
 




CUST_A on Router2

