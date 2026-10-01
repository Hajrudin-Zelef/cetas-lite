---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-106
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [14920, 15028]
sha256: d63e6a48fa374aa0a8829ef6bf43ebd4d97a4d0ade1755bd137dfe8202e40084
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                       1.1.1.9/32 Direct 0 0            D 127.0.0.1                        LoopBack1
                       2.2.2.9/32 OSPF 10 15             D 10.1.1.2                         Vlanif100
                       3.3.3.9/32 OSPF 10 20             D 10.1.4.2                         Vlanif400
                      10.1.1.0/24 Direct 0 0            D 10.1.1.1                         Vlanif100
                      10.1.1.1/32 Direct 0 0            D 127.0.0.1                        Vlanif100
                    10.1.1.255/32 Direct 0 0             D 127.0.0.1                         Vlanif100
                      10.1.2.0/24 OSPF 10 25              D 10.1.1.2                         Vlanif100
                      10.1.3.0/24 OSPF 10 20              D 10.1.4.2                         Vlanif400
                      10.1.4.0/24 Direct 0 0            D 10.1.4.1                         Vlanif400
                      10.1.4.1/32 Direct 0 0            D 127.0.0.1                        Vlanif400
                    10.1.4.255/32 Direct 0 0             D 127.0.0.1                         Vlanif400
                      10.1.5.0/24 OSPF 10 20              D 10.1.4.2                         Vlanif400
                     127.0.0.0/8 Direct 0 0             D 127.0.0.1                        InLoopBack0
                     127.0.0.1/32 Direct 0 0             D 127.0.0.1                        InLoopBack0
                 127.255.255.255/32 Direct 0 0             D 127.0.0.1                          InLoopBack0
                 255.255.255.255/32 Direct 0 0             D 127.0.0.1                          InLoopBack0

                 The command output shows that LSR1 has learned routes to other nodes.
         Step 2 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.
                 To create a TE tunnel from LSR1 to LSR3, enable MPLS, MPLS TE, and RSVP-TE
                 globally on LSR1, LSR2, and LSR3 and on all interfaces along the tunnel, and
                 enable CSPF on the ingress (LSR1).
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


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          253
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 3 Configure OSPF TE.
                 # Configure LSR1.
                 [LSR1] ospf
                 [LSR1-ospf-1] opaque-capability enable
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] mpls-te enable
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.
         Step 4 Create an MPLS TE tunnel on the ingress LSR1.
                 # Configure an explicit path for the tunnel.
                 [LSR1] explicit-path pri-path
                 [LSR1-explicit-path-pri-path] next hop 10.1.1.2
                 [LSR1-explicit-path-pri-path] next hop 10.1.2.2
                 [LSR1-explicit-path-pri-path] next hop 3.3.3.9
                 [LSR1-explicit-path-pri-path] quit

                 # Configure an MPLS TE tunnel and bind it to the explicit path.
                 [LSR1] interface tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 3.3.3.9
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] mpls te path explicit-path pri-path

         Step 5 Configure IGP shortcut.
                 Enable IGP shortcut on the TE tunnel interface and set the IGP metric value to 10
                 (absolute metric value).
                 # Configure LSR1.
                 [LSR1-Tunnel1] mpls te igp shortcut ospf
                 [LSR1-Tunnel1] mpls te igp metric absolute 10
                 [LSR1-Tunnel1] quit
                 [LSR1] ospf 1
                 [LSR1-ospf-1] enable traffic-adjustment
                 [LSR1-ospf-1] quit

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


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         254
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 The command output shows that the next hop of the route to LSR3 (3.3.3.9) is
                 1.1.1.9, the outbound interface is Tunnel1, and the traffic to LSR3 is steered to the
                 TE tunnel.


