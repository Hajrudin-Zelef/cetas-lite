---
id: collect-261001-huawei/huawei/r-networking-comments-a7n1ao-huawei-ar-router-with-mpbgp-and-routereflectors-0ca486c4
title: "r-networking-comments-a7n1ao-huawei-ar-router-with-mpbgp-and-routereflectors-0ca486c4"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/r-networking-comments-a7n1ao-huawei-ar-router-with-mpbgp-and-routereflectors-0ca486c4.md
source_anchor: ""
source_lines: [1, 23]
sha256: a29e982c5f7ec18c4e1cbfb4c3f02cb037d4434c255802ac85f9b17dd5c6e451
---

# r-networking-comments-a7n1ao-huawei-ar-router-with-mpbgp-and-routereflectors-0ca486c4

Huawei AR Router with MP-BGP and RouteReflectors config question 
        
    Hi, we're testing some Huawei devices and I'm struggling with an MPLS test. We have a OSPF core, BGP free, thats working fine. Two dedicated RRs and, for now, two PEs. I created a test VRFs, or as Huawei calls them VPN instances. Now I just want to route between those PEs, right now just with a Loopback interface in the VRF.
First I tried it without the RRs, everything worked fine. Both routers got the routes from the other router and I could ping between the loopbacks in the VRFs. Now, I tried to implement the RRs and I don't get any routes in the VRFs.. or in the vpnv4 BGP. I also don't see BGP peers in the VRF. But I have the neighborship in the "dis bgp vpnv4 all peer" command. And in the global routing-table of course.
For testing, I implemented the RR config for the "ipv4-family unicast" BGP address-family and it is working. I imported the directly connected routes from the PEs into BGP and I get them on the other PE.
Right now, I don't understand what is missing in the VPN-Instance or BGP config.
I'll post comments with the config.
EDIT:
It was the "policy vpn-target" under the ipv4-family vpnv4 BGP address-family. I just deleted that in the lab and it is working fine.
Section des commentaires
Did you try removing policy vpn-target line on rr? Because as I recall, it has similar functions as rt filter has. So maybe rr ignores the received routes.
I‘ll try that tomorrow. Actually I didn‘t even look at it. Thanks for opening my eyes.
You're right. The policy blocked it, as soon as you undo the policy it is working.
[AR9]tracert -a 9.9.9.99 -vpn-instance ROSA -v 9.9.9.10
traceroute to ROSA 9.9.9.10(9.9.9.10), max hops: 30,packet length: 40,press CTRL_C to break
1 1.1.79.7[MPLS Label=1029/1036 Exp=0/0 S=0/1 TTL=1/1] 30 ms 30 ms 30 ms
2 1.1.78.8[MPLS Label=1029/1036 Exp=0/0 S=0/1 TTL=1/2] 40 ms 30 ms 30 ms
3 9.9.9.10 30 ms 30 ms 30 ms
[AR9]
I had the same issue with hp comware. It was like you said. :)
Config RR1, IP: 1.1.1.1
Config of the PE9, IP: 9.9.9.9 (PE10 is identical)
I would make a ticket with Cisco or Extreme TAC. The stolen code has probably been patched by the vendors who made it.
