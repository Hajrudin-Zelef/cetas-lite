---
id: collect-261001-general-networking/general-networking/t5-switching-etherchannel-m-p-885848-highlight-true-020b4d19
title: "t5-switching-etherchannel-m-p-885848-highlight-true-020b4d19"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-switching-etherchannel-m-p-885848-highlight-true-020b4d19.md
source_anchor: ""
source_lines: [1, 125]
sha256: ad635d0582e94de40c2ece8b9f27865e1ae73823de420277827c8c24f431ff7c
---

# t5-switching-etherchannel-m-p-885848-highlight-true-020b4d19

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 06:59 AM - edited 03-05-2019 07:37 PM
I was wondering the differences between the methods of configuring Etherchannels.
Pagp, LACP and "mode on"
We recently had some contractors come through and re-engineer one of our facilities. I'm trying to go through and fix inconsistencies and misconfigurations but before I do I need to be sure of their implementation.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Other Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 11:00 AM
You should use care when using the on mode. This is a manual configuration, and ports on both ends of the EtherChannel must have the same configuration. If the group is misconfigured, packet loss or spanning-tree loops can occur.
When using PAgP or LACP, the switch learns the identity of partners capable of supporting PAgP or LACP and the capabilities of each port.
It then dynamically groups similarly configured ports into a single logical link (channel or aggregate port). Similarly configured ports are grouped based on hardware, administrative, and port parameter constraints. For example, PAgP or LCAP groups the ports with the same speed, duplex mode, native VLAN, VLAN range, and trunking status and type. After grouping the links into an EtherChannel, PAgP or LCAP adds the group to the spanning tree as a single switch port.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 07:06 AM
The main difference is that the Cisco implementation uses a proprietary protocol called Port Aggregation Protocol (PAgP). The IEEE later defined within 802.3ad a new control protocol for link aggregation called Link Aggregate Control Protocol (LACP).
The primary advantage of using LACP (802.3ad) is for interoperability with other vendor switches. Since PAgP is Cisco proprietary, one cannot aggregate channels between a Cisco switch and a switch from another vendor unless LACP is utilized.
Hope that help. Please rate if that helped.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 07:15 AM
Thanks for the reply,
What about the "mode on" it doesn't use either protocols. Various show commands confirm no protocols are running.
In short I'm trying to clean up a contractors work after they were brought in to re-engineer one of our facilities. Ciscoworks is showing logical discrepencies related to their etherchannel implmentation, complaining about spanning-tree. I'm pretty sure I have it narrowed down to the fact they configured the etherchannel to "mode on" which puts the ports into UNI state which is "not compatible with STP."
Before I go through and make configuration changes I need to be sure of what advantages or disadvantages.
Thanks in advance.
Craig
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 07:23 AM
In the LACP Mode On, The link aggregation is forced to be formed without any LACP negotiation .In other words, the switch will neither send the LACP packet nor process any incoming LACP packet. This is similar to the on state for PAgP. Although this works, it is not recommended.
Note: By default, when an LACP channel is configured, the LACP channel mode is passive.
So you should reconfigure the LACP port to passive it more relevant.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 10:26 AM
LACP is not active in the channel. The link is configured without PaGP and LACP. The switch recognizes there is an etherchannel, it shows up/up, spanning-tree calculates the cost correctly.
So LACP and PaGP are just negotiation protocols, nothing else?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 10:52 AM
"So LACP and PaGP are just negotiation protocols, nothing else? "
Exactly.
"On" mode uses neither of those.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 11:00 AM
You should use care when using the on mode. This is a manual configuration, and ports on both ends of the EtherChannel must have the same configuration. If the group is misconfigured, packet loss or spanning-tree loops can occur.
When using PAgP or LACP, the switch learns the identity of partners capable of supporting PAgP or LACP and the capabilities of each port.
It then dynamically groups similarly configured ports into a single logical link (channel or aggregate port). Similarly configured ports are grouped based on hardware, administrative, and port parameter constraints. For example, PAgP or LCAP groups the ports with the same speed, duplex mode, native VLAN, VLAN range, and trunking status and type. After grouping the links into an EtherChannel, PAgP or LCAP adds the group to the spanning tree as a single switch port.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 01:18 PM
That's exactly what I was looking for.
Now tying that into STP, I'm getting logical discrepencies reported on the etherchannels configured by the contractors.
"Fast EtherChannel Port Spanning Tree Not Disabled"
Obviously we don't want spanning tree disabled, and spanning tree is in fact blocking the proper paths.
The way I see it is, this is more of an informational message. Since each etherchannel is not negotiating speed/duplex, spanning tree is making more of an assumption based on local port configurations instead of an informed "best path" selection. If the ports on the other side were misconfigured to a lower speed spanning tree would not know the difference and possibly create a loop because it cannot make an informed decision.
Since it is properly configured at the moment, it is simply not allowing spanning tree to make "complete" decisions.
Does that sound correct?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-26-2007 06:29 PM
>Does that sound correct?
Yes, you will encounter a lot of informational messages when configuring etherchanneling.
Best Practice is to shut down the participating switchports from both link partners, configure etherchanneling and bring them up.
This procedure will prevent switchports going into err-disabled.
