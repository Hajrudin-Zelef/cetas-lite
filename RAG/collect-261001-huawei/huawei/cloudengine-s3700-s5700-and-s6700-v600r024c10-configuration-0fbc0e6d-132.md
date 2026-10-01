---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-132
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [18756, 18912]
sha256: 0ce4b8fb3e8b16456f71aa0507bc8a448bb20566acf9e0396775a1a1632bc558
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 However, the interval for sending RSVP-TE Hello messages to declare a neighbor-
                 down event is three times the interval for sending Hello messages. Therefore, the
                 time for P1 to detect the RSVP neighbor fault is much later than that when no
                 Layer 2 device is deployed. This leads to severe packet loss. To minimize traffic
                 loss, BFD can be configured for P1 to rapidly detect the fault in the link between
                 P2 and Switch and trigger TE FRR switching.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                                        315
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 Figure 4-22 Network diagram for configuring dynamic BFD for RSVP




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Configure IP addresses for interfaces, including the loopback interfaces whose
                        addresses are to be used as MPLS LSR IDs.
                 2.     Configure IS-IS to ensure that nodes can reach each other over public network
                        routes.
                 3.     Configure the MPLS network and basic MPLS TE functions.
                 4.     Configure primary and bypass tunnels and bind them to explicit paths.
                 5.     Create a primary TE tunnel interface and enable TE FRR on PE1. Configure a
                        bypass tunnel on P1.
                 6.     Configure BFD for RSVP on P1 and P2.

Procedure
         Step 1 Configure interface IP addresses for the devices.
                 # Configure PE1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname PE1
                 [PE1] vlan batch 100
                 [PE1] interface vlanif 100
                 [PE1-Vlanif100] ip address 10.1.1.1 30
                 [PE1-Vlanif100] quit
                 [PE1] interface 10ge 1/0/1
                 [PE1-10GE1/0/1] port link-type trunk
                 [PE1-10GE1/0/1] port trunk allow-pass vlan 100


Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                           316
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                 [PE1-10GE1/0/1] quit
                 [PE1] interface loopback 1
                 [PE1-loopback1] ip address 1.1.1.1 32
                 [PE1-loopback1] quit

                 The configurations of P1, P2, P3, and PE2 are similar to the configuration of PE1.
                 For detailed configurations, see Configuration Scripts.
         Step 2 Configure Switch.
                 Configure Switch so that P1 and P2 can communicate. The configuration details
                 are not provided here.
         Step 3 Configure IS-IS to advertise routes.
                 # Configure PE1.
                 [PE1] isis 1
                 [PE1-isis-1] is-level level-2
                 [PE1-isis-1] network-entity 86.4501.0010.0100.1001.00
                 [PE1-isis-1] quit
                 [PE1] interface vlanif 100
                 [PE1-Vlanif100] isis enable 1
                 [PE1-Vlanif100] quit
                 [PE1] interface loopback 1
                 [PE1-LoopBack1] isis enable 1
                 [PE1-LoopBack1] quit

                 The configurations of P1, P2, P3, and PE2 are similar to the configuration of PE1.
                 For detailed configurations, see Configuration Scripts.
         Step 4 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.
                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the tunnel ingress PE1 and the PLR PE2.
                 # Configure PE1.
                 [PE1] mpls lsr-id 1.1.1.1
                 [PE1] mpls
                 [PE1-mpls] mpls te
                 [PE1-mpls] mpls rsvp-te
                 [PE1-mpls] mpls te cspf
                 [PE1-mpls] quit
                 [PE1] interface vlanif 100
                 [PE1-Vlanif100] mpls
                 [PE1-Vlanif100] mpls te
                 [PE1-Vlanif100] mpls rsvp-te
                 [PE1-Vlanif100] quit

                 The configurations of P1, P2, P3, and PE2 are similar to the configuration of PE1.
                 For detailed configurations, see Configuration Scripts.
         Step 5 Configure IS-IS TE.
                 # Configure PE1.
                 [PE1] isis 1
                 [PE1-isis-1] cost-style wide
                 [PE1-isis-1] traffic-eng level-2
                 [PE1-isis-1] quit

                 The configurations of P1, P2, P3, and PE2 are similar to the configuration of PE1.
                 For detailed configurations, see Configuration Scripts.
         Step 6 Configure the primary tunnel and bind it to an explicit path.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                       317
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                 # Configure an explicit path for the primary tunnel on PE1.
                 [PE1] explicit-path tope2
                 [PE1-explicit-path-tope2] next hop 10.1.1.2
                 [PE1-explicit-path-tope2] next hop 10.2.1.2
                 [PE1-explicit-path-tope2] next hop 10.4.1.2
                 [PE1-explicit-path-tope2] next hop 5.5.5.5
                 [PE1-explicit-path-tope2] quit

                 # Create a tunnel interface on PE1, specify an explicit path, and enable TE FRR.
                 [PE1] interface Tunnel 10
                 [PE1-Tunnel10] ip address unnumbered interface loopback 1
                 [PE1-Tunnel10] tunnel-protocol mpls te
                 [PE1-Tunnel10] destination 5.5.5.5
                 [PE1-Tunnel10] mpls te tunnel-id 100
                 [PE1-Tunnel10] mpls te path explicit-path tope2
                 [PE1-Tunnel10] mpls te fast-reroute
                 [PE1-Tunnel10] quit

                 # After the configuration is complete, run the display mpls te tunnel-interface
                 command on PE1. The command output shows that the status of Tunnel10 on PE1
                 is up.

         Step 7 Configure the bypass tunnel and bind it to an explicit path.

                 # Configure an explicit path for the bypass tunnel on P1.
                 [P1] explicit-path tope2
                 [P1-explicit-path-tope2] next hop 10.3.1.2
                 [P1-explicit-path-tope2] next hop 10.5.1.2
                 [P1-explicit-path-tope2] next hop 5.5.5.5
                 [P1-explicit-path-tope2] quit

                 # Configure a bypass tunnel interface on P1. Specify an explicit path and specify
                 the physical interface to be protected by the bypass tunnel.
                 [P1] interface Tunnel 30
                 [P1-Tunnel30] ip address unnumbered interface loopback 1
                 [P1-Tunnel30] tunnel-protocol mpls te
                 [P1-Tunnel30] destination 5.5.5.5
                 [P1-Tunnel30] mpls te tunnel-id 300
                 [P1-Tunnel30] mpls te path explicit-path tope2
                 [P1-Tunnel30] mpls te bypass-tunnel
                 [P1-Tunnel30] mpls te protected-interface vlanif 200
                 [P1-Tunnel30] quit

         Step 8 Configure BFD for RSVP.

                 # Enable BFD for RSVP on VLANIF200 of P1 and P2. Set the minimum intervals for
                 sending and receiving BFD packets as well as the local BFD detection multiplier.

