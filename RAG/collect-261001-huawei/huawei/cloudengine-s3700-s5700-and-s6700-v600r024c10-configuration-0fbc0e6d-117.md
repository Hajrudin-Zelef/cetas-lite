---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-117
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [16548, 16749]
sha256: 4f7c07714d261f1ab344ade7c8a0f54ede0123a41e14da3db4551c4bc53d4163
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100
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
                         next hop 10.1.3.2
                         next hop 4.4.4.9
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


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                              279
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 4.4.4.9
                         mpls te record-route label
                         mpls te fast-reroute
                         mpls te tunnel-id 1
                         mpls te path explicit-path pri-path
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
                         mpls rsvp-te srefresh
                         mpls te cspf
                        #
                        explicit-path by-path
                         next hop 10.1.4.2
                         next hop 10.1.5.2
                         next hop 3.3.3.9
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
                         ip address 10.1.2.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.1 255.255.255.0
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
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      280
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                         isis enable 1
                        #
                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te tunnel-id 2
                         mpls te record-route
                         mpls te path explicit-path by-path
                         mpls te bypass-tunnel
                         mpls te protected-interface Vlanif200
                        #
                        return

                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 200 300 500
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te srefresh
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 00.0005.0000.0000.0003.00
                         traffic-eng level-2
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/5
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                         isis enable 1



Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      281
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

