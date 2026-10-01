---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45-2
title: "t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "throughput"]
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45.md
source_anchor: ""
source_lines: [36, 212]
sha256: b136eaf471020a4faafebd5c646cffa171e69dc962245bdbfe6f1f88176ad05b
---

# t5-routing-and-sd-wan-ask-the-expert-configuring-and-troubleshooting-border-gate-722f8d45

			Routing Protocols
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 12:33 AM
Hi Sandeep,
PFA My NW diagram with proposed link.
We are using three ISPs bandwith with eBGP, we have our own IP address and ASN.
Now we are going to start another site with different location with same ASN.
Router A NW IP : 102.21.20.0/22 advertised with Three ISP
Router A (ASN ) 23456
Router B NW IP 102.21.22.0/24 advertised with another ISP in different location with same ASN.
Router B (ASN) 23456
When another ISP b/w goes down then i need my all the traffic going via iBGP (Router A).
My Requirement when link goes down between Router – B to another ISP (Proposed) then my all the traffic working via iBGP.
So what configuration in my both the Router A & B to fulfill my requirement.
Thanks in ADV,
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 05:00 AM
Hi,
Below is the response on the basis of my understanding to your query:
For giving preference to exit traffic at site A you can use weight attribute as all the ISP's are connected on the same router.
And for influencing the exit traffic for site B you should use the local preference below is the configuration for router B.
router bgp 23456
neighbor 
neighbor 
neighbor 
ip as−path access−list 7 permit ^
route−map setlocalin permit 10
match as−path 7
set local−preference 400
route−map setlocalin permit 20
set local−preference 150
>>>>> you can also use default local preference command in place of using AS-path to simplify and you want to use for whole traffic.
- If you want to infulence the incoming traffic you can use MED attribute.
In case you have any specefic query and not answered here please feel free to ask again.
Thanks & Regards
Sandeep
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 05:08 AM
Hi Sandeep,
Thanks for your great help...
Bellow config i have to config in my Router A right ??
I am bit confused If as per your suggested config in Router A , becasue why ISP-4 configuration in Router A becasue it is not directely conneceted with Router A it is connected directely with Router B.
Great help pl clear my dought.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 05:29 AM
Hi ,
This is router B configuration not router A. I also mentioned in above reply aswell
Please feel free to contact in case you have any furhter query
Thanks & Regards
Sandeep
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 09:49 PM
Hi Sandeep,
Just for clarification you are mention for router B as given bellow.
router bgp 23456
neighbor 
I think " neighbor 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2013 09:44 AM
Hi,
You are correct just a typo, it will be the IBGP peer IP address and for RTRB router peer is RTRA. so this is RTRA-IP .
Thanks & Regards
Sandeep
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 05:05 AM
Hello Sandeep Sir,
Just wanted to know what is a BGP slow-peer and what ar/is the way to mitigate this issue.
/Imran
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 05:43 AM
Hi Imran
BGP Slow peer cases often are reported as "missing update", "slow update", "stopped update" or "session flap due to Hold timer expiry when local BGP is not able to send the updates to neighbor for the time interval of hold time" type issues, rather than being identified as a slow peer issue by the customer.
You can confirm that a case is due to a slow peer by issuing
show ip bgp all summary
and watching the routing table versions associated with various neighbors. The problem neighbor's version will typically increase slowly, if at all, and frequently, but not always, have a large outQ of unsent BGP messages.
The command
show ip bgp all update-group
    show ip bgp 
Will show you which neighbors are in which update-group. A slow peer only impacts neighbors in the same update-group. If there is more than one update-group, you can check and make sure that the impacted neighbors are indeed in the same update-group as the slow peer.
If a * is marked in front of the neighbor then that shows that updates are being sent to the neighbor. If the * mark is not removed for a period of a minute then it must be a slow peer.
One way to find the slow peer is issue
    show ip bgp neighbor 
    show ip bgp 
Look for "Keepalives are temporarily in throttle due to closed TCP window" or TCP receive Window Size is very low or Zero. Repeat this for all the neighbors in the update group. If a neighbor displays above message then it might be a probable slow peer. Coupls of reasons for slow peer might be
There is packet loss and/or high traffic on the link to the peer and the throughput of the BGP TCP connection is very low.
The peer is heavily loaded in terms of its CPU and cannot service the TCP connection at the required frequency.
You can try few wokraround to fix the slo peer issues like :
- If IOS version doesn't support the Slow peer detection & protection feature then identify the slow peer from the steps listed above and move the slow peer to different update-group group by configuring dummy policy or by changing "advertisement-interval" interval different than the rest of the neighbors "neighbor 
How to mitigate the slow peer:
============================
- While fully resolving a slow peer situation requires addressing the issue which is causing it to be slow, such as packet loss between the RR and the peer, or an overloaded CPU on the slow peer, you can mitigate the problem my moving the individual peer into its own update group, so that its slowness does not impact other peers.
- More recent Cisco IOS releases contain automatic slow peer mitigation features which can be turned on.
- for Older releases which do not contain these features. To mitigate a slow peer on these older releases, you need to change the configuration so that the slow peer is forced into its own update-group. you can do this by configuring a dummy route-map and apply it to just the one peer. You may need to remove the peer from a peer-group or other shared configuration in order to do this.
Moving a neighbor into its own update-group causes the router to engage in additional processing, which will increase CPU utilization and memory consumption.
Hope this answers your query, In case you have anny further query please feel free to post.
Thanks & Regards
Sandeep
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-23-2013 06:49 AM
Thanks Sandeep! That was a clear and lucid explanation. It was helpful!!!
/Imran
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-24-2013 02:44 PM
HI Sandeep,
Thanks for open up this discussion on BGP actually I’m looking for a BGP solution, my query as below
If we have 2 WAN routers and a single MPLS connectivity running BGP AS 200, then how we can use our both WAN routers to get hardware redundancy, as service provider is not ready to give duel BGP peer on single link.
Attaching diagram for more clarity
Thanks/SANJEEV
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-26-2013 06:30 AM
Hi Sanjeev,
