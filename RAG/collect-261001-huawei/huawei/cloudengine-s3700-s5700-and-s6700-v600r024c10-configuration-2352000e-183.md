---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-183
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [26578, 26748]
sha256: 4ef0db07d64c168f60d5d777fe4649857351f9001fb3c6207f5bedfe2f9dcf4d
---

                    # CE1 can ping CE2 successfully.
                    <CE1> ping 10.1.1.2
                     PING 10.1.1.2: 56 data bytes, press CTRL_C to break
                      Reply from 10.1.1.2: bytes=56 Sequence=1 ttl=255 time=4 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=2 ttl=255 time=2 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=3 ttl=255 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=4 ttl=255 time=1 ms
                      Reply from 10.1.1.2: bytes=56 Sequence=5 ttl=255 time=1 ms

                     --- 10.1.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 1/1/4 ms


Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan batch 10
                          #
                          interface Vlanif10
                           ip address 10.1.1.1 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 10
                          #
                          return

                    ●     PE
                          #
                          sysname PE
                          #
                          vlan batch 10 20
                          #
                          mpls lsr-id 1.1.1.9


Issue 01 (2025-03-03)                Copyright © Huawei Technologies Co., Ltd.                      422
VPN Configuration
VPN Configuration                                                                          5 VPWS Configuration

                        #
                        mpls
                        #
                        mpls l2vpn
                        #
                        interface Vlanif10
                        #
                        interface Vlanif20
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ccc ce1-ce2 interface Vlanif10 out-interface Vlanif20
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 20
                        #
                        interface Vlanif20
                         ip address 10.1.1.2 255.255.255.0
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        return


5.5.6 Example for Configuring a Remote CCC Connection
Networking Requirements
                    In Figure 5-6, CE1 and CE2 are connected to different PEs. To allow the two CEs to
                    communicate, a remote CCC connection can be established. When establishing the
                    remote CCC connection, configure two static CR-LSPs on the P to transmit packets
                    in both directions.

                    Figure 5-6 Network diagram of configuring a remote CCC connection
                         NOTE

                        In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                    423
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration




Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure bidirectional static CR-LSPs to be used for the CCC connection
                         between PEs only.
                    2.   Enable MPLS L2VPN on PEs but not the P.
                    3.   Configure two connections, one from CE1 to CE2, and the other from CE2 to
                         CE1.
                               NOTE

                         By default, the Link-type Negotiation Protocol (LNP) is enabled globally on the device. If a
                         VLANIF interface is used as an AC interface for L2VPN, the configuration conflicts with LNP.
                         In this case, run the lnp disable command in the system view to disable LNP.


Procedure
         Step 1 Configure the VLANs that interfaces belong to and assign IP addresses to VLANIF
                interfaces.
                    # Configure CE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname CE1
                    [CE1] vlan batch 10
                    [CE1] interface 10ge 1/0/1
                    [CE1-10GE1/0/1] port link-type trunk
                    [CE1-10GE1/0/1] port trunk allow-pass vlan 10
                    [CE1-10GE1/0/1] quit
                    [CE1] interface vlanif 10
                    [CE1-Vlanif10] ip address 10.10.1.1 24
                    [CE1-Vlanif10] quit

                    # Configure CE2.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                      424
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration

                    <HUAWEI> system-view
                    [HUAWEI] sysname CE2
                    [CE2] vlan batch 20
                    [CE2] interface 10ge 1/0/2
                    [CE2-10GE1/0/2] port link-type trunk
                    [CE2-10GE1/0/2] port trunk allow-pass vlan 20
                    [CE2-10GE1/0/2] quit
                    [CE2] interface vlanif 20
                    [CE2-Vlanif20] ip address 10.10.1.2 24
                    [CE2-Vlanif20] quit

                    # Configure PE1.
                    <HUAWEI> system-view
                    [HUAWEI] sysname PE1
                    [PE1] vlan batch 10 20
                    [PE1] interface loopback 1
                    [PE1-LoopBack1] ip address 1.1.1.9 32
                    [PE1-LoopBack1] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] port link-type trunk
                    [PE1-10GE1/0/2] port trunk allow-pass vlan 20
                    [PE1-10GE1/0/2] quit
                    [PE1] interface vlanif 20
                    [PE1-Vlanif20] ip address 10.1.1.1 24
                    [PE1-Vlanif20] quit

