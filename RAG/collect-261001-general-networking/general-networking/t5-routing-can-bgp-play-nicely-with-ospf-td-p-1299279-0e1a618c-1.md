---
id: collect-261001-general-networking/general-networking/t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c-1
title: "t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c.md
source_anchor: ""
source_lines: [1, 16]
sha256: c0578757313fca18a31c20fd403a264e2b6d71ba7738ecac523c39e3b9ae930a
---

# t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 05:23 AM - edited 03-04-2019 05:54 AM
We have a somewhat diverse network. On the "legacy" side (consisting of a dark and leased fiber "MAN", PPP circuits, and a gig fiber based LAN) we have OSPF. On the "New" side we have MPLS as a replacement for our Frame Relay infrastructure. Our telco only supports BGP as a routing protocol so we have BGP across the MPLS. The problem we're having is maintaining a consistent routing table across the BGP/OSPF boundaries. We do not redistribute OSPF into BGP. We control BGP announcements with network statements and prepending. We do however redistribute BGP into OSPF. The tricky part is that we have 4 "shared" entry points into the BGP and OSPF networks. We have 4 DS3 connections coming from the MPLS cloud going to 4 of our major data centers. We also have 4 leased fiber connections going to each of those major data centers. The leased fiber is a rung so the fiber entry point to the data centers is one VLAN. We use 3845s for the DS3 connections. The WAN side has BGP and the LAN side has OSPF. The BGP side uses one AS # for the whole shebang. We've not been able to come up with a consistently stable way to get the BGP routes into OSPF and the OSPF routes into BGP. What typically happens is what I've started calling the "Self Serving Routing Loop". Router A gets a route through OSPF and in turn, via a network statement, injects that route ingo BGP. The MPLS network, via BGP, then sends that route to Router B which dutifully picks it up and drops it into OSPF. OSPF then updates Router A which starts the process all over again. Since Router A thinks Router B is destination for the route through BGP and Router B thinks Router A is a good destination through OSPF, we wind up with a routing loop.
My question is, what is the preferred way of keeping a consistent routing table between OSPF and BGP when there are 4 shared entry points into each network? Using eBGP for the MPLS and iBGP for the LAN/MAN isn't an option due to the costs of upgrading 80+ devices to Advanced Enterprise. I apologize if this my explanation is confusing.
Solved! Go to Solution.
- Labels:
- 
						
							
		
