---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-124
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17467, 17622]
sha256: 0b3f7615c37ed34f2dc52502cdf91edb0c9867f89a1d9413fcbf8c1746ad44c1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 Configure the handshake function so that LSR1 and LSR2 can perform RSVP-TE
                 key authentication on each other to prevent forged RSVP-TE resource reservation
                 requests from occupying network resources. In addition, configure the message
                 window function to prevent RSVP-TE packet disorder.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                              294
MPLS Configuration
MPLS Configuration                                                                        4 MPLS TE Configuration


                 Figure 4-17 Network diagram for configuring RSVP-TE authentication




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign IP addresses to interfaces on nodes and configure OSPF to ensure that
                        public network routes between the nodes are reachable.
                 2.     Configure LSR IDs and enable MPLS, MPLS TE, and MPLS RSVP-TE globally on
                        the nodes and interfaces.
                 3.     Create a tunnel interface on the ingress, specify the tunnel IP address, tunnel
                        protocol, destination address, tunnel ID, and dynamic signaling protocol RSVP-
                        TE, and enable CSPF.
                 4.     Configure RSVP-TE message authentication on LSR1 and LSR2.
                 5.     Configure the handshake function on LSR1 and LSR2 to prevent forged RSVP-
                        TE resource reservation requests from occupying network resources.
                 6.     Configure the message window function on LSR1 and LSR2 to prevent RSVP-
                        TE packet disorder.

                         NOTE

                        You are advised to set the window size to a value greater than 32. If the size is too small,
                        some received RSVP-TE messages may be beyond the window and are discarded. As a
                        result, the RSVP-TE neighbor relationship is terminated.


Procedure
         Step 1 Assign IP addresses to interfaces on nodes and configure OSPF.
                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] interface eth-trunk 1
                 [LSR1] undo portswitch
                 [LSR1-Eth-Trunk1] ip address 10.1.1.1 255.255.255.0
                 [LSR1-Eth-Trunk1] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] eth-trunk 1
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface 10ge 1/0/2
                 [LSR1-10GE1/0/2] eth-trunk 1
                 [LSR1-10GE1/0/2] quit
                 [LSR1] interface 10ge 1/0/3
                 [LSR1-10GE1/0/3] eth-trunk 1
                 [LSR1-10GE1/0/3] quit
                 [LSR1] interface loopback 1
                 [LSR1-LoopBack1] ip address 1.1.1.1 255.255.255.255
                 [LSR1-LoopBack1] quit


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                       295
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                 [LSR1] ospf 1
                 [LSR1-ospf-1] area 0
                 [LSR1-ospf-1-area-0.0.0.0] network 1.1.1.1 0.0.0.0
                 [LSR1-ospf-1-area-0.0.0.0] network 10.1.1.0 0.0.0.255
                 [LSR1-ospf-1-area-0.0.0.0] quit
                 [LSR1-ospf-1] quit

                 The configurations of LSR2 and LSR3 are similar to the configuration of LSR1. For
                 detailed configurations, see Configuration Scripts.

                 After the configuration is complete, run the display ip routing-table command on
                 each node. The command output shows that the nodes have learned routes from
                 each other.

         Step 2 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.

                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the ingress.

                 # Configure LSR1.
                 [LSR1] mpls lsr-id 1.1.1.1
                 [LSR1] mpls
                 [LSR1-mpls] mpls te
                 [LSR1-mpls] mpls rsvp-te
                 [LSR1-mpls] mpls te cspf
                 [LSR1-mpls] quit
                 [LSR1] interface eth-trunk 1
                 [LSR1-Eth-Trunk1] mpls
                 [LSR1-Eth-Trunk1] mpls te
                 [LSR1-Eth-Trunk1] mpls rsvp-te
                 [LSR1-Eth-Trunk1] quit

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

         Step 4 Configure MPLS TE tunnel interfaces.

                 # On LSR1, configure an MPLS TE tunnel to LSR3.
                 [LSR1] interface Tunnel 1
                 [LSR1-Tunnel1] ip address unnumbered interface loopback 1
                 [LSR1-Tunnel1] tunnel-protocol mpls te
                 [LSR1-Tunnel1] destination 3.3.3.3
                 [LSR1-Tunnel1] mpls te tunnel-id 1
                 [LSR1-Tunnel1] quit

                 # After the configuration is complete, run the display interface tunnel command
                 on LSR1. The command output shows that the tunnel interface status is up.
                 [LSR1] display interface tunnel
                 Tunnel1 current state : UP


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                       296
MPLS Configuration
MPLS Configuration                                                                      4 MPLS TE Configuration

                 Line protocol current state : UP
                 ...

         Step 5 Configure RSVP-TE authentication on the interfaces of the MPLS TE link between
                LSR1 and LSR2.
                 # Configure LSR1.
                 [LSR1] interface eth-trunk 1
                 [LSR1-Eth-Trunk1] mpls rsvp-te authentication cipher YsHsjx_202206
                 [LSR1-Eth-Trunk1] mpls rsvp-te authentication handshake
                 [LSR1-Eth-Trunk1] mpls rsvp-te authentication window-size 32
                 [LSR1-Eth-Trunk1] quit

                 # Configure LSR2.
                 [LSR2] interface eth-trunk 1
                 [LSR2-Eth-Trunk1] mpls rsvp-te authentication cipher YsHsjx_202206
                 [LSR2-Eth-Trunk1] mpls rsvp-te authentication handshake
                 [LSR2-Eth-Trunk1] mpls rsvp-te authentication window-size 32
                 [LSR2-Eth-Trunk1] quit

                 ----End

