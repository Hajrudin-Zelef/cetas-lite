---
id: collect-261001-general-networking/general-networking/questions-1949-can-adding-vlans-to-a-vpc-be-done-hot-295ecc11
title: "questions-1949-can-adding-vlans-to-a-vpc-be-done-hot-295ecc11"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/questions-1949-can-adding-vlans-to-a-vpc-be-done-hot-295ecc11.md
source_anchor: ""
source_lines: [1, 34]
sha256: a149af05ca59860e9345b772b6f32d7693c6bc1a9047b140fd0a152403d61bd2
---

# questions-1949-can-adding-vlans-to-a-vpc-be-done-hot-295ecc11

While it is certainly a good idea that a VPC has the same allowed VLAN list on all sides, it is not a requirement. Take a look at the output of "show vpc consistency-parameters global". 
switch# show vpc consistency-parameters global 
Legend:
        Type 1 : vPC will be suspended in case of mismatch
Name                        Type  Local Value            Peer Value             
-------------               ----  ---------------------- -----------------------
QoS                         2     ([], [3], [], [], [],  ([], [3], [], [], [], 
                                  [])                    [])                   
Network QoS (MTU)           2     (1538, 2240, 0, 0, 0,  (1538, 2240, 0, 0, 0, 
                                  0)                     0)                    
Network Qos (Pause)         2     (T, F, F, F, F, F)     (T, F, F, F, F, F)    
Input Queuing (Bandwidth)   2     (50, 50, 0, 0, 0, 0)   (50, 50, 0, 0, 0, 0)  
Input Queuing (Absolute     2     (F, F, F, F, F, F)     (F, F, F, F, F, F)    
Priority)                                                                      
Output Queuing (Bandwidth)  2     (50, 50, 0, 0, 0, 0)   (50, 50, 0, 0, 0, 0)  
Output Queuing (Absolute    2     (F, F, F, F, F, F)     (F, F, F, F, F, F)    
Priority)                                                                      
STP Mode                    1     MST                    MST                   
STP Disabled                1     None                   None                  
STP MST Region Name         1     ""                     ""                    
STP MST Region Revision     1     0                      0                     
STP MST Region Instance to  1                                                  
 VLAN Mapping                                                                  
STP Loopguard               1     Disabled               Disabled              
STP Bridge Assurance        1     Enabled                Enabled               
STP Port Type, Edge         1     Normal, Enabled,       Normal, Enabled,      
BPDUFilter, Edge BPDUGuard        Disabled               Disabled              
STP MST Simulate PVST       1     Enabled                Enabled               
Allowed VLANs               -     1,10,100-101,200-201   1,10,100-101,200-201,2
                                                         000                   
Local suspended VLANs       -     -                      -     
Type 1 parameters absolutely MUST match, else the vPC will be suspended. "Allowed VLANs" is neither a Type 1 or Type 2 parameter.
You should be able to modify the allowed VLAN list without problems, but obviously try to keep things consistent regardless, otherwise you'll still have issues, even though the vPC doesn't get suspended.
EDIT: You probably already know this but just to be sure - include the peer-link in your VLAN changes, as the VLAN (not the entire vPC) will show up as "suspended" on all vPCs until you do.
