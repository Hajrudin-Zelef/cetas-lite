---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-203
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [29753, 29951]
sha256: 78c58cb135a13c848d5b4ed860c93b7ad6f604f1b9627e42dec255f96fdc0308
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                        #
                        mpls ldp remote-peer 3.3.3.9
                         remote-ip 3.3.3.9
                        #
                        interface Vlanif10
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 10000
                         mpls te bandwidth bc0 5000
                         mpls rsvp-te
                        #
                        interface Vlanif20
                         mpls l2vc 3.3.3.9 10 tunnel-policy policy1
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
                        interface Tunnel10
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         mpls te signal-protocol rsvp-te
                         destination 3.3.3.9
                         mpls te bandwidth ct0 2000
                         mpls te tunnel-id 10
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          mpls-te enable
                        #
                        tunnel-policy policy1
                         tunnel select-seq cr-lsp load-balance-number 1
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                          mpls te
                          mpls rsvp-te
                        #
                        interface Vlanif10
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 10000
                         mpls te bandwidth bc0 5000
                         mpls rsvp-te
                        #
                        interface Vlanif20
                         ip address 10.2.1.1 255.255.255.0
                         mpls
                         mpls te


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                   476
VPN Configuration
VPN Configuration                                                            5 VPWS Configuration

                         mpls te bandwidth max-reservable-bandwidth 10000
                         mpls te bandwidth bc0 5000
                         mpls rsvp-te
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
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.2.1.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        mpls l2vpn
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 1.1.1.9
                         remote-ip 1.1.1.9
                        #
                        interface Vlanif10
                         mpls l2vc 1.1.1.9 10 tunnel-policy policy1
                        #
                        interface Vlanif20
                         ip address 10.2.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 10000
                         mpls te bandwidth bc0 5000
                         mpls rsvp-te
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
                         ip address 3.3.3.9 255.255.255.255
                        #
                        interface Tunnel10
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                   477
VPN Configuration
VPN Configuration                                                                5 VPWS Configuration

                         mpls te signal-protocol rsvp-te
                         destination 1.1.1.9
                         mpls te bandwidth ct0 2000
                         mpls te tunnel-id 10
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                          mpls-te enable
                        #
                        tunnel-policy policy1
                         tunnel select-seq cr-lsp load-balance-number 1
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



5.8 Configuring SVC VPWS

5.8.1 Understanding SVC VPWS

Definition
                    SVC VPWS uses inner labels manually configured on PEs for data transmission,
                    whereas LDP VPWS uses LDP to exchange VC labels. The SVC mode can be
                    regarded as the simplified LDP mode.

                    VC labels of SVC VPWS are statically configured and do not require VC label
                    mapping, removing the need to use LDP signaling to transmit VC labels.


