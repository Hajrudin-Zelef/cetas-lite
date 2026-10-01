---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-25
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [5106, 5222]
sha256: 20aef78a0d6e6feb520d110512f03066d524bc2d44e402eeffd854fd4f2c2a69
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

2.2.2.2/32         3/NULL        -/-                                            
10.0.12.0/24       3/NULL        -/-                                            
10.0.1.0/24        3/NULL        -/-                                            
2.2.2.0/24         3/NULL        -/-      
 
Additional Exercises: Analyzing and Verifying 
In which way can only R1 receive label mapping messages from R3? 
Final Configurations
<R1>display current-configuration
[V200R001C00SPC200] 
# 
 sysname R1 
# 
mpls lsr-id 2.2.2.2 
 mpls 
  lsp-trigger all  
# 
mpls ldp 
 inbound peer 3.3.3.3 fec ip-prefix prefix1 
# 
interface Serial1/0/0 
 link-protocol ppp 
 ip address 10.0.12.1 255.255.255.0  
 mpls 
 mpls ldp 
# 
interface GigabitEthernet0/0/1             
 ip address 10.0.1.1 255.255.255.0  
# 
interface LoopBack0 
 ip address 2.2.2.2 255.255.255.0  
# 
ospf 1 router-id 2.2.2.2  
 area 0.0.0.0  
  network 10.0.1.0 0.0.0.255  
  network 10.0.12.0 0.0.0.255  
  network 2.2.2.0 0.0.0.255  
#

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
HC Series HUAWEI TECHNOLOGIES     
Page117 
 
 ip ip-prefix prefix1 index 10 permit 10.0.12.0 24 
# 
return 
 
[R2]display current-configuration  
[V200R001C00SPC200] 
# 
 sysname R2 
# 
mpls lsr-id 3.3.3.3 
 mpls 
  lsp-trigger all  
# 
mpls ldp 
# 
interface Serial1/0/0 
 link-protocol ppp 
 ip address 10.0.12.2 255.255.255.0  
 mpls 
 mpls ldp 
# 
interface Serial2/0/0 
 link-protocol ppp 
 ip address 10.0.23.2 255.255.255.0  
 mpls 
 mpls ldp 
# 
interface LoopBack0 
 ip address 3.3.3.3 255.255.255.0  
# 
ospf 1 router-id 3.3.3.3  
 area 0.0.0.0  
  network 10.0.12.0 0.0.0.255  
  network 10.0.23.0 0.0.0.255  
  network 3.3.3.0 0.0.0.255  
# 
return 
 
[R3]display current-configuration  
[V200R001C00SPC200] 
# 
 sysname R3

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page118 HUAWEI TECHNOLOGIES HC Series 
 
# 
mpls lsr-id 4.4.4.4 
 mpls 
  lsp-trigger all  
# 
mpls ldp 
# 
interface Serial2/0/0 
 link-protocol ppp 
 ip address 10.0.23.3 255.255.255.0  
 mpls 
 mpls ldp 
# 
interface GigabitEthernet0/0/2 
 ip address 10.0.2.1 255.255.255.0  
# 
interface LoopBack0 
 ip address 4.4.4.4 255.255.255.0  
# 
ospf 1 router-id 4.4.4.4  
 area 0.0.0.0  
  network 10.0.2.0 0.0.0.255  
  network 10.0.23.0 0.0.0.255  
  network 4.4.4.0 0.0.0.255  
# 
Return
