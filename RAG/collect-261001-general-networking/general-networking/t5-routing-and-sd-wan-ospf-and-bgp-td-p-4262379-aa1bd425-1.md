---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425-1
title: "t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425.md
source_anchor: ""
source_lines: [1, 94]
sha256: c1a705713ef987e7b2d0fd88c9151b45a10a92512833141c038e868409914485
---

# t5-routing-and-sd-wan-ospf-and-bgp-td-p-4262379-aa1bd425

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
12-22-2020 05:54 AM
FOr one of our customer , they have two router one in each DC
The router in DC1 is connected to COre Switch in DC1 >> This is a Layer 2 connection
Similarly router in DC2 is connected to Core switch in DC 2 ( again L2 connection)
The DCs are connected over a Physical underground link .
Between Router 1 and Router 2 , we have defined a VLAN 11 where we run OSPF . BGP is also configured between two routers.
They use BGP peer Group . I want to understand how the failover will work in case ISP link goes down .
Both the router have a default route pointing to their respective ISP .
In OspF it have defined the redistribute connected and static plus network address for4 subnets also .
Below is the config
interface Port-channel1.11
encapsulation dot1Q 11
ip address 10.2.2.2 255.255.255.248
ip ospf message-digest-key 1 md5 7 XXXXXXXXXXXXXXX
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
router bgp 27272
bgp router-id 10.2.2.10
bgp log-neighbor-changes
bgp graceful-restart
bgp maxas-limit 100
timers bgp 10 30
neighbor TEST_GROUP peer-group
neighbor TEST_GROUP remote-as 27272
neighbor TEST_GROUP password 7 XXXXXXXXXXXXXXXXXX
neighbor TEST_GROUP update-source Loopback0
neighbor 10.2.2.11 peer-group TEST_GROUP   >>>> this IS router 2 loop back address in DC2
neighbor 22.2.22.12 remote-as XXXXX
neighbor 22.2.22.12 description ISP
neighbor 22.2.22.12 password 7 XXXX
!
address-family ipv4
network 10.2.2.0 mask 255.255.255.0 route-map INTERNAL-ROUTES
network 10.2.3.0 mask 255.255.255.0 route-map INTERNAL-ROUTES
network 10.2.4.0 mask 255.255.255.0 route-map INTERNAL-ROUTES
network 10.2.5.0 mask 255.255.255.0 route-map INTERNAL-ROUTES
neighbor TEST_GROUP send-community both
neighbor TEST_GROUP next-hop-self
neighbor TEST_GROUP soft-reconfiguration inbound
neighbor 10.2.2.11 activate
neighbor 22.2.22.12 activate
neighbor 22.2.22.12 send-community both
neighbor 22.2.22.12 prefix-list FILTER-OUT out
exit-address-family
ip route 0.0.0.0 0.0.0.0 22.2.22.12
ip route 10.2.2.0 255.255.255.0 Null0
ip route 10.2.3.0 255.255.255.0 10.2.3.254
ip route 10.2.4.0 255.255.255.0 Null0
ip route 10.2.5.0 255.255.255.0 Null0
ip prefix-list FILTER-OUT seq 10 permit 10.2.2.0/24
ip prefix-list FILTER-OUT seq 20 permit 10.2.3.0/24
ip prefix-list FILTER-OUT seq 30 permit 10.2.4.0/24
ip prefix-list FILTER-OUT seq 40 permit 10.2.5.0/24
interface Loopback0
description Loopback OSPFBGP
ip address 10.2.2.10 255.255.255.255
!
interface Port-channel1.14
encapsulation dot1Q 14
ip address 10.2.3.226 255.255.255.248
glbp 1 ip 10.2.3.227
glbp 1 priority 200
no glbp 1 load-balancing
glbp 1 authentication md5 key-chain glbp
interface GigabitEthernet0/0/4
description ISP
ip address 22.2.22.11 255.255.255.252
no negotiation auto
!
!
Solved! Go to Solution.
- Labels:
- 
						
							
		
