---
id: collect-261001-general-networking/general-networking/t5-switching-help-with-spanning-tree-loop-guard-blocking-m-p-4188648-566f8224
title: "t5-switching-help-with-spanning-tree-loop-guard-blocking-m-p-4188648-566f8224"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-help-with-spanning-tree-loop-guard-blocking-m-p-4188648-566f8224.md
source_anchor: ""
source_lines: [1, 97]
sha256: c7596c7b9d3a0c2e68333f2b482e564055c4c67071b39a413af8d9e9b867cfef
---

# t5-switching-help-with-spanning-tree-loop-guard-blocking-m-p-4188648-566f8224

Help with Spanning tree loop guard blocking
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 02:26 AM
Hi everyone,
Looking for some help with one of our 3850 stacks which is getting this within its logs:
SPANTREE-2-LOOPGUARD_BLOCK: Loop guard blocking port Port-channel1 on VLAN0102.
The stack has 1 uplink to our core switch configured with the below configuration:
interface TenGigabitEthernet2/1/3
description LAG to Core
switchport mode trunk
channel-group 1 mode active
PortChannel 1 config:
interface Port-channel1
description LAG to core
switchport mode trunk
This issue seems to have only started within the last 24hrs, as interface Po1 is the uplink to the core when it is blocking it is affecting users.
Any help would be appreciated greatly !
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
11-25-2020 03:31 AM
Hello @Ollie Colebrook ,
you have probably enabled loop guard at global level as we don't see spanning-tree commands in the port-channel configuration.
In my understanding and experience the error is triggered by the fact that in the affected VLAN 102 the access switch does not receive BPDUs coming from the root bridge on the uplink port. 3 missed in a row are enough for Rapid STP
Instead of promoting itself to root bridge for VLAN 102 the switch for the loop guard feature moves the port in an inconsistent state for Vlan 102 waiting to receive again STP BPDUs from the core / root bridge.
when this happens the port is moved again to normal state.
Users of Vlan 102 are impacted.
Look for input errors and traffic level on the member links and verify that the root bridge for VLAN 102 is the expected one.
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 04:22 AM
Hi @Giuseppe Larosa Thank you for the information.
Yes correct it looks like its enabled at the global level.
Using the commands show spanning-tree inconsistent ports it shows 0 ?
Thank you
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 04:04 AM
Hello
You shouldnt really lose connection when you have a PC even with loopguard intiating as only the affected physical link should become blocked, So can you confirm if you do only have a one physical link in the PC ?
Please rate and mark as an accepted solution if you have found any of the information provided useful.
This then could assist others on these forums to find a valuable answer and broadens the community’s global network.
Kind Regards
Paul
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-25-2020 04:24 AM
Hi @paul driver
Correct there is only 1 physical interface connected to this switch stack from our core switch. There should be two interfaces but it looks like this was missed. If we were to have two interfaces to the core both within the PC would this help the issue ?
Regards,
Ollie
