---
id: collect-261001-general-networking/general-networking/t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c-3
title: "t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c.md
source_anchor: ""
source_lines: [203, 276]
sha256: 887c2fcf174f3001d1ebfa9f9fb93e857f0c3bddd303a37c2d95e7160fcee828
---

# t5-switches-small-business-sg500-vlan-configuration-m-p-2243861-highlight-true-3c41266c

So the switch is acting as a switch (L2 vlans) and router (L3 IPs). You have assigned the switch/router to answer or be the default gateway for each VLAN on 192.168.20.254 (VLAN20), 192.168.60.254 (VLAN60), 192.168.90.254 (VLAN90), 192.168.9.253 (VLAN1). Therefore, each device connected "downstream" of the switch would have a default gateway of the switch's VLAN IP. So everything on VLAN 20 would have an IP of 192.168.20.XXX with a subnet of 255.255.255.0 and a default gateway of 192.168.20.254. Same for 60 (192.168.60.254) and 90 (192.168.90.254).
In other words, everything on the inside LAN (PCs for example) would have the VLAN IP of the switch as their default gateway.
FYI - If your internet router is running a DNS server, then you can set their DNS to that router IP (192.168.9.254). My guess is it is not, so I would set each client to either an internal DNS server (like a windows machine running DNS with our without root hints) or to the ISPs DNS servers which are typically given by the ISP as part of the internet connection. You can also use publically available DNS servers like Google's 8.8.4.4 and 8.8.8.8 but I would avoid these where possible.
On the switch, you would set a default gateway to the internet router (in other words, anything the switch doesn't know about in it's routing table (directly attached VLAN IP subnets), it would send to this "upstream" device - aka the internet router). You can do that either with one or both of the commands given, ip route 0.0.0.0 0.0.0.0 192.168.9.254 (assuming your upstream router/fw has that internal address) or a "default route" command.
Now you need the internet firewall to know how to reach the internal VLANs. It knows about the 192.168.9.xxx subnet because it is directly attached to that. But it does not know about the .20, .60, .90. So on the internet router/fw, make sure the "incoming" packets can reach the internet VLANs/IPs. This is done by adding the static routes in the internet router/fw. Marty wrote above BUT HAS THE WRONG IP FOR THE NEXT HOP ROUTER I THINK. (Should be .253, the VLAN router. He had .254 which I think you have as the internet router's IP (RV042)??):
V042 Config: Setup-> Advanced Routing-> Static Routing
Destination IP: 192.168.20.0
Subnet Mask: 255.255.255.0
Default Gateway: 192.168.9.254 <<< Should be the VLAN router's 192.168.9.xxx address .253?
Hop Count: 1
Interface: LAN
Destination IP: 192.168.60.0
Subnet Mask: 255.255.255.0
Default Gateway: 192.168.9.254 <<< Should be the VLAN router's 192.168.9.xxx address .253?
Hop Count: 1
Interface: LAN
Destination IP: 192.168.90.0
Subnet Mask: 255.255.255.0
Default Gateway: 192.168.9.254 <<< Should be the VLAN router's 192.168.9.xxx address .253?
Hop Count: 1
Interface: LAN
Or you could just add one statement like (192.168.0.0 with a subnet of 255.255.0.0 to 192.168.9.253) if no other 192.168.xxx.xxx addresses are used on the internet router/fw.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-18-2013 03:57 AM
for this case maybe i can suggest you can buy a router cisco RV series support multisubnet. (before buy can ask the store,maybe they can give you suggestion)
on router config needed:
1.config a multi subnet on Router Cis" RV series as you want
2.config port trunk on router(untagged)port as you want 
3.dont forget (check the routing table & try ping)
4.connect the router with sg500(plug on trunk port)
on switch config needed:
1.config sg500 as router mode,
2.setting vlan&ip on each port as you want
3.setting trunk than connect the cable to router
4.tried ping for each port&ip
hope this info helpful
thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 09:48 AM
Did you ever solve this problem? I have the exact same problem and have followed all of the replies and tried all of the steps already mentioned with no success.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 01:46 PM
Jeremy,
What problem are you having specifically? For me, the next hop was incorrect and I could not reach the router from any of the devices in any of the VLANs.
Here is a link to the post I created on the side.
Let me know if you have any troubles with it.
Johnny
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-12-2018 03:55 PM
I have finally fixed it after a year of working on it. As of today everything works! The big thing that I was missing was the multiple subnet feature on the RV 042G. Any configuration using multiple subnets downstream from the RV 042G must have the multiple subnet feature enabled and all subnets added that need edge access. No one seem to be able to provide this answer on the Internet and so I am going to be posting the information somewhere where others can benefit. Many people helped and corrected some things that I had very wrong but that final answer that is specific to the RV was the stumper.
