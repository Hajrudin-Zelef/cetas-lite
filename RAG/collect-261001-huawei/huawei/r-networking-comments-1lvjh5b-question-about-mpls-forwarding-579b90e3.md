---
id: collect-261001-huawei/huawei/r-networking-comments-1lvjh5b-question-about-mpls-forwarding-579b90e3
title: "r-networking-comments-1lvjh5b-question-about-mpls-forwarding-579b90e3"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-huawei/r-networking-comments-1lvjh5b-question-about-mpls-forwarding-579b90e3.md
source_anchor: ""
source_lines: [1, 38]
sha256: 363523b1d9e5b3a3a827248e9d05dd84402a0b89183c77dd0a99e963b0c7e60d
---

# r-networking-comments-1lvjh5b-question-about-mpls-forwarding-579b90e3

Question about MPLS forwarding 
        
        
        
    
    
    Here is the scenario:
CE-A1 --- 1.1.1.1(PE) --- 2.2.2.2(P) --- 3.3.3.3(P) --- 4.4.4.4(PE) --- CE-A2
The providers routers have OSPF and MPLS LDP converged between them, the PE's have eBGP sessions with its connected CE and the PE's have iBGP sessions between themselves.
I want to make the P routers forward packets purely with MPLS
1.1.1.1(PE) has a route to 203.117.8.0 that CE-A2 send to 4.4.4.4(PE) and 4.4.4.4(PE) is advertising it to 1.1.1.1(PE) via iBGP with next-hop-self
1.1.1.1(PE) has this entry in its bgp table:
Network NextHop MED LocPrf PrefVal Path/Ogn
*>i 203.117.8.0/23 4.4.4.4 0 100 0 65001?
1.1.1.1(PE) has this entry in its LSP table:
FEC In/Out Label In/Out IF
4.4.4.4/321028/1028 -/GE0/0/0
The problem is that when CE-A1 tries to ping 203.117.8.1 the 1.1.1.1(PE) forwards the packet to 2.2.2.2(P) but it send the packet with no label, and because 2.2.2.2(P) doesn't participate in BGP it doesn't know how to reach 203.117.8.0/23 and has to drop the packet. But 1.1.1.1(PE) knows that 203.117.8.0/23 next hop is 4.4.4.4, and there is a FEC to 4.4.4.4 in the LSP table, so how do i make 1.1.1.1(PE) add the label to packets whose next hop is 4.4.4.4(PE) when sending them to 2.2.2.2(P) ?
I'm using huawei but i'm not asking for specific configuration commands, just what to do and the name of the functionality that i'm looking for would be nice
Section des commentaires
You need the customers in a vrf and attached to bgp. Then bgp vpnv4 neighborship with the two PES.
Then the bgp will work with MPLS and create VPN labels. The P routers will use two labels to route the traffic. The top label will direct it from p to pe and vice versa. The PE will pop the second label and forward to customer.
The router is using pure IP because you haven't triggered MPLS service. If you uses vpnv4 it will automatically route with bgp-mpls-vpn.
also... he can make Layer 2 point-to-point tunnels or point-to-multipoint VPLS using the MPLS encapsulation...
but if customers are connected to CE it can get kinda "messy" setup... VRFs would be a better solution (as you suggested...)
if you want packets to be forwarded to a target via mpls, then the next hop needs to be via a route that has a label attached. You've only mentioned the bgp next hop. I suspect you're missing some steps here in the routing layer.
Can you post your configuration?
You have to enable mpls ldp on P router and to have ldp sessions with PE routers. P do not have to know about bgp prefix but have to have ipv4 routes for PE loopbacks ( ospf or isis ) and then alocate labels for them. After that it is pure mpls switching , no ip lookup on P router.
BR
Your next hop for PE4 should not be via ge, it should be via tun interface from my memory on huawei. P2/p3 should only see labeled traffic between p1/p4 and not the inner.
sorry for intrusion... but MPLS is still a thing in 2025?
with VPN IPSec and or SD WAN ?
why ?
SD-WAN, IPsec and other tunneling are technologies to use a network.
MPLS is a technology to build a network.
ok, I mean, why I have to use MPLS when I have Internet everywhere and I can use IPSEC ?
According to your description I would have to assume that you are not using VRFs. The proper way to have this working is to use VRFs (L3VPNs).
If your scenario does not allow L3VPNs then the other way for this to work is to populate your FEC table with all your prefixes. The chinese vendor devices only add FECs for Loopbacks by default, so you will have to add the command to populate using all prefixes.
