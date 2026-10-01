---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-116
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [16415, 16547]
sha256: 70904c1d578768fe00d73a79a87185864b3c8e26dab2764ee5cf2386f1b9d42a
---

                 # Disable the protected outbound interface VLANIF200 on the PLR (LSR2).
                 [LSR2] interface vlanif 200
                 [LSR2-Vlanif200] shutdown
                 [LSR2-Vlanif200] quit

                 # Check the status of the primary CR-LSP on LSR1. The command output shows
                 that the tunnel interface is still up.
                 [LSR1] display interface tunnel 1
                 Tunnel1 current state : UP (ifindex: 28)
                 Line protocol current state : UP
                 Last line protocol up time : 2024-04-08 06:20:17
                 Description:
                 ...

                 # Run the tracert lsp te tunnel 1 command on LSR1 to check the path of the
                 tunnel.
                 [LSR1] tracert lsp te tunnel 1
                  LSP Trace Route FEC: TE TUNNEL IPV4 SESSION QUERY Tunnel1 , press CTRL_C to break.
                  TTL Replier           Time Type       Downstream
                  0                       Ingress 10.1.1.2/[17 ]
                  1   10.1.1.2        2 ms Transit 10.1.4.2/[16 ]
                  2   10.1.4.2        2 ms Transit 10.1.5.2/[3 ]
                  3   10.1.5.2        1 ms Transit 10.1.3.2/[3 ]
                  4   4.4.4.9         11     Egress

                 The command output indicates that the link has been switched to the bypass CR-
                 LSP.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                 277
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration


                 # Run the display mpls te tunnel name Tunnel1 verbose command on LSR2.
                 The command output shows that the bypass tunnel has been used.
                 [LSR2] display mpls te tunnel name Tunnel1 verbose
                    No                   : 1
                    Tunnel-Name                   : Tunnel1
                    Tunnel Interface Name : -
                    TunnelIndex               : -
                    Session ID             : 1           LSP ID          : 690
                    LSR Role               : Transit
                    Ingress LSR ID            : 1.1.1.9
                    Egress LSR ID             : 4.4.4.9
                    In-Interface           : Vlanif100
                    Out-Interface             : Vlanif200
                    Sign-Protocol             : RSVP TE       Resv Style      : SE
                    IncludeAnyAff               : 0x0       ExcludeAnyAff      : 0x0
                    IncludeAllAff           : 0x0
                    ER-Hop Table Index              : -      AR-Hop Table Index: -
                    C-Hop Table Index              : -
                    PrevTunnelIndexInSession: -                NextTunnelIndexInSession: -
                    PSB Handle                 : -
                    Created Time                : 2024-04-08 06:20:17
                    RSVP LSP Type                : -
                    --------------------------------
                            DS-TE Information
                    --------------------------------
                    Bandwidth Reserved Flag : Unreserved
                    CT0 Bandwidth(Kbit/sec) : 0                 CT1 Bandwidth(Kbit/sec): 0
                    CT2 Bandwidth(Kbit/sec) : 0                 CT3 Bandwidth(Kbit/sec): 0
                    CT4 Bandwidth(Kbit/sec) : 0                 CT5 Bandwidth(Kbit/sec): 0
                    CT6 Bandwidth(Kbit/sec) : 0                 CT7 Bandwidth(Kbit/sec): 0
                    Setup-Priority           : 7          Hold-Priority        : 7
                    --------------------------------
                              FRR Information
                    --------------------------------
                    Primary LSP Info
                    Bypass In Use              : In Use
                    Bypass Tunnel Id             : 2
                    BypassTunnel                : Tunnel Index[Tunnel2], InnerLabel[16]
                    Bypass LSP ID              : 12         FrrNextHop        : 10.1.5.2
                    ReferAutoBypassHandle : -
                    FrrPrevTunnelTableIndex : -                FrrNextTunnelTableIndex: -
                    Bypass Attribute
                    Setup Priority          : 7           Hold Priority    : 7
                    HopLimit                : 32          Bandwidth         : 0
                    IncludeAnyGroup                : 0       ExcludeAnyGroup : 0
                    IncludeAllGroup              : 0
                    Bypass Unbound Bandwidth Info(Kbit/sec)
                    CT0 Unbound Bandwidth : -                    CT1 Unbound Bandwidth: -
                    CT2 Unbound Bandwidth : -                    CT3 Unbound Bandwidth: -
                    CT4 Unbound Bandwidth : -                    CT5 Unbound Bandwidth: -
                    CT6 Unbound Bandwidth : -                    CT7 Unbound Bandwidth: -
                    --------------------------------
                             BFD Information
                    --------------------------------
                    NextSessionTunnelIndex : -                 PrevSessionTunnelIndex: -
                    NextLspId               : -          PrevLspId        : -

                 # Run the display mpls rsvp-te statistics global command on LSR2 to check
                 Srefresh statistics.
                 [LSR2] display mpls rsvp-te statistics global
                  LSR ID: 2.2.2.9             LSP Count: 2
                  PSB Count: 2                 RSB Count: 2
                  RFSB Count: 1

                 Total Statistics Information:
                  PSB CleanupTimeOutCounter: 0           RSB CleanupTimeOutCounter: 0
                  SendPacketCounter: 122707             RecPacketCounter: 127580


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     278
MPLS Configuration
MPLS Configuration                                                                   4 MPLS TE Configuration

                  SendCreatePathCounter: 27        RecCreatePathCounter: 304
                  SendRefreshPathCounter: 62220      RecRefreshPathCounter: 62122
                  SendCreateResvCounter: 22        RecCreateResvCounter: 32
                  SendRefreshResvCounter: 60111      RecRefreshResvCounter: 64803
                  SendResvConfCounter: 0          RecResvConfCounter: 0
                  SendHelloCounter: 0           RecHelloCounter: 0
                  SendAckCounter: 0             RecAckCounter: 0
                  SendPathErrCounter: 287         RecPathErrCounter: 0
                  SendResvErrCounter: 0          RecResvErrCounter: 0
                  SendPathTearCounter: 11         RecPathTearCounter: 8
                  SendResvTearCounter: 2          RecResvTearCounter: 0
                  SendSrefreshCounter: 13          RecSrefreshCounter: 14
                  SendAckMsgCounter: 14             RecAckMsgCounter: 13
                  SendChallengeMsgCounter: 0        RecChallengeMsgCounter: 0
                  SendResponseMsgCounter: 0         RecResponseMsgCounter: 0
                  SendErrMsgCounter: 0           RecErrMsgCounter: 0
                  SendRecoveryPathMsgCounter: 0       RecRecoveryPathMsgCounter: 0
                  SendGRPathMsgCounter: 0           RecGRPathMsgCounter: 0
                  ResourceReqFaultCounter: 0       RecGRPathMsgFromLSPMCounter: 0
                  Bfd neighbor count: 2         Bfd session count: 0

                 Srefresh takes effect globally on LSR2 and LSR3. Therefore, if the primary tunnel
                 fails, RSVP-TE Srefresh can still work properly on other interfaces of LSR2 and
                 LSR3.

