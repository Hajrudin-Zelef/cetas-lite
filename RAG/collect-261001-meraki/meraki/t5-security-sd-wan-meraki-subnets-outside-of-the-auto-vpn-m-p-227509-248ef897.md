---
id: collect-261001-meraki/meraki/t5-security-sd-wan-meraki-subnets-outside-of-the-auto-vpn-m-p-227509-248ef897
title: "t5-security-sd-wan-meraki-subnets-outside-of-the-auto-vpn-m-p-227509-248ef897"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-security-sd-wan-meraki-subnets-outside-of-the-auto-vpn-m-p-227509-248ef897.md
source_anchor: ""
source_lines: [1, 53]
sha256: 93c0f2cafdb4bf2030ef018ce49f9a485131fd36f14140dee58c1384a5eff750
---

# t5-security-sd-wan-meraki-subnets-outside-of-the-auto-vpn-m-p-227509-248ef897

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-03-2024 07:01 PM
Hi,
Can non VPN subnets, eg a guest VLAN, overlap across multiple MX appliances?
Say I have 3-4 sites that have a guest VLAN that is not in the VPN, use the same subnet.
I dont see why not, as I already a few like that, but thought I'd ask if there are any issues/limitations.
Thanks.
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
03-03-2024 11:25 PM
Yes you can overlap those subnets but you will not be able to include them in auto VPN.
The sanity check dashboard does is to see if that subnet would appear in the routing table of another AutoVPN participant.
You could notice if you clone a network it does have the same networks but the site-2-site feature is disabled then to avoid this scenario.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-03-2024 11:25 PM
Yes you can overlap those subnets but you will not be able to include them in auto VPN.
The sanity check dashboard does is to see if that subnet would appear in the routing table of another AutoVPN participant.
You could notice if you clone a network it does have the same networks but the site-2-site feature is disabled then to avoid this scenario.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-04-2024 06:52 AM
You can also enable Site-to-site VPN Translation (VPN NAT) if you want to add them in the future:
https://documentation.meraki.com/MX/Site-to-site_VPN/Using_Site-to-site_VPN_Translation
