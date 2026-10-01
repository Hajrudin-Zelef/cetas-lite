---
id: collect-261001-cisco/cisco/t5-switching-nexus-5k-vpc-pair-connecting-to-vxrail-nodes-td-p-5052186-b94916cd
title: "t5-switching-nexus-5k-vpc-pair-connecting-to-vxrail-nodes-td-p-5052186-b94916cd"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-switching-nexus-5k-vpc-pair-connecting-to-vxrail-nodes-td-p-5052186-b94916cd.md
source_anchor: ""
source_lines: [1, 103]
sha256: e55720a98c5e6331176ce8c19764b9ba0851506d6a50615b423a4789af10493f
---

# t5-switching-nexus-5k-vpc-pair-connecting-to-vxrail-nodes-td-p-5052186-b94916cd

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-28-2024 08:00 AM
First I'll state i also have an open case with Vmware.
I inherited a vPC pair that had a number of vPC set up, but the multiple links to the VxRail host nodes were trunk links. We lost the secondary so only the trunks to the primary vPC were up. And we lost all connectivity. First thing TAC asked was where are the vPC? Which led me down the road of these need to be vPC links to these hosts.
The fix in the end was to actually abandon the idea that we had any sort of redundancy and I was able to eventually get the secondary up and when I did, without a single VMware change, everything just came back up normally.
Setting these links to vPC is obviously pretty straightforward but there's a lot of talk on the internet about how they should just be trunk links. Clearly if that's the case there's more to the config than just that. This Nexus pair has been up for over 8 years so the redundancy was just assumed and not tested and it failed miserably. The other devices that are connected with vPC pairs remained up on the primary, seemingly unaware there was even an issue.
Anyone have this exact experienc, VxRail nodes connecting to a 5K vPC pair and successfully tested the redundancy by losing a switch?
Solved! Go to Solution.
- Labels:
- 
						
							
		
			LAN Switching
- 
						
							
		
			Other Switches
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-29-2024 01:17 AM
as i suggest good to have Physical diagram to help better along with what logs you have collected when the issueoccur.
This is not like simple steps to offer solution, this required clear understand how your Layer2 connected, where the problem persists ?
vPC only understand by Cisco (other vendor not have any visibility what is vPC).
Since this is Multi vendor integration and you mentioned you already have TAC case (i would pursue with them to trouble shoot and collect the logs)
To be clear, the failure happened with the VxRail ESXi hosts connected to each pair of switches with TRUNK links
I am in assumption sure this caused to Layer 2 Loop some where which i am thinking so far (guess games).
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-28-2024 08:12 AM
As per the switch concerns its just a LAG connection to Dell, You bind the Physical inteface to VXrail - i believe list like esxi dswitch.
(esxi do not care STP) - that should be working as expected - Not looked Dell Vxrail STP side, so better look that and setup STP in vPC as required.
sugges to look best practice vPC
still issue - post you Physical diagram how they connected and what nexus OS running and vPC / STP configuration for the community to help.
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-28-2024 10:20 AM
I'll look through the link. Yes the VMware side is ESXi hosts, vSphere/vCenter setup and dvSwitches on that end.
To be clear, the failure happened with the VxRail ESXi hosts connected to each pair of switches with TRUNK links. My contention is these needed to be vPC since they are Nexus in a vPC set up. I have an ECS storage device that remained accessible during the outage, which is connected via vPC, for example.
The set up as far as connections is pretty straightforward:
NexPrimary 1/5-----vNIC0 ESXi1 vNIC1-----NexSecondary 1/5
There are 8 hosts like this...config of course is the same on pri and sec, for example
interface Ethernet1/5
switchport mode trunk
switchport trunk native vlan 599
switchport trunk allowed vlan 500-550,3939
spanning-tree port type edge trunk
NexSecondary died, and it just looks like the primary didn't know what to do with the traffic on the trunk
With a vPC and the peer-link goes down the remaining switch will absolutely keep passing traffic.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-29-2024 01:17 AM
as i suggest good to have Physical diagram to help better along with what logs you have collected when the issueoccur.
This is not like simple steps to offer solution, this required clear understand how your Layer2 connected, where the problem persists ?
vPC only understand by Cisco (other vendor not have any visibility what is vPC).
Since this is Multi vendor integration and you mentioned you already have TAC case (i would pursue with them to trouble shoot and collect the logs)
To be clear, the failure happened with the VxRail ESXi hosts connected to each pair of switches with TRUNK links
I am in assumption sure this caused to Layer 2 Loop some where which i am thinking so far (guess games).
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-01-2024 11:03 AM
Thanks and sorry about the delay, holiday. It is multi vendor integration. And latest Dell/EMC VxRail code states they fully support vPC from Cisco Nexus devices...I was hoping someone with that specific setup would respond with their experiences. Supposedly these uplinks should have continued working with them just being trunk links but that is not my experience, and knowing what I know about Nexus vPC pairs, if you want redundancy they must be vPC which was also confirmed by at least twp of the TACs on this outage.
Vmware was not much help, but it's not really a VMware issue as they state, since this is a vxrail so I'm going down the path with Dell to better understand the vxrail side. Short of having the switch die i don't actually believe the design flaw to be a Cisco issue, it's the was the ESXi hosts/dvSwitch were configured.
Thanks for the advice though.
