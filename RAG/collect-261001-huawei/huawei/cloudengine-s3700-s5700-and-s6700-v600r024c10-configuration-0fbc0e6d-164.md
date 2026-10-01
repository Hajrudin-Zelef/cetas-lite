---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-164
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [23724, 23843]
sha256: 29831d7d7e7fab084297c0caef413a3045837b17f1d22461ee19c53e0d5aa612
---

                 # Check the hot-standby status of the tunnel on LSR1.
                 [LSR1]display mpls te hot-standby state interface tunnel1
                 (s): same path
                 ----------------------------------------------------------------
                 Verbose information about the Tunnel1 hot-standby state
                 ----------------------------------------------------------------
                    Tunnel name             : Tunnel1
                    Session ID            :1
                    Main LSP index          : 0x0
                    Hot-standby LSP index : 0xA2C1
                    HSB switch result       : hot-standby LSP
                    HSB switch reason          : signal fail
                    WTR config time           : 15 s
                    WTR remain time             :-
                    Using overlapped path : -
                    Fast switch status : no

                 The command output shows that traffic has been switched to the backup CR-LSP.
                 After the removed cable is inserted again, traffic is switched back to the primary
                 CR-LSP about 15s later.
                 # To simulate a backup path fault, remove the cable from 10GE1/0/2 on LSR3 or
                 LSR4 after removing 10GE1/0/1 on LSR1 or LSR2.
                 After the shutdown command is run on 10GE1/0/1 of LSR1, run the shutdown
                 command on 10GE1/0/2 of LSR3 to simulate cable removal.
                 [LSR3] interface 10ge 1/0/2
                 [LSR3-10GE1/0/2] shutdown
                 [LSR3-10GE1/0/2] quit

                 # Wait 5 seconds and check tunnel interface states on LSR1.
                 [LSR1]display mpls te tunnel-interface tunnel 1
                    Tunnel Name       : Tunnel1
                    Signalled Tunnel Name: -
                    Tunnel State Desc : Backup CR-LSP In use and Primary CR-LSP setting Up
                    Tunnel Attributes :


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                 395
MPLS Configuration
MPLS Configuration                                                                                   4 MPLS TE Configuration

                     Active LSP          : BestEffort LSP
                     Traffic Switch       :-
                     Session ID          :1
                     Ingress LSR ID         : 1.1.1.9           Egress LSR ID: 3.3.3.9
                     Admin State            : UP                Oper State : UP
                     Signaling Protocol : RSVP
                     FTid            : 8193
                     Tie-Breaking Policy : None                   Metric Type : TE
                     Bfd Cap            : None
                     Reopt             : Disabled              Reopt Freq : -
                     Inter-area Reopt : Disabled
                     Auto BW               : Disabled            Threshold : -
                     Current Collected BW: -                      Auto BW Freq : -
                     Min BW               :-                 Max BW        :-
                     Offload            : Disabled             Offload Freq : -
                     Low Value             :-                 High Value : -
                     Readjust Value          :-
                     Offload Explicit Path Name: -
                     Tunnel Group            : Primary
                     Interfaces Protected: -
                     Excluded IP Address : -
                     Referred LSP Count : 0
                     Primary Tunnel           :-               Pri Tunn Sum : -
                     Backup Tunnel            :-
                     Group Status           : Up                Oam Status : None
                     IPTN InLabel           :-                Tunnel BFD Status : -
                     BackUp LSP Type           : BestEffort         BestEffort : Enabled
                     Secondary HopLimit : -
                     BestEffort HopLimit : -
                     Secondary Explicit Path Name: backup-path
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
                     PCE Delegate         : No                 LSP Control Status : Local control
                     Path Verification : -
                     Entropy Label       : None
                     Associated Tunnel Group ID: -                 Associated Tunnel Group Type: -
                     Auto BW Remain Time : -                       Reopt Remain Time      :-
                     Metric Inherit IGP : None
                     Binding Sid       :-                   Reverse Binding Sid : -
                     Self-Ping       : Disable               Self-Ping Duration : 1800 sec
                     FRR Attr Source : -                     Is FRR degrade down : -
                     ...

                 The command output shows that the tunnel interface is up, a best-effort path is
                 established, and traffic is switched to the best-effort path.
                 # Check local tunnel path information on LSR1.
                 [LSR1] display mpls te tunnel path
                  Tunnel Interface Name : Tunnel1
                  Lsp ID : 1.1.1.9 :1 :2821
                  Hop Information
                   Hop 0 10.1.5.1
                   Hop 1 10.1.5.2
                   Hop 2 4.4.4.9
                   Hop 3 10.1.4.2
                   Hop 4 10.1.4.1
                   Hop 5 2.2.2.9
                   Hop 6 10.1.2.1
                   Hop 7 10.1.2.2
                   Hop 8 3.3.3.9


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            396
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


                 The command output shows that the current tunnel uses the best-effort path.


