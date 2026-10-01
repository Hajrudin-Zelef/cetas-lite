---
id: collect-261001-general-networking/general-networking/t5-switching-hsrp-not-working-td-p-4794966-996dcb41-1
title: "t5-switching-hsrp-not-working-td-p-4794966-996dcb41"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/t5-switching-hsrp-not-working-td-p-4794966-996dcb41.md
source_anchor: ""
source_lines: [1, 16]
sha256: 95cc23a7c3c2b0cb3a982615b8d296c8f11d9233cc91a13ef8d02cf13785b978
---

# t5-switching-hsrp-not-working-td-p-4794966-996dcb41

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-15-2023 12:59 PM
Hello,
I've configured our two Core Switches to be in HSRP for the VLANs on the network. Upstream from the two Core Switches is an ASA that I configured to have a redundant interface so we have redundancy if an upstream interface fails for the Core Switches. The interfaces failing over works. However, when I test to see if HSRP will failover by pulling the power on Core Switch 1, the entire network goes down. Now, we have some physical limitations so not all access switches all connect to the distribution switches. Some access switches connect directly to the Core Switches and other access switches only have one connection to either Distribution Switch 1 or Distribution Switch 2. I configured the Core Switches to be the Root for the STP for all of the VLANs on the network. Core Switch 1 being primary root and Core Switch 2 being secondary root (spanning-tree vlan x priority 24576 and spanning-tree vlan x priority 28672). Am I missing any configurations that's preventing the HSRP from failing over?
Solved! Go to Solution.
- Labels:
- 
						
							
		
