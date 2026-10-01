---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter-fb4ea8ec-2
title: "t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter--fb4ea8ec"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter--fb4ea8ec.md
source_anchor: ""
source_lines: [29, 184]
sha256: 04dbec9fa38a6733beb5202bc44fb50f00873211e39b9cf8eea14dd008be42ec
---

# t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter--fb4ea8ec

			Other Routing
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-09-2023 12:48 AM - last edited on 07-12-2023 11:41 PM by Translator
Hello @syazwanmarzuki,
there is a misunderstood.
Your server has got
IP 192.18.10.1/24
and its Gateway is itself....... I don't understand the role of R-DSEC Router regarding its configuration.
You don't have anymore L3 equipement on your topology...then no more SVI acted as Gateway for each VLAN.
Then for what you expected is OK ; only PC on the same department can ping each other since they are in the same VLAN no need og Gateway!
But Server in VLAN 10 is only pingable frome other VLAN 10 ressources.....
Also adjust config on S-DSEC Switch if you want VLAN 10 ping Server MAIN-01:
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-09-2023 05:42 AM - edited 07-09-2023 05:43 AM
Adujst neflow configuration on R-SDEC:
You will have sample on netflow configurator on SERVER-MAIN-01:
Follow this for netflow config.:https://www.packettracernetwork.com/tutorials/packet-tracer-netflow.html
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-08-2023 06:55 AM - last edited on 07-12-2023 11:26 PM by Translator
Hello @syazwanmarzuki,
First, configure IP add ont the sub Int ; for example on Router R-BLOK-0:
Ping between laptops from engineerig to Administration should be OK (left side)
Other things, you can not have same IP ADD
(i.e. 192.168.10.0/24 - vlan 10)
on different routing segment. It should not work, you'll need to stretch vlan 10 (L2VPN).
You can have same vlan ID because it is a local ID, but on your context you should do subnetting -- subnetting the IP ADD for Engineering subnet and administration subnet. For example:
R-BLOK-0: vlan 10 - Engineering 192.168.10.0/25 and administration 192.168.20.0/25
R-BLOC-N: vlan 10- Engineering 192.168.10.128/25 and administration 192.168.20.128/25 (keep same for vlan 30 and 40).
On R-BLOK-0 side no need to have sub int for vlan 30-40! Same idea on R-BLOK-N side.
Proposition:
-No need multiple links between L2 switches
-Subnetting (Engineering/administration)
-R-BLOK-0 advertise (redistribute) 192.168.10.0/25 and 192.168.20.0/25
-R-BLOK-N advertise (redistribute) 192.168.10.128/25 - 192.168.20.128/25 - 192.168.30.0/24 and 192.168.40.0/24
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-08-2023 08:08 AM - last edited on 07-12-2023 11:28 PM by Translator
Hello M02@rt37 ,
I'd done configure as you asked on both Router R-BLOK-O and Router R-BLOK-N. I'd keep sub int for VLAN 30-40 for R-BLOK-0 for future implementation of the other two departments. Take note that I want to ping device in same department only (example: Left-side PC Engineering can ping right-side PC Engineering and vice versa).
Result:
- The laptops can ping/reach the switches, routers and the server but they can't reach the laptops from other building (same VLAN laptops)
- All Engineering and Administration devices can ping the server but the other two departments can't ping the server
- The server (source) can ping/reaches the switches and routers but it can't reach any laptop (destination)
Hello @Martin L ,
these are my remaining objectives while other objectives were completed.
Remaining Objective:
- LAPTOP-O-01 and LAPTOP-O-02 can ping LAPTOP-N-05 and LAPTOP-N-06 only (all under VLAN 10)
- LAPTOP-O-03 and LAPTOP-O-04 can ping LAPTOP-N-07 and LAPTOP-N-08 only (all under VLAN 20)
- LAPTOP-N-01 and LAPTOP-N-02 (VLAN 30 Account), LAPTOP-N-03 and LAPTOP-N-04 (VLAN 40 Sales & Marketing) can ping the server SERVER-MAIN-01
- The server SERVER-MAIN-01 can ping all laptops (end devices)
I did attached the updated file with these inter-VLAN configurations here.
Thanks in advance, again 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-08-2023 08:17 AM - last edited on 07-12-2023 11:30 PM by Translator
Delete SVI on you L2 Switches. For example, you have on the left side, engineering vlan,
IP add 192.168.10.1/24 on the Sw
and on the R-BLOK-O too!!!! Let's this IP on your L3 equipement, such as R-BLOK-R ; it serves as Gateway also for your vlan!
Why you wan SVI on your L2 switch? The only one should be an SVI for management purpose!
I delete on SW-BLOK-O-1 SVI and ip default gateway
LAPTOP can ping its Gateway hosted on R-BLOC-O
I delete also SVI on S-BLOC-O-2 and now vlan 10 ping vlan 20 on left side !
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-08-2023 08:29 AM - last edited on 07-12-2023 11:33 PM by Translator
you say "ake note that I want to ping device in same department only (example: Left-side PC Engineering can ping right-side PC Engineering and vice versa)."
it's impossible, in your topology like this, that Left-side PC Engineering can ping right-side PC Engineering and vice versa. As @Martin L explain to you. In other word, how you want R-BLOK-O router route packet from left-side
192.168.10.0/24 to its WAN Interface
towards Right-side engineering
subnet 192.168.10.0/24
? It's directly connected! You should NAT or do subnetting !
.ı|ı.ı|ı. If This Helps, Please Rate .ı|ı.ı|ı.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-08-2023 06:08 PM
What is objective of this lab? purpose or goals? is it just VSLM subnetting? did u study NAT/PAT ?
if u must follow pdf subnet table and there is no mistake/typo in pdf, then u must do PAT on edge routers. because you have same network and subnet on opposite sides, aka 2 companies. it just like real world example where u and I have 192.168.10.1 on home PC, yours and mine ISP does translation PAT for us so that e can communicate over the Internet.
Regards, ML
**Please Rate All Helpful Responses **
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-08-2023 07:03 AM - last edited on 07-12-2023 11:35 PM by Translator
That is normal ! By default, PCs in vlan 10 will not reach PCs in the same vlan 10 across L3 Broadcast domain. This is true even if they are on the same subnet. L3 B-cast domain is separating L2 and L3 domains. PCs on the same subnet will need some kind of NAT translation in your case on the edges. OR change IP subnet for vlan 10 and 20 on one side; then fix routing.
Routing with RIP is obsolete but if u want use RIP version 2 and
no auto-summary
commands.
there might be other issues ....
Regards, ML
**Please Rate All Helpful Responses **
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-09-2023 12:19 AM - last edited on 07-12-2023 11:38 PM by Translator
Hello M02@rt37 and @Martin L ,
First of all, thanks for spending time on configuring and troubleshooting my
.pkt file
However, there are a lot of misunderstanding and misconception. I beg for apologise from both of you.
My goal and objective is to create fastest topology and fastest internet connection since the question did told me:
The minimum bandwidth for internet access is 100 Mbps and is expected to increase significantly in few years ahead.
Discuss the most appropriate physical layer medium for the network connectivity of the following cases:
1) Within the Blok-O and Blok-N
2) All the way from Blok-O and Blok-N to the main server room in SDEC.
You may assume that the switches can support any transmission medium.
I didn't insert in the first place (the first .zip file), I'm sorry again.
