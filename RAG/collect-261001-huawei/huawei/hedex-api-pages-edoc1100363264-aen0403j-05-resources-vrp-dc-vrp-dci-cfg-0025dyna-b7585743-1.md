---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100363264-aen0403j-05-resources-vrp-dc-vrp-dci-cfg-0025dyna-b7585743-1
title: "Configure PE1."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100363264-aen0403j-05-resources-vrp-dc-vrp-dci-cfg-0025dyna-b7585743.md
source_anchor: ""
source_lines: [1, 198]
sha256: fac83e4945ff91bd9650d17549ffd15342f16425251d797f26960b48231fbf8a
---

# Configure PE1.

In a scenario where a DC needs to communicate with an enterprise site at Layer 3 and a CE is dual-homed to the VXLAN network, the carrier can enhance VXLAN access reliability to improve the stability of user services and enable rapid convergence in the case of a fault.
On the network shown in Figure 1, a VXLAN tunnel is required to be dynamically deployed using BGP EVPN between the CPE and the PE1-PE2 pair. An EVPN peer relationship needs to be established between PE1 and PE2 to deploy a bypass VXLAN tunnel. In addition, CE1 needs to be dual-homed to PE1 and PE2 (which are configured with the same anycast VTEP address) to implement the active-active function. If one PE fails, this function allows traffic to be quickly switched to the other PE.
In this example, interfaces 1 through 3 represent GigabitEthernet0/1/1, GigabitEthernet0/1/2, and GigabitEthernet0/1/3, respectively.
| Table 1 Interface IP addresses |  |  | 
|---|---|---|
| Device | Interface | IP Address | 
|---|---|---|
| PE1 | GigabitEthernet0/1/1 | 10.1.20.1/24 | 
|  | GigabitEthernet0/1/2 | - | 
|  | GigabitEthernet0/1/3 | 10.1.1.1/24 | 
|  | Loopback 1 | 1.1.1.1/32 | 
|  | Loopback 2 | 3.3.3.3/32 | 
| PE2 | GigabitEthernet0/1/1 | 10.1.20.2/24 | 
|  | GigabitEthernet0/1/2 | - | 
|  | GigabitEthernet0/1/3 | 10.1.2.1/24 | 
|  | Loopback 1 | 2.2.2.2/32 | 
|  | Loopback 2 | 3.3.3.3/32 | 
| CE1 | GigabitEthernet0/1/1 | - | 
|  | GigabitEthernet0/1/2 | - | 
| CPE | GigabitEthernet0/1/1 | 10.1.1.2/24 | 
|  | GigabitEthernet0/1/2 | 10.1.2.2/24 | 
|  | Loopback 1 | 4.4.4.4/32 | 
The configuration roadmap is as follows:
For detailed configurations, see Configuration Files.
# Configure PE1.
<PE1> system-view
[~PE1] interface eth-trunk 10
[*PE1-Eth-Trunk10] trunkport gigabitethernet 0/1/2
[*PE1-Eth-Trunk10] esi 0000.0000.0000.0000.1111
[*PE1-Eth-Trunk10] quit
[*PE1] bridge-domain 10
[*PE1-bd10] vxlan vni 10 split-horizon-mode
[*PE1-bd10] quit
[*PE1] interface eth-trunk 10.1 mode l2
[*PE1-Eth-Trunk10.1] encapsulation dot1q vid 10
[*PE1-Eth-Trunk10.1] rewrite pop single
[*PE1-Eth-Trunk10.1] bridge-domain 10
[*PE1-Eth-Trunk10.1] quit
[*PE1] commit
# Configure PE2.
<PE2> system-view
[~PE2] interface eth-trunk 10
[*PE2-Eth-Trunk10] trunkport gigabitethernet 0/1/2
[*PE2-Eth-Trunk10] esi 0000.0000.0000.0000.1111
[*PE2-Eth-Trunk10] quit
[*PE2] bridge-domain 10
[*PE2-bd10] vxlan vni 10 split-horizon-mode
[*PE2-bd10] quit
[*PE2] interface eth-trunk 10.1 mode l2
[*PE2-Eth-Trunk10.1] encapsulation dot1q vid 10
[*PE2-Eth-Trunk10.1] rewrite pop single
[*PE2-Eth-Trunk10.1] bridge-domain 10
[*PE2-Eth-Trunk10.1] quit
[*PE2] commit
[~PE1] evpn
[*PE1-evpn] bypass-vxlan enable
[*PE1-evpn] quit
[*PE1] commit
The configuration of PE2 is similar to that of PE1. For detailed configurations, see Configuration Files.
# Configure the CPE.
<CPE> system-view
[~CPE] evpn vpn-instance evrf1 bd-mode
[*CPE-evpn-instance-evrf1] route-distinguisher 11:11
[*CPE-evpn-instance-evrf1] vpn-target 1:1 import-extcommunity
[*CPE-evpn-instance-evrf1] vpn-target 1:1 export-extcommunity
[*CPE-evpn-instance-evrf1] quit
[*CPE] bridge-domain 20
[*CPE-bd20] vxlan vni 20 split-horizon-mode
[*CPE-bd20] evpn binding vpn-instance evrf1
[*CPE-bd20] quit
[*CPE] commit
[~PE1] evpn vpn-instance evrf1 bd-mode
[*PE1-evpn-instance-evrf1] route-distinguisher 11:11
[*PE1-evpn-instance-evrf1] vpn-target 1:1 import-extcommunity
[*PE1-evpn-instance-evrf1] vpn-target 1:1 export-extcommunity
[*PE1-evpn-instance-evrf1] quit
[*PE1] bridge-domain 10
[*PE1-bd10] evpn binding vpn-instance evrf1
[*PE1-bd10] quit
[*PE1] commit
[~CPE] ip vpn-instance vpn1
[*CPE-vpn-instance-vpn1] vxlan vni 100
[*CPE-vpn-instance-vpn1] ipv4-family
[*CPE-vpn-instance-vpn1-af-ipv4] route-distinguisher 1:1
[*CPE-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 import-extcommunity
[*CPE-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 export-extcommunity
[*CPE-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 import-extcommunity evpn
[*CPE-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 export-extcommunity evpn
[*CPE-vpn-instance-vpn1-af-ipv4] quit
[*CPE-vpn-instance-vpn1] quit
[*CPE] commit
[~PE1] ip vpn-instance vpn1
[*PE1-vpn-instance-vpn1] vxlan vni 100
[*PE1-vpn-instance-vpn1] ipv4-family
[*PE1-vpn-instance-vpn1-af-ipv4] route-distinguisher 1:1
[*PE1-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 import-extcommunity
[*PE1-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 export-extcommunity
[*PE1-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 import-extcommunity evpn
[*PE1-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 export-extcommunity evpn
[*PE1-vpn-instance-vpn1-af-ipv4] quit
[*PE1-vpn-instance-vpn1] quit
[~CPE] bgp 100
[*CPE-bgp] peer 1.1.1.1 as-number 100
[*CPE-bgp] peer 1.1.1.1 connect-interface LoopBack 1
[*CPE-bgp] peer 2.2.2.2 as-number 100
[*CPE-bgp] peer 2.2.2.2 connect-interface LoopBack 1
[*CPE-bgp] ipv4-family unicast
[*CPE-bgp-af-ipv4] undo synchronization
[*CPE-bgp-af-ipv4] peer 1.1.1.1 enable
[*CPE-bgp-af-ipv4] peer 2.2.2.2 enable 
[*CPE-bgp-af-ipv4] quit
[*CPE-bgp] ipv4-family vpn-instance vpn1
[*CPE-bgp-vpn1] advertise l2vpn evpn
[*CPE-bgp-vpn1] import-route direct
[*CPE-bgp-vpn1] quit
[*CPE-bgp] l2vpn-family evpn
[*CPE-bgp-af-evpn] undo policy vpn-target
[*CPE-bgp-af-evpn] peer 1.1.1.1 enable
[*CPE-bgp-af-evpn] peer 1.1.1.1 advertise irb
[*CPE-bgp-af-evpn] peer 1.1.1.1 advertise encap-type vxlan
[*CPE-bgp-af-evpn] peer 2.2.2.2 enable
[*CPE-bgp-af-evpn] peer 2.2.2.2 advertise irb
[*CPE-bgp-af-evpn] peer 2.2.2.2 advertise encap-type vxlan
[*CPE-bgp-af-evpn] quit
[*CPE-bgp] quit
[*CPE] commit
[~PE1] bgp 100
[*PE1-bgp] peer 2.2.2.2 as-number 100
[*PE1-bgp] peer 2.2.2.2 connect-interface LoopBack 1
[*PE1-bgp] peer 4.4.4.4 as-number 100
[*PE1-bgp] peer 4.4.4.4 connect-interface LoopBack 1
[*PE1-bgp] ipv4-family unicast
[*PE1-bgp-af-ipv4] undo synchronization
[*PE1-bgp-af-ipv4] peer 2.2.2.2 enable
[*PE1-bgp-af-ipv4] peer 4.4.4.4 enable 
[*PE1-bgp-af-ipv4] quit
[*PE1-bgp] ipv4-family vpn-instance vpn1
[*PE1-bgp-vpn1] import-route direct
[*PE1-bgp-vpn1] auto-frr
[*PE1-bgp-vpn1] advertise l2vpn evpn
[*PE1-bgp-vpn1] quit
[*PE1-bgp] l2vpn-family evpn
[*PE1-bgp-af-evpn] undo policy vpn-target
[*PE1-bgp-af-evpn] peer 2.2.2.2 enable
[*PE1-bgp-af-evpn] peer 2.2.2.2 advertise irb
[*PE1-bgp-af-evpn] peer 2.2.2.2 advertise encap-type vxlan
[*PE1-bgp-af-evpn] peer 4.4.4.4 enable
[*PE1-bgp-af-evpn] peer 4.4.4.4 advertise irb
[*PE1-bgp-af-evpn] peer 4.4.4.4 advertise encap-type vxlan
[*PE1-bgp-af-evpn] quit
[*PE1-bgp] quit
[~CPE] interface nve 1
[*CPE-Nve1] source 4.4.4.4
[*CPE-Nve1] vni 20 head-end peer-list protocol bgp
[*CPE-Nve1] quit
[*CPE] commit
[~PE1] interface nve 1
[*PE1-Nve1] source 3.3.3.3
[*PE1-Nve1] bypass source 1.1.1.1
[*PE1-Nve1] mac-address 00e0-fc12-7890
[*PE1-Nve1] vni 10 head-end peer-list protocol bgp
[*PE1-Nve1] quit
[~CPE] interface vbdif20
[*CPE-Vbdif20] ip binding vpn-instance vpn1
[*CPE-Vbdif20] ip address 10.1.30.1 24
[*CPE-Vbdif20] arp collect host enable
[*CPE-Vbdif20] vxlan anycast-gateway enable
[*CPE-Vbdif20] quit
[*CPE] commit
[~PE1] interface vbdif10
[*PE1-Vbdif10] ip binding vpn-instance vpn1
[*PE1-Vbdif10] ip address 10.1.10.1 24
[*PE1-Vbdif10] arp collect host enable
[*PE1-Vbdif10] vxlan anycast-gateway enable
[*PE1-Vbdif10] mac-address 00e0-fc12-3456
[*PE1-Vbdif10] quit
Run the display vxlan tunnel command on PE1 to check VXLAN tunnel information. The following example uses the command output on PE1.
[~PE1] display vxlan tunnel
Number of vxlan tunnel : 1
Tunnel ID   Source           Destination      State  TyPE    Uptime
-------------------------------------------------------------------
4026531841  1.1.1.1          2.2.2.2          up     dynamic       0033h12m
4026531842  3.3.3.3          4.4.4.4          up     dynamic       0033h12m
PE1 configuration file
#
sysname PE1
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
