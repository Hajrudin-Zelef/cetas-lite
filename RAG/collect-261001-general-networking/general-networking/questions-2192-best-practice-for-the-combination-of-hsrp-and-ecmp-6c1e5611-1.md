---
id: collect-261001-general-networking/general-networking/questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611-1
title: "questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "research"]
source: docs/RAG/collect-261001-general-networking/questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611.md
source_anchor: ""
source_lines: [1, 75]
sha256: 7599eb115ca816003040476ad00bed0a376ad18edd56db0e2c1ada70a8647dbb
---

# questions-2192-best-practice-for-the-combination-of-hsrp-and-ecmp-6c1e5611

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
18
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
The combination of ECMP (or other causes of asymmetric paths) and HSRP is broken by default in Cisco IOS; the default behaviour with this design floods unicast traffic excessively.
What is the best-practice for using HSRP with ECMP to prevent unknown unicast flooding?
Details / Background
We have a HSRP topology similar to the first diagram below for many of our facilities. Our Cisco WAN routers have equal-cost routes to all other sites; thus we can see asymmetric routing effects all the time. Normally we assign R1 to be the HSRP primary, but ECMP allow return traffic through either R1 or R2.
The issue is that when PC1 mounts a remote iSCSI drive across the WAN, the traffic leaves the site via R1, but could return via R2. As long as the iSCSI traffic returns via R1, there are no issues.
The problem occurs when PC1's traffic returns via R2. Assume the iSCSI session starts at 8:00:00, and both routers and both switches learn PC1's mac simultaneously. Between 8:00:00 and 8:00:05, there are no flooding problems because both switches still have PC1's mac-address in their CAM table.
Five minutes after the iSCSI session starts, S2's CAM entry for PC1's mac expires out of the CAM table and S2 floods PC1's traffic out all ports (in this case to Po1, Gi0/3 and Gi0/4). If PC1's iSCSI session consumes a lot of bandwidth, this unknown unicast flooding can suck non-trivial capacity from the links to PC3 and PC4.
Cisco IOS switches have a default CAM timer of 300 seconds...
S2# show mac address-table aging-time
Vlan Aging Time
---- ----------
1 300
17 300
However, Cisco IOS' default interface ARP timer is 4 hours...
R2# show interface gi0/0
GigabitEthernet0/0 is up, line protocol is up
Hardware is AmdP2, address is 000a.dead.beef (bia 000a.dead.beef)
Internet address is 172.17.1.252/24
MTU 1500 bytes, BW 10000 Kbit, DLY 1000 usec,
reliability 255/255, txload 1/255, rxload 1/255
Encapsulation ARPA, loopback not set
Keepalive set (10 sec)
ARP type: ARPA, ARP Timeout 04:00:00 <--------------
Therefore, S2 starts flooding PC1's iSCSI traffic after five minutes.
The simple answer is to make the CAM timer equal or slightly longer than the corresponding interface ARP timer, but there are at least three different options to select from...
Option 1: Lower all interface ARP Timers
This option works best if you have a decent-sized layer2 switched network, a reasonable number of ARP entries and few routed interfaces. This method also is preferable if you like to see PC mac entries age out of the topology quickly.
On all IOS ethernet interfaces facing an ethernet switch: arp timeout 240
On all IOS ethernet interfaces facing an ethernet switch: hold-queue 200 in and hold-queue 200 out to avoid dropping ARP packets during periodic ARP-refreshes (these limits could be higher, or lower depending on how many ARP refreshes you think that you'll need to handle at once). If you are adjusting Selective Packet Discard values, then you should follow the guidelines in the paper I linked.
This forces Cisco IOS to refresh the ARP table within four minutes, if it hasn't happened otherwise for a given ARP entry. The obvious disadvantage is that this doesn't scale well if you have lots of ARP entries... the limits vary by platform. I have used this with a few hundred ARPs per router on Catalyst 4500 / 6500 (the Layer3 SVIs) without any issues.
Option 2: Increase the switch CAM Timers
This option works best if you have a large number of ARP entries (i.e. thousands, such as an intense VMWare environment could see).
On all IOS switches: mac address-table aging-time 14400, or mac address-table aging-time 14400 vlan <vlan-id> for any Vlan that is of concern.
This change adjusts timers that most people assume are fixed at 300 seconds (on Cisco IOS), so be sure to include this in continuity docs. The side-effect of this is that CAM table entries linger for 4 hours after the PC is removed (which can be either good or bad, depending on your PoV). If 4 hours is too long, see the next option...
Option 3: Change both the interface ARP timers, and the switch CAM Timers
This option avoids hideously-long CAM timers in Option 2 at the expense of more configuration. You can choose whether you need 900 seconds, 1800 seconds, or whatever... just make sure your CAM and ARP timers match; thus, you will need to configure both Option 1 and Option 2 in your topologies.
To me, ECMP is the real issue here - so in addition to the above steps to limit unknown unicast flooding, you can also tune the route metrics toward the WAN so that R1 is preferred over R2 for the return traffic. One way to achieve this is via distribute list on R2 as follows:
(EIGRP used for example only, the same can be achieved with OSPF or BGP with other commands)
!
ip prefix-list R1-PREFER seq 5 permit 172.17.1.0/24
!
route-map R1-PREFER-MAP permit 10
match ip address prefix-list R1-PREFER
set metric 1 1 1 1 1
... (allow all other routes)
!
router eigrp 1
....
distribute-list route-map R1-PREFER-MAP out Ser1/0
....
!
This will result in the WAN forwarding all traffic for 172.17.1.0 to R1. If R1 Se1/0 fails, the route will be installed toward R2. You can further tune these metrics so the backup route to R2 is actually a feasible successor for faster failover. HSRP and tracking will take care of the egress traffic.
The idea if not use ECMP if HSRP is in use may be ok for SERVERS where ingress traffic may be higher than egress traffic, in a PC situation IN GENERAL ingress traffic from the WAN (responses) is higher than egress traffic (ingress). We like most people just set the ARP timers. you can mess with CAM timers BUT if you have say an MDF with the layer 3 switch and an IDF with 2 collection switches and say 5 access switches, it is a LOT easer to configure on the L3 SVI than doing all access switches.
Ah, I remember this one. Weeks of fun was had dealing with this back a few jobs ago. One wrinkle is that STP events will put the vlans in fast-aging mode, so setting the MAC timer longer than the ARP timer doesn't help
I solved the issue by forcing ECMP back from the servers, by creating two floating HSRP gateways, with one primary on each router. We then configured both gateways on each host. By forcing host traffic to both R1 and R2 in this fashion, we would be sure that R2 would never age out the MAC addresses.
Ideally, this wouldn't be an issue if L2/3 switches purged ARP entries associated with aged-out MAC addresses. The next packet to the IP would then result in a new ARP request, populating both the ARP cache and the MAC table. I think Cisco eventually implemented this, but I can't say for sure.
I've never been a fan of tweaking timers. Timers are set a certain way usually for many reasons. Altering them:
is potentially operationally intensive to maintain everywhere the same
makes things more complicated and difficult to troubleshoot
as a recent commentor showed, can have unexpected side effects
may not "play nice" with future Cisco enhancements
Alternately:
Use MC-LAG (aka "MEC" in Cisco documentation). This is your best option, though you should understand the deployment scenarios where MC-LAG can be used (its not a universal solution, and should only be deployed after appropriate research and testing). MC-LAG variants are hardware dependent. Examples are:
a. Stacking (Cat 3k)
b. VSS (Cat4k/6k)
c. VPC (Nexus)
d. Pseudo mLACP (ASR1k)
e. MC-LAG (ASR9k)
f. Clustering (ASA)
