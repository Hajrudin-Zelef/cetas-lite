---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-308
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [45627, 45818]
sha256: 49e213c78a25396f70b35e0fdb126f940cfe801f952ad720c829b402d043e6e2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        interface Vlanif30
                         ip address 1.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 20
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 1.1.1.0 0.0.0.255
                          network 3.1.1.0 0.0.0.255
                        #
                        static-lsp ingress SPE1toUPE1 destination 4.4.4.9 32 nexthop 3.1.1.2 out-label 30
                        static-lsp egress UPE1toSPE1 incoming-interface Vlanif20 in-label 20
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 30 40
                        #
                        mpls lsr-id 2.2.2.9
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif30
                         ip address 1.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif40
                         ip address 2.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 1.1.1.0 0.0.0.255
                          network 2.1.1.0 0.0.0.255
                        #
                        return
                    ●   SPE2
                        #
                        sysname SPE2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          733
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration

                        #
                        vlan batch 40 50
                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        vsi V100 static
                         pwsignal ldp
                          vsi-id 100
                          mac-withdraw enable
                          peer 1.1.1.9
                          peer 5.5.5.9 static-upe trans 100 recv 100
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 1.1.1.9
                         remote-ip 1.1.1.9
                        #
                        interface Vlanif40
                         ip address 2.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif50
                         ip address 4.1.1.1 255.255.255.0
                         mpls
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 40
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 2.1.1.0 0.0.0.255
                          network 4.1.1.0 0.0.0.255
                        #
                        static-lsp ingress SPE2toUPE2 destination 5.5.5.9 32 nexthop 4.1.1.2 out-label 50
                        static-lsp egress UPE2toSPE2 incoming-interface Vlanif50 in-label 40
                        #
                        return
                    ●   UPE2
                        #
                        sysname UPE2
                        #
                        vlan batch 50 60
                        #
                        mpls lsr-id 5.5.5.9
                        mpls
                        #
                        mpls l2vpn
                        #
                        interface Vlanif50
                         ip address 4.1.1.2 255.255.255.0
                         mpls
                        #
                        interface Vlanif60
                         mpls static-l2vc destination 3.3.3.9 transmit-vpn-label 100 receive-vpn-label 100
                        #
                        interface 10GE1/0/1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          734
VPN Configuration
VPN Configuration                                                                                    6 VPLS Configuration

                         port link-type trunk
                         port trunk allow-pass vlan 50
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 60
                        #
                        interface LoopBack1
                         ip address 5.5.5.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 5.5.5.9 0.0.0.0
                          network 4.1.1.0 0.0.0.255
                        #
                        static-lsp ingress UPE2toSPE2 destination 3.3.3.9 32 nexthop 4.1.1.1 out-label 40
                        static-lsp egress SPE2toUPE2 incoming-interface Vlanif50 in-label 50
                        #
                        return


6.13.6 Example for Configuring Dynamic VPWS Accessing
VPLS
Networking Requirements
                    Figure 6-36 shows a backbone network built by an enterprise. UPEs support
                    dynamic VPWS for accessing SPEs. Site1 connects to the backbone network by
                    connecting CE1 to UPE1, while Site2 connects to the backbone network by
                    connecting CE2 to UPE2. Users at Site1 and Site2 need to communicate at Layer 2,
                    and user information needs to be retained in Layer 2 packets when the packets
                    are transmitted over the backbone network.

                    Figure 6-36 Network diagram of configuring dynamic VPWS accessing VPLS
                         NOTE

                        In this example, interface1 and interface2 represent 10GE1/0/1 and 10GE1/0/2, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                          735

