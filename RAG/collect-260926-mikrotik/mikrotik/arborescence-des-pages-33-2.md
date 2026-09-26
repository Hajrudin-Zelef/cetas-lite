---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-33-2
title: "Description"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-33.md
source_anchor: ""
source_lines: [144, 302]
sha256: c70ac1d19741f79ca91f78ae5e2dbe2493eec6e646f3153586e2e141d2806978
---

# Description

```
# mark new customer connections
/ip firewall mangle 
add action=mark-connection chain=prerouting connection-state=new new-connection-mark=\
    cust_a_conn src-address=192.168.1.0/24 passthrough=no
add action=mark-connection chain=prerouting connection-state=new new-connection-mark=\
    cust_b_conn src-address=192.168.2.0/24 passthrough=no 
# mark routing
/ip firewall mangle  
add action=mark-routing chain=prerouting connection-mark=cust_a_conn \
    in-interface=public new-routing-mark=cust_a
add action=mark-routing chain=prerouting connection-mark=cust_b_conn \
    in-interface=public new-routing-mark=cust_b
```
 ## Static inter-VRF routes

In general, it is recommended that all routes between VRF should be exchanged using BGP local import and export functionality. If that is not enough, static routes can be used to achieve this so-called route leaking.

There are two ways to install a route that has a gateway in a different routing table than the route itself.

The first way is to explicitly specify the routing table in the gateway field when adding a route. This is only possible when leaking a route and gateway from the "main" routing table to a different routing table (VRF). Example:

# add route to 5.5.5.0/24 in 'vrf1' routing table with gateway in the main routing table 
add dst-address=5.5.5.0/24 gateway=10.3.0.1@main routing-table=vrf1

 


The second way is to explicitly specify the interface in the gateway field. The interface specified can belong to a VRF instance. Example:

# add route to 5.5.5.0/24 in the main routing table with gateway at 'ether2' VRF interface 
add dst-address=5.5.5.0/24 gateway=10.3.0.1%ether2 routing-table=main 
# add route to 5.5.5.0/24 in the main routing table with 'ptp-link-1' VRF interface as gateway 
add dst-address=5.5.5.0/24 gateway=ptp-link-1 routing-table=main

 


As can be observed, there are two variations possible - to specify gateway as *ip_address%interface* or to simply specify an *interface*. The first should be used for broadcast interfaces in most cases. The second should be used for point-to-point interfaces, and also for broadcast interfaces, if the route is a connected route in some VRF. For example, if you have an address `1.2.3.4/24` on interface *ether2* that is put in a VRF, there will be a connected route to `1.2.3.0/24` in that VRF's routing table. It is acceptable to add a static route `1.2.3.0/24` in a different routing table with an interface-only gateway, even though *ether2* is a broadcast interface:

 

add dst-address=1.2.3.0/24 gateway=ether2 routing-table=main

 ## Static VRF-Lite Connected route leaking

Sometimes it is necessary to access directly connected resources from another vrf. In our example setup we have two connected networks each in its own VRF. And we want to allow client1 to be able to access client2.

```
                   +-----------------+
                   |+-vrf1-+ +-vrf2-+|
client1(*.2)-------||ip *.1| |ip *.1||-------client2(*.2)
   (10.11.0.0/24)  |+------+ +------+|   (10.12.0.0/24)
                   +-----------------+
```
 


/ip address
add address=10.11.0.1/24 interface=sfp-sfpplus1
add address=10.12.0.1/24 interface=sfp-sfpplus2
# add VRF configuration
/ip vrf
add name=vrfTest1 interface=sfp-sfpplus1 place-before 0
add name=vrfTest2 interface=sfp-sfpplus2 place-before 0

 We can say that connected network is reachable on specific vrf by setting gateway "interface@vrf"

# add vrf routes
/ip route
add dst-address=10.11.0.0/24 gateway="sfp-sfpplus1@vrfTest1" routing-table=vrfTest2
add dst-address=10.12.0.0/24 gateway="sfp-sfpplus2@vrfTest2" routing-table=vrfTest1

 Verify routes and reachability:

