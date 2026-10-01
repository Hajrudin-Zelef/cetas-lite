---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0028-html-305059cd-2
title: "Configure CE1."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-gre-0028-html-305059cd.md
source_anchor: ""
source_lines: [201, 380]
sha256: 40a1ee6869f80b0419692d8fb70f6028e7ab067d295ce5b8b884f347a9f4d12c
---

# Configure CE1.

Route Flags: R - relay, D - download to fib, T - to vpn-instance
------------------------------------------------------------------------------
Routing Table : Public
Summary Count : 1
Destination/Mask    Proto   Pre  Cost      Flags NextHop         Interface
       41.1.1.0/24  ISIS-L2 15   74          D   2.2.2.2         Tunnel0/0/1
Configuration file of CE1
#
 sysname CE1
#
isis 50
 network-entity 50.0000.0000.0001.00
#
interface GigabitEthernet1/0/0
 ip address 10.1.1.2 255.255.255.0
 isis enable 50
#
interface GigabitEthernet2/0/0
 ip address 30.1.1.1 255.255.255.0
#
interface Tunnel0/0/1
 ip address 2.2.2.1 255.255.255.0
 tunnel-protocol gre
 source 30.1.1.1
 destination 50.1.1.2
 isis enable 50
#
ospf 20
 area 0.0.0.0
  network 30.1.1.0 0.0.0.255
#
return
Configurations file of R1
#
 sysname R1
#
interface GigabitEthernet1/0/0
ip address 30.1.1.2 255.255.255.0
#
interface GigabitEthernet2/0/0
ip address 50.1.1.1 255.255.255.0
#
ospf 20
 area 0.0.0.0
  network 30.1.1.0 0.0.0.255
  network 50.1.1.0 0.0.0.255
#
return
Configuration file of PE1
#
 sysname PE1
#
ip vpn-instance vpn1
 route-distinguisher 100:1
 vpn-target 111:1 export-extcommunity
 vpn-target 111:1 import-extcommunity
#
mpls lsr-id 1.1.1.9
mpls
 lsp-trigger all
#
mpls ldp
#
isis 50 vpn-instance vpn1
 network-entity 50.0000.0000.0002.00
 import-route bgp
#
interface GigabitEthernet1/0/0
 ip address 50.1.1.2 255.255.255.0
#
interface GigabitEthernet2/0/0
 ip address 110.1.1.1 255.255.255.0
 mpls
 mpls ldp
#
interface LoopBack1
 ip address 1.1.1.9 255.255.255.255
#
interface Tunnel0/0/1
 ip binding vpn-instance vpn1
 ip address 2.2.2.2 255.255.255.0
 tunnel-protocol gre
 source 50.1.1.2
 destination 30.1.1.1
 isis enable 50
#
bgp 100
 peer 3.3.3.9 as-number 100
 peer 3.3.3.9 connect-interface LoopBack1
 #
 ipv4-family unicast
  undo synchronization
  peer 3.3.3.9 enable
 #
 ipv4-family vpnv4
  policy vpn-target
  peer 3.3.3.9 enable
 #
 ipv4-family vpn-instance vpn1
  import-route isis 50
#
ospf 10
 area 0.0.0.0
  network 1.1.1.9 0.0.0.0
  network 110.1.1.0 0.0.0.255
#
ospf 20
 area 0.0.0.0
  network 50.1.1.0 0.0.0.255
#
return
Configuration file of PE2
#
 sysname PE2
#
ip vpn-instance vpn1
 route-distinguisher 200:1
 vpn-target 111:1 export-extcommunity
 vpn-target 111:1 import-extcommunity
#
mpls lsr-id 3.3.3.9
mpls
 lsp-trigger all
#
mpls ldp
#
isis 50 vpn-instance vpn1
 network-entity 50.0000.0000.0003.00
 import-route bgp
#
interface GigabitEthernet1/0/0
 ip address 110.1.1.2 255.255.255.0
 mpls
 mpls ldp
#
interface GigabitEthernet2/0/0
 ip binding vpn-instance vpn1
 ip address 11.1.1.2 255.255.255.0
 isis enable 50
#
interface LoopBack1
 ip address 3.3.3.9 255.255.255.255
#
bgp 100
 peer 1.1.1.9 as-number 100
 peer 1.1.1.9 connect-interface LoopBack1
 #
 ipv4-family unicast
  undo synchronization
  peer 1.1.1.9 enable
 #
 ipv4-family vpnv4
  policy vpn-target
  peer 1.1.1.9 enable
 #
 ipv4-family vpn-instance vpn1
  import-route isis 50
#
ospf 10
 area 0.0.0.0
  network 3.3.3.9 0.0.0.0
  network 110.1.1.0 0.0.0.255
#
return
Configuration file of CE2
#
 sysname CE2
#
isis 50
 network-entity 50.0000.0000.0004.00
#
interface GigabitEthernet1/0/0
 ip address 11.1.1.1 255.255.255.0
 isis enable 50
#
interface GigabitEthernet2/0/0
 ip address 10.2.1.2 255.255.255.0
 isis enable 50
#
return
