---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-291
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [42959, 43101]
sha256: bfaa16f29890d6380fd6eb7c48121e8b46aa8ce5e621587b770e58eb3bbf4226
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPLS to transparently transmit Layer 2 packets on the backbone
                         network to implement Layer 2 communication between Site1 and Site2 and
                         to retain user information in Layer 2 packets when the packets are
                         transmitted over the backbone network.
                    2.   Use BGP VPLS to implement Layer 2 communication between CEs because
                         the network environments of the branch sites are unstable.
                    3.   Configure an IGP on the backbone network to implement communication
                         between devices within an AS on the public network.
                    4.   Configure basic MPLS functions and LDP on PEs on the backbone network to
                         implement VPLS.
                    5.   Establish tunnels for transmitting data between PEs within an AS to prevent
                         data from being accessed by the public network.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Enable PEs within an AS to function as BGP peers to exchange VPLS
                         information, create VSIs on PEs, specify BGP as the signaling protocol, specify
                         the RD, VPN targets, and site ID, and bind AC interfaces to the VSIs to
                         implement BGP VPLS.
                    8.   Configure the peer ASBR as a CE on ASBR_PEs, and bind VSIs to peer
                         interfaces to implement inter-AS VPLS OptionA.

Procedure
         Step 1 Configure VLANs to which interfaces belong and assign IP addresses to the
                corresponding VLANIF interfaces.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan 10
                    [CE1-vlan10] quit
                    [CE1] interface vlanif 10
                    [CE1-Vlanif10] ip address 10.1.1.1 255.255.255.0
                    [CE1-Vlanif10] quit
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] quit

                    The configurations of PE1, PE2, ASBR_PE1, ASBR_PE2, and CE2 are similar to the
                    configuration of CE1. For detailed configurations, see Configuration Scripts.

                          NOTE

                         Do not add AC-side and PW-side physical interfaces on a PE to the same VLAN. Otherwise,
                         a loop may occur.

         Step 2 Configure a routing protocol for communication between devices.
                    Configure an IGP to achieve connectivity between PEs and ASBR_PEs on the MPLS
                    backbone network. In this example, IS-IS is configured.
                    # Configure PE1.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 690
VPN Configuration
VPN Configuration                                                                                           6 VPLS Configuration

                    [PE1] isis 1
                    [PE1-isis-1] network-entity 10.0000.0000.0001.00
                    [PE1-isis-1] quit
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.1 255.255.255.255
                    [PE1-LoopBack1] isis enable 1
                    [PE1-LoopBack1] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] isis enable 1
                    [PE1-Vlanif20] quit

                    The configurations of ASBR_PE1, ASBR_PE2, and PE2 are similar to the
                    configuration of PE1. For detailed configurations, see Configuration Scripts.
                    After the configurations are complete, an IS-IS neighbor relationship can be
                    established between the ASBR_PE and PE in the same AS.
                    # Check IS-IS neighbor information on PE1.
                    [PE1] display isis peer

                                        Peer information for ISIS(1)

                      System Id      Interface         Circuit Id      State HoldTime Type         PRI
                    -------------------------------------------------------------------------------
                    0000.0000.0002 Vlanif20               0000.0000.0002.01 Up 8s             L1(L1L2) 64
                    0000.0000.0002 Vlanif20               0000.0000.0002.01 Up 8s             L2(L1L2) 64

                    Total Peer(s): 2

                    The command output shows that an IS-IS neighbor relationship has been
                    established between the ASBR_PE and PE in the same AS, and the neighbor status
                    is up.
                    # Perform a ping test to check the connectivity.
                    [ASBR_PE1] ping 1.1.1.1
                     PING 1.1.1.1: 56 data bytes, press CTRL_C to break
                      Reply from 1.1.1.1: bytes=56 Sequence=1 ttl=255 time=47 ms
                      Reply from 1.1.1.1: bytes=56 Sequence=2 ttl=255 time=31 ms
                      Reply from 1.1.1.1: bytes=56 Sequence=3 ttl=255 time=31 ms
                      Reply from 1.1.1.1: bytes=56 Sequence=4 ttl=255 time=31 ms
                      Reply from 1.1.1.1: bytes=56 Sequence=5 ttl=255 time=31 ms

                     --- 1.1.1.1 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 31/34/47 ms

                    The ASBR_PE and PE in the same AS can ping Loopback1 of each other
                    successfully.
         Step 3 Configure basic MPLS functions and LDP.
                    Configure basic MPLS functions on the MPLS backbone network. Establish a
                    dynamic LDP LSP between the PE and ASBR_PE in the same AS.
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] mpls


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                            691
VPN Configuration
VPN Configuration                                                                                     6 VPLS Configuration

                    [PE1-Vlanif20] mpls ldp
                    [PE1-Vlanif20] quit

                    The configurations of ASBR_PE1, ASBR_PE2, and PE2 are similar to the
                    configuration of PE1. For detailed configurations, see Configuration Scripts.

                    # Check LDP LSP information on PE1.
                    [PE1] display mpls lsp
                     Flag after Out IF: (I) - LSP Is Only Iterated by RLFA
                    -------------------------------------------------------------------------------
                                 LSP Information: LDP LSP
                    -------------------------------------------------------------------------------
                    FEC              In/Out Label In/Out IF                      Vrf Name
                    1.1.1.1/32         3/NULL         -/-
                    2.2.2.2/32         NULL/3         -/Vlanif20
                    2.2.2.2/32         1025/3        -/Vlanif20

                    The command output shows that an LSP has been established between the PE
                    and ASBR_PE in the same AS.

         Step 4 Establish an MP-IBGP connection within an AS.

