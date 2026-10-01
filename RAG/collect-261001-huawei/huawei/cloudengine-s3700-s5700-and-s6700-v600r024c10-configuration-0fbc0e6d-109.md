---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-109
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [15324, 15460]
sha256: 9f803c8f7bc1900b8f687ffc793edc7a0f8df7fef624a11f08da45e4a0f01c00
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
         Step 1 Assign an IP address to each interface and configure OSPF and OSPF costs for
                links.
                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100 400 600
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 255.255.255.0
                 [LSR1-Vlanif100] ospf cost 15
                 [LSR1-Vlanif100] quit
                 [LSR1] interface vlanif 400
                 [LSR1-Vlanif400] ip address 10.1.4.1 255.255.255.0
                 [LSR1-Vlanif400] ospf cost 10
                 [LSR1-Vlanif400] quit


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                         259
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration

                 [LSR1] interface vlanif 600
                 [LSR1-Vlanif600] ip address 10.1.6.1 255.255.255.0
                 [LSR1-Vlanif600] ospf cost 10
                 [LSR1-Vlanif600] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface 10ge 1/0/2
                 [LSR1-10GE1/0/2] port link-type trunk
                 [LSR1-10GE1/0/2] port trunk allow-pass vlan 400
                 [LSR1-10GE1/0/2] quit
                 [LSR1] interface 10ge 1/0/3
                 [LSR1-10GE1/0/3] port link-type trunk
                 [LSR1-10GE1/0/3] port trunk allow-pass vlan 600
                 [LSR1-10GE1/0/3] quit
                 [LSR1] interface loopback 1
                 [LSR1-LoopBack1] ip address 1.1.1.9 255.255.255.255
                 [LSR1-LoopBack1] quit
                 [LSR1] ospf 1
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] network 1.1.1.9 0.0.0.0
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.4.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.6.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2, LSR3, LSR4, and LSR5 are similar to the configuration
                 of LSR1. For detailed configurations, see Configuration Scripts.

                 # After the configuration is complete, check the IP routing tables on LSR1, LSR2,
                 and LSR3. The following example uses the command output on LSR1.
                 [LSR1] display ip routing-table
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 19        Routes : 20

                 Destination/Mask     Proto Pre Cost        Flags NextHop                          Interface

                       1.1.1.9/32 Direct 0 0            D 127.0.0.1                        LoopBack1
                       2.2.2.9/32 OSPF 10 15             D 10.1.1.2                          Vlanif100
                       3.3.3.9/32 OSPF 10 20             D 10.1.4.2                          Vlanif400
                      10.1.1.0/24 Direct 0 0            D 10.1.1.1                         Vlanif100
                      10.1.1.1/32 Direct 0 0            D 127.0.0.1                         Vlanif100
                    10.1.1.255/32 Direct 0 0             D 127.0.0.1                          Vlanif100
                      10.1.2.0/24 OSPF 10 25              D 10.1.1.2                          Vlanif100
                      10.1.3.0/24 OSPF 10 20              D 10.1.4.2                          Vlanif400
                      10.1.4.0/24 Direct 0 0            D 10.1.4.1                         Vlanif400
                      10.1.4.1/32 Direct 0 0            D 127.0.0.1                         Vlanif400
                    10.1.4.255/32 Direct 0 0             D 127.0.0.1                          Vlanif400
                      10.1.5.0/24 OSPF 10 20              D 10.1.6.2                          Vlanif600
                                OSPF 10 20             D 10.1.4.2                         Vlanif400
                      10.1.6.0/24 Direct 0 0            D 10.1.6.1                         Vlanif600
                      10.1.6.1/32 Direct 0 0            D 127.0.0.1                         Vlanif600
                    10.1.6.255/32 Direct 0 0             D 127.0.0.1                          Vlanif600
                     127.0.0.0/8 Direct 0 0             D 127.0.0.1                         InLoopBack0
                     127.0.0.1/32 Direct 0 0             D 127.0.0.1                         InLoopBack0
                 127.255.255.255/32 Direct 0 0             D 127.0.0.1                           InLoopBack0
                 255.255.255.255/32 Direct 0 0             D 127.0.0.1                           InLoopBack0

                 The command output shows that LSR1 has learned routes to other nodes.

         Step 2 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          260
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


                 To create a TE tunnel on LSR1 and LSR3, enable MPLS, MPLS TE, and RSVP-TE
                 globally on LSR1, LSR2, and LSR3 and on all interfaces along the tunnel, and
                 enable CSPF on the ingresses (LSR1 and LSR3).
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
                 [LSR1-Vlanif100] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 3 Configure OSPF TE.
                 To create a TE tunnel on LSR1 and LSR3, configure OSPF TE on LSR1, LSR2, and
                 LSR3.
                 # Configure LSR1.
                 [LSR1] ospf
                 [LSR1-ospf-1] opaque-capability enable
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] mpls-te enable
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 4 Create an MPLS TE tunnel on the ingress.
                 Create an MPLS TE tunnel interface on LSR1 and LSR3, and configure and apply
                 an explicit path.
                 # Configure an explicit path on LSR1.
                 [LSR1] explicit-path pri-path
                 [LSR1-explicit-path-pri-path] next hop 10.1.1.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.2.2
                 [LSR1-explicit-path-pri-path] next hop 3.3.3.9
                 [LSR1-explicit-path-pri-path] quit

