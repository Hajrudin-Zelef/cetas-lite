---
id: collect-261001-general-networking/general-networking/t5-sd-wan-and-cloud-networking-aaa-tacacs-issue-td-p-3872285-880379ea
title: "t5-sd-wan-and-cloud-networking-aaa-tacacs-issue-td-p-3872285-880379ea"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-sd-wan-and-cloud-networking-aaa-tacacs-issue-td-p-3872285-880379ea.md
source_anchor: ""
source_lines: [1, 100]
sha256: 22bcb340a6fec066f17982a900c01f4b8107280afb680e175b1295a6b8833d6e
---

# t5-sd-wan-and-cloud-networking-aaa-tacacs-issue-td-p-3872285-880379ea

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-13-2019 01:45 AM
Hi, im having trouble connecting all SD-WAN components with my acs version 5.8.1.4-B.462
I think its configured alright. Got set up, authentication order tacacs->local, next add two acs servers, by IP, secret key, and vpn.
%AAA-3-BADSERVERTYPEERROR: Cannot process authentication server type tacacs+ (UNKNOWN)
Solved! Go to Solution.
- Labels:
- 
						
							
		
			SD-WAN Infrastructure
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-14-2019 01:01 AM
Most likely you are running controllers version 18.4.1 or older and this is caused by CSCvn38487, please upgrade to 19.1.0
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-14-2019 01:01 AM
Most likely you are running controllers version 18.4.1 or older and this is caused by CSCvn38487, please upgrade to 19.1.0
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-14-2019 02:38 AM
Thanks, 19.1 got other critical bug, so gonna wait till next major release.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-14-2019 04:05 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-14-2019 04:23 AM
no ip address negotiated, on cellular interface
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-29-2019 02:28 PM
Is that for the vedges or the cedges ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-01-2019 02:09 AM
vedge
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-01-2019 06:45 AM
Ah, good to know. We wanted to start testing 19.1 but that bug would be a show-stopper for us so I guess we don't need to bother with it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-02-2019 12:22 AM
Yes, 19.1 its great, got some nice improvements, like eigrp, but this little detail got me few days extra work with downgrade
https://community.cisco.com/t5/sd-wan/vedge-isr1100-cellular-interface-issue/td-p/3844060
