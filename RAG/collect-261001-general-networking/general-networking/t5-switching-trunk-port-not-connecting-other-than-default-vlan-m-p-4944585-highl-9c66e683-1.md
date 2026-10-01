---
id: collect-261001-general-networking/general-networking/t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683-1
title: "t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683.md
source_anchor: ""
source_lines: [1, 47]
sha256: b705fa552f2e825e15c65b45022b77dbd4707f9ef20d25de0171bb686c2a996d
---

# t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 08:02 PM
I have a router on a stick set-up. Router is connected to the main switch (cisco 2960x series) via trunk port. Which accepts default vlan1 and vlan30,40 etc. I want to connect multiple vlans on one switchport and still connect to the network.
Router configuration below.
interface GigabitEthernet0/0/1
ip address 192.168.2.1 255.255.255.0
ip nat inside
duplex auto
speed auto
!
interface GigabitEthernet0/0/1.30
description vlan 30
encapsulation dot1Q 30
ip address 192.168.3.1 255.255.255.0
!
interface GigabitEthernet0/0/1.40
encapsulation dot1Q 40
ip address 192.168.4.1 255.255.255.0
Switch configuration below
interface GigabitEthernet0/1
switchport trunk allowed vlan 1,30,40
switchport mode trunk
Access ports which allowed vlan 30 or 40 (access ports accept one vlan only) connect vlan 30 or 40 devices just fine.
As soon as I configure one of the switchports to trunk and allow vlan 1,30,40 only devices from vlan1 connects to the network but not from vlan 30 or 40 devices.
Port configuration below
interface FastEthernet0/4
switchport trunk allowed vlan 1-2,30,40
switchport mode trunk
Essentially, I want to connect any vlan device to any switch port and still connect to my network.
Thank you
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Catalyst 2000
- 
						
							
		
