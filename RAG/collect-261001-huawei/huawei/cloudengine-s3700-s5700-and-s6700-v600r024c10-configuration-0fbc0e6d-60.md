---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-60
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [8344, 8469]
sha256: ca8cad4a30b2df32a74c35a5bce23c6d981a41bb0631949b3a6c35bb9e0c776c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Precautions
                 During the configuration, note the following:
                 ●      Each LSR must have route entries that exactly match FECs for the LSPs to be
                        established.
                 ●      By default, the triggering policy is host, allowing a device to use host IP
                        routes with 32-bit addresses to trigger LDP LSP establishment.
                 ●      If the triggering policy is all, all IGP routes are used to trigger LDP LSP
                        establishment. The device does not use public network BGP routes to trigger
                        LDP LSP establishment.

Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign an IP address to each interface of each LSR according to Figure 3-27.
                 2.     Configure OSPF to implement network layer connectivity Adjust interface
                        costs to have the path PE1 -> P1 -> PE2 become the primary path and the
                        path PE1 -> P2 -> PE2 become the backup path.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     140
MPLS Configuration
MPLS Configuration                                                                      3 MPLS LDP Configuration


                 3.     Enable each node to use all IGP routes to establish LDP LSPs.
                 4.     Configure OSPF FRR and LDP Auto FRR to generate a backup LSP between
                        PE1 and PE2.
                 5.     Configure dynamic BFD for LDP to use FEC lists to create BFD sessions.

Procedure
         Step 1 Assign an IP address to each interface.
                 Assign an IP address to each interface according to Figure 3-27 and create a
                 loopback interface on each node. For details, see the configuration scripts.
         Step 2 Configure OSPF.
                 Configure OSPF on each node to implement network layer connectivity. For
                 details, see the configuration scripts.
         Step 3 Configure LDP LSPs.
                 Configure MPLS LDP on each node and enable the nodes to use all IGP routes to
                 establish LDP LSPs. For details, see the configuration scripts.
         Step 4 Configure OSPF FRR and LDP Auto FRR.
                 # Configure PE1.
                 [PE1] ospf 1
                 [PE1-ospf-1] frr
                 [PE1-ospf-1-frr] loop-free-alternate
                 [PE1-ospf-1-frr] quit
                 [PE1-ospf-1] quit
                 [PE1] mpls ldp
                 [PE1-mpls-ldp] auto-frr lsp-trigger all
                 [PE1-mpls-ldp] quit

                 # Configure PE2.
                 [PE2] ospf 1
                 [PE2-ospf-1] frr
                 [PE2-ospf-1-frr] loop-free-alternate
                 [PE2-ospf-1-frr] quit
                 [PE2-ospf-1] quit
                 [PE2] mpls ldp
                 [PE2-mpls-ldp] auto-frr lsp-trigger all
                 [PE2-mpls-ldp] quit

         Step 5 Configure dynamic BFD sessions to monitor LDP LSPs.
                 # Enable BFD, specify a FEC list used to establish a BFD session, and set BFD
                 parameters on PE1.
                 [PE1] bfd
                 [PE1-bfd] mpls-passive
                 [PE1-bfd] quit
                 [PE1] fec-list l1
                 [PE1-fec-list-l1] fec-node 4.4.4.4
                 [PE1-fec-list-l1] quit
                 [PE1] mpls
                 [PE1-mpls] mpls bfd enable
                 [PE1-mpls] mpls bfd-trigger fec-list l1
                 [PE1-mpls] mpls bfd min-tx-interval 100 min-rx-interval 600 detect-multiplier 4
                 [PE1-mpls] quit

                 # Enable BFD, specify a FEC list used to establish a BFD session, and set BFD
                 parameters on PE2.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                  141
MPLS Configuration
MPLS Configuration                                                                                  3 MPLS LDP Configuration

                 [PE2] bfd
                 [PE2-bfd] mpls-passive
                 [PE2-bfd] quit
                 [PE2] fec-list l2
                 [PE2-fec-list-l2] fec-node 1.1.1.1
                 [PE2-fec-list-l2] quit
                 [PE2] mpls
                 [PE2-mpls] mpls bfd enable
                 [PE2-mpls] mpls bfd-trigger fec-list l2
                 [PE2-mpls] mpls bfd min-tx-interval 100 min-rx-interval 600 detect-multiplier 4
                 [PE2-mpls] quit

                 ----End


Result
                 # Run the display bfd session all verbose command to check the dynamic BFD
                 session status.
                 [PE1] display bfd session all verbose
                 (w): State in WTR
                 (*): State is invalid
                 --------------------------------------------------------------------------------
                   State : Up                    Name : dyn_16388
                 --------------------------------------------------------------------------------
                   Local Discriminator : 16388                Remote Discriminator : 16386
                   Session Detect Mode : Asynchronous Mode Without Echo Function
                   BFD Bind Type             : LDP_LSP
                   Bind Session Type          : Dynamic
                   Bind Peer IP Address : 4.4.4.4
                   NextHop Ip Address           : 10.1.1.2
                   Bind Interface         : Vlanif100
                   FSM Board Id             :3             TOS-EXP               :7
                   Min Tx Interval (ms) : 600                 Min Rx Interval (ms) : 100
                   Actual Tx Interval (ms): 600               Actual Rx Interval (ms): 100
                   WTR Interval (ms)           :-           Detect Interval (ms) : 300
                   Local Detect Multi         :4            Active Multi          :3
                   Destination Port         : 4784           TTL                :1
                   Proc Interface Status : Disable
                   Last Local Diagnostic : No Diagnostic
                   Bind Application         : LDP
                   Session Not Up Reason : In negotiation
                   Session Description : -
                 --------------------------------------------------------------------------------

