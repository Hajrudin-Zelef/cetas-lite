---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-212
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2013-01-21", "2013-09-16", "2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [30790, 30918]
sha256: 6fec5ccf0cdb4008bcaf202614ef744c48d2a02e6b742ad51a169a30f06516ff
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 [LSR1-Tunnel1] mpls te fast-reroute
                 [LSR1-Tunnel1] quit

                 # After the configuration is complete, run the display interface tunnel command
                 on LSR1. The command output shows that Tunnel1 is up.
                 [LSR1] display interface tunnel 1
                 Tunnel1 current state : UP
                 Line protocol current state : UP
                 Last line protocol up time : 2013-01-21 10:58:49
                 Description:
                 ...

         Step 6 Configure a bypass CR-LSP on LSR2 that functions as a PLR.
                 # Configure an explicit path for the bypass CR-LSP.
                 [LSR2] explicit-path by-path
                 [LSR2-explicit-path-by-path] next hop 10.1.4.2
                 [LSR2-explicit-path-by-path] next hop 10.1.5.2
                 [LSR2-explicit-path-by-path] next hop 3.3.3.9
                 [LSR2-explicit-path-by-path] quit

                 # Configure a tunnel interface for the bypass CR-LSP.
                 [LSR2] interface tunnel 2
                 [LSR2-Tunnel2] ip address unnumbered interface loopback 1
                 [LSR2-Tunnel2] tunnel-protocol mpls te
                 [LSR2-Tunnel2] destination 3.3.3.9
                 [LSR2-Tunnel2] mpls te tunnel-id 300
                 [LSR2-Tunnel2] mpls te path explicit-path by-path
                 [LSR2-Tunnel2] mpls te bypass-tunnel

                 # Bind the bypass CR-LSP to the protected interface.
                 [LSR2-Tunnel2] mpls te protected-interface vlanif 200
                 [LSR2-Tunnel2] quit

                 # After the configuration is complete, check LSP entries. The command output
                 shows that two tunnels pass through LSR2 and LSR3. The following example uses
                 the command output on LSR2.
                 [LSR2] display mpls lsp
                 -------------------------------------------------------------------------------
                              LSP Information: RSVP LSP
                 -------------------------------------------------------------------------------
                 FEC              In/Out Label In/Out IF                      Vrf Name
                 4.4.4.9/32         1059/1068       Vlanif100/Vlanif200
                 3.3.3.9/32         NULL/1036       -/Vlanif400

                 # Check tunnel establishment. The command output shows that two tunnels pass
                 through LSR2 and LSR3. The following example uses the command output on
                 LSR2.
                 [LSR2] display mpls te tunnel
                 ------------------------------------------------------------------------------
                 Ingress LsrId Destination           LSPID In/Out Label         R Tunnel-name
                 ------------------------------------------------------------------------------
                 1.1.1.9        4.4.4.9         40      1059/1068         T Tunnel1
                 2.2.2.9        3.3.3.9         4      --/1036         I Tunnel2

                 # Check detailed information about the tunnel. The command output shows that
                 the bypass tunnel is bound to the outbound interface VLANIF200 and is not in use.
                 [LSR2] display mpls te tunnel name Tunnel1 verbose
                    No                : 1
                    Tunnel-Name           : Tunnel1
                    Tunnel Interface Name : -


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          510
MPLS Configuration
MPLS Configuration                                                                             4 MPLS TE Configuration

                     TunnelIndex               : 3         LSP Index        : 2048
                     Session ID             : 100          LSP ID         : 40
                     LSR Role               : Transit
                     Ingress LSR ID            : 1.1.1.9
                     Egress LSR ID             : 4.4.4.9
                     In-Interface           : Vlanif100
                     Out-Interface             : Vlanif200
                     Sign-Protocol             : RSVP TE       Resv Style      : SE
                     IncludeAnyAff               : 0x0       ExcludeAnyAff      : 0x0
                     IncludeAllAff           : 0x0
                     ER-Hop Table Index              : 3      AR-Hop Table Index: 1
                     C-Hop Table Index              : -
                     PrevTunnelIndexInSession: -                NextTunnelIndexInSession: -
                     PSB Handle                 : 8200
                     Created Time                : 2013-09-16 12:52:03+00:00
                     RSVP LSP Type                : -
                     --------------------------------
                             DS-TE Information
                     --------------------------------
                     Bandwidth Reserved Flag : Unreserved
                     CT0 Bandwidth(Kbit/sec) : 0                 CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec) : 0                 CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec) : 0                 CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec) : 0                 CT7 Bandwidth(Kbit/sec): 0
                     Setup-Priority           : 7          Hold-Priority    : 7
                     --------------------------------
                               FRR Information
                     --------------------------------
                     Primary LSP Info
                     TE Attribute Flag           : 0x63       Protected Flag : 0x1
                     Bypass In Use              : Not Used
                     Bypass Tunnel Id             : 15
                     BypassTunnel                  : Tunnel Index[Tunnel2], InnerLabel[1068]
                     Bypass LSP ID              : 4         FrrNextHop         : 10.1.5.2
                     ReferAutoBypassHandle : -
                     FrrPrevTunnelTableIndex : -                FrrNextTunnelTableIndex: -
                     Bypass Attribute(Not configured)
                     Setup Priority          : -          Hold Priority    : -
                     HopLimit                : -          Bandwidth         : -
                     IncludeAnyGroup                : -      ExcludeAnyGroup : -
                     IncludeAllGroup              : -
                     Bypass Unbound Bandwidth Info(Kbit/sec)
                     CT0 Unbound Bandwidth : -                   CT1 Unbound Bandwidth: -
                     CT2 Unbound Bandwidth : -                   CT3 Unbound Bandwidth: -
                     CT4 Unbound Bandwidth : -                   CT5 Unbound Bandwidth: -
                     CT6 Unbound Bandwidth : -                   CT7 Unbound Bandwidth: -
                     --------------------------------
                              BFD Information
                     --------------------------------
                     NextSessionTunnelIndex : -                 PrevSessionTunnelIndex: -
                     NextLspId               : -          PrevLspId       : -

         Step 7 Configure an ordinary backup CR-LSP on the ingress LSR1.

                 # Specify an explicit path for the backup CR-LSP.
                 [LSR1] explicit-path backup-path
                 [LSR1-explicit-path-backup-path] next hop 10.1.6.1
                 [LSR1-explicit-path-backup-path] next hop 10.1.7.2
                 [LSR1-explicit-path-backup-path] next hop 10.1.3.2
                 [LSR1-explicit-path-backup-path] next hop 4.4.4.9
                 [LSR1-explicit-path-backup-path] quit

