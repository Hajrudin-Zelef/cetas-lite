---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-direct-vlan-configuration-on-mx250-m-p-11525-highlight-true-e62405c5
title: "t5-security-sd-wan-direct-vlan-configuration-on-mx250-m-p-11525-highlight-true-e62405c5"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-direct-vlan-configuration-on-mx250-m-p-11525-highlight-true-e62405c5.md
source_anchor: ""
source_lines: [1, 66]
sha256: 9dd03fa391ea754bb5bbda5d0e5e1c2392b0a90360996fdb61f06053b27c46b3
---

# t5-security-sd-wan-direct-vlan-configuration-on-mx250-m-p-11525-highlight-true-e62405c5

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-07-2018 12:28 AM
We are refreshing our network with Meraki gear and I am in the planning stages of moving over all the firewall rules, VLAN information, etc. to a MX250. We have a point to point vlan connection to a branch office that is local to our corporate headquarters. On our current WatchGuard XTM 5 the port is configured as a trusted subnet. I am assuming on the MX250 I just need to configure the branch office subnet vlan along with the others (in Addressing & VLANs) and in the section below that (per Port VLAN configuration) configure a port for the branch office vlan?
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
02-07-2018 05:05 AM
@OSPF71Can you provide more information as to the point to point. Is is a layer 2 through the PTP private link to your branch office? everything is flat from MX/watchguard to that office or is there routing being done between the sites? If it is L2 like your post is semi-alluding to, your on the right path, if you have L3 your going to need some routes on the MX250. I also assume this site comes back through that private link to get to the internet/services?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-07-2018 05:05 AM
@OSPF71Can you provide more information as to the point to point. Is is a layer 2 through the PTP private link to your branch office? everything is flat from MX/watchguard to that office or is there routing being done between the sites? If it is L2 like your post is semi-alluding to, your on the right path, if you have L3 your going to need some routes on the MX250. I also assume this site comes back through that private link to get to the internet/services?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-07-2018 06:38 AM
Hi @Dcooper. The Watchguard has the branch office VLAN configured on one of the interfaces and is directly connected to the branch via a Dell switch. After looking at the network information on the WG again under "Routes" I do see static routes for the subnets at the branch office (user, phone vlan's). Would I need to add those static routes under "Add a Static Route" on the Addressing & VLAN's section of the MX250?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-07-2018 06:45 AM
Yes sir, your correct on adding the routes. Copy them over.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-07-2018 08:09 AM
So to summarize...under Addressing & VLAN's-Routing:
-Configure branch office VLAN/subnet
-Add static routes for user & phone subnets at branch
-Configure Per port VLAN for branch office ethernet handoff
Thanks again for clarifying!
