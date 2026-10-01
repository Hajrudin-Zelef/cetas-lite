---
id: collect-261001-cisco/cisco/t5-switching-connecting-active-passive-palo-alto-pair-850-to-nexus-vpc-7k-m-p-39-5296b659
title: "t5-switching-connecting-active-passive-palo-alto-pair-850-to-nexus-vpc-7k-m-p-39-5296b659"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-switching-connecting-active-passive-palo-alto-pair-850-to-nexus-vpc-7k-m-p-39-5296b659.md
source_anchor: ""
source_lines: [1, 166]
sha256: 25bd412d2bbc59c4b14d9e3291f2209b3ddfaeea5cbacdfef769686d8b8f2172
---

# t5-switching-connecting-active-passive-palo-alto-pair-850-to-nexus-vpc-7k-m-p-39-5296b659

Connecting Active/Passive Palo Alto Pair(850) To Nexus VPC 7K Pair
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-26-2017 11:55 AM - edited 03-08-2019 11:29 AM
Hello,
Palo1(Active)(Inside seg) >>>(L2? L3-p2p?)7K1(VPC)
Palo2(Passive)(Inside seg) >>> (L2? L3-p2p?)7K2(VPC)
How should this be done in order to maintain redundancy?
Create a new SVI and VPC for the inside firewall segment, then configure the firewall facing link on each 7K as an access port? This would break the VPC design though, as the the endpoints(Palo Altos) are not capable of VPC or PC technology, right?
What about configuring the interfaces as L3 point to point links? But how would state knowledge of the neighboring Nexus be shared?
Finally, I thought about using a small switch like the 2960CG, port-channeling it up to the 7Ks, then connecting the PAs to the designated inside VLAN.
All support is appreciated.
- Labels:
- 
						
							
		
			Other Switches
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-26-2017 01:04 PM
Hi,
I have never deployed PA firewalls but if they function the same as Juniper and Cisco firewalls, you can connect the active firewall to one nexus and passive to the other nexus, put them in one vlan (access) with a /29 or 28 subnet with IP on each device. Nexus-1 one IP, Nexus-2 one IP and firewalls one IP if they are clustered, if not one each.
HTH
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-11-2019 06:08 AM
hi,
i know this was 1 year before but if you need any help in deploying the PA with Cisco network gear, nexus or CAT family. please respond and i will provide you the configurations for VPC or PO's as i have deployed them in both environments.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-18-2019 12:40 PM
Hi Usman,
Any information you may have on connecting Active/Passive pair of Palos to Nexus5K; would be great!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-10-2019 02:44 PM
Could you please share the recommended configuration on Nexus side for:
Nexus VPC to PA active/passive in L2 mode.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-10-2019 05:43 PM
sure why no. we have a multi zone config. I will post the config and a diagram if I can here otherwise send me a buz on usmanalidar@outlook.com and I will share the complete step by step doc with diagram that we have.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-17-2019 01:07 PM
apology for this delay got into some work and forgot to respond you. anyhow please find the configurations below for your understanding.
Please find the attached diagram also for your review
so VPC 81 is trusted zone or inside and VPC 91 is untrusted zone or outside
NEXUS 1: PRIMARY
==========================
interface port-channel81
description WIRED INSIDE
switchport mode trunk
switchport trunk allowed vlan "whatever is required"
speed 10000
vpc 81
interface port-channel91
description PA-WLL-outside
switchport mode trunk
switchport trunk allowed vlan "whatever vlan required"
speed 10000
vpc 90
interface Ethernet1/9
description WIRED INSIDE PA-1(ACTIVE)
switchport mode trunk
switchport trunk allowed vlan XXXXX
channel-group 81 mode active
interface Ethernet1/10
description WIRED INSIDE PA-2 (PASSIVE-FIREWALL)
switchport mode trunk
switchport trunk allowed vlan XXXX
channel-group 81 mode active
interface Ethernet1/14
description WIRED OUTSIDE PA1 (ACTIVE FIREWALL)
switchport mode trunk
switchport trunk allowed vlan XXX
channel-group 91 mode active
interface Ethernet1/15
description WIRED OUTSIDE PA2 (PASSIVE FIREWALL / STANDBY)
switchport mode trunk
switchport trunk allowed vlan XXXX
channel-group 91 mode active
NEXUS 2: SECONDARY
interface port-channel81
description WIRED INSIDE
switchport mode trunk
switchport trunk allowed vlan XX
speed 10000
vpc 81
interface port-channel91
description WIRED OUTSIDE PA
switchport mode trunk
switchport trunk allowed vlan XX
speed 10000
vpc 91
interface Ethernet1/9
description WIRED INSIDE PA-1(ACTIVE)
switchport mode trunk
switchport trunk allowed vlan XXXXX
channel-group 81 mode active
interface Ethernet1/10
description WIRED INSIDE PA-2 (PASSIVE-FIREWALL)
switchport mode trunk
switchport trunk allowed vlan XXXX
channel-group 81 mode active
interface Ethernet1/14
description WIRED OUTSIDE PA1 (ACTIVE FIREWALL)
switchport mode trunk
switchport trunk allowed vlan XXX
channel-group 91 mode active
interface Ethernet1/15
description WIRED OUTSIDE PA2 (PASSIVE FIREWALL / STANDBY)
switchport mode trunk
switchport trunk allowed vlan XXXX
channel-group 91 mode active
Hope that helps, if you need anything else please feel free to ask
Regards
