---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-163
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [23590, 23723]
sha256: c5b0e0ae343482c035028c76af23aebbf11e45c6718b6241b5b2d4baaca77be1
---

                 # Configure hot standby on the tunnel interface, set the switchback delay to 15s,
                 and specify an explicit path. Enable best-effort path setup.
                 [LSR1-Tunnel1] mpls te backup hot-standby mode revertive wtr 15
                 [LSR1-Tunnel1] mpls te path explicit-path backup-path secondary
                 [LSR1-Tunnel1] mpls te backup ordinary best-effort
                 [LSR1-Tunnel1] quit

                 # After the configuration is complete, check tunnel interface states on LSR1.
                 [LSR1] display mpls te tunnel-interface tunnel
                    Tunnel Name          : Tunnel1
                    Signalled Tunnel Name: -
                    Tunnel State Desc : Primary CR-LSP Up and HotBackup CR-LSP Up
                    Tunnel Attributes :
                    Active LSP         : Primary LSP
                    Traffic Switch      :-
                    Session ID         :1
                    Ingress LSR ID       : 1.1.1.9     Egress LSR ID: 3.3.3.9
                    Admin State           : UP         Oper State : UP
                    Signaling Protocol : RSVP
                    FTid           : 8193
                    Tie-Breaking Policy : None           Metric Type : TE
                    Bfd Cap           : None
                    Reopt            : Disabled       Reopt Freq : -
                    Inter-area Reopt : Disabled
                    Auto BW             : Disabled      Threshold : -
                    Current Collected BW: -              Auto BW Freq : -
                    Min BW              :-           Max BW       :-
                    Offload           : Disabled      Offload Freq : -
                    Low Value           :-           High Value : -


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           393
MPLS Configuration
MPLS Configuration                                                                                4 MPLS TE Configuration

                     Readjust Value         :-
                     Offload Explicit Path Name: -
                     Tunnel Group           : Primary
                     Interfaces Protected: -
                     Excluded IP Address : -
                     Referred LSP Count : 0
                     Primary Tunnel         :-              Pri Tunn Sum : -
                     Backup Tunnel           :-
                     Group Status          : Up              Oam Status : None
                     IPTN InLabel         :-               Tunnel BFD Status : -
                     BackUp LSP Type          : Hot-Standby        BestEffort : Enabled
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
                     PCE Delegate        : No               LSP Control Status : Local control
                     Path Verification : -
                     Entropy Label      : None
                     Associated Tunnel Group ID: -              Associated Tunnel Group Type: -
                     Auto BW Remain Time : -                    Reopt Remain Time      :-
                     Metric Inherit IGP : None
                     Binding Sid       :-                Reverse Binding Sid : -
                     Self-Ping       : Disable            Self-Ping Duration : 1800 sec
                     FRR Attr Source : -                  Is FRR degrade down : -
                     ...

                 The command output shows that the primary and backup CR-LSPs have been
                 established.
                 # Check the hot-standby state of the tunnel on LSR1.
                 [LSR1]display mpls te hot-standby state interface Tunnel 1
                 (s): same path
                 ----------------------------------------------------------------
                 Verbose information about the Tunnel1 hot-standby state
                 ----------------------------------------------------------------
                    Tunnel name             : Tunnel1
                    Session ID            :1
                    Main LSP index          : 0xA002
                    Hot-standby LSP index : 0xA2C1
                    HSB switch result       : main LSP
                    HSB switch reason          :-
                    WTR config time           : 15 s
                    WTR remain time            :-
                    Using overlapped path : no
                    Fast switch status : no

                 # Run the ping command to check the connectivity of the hot-standby CR-LSP on
                 LSR1.
                 [LSR1] ping lsp te tunnel 1 hot-standby
                  LSP PING FEC: TE TUNNEL IPV4 SESSION QUERY Tunnel1 : 100 data bytes, press CTRL_C to break
                    Reply from 3.3.3.9: bytes=100 Sequence=1 time=22 ms
                    Reply from 3.3.3.9: bytes=100 Sequence=2 time=19 ms
                    Reply from 3.3.3.9: bytes=100 Sequence=3 time=15 ms
                    Reply from 3.3.3.9: bytes=100 Sequence=4 time=12 ms
                    Reply from 3.3.3.9: bytes=100 Sequence=5 time=17 ms

                  --- FEC: TE TUNNEL IPV4 SESSION QUERY Tunnel1 ping statistics ---
                    5 packet(s) transmitted


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          394
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration

                     5 packet(s) received
                     0.00% packet loss
                     round-trip min/avg/max = 12/17/22 ms

                 # Run the tracert command on LSR1 to check the hot-standby CR-LSP path.
                 [LSR1] tracert lsp te tunnel1 hot-standby
                  LSP Trace Route FEC: TE TUNNEL IPV4 SESSION QUERY Tunnel1 , press CTRL_C to break.
                  TTL Replier           Time Type      Downstream
                  0                       Ingress 10.1.5.2/[46 ]
                  1    10.1.5.2        6 ms Transit 10.1.3.1/[3 ]
                  2    3.3.3.9         18 ms Egress

                 ----End

Verifying the Configuration
                 # Connect two ports on a tester (such as Port1 and Port2) to LSR1 and LSR3,
                 respectively. Inject MPLS traffic from Port1 to Port2. Ensure that label values are
                 set correctly. When the cable is removed from 10GE1/0/1 on LSR1 or LSR2, check
                 whether fault-trigger service convergence is complete in milliseconds.
                 Run the shutdown command on 10GE1/0/1 of LSR1 to simulate cable removal.
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] shutdown
                 [LSR1-10GE1/0/1] quit

