---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-197
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [28431, 28610]
sha256: 43d19e3d449ccc53d0a47a31ecb14940a79414ade2b7816a26d7a840aae32c9d
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The command output shows that the automatic bypass tunnel protects the
                 outbound interface VLANIF200 of the primary tunnel and excludes two addresses
                 (10.1.2.1 and 10.1.2.2) on the primary tunnel path to provide link protection.
                 # Run the display mpls te tunnel path command to check tunnel path
                 information on the local node. The following example uses the command output
                 on LSR1.
                 [LSR1] display mpls te tunnel path
                  Tunnel Interface Name : Tunnel1
                  Lsp ID : 1.1.1.9 :1 :1099
                  Hop Information
                   Hop 0 10.1.1.1 Local-Protection available | node
                   Hop 1 10.1.1.2 Label 18
                   Hop 2 2.2.2.9 Label 18
                   Hop 3 10.1.2.1 Local-Protection available


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       473
MPLS Configuration
MPLS Configuration                                                                 4 MPLS TE Configuration

                  Hop 4    10.1.2.2 Label 18
                  Hop 5    3.3.3.9 Label 18
                  Hop 6    10.1.3.1
                  Hop 7    10.1.3.2 Label 3
                  Hop 8    4.4.4.9 Label 3

                  Tunnel Interface Name : AutoBypassTunnel_1.1.1.9_3.3.3.9_32769
                  Lsp ID : 1.1.1.9 :32769 :1100
                  Hop Information
                   Hop 0 10.1.6.2
                   Hop 1 10.1.6.1 Label 16
                   Hop 2 6.6.6.9 Label 16
                   Hop 3 10.1.7.1
                   Hop 4 10.1.7.2 Label 3
                   Hop 5 3.3.3.9 Label 3

                 The information about the outbound interface of the primary tunnel on LSR1
                 shows that node protection and link protection are available to the primary
                 tunnel. The information about the outbound interface of the primary tunnel on
                 LSR2 shows that the link protection is available to the primary tunnel.

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 600
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls te cspf
                         mpls rsvp-te
                         mpls te auto-frr
                        #
                        explicit-path pri-path
                         next hop 10.1.1.2
                         next hop 10.1.2.2
                         next hop 10.1.3.2
                         next hop 4.4.4.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif600
                         ip address 10.1.6.2 255.255.255.0
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
                         port trunk allow-pass vlan 600
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            474
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         tunnel-protocol mpls te
                         destination 4.4.4.9
                         mpls te record-route label
                         mpls te fast-reroute
                         mpls te tunnel-id 1
                         mpls te priority 4 3
                         mpls te path explicit-path pri-path
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.6.0 0.0.0.255
                          mpls-te enable
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
                         mpls te cspf
                         mpls te auto-frr
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
                         mpls te auto-frr link
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
                         isis enable 1
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      475
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          mpls-te enable
                        #
                        return

