---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-183
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [26560, 26695]
sha256: 109bcca45fac719f7a9177bf0e7b2af85dea48d5433e83786b22b46358643268
---

                 # Run the display interface tunnel 1 command on LSR1 to check the status of
                 the primary CR-LSP. The command output shows that Tunnel1 is still up.
                 [LSR1] display interface tunnel 1
                 Tunnel1 current state : UP (ifindex: 28)
                 Line protocol current state : UP
                 Last line protocol up time : 2024-04-08 06:20:17
                 Description:
                 ...

                 # Run the display mpls te tunnel path command on LSR1 to view tunnel path
                 information on the local node.
                 [LSR1] display mpls te tunnel path
                  Tunnel Interface Name : Tunnel1
                  Lsp ID : 1.1.1.9 :1 :776
                  Hop Information
                   Hop 0 1.1.1.9
                   Hop 1 10.1.1.1
                   Hop 2 10.1.1.2 Label 17
                   Hop 3 2.2.2.9 Label 17
                   Hop 4 10.1.4.1 Local-Protection in use
                   Hop 5 10.1.5.2 Label 17
                   Hop 6 3.3.3.9 Label 17
                   Hop 7 10.1.3.1
                   Hop 8 10.1.3.2 Label 3
                   Hop 9 4.4.4.9 Label 3

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

                 The preceding information indicates that the link has been switched to the bypass
                 tunnel (bypass CR-LSP).

                         NOTE

                        If you run the display mpls te tunnel-interface command immediately after TE FRR
                        switching to view detailed information about the tunnel interface, you can view two CR-
                        LSPs in the up state. This is because TE FRR sets up a new CR-LSP in make-before-break
                        mode and the old CR-LSP is not deleted until the new CR-LSP is set up.

                 # Run the display mpls te tunnel bypass-inuse inuse command on LSR2 to view
                 information about TE FRR tunnels that are using the bypass tunnel.
                 [LSR2]display mpls te tunnel bypass-inuse exists-not-used
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------
                 1.1.1.9        4.4.4.9       776 17/17             T Tunnel1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          443
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration

                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress

                 # Run the display mpls te tunnel name Tunnel1 verbose command on LSR2.
                 The command output shows that the bypass tunnel is in use.
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

                 # Make the protected outbound interface on the PLR take effect again.
                 [LSR2] interface vlanif 200
                 [LSR2-Vlanif200] undo shutdown

                 Run the display interface tunnel 1 command on LSR1 to check the status of the
                 primary CR-LSP. The command output shows that the tunnel interface is up.


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          444
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


