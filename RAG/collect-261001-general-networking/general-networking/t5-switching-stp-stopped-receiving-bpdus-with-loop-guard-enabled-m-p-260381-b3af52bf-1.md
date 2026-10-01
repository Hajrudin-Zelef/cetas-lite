---
id: collect-261001-general-networking/general-networking/t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf-1
title: "t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition"]
source: docs/RAG/collect-261001-general-networking/t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf.md
source_anchor: ""
source_lines: [1, 33]
sha256: b85a779eb6f02f0371878ace3618ba28aabf31e510d9ba030bbc77141c40ca98
---

# t5-switching-stp-stopped-receiving-bpdus-with-loop-guard-enabled-m-p-260381-b3af52bf

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-22-2025 01:42 PM
Hi All,
I hope we are well.
I’ve recently inherited a Cisco Meraki network following an acquisition and have noticed some alerts.
Edit - Added Network Topology/Switch overview (below)
Toplogy
Switch Overview
The network Consists of:
2 x MX105 in HA config.
5 x MS130-48P
4 x CW9166I
I've documented the network as best as possible.. (minus the access points) will re-review tomorrow and update asap.
Version 1
Version 2 - Since Reviewing Meraki Live Topology
- Root Ports mislabelled / corrected
DUG00-SW01 Port 51
DUG00-SW03 Port 50
The affected ports seem to be alternate uplink ports, so the network is still operational. I’d just like to understand the issue and get the network’s RAG status back to healthy.
Any thoughts on the best place to start troubleshooting?
Thanks in advance - let me know if you need more information!
Solved! Go to Solution.
- Labels:
- 
						
							
		
