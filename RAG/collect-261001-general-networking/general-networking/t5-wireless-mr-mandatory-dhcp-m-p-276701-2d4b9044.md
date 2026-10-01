---
id: collect-261001-general-networking/general-networking/t5-wireless-mr-mandatory-dhcp-m-p-276701-2d4b9044
title: "t5-wireless-mr-mandatory-dhcp-m-p-276701-2d4b9044"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-wireless-mr-mandatory-dhcp-m-p-276701-2d4b9044.md
source_anchor: ""
source_lines: [1, 154]
sha256: 0154ff41012c24b322a3dc531756caf26e93cce8b817db801539c55db634ba6d
---

# t5-wireless-mr-mandatory-dhcp-m-p-276701-2d4b9044

MR - Mandatory DHCP
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-14-2024 08:37 AM
Hi ,
We had issues with Zebra wireless scanners. After days of troubleshooting we found that article :
and that post :
https://community.meraki.com/t5/Wireless/Mandatory-DHCP/m-p/102573
https://documentation.meraki.com/MR/Access_Control#Mandatory_DHCP
which is not required by AOSP or IEEE standards and Zebra devices do not support this.
Why would Meraki go that way ?
Still running MR28 , does anyone knows if that behavior has changed with MR29 or MR30 ?
I can't test it at the moment.
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
03-14-2024 08:42 AM
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-14-2024 09:54 AM
No that wouldn't make any sense.
Devices are DHCP. but Mandatory DHCP requires a DHCP transaction per roaming , which is not part of the IEEE standards so some devices do not have to support it. Zebra are one of those devices and they won't support it like mentionned in their KB
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-14-2024 09:58 AM
I understand, well I believe that on Meraki's side they are unlikely to make any changes, so the only alternative is to disable the feture itself, or create an exclusive SSID for these devices (I'm not particularly a fan of this approach).
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-14-2024 06:01 PM
Thanks for posting about this. But I'm a little unclear on the context: Is mandatory DHCP a requirement for this SSID/network/org?
And was mandatory DHCP enabled by default on a SSID/network/org? I'm pretty sure that it's always been off by default for me.
Meanwhile, I would recommend three sets of changes to the Meraki team:
1. Update all documentation related to mandatory DHCP
2. Update applicable best practices and CVD-like docs with a note about this
3. Add an 'i' with a  circle hover label near the toggles for this setting to link to the appropriate documentation and not potential compatibility issues with roaming on some devices.
I'm going to make internal notes to keep this feature disabled for VLANs with wireless devices.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-14-2024 07:04 PM
Hi ,
This is not enabled by default. It was a special request from our security teams. This has been working perfectly in 99.9% of our sites expect for the ones which contains Zebra devices ( however this issue can be replicated with other devices )
It is specified in the documentation , but can still bite you ( like it did for me )
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-26-2024 08:50 AM
We've been banging our heads against the wall for the last couple days trying to solve the issue with our Zebra scanners and roaming between APs. Support advised us to disable Mandatory DHCP and it worked, the first thing we said after that is that that setting needs an 'i' info bubble!
I found this in the documentation today, and it's exactly what needs to be in that info bubble.
"Enabled: Wireless clients associated to an AP (either new associations or clients that roamed from another AP) that have not requested a DHCP address are placed in a blocked state and are not able to send any traffic on LAN and WAN."
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-15-2024 01:16 AM
I was "caught" in the exact same situation a few years back. The exact same thing.
I was also quite "surprised" until I, almost by chance, read the documentation, and thought about it.
I mean, in theory, the Meraki APs could exchange the information with the other APs on the site, that a client already had done DHCP once, and the problem would be "solved".
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-15-2024 08:49 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2024 05:42 AM
I know the network i was working on at the time was running MR29 (dont know the exact release).
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-26-2024 08:48 AM
I just found this thread after experiencing the exact same issue with Zebra scanners and our brand new Meraki APs. Meraki support advised us to disable Mandatory DHCP and it fixed the issue right away. We're on MR30.6 firmware, so I can report that it hasn't changed. Extremely happy that we figured out what was causing our problems and how to fix it though.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-07-2025 10:26 AM
Woohoo !
https://documentation.meraki.com/MR/Wi-Fi_Basics_and_Best_Practices/Roaming_Technologies
Mandatory DHCP + Fast roaming (802.11r)
This feature enhances client mobility by enabling support for Fast Roaming on SSIDs where Mandatory DHCP is enforced. Starting with firmware version MR 32.1.X, DHCP state information is now shared across APs, allowing seamless client transitions between APs without reinitiating DHCP requests.
Prior to MR32, client roaming to a new AP was required to send new DHCP request before traffic would be allowed. Even if the client had successfully obtained a DHCP lease from the previous AP, the new AP would block traffic until it saw another DHCP exchange.
RFC compliant !
