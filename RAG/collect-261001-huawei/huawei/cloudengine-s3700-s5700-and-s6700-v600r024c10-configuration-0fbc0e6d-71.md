---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-71
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [9946, 10086]
sha256: dcaf1cf51a9d81ba39ed9c70189513e5ceecb8aef25307c21e8a79f575fd28d1
---

                 # Configure LSRB.
                 [LSRB] mpls lsr-id 2.2.2.9
                 [LSRB] mpls
                 [LSRB-mpls] quit
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] quit

                 # Configure LSRC.
                 [LSRC] mpls lsr-id 3.3.3.9
                 [LSRC] mpls
                 [LSRC-mpls] quit
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] quit

         Step 3 Enable MPLS and MPLS LDP on each interface.
                 # Configure LSRA.
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] mpls ldp
                 [LSRA-Vlanif100] quit

                 # Configure LSRB.
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] mpls
                 [LSRB-Vlanif100] mpls ldp
                 [LSRB-Vlanif100] quit
                 [LSRB] interface vlanif 200
                 [LSRB-Vlanif200] mpls
                 [LSRB-Vlanif200] mpls ldp
                 [LSRB-Vlanif200] quit

                 # Configure LSRC.
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] mpls
                 [LSRC-Vlanif100] mpls ldp
                 [LSRC-Vlanif100] quit

                 After completing the preceding configuration, check whether a local LDP session is
                 successfully established between LSRA and LSRB and between LSRB and LSRC.
                 # Run the display mpls ldp session command on each node to check LDP session
                 information. The following example uses the command output on LSRA.
                 [LSRA] display mpls ldp session
                  LDP Session(s) in Public Network
                   Codes: LAM(Label Advertisement Mode), SsnAge Unit(DDD:HH:MM)
                   An asterisk (*) before a session means the session is being deleted.
                  --------------------------------------------------------------------------
                   PeerID            Status      LAM SsnRole SsnAge             KASent/Rcv
                  --------------------------------------------------------------------------
                  2.2.2.9:0        Operational DU Passive 000:00:02 9/9
                  --------------------------------------------------------------------------
                  TOTAL: 1 Session(s) Found.


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        168
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration


         Step 4 Enable LDP GR.
                 # Configure LSRA.
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] graceful-restart
                 Warning: All the related sessions will be deleted if the operation is performed!
                  Continue? [Y/N]:y
                 [LSRA-mpls-ldp] quit

                 # Configure LSRB.
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] graceful-restart
                 Warning: All the related sessions will be deleted if the operation is performed!
                  Continue? [Y/N]:y
                 [LSRB-mpls-ldp] quit

                 # Configure LSRC.
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] graceful-restart
                 Warning: All the related sessions will be deleted if the operation is performed!
                  Continue? [Y/N]:y
                 [LSRC-mpls-ldp] quit

                 ----End

Result
                 # After completing the preceding configuration, run the display mpls ldp session
                 verbose command. The command output shows that the value of the Session FT
                 Flag field is On. The following example uses the command output on LSRA.
                 [LSRA] display mpls ldp session verbose
                  LDP Session(s) in Public Network
                  ------------------------------------------------------------------------
                  Peer LDP ID       : 2.2.2.9:0         Local LDP ID : 1.1.1.9:0
                  TCP Connection : 1.1.1.9 <- 2.2.2.9
                  Session State : Operational              Session Role : Passive
                  Session FT Flag : On                    MD5 Flag        : Off
                  Reconnect Timer : 300 Sec                 Recovery Timer : 300 Sec
                  Keychain Name : kc1
                  Tcpao Name          : ---
                  Authentication applied : Peer-group list1

                  Negotiated Keepalive Hold Timer : 45 Sec
                  Configured Keepalive Send Timer : 30 Sec
                  Keepalive Message Sent/Rcvd       : 1/1 (Message Count)
                  Label Advertisement Mode         : Downstream Unsolicited
                  Label Resource Status(Peer/Local) : Available/Available
                  Session Age                : 0000:00:00 (DDDD:HH:MM)

                  Capability:
                   Capability-Announcement              : Off

                  Outbound Policies applied: NULL

                  Addresses received from peer: ( Count: 3 )
                  2.2.2.9         10.1.1.2         10.2.1.1
                  ------------------------------------------------------------------------

                 Run the display mpls ldp peer verbose command on each LSR. The command
                 output shows that the value of the Peer FT Flag field is On. The following
                 example uses the command output on LSRA.
                 [LSRA] display mpls ldp peer verbose
                  LDP Peer Information in Public network


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                      169
MPLS Configuration
MPLS Configuration                                                                                  3 MPLS LDP Configuration

                  -------------------------------------------------------------------------------
                  Peer LDP ID             : 2.2.2.9:0
                  Peer Max PDU Length : 4096                 Peer Transport Address : 2.2.2.9
                  Peer Loop Detection : Off                Peer Path Vector Limit : --
                  Peer FT Flag             : On          Peer Keepalive Timer : 45 Sec
                  Recovery Timer             : 300 Sec     Reconnect Timer          : 300 Sec
                  Peer Type             : Local
                  Peer Label Advertisement Mode : Downstream Unsolicited
                  Distributed ID          :0
                  Peer Discovery Source : Vlanif100
                  Capability-Announcement              : On
                  -------------------------------------------------------------------------------


