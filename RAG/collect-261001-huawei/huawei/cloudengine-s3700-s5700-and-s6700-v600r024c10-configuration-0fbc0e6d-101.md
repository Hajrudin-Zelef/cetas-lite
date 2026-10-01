---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-101
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [14172, 14322]
sha256: 26bc21e464e4e4c0debef03e859b7fbfacd944c673b3212cdd60c0b49b7d4649
---

                 # Check detailed tunnel information.
                 [LSR1] display mpls te tunnel-interface tunnel1
                    Tunnel Name            : Tunnel1
                    Signalled Tunnel Name: -
                    Tunnel State Desc : CR-LSP is Up
                    Tunnel Attributes :
                    Active LSP          : Primary LSP
                    Traffic Switch       :-
                    Session ID          :1
                    Ingress LSR ID         : 1.1.1.9       Egress LSR ID: 4.4.4.9
                    Admin State             : UP           Oper State : UP
                    Signaling Protocol : RSVP
                    FTid            : 16385
                    Tie-Breaking Policy : None               Metric Type : TE
                    Bfd Cap            : None
                    Reopt             : Disabled          Reopt Freq : -
                    Inter-area Reopt : Disabled
                    Auto BW               : Disabled        Threshold : -
                    Current Collected BW: -                  Auto BW Freq : -
                    Min BW               :-             Max BW        :-
                    Offload            : Disabled         Offload Freq : -
                    Low Value             :-             High Value : -
                    Readjust Value           :-
                    Offload Explicit Path Name: -
                    Tunnel Group             : Primary
                    Interfaces Protected: -
                    Excluded IP Address : -
                    Referred LSP Count : 0
                    Primary Tunnel           :-           Pri Tunn Sum : -
                    Backup Tunnel             :-
                    Group Status            : Up           Oam Status : None
                    IPTN InLabel           :-            Tunnel BFD Status : -
                    BackUp LSP Type            : None         BestEffort : Disabled
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
                    PCE Delegate         : No             LSP Control Status : Local control
                    Path Verification : -
                    Entropy Label       : None
                    Associated Tunnel Group ID: -             Associated Tunnel Group Type: -
                    Auto BW Remain Time : -                   Reopt Remain Time      :-
                    Metric Inherit IGP : None
                    Binding Sid       :-               Reverse Binding Sid : -
                    Self-Ping       : Disable           Self-Ping Duration : 1800 sec
                    FRR Attr Source : -                 Is FRR degrade down : -

                     Primary LSP ID       : 1.1.1.9:40
                     LSP State        : UP             LSP Type     : Primary
                     Configured Attribute Information:
                     Setup Priority     :7             Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny         : 0x0
                     ExcludeAny          : 0x0
                     Affinity Prop/Mask : 0x0/0x0          Resv Style : SE
                     Actual Attribute Information:
                     Setup Priority     :7             Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny         : 0x0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        241
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration

                      ExcludeAny          : 0x0
                      Metric Type        : TE          Metric    : 30
                      Configured Bandwidth Information:
                      CT0 Bandwidth(Kbit/sec): 20000       CT1 Bandwidth(Kbit/sec): 0
                      CT2 Bandwidth(Kbit/sec): 0        CT3 Bandwidth(Kbit/sec): 0
                      CT4 Bandwidth(Kbit/sec): 0        CT5 Bandwidth(Kbit/sec): 0
                      CT6 Bandwidth(Kbit/sec): 0        CT7 Bandwidth(Kbit/sec): 0
                      Actual Bandwidth Information:
                      CT0 Bandwidth(Kbit/sec): 20000       CT1 Bandwidth(Kbit/sec): 0
                      CT2 Bandwidth(Kbit/sec): 0        CT3 Bandwidth(Kbit/sec): 0
                      CT4 Bandwidth(Kbit/sec): 0        CT5 Bandwidth(Kbit/sec): 0
                      CT6 Bandwidth(Kbit/sec): 0        CT7 Bandwidth(Kbit/sec): 0
                      Explicit Path Name : -                   Hop Limit: -
                      Record Route         : Enabled     Record Label : Disabled
                      Route Pinning        : Disabled
                      FRR Flag         : Disabled
                      IdleTime Remain        :-
                      BFD Status        :-
                      Soft Preemption       : Disabled
                      Reroute Flag        : Enabled
                      Pce Flag        : Normal
                      Path Setup Type       : CSPF
                      Create Modify LSP Reason: -
                      Self-Ping Status : -

                 # Check link information in the TEDB on LSR1.
                 [LSR1] display mpls te cspf tedb all
                 Current Total Node Number: 4
                 Current Total Link Number: 6
                 Current Total SRLG Number: 0

                 Id      Router-Id      IGP       Process-Id    Area       Link-Count
                 1       1.1.1.9      ISIS    1             Level-2    1
                 2       2.2.2.9      ISIS    1             Level-2    2
                 3       3.3.3.9      ISIS    1             Level-2    2
                 4       4.4.4.9      ISIS    1             Level-2    1


Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0001.00
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                             242
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

