---
id: collect-261001-general-networking/general-networking/t5-switching-layer-3-switch-supporting-nat-pat-needed-td-p-2292566-39e3cd82
title: "t5-switching-layer-3-switch-supporting-nat-pat-needed-td-p-2292566-39e3cd82"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-switching-layer-3-switch-supporting-nat-pat-needed-td-p-2292566-39e3cd82.md
source_anchor: ""
source_lines: [1, 64]
sha256: e157fbcec08dbb82f44b75a799f96d5a0dda20d363803ac03236626c4f6e646a
---

# t5-switching-layer-3-switch-supporting-nat-pat-needed-td-p-2292566-39e3cd82

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-25-2013 12:50 AM - edited 03-07-2019 02:34 PM
Hi all,
I'm looking for a low cost switch that supports layer 3 IP routing along with NAT/PAT functionality.
Ideally it will have 24 RJ45 GbEth ports.
Can anyone suggest a suitable model.
Also, I have read that the 2960 model switches support basic layer 3 functionality but it is not clear if I can configure IP addresses on the physical port interfaces and also configure NAT/PAT. I'm doubting this is supported but can someone confirm this.
Thanks in advance.
Peter
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-25-2013 01:08 AM
Hi,
only 6000/6500 and 5500 switches support NAT.
on a 2960 you can configure an IP address on a port if you make this port a L3 port with the no switchport interface command.
Regards
Alain
Don't forget to rate helpful posts.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-25-2013 01:08 AM
Hi,
only 6000/6500 and 5500 switches support NAT.
on a 2960 you can configure an IP address on a port if you make this port a L3 port with the no switchport interface command.
Regards
Alain
Don't forget to rate helpful posts.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-25-2013 01:58 AM
Hi Alain,
Your response is very helpful. Thanks very much for your reply.
Have a nice day.
Cheers
Peter
