---
id: collect-261001-general-networking/general-networking/t5-wireless-mandatory-dhcp-m-p-102573-57866f4f-3
title: "t5-wireless-mandatory-dhcp-m-p-102573-57866f4f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-wireless-mandatory-dhcp-m-p-102573-57866f4f.md
source_anchor: ""
source_lines: [163, 183]
sha256: 30fa42475f2b0536bee0848099f29b125104d088c0daf538071170a1b2473317
---

# t5-wireless-mandatory-dhcp-m-p-102573-57866f4f

- LAN connection blocked from wired to wireless completely i.e. I couldn’t see the ICMP echo from the wireless client on the wired capture
So from this I think we can say that mandatory DHCP can be enabled but it is reliant on certain client behaviours. e.g. roaming with a reassociation frame (which then permits wireless->wired traffic) and also refreshing DHCP when the roam is complete (which then permits wired->wireless traffic).
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2020 06:11 AM
Testing was with a pair of MR42s by the way.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-01-2020 12:18 AM
The Meraki documentation has now been updated on this feature to reflect the fact that it only works when roaming if the client refreshes its DHCP when the roam is complete.
https://documentation.meraki.com/MR/Access_Control#Mandatory_DHCP
