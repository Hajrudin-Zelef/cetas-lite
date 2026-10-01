---
id: collect-261001-cisco/cisco/t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td-aa260c21-3
title: "t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21.md
source_anchor: ""
source_lines: [203, 223]
sha256: 809c7b50f09033429871d4cf8bb1c8499e15d331ab0728ed5694804d466f17eb
---

# t5-switching-l3-routing-problem-with-nexus-vpc-pair-in-typical-core-position-td--aa260c21

So in your environment, frames entered the cat3k in vl 100. The cat3k will forward the frames out on the port-channel, so it will do a hash calculation to see wish of the ports to select for forwarding. This is done regardless of the status of the interface vlan 100 (shut or no shut) If it is the link towards n9k-2 then the n9k-2 will receive it. If the vlan 100 is up in n9k-2 it will do a routing lookup, see it should be sent to the firewall on vl 200, and then send it directly to the firewall on that link. If, however you have done a shutdown on int vlan 100, the n9k-2 can not do a routing lookup, so it just look in the mac-address table to see where to forward the frame. It sees it should be sent over the vpc-link towards n9k-1, and when it sends the frame it also inserts the marking that tells the neighbor that this frame was originally received from an "vpc port-channel. When n9k-1 receives the frame it does a lookup and sees that this frame has to be forwarded out to the firewall on a "vpc port-channel", but because the frame was marked as being received from a port-channel, it will drop the frame.
/Mikael
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2022 09:57 AM
thx mlund. the frame forwarded out to the firewall being dropped doesn't make sense to me. Shouldn't the N9k-1 see this as an L3 frame and realize, it has routed it to vlan200, and its now safe to forward to the Firewall next hop address? There is no reason to use some sort of L2 loop prevention on this frame at this point. Seems like the nexus cannot handle or doesn't have the smarts to run L3 SVI and L2 at the same time.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2022 10:08 AM
I think I found where you are misunderstanding, 
when the traffic receive for V100 it can not L3 routing using SVI of V200 it must route use V100 in other nexus Peer.
