---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-eigrp-hsrp-m-p-695206-highlight-true-7770ce1b
title: "t5-routing-and-sd-wan-eigrp-hsrp-m-p-695206-highlight-true-7770ce1b"
domain: general-networking
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-eigrp-hsrp-m-p-695206-highlight-true-7770ce1b.md
source_anchor: ""
source_lines: [1, 171]
sha256: ec2339e5da02b6107046470cb160662ed9fe31383473c029ce44d44e1e72efd9
---

# t5-routing-and-sd-wan-eigrp-hsrp-m-p-695206-highlight-true-7770ce1b

EIGRP & HSRP
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-21-2006 11:18 PM - edited 03-03-2019 02:05 PM
Can you please help.
Our satellite office is to be split into 4 subnets.
We have a frame-relay network and are suing EIGRP, and are looking to load share using EIGRP.
2 routers at the satellite connecting to 3 routers at the head office over 4 pvcs.
We intend to route 2 subnets down each router. i.e. 2 subnets down r1 to r3 and r4 and 2 subnets down r2 to r4 and r5.
Using HSRP we will have 2 groups so that 2 subnets are active on one router and passive on the other and vice versa.
What I want to know is will the routing tables in r1 and r2 advertise just the active subnets that it is routing or because it knows of all 4 subnets will it advertise all 4 of them.
Hope this make sense.
Any pointers to documentation much appreciated.
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
09-22-2006 01:32 AM
Hello,
basically, EIGRP will advertise whatever interface you specify with the 'network' statement, provided the interface is up/up.
From what you are saying, you have r1 and r2 at the satellite location, and r3, r4, and r5 at the head office location ?
If that is correct, your configuration could look like this:
hostname r1
!
interface ethernet 0
ip address 192.168.1.3 255.255.255.0
standby 1 ip 192.168.1.1
standby 1 priority 110
standby 1 preempt
standby 2 ip 192.168.1.2
standby 2 preempt
!
hostname r2
!
interface ethernet 0
ip address 192.168.1.4 255.0.0.0
standby 1 ip 192.168.1.1
standby 1 preempt
standby 2 ip 192.168.1.2
standby 2 priority 110
standby 2 preempt
Now, if you give half of your hosts at the satellite connection the IP address of standby group 1 as their default gateway, and the other half of your hosts the IP address of standby group 2, you will have a fault-tolerant load balancing configuration.
Have a look at the document below as well. If possible, can you post the full configurations of your routers ?
Using HSRP for Fault-Tolerant IP Routing
http://www.cisco.com/en/US/tech/tk1330/technologies_design_guide_chapter09186a008066670b.html
HTH,
GNT
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-22-2006 02:02 AM
Thank you for your reply.
I can see how we could use the configurations you have included.
Unfortunately, behind each router there are switches and we need these in different vlans, due to some legacy stuff, also the PCs in the satellite office use dhcp.
Just looking whether responses from the server in the head office will be routed back to the satellite office correctly bu just using the routing tables forwarded by the satellite office.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-22-2006 02:54 AM
Andrew
There is not an impact on what EIGRP will advertise from running HSRP. EIGRP will advertise subnets based on the network statements that are specified and has no consideration for whether HSRP is leader or standby.
You do raise a valid point that HSRP will help steer traffic to R1 or R2 to be forwarded to Head Office. But routing traffic from HO to the Satellite Office will not be affected by HSRP. So you need something to modify routing at HO. I would suggest offset lists on R1 and R2. Your config might look something like this:
hostname R1
interface fastethernet 0/0.32
ip address 172.18.32.2 255.255.224.0
standby 1 ip 172.18.32.1
standby 1 priority 105
standby 1 preempt
interface fastethernet 0/0.64
ip address 172.18.64.2 255.255.224.0
standby 2 ip 172.18.64.1
standby 2 priority 105
standby 2 preempt
interface fastethernet 0/0.96
ip address 172.18.96.2 255.255.224.0
standby 3 ip 172.18.96.1
standby 3 priority 101
standby 3 preempt
interface fastethernet 0/0.128
ip address 172.18.128.2 255.255.224.0
standby 4 ip 172.18.128.1
standby 4 priority 101
standby 4 preempt
!
router eigrp 100
network 172.18.32.0 0.0.31.255
network 172.18.64.0 0.0.31.255
network 172.18.96.0 0.0.31.255
network 172.18.128.0 0.0.31.255
offset-list 10 out 1000
access-list 10 permit 172.18.96.0
access-list 10 permit 172.18.128.0
!
hostname R2
interface fastethernet 0/0.32
ip address 172.18.32.3 255.255.224.0
standby 1 ip 172.18.32.1
standby 1 priority 101
standby 1 preempt
interface fastethernet 0/0.64
ip address 172.18.64.3 255.255.224.0
standby 2 ip 172.18.64.1
standby 2 priority 101
standby 2 preempt
interface fastethernet 0/0.96
ip address 172.18.96.3 255.255.224.0
standby 3 ip 172.18.96.1
standby 3 priority 105
standby 3 preempt
interface fastethernet 0/0.128
ip address 172.18.128.3 255.255.224.0
standby 4 ip 172.18.128.1
standby 4 priority 105
standby 4 preempt
!
router eigrp 100
network 172.18.32.0 0.0.31.255
network 172.18.64.0 0.0.31.255
network 172.18.96.0 0.0.31.255
network 172.18.128.0 0.0.31.255
offset-list 10 out 1000
access-list 10 permit 172.18.32.0
access-list 10 permit 172.18.64.0
HTH
Rick
Rick
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-22-2006 04:24 AM
Thank you for the reply and advice, much appreciated.
Andy
