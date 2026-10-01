---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-86
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11985, 12127]
sha256: 93c1ebe996b0d5c73d6a180726121b1fed8fc4595918795698cb5ef0a36883b6
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                  TIME_VALU        5       1                       Interval at which the message is
                  E                                                sent.

                  STYLE            8       1                       Resource reservation style, which is
                                                                   specified on the ingress.

                  FLOW_SPEC        9       ● 1: Reserved           QoS parameters of a data flow.
                                             (obsolete)
                                             flowspec object
                                           ● 2: Inv-serv
                                             flowspec object

                  FILTER_SPEC      10      1                       IP address and LSP ID of the node
                                                                   that sends the message.

                  RECORD_RO        21      1                       RRO that records the actual path of
                  UTE                                              an LSP

                  LABEL            16      1                       Label allocated by a downstream
                                                                   node to an upstream node

                  RESV_CONF        15      1                       IP address of the node that
                  IRM                                              requests resource reservation
                                                                   confirmation




                 A reservation style carried in a Resv message defines how an RSVP-TE node
                 reserves resource after receiving a request sent by an upstream node. Common
                 reservation styles are as follows:

                 ●      Fixed Filter (FF) style: reserves exclusive resources for each node that requests
                        resource reservation. This means that reserved resources vary according to CR-
                        LSPs on the same link.
                 ●      Shared Explicit (SE) style: reserves shared resources for a series of nodes that
                        request resource reservation. This means that reserved resources are shared
                        by certain CR-LSPs on the same link.

4.2.4 RSVP-TE Fundamentals
                 RSVP-TE can dynamically establish and maintain CR-LSPs, mainly including:


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             204
MPLS Configuration
MPLS Configuration                                                               4 MPLS TE Configuration


                 ●      Path establishment: CR-LSPs are dynamically established based on the path
                        calculation results of CSPF. Path establishment is initiated by tunnel ingresses.
                 ●      Path maintenance: After CR-LSPs are established, RSVP-TE sends messages to
                        maintain path status on each node.
                 ●      Error notification: If a message processing error or fault occurs during path
                        establishment or maintenance, an RSVP-TE node notifies its upstream and
                        downstream nodes of the error.
                 ●      Path teardown: A CR-LSP is torn down, and labels and bandwidth related to
                        the CR-LSP are released on each node involved. Path teardown is also
                        initiated by tunnel ingresses.


Path Establishment
                 To establish a dynamic CR-LSP using RSVP-TE, the ingress sends a Path message
                 to the egress, and the egress sends a Resv message to the ingress. The Path
                 message is used to create an RSVP-TE session and associate path status. Each
                 node that receives the Path message creates a PSB. The Resv message carries
                 resource reservation information. Each node that receives the Resv message
                 creates an RSB and allocates a label.

                 Figure 4-3 shows how RSVP-TE dynamically establishes a CR-LSP.

                 Figure 4-3 Dynamic CR-LSP establishment process




                 1.     PE1 triggers CSPF to calculate a path from PE1 to PE2. The IP address of each
                        hop along the path is specified. PE1 generates an ERO with the IP address of
                        each hop and adds the ERO in a Path message. PE1 then creates a PSB and
                        sends the Path message to P1 according to information in the ERO. Table 4-5
                        describes the objects carried in the Path message.

                        Table 4-5 Path message on PE1

                         Object                                   Value

                         SESSION                                  Source: PE1-if1; Destination: PE2-if0

                         RSVP_HOP                                 PE1-if1

                         EXPLICIT_ROUTE                           P1-if0; P2-if0; PE2-if0

                         LABEL                                    LABEL_REQUEST


                 2.     After receiving the Path message from PE1, P1 parses the message and
                        constructs a PSB based on the Path message. Then P1 updates the Path
                        message and sends the message to P2 according to the ERO. Table 4-6
                        describes the objects in the Path message.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            205
MPLS Configuration
MPLS Configuration                                                                     4 MPLS TE Configuration


                        –   The RSVP_HOP object specifies the IP address of the outbound interface
                            through which a Path message is sent. Therefore, PE1 sets the RSVP_HOP
                            object to the IP address of the outbound interface toward P1, and P1 sets
                            the RSVP_HOP field to the IP address of the outbound interface toward
                            P2.
                        –   P1 deletes its own inbound and outbound interface addresses and LSR ID
                            to update the ERO of the Path message.

                        Table 4-6 Path message on P1
                         Object                                      Value

                         SESSION                                     Source: PE1-if1; Destination: PE2-if0

                         RSVP_HOP                                    P1-if1

                         EXPLICIT_ROUTE                              P2-if0; PE2-if0

                         LABEL                                       LABEL_REQUEST


                 3.     After receiving the Path message from P1, P2 creates a PSB according to the
                        Path message, updates the Path message, and sends the message to PE2
                        according to the ERO field. Table 4-7 describes the objects in the Path
                        message.

                        Table 4-7 Path message on P2
                         Object                                      Value

                         SESSION                                     Source: PE1-if1; Destination: PE2-if0

                         RSVP_HOP                                    P2-if1

                         EXPLICIT_ROUTE                              PE2-if0

                         LABEL                                       LABEL_REQUEST


