---
id: collect-261001-huawei/huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d-23
title: "ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d.md
source_anchor: ""
source_lines: [4811, 4967]
sha256: 6def6a79f9356e5a0edaaa4db4049cd47c7271e839d6cae3558fc8db7425b88c
---

# ipcompro-training-materials-huawei-hcip-rs-switching-hcnp-iesn-20en-20lab-20guid-ce34e77d

[R2]display mpls ldp session  
 LDP Session(s) in Public Network  
 Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM) 
 A '*' before a session means the session is being deleted. 
 ---------------------------------------------------------------------------- 
 PeerID             Status      LAM  SsnRole  SsnAge      KASent/Rcv 
 ---------------------------------------------------------------------------- 
 2.2.2.2:0          Operational DU   Active   0000:00:11  46/46 
 4.4.4.4:0          Operational DU   Passive  0000:00:10  43/43 
 ---------------------------------------------------------------------------- 
 TOTAL: 2 session(s) Found. 
 
[R3]display mpls ldp session 
 LDP Session(s) in Public Network

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page110 HUAWEI TECHNOLOGIES HC Series 
 
 Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM) 
 A '*' before a session means the session is being deleted. 
 ---------------------------------------------------------------------------- 
 PeerID             Status      LAM  SsnRole  SsnAge      KASent/Rcv 
 ---------------------------------------------------------------------------- 
 3.3.3.3:0          Operational DU   Active   0000:00:11  46/46 
 ---------------------------------------------------------------------------- 
 TOTAL: 1 session(s) Found. 
 
Step 4 Set up LSPs using LDP. 
All LSRs are triggered to establis h LDP LSPs based on the host route, 
which is the default trigger policy. 
Run the display mpls ldp lsp  command on LSRs. All host routes are 
triggered to establish LDP LSPs.
[R1]display mpls ldp lsp 
  LDP LSP Information   
 ---------------------------------------------------------------------------- 
 DestAddress/Mask   In/OutLabel    UpstreamPeer    NextHop         OutInterface  
 ---------------------------------------------------------------------------- 
 2.2.2.2/32         3/NULL         3.3.3.3         127.0.0.1       InLoop0 
*2.2.2.2/32         Liberal/1024                   DS/3.3.3.3  
 3.3.3.3/32         NULL/3         -               10.0.12.2       S1/0/0 
 3.3.3.3/32         1024/3         3.3.3.3         10.0.12.2       S1/0/0 
 4.4.4.4/32         NULL/1025      -               10.0.12.2       S1/0/0 
 4.4.4.4/32         1025/1025      3.3.3.3         10.0.12.2       S1/0/0 
 ---------------------------------------------------------------------------- 
 TOTAL: 5 Normal LSP(s) Found. 
 TOTAL: 1 Liberal LSP(s) Found. 
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

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
HC Series HUAWEI TECHNOLOGIES     
Page111 
 
 ---------------------------------------------------------------------------- 
 2.2.2.2/32         NULL/3         -               10.0.12.1       S1/0/0 
 2.2.2.2/32         1024/3         2.2.2.2         10.0.12.1       S1/0/0 
 2.2.2.2/32         1024/3         4.4.4.4         10.0.12.1       S1/0/0 
*2.2.2.2/32         Liberal/1024                   DS/4.4.4.4  
 3.3.3.3/32         3/NULL         2.2.2.2         127.0.0.1       InLoop0 
 3.3.3.3/32         3/NULL         4.4.4.4         127.0.0.1       InLoop0 
*3.3.3.3/32         Liberal/1024                   DS/2.2.2.2  
*3.3.3.3/32         Liberal/1025                   DS/4.4.4.4  
 4.4.4.4/32         NULL/3         -               10.0.23.3       S2/0/0 
 4.4.4.4/32         1025/3         2.2.2.2         10.0.23.3       S2/0/0 
 4.4.4.4/32         1025/3         4.4.4.4         10.0.23.3       S2/0/0 
*4.4.4.4/32         Liberal/1025                   DS/2.2.2.2  
 ---------------------------------------------------------------------------- 
 TOTAL: 8 Normal LSP(s) Found. 
 TOTAL: 4 Liberal LSP(s) Found. 
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
 3.3.3.3/32         NULL/3         -               10.0.23.2       S2/0/0 
 3.3.3.3/32         1025/3         3.3.3.3         10.0.23.2       S2/0/0 
 4.4.4.4/32         3/NULL         3.3.3.3         127.0.0.1       InLoop0 
*4.4.4.4/32         Liberal/1025                   DS/3.3.3.3  
 ---------------------------------------------------------------------------- 
 TOTAL: 5 Normal LSP(s) Found. 
 TOTAL: 1 Liberal LSP(s) Found. 
 TOTAL: 0 Frr LSP(s) Found. 
 A '*' before an LSP means the LSP is not established  
 A '*' before a Label means the USCB or DSCB is stale  
 A '*' before a UpstreamPeer means the session is in GR state

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
Page112 HUAWEI TECHNOLOGIES HC Series 
 
 A '*' before a DS means the session is in GR state  
 A '*' before a NextHop means the LSP is FRR LSP 
 
In most cases, the default trigger policy is used. The establishment of an 
LDP LSP is triggered in Host mode. 
Change the trigger policy to All on LSRs so that all static routes and IGP 
entries can trigger the establishment of the LDP LSPs.
[R1]mpls  
[R1-mpls]lsp-trigger all 
 
[R2]mpls 
[R2-mpls]lsp-trigger all 
 
[R3]mpls 
[R3-mpls]lsp-trigger all 
 
Run the display mpls ldp lsp  command. Information about the 
established LDP LSPs is displayed. 
[R1]display mpls ldp lsp 
  LDP LSP Information   
 ---------------------------------------------------------------------------- 
 DestAddress/Mask   In/OutLabel    UpstreamPeer    NextHop         OutInterface  
 ---------------------------------------------------------------------------- 
 2.2.2.0/24         3/NULL         3.3.3.3         2.2.2.2         Loop0 
 2.2.2.2/32         3/NULL         3.3.3.3         127.0.0.1       InLoop0 
*2.2.2.2/32         Liberal/1024                   DS/3.3.3.3  
*3.3.3.0/24         Liberal/3                      DS/3.3.3.3  
 3.3.3.3/32         NULL/3         -               10.0.12.2       S1/0/0 
 3.3.3.3/32         1024/3         3.3.3.3         10.0.12.2       S1/0/0 
 4.4.4.4/32         NULL/1025      -               10.0.12.2       S1/0/0 
 4.4.4.4/32         1025/1025      3.3.3.3         10.0.12.2       S1/0/0 
 10.0.1.0/24        3/NULL         3.3.3.3         10.0.1.1        GE0/0/1 
*10.0.1.0/24        Liberal/1026                   DS/3.3.3.3  
 10.0.2.0/24        NULL/1027      -               10.0.12.2       S1/0/0 
 10.0.2.0/24        1027/1027      3.3.3.3         10.0.12.2       S1/0/0 
 10.0.12.0/24       3/NULL         3.3.3.3         10.0.12.1       S1/0/0 
*10.0.12.0/24       Liberal/3                      DS/3.3.3.3  
 10.0.23.0/24       NULL/3         -               10.0.12.2       S1/0/0 
 10.0.23.0/24       1026/3         3.3.3.3         10.0.12.2       S1/0/0 
 ---------------------------------------------------------------------------- 
 TOTAL: 12 Normal LSP(s) Found.

HCDP-IESN  Chapter 3 Implementing MPLS technologies 
 
HC Series HUAWEI TECHNOLOGIES     
Page113 
 
