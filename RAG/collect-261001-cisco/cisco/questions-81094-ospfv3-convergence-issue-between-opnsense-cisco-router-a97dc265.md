---
id: collect-261001-cisco/cisco/questions-81094-ospfv3-convergence-issue-between-opnsense-cisco-router-a97dc265
title: "OSPFv3 Convergence Issue between OPNSense &amp; Cisco Router"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/questions-81094-ospfv3-convergence-issue-between-opnsense-cisco-router-a97dc265.md
source_anchor: ""
source_lines: [1, 153]
sha256: 73a8e0a2c3af88f72bcd8780bb678a2bb98bbc5befd6d7df68280e318cd612a6
---

# OSPFv3 Convergence Issue between OPNSense &amp; Cisco Router

*Score : 1 | Source : https://networkengineering.stackexchange.com/questions/81094/ospfv3-convergence-issue-between-opnsense-cisco-router*

I have been digging around the internet and have been trying to get IPv6 up and running on my lab. Here is my current layout:
I have been trying to get OPNSense and the Cisco 1900 router to give each other's networks via OSPFv3. DHCPv6 is forwarding all the correct prefixes to the Home Lab and the router is distributing them properly. The OPNSense FW and the home lab can ping each other. However, they never seem to converge and establish a DR/BR relationship.
Here are my configurations:
CISCO 1900 CONFIGURATION
ipv6 unicast-routing
ipv6 cef
! interface GigabitEthernet0/0 
 ip address 10.0.1.2 255.255.255.252
 ip ospf 1 area 0
 duplex auto
 speed auto
 ipv6 address autoconfig default
 ipv6 enable
 ipv6 dhcp client pd prefix-from-provider
 ipv6 ospf 1 area 0
 ipv6 ospf priority 2
! 
interface GigabitEthernet0/1
 no ip address
 duplex full
 speed auto 
! 
interface GigabitEthernet0/1.10
 description DMZ
 encapsulation dot1Q 10
 ip address 10.1.10.1 255.255.255.0
 ipv6 address prefix-from-provider ::1:0:0:0:1/64
 ipv6 enable 
! 
interface GigabitEthernet0/1.100
 description MGMT VLAN
 encapsulation dot1Q 100 native
 ip address 10.1.100.1 255.255.255.0
 ipv6 address prefix-from-provider ::2:0:0:0:1/64
 ipv6 enable
 
! 
router ospfv3 1
 router-id 2.2.2.2
 log-adjacency-changes detail
! 
 address-family ipv6 unicast
 area 0 normal
 exit-address-family 
! 
! 
router ospf 1 
 network 10.0.1.0 0.0.0.255 area 0
 network 10.1.0.0 0.0.255.255 area 0 
!
OPNSENSE CONFIGURATION
Cisco Router Diagnostics
EdgeRTR#sh ipv6 ospf
 Routing Process "ospfv3 1" with ID 2.2.2.2
 Supports NSSA (compatible with RFC 3101)
 Supports Database Exchange Summary List Optimization (RFC 5243)
 Event-log enabled, Maximum number of events: 1000, Mode: cyclic
 Router is not originating router-LSAs with maximum metric
 Initial SPF schedule delay 5000 msecs
 Minimum hold time between two consecutive SPFs 10000 msecs
 Maximum wait time between two consecutive SPFs 10000 msecs
 Minimum LSA interval 5 secs
 Minimum LSA arrival 1000 msecs
 LSA group pacing timer 240 secs
 Interface flood pacing timer 33 msecs
 Retransmission pacing timer 66 msecs
 Retransmission limit dc 24 non-dc 24
 EXCHANGE/LOADING adjacency limit: initial 300, process maximum 300
 Number of external LSA 0. Checksum Sum 0x000000
 Number of areas in this router is 1. 1 normal 0 stub 0 nssa
 Graceful restart helper support enabled
 Reference bandwidth unit is 100 mbps
 RFC1583 compatibility enabled
    Area BACKBONE(0) (Inactive)
        Number of interfaces in this area is 1
        SPF algorithm executed 4 times
        Number of LSA 3. Checksum Sum 0x00E61D
        Number of DCbitless LSA 0
        Number of indication LSA 0
        Number of DoNotAge LSA 0
        Flood list length 0
EdgeRTR#sh ipv6 ospf neigh
            OSPFv3 Router with ID (2.2.2.2) (Process ID 1)
Neighbor ID     Pri   State           Dead Time   Interface ID    Interface
3.3.3.3           1   INIT/DROTHER    00:00:38    3               GigabitEthernet0/0
EdgeRTR#sh ipv6 ospf int gi0/0
GigabitEthernet0/0 is up, line protocol is up
  Link Local Address FE80::B6DE:31FF:FEA3:AC60, Interface ID 3
  Area 0, Process ID 1, Instance ID 0, Router ID 2.2.2.2
  Network Type BROADCAST, Cost: 1
  Transmit Delay is 1 sec, State DR, Priority 2
  Designated Router (ID) 2.2.2.2, local address FE80::B6DE:31FF:FEA3:AC60
  No backup designated router on this network
  Timer intervals configured, Hello 10, Dead 40, Wait 40, Retransmit 5
    Hello due in 00:00:06
  Graceful restart helper support enabled
  Index 1/1/1, flood queue length 0
  Next 0x0(0)/0x0(0)/0x0(0)
  Last flood scan length is 0, maximum is 0
  Last flood scan time is 0 msec, maximum is 0 msec
  Neighbor Count is 1, Adjacent neighbor count is 0
  Suppress hello for 0 neighbor(s)
Any help with this would be greatly appreciated. I know I am missing something, but what is it?
Cheers!
Raksul

---

### Reponse — score 1

What you have on the Cisco router is incomplete and incorrect. OSPFv3 for IPv6 needs to be configured on each interface. You can also use OSPFv3 for IPv4 and configure it on each interface. Remember that the network statements only tell OSPF which interfaces to include in the OSPF process, not which networks to advertise, and OSPFv3 has simplified that by having you configure on the interfaces instead of under the OSPF router command.
You can also configure IPv4 with OSPFv3 the same way using the IPv4 address family under the OSPFv3 router, or using OSPFv2 on the interfaces instead of network statements under the OSPFv2 router.
interface GigabitEthernet0/0 
 ip address 10.0.1.2 255.255.255.252
 ip ospf 1 area 0
 duplex auto
 speed auto
 ipv6 address autoconfig default
 ipv6 enable
 ipv6 dhcp client pd prefix-from-provider
 ipv6 ospf 1 area 0
 ipv6 ospf priority 2
 ospfv3 1 ipv6 area 0  !*** Add interface to OSPFv3  ***
! 
interface GigabitEthernet0/1.10
 description DMZ
 encapsulation dot1Q 10
 ip address 10.1.10.1 255.255.255.0
 ipv6 address prefix-from-provider ::1:0:0:0:1/64
 ipv6 enable 
 ospfv3 1 ipv6 area 0  !*** Add interface to OSPFv3  ***
! 
interface GigabitEthernet0/1.100
 description MGMT VLAN
 encapsulation dot1Q 100 native
 ip address 10.1.100.1 255.255.255.0
 ipv6 address prefix-from-provider ::2:0:0:0:1/64
 ipv6 enable
 ospfv3 1 ipv6 area 0  !*** Add interface to OSPFv3  ***
! 
router ospfv3 1
 router-id 2.2.2.2
 log-adjacency-changes detail
! 
 address-family ipv6 unicast
 area 0 normal
 exit-address-family 
! 
Unfortunately, the OPNSense manufacturer does not offer optional, paid support, which is a requirement to be on-topic here.
