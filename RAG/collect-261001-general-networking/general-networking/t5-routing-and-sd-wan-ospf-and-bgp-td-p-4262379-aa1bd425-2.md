---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425-2
title: "t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425.md
source_anchor: ""
source_lines: [95, 202]
sha256: 9c5c69ee39e53a73b0df3cbc0e1f678ab06efd232eaf82887b5d348f8cb96738
---

# t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425

			Routing Protocols
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-27-2020 03:22 PM
There are several things I would like to address. First you ask this question "Cant we use simply the VLAN 15 address as loopback." Perhaps it is simply a confusion about terminology. A loopback is a particular type of virtual interface. A vlan interface is a particular type of virtual interface. And they are not interchangeable. A vlan interface can not be a loopback and a loopback interface can not be a vlan.
Then let me address this statement that you make "OSPF is having two routes - the vlan 15 interface and ISP address" This is not correct. It is not the vlan 15 address that is advertised but is the loopback interface address that is used. (note that the loopback interface address is used both as the OSPF Router ID and used as the IBGP peer address.
Then let me try to clarify my statement about OSPF and IBGP. Looking into the configuration of IBGP we find that one router defines the IBGP neighbor as 10.2.52.240. This address is the loopback interface address of the peer router. To form the IBGP peer relationship the router must know how to reach the peer address. So how does this router know how to reach 10.2.52.240? That address is advertised by OSPF. So this is the basis of my statement about OSPF and IBGP. If OSPF did not advertise that address then IBGP would not work. So OSPF is necessary for this implementation of IBGP.
It does not have to be done this way. I suggest an alternative: configure IBGP so that the IBGP neighbor address is the vlan 15 address of the peer router (rather than as the loopback address). If the neighbor address is the vlan 15 address then the router knows how to reach that address (it is in a directly connected subnet) and there is no need for OSPF.
This is one example of making things more complex than they need to be. There are some good reasons why you might choose to make the IBGP neighbor address be on a loopback interface. But those reasons are not present in this network. By changing the IBGP neighbor address to the vlan interface rather than the loopback interface we remove the requirement for OSPF and make the configuration more simple.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-28-2020 09:36 AM
You are welcome. It has been a long discussion and I am glad that our explanations have been helpful. This community is an excellent place to ask questions and to learn about networking. I hope to see you continue to be active in the community.
I tried to explain about BGP neighbors using loopback addresses or not in a previous response. It seems that was not so clear so let me try again from a slightly different perspective. Let us think about 2 routers, routerA and routerB who want to become BGP neighbors (it might be IBGP or it might be EBGP, same things apply to both). Let us assume that the routers are connected using vlan 15. The simple way is for both routers in their bgp neighbor commands to use the vlan 15 interface IP address. When it is configured this way there is no need for any other routing protocol to be running on the routers. The routers are attempting to access the neighbor on a connected subnet and all they need to do is to arp for each other. When this is configured the BGP neighbor relationship will be negotiated and they become active neighbors and remain active neighbors as long as vlan 15 continues to work. But if something happens on vlan 15 and it stops working then the BGP neighbor relationship is terminated.
That is the more simple approach for BGP neighbors (to use directly connected subnet addresses as neighbor address). But let us think about routerA and routerB and what might happen if both routers are also connected using vlan 10 (in addition to vlan 15). If the routers continue to use the subnet address of vlan 15 as the neighbor address then they are dependent on vlan 15 to maintain the neighbor relationship. But if the routers were to use some IP address on the other router that was reachable using either vlan 10 or vlan 15 then the routers are not dependent on a single connection and can take advantage of the redundancy of their connections. A loopback interface address is frequently used for this.
So using a loopback interface address as the BGP neighbor address takes advantage of possible redundant connections. Since the routers are now using an address that is not directly connected they can no longer simply arp for the neighbor address but must have some routing information about how to reach the neighbor. It might use some routing protocol like OSPF or might use something simple like static routes.
So configuring BGP neighbor statements with loopback interfaces is more complicated but takes advantage of potential redundancy. And configuring BGP neighbor statements with a connected subnet address is more simple, but does not provide redundancy. It is a choice to be made when the network design is being done.
So let us now think about your discussion. Your routers were configured with loopback interface addresses in the BGP neighbor statements, which suggests redundancy. And OSPF was being used and the loopback interface addresses were being advertised which also suggests redundancy. But you were using only vlan 15. It was quite clear that in the network design of your network that vlan 14 was for inside traffic and vlan 15 was for outside traffic (like BGP). So you had the more complex configuration (loopback interfaces and OSPF) but were using only a single connection. So my advice was to use the more simple configuration of BGP neighbor using the vlan 15 interface addresses.
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 06:50 AM
Hello,
what are the two DC routers connected to ? An overview of your entire topology would help. If there is only one uplink to one ISP, there is no failover...everything will fail regardless of which routing protocol is being used.
In case there are different links connecting through OSPF and iBGP, OSPF will be selected first, if that (link) fails, iBGP will be selected, and if that fails, your static default route.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 07:17 AM
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 07:20 AM
Subinterface 11 is configured for OSPF
encapsulation dot1Q 11
ip address 10.2.2.2 255.255.255.248
ip ospf message-digest-key 1 md5 7 0468031357205E1F2D18
!
router ospf 11
router-id 10.2.2.10
area 0 authentication message-digest
redistribute connected subnets
redistribute static subnets
passive-interface default
no passive-interface Port-channel1.11
network 10.2.2.0 0.0.0.255 area 0
network 10.2.3.0 0.0.0.255 area 0
network 10.2.4.0 0.0.0.255 area 0
network 10.2.5.0 0.0.0.255 area 0
!
My query is if ISP link in DC1 goes down , what OSPF will do in that case
Also , BGP is between Router peer group and ISP . and BGP is in between two routers also
how OSPF and BGP work in this case in case ISP1 goes down .
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 08:11 AM
Hello,
I am not sure I fully understand your topology, sorry for that. Do you have only one link to ISP1 and ISP2 respectively from each router ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 08:19 AM
Hello,
Yes link to isp1 from router1
Link to isp2 from router 2
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 11:29 AM
Hello,
