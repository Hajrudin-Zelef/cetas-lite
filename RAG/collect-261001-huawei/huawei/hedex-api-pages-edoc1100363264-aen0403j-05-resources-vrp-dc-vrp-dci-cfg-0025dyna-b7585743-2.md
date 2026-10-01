---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100363264-aen0403j-05-resources-vrp-dc-vrp-dci-cfg-0025dyna-b7585743-2
title: "Configure PE1."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100363264-aen0403j-05-resources-vrp-dc-vrp-dci-cfg-0025dyna-b7585743.md
source_anchor: ""
source_lines: [199, 477]
sha256: 12163a3480bea710709d48da213c8503586d4b47ea2cbb0440a129739990633e
---

# Configure PE1.

  apply-label per-instance
  vpn-target 1:1 export-extcommunity
  vpn-target 1:1 import-extcommunity
  vpn-target 1:1 export-extcommunity evpn
  vpn-target 1:1 import-extcommunity evpn
 vxlan vni 100
#
bridge-domain 10
 vxlan vni 10 split-horizon-mode
 evpn binding vpn-instance evrf1
#
isis 1
 network-entity 10.0000.0000.0010.00
 frr
#
interface Eth-Trunk10
 esi 0000.0000.0000.0000.1111
#
interface Eth-Trunk10.1 mode l2
 encapsulation dot1q vid 10
 rewrite pop single
 bridge-domain 10
#
interface Vbdif10
 ip binding vpn-instance vpn1
 ip address 10.1.10.1 255.255.255.0
 mac-address 00e0-fc12-3456
 vxlan anycast-gateway enable
 arp collect host enable
#
interface GigabitEthernet0/1/1
 undo shutdown
 ip address 10.1.20.1 255.255.255.0
 isis enable 1
#
interface GigabitEthernet0/1/2
 undo shutdown
 eth-trunk 10
#
interface GigabitEthernet0/1/3
 undo shutdown
 ip address 10.1.1.1 255.255.255.0
 isis enable 1
#
interface LoopBack1
 ip address 1.1.1.1 255.255.255.255
 isis enable 1
#
interface LoopBack2
 ip address 3.3.3.3 255.255.255.255
 isis enable 1
#
interface Nve1
 source 3.3.3.3
 bypass source 1.1.1.1
 mac-address 00e0-fc12-7890
 vni 10 head-end peer-list protocol bgp
#
bgp 100
 peer 2.2.2.2 as-number 100
 peer 2.2.2.2 connect-interface LoopBack1
 peer 4.4.4.4 as-number 100
 peer 4.4.4.4 connect-interface LoopBack1
 #
 ipv4-family unicast
  undo synchronization
  peer 2.2.2.2 enable
  peer 4.4.4.4 enable
 #
 ipv4-family vpn-instance vpn1
  import-route direct
  auto-frr
  advertise l2vpn evpn
 #
 l2vpn-family evpn
  undo policy vpn-target
  peer 2.2.2.2 enable
  peer 2.2.2.2 advertise irb
  peer 2.2.2.2 advertise encap-type vxlan
  peer 4.4.4.4 enable
  peer 4.4.4.4 advertise irb
  peer 4.4.4.4 advertise encap-type vxlan
#
return
PE2 configuration file
#
sysname PE2
#
evpn
 bypass-vxlan enable
#
evpn vpn-instance evrf1 bd-mode
 route-distinguisher 11:11
 vpn-target 1:1 export-extcommunity
 vpn-target 1:1 import-extcommunity
#
ip vpn-instance vpn1
 ipv4-family
  route-distinguisher 1:1
  apply-label per-instance
  vpn-target 1:1 export-extcommunity
  vpn-target 1:1 import-extcommunity
  vpn-target 1:1 export-extcommunity evpn
  vpn-target 1:1 import-extcommunity evpn
 vxlan vni 100
#
bridge-domain 10
 vxlan vni 10 split-horizon-mode
 evpn binding vpn-instance evrf1
#
isis 1
 network-entity 10.0000.0000.0020.00
 frr
#
interface Eth-Trunk10
 esi 0000.0000.0000.0000.1111
#
interface Eth-Trunk10.1 mode l2
 encapsulation dot1q vid 10
 rewrite pop single
 bridge-domain 10
