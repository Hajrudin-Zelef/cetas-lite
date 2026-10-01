---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-mx-in-warmspare-doesnt-trigger-failover-when-only-lan-link-is-7c0184c8
title: "t5-security-sd-wan-mx-in-warmspare-doesnt-trigger-failover-when-only-lan-link-is-7c0184c8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-mx-in-warmspare-doesnt-trigger-failover-when-only-lan-link-is-7c0184c8.md
source_anchor: ""
source_lines: [1, 159]
sha256: 540aac8260f482b8072560cf1eaae8e83e7d9f392012bb7cc7d3ef7f89d78a5e
---

# t5-security-sd-wan-mx-in-warmspare-doesnt-trigger-failover-when-only-lan-link-is-7c0184c8

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 08:27 AM
Hello,
We have 2 MXs in Warm Spare. Each MX has connection to ISP-1, and both MXs can talk on LAN. When i unplug LAN cable on primary MX, secondary stop recieving VRRP message and promote himslef as active/master. He continue to processing LAN traffic but it doesn process WAN traffic. Primary MX still have WAN connection so i have in dashboard 2 masters and problem with connection coming from outside. Does anyone know why MX failover full so secondary take response for Lan and Wan traffic? Can you give some advices?
Best regards,
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
11-10-2024 08:09 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 08:45 AM
MX Warm Spare - High-Availability Pair has a troubleshooting section towards the bottom that might assist you. Not sure if you have already read this or are following the suggested implementation practices.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 09:29 AM
Hello,
Yes ,i know troubleshooting and i know that spare doesnt recieve vrrp message. I know regarding design but what i want to say if lan connection on primary mx fail, or switch or switches where primary connected, we dont have full failover. In the docuemntation it says when stop recieving heartbeats it will become active. It still becomes active, but also active one stays active because it has wan online. If i need to explain more please let me know
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 10:26 AM
VRRP only occurs on the LAN side of the MX. Do you have redundant links between the MXs and the Switch(es)? If you only have a single LAN link from each MX you'd want to add another to provide some redundancy.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 10:30 AM
Hello,
Yes i have redundant links, but unplugg both lan links, and the situatiom is like i described above. Both mx is master because both have access to cloud, only one mx lose its lan cables. I am courios if there is mechanisam if lan communication is down,that mx goes in standby instead staying in master state
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 10:39 AM
The only option at that point is to connect the MXs to each other directly so there's a remaining path for VRRP.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 10:48 AM
Hello,
Need a little clarification regarding this please . If i provide directly path for VRRP on LAN 3 on both MXs (LAN 1 and LAN2 connected to the stack), when LAN 1 and LAN 2 is disconnected VRRP message will flow through LAN 3 and how failover in this case will be triggered?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 11:32 AM
VRRP is sent on all links and all VLANs. So yes, in that example of losing LAN 1 & 2 LAN 3 would still carry VRRP packets between the MXs and keep MX1 as primary and MX2 as spare.
Traffic flow of clients downstream would be from the switch up to MX2 then over to MX1 as that would still be the primary unit in the HA pair.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-09-2024 11:28 AM
Hello,
Thank you for your answer. So if i understand good, secondary MX2 will start to answer on LAN virtual MAC address even it is in spare state?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-09-2024 06:01 PM
If the MXs are directly connected to each other MX1 would remain primary and therefore retain answering to the virtual MAC.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-10-2024 07:36 AM
Hello,
Ok, but MX1 lose LAN connections to the stack. So traffic from client pointing on LAN virtual MAC address doesnt have path to the MX1 if i understand good.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-10-2024 08:09 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 10:26 AM
How do you have the setup cabled?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 10:34 AM
Hello,
Both MXs redunandt cable to one of the stack member.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-07-2024 10:36 AM
Why remove two cables in a failover test? I don't know of anyone who designs for two simultaneous and separate failures. I wouldn't worry about the active : active status of both MXs if the (usual) primary is entirely disconnected from the LAN - it can't confuse the clients, in that scenario. But you said the (usual) spare MX is not carrying user traffic, even though it's now active and communicating over the WAN...?
