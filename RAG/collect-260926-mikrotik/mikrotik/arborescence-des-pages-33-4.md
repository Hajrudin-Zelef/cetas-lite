---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-33-4
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-33.md
source_anchor: ""
source_lines: [462, 644]
sha256: 58a817f51b5de646074e1d5f44a8cd3d9b8fcd0038f2528c5069e6c2adcc4479
---

# Description

ip vrf cust-one
rd 1.1.1.1:111
route-target export 1.1.1.1:111
route-target import 1.1.1.1:111
exit
interface Loopback0
ip address 10.5.5.3 255.255.255.255
mpls ldp router-id Loopback0 force
mpls label protocol ldp
interface FastEthernet0/0
ip address 10.2.2.3 255.255.255.0
mpls ip
interface FastEthernet1/0
ip vrf forwarding cust-one
ip address 10.3.3.3 255.255.255.0
router bgp 65000
neighbor 10.5.5.2 remote-as 65000
neighbor 10.5.5.2 update-source Loopback0
address-family vpnv4
neighbor 10.5.5.2 activate
neighbor 10.5.5.2 send-community both
exit-address-family
address-family ipv4 vrf cust-one
redistribute connected
exit-address-family
ip route 10.5.5.2 255.255.255.255 10.2.2.2

 Results

Check that VPNv4 route redistribution is working:

 

```
[admin@PE1] /routing/route> print detail where afi="vpn4" 
Flags: X - disabled, F - filtered, U - unreachable, A - active; 
c - connect, s - static, r - rip, b - bgp, o - ospf, d - dhcp, v - vpn, m - modem, a - ldp-address, l - l
dp-mapping, g - slaac, y - bgp-mpls-vpn; 
H - hw-offloaded; + - ecmp, B - blackhole 
 Ab   afi=vpn4 contribution=active dst-address=111.16.0.0/24&1.1.1.1:111 routing-table=main label=16 
       gateway=111.111.111.4 immediate-gw=111.13.0.2%ether9 distance=200 scope=40 target-scope=30 
       belongs-to="bgp-VPN4-111.111.111.4" 
       bgp.peer-cache-id=*2C00011 .as-path="65511" .ext-communities=rt:1.1.1.1:111 .local-pref=100 
       .atomic-aggregate=yes .origin=igp 
       debug.fwp-ptr=0x202427E0 
[admin@PE1] /routing/bgp/advertisements> print 
 0 peer=to-pe2-1 dst=10.1.1.0/24 local-pref=100 origin=2 ext-communities=rt:1.1.1.1:111 atomic-aggregate=yes 
  
```
 Check that the 10.3.3.0 is installed in IP routes, in the cust-one route table:

[admin@PE1] > /ip route print where routing-table="cust-one" 
Flags: D - DYNAMIC; A - ACTIVE; c, b, y - BGP-MPLS-VPN
Columns: DST-ADDRESS, GATEWAY, DISTANCE
# DST-ADDRESS     GATEWAY         DISTANCE 
0 ADC 10.1.1.0/24 ether1@cust-one        0 
1 ADb 10.3.3.0/24 10.5.5.3              20 

 


Let's take a closer look at IP routes in cust-one VRF. The 10.1.1.0/24 IP prefix is a connected route that belongs to an interface that was configured to belong to cust-one VRF. The 10.3.3.0/24 IP prefix was advertised via BGP as a VPNv4 route from PE2 and is imported in this VRF routing table, because our configured **import-route-targets** matched the BGP extended communities attribute it was advertised with.

