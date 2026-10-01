---
id: collect-261001-general-networking/general-networking/questions-16351-ospf-backbone-area-and-the-wan-ae71f4ca
title: "questions-16351-ospf-backbone-area-and-the-wan-ae71f4ca"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-16351-ospf-backbone-area-and-the-wan-ae71f4ca.md
source_anchor: ""
source_lines: [1, 22]
sha256: 4dbad6f1160a2799202d6d9406da0c857f941b6674ce5faa15e10fa04ea3d9c1
---

# questions-16351-ospf-backbone-area-and-the-wan-ae71f4ca

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
8
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am working on a new network deployment, and have run into an interesting question... Should the OSPF area 0 be extended across a WAN's point to point links (thereby creating more border routers), or should one's WAN Aggregation router serve as the ABR for each area? Consider the following:
In Scenario 1, R2, assuming area 1 is a totally stubby area receives only a default route, thereby reducing WAN bandwidth... I think? In Scenario 2, R2 receives many more routes. This seems like a trivial question, but I feel it may have a performance impact at scale.
What is the best practice for extending area 0 across a WAN?
There is no one "best practice," but rather several "good practices" and a few "not-so-good practices."
Generally speaking, if you have multiple areas, you want the hub and spokes in the same area, and summarize between the hub and the rest of the LAN network. You can make the whole hub and spoke area totally stubby.
Another possibility is to break up the spokes into a few (no more than 2-3) areas, with the hub as the ABR. The most important factor that would make this design preferable is the relative instability of your WAN and remote LAN links. If you have a lot of links going up and down, this will affect the amount of flooding over the WAN and therefore the link utilization. But for 50 routes or so, that shouldn't be a problem for a T1.
I prefer scenario 2, extend area 0 into the branch, and then stub area into branch to l3 switch for vlan routing.
Making the branch router ABR.
The reason is that we need high availability for the branches, with a couple of redundant links from different providers, as well as 3G, and dial-up backup.
The bigger branches also have dual routers to cater for a router failure, here we do provision a area 0 link between the routers, using a vlan.
Allowing Area 0 up to the edge of the branch allows for better routing decisions.
We are also able to deploy links between branches that are close together, using the area 0 interface on each branch router, this would not be possible if each branch was in a different area.
Dialup / DMVPN etc also causes problems if the branches routers are only in a area >1 as the central router now needs to have the template/mutipoint-gre in more than one area.
Summary:
I think the WAN should be in area 0, if the wan is not extremely simple/standard with just one or two links to each site.
Scenario 1 would be better in this case since WAN bandwidth will be reduced by a bit and and R2 would be spared of spending its resources if it would be an ABR (scenario 2).
