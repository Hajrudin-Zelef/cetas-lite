---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-196
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "preemption"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [28303, 28430]
sha256: 6f66cb94b41bb34a1f55fcf9c6be35d9621e482e5dcdeac0a87cf73ac1700468
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                     Bypass LSP ID            : -      FrrNextHop        : -
                     ReferAutoBypassHandle : -
                     FrrPrevTunnelTableIndex : -           FrrNextTunnelTableIndex: -
                     Bypass Attribute(Not configured)
                     Setup Priority          : -      Hold Priority   : -
                     HopLimit               : -       Bandwidth        : -
                     IncludeAnyGroup             : -     ExcludeAnyGroup : -
                     IncludeAllGroup           : -
                     Bypass Unbound Bandwidth Info(Kbit/sec)
                     CT0 Unbound Bandwidth : 0               CT1 Unbound Bandwidth: 0
                     CT2 Unbound Bandwidth : 0               CT3 Unbound Bandwidth: 0
                     CT4 Unbound Bandwidth : 0               CT5 Unbound Bandwidth: 0
                     CT6 Unbound Bandwidth : 0               CT7 Unbound Bandwidth: 0
                     --------------------------------
                              BFD Information
                     --------------------------------
                     NextSessionTunnelIndex : -            PrevSessionTunnelIndex: -
                     NextLspId              : -       PrevLspId      : -

                 The command output shows that LSR2 has a primary tunnel named Tunnel1 and
                 an automatic bypass tunnel named AutoBypassTunnel_2.2.2.9_3.3.3.9_32769, and
                 the primary tunnel is bound with the automatic bypass tunnel. The outbound
                 interface of the automatic bypass tunnel is VLANIF400, which protects the link
                 LSR2 -> LSR3.

                 # Run the display mpls te tunnel-interface auto-bypass-tunnel command on
                 LSR2 to check detailed information about the automatic bypass tunnel.
                 [LSR2]display mpls te tunnel-interface auto-bypass-tunnel
                    Tunnel Name           : AutoBypassTunnel_2.2.2.9_3.3.3.9_32769
                    Signalled Tunnel Name: -
                    Tunnel State Desc : CR-LSP is Up
                    Tunnel Attributes :
                    Active LSP          : Primary LSP
                    Traffic Switch       :-
                    Session ID          : 32769
                    Ingress LSR ID        : 2.2.2.9        Egress LSR ID: 3.3.3.9
                    Admin State            : UP            Oper State : UP
                    Signaling Protocol : RSVP
                    FTid            : 8193
                    Tie-Breaking Policy : None               Metric Type : TE
                    Bfd Cap            : None
                    Reopt             : Disabled          Reopt Freq : -
                    Inter-area Reopt : Disabled
                    Auto BW              : Disabled         Threshold : -
                    Current Collected BW: -                  Auto BW Freq : -
                    Min BW               :-              Max BW       :-
                    Offload            : Disabled         Offload Freq : -
                    Low Value             :-             High Value : -
                    Readjust Value          :-
                    Offload Explicit Path Name: -
                    Tunnel Group            : Primary
                    Interfaces Protected: Vlanif200
                    Excluded IP Address : 10.1.2.1
                                    10.1.2.2
                    Referred LSP Count : 1
                    Primary Tunnel          :-            Pri Tunn Sum : -
                    Backup Tunnel            :-
                    Group Status           : Down            Oam Status : None
                    IPTN InLabel          :-             Tunnel BFD Status : -
                    BackUp LSP Type           : None          BestEffort : Disabled
                    Secondary HopLimit : -
                    BestEffort HopLimit : -
                    Secondary Explicit Path Name: -
                    Secondary Affinity Prop/Mask: 0x0/0x0
                    BestEffort Affinity Prop/Mask: 0x0/0x0
                    IsConfigLspConstraint: -
                    Hot-Standby Revertive Mode: Revertive
                    Hot-Standby Overlap-path: Disabled


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                472
MPLS Configuration
MPLS Configuration                                                                             4 MPLS TE Configuration

                     Hot-Standby Switch State: CLEAR
                     Bit Error Detection: Disabled
                     Bit Error Detection Switch Threshold: -
                     Bit Error Detection Resume Threshold: -
                     Ip-Prefix Name : -
                     P2p-Template Name : -
                     PCE Delegate        : No            LSP Control Status : Local control
                     Path Verification : -
                     Entropy Label      : None
                     Associated Tunnel Group ID: -           Associated Tunnel Group Type: -
                     Auto BW Remain Time : -                 Reopt Remain Time      :-
                     Metric Inherit IGP : None
                     Binding Sid       :-            Reverse Binding Sid : -
                     Self-Ping       : Disable         Self-Ping Duration : 1800 sec
                     FRR Attr Source : Tunnel Cfg          Is FRR degrade down : -

                     Primary LSP ID        : 2.2.2.9:13
                     LSP State         : UP             LSP Type     : Primary
                     Configured Attribute Information:
                     Setup Priority     :7              Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Affinity Prop/Mask : 0x0/0x0           Resv Style : SE
                     Actual Attribute Information:
                     Setup Priority     :7              Hold Priority: 7
                     IncludeAll       : 0x0
                     IncludeAny          : 0x0
                     ExcludeAny           : 0x0
                     Metric Type         : TE           Metric      :2
                     Configured Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 0            CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0            CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0            CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0            CT7 Bandwidth(Kbit/sec): 0
                     Actual Bandwidth Information:
                     CT0 Bandwidth(Kbit/sec): 0            CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec): 0            CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec): 0            CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec): 0            CT7 Bandwidth(Kbit/sec): 0
                     Explicit Path Name : -                       Hop Limit: -
                     Record Route          : Enabled       Record Label : Enabled
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

