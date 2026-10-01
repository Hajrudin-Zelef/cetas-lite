---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-mx-questions-m-p-210180-f5ddb7f0
title: "t5-security-sd-wan-mx-questions-m-p-210180-f5ddb7f0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-mx-questions-m-p-210180-f5ddb7f0.md
source_anchor: ""
source_lines: [1, 65]
sha256: 95df724425a1ab674da5a5626cbd406ed184b421a0da7b2da3569b9429111edd
---

# t5-security-sd-wan-mx-questions-m-p-210180-f5ddb7f0

MX questions
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-28-2023 06:33 AM
Hi All
I cant see any mention of this in the in the documentation pages....Is it possible to use a MX just to act like a normal router connecting to a MPLS PE? So in this scenario there would be no Auto VPN requirement, it would need to advertise its LAN routes into BGP only to the PE.....Cloud connectivity would be done centrally
The other thing is....say we we we have a central hub MX with reachability to the internet, and also to a MPLS network. We have a spoke on each transport only. Is that a design Meraki has no problem with? The S2S traffic has to traverse the hub obviously, but is this something Meraki has no issue with as long as the underlay reachability is there?
thanks
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
09-28-2023 06:46 AM
Maybe this will help you but remember that you still need a license to operate.
https://documentation.meraki.com/MX/Networks_and_Routing/Border_Gateway_Protocol_(BGP)
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-28-2023 02:56 PM
To answer the second part, we have MXs in VPN concentrator mode in our DC with spokes on the internet and over MPLS networks. They all connect using the SD-WAN to the DC's services and to each other.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-02-2023 08:54 AM
Hi, regarding your second question: Meraki Hub can properly communicate spokes even when they do not share a common underlay. AutoVPN in hub & spoke model covers that scenario without nothing special to be configured. You only need Hub connectivity to both underlays. Imagine you have dual underlay spoke and one loses underlay1 at the same time the other loses underlay 2. They could still communicate each other.
Regarding your first question: An MX hub in concentrator mode does not allow you to configure static routes. You could theorically run ebgp in the underlay against the PE (asking Meraki to make it available). However, IMHO it would be a case that does not fit with the purpose of positioning a MX. eBGP in Meraki MX is intended for learning local subnets and propagate AutoVPN learnt routes towards your eBGP local peer and still use AutoVPN (and its embbeded ibgp for wan connectivity)
MX hub in route mode can be configured with lan side static routes. However I guess it does not even support eBGP over lan yet.
Regardless of the hub type, uplinks do not support eBGP. Only AutoVPN runs ibgp among meraki MX peers.
Regards.
