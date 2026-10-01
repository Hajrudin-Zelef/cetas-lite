---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-153
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [22043, 22214]
sha256: af9a0eb2f051241f8d73c8ee392fe74830a5540b2ddf33860fc20edde66b7f9b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

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
                 [PE1-10GE1/0/1] quit
                 [PE1] interface loopback 1
                 [PE1-loopback1] ip address 4.4.4.4 32
                 [PE1-loopback1] quit

                 The configurations of P1, PE2, and P2 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.
         Step 2 Configure IS-IS to advertise routes.
                 # Configure PE1.
                 [PE1] isis 1
                 [PE1-isis-1] network-entity 10.0000.0000.0004.00
                 [PE1-isis-1] quit
                 [PE1] interface vlanif 100
                 [PE1-Vlanif100] isis enable 1
                 [PE1-Vlanif100] quit
                 [PE1] interface loopback 1
                 [PE1-LoopBack1] isis enable 1
                 [PE1-LoopBack1] quit

                 The configurations of P1, PE2, and P2 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.
         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.
                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the tunnel ingress PE1 and the PLR.
                 # Configure PE1.
                 [PE1] mpls lsr-id 4.4.4.4
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

                 The configurations of P1, PE2, and P2 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.
         Step 4 Configure IS-IS TE.
                 # Configure PE1.
                 [PE1] isis 1
                 [PE1-isis-1] cost-style wide


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                       368
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration

                 [PE1-isis-1] traffic-eng level-1-2
                 [PE1-isis-1] quit

                 The configurations of P1, PE2, and P2 are similar to the configuration of PE1. For
                 detailed configurations, see Configuration Scripts.
         Step 5 Configure an explicit path for the primary tunnel on its ingress PE1.
                 # Configure an explicit path for the primary tunnel on PE1.
                 <PE1> system-view
                 [PE1] explicit-path main
                 [PE1-explicit-path-main] next hop 10.1.1.2
                 [PE1-explicit-path-main] next hop 10.2.1.2
                 [PE1-explicit-path-main] next hop 5.5.5.5
                 [PE1-explicit-path-main] quit

                 # After the configuration is complete, check information about the explicit path of
                 the primary tunnel.
                 [PE1] display explicit-path main
                 Path Name : main        Path Status : Enabled
                  1    10.1.1.2      Strict    Include
                  2    10.2.1.2      Strict    Include
                  3    5.5.5.5       Strict   Include

         Step 6 Configure a tunnel interface for the primary MPLS TE tunnel.
                 # On PE1, configure a primary MPLS TE tunnel from PE1 to PE2 and specify the
                 explicit path and tunnel bandwidth.
                 [PE1] interface tunnel1
                 [PE1-Tunnel1] ip address unnumbered interface loopback 1
                 [PE1-Tunnel1] tunnel-protocol mpls te
                 [PE1-Tunnel1] destination 5.5.5.5
                 [PE1-Tunnel1] mpls te tunnel-id 100
                 [PE1-Tunnel1] mpls te path explicit-path main
                 [PE1-Tunnel1] mpls te bandwidth ct0 10000

                 # After the configuration is complete, check the tunnel status.
                 [PE1] display interface tunnel1
                 Tunnel1 current state : UP (ifindex: 26)
                 Line protocol current state : UP
                 ...

                 The command output shows that the tunnel is up. The preceding information is
                 only a part of the command output. The ... symbol is an ellipsis, indicating that
                 some information is not presented.
         Step 7 Configure an SRLG.
                 # Add the links with network segment addresses 10.2.1.0/30 and 10.5.1.0/30 to
                 SRLG1 on P1.
                 [P1] interface vlanif 200
                 [P1-Vlanif200] mpls te srlg 1
                 [P1-Vlanif200] quit
                 [P1] interface vlanif 500
                 [P1-Vlanif500] mpls te srlg 1
                 [P1-Vlanif500] quit

                 # Configure an SRLG-based path calculation mode on P1.
                 [P1] mpls
                 [P1-mpls] mpls te srlg path-calculation preferred
                 [P1-mpls] quit


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        369
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


                 # After the configuration is complete, check SRLG information and SRLG member
                 interfaces. The following example uses the command output on P1.
                 [P1] display mpls te srlg 1
                  SRLG 1:            Vlanif200          Vlanif500

                 # Check the SRLGs to which an interface belongs. The following example uses the
                 command output on P1.
                 [P1] display mpls te link-administration srlg-information

                  SRLGs on Vlanif200:
                          1

                  SRLGs on Vlanif500:
                          1

                 # Check SRLG TEDB information. The following example uses the command
                 output on P1.
                 [P1] display mpls te cspf tedb srlg 1
                 Interface-Address IGP-Type          Area
                 10.2.1.1       ISIS          1
                 10.5.1.1       ISIS          1
                 10.2.1.1       ISIS          2
                 10.5.1.1       ISIS          2

         Step 8 Configure TE auto FRR
                 # Enable TE Auto FRR on VLANIF200 of P1.
                 [P1] interface vlanif 200
                 [P1-Vlanif200] mpls te auto-frr link
                 [P1-Vlanif200] quit

                 # Enable TE FRR on the tunnel interface of the primary tunnel's ingress PE1.
                 [PE1] interface tunnel1
                 [PE1-Tunnel1] mpls te fast-reroute

                 # After the configuration is complete, check information about Tunnel1 on PE1.
                 [PE1] display mpls te tunnel path Tunnel1
                  Tunnel Interface Name : Tunnel1
                  Lsp ID : 5.5.5.5 :100 :1
                  Hop Information
                   Hop 0 10.1.1.1
                   Hop 1 10.1.1.2 Label 65536
                   Hop 2 1.1.1.1 Label 65536
                   Hop 3 10.2.1.1 Local-Protection available
                   Hop 4 10.2.1.2 Label 3
                   Hop 5 5.5.5.5 Label 3

