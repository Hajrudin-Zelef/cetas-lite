---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-10
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [719, 850]
sha256: 8d1be8707d268d6f3e27860032d52d620bdb9fb78698510750b0df979006aa50
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Label Space
                 A label space indicates a label value range. Label spaces can be categorized as
                 follows:
                 ●      Spaces for special labels. For details about special labels, see Table 2-1.
                 ●      Label spaces of dynamic signaling protocols, such as LDP
                        The label spaces of different dynamic signaling protocols are not shared;
                        instead, they are independent and continuous.

                 Table 2-1 Special labels
                  Label Value       Description              Description

                  0                 IPv4 explicit NULL       This label must be popped. Packets
                                    label                    carrying this label must be forwarded
                                                             based on IPv4. If the egress allocates a
                                                             label with the value 0 to the penultimate
                                                             LSR, the penultimate LSR pushes the label
                                                             onto the top of the packet label stack and
                                                             forwards the packet to the last hop. After
                                                             receiving the packet, the last hop pops the
                                                             label.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                                    8
MPLS Configuration
MPLS Configuration                                                          2 Basic MPLS Configuration


                  Label Value    Description             Description

                  1              Router Alert label      This label takes effect only when it is not
                                                         at the bottom of a label stack. If a node
                                                         receives a packet carrying the Router Alert
                                                         label (similar to the Router Alert Option
                                                         field in an IP packet), the node sends the
                                                         packet to its local software module for
                                                         processing. The node determines whether
                                                         to forward the packet based on the next-
                                                         layer label. If the packet needs to be
                                                         forwarded, the node pushes the Router
                                                         Alert label back onto the top of the label
                                                         stack.

                  2              IPv6 explicit NULL      This label must be popped. Packets
                                 label                   carrying this label must be forwarded
                                                         based on IPv6. If the egress allocates a
                                                         label with the value 2 to the penultimate
                                                         LSR, the penultimate LSR pushes the label
                                                         onto the top of the packet label stack and
                                                         forwards the packet to the last hop. After
                                                         receiving the packet, the last hop pops the
                                                         label.

                  3              Implicit NULL label     This label must be popped on the
                                                         penultimate LSR. Packets carrying this
                                                         label need to be forwarded to the last hop.
                                                         After receiving a packet carrying this label,
                                                         the last hop forwards the packet over IP or
                                                         based on the next-layer label.

                  4 to 13        Reserved                -

                  14             OAM Router Alert        Packets carrying this label are operation,
                                 label                   administration and maintenance (OAM)
                                                         packets for monitoring LSPs and notifying
                                                         LSP faults. OAM packets are carried over
                                                         MPLS. They are transparent to transit LSRs
                                                         and penultimate LSRs.

                  15             Reserved                -




Label Stack
                 A label stack contains an ordered set of labels. In the MPLS label stack, the label
                 next to the link layer header is the top or outer label, and the label next to the
                 Layer 3 header is the bottom or inner label. Theoretically, there is no limit to the
                 number of MPLS labels in a stack.




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                                   9
MPLS Configuration
MPLS Configuration                                                          2 Basic MPLS Configuration


                 Figure 2-5 Label stack




                 The labels are processed from the top of the label stack based on the last in, first
                 out rule.


Label Operation Types
                 MPLS involves the following label operations:

                 ●      Push: label add operation. When an IP packet enters an MPLS domain, the
                        ingress adds a label between the link layer header and the Layer 3 header of
                        the packet. When the packet reaches a transit node, the transit node can also
                        add a label to the top of the label stack (for label nesting) as needed.
                 ●      Swap: label replacement operation. When a packet is forwarded inside an
                        MPLS domain, a transit node searches the label forwarding information base
                        (LFIB) and replaces the label on the top of the stack in the MPLS packet with
                        the label that is assigned by the next hop.
                 ●      Pop: label removal operation. When a packet leaves an MPLS domain, the
                        egress removes the MPLS label. The penultimate MPLS node may also remove
                        the label on the top of the stack to reduce the number of labels in the stack.


Penultimate Hop Popping
                 Labels become unnecessary on the last hop. To reduce the workload of the last
                 hop, labels can be popped at the penultimate hop, a function known as
                 penultimate hop popping (PHP). In this case, when receiving a packet, the last
                 hop directly forwards it over IP or based on the next-layer label.

                 PHP can be configured on an egress. The PHP-enabled egress distributes only one
                 type of label to the penultimate hop — implicit NULL label.

                 The label value 3 indicates the implicit NULL label. This label value does not
                 appear in the label stack. When an implicit NULL label is assigned to an LSR, the
                 LSR pops the label, without replacing the label at the top of the stack with this
                 implicit-null label. The egress then forwards the packet over IP or based on the
                 next-layer label.


