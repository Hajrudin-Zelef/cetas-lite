---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-61
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [8470, 8574]
sha256: db68976029009732a96e753858a497b81ba8819fb57bd9c3c1c47908fa7300de
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 --------------------------------------------------------------------------------
                   (Multi Hop) State : Up                    Name : dyn_16390
                 --------------------------------------------------------------------------------
                   Local Discriminator : 16390                Remote Discriminator : 16387
                   Session Detect Mode : Asynchronous Mode Without Echo Function
                   BFD Bind Type            : Peer IP Address
                   Bind Session Type         : Entire_Dynamic
                   Bind Peer IP Address : 4.4.4.4
                   Bind Interface         :-
                   Track Interface        :-
                   Bind Source IP Address : 1.1.1.1
                   FSM Board Id            :3             TOS-EXP                :7
                   Min Tx Interval (ms) : 1000                Min Rx Interval (ms) : 1000
                   Actual Tx Interval (ms): 1000              Actual Rx Interval (ms): 600
                   WTR Interval (ms)          :-            Detect Interval (ms) : 2400
                   Local Detect Multi        :3            Active Multi           :4
                   Destination Port        : 4784           TTL                 : 253
                   Proc Interface Status : Disable
                   Last Local Diagnostic : No Diagnostic
                   Bind Application        : LDP
                   Session Not Up Reason : In negotiation
                   Session Description : -



Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            142
MPLS Configuration
MPLS Configuration                                                                                  3 MPLS LDP Configuration


                 # Run the display bfd session passive-dynamic verbose command on PE2 to
                 check the status of the passively created BFD session. The BFD session status is
                 Up. The BFD Bind Type field displays Peer IP Address, indicating that PE2 sends
                 BFD packets over IP routes.
                 [PE2] display bfd session passive-dynamic verbose
                 (w): State in WTR
                 (*): State is invalid
                 --------------------------------------------------------------------------------
                   (Multi Hop) State : Up                    Name : dyn_16386
                 --------------------------------------------------------------------------------
                   Local Discriminator : 16386                Remote Discriminator : 16388
                   Session Detect Mode : Asynchronous Mode Without Echo Function
                   BFD Bind Type             : Peer IP Address
                   Bind Session Type          : Entire_Dynamic
                   Bind Peer IP Address : 1.1.1.1
                   Bind Interface         :-
                   Track Interface        :-
                   Bind Source IP Address : 4.4.4.4
                   FSM Board Id             :3            TOS-EXP                :7
                   Min Tx Interval (ms) : 10                 Min Rx Interval (ms) : 10
                   Actual Tx Interval (ms): 100               Actual Rx Interval (ms): 600
                   Local Detect Multi         :3            Detect Interval (ms) : 2400
                   Echo Passive           : Disable         Acl Number               :-
                   Destination Port         : 4784           TTL                : 253
                   Proc Interface Status : Disable            Process PST             : Disable
                   WTR Interval (ms)           :-           Config PST             : Disable
                   Active Multi          :4
                   Last Local Diagnostic : No Diagnostic
                   Bind Application         : No Application Bind
                   Session TX TmrID            :-           Session Detect TmrID : -
                   Session Init TmrID        :-            Session WTR TmrID           :-
                   Session Echo Tx TmrID : -
                   Session Description : -
                 --------------------------------------------------------------------------------

                     Total UP/DOWN Session Number : 1/0


Configuration Scripts
                 ●      PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100 200
                        #
                        bfd
                         mpls-passive
                        #
                        mpls lsr-id 1.1.1.1
                        #
                        mpls
                         lsp-trigger all
                         mpls bfd enable
                         mpls bfd-trigger fec-list l1
                         mpls bfd min-tx-interval 100 min-rx-interval 600 detect-multiplier 4
                        #
                        fec-list l1
                         fec-node 4.4.4.4
                        #
                        mpls ldp
                         #
                         ipv4-family
                          auto-frr lsp-trigger all
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            143
MPLS Configuration
MPLS Configuration                                                                              3 MPLS LDP Configuration