```
[admin@CCR2004_2XS] /ip/route> print detail 
Flags: D - dynamic; X - disabled, I - inactive, A - active; 
c - connect, s - static, r - rip, b - bgp, o - ospf, i - is-is, d - dhcp, v - vpn, m - modem, y - bgp-mpls-vpn; H - hw-offloaded; + - ecmp 
   DAc   dst-address=111.11.0.0/24 routing-table=vrfTest1 gateway=sfp-sfpplus1@vrfTest1 immediate-gw=sfp-sfpplus1 distance=0 scope=10 suppress-hw-offload=no 
         local-address=111.11.0.1%sfp-sfpplus1@vrfTest1 
 1  As   dst-address=111.12.0.0/24 routing-table=vrfTest1 pref-src="" gateway=vrfTest2 immediate-gw=vrfTest2 distance=1 scope=30 target-scope=10 
         suppress-hw-offload=no 
 2  As   dst-address=111.11.0.0/24 routing-table=vrfTest2 pref-src="" gateway=vrfTest1 immediate-gw=vrfTest1 distance=1 scope=30 target-scope=10 
         suppress-hw-offload=no 
   DAc   dst-address=111.12.0.0/24 routing-table=vrfTest2 gateway=sfp-sfpplus2@vrfTest2 immediate-gw=sfp-sfpplus2 distance=0 scope=10 suppress-hw-offload=no 
         local-address=111.12.0.1%sfp-sfpplus2@vrfTest2 
```
 


```
[admin@cl2] > /ping 111.11.0.2 src-address=111.12.0.2
  SEQ HOST                                     SIZE TTL TIME       STATUS                                                                                         
    0 111.11.0.2                                 56  64 67us      
    1 111.11.0.2                                 56  64 61us      
    sent=2 received=2 packet-loss=0% min-rtt=61us avg-rtt=64u
```
 


But now what if we want to access routers local address located in another vrf?

```
[admin@cl2] > /ping 111.11.0.1 src-address=111.12.0.2
  SEQ HOST                                     SIZE TTL TIME       STATUS                                                                                         
    0 111.11.0.1                                                   timeout                                                                                        
    1 111.11.0.1                                                   timeout                                                                                        
    sent=2 received=0 packet-loss=100% 
```
 Approach with "interface@vrf" gateways works only when router is forwarding packets. To access local vrf addresses we need to route to the vrf interface.

# add vrf routes
/ip route
add dst-address=10.11.0.0/24 gateway=vrfTest1@vrfTest1 routing-table=vrfTest2
add dst-address=10.12.0.0/24 gateway=vrfTest2@vrfTest2 routing-table=vrfTest1

 


```
[admin@cl2] > /ping 111.11.0.1 src-address=111.12.0.2
  SEQ HOST                                     SIZE TTL TIME       STATUS                                                                                         
    0 111.11.0.1                                 56  64 67us      
    1 111.11.0.1                                 56  64 61us      
    sent=2 received=2 packet-loss=0% min-rtt=61us avg-rtt=64u
```
 


## Dynamic Vrf-Lite route leaking

With large enough setups static route leaking is not sufficient. Let's consider we have the same setup as in static route leaking example plus ipv6 addresses, just for demonstration.

/ip address
add address=10.11.0.1/24 interface=sfp-sfpplus1
add address=10.12.0.1/24 interface=sfp-sfpplus2
# add VRF configuration
/ip vrf
add name=vrfTest1 interface=sfp-sfpplus1 place-before 0
add name=vrfTest2 interface=sfp-sfpplus2 place-before 0
/ipv6 address
add address=2001:1::1 advertise=no interface=sfp-sfpplus1
add address=2001:2::1 advertise=no interface=sfp-sfpplus2

 We can use BGP VPN to leak local routes without actually establishing BGP session.

```
/routing bgp vpn
add export.redistribute=connected .route-targets=1:1 import.route-targets=1:2 label-allocation-policy=per-vrf name=bgp-mpls-vpn-1 \
    route-distinguisher=1.2.3.4:1 vrf=vrfTest1
add export.redistribute=connected .route-targets=1:2 import.route-targets=1:1 label-allocation-policy=per-vrf name=bgp-mpls-vpn-2 \
    route-distinguisher=1.2.3.4:1 vrf=vrfTest2
```
 


Now we can see that connected routes between VRFs are exchanged

