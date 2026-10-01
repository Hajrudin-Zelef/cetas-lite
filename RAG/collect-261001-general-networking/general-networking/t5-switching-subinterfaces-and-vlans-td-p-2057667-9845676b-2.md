---
id: collect-261001-general-networking/general-networking/t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b-2
title: "t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b.md
source_anchor: ""
source_lines: [19, 175]
sha256: 9d8bf0923eca44bf4fd9b603104e81ddf4f8197711a6a76bd6cc0453031f9a7a
---

# t5-switching-subinterfaces-and-vlans-td-p-2057667-9845676b

			LAN Switching
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2012 01:59 PM
Yes you are correct. if you are using layer 2 switch and you want to do intervlan routing then you need layer 3 device like router. But you need to configure sub interfaces with default gateway to route the traffic. Because there is one trunk between swich and router so we need sub interfaces for multiple vlans.
Interface FastEthernet0/0.1
Encapsulation dot1q 10 (10 represent VLAN ID 10 )
IP address 10.1.1.1 255.255.255.0
If you are using a layer 3 switch then you dot need any sub interfaces so then you can create vlan interface with the default gateway. You need to enable ip routing first.
Interface vlan 10
IP address 10.1.1.1 255.255.255.0
Hope this will help.
Please rate if this helps.
thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2012 01:59 PM
Yes you are correct. if you are using layer 2 switch and you want to do intervlan routing then you need layer 3 device like router. But you need to configure sub interfaces with default gateway to route the traffic. Because there is one trunk between swich and router so we need sub interfaces for multiple vlans.
Interface FastEthernet0/0.1
Encapsulation dot1q 10 (10 represent VLAN ID 10 )
IP address 10.1.1.1 255.255.255.0
If you are using a layer 3 switch then you dot need any sub interfaces so then you can create vlan interface with the default gateway. You need to enable ip routing first.
Interface vlan 10
IP address 10.1.1.1 255.255.255.0
Hope this will help.
Please rate if this helps.
thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2012 09:46 PM
Amrinder,
Thanks for your brief explanation! It really cleared things up.. So let me get this right! The trunking protocol tags frames and sends them through the trunked port but the router doesn't know what to do with the tagged frames. But by creating subinterfaces, encapsulating them with 802.1Q, and using the VLAN ID.. this tells the router which vlan tagged frames belong to? Then with a routing protocol (OSPF) I would need to advertise all the VLAN networks within the same router to make intervlan routing possible. Is this correct?
You also saved me another question because we have a layer 3 core switch. It was going to be how intervlan routing will work on a layer 3 switch.. but I see that when you create the VLANS you also add an IP address.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2012 10:35 PM
Hello Miguel,
Yes 802.1q and vlan ID route the traffic to its relevant vlan. You can use OSPF, when a Host can ping its default gateway because it is on the same local subnet. Host can ping both switches because the management interfaces are set to VLAN1. Because a host does not have a route to get to the other VLANs/subnets, it forwards the packets to its default gateway,. Although router has a route to get to the majority of the other subnets, remember that the Internet Control Message Protocol (ICMP) packets need to return as well. if a router has two directly connected routes in the routing table, but no routing protocols or static routes are configured to facilitate communicating from one network to another. You should configure Open Shortest Path First (OSPF) as the routing protocol to allow inter-VLAN routing.
http://www.informit.com/library/content.aspx?b=CCNP_Studies_Troubleshooting&seqNum=77
For layer 3 switch. You need to enable routing first with command IP Routing. You can’t assign IP address to an interface unless you enter No switchport to a interface where you want to use Ip address. By default every port is in layer 2 mode and you need to change it to layer 3 by issue no switchport command . then you will be able to assign IP.
thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2012 03:16 PM
Hi Miguel De Santiago
My name is Johnnatan Rodriguez, your information about the vlan's is correct, now I'm goint to explain to you, why does the router need subinterfaces?
Subinterfaces: we need it when we have more vlans than physical links
As you said, we have setup a trunking port between the switch and router, then configure the configure the interface and sub interfaces in your router:
Switch(config)#interface fa0/2
Switch(config-if)#switchport mode trunk
Router(config)#interface fa0/1
Router(config-if)#no shutdown
Router(config)#interface fa0/1.1
Router(config-subif)#encapsulation dot1q 1
Router(config-subif)#ip address 192.168.10.1 255.255.255.0
Router(config)#interface fa0/1.2
Router(config-subif)#encapsulation dot1q 15
Router(config-subif)#ip address 192.168.15.1 255.255.255.0
Router(config)#interface fa0/1.3
Router(config-subif)#encapsulation dot1q 35
Router(config-subif)#ip address 192.168.20.1 255.255.255.0
We configure sub interfaces because we have 3 vlan and just one physical link, for this reason we need that all data passing through a single link, how do we fix that?
Creating 3 sub-interfaces, one for each vlan, however it reduces the bandwidth one third.
Physical Interfaces: we need it when we have more physical links than vlans
We configure a normal ip in each interface, (one per vlan), now in the switch instead of create trunk links, we create access links (one per vlan).
Router(config)#interface FastEthernet0/0
Router(config-subif)# ip address 10.10.10.1 255.255.255.0
Router(config)interface FastEthernet0/1
Router(config-subif)# ip address 10.10.20.1 255.255.255.0
Router(config)#interface FastEthernet0/3
Router(config-subif)#ip address 10.10.30.1 255.255.255.0
Switch(config)#interface range fa0/1, fa0/10
Switch(config-if)#switchport mode access
Switch(config-if)#switchport access vlan 10
Switch(config)#interface range fa0/5,fa0/20
Switch(config-if)#switchport mode access
Switch(config-if)#switchport access vlan 20
Switch(config)#interface range fa0/8, fa0/30
Switch(config-if)#switchport mode access
Switch(config-if)#switchport access vlan 30
Here we have one link per vlan and we can use the full bandwidth of each interface.
I hope you find this answer useful, we will help you with any doubt that you have, if you found this answer useful please mark the question as Answered and rate the anwer.
Thanks for using our forum.
Greetings,
Johnnatan Rodriguez Miranda.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2012 08:35 PM
hi johnnatan,
nicely done! +5
i felt like attending a CCNA class. keep it up!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-01-2012 09:51 PM
Johnnatan,
You post was very helpful also and it all makes sence. I was, however, not aware that by creating subinterfaces on one physical link, it would reduce or split the bandwidth! It makes sense that this would occur but I guess overlooking the easy stuff is easy to!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-04-2013 04:37 AM
Johnnatan, great answer.
One question:
Do you configure fa 0/1,5 and 8 as access mode because there is only one VLAN per port? And if the answer is yes, am I right in saying this is a valid configuration because only one adjacency is made (maximum for this port)?
Thank you for clarifying in advance. I am just confused as to why you wouldn't configure the port as a trunk, even if it only does have one VLAN across it
Regards,
pp
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-08-2013 12:10 PM
Hi,
Explaining the scenario:
When a port connected to 1 host belonging to a VLAN is access and not trunk.
