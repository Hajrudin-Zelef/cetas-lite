---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-134
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [19079, 19272]
sha256: ba51b74cc6ddd8f72a478e69240a2e58370f75d93437e0b6104b6be8d75d54bf
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                         isis enable 1
                        #
                        interface Tunnel30
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 5.5.5.5
                         mpls te record-route
                         mpls te tunnel-id 300
                         mpls te path explicit-path tope2
                         mpls te bypass-tunnel
                         mpls te protected-interface vlanif 200
                        #
                        return
                 ●      P2
                        #
                        sysname P2
                        #
                        vlan batch 200 400
                        #
                        bfd
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                         mpls te cspf
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 86.4501.0030.0300.3003.00
                         traffic-eng level-2
                        #
                        interface vlanif 400
                         ip address 10.4.1.1 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
                        interface vlanif 200
                         ip address 10.2.1.2 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te bfd enable
                         mpls rsvp-te bfd min-tx-interval 100 min-rx-interval 100 detect-multiplier 3
                         mpls rsvp-te hello
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        321
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      P3
                        #
                        sysname P3
                        #
                        vlan batch 300 500
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                         mpls te cspf
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 86.4501.0040.0400.4004.00
                         traffic-eng level-2
                        #
                        interface vlanif 300
                         ip address 10.3.1.2 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
                        interface vlanif 500
                         ip address 10.5.1.1 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      PE2
                        #
                        sysname PE2
                        #
                        vlan batch 400 500
                        #
                        mpls lsr-id 5.5.5.5
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                         mpls te cspf


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      322
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration

                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 86.4501.0050.0500.5005.00
                         traffic-eng level-2
                        #
                        interface vlanif 400
                         ip address 10.4.1.2 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
                        interface vlanif 500
                         ip address 10.5.1.2 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
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
                         ip address 5.5.5.5 255.255.255.255
                         isis enable 1
                        #
                        return



4.13 Configuring the MPLS TE Bandwidth Flooding
Threshold

