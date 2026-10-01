---
id: collect-261001-general-networking/general-networking/t5-switching-high-rate-of-stp-topology-changes-m-p-69891-1a668023
title: "t5-switching-high-rate-of-stp-topology-changes-m-p-69891-1a668023"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-high-rate-of-stp-topology-changes-m-p-69891-1a668023.md
source_anchor: ""
source_lines: [1, 87]
sha256: 7cf7fa6f440acd0398b8e850d6d64adfb66edb672345e2f1d0fa0ce9c931f60f
---

# t5-switching-high-rate-of-stp-topology-changes-m-p-69891-1a668023

High Rate of STP Topology Changes
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-24-2019 02:17 PM
I get "High Rate of STP Topology Changes" on all of my Meraki Core switchports that are connected to other non-Meraki switches. Should I be concerned and what would be my next steps in troubleshooting this?
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-24-2019 04:30 PM
Found this helpful? Give me some Kudos! (click on the little up-arrow below)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-24-2019 08:08 PM
Agreed, I'd confirm all of the physical connectivity and do some quick packet captures or work with Support to find the source(s) of the TCN BPDUs and start from there, possibly even temporarily admin-downing certain redundant/blocking links to isolate the issue.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-02-2019 12:16 PM
I'm having the same issue on my Meraki. It only has one uplink to the Cisco core switch. None of my other Meraki switches are having the same issue. I'm fairly new to Meraki, so any help is much appreciated.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-24-2019 09:12 PM
Have you configured the root priority for the switch that should be the core of your network?
https://documentation.meraki.com/MS/Other_Topics/Switch_Settings#Setting_the_STP_Root_Bridge
Next the other switches, what are they? If they are Cisco Enterprise switches I would configure them to use mst "spanning-tree mode mst".
If they are dumb layer 2 switches it is possible they don't even run spanning tree.
Look for links between edge/leaf switches that should be there. You ideally want a loop free design where the downstream switches only connect to the core switch and no other switches.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-04-2019 11:24 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-07-2019 11:38 AM
Also make sure every single access port (connecting end devices) on your switches in your entire network is running in portfast/edge mode.
If you don't every time a device turns off or on you will get a TCN flood in your network causing shortened MAC address lifetimes.
