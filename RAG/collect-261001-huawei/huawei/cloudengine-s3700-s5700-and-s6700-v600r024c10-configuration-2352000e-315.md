---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-315
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [46804, 46961]
sha256: 2a9bb7c8fe6b45d6bb04b5e5599f05c9d18839119ac16b6b672242e65cc62708
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    3.   Set the VSI attribute of the AC interface to hub.
                         hub-mode enable

                         By default, the VSI attribute of an AC interface is spoke.

                    ----End

6.15.3 Example for Configuring VPLS Service Isolation
Networking Requirements
                    Figure 6-41 shows a backbone network built by an enterprise. Site1 connects to
                    the backbone network by connecting CE1 to PE1, while Site2 connects to the
                    backbone network by connecting CE2, CE3, and CE4 to PE2. LDP VPLS needs to be
                    configured between PE1 and PE2 to implement Layer 2 service interworking. CE2,
                    CE3, and CE4 are used to connect different user services to the network. The
                    enterprise requires forwarding isolation between CE3 and CE4, while allowing
                    communication between CE2 and CE3 and between CE2 and CE4.

                    Figure 6-41 Network diagram of configuring VPLS service isolation
                          NOTE

                         In this example, interface1, interface2, interface3, interface4 and interface5 represent
                         10GE1/0/1, 10GE1/0/2, 10GE1/0/3, 10GE1/0/4, and 10GE1/0/5, respectively.




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure VLANs to which interfaces belong and assign IP addresses to the
                         corresponding VLANIF interfaces.

Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                        753
VPN Configuration
VPN Configuration                                                                  6 VPLS Configuration


                    2.   Configure OSPF.
                    3.   Configure MPLS LDP.
                    4.   Establish a remote MPLS LDP session.
                    5.   Configure LDP VPLS.
                    6.   Configure VPLS service isolation to implement forwarding isolation between
                         CE3 and CE4, while allowing communication between CE2 and CE3 and
                         between CE2 and CE4.

Procedure
         Step 1 Configure VLANs to which interfaces belong and assign IP addresses to the
                corresponding VLANIF interfaces.
                    # Configure CE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] vlan batch 100
                    [CE2] interface vlanif 100
                    [CE2-Vlanif100] ip address 10.1.1.2 255.255.255.0
                    [CE2-Vlanif100] quit
                    [CE2] interface 10ge 1/0/1
                    [CE2-10GE1/0/1] port link-type trunk
                    [CE2-10GE1/0/1] port trunk allow-pass vlan 100
                    [CE2-10GE1/0/1] quit

                    The configurations of CE1, CE3, and CE4 are similar to the configuration of CE2.
                    For detailed configurations, see Configuration Scripts.
                    # Configure the switch.
                    <HUAWEI> system-view
                    [HUAWEI] sysname Switch
                    [Switch] vlan batch 100 200 300
                    [Switch] interface 10ge 1/0/1
                    [Switch-10GE1/0/1] port link-type trunk
                    [Switch-10GE1/0/1] port trunk allow-pass vlan 100 200 300
                    [Switch-10GE1/0/1] quit
                    [Switch] interface 10ge 1/0/3
                    [Switch-10GE1/0/3] port link-type trunk
                    [Switch-10GE1/0/3] port trunk allow-pass vlan 100
                    [Switch-10GE1/0/3] quit
                    [Switch] interface 10ge 1/0/4
                    [Switch-10GE1/0/4] port link-type trunk
                    [Switch-10GE1/0/4] port trunk allow-pass vlan 200
                    [Switch-10GE1/0/4] quit
                    [Switch] interface 10ge 1/0/5
                    [Switch-10GE1/0/5] port link-type trunk
                    [Switch-10GE1/0/5] port trunk allow-pass vlan 300
                    [Switch-10GE1/0/5] quit

                    # Configure PE2.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE2
                    [PE2] vlan batch 30 100 200 300
                    [PE2] interface vlanif 30
                    [PE2-Vlanif30] ip address 9.1.1.2 255.255.255.0
                    [PE2-Vlanif30] quit
                    [PE2] interface 10ge 1/0/1
                    [PE2-10GE1/0/1] port link-type trunk
                    [PE2-10GE1/0/1] port trunk allow-pass vlan 30
                    [PE2-10GE1/0/1] quit
                    [PE2] interface 10ge 1/0/2


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                       754
VPN Configuration
VPN Configuration                                                                         6 VPLS Configuration

                    [PE2-10GE1/0/2] port link-type trunk
                    [PE2-10GE1/0/2] port trunk allow-pass vlan 100 200 300
                    [PE2-10GE1/0/2] quit

                    The configurations of PE1 and the P are similar to the configuration of PE2. For
                    detailed configurations, see Configuration Scripts.

                          NOTE

                         Do not add AC-side and PW-side physical interfaces on a PE to the same VLAN. Otherwise,
                         a loop may occur.

         Step 2 Configure a routing protocol for communication between devices.
                    OSPF is used as an example.
                    # Configure PE1.
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [PE1-LoopBack1] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] network 8.1.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display ip routing-table command
                    on PE1, PE2, and the P. The command outputs show that these devices have
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
                    [PE1-Vlanif20] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
                    After the configurations are complete, run the display mpls ldp session command
                    on PE1, PE2, and the P. The command outputs show that the status of the peer
                    relationship between PE1 and the P and between PE2 and the P is Operational,
                    indicating that the peer relationships have been established. Run the display mpls
                    lsp command. The command output shows LSP establishment information.
         Step 4 Establish a remote LDP session between PEs.
                    # Configure PE1.
                    [PE1] mpls ldp remote-peer 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] remote-ip 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] quit


