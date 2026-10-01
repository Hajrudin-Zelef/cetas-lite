---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-202
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [29236, 29442]
sha256: 0afbb3e1b989bfb1af5a3b5e8dd889c23c813efc1ee5d48058fd5b6527a3ad83
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.41.1.2 255.255.255.0
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
                         ip address 1.1.1.1 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.3
                         mpls te record-route label
                         mpls te fast-reroute
                         mpls te tunnel-id 1
                         mpls te path explicit-path master
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.21.1.0 0.0.0.255
                          network 10.41.1.0 0.0.0.255
                          mpls-te enable
                        #
                        return

                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 200 300 500
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.21.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.31.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      485
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         ip address 10.32.1.1 255.255.255.0
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
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                         isis enable 1
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.21.1.0 0.0.0.255
                          network 10.31.1.0 0.0.0.255
                          network 10.32.1.0 0.0.0.255
                          mpls-te enable
                        #
                        return

                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 100 300
                        #
                        mpls lsr-id 3.3.3.3
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
                        interface Vlanif300
                         ip address 10.31.1.2 255.255.255.0
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
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      486
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                          network 3.3.3.3 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.31.1.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 400 500
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.41.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.32.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.32.1.0 0.0.0.255
                          network 10.41.1.0 0.0.0.255
                          mpls-te enable
                        #
                        return



