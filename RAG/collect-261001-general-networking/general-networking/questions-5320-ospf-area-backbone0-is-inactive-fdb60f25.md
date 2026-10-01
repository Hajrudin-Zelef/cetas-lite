---
id: collect-261001-general-networking/general-networking/questions-5320-ospf-area-backbone0-is-inactive-fdb60f25
title: "questions-5320-ospf-area-backbone0-is-inactive-fdb60f25"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "voice"]
source: docs/RAG/collect-261001-general-networking/questions-5320-ospf-area-backbone0-is-inactive-fdb60f25.md
source_anchor: ""
source_lines: [1, 247]
sha256: dad028178abc85d16327cd5396932c5b335d1edb7bf3d0bb5943d8947e271719
---

# questions-5320-ospf-area-backbone0-is-inactive-fdb60f25

I am having a puzzling issue with OSPF connectivity on one of my routers. The other routers in my network have formed an adjacency and are working correctly, but this last one is a case I've never seen before. All of the interfaces on the remaining router are up and functional and are connected to AREA 0 just as the other two routers are. However, a show ip ospf command results in the following:
 Area BACKBONE(0) (Inactive)
        Number of interfaces in this area is 3
Additionally, I get nothing if I try to ping 224.0.0.5 or 224.0.0.6 from this router. I've tried restarting the OSPF process on the router as well as rebooting the router itself, but it makes no difference. When I turn debugging on, I can see the router isn't even sending OSPF hello messages. What could be the cause of something like this?
The router interfaces are configured as follows for the router in question:
f0/0 - 111.1.0.2/8
f0/1 - 112.1.0.2/8
f0/2/0 - (Vlan5) 122.0.0.1/8
f0/2/3 - (Vlan88) 10.10.10.1/28
Here is the entire router config along with OSPF process info and interface statuses:
router_1#sh ip ospf 
 Routing Process "ospf 10" with ID 122.0.0.1
 Start time: 00:00:19.304, Time elapsed: 18:57:25.596
 Supports only single TOS(TOS0) routes
 Supports opaque LSA
 Supports Link-local Signaling (LLS)
 Supports area transit capability
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
 Number of external LSA 0. Checksum Sum 0x000000
 Number of opaque AS LSA 0. Checksum Sum 0x000000
 Number of DCbitless external and opaque AS LSA 0
 Number of DoNotAge external and opaque AS LSA 0
 Number of areas in this router is 1. 1 normal 0 stub 0 nssa
 Number of areas transit capable is 0
 External flood list length 0
 IETF NSF helper support enabled
 Cisco NSF helper support enabled
    Area BACKBONE(0) (Inactive)
    Number of interfaces in this area is 3
    Area has no authentication
    SPF algorithm last executed 01:36:13.484 ago
    SPF algorithm executed 3 times
    Area ranges are
    Number of LSA 1. Checksum Sum 0x00F9B7
    Number of opaque link LSA 0. Checksum Sum 0x000000
    Number of DCbitless LSA 0
    Number of indication LSA 0
    Number of DoNotAge LSA 0
    Flood list length 0
 router_1#sh ip int br
    Interface                  IP-Address      OK? Method Status                Protocol
    FastEthernet0/0            111.1.0.2       YES NVRAM  up                    up      
    FastEthernet0/1            112.1.0.2       YES NVRAM  up                    up      
    FastEthernet0/2/0          unassigned      YES unset  up                    up      
    FastEthernet0/2/1          unassigned      YES unset  administratively down down    
    FastEthernet0/2/2          unassigned      YES unset  administratively down down    
    FastEthernet0/2/3          unassigned      YES unset  up                    up      
    Vlan1                      unassigned      YES NVRAM  administratively down down    
    Vlan5                      122.0.0.1       YES NVRAM  up                    up      
    Vlan88                     10.10.10.1      YES NVRAM  up                    up      
router_1#sh run
Building configuration...
Current configuration : 3904 bytes
!
version 12.4
service timestamps debug datetime msec
service timestamps log datetime msec
no service password-encryption
!
hostname router_1
!
boot-start-marker
boot-end-marker
!
enable secret 5 $1$iXM1$Rcki3hrzpXULV8lCJXDqY1
!
no aaa new-model
!
!
ip cef
!
!
no ip domain lookup
multilink bundle-name authenticated
!
!
!
username thedoctor privilege 15 secret 5 $1$KUwA$VWYQeI/fm8LacwgysL7RK/
archive
 log config
  hidekeys
