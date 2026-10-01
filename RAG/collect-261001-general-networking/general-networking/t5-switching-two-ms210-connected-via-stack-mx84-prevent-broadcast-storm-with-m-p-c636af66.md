---
id: collect-261001-general-networking/general-networking/t5-switching-two-ms210-connected-via-stack-mx84-prevent-broadcast-storm-with-m-p-c636af66
title: "t5-switching-two-ms210-connected-via-stack-mx84-prevent-broadcast-storm-with-m-p-c636af66"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-two-ms210-connected-via-stack-mx84-prevent-broadcast-storm-with-m-p-c636af66.md
source_anchor: ""
source_lines: [1, 109]
sha256: 97bb30dbcdaa442bad603d19227e037a7747b20bda633daedb14ab7576fb7512
---

# t5-switching-two-ms210-connected-via-stack-mx84-prevent-broadcast-storm-with-m-p-c636af66

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 05:45 AM
Hi,
i have two ms210 switches connected via stack ports and a MX84.
My Idea:
Both of them are connected to one mx84 (uplink). (sw1.port24 goes to mx.port9 and sw2.port24 goes to mx.port10)
I want to configure a "semi" failover, where, when one switch goes down, we only need to replug a few ports on the switch. Does STP guard activated on the switch uplink ports to the mx prevent storm broadcasting or flooding? or do i have to disconnect one port?
We also have APs distributed equally and some servers, which have multiple ports, are plugged in on both switches (Server Port1 is connected to Switch1.Port10 and Server Port2 is connected to Switch2.Port10)
Is it possible to configure port-trunking/link aggregation on a server and enable port trunking for Switch1.port10 and Switch2.port10 also so that the traffic is distributed or if one switch goes down the other switch takes over?
Thanks for your help/input!
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 06:17 AM
@merda_sottile wrote:
if the switch is down, this is no problem, but if both switches are up, i have usually the problem having a loop between mx <- switch1 - stackconnection - switch2 -> mx, and then the connection goes havoc.
This typically happens when BPDUs from switch1 to the MX are not received back on switch2. One reason could be that you don't have a VLAN1 on the MX that can pass the BPDUs.
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 05:52 AM
You dont have to disconnect one port. Because if the switch is down the port is also down.
Yes you can create a lag between server and switch stack
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 05:59 AM
if the switch is down, this is no problem, but if both switches are up, i have usually the problem having a loop between mx <- switch1 - stackconnection - switch2 -> mx, and then the connection goes havoc.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 06:17 AM
@merda_sottile wrote:
if the switch is down, this is no problem, but if both switches are up, i have usually the problem having a loop between mx <- switch1 - stackconnection - switch2 -> mx, and then the connection goes havoc.
This typically happens when BPDUs from switch1 to the MX are not received back on switch2. One reason could be that you don't have a VLAN1 on the MX that can pass the BPDUs.
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-09-2021 12:16 AM
ok, .. there are multiple vlans configured, but primarily they are all in mgmt vlan 1.
i have now activated loop guard (just in case) on second switch port to mx, and it now shows "blocked" on that port, which is ok i think. i also assume now, that if the first switch goes down, the second port gets unblocked automatically?
In port configuration, the switch also recognized that it is connected to MX. (cdp/lldp), but is not recognized as uplink.
Usually i assume, (i have rtsp enabled on all ports), that the mx+switch configuration finds out that there is a ring and it uses both ports to communicate with the mx, or am i wrong with my knowledge?
all ports an congured as trunk, and there is no other port configured with bdpu or loop guard. they also have access to all vlans (but default is vlan 1).
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 05:54 AM
For the Link-aggregation, be aware that Meraki switches only support LACP and no static LAG config.
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 05:56 AM
yeah, but does the LACP trunkin has to be on one switch or is it possible to have ports on different switches and configure a stack aggregation?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2021 06:15 AM
On a stack you can have the links on different switch-members.
If you found this post helpful, please give it Kudos. If my answer solves your problem, please click Accept as Solution so others can benefit from it.
