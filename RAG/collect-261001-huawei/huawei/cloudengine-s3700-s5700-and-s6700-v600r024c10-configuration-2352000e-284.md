---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-284
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [41723, 41887]
sha256: c16642fbde5469776dce20ef8ecd40ba732a7592a4db549903d83b19d10bcbd8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        In this example, interface1, interface2, and interface3 represent 10GE1/0/1, 10GE1/0/2, and
                        10GE1/0/3, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     670
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration


Configuration Roadmap
                    The configuration roadmap is as follows:

                    1.   Configure VPLS to transparently transmit Layer 2 packets over the backbone
                         network to implement Layer 2 communication between Site1, Site2, and Site3
                         and to retain user information in Layer 2 packets when the packets are
                         transmitted over the backbone network.
                    2.   Use LDP HVPLS to form a layered network topology and implement Layer 2
                         communication between CEs.
                    3.   Configure an IGP on the backbone network for data transmission between
                         PEs on the public network.
                    4.   Configure basic MPLS functions and LDP on devices on the backbone network
                         to implement VPLS.
                    5.   Establish tunnels for transmitting data between PEs to prevent data from
                         being accessed by the public network.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Create VSIs on PEs, configure LDP as the signaling protocol, and bind VSIs to
                         AC interfaces on the UPE and PE1 to implement LDP VPLS.
                    8.   Specify the UPE as the under-layer PE and PE1 as a VSI peer on the SPE, and
                         specify the SPE as a VSI peer on both the UPE and PE1 to implement HVPLS.

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

                    The configurations of the UPE, SPE, PE1, CE2, and CE3 are similar to the
                    configuration of CE1. For detailed configurations, see Configuration Scripts.

                          NOTE

                         Do not add AC-side and PW-side physical interfaces on a PE to the same VLAN. Otherwise,
                         a loop may occur.

         Step 2 Configure a routing protocol for communication between devices.

                    In this example, OSPF is configured. When configuring OSPF, configure the UPE,
                    SPE, and PE1 to advertise their loopback interfaces' 32-bit IP addresses (used as
                    LSR IDs).

                    # Configure the UPE.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 671
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration

                    [UPE] interface loopback 1
                    [UPE-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [UPE-LoopBack1] quit
                    [UPE] ospf 1
                    [UPE-ospf-1] area 0.0.0.0
                    [UPE-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [UPE-ospf-1-area-0.0.0.0] network 192.0.2.0 0.0.0.255
                    [UPE-ospf-1-area-0.0.0.0] quit
                    [UPE-ospf-1] quit

                    The configurations of the SPE and PE1 are similar to the configuration of the UPE.
                    For detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display ip routing-table command
                    on the UPE, SPE, and PE1. The command outputs show that the UPE, SPE, and PE1
                    have learned each other's loopback interface address.
         Step 3 Configure basic MPLS functions and LDP.
                    # Configure the UPE.
                    [UPE] mpls lsr-id 1.1.1.9
                    [UPE] mpls
                    [UPE-mpls] quit
                    [UPE] mpls ldp
                    [UPE-mpls-ldp] quit
                    [UPE] interface vlanif 30
                    [UPE-Vlanif30] mpls
                    [UPE-Vlanif30] mpls ldp
                    [UPE-Vlanif30] quit

                    The configurations of the SPE and PE1 are similar to the configuration of the UPE.
                    For detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display mpls ldp session command
                    on the UPE, SPE, and PE1. The command outputs show that the status of the peer
                    relationship between the UPE and SPE and between PE1 and the SPE is
                    Operational, which indicates that the peer relationship has been established. Run
                    the display mpls lsp command. The command output shows LSP establishment
                    information.
         Step 4 Enable MPLS L2VPN and configure a VSI.
                    # Configure the UPE.
                    [UPE] mpls l2vpn
                    [UPE-l2vpn] quit
                    [UPE] vsi v123 static
                    [UPE-vsi-v123] pwsignal ldp
                    [UPE-vsi-v123-ldp] vsi-id 123
                    [UPE-vsi-v123-ldp] peer 2.2.2.9
                    [UPE-vsi-v123-ldp] quit
                    [UPE-vsi-v123] quit

                    # Configure the SPE.
                    [SPE] mpls l2vpn
                    [SPE-l2vpn] quit
                    [SPE] vsi v123 static
                    [SPE-vsi-v123] pwsignal ldp
                    [SPE-vsi-v123-ldp] vsi-id 123
                    [SPE-vsi-v123-ldp] peer 3.3.3.9
                    [SPE-vsi-v123-ldp] peer 1.1.1.9 upe
                    [SPE-vsi-v123-ldp] quit
                    [SPE-vsi-v123] quit

                    # Configure PE1.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       672
VPN Configuration
VPN Configuration                                                                     6 VPLS Configuration

                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] vsi v123 static
                    [PE1-vsi-v123] pwsignal ldp
                    [PE1-vsi-v123-ldp] vsi-id 123
                    [PE1-vsi-v123-ldp] peer 2.2.2.9
                    [PE1-vsi-v123-ldp] quit
                    [PE1-vsi-v123] quit

         Step 5 Bind the VSI to an interface.

                    # Configure the UPE.
                    [UPE] interface vlanif 10
                    [UPE-Vlanif10] l2 binding vsi v123
                    [UPE-Vlanif10] quit
                    [UPE] interface vlanif 20
                    [UPE-Vlanif20] l2 binding vsi v123
                    [UPE-Vlanif20] quit

                    # Configure PE1.
                    [PE1] interface vlanif 50
                    [PE1-Vlanif50] l2 binding vsi v123
                    [PE1-Vlanif50] quit

                    ----End

Verifying the Configuration
                    # Check detailed information about the VSI named v123 on the SPE.
                    [SPE] display vsi name v123 verbose

