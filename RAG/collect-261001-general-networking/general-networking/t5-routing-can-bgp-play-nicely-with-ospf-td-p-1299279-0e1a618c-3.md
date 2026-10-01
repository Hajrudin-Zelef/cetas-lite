---
id: collect-261001-general-networking/general-networking/t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c-3
title: "t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c.md
source_anchor: ""
source_lines: [159, 197]
sha256: 64bc1c64cbdf3feb9a0137c8ed31197fd54ecc7d6fa3e5201d9ab371b1665b19
---

# t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c

I know when i was doing something similiar ie. multiple entry points with BGP and EIGRP I ended up modifying the weight in BGP but that was to make sure BGP was used over EIGRP so it may not be applicable. Like i say, an example would be helpful.
Having said that it may also be worth waiting for Edison to come back as i may well have misunderstood what he is proposing and i've learnt it's never a good idea to underestimate Edison !!
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 08:47 AM
I've just about got the picture ready to go. I think when you see that, you'll have a better understanding of what I mean by the self serving routing loop. Thanks!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 09:25 AM
OK, here's the diagram.
We started out with just the leased fiber connections on the back side of locations A thru D. Unfortunately, our leased fiber provider has a history of outages so management asked that we provide a good sized backup pipe. That's where the MPLS connections came into play. They're primarily there to provide connectivity to the MPLS connected sites (remote locations). However, if the leased fiber from Location A should go down, we need all traffic from Locations B thru D to fail over to the MPLS cloud to get to location A. The same is true for the other major locations, the MPLS connection needs to be a backup for the leased fiber. Note that the leased fiber and the DS3 connections do not terminate in the same device. The L3 switches are a combination of 6509s, 3750s, and a 4948. There are multiple subnets at each of the major locations A thru D.
Also, there's a multi-point GRE VPN tunnel from each of the remote locations to a 2811 at one major location. We use the floating static route method to handle traffic to and from the remote locations if their primary connection ever goes down. The remote location has a default route pointing down the tunnel. The VPN router has a set of static routes, pointing back to the remote location's local subnets with the IP address of the remote router's tunnel as the next hop. These static routes have a metric of 111 so that OSPF is preferred. If the route disapears from OSPF the VPN router injects it's static route. This part has worked flawlessly over the years.
Please let me know if you have any other questions about what we've got and/or what we're trying to do. Thanks.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 09:55 AM
Jon,
Picture a remote location with 172.16.1.0/24 connected to the MPLS Cloud. There are multiple BGP routers at the DC connected to the MPLS.
Router A picks the 172.16.1.0/24 and redistribute this route into OSPF.
Router B has a network statement for this route (I'm assuming based on the route-loop Terry has experienced) and advertises this route back to BGP.
Router A may use this route as Best Path based on the BGP attributes hence causing the loop.
We need to eliminate Router B from advertising 172.16.1.0/24 back into BGP and the most scalable way of avoiding this would be with redistribution with route-tagging.
I guess you haven't come across those redistribution nightmares from the INE labs on your CCIE studies yet?
