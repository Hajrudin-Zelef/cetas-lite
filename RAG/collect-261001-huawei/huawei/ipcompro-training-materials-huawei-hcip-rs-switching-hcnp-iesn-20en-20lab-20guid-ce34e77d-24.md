---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-24
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [4968, 5105]
sha256: 325e3cd57eb35cd921ce1b015d06b31dd3d2621652de62e18762c67f90c12f0d
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

 TOTAL: 4 Liberal LSP(s) Found. 
 TOTAL: 0 Frr LSP(s) Found. 
 A '*' before an LSP means the LSP is not established  
 A '*' before a Label means the USCB or DSCB is stale  
 A '*' before a UpstreamPeer means the session is in GR state  
 A '*' before a DS means the session is in GR state  
 A '*' before a NextHop means the LSP is FRR LSP 
 
[R2]display mpls ldp lsp 
  LDP LSP Information   
 ---------------------------------------------------------------------------- 
 DestAddress/Mask   In/OutLabel    UpstreamPeer    NextHop         OutInterface  
 ---------------------------------------------------------------------------- 
*2.2.2.0/24         Liberal/3                      DS/2.2.2.2  
 2.2.2.2/32         NULL/3         -               10.0.12.1       S1/0/0 
 2.2.2.2/32         1024/3         2.2.2.2         10.0.12.1       S1/0/0 
 2.2.2.2/32         1024/3         4.4.4.4         10.0.12.1       S1/0/0 
*2.2.2.2/32         Liberal/1024                   DS/4.4.4.4  
 3.3.3.0/24         3/NULL         2.2.2.2         3.3.3.3         Loop0 
 3.3.3.0/24         3/NULL         4.4.4.4         3.3.3.3         Loop0 
 3.3.3.3/32         3/NULL         2.2.2.2         127.0.0.1       InLoop0 
 3.3.3.3/32         3/NULL         4.4.4.4         127.0.0.1       InLoop0 
*3.3.3.3/32         Liberal/1024                   DS/2.2.2.2  
*3.3.3.3/32         Liberal/1025                   DS/4.4.4.4  
*4.4.4.0/24         Liberal/3                      DS/4.4.4.4  
 4.4.4.4/32         NULL/3         -               10.0.23.3       S2/0/0 
 4.4.4.4/32         1025/3         2.2.2.2         10.0.23.3       S2/0/0 
 4.4.4.4/32         1025/3         4.4.4.4         10.0.23.3       S2/0/0 
*4.4.4.4/32         Liberal/1025                   DS/2.2.2.2  
 10.0.1.0/24        NULL/3         -               10.0.12.1       S1/0/0 
 10.0.1.0/24        1026/3         2.2.2.2         10.0.12.1       S1/0/0 
 10.0.1.0/24        1026/3         4.4.4.4         10.0.12.1       S1/0/0 
*10.0.1.0/24        Liberal/1026                   DS/4.4.4.4  
 10.0.2.0/24        NULL/3         -               10.0.23.3       S2/0/0 
 10.0.2.0/24        1027/3         2.2.2.2         10.0.23.3       S2/0/0 
 10.0.2.0/24        1027/3         4.4.4.4         10.0.23.3       S2/0/0 
*10.0.2.0/24        Liberal/1027                   DS/2.2.2.2  
 10.0.12.0/24       3/NULL         2.2.2.2         10.0.12.2       S1/0/0 
 10.0.12.0/24       3/NULL         4.4.4.4         10.0.12.2       S1/0/0 
*10.0.12.0/24       Liberal/3                      DS/2.2.2.2  
*10.0.12.0/24       Liberal/1027                   DS/4.4.4.4  
 10.0.23.0/24       3/NULL         2.2.2.2         10.0.23.2       S2/0/0

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page114 HUAWEI TECHNOLOGIES HC Series 
 
 10.0.23.0/24       3/NULL         4.4.4.4         10.0.23.2       S2/0/0 
*10.0.23.0/24       Liberal/1026                   DS/2.2.2.2  
*10.0.23.0/24       Liberal/3                      DS/4.4.4.4  
 ---------------------------------------------------------------------------- 
 TOTAL: 20 Normal LSP(s) Found. 
 TOTAL: 12 Liberal LSP(s) Found. 
 TOTAL: 0 Frr LSP(s) Found. 
 A '*' before an LSP means the LSP is not established  
 A '*' before a Label means the USCB or DSCB is stale  
 A '*' before a UpstreamPeer means the session is in GR state  
 A '*' before a DS means the session is in GR state  
 A '*' before a NextHop means the LSP is FRR LSP 
 
