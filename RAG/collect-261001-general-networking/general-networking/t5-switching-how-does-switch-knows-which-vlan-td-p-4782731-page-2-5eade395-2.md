---
id: collect-261001-general-networking/general-networking/t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395-2
title: "t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395.md
source_anchor: ""
source_lines: [15, 153]
sha256: e036be741cb83034a4a3d35dff0fe6508ab0976859f5308ca058422ed891915d
---

# t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-25-2023 11:32 AM
I understand VLAN creates a broadcast domain and L2 switch only fwd within the VLAN without an L3 device. However, how does the L2 switch knows which VLAN the frame needs to go without a dot1q Header?
For eg. I did a pcap and I don't see any VLAN ID or info on the ethernet header for access ports so when access ports send a broadcast or unicast how does the switch know which VLAN it's coming from and the other host is on the same VLAN?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-27-2023 03:00 AM
"However, when the frame comes in it doesn't say anywhere this frame belongs to VLAN 5 (at least I couldn't see it on a Packet Capture)."
Incorrect, it does "say". See my posting on how switch knows frame's VLAN membership even without VLAN tag.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-27-2023 03:06 AM
"So If the port is configured as an access VLAN 5 then the source mac = VLAN 5?"
MAC alone, no, because they only need to be unique per L2 domain.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-25-2023 02:25 PM
Adding to what @Richard Burts mentioned, by default nothing configured on the switch al in VLAN 1 (that is cisco default)
you can also view where the port connected end device - is this access port configured with access port vlan x ?
you did pcap where ? on the Access port ? can you share here your PCAP to understand ?
=====️ Preenayamo Vasudevam ️=====
***** Rate All Helpful Responses *****
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-25-2023 02:54 PM
Basically it's a simple scenario and I am trying to cement down the basic switching foundtation:
For Eg.
Host A in VLAN 5 is connected to an L2 switch access port VLAN 5. It sends a broadcast frame or an ICMP to a host on Access vlan 10 same switch. From what I saw on PCAP, there is nth on L2 header that tells a switch ok this frame is coming from VLAN 5 unlike if there is a dot1q tag where it adds the header. So I am trying to understand how switch makes a decision without that info.
When I took packet capture I did it on the actual wire. Does smth else happen on the backplane of the switch and changes while sending out to the wire?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 05:00 AM
I’m not an expert, but maybe my understanding of this will help.
Technically, Host A may be in VLAN 5, but the host itself does not have a clue what VLAN it is part of. Actually, it does not even know that there is some kind of VLAN. It has no concept of the VLAN.
A host sends and receives frames that do not contain a VLAN tag. A host is just connected to some access port. However, that access port is a member of a certain VLAN. And, only one VLAN, not many. It is that membership that tells the switch what VLAN a frame should go to. So, a switch does not need a VLAN tag on an access port at all. VLAN tags play role on trunk ports, not access ports. Although it was suggested, I do not think MAC address tables play any role in that, either. To me, MAC addresses are associated with VLAN IDs only to identify the VLANs MAC addresses belongs to. That’s just for switching purposes. When a frame exits a switch, the switch also looks at the port-VLAN membership and can send out a frame only through ports that are members of the VLAN the frame is in. The port-VLAN membership defines the boundaries of a VLAN broadcast domain.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 06:15 AM
Kris K
You may not be expert but you make some good points. So +5 for you. You are correct that a typical host in the network does not know about vlan and that there is not anything in the ethernet frame that the typical host sends. Knowledge of the vlan is in the switch. So the typical host sends a ethernet frame with no indication of vlan membership. The frame arrives on an access port of the switch. Based on the access port the switch determines which vlan this frame is associated with. If it is a layer 2 switch it can forward the frame only to other devices in that vlan.
There has been some discussion about vlan 5 and vlan 10. If it is a layer 2 switch it can not forward from vlan 5 to vlan 10. A layer 2 switch can only forward within the vlan that it received the frame from. To get from vlan 5 to vlan 10 the switch would need to forward the frame to some gateway/router.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 07:07 AM - edited 02-26-2023 07:10 AM
Yes, def +5 for Kris K. Sometimes a non-expert explanation is the best way to understand
So why cannot a layer 2 switch forward frame between 2 VLANs (VLAN 5<>VLAN 10) as it has knowledge of all port mappings & VLAN ID?
Is it bcoz a Host will send the frame to its default GW rather than directly to Host B to reach another VLAN which will be an L3 device?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 07:35 AM
If arp source from port with vlan 5' what port that SW will forward to??
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 08:05 AM
“So why cannot a layer 2 switch forward frame between 2 VLANs (VLAN 5<>VLAN 10) as it has knowledge of all port mappings & VLAN ID?”
Well, an L2 switch just does not want to do that . The IP subnet in VLAN10 would be different than the IP subnet in VLAN5. Changing subnets is routing, not switching. L2 simply does not deal with routing. Routing is dealt with in L3. That’s how the standards are defined.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 02:08 PM
Somehow switch associates frame with VLAN.
Exactly how this is done would depend on the switch's architecture and likely considered proprietary.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 03:24 PM - edited 02-26-2023 03:33 PM
Interesting, when I responded, initially, I didn't notice there have been a whole lot of replies, already.
@Gucamole you appear to understand that on a shared wire, VLAN tags are needed to identify what VLAN a frame belong to. But wonder, how a switch knows what VLAN a switch belongs to within the switch.
Well, as I wrote, "somehow" switch keeps track of what VLAN a frame belongs to, using some method that, generally, a switch manufacturer doesn't reveal. For all we know, it's using .Q VLAN frame tags, or ISL frame tags, or some other method.
How the switch keeps track of what VLAN a frame belongs to, can be the "secret sauce", but how the switch knows which VLAN a frame should belong to is done in different ways.
On a trunk link, all the frames will have a .Q or ISL (unlikely now a-days) frame tag, except for on Cisco trunks which have the concept of a "native" VLAN. The latter, which are untagged frames, by default, are assumed part of VLAN 1, unless we explicitly config what other VLAN should be used for any untagged frames:
E.g.
interface GigabitEthernet1/0/1
switchport trunk native vlan 10 !non-default explicit assignment
switchport mode trunk
An "access" ports, again the switch knows what VLAN frames entering/exiting that port should be, either, again, by default to VLAN 1, or explicitly assigned:
E.g.
interface GigabitEthernet1/0/1
switchport access vlan 10 !non-default explicit assignment
