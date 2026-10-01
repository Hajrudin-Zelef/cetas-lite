---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-130
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [18400, 18525]
sha256: 7a30b1adde363bfd9df695640896ec54b6e717ab482970fe7805ef5e547a9b0c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         292
VPN Configuration
VPN Configuration                                                                               4 IPv6 L3VPN Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 100
                    [CE1] interface 10GE 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 100
                    [CE1-10GE1/0/1] quit
                    [CE1] interface Vlanif 100
                    [CE1-Vlanif100] ipv6 enable
                    [CE1-Vlanif100] ipv6 address 2001:db8:1::1 64
                    [CE1-Vlanif100] quit

                    The configurations of CE2, CE3, CE4, PE1, PE2, and the P are similar to the
                    configuration of CE1. For detailed configurations, see Configuration Scripts.
         Step 2 Configure IGP on the IPv4 backbone network for PEs to communicate. IS-IS is used
                as IGP in this example.
                    # Configure PE1.
                    [PE1] isis 1
                    [PE1-isis-1] network-entity 10.1111.1111.1111.00
                    [PE1-isis-1] quit
                    [PE1] interface Vlanif 300
                    [PE1-Vlanif300] isis enable 1
                    [PE1-Vlanif300] quit
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] isis enable 1
                    [PE1-LoopBack1] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
                    After the configuration is complete, PE1, PE2, and the P can learn routes, including
                    the routes to loopback interfaces, from one another. You can run the display ip
                    routing-table command to check route information. The following example uses
                    the command output on PE1.
                    [PE1] display ip routing-table
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table: _public_
                           Destinations : 11         Routes : 11

                    Destination/Mask     Proto Pre Cost       Flags NextHop        Interface

                    1.1.1.9/32      Direct 0 0           D 127.0.0.1    InLoopBack0
                    2.2.2.9/32      ISIS-L2 15 10         D 10.11.11.2     Vlanif300
                    3.3.3.9/32      ISIS-L2 15 20         D 10.11.11.2     Vlanif300
                    127.0.0.0/8      Direct 0 0          D 127.0.0.1     InLoopBack0
                    127.0.0.1/32      Direct 0 0          D 127.0.0.1     InLoopBack0
                    127.255.255.255/32 Direct 0 0           D 127.0.0.1       InLoopBack0
                    10.11.11.0/24     Direct 0 0          D 10.11.11.1     Vlanif300
                    10.11.11.1/32     Direct 0 0          D 127.0.0.1      Vlanif300
                    10.11.11.255/32    Direct 0 0          D 127.0.0.1      Vlanif300
                    10.12.12.0/24     ISIS-L2 15 20        D 10.11.11.2      Vlanif300
                    255.255.255.255/32 Direct 0 0           D 127.0.0.1       InLoopBack0

         Step 3 Enable MPLS and MPLS LDP both globally and per interface on each device of the
                IPv4 backbone network to establish an LDP LSP between PE1 and PE2.
                    # Enable MPLS and MPLS LDP on PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                         293
VPN Configuration
VPN Configuration                                                                                       4 IPv6 L3VPN Configuration

                    [PE1-mpls-ldp] quit
                    [PE1] interface Vlanif 300
                    [PE1-Vlanif300] mpls
                    [PE1-Vlanif300] mpls ldp
                    [PE1-Vlanif300] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
                    After the configuration is complete, an LDP LSP should exist between PE1 and
                    PE2. Run the display mpls ldp lsp command. The command output shows that an
                    LDP LSP has been established. The following example uses the command output
                    on PE1.
                    [PE1] display mpls ldp lsp
                     LDP LSP Information
                    -------------------------------------------------------------------------------
                    Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                    -------------------------------------------------------------------------------
                     DestAddress/Mask In/OutLabel UpstreamPeer NextHop                             OutInterface
                    -------------------------------------------------------------------------------
                     1.1.1.9/32        3/NULL          2.2.2.9        127.0.0.1       InLoop0
                    *1.1.1.9/32         Liberal/1024                 DS/2.2.2.9
                     2.2.2.9/32        NULL/3          -            10.11.11.2       Vlanif300
                     2.2.2.9/32        1024/3          2.2.2.9       10.11.11.2       Vlanif300
                     3.3.3.9/32        NULL/1025         -            10.11.11.2       Vlanif300
                     3.3.3.9/32        1025/1025         2.2.2.9       10.11.11.2       Vlanif300
                     -------------------------------------------------------------------------------
                     TOTAL: 5 Normal LSP(s) Found.
                     TOTAL: 1 Liberal LSP(s) Found.
                     TOTAL: 0 Frr LSP(s) Found.
                     An asterisk (*) before an LSP means the LSP is not established
                     An asterisk (*) before a Label means the USCB or DSCB is stale
                     An asterisk (*) before an UpstreamPeer means the session is stale
                     An asterisk (*) before a DS means the session is stale
                     An asterisk (*) before a NextHop means the LSP is FRR LSP

         Step 4 Configure an IPv6-address-family-enabled VPN instance on PEs and bind the
                interfaces connecting PEs to CEs to the corresponding VPN instances.
                    # On PE1, configure an IPv6-address-family-enabled VPN instance named vpna.
                    [PE1] ip vpn-instance vpna
                    [PE1-vpn-instance-vpna] ipv6-family
                    [PE1-vpn-instance-vpna-af-ipv6] route-distinguisher 100:1
                    [PE1-vpn-instance-vpna-af-ipv6] vpn-target 22:22 export-extcommunity
                    [PE1-vpn-instance-vpna-af-ipv6] vpn-target 33:33 import-extcommunity
                    [PE1-vpn-instance-vpna-af-ipv6] quit
                    [PE1-vpn-instance-vpna] quit

                    # Bind the interface that directly connects PE1 to CE1 to the VPN instance named
                    vpna.
                    [PE1] interface Vlanif 100
                    [PE1-Vlanif100] ip binding vpn-instance vpna
                    [PE1-Vlanif100] ipv6 enable
                    [PE1-Vlanif100] ipv6 address 2001:db8:1::2 64
                    [PE1-Vlanif100] quit

