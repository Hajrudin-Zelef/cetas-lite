---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-184
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [26696, 26895]
sha256: 2ea35a10e3090147f20a582f8415d0b2d7db350bef98aa8a296d8f48190831d7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 After a period of time, run the display mpls te tunnel name Tunnel1 verbose
                 command on LSR2. The command output shows that Bypass In Use of Tunnel1 is
                 Not Used, indicating that data flows have been switched back to VLANIF 200 for
                 transmission.

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
                         mpls te cspf
                         mpls rsvp-te
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


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      445
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                         mpls rsvp-te
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
                         isis enable 1
                        #
                        interface Tunnel2
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te record-route
                         mpls te tunnel-id 2
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


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      446
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

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
                        #
                        return

