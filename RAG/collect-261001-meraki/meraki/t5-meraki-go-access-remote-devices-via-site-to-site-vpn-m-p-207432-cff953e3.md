---
id: collect-261001-meraki/meraki/t5-meraki-go-access-remote-devices-via-site-to-site-vpn-m-p-207432-cff953e3
title: "t5-meraki-go-access-remote-devices-via-site-to-site-vpn-m-p-207432-cff953e3"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-meraki-go-access-remote-devices-via-site-to-site-vpn-m-p-207432-cff953e3.md
source_anchor: ""
source_lines: [1, 140]
sha256: 066c5b2bf3c5dba1cd0e561b23f6b16a6d005441caa999bd47c7feb1c420ab52
---

# t5-meraki-go-access-remote-devices-via-site-to-site-vpn-m-p-207432-cff953e3

Access Remote Devices via Site to Site VPN
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:33 AM
I'm unable to access remote devices after setting up site to site with 2 GX50's. Status shows the VPN active on both GX50's, but IP/UNC paths not working. What am I missing?
- Labels:
- 
						
							
		
			Meraki Go
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:37 AM
Do you use different ip-address-ranges on both sites?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:38 AM
Yes. LAN ranges on device 1 is 192.168.254.0/24 and device 2 is 192.168.1.0/24
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:44 AM
I'm concerned that DNS is an issue, even with directly trying to use the remote subnet IP isn't working also. Should the default VLAN 1 use custom DNS instead of Upstream DNS?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:44 AM
Are you able to ping each GX50 from the other site?
If not, did you reboot both GX50 and renew the VPN connection?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:48 AM
I'm able to ping each WAN remote IP across sites
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:52 AM
WAN address or local IP(192.168.254.1 or 192.168.1.1)?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 07:57 AM
I can ping the 192.168.1.1 from the 192.168.254.0/24 network, but not vice versa. Both WAN addresses are pingable from each network. I'm also able to bring up the Meraki Go page of 192.168.1.1 on the remote network.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 08:00 AM
Looks like there's something wrong with the VPN-tunnel, you should be able to ping 192.168.1.1 from 192.168.254.0/24 network and the other way around.
Did you follow this guide Site to Site VPN with Meraki Go Router Firewalls - Cisco Meraki to setup the VPN?
Did you open a support-ticket?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 08:04 AM
I did, but was hoping for a quicker option.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 08:15 AM
not the type of info I like to read on a documentation.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 08:19 AM
Well, that’s not good.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-07-2023 08:20 AM
Guess I’ll see what support has to say.
