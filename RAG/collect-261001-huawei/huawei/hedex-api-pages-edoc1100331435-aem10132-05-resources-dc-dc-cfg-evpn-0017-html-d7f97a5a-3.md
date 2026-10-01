---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-evpn-0017-html-d7f97a5a-3
title: "Configure Router1. The configurations of Router2 and Router3 are similar to that of Router1, and are not mentioned here. When OSPF is used, the 32-bit loopback address of each router must be advertised."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-evpn-0017-html-d7f97a5a.md
source_anchor: ""
source_lines: [166, 219]
sha256: f157fef9e19123537d32fe0965d6c2d57458745aaa8727c29a7e2d5ab481d76c
---

# Configure Router1. The configurations of Router2 and Router3 are similar to that of Router1, and are not mentioned here. When OSPF is used, the 32-bit loopback address of each router must be advertised.

interface LoopBack1                                                             
 ip address 10.2.2.2 255.255.255.255  
#                                                                               
interface Vbdif20                                                               
 ip binding vpn-instance vpn1                                                   
 ip address 192.168.20.10 255.255.255.0                                         
#                                                                               
interface Nve1                                                                  
 source 10.2.2.2  
 vni 2020 head-end peer-list 10.1.1.2                                                                
#                                                                               
bgp 100    
 router-id 10.2.2.2                                                                     
 peer 10.1.1.2 as-number 100                                                     
 peer 10.1.1.2 connect-interface LoopBack1                                       
 #                                                                              
 ipv4-family unicast                                                            
  undo synchronization                                                          
  peer 10.1.1.2 enable                                                           
 #                                                                              
 l2vpn-family evpn                                                              
  policy vpn-target                                                             
  peer 10.1.1.2 enable                                                           
 #                                                                              
 ipv4-family vpn-instance vpn1                                                  
  import-route direct                                                           
  advertise l2vpn evpn                                                          
#  
ospf 1 router-id 10.2.2.2                                                                           
 area 0.0.0.0                                                                   
  network 10.2.2.2 0.0.0.0                                                       
  network 192.168.3.0 0.0.0.255 
#                                                                               
return 
Router3 configuration file
#
sysname Router3
#                                                                               
interface Ethernet2/0/1                                                                                              
 ip address 192.168.2.2 255.255.255.0                                           
#                                                                               
interface Ethernet2/0/2                                                                                           
 ip address 192.168.3.2 255.255.255.0                                           
#                                                                               
interface LoopBack1                                                             
 ip address 10.3.3.2 255.255.255.255  
#  
ospf 1 router-id 10.3.3.2                                                                         
 area 0.0.0.0                                                                   
  network 10.3.3.2 0.0.0.0                                                       
  network 192.168.2.0 0.0.0.255 
  network 192.168.3.0 0.0.0.255 
#                                                                               
return
