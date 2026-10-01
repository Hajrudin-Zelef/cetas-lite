---
id: collect-261001-general-networking/general-networking/questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing-60a06b3c-1
title: "questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing--60a06b3c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-general-networking/questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing--60a06b3c.md
source_anchor: ""
source_lines: [1, 163]
sha256: 54a2b8151f5d94ac97b67b7ca78e1691ef8a18f96729084f35f45d93c20707bf
---

# questions-5678-network-not-participating-in-ospf-does-not-appear-in-the-routing--60a06b3c

I am running OSPF on 2 router in GNS simulation for practice. Adjacencies are ok. Router R2 has in the routing table all networks of R3 . But R3 fails to add R2 network in its routing table . The network that is not participating in ospf is the 172.16.100.0/24 as shown below Any idea why?
**R2**
Routing Protocol is "ospf 1"
  Outgoing update filter list for all interfaces is not set
  Incoming update filter list for all interfaces is not set
  Router ID 172.168.100.1
  It is an autonomous system boundary router
  Redistributing External Routes from,
    eigrp 1, includes subnets in redistribution
  Number of areas in this router is 2. 2 normal 0 stub 0 nssa
  Maximum path: 4
  Routing for Networks:
    172.16.23.0 0.0.0.255 area 0
    172.16.100.0 0.0.0.255 area 10
 Reference bandwidth unit is 100 mbps
  Routing Information Sources:
    Gateway         Distance      Last Update
    192.168.40.1         110      00:14:35
  Distance: (default is 110)
R2#show run | s ospf
 ip ospf network point-to-point
 redistribute ospf 1 metric 10000 100 255 1 1500
router ospf 1
 log-adjacency-changes
 redistribute eigrp 1 subnets
 network 172.16.23.0 0.0.0.255 area 0
 network 172.16.100.0 0.0.0.255 area 10
R2#show ip ospf interface
Serial0/1 is up, line protocol is up
  Internet Address 172.16.23.2/24, Area 0
  Process ID 1, Router ID 172.168.100.1, Network Type POINT_TO_POINT, Cost: 64
  Transmit Delay is 1 sec, State POINT_TO_POINT,
  Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    oob-resync timeout 40
    Hello due in 00:00:07
  Supports Link-local Signaling (LLS)
  Index 1/1, flood queue length 0
  Next 0x0(0)/0x0(0)
  Last flood scan length is 1, maximum is 6
  Last flood scan time is 0 msec, maximum is 0 msec
  Neighbor Count is 1, Adjacent neighbor count is 1
    Adjacent with neighbor 192.168.40.1
  Suppress hello for 0 neighbor(s)
R2#show ip int b
Interface                  IP-Address      OK? Method Status                Protocol
FastEthernet0/0            unassigned      YES NVRAM  administratively down down 
Serial0/0                  172.16.12.2     YES NVRAM  up                    up   
FastEthernet0/1            unassigned      YES NVRAM  administratively down down 
Serial0/1                  172.16.23.2     YES NVRAM  up                    up   
Loopback0                  172.16.2.1      YES NVRAM  up                    up   
Loopback100                172.168.100.1   YES manual up                    up
R2#show ip route
Codes: C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route
Gateway of last resort is not set
     192.168.30.0/32 is subnetted, 1 subnets
O       192.168.30.1 [110/65] via 172.16.23.3, 00:36:47, Serial0/1
     192.168.25.0/32 is subnetted, 1 subnets
O       192.168.25.1 [110/65] via 172.16.23.3, 00:36:47, Serial0/1
     192.168.40.0/32 is subnetted, 1 subnets
O       192.168.40.1 [110/65] via 172.16.23.3, 00:36:47, Serial0/1
     172.168.0.0/24 is subnetted, 1 subnets
C       172.168.100.0 is directly connected, Loopback100
     172.16.0.0/16 is variably subnetted, 5 subnets, 2 masks
C       172.16.23.0/24 is directly connected, Serial0/1
C       172.16.12.0/24 is directly connected, Serial0/0
D       172.16.1.0/24 [90/2297856] via 172.16.12.1, 01:06:07, Serial0/0
O       172.16.3.1/32 [110/65] via 172.16.23.3, 00:36:48, Serial0/1
C       172.16.2.0/24 is directly connected, Loopback0
     192.168.20.0/32 is subnetted, 1 subnets
O       192.168.20.1 [110/65] via 172.16.23.3, 00:36:49, Serial0/1
D    192.168.51.0/24 [90/2297856] via 172.16.12.1, 01:06:08, Serial0/0
D    192.168.50.0/24 [90/2297856] via 172.16.12.1, 01:06:08, Serial0/0
     192.168.35.0/32 is subnetted, 1 subnets
