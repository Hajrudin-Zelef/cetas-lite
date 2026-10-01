---
id: collect-261001-general-networking/general-networking/t5-switching-root-bridge-id-spanning-tree-newbie-avoid-loop-td-p-1823027-099e9b31
title: "t5-switching-root-bridge-id-spanning-tree-newbie-avoid-loop-td-p-1823027-099e9b31"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-root-bridge-id-spanning-tree-newbie-avoid-loop-td-p-1823027-099e9b31.md
source_anchor: ""
source_lines: [1, 71]
sha256: 42ea988de0a51718cb6197bfaf88ef521287009b4985a69e192db57fc351622f
---

# t5-switching-root-bridge-id-spanning-tree-newbie-avoid-loop-td-p-1823027-099e9b31

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2012 08:02 AM - edited 03-07-2019 04:26 AM
Hello Experts,
I have an Extremely Old switch that I need to connect to my network. Because it is so old I don't want it to become the Root Switch.
Can anyone tell me what is the command to change the priority. (Honestly I don't remember if it has to be a lower number 1 or a higher number ). Always get that mixed up.
I've read about root guard, but I would like to prevent it manually. (It is a small network after all)
It is a Cisco 2950.
Thank you!
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2012 08:07 AM
Randall,
Use the following command in your switch global configuration mode:
spanning-tree vlan 1-4094 priority 61440
The root is the switch with the lowest priority setting, and the value indicated here (61440) is the highest configurable value. All other switches should therefore be more preferred.
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2012 08:07 AM
Randall,
Use the following command in your switch global configuration mode:
spanning-tree vlan 1-4094 priority 61440
The root is the switch with the lowest priority setting, and the value indicated here (61440) is the highest configurable value. All other switches should therefore be more preferred.
Best regards,
Peter
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2012 08:15 AM
Hey Peter, thank you very much for the fast response! Rating of 5 for speed and accuracy.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2012 09:40 AM
Hi Randall,
Thank you very much!
Best regards,
Peter
