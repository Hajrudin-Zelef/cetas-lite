---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-107
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [15029, 15173]
sha256: 72b190052bcb39eadfe64e7a9386e1e82a5673295c69e9cafcb167a6ea9e761d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 400
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf
                        #
                        explicit-path pri-path
                         next hop 10.1.1.2
                         next hop 10.1.2.2
                         next hop 3.3.3.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         ospf cost 15
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.1 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te tunnel-id 1
                         mpls te path explicit-path pri-path
                         mpls te igp shortcut ospf
                         mpls te igp metric absolute 10
                        #
                        ospf 1
                         opaque-capability enable
                         enable traffic-adjustment
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          mpls-te enable
                        #
                        return

                 ●      LSR2


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         255
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        sysname LSR2
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         ospf cost 15
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         ospf cost 10
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
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
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
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         ospf cost 10
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.1 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      256
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

