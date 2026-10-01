---
id: collect-261001-general-networking/general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f-3
title: "t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f.md
source_anchor: ""
source_lines: [184, 294]
sha256: 9a255b51cdbc695ebf50f05c076a4d81eb453282013e83ab9585be91743785eb
---

# t5-switching-multiple-ospf-processes-td-p-1335402-24efa88f

If I were in router 1 and did a "sh ip ro" for one of the /29 subnets hanging off of router 2, what would I see? Would it be a directly connected route, since it falls within the range of the supernet under the OSPF 599 process of router 1? Or would it see a specific route pointing to router 2 for the specific subnet?
Does the answer to this question have to do with whether the routers are creating and sending out specific LSAs for each /29 subnet or if they are creating and sending out LSAs for the summarized (supernet) addresses only?
If the routers are sending specific LSAs for each /29 subnet, then router 1's routing table should point to router 2 as the next-hop for that subnet hanging off of router 2. However, if only an LSA is generated for each supernet, then router 1 will see the route to the /29 hanging off of r2 as a "directly connected" route."
My guess is that the routers are indeed creating specific LSAs for EACH /29 subnet because there is no summary route configured under the OSPF 599 process. If there were, then there would be only ONE LSA for the entire summarized address range.
Is the above analysis and assumptions correct?
Unfortunately, I do NOT have access to the routers, or I would have done all the discovery myself. We have a situation where a remote engineer has access and me and my engineer have to walk him through some routing issues.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-14-2009 07:27 AM
Joe
There are some things about the topology of the network and what is going on that I do not yet understand very well. But I believe that enough is clear to answer your main question. If we need to go deeper then you may need to clarify some things.
Each router should see individual routes for each subnet. For the VLAN/subnets to which it is connected they will appear as connected routes. For the subnets for which it is not directly connected the route should be intra area routes and should have the other router as the next hop. This would be true even if the routers were configured to generate summary routes - and they are not configured to generate summaries.
The important concept here is that the VLAN/subnets are all in area 2. And within an area all OSPF routers will see all of the detail for all the prefixes within the area.
Let me also clarify that just because some subnet falls into a range used in a network statement on a router, it does not mean that the subnet would show as connected. The only "connected" routes are the ones to which your router is actually physically connected.
Let me also clarify that just because the OSPF configuration has a range in the network statement, that does not mean that OSPF will generate a summary for the range. The range in the network statement is only to simplify the process of determining which interfaces get included within the routing process. If you want summary routes in OSPF then you must specifically configure it to generate summary routes.
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-15-2009 08:25 PM
Rick:
I understand much of what you have said...
So, just to summarize, each router will create and send an LSA for EACH routed interface configured on it that participtes in OSPF. Correct?
So, each of these routers has about 125 SVIs (/29 subnets) configured on them, so each router should generate an LSA for each of those /29 subnets. Correct?
So, this is why, router 1 will point to router 2 as the next hop for a subnet configured on router 2, and NOT view it as a directly connected route - because it is not the fact that the network statement under OSPF encompasses the whole range of subnets that determines which LSAs are created and sent, but the actual interfaces that are configured on each router. Correct?
Thanks and sorry for taking so long to get ack to you. Im getting slammmed with 14 hour days this new client.
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-16-2009 03:51 AM
Joe
You are pretty close. Let me try to clarify one of the things that you say:"it is not the fact that the network statement under OSPF encompasses the whole range of subnets that determines which LSAs are created and sent, but the actual interfaces that are configured on each router."
It is not just the network statement and it is not just the interfaces on the router that are configured. It is the combination of the network statement and the configured interfaces that determine what is advertised by OSPF.
To explain in a bit more detail:
When the OSPF process starts it looks at its configured network statements (and the address ranges defined by the address and the mask) and it looks at every interface on the router (that is in up/up state) and if an interface falls into the range defined by a network statement then that interface is included into OSPF. Then OSPF looks at the subnet defined on that interface (including its mask) and advertises that subnet.
So to clarify a couple of points:
- the network statement does not tell OSPF what to advertise but tells OSPF what interfaces to process.
- the network statement does not tell OSPF to summarize (there are separate commands to control summarization).
- OSPF will determine what to advertise based on the configured subnets on the interfaces that it includes in its processing.
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-16-2009 09:28 AM
Thanks, Rick..excellent exlanation...got it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-05-2018 10:47 PM
Hi I have an scenario using two ABR routers where there two OSPF Process ID using Area 0? How can connect the Process ID that is part of the Aggregation ring? Since I'm facing some issues when the link goes down in the AGG.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-06-2018 12:46 PM
I have looked at the diagram that you posted but am still confused about what is going on. I see ABR1 and ABR2 And a router between them. What is that middle router, is it in both OSPF processes?
I see a dashed blue line that seems to indicate what is the area 0 for OSPF10 It shows a direct connection between the ABRs and connection through the core nodes. And I see a dashed green line but am not sure about it and what it means. It would seem to indicate the area 0 for OSPF 100. But both lines run through the agg routers and there does not seem to be any direct connection for these routers in OSPF 100. Can you provide clarification?
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-13-2018 09:55 AM
Hi the router between RR participate in OSPF10, is not part of OSPF100. The OSPF100 participate between AGG and RR nodes. I understand the best solution is create a link between RR and be part of OSPF 100, sometimes is not posible in the real scenario. Some issues are present if there a link donw in the AGG nodes, to reduce this I redistribute the OSPF into BGP. I'm using MP BGP, but when I show the routes The AGG nodes learn by OSPF the Lo0 of RR, but in the RR the Lo0 is learned via BGP (If I generate a link down between AGG ABR04 and ABR2 RR inline)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-21-2013 05:01 PM
Richard:
You wrote (I changed the hyphens to a, b, c):
The implications of running separate OSPF processes include these:
a) an interface can be active in only 1 interface. So each OSPF process will have a unique set of neighbors.
b) each OSPF process will learn its own prefixes and maintain those prefixes in its own database.
