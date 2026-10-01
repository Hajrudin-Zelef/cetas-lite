---
id: collect-261001-general-networking/general-networking/t5-switching-vlan-shows-down-down-td-p-1980251-e62c0237
title: "t5-switching-vlan-shows-down-down-td-p-1980251-e62c0237"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-vlan-shows-down-down-td-p-1980251-e62c0237.md
source_anchor: ""
source_lines: [1, 121]
sha256: 9cbe8e03760183441a0b1abe62b4cbb79a093cbbdbc38f46c9d012562c787d83
---

# t5-switching-vlan-shows-down-down-td-p-1980251-e62c0237

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-25-2012 12:51 PM - edited 03-07-2019 07:27 AM
Looking for a little help in why a newly created Vlan shows down/down when I issue a sh ip int brief command.
My setup is an AP with two SSIDs on two different Vlans (4 and 7). Traffic flows properly across Vlan 4, but Vlan 7 shows it is in a down state and cannot be pinged from the local switch. From documentation I've read, the Vlan should become active when a port, trunked or attached to the Vlan, connects to the Vlan. I have at least two ports trunked on the switch, (including the one the AP is connected to) but Vlan 7 does not appear to recognize them. Spanning-tree information comes back back with no instances for the Vlan. I presume this is due to the Vlan not being active.
What am I missing? Why will the Vlan not see the trunked ports, create the spanning tree instances, and activate so traffic can flow?
Solved! Go to Solution.
- Labels:
- 
						
							
		
			LAN Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-26-2012 08:09 AM
Interface vlan attaches a L3 svi to an already existing L2 vlan. Vlan database is being deprecated and you can create the L2 vlan with "vlan 7 
HTH,
John
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-25-2012 02:44 PM
Please post config along with "show vlan" , show ip int brief and show int trunk .
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-25-2012 03:13 PM
Traffic flows properly across Vlan 4, but Vlan 7 shows it is in a down state and cannot be pinged from the local switch.
What appliance is doing your inter-VLAN routing?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-26-2012 07:24 AM
Glen - I've attached zipped txt files with the requested information. Thanks for your help.
Leo - It's a 3750 switch.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-26-2012 07:35 AM
From what I'm seeing, the vlan doesn't exist on the switch at all which would cause the down/down state. If you're not running vtp, manually create vlan 7 and it should come up.
HTH,
John
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-26-2012 07:58 AM
I've seen creating Vlans two ways, through vlan database and through interface vlan. It was my understanding that vlan database was being phased out and that interface vlan should be the method used. In the running config, Vlan 7 does exist, though in the show vlan, it doesn't. I believe that's because it isn't active. Which is what I'm trying to correct.
At another site, this exact configuration works as needed, but if I remove the AP form the network (and thus any vlan 7 connections), Vlan 7 goes down in sh ip int brief, and does not appear in a show vlan.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-26-2012 08:09 AM
Interface vlan attaches a L3 svi to an already existing L2 vlan. Vlan database is being deprecated and you can create the L2 vlan with "vlan 7 
HTH,
John
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-26-2012 10:07 AM
Simply put , conf t , vlan 7 , exit ...
You have your layer 2 vlan which is created as shown above. And you have your layer 3 definition (SVI) which allows you to route packets off that vlan. You were missing the layer 2 definition so the L3 SVI could never come up . Also as an fyi if you are not running vtp client/server then change your vtp mode to tranparent and all your layer 2 definitions will also show up in your running config instead of just with a show vlan command.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-26-2012 11:52 AM
Yep, I can see the Vlan now showing as active. I'll test later today when I have the free time and supply the correct answer when it works out.
Thanks!
Edit - That did the trick! Tested and traffic is flowing as needed. Thanks for the help and quick replies!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
01-16-2023 10:28 PM
probably if you have configured a vpc so it need to be created in all switches in vpc domain and it will comes up
