---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-199
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [29083, 29289]
sha256: 65c8b39333b85c77daad5f2392bc338c917999e8d4c266abc9efe7952a428f28
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                     --- 10.10.1.2 ping statistics ---
                       5 packet(s) transmitted
                       5 packet(s) received
                       0.00% packet loss
                       round-trip min/avg/max = 2/3/6 ms


Configuration Scripts
                    ●    CE1
                         #
                         sysname CE1
                         #
                         vlan batch 20
                         #
                         interface Vlanif20
                          ip address 10.10.1.1 255.255.255.0
                         #
                         interface 10GE1/0/2
                          port link-type trunk
                          port trunk allow-pass vlan 20
                         #
                         return

                    ●    PE1
                         #
                         sysname PE1
                         #
                         vlan batch 10 20
                         #
                         mpls lsr-id 192.168.2.2
                         #
                         mpls
                         #
                         mpls l2vpn
                         #
                         mpls ldp
                         #
                         mpls ldp remote-peer 192.168.3.3
                          remote-ip 192.168.3.3
                         #
                         interface Vlanif10
                          ip address 10.1.1.1 255.255.255.0
                          mpls
                          mpls ldp
                         #


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                       466
VPN Configuration
VPN Configuration                                                             5 VPWS Configuration

                        interface Vlanif20
                         mpls l2vc 192.168.3.3 100
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack0
                         ip address 192.168.2.2 255.255.255.255
                        #
                        ospf 1
                        area 0.0.0.0
                          network 192.168.2.2 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 192.168.4.4
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif10
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif20
                         ip address 10.2.2.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack0
                         ip address 192.168.4.4 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 192.168.4.4 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.2.2.0 0.0.0.255
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 192.168.3.3
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                   467
VPN Configuration
VPN Configuration                                                              5 VPWS Configuration

                        mpls
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 192.168.2.2
                         remote-ip 192.168.2.2
                        #
                        interface Vlanif10
                         mpls l2vc 192.168.2.2 100
                        #
                        interface Vlanif20
                         ip address 10.2.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack0
                         ip address 192.168.3.3 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 192.168.3.3 0.0.0.0
                          network 10.2.2.0 0.0.0.255
                        #
                        return

                    ●   CE2
                        #
                        sysname CE2
                        #
                        vlan batch 10
                        #
                        interface Vlanif10
                         ip address 10.10.1.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        return


5.7.7 Example for Configuring LDP VPWS (Using an MPLS TE
Tunnel)

Networking Requirements
                    In Figure 5-25, CE1 and CE2 belong to the same VPN and access the MPLS
                    backbone network through PE1 and PE2, respectively. OSPF is used as the IGP on
                    the MPLS backbone network.

                    An LDP VPWS connection needs to be configured. The dynamic signaling protocol
                    Resource Reservation Protocol-Traffic Engineering (RSVP-TE) needs to be used to
                    establish an MPLS TE tunnel between PE1 and PE2 to forward VPWS traffic. It is
                    required that the bandwidth be 2 Mbit/s, the maximum link bandwidth of the
                    tunnel be 5 Mbit/s, and the maximum reservable bandwidth be 10 Mbit/s.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                     468
VPN Configuration
VPN Configuration                                                                           5 VPWS Configuration


                    Figure 5-25 Network diagram of configuring LDP VPWS (using an MPLS TE
                    tunnel)
                          NOTE

                         In this example, interface1 and interface2 represent VLANIF 10 and VLANIF 20, respectively.




