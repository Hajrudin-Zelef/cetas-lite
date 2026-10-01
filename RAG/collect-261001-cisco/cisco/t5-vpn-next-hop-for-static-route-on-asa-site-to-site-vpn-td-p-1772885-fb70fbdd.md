---
id: collect-261001-cisco/cisco/t5-vpn-next-hop-for-static-route-on-asa-site-to-site-vpn-td-p-1772885-fb70fbdd
title: "t5-vpn-next-hop-for-static-route-on-asa-site-to-site-vpn-td-p-1772885-fb70fbdd"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-vpn-next-hop-for-static-route-on-asa-site-to-site-vpn-td-p-1772885-fb70fbdd.md
source_anchor: ""
source_lines: [1, 87]
sha256: 824d6241b813c5f51c2fcf34e902d79d9bc4486010694c9148e59eb7875a3e2b
---

# t5-vpn-next-hop-for-static-route-on-asa-site-to-site-vpn-td-p-1772885-fb70fbdd

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-09-2011 04:34 AM
Hi All,
I would be grateful if someone could help me with my ASA issue/misunderstanding. I have a site to site VPN on an ASA. I want to add a floating static route to point to the VPN on this ASA. Note the traffic via this route is not with in the crypto ACL subnets which is used to bring up the VPN. This VPN is used only as a backup.
Do I add the static route with the next hop as the local public address or the remote public address of the VPN? Can the next hop be the local ASA interface facing the internet isp? I plan to do this on the ASDM. I'm sorry if this is a simple question but I've not found material that explains this?
Regards
Solved! Go to Solution.
- Labels:
- 
						
							
		
			VPN
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-15-2011 04:47 AM
Ahh, ok, makes sense.
The next hop needs to be the next hop of the interface that terminates the VPN connection, essentially the same as your Internet/outside interface next hop.
Example topology:
Site B (outside interface - 1.1.1.1) - (next hop: 1.1.1.2) Internet
The static route should say:
route outside 10.2.2.2 255.255.255.255 1.1.1.2 200
Hope this helps.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-13-2011 03:10 AM
Not quite sure if I understand your question.
If the floating static route that you are trying to configure is not configured in the crypto ACL subnet, then what are you trying to achieve with the floating static route?
This subnet will never be routed towards the VPN even in the event of the floating static route is triggered if it's not in the crypto ACL as traffic which is not in the crypto ACl will not be encrypted/decrypted.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-15-2011 01:01 AM
Hi Jennifer,
Many thanks for your reply, let me explain a little more. The crypto ACL has includes a large range, lets say 10.1.0.0/8 is site A and 10.2.0.0/8 is site B
On Site B, ASA, I want to add a static to point to a host on site A, ( route 10.2.2.2 255.255.255.255 to VPN tunnel admin distance 200). The VPN tunnel only comes up as a backup. I'm trying to find out what next hop I needsto point the static route to? Is it the IP address of the public side of site A, IP address of public side to site B?
Regards
D.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-15-2011 04:47 AM
Ahh, ok, makes sense.
The next hop needs to be the next hop of the interface that terminates the VPN connection, essentially the same as your Internet/outside interface next hop.
Example topology:
Site B (outside interface - 1.1.1.1) - (next hop: 1.1.1.2) Internet
The static route should say:
route outside 10.2.2.2 255.255.255.255 1.1.1.2 200
Hope this helps.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-15-2011 09:06 AM
Hi Jennifer,
Many thanks for that, I'll therefore use the internet service providers next hop address for the static route.
Many thanks.
Regards
Dinesh
