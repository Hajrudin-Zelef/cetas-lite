---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-211
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [30648, 30789]
sha256: 1e4a819d99b7d13f953d60ea12d679e13a4f37504c79d4af8630bc2736d88239
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 [LSR1-Vlanif100] quit
                 [LSR1] interface vlanif 600
                 [LSR1-Vlanif600] ip address 10.1.6.2 24
                 [LSR1-Vlanif600] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface 10ge 1/0/2
                 [LSR1-10GE1/0/2] port link-type trunk
                 [LSR1-10GE1/0/2] port trunk allow-pass vlan 600
                 [LSR1-10GE1/0/2] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.9 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2, LSR3, LSR4, LSR5 and LSR6 are similar to the
                 configuration of LSR1. For detailed configurations, see Configuration Scripts.

         Step 2 Configure OSPF to advertise routes.

                 # Configure LSR1.
                 [LSR1] ospf 1
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.6.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2, LSR3, LSR4, LSR5 and LSR6 are similar to the
                 configuration of LSR1. For detailed configurations, see Configuration Scripts.

                 # After the configuration is complete, run the display ip routing-table command
                 on each node to check whether the nodes have learned routes from each other.
                 The following example uses the command output on LSR1.
                 [LSR1] display ip routing-table
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 21        Routes : 25

                 Destination/Mask     Proto Pre Cost        Flags NextHop                          Interface

                       1.1.1.9/32 Direct 0 0           D 127.0.0.1                         LoopBack1
                       2.2.2.9/32 OSPF 10 1             D 10.1.1.2                          Vlanif100
                       3.3.3.9/32 OSPF 10 2             D 10.1.6.1                          Vlanif600
                                OSPF 10 2             D 10.1.1.2                         Vlanif100
                       4.4.4.9/32 OSPF 10 3             D 10.1.6.1                          Vlanif600
                                OSPF 10 3             D 10.1.1.2                         Vlanif100
                       5.5.5.9/32 OSPF 10 2             D 10.1.1.2                          Vlanif100
                       6.6.6.9/32 OSPF 10 1             D 10.1.6.1                          Vlanif600
                      10.1.1.0/24 Direct 0 0           D 10.1.1.1                          Vlanif100
                      10.1.1.1/32 Direct 0 0           D 127.0.0.1                          Vlanif100
                     10.1.1.255/32 Direct 0 0           D 127.0.0.1                          Vlanif100
                      10.1.2.0/24 OSPF 10 2             D 10.1.1.2                           Vlanif100
                      10.1.3.0/24 OSPF 10 3             D 10.1.6.1                           Vlanif600
                                OSPF 10 3             D 10.1.1.2                         Vlanif100
                      10.1.4.0/24 OSPF 10 2             D 10.1.1.2                           Vlanif100
                      10.1.5.0/24 OSPF 10 3             D 10.1.6.1                           Vlanif600
                                OSPF 10 3             D 10.1.1.2                         Vlanif100
                      10.1.6.0/24 Direct 0 0           D 10.1.6.2                          Vlanif600
                      10.1.6.2/32 Direct 0 0           D 127.0.0.1                          Vlanif600
                     10.1.6.255/32 Direct 0 0           D 127.0.0.1                          Vlanif600
                      10.1.7.0/24 OSPF 10 2             D 10.1.6.1                           Vlanif600


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          508
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration

                     127.0.0.0/8 Direct 0 0            D 127.0.0.1           InLoopBack0
                     127.0.0.1/32 Direct 0 0           D 127.0.0.1            InLoopBack0
                 127.255.255.255/32 Direct 0 0           D 127.0.0.1             InLoopBack0
                 255.255.255.255/32 Direct 0 0           D 127.0.0.1             InLoopBack0

         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.

                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along tunnels. Enable CSPF on the ingress of the primary tunnel and the PLR of
                 the bypass tunnel.

                 # Configure LSR1.
                 [LSR1] mpls lsr-id 1.1.1.9
                 [LSR1] mpls
                 [LSR1-mpls] mpls te
                 [LSR1-mpls] mpls rsvp-te
                 [LSR1-mpls] mpls te cspf
                 [LSR1-mpls] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls
                 [LSR1-Vlanif100] mpls te
                 [LSR1-Vlanif100] mpls rsvp-te
                 [LSR1] interface vlanif 600
                 [LSR1-Vlanif600] mpls
                 [LSR1-Vlanif600] mpls te
                 [LSR1-Vlanif600] mpls rsvp-te
                 [LSR1-Vlanif600] quit

                 The configurations of LSR2, LSR3, LSR4, LSR5 and LSR6 are similar to the
                 configuration of LSR1. For detailed configurations, see Configuration Scripts.

         Step 4 Configure OSPF TE.

                 # Configure LSR1.
                 [LSR1] ospf
                 [LSR1-ospf-1] opaque-capability enable
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] mpls-te enable
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2, LSR3, LSR4, LSR5 and LSR6 are similar to the
                 configuration of LSR1. For detailed configurations, see Configuration Scripts.

         Step 5 Create the primary MPLS TE tunnel on the ingress LSR1.

                 # Configure an explicit path for the primary tunnel.
                 [LSR1] explicit-path pri-path
                 [LSR1-explicit-path-pri-path] next hop 10.1.1.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.2.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.3.2
                 [LSR1-explicit-path-pri-path] next hop 4.4.4.9
                 [LSR1-explicit-path-pri-path] quit

                 # Configure the primary MPLS TE tunnel and bind it to the explicit path.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 4.4.4.9
                 [LSR1-Tunnel1] mpls te tunnel-id 100
                 [LSR1-Tunnel1] mpls te path explicit-path pri-path

                 # Enable TE FRR.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                          509
MPLS Configuration
MPLS Configuration                                                                                 4 MPLS TE Configuration

