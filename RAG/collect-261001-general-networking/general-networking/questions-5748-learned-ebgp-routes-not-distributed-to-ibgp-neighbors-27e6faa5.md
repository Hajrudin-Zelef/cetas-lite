---
id: collect-261001-general-networking/general-networking/questions-5748-learned-ebgp-routes-not-distributed-to-ibgp-neighbors-27e6faa5
title: "questions-5748-learned-ebgp-routes-not-distributed-to-ibgp-neighbors-27e6faa5"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-5748-learned-ebgp-routes-not-distributed-to-ibgp-neighbors-27e6faa5.md
source_anchor: ""
source_lines: [1, 62]
sha256: 0601de91d3272543bc6b2fc1cc096a7b16256f790e25b8cfc8763c7771066886
---

# questions-5748-learned-ebgp-routes-not-distributed-to-ibgp-neighbors-27e6faa5

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
9
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have (in GNS3) three Cisco 3640s running 12.4(23) connected in series (R1 -> R2 -> R3).
R1 and R2 are eBGP peers, R2 and R3 are iBGP peers.
R1 advertises network 192.168.1.0/24. R2 receives this route, but does not advertise it to R3.
Shouldn't routes learned from eBGP be advertised to iBGP?
Here's the full configuration as entered:
loop0(R1)s0/0 <--> s0/0(R2)s0/1 <--> s0/1(R3)
R1:
configure terminal
interface s0/0
ip address 172.16.1.1 255.255.255.252
no shutdown
interface loopback0
ip address 192.168.1.1 255.255.255.0
router bgp 1
neighbor 172.16.1.2 remote-as 2
network 192.168.1.0 mask 255.255.255.0
R2:
configure terminal
interface s0/0
ip address 172.16.1.2 255.255.255.252
no shutdown
interface s0/1
ip address 172.16.1.5 255.255.255.252
no shutdown
router bgp 2
neighbor 172.16.1.1 remote-as 1
neighbor 172.16.1.6 remote-as 2
R3:
configure terminal
interface s0/1
ip address 172.16.1.6 255.255.255.252
no shutdown
router bgp 2
neighbor 172.16.1.5 remote-as 2
Assuming you have no IGP configured (such as EIGRP / OSPF / ISIS / RIP), then the simplest explanation is that R3 hasn't got a route to 192.168.1.0/24's next-hop when the iBGP update arrives at R3.
loop0(R1)s0/0 <-----------> s0/0(R2)s0/1 <-------------> s0/1(R3)
AS 1 AS 2 AS 2
---------> ----------->
Prefix: 192.168.1.0/24 Prefix: 192.168.1.0/24
AS-path: 1 AS-path: 1
Next-hop: 172.16.1.1 Next-hop: 172.16.1.1
*via eBGP* *via iBGP*
Since iBGP doesn't reset the next-hop when it receives the update from R1, 192.168.1.0/24's next-hop (172.16.1.1) has to be reachable (see Why routers ignore BGP paths for more details).
The most basic way to test this is to configure a static on R3:
ip route 172.16.1.0 0.0.0.3 172.16.1.5 name BAD_HACK_FOR_IBGP
That's obviously the wrong solution, but it's a dead-simple test to illustrate why things are broken (keep in mind that you may need to wait a bit for the BGP next-hop scanner to run before 192.168.1.0/24's route is installed).
There are two possible solutions that come to mind, but only one really makes sense in most networks...
Best solution: Configure an IGP... pick any IGP you like, and advertise 172.16.1.0/30 throughout AS 2 in that IGP
Optional solution: Configure the peering session between R2 and R3 to set next-hop-self
Next-hop reachability is one of the most basic issues when understanding BGP; almost everyone runs into this issue when they're experimenting with the protocol.
You should have to anounce your connected interfaces on R2. because the only network 192.168.1.0 does not now what is the next hop.
You can checking with "show ip bgp" on R3.
As you can see 192.168.1.0 in R3 bgp table but not inserted to routing table. Because it does not know what is the next-hop
Solution:
Redistribute connected interfaces on R2
Say EBGP neighbor that who is next hop as "neighbor 172.16.1.2 next-hop-self" on R1.
