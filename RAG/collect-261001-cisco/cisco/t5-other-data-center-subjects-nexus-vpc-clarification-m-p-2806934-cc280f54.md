---
id: collect-261001-cisco/cisco/t5-other-data-center-subjects-nexus-vpc-clarification-m-p-2806934-cc280f54
title: "t5-other-data-center-subjects-nexus-vpc-clarification-m-p-2806934-cc280f54"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-other-data-center-subjects-nexus-vpc-clarification-m-p-2806934-cc280f54.md
source_anchor: ""
source_lines: [1, 43]
sha256: 1725a5a1474e8bb59d8c0afa8dbc2d8ab34b2733c9d7f786e816e72c16e9a487
---

# t5-other-data-center-subjects-nexus-vpc-clarification-m-p-2806934-cc280f54

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-08-2015 01:00 PM - edited 03-01-2019 08:05 AM
After a recent switch deployment, I'm seeing some very strange layer 2 forwarding issues. I am uncertain if this is related to a configuration issue on my part. I would like some clarification about VPC, and I'm hoping you can help me.
It is my understanding that when you have a series of double-sided VPCs between, say, a pair of N7700 and multiple pairs of Nexus 9K or 5K switches, that all the VPC domain IDs must be different. However, a contractor told me that each VPC number - associated with the etherchannel uplink to the Nexus 7Ks - must also be unique. Attached is a diagram which illustrates my dilemma.
Question - Am I ok to reuse port-channel number Po3 with VPC number 3 for all leaf pairs, or must these be different?
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Other Data Center Topics
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-09-2015 10:45 AM
Hi,
What you show is a perfectable acceptable solution. The two vPC domains at the Nexus 9000 layer, vPC domain 117 and 118 in your diagram, have no knowledge of each other and so the same vPC and port-channel number may be used.
As you've indicated, you obviously need unique numbering at the Nexus 7000 layer.
Regards
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-09-2015 10:45 AM
Hi,
What you show is a perfectable acceptable solution. The two vPC domains at the Nexus 9000 layer, vPC domain 117 and 118 in your diagram, have no knowledge of each other and so the same vPC and port-channel number may be used.
As you've indicated, you obviously need unique numbering at the Nexus 7000 layer.
Regards
