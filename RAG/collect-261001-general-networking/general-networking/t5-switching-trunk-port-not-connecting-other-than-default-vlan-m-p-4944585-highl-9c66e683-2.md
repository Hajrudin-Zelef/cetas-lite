---
id: collect-261001-general-networking/general-networking/t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683-2
title: "t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683.md
source_anchor: ""
source_lines: [48, 219]
sha256: bf06794b8f957e85aaea7eb7c582338734969c2b1f8d792517621a72b9d267dd
---

# t5-switching-trunk-port-not-connecting-other-than-default-vlan-m-p-4944585-highl-9c66e683

			LAN Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 08:18 PM - edited 10-18-2023 08:19 PM
Hello,
If I understand you correctly you want to plug in a PC on interface FastEthernet0/4 (while its configured as a trunk) and allow it to access the VLANs you mention (vlan 1-2,30,40), essentially making the PC a part of all the VLANS on the trunk and be able to communicate on all VLANs? If that's the case that is not doable and would undermine the reason VLANs exist. That along with the PC wouldn't be able to understand. here's why:
The PC connects on VLAN 1 because by default VLAN 1 is not tagged and is the native VLAN on a switchport. Since its untagged the PC can comprehend frames on that VLAN because its not being sent with a tag. Traffic form VLAN 30 and 40 are sent with tags on the trunk port towards the PC which does not understand tags.
The only instance you would use this setup is that port connects to a servers post that acts like a switch and CAN understand VLAN tags from a trunk port.
Hope that helps
-David
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 08:17 PM
interface GigabitEthernet0/0/1
No ip address 192.168.2.1 255.255.255.0
ip nat inside
duplex auto
speed auto
!
interface GigabitEthernet0/0/1.1
description vlan 1
encapsulation dot1Q 1 native
ip address 192.168.2.1 255.255.255.0
Try this way' use subinterface for vlan1 and make it native if vlan1 is native vlan of SW.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 08:49 PM
Yes vlan1 is native for SW. But solution did not work. Same result as before. Thank you.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 09:10 PM
Add
Switchport trunk encapsulate dot1x
Switchport mode trunk
To trunk config of SW and check again.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 10:24 PM
MainSwitch(config-if)#switchport trunk encapsulate dot1x
% Invalid input detected at 'e' marker.
Thank you. Command seems to be specific for other SW
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 10:41 PM - edited 10-18-2023 10:48 PM
Can you share topolgy here
Show interface x/x switchport
For port connect to router and other unmgmt sw
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 08:18 PM - edited 10-18-2023 08:19 PM
Hello,
If I understand you correctly you want to plug in a PC on interface FastEthernet0/4 (while its configured as a trunk) and allow it to access the VLANs you mention (vlan 1-2,30,40), essentially making the PC a part of all the VLANS on the trunk and be able to communicate on all VLANs? If that's the case that is not doable and would undermine the reason VLANs exist. That along with the PC wouldn't be able to understand. here's why:
The PC connects on VLAN 1 because by default VLAN 1 is not tagged and is the native VLAN on a switchport. Since its untagged the PC can comprehend frames on that VLAN because its not being sent with a tag. Traffic form VLAN 30 and 40 are sent with tags on the trunk port towards the PC which does not understand tags.
The only instance you would use this setup is that port connects to a servers post that acts like a switch and CAN understand VLAN tags from a trunk port.
Hope that helps
-David
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 08:45 PM - edited 10-18-2023 08:54 PM
Sorry if my description was not clear, what I want to achieve is that fa0/4 port is connected to unmanaged switch, which is connected to multiple vlan devices. Both vlan1 and vlan30 computer are connected to the unmanaged switch and still connect to the network.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 10:24 PM
one more thing; if u cannot add vlans to unmanaged switch, those ports should be in vlan 1 by default; All PCs connected to unmanaged switch should be able to communicate. All traffic in vlan 1 should reach other switch and router in vlan1 but not in other vlans.
Regards, ML
**Please Rate All Helpful Responses **
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-19-2023 05:43 PM
Then the devices need to tag the packets with the VLAN they want to communicate with. When a PC sends a frame and the switch on the trunkport receives it untagged it will assume it belongs to the default native VLAN of a trunk which is 1. You can test this by changing the native VLAN to 30 and then all devices will communicate on VLAN 30.
-David
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-19-2023 06:23 PM
Yes, that is my understanding. But there is multiple devices without VLAN tagging capabilities in my network. As I mentioned if RV082 can do this, ISR4331 combined with catalyst2960 should do the same.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-19-2023 06:48 PM
Keep in mind this is Packet Tracer and more advanced functionality like that may not be supported. I use the term advanced as I advanced for PT.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 10:05 PM
@Martin L Thank you. Here is the simplified PT version of my network.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 10:21 PM - edited 10-18-2023 10:26 PM
Vlans are missing on un-managed switch; PCs connected to 2nd switch will need to be in correct access ports; and in correct vlan x in order to communicate to PC on vlan x in other switches.
all links between switches must be trunk mode and of course switch to router must be trunking.
also vlans must be present in all switches; i.e. Sw1 == SW2 == Sw3, if vlan x is present on sw1 and sw 3, vlan x must be present on Sw2 in order to allow traffic
Cross-over cable between switches!
Regards, ML
**Please Rate All Helpful Responses **
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-18-2023 10:26 PM
Thank you, but unmanaged switch is plug and play device, no available configuration methods.
