---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-122
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [17273, 17356]
sha256: 112d76a5f6a5ae20f06c55f74a5f7e6af40a4e153861951f0faa11e014f75a1b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                  Configure the handshake     mpls rsvp-te                If the local end
                  function.                   authentication              configured with the
                                              handshake                   handshake function
                                                                          receives an RSVP-TE
                                                                          message from a
                                                                          neighbor that does not
                                                                          establish an RSVP-TE
                                                                          authentication
                                                                          relationship with the
                                                                          local end, the local end
                                                                          sends a Challenge
                                                                          message carrying the
                                                                          local identifier to the
                                                                          neighbor. After receiving
                                                                          the Challenge message,
                                                                          the neighbor returns a
                                                                          Response message,
                                                                          which carries the
                                                                          identifier in the received
                                                                          Challenge message. If
                                                                          the identifier in the
                                                                          Response messages is
                                                                          the same as the local
                                                                          identifier, the local end
                                                                          determines to establish
                                                                          an RSVP-TE
                                                                          authentication
                                                                          relationship with the
                                                                          neighbor.




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                           291
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


                  Operation                   Command                     Purpose
                                                                          NOTE
                                                                           If the handshake function
                                                                           is configured between
                                                                           neighbors and the mpls
                                                                           rsvp-te authentication
                                                                           lifetime lifetime command
                                                                           needs to be run to
                                                                           configure an
                                                                           authentication lifetime, the
                                                                           lifetime must be longer
                                                                           than the interval for
                                                                           sending RSVP-TE Refresh
                                                                           messages, which is
                                                                           configured using the mpls
                                                                           rsvp-te timer refresh
                                                                           command. If the
                                                                           authentication lifetime is
                                                                           shorter than the interval
                                                                           for sending RSVP-TE
                                                                           Refresh messages, the
                                                                           authentication relationship
                                                                           is deleted because no
                                                                           RSVP-TE Refresh message
                                                                           is received within the
                                                                           authentication lifetime. As
                                                                           a result, the handshake
                                                                           mechanism is triggered
                                                                           again when a new Refresh
                                                                           message is received. This
                                                                           process repeats, causing a
                                                                           failure of TE tunnel
                                                                           establishment or deletion.




Issue 01 (2025-03-03)         Copyright © Huawei Technologies Co., Ltd.                             292
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                  Operation                    Command                      Purpose

