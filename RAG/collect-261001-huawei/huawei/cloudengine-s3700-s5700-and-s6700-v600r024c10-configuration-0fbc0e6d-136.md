---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-136
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [19426, 19561]
sha256: d59abe16d053cd97744dbdf436611ee789bed3a4ed38a9f933dde4d13711e790
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         326
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


                 # After the configuration is complete, check the IP routing table on each node.
                 The following example uses the command output on LSR1.
                 [LSR1] display ip routing-table
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 13        Routes : 13

                 Destination/Mask     Proto Pre Cost        Flags NextHop                          Interface

                       1.1.1.9/32 Direct 0 0            D 127.0.0.1                        LoopBack1
                       2.2.2.9/32 ISIS-L2 15 10          D 10.1.1.2                         Vlanif100
                       3.3.3.9/32 ISIS-L2 15 20          D 10.1.1.2                         Vlanif100
                       4.4.4.9/32 ISIS-L2 15 30          D 10.1.1.2                         Vlanif100
                      10.1.1.0/24 Direct 0 0            D 10.1.1.1                         Vlanif100
                      10.1.1.1/32 Direct 0 0            D 127.0.0.1                        Vlanif100
                    10.1.1.255/32 Direct 0 0             D 127.0.0.1                         Vlanif100
                      10.1.2.0/24 ISIS-L2 15 20          D 10.1.1.2                         Vlanif100
                      10.1.3.0/24 ISIS-L2 15 30          D 10.1.1.2                         Vlanif100
                     127.0.0.0/8 Direct 0 0             D 127.0.0.1                        InLoopBack0
                     127.0.0.1/32 Direct 0 0             D 127.0.0.1                        InLoopBack0
                 127.255.255.255/32 Direct 0 0             D 127.0.0.1                          InLoopBack0
                 255.255.255.255/32 Direct 0 0             D 127.0.0.1                          InLoopBack0

                 The command output shows that LSR1 has learned routes to other nodes.
         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.
                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the ingress.
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

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.
         Step 4 Configure IS-IS TE.
                 # Configure LSR1.
                 [LSR1] isis 1
                 [LSR1-isis-1] cost-style wide
                 [LSR1-isis-1] traffic-eng level-2
                 [LSR1-isis-1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.
         Step 5 Configure MPLS TE bandwidth attributes for links.
                 Configure the maximum reservable bandwidth and BC0 bandwidth for links on
                 each interface along the tunnel.
                 # Configure LSR1.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          327
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration

                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls te bandwidth max-reservable-bandwidth 100000
                 [LSR1-Vlanif100] mpls te bandwidth bc0 100000
                 [LSR1-Vlanif100] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 6 Configure the bandwidth flooding threshold.
                 Set the bandwidth flooding threshold to 20 on link interfaces. If the ratio of the
                 bandwidth used or released by the MPLS TE tunnel to the available bandwidth in
                 the TEDB is greater than or equal to 20%, IS-IS floods the bandwidth information,
                 and CSPF updates the TEDB.
                 # Configure LSR1.
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls te bandwidth change thresholds up 20
                 [LSR1-Vlanif100] mpls te bandwidth change thresholds down 20
                 [LSR1-Vlanif100] quit

                 The configuration of LSR2 is similar to the configuration of LSR1. For detailed
                 configurations, see Configuration Scripts.
         Step 7 Configure MPLS TE tunnel interfaces.
                 # On LSR1, configure an MPLS TE tunnel to LSR4.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 4.4.4.9
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] mpls te bandwidth ct0 10000
                 [LSR1-Tunnel1] quit

                 # Display tunnel interface information on LSR1.
                 [LSR1] display mpls te tunnel-interface tunnel1
                    Tunnel Name           : Tunnel1
                    Signalled Tunnel Name: -
                    Tunnel State Desc : CR-LSP is Up
                    Tunnel Attributes :
                    Active LSP          : Primary LSP
                    Traffic Switch       :-
                    Session ID          :1
                    Ingress LSR ID        : 1.1.1.9     Egress LSR ID: 4.4.4.9
                    Admin State            : UP         Oper State : UP
                    Signaling Protocol : RSVP
                    FTid            : 16385
                    Tie-Breaking Policy : None            Metric Type : TE
                    Bfd Cap            : None
                    Reopt             : Disabled       Reopt Freq : -
                    Inter-area Reopt : Disabled
                    Auto BW              : Disabled      Threshold : -
                    Current Collected BW: -               Auto BW Freq : -
                    Min BW               :-           Max BW       :-
                    Offload            : Disabled      Offload Freq : -
                    Low Value            :-           High Value : -
                    Readjust Value          :-
                    Offload Explicit Path Name: -
                    Tunnel Group            : Primary
                    Interfaces Protected: -
                    Excluded IP Address : -
                    Referred LSP Count : 0
                    Primary Tunnel          :-         Pri Tunn Sum : -
                    Backup Tunnel            :-


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                               328
MPLS Configuration
MPLS Configuration                                                                             4 MPLS TE Configuration

