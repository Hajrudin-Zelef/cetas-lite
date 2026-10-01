---
id: collect-261001-general-networking/general-networking/t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c-2
title: "t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c.md
source_anchor: ""
source_lines: [17, 158]
sha256: f0e336d367060b8b8f2dfad7157aba82b87fb3467e5105bf5cf865e9005f04dc
---

# t5-routing-can-bgp-play-nicely-with-ospf-td-p-1299279-0e1a618c

			Routing Protocols
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 01:10 PM
Terry,
You got the concept, nice job.
BTW, no need to include the metric-type 2 on the redistribution, it's done by default.
OSPF into BGP does not need 'subnets' you need 'subnets' from BGP into OSPF.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 05:58 AM
If you could make a diagram and draw out the BGP/OSPF boudary points that would help. What you would want to do is deny duplicate routes when redistributing. Its fine if your not doing mutual redistribution but you also have to be carefull about what your advertising and you Administrative distance on the devices.
Typically when redistributing from B-->O you would use a route-map permitting the specific routes that you need (either ACL, tag value,etc) and just bring those into the network. If you have Dual redistribution points than you would typically deny those tagged routes from the other device that it doing redistribution. If you could draw up a diagram this would be helpful. thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 06:47 AM
Thanks. We've got a far too detailed map that I'll try to simplify and post. If I'm understanding what you're saying, I think that's where we're having a problem. We do still need the duplicate routes as the BGP network is supposed to be a backup connection in case the OSPF network connections fail at the data center. In other words, Router A still needs to somehow or another know that Router B's subnets are available via BGP just in case the leased fiber connection at Router B's location fails. We've come up with a configuration that uses a route-map to limit what can come in dynamically through BGP and a static route with a higher metric for all the subnets behind the other redistribution routers. We haven't implemented it yet as we wanted to make sure that this change will be the final change and we're not overlooking anything.
I'll see if I can get the diagram posted later this afternoon. Thanks.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 06:53 AM
Awesome...If you could also include the configurations that your planning on using I could help you out with this as well. I will be waiting for the digram and your response. thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 06:56 AM
When you redistribute a route from BGP into OSPF a tag from the BGP AS is inserted into the route.
You can create a route-map to match on that tag and deny the redistribution from OSPF back into BGP hence avoiding the "Self Serving Routing Loop".
This design will work better if all BGP speaking routers had their own AS #. Having their own AS # will help you determine what router redistributed that route from BGP into OSPF.
Route Tagging is the most scalable solution to avoid routing loops in a complex network as yours.
__
Edison.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 07:23 AM
Is this Edison Ortiz from NYC that either is or was a Novell SysOp?
More on topic, since we use the same AS everywhere, if I understand you correctly, then we couldn't use the default tag. Is there a way to apply a tag in either in the BGP config or the network statements so that we can create a route map statement based off of that?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 07:28 AM
Oh my, Terry Rodecker - Okie boy How are you man?
You can apply a tag on redistribution from BGP into OSPF.
router ospf x
redistribute bgp xxx subnets tag xxx
Then this redistributed route is unique within the OSPF domain and be able to block on other BGP routers doing the redistribution back into BGP.
Good to hear from you man..
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 07:35 AM
Same here Edison. E-mail me at terry dot rodecker at gmail dot com.
I apologize but this is coming to me not as fast as I'd like, if we add custom tags to the BGP redistributed routes we can block those routes from getting into the OSPF database on other routers using a route map statement? Is there a way to add tags to only some of the routes getting redistributed? In other words, tag the non-mpls routes coming in from one of the other 3845s but not the routes coming in from the MPLS only connected sites?
Boy, I really need to get that diagram posted. I'm getting confused talking about it and it's my network.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 07:42 AM
I apologize but this is coming to me not as fast as I'd like,
It's Ok
if we add custom tags to the BGP redistributed routes we can block those routes from getting into the OSPF database on other routers using a route map statement?
You won't block those routes from entering your OSPF database but you will be able to color them with a tags. With this design, a BGP router will bring those routes from BGP into OSPF but another BGP router won't take these routes back into BGP causing this loop. You want these routes to remain in OSPF, not advertised back into BGP.
Is there a way to add tags to only some of the routes getting redistributed?
Yes, you can - with route-maps:
route-map SET-TAG permit 10
match ....
set tag xxx
route-map SET-TAG permit 20
router ospf x
redistribute bgp xxx subnets route-map SET-TAG
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 08:15 AM
Edison
Apologies for butting in but the OP is not redistributing OSPF into BGP, he is using network statements so how will the tags help him ?
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 08:31 AM
Initially we were redistributing OSPF into BGP but as you can imagine, we had less than stellar results as soon as we brought up the second shared connection. That's when we went to network statements. We still had issues so we added prepending using a route-map statement - any "local" routes go out with just the AS added to it, any "backup" routes get the AS prepended a certain number of times depending on if that router is the main backup, the secondary backup, or the tertiary backup for those routes. It would be nice if we could simply go back to automatically redistributing OSPF into BGP and vice versa. Over the years I've found that while you gain a large measure of control over things doing it manually, you also gain a large potential for royally messing things up. We're human and humans make mistakes.
At this point in time, I'm still open to any and all ideas. To control the issue now we simply only have one DS3 currently enabled. We'll manually enable any other DS3 if we need to fail to that site. As bad as that is it's preferable to not knowing if a route to a remote site is going to be working or not when we come to work in the morning.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-01-2009 08:44 AM
Terry it would be helpful if you could provide a small example of what is happening in terms of route choice and what you want to happen ie. router A sees BGP as the best path, why does router B see OSPF as the best path.
