---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter-fb4ea8ec-3
title: "t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter--fb4ea8ec"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter--fb4ea8ec.md
source_anchor: ""
source_lines: [185, 245]
sha256: 0cd752f02d4e1573c6e0e35ddc5f49a0bb91dbe01067d4f3136b601084afcfe7
---

# t5-routing-and-sd-wan-pc-server-ping-problem-after-vlan-configuration-amp-inter--fb4ea8ec

So I tried to implement mesh topologies in both LAN and MAN connection. So I tried to connect mesh topology with mesh topology. As M02@rt37 and @Martin L said earlier, I suppose that this topology is impossible. Therefore, I change the topology from connecting mesh topology with another mesh topology to connect star topology (MAN) with another star topology (LAN).
It works! Here's the design I made:
Achieved wanted objectives:
- LAPTOP-O-01 and LAPTOP-O-02 able to ping LAPTOP-N-05 and LAPTOP-N-06 only (all under VLAN 10)
- LAPTOP-O-03 and LAPTOP-O-04 able to ping LAPTOP-N-07 and LAPTOP-N-08 only (all under VLAN 20)
However,
- The server SERVER-MAIN-01 cannot ping all laptops (end devices) and all laptops cannot ping the server
How to fix this? Because after fixing this server ping problem, I want to make the server for network monitoring. I want to implement Net flow for network monitoring.
@Martin L , I didn't study about VSLM subnetting and NAT/PAT. I just follow what the YouTube says, I'm sorry. And I already adjust all end devices to use the 192.168.10.0/24 IP address block. I leave the updated
.pkt file
and the whole question below.
Thanks in advance and sorry for miscommunication.
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
07-09-2023 05:12 AM - last edited on 07-12-2023 11:44 PM by Translator
Hello M02@rt37 ,
Thanks a lot! It works now, for practical purposes, laptops under VLAN 10 can ping the server.
I also adjust the switch where as the router are connected to VLAN 10 too. Actually, the router is used for NetFlow configuration.
I have done the router R-SDEC settings
 ip flow destination 192.168.10.1 . Also, ip flow ingress and egress for Gig 3/0 R-SDEC
I can ping 192.168.10.1 (the server) and 192.168.10.2 (the router R-SDEC) from LAPTOP-0-01 under VLAN 10.
But there is no any output when I use the server's NetFlow app. If the server's NetFlow app got output, my whole assignment is complete.
How to fix this? Latest
.pkt
file below.
Thanks a lot in advance.
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
