---
id: collect-261001-meraki/meraki/t5-security-sd-wan-meraki-mx85-how-to-configure-to-limit-inbound-1-1-and-1-many-85cc9cb9
title: "t5-security-sd-wan-meraki-mx85-how-to-configure-to-limit-inbound-1-1-and-1-many--85cc9cb9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-security-sd-wan-meraki-mx85-how-to-configure-to-limit-inbound-1-1-and-1-many--85cc9cb9.md
source_anchor: ""
source_lines: [1, 95]
sha256: 0102448b0843623c06a005e9567d50fb4508e96b1b18230e5dffe90c86d4cf92
---

# t5-security-sd-wan-meraki-mx85-how-to-configure-to-limit-inbound-1-1-and-1-many--85cc9cb9

Meraki MX85 - How to configure to limit inbound 1:1 and 1:Many NAT with GEO Location
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2025 02:26 AM
How to configure to limit inbound 1:1 and 1:Many NAT with GEO Location ?
As i understand, the layer 7 geo location blocking does not inspect inbound NAT traffic. With this how we can configure to perform above or is this a limitation in the MX ?
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
07-28-2025 02:52 AM
From my understanding it's not possible unless you use NAT Exemptions and L3 firewall rules with CIDR's you want to block.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-28-2025 03:48 AM
Meraki’s Layer 7 firewall and GEO-IP filtering are designed primarily for outbound traffic from LAN to WAN. When traffic comes from the internet into your network via NAT, it bypasses those content-aware inspection engines.
In short, you need another device in front of the MX to perform this function given this limitation of the MX.
https://www.f5.com/products/big-ip-services/advanced-waf
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-29-2025 06:10 AM
Then explain me this: 
Why do the L7 firewall rules say traffic TO/FROM .. countries..
I can't test this but if you block a country it should be both ways if you read the configuration correctly.
Having to need another device in front of the MX to do a function the MX itself should do is a bit redundant 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-29-2025 08:12 AM
I partially agree; the problem is that you think a UTM has to do everything, but that's a misconception, to say the least.
Having specific equipment for this function is the most appropriate.
The truth is that the MX leaves nothing to be desired in this regard.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-29-2025 12:46 PM
Country blocking for both traffic passing through the firewall or limiting a service running on the UTM box should both be done by the box itself.  These are essential basic functions of an UTM.  Having to put another box in front of an UTM just to block it is quite alien to me.
I do get the usecase for SASE, SSE where the truly advanced stuff is done by the cloud delivered security and have the MX be more of an advanced edge router.  However the MX should be able to stand on it's own for this one simple feature.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-29-2025 02:33 PM
Okay, that's your point of view, and I respect it, but I still believe that MX can't handle it well (I speak from experience, so it's not an unsubstantiated statement). I hope you don't misunderstand me. Best regards.
Please, if this post was useful, leave your kudos and mark it as solved.
