---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-ibgp-neighborship-using-loopback-td-p-4843273-b6d31c83
title: "t5-routing-and-sd-wan-ibgp-neighborship-using-loopback-td-p-4843273-b6d31c83"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-ibgp-neighborship-using-loopback-td-p-4843273-b6d31c83.md
source_anchor: ""
source_lines: [1, 169]
sha256: 05b3428ae8bba26e54adb86426de8b3390456581f99c64519cc7b594023f5d36
---

# t5-routing-and-sd-wan-ibgp-neighborship-using-loopback-td-p-4843273-b6d31c83

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-25-2023 10:58 PM
Hello, 
I want to know how iBGP manipulates OSPF routes to R2 loopback to go via R4 when the link between R1 and R2 goes down.
Here I am establishing iBGP neighborship using loopbacks. I have shut down the fa0/0 interface of R1.
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
05-27-2023 12:02 AM
Thanks to your ospf, acting as an underlay, you still get path to R2's loopback via R4. Then iBGP session beween R1 and R2, based on loopback routed via ospf, is always established.
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-25-2023 11:16 PM - edited 05-25-2023 11:16 PM
Hello @Karan Belani,
I perhaps not understand your need so sorry.
You seem to have configured OSPF to route Routers' loopback and you configure iBGP based on these loopback.
Then when the link between R1 and R2 goes down, R1 loopback is known via R4. IBGP does not manipulate OSPF routes.
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-26-2023 11:56 PM
Hello, 
I have established neighborship between all routers using their loopback and for the connectivity to loopback, I have used OSPF. 
Now when I shut down the link between R1 and R2 the routes to R2 loopback from R1 via R4. I want to know how did this happen?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-27-2023 12:02 AM
Thanks to your ospf, acting as an underlay, you still get path to R2's loopback via R4. Then iBGP session beween R1 and R2, based on loopback routed via ospf, is always established.
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-27-2023 12:11 AM
Hello M02@rt37 , 
Thanks, I understood it now, because of OSPF the neighborship is UP.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-27-2023 12:13 AM
You're right @Karan Belani
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-26-2023 12:13 AM - edited 05-26-2023 12:13 AM
the BGP dont care about the path IGP use direct or via R4
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-26-2023 11:57 PM
Hello, 
So here even if I didn't use BGP, I would still get path to R2's loopback via R4 because of OSPF?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-27-2023 02:05 AM - edited 05-27-2023 03:01 AM
Yes' correct
Here you can configure in each router two LO' one advertise via ospf and use as update source of bgp
Other LO advertise via bgp.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-27-2023 09:22 PM
Hello, 
Thank you for helping.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-26-2023 03:32 AM
Hello @Karan Belani ,
iBGP sessions are not limited to TTL =1 like a standard eBGP session does. By using iBGP sessions on loopbacks the iBGP session is able to survive the shutdown of R1 LAN port by using OSPF to reroute.
You can see iBGP like an application using OSPF IGP routing services.
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-27-2023 12:00 AM
Hello, 
Thank you for explaining, so here BGP neighborship is maintained with the help of OSPF reroute am I right?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-28-2023 11:10 PM
Hello @Karan Belani ,
>> so here BGP neighborship is maintained with the help of OSPF reroute am I right?
Yes you are right
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-29-2023 04:19 PM
Hello, 
Thank you, I understood it now, I was so focused on BGP that I forgot about OSPF.
