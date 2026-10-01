---
id: collect-261001-huawei/huawei/t5-security-sd-wan-performance-issues-on-mx100-running-versions-16-6-or-17-6-m-p-ec831cd9
title: "t5-security-sd-wan-performance-issues-on-mx100-running-versions-16-6-or-17-6-m-p-ec831cd9"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "throughput"]
source: docs/RAG/collect-261001-huawei/t5-security-sd-wan-performance-issues-on-mx100-running-versions-16-6-or-17-6-m-p-ec831cd9.md
source_anchor: ""
source_lines: [1, 155]
sha256: 6676da55761cf03d746b5378ec565dc9793c9d97ad75679db6a2e4a3ea10b4b5
---

# t5-security-sd-wan-performance-issues-on-mx100-running-versions-16-6-or-17-6-m-p-ec831cd9

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-12-2022 02:51 AM
I'm wondering if anyone else has had performance issues on the MX100 once upgraded to 16.6 or 17.6. We've got an MX100 that runs great on 15.44 but plain internet throughput seems to drop by about half on version 16.6 or 17.6. We need to upgrade from 15.44 due to SNORT vulnerability. This MX100 has content filtering and AMP and IDS enabled. 
We've rolled back to 15.44 from 16.6 and 17.6 a few times now and performance always returns to normal.
Thanks
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
08-19-2022 09:35 AM
That is a good question, I think Meraki support would have to check into that, a follow up for performance, my MX100's did handle near 1gbps the entire night without issue.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-12-2022 01:29 PM
I don't have a lot of customers with MX100s, but I have not seen that issue on 16.16. I don't have anyone with an MX100 running 17.x code.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-13-2022 01:58 AM
@martins@netxuk.com 17.7 is now released and includes more performance fixes, please try that.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-16-2022 02:04 AM
@CMR Thanks for the heads up on 17.7. We upgraded the MX100 to 17.7 over the weekend but throughput still comes out about half of what we had on 15.44.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-13-2022 05:14 AM
It is better for you to use 15.44 until the 17.7 gets more time to be tested from the other users. 
Both firmware 16.16 and 17.6 are not stable firmware even for the other Models like MX450 or MX84. We had last month a lot of problems with 16.16 which included: VPN interruption, High Resource Utilization, packet loss and high latency (and for specific MX model throughput degradation too). So, my suggestion would be to keep your devices running on 15.44 Firmware. 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-16-2022 02:19 AM
@Nicholas Kule, thanks, would have been happy to stay on 15.44 if it did not have the SNORT vulnerability.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-20-2022 02:52 PM
We have the same issue. We have 2 MX100 and 15.x is the highest release we can use. You estimated correctly, there is about a 50% performance hit. We have tried RMA the MX100 and same issue. Sending the new MX back to Meraki and staying on 15.44 My thoughts are 16.x and 17. are not ready for prime time.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-21-2022 12:46 AM
Meraki TAC have told me that it's a recognised issue and they are working on a fix. I'll update this thread as and when I get more news.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-02-2022 08:25 AM
I think the issue is that in the firmware upgrades starting with 16.x, someone hardcoded the 500mb setting on the Uplink Configuration under SD-WAN & traffic shaping. Looking at the dashboard, it seems the dashboard setting for this makes no difference no matter how you set it so I think the backend code has disconnected from the GUI and hardcoded at a particular level. Here's hoping Cisco TAC finds this and fixes it soon. We now have both our MX100s back down to 15.44 and the Uplink configuration can be altered via the GUI to any setting.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-03-2022 12:54 AM
That's intereasting dgander. For us it looks like it's related to Snort v3. When running versions 16 or 17, performance seems OK as long as IDS is switched off. As soon as IDS is enabled again the throughput drops by about half. We tried 16.16.4 most recently but still had the same issue.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-03-2022 05:05 AM
I haven't tried turning that off or on, we now are staying on 15.44 until it is identified so I will let Cisco play with it in their lab. Not a solution either to turn it off but interesting that you saw a drop with it on and full throughput with it off.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-18-2022 01:13 PM
I'm running 17.9, (MX 100's HA) I just worked with support and the performance issue was resolved after they downgraded snort from Version 3 to Version 2.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-19-2022 12:59 AM
Hi Rob2041, that's very interesting to hear. I was unaware that a version 17 variant could be run with Snort on V2. Do you know if the snort vulnerability in V2 has been addressed in this particular scenario?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-19-2022 09:35 AM
That is a good question, I think Meraki support would have to check into that, a follow up for performance, my MX100's did handle near 1gbps the entire night without issue.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-24-2022 01:17 AM
Hi Rob2041, Meraki TAC tell me that the MX running v17 and Snort v2 still has the Snort vulnerability. They also tell me that the performance issue is now sorted for MX250/450 in v17 and they are continuing to work on a fix for other MX models running newer firmware than v15.
