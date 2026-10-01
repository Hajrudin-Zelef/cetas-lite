---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-151
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [21704, 21854]
sha256: e662f33e758f44d8f28fcd2349ca506eb12eb00c76b6975d76d9245b6c75cadf
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         port trunk allow-pass vlan 800
                        #
                        interface LoopBack1
                         ip address 5.5.5.5 255.255.255.255
                         isis enable 1
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 6.6.6.6
                         mpls te tunnel-id 100
                         mpls te record-route
                         mpls te bandwidth ct0 10000
                         mpls te backup hot-standby
                         mpls te path explicit-path main
                        #
                        return
                 ●      P1
                        #
                        sysname P1
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0001.00
                         traffic-eng level-1-2
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
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
                         ip address 1.1.1.1 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      P2
                        #
                        sysname P2
                        #
                        vlan batch 200 400 500
                        #
                        mpls lsr-id 2.2.2.2


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      363
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0002.00
                         traffic-eng level-1-2
                        #
                        interface Vlanif200
                         ip address 10.2.1.2 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.4.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.5.1.1 255.255.255.252
                         mpls
                         mpls te
                         isis enable 1
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      P3
                        #
                        sysname P3
                        #
                        vlan batch 400 700
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         cost-style wide
                         network-entity 10.0000.0000.0003.00
                         traffic-eng level-1-2
                        #
                        interface Vlanif400
                         ip address 10.4.1.2 255.255.255.252
                         mpls
                         mpls te


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      364
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

