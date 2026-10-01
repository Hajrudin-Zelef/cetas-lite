---
id: collect-261001-general-networking/general-networking/t5-routing-ospf-could-not-establish-adjacency-manually-in-frame-relay-td-p-15793-2934d03d
title: "t5-routing-ospf-could-not-establish-adjacency-manually-in-frame-relay-td-p-15793-2934d03d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-ospf-could-not-establish-adjacency-manually-in-frame-relay-td-p-15793-2934d03d.md
source_anchor: ""
source_lines: [1, 78]
sha256: 5a264648bfdb69b684862f54c874043dc3368bc1b2b839351b8c596b53974245
---

# t5-routing-ospf-could-not-establish-adjacency-manually-in-frame-relay-td-p-15793-2934d03d

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-03-2010 11:29 AM - edited 03-04-2019 09:58 AM
Hi
I have problem establishing adjacency with R2 se0/1.1 and R3 se0/0.1 from R1 se0/0.1
Attached are the relevant configuration and a partial network diagram.
I am trying to manually configure neighbour relationship using unicast (neighbor command) but stuck at attempt state until the dead timer expires...Please give advice on what's missing in my configuration on the 3 routers (the zipped also contain FRS1 configuration for further reference)
Thanks.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Routing Protocols
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-04-2010 08:34 AM
OK = firstly let me state an OSPF fact. The OSPF neighbor adjacencies have a TTL of 1.
So with that in mind trying to establish a neighbor between R2 and R3 going thru R1 is never going to happen on Point-to-Point/Broadcast and Point-to-Multipoint networks.
Now to be honest in your topology this is the worst thing you would want to do, as you may force a DR/BDR election, and R2/R3 will win (.2 & .3 are higher that .1). Now since the DR controls the database - it's highly unlikely R3 will ever get any routes. So in your topology R1 is the best device to become DR. With that said if you are really intent on making R2 & R3 neighbors - the only way to get over the TTL 1 is configure a tunnel between R2 & R3 traversing R1. Since the tunnel will originate on R2 and end R3 = 1 hop TTL 1 = solved.
HTH>
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-04-2010 08:34 AM
OK = firstly let me state an OSPF fact. The OSPF neighbor adjacencies have a TTL of 1.
So with that in mind trying to establish a neighbor between R2 and R3 going thru R1 is never going to happen on Point-to-Point/Broadcast and Point-to-Multipoint networks.
Now to be honest in your topology this is the worst thing you would want to do, as you may force a DR/BDR election, and R2/R3 will win (.2 & .3 are higher that .1). Now since the DR controls the database - it's highly unlikely R3 will ever get any routes. So in your topology R1 is the best device to become DR. With that said if you are really intent on making R2 & R3 neighbors - the only way to get over the TTL 1 is configure a tunnel between R2 & R3 traversing R1. Since the tunnel will originate on R2 and end R3 = 1 hop TTL 1 = solved.
HTH>
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-04-2010 09:07 AM
Hello Andrew.
I like your style - short, candid and sweet
Btw, I solved the problem, virtual routers never beat the real hardware I tried configuring with my routers I got no problems.
R1#sh ip ospf neighbor
Neighbor ID     Pri   State           Dead Time   Address         Interface
3.3.3.3           0   FULL/DROTHER    00:01:58    172.30.0.3      Serial0/0.1
2.2.2.2           0   FULL/DROTHER    00:01:52    172.30.0.2      Serial0/0.1
R2#sh ip ospf neigh
Neighbor ID     Pri   State           Dead Time   Address         Interface
1.1.1.1           1   FULL/DR         00:01:47    172.30.0.1      Serial0/1.1
R3#sh ip ospf neigh
Neighbor ID     Pri   State           Dead Time   Address         Interface
1.1.1.1           1   FULL/DR         00:01:52    172.30.0.1      Serial0/0.1
The topology is not for real deployment, it is for demonstration only, originally I wanted to try on simulation first due to number of routers constraint. But it seems GNS3 never works for me.
Thanks for looking into it
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-04-2010 09:12 AM
np - glad to help.
Sorry about the short answer - was dealing with a college who was having a bad "brain won't work" day, my furstration kinda rubbed off on my post to you.
