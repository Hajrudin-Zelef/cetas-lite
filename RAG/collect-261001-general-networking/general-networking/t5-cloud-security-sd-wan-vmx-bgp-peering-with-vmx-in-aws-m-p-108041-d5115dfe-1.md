---
id: collect-261001-general-networking/general-networking/t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe-1
title: "t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe.md
source_anchor: ""
source_lines: [1, 21]
sha256: 8371ab4455c3e07bbb38e52648f4d2fee1057f7619457811e2babbc760375170
---

# t5-cloud-security-sd-wan-vmx-bgp-peering-with-vmx-in-aws-m-p-108041-d5115dfe

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-28-2021 07:20 AM
There was another post about BGP with the MX that briefly talked about a cloud deployment.
https://community.meraki.com/t5/Security-SD-WAN/BGP-Configuration-on-MX/td-p/28710
We have 2 physical data centers that we are decommissioning and moving everything into AWS. In both of those we have BGP peering enable and working great.
We have spun up 2 vMXs in AWS. I am about to enable those as VPN Hubs this evening, but I am concerned that they will not route correctly.
I have been looking for documentation about how to do BGP peering in AWS. The only documentation I have been able to find pertaining to BGP in AWS deals with Direct Connects, which is not what I need.
Has anyone gotten the vMX to work in AWS as a One Armed Concentrator for VPN termination using BGP?
If not using BGP, how did you deploy it so that it would route traffic to the internal networks, but still be accessible from the outside?
Solved! Go to Solution.
- Labels:
- 
						
							
		
