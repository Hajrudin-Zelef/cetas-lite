---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-native-vlan-configuration-td-p-235627-8a977959
title: "t5-security-sd-wan-native-vlan-configuration-td-p-235627-8a977959"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-native-vlan-configuration-td-p-235627-8a977959.md
source_anchor: ""
source_lines: [1, 58]
sha256: 2871bb63bd9998442345f90f14db236cc612b100657a5da0fa2aaa3d32e93daa
---

# t5-security-sd-wan-native-vlan-configuration-td-p-235627-8a977959

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-14-2024 10:46 AM
I've seen articles about not configuring native VLAN for security purposes.
I have a questions as to how to do so in terms of the MX appliances, can anyone provide guidance?
Should I use drop untagged traffic, VLAN 50 for management or create a VLAN that I won't use for this purpose?
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-14-2024 10:56 AM
I believe this is not the meaning, generally avoiding using VLAN 1, and only allowing VLANs that will actually be used in the Trunk, but you do not need to avoid the native VLAN.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-14-2024 10:56 AM
I believe this is not the meaning, generally avoiding using VLAN 1, and only allowing VLANs that will actually be used in the Trunk, but you do not need to avoid the native VLAN.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-14-2024 11:14 AM
So avoiding the use of the default VLAN 1 and using a created VLAN for the Native VLAN would be fine?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
05-14-2024 11:36 AM
Yes, it doesn't mean you'll have a super secure network, these are just best practices, but security goes far beyond that.
Please, if this post was useful, leave your kudos and mark it as solved.
