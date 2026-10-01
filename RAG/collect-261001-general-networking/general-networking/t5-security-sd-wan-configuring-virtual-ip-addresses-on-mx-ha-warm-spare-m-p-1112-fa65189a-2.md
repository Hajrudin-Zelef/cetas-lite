---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a-2
title: "t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a.md
source_anchor: ""
source_lines: [34, 174]
sha256: a05b6bab90f24d7416af9929709b03f6326fb66af709e457862d2728e368e3da
---

# t5-security-sd-wan-configuring-virtual-ip-addresses-on-mx-ha-warm-spare-m-p-1112-fa65189a

			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-08-2021 02:45 PM
Hi @hmc250000
You can use either public or private IPs for the VRRP addresses but would depend on your use case. VRRP addresses are only configured on the WAN Interfaces.
Taken from the below document:
Use MX uplink IPs: When using this option, the current Active MX will use its distinct uplink IP or IPs when sending traffic out to the Internet. This option does not require additional public IPs for Internet-facing MXs, but also results in more disruptive failover because the source IP of outbound flows will change.
Use virtual uplink IPs: When using this option, both MXs will use a shared virtual IP (VIP) when sending traffic out to the Internet. This option requires an additional public IP per uplink but allows for seamless failover because the IP address the network is using to communicate with the Internet will be consistent. The VIP for each uplink must be in the same subnet as the IPs of the MXs themselves for that uplink, and the VIP must be different from both MX uplink IPs.
Only one MX is Active at a time, when the Primary fails everything moves over to the secondary. So no need to change default routes etc.
https://www.linkedin.com/in/darrenoconnor
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-08-2021 02:54 PM
Sorry I'm still confused. Are both MX's (active and passive) not supposed to have unique ip addresses for the LAN?
We use static routes on our L3 switch.
Which article is more up to date?
this
https://documentation.meraki.com/MX/Deployment_Guides/MX_Warm_Spare_-_High_Availability_Pair
or this?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-08-2021 03:44 PM
Greetings @hmc250000!
What @DarrenOC is true!
HA configuration is quite straightforward and the doc you originally posted should do!
To your questions:
- "Where do we configure the VIP?" The VIP can only be configured on the WAN interfaces
- "Can we use private ip addresses?" Private address space can be used for your WAN interfaces including your VIP - likewise, these must be in the same subnet. So just make sure there is at least a /29 to accommodate the 3 address. A-WAN1/B-WAN1/VIP
- How would this work for the LAN?" Typically, we see, depending on the placement of the MXs and the gateways, the LAN ports configured as transit or management networks. There also needs to be communication between the LAN ports of each MX via a switch downstream in order to exchange VRRP heartbeats. The MX doesn't participate in STP, so it is best like Willette.Works' doc mentions to prune out your Mgmt/Heartbeat VLAN on a specific LAN port to prevent unintended BPDU flooding.
Hope that helps!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 03:34 AM
Hi Meraki Experts,
Greetings and hats-off for the forum replies. Appreciated!
We have a use-case to aggregate the ISP-1 and ISP-2 links in MX HA pair as per my attached design. Here we go for the physical connectivity for WAN and LAN using switches. I've following queries to get the response from you guys..
1) Should we manage to aggregate both ISPs in MX irrespective of MX hardware level failover?
2) Should we end-up landing into STP issues by having same Switch for both Uplink & Downlink?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 07:21 AM
You should really create a new post for this separate enquiry...
Leaving that aside - you use the term 'aggregate'; MX does not support link aggregation (802.3ad / LACP or similar). With MX you provide a separate WAN (Internet) uplink to each of WAN1 and WAN2 ports and the MX can make use of both links simultaneously, by allocating different sessions to different links, based upon preference and performance criteria you configure. so; don't configure any link aggregation on your switches, for connection to the MXs.
Each of your MXs should be connected to both switches, on the LAN side, for resilience (your diagram shows only one such link per MX). Note however that MX does not participate in spanning tree, but will forward STP BPDUs, to allow your switches to resolve the physical loops that will create. So; configure STP on the switches to match the desired topology (e.g. bridge priority).
Provided you do this properly and use different VLANs for WAN1, WAN2 and LAN-side VLANs, you should not get into STP-related problems.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 08:45 AM
My apologies if this post confusing the earlier topic. However, let me reply for your queries to make 100% clear within us.
I agree with you that i meant "aggregate" as a term refers to utilize both the ISP links together in MX on account of load sharing. Of course i don't enable any LACP between switches and MX WAN interfaces.
Considering the shortage of switch ports, we would plan to use 1x LAN port/downlink per MX. I understand we are compromising the resiliency but still it's OK to move on i believe. As a caveat, this would be a single point of failure for LAN/VRRP communications that i understand...
In conclusion, should we proceed to deploy the MX HA architecture using mentioned topology? If yes, i hope we can get following results around this design...
1) Internet link load sharing by using both ISP links in MX
2) Internet link failover
3) MX failover
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 09:20 AM
@Guru-Hexaware you should be just fine with one LAN link per MX. In fact, unless you are connecting the MXs to a switch stack I'd recommend one link over two.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 09:31 AM
Thank you for the reply mate!
Could you please confirm the below points as a conclusion of this query?
Should we proceed to deploy the MX HA architecture using the mentioned topology in Routed mode? If yes, i hope we can get following results around this design...
1) Internet link load sharing by using both ISP links in MX with /29 IP subnet each
2) Internet link failover
3) MX failover
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 09:37 AM
@Guru-Hexaware I personally would not use the same switch for internet splitting and LAN, though it should work as you have drawn it if you get the VLANs right.
We always use small unmanaged 5 port gigabit switches for the ISP splitting as they cost next to nothing and keep the design simple.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-09-2021 09:46 AM
Thanks again mate.
Yes, we have planned for unique VLANs in respect to the ISP-1, ISP-2 and LAN connectivity for sure.
I totally agreed to have a separate switches for WAN and LAN termination however this is one of small branch site for customer. Still, we look forward to check the different switch in subject to the availability factor.
Sincerely, i thank you for the response on time. Highly appreciated for your dedication & efforts.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-08-2021 10:08 PM
Just in case you still need a hand....
- WAN interfaces can either use only their dedicated WAN IP address (in routed/NAT mode), or you can create a virtual IP address that moves between the primary and standby MX when failover occurs
