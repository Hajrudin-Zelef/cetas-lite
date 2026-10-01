---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-149
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [21396, 21567]
sha256: f6f4ca3f82e47ee0a03345348b73d9d547a2de038cc57e357b07f9b95b4268c7
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


                 ●      Run the display mpls te link-administration srlg-information [ interface
                        port-type port-num ] command to check the SRLGs to which an interface
                        belongs.

4.19.3 Example for Configuring an SRLG in a CR-LSP Hot
Standby Scenario

Networking Requirements
                 On the network shown in Figure 4-31, an RSVP-TE tunnel is established from PE1
                 to PE2 over the path PE1 -> P4 -> PE2. The link PE1 -> P1 -> P2 -> P4 and the link
                 PE1 -> P4 belong the same SRLG (SRLG1). Configure CR-LSP hot standby to
                 improve reliability. The primary and backup paths should not belong to the same
                 SRLG.

                 Figure 4-31 Network diagram of an SRLG in a CR-LSP hot standby scenario




Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Configure IP addresses for interfaces, including the loopback interfaces whose
                        addresses are to be used as MPLS LSR IDs.
                 2.     Configure IS-IS to ensure that nodes can reach each other over public network
                        routes.
                 3.     Enable MPLS, MPLS TE, and MPLS RSVP-TE on all nodes and interfaces
                        involved.
                 4.     Configure IS-IS TE on all nodes.
                 5.     Configure a dynamic primary MPLS TE tunnel from PE1 to PE2 over the
                        explicit path PE1 -> P4 -> PE2.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          358
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


                 6.     Set the SRLG number on the member interfaces of the SRLG.
                 7.     Configure an SRLG-based path calculation mode on the ingress of the tunnel.
                 8.     Configure CR-LSP hot standby.

Procedure
         Step 1 Configure interface IP addresses for the devices.
                 # Configure P1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname P1
                 [P1] vlan batch 100 200
                 [P1] interface vlanif 100
                 [P1-Vlanif100] ip address 10.1.1.2 30
                 [P1-Vlanif100] quit
                 [P1] interface vlanif 200
                 [P1-Vlanif200] ip address 10.2.1.1 30
                 [P1-Vlanif200] quit
                 [P1] interface 10ge 1/0/1
                 [P1-10GE1/0/1] port link-type trunk
                 [P1-10GE1/0/1] port trunk allow-pass vlan 100
                 [P1-10GE1/0/1] quit
                 [P1] interface 10ge 1/0/2
                 [P1-10GE1/0/2] port link-type trunk
                 [P1-10GE1/0/2] port trunk allow-pass vlan 200
                 [P1-10GE1/0/2] quit
                 [P1] interface loopback 1
                 [P1-loopback1] ip address 1.1.1.1 32
                 [P1-loopback1] quit

                 The configurations of P2, P3, P4, PE1, and PE2 are similar to the configuration of
                 P1. For detailed configurations, see Configuration Scripts.
         Step 2 Configure IS-IS to advertise routes.
                 # Configure P1.
                 [P1] isis 1
                 [P1-isis-1] network-entity 10.0000.0000.0001.00
                 [P1-isis-1] quit
                 [P1] interface vlanif 100
                 [P1-Vlanif100] isis enable 1
                 [P1-Vlanif100] quit
                 [P1] interface vlanif 200
                 [P1-Vlanif200] isis enable 1
                 [P1-Vlanif200] quit
                 [P1] interface loopback 1
                 [P1-LoopBack1] isis enable 1
                 [P1-LoopBack1] quit

                 The configurations of P2, P3, P4, PE1, and PE2 are similar to the configuration of
                 P1. For detailed configurations, see Configuration Scripts.
         Step 3 Configure basic MPLS functions and enable MPLS TE, RSVP-TE, and CSPF.
                 Enable MPLS, MPLS TE, and RSVP-TE globally on each node and on all interfaces
                 along the tunnel. Enable CSPF on the tunnel ingress PE1.
                 # Configure P1.
                 [P1] mpls lsr-id 1.1.1.1
                 [P1] mpls
                 [P1-mpls] mpls te
                 [P1-mpls] mpls rsvp-te


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                       359
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration

                 [P1-mpls] quit
                 [P1] interface vlanif 100
                 [P1-Vlanif100] mpls
                 [P1-Vlanif100] mpls te
                 [P1-Vlanif100] mpls rsvp-te
                 [P1-Vlanif100] quit
                 [P1] interface vlanif 200
                 [P1-Vlanif200] mpls
                 [P1-Vlanif200] mpls te
                 [P1-Vlanif200] mpls rsvp-te
                 [P1-Vlanif200] quit

                 The configurations of P2, P3, P4, PE1, and PE2 are similar to the configuration of
                 P1. For detailed configurations, see Configuration Scripts.

         Step 4 Configure IS-IS TE.

                 # Configure P1.
                 [P1] isis 1
                 [P1-isis-1] cost-style wide
                 [P1-isis-1] traffic-eng level-1-2
                 [P1-isis-1] quit

                 The configurations of P2, P3, P4, PE1, and PE2 are similar to the configuration of
                 P1. For detailed configurations, see Configuration Scripts.

         Step 5 Configure an explicit path for the primary tunnel on its ingress PE1.

                 # Configure an explicit path for the primary tunnel on PE1.
                 <PE1> system-view
                 [PE1] explicit-path main
                 [PE1-explicit-path-main] next hop 10.3.1.2
                 [PE1-explicit-path-main] next hop 10.6.1.2
                 [PE1-explicit-path-main] next hop 6.6.6.6
                 [PE1-explicit-path-main] quit

                 # After the configuration is complete, check information about the explicit path.
                 [PE1] display explicit-path main
                 Path Name : main        Path Status : Enabled
                  1    10.3.1.2      Strict    Include
                  2    10.6.1.2      Strict    Include
                  3    6.6.6.6       Strict   Include

         Step 6 Configure MPLS TE tunnel interfaces.

                 # On PE1, configure an MPLS TE tunnel from PE1 to PE2 and specify the explicit
                 path and tunnel bandwidth.
                 [PE1] interface tunnel1
                 [PE1-Tunnel1] ip address unnumbered interface loopback 1
                 [PE1-Tunnel1] tunnel-protocol mpls te
                 [PE1-Tunnel1] destination 6.6.6.6
                 [PE1-Tunnel1] mpls te tunnel-id 100
                 [PE1-Tunnel1] mpls te path explicit-path main
                 [PE1-Tunnel1] mpls te bandwidth ct0 10000

                 # After the configuration is complete, check the tunnel status.
                 [PE1] display interface tunnel1
                 Tunnel1 current state : UP (ifindex: 26)
                 Line protocol current state : UP
                 ...


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                        360
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


