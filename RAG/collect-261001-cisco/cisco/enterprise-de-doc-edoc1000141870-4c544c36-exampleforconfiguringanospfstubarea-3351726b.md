---
id: collect-261001-cisco/cisco/enterprise-de-doc-edoc1000141870-4c544c36-exampleforconfiguringanospfstubarea-3351726b
title: "Configure SwitchA. The configurations of SwitchB and SwitchC are similar."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-cisco/enterprise-de-doc-edoc1000141870-4c544c36-exampleforconfiguringanospfstubarea-3351726b.md
source_anchor: ""
source_lines: [1, 172]
sha256: 4ddf92da09286c9f5fc02321cc83c974c22f61a1fa53976046fc39a0e94dabcf
---

# Configure SwitchA. The configurations of SwitchB and SwitchC are similar.

Unternehmen
A stub area is an area that does not allow an ABR to advertise received AS external routes. In a stub area, the routing table size and transmitted routing information volume of routers are greatly reduced. A stub area is often placed at the edge of an AS. To ensure the reachability of a destination outside the AS, the ABR in the stub area generates a default route and advertises it to the non-ABR routers in the stub area.
Assume that a device of Company H connects to the backbone area through a single link. The device has low performance and a small routing table. The area where the device resides needs to access other areas or network segments outside the OSPF area, and the next-hop address of routes of the device is the IP address of the next-hop core device of the link. Therefore, the area where the device resides does not need to learn a large number of OSPF external routes and can be configured as a stub area. This configuration can reduce the routing table size of the area and resource consumption of the device.
For details about the hardware specifications and matched parts of the switch, visit Hardware Center. For details about the key specifications and full software specifications of the switch, visit Specifications Query.
As shown in Figure 8-7, SwitchA, SwitchB, and SwitchC run OSPF, and the OSPF network is divided into Area 0 and Area 1. SwitchB functions as an ASBR to communicate with external networks. The OSPF routing table size on SwitchC needs to be reduced without affecting communication.
The configuration roadmap is as follows:
Configure basic OSPF functions on each switch to implement interworking in the OSPF network.
Configure a static route on SwitchB and import the route to the OSPF routing table to ensure that there is a reachable route from the OSPF network to external networks.
Configure Area 1 as a stub area to reduce the OSPF routing table size on SwitchC.
Prohibit the ABR (SwitchA) in Area 1 from advertising Type 3 LSAs within the stub area to configure Area 1 as a totally stub area. This configuration minimizes the OSPF routing table size on SwitchC.
# Configure SwitchA. The configurations of SwitchB and SwitchC are similar.
<HUAWEI> system-view
[HUAWEI] sysname SwitchA
[SwitchA] vlan batch 10 20
[SwitchA] interface gigabitethernet 0/0/1
[SwitchA-GigabitEthernet0/0/1] port link-type trunk
[SwitchA-GigabitEthernet0/0/1] port trunk allow-pass vlan 10
[SwitchA-GigabitEthernet0/0/1] quit
[SwitchA] interface gigabitethernet 0/0/2
[SwitchA-GigabitEthernet0/0/2] port link-type trunk
[SwitchA-GigabitEthernet0/0/2] port trunk allow-pass vlan 20
[SwitchA-GigabitEthernet0/0/2] quit
# Configure SwitchA. The configurations of Switch B and SwitchC are similar.
[SwitchA] interface vlanif 10
[SwitchA-Vlanif10] ip address 192.168.0.1 24
[SwitchA-Vlanif10] quit
[SwitchA] interface vlanif 20
[SwitchA-Vlanif20] ip address 192.168.1.1 24
[SwitchA-Vlanif20] quit
# Configure SwitchA.
[SwitchA] ospf 1 router-id 10.1.1.1
[SwitchA-ospf-1] area 0
[SwitchA-ospf-1-area-0.0.0.0] network 192.168.0.0 0.0.0.255
[SwitchA-ospf-1-area-0.0.0.0] quit
[SwitchA-ospf-1] area 1
[SwitchA-ospf-1-area-0.0.0.1] network 192.168.1.0 0.0.0.255
[SwitchA-ospf-1-area-0.0.0.1] quit
[SwitchA-ospf-1] quit
# Configure SwitchB.
[SwitchB] ospf 1 router-id 10.2.2.2
[SwitchB-ospf-1] area 0
[SwitchB-ospf-1-area-0.0.0.0] network 192.168.0.0 0.0.0.255
[SwitchB-ospf-1-area-0.0.0.0] quit
[SwitchB-ospf-1] quit
# Configure SwitchC.
[SwitchC] ospf 1 router-id 10.3.3.3
[SwitchC-ospf-1] area 1
[SwitchC-ospf-1-area-0.0.0.1] network 192.168.1.0 0.0.0.255
[SwitchC-ospf-1-area-0.0.0.1] quit
[SwitchC-ospf-1] quit
[SwitchB] ip route-static 10.0.0.0 8 null 0
[SwitchB] ospf 1
[SwitchB-ospf-1] import-route static type 1   //SwitchB functions as an ASBR and imports external routes.
[SwitchB-ospf-1] quit
Check the OSPF routing table on SwitchC. The command output shows that the OSPF routing table contains an AS external route.
[SwitchC] display ospf routing
     OSPF Process 1 with Router ID 10.3.3.3
                   Routing Tables
 Routing for Network
 Destination        Cost  Type         NextHop         AdvRouter       Area
 192.168.1.0/24     1     Transit      192.168.1.2     10.3.3.3        0.0.0.1
 192.168.0.0/24     2     Inter-area   192.168.1.1     10.1.1.1        0.0.0.1
 Routing for ASEs
 Destination    Cost      Type       Tag         NextHop         AdvRouter
 10.0.0.0/8    3         Type1      1           192.168.1.1     10.2.2.2
 Total Nets: 3
 Intra Area: 1  Inter Area: 1  ASE: 1  NSSA: 0
