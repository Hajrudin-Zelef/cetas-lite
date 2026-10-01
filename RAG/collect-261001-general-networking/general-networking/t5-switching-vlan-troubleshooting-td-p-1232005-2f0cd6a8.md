---
id: collect-261001-general-networking/general-networking/t5-switching-vlan-troubleshooting-td-p-1232005-2f0cd6a8
title: "t5-switching-vlan-troubleshooting-td-p-1232005-2f0cd6a8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-vlan-troubleshooting-td-p-1232005-2f0cd6a8.md
source_anchor: ""
source_lines: [1, 53]
sha256: 6ae1441d00bc04e3db2bd4818be19868f8bfaed040e6107386293fd8eed20004
---

# t5-switching-vlan-troubleshooting-td-p-1232005-2f0cd6a8

Vlan troubleshooting
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2009 09:20 AM - edited 03-06-2019 03:31 AM
Hi,
My company has 6500s as backbone and 3560s as end switches. We have around 100 vlan. And 6500 is routing between them. Noting else..No access-list between vlans..
Sometimes, users (2-3) complain about a program that uses network, that is very slow. And my network managers always tells me to check that vlan if something wrong. I check switch cpu, port utilization and i know that problem is not related to network. It always ends like service stop or server problem. What i want is, is there a tool that measures inter vlan routing or availability of switch ports etc.. I am sick of telling no problem on the network.
Thank you.
- Labels:
- 
						
							
		
			LAN Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2009 09:47 AM
You should configure NetFlow, please refer to the documentation:
HTH,
__
Edison.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-19-2009 09:54 AM
Does this happen to the same specific users all the time while others continue to access the server?
