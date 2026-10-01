---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-100
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2024-04-01", "2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [14060, 14171]
sha256: 8bf53a34a0fa30815ac32d1d1499623f7c5cd36118d5d24dc3f7f4e229517c2c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                       1.1.1.9/32 Direct 0 0            D 127.0.0.1                        LoopBack1
                       2.2.2.9/32 ISIS-L2 15 10          D 10.1.1.2                         Vlanif100
                       3.3.3.9/32 ISIS-L2 15 20          D 10.1.1.2                         Vlanif100
                       4.4.4.9/32 ISIS-L2 15 30          D 10.1.1.2                         Vlanif100
                      10.1.1.0/24 Direct 0 0            D 10.1.1.1                         Vlanif100
                      10.1.1.1/32 Direct 0 0            D 127.0.0.1                        Vlanif100
                    10.1.1.255/32 Direct 0 0             D 127.0.0.1                         Vlanif100
                      10.1.2.0/24 ISIS-L2 15 20          D 10.1.1.2                         Vlanif100
                      10.1.3.0/24 ISIS-L2 15 30          D 10.1.1.2                         Vlanif100
                     127.0.0.0/8 Direct 0 0             D 127.0.0.1                        InLoopBack0
                     127.0.0.1/32 Direct 0 0             D 127.0.0.1                        InLoopBack0
                 127.255.255.255/32 Direct 0 0             D 127.0.0.1                          InLoopBack0
                 255.255.255.255/32 Direct 0 0             D 127.0.0.1                          InLoopBack0

                 The command output shows that LSR1 has learned routes to other nodes.

         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.

                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the ingress.

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

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          239
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


         Step 4 Configure IS-IS TE.

                 # Configure LSR1.
                 [LSR1] isis 1
                 [LSR1-isis-1] cost-style wide
                 [LSR1-isis-1] traffic-eng level-2
                 [LSR1-isis-1] quit

                 The configurations of LSR2, LSR3, and LSR4 are similar to the configuration of
                 LSR1. For detailed configurations, see Configuration Scripts.

         Step 5 Set MPLS TE bandwidth attributes for links.

                 Configure the maximum reservable bandwidth and BC0 bandwidth for links on
                 each interface along the tunnel.

                 # Configure LSR1.
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] mpls te bandwidth max-reservable-bandwidth 100000
                 [LSR1-Vlanif100] mpls te bandwidth bc0 100000
                 [LSR1-Vlanif100] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.

         Step 6 Configure MPLS TE tunnel interfaces.

                 # On LSR1, configure an MPLS TE tunnel to LSR4.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 4.4.4.9
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] mpls te bandwidth ct0 20000
                 [LSR1-Tunnel1] quit

                 ----End

Verifying the Configuration
                 # Check the tunnel interface state on LSR1.
                 [LSR1] display interface tunnel 1
                 Tunnel1 current state : UP (ifindex: 29)
                 Line protocol current state : UP
                 Last line protocol up time : 2024-04-01 07:46:29
                 Description:
                 Route Port,The Maximum Transmit Unit is 1500, Current BW: 20Mbps
                 Internet Address is unnumbered, using address of LoopBack1(1.1.1.9/32)
                 Encapsulation is TUNNEL, loopback not set
                 Tunnel destination 4.4.4.9
                 Tunnel up/down statistics 5
                 Tunnel ct0 bandwidth is 20000 Kbit/sec
                 Tunnel protocol/transport MPLS/MPLS, ILM is available
                 primary tunnel id is 0xA001, secondary tunnel id is 0x0
                 Current system time: 2024-04-01 07:49:46
                    0 seconds output rate 0 bits/sec, 0 packets/sec
                    0 seconds output rate 0 bits/sec, 0 packets/sec
                    0 packets output, 0 bytes
                    0 output error
                    0 output drop
                    Last 300 seconds input utility rate: 0.00%
                    Last 300 seconds output utility rate: 0.00%


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                   240
MPLS Configuration
MPLS Configuration                                                                              4 MPLS TE Configuration


