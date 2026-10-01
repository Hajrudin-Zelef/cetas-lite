---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-154
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [22215, 22382]
sha256: 038c1626e1f1894181e80eae31e395ea497711fd85050e90a7d062f0fb477ec9
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The command output shows that Local-Protection is configured on the outbound
                 interface (10.2.1.1) of the primary tunnel on P1.

                 ----End

Verifying the Configuration
                 # Check information about the primary tunnel and its bound bypass tunnel on P1.
                 [P1] display mpls te tunnel name Tunnel1 verbose
                   No                : 1
                   Tunnel-Name            : Tunnel1
                   Tunnel Interface Name : Tunnel1
                   TunnelIndex          : -


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                       370
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration

                     Session ID             : 100          LSP ID        : 1
                     Lsr Role              : Transit
                     Ingress LSR ID            : 4.4.4.4
                     Egress LSR ID             : 5.5.5.5
                     In-Interface           : Vlanif100
                     Out-Interface             : Vlanif200
                     Sign-Protocol             : RSVP TE      Resv Style       : SE
                     IncludeAnyAff               : 0x0       ExcludeAnyAff      : 0x0
                     IncludeAllAff           : 0x0
                     ER-Hop Table Index              : -      AR-Hop Table Index: 2
                     C-Hop Table Index              : -
                     PrevTunnelIndexInSession: -                NextTunnelIndexInSession: -
                     PSB Handle                 : 65546
                     Created Time                : 2010/10/15 09:52:03
                     --------------------------------
                             DS-TE Information
                     --------------------------------
                     Bandwidth Reserved Flag : Reserved
                     CT0 Bandwidth(Kbit/sec) : 10000              CT1 Bandwidth(Kbit/sec): 0
                     CT2 Bandwidth(Kbit/sec) : 0                CT3 Bandwidth(Kbit/sec): 0
                     CT4 Bandwidth(Kbit/sec) : 0                CT5 Bandwidth(Kbit/sec): 0
                     CT6 Bandwidth(Kbit/sec) : 0                CT7 Bandwidth(Kbit/sec): 0
                     Setup-Priority           : 7          Hold-Priority   : 7
                     --------------------------------
                               FRR Information
                     --------------------------------
                     Primary LSP Info
                     Bypass In Use              : Not Used
                     Bypass Tunnel Id             : 67141670
                     BypassTunnel                  : Tunnel Index[AutoTunnel67141670], InnerLabel[3]
                     Bypass Lsp ID              : 1        FrrNextHop         : 5.5.5.5
                     ReferAutoBypassHandle : -
                     FrrPrevTunnelTableIndex : -               FrrNextTunnelTableIndex: -
                     Bypass Attribute(Not configured)
                     Setup Priority          : -          Hold Priority   : -
                     HopLimit                : -          Bandwidth        : -
                     IncludeAnyGroup                : -      ExcludeAnyGroup : -
                     IncludeAllGroup              : -
                     Bypass Unbound Bandwidth Info(Kbit/sec)
                     CT0 Unbound Bandwidth : -                  CT1 Unbound Bandwidth: -
                     CT2 Unbound Bandwidth : -                  CT3 Unbound Bandwidth: -
                     CT4 Unbound Bandwidth : -                  CT5 Unbound Bandwidth: -
                     CT6 Unbound Bandwidth : -                  CT7 Unbound Bandwidth: -
                     --------------------------------
                              BFD Information
                     --------------------------------
                     NextSessionTunnelIndex : -                PrevSessionTunnelIndex: -
                     NextLspId               : -          PrevLspId      : -

                 The command output shows that a bypass tunnel is bound to the primary tunnel
                 and the FrrNextHop value is 5.5.5.5.
                 # Check the path of the bypass tunnel.
                 [P1] display mpls te tunnel path AutoTunnel67141670
                  Tunnel Interface Name : AutoTunnel67141670
                  Lsp ID : 1.1.1.1 :67141670 :1
                  Hop Information
                   Hop 0 10.3.1.1
                   Hop 1 10.3.1.2
                   Hop 2 2.2.2.2
                   Hop 3 10.4.1.1
                   Hop 4 10.4.1.2
                   Hop 5 5.5.5.5

                 The command output shows that the path of the automatic bypass tunnel is P1 ->
                 P2 ->PE2.



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         371
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


Configuration Scripts
                 ●      PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        explicit-path main
                         next hop 10.1.1.2
                         next hop 10.2.1.2
                         next hop 5.5.5.5
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0004.00
                         traffic-eng level-1-2
                        #
                        interface vlanif 100
                         ip address 10.1.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                         isis enable 1
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 5.5.5.5
                         mpls te tunnel-id 100
                         mpls te record-route
                         mpls te bandwidth ct0 10000
                         mpls te path explicit-path main
                         mpls te fast-reroute
                        #
                        return

                 ●      P1
                        #
                        sysname P1
                        #
                        vlan batch 100 200 300 500
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                         mpls te srlg path-calculation preferred
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0001.00


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      372
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

