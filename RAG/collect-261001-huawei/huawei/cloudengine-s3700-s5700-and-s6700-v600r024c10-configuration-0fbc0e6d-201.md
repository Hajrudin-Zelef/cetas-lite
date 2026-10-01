---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-201
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [29088, 29235]
sha256: 0cc42de8375b9f5501052594ad67f66a83b86619740785ebbcf8d6088e40cbcb
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                     Readjust Value         :-
                     Offload Explicit Path Name: -
                     Tunnel Group           : Primary
                     Interfaces Protected: Vlanif200
                     Excluded IP Address : 10.21.1.1
                                     10.21.1.2
                     Referred LSP Count : 1
                     Primary Tunnel         :-           Pri Tunn Sum : -
                     Backup Tunnel           :-
                     Group Status          : Down          Oam Status : None
                     IPTN InLabel         :-            Tunnel BFD Status : -
                     BackUp LSP Type          : None         BestEffort : Disabled
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
                     PCE Delegate        : No            LSP Control Status : Local control
                     Path Verification : --
                     Entropy Label      : None
                     Associated Tunnel Group ID: -          Associated Tunnel Group Type: -
                     Auto BW Remain Time : -                 Reopt Remain Time       :-
                     Metric Inherit IGP : None
                     Binding Sid       :-             Reverse Binding Sid : -
                     Self-Ping       : Disable         Self-Ping Duration : 1800 sec
                     FRR Attr Source : Tunnel Cfg          Is FRR degrade down : -

                     Primary LSP ID        : 1.1.1.1:7
                     LSP State         : UP            LSP Type     : Primary
                     Configured Attribute Information:
                     Setup Priority : 7                Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Affinity Prop/Mask : 0x0/0x0           Resv Style : SE
                     Actual Attribute Information:
                     Setup Priority     :7             Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Metric Type         : TE           Metric      : 20
                     Configured Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 0           CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0           CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0           CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0           CT7 Bandwidth(Kbit/sec): 0
                     Actual Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 0           CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0           CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0           CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0           CT7 Bandwidth(Kbit/sec): 0
                     Explicit Path Name : -                  Hop Limit: -
                     Record Route          : Enabled       Record Label : Enabled
                     Route Pinning         : Disabled
                     FRR Flag          : Disabled
                     IdleTime Remain         :-
                     BFD Status          :-
                     Soft Preemption        : Disabled
                     Reroute Flag         : Enabled
                     Pce Flag         : Normal



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      483
MPLS Configuration
MPLS Configuration                                                                 4 MPLS TE Configuration

                     Path Setup Type    : CSPF
                     Create Modify LSP Reason: -
                     Self-Ping Status : -

                 The command output shows that the automatic bypass tunnel protects the
                 outbound interface VLANIF200 of the primary tunnel. Two addresses (10.21.1.1
                 and 10.21.1.2) on the primary tunnel path are excluded, implementing link
                 protection.
                 # Run the display mpls te tunnel path command to check tunnel path
                 information on the local node. The following example uses the command output
                 on LSR1.
                 [LSR1] display mpls te tunnel path
                  Tunnel Interface Name : Tunnel1
                  Lsp ID : 1.1.1.1 :1 :6
                  Hop Information
                   Hop 0 10.21.1.1 Local-Protection available
                   Hop 1 10.21.1.2 Label 48093
                   Hop 2 2.2.2.2 Label 48093
                   Hop 3 10.31.1.1
                   Hop 4 10.31.1.2 Label 3
                   Hop 5 3.3.3.3 Label 3

                  Tunnel Interface Name : AutoBypassTunnel_1.1.1.1_2.2.2.2_40961
                  Lsp ID : 1.1.1.1 :40961 :7
                  Hop Information
                   Hop 0 10.41.1.2
                   Hop 1 10.41.1.1 Label 48090
                   Hop 2 4.4.4.4 Label 48090
                   Hop 3 10.32.1.2
                   Hop 4 10.32.1.1 Label 3
                   Hop 5 2.2.2.2 Label 3

                 The information about the outbound interface of the primary tunnel on LSR1
                 shows that link protection is available to the primary tunnel.

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 200 400
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         mpls te
                         mpls te auto-frr
                         mpls te plr apply node-protection-flag
                         mpls te frr no-node-protection
                         mpls te cspf
                         mpls rsvp-te
                        #
                        explicit-path master
                         next hop 10.21.1.2 index 1
                         next hop 10.31.1.2 index 2
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.21.1.1 255.255.255.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                           484
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

