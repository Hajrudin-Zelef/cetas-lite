---
id: collect-261001-general-networking/general-networking/t5-routing-how-to-make-isp-b-active-using-eigrp-hsrp-on-the-below-scenario-m-p-1-7100d200
title: "t5-routing-how-to-make-isp-b-active-using-eigrp-hsrp-on-the-below-scenario-m-p-1-7100d200"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/t5-routing-how-to-make-isp-b-active-using-eigrp-hsrp-on-the-below-scenario-m-p-1-7100d200.md
source_anchor: ""
source_lines: [1, 187]
sha256: 93600e358246f41507c75222e5aeca8267fe56174f7412fbfd0a59d741f87738
---

# t5-routing-how-to-make-isp-b-active-using-eigrp-hsrp-on-the-below-scenario-m-p-1-7100d200

How to make ISP B active using EIGRP & HSRP on the below scenario
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2009 08:31 AM - edited 03-04-2019 01:06 AM
Hi,
Please help me to configure ISP B using Eigrp ( metirc) and HSRP as the primary Link , such that traffic should only pass through it.
Current scenario
L2 Core switch --vlan 905-----ISP A
L2 Core Switch --vlan 906---------ISP B
Now L2 Core switch connected to HSRP Two Routers ( DDC Router A& Router B ).
HSRP is configured between Router A & B.
Currently Router A is given higher priority of 110.
Also let me know in this current scenario Eigrp is performing Load balancing ? if not please let me know the config.
+++++++++++++++++++++++++++++++++++++++++
L2 Core switchDDC1#interface GigabitEthernet5/7
description ISPA MAN link DDC1-DDC3
switchport access vlan 905
switchport mode access
speed 100
duplex full
no cdp enable
!
interface GigabitEthernet5/8
description ISPB Link DDC1-DDC3
switchport access vlan 906
switchport mode access
speed 100
duplex full
no cdp enable
++++++++++++++++++++++++++++++++++++
Router A#HSRP Priamry IP - 10.189.40.44 ( Standby ip - 10.189.40.40)-----
interface GigabitEthernet0/1.905
description "Primary Link-DDC1-DDC3_ISP A
encapsulation dot1Q 905
ip address 10.189.223.209 255.255.255.252
no ip redirects
no ip unreachables
no ip proxy-arp
ip nbar protocol-discovery
ip virtual-reassembly
no snmp trap link-status
no cdp enable
router eigrp 140
redistribute connected
redistribute static
network 10.189.40.44 0.0.0.1
network 10.189.223.192 0.0.0.3
network 10.189.223.200 0.0.0.3
network 10.189.223.208 0.0.0.3
no auto-summary
++++++++++++++++++++++++++++++++++++
RouterB# HSRP secondary IP -10.189.40.45
interface GigabitEthernet0/1.906
description "Secondary Link-DDC1-DDC3_ISP B
bandwidth 100000
encapsulation dot1Q 906
ip address 10.189.223.213 255.255.255.252
no ip redirects
no ip unreachables
no ip proxy-arp
ip nbar protocol-discovery
ip virtual-reassembly
no snmp trap link-status
no cdp enable
router eigrp 140
redistribute connected
redistribute static
network 10.189.40.44 0.0.0.1
network 10.189.223.192 0.0.0.31
no auto-summary
- Labels:
- 
						
							
		
			Routing Protocols
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2009 09:02 AM
Hello Mirza,
for HSRP you should use tracking but with object tracking to track ISPA next hop ip address or ip routing on interface g0/1.906
for EIGRP you can make R2 the less preferred path by increasing delay on the client vlan on R2.
In this way R2 EIGRP update will be worse then R1 update about the same subnets (this implies both providers can carry EIGRP routes to remote/central site)
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2009 10:51 AM
Thanks for your comments!
I could not get your above Post!
L2 Core switch both the routers are connected to the same switch .
L2Core switch1 ---> Router A with Vlan 905
L2Core switch 1---> RouterB with Vlan 906.
Can we do perform some thing on Eigrp ( metric/cost part) by keeping Vlan config same.
Bcoz at persent i have some intermittent issue with ISP A and hence want all traiffc on ISPB.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2009 12:03 PM
Hello Mirza,
I mean on the one you want less preferred
R1 - ISPA
let's suppose vlan 200 is your client vlan
int gi0/0.200
delay 10000
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2009 12:28 PM
Hello Giuslar,
I want for some period both Vlan 905 & 906 at L2 shld pass on ISP B to reach their same destination.
Vlan 905 & 906 are reaching to pa particular destination.
Can you expalin me in detail!
For you above comments - what shld be the delay i have to setup for ISP B ....
shld Delay on ISP B be more or less to get Priority.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2009 12:35 PM
Hello Mirza,
EIGRP metric is cumulative on delay so you need to increase delay on the interfaces you want to make less preferred
you can do this on the client Vlans and/or the core facing links
in this way you should make the routes via R2 and via ISPB preferred
it can help to increase delay also on the interface from ISP1 on remote site
Hope to help
Giuseppe
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
02-04-2009 01:01 PM
Hi Guislar,
Now i am able to catch your point on the delay part of Eigrp.
So you mean my both vlans will not at all go to ISP A path to reach destination?
Also can we achieve my task using Eigrp Metrics?
Can we achive this by any chance with HSRP.
Based on my first Post config, is it a Loadbalancing configured at l2 by having two seperate vlans for both ISP ...from switch level
Here i get a question like how does HSRP works between two routers when both the routers are allowed only by a single Vlan at a time.
Router A - vlan 905
Router B - vlan 906.
Please clariefy me.
Many thanks for the posts.
