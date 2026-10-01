---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-84
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [11743, 11849]
sha256: 7ffd12f6b06b3fce6d5a5af2ec4e04fe9b33e87bf39b6e6ee9bf417ce360e058
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The process of establishing a dynamic MPLS TE tunnel involves the following
                 steps:
                 1.     The information advertisement component advertises the attributes of each
                        link on each node, enabling all nodes in the local area to acquire information
                        about the links on other nodes and use the information to create a TEDB.
                 2.     The path calculation component on the ingress uses the CSPF algorithm to
                        calculate an optimal path that meets tunnel attributes based on the
                        information stored in the TEDB.
                 3.     The path establishment component uses RSVP-TE to reserve resources and
                        allocate labels based on the optimal path calculated by CSPF. The component
                        establishes a CR-LSP and associates it with a tunnel interface.
                 To implement dynamic MPLS TE, a network administrator only needs to configure
                 tunnel attributes and link attributes according to service requirements and
                 network planning. MPLS TE will then automatically create tunnels in line with the
                 configurations specified by the network administrator.

4.2.3 RSVP-TE Message Format
                 RSVP is designed for the integrated services model, reserving resources, such as
                 bandwidth, for nodes along a specified path. The ability to reserve bandwidth
                 makes RSVP-TE an ideal signaling protocol for establishing MPLS TE paths.
                 RSVP-TE introduces the following extensions to RSVP to facilitate the
                 implementation of MPLS TE:
                 ●      RSVP-TE adds LABEL_REQUEST objects to Path messages to request labels,
                        and LABEL objects to Resv messages to allocate labels.
                 ●      An extended RSVP message can convey path constraints, in addition to label
                        binding information.
                 ●      The extended objects carry MPLS TE bandwidth constraints to implement
                        resource reservation.

RSVP-TE Message
                 An RSVP-TE message includes a common header, followed by multiple variable-
                 length and variable-type objects. Figure 4-2 shows the RSVP message format.

                 Figure 4-2 RSVP-TE message format




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                            200
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 For the description of each field, see 4.2.3 RSVP-TE Message Format.

                 Table 4-2 RSVP-TE message format
                  Field           Length        Description

                  Vers            4 bits        Indicates the RSVP version number. Currently, the
                                                value is 1.

                  Flags           4 bits        Indicates whether summary refresh (Srefresh) is
                                                supported. It is extended by RFC 2961, and its value
                                                is generally 0. If Srefresh is supported, the Flags
                                                field is set to 0x01.

                  Message         8 bits        Indicates the RSVP-TE message type. The value 1
                  Type                          indicates a Path message, and 2 indicates a Resv
                                                message.

                  RSVP            16 bits       Indicates the RSVP-TE checksum. If the value is 0,
                  Checksum                      the checksum is not checked during message
                                                transmission.

                  Send TTL        8 bits        Indicates the message TTL. When a node receives
                                                an RSVP-TE message, it compares the Send_TTL and
                                                the TTL in the IP header to calculate the number of
                                                hops that the message has passed in non-RSVP-TE
                                                areas.

                  Reserved        8 bits        Reserved.

                  RSVP            16 bits       Indicates the total length of an RSVP-TE message, in
                  Length                        bytes.

                  Objects         Variable      Indicates the objects in an RSVP-TE message. Each
                                                RSVP-TE message contains multiple objects. The
                                                carried objects vary according to the types of
                                                messages.

                  Length          16 bits       Indicates the total length of objects, in bytes. The
                  (bytes)                       value must be a multiple of 4, and the smallest
                                                value is 4.

                  Class-Num       8 bits        Identifies an object class. Each object class has a
                                                name, such as SESSION, SENDER_TEMPLATE, or
                                                TIME_VALUE.

                  C-Type          8 bits        Indicates an object type. Class-Number and C-Type
                                                together identify an object.

                  Object          Variable      Indicates the content of an object.
                  contents



                 RSVP-TE uses the following message types:
                 ●      Path message: used to request downstream nodes to distribute labels. A Path
                        message records path information on each node that the message passes

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                               201
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration


