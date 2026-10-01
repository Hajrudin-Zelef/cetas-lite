---
id: collect-261001-huawei/huawei/t5-security-sd-wan-how-to-troubleshoot-performance-issues-in-a-auto-vpn-connecti-6d861a5e
title: "t5-security-sd-wan-how-to-troubleshoot-performance-issues-in-a-auto-vpn-connecti-6d861a5e"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/t5-security-sd-wan-how-to-troubleshoot-performance-issues-in-a-auto-vpn-connecti-6d861a5e.md
source_anchor: ""
source_lines: [1, 144]
sha256: 7d94ca7a758cb95e564b6cc4bfd75c410952657573ca998db6da820ff9d3e429
---

# t5-security-sd-wan-how-to-troubleshoot-performance-issues-in-a-auto-vpn-connecti-6d861a5e

How to troubleshoot performance issues in a Auto-VPN connection?
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-31-2023 01:34 AM
Hi,
a customer of mine reports sporadic disconnections during RDP sessions, but only from locations which are connected through Hub&Spoke Auto-VPN...
Does anyone of you have a tip on how best to troubleshoot such performance problems between locations that are connected via Meraki Auto-VPN out from the Meraki Dashboard with no 3rd party tools?
Every help is appreciated, thank you all in advance!
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-31-2023 03:06 AM
There is a documentation for It.
https://documentation.meraki.com/MX/Site-to-site_VPN/Site-to-Site_VPN_Troubleshooting
3rd party VPN is limited.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-02-2023 11:53 PM
@alessandrodematosI know that documentation, but unfortunately I´m not sure what you mean exactly what could help me with that - finding out what causes the interruption?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-03-2023 06:01 AM
Why don't you just do a packet capture?
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-03-2023 09:11 AM
easier said than done... first of all as already replied to @Raphletourn it`s really difficult to handle that kind of problem with the packet capture on the dashboard, because the problem occur only sporadically and for a short period of time (user reports that the session had disconncted and automatically re-connected), so when to start the capture? Next thing, what will I see in the capture? If a performance related problem and the Application is running over TCP/3389 than maybe restransmissions and out-of-order packets followed by Resets , but the cause of that remains still unclear
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-03-2023 09:22 AM
You're assuming it's a Meraki problem but you're not sure.
I don't remember if you informed the software version that is running on the MX, but if I were to bet on something, I would say that it is a version problem. I myself have had issues printing over the network and RDP with any 18.x version and had to downgrade.
In your place I would open a support case.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-03-2023 09:30 AM
honestly speaking at the moment I don`t know what the cause of the problem is, I´m responsible for the WAN and only got the information that there are interruptions over the Auto-VPN connections! So, no clue if the Meraki Software is the reason or if the access in the underlay which is hybrid (MPLS and Internet) is overloaded with bursty traffic... it`s really frustrating that there`re no helpful informations on the dashboard to troubleshoot that in a serious manner
thank you @alessandrodematos for the hint with the software - I´ve checked back and all MX´s are currently using 18.107.2 - so I´ll follow your suggestion and open a support-case!
One more question to you regarding the downgrade to 17.x... did you get more or detailed informations in the response, in regard of what was the problem in your case with 18.x?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-03-2023 11:24 AM
No, as it was something critical I didn't take the time to investigate further.
But I've noticed complaints from other members and Meraki never spoke or gave more details.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-31-2023 04:55 AM
Running multi point packet capture would be my way to troubleshoot that.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-02-2023 11:56 PM
@RaphletournI think that would be the best way to figure it out, but my problem is that the issues occur sporadically and so fast... what is really difficult to handle with the packet capture on the dashboard! In my opinion it`s also hard to analyze and find out in behalf of the Capture-files if potential performance problems are the cause of the problem and of course what traffic is the reason for them
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-31-2023 02:35 PM
This is not a great solution for your problem, but you could consider buying a couple of Meraki Insight licences.
https://documentation.meraki.com/MI/Intro/Meraki_Insight_Introduction
Insight can not monitor RDP sessions. It can only monitor web apps. Hopefully, they access web apps over AutoVPN as well, and if so, you could use the monitoring of those to help form a conclusion about RDP.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-02-2023 11:58 PM
@Philip D'Athwe´ve already licensed MI but as you mentioned, only Web Apps are supported here but I´ve to deal with the native RDP application from windows
