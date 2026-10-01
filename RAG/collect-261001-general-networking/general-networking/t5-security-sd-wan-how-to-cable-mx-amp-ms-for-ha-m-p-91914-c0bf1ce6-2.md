---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-how-to-cable-mx-amp-ms-for-ha-m-p-91914-c0bf1ce6-2
title: "t5-security-sd-wan-how-to-cable-mx-amp-ms-for-ha-m-p-91914-c0bf1ce6"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-how-to-cable-mx-amp-ms-for-ha-m-p-91914-c0bf1ce6.md
source_anchor: ""
source_lines: [23, 162]
sha256: 971866ff3eb5f782264be1d8f2d2bc1550f042ad6330887fc09c8ad7e66463e0
---

# t5-security-sd-wan-how-to-cable-mx-amp-ms-for-ha-m-p-91914-c0bf1ce6

			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-03-2020 12:56 PM
H
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-04-2020 11:31 AM
does the LTE failover only work when the MX WAN1 and WAN2 physical links are down only?
I am seeing weird instances when the WAN links can still connect and ping their gateway to their connection to the cable modem or circuit ISP router that is still online but the internet peering next hop is down. From there it does not failover to the LTE.
what steps does the MX try before it determines this? Does it do some sort of SLA?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-03-2020 12:48 PM
I tested many scenarios but the best to get this to work is to have a designated vlan set on the MX for the heartbeat for example vlan 1111 ip 1.1.1.1/30.
then connect a cable between each MX for example MX1-e3 to MX2-e3 and set it as a trunk with the native vlan 1111 and allow all vlans (do not prune or pick any other vlans specified in your MX)
then your connections from the MX to the Skitch facing switches. Also if you are not using a Meraki switches like Cisco Catalyst or Arista then I set the Spanning of tree cost for the primary port to be 100 and the second link to the MX to 200
MX1-E4 —> SW1-E4
MX1-E5 —> SW2-E4
MX2-E4 —> SW1-E5
MX2-E5 —> SW2-E5
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-03-2020 12:50 PM
@lsborea wrote:
I tested many scenarios but the best to get this to work is to have a designated vlan set on the MX for the heartbeat for example vlan 1111 ip 1.1.1.1/30.
Can you please show me exactly how you designate a VLAN as heartbeat?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-03-2020 01:44 PM
like this
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 09:55 AM
Hi @lsborea ,
No where in there did you designate that VLAN as heartbeat. You simply created a VLAN, assigned it a subnet. Can you please clarify what makes this VLAN a heartbeat VLAN different from all the other VLANs?
How do you control which VLAN is used for hearbeats?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 10:09 AM
You can't designate. Heartbeats for VRRP are sent out across all VLANs by default, and there is no way to configure them to do otherwise. At this point that designated VLAN is probably redundant logically, maybe not physically if you separated it out (not sure if that is worth it either).
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 10:59 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 11:23 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 12:18 PM
Like I said I use a direct link between the two MX units.
I specified a direct VLAN for only that link and setting it was the native vlan trusting all vlans to be used as the heartbeat. maybe it only works logically so be it but it works for over 15 setups flawlessly once I did it this way and the LAN base links if one of them goes down and helps prevents the Master to Master HA issues..
Try it, if you don't like so be it figure it out another way.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 08:47 PM
Hi @lsborea,
I'm not saying it won't work. What I'm saying is that it's not possible to configure a "heartbeat VLAN" as described in that document. It's a completely false pretense to think that this supposedly magical VLAN is any different from any other VLAN you have configured. As @Nolan Herring pointed out, VRRP hellos (the "heartbeats") are sent out on every VLAN. Further, Meraki's implmentation of VRRP has elected to use a single Virtual Router with multiple Virtual IP addresses as opposed to a Virtual Router for each VLAN, meaning, that as long as a VRRP heartbeat makes is to the standby unit on ANY vlan the current active master will remain the master for all VLANs. It's not possible for only one VLAN to fail over. This is truly and Active/Standby only implementation.
So please forgive me if this sounds harsh as it's not meant to, but your 15 installations are not working because you followed that guide, and they also are not working the way you think they are working.
This is actually my objection to this configuration, and Aaron Willette's design. It's based on a falsity, and that falsity has been propagated to many corners of the Internet. Quite often people chime in here in the Community and reference that design, and nearly every time the person making the reference has no idea that it's simply not possible to actually designate a VLAN as a "Heartbeat VLAN".
My preferred way to do Meraki Warm Spare, which has come from my own experience deploying Warm Spare for numerous customers, and through the very insightful advice of @Philip D'Ath, is to NEVER EVER directly connect the two MX appliances together, and in stead singly connect each MX to a single switch or switch stack (the latter preferred for redundancy). See, having the MX's directly connected actually can cause issues under certain failures with Spanning Tree convergence. Further, it's also desirable to have the VRRP heartbeats traverse a path more representative of the path that user data actually takes. By creating a shortcut that heartbeats can take, but not user traffic, you are actually exposing yourself to a failure scenario where user traffic is blackholed, but VRRP is operation just fine and not allowing a failover to occur.
Failovers are a good thing! You want failover to occur when there's a failure. A direct link cheats the system and can actually prevent a failover when you actually want one.
I'll close this by pointing out that Meraki has actually changed their recommended topology for Warm Spare by removing the direct connection between MX's. It took me some time to convince them to do it, but they did come around
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 09:24 PM
jdsilva, so what do you mean my 15 installations are not working please explain..
In my setups I am not using Meraki switches but either Arista or Cisco Nexus not using Stacking but using MLAG or VTP.
Sure I get it, you do not need a set a VLAN to determine the heartbeat because VRRP is used on all VLANs sure but what if you connect the MX pairs to the south facing switches with only certain VLANs only and not the one VLAN you set for the heartbeat like in my example..
I know Meraki removed the peer link between MX HA pair but no one can explain to me why this is bad and why the failover will not occur when it happens..
You seem to know all the answers so explain it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-06-2020 09:28 PM
>My preferred way to do Meraki Warm Spare, which has come from my own experience deploying Warm Spare for numerous customers, and through the very insightful advice of @Philip D'Ath, is to NEVER EVER directly connect the two MX appliances together, and in stead singly connect each MX to a single switch or switch stack (the latter preferred for redundancy).