#
interface Vbdif10
 ip binding vpn-instance vpn1
 ip address 10.1.10.1 255.255.255.0
 mac-address 00e0-fc12-3456
 vxlan anycast-gateway enable
 arp collect host enable
#
interface GigabitEthernet0/1/1
 undo shutdown
 ip address 10.1.20.2 255.255.255.0
 isis enable 1
#
interface GigabitEthernet0/1/2
 undo shutdown
 eth-trunk 10
#
interface GigabitEthernet0/1/3
 undo shutdown
 ip address 10.1.2.1 255.255.255.0
 isis enable 1
#
interface LoopBack1
 ip address 2.2.2.2 255.255.255.255
 isis enable 1
#
interface LoopBack2
 ip address 3.3.3.3 255.255.255.255
 isis enable 1
#
interface Nve1
 source 3.3.3.3
 bypass source 2.2.2.2
 mac-address 00e0-fc12-7890
 vni 10 head-end peer-list protocol bgp
#
bgp 100
 peer 1.1.1.1 as-number 100
 peer 1.1.1.1 connect-interface LoopBack1
 peer 4.4.4.4 as-number 100
 peer 4.4.4.4 connect-interface LoopBack1
 #
 ipv4-family unicast
  undo synchronization
  peer 1.1.1.1 enable
  peer 4.4.4.4 enable
 #
 ipv4-family vpn-instance vpn1
  import-route direct
  auto-frr
  advertise l2vpn evpn
 #
 l2vpn-family evpn
  undo policy vpn-target
  peer 1.1.1.1 enable
  peer 1.1.1.1 advertise irb
  peer 1.1.1.1 advertise encap-type vxlan
  peer 4.4.4.4 enable
  peer 4.4.4.4 advertise irb
  peer 4.4.4.4 advertise encap-type vxlan
#
return
CE1 configuration file
# 
sysname CE1
#
vlan batch 10
#
interface Eth-Trunk10
 portswitch
 port link-type trunk
 port trunk allow-pass vlan 10
#
interface GigabitEthernet0/1/1
 undo shutdown
 eth-trunk 10
#
interface GigabitEthernet0/1/2
 undo shutdown
 eth-trunk 10
#
return
CPE configuration file
#
sysname CPE
#
evpn vpn-instance evrf1 bd-mode
 route-distinguisher 11:11
 vpn-target 1:1 export-extcommunity
 vpn-target 1:1 import-extcommunity
#
ip vpn-instance vpn1
 ipv4-family
  route-distinguisher 1:1
  apply-label per-instance
  vpn-target 1:1 export-extcommunity
  vpn-target 1:1 import-extcommunity
  vpn-target 1:1 export-extcommunity evpn
  vpn-target 1:1 import-extcommunity evpn
 vxlan vni 100
#
bridge-domain 20
 vxlan vni 20 split-horizon-mode
 evpn binding vpn-instance evrf1
#
isis 1
 network-entity 20.0000.0000.0001.00
 frr
#
interface Vbdif20
 ip binding vpn-instance vpn1
 ip address 10.1.30.1 255.255.255.0
 vxlan anycast-gateway enable
 arp collect host enable
#
interface GigabitEthernet0/1/1
 undo shutdown
 ip address 10.1.1.2 255.255.255.0
 isis enable 1
#
interface GigabitEthernet0/1/2
 undo shutdown
 ip address 10.1.2.2 255.255.255.0
 isis enable 1
#
interface LoopBack1
 ip address 4.4.4.4 255.255.255.255
 isis enable 1
#
interface Nve1
 source 4.4.4.4 
 vni 20 head-end peer-list protocol bgp
#
bgp 100
 peer 1.1.1.1 as-number 100
 peer 1.1.1.1 connect-interface LoopBack1
 peer 2.2.2.2 as-number 100
 peer 2.2.2.2 connect-interface LoopBack1
 #
 ipv4-family unicast
  undo synchronization
  peer 1.1.1.1 enable
  peer 2.2.2.2 enable
 #
 ipv4-family vpn-instance vpn1
  import-route direct
  advertise l2vpn evpn
 #
 l2vpn-family evpn
  undo policy vpn-target
  peer 1.1.1.1 enable
  peer 1.1.1.1 advertise irb
  peer 1.1.1.1 advertise encap-type vxlan
  peer 2.2.2.2 enable
  peer 2.2.2.2 advertise irb
  peer 2.2.2.2 advertise encap-type vxlan
#
return
