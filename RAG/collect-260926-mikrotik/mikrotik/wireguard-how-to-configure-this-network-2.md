---
id: collect-260926-mikrotik/mikrotik/wireguard-how-to-configure-this-network-2
title: "wireguard-how-to-configure-this-network"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/wireguard-how-to-configure-this-network.md
source_anchor: ""
source_lines: [189, 323]
sha256: 27de5cb954bc036f2e036a7074aea3e70a0fac0fa5c99471e06504d7bbd0dc74
---

# wireguard-how-to-configure-this-network

*Q1 - some users from R0 need internet via VPS or entire subnet? if some, how many?*

*Q2 - some users from R1 need internet via VPS or entire subnet? if some, how many?*

Entire subnet 172.20.2x.0/24 need internet via VPS (with address-lists and routing marks)

*Q3. For R2 router, is there only one group of users needing special routing or two?*

Also the entire subnet, but priority is to have access through P0, and in second place through the VPS (just distance=1 and distance=2: VPS as backup route)

*Q4. If it is only one group is it some users or the entire subnet?*

Entire subnet as above

*Q5. If one, what do you mean priority 0 to R0 and priority 1 to VPS, ie, primary and secondary in case R0 is not linked to VPS for some reason and not reachable???*

Yes, if R0 not connected to the internet, R2 users will user VPS-route.

             
            
           
          
            
            
              Good to know, we can probably dispense with any mangling!!

Any issues with pinging is at your VPS see below!

(1) As indicated you need to ensure you have an equivalent rule on VPS

*add action=accept chain=forward in-interface=wireguard-interface  out-interface=wireguard-interface*

(2) Route to all subnets on VPS

(3) Route for wireguard IPs if the VPS does provide automagic routes like MT.

MT–>  **dst-address=10.66.66.0/24  gateway=wireguard-interface table=main**

             
            
           
          
            
            
              **R0** -  Three items ( table, route, routing rule)

***/routing table add fib name=useVPS***

*/ip route
add dst–address=0.0.0.0./0 gateway=RO-WAN table=main
**add dst-address=0.0.0.0/0 gateway=wireguard-interface table=useVPS***

*/routing rule add action=lookup src-address=172.16.20.0/24  table=useVPS*

Note1:  If you didnt want the subnet to fall back to main table and find RO WAN, change action to  *action=lookup-only-in-table*

Note2:  If there are other subnets on R0 that the current subnet needs access to prior to being forced out WG for internet, then you need

to add another routing rule prior to the WG one, that is very similar except it states ***dst-address=othersubnet  table=main***

**SAME ON R1**

*/ip route
add dst–address=0.0.0.0./0 gateway=R1-WAN table=main
**add dst-address=0.0.0.0/0 gateway=wireguard-interface table=useVPS***

*/routing rule add action=lookup src-address=172.16.21.0/24  table=useVPS*

Notes1&2 also apply.

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

**R2 - A different Beast**

In this case you want all users to use R2 WAN as per normal.

You dont want to force users to use VPS for WAN, so its more of acting as a secondary WAN.

Hence, I believe this should do it.

*/ip route
add distance=5  dst-address=0.0.0.0/0  gateway=R2 WAN  table=main check-gateway=ping.
add distance=10 dst-address=0.0.0.0/0 gateway=wireguard interface table=main*

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++

No mangling, marking required on MT routers.

             
            
           
          
            
            
              Excuse me, I think you misunderstood me. All routers access the Internet through their regular WAN. But: everyone has an ADDITIONAL way to get Internet through the VPS. So it seems to me that this rule is not quite suitable for me, if I understand ROS correctly.

About the mangles, I meant this option:

```
/ip firewall mangle
add action=mark-routing chain=prerouting dst-address-list=to-useVPS-list in-interface=bridge-local new-routing-mark=useVPS passthrough=yes
```

             
            
           
          
            
            
              Can you clarify  Router2.   It seems  you want it to be able to go out internet via 3 locations, local, vps and Router 0.

Do you mean different subnets on Router2 or the same single subnet?

If the latter this will not be possible I dont think.

If router2 requests internet, its first peer to peer link will be to VPS.

VPS has no way to distinguish an incoming internet request ( should it go out VPS or router 0 )

You would need a separate WG interface on VPS to  handle that requirement to avoid overlap issues.

             
            
           
          
            
            
              Router 2 should have three routes (main and two “marked”):

1. Main WAN 0.0.0.0/0
2. Through VPS, with “vpn-mark” routing mark and distance 2
3. Through router 0 with “vpn-mark” routing mark and distance  1
It is enough that I can manually select which of the “marked” routes the “marked” traffic will go through.

Is it possible to use one WG interface in such a configuration, or is it better to make 2 WG interfaces: for connecting to the VPS and for connecting to router 0?
