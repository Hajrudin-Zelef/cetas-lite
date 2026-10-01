---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-198
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [28611, 28823]
sha256: 6541640951cfd1d2da71c2630b96eb3760d157fcf0c08e88386d7d6dc7269991
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 200 300 500 700
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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      476
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

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

                 ●      LSR5
                        #
                        sysname LSR5
                        #
                        vlan batch 400 500
                        #
                        mpls lsr-id 5.5.5.9
                        #
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
                         ip address 10.1.5.1 255.255.255.0
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



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      477
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        interface LoopBack1
                         ip address 5.5.5.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 5.5.5.9 0.0.0.0
                          network 10.1.4.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          mpls-te enable
                        #
                        return
                 ●      LSR6
                        #
                        sysname LSR6
                        #
                        vlan batch 600 700
                        #
                        mpls lsr-id 6.6.6.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif600
                         ip address 10.1.6.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif700
                         ip address 10.1.7.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 600
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 700
                        #
                        interface LoopBack1
                         ip address 6.6.6.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 6.6.6.9 0.0.0.0
                          network 10.1.6.0 0.0.0.255
                          network 10.1.7.0 0.0.0.255
                          mpls-te enable
                        #
                        return



4.27.7 Example for Configuring MPLS TE Auto FRR (TE FRR
Link Protection on the Ingress)
Networking Requirements
                 On the network shown in Figure 4-51, establish a primary tunnel along the
                 explicit path LSR1 -> LSR2 -> LSR3. The ingress (LSR1) also functions as a PLR and

