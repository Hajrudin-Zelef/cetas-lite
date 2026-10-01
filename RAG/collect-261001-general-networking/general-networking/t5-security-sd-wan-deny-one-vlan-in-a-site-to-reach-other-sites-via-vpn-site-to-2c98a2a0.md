---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-deny-one-vlan-in-a-site-to-reach-other-sites-via-vpn-site-to-2c98a2a0
title: "t5-security-sd-wan-deny-one-vlan-in-a-site-to-reach-other-sites-via-vpn-site-to--2c98a2a0"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-deny-one-vlan-in-a-site-to-reach-other-sites-via-vpn-site-to--2c98a2a0.md
source_anchor: ""
source_lines: [1, 101]
sha256: 7eebe1d4d4026785999fb3921f3a35e702af8b67a27b31e1f57039fa91d2c079
---

# t5-security-sd-wan-deny-one-vlan-in-a-site-to-reach-other-sites-via-vpn-site-to--2c98a2a0

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 01:58 AM
is there any way to deny one VLAN inside a site to reach the network in other sites via vpn site to site because the auto vpn is always on
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
09-06-2024 02:22 AM
Yes - use the site-to-site outbound firewall to create a deny policy matching the subnets.
Site-to-site VPN Firewall Rule Behavior - Cisco Meraki Documentation
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 02:22 AM
Yes - use the site-to-site outbound firewall to create a deny policy matching the subnets.
Site-to-site VPN Firewall Rule Behavior - Cisco Meraki Documentation
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 02:25 AM
in the Site-to-site outbound firewall if i choose the subnet and deny all, it still can connect inside network and internet and stop connecting to other sites ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 02:26 AM
That's right, it only affects VPN traffic. You'd have to create Layer 3 outbound rules (on the Firewall page) if you wanted to restrict traffic to another on-site VLAN or the internet.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 02:27 AM
thank you , what i want is to deny access to other sites so i shoose my subnet and i deny all
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 02:28 AM
If the VLAN requires absolutely no site-to-site traffic at all then the easier solution would to just stop advertising it to other sites by setting VPN Mode to Disabled...
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 02:29 AM
i would like to reach this vlan from outside site but i dont want this vlan to reach other sites
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 02:57 AM
As @JamesT91 said, if you don't want the VLAN to communicate to any sites across the AutoVPN, it's easiest to disable it for that vlan under Sd-WAN -> Site to Site VPN settings.
Otherwise if it's just specific sites, the site to site outbound firewall rules need to be used.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-06-2024 04:11 AM
i would like to reach this vlan from outside site but i dont want this vlan to reach other sites so i think deny source this vlan destination all
