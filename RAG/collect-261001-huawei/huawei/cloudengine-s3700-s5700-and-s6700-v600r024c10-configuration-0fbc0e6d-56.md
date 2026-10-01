---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-56
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [7717, 7851]
sha256: 2920087562aec55ae230e0247cb02179ddf1396edeb1513b1f8612e39417756f
---

                 # Run the display mpls ldp lsp command to check whether an LDP LSP destined
                 for 4.4.4.4/32 has been established on PE1.
                 <PE1> display mpls ldp lsp
                  LDP LSP Information
                  -------------------------------------------------------------------------------
                  Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                  -------------------------------------------------------------------------------
                  DestAddress/Mask In/OutLabel UpstreamPeer NextHop                             OutInterface
                  -------------------------------------------------------------------------------
                  1.1.1.1/32        3/NULL          2.2.2.2        127.0.0.1        LoopBack1
                 *1.1.1.1/32         Liberal/21                  DS/2.2.2.2
                  2.2.2.2/32        NULL/3          -            10.1.1.2         Vlanif100
                  2.2.2.2/32        16/3          2.2.2.2        10.1.1.2         Vlanif100
                  4.4.4.4/32        NULL/22          -            10.1.1.2         Vlanif100
                  4.4.4.4/32        17/22          2.2.2.2        10.1.1.2        Vlanif100
                  -------------------------------------------------------------------------------
                  TOTAL: 5 Normal LSP(s) Found.
                  TOTAL: 1 Liberal LSP(s) Found.
                  TOTAL: 0 Frr LSP(s) Found.
                  An asterisk (*) before an LSP means the LSP is not established
                  An asterisk (*) before a Label means the USCB or DSCB is stale
                  An asterisk (*) before an UpstreamPeer means the session is stale
                  An asterisk (*) before a DS means the session is stale
                  An asterisk (*) before a NextHop means the LSP is FRR LSP

         Step 3 Enable BFD globally on the two ends of a link to be monitored.

                 # Configure PE1.
                 <PE1> system-view
                 [PE1] bfd
                 [PE1-bfd] quit

                 # Configure PE2.
                 <PE2> system-view
                 [PE2] bfd
                 [PE2-bfd] quit

         Step 4 On the ingress, configure a BFD session and bind it to the LDP LSP.

                 # Configure PE1.
                 <PE1> system-view
                 [PE1] bfd 1to4 bind ldp-lsp peer-ip 4.4.4.4 nexthop 10.1.1.2 interface Vlanif100
                 [PE1-bfd-lsp-session-1to4] discriminator local 1
                 [PE1-bfd-lsp-session-1to4] discriminator remote 2
                 [PE1-bfd-lsp-session-1to4] process-pst
                 [PE1-bfd-lsp-session-1to4] quit


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                              130
MPLS Configuration
MPLS Configuration                                                                                  3 MPLS LDP Configuration


         Step 5 On the egress, configure a BFD session and bind it to the IP link, enabling the
                egress to notify the ingress of LDP LSP faults.
                 # Configure PE2.
                 <PE2> system-view
                 [PE2] bfd 4to1 bind peer-ip 1.1.1.1
                 [PE2-bfd-session-4ot1] discriminator local 2
                 [PE2-bfd-session-4ot1] discriminator remote 1
                 [PE2-bfd-session-4ot1] quit

                 ----End

Result
                 # After completing the configuration, run the display bfd session all verbose
                 command on the ingress. The State field displays Up, and the BFD Bind Type field
                 displays LDP_LSP.
                 <PE1> display bfd session all verbose
                 (w): State in WTR
                 (*): State is invalid
                 --------------------------------------------------------------------------------
                   State : Up                   Name : 1to4
                 --------------------------------------------------------------------------------
                   Local Discriminator : 1                  Remote Discriminator : 2
                   Session Detect Mode : Asynchronous Mode Without Echo Function
                   BFD Bind Type            : LDP_LSP
                   Bind Session Type         : Static
                   Bind Peer IP Address : 4.4.4.4
                   NextHop Ip Address          : 10.1.1.2
                   Bind Interface         : Vlanif100
                   FSM Board Id             :-            TOS-EXP                :7
                   Min Tx Interval (ms) : 1000                Min Rx Interval (ms) : 1000
                   Actual Tx Interval (ms): 1000              Actual Rx Interval (ms): 1000
                   WTR Interval (ms)          :-            Detect Interval (ms) : -
                   Local Detect Multi        :3            Active Multi           :-
                   Echo Passive           : Disable         Acl Number              :-
                   Destination Port         :-            TTL                 :1
                   Proc Interface Status : Disable            Process PST            : Enable
                   Config PST            : Enable
                   Active Multi          :3
                   Last Local Diagnostic : No Diagnostic
                   Bind Application         : No Application Bind
                   Session Not Up Reason : In negotiation
                   Session Description : -

                 Run the display bfd session all verbose command on the egress. The (Multi
                 Hop) State field displays Up, and the BFD Bind Type field displays Peer IP
                 Address.
                 <PE2> display bfd session all verbose
                 (w): State in WTR
                 (*): State is invalid
                 --------------------------------------------------------------------------------
                   (Multi Hop) State : Up                     Name : 4to1
                 --------------------------------------------------------------------------------
                   Local Discriminator : 2                  Remote Discriminator : 1
                   Session Detect Mode : Asynchronous Mode Without Echo Function
                   BFD Bind Type           : Peer IP Address
                   Bind Session Type        : Static
                   Bind Peer IP Address : 1.1.1.1
                   Bind Interface         :-
                   Track Interface        :-
                   Bind Source IP Address : 4.4.4.4
                   FSM Board Id            :3             TOS-EXP                :7
                   Min Tx Interval (ms) : 1000                Min Rx Interval (ms) : 1000


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                            131
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration

                  Actual Tx Interval (ms): 1000           Actual Rx Interval (ms): 1000
                  Local Detect Multi      :3           Detect Interval (ms) : 3000
                  Echo Passive         : Disable        Acl Number           :-
                  Destination Port       : 3784         TTL               : 254
                  Proc Interface Status : Disable         Process PST         : Disable
                  Config PST          : Disable
                  Last Local Diagnostic : No Diagnostic
                  Bind Application       : No Application Bind
                  Session Not Up Reason : In negotiation
                  Session Description : -


