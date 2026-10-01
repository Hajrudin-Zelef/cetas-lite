---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-309
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [45819, 45989]
sha256: cbfea1fc6bd8d6fb87fb9cf5db5f9aeaab18eec876c31bf0db9a62d71205bf14
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPLS to transparently transmit Layer 2 packets on the backbone
                         network to implement Layer 2 communication between Site1 and Site2 and
                         to retain user information in Layer 2 packets when the packets are
                         transmitted over the backbone network.
                    2.   Configure LDP VPLS to implement Layer 2 communication between CEs as
                         planned.
                    3.   Configure an IGP on the UPEs, SPEs, and P to implement data transmission
                         between PEs on the public network.
                    4.   Configure basic MPLS functions and LDP on the UPEs, SPEs, and P to
                         implement VPLS.
                    5.   Establish tunnels for transmitting data between PEs, including dynamic LSPs
                         between SPEs and dynamic LSPs between UPEs and SPEs, to prevent data
                         from being accessed by the public network.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Create LDP VPWS connections between UPEs and SPEs to implement dynamic
                         VPWS accessing VPLS.
                    8.   Create VSIs on SPEs, configure LDP as the signaling protocol, and bind VSIs to
                         AC interfaces to implement LDP VPLS.

Procedure
         Step 1 Configure VLANs to which interfaces belong and assign IP addresses to the
                corresponding VLANIF interfaces.
                    # Configure CE1.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           736
VPN Configuration
VPN Configuration                                                                        6 VPLS Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan 50
                    [CE1-vlan10] quit
                    [CE1] interface vlanif 50
                    [CE1-Vlanif10] ip address 10.1.1.1 255.255.255.0
                    [CE1-Vlanif10] quit
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 50
                    [CE1-10GE1/0/1] quit

                    The configurations of UPE1, UPE2, SPE1, SPE2, the P, and CE2 are similar to the
                    configuration of CE1. For detailed configurations, see Configuration Scripts.

                          NOTE

                        Do not add AC-side and PW-side physical interfaces on a PE to the same VLAN. Otherwise,
                        a loop may occur.

         Step 2 Configure a routing protocol for communication between devices.
                    OSPF is used as an example.
                    # Configure SPE1.
                    [SPE1] interface loopback 1
                    [SPE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [SPE1-LoopBack1] quit
                    [SPE1] ospf
                    [SPE1-ospf-1] area 0
                    [SPE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [SPE1-ospf-1-area-0.0.0.0] network 10.2.1.0 0.0.0.255
                    [SPE1-ospf-1-area-0.0.0.0] network 10.2.3.0 0.0.0.255
                    [SPE1-ospf-1-area-0.0.0.0] quit
                    [SPE1-ospf-1] quit

                    The configurations of UPE1, UPE2, SPE2, and the P are similar to the configuration
                    of SPE1. For detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display ip routing-table command
                    on PE1, PE2, and the P. The command outputs show that PE1, PE2, and the P have
                    learned routes from each other.
         Step 3 Configure basic MPLS functions and LDP.
                    # Configure UPE1.
                    [UPE1] mpls lsr-id 4.4.4.9
                    [UPE1] mpls
                    [UPE1-mpls] quit
                    [UPE1] mpls ldp
                    [UPE1-mpls-ldp] quit
                    [UPE1] interface vlanif 30
                    [UPE1-Vlanif30] mpls
                    [UPE1-Vlanif30] mpls ldp
                    [UPE1-Vlanif30] quit

                    # Configure UPE2.
                    [UPE2] mpls lsr-id 5.5.5.9
                    [UPE2] mpls
                    [UPE2-mpls] quit
                    [UPE2] mpls ldp
                    [UPE2-mpls-ldp] quit
                    [UPE2] interface vlanif 40
                    [UPE2-Vlanif40] mpls


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                737
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration

                    [UPE2-Vlanif40] mpls ldp
                    [UPE2-Vlanif40] quit

                    # Configure SPE1.
                    [SPE1] mpls lsr-id 1.1.1.9
                    [SPE1] mpls
                    [SPE1-mpls] quit
                    [SPE1] mpls ldp
                    [SPE1-mpls-ldp] quit
                    [SPE1] interface vlanif 10
                    [SPE1-Vlanif10] mpls
                    [SPE1-Vlanif10] mpls ldp
                    [SPE1-Vlanif10] quit
                    [SPE1] interface vlanif 30
                    [SPE1-Vlanif30] mpls
                    [SPE1-Vlanif30] mpls ldp
                    [SPE1-Vlanif30] quit

                    # Configure the P.
                    [P] mpls lsr-id 2.2.2.9
                    [P] mpls
                    [P-mpls] quit
                    [P] mpls ldp
                    [P-mpls-ldp] quit
                    [P] interface vlanif 10
                    [P-Vlanif10] mpls
                    [P-Vlanif10] mpls ldp
                    [P-Vlanif10] quit
                    [P] interface vlanif 20
                    [P-Vlanif20] mpls
                    [P-Vlanif20] mpls ldp
                    [P-Vlanif20] quit

                    # Configure SPE2.
                    [SPE2] mpls lsr-id 3.3.3.9
                    [SPE2] mpls
                    [SPE2-mpls] quit
                    [SPE2] mpls ldp
                    [SPE2-mpls-ldp] quit
                    [SPE2] interface vlanif 20
                    [SPE2-Vlanif20] mpls
                    [SPE2-Vlanif20] mpls ldp
                    [SPE2-Vlanif20] quit
                    [SPE2] interface vlanif 40
                    [SPE2-Vlanif40] mpls
                    [SPE2-Vlanif40] mpls ldp
                    [SPE2-Vlanif40] quit

                    After the configurations are complete, run the display mpls ldp session command
                    on UPEs, SPEs, and the P. The command outputs show that Status of the peer
                    relationships between SPEs and UPEs and between SPEs and the P is Operational,
                    indicating that the peer relationships have been established. Run the display mpls
                    lsp command. The command output shows LSP establishment information.

         Step 4 Establish a remote LDP session between SPEs.

                    # Configure SPE1.
                    [SPE1] mpls ldp remote-peer 3.3.3.9
                    [SPE1-mpls-ldp-remote-3.3.3.9] remote-ip 3.3.3.9
                    [SPE1-mpls-ldp-remote-3.3.3.9] quit

                    # Configure SPE2.

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                     738
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration

