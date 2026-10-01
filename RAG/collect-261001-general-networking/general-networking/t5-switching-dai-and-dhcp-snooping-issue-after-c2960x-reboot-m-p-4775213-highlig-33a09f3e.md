---
id: collect-261001-general-networking/general-networking/t5-switching-dai-and-dhcp-snooping-issue-after-c2960x-reboot-m-p-4775213-highlig-33a09f3e
title: "t5-switching-dai-and-dhcp-snooping-issue-after-c2960x-reboot-m-p-4775213-highlig-33a09f3e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/t5-switching-dai-and-dhcp-snooping-issue-after-c2960x-reboot-m-p-4775213-highlig-33a09f3e.md
source_anchor: ""
source_lines: [1, 159]
sha256: 300effcd36af627c48edc4da599b4defc0c649cb7a5156e306a2c2db30d5c1d4
---

# t5-switching-dai-and-dhcp-snooping-issue-after-c2960x-reboot-m-p-4775213-highlig-33a09f3e

DAI and DHCP snooping issue after C2960X reboot
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 04:03 AM
Hi all,
we encountered a dhcp snooping issue after a reboot on a C2960X stack. After a reboot, some of the IP Phones would not register, unless a shut/no shut was issued in the access port. The logs on the switch were numerous:
%SW_DAI-4-DHCP_SNOOPING_DENY: 1 Invalid ARPs (Req) on Gi4/0/23, vlan 4 [xxxx.xxxx.xxxx/10.10.4.24/0000.0000.0000/10.10.4.1
It affected the majority of the 7800 series IP Phones, not the 3905.
The config of the switch is:
ip arp inspection vlan 4
ip dhcp snooping vlan 4
ip dhcp snooping
The config of the access port has the commands:
ip arp inspection limit rate 100 and ip dhcp snooping limit rate 100
From my understanding, there was a mismatch between the ARP request and the DHCP snooping database, but I cannot understand the reason. After the reboot, shouldn't the DHCP snooping database be empty and no mismatch occur?
Also note that on our environment we have deployed many C2960X switches and we haven't faced again such an issue after reboot.
Thank you,
Alex
- Labels:
- 
						
							
		
			LAN Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 04:14 AM
I think it bug check below bug
DHCP SNOOPING DATABASE IS NOT REFRESHED AFTER RELOAD
CSCvp49518 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 04:19 AM
The model and version is not the same (we have C2960X with 15.2(7)E5) but the behavior is similar. Maybe we should open a TAC for further investigation
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 04:34 AM
I know it not same model and version but since there is bug with similar behaviour then you can face same bug to. 
open TAC is very good 
because it sure confirm that it bug or not. 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 04:53 AM
If this is your complete DHCP snooping configuration, you have no database configuration configured. You should have this to make sure the switch starts with a populated database after a reboot. Not having the database location would result in your symptoms.
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 05:18 AM
Hi Karsten,
Thank's for your response. What do you mean by no database configuration?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 05:29 AM
he mean you specify other location for database like URL. 
if you not  config other location then it save in your SW memory and must be clear after reboot (like you mention what happened for other SW). but I think because bug this database not clear. 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 05:34 AM
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 05:44 AM
Ok, that's helpful and we might use it as workaround, however the default location shouldn't be the switch and not an external location?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 05:54 AM
it copy, the default is in SW and SW will writing the DB to external.
ip dhcp snooping database write-delay xx <<- this time is DB delay writing. 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 05:55 AM
It's not a workaround. It's part of a DHCP-Snooping implementation. A stable remote location is certainly the best, but often I use the local flash because there I am not dependant on external systems. Constantly writing (ok, there is the write-delay) onto local flash is nothing I really like, but up to now the systems didn't complain even after some years in operation.
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 06:03 AM
I will try to reproduce the issue and use this configuration. We have also opened a TAC to examine whether this is a bug. Anyway, I will update this thread with the results. Thank you both for the very helpful support!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-14-2023 06:04 AM
You are so so welcome
