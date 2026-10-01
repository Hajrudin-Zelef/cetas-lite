---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-35
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [4501, 4628]
sha256: a4f24ff6486889f63e58c8756ce8740662f50fa3b632add3050c403e683eb08e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        #
                        vlan batch 100
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                        #
                        mpls ldp
                         #
                          ipv4-family
                        #
                        interface Vlanif100
                         ip address 192.168.3.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 192.168.3.0 0.0.0.255
                        #
                        return



3.10 Configuring LDP Label Advertisement and
Management Modes

3.10.1 Understanding LDP Label Advertisement and
Management Modes
                 After an LDP session is established, LDP starts to exchange messages, such as
                 Label Mapping messages, to establish LSPs. Related standards have defined label
                 advertisement, distribution control, and retention modes to determine how LSRs
                 advertise and manage labels.

                 The device currently supports the following combinations of modes:

                 ●      Downstream unsolicited (DU) label advertisement + ordered label distribution
                        control + liberal label retention
                 ●      Downstream on demand (DoD) label advertisement + ordered label
                        distribution control + conservative label retention

                 The default label advertisement and management modes are DU label
                 advertisement + ordered label distribution control + liberal label retention.


Label Advertisement Modes
                 In MPLS, downstream LSRs distribute labels to specific FECs and notify their
                 upstream LSRs of the mappings between labels and FECs. That is, labels are
                 designated by downstream LSRs and advertised in the downstream-to-upstream
                 direction.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        76
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


                 The label advertisement modes on the upstream and downstream LSRs of an
                 adjacency must be the same. There are two label advertisement modes:

                 ●      DU mode: An LSR does not wait for a Label Request message from the
                        upstream LSR before advertising the label of a specific FEC.
                        In Figure 3-13, the downstream node (egress) triggers the establishment of
                        an LSP destined for FEC 192.168.1.1/32 in host mode by advertising the label
                        of the host route 192.168.1.1/32 over a Label Mapping message.

                        Figure 3-13 DU mode




                 ●      DoD mode: An LSR waits for a Label Request message before advertising the
                        label of a specific FEC.
                        In Figure 3-14, the downstream egress triggers the establishment of an LSP
                        destined for FEC 192.168.1.1/32 in host mode. The upstream node (ingress)
                        sends a Label Request message to the downstream node (egress). The
                        downstream egress sends a Label Mapping message to the upstream LSR only
                        after receiving the Label Request message.

                        Figure 3-14 DoD mode




Label Distribution Control Modes
                 A label distribution control mode defines how an LSR distributes labels during the
                 establishment of an LSP. There are two label distribution control modes:

                 ●      Independent mode: A local LSR distributes a label to a FEC and notifies the
                        upstream LSR of the FEC-label binding without waiting to receive a label from
                        a downstream LSR.
                        –   In Figure 3-13, if the label advertisement mode is DU and the label
                            distribution control mode is independent, the transit LSR distributes labels
                            to the upstream ingress without waiting to receive a label from the
                            downstream egress.
                        –   In Figure 3-14, if the label advertisement mode is DoD and the label
                            distribution control mode is independent, the downstream transit LSR
                            directly connected to the ingress LSR that sends a Label Request message
                            replies with labels without waiting to receive a label from the
                            downstream egress.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                             77
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


                 ●      Ordered mode: An LSR advertises the mapping between a label and an FEC to
                        its upstream LSR only when this LSR has received a Label Mapping message
                        from the next hop of the FEC or the LSR is the egress of the FEC.
                        –   In Figure 3-13, if the label advertisement mode is DU and the label
                            distribution control mode is ordered, the transit LSR distributes a label to
                            the upstream ingress only after receiving a Label Mapping message from
                            the downstream egress.
                        –   In Figure 3-14, if the label advertisement mode is DoD and the label
                            distribution control mode is ordered, the transit LSR directly connected to
                            the ingress LSR that sends a Label Request message distributes a label to
                            the upstream ingress only after receiving a Label Mapping message from
                            the downstream egress.

