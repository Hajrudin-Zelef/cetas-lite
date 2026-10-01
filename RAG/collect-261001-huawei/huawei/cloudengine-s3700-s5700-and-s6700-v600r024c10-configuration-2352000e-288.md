---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-288
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [42409, 42582]
sha256: 12492629f1ecf175b9fa60ef20023b7d4a6eda39ea1bcfd0f7e02a8a59b9adc0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                     681
VPN Configuration
VPN Configuration                                                                                            6 VPLS Configuration


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
                    0000.0000.0002 Vlanif20               0000.0000.0001.01 Up 27s             L1(L1L2) 64
                    0000.0000.0002 Vlanif20               0000.0000.0001.01 Up 27s             L2(L1L2) 64

                    Total Peer(s): 2


Issue 01 (2025-03-03)                  Copyright © Huawei Technologies Co., Ltd.                                             682
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration


                    The command output shows that the IS-IS neighbor relationship has been
                    established between the ASBR_PE and the PE in the local AS, and the neighbor
                    relationship is up.

                    # Perform a ping test to check the connectivity.
                    [PE1] ping 2.2.2.2
                     PING 2.2.2.2: 56 data bytes, press CTRL_C to break
                       Reply from 2.2.2.2: bytes=56 Sequence=1 ttl=255 time=180 ms
                       Reply from 2.2.2.2: bytes=56 Sequence=2 ttl=255 time=90 ms
                       Reply from 2.2.2.2: bytes=56 Sequence=3 ttl=255 time=60 ms
                       Reply from 2.2.2.2: bytes=56 Sequence=4 ttl=255 time=60 ms
                       Reply from 2.2.2.2: bytes=56 Sequence=5 ttl=255 time=100 ms

                     --- 2.2.2.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 60/98/180 ms

                    The ASBR_PE and PE in the same AS can ping each other successfully.

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
                    [PE1-Vlanif20] mpls ldp
                    [PE1-Vlanif20] quit

                    The configurations of ASBR_PE1, ASBR_PE2, and PE2 are similar to the
                    configuration of PE1. For detailed configurations, see Configuration Scripts.

                    # Check LDP session information on ASBR_PE1.
                    [ASBR_PE1] display mpls ldp session

                    LDP Session(s) in Public Network
                    Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                    A '*' before a session means the session is being deleted.
                    ------------------------------------------------------------------------------
                    PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                    ------------------------------------------------------------------------------
                    1.1.1.1:0        Operational DU Active 0000:00:08 34/34
                    ------------------------------------------------------------------------------
                    TOTAL: 1 session(s) Found.

         Step 4 Enable MPLS L2VPN on PEs.

                    # Configure PE1.
                    [PE1] mpls l2vpn
                    [PE1-l2vpn] quit

                    # Configure PE2.
                    [PE2] mpls l2vpn
                    [PE2-l2vpn] quit


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                                       683
VPN Configuration
VPN Configuration                                                                 6 VPLS Configuration


                    # Configure ASBR_PE1.
                    [ASBR_PE1] mpls l2vpn
                    [ASBR_PE1-l2vpn] quit

                    # Configure ASBR_PE2.
                    [ASBR_PE2] mpls l2vpn
                    [ASBR_PE2-l2vpn] quit

         Step 5 Configure LDP VPLS and bind VSIs to interfaces.
                    # Configure PE1.
                    [PE1] vsi a1 static
                    [PE1-vsi-a1] pwsignal ldp
                    [PE1-vsi-a1-ldp] vsi-id 2
                    [PE1-vsi-a1-ldp] peer 2.2.2.2
                    [PE1-vsi-a1-ldp] quit
                    [PE1-vsi-a1] quit
                    [PE1] interface vlanif 10
                    [PE1-Vlanif10] l2 binding vsi a1
                    [PE1-Vlanif10] quit

                    # Configure ASBR_PE1.
                    [ASBR_PE1] vsi a1 static
                    [ASBR_PE1-vsi-a1] pwsignal ldp
                    [ASBR_PE1-vsi-a1-ldp] vsi-id 2
                    [ASBR_PE1-vsi-a1-ldp] peer 1.1.1.1
                    [ASBR_PE1-vsi-a1-ldp] quit
                    [ASBR_PE1-vsi-a1] quit
                    [ASBR_PE1] interface vlanif 30
                    [ASBR_PE1-Vlanif30] l2 binding vsi a1
                    [ASBR_PE1-Vlanif30] quit

