---
id: collect-260926-mikrotik/mikrotik/pcc-across-two-dhcp-wan-failover-help
title: "pcc-across-two-dhcp-wan-failover-help"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "optics"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/pcc-across-two-dhcp-wan-failover-help.md
source_anchor: ""
source_lines: [1, 140]
sha256: 7292f7f26112604969a0ebfa4bfbde40f9842ddad93c22bee3ace3a690be7bf4
---

# pcc-across-two-dhcp-wan-failover-help

Hi everyone,

I’m working on load-balancing a pair of WAN connections on my hEX S router. Unlike the Manual:PCC guide, though, both of my WAN connections have a dynamic IP address assigned with DHCP.

To get around this, I’m using a slight twist on the DHCP script example found on the wiki:

:local rmark “to_ISP1”

:local count [/ip route print count-only where comment=“to_ISP1”]

:if ($bound=1) do={

:if ($count = 0) do={

/ip route add gateway=$“gateway-address” comment=“to_ISP1” routing-mark=$rmark check-gateway=ping

} else={

:if ($count = 1) do={

:local test [/ip route find where comment=“to_ISP1”]

:if ([/ip route get $test gateway] != $“gateway-address”) do={

/ip route set $test gateway=$“gateway-address”

}

} else={

:error “Multiple routes found”

}

}

} else={

/ip route remove [find comment=“to_ISP1”]

}The other WAN port’s script replaces “to_ISP1” with “to_ISP2”.

I’m also using the router’s ‘bridge’ port instead of a ‘LAN’ port as listed in the guide.

Finally, I’ve replaced the two ‘accept’ rules with a single rule specifying a destination address type of ‘local’ instead.

/ip firewall mangle

add action=accept chain=prerouting dst-address-type=local in-interface=bridge

add action=mark-connection chain=prerouting connection-mark=no-mark 

in-interface=ISP1 new-connection-mark=ISP1_conn

add action=mark-connection chain=prerouting connection-mark=no-mark 

in-interface=ISP2 new-connection-mark=ISP2_conn

add action=mark-connection chain=prerouting connection-mark=no-mark 

dst-address-type=!local in-interface=bridge new-connection-mark=ISP1_conn 

per-connection-classifier=both-addresses:2/0

add action=mark-connection chain=prerouting connection-mark=no-mark 

dst-address-type=!local in-interface=bridge new-connection-mark=ISP2_conn 

per-connection-classifier=both-addresses:2/1

add action=mark-routing chain=prerouting connection-mark=ISP1_conn 

in-interface=bridge new-routing-mark=to_ISP1

add action=mark-routing chain=prerouting connection-mark=ISP2_conn 

in-interface=bridge new-routing-mark=to_ISP2

add action=mark-routing chain=output connection-mark=ISP1_conn 

new-routing-mark=to_ISP1

add action=mark-routing chain=output connection-mark=ISP2_conn 

new-routing-mark=to_ISP2

/ip firewall nat

add action=masquerade chain=srcnat comment=“defconf: masquerade” 

ipsec-policy=out,none out-interface-list=WAN

add action=masquerade chain=srcnat out-interface=ISP1

add action=masquerade chain=srcnat out-interface=ISP2This is working well! I can pull the combined speed of both of my WAN connections when on a multi-connection download test like fast.com. 42Mbps might not sound like a lot, but it is when you’re used to 20. 

However, the problem comes when I try to use failover routing as discussed in the PCC guide. It suggests creating a route with a cost of 1 and 2 on WAN1 and WAN2, respectively.

add dst-address=0.0.0.0/0 gateway=10.111.0.1 distance=1 check-gateway=ping

add dst-address=0.0.0.0/0 gateway=10.112.0.1 distance=2 check-gateway=pingI’ve tried a couple of different options to fill this in with the correct information. One was adding another route inside the DHCP script:

/ip route add gateway=$“gateway-address” comment=“to_ISP1” routing-mark=$rmark check-gateway=ping

/ip route add gateway=$“gateway-address” comment=“failover_ISP1” distance=1 check-gateway=pingI also tried to use the DHCP Client’s default route.

When I try either of these options, I’m basically unable to use HTTP(S). While ICMP traffic seems to pass undisturbed, any attempts to browse the web take a very long time to connect and only pass traffic at about 100Kbps. Once the routes without marks are deleted again, traffic is able to pass at 42Mbps.

But that still leaves me without failover… Does anyone know why having the so-called “Failover” routes in this configuration would cause such strange traffic patterns? I suspect that changing the “accept” rules is the most likely problem, but I’m not sure how I’d dynamically create these rules correctly.

A full /export hide-sensitive with my router’s IPs masked can be found at https://paste.ubuntu.com/p/XxXPJgYmzd/

             
            
           
          
            
            
              You can set much easier load balancing even with DHCP client ISPs - no need for scripts.

Basically, with two static ISP routes, you can do following (which is simplification of Mikrotik guide, by only marking ISP2 connections):

** in IP/Firewall/Mangle

1. mark-connection on input/ISP2 ( and maybe prerouting if you allow incoming connections) any new connection that enters via ISP2 as ‘for ISP2’
2. mark-connection on prerouting/bridge any new outgoing connection you want to loadbalance ( for example, using PCC/2:0 or NTH 2,1) as ‘for ISP2’
3. mark-routing anything with connection mark ‘for ISP2’ on prerouting/bridge and output  with routing-mark ‘to ISP2’
** in IP/Routes
4. have highest priority (say distance 10) default route 0.0.0.0 point to ISP2 if routing mark is ‘to ISP2’ ( with check:ping, to all use next ISP1 default route if ISP2 is down)
5. have second priority (say distance 20) default route 0.0.0.0 point to ISP1 ( with check:ping to all use next ISP2 default route if ISP1 is down)
6. have third priority (say distance 30) default route 0.0.0.0 point to ISP2 ( check:ping not necessary, if this one fails too, no route works anyway )

Now, if first ISP have DHCP client, you have #5 created for you dynamically, but with wrong distance and without “check:ping”, and thus without failover. To fix that, you can do:

- in IP/DHCP client, set ‘default route distance’ to high number ( say 55 ) - this would make dynamically created ISP1 route as last priority and still without check, which is not exactly what we want ( but important is that distance 55 is unique and no other route has it )
- in Routing/Filter , add new filter for “dynamic-in” chain and distance ‘55’ , and with  ‘set check gateway’=‘ping’ and ‘set distance’=‘20’ - this will “fix” above dynamic route to have failover ping check and correct priority/distance  ( and since 55 was unique, it will fix only that dynamic DHCP route )

If you have both ISPs as DHCP client, just can repeat above for ISP2 also - but you do NOT have to , since dynamic route for ISP2 will be last one where checking for failover is moot point anyway. You do need to go and change DHCP client default route distance to 30 for ISP2 though.

I believe that DHCP client dynamic routes are even better for failover than statically created routes - configuration has just one or two more commands as described above, but benefit is that ping is checking NEXT ISP router ( at ISP premises, one that gave DHCP data), and thus can detect when your cable is down. That is much better than using check:ping on static routes , where that ping will work even if your cable/optics is down between you and ISP, since you are pinging ISP router inside your home. Only way for that to fail is if you turn off their router in your home. Now, if Mikrotik would introduce option to set checkPingAddress, it would solve this issue ( if we set checkPingAddress=8.8.8.8, and Mikrotik sends ping over gateway we are checking, it would check not only cable but also ISP network as well).  But I digress… that has nothing to do with your question , and I bet similar thing was suggested ages ago and if Mikrotik did not introduce it by now, I doubt they will do it now.
