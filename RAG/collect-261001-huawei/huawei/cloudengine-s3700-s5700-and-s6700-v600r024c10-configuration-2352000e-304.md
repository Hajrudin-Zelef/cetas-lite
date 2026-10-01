---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-304
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [45027, 45176]
sha256: 7394043875e877d57b9dd430e87040a44f4e13656666703770854e72c6c2b367
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

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
                    4.   Configure basic MPLS functions and LDP on SPEs and the P to implement
                         VPLS.
                    5.   Establish tunnels for transmitting data between PEs, including dynamic LSPs
                         between SPEs and static LSPs between UPEs and SPEs, to prevent data from
                         being accessed by the public network.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Create SVC VPWS connections between UPEs and SPEs, configure static VPWS
                         and VSIs on SPEs, enable the MAC Withdraw function for VSIs, and configure
                         static VPWS on UPEs for accessing SPEs to implement static VPWS accessing
                         VPLS.
                    8.   Create VSIs on SPEs, configure LDP as the signaling protocol, and bind VSIs to
                         AC interfaces to implement LDP VPLS.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           724
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


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
                    [SPE1-ospf-1-area-0.0.0.0] network 1.1.1.0 0.0.0.255
                    [SPE1-ospf-1-area-0.0.0.0] network 3.1.1.0 0.0.0.255
                    [SPE1-ospf-1-area-0.0.0.0] quit
                    [SPE1-ospf-1] quit

                    The configurations of UPE1, UPE2, SPE2, and the P are similar to the configuration
                    of SPE1. For detailed configurations, see Configuration Scripts.

                    After the configurations are complete, run the display ip routing-table command
                    on PE1, PE2, and the P. The command outputs show that PE1, PE2, and the P have
                    learned routes from each other.

         Step 3 Configure basic MPLS functions and LDP.

                    # Configure SPE1.
                    [SPE1] mpls lsr-id 1.1.1.9
                    [SPE1] mpls
                    [SPE1-mpls] quit
                    [SPE1] mpls ldp
                    [SPE1-mpls-ldp] quit
                    [SPE1] interface vlanif 30
                    [SPE1-Vlanif30] mpls
                    [SPE1-Vlanif30] mpls ldp
                    [SPE1-Vlanif30] quit
                    [SPE1] interface vlanif 20


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 725
VPN Configuration
VPN Configuration                                                                                     6 VPLS Configuration

                    [SPE1-Vlanif20] mpls
                    [SPE1-Vlanif20] quit

                    # Configure the P.
                    [P] mpls lsr-id 2.2.2.9
                    [P] mpls
                    [P-mpls] quit
                    [P] mpls ldp
                    [P-mpls-ldp] quit
                    [P] interface vlanif 30
                    [P-Vlanif30] mpls
                    [P-Vlanif30] mpls ldp
                    [P-Vlanif30] quit
                    [P] interface vlanif 40
                    [P-Vlanif40] mpls
                    [P-Vlanif40] mpls ldp
                    [P-Vlanif40] quit

                    # Configure SPE2.
                    [SPE2] mpls lsr-id 3.3.3.9
                    [SPE2] mpls
                    [SPE2-mpls] quit
                    [SPE2] mpls ldp
                    [SPE2-mpls-ldp] quit
                    [SPE2] interface vlanif 40
                    [SPE2-Vlanif40] mpls
                    [SPE2-Vlanif40] mpls ldp
                    [SPE2-Vlanif40] quit
                    [SPE2] interface vlanif 50
                    [SPE2-Vlanif50] mpls
                    [SPE2-Vlanif50] quit

                    # Check LDP session information on SPE1.
                    [SPE1] display mpls ldp session

                    LDP Session(s) in Public Network
                    Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                    A '*' before a session means the session is being deleted.
                    ------------------------------------------------------------------------------
                    PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                    ------------------------------------------------------------------------------
                    2.2.2.9:0        Operational DU Passive 0000:00:00 2/2
                    ------------------------------------------------------------------------------
                    TOTAL: 1 session(s) Found.

                    After the configurations are complete, run the display mpls ldp session command
                    on SPE1, SPE2, and the P. The command outputs show that Status of the peer
                    relationships between SPE1 and the P and between SPE2 and the P is
                    Operational, indicating that the peer relationships have been established.

