---
id: collect-261001-general-networking/general-networking/t5-switching-default-static-route-and-ospf-routing-on-ms425-m-p-134592-a79f7a6d
title: "t5-switching-default-static-route-and-ospf-routing-on-ms425-m-p-134592-a79f7a6d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "voice"]
source: docs/RAG/collect-261001-general-networking/t5-switching-default-static-route-and-ospf-routing-on-ms425-m-p-134592-a79f7a6d.md
source_anchor: ""
source_lines: [1, 92]
sha256: 80ce70111c5a4fb1ffed39a43ebcc2999efe3b3172b4e4eaad55a5b713887f8e
---

# t5-switching-default-static-route-and-ospf-routing-on-ms425-m-p-134592-a79f7a6d

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2021 04:01 AM
We are just introducing MS425 as L3 switches into our network. we initially just had a default static route. Now we have configured OSPF, we want the default route to come from the primary or secondary WAN routers via OSPF.
1) I can't remove the Static route
2) What ever I do it still sees the route as static
though it says this
How do I fix!
On a Classic Cisco I would just increase the cost of the Static route
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2021 09:53 AM
Is there a default route in EIGRP or BGP that should be distributed via OSPF? I would typically expect to see default-information originate in the OSPF section to propagate the default route.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2021 03:08 PM
what is the ospf config of the routers?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2021 04:22 PM
Here is the OSPF for WAN routers
RT01-Primary for Data
**********************
router ospf 1
router-id 172.16.98.42
redistribute eigrp 1 subnets
redistribute bgp 64589 subnets route-map BGP-to-OSPF
network 172.16.100.216 0.0.0.3 area 0
route-map BGP-to-OSPF, permit, sequence 10
Match clauses:
ip address prefix-lists: vocus-sip-service
Set clauses:
metric 70                                   <<<<< higher cost for voice service
Policy routing matches: 0 packets, 0 bytes
route-map BGP-to-OSPF, permit, sequence 20
Match clauses:
ip address prefix-lists: Prefix-in
Set clauses:
metric 20                          <lower cost default route here
RT02-Primary for Voice
***************************
router ospf 1
router-id 172.16.97.42
redistribute eigrp 1 metric 160 subnets
redistribute bgp 65032 subnets route-map BGP-to-OSPF
network 172.16.100.224 0.0.0.3 area 0
route-map BGP-to-OSPF permit 10
match ip address prefix-list vocus-sip-service
set metric 20                  <<<<Lower cost for Voice
route-map BGP-to-OSPF permit 20
match ip address prefix-list Prefix-in
set metric 70  <<< higher cost for default traffic
Normally default route via RT01, but if problem on RT01 will go thru RT02,
But the default route is not advertised on the Meraki, has static going to RT01 so would cause problems if RT01 or WAN failed
On a cisco router would normally increase the cost of the static route to give preference to OSPF
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2021 09:53 AM
Is there a default route in EIGRP or BGP that should be distributed via OSPF? I would typically expect to see default-information originate in the OSPF section to propagate the default route.
