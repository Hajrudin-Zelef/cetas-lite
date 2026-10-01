---
id: collect-261001-general-networking/general-networking/t5-wireless-wlc-dhcp-proxy-mode-and-dhcp-snooping-on-upstream-switch-td-p-492875-e4d10d76
title: "t5-wireless-wlc-dhcp-proxy-mode-and-dhcp-snooping-on-upstream-switch-td-p-492875-e4d10d76"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-wireless-wlc-dhcp-proxy-mode-and-dhcp-snooping-on-upstream-switch-td-p-492875-e4d10d76.md
source_anchor: ""
source_lines: [1, 120]
sha256: 01c0e4e783ef1fd5eb46ebcd3bdeeb80acc435496055480e5ed31da14a52e41f
---

# t5-wireless-wlc-dhcp-proxy-mode-and-dhcp-snooping-on-upstream-switch-td-p-492875-e4d10d76

WLC DHCP Proxy mode and DHCP Snooping on upstream switch
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-25-2023 05:20 AM
Hello,
is there an issue when we have WLC DHCP proxy mode and upstream switch with DHCP snooping enabled?
Based on docs, WLC in proxy mode changes giaddr field (and can insert option-82 as well) and switch ignores DHCP messages over untrusted ports if it has non-zero giaddr field or option-82 (like relay info inserted).
Then, it should be problematic for DHCP snooping enabled environment, right? We need to make trust WLC connected ports (which disables snooping checks for those ports, in reality) or configure L2 ports as ip dhcp relay trusted. Did anyone had
Please rate and mark as an accepted solution if you have found any of the information provided useful.
- Labels:
- 
						
							
		
			Catalyst 9000
- 
						
							
		
			Wireless LAN Controller
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-25-2023 05:53 AM - edited 09-25-2023 05:54 AM
snooping dictates where offer comes from not where discover comes from, so dont think this should be an issue
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-25-2023 06:50 AM
No, snooping has some checks for client messages as well.
For example, when you have access and distro switch with both snooping enabled where access inserts option82, then distro switch ignores client messages. We normally either remove option82 on access OR allow it on untrusted port on distro switch.
I assume then same happens in WLC, but can not get confirmation since I dont have WLC Lab
Please rate and mark as an accepted solution if you have found any of the information provided useful.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-25-2023 07:16 AM - edited 09-25-2023 07:18 AM
you are right you have to have ip dhcp snooping information option allow-untrusted.
by default its disabled.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2023 03:58 PM
DHCP proxy only applies to the old AireOS based WLCs which are almost end of life. If you are designing for future then you should be looking at the 9800 series WLCs.
If you use 9800 series WLC as per the Best Practice guide (link below) then you should not configure SVI on the 9800 at all and leave the snooping/forwarding/relaying to the attached infrastructure. If you do configure SVI with helper address/dhcp relay then it will be doing standards based DHCP relay not DHCP proxy.
Please click Helpful if this post helped you and Accept as Solution if this answered your query.
------------------------------
TAC recommended codes for AireOS WLC's and TAC recommended codes for 9800 WLC's
Best Practices for AireOS WLC's, Best Practices for 9800 WLC's and Cisco Wireless compatibility matrix
Check your 9800 WLC config with Wireless Config Analyzer using "show tech wireless" output or "config paging disable" then "show run-config" output on AireOS and use Wireless Debug Analyzer to analyze your WLC client debugs
Field Notice: FN63942 APs and WLCs Fail to Create CAPWAP Connections Due to Certificate Expiration
Field Notice: FN72424 Later Versions of WiFi 6 APs Fail to Join WLC - Software Upgrade Required
Field Notice: FN72524 IOS APs stuck in downloading state after 4 Dec 2022 due to Certificate Expired
- Fixed in 8.10.196.0, latest 9800 releases, 8.5.182.12 (8.5.182.13 for 3504) and 8.5.182.109 (IRCM, 8.5.182.111 for 3504)
Field Notice: FN70479 AP Fails to Join or Joins with 1 Radio due to Country Mismatch, RMA needed
Field Notice: FN74383 APs Running 17.12.4/5/6/6a May Run Out of Flash Space Preventing Upgrades
How to avoid boot loop due to corrupted image on Wave 2 and Catalyst 11ax Access Points (CSCvx32806)
Field Notice: FN74035 - Wave2 APs DFS May Not Detect Radar After Channel Availability Check Time
Leo's list of bugs affecting 2800/3800/4800/1560 APs
Default AP console baud rate from 17.12.x is 115200 - introduced by CSCwe88390
AP supported channel lookup: https://apchannels.cisco.com/
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-02-2023 11:42 AM
Thank you, but since it is DHCP relay then giaddr will be modified and infrastructure dhcp snooping enabled switch will ignore these messages over untrusted.
Seems, if it is not bridge mode then ip dhcp snooping trust is needed
Please rate and mark as an accepted solution if you have found any of the information provided useful.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-02-2023 12:01 PM
Ip dhcp snooping trust toward wlc is not needed since the wlc is represent client here.
The modify of dhcp and add op82 is need I think.
Now
Wlc add op82 send to SW (with dhcp snooping) what you need is
Ip dhcp snooping information option allow-untrusted
Why untrust ? Since the port is untrust and wlc add op82 then this need.