[SwitchA] ospf 1
[SwitchA-ospf-1] area 1
[SwitchA-ospf-1-area-0.0.0.1] stub   //Configure Area 1 as a stub area. All the routers in Area 1 must have the stub command configured.
[SwitchA-ospf-1-area-0.0.0.1] quit
[SwitchA-ospf-1] quit
[SwitchC] ospf 1
[SwitchC-ospf-1] area 1
[SwitchC-ospf-1-area-0.0.0.1] stub   //Configure Area 1 as a stub area. All the routers in Area 1 must have the stub command configured.
[SwitchC-ospf-1-area-0.0.0.1] quit
[SwitchC-ospf-1] quit
# Check the OSPF routing table on SwitchC. The command output shows that the OSPF routing table does not contain the AS external route 10.0.0.0/8 but contains a default route to external networks.
[SwitchC] display ospf routing
          OSPF Process 1 with Router ID 10.3.3.3
                   Routing Tables
 Routing for Network
 Destination        Cost  Type       NextHop         AdvRouter       Area
 192.168.1.0/24     1     Transit    192.168.1.2     10.3.3.3         0.0.0.1
 0.0.0.0/0          2     Inter-area 192.168.1.1     10.1.1.1         0.0.0.1
 192.168.0.0/24     2     Inter-area 192.168.1.1     10.1.1.1         0.0.0.1
 Total Nets: 3
 Intra Area: 1  Inter Area: 2  ASE: 0  NSSA: 0
[SwitchA] ospf 1
[SwitchA-ospf-1] area 1
[SwitchA-ospf-1-area-0.0.0.1] stub no-summary   //Configure Area 1 as a totally stub area. An ABR in Area 1 must have the stub no-summary command configured, while other routers in Area 1 must have the stub command configured.
[SwitchA-ospf-1-area-0.0.0.1] quit
[SwitchA-ospf-1] quit
# Check the OSPF routing table on SwitchC. The command output shows that the OSPF routing table contains only an intra-area OSPF route and a default route to external networks but does not contain the Inter-Area route 192.168.0.0/24.
[SwitchC] display ospf routing
          OSPF Process 1 with Router ID 10.3.3.3
                   Routing Tables
 Routing for Network
 Destination        Cost  Type       NextHop         AdvRouter       Area
 192.168.1.0/24     1     Transit    192.168.1.2     10.3.3.3         0.0.0.1
 0.0.0.0/0          2     Inter-area 192.168.1.1     10.1.1.1         0.0.0.1
 Total Nets: 2
 Intra Area: 1  Inter Area: 1  ASE: 0  NSSA: 0
SwitchA configuration file
#
sysname SwitchA
#
vlan batch 10 20
#
interface Vlanif10
 ip address 192.168.0.1 255.255.255.0
#
interface Vlanif20
 ip address 192.168.1.1 255.255.255.0
#
interface GigabitEthernet0/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface GigabitEthernet0/0/2
 port link-type trunk
 port trunk allow-pass vlan 20
#
ospf 1 router-id 10.1.1.1
 area 0.0.0.0
  network 192.168.0.0 0.0.0.255
 area 0.0.0.1
  network 192.168.1.0 0.0.0.255
  stub no-summary
#
return
SwitchB configuration file
#
sysname SwitchB
#
vlan batch 10
#
interface Vlanif10
 ip address 192.168.0.2 255.255.255.0
#
interface GigabitEthernet0/0/1
 port link-type trunk
 port trunk allow-pass vlan 10
#
ospf 1 router-id 10.2.2.2
 import-route static type 1
 area 0.0.0.0
  network 192.168.0.0 0.0.0.255
#
ip route-static 10.0.0.0 255.0.0.0 NULL0
#
return
SwitchC configuration file
#
sysname SwitchC
#
vlan batch 20
#
interface Vlanif20
 ip address 192.168.1.2 255.255.255.0
#
interface GigabitEthernet0/0/1
 port link-type trunk
 port trunk allow-pass vlan 20
#
ospf 1 router-id 10.3.3.3
 area 0.0.0.1
  network 192.168.1.0 0.0.0.255
  stub
#
return
Select the content with the mouse pointer to quickly report the problem.
