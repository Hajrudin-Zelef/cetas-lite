---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-307
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [45486, 45626]
sha256: 65340311456002edcf03597789ac02fbecbafde9f47115ad75a598d4a73cb00d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    After the shutdown command is run on VLANIF 10 that is bound to the VSI on
                    UPE1, CE1 and CE2 cannot ping each other. This indicates that data is transmitted
                    through the PW of this VSI.
                    # Before shutting down 10GE1/0/2 of SPE1, check the MAC addresses learned by
                    the VSI on SPE2. (The following displays only the dynamic MAC addresses of the
                    VSI named V100.)
                    [SPE2] display mac-address vsi V100
                    -------------------------------------------------------------------------------
                    MAC Address VLAN/VSI                              Learned-From           Type
                    -------------------------------------------------------------------------------
                    00e0-fc12-3456 -/V100                            10GE1/0/1             dynamic
                    00e0-fc34-5678 -/V100                            10GE1/0/1             dynamic

                    -------------------------------------------------------------------------------
                    Total items displayed = 2

                    Run the shutdown command on 10GE1/0/2 of SPE1 to enable the VSI bound to
                    static VPWS to go down.
                    [SPE1] interface 10GE1/0/2
                    [SPE1-10GE1/0/2] shutdown
                    [SPE1-10GE1/0/2] quit

                    # Check the MAC addresses learned by the VSI on SPE2. The command output
                    shows that the MAC address learned from 10GE1/0/2 has been deleted.
                    [SPE2] display mac-address vsi V100
                    -------------------------------------------------------------------------------
                    MAC Address VLAN/VSI                              Learned-From           Type
                    -------------------------------------------------------------------------------

                    -------------------------------------------------------------------------------
                    Total items displayed = 0


Configuration Scripts
                    ●     CE1
                          #
                          sysname CE1
                          #
                          vlan 10
                          #
                          interface Vlanif10
                           ip address 10.1.1.1 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk
                           port trunk allow-pass vlan 10
                          #
                          return
                    ●     CE2
                          #
                          sysname CE2
                          #
                          vlan 60
                          #
                          interface Vlanif60
                           ip address 10.1.1.2 255.255.255.0
                          #
                          interface 10GE1/0/1
                           port link-type trunk


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       731
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration

                         port trunk allow-pass vlan 60
                        #
                        return
                    ●   UPE1
                        #
                        sysname UPE1
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 4.4.4.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        interface Vlanif10
                         mpls static-l2vc destination 1.1.1.9 transmit-vpn-label 100 receive-vpn-label 100
                        #
                        interface Vlanif20
                         ip address 3.1.1.2 255.255.255.0
                         mpls
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 3.1.1.0 0.0.0.255
                        #
                        static-lsp ingress UPE1toSPE1 destination 1.1.1.9 32 nexthop 3.1.1.1 out-label 20
                        static-lsp egress SPE1toUPE1 incoming-interface Vlanif20 in-label 30
                        #
                        return
                    ●   SPE1
                        #
                        sysname SPE1
                        #
                        vlan batch 20 30
                        #
                        mpls lsr-id 1.1.1.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi V100 static
                         pwsignal ldp
                          vsi-id 100
                          mac-withdraw enable
                          peer 3.3.3.9
                          peer 4.4.4.9 static-upe trans 100 recv 100
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 3.3.3.9
                         remote-ip 3.3.3.9
                        #
                        interface Vlanif20
                         ip address 3.1.1.1 255.255.255.0
                         mpls
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          732
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration

