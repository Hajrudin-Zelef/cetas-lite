---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-evpn-0017-html-d7f97a5a-2
title: "Configure Router1. The configurations of Router2 and Router3 are similar to that of Router1, and are not mentioned here. When OSPF is used, the 32-bit loopback address of each router must be advertised."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-evpn-0017-html-d7f97a5a.md
source_anchor: ""
source_lines: [34, 165]
sha256: 3b03a4faac9619d9a7e6891e6a472b3c9622c2af69fdbba1c027db0040fef1d3
---

# Configure Router1. The configuration of Router2 is similar to that of Router1, and is not mentioned here.
[Router1] bridge-domain 10
[Router1-bd10] quit
[Router1] interface ethernet 2/0/1.1 mode l2
[Router1-Ethernet2/0/1.1] encapsulation dot1q vid 10
[Router1-Ethernet2/0/1.1] bridge-domain 10
[Router1-Ethernet2/0/1.1] quit
# Establish a BGP EVPN peer relationship on Router1. The configuration of Router2 is similar to that of Router1, and is not mentioned here.
[Router1] bgp 100
[Router1-bgp] router-id 10.1.1.2
[Router1-bgp] peer 10.2.2.2 as-number 100
[Router1-bgp] peer 10.2.2.2 connect-interface LoopBack1
[Router1-bgp] l2vpn-family evpn
[Router1-bgp-af-evpn] peer 10.2.2.2 enable
[Router1-bgp-af-evpn] quit
[Router1-bgp] quit
[Router1] interface nve 1
[Router1-Nve1] source 10.1.1.2
[Router1-Nve1] vni 2010 head-end peer-list 10.2.2.2
[Router1-Nve1] quit
[Router1] ip vpn-instance vpn1
[Router1-vpn-instance-vpn1] ipv4-family
[Router1-vpn-instance-vpn1-af-ipv4] route-distinguisher 100:1
[Router1-vpn-instance-vpn1-af-ipv4] vpn-target 1:1 evpn
[Router1-vpn-instance-vpn1-af-ipv4] quit
[Router1-vpn-instance-vpn1] vxlan vni 5010
[Router1-vpn-instance-vpn1] quit
[Router1] bridge-domain 10
[Router1-bd10] vxlan vni 2010
[Router1-bd10] quit
[Router1] interface vbdif 10
[Router1-Vbdif10] ip binding vpn-instance vpn1
[Router1-Vbdif10] ip address 192.168.10.10 24
[Router1-Vbdif10] quit
[Router1] bgp 100
[Router1-bgp] ipv4-family vpn-instance vpn1
[Router1-bgp-vpn1] import-route direct
[Router1-bgp-vpn1] advertise l2vpn evpn
[Router1-bgp-vpn1] quit
[Router1-bgp] quit
# After the configuration is complete, run the display vxlan tunnel command on Router1 and Router2 to view VXLAN tunnel information. The command output on Router1 is used as an example.
[Router1] display vxlan tunnel
 Tunnel ID       Source              Destination         State     Type         
 ----------------------------------------------------------------------------   
 4               10.1.1.2            10.2.2.2            up        l3 dynamic      
  ----------------------------------------------------------------------------   
 Number of vxlan tunnel : 2  
# After the configuration is complete, run the display vxlan vni command on Router1 and Router2, view the VNI status is up. The command output on Router1 is used as an example.
<Router> display vxlan vni 2010 verbose  
BD ID               :10                                                          
State               :up                                                          
Source              :10.1.1.2     
Source IPv6 Address :-                                                 
UDP Port            :4789                                                        
Peer List           :10.2.2.2     
IPv6 Peer List      :-                    
Router1 configuration file
#
sysname Router1
#                                                                               
ip vpn-instance vpn1                                                            
 ipv4-family                                                                    
  route-distinguisher 100:1                                                    
  vpn-target 1:1 export-extcommunity evpn                                       
  vpn-target 1:1 import-extcommunity evpn                                       
 vxlan vni 5010                                                                 
#
bridge-domain 10                                                                
 vxlan vni 2010
#                                                                               
interface Ethernet2/0/0                                                                                                                         
 ip address 192.168.2.1 255.255.255.0                                           
#                                                                               
interface Ethernet2/0/1.1 mode l2                                               
 encapsulation dot1q vid 10                                                     
 bridge-domain 10
#                                                                               
interface LoopBack1                                                             
 ip address 10.1.1.2 255.255.255.255  
#                                                                               
interface Vbdif10                                                               
 ip binding vpn-instance vpn1                                                   
 ip address 192.168.10.10 255.255.255.0                                         
#                                                                               
interface Nve1                                                                  
 source 10.1.1.2   
 vni 2010 head-end peer-list 10.2.2.2                                                             
#                                                                               
bgp 100  
 router-id 10.1.1.2                                                                       
 peer 10.2.2.2 as-number 100                                                     
 peer 10.2.2.2 connect-interface LoopBack1                                       
 #                                                                              
 ipv4-family unicast                                                            
  undo synchronization                                                          
  peer 10.2.2.2 enable                                                           
 #                                                                              
 l2vpn-family evpn                                                              
  policy vpn-target                                                             
  peer 10.2.2.2 enable                                                           
 #                                                                              
 ipv4-family vpn-instance vpn1                                                  
  import-route direct                                                           
  advertise l2vpn evpn                                                          
#  
ospf 1 router-id 10.1.1.2                                                                          
 area 0.0.0.0                                                                   
  network 10.1.1.2 0.0.0.0                                                       
  network 192.168.2.0 0.0.0.255 
#                                                                               
return 
Router2 configuration file
#
sysname Router2
#                                                                               
ip vpn-instance vpn1                                                            
 ipv4-family                                                                    
  route-distinguisher 100:1                                                    
  vpn-target 1:1 export-extcommunity evpn                                       
  vpn-target 1:1 import-extcommunity evpn                                       
 vxlan vni 5020                                                                 
#
bridge-domain 20                                                                
 vxlan vni 2020
#                                                                               
interface Ethernet2/0/0                                                                                                                        
 ip address 192.168.3.1 255.255.255.0                                           
#                                                                               
interface Ethernet2/0/1.1 mode l2                                               
 encapsulation dot1q vid 20                                                     
 bridge-domain 20
#                                                                               
