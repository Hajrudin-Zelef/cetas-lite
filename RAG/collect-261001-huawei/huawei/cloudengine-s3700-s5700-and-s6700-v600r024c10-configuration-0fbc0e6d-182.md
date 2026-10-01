---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-182
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-08", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [26423, 26559]
sha256: d5e224f144309b98486b1fa4afae7625517c209d9e2b9eb6dba335a329d40d2f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

         Step 6 Create a bypass tunnel on LSR2 that functions as a PLR.
                 # Configure an explicit path for the bypass tunnel.
                 [LSR2] explicit-path by-path
                 [LSR2-explicit-path-by-path] next hop 10.1.4.2
                 [LSR2-explicit-path-by-path] next hop 10.1.5.2
                 [LSR2-explicit-path-by-path] next hop 3.3.3.9
                 [LSR2-explicit-path-by-path] quit

                 # Create a bypass tunnel on LSR2 and bind the bypass tunnel to the explicit path.
                 [LSR2] interface Tunnel 2
                 [LSR2-Tunnel2] ip address unnumbered interface loopback 1
                 [LSR2-Tunnel2] tunnel-protocol mpls te
                 [LSR2-Tunnel2] destination 3.3.3.9
                 [LSR2-Tunnel2] mpls te tunnel-id 2
                 [LSR2-Tunnel2] mpls te path explicit-path by-path
                 [LSR2-Tunnel2] mpls te bypass-tunnel

                 # Bind the bypass tunnel to the protected interface.
                 [LSR2-Tunnel2] mpls te protected-interface vlanif 200
                 [LSR2-Tunnel2] quit

                 # After the configuration is complete, run the display interface tunnel command
                 on LSR2. The command output shows that Tunnel2 is up.
                 [LSR2] display interface tunnel
                 Tunnel2 current state : UP (ifindex: 33)
                 Line protocol current state : UP
                 Last line protocol up time : 2024-04-08 06:19:56
                 Description:
                 ...

                 # Run the display mpls te tunnel command on LSR2. The command output
                 shows that two tunnels pass through LSR2.
                 [LSR2]display mpls te tunnel
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name
                 -------------------------------------------------------------------------------
                 1.1.1.9        4.4.4.9       690 16/16              T Tunnel1
                 2.2.2.9        3.3.3.9       12 -/16              I Tunnel2
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress

                 # Run the display mpls te tunnel bypass-inuse exists-not-used command on
                 LSR2 to view information about the tunnels that are enabled with TE FRR and
                 have unused bypass tunnels.
                 [LSR2]display mpls te tunnel bypass-inuse exists-not-used
                 * means the LSP is detour LSP
                 -------------------------------------------------------------------------------
                 Ingress LsrId Destination         LSPID In/OutLabel        R Tunnel-name


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          441
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration

                 -------------------------------------------------------------------------------
                 1.1.1.9        4.4.4.9       776 17/17             T Tunnel1
                 -------------------------------------------------------------------------------
                 R: Role, I: Ingress, T: Transit, E: Egress

                 # Run the display mpls te tunnel name Tunnel1 verbose or display mpls te
                 tunnel bypass-inuse exists-not-used verbose command on LSR2. The command
                 output shows that the bypass tunnel is bound to the outbound interface VLANIF
                 200 and is not in use.
                 [LSR2]display mpls te tunnel name Tunnel1 verbose
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
                    Bypass In Use              : Not Used
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

                 ----End




Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          442
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration


Verifying the Configuration
                 # Shut down the protected outbound interface on the PLR (LSR2).
                 [LSR2] interface vlanif 200
                 [LSR2-Vlanif200] shutdown
                 [LSR2-Vlanif200] quit

