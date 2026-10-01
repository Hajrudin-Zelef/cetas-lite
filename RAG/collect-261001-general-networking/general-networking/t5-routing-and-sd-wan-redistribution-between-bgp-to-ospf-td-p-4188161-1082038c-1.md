---
id: collect-261001-general-networking/general-networking/t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c-1
title: "t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c.md
source_anchor: ""
source_lines: [1, 159]
sha256: 22b1e2d866180c9e1370d14b8e2f5ef6a14d381a3d573b3c486e7beb8b889174
---

# t5-routing-and-sd-wan-redistribution-between-bgp-to-ospf-td-p-4188161-1082038c

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-24-2020 07:52 AM
Hi, I am new in redistrution. I have configured the routers. I can see OSPF routes into BGP. However not able to see the BGP router in OSPF as external route. Kindly suggest me how should I configure to get it.
Topology: Attached.
R1:
R1#sh run | s r ospf
router ospf 1
log-adjacency-changes
network 1.1.1.1 0.0.0.0 area 0
network 12.1.1.0 0.0.0.255 area 0
network 14.1.1.0 0.0.0.255 area 0
R1#
R1#
R1#sh ip int br | ex down|unass
Interface IP-Address OK? Method Status Protocol
Ethernet1/0 12.1.1.1 YES NVRAM up up
Ethernet1/1 14.1.1.1 YES NVRAM up up
Loopback0 1.1.1.1 YES NVRAM up up
R1#
R2:
*Nov 24 21:11:40.607: %SYS-5-CONFIG_I: Configured from console by console
R2#sh run | s r bgp
router bgp 235
no synchronization
bgp log-neighbor-changes
network 2.2.2.2 mask 255.255.255.255
redistribute ospf 1
neighbor 23.1.1.3 remote-as 235
neighbor 23.1.1.3 next-hop-self
no auto-summary
R2#sh run | s r ospf
router ospf 1
log-adjacency-changes
redistribute connected subnets
redistribute bgp 235 subnets
R2#
R2#sh ip int br | ex down|unass
Interface IP-Address OK? Method Status Protocol
Ethernet1/0 12.1.1.2 YES NVRAM up up
Ethernet1/1 23.1.1.2 YES NVRAM up up
Loopback0 2.2.2.2 YES NVRAM up up
R5:
R5#sh run | s r bgp
router bgp 235
no synchronization
bgp log-neighbor-changes
network 5.5.5.5 mask 255.255.255.255
redistribute ospf 1 match internal external 1 external 2
neighbor 53.1.1.3 remote-as 235
neighbor 53.1.1.3 next-hop-self
no auto-summary
R5#sh run | s r ospf
router ospf 1
log-adjacency-changes
redistribute connected subnets
redistribute bgp 235 subnets
R5#sh ip int br | ex down|unass
Interface IP-Address OK? Method Status Protocol
Ethernet1/0 45.1.1.5 YES NVRAM up up
Ethernet1/1 53.1.1.5 YES NVRAM up up
Loopback0 5.5.5.5 YES NVRAM up up
R5#
R3:
R3#
R3#sh run | s r bgp
router bgp 235
no synchronization
bgp log-neighbor-changes
network 3.3.3.3 mask 255.255.255.255
neighbor 23.1.1.2 remote-as 235
neighbor 53.1.1.5 remote-as 235
no auto-summary
R3#
R3#
R3#
R3#sh ip bgp su
BGP router identifier 3.3.3.3, local AS number 235
BGP table version is 28, main routing table version 28
9 network entries using 1080 bytes of memory
15 path entries using 780 bytes of memory
8/5 BGP path/bestpath attribute entries using 992 bytes of memory
0 BGP route-map cache entries using 0 bytes of memory
0 BGP filter-list cache entries using 0 bytes of memory
Bitfield cache entries: current 1 (at peak 2) using 32 bytes of memory
BGP using 2884 total bytes of memory
BGP activity 16/7 prefixes, 23/8 paths, scan interval 60 secs
Neighbor V AS MsgRcvd MsgSent TblVer InQ OutQ Up/Down State/PfxRcd
23.1.1.2 4 235 53 42 28 0 0 00:17:44 6
53.1.1.5 4 235 46 38 28 0 0 00:13:54 8
R4:
e
R4#sh ip int br
Interface IP-Address OK? Method Status Protocol
FastEthernet0/0 unassigned YES NVRAM administratively down down
Ethernet1/0 14.1.1.4 YES NVRAM up up
Ethernet1/1 45.1.1.4 YES NVRAM up up
Ethernet1/2 unassigned YES NVRAM administratively down down
Ethernet1/3 unassigned YES NVRAM administratively down down
GigabitEthernet2/0 unassigned YES NVRAM administratively down down
Loopback0 4.4.4.4 YES NVRAM up up
R4#sh ip int br | ex down
Interface IP-Address OK? Method Status Protocol
Ethernet1/0 14.1.1.4 YES NVRAM up up
Ethernet1/1 45.1.1.4 YES NVRAM up up
Loopback0 4.4.4.4 YES NVRAM up up
R4#sh run | s r os
router ospf 1
log-adjacency-changes
network 4.4.4.4 0.0.0.0 area 0
network 14.1.1.0 0.0.0.255 area 0
network 45.1.1.0 0.0.0.255 area 0
R4#sh ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
E1 - OSPF external type 1, E2 - OSPF external type 2
i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
ia - IS-IS inter area, * - candidate default, U - per-user static route
o - ODR, P - periodic downloaded static route
Gateway of last resort is not set
1.0.0.0/32 is subnetted, 1 subnets
O 1.1.1.1 [110/11] via 14.1.1.1, 00:39:07, Ethernet1/0
2.0.0.0/32 is subnetted, 1 subnets
O E2 2.2.2.2 [110/20] via 14.1.1.1, 00:03:55, Ethernet1/0
4.0.0.0/32 is subnetted, 1 subnets
C 4.4.4.4 is directly connected, Loopback0
5.0.0.0/32 is subnetted, 1 subnets
O E2 5.5.5.5 [110/20] via 45.1.1.5, 00:13:30, Ethernet1/1
53.0.0.0/24 is subnetted, 1 subnets
O E2 53.1.1.0 [110/20] via 45.1.1.5, 00:13:30, Ethernet1/1
23.0.0.0/24 is subnetted, 1 subnets
O E2 23.1.1.0 [110/20] via 14.1.1.1, 00:03:55, Ethernet1/0
12.0.0.0/24 is subnetted, 1 subnets
O 12.1.1.0 [110/20] via 14.1.1.1, 00:21:17, Ethernet1/0
14.0.0.0/24 is subnetted, 1 subnets
C 14.1.1.0 is directly connected, Ethernet1/0
45.0.0.0/24 is subnetted, 1 subnets
C 45.1.1.0 is directly connected, Ethernet1/1
The redistribution happened on R2 and R5. Kindly suggest how I can see the bgp routes in pspf
R1----------R2
| R3
R4 --------R5
Solved! Go to Solution.
- Labels:
- 
						
							
		
			LAN Switching
- 
						
							
		
