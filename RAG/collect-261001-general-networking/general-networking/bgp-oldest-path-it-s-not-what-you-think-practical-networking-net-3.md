---
id: collect-261001-general-networking/general-networking/bgp-oldest-path-it-s-not-what-you-think-practical-networking-net-3
title: "bgp-oldest-path-it-s-not-what-you-think-practical-networking-net"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/bgp-oldest-path-it-s-not-what-you-think-practical-networking-net.md
source_anchor: ""
source_lines: [408, 477]
sha256: 9362ca53554c03d76698f39db31f6879e12a736b31f9704325259819757c3445
---

# bgp-oldest-path-it-s-not-what-you-think-practical-networking-net

Consider what could happen if Step 10 did not exist, and instead only Step 11 (prefer the path with the lowest router-id) was breaking ties (presuming steps 1-9 were equal). The Router-ID of the next eBGP peer you build an adjacency with cannot be known, therefore the Router-ID is essentially random.

If at some point in the future you buildÂ a redundant ISP connection with another eBGP peer which happens to have a better (lower) Router-ID than your current ISPâs eBGP peers, all your traffic would shift off of the current âknown goodâ path, to the new âuntestedâ path.

Moreover, if you happened to have a neighbor adjacency start to flap, and that neighbor happened to have the best Router-ID, then you would continually have traffic âflapâ to the peer when it is up, and âflapâ away when it goes down. This could cause all sorts of instability if all of your egress traffic kept changing its path whenever an eBGP peer was having issues.


In the end, Step 10 does not necessarily prefer the *oldest* path, so much as it prefers a *current* path if a new path isÂ learned.

Of course, if the new path is actually desired, you can force traffic along that path by modifying some of the attributes in Step 1-9, but if all these attributes are identical then both paths will have an identical preference. BGP will therefore prefer the current, stable, âknown goodâ path instead of the newly learned, identically preferred path.

Step 10 assures that a new, identically preferred, path will not supersede a path that is already known. Step 10 does not prefer the oldest path as an absolute age, it simply prefers theÂ current path if a new one is learned.


Have you ever wondered why there is no Cisco command that can tell you the absolute age that a particular path to a prefix was learned (not a neighbor adjacency, but the specific age of a specific path advertised by a neighbor)? It is because that absolute age is not used anywhere, and therefore not tracked.

When a new path is learned and nothing in Steps 1-9 causes the new path to be more preferred, Step 10 assures the current best path remains marked as the best path.

When a best path is lost and BGP runs through the path selection process to elect a successor, all *remaining* paths to the prefix *already* exist in the BGP table. Therefore, they all have the *same* age and Step 10 cannot break the tie. Instead, Step 11 breaks the tie.


### Conclusion

BGPâs oldest path is often stated as âprefer the oldest pathâ, but as we’ve demonstrated, this isnât entirely accurate. A more accurate way of stating it would be â*when learning of a new path*, prefer the current (stable) path over the newly learned path (if nothing in step 1-9 explicitly implies the new path is more desirable).â

Awesome explanation in a very different way. Thank-you so much

If you could get some time to post for IGMP, Multicast it would be great help

Great job with the detailed explanation!

Possibly another way to present this behavior would be that BGP process does not maintain historical state information on the peers it already knows about. It is a stateless comparison between the ones already known and the one(s) newly learnt. This is being specific to the step 10 (Oldest peer). Hence it will always ends up preferring the “already” known and if there are more than one in the set known then it proceeds to the next step to select one form them.

Please post about the OSPF also.

This is such a brilliant article. Thank you!

Very good, thank you.

Excellent article.

is there a way to force a preferred path by using “clear IP bgp …” to prefer a different path if there is a tie break?

Brilliant article! Thanks very much for the efforts. Very well-explained.

can not see the BGP, OSPF, EIGRP videos

Hi Ed! How does Trial 2 differ from Trial 1? When you shut down the 9.22.11.2 peer, both remaining paths were also present in the topology table, so following this logic – there should be a tie, with step 11 as a tie-breaker.

That is exactly what happened. When the link to R2 was shut down, the remaining paths through R3 and R4 were tied since they were both present in the BGP table when R1 lost the path through R2. R1 used Step 11 (lowest Router-ID) to break the tie to select R3 as the next best path.

Often people consider that R3 was selected because the R3 path was learned before R4, meaning it had an older absolute age. But Trial 2 proves that wasn’t the case. I elaborate on it further in the “What Happened” section.

Hi Ed !

Can you publish a series on OSPF topics.

OSPF is on my list =)

If OSPF is on your list, Forwarding Address would be a great topic for discussion. Great Article on BGP.

Noted, thanks Dinesh =)

Hello Ed,

Can you publish BGP video series in Youtube.

Ed – this was excellent! Thanx

Hey Ed. How about make a video of such a beautiful content. I think this will be excellent addition for the networking fundamentals course
