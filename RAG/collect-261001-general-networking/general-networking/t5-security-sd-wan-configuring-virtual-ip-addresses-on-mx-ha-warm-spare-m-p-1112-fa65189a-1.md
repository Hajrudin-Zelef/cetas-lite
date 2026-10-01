---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a-1
title: "t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a.md
source_anchor: ""
source_lines: [1, 33]
sha256: f35b51b6c83e148dbc5d6f74ee1e1f4d909d576c1d0a1daad0b13c0d79086c4a
---

# t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a

Configuring virtual ip addresses on MX HA warm spare
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-08-2021 02:23 PM
If we want to configure virtual ip address for HA do we configure it on both the WAN interfaces and LAN interfaces of each MX or only the WAN interface?
I'm looking at this configuration document and only see steps to configure a virtual ip address for the WAN interface.
https://documentation.meraki.com/MX/Deployment_Guides/MX_Warm_Spare_-_High_Availability_Pair
- How would this work for the LAN? We have a default route configured on the local L3 switch and do not want to have to change that each time a failover occurs.
- And can we use private ip addresses for the virtual ip addresses?
Thanks.
- Labels:
- 
						
							
		
