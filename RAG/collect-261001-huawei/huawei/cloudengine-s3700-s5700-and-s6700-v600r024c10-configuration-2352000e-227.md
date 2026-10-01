---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-227
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [33350, 33537]
sha256: dd77d93795c5c4aa2d169c5301f629645611a5091ef4e01065ada3d724aecd28
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                         mpls te tunnel-id 13
                         mpls te
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.0.2.0 0.0.0.3
                          network 10.2.1.0 0.0.0.3
                          mpls-te enable
                        #
                        return
                    ●   P
                        #
                        sysname P
                        #
                        vlan batch 10 20
                        #
                        mpls lsr-id 4.4.4.4
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif10
                         ip address 10.0.3.1 255.255.255.252
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif20
                         ip address 10.0.2.2 255.255.255.252
                         mpls
                         mpls te
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
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.0.2.0 0.0.0.3
                          network 10.0.3.0 0.0.0.3
                          mpls-te enable
                        #
                        return
                    ●   PE3
                        #
                        sysname PE3
                        #
                        vlan batch 10
                        #
                        bfd
                        #
                        mpls lsr-id 3.3.3.3
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                   532
VPN Configuration
VPN Configuration                                                                    5 VPWS Configuration

                        #
                        mpls l2vpn
                        #
                        pw-template 3to1
                         peer-address 1.1.1.1
                         control-word
                        #
                        mpls ldp
                        #
                        mpls ldp remote-peer 1.1.1.1
                         remote-ip 1.1.1.1
                        #
                        interface Vlanif10
                         ip address 10.0.3.2 255.255.255.252
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        tunnel-policy p1
                         tunnel select-seq cr-lsp load-balance-number 1
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         mpls l2vc pw-template 3to1 100 tunnel-policy p1
                         mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100
                         mpls l2vpn trigger if-down
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 10
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 1.1.1.1
                         mpls te tunnel-id 31
                         mpls te
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.0.3.0 0.0.0.3
                          mpls-te enable
                        #
                        return
                    ●   PE2
                        #
                        sysname PE2
                        #
                        vlan batch 30
                        #
                        bfd
                        #
                        mpls lsr-id 2.2.2.2
                        mpls
                        #
                        mpls l2vpn
                        #
                        pw-template 2to1
                         peer-address 1.1.1.1
                         control-word
                        #
                        mpls ldp
                        #
                        interface Vlanif30


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                          533
VPN Configuration
VPN Configuration                                                                    5 VPWS Configuration

                         ip address 10.2.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         undo portswitch
                         mpls l2vc pw-template 2to1 200
                         mpls l2vpn pw bfd min-rx-interval 100 min-tx-interval 100
                         mpls l2vpn trigger if-down
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 30
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.2.1.0 0.0.0.3
                        #
                        return



5.11 Configuring BFD for VPWS
Context
                    If a failure occurs on a VPWS network, devices can detect link failures through
                    route convergence. However, the detection is slow and cannot meet the
                    requirements of delay-sensitive services, such as VoIP. To speed up failure
                    detection on the VPWS network, deploy BFD on PEs to rapidly detect PW failures
                    and trigger a rapid switchover of upper-layer applications. BFD provides low-
                    overhead failure detection within milliseconds.
                    A BFD session can be either static or dynamic. Determine which one to use as
                    needed.

