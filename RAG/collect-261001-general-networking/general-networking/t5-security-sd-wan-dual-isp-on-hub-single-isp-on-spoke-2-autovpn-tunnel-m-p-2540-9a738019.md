---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-dual-isp-on-hub-single-isp-on-spoke-2-autovpn-tunnel-m-p-2540-9a738019
title: "t5-security-sd-wan-dual-isp-on-hub-single-isp-on-spoke-2-autovpn-tunnel-m-p-2540-9a738019"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-dual-isp-on-hub-single-isp-on-spoke-2-autovpn-tunnel-m-p-2540-9a738019.md
source_anchor: ""
source_lines: [1, 84]
sha256: f9937a354d3d527eb79dbbee8bb421f1b2f952a4ab271c99624aafc17ba98917
---

# t5-security-sd-wan-dual-isp-on-hub-single-isp-on-spoke-2-autovpn-tunnel-m-p-2540-9a738019

Dual ISP on Hub, single ISP on Spoke-- 2 autoVPN tunnel ?
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-10-2024 11:49 PM
Hello,
I have scenario like on picture below:
So like you see we only have 2 different ISP on Hub/central location but only 1 ISP on spoke location. Is it possible to bring 2 autoVPN tunnel from Spoke Wan1 interface and play with sd-wan policises for vpn traffic? For example from 10 network to 20 go through autovpn tunnel 1, and from 10 to 30 through autovpn tunnel 2 and also one will be redundant to the other. For now i dont see option in sd-wan policy--vpn traffic that we can choose autovpn tunnel, only exit interface (WAN 1 and WAN 2).
Does anyone have experience with this scenario with MXs firewalls?
Best regards,
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
11-11-2024 12:43 AM
If your Hub and spoke locations are running in Routed/NAT mode and you have 2 ISP's at the hub and 1 ISP at the spoke you will essentially have 2 AutoVPN tunnels.  That means you will be able to control the traffic at the Hub but not at the spoke.
If you would have 2 ISP's at both locations you would have 4 tunnels.
As I have found earlier this year the SD-WAN uplink policies have effect in both upstream and downstream direction. So if you match certain traffic at the hub site to take WAN1 it should leave via WAN1 but also arrive at WAN1. This had been tested with simple L3/4 matching criteria.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2024 12:54 AM
Hello,
Thank you for your answer and explanation. Most of our traffic has direction from spokes to Hub. So if i understand you good, if we want that traffic from 10 network to 20 (on topology above from spoke to Hub ) go through autovpn tunnel 1 we need to create sd-wan policy--vpn traffic on Hub in oposite director from 20 to 10 and choose isp1 as exit point? Is there any place on Meraki dashboard that we can confirm this behavior?
Best regards,
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-18-2024 10:54 AM
I can only confirm that this was the behavior I noticed.
However it is important to note that this was matching traffic by regular L3/4 rules, not the L7 application layer detection.
So I cannot guarantee it 100%.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2024 10:55 AM
You can disable Active-Active AutoVPN. Then the tunnel is built with the primary uplink. The secondary uplink is used, if the primary uplink fails.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-17-2024 10:40 AM
I want active-active VPN from spoke with ISP1 and ISP2 on HUB location. I want some traffic from branch sent to the HUB via ISP1, and other traffic from branch sent to the HUB via ISP2 if this is possible.
