---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-37
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [4267, 4406]
sha256: aa660885eb6d4c699d67646dd2c673a759d71020a3079cc3c6db460be582189b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    ●    On the same VPN, the export VPN target list of a site shares VPN targets with
                         the import VPN target lists of the other sites. Conversely, the import VPN
                         target list of a site shares VPN targets with the export VPN target lists of the
                         other sites.
                    ●    After a PE interface connected to a CE is bound to a VPN instance, Layer 3
                         configurations on this interface are automatically deleted. Such configurations
                         include IP address and routing protocol configurations, and must be added
                         again if needed.


Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Enable OSPF on the backbone network to ensure that PEs can communicate.
                    2.   Configure basic MPLS capabilities and MPLS LDP to establish LDP LSPs on the
                         backbone network.
                    3.   Configure a VPN instance on each PE, enable the IPv4 address family for the
                         instance, and bind the interface that connects each PE to a CE to the VPN
                         instance on that PE.
                    4.   Enable MP-IBGP on PEs to exchange VPN routing information.
                    5.   Configure EBGP between CEs and PEs to exchange VPN routing information.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                             68
VPN Configuration
VPN Configuration                                                              3 IPv4 L3VPN Configuration


Procedure
         Step 1 Configure IGP to achieve connectivity between devices, including PEs and the P, on
                the MPLS backbone network. OSPF is used as IGP in this example.
                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 32
                    [PE1-LoopBack1] quit
                    [PE1] vlan batch 100 200 300
                    [PE1] interface 10GE1/0/3
                    [PE1-10GE1/0/3] port link-type trunk
                    [PE1-10GE1/0/3] port trunk allow-pass vlan 300
                    [PE1-10GE1/0/3] quit
                    [PE1] interface Vlanif 300
                    [PE1-Vlanif300] ip address 11.11.11.1 24
                    [PE1-Vlanif300] quit
                    [PE1] ospf
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 11.11.11.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    # Configure the P.
                    <HUAWEI> system-view
                    [HUAWEI] sysname P
                    [P] interface loopback 1
                    [P-LoopBack1] ip address 2.2.2.9 32
                    [P-LoopBack1] quit
                    [P] vlan batch 200 300
                    [P] interface 10GE
                    [P-10GE1/0/3] port link-type trunk
                    [P-10GE1/0/3] port trunk allow-pass vlan 300
                    [P-10GE1/0/3] quit
                    [P] interface Vlanif 300
                    [P-Vlanif300] ip address 11.11.11.2 24
                    [P-Vlanif300] quit
                    [P] interface 10GE1/0/2
                    [P-10GE1/0/2] port link-type trunk
                    [P-10GE1/0/2] port trunk allow-pass vlan 200
                    [P-10GE1/0/2] quit
                    [P] interface Vlanif 200
                    [P-Vlanif200] ip address 12.12.12.1 24
                    [P-Vlanif200] quit
                    [P] ospf
                    [P-ospf-1] area 0
                    [P-ospf-1-area-0.0.0.0] network 11.11.11.0 0.0.0.255
                    [P-ospf-1-area-0.0.0.0] network 12.12.12.0 0.0.0.255
                    [P-ospf-1-area-0.0.0.0] network 2.2.2.9 0.0.0.0
                    [P-ospf-1-area-0.0.0.0] quit
                    [P-ospf-1] quit

                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] interface loopback 1
                    [PE2-LoopBack1] ip address 3.3.3.9 32
                    [PE2-LoopBack1] quit
                    [PE2] vlan batch 100 200 300
                    [PE2] interface 10GE
                    [PE2-10GE1/0/2] port link-type trunk
                    [PE2-10GE1/0/2] port trunk allow-pass vlan 200
                    [PE2-10GE1/0/2] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                          69
VPN Configuration
VPN Configuration                                                                               3 IPv4 L3VPN Configuration

                    [PE2] interface Vlanif 200
                    [PE2-Vlanif200] ip address 12.12.12.2 24
                    [PE2-Vlanif200] quit
                    [PE2] ospf
                    [PE2-ospf-1] area 0
                    [PE2-ospf-1-area-0.0.0.0] network 12.12.12.0 0.0.0.255
                    [PE2-ospf-1-area-0.0.0.0] network 3.3.3.9 0.0.0.0
                    [PE2-ospf-1-area-0.0.0.0] quit
                    [PE2-ospf-1] quit

                    After the configuration is complete, OSPF neighbor relationships can be
                    established between PE1, the P, and PE2. Run the display ospf peer command.
                    The command output shows that the neighbor status is Full. Run the display ip
                    routing-table command. The command output shows that the PEs have learned
                    the routes to each other's Loopback 1.
                    The following example uses the command output on PE1.
                    [PE1] display ip routing-table
                    Proto: Protocol        Pre: Preference
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table: _public_
                           Destinations : 11        Routes : 11


                    Destination/Mask Proto Pre Cost             Flags NextHop      Interface
                        1.1.1.9/32 Direct 0 0               D 127.0.0.1    LoopBack1
                        2.2.2.9/32 OSPF 10 2                 D 11.11.11.2     Vlanif300
                        3.3.3.9/32 OSPF 10 3                 D 11.11.11.2     Vlanif300
                      11.11.11.0/24 Direct 0 0               D 11.11.11.1    Vlanif300
                      11.11.11.1/32 Direct 0 0               D 127.0.0.1     Vlanif300
                     11.11.11.255/32 Direct 0 0               D 127.0.0.1     Vlanif300
                      12.12.12.0/24 OSPF 10 2                 D 11.11.11.2     Vlanif300
                       127.0.0.0/8  Direct 0 0              D 127.0.0.1    InLoopBack0
                       127.0.0.1/32 Direct 0 0               D 127.0.0.1    InLoopBack0
                    127.255.255.255/32 Direct 0 0              D 127.0.0.1      InLoopBack0
                    255.255.255.255/32 Direct 0 0              D 127.0.0.1      InLoopBack0

                    [PE1] display ospf peer
                    (M) Indicates MADJ neighbor
                           OSPF Process 1 with Router ID 1.1.1.9

