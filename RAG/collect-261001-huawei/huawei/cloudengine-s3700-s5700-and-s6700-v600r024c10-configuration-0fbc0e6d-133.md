---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-133
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [18913, 19078]
sha256: a8e10cbfd1d2d8fefde2ac62f11bd3f549747efdf4dd301f176dabaa192c29d9
---

                 # Configure P1.
                 [P1] bfd
                 [P1-bfd] quit
                 [P1] interface vlanif 200
                 [P1-Vlanif200] mpls rsvp-te bfd enable
                 [P1-Vlanif200] mpls rsvp-te bfd min-tx-interval 100 min-rx-interval 100 detect-multiplier 3
                 [P1-Vlanif200] quit

                 # Configure P2.
                 [P2] bfd
                 [P2-bfd] quit
                 [P2] interface vlanif 200
                 [P2-Vlanif200] mpls rsvp-te bfd enable


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                      318
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration

                 [P2-Vlanif200] mpls rsvp-te bfd min-tx-interval 100 min-rx-interval 100 detect-multiplier 3
                 [P2-Vlanif200] quit

                 ----End

Verifying the Configuration
                 # After the configuration is complete, check BFD session information on P1 and
                 P2. The following example uses the command output on P1.
                 <P1> display mpls rsvp-te bfd session all
                 Total Nbrs/Rsvp triggered sessions : 3/1
                 -------------------------------------------------------------------------------
                 Local      Remote       Local         Peer           Interface       Session
                 Discr     Discr      Addr           Addr           Name             State
                 -------------------------------------------------------------------------------
                 16385       16385       10.2.1.1       10.2.1.2       Vlanif200         UP

                 The command output shows that the BFD session is up.

Configuration Scripts
                         NOTE

                        The configuration script of Switch is not provided in this example.
                 ●      PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                         mpls te cspf
                        #
                        explicit-path tope2
                         next hop 10.1.1.2
                         next hop 10.2.1.2
                         next hop 10.4.1.2
                         next hop 5.5.5.5
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 86.4501.0010.0100.1001.00
                         traffic-eng level-2
                        #
                        interface vlanif 100
                         ip address 10.1.1.1 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                         isis enable 1


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                          319
MPLS Configuration
MPLS Configuration                                                                              4 MPLS TE Configuration

                        #
                        interface Tunnel10
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 5.5.5.5
                         mpls te record-route label
                         mpls te fast-reroute
                         mpls te tunnel-id 100
                         mpls te path explicit-path tope2
                        #
                        return
                 ●      P1
                        #
                        sysname P1
                        #
                        vlan batch 100 200 300
                        #
                        bfd
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                         mpls te cspf
                        #
                        explicit-path tope2
                         next hop 10.3.1.2
                         next hop 10.5.1.2
                         next hop 5.5.5.5
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 86.4501.0020.0200.2002.00
                         traffic-eng level-2
                        #
                        interface vlanif 100
                         ip address 10.1.1.2 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
                        interface vlanif 200
                         ip address 10.2.1.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te bfd enable
                         mpls rsvp-te bfd min-tx-interval 100 min-rx-interval 100 detect-multiplier 3
                         mpls rsvp-te hello
                        #
                        interface vlanif 300
                         ip address 10.3.1.1 255.255.255.252
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                         mpls rsvp-te hello
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        320
MPLS Configuration
MPLS Configuration                                                                              4 MPLS TE Configuration

