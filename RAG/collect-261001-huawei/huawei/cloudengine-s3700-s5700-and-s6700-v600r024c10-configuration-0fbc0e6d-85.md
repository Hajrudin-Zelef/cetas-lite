---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-85
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11850, 11984]
sha256: 693f4c8dfb1fb34880c17376b1641fcc9d687b85a22b9350c7f8ec5435985c6a
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        through. The path information is used to establish a path state block (PSB) on
                        a node.
                 ●      Resv message: used to reserve resources at each hop of a path. A Resv
                        message is transmitted along the reverse paths of a data flow. Each node that
                        receives the Resv message reserves resources based on reservation
                        information carried in the message. The reservation information is used to
                        establish a reservation state block (RSB) and to record information about
                        distributed labels.
                 ●      PathErr message: sent upstream by an RSVP-TE node if an error occurs during
                        the processing of a Path message. A PathErr message is forwarded upstream
                        by each transit node until it arrives at the ingress.
                 ●      ResvErr message: sent downstream by an RSVP-TE node if an error occurs
                        during the processing of a Resv message. After receiving the ResvErr message,
                        the transit node continues to forward the message to the downstream node
                        until the message reaches the egress.
                 ●      PathTear message: instructs a node to remove path information, performing
                        the reverse action of a Path message.
                 ●      ResvTear message: instructs a node to remove resource reservation status
                        information, performing the reverse action of a Resv message.
                 ●      ResvConf message: sent downstream from the sender hop by hop to confirm
                        a resource reservation request. It is sent only when the Resv message contains
                        the RESV_CONFIRM object.
                 ●      Srefresh message: used to update the RSVP-TE state.


Path Message
                 An RSVP-TE Path message is used to create an RSVP-TE session and associate path
                 status. A Path message travels from the ingress to the egress. Each node that
                 receives the message creates a PSB. The source IP address of a Path message is
                 the LSR ID of the ingress, and the destination IP address is the LSR ID of the
                 egress.

                 4.2.3 RSVP-TE Message Format lists some objects carried in a Path message.


                 Table 4-3 Path message objects

                  Message          Obje    Object Type           Object Content
                  Object           ct
                                   Class

                  SESSION          1       1                     RSVP-TE session information,
                                                                 including the destination address,
                                                                 tunnel ID, and extend tunnel ID.

                  RSVP-            3       1                     Address and index of the outbound
                  TE_HOP                                         interface of the previous hop that
                                                                 sends the Path message.

                  TIME_VALU        5       1                     Interval at which the message is
                  E                                              sent.


Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            202
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                  Message         Obje    Object Type            Object Content
                  Object          ct
                                  Class

                  SENDER_TE       11      1                      IP address and LSP ID of the node
                  MPLATE                                         that sends the message.

                  SENDER_TS       12      2                      Traffic characteristics of a data
                  PEC                                            flow.

                  LABEL_REQ       19      1                      Label request object, which is
                  UEST                                           carried only in a Path message.

                  ADSPEC          13      2                      Actual QoS parameters about a
                                                                 path, such as desired bandwidth,
                                                                 minimum path delay, and path
                                                                 maximum transmission unit (MTU).

                  EXPLICIT_R      20      1                      Explicit route object (ERO), carrying
                  OUTE                                           information about the path
                                                                 computed by the ingress. A Path
                                                                 message is forwarded over a path
                                                                 specified by the ERO, ignoring the
                                                                 shortest path determined by IGP.

                  RECORD_RO       21      1                      Record route object (RRO) that
                  UTE                                            records the actual path of an LSP

                  SESSION_AT      207     ● 1:                   MPLS TE tunnel attributes, such as
                  TRIBUTE                   LSP_TUNNEL_R         the setup priority, holding priority,
                                            A                    resource reservation style, and
                                          ● 7: LSP Tunnel        affinity.




Resv Message
                 After receiving a Path message, an egress returns a Resv message in response. A
                 Resv message carries resource reservation information and is sent from the egress
                 to the ingress hop by hop. After receiving the Resv message, each transit node
                 creates and maintains an RSB and allocates a label to the Resv message. When
                 the Resv message reaches the ingress, a CR-LSP is established successfully.

                 Table 4-4 lists objects carried in a Resv message.


                 Table 4-4 Resv message objects

                  Message         Obj     Object Type            Object Content
                  Object          ect
                                  Clas
                                  s

                  INTEGRITY       4       1                      Authentication key of an RSVP
                                                                 message.


Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             203
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


                  Message          Obj     Object Type             Object Content
                  Object           ect
                                   Clas
                                   s

                  SESSION          1       1                       RSVP-TE session information, such
                                                                   as the destination address, tunnel
                                                                   ID, and extend tunnel ID.

                  RSVP-            3       1                       IP address and index of the
                  TE_HOP                                           outbound interface at the previous
                                                                   hop that sends the Resv message.

