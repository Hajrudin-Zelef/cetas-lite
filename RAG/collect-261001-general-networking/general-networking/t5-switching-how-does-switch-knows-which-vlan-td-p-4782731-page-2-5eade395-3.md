---
id: collect-261001-general-networking/general-networking/t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395-3
title: "t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395.md
source_anchor: ""
source_lines: [154, 198]
sha256: fdda9544b944a8f5308def59f5b22dd2db57ad7bc6dfe4252230e3cb1c4b3d97
---

# t5-switching-how-does-switch-knows-which-vlan-td-p-4782731-page-2-5eade395

An access port, though, doesn't used tagged frames, but because the prior example configures the port as a VLAN 10 port member, frames entering this port would be considered/treated as part of VLAN 10, and frames exiting this port, would only come from VLAN 10.
Again, given 4 access ports, two in VLAN X and two in VLAN Y, none with VLAN tagged frames, how does the switch "know" what VLAN they belong to, is accomplished by what access port they entered on. What access port the frame may exit on, is accomplished by access port being a member of the same VLAN. Lastly, how is the switch not confused by multiple VLAN frames, within it, we don't know, unless the manufacturer documents that. Whatever/however the switch does what it does, VLAN X and VLAN Y frames are "known" internally.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 03:43 PM
You ask "So why cannot a layer 2 switch forward frame between 2 VLANs (VLAN 5<>VLAN 10) as it has knowledge of all port mappings & VLAN ID?". There is a simple answer to this question and a more complicated answer.
First the simple answer: it is a basic limitation of layer 2 switch that it can not forward between vlans. To forward between vlans you need a layer 3 switch.
The more complicated answer: a vlan is a broadcast domain and a layer 2 switch can forward traffic to any destination in that broadcast domain but can not forward to a different broadcast domain. So if a switch has an access port in vlan 5 and receives a frame on that port with a destination that is in vlan 5 the switch can forward with no problem. But if that switch receives a frame on that access port and the destination is in vlan 10 then the layer 2 switch can not forward to that destination. (remember that right now we are dealing with mac actresses for layer 2 and not IP addresses for layer 3. note that the host device that originated that frame should not send the frame with a destination mac address in vlan 10).
If we have vlan 5 and vlan 10 and we want them to communicate then we need something that can operate between the vlans and that would be some layer 3 device (could be layer 3 switch, or could be router, or might be firewall). vlan 5 would have its own subnet (perhaps 10.5.5.0) and vlan 10 would have its own subnet (perhaps 10.10.10.0). If we think about the device that is connected to an access port in vlan 5 it would have an IP address in that subnet (perhaps 10.5.5.55). If it wanted to communicate with some device in vlan 10 (perhaps 10.10.10.10) its processing logic would recognize that it want to communicate with a "remote" device and would not look for a mac address of 10.10.10.10 but would look for a destination mac address of a device that can communicate between subnets (its default gateway).
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-26-2023 03:53 PM
Having sent that response it occurs to me that there is a different approach to the question "So why cannot a layer 2 switch forward frame between 2 VLANs (VLAN 5<>VLAN 10) as it has knowledge of all port mappings & VLAN ID?"." The question assumes that a frame has arrived on an access port in vlan 5 and the destination mac address is in vlan 10. But how/why would a frame be sent in vlan 5 whose destination is in vlan 10? That assumes a serious bug in the device that originated the frame.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-27-2023 12:13 AM
When we think about networking at layer 3 there is the concept of networking that is local and networking that is remote. So from vlan 5 to vlan 10 is possible. But at layer 2 there is not any concept of remote. The original post specifically asks about layer 2 switch. In a layer 2 switch if a frame enters in vlan 5 it can not be forwarded anywhere that is not in vlan 5. The switch may know about vlan 10 and be forwarding traffic in vlan 10 as well as in vlan 5. But a layer 2 switch can not forward from vlan 5 to vlan 10. To get between the vlans you need layer 3.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-27-2023 02:48 AM
Yeah,
Yeah, the funny thing about all of this is that Host A will not even try to reach Host B using L2. Host A and Host B being in two different VLANs will be in two different networks. Host A will recognize that Host B is in a different network and it will send a L3 packet to its gateway for delivery to Host B.
