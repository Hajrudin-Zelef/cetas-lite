---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-38
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [4875, 5003]
sha256: e4ac6a412da49ad5d6412f8d43bd9cabfc39aa1fb49676b16af7067b926cd1dd
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      Configure an LDP split horizon policy.
                        DSLAMs deployed on an MPLS network for user access have low performance.
                        If LDP distributes labels to all peers, a large number of LSPs will be
                        established, which may result in large overhead on devices. It is recommended
                        that split horizon be configured for LDP peers, so that an LSR distributes
                        labels only to upstream LDP peers.
                 ●      Configure an inbound LDP policy.
                        Configure an inbound LDP policy to restrict the receiving of Label Mapping
                        messages.
                 ●      Configure an outbound LDP policy.
                        Configure an outbound LDP policy to restrict the sending of Label Mapping
                        messages.


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            82
MPLS Configuration
MPLS Configuration                                                                           3 MPLS LDP Configuration


Procedure
                 ●      Configure an LDP split horizon policy.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Configure a split horizon policy for an LDP peer to allow the LSR to
                             distribute labels only to the upstream LDP peer.
                             outbound peer { peer-id | all } split-horizon

                             By default, split horizon is not configured for LDP peers. That is, an LSR
                             distributes labels to both upstream and downstream LDP peers.
                 ●      Configure an inbound LDP policy.
                        a.   Enter the system view.
                             system-view

                        b.   Enter the MPLS-LDP view.
                             mpls ldp

                        c.   Enter the MPLS-LDP-IPv4 view.
                             ipv4-family

                        d.   Configure an inbound LDP policy for specified IGP routes of a specified
                             peer.
                             inbound peer { peer-id | peer-group peer-group-name | all } fec { none | host | ip-prefix prefix-
                             name }

                             An inbound LDP policy restricts the receiving of LDP Label Mapping
                             messages based on the selected parameter:

                             ▪    none: filters out all FECs. If this parameter is set, the specified peer
                                  does not receive Label Mapping messages on any IGP route.

                             ▪    host: allows only the FECs on host routes to pass. If this parameter is
                                  set, the specified peer receives Label Mapping messages on host
                                  routes.

                             ▪    ip-prefix: allows only the FECs on routes in a specified IP prefix list. If
                                  this parameter is set, the specified peer receives Label Mapping
                                  messages on IGP routes in the specified IP prefix list.
                             To apply a policy associated with the same FEC range to an LDP peer
                             group or all LDP peers receiving Label Mapping messages, specify either
                             peer-group peer-group-name or all in the command.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                                 83
MPLS Configuration
MPLS Configuration                                                                          3 MPLS LDP Configuration


                                   NOTE

                                  If multiple inbound policies are configured for a specified peer, the earliest
                                  configuration takes effect. For example, the following configurations are
                                  performed, in this sequence:
                                  inbound peer 2.2.2.2 fec host
                                  inbound peer peer-group group1 fec none
                                  As group1 also contains an LDP peer with peer-id of 2.2.2.2, the following
                                  inbound policy takes effect:
                                  inbound peer 2.2.2.2 fec host
                                  If two inbound policies are configured one after the other and the peer
                                  parameter settings in the two commands are the same, the latter configuration
                                  overwrites the former. For example, the following configurations are performed,
                                  in this sequence:
                                  inbound peer 2.2.2.2 fec host
                                  inbound peer 2.2.2.2 fec none
                                  The second configuration overwrites the first one. This means that the following
                                  inbound policy takes effect for the LDP peer with peer-id of 2.2.2.2:
                                  inbound peer 2.2.2.2 fec none
                                  If an inbound policy for all peers is configured and another inbound policy for a
                                  specified peer or peer group is configured, the former policy has a higher priority,
                                  and the latter policy does not take effect. For example, the following
                                  configurations are performed:
                                  inbound peer all fec none
                                  inbound peer 2.2.2.2 fec host
                                  The following inbound policy takes effect:
                                  inbound peer all fec none
                                  MPLS and MPLS LDP must be enabled globally before an inbound policy is
                                  configured.
                                  To delete all inbound policies simultaneously, run the undo inbound peer all
                                  command.
                 ●      Configure an outbound LDP policy.
                        a.   Enter the system view.
                             system-view
                        b.   Enter the MPLS-LDP view.
                             mpls ldp
                        c.   (Optional) Enter the MPLS-LDP-IPv4 view.
                             ipv4-family
                        d.   Configure an outbound LDP policy for specified IGP routes of a specified
                             peer.
                             outbound peer { peer-id | peer-group peer-group-name | all } fec { none | host | ip-prefix
                             prefix-name }

                             The outbound LDP policy helps an LSR provide a specific function based
                             on one of the following parameters:

                             ▪    none: filters out all FECs. If this parameter is specified, the device
                                  does not send Label Mapping messages for IGP routes to specified
                                  peers.

                             ▪    host: allows only the FECs on host routes to pass. If this parameter is
                                  specified, the device sends Label Mapping messages only for host
                                  routes to specified peers.

