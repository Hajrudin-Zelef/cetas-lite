---
id: collect-261001-general-networking/general-networking/r-networking-comments-38x7db-ospf-adjacency-randomly-dropping-after-forming-a527c626
title: "r-networking-comments-38x7db-ospf-adjacency-randomly-dropping-after-forming-a527c626"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/r-networking-comments-38x7db-ospf-adjacency-randomly-dropping-after-forming-a527c626.md
source_anchor: ""
source_lines: [1, 101]
sha256: ac7f946a514c65a142fe572d1e21f52da07ea72eec71ff4f1d49b06801c6210a
---

# r-networking-comments-38x7db-ospf-adjacency-randomly-dropping-after-forming-a527c626

OSPF Adjacency randomly dropping after forming between Cisco 3750-X and Corente Gateway. 
        
    We have an issue of an OSPF adjacency forming and dropping at random times. Initially configured and the adjacency was formed and the Cisco 3750-X installed the routes fine. After a couple minutes the routes were gone and noticed the adjacency was down. After a little research I was told the issue was a possible MTU mismatch between the two devices.
I entered the ip ospf mtu-ignore command as the fix. It was working fine for about a day and the same exact thing happened. Neighbor dropped and lost the routes. The only fix for this I have seen is to use the ip ospf mtu-ignore command which I have done.
Exact Configuration: 3750-X running ip base, and a Corente Red Hat Virtual Machine. OSPF configured on both
Below is some debugging I have done with this configured but no adjacency is up. Anyone ever ran into this before or know where to start looking next?
#show ip ospf neighbor
 
Neighbor ID     Pri   State           Dead Time   Address         Interface
10.0.11.246       1   DOWN/DROTHER       -        192.168.60.18   Vlan5
192.168.60.34     1   FULL/DR         00:00:32    192.168.60.34   Vlan99
192.168.60.66     1   FULL/DR         00:00:35    192.168.60.49   GigabitEthernet1/0/4
 #show ip ospf 1
 Routing Process "ospf 1" with ID 192.168.63.1
 Start time: 1y30w, Time elapsed: 2d02h
 Supports only single TOS(TOS0) routes
 Supports opaque LSA
 Supports Link-local Signaling (LLS)
 Supports area transit capability
 Event-log enabled, Maximum number of events: 1000, Mode: cyclic
 Router is not originating router-LSAs with maximum metric
 Initial SPF schedule delay 5000 msecs
 Minimum hold time between two consecutive SPFs 10000 msecs
 Maximum wait time between two consecutive SPFs 10000 msecs
 Incremental-SPF disabled
 Minimum LSA interval 5 secs
 Minimum LSA arrival 1000 msecs
 LSA group pacing timer 240 secs
 Interface flood pacing timer 33 msecs
 Retransmission pacing timer 66 msecs
 Number of external LSA 15. Checksum Sum 0x2522C7
 Number of opaque AS LSA 0. Checksum Sum 0x000000
 Number of DCbitless external and opaque AS LSA 15
 Number of DoNotAge external and opaque AS LSA 0
 Number of areas in this router is 1. 1 normal 0 stub 0 nssa
 Number of areas transit capable is 0
 External flood list length 0
 IETF NSF helper support enabled
 Cisco NSF helper support enabled
 Reference bandwidth unit is 100 mbps
    Area BACKBONE(0)
        Number of interfaces in this area is 5
  Area has no authentication
  SPF algorithm last executed 00:01:27.308 ago
  SPF algorithm executed 2 times
  Area ranges are
  Number of LSA 6. Checksum Sum 0x0EFFB0
  Number of opaque link LSA 0. Checksum Sum 0x000000
  Number of DCbitless LSA 1
  Number of indication LSA 0
  Number of DoNotAge LSA 0
  Flood list length 0