O       192.168.35.1 [110/65] via 172.16.23.3, 00:36:49, Serial0/1
D    192.168.70.0/24 [90/2297856] via 172.16.12.1, 01:06:08, Serial0/0
O IA 192.168.8.0/22 [110/65] via 172.16.23.3, 00:36:49, Serial0/1
D    192.168.48.0/23 [90/2297856] via 172.16.12.1, 01:06:08, Serial0/0
"HERE the interface loopback 100 does not appear , i added it using network command"
**R3**
Routing Protocol is "ospf 1"
  Outgoing update filter list for all interfaces is not set
  Incoming update filter list for all interfaces is not set
  Router ID 192.168.40.1
  It is an area border router
  Number of areas in this router is 2. 2 normal 0 stub 0 nssa
  Maximum path: 4
  Routing for Networks:
    172.16.3.0 0.0.0.255 area 0
    172.16.23.0 0.0.0.255 area 0
    192.168.8.0 0.0.0.255 area 20
    192.168.9.0 0.0.0.255 area 20
    192.168.10.0 0.0.0.255 area 20
    192.168.11.0 0.0.0.255 area 20
    192.168.20.0 0.0.0.255 area 0
    192.168.25.0 0.0.0.255 area 0
    192.168.30.0 0.0.0.255 area 0
    192.168.35.0 0.0.0.255 area 0
    192.168.40.0 0.0.0.255 area 0
 Reference bandwidth unit is 100 mbps
  Routing Information Sources:
    Gateway         Distance      Last Update
    (this router)        110      00:35:11
    Gateway         Distance      Last Update
    172.168.100.1        110      00:15:09
    172.16.2.1           110      02:03:26
  Distance: (default is 110)
R3#show run | s ospf
 ip ospf network point-to-point
 ip ospf network point-to-point
 ip ospf network point-to-point
 ip ospf network point-to-point
router ospf 1
 log-adjacency-changes
 area 20 range 192.168.8.0 255.255.252.0
 network 172.16.3.0 0.0.0.255 area 0
 network 172.16.23.0 0.0.0.255 area 0
 network 192.168.8.0 0.0.0.255 area 20
 network 192.168.9.0 0.0.0.255 area 20
 network 192.168.10.0 0.0.0.255 area 20
 network 192.168.11.0 0.0.0.255 area 20
 network 192.168.20.0 0.0.0.255 area 0
 network 192.168.25.0 0.0.0.255 area 0
 network 192.168.30.0 0.0.0.255 area 0
 network 192.168.35.0 0.0.0.255 area 0
 network 192.168.40.0 0.0.0.255 area 0
R3#show ip ospf interface
Loopback40 is up, line protocol is up
  Internet Address 192.168.40.1/24, Area 0
  Process ID 1, Router ID 192.168.40.1, Network Type LOOPBACK, Cost: 1
  Loopback interface is treated as a stub Host
Loopback35 is up, line protocol is up
  Internet Address 192.168.35.1/24, Area 0
  Process ID 1, Router ID 192.168.40.1, Network Type LOOPBACK, Cost: 1
  Loopback interface is treated as a stub Host
Loopback30 is up, line protocol is up
  Internet Address 192.168.30.1/24, Area 0
  Process ID 1, Router ID 192.168.40.1, Network Type LOOPBACK, Cost: 1
  Loopback interface is treated as a stub Host
Loopback25 is up, line protocol is up
  Internet Address 192.168.25.1/24, Area 0
  Process ID 1, Router ID 192.168.40.1, Network Type LOOPBACK, Cost: 1
  Loopback interface is treated as a stub Host
Loopback20 is up, line protocol is up
  Internet Address 192.168.20.1/24, Area 0
  Process ID 1, Router ID 192.168.40.1, Network Type LOOPBACK, Cost: 1
  Loopback interface is treated as a stub Host
R3#show ip int b
Interface                  IP-Address      OK? Method Status                Protocol
FastEthernet0/0            unassigned      YES NVRAM  administratively down down
Serial0/0                  172.16.23.3     YES NVRAM  up                    up
FastEthernet0/1            unassigned      YES NVRAM  administratively down down
Serial0/1                  unassigned      YES NVRAM  administratively down down
Loopback0                  172.16.3.1      YES NVRAM  up                    up
Loopback8                  192.168.8.1     YES manual up                    up
Loopback9                  192.168.9.1     YES manual up                    up
Loopback10                 192.168.10.1    YES manual up                    up
Loopback11                 192.168.11.1    YES manual up                    up
