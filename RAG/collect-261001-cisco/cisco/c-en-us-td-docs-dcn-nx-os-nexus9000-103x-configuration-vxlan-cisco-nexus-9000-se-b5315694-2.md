---
id: collect-261001-cisco/cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-b5315694-2
title: "c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-b5315694"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-b5315694.md
source_anchor: ""
source_lines: [234, 275]
sha256: fb2d6a0ca97fc22f1a06190cfb1d98e7f9346af78279428b8d224dd3b2c1642a
---

# c-en-us-td-docs-dcn-nx-os-nexus9000-103x-configuration-vxlan-cisco-nexus-9000-se-b5315694

interface Vlan3300
  no shutdown
  vrf member vrf2
  ip forward
  ip nat outside
 
interface Ethernet1/16
  switchport
  switchport mode trunk
 
interface Ethernet1/43
  switchport
  switchport mode trunk
 
vrf context vrf1
  vni 33200
  rd auto
  address-family ipv4 unicast
    route-target both auto
    route-target both auto evpn
vrf context vrf2
  vni 33300
  rd auto
  address-family ipv4 unicast
    route-target both auto
    route-target both auto evpn
 
 
router bgp 100
  vrf vrf1
    address-family ipv4 unicast
      network 172.21.1.20/32
      advertise l2vpn evpn
vrf vrf2
    address-family ipv4 unicast
      network 172.31.1.20/32
     advertise l2vpn evpn
The following show command provides the display of insulation policies configured in the switch for EVPN Distributed NAT.
show ip nat translations 
Pro Inside global Inside local Outside local Outside global
any 174.2.216.2 42.2.216.2 --- ---
any 174.3.217.2 42.3.217.2 --- ---
