---
id: collect-261001-meraki/meraki/t5-cloud-security-sd-wan-vmx-vmx100-non-meraki-peer-to-auto-vpn-site-m-p-109064-1eacdfe2
title: "t5-cloud-security-sd-wan-vmx-vmx100-non-meraki-peer-to-auto-vpn-site-m-p-109064--1eacdfe2"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-cloud-security-sd-wan-vmx-vmx100-non-meraki-peer-to-auto-vpn-site-m-p-109064--1eacdfe2.md
source_anchor: ""
source_lines: [1, 87]
sha256: 71fa8a15733d0ba7fc7107e0c05e3e0045a10d8a0d42097936d774d719ced47d
---

# t5-cloud-security-sd-wan-vmx-vmx100-non-meraki-peer-to-auto-vpn-site-m-p-109064--1eacdfe2

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-10-2021 02:53 PM
Trying to get traffic from a non-Meraki network to an Auto VPN network....that goes across an Azure vMX. The vMX has routes to both networks....1.1.1.X and 2.2.2.X and both VPNs are online.
1.1.1.X-----Non-Meraki Peer---VPN---vMX100---Auto VPN---MX84---2.2.2.X
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
02-11-2021 10:44 AM
It would not make any difference. Both can do local routing fine.
What you are asking for is VPN hair pinning between a non-Meraki VPN and a Meraki VPN - and they don't do that.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-10-2021 03:25 PM
That does not work. You need also the 3rd party vpn to the mx84
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-10-2021 03:45 PM
So, the vMX100 will not route between the two tunnels? So, I will need to create site-to-site tunnels from every Auto VPN spoke to the 3rd party VPN?
Doesn't that waste the whole hub-spoke benefit of Meraki? It seems creating one VPN to the 3rd party from the vMX100 would be a lot simpler than hundreds of 3rd party VPNs from the MX84s.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-10-2021 11:50 PM
>It seems creating one VPN to the 3rd party from the vMX100 would be a lot simpler than hundreds of 3rd party VPNs from the MX84s.
if you want to do that then terminate the third party VPN on the Azure VPN gateway for Strongswan on Ubuntu, and then just put routes between the two systems. You can include static routes into AutoVPN.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2021 03:57 AM
Thanks! One final question... Would a similar workaround be needed if the headend was a MX450 instead of a vMX100...physcial device vs virtual?
I have been routing since the Novell days, including 6 years at Cisco HTTS/TAC, and I am trying to wrap my head around why the vMX100 cannot route between two subnets in its routing table.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-11-2021 10:44 AM
It would not make any difference. Both can do local routing fine.
What you are asking for is VPN hair pinning between a non-Meraki VPN and a Meraki VPN - and they don't do that.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-13-2021 01:54 PM
Ok....thanks! I will "make a wish "to make VPN hair pinning possible, especially from the headends.
I did it on a PIX 520 probably 20 years ago now when VPNs first came out.
