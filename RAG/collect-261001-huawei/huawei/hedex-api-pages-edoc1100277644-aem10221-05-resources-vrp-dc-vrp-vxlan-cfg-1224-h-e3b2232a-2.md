---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100277644-aem10221-05-resources-vrp-dc-vrp-vxlan-cfg-1224-h-e3b2232a-2
title: "hedex-api-pages-edoc1100277644-aem10221-05-resources-vrp-dc-vrp-vxlan-cfg-1224-h-e3b2232a"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100277644-aem10221-05-resources-vrp-dc-vrp-vxlan-cfg-1224-h-e3b2232a.md
source_anchor: ""
source_lines: [243, 301]
sha256: 69eb7b1c67f546e4674d117f08fe5fcce7ba2219788738ecf44e38f2e77d20f6
---

#               
ospf 1          
 area 0.0.0.0   
  network 3.3.3.3 0.0.0.0
  network 192.168.30.0 0.0.0.255
#
ip route-static 2.2.2.2 255.255.255.255 192.168.50.1
#               
return
Leaf4 configuration file
#
sysname Leaf4
#
evpn vpn-instance evrf1 bd-mode
 route-distinguisher 10:1
 vpn-target 11:1 export-extcommunity
 vpn-target 11:1 import-extcommunity
#
bridge-domain 10
 vxlan vni 20 split-horizon-mode
 evpn binding vpn-instance evrf1
#               
interface GE0/1/0
 undo shutdown  
 ip address 192.168.40.2 255.255.255.0
#                              
interface GE0/1/8
 undo shutdown       
#               
interface GE0/1/8.1 mode l2
 encapsulation dot1q vid 10
 rewrite pop single
 bridge-domain 10
#               
interface LoopBack1
 ip address 4.4.4.4 255.255.255.255
#               
interface Nve1  
 source 4.4.4.4 
 vni 20 head-end peer-list protocol bgp
#               
bgp 200
 peer 3.3.3.3 as-number 200
 peer 3.3.3.3 connect-interface LoopBack1
 #
 ipv4-family unicast
  peer 3.3.3.3 enable
 #              
 l2vpn-family evpn
  undo policy vpn-target
  peer 3.3.3.3 enable
  peer 3.3.3.3 advertise encap-type vxlan
#               
ospf 1          
 area 0.0.0.0   
  network 4.4.4.4 0.0.0.0
  network 192.168.40.0 0.0.0.255
#               
return
