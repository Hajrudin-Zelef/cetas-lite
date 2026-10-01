---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-36
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [4629, 4745]
sha256: 816a1d9ba4849c8f38c84bd89ffc12e5dc868be2e7db90751e5602fe6fad7e7e
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Label Retention Modes
                 A label retention mode defines how an LSR processes label mappings from non-
                 preferred next hops. The label mappings that an LSR receives may or may not
                 originate from the next hop. There are two label retention modes:
                 ●      Liberal mode: An LSR retains the label mappings received from a neighbor
                        LSR, regardless of whether the neighbor LSR is its next hop.
                 ●      Conservative mode: An LSR retains the label mappings received from a
                        neighbor LSR only when the neighbor LSR is its next hop.
                 When the next hop of an LSR changes due to a network topology change:
                 ●      In liberal mode, the LSR can use the labels advertised by a non-next hop LSR
                        to quickly reestablish an LSP. This mode, however requires more memory and
                        label space than the conservative mode. An LSP that has been assigned a
                        label but fails to be established is called a liberal LSP.
                 ●      In conservative mode, the LSR retains the labels advertised by the next hop
                        only. This mode saves memory and label space but takes more time to
                        reestablish an LSP. The conservative label retention mode is usually used
                        together with DoD on the LSRs that have limited label space.

Outbound and Inbound LDP Policies
                 Generally, LSRs do not have specific requirements for receiving or sending Label
                 Mapping messages. As a result, a large number of LSPs may be established. When
                 excessive amounts of system resources are consumed on an LSR, the system fails
                 to run stably. To address this issue, an outbound or inbound LDP policy can be
                 configured to limit the number of Label Mapping messages to be sent or received,
                 reducing the number of LSPs to be established and memory consumption.
                 ●      Outbound LDP policy
                        An outbound LDP policy filters Label Mapping messages to be sent. The
                        outbound LDP policy does not take effect for Label Mapping messages of an
                        L2VPN, which means that all Label Mapping messages of the L2VPN can be
                        sent. In addition, the ranges of FECs to which routes are mapped can be
                        configured.
                        If FECs in the Label Mapping messages to be sent to an LDP peer group or all
                        LDP peers are in the same range, the same outbound policy applies to the
                        LDP peer group or all LDP peers.

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                              78
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration


                        In addition, the outbound LDP policy supports split horizon. After split horizon
                        is configured, an LSR distributes labels only to its upstream LDP peer.
                        An LSR checks whether an outbound policy mapped to the route type
                        (labeled BGP route or non-BGP route) is configured before sending a Label
                        Mapping message for a FEC.
                        –   If no outbound policy is configured, the LSR sends the Label Mapping
                            message.
                        –   If an outbound policy is configured, the LSR checks whether the FEC in
                            the Label Mapping message is within the range defined in the outbound
                            policy. If it is within the FEC range, the LSR sends the Label Mapping
                            message for the FEC; if it is outside the FEC range, the LSR does not send
                            the Label Mapping message.
                 ●      Inbound LDP policy
                        An inbound LDP policy filters Label Mapping messages to be received. The
                        inbound LDP policy does not take effect for Label Mapping messages of an
                        L2VPN, which means that all Label Mapping messages of the L2VPN can be
                        received. In addition, the range of FECs to which non-BGP routes are mapped
                        can be configured.
                        If FECs in the Label Mapping messages to be received by an LDP peer group
                        or all LDP peers are in the same range, the same inbound policy applies to
                        the LDP peer group or all LDP peers.
                        An LSR checks whether an inbound policy mapped to a FEC is configured
                        before receiving a Label Mapping message for the FEC.
                        –   If no inbound policy is configured, the LSR receives the Label Mapping
                            message.
                        –   If an inbound policy is configured, the LSR checks whether the FEC in the
                            Label Mapping message is within the range defined in the inbound policy.
                            If it is within the FEC range, the LSR receives the Label Mapping message
                            for the FEC; if it is outside the FEC range, the LSR does not receive the
                            Label Mapping message.
                        If the FEC fails to pass an inbound policy on an LSR, the LSR receives no Label
                        Mapping message for the FEC.
                        One of the following results may occur:
                        –   If a DU LDP session is established between an LSR and its peer, a liberal
                            LSP is established. This liberal LSP cannot function as a backup LSP after
                            LDP FRR is enabled.
                        –   If a DoD LDP session is established between an LSR and its peer, the LSR
                            sends a Label Release message to tear down the label binding.

3.10.2 Configuring PHP
Context
                 The penultimate hop popping (PHP) function can be enabled after you configure
                 the label to be assigned to the penultimate hop.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            79
MPLS Configuration
MPLS Configuration                                                                     3 MPLS LDP Configuration


Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS view.
                 mpls

         Step 3 Configure the label to be assigned to the penultimate hop.
                 label advertise { explicit-null | implicit-null | non-null }

                 The default configuration (implicit-null) is recommended because it can reduce
                 the forwarding pressure on the egress and improve forwarding efficiency. You can
                 also choose a value as required.
                 ●      If implicit-null is specified, PHP is supported, allowing the penultimate hop to
                        pop the label from an MPLS packet, which reduces the overhead on the
                        egress. After receiving the packet, the egress directly forwards it over IP or
                        based on the next-layer label.
                 ●      If non-null is specified, PHP is not supported. This parameter setting may
                        cause high resource consumption on the egress and therefore is not
                        recommended. This parameter can be specified when the egress needs to
                        identify services based on labels.
                         NOTE

