---
id: collect-261001-general-networking/general-networking/t5-wireless-mandatory-dhcp-m-p-102573-57866f4f-1
title: "t5-wireless-mandatory-dhcp-m-p-102573-57866f4f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-wireless-mandatory-dhcp-m-p-102573-57866f4f.md
source_anchor: ""
source_lines: [1, 19]
sha256: 9fcc425d42c8c16cf728409a3526d616678848f0dd15b632abb942e43202b808
---

# t5-wireless-mandatory-dhcp-m-p-102573-57866f4f

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-11-2020 08:26 AM
Hi, has anyone turned on mandatory DHCP on an SSID?
I'm having problems with it.
If a client roams to a different AP to the one it originally connected to, the traffic isn't bridged to the LAN side of the "roamed to" AP. If you then disconnect and reconnect it works because the new AP has witnessed the DHCP exchange.
Meraki support have told me that the client should refresh DHCP on every roam but that's just not feasible surely?
Anyone else had any experiences to share? I have yet to witness ANY client refreshing DHCP on roaming.
Solved! Go to Solution.
- Labels:
- 
						
							
		
