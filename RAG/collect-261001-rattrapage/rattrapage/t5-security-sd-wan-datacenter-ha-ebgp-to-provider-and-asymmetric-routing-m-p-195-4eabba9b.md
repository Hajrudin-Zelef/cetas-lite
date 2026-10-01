---
id: collect-261001-rattrapage/rattrapage/t5-security-sd-wan-datacenter-ha-ebgp-to-provider-and-asymmetric-routing-m-p-195-4eabba9b
title: "t5-security-sd-wan-datacenter-ha-ebgp-to-provider-and-asymmetric-routing-m-p-195-4eabba9b"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/t5-security-sd-wan-datacenter-ha-ebgp-to-provider-and-asymmetric-routing-m-p-195-4eabba9b.md
source_anchor: ""
source_lines: [1, 60]
sha256: ee7bb4293b12af6e4047a9f16c12023a02005321b0ff80a1151cfdbb9297b7e0
---

# t5-security-sd-wan-datacenter-ha-ebgp-to-provider-and-asymmetric-routing-m-p-195-4eabba9b

Datacenter HA, eBGP to Provider and asymmetric routing
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-24-2023 01:56 AM
Hi,
I have two datacenter hubs and an One-Armed-MX setup in every DC. From every DC there is an eBGP session to another provider/AS. Is there an issue with asymmetric routing, if traffic goes from DC1 to the eBGP nexthop connected to DC1 and the provider routes the reverse traffic to DC2? If so, how can I resolve this?
Best regards,
Rene
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
05-24-2023 02:34 AM
I think, I found the point: routes exported to the AS65001 via the secondary DC will have an additional AS65002 prepended so the incoming (reverse) traffic comes via the primary DC.
I'm still in the design phase and hope, this works as documented
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-24-2023 03:39 AM
Correct - AS prepend ensures that the Hub preference for a given spoke is reflected in that spoke's subnet advertisements from the configured Hubs. This works fine, from my previous testing.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-24-2023 10:39 AM
In the PVT today they mentioned a list of hidden features which support can enable. One of them was disabling peering between hubs. If you are receiving the same routes in both dcs it is recommended to enable this feature to avoid asymmetric routing.
