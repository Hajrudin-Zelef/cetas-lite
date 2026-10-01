---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-195
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [28191, 28302]
sha256: 8ecabafcd221c291bc597c873783ad7930b9e0b1f5f70ad971b2e8655f1d4b84
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The command output shows that the automatic bypass tunnel protects the
                 outbound interface VLANIF100 of the primary tunnel and excludes the three
                 addresses (10.1.1.1, 10.1.1.2, and 2.2.2.9) on the primary tunnel path to provide
                 node protection.
                 # Run the display mpls te tunnel verbose command on LSR2. The command
                 output shows information about the primary tunnel and its bound bypass tunnel.
                 [LSR2] display mpls te tunnel verbose
                    No                   : 1
                    Tunnel-Name                  : Tunnel1
                    Tunnel Interface Name : -
                    TunnelIndex              : -
                    Session ID             : 1          LSP ID         : 1099
                    LSR Role               : Transit
                    Ingress LSR ID           : 1.1.1.9
                    Egress LSR ID            : 4.4.4.9
                    In-Interface           : Vlanif100
                    Out-Interface            : Vlanif200
                    Sign-Protocol            : RSVP TE      Resv Style     : SE
                    IncludeAnyAff              : 0x0       ExcludeAnyAff    : 0x0
                    IncludeAllAff           : 0x0
                    ER-Hop Table Index             : -      AR-Hop Table Index: -
                    C-Hop Table Index             : -
                    PrevTunnelIndexInSession: -               NextTunnelIndexInSession: -
                    PSB Handle                : -
                    Created Time               : 2024-04-08 14:31:35
                    RSVP LSP Type               : -
                    --------------------------------
                            DS-TE Information


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    470
MPLS Configuration
MPLS Configuration                                                                            4 MPLS TE Configuration

                     --------------------------------
                     Bandwidth Reserved Flag : Unreserved
                     CT0 Bandwidth(Kbit/sec) : 0              CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec) : 0              CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec) : 0              CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec) : 0              CT7 Bandwidth(Kbit/sec): 0
                     Setup-Priority          : 4         Hold-Priority      : 3
                     --------------------------------
                               FRR Information
                     --------------------------------
                     Primary LSP Info
                     Bypass In Use            : Not Used
                     Bypass Tunnel Id           : 32769
                     BypassTunnel              : Tunnel Index[AutoBypassTunnel_2.2.2.9_3.3.3.9_32769], InnerLabel[18]
                     Bypass LSP ID            : 13        FrrNextHop       : 10.1.5.2
                     ReferAutoBypassHandle : -
                     FrrPrevTunnelTableIndex : -             FrrNextTunnelTableIndex: -
                     Bypass Attribute
                     Setup Priority          : 7        Hold Priority   : 7
                     HopLimit               : 32         Bandwidth       : 0
                     IncludeAnyGroup             : 0       ExcludeAnyGroup : 0
                     IncludeAllGroup            : 0
                     Bypass Unbound Bandwidth Info(Kbit/sec)
                     CT0 Unbound Bandwidth : -                 CT1 Unbound Bandwidth: -
                     CT2 Unbound Bandwidth : -                 CT3 Unbound Bandwidth: -
                     CT4 Unbound Bandwidth : -                 CT5 Unbound Bandwidth: -
                     CT6 Unbound Bandwidth : -                 CT7 Unbound Bandwidth: -
                     --------------------------------
                              BFD Information
                     --------------------------------
                     NextSessionTunnelIndex : -              PrevSessionTunnelIndex: -
                     NextLspId              : -         PrevLspId      : -


                     No                   : 2
                     Tunnel-Name                  : AutoBypassTunnel_2.2.2.9_3.3.3.9_32769
                     Tunnel Interface Name : -
                     TunnelIndex              : -
                     Session ID             : 32769        LSP ID        : 13
                     LSR Role               : Ingress
                     Ingress LSR ID           : 2.2.2.9
                     Egress LSR ID            : 3.3.3.9
                     In-Interface           : -
                     Out-Interface            : Vlanif400
                     Sign-Protocol            : RSVP TE      Resv Style     : SE
                     IncludeAnyAff              : 0x0       ExcludeAnyAff    : 0x0
                     IncludeAllAff           : 0x0
                     ER-Hop Table Index             : -      AR-Hop Table Index: 24578
                     C-Hop Table Index             : 8193
                     PrevTunnelIndexInSession: -               NextTunnelIndexInSession: -
                     PSB Handle                : -
                     Created Time               : 2024-04-08 14:34:45
                     RSVP LSP Type               : -
                     --------------------------------
                             DS-TE Information
                     --------------------------------
                     Bandwidth Reserved Flag : Unreserved
                     CT0 Bandwidth(Kbit/sec) : 0               CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec) : 0               CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec) : 0               CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec) : 0               CT7 Bandwidth(Kbit/sec): 0
                     Setup-Priority          : 7          Hold-Priority      : 7
                     --------------------------------
                               FRR Information
                     --------------------------------
                     Primary LSP Info
                     Bypass In Use             : Not Exists
                     Bypass Tunnel Id            : 0
                     BypassTunnel               : -



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                             471
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration

