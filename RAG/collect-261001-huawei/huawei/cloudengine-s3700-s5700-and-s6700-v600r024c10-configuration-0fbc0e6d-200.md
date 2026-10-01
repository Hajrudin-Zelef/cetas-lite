---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-200
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-05-24", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [28990, 29087]
sha256: c2e5f86acaad20f59d1853bdacac677b10f4c72a0171fd5d06908c1be509671a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Verifying the Configuration
                 # Run the display mpls te tunnel name Tunnel1 verbose command on LSR1.
                 The command output shows information about the primary tunnel and its bound
                 bypass tunnel.
                 [LSR1] display mpls te tunnel name Tunnel1 verbose
                    No                : 1
                    Tunnel-Name             : Tunnel1
                    Tunnel Interface Name : Tunnel1
                    TunnelIndex           : -
                    Session ID          : 1         LSP ID         : 6
                    LSR Role            : Ingress
                    Ingress LSR ID        : 1.1.1.1
                    Egress LSR ID         : 3.3.3.3
                    In-Interface        : -
                    Out-Interface         : Vlanif200
                    Sign-Protocol         : RSVP TE     Resv Style     : SE


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        481
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration

                     IncludeAnyAff             : 0x0       ExcludeAnyAff      : 0x0
                     IncludeAllAff           : 0x0
                     ER-Hop Table Index           : 1       AR-Hop Table Index: 8323
                     C-Hop Table Index           : 8193
                     PrevTunnelIndexInSession: -              NextTunnelIndexInSession: -
                     PSB Handle               : -
                     Created Time              : 2024-05-24 10:43:51
                     RSVP LSP Type              : -
                     --------------------------------
                             DS-TE Information
                     --------------------------------
                     Bandwidth Reserved Flag : Unreserved
                     CT0 Bandwidth(Kbit/sec) : 0              CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec) : 0              CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec) : 0              CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec) : 0              CT7 Bandwidth(Kbit/sec): 0
                     Setup-Priority          : 7         Hold-Priority        : 7
                     --------------------------------
                               FRR Information
                     --------------------------------
                     Primary LSP Info
                     Bypass In Use            : Not Used
                     Bypass Tunnel Id           : 40961
                     BypassTunnel              : Tunnel Index[AutoBypassTunnel_1.1.1.1_2.2.2.2_40961], InnerLabel[48093]
                     Bypass LSP ID            : 7         FrrNextHop        : 10.32.1.1
                     ReferAutoBypassHandle : -
                     FrrPrevTunnelTableIndex : -             FrrNextTunnelTableIndex: -
                     Bypass Attribute
                     Setup Priority          : 7        Hold Priority    : 7
                     HopLimit               : 32         Bandwidth        : 0
                     IncludeAnyGroup             : 0        ExcludeAnyGroup : 0
                     IncludeAllGroup            : 0
                     Bypass Unbound Bandwidth Info(Kbit/sec)
                     CT0 Unbound Bandwidth : -                  CT1 Unbound Bandwidth: -
                     CT2 Unbound Bandwidth : -                  CT3 Unbound Bandwidth: -
                     CT4 Unbound Bandwidth : -                  CT5 Unbound Bandwidth: -
                     CT6 Unbound Bandwidth : -                  CT7 Unbound Bandwidth: -
                     --------------------------------
                              BFD Information
                     --------------------------------
                     NextSessionTunnelIndex : -               PrevSessionTunnelIndex: -
                     NextLspId              : -         PrevLspId       : -

                 The command output shows a primary tunnel named Tunnel1 and a bound
                 automatic bypass tunnel named AutoBypassTunnel_1.1.1.1_2.2.2.2_40961 on LSR1.
                 # Run the display mpls te tunnel-interface auto-bypass-tunnel command on
                 LSR1 to check detailed information about the automatic bypass tunnel.
                 [LSR1]display mpls te tunnel-interface auto-bypass-tunnel AutoBypassTunnel_1.1.1.1_2.2.2.2_40961
                    Tunnel Name          : AutoBypassTunnel_1.1.1.1_2.2.2.2_40961
                    Signalled Tunnel Name: -
                    Tunnel State Desc : CR-LSP is Up
                    Tunnel Attributes :
                    Active LSP         : Primary LSP
                    Traffic Switch      :-
                    Session ID         : 40961
                    Ingress LSR ID       : 1.1.1.1        Egress LSR ID: 2.2.2.2
                    Admin State           : UP            Oper State : UP
                    Signaling Protocol : RSVP
                    FTid           : 65
                    Tie-Breaking Policy : None              Metric Type : TE
                    Bfd Cap           : None
                    Reopt            : Disabled          Reopt Freq : -
                    Inter-area Reopt : Disabled
                    Auto BW             : Disabled         Threshold : -
                    Current Collected BW: -                 Auto BW Freq : -
                    Min BW              :-              Max BW       :-
                    Offload           : Disabled         Offload Freq : -
                    Low Value            :-             High Value : -


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                            482
MPLS Configuration
MPLS Configuration                                                                            4 MPLS TE Configuration

