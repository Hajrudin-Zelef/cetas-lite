---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-vxlan-configuration-guide-b-da751f1e
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-vxlan-configuration-guide-b-da751f1e"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-vxlan-configuration-guide-b-da751f1e.md
source_anchor: ""
source_lines: [1, 101]
sha256: b7c2cc85c05fd004ad7b90cb70296caaed76acfdbfe8d71627cdf01f969e0e41
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-7-x-vxlan-configuration-guide-b-da751f1e

Show commands 
                                 			 
tor1# sh nve vni 
Codes: CP - Control Plane        DP - Data Plane          
       UC - Unconfigured         SA - Suppress ARP
       
Interface VNI      Multicast-group   State Mode Type [BD/VRF]      Flags
--------- -------- ----------------- ----- ---- ------------------ -----
nve1      10001    225.1.1.1         Up    DP   L2 [101]                
nve1      10002    225.1.1.1         Up    DP   L2 [102]                
nve1      10003    225.1.1.1         Up    DP   L2 [103]               
tor1# sh nve peers 
Interface Peer-IP          State LearnType Uptime   Router-Mac       
--------- ---------------  ----- --------- -------- -----------------
nve1      10.200.1.1       Up    DP        00:07:23 n/a              
nve1      10.200.1.2       Up    DP        00:07:18 n/a              
nve1      102.102.102.102  Up    DP        00:07:23 n/a      
tor1# sh ip mroute 225.1.1.1
IP Multicast Routing Table for VRF "default"
(*, 225.1.1.1/32), uptime: 00:07:41, ip pim nve static igmp 
  Incoming interface: Ethernet2/1, RPF nbr: 10.1.5.2
  Outgoing interface list: (count: 3)
    Vlan2, uptime: 00:07:23, igmp
    Vlan2000, uptime: 00:07:31, static
    nve1, uptime: 00:07:41, nve
(10.200.1.1/32, 225.1.1.1/32), uptime: 00:07:40, ip mrib pim nve 
  Incoming interface: Vlan2, RPF nbr: 10.200.1.1
  Outgoing interface list: (count: 3)
    Vlan2, uptime: 00:07:23, mrib, (RPF)
    Vlan2000, uptime: 00:07:31, mrib
    nve1, uptime: 00:07:40, nve
(10.200.1.2/32, 225.1.1.1/32), uptime: 00:07:41, ip mrib pim nve 
  Incoming interface: Vlan2, RPF nbr: 10.200.1.2
  Outgoing interface list: (count: 3)
    Vlan2, uptime: 00:07:23, mrib, (RPF)
    Vlan2000, uptime: 00:07:31, mrib
    nve1, uptime: 00:07:41, nve
(99.99.99.99/32, 225.1.1.1/32), uptime: 00:07:41, ip mrib pim nve 
  Incoming interface: loopback0, RPF nbr: 99.99.99.99
  Outgoing interface list: (count: 3)
    Vlan2, uptime: 00:07:23, mrib
    Vlan2000, uptime: 00:07:31, mrib
    Ethernet2/5, uptime: 00:07:39, pim
(102.102.102.102/32, 225.1.1.1/32), uptime: 00:07:40, ip mrib pim nve 
  Incoming interface: Ethernet2/1, RPF nbr: 10.1.5.2
  Outgoing interface list: (count: 1)
    nve1, uptime: 00:07:40, nve
tor1# sh vpc
Legend:
                 - local vPC is down, forwarding via vPC peer-link
vPC domain id                     : 1   
Peer status                       : peer adjacency formed ok      
vPC keep-alive status             : peer is alive                 
Configuration consistency status  : success 
Per-vlan consistency status       : success                       
Type-2 consistency status         : success 
vPC role                          : secondary, operational primary
Number of vPCs configured         : 4   
Peer Gateway                      : Enabled
Dual-active excluded VLANs        : -
Graceful Consistency Check        : Enabled
Auto-recovery status              : Disabled
Delay-restore status              : Timer is off.(timeout = 180s)
Delay-restore SVI status          : Timer is off.(timeout = 10s)
vPC Peer-link status
---------------------------------------------------------------------
id   Port   Status Active vlans    
--   ----   ------ --------------------------------------------------
1    Po100  up     1-4,101-103,2000                                       
vPC status
----------------------------------------------------------------------
id   Port   Status Consistency Reason                     Active vlans
--   ----   ------ ----------- ------                     ------------
1    Po1    up     success     success                    2               
                                                                          
2    Po2    up     success     success                    2  
tor1# sh vpc consistency-parameters global 
    Legend:
        Type 1 : vPC will be suspended in case of mismatch
Name                        Type  Local Value            Peer Value             
-------------               ----  ---------------------- -----------------------
Vlan to Vn-segment Map      1     3 Relevant Map(s)      3 Relevant Map(s)     
STP Mode                    1     Rapid-PVST             Rapid-PVST            
STP Disabled                1     None                   None                  
STP MST Region Name         1     ""                     ""                    
STP MST Region Revision     1     0                      0                     
STP MST Region Instance to  1                                                  
 VLAN Mapping                                                                  
STP Loopguard               1     Disabled               Disabled              
STP Bridge Assurance        1     Enabled                Enabled               
STP Port Type, Edge         1     Normal, Disabled,      Normal, Disabled,     
BPDUFilter, Edge BPDUGuard        Disabled               Disabled              
STP MST Simulate PVST       1     Enabled                Enabled               
Nve Oper State, Secondary   1     Up, 99.99.99.99, DP    Up, 99.99.99.99, DP   
IP, Host Reach Mode                                                            
Nve Vni Configuration       1     10001-10003            10001-10003           
Interface-vlan admin up     2     2,2000                 2,2000                
Interface-vlan routing      2     1-4,2000               1-4,2000              
capability                                                                     
Allowed VLANs               -     1-4,101-103,2000       1-4,101-103,2000      
Local suspended VLANs       -     -
