---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-214
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [31073, 31248]
sha256: 26946661adae2306274edadffe7a52e5443cf41c51c4998ba6ccbffd4892c450
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        mpls lsr-id 2.2.2.9
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        explicit-path by-path
                         next hop 10.1.4.2
                         next hop 10.1.5.2
                         next hop 3.3.3.9
                        #
                        interface Vlanif100
                         ip address 1.1.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te tunnel-id 300
                         mpls te record-route
                         mpls te path explicit-path by-path
                         mpls te bypass-tunnel
                         mpls te protected-interface Vlanif200
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 200 300 500 700


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      514
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration

                        #
                        mpls lsr-id 3.3.3.9
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif700
                         ip address 10.1.7.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 700
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          network 10.1.7.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 300
                        #
                        mpls lsr-id 4.4.4.9
                        mpls
                         mpls te


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                      515
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 10.1.3.0 0.0.0.255
                          mpls-te enable
                        #
                        return