```
[admin@PE1] /routing/route> print detail where routing-table="cust-one"
Flags: X - disabled, F - filtered, U - unreachable, A - active; 
c - connect, s - static, r - rip, b - bgp, o - ospf, d - dhcp, v - vpn, m - modem, a - ldp-address, l - l
dp-mapping, g - slaac, y - bgp-mpls-vpn; 
H - hw-offloaded; + - ecmp, B - blackhole 
 Ac   afi=ip4 contribution=active dst-address=10.1.1.0/24 routing-table=cust-one 
       gateway=ether1@cust-one immediate-gw=ether1 distance=0 scope=10 belongs-to="connected" 
       local-address=10.1.1.2%ether1@cust-one
       debug.fwp-ptr=0x202420C0 
 Ay   afi=ip4 contribution=active dst-address=10.3.3.0/24 routing-table=cust-one label=16 
       gateway=10.5.5.3 immediate-gw=10.2.2.3%ether2 distance=20 scope=40 target-scope=30 
       belongs-to="bgp-mpls-vpn-1-bgp-VPN4-10.5.5.3-import" 
       bgp.peer-cache-id=*2C00011 .ext-communities=rt:1.1.1.1:111 .local-pref=100 
       .atomic-aggregate=yes .origin=igp 
       debug.fwp-ptr=0x20242840 
[admin@PE1] /routing/route> print detail where afi="vpn4"                 
Flags: X - disabled, F - filtered, U - unreachable, A - active; 
c - connect, s - static, r - rip, b - bgp, o - ospf, d - dhcp, v - vpn, m - modem, a - ldp-address, l - l
dp-mapping, g - slaac, y - bgp-mpls-vpn; 
H - hw-offloaded; + - ecmp, B - blackhole 
 Ay   afi=vpn4 contribution=active dst-address=10.1.1.0/24&1.1.1.1:111 routing-table=main label=19 
       gateway=ether1@cust-one immediate-gw=ether1 distance=200 scope=40 target-scope=10 
       belongs-to="bgp-mpls-vpn-1-connected-export" 
       bgp.ext-communities=rt:1.1.1.1:1111 .atomic-aggregate=no .origin=incomplete 
       debug.fwp-ptr=0x202426C0 
 Ab   afi=vpn4 contribution=active dst-address=10.3.3.0/24&1.1.1.1:111 routing-table=main label=16 
       gateway=10.5.5.3 immediate-gw=10.2.2.3%ether2 distance=200 scope=40 target-scope=30 
       belongs-to="bgp-VPN4-10.5.5.3" 
       bgp.peer-cache-id=*2C00011 .ext-communities=rt:1.1.1.1:111 .local-pref=100 
       .atomic-aggregate=yes .origin=igp 
       debug.fwp-ptr=0x202427E0 
```
 


The same for Cisco:

PE2#show ip bgp vpnv4 all 
BGP table version is 5, local router ID is 10.5.5.3 
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal, 
r RIB-failure, S Stale 
Origin codes: i - IGP, e - EGP, ? - incomplete 
Network Next Hop Metric LocPrf Weight Path 
Route Distinguisher: 1.1.1.1:111 (default for vrf cust-one) 
*>i10.1.1.0/24 10.5.5.2 100 0 ? 
*> 10.3.3.0/24 0.0.0.0 0 32768 ? 
PE2#show ip route vrf cust-one 
Routing Table: cust-one 
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP 
D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2 
E1 - OSPF external type 1, E2 - OSPF external type 2 
i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2 
ia - IS-IS inter area, * - candidate default, U - per-user static route 
o - ODR, P - periodic downloaded static route 
Gateway of last resort is not set 
10.0.0.0/24 is subnetted, 1 subnets
B 10.1.1.0 [200/0] via 10.5.5.2, 00:05:33 
10.0.0.0/24 is subnetted, 1 subnets 
C 10.3.3.0 is directly connected, FastEthernet1/0

 


You should be able to ping from CE1 to CE2 and vice versa.

[admin@CE1] > /ping 10.3.3.4 
10.3.3.4 64 byte ping: ttl=62 time=18 ms 
10.3.3.4 64 byte ping: ttl=62 time=13 ms 
10.3.3.4 64 byte ping: ttl=62 time=13 ms 
10.3.3.4 64 byte ping: ttl=62 time=14 ms 
4 packets transmitted, 4 packets received, 0% packet loss 
round-trip min/avg/max = 13/14.5/18 ms

 


## A more complicated setup (changes only)

As opposed to the simplest setup, in this example, we have two customers: cust-one and cust-two.

We configure two VPNs for them, cust-one and cust-two respectively, and exchange all routes between them. (This is also called "route leaking").

Note that this could be not the most typical setup, because routes are usually not exchanged between different customers. In contrast, by default, it should not be possible to gain access from one VRF site to a different VRF site in another VPN. (This is the "Private" aspect of VPNs.) Separate routing is a way to provide privacy, and it is also required to solve the problem of overlapping IP network prefixes. Route exchange is in direct conflict with these two requirements but may sometimes be needed (e.g. temp. solution when two customers are migrating to a single network infrastructure).

### CE1 Router, *cust-one*

/ip route add dst-address=10.4.4.0/24 gateway=10.1.1.2

 


### CE2 Router, *cust-one*

 

/ip route add dst-address=10.4.4.0/24 gateway=10.3.3.3

 CE1 Router,*cust-two*

/ip address add address=10.4.4.5 interface=ether1 
/ip route add dst-address=10.1.1.0/24 gateway=10.3.3.3 
/ip route add dst-address=10.3.3.0/24 gateway=10.3.3.3

 


### PE1 Router

# replace the old BGP VPN with this:
/routing bgp vpn 
add vrf=cust-one \
  export.redistribute=connected \
  route-distinguisher=1.1.1.1:111 \
  import.route-targets=1.1.1.1:111,2.2.2.2:222  \
  export.route-targets=1.1.1.1:111

 PE2 Router (Cisco)

