---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-143
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [20525, 20631]
sha256: 0cf4f2d288c8243d63a2e0c75003b12e07de1ab8cde48ee750d5e5c00c0e6526
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            343
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration


                 Tunnel1 is 0x10101, this tunnel selects the link with the second and fourth bits of
                 the administrative group attribute being 0 and at least one of the first and fifth
                 bits being 1. According to the preceding rules, if the value of the administrative
                 group attribute is 0x10001, 0x10000, 0x00001, 0x10101, 0x10100, or 0x00101, the
                 value meets requirements. Finally, Tunnel1 selects VLANIF100 of LSR1 (the
                 administrative group attribute value is 0x10001) and VLANIF200 of LSR2 (the
                 administrative group attribute value is 0x10101).
                 # After the configuration is complete, check tunnel states on LSR1.
                 [LSR1] display mpls te tunnel-interface
                    Tunnel Name            : Tunnel1
                    Signalled Tunnel Name: -
                    Tunnel State Desc : CR-LSP is Up
                    Tunnel Attributes :
                    Active LSP          : Primary LSP
                    Traffic Switch       :-
                    Session ID          :1
                    Ingress LSR ID         : 1.1.1.1        Egress LSR ID: 3.3.3.3
                    Admin State             : UP            Oper State : UP
                    Signaling Protocol : RSVP
                    FTid            :1
                    Tie-Breaking Policy : None                Metric Type : None
                    Bfd Cap            : None
                    Reopt             : Disabled           Reopt Freq : -
                    Inter-area Reopt : Disabled
                    Auto BW               : Disabled         Threshold : 0 percent
                    Current Collected BW: 0 kbps                Auto BW Freq : 0
                    Min BW               : 0 kbps           Max BW       : 0 kbps
                    Offload            : Disabled          Offload Freq : -
                    Low Value             :-              High Value : -
                    Readjust Value           :-
                    Offload Explicit Path Name:
                    Tunnel Group             :-
                    Interfaces Protected: -
                    Excluded IP Address : -
                    Referred LSP Count : 0
                    Primary Tunnel           :-            Pri Tunn Sum : -
                    Backup Tunnel             :-
                    Group Status            : Up            Oam Status : -
                    IPTN InLabel           :-             Tunnel BFD Status : -
                    BackUp LSP Type            : None          BestEffort : Enabled
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
                    PCE Delegate         : No         LSP Control Status : Local control
                    Path Verification : --
                    Entropy Label       : None
                    Associated Tunnel Group ID: -                Associated Tunnel Group Type: -
                    Auto BW Remain Time : 200 s Reopt Remain Time : 100 s
                    Metric Inherit IGP : None
                    Binding Sid       :-                Reverse Binding Sid : -
                    Self-Ping       : Disable            Self-Ping Duration : 1800 sec
                    FRR Attr Source : -                  Is FRR degrade down : No

                     Primary LSP ID    : 1.1.1.1:19
                     LSP State      : UP               LSP Type     : Primary


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                           344
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration

                     Setup Priority     :7            Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Affinity Prop/Mask : 0x0/0x0         Resv Style : SE
                     Configured Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 10000        CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0         CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0         CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0         CT7 Bandwidth(Kbit/sec): 0
                     Actual Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 10000        CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0         CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0         CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0         CT7 Bandwidth(Kbit/sec): 0
                     Explicit Path Name : -                     Hop Limit: -
                     Record Route          : Disabled    Record Label : Disabled
                     Route Pinning         : Disabled
                     FRR Flag          : Disabled
                     IdleTime Remain         :-
                     BFD Status         :-
                     Soft Preemption        : Enabled
                     Reroute Flag         : Disabled
                     Pce Flag     : Normal
                     Path Setup Type        : CSPF
                     Create Modify LSP Reason: -
                     Self-Ping Status : -

