---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425-3
title: "t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425.md
source_anchor: ""
source_lines: [203, 313]
sha256: 3794291f17ea525412cc6a5d5e1e17cfb358c809e8af968ef8771dc79d7d393d
---

# t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425

if that is the case, one link to the ISP, the routing will execute in the order I outlined in my first answer. OSPF first, then BGP, then the static route. As everything goes over one physical link, you don't really have a 'failover'.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 11:43 AM
Hi,
I need to understand this ,
If isp1 link goes down , this means static default route configured in router 1 will no longer works.
In ospf , we are redistributing static and connected but we are also using specific network by mentioning them
redistribute connected subnets
redistribute static subnets
passive-interface default
no passive-interface Port-channel1.11
network 10.2.2.0 0.0.0.255 area 0
network 10.2.3.0 0.0.0.255 area 0
network 10.2.4.0 0.0.0.255 area 0
network 10.2.5.0 0.0.0.255 area 0
So how does.default route will redistrbute because if I run show ip route ospf ,it does not show default route.
If I run show ip route , default route is static pointing to isp1.
From the setup , i understand that purpose of setting ospf is to act as a carrier for bgp.
I am still not able to understand , if isp1 goes down , what will happen to
Ospf.
Bgp.
Default route .
Traffic flows through below order if isp1 goes down.
Switch in dc1 >router1vlan 14>backto switch dc1 ( tunneled in vlan 14).>Switch dc2 ( tunnels in vlan 14) > router in dc2 vlan 14> isp2
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 12:37 PM
Hello,
in OSPF, you cannot redistribute a static default route. You need to use the 'default-information originate' command (since you have a default route configured), or the 'default-information originate always' command (which will generate a default route without one actually existing).
If your physical link goes down, neither your BGP, OSPF, or the static route will work.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 12:45 PM
Hello , This setup did work .
It was tested 2 years back and ISP1 link was switched off and traffic moved dynamically to ISP2
I know my knowledge of routing is limited . but what is the point of OSPF and BGP if traffic wont shift to another DC .
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 01:20 PM
Hello,
as I said earlier, it is unclear from your topology drawing how (and if) the two routers are connected. If they are, the backup (failover) will work. If possible, post the full configs of both routers so we can see how this is set up.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 01:53 PM
OSPF will redistribute whatever you'll tell it, however Georg is right and it wouldn't redistribute default route. You have to have "default-information originate" under OSPF configuration to make sure default route got advertised to neighbors. Moreover - static route is not reliable as it is not detecting the peer status. You would have a valid static route as long as interface is in the up/up state.
Ideally you want to run BGP with your ISP and then make sure the 0.0.0.0/0 goes into OSPF and get redistributed between DCs. It is a bit tricky since OSPF has a better administrative distance so in one of the DCs it would prefer OSPF route over BGP route. You can overcome this by altering AD from specific gateway using route-map to match specific route.
Another option is to use "ip sla" with "tracking" option. In that case you can detect your ISP's PE router failure even if the interface status is in the "up/up" state.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 02:47 PM
Hello , I am attaching the config of both the routers . I have removed the Public IP address of ISP and all other private details .
Regarding connection :
Router 1 connect to physically core switch 1 in DC1 . Router 1 also connected to ISP1 on a separate physical interface
Router 2 connect to Physically core switch 2 in DC2 Router 2 also connected to ISP2 on a separate physical interface
and both the core switches connect via Physical links .
So any vlan created on router1 will flow via cross DC physical link and extend over to Router 2 .
In this case vlan 14( GLB vlan ) and vlan 15 ( OSPF vlan) flows over cross DC link . The BGP uses the same loopback and flows over vlan 15
The goal of this conversation is to validate what happen if only and only ISP1 link goes down - Router 1 and Router 2 will still remain active and how traffic flows via OSPF/BGP . Will it dynamically send default route to Router 2 or manual intervention is required
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 03:07 PM
There is much that is not clear about this environment and that makes it difficult to give good advice. Some clarification of the environment would be helpful. But one thing stands out to me and that is that the posted partial config shows that the router has a static default route. In my experience the main reason for running BGP with an ISP to to allow the ISP to advertise a default route. And loss of the advertised default route from the ISP is the main mechanism for failover. But with this router and its static default route you could lose access to the ISP but the static default route might very well remain in the routing table and you would not fail over.
The first thing that I would suggest is to remove the static default route. Verify that your ISP is advertising a default route to you, and if so trust the ISP default route. Beyond that there are other things to figure out including whether you want the ISPs to operate as Primary and Backup or as both active and sharing the load, and how OSPF will treat the default route.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 03:13 PM
Hi Richard , i have attached the router config in the post trail
The special thing is that BGP peer group is being used here . So if ISP1 goes down ; BGP peer group ( having both R1 and R2 loopbacks) will be updated .
Also , how the default route will be propagated ( if ) via OSPF/BGP if ISP1 goes down .
