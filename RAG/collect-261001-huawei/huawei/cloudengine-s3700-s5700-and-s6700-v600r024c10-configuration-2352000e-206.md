---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-206
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [30236, 30388]
sha256: 806fa777612825aec85238a0aa1c9531486864df335b92236e4e4615c01c92ee
---

                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 10
                    [CE1] interface vlanif 10
                    [CE1-Vlanif10] ip address 10.10.1.1 255.255.255.0
                    [CE1-Vlanif10] quit
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] quit

                    The configuration of CE2 is similar to that of CE1. For detailed configurations, see
                    Configuration Scripts.

                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] vlan batch 10 20
                    [PE1] interface 10ge 1/0/1
                    [PE1-10GE1/0/1] port link-type trunk
                    [PE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [PE1-10GE1/0/1] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [PE1-10GE1/0/2] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] ip address 10.1.1.1 255.255.255.0
                    [PE1-Vlanif20] quit
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                    [PE1-LoopBack1] quit
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit

                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.

                    The configurations of CE2, PE1, PE2, and the P are similar to the configuration of
                    CE1. For detailed configurations, see Configuration Scripts.

         Step 2 Configure an IGP on the MPLS backbone network. In this example, OSPF is used.

                    When configuring OSPF, configure PE1, PE2, and the P to advertise their 32-bit IP
                    addresses (used as LSR IDs) of loopback interfaces.

                    # Configure PE1.
                    [PE1] ospf 1
                    [PE1-ospf-1] area 0
                    [PE1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                    [PE1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                    [PE1-ospf-1-area-0.0.0.0] quit
                    [PE1-ospf-1] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                           484
VPN Configuration
VPN Configuration                                                                                     5 VPWS Configuration


                    The configurations of the P and PE2 are similar to the configuration of PE1. For
                    detailed configurations, see Configuration Scripts.
         Step 3 Configure basic MPLS functions and LDP on the MPLS backbone network, and
                establish LDP LSPs.
                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.9
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] mpls ldp remote-peer 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] remote-ip 3.3.3.9
                    [PE1-mpls-ldp-remote-3.3.3.9] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] mpls
                    [PE1-Vlanif20] mpls ldp
                    [PE1-Vlanif20] quit

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

                    # Configure PE2.
                    [PE2] mpls lsr-id 3.3.3.9
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] mpls ldp remote-peer 1.1.1.9
                    [PE2-mpls-ldp-remote-1.1.1.9] remote-ip 1.1.1.9
                    [PE2-mpls-ldp-remote-1.1.1.9] quit
                    [PE2] interface vlanif 10
                    [PE2-Vlanif10] mpls
                    [PE2-Vlanif10] mpls ldp
                    [PE2-Vlanif10] quit

                    After the configurations are complete, LDP sessions are established between PE1,
                    P, and PE2. Run the display mpls ldp session command. The command output
                    shows that the session status is Operational.
                    The following example uses the command output on PE1.
                    <PE1> display mpls ldp session
                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                     ------------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                     ------------------------------------------------------------------------------
                     2.2.2.9:0        Operational DU Passive 000:02:22 572/572
                     3.3.3.9:0        Operational DU Passive 000:02:21 566/566
                     ------------------------------------------------------------------------------
                     TOTAL: 2 session(s) Found.


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       485
VPN Configuration
VPN Configuration                                                                                5 VPWS Configuration


         Step 4 Enable MPLS L2VPN and create static VCs on PEs.
                    # Configure PE1: Create a static VC on VLANIF 10 connected to CE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] mpls static-l2vc destination 3.3.3.9 transmit-vpn-label 100 receive-vpn-label 200
                    [PE1-Vlanif10] quit

                    # Configure PE2: Create a static VC on VLANIF 20 connected to CE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit
                    [PE2] interface vlanif 20
                    [PE2-Vlanif20] mpls static-l2vc destination 1.1.1.9 transmit-vpn-label 200 receive-vpn-label 100
                    [PE2-Vlanif20] quit

                    ----End

