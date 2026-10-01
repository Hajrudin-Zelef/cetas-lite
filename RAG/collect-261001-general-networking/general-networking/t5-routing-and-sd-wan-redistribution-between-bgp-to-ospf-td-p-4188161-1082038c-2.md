---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c-2
title: "t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c.md
source_anchor: ""
source_lines: [160, 301]
sha256: 26eefbec41734734d58e5d5a0d055683a229dd08f384d255fe96c942b3a856f0
---

# t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c

			Other Routers
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 12:22 AM
Hello,
to be honest, I cannot figure out your topology either. You redistribute BGP into OSPF on two routers ?
R2#sh run | s r ospf
router ospf 1
log-adjacency-changes
redistribute connected subnets
redistribute bgp 235 subnets
R5#sh run | s r ospf
router ospf 1
log-adjacency-changes
redistribute connected subnets
redistribute bgp 235 subnets
Either way, in the BGP process on the redistributing routers, you need to configure 'redistribute bgp-internal:
R2
router bgp 235
redistribute bgp-internal
R5
router bgp 235
redistribute bgp-internal
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 02:11 AM
Hello @kousikdutta ,
Georg is right you have iBGP sessions = same BGP AS number so by default iBGP routes are not redistributed into an IGP like OSPF unless in router bgp configuration you add the command
router bgp 235
bgp redistribute-internal
To be noted this is not common practice in real world where the IGP is used to be able to setup iBGP sessions and all service related advertisements are carried in BGP,
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-24-2020 08:37 AM
I find the topology diagram not clear. It seems clear that R1 connects to R2 and that R4 connects to R5. But what does R3 connect to? Is the connection R1 to R3 to R4 or is it R2 to R3 to R5? Does this suggest that R2 and R5 are on the outside edge of the organization network?
In looking at what is posted I believe that I see at least one issue. In the configuration of ospf I do not see any network statements:
R2#sh run | s r ospf
router ospf 1
log-adjacency-changes
redistribute connected subnets
redistribute bgp 235 subnets
If the bgp routes are being redistributed where would you expect to see the redistributed routes? Who would R2 advertise the redistributed routes to? on R2 the routes would be BGP routes. They would only be OSPF external on some router that R2 advertises to.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-24-2020 05:42 PM
R2 R3 R5
you need command no synch under each ibgp router.
try and see result. 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 12:22 AM
Hello,
to be honest, I cannot figure out your topology either. You redistribute BGP into OSPF on two routers ?
R2#sh run | s r ospf
router ospf 1
log-adjacency-changes
redistribute connected subnets
redistribute bgp 235 subnets
R5#sh run | s r ospf
router ospf 1
log-adjacency-changes
redistribute connected subnets
redistribute bgp 235 subnets
Either way, in the BGP process on the redistributing routers, you need to configure 'redistribute bgp-internal:
R2
router bgp 235
redistribute bgp-internal
R5
router bgp 235
redistribute bgp-internal
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 02:11 AM
Hello @kousikdutta ,
Georg is right you have iBGP sessions = same BGP AS number so by default iBGP routes are not redistributed into an IGP like OSPF unless in router bgp configuration you add the command
router bgp 235
bgp redistribute-internal
To be noted this is not common practice in real world where the IGP is used to be able to setup iBGP sessions and all service related advertisements are carried in BGP,
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 06:49 AM
Hi Giuseppe,
Thanks for yor comment. That is the reason I have tried to use Prefix-list , then configured the route-map and this route-map added in the interface to get the BGP route in OSPF.
Could you please suggest wheather is there any other option that I can use to get this BGP route in OSPF.
Thanks,
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 08:11 AM
Hello @kousikdutta ,
the command in router bgp to allow redistribution of iBGP routes into IGP like OSPF can be seen as a protection mechanism:
a) from a scalability point of view BGP can support much more routes then any IGP hundreds of thousands of routes instead of tens of thousands so it is a protection for OSPF
b) if all routers are running both BGP and OSPF this redistribution is not necessary and OSPF is used to provide loopback advertisements and inter router links (infrastructure) and all everything else (services ) is carried in BGP
>> Could you please suggest wheather is there any other option that I can use to get this BGP route in OSPF.
As far as I know as I have explained above there is no other way to get an iBGP route redistributed into OSPF.
Hope to help
Giuseppe
