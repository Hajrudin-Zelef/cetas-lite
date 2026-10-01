---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-138
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [19707, 19913]
sha256: 313886ccdf0da56a28b3c86ceb12994445880a157da384df5496ba1186a33d1f
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        #
                        vlan batch 100
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0001.00
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls te bandwidth change thresholds up 20
                         mpls te bandwidth change thresholds down 20
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                         isis enable 1
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 4.4.4.9
                         mpls te bandwidth ct0 10000
                         mpls te tunnel-id 1
                        #
                        return

                 ●      LSR2
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
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 00.0005.0000.0000.0002.00
                         traffic-eng level-2
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      331
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                         ip address 10.1.2.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls te bandwidth change thresholds up 20
                         mpls te bandwidth change thresholds down 20
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
                         isis enable 1
                        #
                        return

                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0003.00
                        #
                        interface Vlanif100
                         ip address 10.1.3.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         isis enable 1
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
                         ip address 3.3.3.9 255.255.255.255
                         isis enable 1




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      332
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                        #
                        return

                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0004.00
                        #
                        interface Vlanif100
                         ip address 10.1.3.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                         isis enable 1
                        #
                        return



4.14 Configuring a Metric for MPLS TE Tunnel Path
Selection
Prerequisites
                 Before configuring a metric for MPLS TE tunnel path selection, you have
                 completed the following task:
                 ●      Configure a dynamic MPLS TE tunnel.

Context
                 In addition to the IGP metric, the TE metric can be used for MPLS TE path
                 computation. This enables TE tunnel path computation to be more independent of
                 IGP route selection, making tunnel path control more flexible.

