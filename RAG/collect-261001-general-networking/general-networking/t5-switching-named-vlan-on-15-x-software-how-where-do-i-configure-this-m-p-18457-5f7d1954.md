---
id: collect-261001-general-networking/general-networking/t5-switching-named-vlan-on-15-x-software-how-where-do-i-configure-this-m-p-18457-5f7d1954
title: "t5-switching-named-vlan-on-15-x-software-how-where-do-i-configure-this-m-p-18457-5f7d1954"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-named-vlan-on-15-x-software-how-where-do-i-configure-this-m-p-18457-5f7d1954.md
source_anchor: ""
source_lines: [1, 104]
sha256: d4bcd558f70e6c8edbd6e09d354aeac79f5861ff19bcd069bbee64351e51ceff
---

# t5-switching-named-vlan-on-15-x-software-how-where-do-i-configure-this-m-p-18457-5f7d1954

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2022 06:47 AM
Hi all
I saw that there was a named VLAN feature on the 15.x software train.
And I upgraded some switches with that software, but I cannot seem to find any new fields or meny points where I can configure this. And there does not seem to be a documentation page about it (except on a "Beta" page that requires a Meraki login ).
Do I have to ask support for the feature ?
Thanks
Thomas
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
06-27-2022 07:22 AM
I'd contact Meraki Support. Seeing the same as yourself
https://www.linkedin.com/in/darrenoconnor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2022 07:03 AM
There’s a previous post which mentions named VLANs using RADIUS assignment
https://community.meraki.com/t5/Switching/Named-VLAN-for-port-config/m-p/133469#M9864
https://www.linkedin.com/in/darrenoconnor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2022 07:07 AM
Yeah I found that one, but it does not really mention anything about configuration.
And it is actually for VLAN assignment from Radius that Im looking for.
I mean, I could do it with a group policy I suppose, but that is not really the same thing as named VLAN.
So I still need some documentation and / or config guide on this.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-27-2022 07:22 AM
I'd contact Meraki Support. Seeing the same as yourself
https://www.linkedin.com/in/darrenoconnor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-30-2022 07:11 AM
Just to "end" this thread. Yes, you have to contact Meraki support to get the feature enabled. And Yes, it works, when returning a name for a VLAN from ISE instead of the number. So all good.
(Perhaps Ill just accept my own answer as the solution ? )
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-01-2023 06:09 AM
So how has it been running?
I wonder why this is still not public or early access via the dashboard.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2023 01:04 AM
I don't remember that we have encountered any issues. So I have no idea why it is still not public.
Perhaps it's just one of those features that is being worked on because it might not work on "all" switches, or I'm just the lucky one - MS2xx and 12x switches here. (We all know who the "problem" switch is )
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-01-2022 07:12 AM
Would be nice to keep a VLAN list with names on the switches like on Cisco "classic" switches with the command: sh vlan
