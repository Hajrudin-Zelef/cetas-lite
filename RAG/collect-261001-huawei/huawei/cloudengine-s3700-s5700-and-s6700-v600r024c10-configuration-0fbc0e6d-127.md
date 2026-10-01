---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-127
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17947, 18069]
sha256: d24599836463ec9ae00b180b6c8bf5ef8001923d74c09bad4b73f18dcdc53952
---

                 # Run the display mpls te tunnel name Tunnel1 verbose command on LSR2.
                 The command output shows that the bypass tunnel has been used.
                 [LSR2] display mpls te tunnel name Tunnel1 verbose
                    No                   : 1
                    Tunnel-Name                  : Tunnel1
                    Tunnel Interface Name : -
                    TunnelIndex              : -
                    Session ID             : 1          LSP ID          : 690
                    LSR Role               : Transit
                    Ingress LSR ID           : 1.1.1.9
                    Egress LSR ID            : 4.4.4.9
                    In-Interface           : Vlanif100
                    Out-Interface            : Vlanif200
                    Sign-Protocol            : RSVP TE       Resv Style      : SE
                    IncludeAnyAff              : 0x0       ExcludeAnyAff      : 0x0
                    IncludeAllAff           : 0x0
                    ER-Hop Table Index             : -      AR-Hop Table Index: -
                    C-Hop Table Index             : -
                    PrevTunnelIndexInSession: -               NextTunnelIndexInSession: -
                    PSB Handle                : -
                    Created Time               : 2024-04-08 06:20:17
                    RSVP LSP Type               : -
                    --------------------------------
                            DS-TE Information
                    --------------------------------
                    Bandwidth Reserved Flag : Unreserved
                    CT0 Bandwidth(Kbit/sec) : 0                CT1 Bandwidth(Kbit/sec): 0
                    CT2 Bandwidth(Kbit/sec) : 0                CT3 Bandwidth(Kbit/sec): 0
                    CT4 Bandwidth(Kbit/sec) : 0                CT5 Bandwidth(Kbit/sec): 0
                    CT6 Bandwidth(Kbit/sec) : 0                CT7 Bandwidth(Kbit/sec): 0
                    Setup-Priority          : 7          Hold-Priority        : 7
                    --------------------------------
                              FRR Information
                    --------------------------------
                    Primary LSP Info
                    Bypass In Use             : In Use
                    Bypass Tunnel Id            : 2
                    BypassTunnel               : Tunnel Index[Tunnel2], InnerLabel[16]
                    Bypass LSP ID             : 12         FrrNextHop        : 10.1.5.2
                    ReferAutoBypassHandle : -
                    FrrPrevTunnelTableIndex : -               FrrNextTunnelTableIndex: -
                    Bypass Attribute
                    Setup Priority          : 7          Hold Priority    : 7


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    302
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration

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

                 # Run the display mpls rsvp-te peer command on LSR2 to check whether the
                 bypass CR-LSP is successfully established.
                 [LSR2] display mpls rsvp-te peer
                  Remote Node id Neighbor
                  Neighbor Addr: -----
                  SrcInstance: 0x60128590         NbrSrcInstance: 0x0
                  PSB Count: 1                 RSB Count: 0
                  Hello Type Sent: NONE
                  SRefresh Enable: NO
                  Last valid seq # rcvd: NULL

                  Remote Node id Neighbor
                  Neighbor Addr: 3.3.3.9
                  SrcInstance: 0x60128590           NbrSrcInstance: 0x0
                  PSB Count: 0                   RSB Count: 1
                  Hello Type Sent: NONE
                  SRefresh Enable: NO
                  Last valid seq # rcvd: NULL

                 Interface: Vlanif100
                  Neighbor Addr: 10.1.1.1
                  SrcInstance: 0x60128590           NbrSrcInstance: 0x0
                  PSB Count: 1                   RSB Count: 0
                  Hello Type Sent: NONE
                  SRefresh Enable: NO
                  Last valid seq # rcvd: NULL

                 Interface: Vlanif400
                  Neighbor Addr: 10.1.4.2
                  SrcInstance: 0x60128590           NbrSrcInstance: 0x0
                  PSB Count: 0                   RSB Count: 1
                  Hello Type Sent: NONE
                  SRefresh Enable: NO
                  Last valid seq # rcvd: NULL


                 The RSB count on LSR2's neighbor at 3.3.3.9 is not zero, indicating that RSVP-TE
                 key authentication between LSR2 and its neighbor LSR3 is successful and
                 resources are successfully reserved.

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
                         mpls rsvp-te


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                               303
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration

