---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-routing-between-auto-vpn-and-mpls-on-mx-m-p-237186-4a8f0e0c
title: "t5-security-sd-wan-routing-between-auto-vpn-and-mpls-on-mx-m-p-237186-4a8f0e0c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-routing-between-auto-vpn-and-mpls-on-mx-m-p-237186-4a8f0e0c.md
source_anchor: ""
source_lines: [1, 50]
sha256: 8b4126dceee4cac28919ed96813730effa9592b67c539fc75e7e99a52594602f
---

# t5-security-sd-wan-routing-between-auto-vpn-and-mpls-on-mx-m-p-237186-4a8f0e0c

Routing between auto VPN and MPLS on MX
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-28-2024 12:17 AM
Can Meraki Hub advertise subnets learned from static route through L3VPN-MPLS to its Spokes and vice versa as below block diagram?
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
05-28-2024 12:56 AM
You can configure the hub MX's to talk BGP to your FG-80F.
https://documentation.meraki.com/MX/Networks_and_Routing/Border_Gateway_Protocol_(BGP)
As long as the FG-80F advertises the static routes, they will be learned.
Note that the BGP functionality is more extensive if the MXs are in VPN concentrator (or "one-armed") mode.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-30-2024 03:03 AM
Can we use both MXs as active HUBs (not with HA) in routed mode and advertise same subnet X through HUBs to spokes to distribute/split spokes (EX, 2 spokes use HQ MX and 2 other spokes use DR MX) and also MXs backup each others in the same time?
