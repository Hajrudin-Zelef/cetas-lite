---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-110
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [15461, 15622]
sha256: 329be2825db0d106136b0feeda8f87540d8a613888118ec70bca52008551bc17
---

                 # Configure an MPLS TE tunnel on LSR1 and bind it to the explicit path.
                 [LSR1] interface tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 3.3.3.9
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] mpls te path explicit-path pri-path

                 The configuration of LSR3 is similar to the configuration of LSR1. For detailed
                 configurations, see Configuration Scripts.
         Step 5 Configure forwarding adjacency.
                 Enable forwarding adjacency on the TE tunnel interface and set the IGP metric
                 value to 10 (absolute metric value).

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         261
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


                 # Configure LSR1.
                 [LSR1-Tunnel1] mpls te igp advertise
                 [LSR1-Tunnel1] mpls te igp metric absolute 10
                 [LSR1-Tunnel1] quit
                 [LSR1] ospf 1
                 [LSR1-ospf-1] enable traffic-adjustment advertise
                 [LSR1-ospf-1] quit

                 The configuration of LSR3 is similar to the configuration of LSR1. For detailed
                 configurations, see Configuration Scripts.

                 ----End

Verifying the Configuration
                 # On LSR1, check the route to LSR3.
                 [LSR1] display ip routing-table 3.3.3.9
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                 Summary Count : 1

                 Destination/Mask     Proto Pre Cost       Flags NextHop        Interface

                     3.3.3.9/32   OSPF    10 10         D 1.1.1.9        Tunnel1

                 The command output shows that the next hop of the route to LSR3 (3.3.3.9) is
                 1.1.1.9, the outbound interface is Tunnel1, and the traffic to LSR3 is steered to the
                 TE tunnel.
                 # On LSR5, check the route to LSR3.
                 [LSR5] display ip routing-table 3.3.3.9
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                 Summary Count : 2

                 Destination/Mask     Proto Pre Cost       Flags NextHop        Interface

                        3.3.3.9/32 OSPF 10 20          D 10.1.5.1         Vlanif500
                                 OSPF 10 20          D 10.1.6.1        Vlanif600

                 The command output shows that there are two equal-cost routes to LSR3
                 (3.3.3.9). Some traffic destined for LSR3 is forwarded through LSR4, and the other
                 traffic is sent to LSR1 and forwarded through the TE tunnel.

Configuration Scripts
                 ●      LSR1
                        #
                        sysname LSR1
                        #
                        vlan batch 100 400 600
                        #
                        mpls lsr-id 1.1.1.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                         mpls te cspf


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         262
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                        #
                        explicit-path pri-path
                         next hop 10.1.1.2
                         next hop 10.1.2.2
                         next hop 3.3.3.9
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         ospf cost 15
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.1 255.255.255.0
                         ospf cost 10
                        #
                        interface Vlanif600
                         ip address 10.1.6.1 255.255.255.0
                         ospf cost 10
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 600
                        #
                        interface LoopBack1
                         ip address 1.1.1.9 255.255.255.255
                        #
                        interface Tunnel1
                         ip address unnumbered interface LoopBack1
                         tunnel-protocol mpls te
                         destination 3.3.3.9
                         mpls te tunnel-id 1
                         mpls te path explicit-path pri-path
                         mpls te igp advertise
                         mpls te igp metric absolute 10
                        #
                        ospf 1
                         opaque-capability enable
                         enable traffic-adjustment advertise
                         area 0.0.0.0
                          network 1.1.1.9 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          network 10.1.6.0 0.0.0.255
                          mpls-te enable
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
                        interface Vlanif100


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      263
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

