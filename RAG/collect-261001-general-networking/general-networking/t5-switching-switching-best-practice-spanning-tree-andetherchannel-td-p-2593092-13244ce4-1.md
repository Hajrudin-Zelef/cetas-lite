---
id: collect-261001-general-networking/general-networking/t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4-1
title: "t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4.md
source_anchor: ""
source_lines: [1, 22]
sha256: e836c509708a4c3145f89c5cb5e8eac50d702ab196231b65b310cefd7a83b897
---

# t5-switching-switching-best-practice-spanning-tree-andetherchannel-td-p-2593092-13244ce4

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-22-2014 11:21 AM - edited 03-07-2019 09:37 PM
Dear All,
Regarding best practice related to Spanning Tree and Etherchannel, we have decided to configure following.
1. Manually configure STP Root Bridge.
2. On end ports, enable portfast and bpduguard.
3. On ports connecting to other switches enable root guard.
In etherchannel config, we have kept mode on on both side, need to change to Active and desirable as I have read that mode on may create loops? Please let me know if this is OK and suggest if something missing.
Thank You,
Abhisar.
Solved! Go to Solution.
- Labels:
- 
						
							
		
