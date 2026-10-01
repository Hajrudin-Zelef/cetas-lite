---
id: collect-261001-general-networking/general-networking/t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf-2
title: "t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf.md
source_anchor: ""
source_lines: [34, 230]
sha256: 2ecb52bb67873e93af28b97ab459d0aa8a36a6d4bb5f61f48fe745248c9b59cf
---

# t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf

			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-23-2025 01:37 AM
Not sure why 3 downlinks are needed….seems a bit overkill or over engineered.
Too many redundant paths likely intended for a failover situation but the errors seem to be the redundant paths (3 downlinks).
If it were my network to remedy, each switch would have only two downlinks with all root traffic from SW00.
SW01 and SW02 would be one alt path (loop guard)
SW03 and SW04 would be one alt path (loop guard)
I would remove the link between SW02 and SW03.
I would remove the link between SW01 and SW04.
Thanks for sharing the topology to better understand. 
That would be my suggestion/recommendation. 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-22-2025 01:56 PM
I realize you mentioned inheriting this network but the device health notice indicates SW01 and SW03 error notices yet the lower images you share indicate SW04 and the port descriptions don’t seem to jive. Seems like something doesn't quite add up.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-22-2025 01:59 PM
Can you share a screen capture of the topology from the meraki dashboard to see what that shows....I'm curious to see if that might reveal more insight. Not meant as critical at / toward you....just tried to understand the above and I looked at it like 10 times.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-22-2025 10:37 PM
Hi @Sweemz ,
Not sure I follow the Topology Diagram either. Port 52 is directly connected to the Root switch. So that would be the Root port. Port 49 & 50 would be alternate ports.
I agree with @RWelch2. The Topology map in the Dashbaord is 'live". So if the Tree changes so will the map. I use it daily. Best tool to use for these kind of queries.
https://documentation.meraki.com/MS/Monitoring_and_Reporting/Network_Topology
If you can share a screenshot, that'll be helpful.
Cheers,
Ivan Jukić,
Meraki APJC
If you found this post helpful, please give it kudos. If it solved your problem, click "accept as solution" so that others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-23-2025 01:43 AM
Hi @IvanJukic
Thank you for looking into this.
2nd glance i've clearly made some mistakes with my labels
Using the Network topology tool I have re-confirmed the uplinks/ports and updated my documentation to reflect the live links in the Meraki Toplogy.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-23-2025 07:48 PM
No problem at all @Sweemz. 
Given the devices have already been purchased you could increase efficiency and redundancy. My preference would be to clean up the design, rather than chasing errors and warnings. 
Judging by that Topology Screenshot. It Seems like, RSTP Priority is set to the default of 32768, and SW00 happens to be the Root Bridge. However, without actually looking into the Network, I can't verify.
Also, it looks like SW01 could be a Backup Root or potentially used as one. If so you could have; 
A) SW00 & SW01 connected together as a Core/Distribution. With uplinks to the MX-HA Pair.
B) SW00 as Root Priority 0
C) SW01 backup Priority 4096
D) SW02,3,4 leave Priority as is. Then only have 2 uplinks, 1 each to SW00 & SW01.
E) All other Devices connected to only SW02,3,4
Cheers,
Ivan Jukić,
Meraki APJC
If you found this post helpful, please give it kudos. If it solved your problem, click "accept as solution" so that others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-23-2025 01:20 AM
Not sure why I didn't include these originally .. Thank you for looking into this btw
Toplogy
Switch Overview
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-23-2025 01:37 AM
Not sure why 3 downlinks are needed….seems a bit overkill or over engineered.
Too many redundant paths likely intended for a failover situation but the errors seem to be the redundant paths (3 downlinks).
If it were my network to remedy, each switch would have only two downlinks with all root traffic from SW00.
SW01 and SW02 would be one alt path (loop guard)
SW03 and SW04 would be one alt path (loop guard)
I would remove the link between SW02 and SW03.
I would remove the link between SW01 and SW04.
Thanks for sharing the topology to better understand. 
That would be my suggestion/recommendation. 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-23-2025 03:26 AM
Thank you @RWelch2
Makes absolute sense.
I'll have a play with this network this evening and provide an update on the topology asap.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-22-2025 02:12 PM
APs do not participate to the STP topology so that shouldn't matter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-23-2025 07:08 AM
I agree with RWelch too many redundant paths for the failover situation and the errors do seem to be the redundant 3 paths.
I see you are on MS 16.9 and there is an update firmware for these switches as well MS 17+ would help as well.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-27-2025 08:11 AM
Update - 27/01/25
2 x Links have been removed
- Link between SW02 and SW03
- Link between SW01 and SW04.
Updated Topology
Observation
Alert on DUG00-SW01 has disappered
Alert on DUG00-SW04 persists
Going to attempt port cycles on DUG00-SW03, Port 50 & DUG00-SW04, Port 49 to see if it gives it a kick. If the Alert persists I will also try a reboot of switches this evening to see if this clears out the alert.
IF all goes well I will then schedule in a FW update later this week.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-27-2025 09:57 AM
If an error or alert remains, I could simply double check the following port STP guard settings: 
STP Guard setting between SW01 and SW02 ports (on both switches) would be loopback guard.
STP Guard setting between SW03 and SW04 ports (on both switches) would be loopback guard.
STP Guard setting (uplink) on SW01, SW02, SW03 and SW04 going up to SW00 would all be disabled.
STP Guard setting (downlink) on SW00 going down to SW01, SW02, SW03 and SW04 would all be root guard.
And as @IvanJukic mentioned, setting the SW00 as bridge prioirity 0.  I would personally set all the other switches (SW01-SW04) at root bridge priority of 61440, however he mentioned leaving at the default 32768 (he might know something that I don't in that regard).
Dashboard > Switching > Configure > Switch Settings where you can set STP Configuration (STP bridge priority)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-30-2025 04:27 AM
Thanks @RWelch2 I can confirm that the STP Guard settings are configured exactly as you suggested/described.
For the STP bridge priority I have DUG00-SW00 configured as '4096' and then anything else falls under the default of '32768'.
Should I explicitly define all other switches in here as a higher priority e.g '32768'? All other switches do show DUG00-SW00 as the RSTP root.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-30-2025 04:49 AM
