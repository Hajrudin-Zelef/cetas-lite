---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0028-html-305059cd-1
title: "Configure CE1."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0028-html-305059cd.md
source_anchor: ""
source_lines: [1, 200]
sha256: 95f5d81456861e1808b7d6f43915b218a89cdefc4e61240e403a6d8acc7ffce6
---

# Configure CE1.

In Figure 1:
PE1 and PE2 reside on the MPLS backbone network.
R1 connects CE1 and PE1 over the public network.
CE2 is directly connected to PE2.
CE1 and CE2 reside on the same VPN and are reachable to each other.
PE1 is indirectly connected to CE1. Therefore, no VPN instance can be bound to the physical interface of PE1. A GRE tunnel is set up between CE1 and PE1 and this tunnel traverses the public network. On PE1, bind the GRE tunnel to a VPN to connect CE1 to the VPN using the GRE tunnel.
AR611-S, AR611W-S, AR611E-S, AR611-LTE4EA, AR611, AR611W, AR611W-LTE4CN, AR611W-LTE6EA, AR617VW, AR617VW-LTE4, AR617VW-LTE4EA, and AR651C do not support this function.
The RU-5G-101 does not support this function.
The configuration roadmap is as follows:
Run OSPF process 10 on PE1 and PE2 to implement interworking between them, and enable MPLS.
Run OSPF process 20 on CE1, R1, and PE1 to implement interworking among them.
Set up a GRE tunnel between CE1 and PE1.
Create vpn1 on PE1 and PE2. On PE1, bind vpn1 to the GRE tunnel interface. On PE2, bind vpn1 to the physical interface connected to CE2.
Configure IS-IS routes between CE1 and PE1 and between CE2 and PE2, so that each pair of devices can communicate.
Run BGP on the PEs to implement interworking between CE1 and CE2.
# Configure CE1.
<Huawei> system-view
[Huawei] sysname CE1
[CE1] interface gigabitethernet 1/0/0
[CE1-GigabitEthernet1/0/0] ip address 10.1.1.2 24
[CE1-GigabitEthernet1/0/0] quit
[CE1] interface gigabitethernet 2/0/0
[CE1-GigabitEthernet2/0/0] ip address 30.1.1.1 24
[CE1-GigabitEthernet2/0/0] quit
# Configure R1.
<Huawei> system-view
[Huawei] sysname R1
[R1] interface gigabitethernet 1/0/0
[R1-GigabitEthernet1/0/0] ip address 30.1.1.2 24
[R1-GigabitEthernet1/0/0] quit
[R1] interface gigabitethernet 2/0/0
[R1-GigabitEthernet2/0/0] ip address 50.1.1.1 24
[R1-GigabitEthernet2/0/0] quit
# Configure PE1.
<Huawei> system-view
[Huawei] sysname PE1
[PE1] interface gigabitethernet 1/0/0
[PE1-GigabitEthernet1/0/0] ip address 50.1.1.2 24
[PE1-GigabitEthernet1/0/0] quit
[PE1] interface gigabitethernet 2/0/0
[PE1-GigabitEthernet2/0/0] ip address 110.1.1.1 24
[PE1-GigabitEthernet2/0/0] quit
[PE1] interface loopback 1
[PE1-LoopBack1] ip address 1.1.1.9 32
[PE1-LoopBack1] quit
# Configure IP addresses for interfaces on PE2 except the interface to be bound to a VPN instance, because all configurations on this interface are deleted when the interface is bound to a VPN instance.
<Huawei> system-view
[Huawei] sysname PE2
[PE2] interface gigabitethernet 1/0/0
[PE2-GigabitEthernet1/0/0] ip address 110.1.1.2 24
[PE2-GigabitEthernet1/0/0] quit
[PE2] interface loopback 1
[PE2-LoopBack1] ip address 3.3.3.9 32
[PE2-LoopBack1] quit
# Configure CE2.
<Huawei> system-view
[Huawei] sysname CE2
[CE2] interface gigabitethernet 1/0/0
[CE2-GigabitEthernet1/0/0] ip address 11.1.1.1 24
[CE2-GigabitEthernet1/0/0] quit
[CE2] interface gigabitethernet 2/0/0
[CE2-GigabitEthernet2/0/0] ip address 10.2.1.2 24
[CE2-GigabitEthernet2/0/0] quit
# On PE1, enable MPLS LDP, and run OSPF process 10 to configure reachable routes between the PEs. LSPs are set up automatically.
[PE1] mpls lsr-id 1.1.1.9
[PE1] mpls
[PE1-mpls] lsp-trigger all
[PE1-mpls] quit
[PE1] mpls ldp
[PE1-mpls-ldp] quit
[PE1] ospf 10
[PE1-ospf-10] area 0
[PE1-ospf-10-area-0.0.0.0] network 1.1.1.9 0.0.0.0
[PE1-ospf-10-area-0.0.0.0] network 110.1.1.0 0.0.0.255
[PE1-ospf-10-area-0.0.0.0] quit
[PE1-ospf-10] quit
[PE1] interface gigabitethernet 2/0/0
[PE1-GigabitEthernet2/0/0] mpls
[PE1-GigabitEthernet2/0/0] mpls ldp
[PE1-GigabitEthernet2/0/0] quit
# On PE2, enable MPLS LDP, and run OSPF process 10 to configure reachable routes between the PEs. LSPs are set up automatically.
[PE2] mpls lsr-id 3.3.3.9
[PE2] mpls
[PE2-mpls] lsp-trigger all
[PE2-mpls] quit
[PE2] mpls ldp
[PE2-mpls-ldp] quit
[PE2] ospf 10
[PE2-ospf-10] area 0
[PE2-ospf-10-area-0.0.0.0] network 3.3.3.9 0.0.0.0
[PE2-ospf-10-area-0.0.0.0] network 110.1.1.0 0.0.0.255
[PE2-ospf-10-area-0.0.0.0] quit
[PE2-ospf-10] quit
[PE2] interface gigabitethernet 1/0/0
[PE2-GigabitEthernet1/0/0] mpls
[PE2-GigabitEthernet1/0/0] mpls ldp
[PE2-GigabitEthernet1/0/0] quit
[PE1] ip vpn-instance vpn1
[PE1-vpn-instance-vpn1] route-distinguisher 100:1
[PE1-vpn-instance-vpn1-af-ipv4] vpn-target 111:1 export-extcommunity
[PE1-vpn-instance-vpn1-af-ipv4] vpn-target 111:1 import-extcommunity
[PE1-vpn-instance-vpn1-af-ipv4] quit
[PE1-vpn-instance-vpn1] quit
[PE1] interface tunnel 0/0/1
[PE1-Tunnel0/0/1] ip binding vpn-instance vpn1 
[PE1-Tunnel0/0/1] ip address 2.2.2.2 255.255.255.0
[PE1-Tunnel0/0/1] quit
[PE2] ip vpn-instance vpn1
[PE2-vpn-instance-vpn1] route-distinguisher 200:1
[PE2-vpn-instance-vpn1-af-ipv4] vpn-target 111:1 export-extcommunity
[PE2-vpn-instance-vpn1-af-ipv4] vpn-target 111:1 import-extcommunity
[PE2-vpn-instance-vpn1-af-ipv4] quit
[PE2-vpn-instance-vpn1] quit
[PE2] interface gigabitethernet 2/0/0
[PE2-GigabitEthernet2/0/0] ip binding vpn-instance vpn1 
[PE2-GigabitEthernet2/0/0] ip address 11.1.1.2 255.255.255.0
[PE2-GigabitEthernet2/0/0] quit
[CE1] interface tunnel 0/0/1
[CE1-Tunnel0/0/1] tunnel-protocol gre
[CE1-Tunnel0/0/1] source 30.1.1.1
[CE1-Tunnel0/0/1] destination 50.1.1.2
[CE1-Tunnel0/0/1] ip address 2.2.2.1 24
[CE1-Tunnel0/0/1] quit
[PE1] interface tunnel 0/0/1
[PE1-Tunnel0/0/1] tunnel-protocol gre
[PE1-Tunnel0/0/1] source 50.1.1.2
[PE1-Tunnel0/0/1] destination 30.1.1.1
[PE1-Tunnel0/0/1] quit
[CE1] ospf 20
[CE1-ospf-20] area 0
[CE1-ospf-20-area-0.0.0.0] network 30.1.1.0 0.0.0.255
[CE1-ospf-20-area-0.0.0.0] quit
[CE1-ospf-20] quit
[R1] ospf 20
[R1-ospf-20] area 0
[R1-ospf-20-area-0.0.0.0] network 30.1.1.0 0.0.0.255
[R1-ospf-20-area-0.0.0.0] network 50.1.1.0 0.0.0.255
[R1-ospf-20-area-0.0.0.0] quit
[R1-ospf-20] quit
[PE1] ospf 20
[PE1-ospf-20] area 0
[PE1-ospf-20-area-0.0.0.0] network 50.1.1.0 0.0.0.255
[PE1-ospf-20-area-0.0.0.0] quit
[PE1-ospf-20] quit
[CE1] isis 50
[CE1-isis-50] network-entity 50.0000.0000.0001.00
[CE1-isis-50] quit
[CE1] interface gigabitethernet 1/0/0
[CE1-GigabitEthernet1/0/0] isis enable 50
[CE1-GigabitEthernet1/0/0] quit
[CE1] interface tunnel 0/0/1
[CE1-Tunnel0/0/1] isis enable 50
[CE1-Tunnel0/0/1] quit
[PE1] isis 50 vpn-instance vpn1
[PE1-isis-50] network-entity 50.0000.0000.0002.00
[PE1-isis-50] quit
[PE1] interface tunnel 0/0/1
[PE1-Tunnel0/0/1] isis enable 50
[PE1-Tunnel0/0/1] quit
[CE2] isis 50
[CE2-isis-50] network-entity 50.0000.0000.0004.00
[CE2-isis-50] quit
[CE2] interface gigabitethernet 1/0/0
[CE2-GigabitEthernet1/0/0] isis enable 50
[CE2-GigabitEthernet1/0/0] quit
[CE2] interface gigabitethernet 2/0/0
[CE2-GigabitEthernet2/0/0] isis enable 50
[CE2-GigabitEthernet2/0/0] quit
# Configure PE2.
[PE2] isis 50 vpn-instance vpn1
[PE2-isis-50] network-entity 50.0000.0000.0003.00
[PE2-isis-50] quit
[PE2] interface gigabitethernet 2/0/0
[PE2-GigabitEthernet2/0/0] isis enable 50
[PE2-GigabitEthernet2/0/0] quit
# On PE1, configure an IBGP peer relationship with PE2 using a loopback interface to exchange VPN IPv4 route information.
[PE1] bgp 100
[PE1-bgp] peer 3.3.3.9 as-number 100
[PE1-bgp] peer 3.3.3.9 connect-interface loopback 1
[PE1-bgp] ipv4-family vpnv4
[PE1-bgp-af-vpnv4] peer 3.3.3.9 enable
[PE1-bgp-af-vpnv4] quit
# Import IS-IS routes to vpn1.
[PE1-bgp] ipv4-family vpn-instance vpn1
[PE1-bgp-vpn1] import-route isis 50
# On PE2, configure an IBGP peer relationship with PE1 using a loopback interface to exchange VPN IPv4 route information.
[PE2] bgp 100
[PE2-bgp] peer 1.1.1.9 as-number 100
[PE2-bgp] peer 1.1.1.9 connect-interface loopback 1
[PE2-bgp] ipv4-family vpnv4
[PE2-bgp-af-vpnv4] peer 1.1.1.9 enable
[PE2-bgp-af-vpnv4] quit
[PE2-bgp] ipv4-family vpn-instance vpn1
[PE2-bgp-vpn1] import-route isis 50
[PE1] isis 50
[PE1-isis-50] import-route bgp
[PE2] isis 50
[PE2-isis-50] import-route bgp
# After the configuration is complete, CE1 and CE2 have reachable routes to each other. The command output on CE1 is used as an example.
<CE1> display ip routing-table 41.1.1.0
