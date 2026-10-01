---
id: collect-261001-cisco/cisco/r-networking-comments-rqynzh-eigrpospf-multicast-troubleshooting-cisco-7afead04
title: "r-networking-comments-rqynzh-eigrpospf-multicast-troubleshooting-cisco-7afead04"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/r-networking-comments-rqynzh-eigrpospf-multicast-troubleshooting-cisco-7afead04.md
source_anchor: ""
source_lines: [1, 26]
sha256: 801eb72ee824ac83fc0d327c7db6718ec84673fd8c192a847922096fe146156f
---

# r-networking-comments-rqynzh-eigrpospf-multicast-troubleshooting-cisco-7afead04

EIGRP/OSPF Multicast Troubleshooting - Cisco
Hey all, I have a pretty noob question that I stumbled upon while troubleshooting an EIGRP neighbor issue. I was using the "Troubleshooting IP Routing Protocols" book and in the EIGRP section it has a flowchart for issues, and one of those steps was pinging the multicast address from both routers to ensure EIGRP packets are flowing. I couldn't ping the multicast address no matter which interface I sourced it from, which I thought was strange.
For funsies I made an EIGRP/OSPF lab and tried pinging the .5 and .6 addresses, but that wasn't getting replies either. Multicast routing is enabled, adjacencies have formed and everything is working, but still no pings to the multicast addresses.
I have a strong feeling that I'm missing something fundamental about multicast in general and I'm about to dive into a few multicast Cisco Live breakout sessions, but what's my disconnect on why I can't ping those two groups?
Flowchart for reference:
Troubleshooting EIGRP > Troubleshooting EIGRP Neighbor Relationships | Cisco Press
Section des commentaires
I just setup a quick lab between two Cisco IOS routers and with an active EIGRP adjacency I can ping the 224.0.0.10 address from either side and get a response.
R1:
R2:
R1 ping:
R2 ping:
So, as for why you can't ping between, no clue. Are the two routers directly connected or is there some device in between that may drop the packets?
This was it, but I'm not entirely sure why. I have a test MPLS switch that I threw in so that everything could connect to that, once I removed that and made direct connections from Router -> Router I was able to ping the multicast address.
When I get some free cycles I'll take a closer look at the switches config.
This won't be relevant as the groups used by EIGRP and OSPF are link-local. You should only be able to ping them from the same broadcast domain by design.
Would be interesting if you could route it though. I wonder how OSPF would work with a single DR for the whole network :D
Kind of like a BGP RR, except much more bodge-tastic.
Are you defining static neighbor relationships between routers? If you do this, EIGRP will only use unicast and disables EIGRP multicast on the selected interface.
Were you specifying the source interface? eg. "ping ... source Gi 0/1" When pinging a multicast address that's not routed you need to ensure it's going out the right interface.
Apparently this doesn't work, and "source" will only change the source IP and not the source/destination interface. So I have no idea how you're supposed to ping a link-local multicast address when the router may decide to send it out an arbitrary interface. In other words it sounds like a bad test that you should ignore.
source interface specifies the source IP in the packets, not the interface they will be sent out.
Oh wow you're right, I'm too used to linux ping which does "source interface" properly.
Commentaire supprimé par le membre
Lol. I know what EIGRP's multicast address is, I'm asking why I can't ping either routing protocol's address when I'm using both and am forming adjacencies.
I brought up OSPF to further illustrate that it wasn't just EIGRP's address that I couldn't ping.
