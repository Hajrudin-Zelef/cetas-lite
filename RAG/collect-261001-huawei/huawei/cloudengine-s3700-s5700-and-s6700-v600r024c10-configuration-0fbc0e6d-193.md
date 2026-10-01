---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-193
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [27954, 28080]
sha256: e853ab922cf50862ee9e0952d87f7291ecfdccca71246617bf74c477784441e2
---

                 # Run the display mpls te tunnel verbose command on LSR1. The command
                 output shows information about the primary tunnel and its bound bypass tunnel.
                 [LSR1] display mpls te tunnel verbose
                    No                   : 1
                    Tunnel-Name                  : Tunnel1
                    Tunnel Interface Name : Tunnel1
                    TunnelIndex              : -
                    Session ID             : 1          LSP ID         : 1099
                    LSR Role               : Ingress
                    Ingress LSR ID           : 1.1.1.9
                    Egress LSR ID            : 4.4.4.9
                    In-Interface           : -
                    Out-Interface            : Vlanif100
                    Sign-Protocol            : RSVP TE      Resv Style     : SE
                    IncludeAnyAff              : 0x0       ExcludeAnyAff    : 0x0
                    IncludeAllAff           : 0x0
                    ER-Hop Table Index             : 1      AR-Hop Table Index: 193
                    C-Hop Table Index             : 961
                    PrevTunnelIndexInSession: -               NextTunnelIndexInSession: -
                    PSB Handle                : -
                    Created Time               : 2024-04-08 14:31:35
                    RSVP LSP Type               : -
                    --------------------------------
                            DS-TE Information
                    --------------------------------
                    Bandwidth Reserved Flag : Unreserved
                    CT0 Bandwidth(Kbit/sec) : 0               CT1 Bandwidth(Kbit/sec): 0
                    CT2 Bandwidth(Kbit/sec) : 0               CT3 Bandwidth(Kbit/sec): 0
                    CT4 Bandwidth(Kbit/sec) : 0               CT5 Bandwidth(Kbit/sec): 0
                    CT6 Bandwidth(Kbit/sec) : 0               CT7 Bandwidth(Kbit/sec): 0
                    Setup-Priority          : 4          Hold-Priority      : 3
                    --------------------------------
                              FRR Information
                    --------------------------------
                    Primary LSP Info
                    Bypass In Use             : Not Used
                    Bypass Tunnel Id            : 32769
                    BypassTunnel               : Tunnel Index[AutoBypassTunnel_1.1.1.9_3.3.3.9_32769], InnerLabel[18]
                    Bypass LSP ID             : 1100       FrrNextHop       : 10.1.7.2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                             467
MPLS Configuration
MPLS Configuration                                                                            4 MPLS TE Configuration

                     ReferAutoBypassHandle : -
                     FrrPrevTunnelTableIndex : -           FrrNextTunnelTableIndex: -
                     Bypass Attribute
                     Setup Priority          : 7      Hold Priority   : 7
                     HopLimit               : 32       Bandwidth       : 0
                     IncludeAnyGroup            : 0      ExcludeAnyGroup : 0
                     IncludeAllGroup           : 0
                     Bypass Unbound Bandwidth Info(Kbit/sec)
                     CT0 Unbound Bandwidth : -               CT1 Unbound Bandwidth: -
                     CT2 Unbound Bandwidth : -               CT3 Unbound Bandwidth: -
                     CT4 Unbound Bandwidth : -               CT5 Unbound Bandwidth: -
                     CT6 Unbound Bandwidth : -               CT7 Unbound Bandwidth: -
                     --------------------------------
                              BFD Information
                     --------------------------------
                     NextSessionTunnelIndex : -            PrevSessionTunnelIndex: -
                     NextLspId              : -       PrevLspId      : -


                     No                   : 2
                     Tunnel-Name                   : AutoBypassTunnel_1.1.1.9_3.3.3.9_32769
                     Tunnel Interface Name : -
                     TunnelIndex               : -
                     Session ID             : 32769         LSP ID        : 1100
                     LSR Role               : Ingress
                     Ingress LSR ID            : 1.1.1.9
                     Egress LSR ID             : 3.3.3.9
                     In-Interface           : -
                     Out-Interface             : Vlanif600
                     Sign-Protocol             : RSVP TE      Resv Style     : SE
                     IncludeAnyAff               : 0x0       ExcludeAnyAff    : 0x0
                     IncludeAllAff           : 0x0
                     ER-Hop Table Index              : -      AR-Hop Table Index: 130
                     C-Hop Table Index              : 962
                     PrevTunnelIndexInSession: -                NextTunnelIndexInSession: -
                     PSB Handle                 : -
                     Created Time                : 2024-04-08 14:31:35
                     RSVP LSP Type                : -
                     --------------------------------
                             DS-TE Information
                     --------------------------------
                     Bandwidth Reserved Flag : Unreserved
                     CT0 Bandwidth(Kbit/sec) : 0                CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec) : 0                CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec) : 0                CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec) : 0                CT7 Bandwidth(Kbit/sec): 0
                     Setup-Priority           : 7          Hold-Priority      : 7
                     --------------------------------
                               FRR Information
                     --------------------------------
                     Primary LSP Info
                     Bypass In Use              : Not Exists
                     Bypass Tunnel Id             : 0
                     BypassTunnel                : -
                     Bypass LSP ID              : -        FrrNextHop       : -
                     ReferAutoBypassHandle : -
                     FrrPrevTunnelTableIndex : -               FrrNextTunnelTableIndex: -
                     Bypass Attribute(Not configured)
                     Setup Priority          : -          Hold Priority  : -
                     HopLimit                : -          Bandwidth       : -
                     IncludeAnyGroup                : -      ExcludeAnyGroup : -
                     IncludeAllGroup              : -
                     Bypass Unbound Bandwidth Info(Kbit/sec)
                     CT0 Unbound Bandwidth : 0                    CT1 Unbound Bandwidth: 0
                     CT2 Unbound Bandwidth : 0                    CT3 Unbound Bandwidth: 0
                     CT4 Unbound Bandwidth : 0                    CT5 Unbound Bandwidth: 0
                     CT6 Unbound Bandwidth : 0                    CT7 Unbound Bandwidth: 0
                     --------------------------------
                              BFD Information



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                     468
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration

                     --------------------------------
                     NextSessionTunnelIndex : -              PrevSessionTunnelIndex: -
                     NextLspId              : -         PrevLspId      : -

