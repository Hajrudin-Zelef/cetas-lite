---
id: collect-261001-cisco/cisco/t5-cisco-one-discussions-troubleshooting-vlan-connectivity-issue-on-cisco-switch-4ceac22c
title: "t5-cisco-one-discussions-troubleshooting-vlan-connectivity-issue-on-cisco-switch-4ceac22c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-cisco-one-discussions-troubleshooting-vlan-connectivity-issue-on-cisco-switch-4ceac22c.md
source_anchor: ""
source_lines: [1, 90]
sha256: 4a383a064e8f5578b85905f8e91ebf8f5c44c67fd2cd426cc7a711fc5d45f51b
---

# t5-cisco-one-discussions-troubleshooting-vlan-connectivity-issue-on-cisco-switch-4ceac22c

Troubleshooting VLAN Connectivity Issue on Cisco Switch
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-17-2023 02:12 AM - edited 12-25-2023 10:58 PM
I'm currently facing a VLAN connectivity issue on my Cisco switch. I've configured the VLANs, but devices in different VLANs can't seem to communicate. I've checked the configurations, and everything seems fine. What steps can I take to troubleshoot and resolve this VLAN connectivity issue?
My friend told me through whatsapp apk that he's also facing the same issue.
- Labels:
- 
						
							
		
			Cisco ONE
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-17-2023 02:44 AM
You need
Config SVI for each vlan
Make SW L3 via
Ip routing
MHM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-17-2023 02:47 AM
Hello @JackBDMN
Which switch is it ? If a traditional L2 as C2960, by default inter vlan communication is not possible since It works as Layer 2! You need a L3 equipement to achieve this. Or of you enable "L3 capabilities" on that Switch...depend on the model.
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-17-2023 03:25 AM
There are lot of information missing here for the community to help in best way :
1. Always provide the environment - where other peers can understand the setup "rather generic not work ?"
2. what Device models and what IOS code running /
3. how is these devices conencted each other - Layer 2 running or Layer 3 p2p configuration ?
but devices in different VLANs can't seem to communicate. What VLAN those are give the numbers - from what VLAN you trying to access what VLAN - do you have IP address of that VLANS ?
For VLAN to VLAN (different VLAN) communication, you need Layer 3 routing should take place, what device doing this Layer 3 Routing here ?
I've checked the configurations, and everything seems fine
what basis this was fine, since you have issue, so there may be some configuration missing that where it is not working, hence this post here.
post below information to help better :
Give high level picture of your network
show vlan (from devices)
show ip interface brief
show ip route
show cdp neigh
show spann summary
ahow ip arp
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-17-2023 04:03 PM
What troubleshooting has been done?
NOTE: My response is to elicit a response from the bot scrubber.
