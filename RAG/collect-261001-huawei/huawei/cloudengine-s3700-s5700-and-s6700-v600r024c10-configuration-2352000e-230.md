---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-230
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [33876, 34022]
sha256: b2d84b6b15354833c38cedfa5d77197c0ee4d0f90c51dbb7ae1e895dcd467324
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                    The following example uses the command output on PE1.
                    <PE1> display ip routing-table
                    Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                    ------------------------------------------------------------------------------
                    Routing Table: Public
                            Destinations : 15       Routes : 15
                    Destination/Mask Proto Pre Cost                Flags NextHop           Interface
                           1.1.1.1/32 Direct 0 0               D 127.0.0.1        InLoopBack0
                          2.2.2.2/32 OSPF 10 2                D 10.2.1.2       10GE1/0/2
                          3.3.3.3/32 OSPF 10 2                D 10.4.1.2       10GE1/0/3
                          4.4.4.4/32 OSPF 10 3                D 10.2.1.2       10GE1/0/2
                          5.5.5.5/32 OSPF 10 3                D 10.4.1.2       10GE1/0/3
                           10.2.1.0/30 Direct 0 0                D 10.2.1.1        10GE1/0/2
                           10.2.1.1/32 Direct 0 0                D 127.0.0.1       InLoopBack0
                           10.2.1.2/32 Direct 0 0                D 10.2.1.2        10GE1/0/2
                           10.3.1.0/30 OSPF        10 2           D 10.2.1.2        10GE1/0/2
                           127.0.0.0/8 Direct 0 0                D 127.0.0.1      InLoopBack0
                           127.0.0.1/32 Direct 0 0               D 127.0.0.1       InLoopBack0
                           10.4.1.0/30 Direct 0 0                D 10.4.1.1        10GE1/0/3
                           10.4.1.1/32 Direct 0 0                D 127.0.0.1       InLoopBack0
                           10.4.1.2/32 Direct 0 0                D 10.4.1.2        10GE1/0/3
                           10.5.1.0/30 OSPF 10 2                  D 10.4.1.2        10GE1/0/3

         Step 3 Configure basic MPLS functions on the MPLS backbone network.

                    # Enable MPLS and configure the Loopback1 interface IP address as an LSR ID.
                    Enable MPLS and MPLS LDP on interfaces.

                    # Configure PE1.
                    [PE1] mpls lsr-id 1.1.1.1
                    [PE1] mpls
                    [PE1-mpls] quit
                    [PE1] mpls ldp
                    [PE1-mpls-ldp] quit
                    [PE1] interface 10ge 1/0/2
                    [PE1-10GE1/0/2] mpls
                    [PE1-10GE1/0/2] mpls ldp
                    [PE1-10GE1/0/2] quit
                    [PE1] interface 10ge 1/0/3
                    [PE1-10GE1/0/3] mpls
                    [PE1-10GE1/0/3] mpls ldp
                    [PE1-10GE1/0/3] quit


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                                          540
VPN Configuration
VPN Configuration                                                                            5 VPWS Configuration


                    # Configure P1.
                    [P1] mpls lsr-id 2.2.2.2
                    [P1] mpls
                    [P1-mpls] quit
                    [P1] mpls ldp
                    [P1-mpls-ldp] quit
                    [P1] interface 10ge 1/0/1
                    [P1-10GE1/0/1] mpls
                    [P1-10GE1/0/1] mpls ldp
                    [P1-10GE1/0/1] quit
                    [P1] interface 10ge 1/0/2
                    [P1-10GE1/0/2] mpls
                    [P1-10GE1/0/2] mpls ldp
                    [P1-10GE1/0/2] quit

                    # Configure P2.
                    [P2] mpls lsr-id 3.3.3.3
                    [P2] mpls
                    [P2-mpls] quit
                    [P2] mpls ldp
                    [P2-mpls-ldp] quit
                    [P2] interface 10ge 1/0/1
                    [P1-10GE1/0/1] mpls
                    [P1-10GE1/0/1] mpls ldp
                    [P1-10GE1/0/1] quit
                    [P2] interface 10ge 1/0/2
                    [P1-10GE1/0/2] mpls
                    [P1-10GE1/0/2] mpls ldp
                    [P1-10GE1/0/2] quit

                    # Configure PE2.
                    [PE2] mpls lsr-id 4.4.4.4
                    [PE2] mpls
                    [PE2-mpls] quit
                    [PE2] mpls ldp
                    [PE2-mpls-ldp] quit
                    [PE2] interface 10ge 1/0/2
                    [PE2-10GE1/0/2] mpls
                    [PE2-10GE1/0/2] mpls ldp
                    [PE2-10GE1/0/2] quit

                    # Configure PE3.
                    [PE3] mpls lsr-id 5.5.5.5
                    [PE3] mpls
                    [PE3-mpls] quit
                    [PE3] mpls ldp
                    [PE3-mpls-ldp] quit
                    [PE3] interface 10ge 1/0/1
                    [PE3-10GE1/0/1] mpls
                    [PE3-10GE1/0/1] mpls ldp
                    [PE310GE1/0/1] quit

                    After the configurations are complete, run the display tunnel all command on
                    PEs. The command outputs show that an MPLS LSP has been established between
                    PE1 and PE2 and between PE1 and PE3.
                    The following example uses the command output on PE1.
                    <PE1> display tunnel all
                    Tunnel ID                Type            Destination          Status
                    ----------------------------------------------------------------------
                    0x000000000300000000             ldp             2.2.2.2           UP
                    0x000000000300000001             ldp             --               UP
                    0x000000000300000002             ldp             3.3.3.3           UP
                    0x000000000300000003             ldp             --               UP


Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                              541
VPN Configuration
VPN Configuration                                                                                     5 VPWS Configuration

                    0x000000000300000004 ldp                   4.4.4.4          UP
                    0x000000000300000005  ldp                    --              UP
                    0x000000000300000006 ldp                   5.5.5.5          UP
                    0x000000000300000007  ldp                    --              UP

                    Run the display mpls ldp session command on PEs. The command outputs show
                    that the Status field displays Operational, indicating that an LDP peer
                    relationship has been established between the PEs and their neighboring P.
                    The following example uses the command output on PE1.
                    <PE1> display mpls ldp session
                     LDP Session(s) in Public Network
                     Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDDD:HH:MM)
                     An asterisk (*) before a session means the session is being deleted.
                     ------------------------------------------------------------------------------
                     PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                     ------------------------------------------------------------------------------
                     2.2.2.2:0        Operational DU Passive 000:00:03 16/16
                     3.3.3.3:0        Operational DU Passive 000:00:03 13/13
                     ------------------------------------------------------------------------------
                     TOTAL: 2 session(s) Found.

         Step 4 Establish a remote LDP session between PEs.
                    # Configure a remote LDP session, and use the loopback interface address of a
                    remote LDP peer as the remote peer IP address.

                           NOTE

