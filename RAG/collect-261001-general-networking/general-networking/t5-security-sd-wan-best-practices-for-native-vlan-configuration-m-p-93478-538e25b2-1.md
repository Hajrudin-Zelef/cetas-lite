---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2-1
title: "t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2.md
source_anchor: ""
source_lines: [1, 20]
sha256: 1d78f5e7352fbec716a6a0685e16c8529435d59afafb5ca802d24ec7e83de729
---

# t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-23-2019 02:29 AM
Hi there,
- A good security practice is to separate management and user data traffic. The management VLAN, which is VLAN 1 by default, should be changed to a separate, distinct VLAN.
- A recommended security practice is to change the native VLAN to a different VLAN than VLAN 1. The native VLAN should also be distinct from all user VLANs.
I would like to know what are the best practices which you usually implement in the Meraki world. Could you please share some insights? I am basically planning to use a random number for the native VLAN in a new environment which is going to be deployed in some days.
Thanks!
Federico
Solved! Go to Solution.
- Labels:
- 
						
							
		
