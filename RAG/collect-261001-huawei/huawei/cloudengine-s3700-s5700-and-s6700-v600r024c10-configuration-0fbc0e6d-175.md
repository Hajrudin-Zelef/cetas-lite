---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-175
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [25401, 25593]
sha256: b92b5bfe62e6f8407beb431ccf3f5979d1faf7f76b1ddf076f032ed8606231d2
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

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
                        bfd
                         mpls-passive
                        #
                        mpls lsr-id 3.3.3.9
                        #
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      423
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

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

                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 300 400 500
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.2 255.255.255.0
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
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      424
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         ip address 4.4.4.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 10.1.3.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          mpls-te enable
                        #
                        return




4.26 Configuring Manual MPLS TE FRR

4.26.1 Understanding Manual MPLS TE FRR
                 MPLS TE FRR protects links and nodes on MPLS TE tunnels. If a link or node fails,
                 TE FRR rapidly switches traffic to a backup path, minimizing traffic loss.
                 Manual FRR allows you to manually create a bypass tunnel that meets
                 requirements and bind the bypass tunnel to the primary tunnel.

Context
                 Generally, a link or node failure on an MPLS TE tunnel triggers a tunnel switchover
                 from the primary path to the backup path. The switching process involves IGP
                 route convergence on the backup path, CSPF path recalculation, and CR-LSP re-
                 establishment. The switching process is slow, and traffic loss may occur.
                 TE FRR can solve this problem. TE FRR establishes a backup path that bypasses the
                 faulty link or node in advance. When a link or node on an MPLS TE tunnel fails,
                 traffic can be quickly switched to the backup path, preventing traffic loss. When
                 traffic continues to be transmitted over the backup path, the ingress continues to
                 reestablish the primary path.

Benefits
                 TE FRR provides carrier-class local protection capabilities for MPLS TE, improving
                 the reliability of an entire network.