!
!
!
class-map match-all data-inmmediate
 match ip dscp af13 
class-map match-all class-all
 match any 
class-map match-all data-priority
 match ip dscp af23 
class-map match-all data-routine
 match ip dscp default 
class-map match-all data-flash
 match ip dscp default 
 match ip dscp af33 
class-map match-all video
 match ip dscp cs4  33  af41  35  af42 
class-map match-all data-flash-override
 match ip dscp 31 
class-map match-all voice
 match ip dscp cs5  41  42  43  44 
!
!
policy-map ecn-llq-enable
 class voice
  priority 48
 class video
  bandwidth percent 1
  random-detect dscp-based
  random-detect exponential-weighting-constant 1
  random-detect ecn
  random-detect dscp 32   8     16    16   
  random-detect dscp 33   16    24    16   
  random-detect dscp 34   24    32    16   
  random-detect dscp 35   32    40    16   
  random-detect dscp 36   40    48    16   
 class data-flash-override
  bandwidth percent 35
  random-detect dscp-based
  random-detect exponential-weighting-constant 1
  random-detect ecn
  random-detect dscp 31   70    4000  2    
 class data-flash
  bandwidth percent 1
  random-detect dscp-based
  random-detect exponential-weighting-constant 1
  random-detect ecn
  random-detect dscp 30   80    96    16   
 class data-priority
  bandwidth percent 1
  random-detect dscp-based
  random-detect exponential-weighting-constant 1
  random-detect ecn
  random-detect dscp 22   96    112   12   
 class data-inmmediate
  bandwidth percent 1
  random-detect dscp-based
  random-detect exponential-weighting-constant 1
  random-detect ecn
  random-detect dscp 14   112   128   8    
 class data-routine
  bandwidth percent 25
  random-detect dscp-based
  random-detect exponential-weighting-constant 1
  random-detect ecn
  random-detect dscp 0   2     4000  2    
policy-map shape-all-ncw
 class class-all
  shape average 128000
  service-policy ecn-llq-enable
policy-map shape-all-hnw
 class class-all
  shape average 1000000
  service-policy ecn-llq-enable
!
!
!
!         
interface FastEthernet0/0
 description Intermidiate_Router_GW
 ip address 111.1.0.2 255.0.0.0
 ip ospf cost 400
 duplex auto
 speed auto
 service-policy output shape-all-hnw
!
interface FastEthernet0/1
 description Router2
 ip address 112.1.0.2 255.0.0.0
 ip ospf cost 600
 duplex auto
 speed auto
 service-policy output shape-all-hnw
!
interface FastEthernet0/2/0
 description ENCRYPTION1
 switchport access vlan 5
!
interface FastEthernet0/2/1
 shutdown
!
interface FastEthernet0/2/2
 shutdown
!
interface FastEthernet0/2/3
 description AF_NM
 switchport access vlan 88
!
interface Vlan1
 no ip address
 shutdown
!
interface Vlan5
 description ENCRYPTION1
 ip address 122.0.0.1 255.0.0.0
!
interface Vlan88
 description AF_NM
 ip address 10.10.10.1 255.255.255.248
!
router ospf 10
 log-adjacency-changes
 network 111.0.0.0 0.255.255.255 area 0
 network 112.0.0.0 0.255.255.255 area 0
 network 122.0.0.0 0.255.255.255 area 0
!
ip forward-protocol nd
ip route 0.0.0.0 0.0.0.0 120.0.0.1
ip route 122.1.1.0 255.255.255.0 122.0.0.2
ip route 122.1.2.0 255.255.255.0 122.0.0.2
!         
!
ip http server
ip http access-class 23
ip http authentication local
ip http timeout-policy idle 60 life 86400 requests 10000
!
access-list 23 permit 10.10.10.0 0.0.0.7
!
!
control-plane
!
!
line con 0
 login local
line aux 0
line vty 0 4
 privilege level 15
 login local
 transport input telnet
line vty 5 15
 privilege level 15
 login local
 transport input telnet
!
scheduler allocate 20000 1000
!
end
