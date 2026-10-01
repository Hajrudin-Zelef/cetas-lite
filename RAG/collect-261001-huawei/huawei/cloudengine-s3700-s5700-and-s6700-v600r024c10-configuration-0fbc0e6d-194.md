---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-194
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [28081, 28190]
sha256: f4d6025a49ac607aaac63c77dac0749ba83850751e891a3a9958fa89d32e1903
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The command output shows that LSR1 has a primary tunnel named Tunnel1 and
                 an automatic bypass tunnel named AutoBypassTunnel_1.1.1.9_3.3.3.9_32769, and
                 the primary tunnel is bound with the automatic bypass tunnel. The outbound
                 interface of the automatic bypass tunnel is VLANIF600, and LSR2 is excluded from
                 the primary tunnel path, implementing node protection.

                 # Run the display mpls te tunnel-interface auto-bypass-tunnel command on
                 LSR1 to check detailed information about the automatic bypass tunnel.
                 [LSR1]display mpls te tunnel-interface auto-bypass-tunnel
                    Tunnel Name            : AutoBypassTunnel_1.1.1.9_3.3.3.9_32769
                    Signalled Tunnel Name: -
                    Tunnel State Desc : CR-LSP is Up
                    Tunnel Attributes :
                    Active LSP          : Primary LSP
                    Traffic Switch       :-
                    Session ID          : 32769
                    Ingress LSR ID         : 1.1.1.9        Egress LSR ID: 3.3.3.9
                    Admin State             : UP            Oper State : UP
                    Signaling Protocol : RSVP
                    FTid            :2
                    Tie-Breaking Policy : None                Metric Type : TE
                    Bfd Cap            : None
                    Reopt             : Disabled           Reopt Freq : -
                    Inter-area Reopt : Disabled
                    Auto BW               : Disabled         Threshold : -
                    Current Collected BW: -                   Auto BW Freq : -
                    Min BW               :-               Max BW       :-
                    Offload            : Disabled          Offload Freq : -
                    Low Value             :-              High Value : -
                    Readjust Value           :-
                    Offload Explicit Path Name: -
                    Tunnel Group             : Primary
                    Interfaces Protected: Vlanif100
                    Excluded IP Address : 10.1.1.1
                                    10.1.1.2
                                    2.2.2.9
                    Referred LSP Count : 1
                    Primary Tunnel           :-            Pri Tunn Sum : -
                    Backup Tunnel             :-
                    Group Status            : Down            Oam Status : None
                    IPTN InLabel           :-             Tunnel BFD Status : -
                    BackUp LSP Type            : None          BestEffort : Disabled
                    Secondary HopLimit : -
                    BestEffort HopLimit : -
                    Secondary Explicit Path Name: -
                    Secondary Affinity Prop/Mask: 0x0/0x0
                    BestEffort Affinity Prop/Mask: 0x0/0x0
                    IsConfigLspConstraint: -
                    Hot-Standby Revertive Mode: Revertive
                    Hot-Standby Overlap-path: Disabled
                    Hot-Standby Switch State: CLEAR
                    Bit Error Detection: Disabled
                    Bit Error Detection Switch Threshold: -
                    Bit Error Detection Resume Threshold: -
                    Ip-Prefix Name : -
                    P2p-Template Name : -
                    PCE Delegate         : No              LSP Control Status : Local control
                    Path Verification : -
                    Entropy Label       : None
                    Associated Tunnel Group ID: -              Associated Tunnel Group Type: -
                    Auto BW Remain Time : -                    Reopt Remain Time      :-
                    Metric Inherit IGP : None
                    Binding Sid       :-                Reverse Binding Sid : -
                    Self-Ping       : Disable            Self-Ping Duration : 1800 sec


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                       469
MPLS Configuration
MPLS Configuration                                                                          4 MPLS TE Configuration

                     FRR Attr Source : Tunnel Cfg           Is FRR degrade down : -

                     Primary LSP ID        : 1.1.1.9:1100
                     LSP State         : UP               LSP Type     : Primary
                     Configured Attribute Information:
                     Setup Priority     :7                Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Affinity Prop/Mask : 0x0/0x0             Resv Style : SE
                     Actual Attribute Information:
                     Setup Priority     :7                Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Metric Type         : TE             Metric      :2
                     Configured Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 0              CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0              CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0              CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0              CT7 Bandwidth(Kbit/sec): 0
                     Actual Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 0              CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0              CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0              CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0              CT7 Bandwidth(Kbit/sec): 0
                     Explicit Path Name : -                         Hop Limit: -
                     Record Route          : Enabled         Record Label : Enabled
                     Route Pinning         : Disabled
                     FRR Flag          : Disabled
                     IdleTime Remain         :-
                     BFD Status          :-
                     Soft Preemption        : Disabled
                     Reroute Flag         : Enabled
                     Pce Flag         : Normal
                     Path Setup Type        : CSPF
                     Create Modify LSP Reason: -
                     Self-Ping Status : -

