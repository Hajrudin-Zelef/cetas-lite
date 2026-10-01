---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-181
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [26277, 26422]
sha256: cec5160583203bd43d5acf8c283889bb995990b29d8bc165d11739c6f631d373
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.
                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along tunnels. Enable CSPF on the ingress (LSR1) of the primary tunnel and the
                 ingress (LSR2) of the bypass tunnel.
                 # Configure LSR1.
                 [LSR1] mpls lsr-id 1.1.1.9
                 [LSR1] mpls
                 [LSR1-mpls] mpls te
                 [LSR1-mpls] mpls rsvp-te
                 [LSR1-mpls] mpls te cspf
                 [LSR1-mpls] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls
                 [LSR1-Vlanif100] mpls te
                 [LSR1-Vlanif100] mpls rsvp-te
                 [LSR1-Vlanif100] quit

                 The configurations of LSR2, LSR3, LSR4, and LSR5 are similar to the configuration
                 of LSR1. For detailed configurations, see Configuration Scripts.
         Step 4 Configure IS-IS TE.
                 # Configure LSR1.
                 [LSR1] isis 1
                 [LSR1-isis-1] cost-style wide
                 [LSR1-isis-1] traffic-eng level-2
                 [LSR1-isis-1] quit

                 The configurations of LSR2, LSR3, LSR4, and LSR5 are similar to the configuration
                 of LSR1. For detailed configurations, see Configuration Scripts.
         Step 5 Create the primary MPLS TE tunnel on the ingress LSR1.
                 # Configure an explicit path for the primary tunnel.
                 [LSR1] explicit-path pri-path
                 [LSR1-explicit-path-pri-path] next hop 10.1.1.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.2.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.3.2
                 [LSR1-explicit-path-pri-path] next hop 4.4.4.9
                 [LSR1-explicit-path-pri-path] quit

                 # Configure the primary MPLS TE tunnel and bind it to the explicit path.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 4.4.4.9
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] mpls te path explicit-path pri-path

                 # Enable TE FRR.
                 [LSR1-Tunnel1] mpls te fast-reroute
                 [LSR1-Tunnel1] quit

                 # After the configuration is complete, run the display interface tunnel command
                 to check the working status of the tunnel interface. The command output shows
                 that the tunnel interface is up.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         439
MPLS Configuration
MPLS Configuration                                                                           4 MPLS TE Configuration

                 [LSR1] display interface tunnel
                 Tunnel1 current state : UP (ifindex: 28)
                 Line protocol current state : UP
                 Last line protocol up time : 2024-04-08 06:20:17
                 Description:
                 ...

                 # Run the display mpls te tunnel verbose command to view detailed information
                 about the tunnel interface.
                 [LSR1] display mpls te tunnel verbose
                    No                   : 1
                    Tunnel-Name                   : Tunnel1
                    Tunnel Interface Name : Tunnel1
                    TunnelIndex               : -
                    Session ID             : 1           LSP ID          : 690
                    LSR Role               : Ingress
                    Ingress LSR ID            : 1.1.1.9
                    Egress LSR ID             : 4.4.4.9
                    In-Interface           : -
                    Out-Interface             : Vlanif100
                    Sign-Protocol             : RSVP TE       Resv Style       : SE
                    IncludeAnyAff               : 0x0       ExcludeAnyAff       : 0x0
                    IncludeAllAff           : 0x0
                    ER-Hop Table Index              : 1       AR-Hop Table Index: 24578
                    C-Hop Table Index              : 16770
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
                    Setup-Priority           : 7          Hold-Priority         : 7
                    --------------------------------
                              FRR Information
                    --------------------------------
                    Primary LSP Info
                    Bypass In Use              : Not Exists
                    Bypass Tunnel Id             : 0
                    BypassTunnel                : -
                    Bypass LSP ID              : -         FrrNextHop         : -
                    ReferAutoBypassHandle : -
                    FrrPrevTunnelTableIndex : -                FrrNextTunnelTableIndex: -
                    Bypass Attribute
                    Setup Priority          : 7           Hold Priority    : 7
                    HopLimit                : 32           Bandwidth        : 0
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

                 # Run the display mpls te tunnel path command to check tunnel path
                 information on the local node.
                 [LSR1] display mpls te tunnel path
                  Tunnel Interface Name : Tunnel1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     440
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration

                  Lsp ID : 1.1.1.9 :1 :776
                  Hop Information
                   Hop 0 1.1.1.9
                   Hop 1 10.1.1.1
                   Hop 2 10.1.1.2 Label 17
                   Hop 3 2.2.2.9 Label 17
                   Hop 4 10.1.2.1 Local-Protection available
                   Hop 5 10.1.2.2 Label 17
                   Hop 6 3.3.3.9 Label 17
                   Hop 7 10.1.3.1
                   Hop 8 10.1.3.2 Label 3
                   Hop 9 4.4.4.9 Label 3

