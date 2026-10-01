---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-165
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [23844, 24027]
sha256: 8e4d70364057c3d5d81fa03f92773cb1e623e267628f52005305aae7dfe77c60
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 500
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        explicit-path backup-path
                         next hop 10.1.5.2
                         next hop 10.1.3.1
                         next hop 3.3.3.9
                        #
                        explicit-path pri-path
                         next hop 10.1.1.2
                         next hop 10.1.2.2
                         next hop 3.3.3.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
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
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te tunnel-id 1
                         mpls te record-route
                         mpls te path explicit-path pri-path
                         mpls te path explicit-path backup-path secondary
                         mpls te backup hot-standby mode revertive wtr 15
                         mpls te backup ordinary best-effort
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          mpls-te enable


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      397
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        return

                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100 200 400
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
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
                        vlan batch 200 300
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      398
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

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
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 300
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
                          mpls-te enable
                        #
                        return