*Sep 29 04:59:29.922: OSPF: Send DBD to 10.0.11.246 on Vlan5 seq 0x1828 opt 0x52 flag 0x7 len 32
*Sep 29 04:59:29.922: OSPF: Retransmitting DBD to 10.0.11.246 on Vlan5 [25]
*Sep 29 04:59:33.890: OSPF: Rcv DBD from 10.0.11.246 on Vlan5 seq 0xDCC opt 0x2 flag 0x7 len 32  mtu 1500 state EXSTART
*Sep 29 04:59:33.890: OSPF: First DBD and we are not SLAVE
*Sep 29 04:59:34.888: OSPF: Killing nbr 10.0.11.246 on Vlan5 due to excessive (25) retransmissions
*Sep 29 04:59:34.888: OSPF: 10.0.11.246 address 192.168.60.18 on Vlan5 is dead, state DOWN
*Sep 29 04:59:34.888: %OSPF-5-ADJCHG: Process 1, Nbr 10.0.11.246 on Vlan5 from EXSTART to DOWN, Neighbor Down: Too many retransmissions
*Sep 29 04:59:34.888: OSPF: Vlan5 Nbr 10.0.11.246: Clean-up dbase exchange
*Sep 29 04:59:34.888: OSPF: Neighbor change Event on interface Vlan5
*Sep 29 04:59:34.888: OSPF: DR/BDR election on Vlan5 
*Sep 29 04:59:34.888: OSPF: Elect BDR 0.0.0.0
*Sep 29 04:59:34.888: OSPF: Elect DR 192.168.63.1
*Sep 29 04:59:34.888:        DR: 192.168.63.1 (Id)   BDR: none 
*Sep 29 04:59:38.898: OSPF: Nbr 10.0.11.246 192.168.60.18 Vlan5 is currently ignored
*Sep 29 05:00:37.182: OSPF: Rcv DBD from 10.0.11.246 on Vlan5 seq 0xDCD opt 0x2 flag 0x7 len 32  mtu 1500 state EXSTART
*Sep 29 05:00:37.182: OSPF: First DBD and we are not SLAVE
*Sep 29 05:00:37.711: OSPF: Build router LSA for area 0, router ID 192.168.63.1, seq 0x80000071, process 1
*Sep 29 05:00:39.548: OSPF: Send DBD to 10.0.11.246 on Vlan5 seq 0x10B opt 0x52 flag 0x7 len 32
*Sep 29 05:00:39.548: OSPF: Retransmitting DBD to 10.0.11.246 on Vlan5 [2]
*Sep 29 05:00:40.328: OSPF: Neighbor change Event on interface Vlan5
*Sep 29 05:00:40.328: OSPF: DR/BDR election on Vlan5 
*Sep 29 05:00:40.328: OSPF: Elect BDR 192.168.63.1
*Sep 29 05:00:40.328: OSPF: Elect DR 10.0.11.246
*Sep 29 05:00:40.328:        DR: 10.0.11.246 (Id)   BDR: 192.168.63.1 (Id)
*Sep 29 05:00:42.190: OSPF: Rcv DBD from 10.0.11.246 on Vlan5 seq 0xDCD opt 0x2 flag 0x7 len 32  mtu 1500 state EXSTART
*Sep 29 05:00:42.190: OSPF: First DBD and we are not SLAVE
*Sep 29 05:00:44.380: OSPF: Send DBD to 10.0.11.246 on Vlan5 seq 0x10B opt 0x52 flag 0x7 len 32
*Sep 29 05:00:44.380: OSPF: Retransmitting DBD to 10.0.11.246 on Vlan5 [3]
*Sep 29 05:00:44.707: OSPF: Neighbor change Event on interface Vlan99
*Sep 29 05:00:44.707: OSPF: DR/BDR election on Vlan99 
*Sep 29 05:00:44.707: OSPF: Elect BDR 192.168.63.1
*Sep 29 05:00:44.707: OSPF: Elect DR 192.168.60.34
*Sep 29 05:00:44.707:        DR: 192.168.60.34 (Id)   BDR: 192.168.63.1 (Id)
*Sep 29 05:00:47.198: OSPF: Rcv DBD from 10.0.11.246 on Vlan5 seq 0xDCD opt 0x2 flag 0x7 len 32  mtu 1500 state EXSTART
*Sep 29 05:00:47.198: OSPF: First DBD and we are not SLAVE
*Sep 29 05:00:49.379: OSPF: Send DBD to 10.0.11.246 on Vlan5 seq 0x10B opt 0x52 flag 0x7 len 32
*Sep 29 05:00:49.379: OSPF: Retransmitting DBD to 10.0.11.246 on Vlan5 [4]
*Sep 29 05:00:52.215: OSPF: Rcv DBD from 10.0.11.246 on Vlan5 seq 0xDCD opt 0x2 flag 0x7 len 32  mtu 1500 state EXSTART
*Sep 29 05:00:52.215: OSPF: First DBD and we are not SLAVE
    
Section des commentaires
Fix the remote device. It's either losing packets, or not replying.
Also, MTU is checked when adjacency is formed, so ignore-mtu is a red herring.
Checked the MTU and both devices are 1500.
Eh, that's what I said. MTU is not your issue.
Commentaire supprimé par le membre
Shouldn't the ip ospf mtu-ignore command fix this problem all together? When I applied it the adjacency stayed up for a solid day compared to a couple minutes before I applied the command.
If the MTUs are different sizes, one host is going to try to send packets too big for the other host. This difference is supposed to be detected, and considered an error, during adjacency formation. Suppressing the error during adjacency formation doesn't actually fix the problem that two hosts have different ideas about the MTU size on the link they share.
Look at the MTU on each 3750-x and the Corente gateway. If they don't match, make them match.
