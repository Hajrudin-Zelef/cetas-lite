---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-278
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [40845, 41000]
sha256: 366882d1029a4be667277e3231ba41c5086302a98e5bf7f9ebb77bbf5a6a1c8f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VPLS to transparently transmit Layer 2 packets over the backbone
                         network to implement Layer 2 communication between Site1, Site2, and Site3
                         and to retain user information in Layer 2 packets when the packets are
                         transmitted over the backbone network.
                    2.   Configure BGP AD VPLS to implement Layer 2 communication between CEs
                         on the enterprise network with many sites and a complex network
                         environment.
                    3.   Configure an IGP on the backbone network for data transmission between
                         PEs on the public network.
                    4.   Configure basic MPLS functions and LDP on devices on the backbone network
                         to implement VPLS.
                    5.   Establish tunnels for transmitting data between PEs to prevent data from
                         being accessed by the public network.
                    6.   Enable MPLS L2VPN on PEs to implement VPLS.
                    7.   Enable PEs to function as BGP peers to exchange VPLS information, create
                         VSIs on the PEs, specify BGP as the signaling protocol, specify the VPLS ID and
                         VPN targets, and bind AC interfaces to the VSIs to implement BGP AD VPLS.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           656
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

                    The configurations of PE1, PE2, PE3, CE2, and CE3 are similar to the configuration
                    of CE1. For detailed configurations, see Configuration Scripts.

                          NOTE

                         Do not add AC-side and PW-side physical interfaces on a PE to the same VLAN. Otherwise,
                         a loop may occur.

         Step 2 Configure a routing protocol for communication between devices.
                    In this example, OSPF is configured. When configuring OSPF, configure PE1, PE2,
                    and PE3 to advertise their loopback interfaces' 32-bit IP addresses (used as LSR
                    IDs).
                    # Configure PE1.
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [PE1-LoopBack1] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 172.16.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 172.16.2.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    The configurations of PE2 and PE3 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display ip routing-table command
                    on PE1, PE2, and PE3. The command outputs show that PE1, PE2, and PE3 have
                    learned routes from each other.
         Step 3 Configure basic MPLS functions and LDP.
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] mpls
                    [PE1-Vlanif20] mpls ldp


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                 657
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration

                    [PE1-Vlanif20] quit
                    [PE1] interface vlanif 30
                    [PE1-Vlanif30] mpls
                    [PE1-Vlanif30] mpls ldp
                    [PE1-Vlanif30] quit

                    The configurations of PE2 and PE3 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.

                    After the configurations are complete, run the display mpls ldp peer command
                    on PE1, PE2, and PE3. The command outputs show that LDP peer relationships
                    have been established between PE1 and PE2, between PE1 and PE3, and between
                    PE2 and PE3. Run the display mpls ldp session command on PE1, PE2, and PE3.
                    The command outputs show that LDP sessions have been established between
                    them. Run the display mpls lsp command. The command output shows LSP
                    establishment information.

         Step 4 Enable BGP peers to exchange VPLS member information.

                    # Configure PE1.
                    [PE1] bgp 100
                    [PE1-bgp] peer 2.2.2.9 as-number 100
                    [PE1-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [PE1-bgp] peer 3.3.3.9 as-number 100
                    [PE1-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [PE1-bgp] l2vpn-ad-family
                    [PE1-bgp-af-l2vpn-ad] peer 2.2.2.9 enable
                    [PE1-bgp-af-l2vpn-ad] peer 3.3.3.9 enable
                    [PE1-bgp-af-l2vpn-ad] quit
                    [PE1-bgp] quit

                    # Configure PE2.
                    [PE2] bgp 100
                    [PE2-bgp] peer 1.1.1.9 as-number 100
                    [PE2-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [PE2-bgp] peer 3.3.3.9 as-number 100
                    [PE2-bgp] peer 3.3.3.9 connect-interface loopback 1
                    [PE2-bgp] l2vpn-ad-family
                    [PE2-bgp-af-l2vpn-ad] peer 1.1.1.9 enable
                    [PE2-bgp-af-l2vpn-ad] peer 3.3.3.9 enable
                    [PE2-bgp-af-l2vpn-ad] quit
                    [PE2-bgp] quit

                    # Configure PE3.
                    [PE3] bgp 100
                    [PE3-bgp] peer 1.1.1.9 as-number 100
                    [PE3-bgp] peer 1.1.1.9 connect-interface loopback 1
                    [PE3-bgp] peer 2.2.2.9 as-number 100
                    [PE3-bgp] peer 2.2.2.9 connect-interface loopback 1
                    [PE3-bgp] l2vpn-ad-family
                    [PE3-bgp-af-l2vpn-ad] peer 1.1.1.9 enable
                    [PE3-bgp-af-l2vpn-ad] peer 2.2.2.9 enable
                    [PE3-bgp-af-l2vpn-ad] quit
                    [PE3-bgp] quit

         Step 5 Enable MPLS L2VPN.

                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit

                    # Configure PE2.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                         658