[R3]display mpls ldp lsp 
  LDP LSP Information   
 ---------------------------------------------------------------------------- 
 DestAddress/Mask   In/OutLabel    UpstreamPeer    NextHop         OutInterface  
 ---------------------------------------------------------------------------- 
 2.2.2.2/32         NULL/1024      -               10.0.23.2       S2/0/0 
 2.2.2.2/32         1024/1024      3.3.3.3         10.0.23.2       S2/0/0 
*3.3.3.0/24         Liberal/3                      DS/3.3.3.3  
 3.3.3.3/32         NULL/3         -               10.0.23.2       S2/0/0 
 3.3.3.3/32         1025/3         3.3.3.3         10.0.23.2       S2/0/0 
 4.4.4.0/24         3/NULL         3.3.3.3         4.4.4.4         Loop0 
 4.4.4.4/32         3/NULL         3.3.3.3         127.0.0.1       InLoop0 
*4.4.4.4/32         Liberal/1025                   DS/3.3.3.3  
 10.0.1.0/24        NULL/1026      -               10.0.23.2       S2/0/0 
 10.0.1.0/24        1026/1026      3.3.3.3         10.0.23.2       S2/0/0 
 10.0.2.0/24        3/NULL         3.3.3.3         10.0.2.1        GE0/0/2 
*10.0.2.0/24        Liberal/1027                   DS/3.3.3.3  
 10.0.12.0/24       NULL/3         -               10.0.23.2       S2/0/0 
 10.0.12.0/24       1027/3         3.3.3.3         10.0.23.2       S2/0/0 
 10.0.23.0/24       3/NULL         3.3.3.3         10.0.23.3       S2/0/0 
*10.0.23.0/24       Liberal/3                      DS/3.3.3.3  
 ---------------------------------------------------------------------------- 
 TOTAL: 12 Normal LSP(s) Found. 
 TOTAL: 4 Liberal LSP(s) Found. 
 TOTAL: 0 Frr LSP(s) Found. 
 A '*' before an LSP means the LSP is not established  
 A '*' before a Label means the USCB or DSCB is stale  
 A '*' before a UpstreamPeer means the session is in GR state  
 A '*' before a DS means the session is in GR state  
 A '*' before a NextHop means the LSP is FRR LSP

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
HC Series HUAWEI TECHNOLOGIES     
Page115 
 
 
Step 5 Configure an inbound LDP policy. 
If labels received on R1 are not controlled, R1 will establish a large number 
of LSPs, consuming large memory. 
After an inbound LDP policy is configured, R1 receives label mapping 
messages only from R2 and establishes LSPs to R2, saving resources. 
Run the display mpls lsp command on R1. Information about established 
LSPs is displayed.
[R1]display mpls lsp 
---------------------------------------------------------------------------- 
                 LSP Information: LDP LSP 
---------------------------------------------------------------------------- 
FEC                In/Out Label  In/Out IF                      Vrf Name        
3.3.3.3/32         NULL/3        -/S1/0/0                                       
3.3.3.3/32         1024/3        -/S1/0/0                                       
2.2.2.2/32         3/NULL        -/-                                            
4.4.4.4/32         NULL/1025     -/S1/0/0                                       
4.4.4.4/32         1025/1025     -/S1/0/0                                       
10.0.12.0/24       3/NULL        -/-                                            
10.0.1.0/24        3/NULL        -/-                                            
2.2.2.0/24         3/NULL        -/-                                            
10.0.23.0/24       NULL/3        -/S1/0/0                                       
10.0.23.0/24       1026/3        -/S1/0/0                                       
10.0.2.0/24        NULL/1027     -/S1/0/0                                       
10.0.2.0/24        1027/1027     -/S1/0/0    
 
LSPs on R1 to R2 and R3 are displayed. If the inbound policy is configured 
on R1, only routes to R2 are allowed. 
[R1]ip ip-prefix prefix1 permit 10.0.12.0 24 
[R1]mpls ldp 
[R1-mpls-ldp]inbound peer 3.3.3.3 fec ip-prefix prefix1 
[R1-mpls-ldp]quit 
[R1]display mpls lsp 
---------------------------------------------------------------------------- 
                 LSP Information: LDP LSP 
---------------------------------------------------------------------------- 
FEC                In/Out Label  In/Out IF                      Vrf Name

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page116 HUAWEI TECHNOLOGIES HC Series 
 
