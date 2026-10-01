---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-129
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [18259, 18434]
sha256: 08f0bc5cb129957042167ab5e8a0116b8c231e5df829b2636afb70b2d0ab7e13
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 300
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 00.0005.0000.0000.0004.00
                         traffic-eng level-2
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                         isis enable 1



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      306
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration

                        #
                        return

                 ●      LSR5
                        #
                        sysname LSR5
                        #
                        vlan batch 400 500
                        #
                        mpls lsr-id 5.5.5.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        isis 1
                         is-level level-2
                         cost-style wide
                         network-entity 00.0005.0000.0000.0005.00
                         traffic-eng level-2
                        #
                        interface Vlanif400
                         ip address 10.1.4.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/4
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 5.5.5.9 255.255.255.255
                         isis enable 1
                        #
                        return



4.11 Configuring an RSVP-TE GR Helper

4.11.1 Understanding RSVP-TE GR
                 RSVP-TE graceful restart (GR) ensures uninterrupted data transmission in the
                 forwarding plane when an active/standby switchover is performed in the control
                 plane in the case of a node failure.
                 RSVP-TE GR is a fast state recovery mechanism of RSVP-TE. As a high reliability
                 technology, RSVP-TE GR is designed based on the concept of non-stop forwarding
                 (NSF). If a fault occurs in the control plane of a node, the upstream and
                 downstream neighbors send messages to restore the RSVP-TE soft state of the
                 node. The forwarding plane is unaware of the fault and is not affected by the
                 fault, ensuring traffic stability and reliability.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      307
MPLS Configuration
MPLS Configuration                                                                    4 MPLS TE Configuration


Context
                 GR typically applies to PEs, particularly in scenarios where users connect to the
                 backbone network through a single point, such as PE3 in Figure 4-19. If PE3 fails
                 or encounters an active/standby switchover for maintenance purposes (such as a
                 version upgrade), traffic may be interrupted. However, with an MPLS TE tunnel
                 deployed to implement traffic engineering or function as a VPN public network
                 tunnel, you can configure RSVP-TE GR on this PE to ensure uninterrupted
                 forwarding of key services.

                 Figure 4-19 RSVP-TE GR application




Related Concepts
                 GR involves the GR restarter and GR helper, which performs a GR and helps
                 another device to perform a GR, respectively.
                 RSVP-TE GR messages are mainly classified as follows:
                 ●      Hello message with GR extensions: is used to detect the GR status of a
                        neighbor.
                 ●      GR Path message: is sent from an upstream node and carries the content of
                        the last Path Refresh message.
                 ●      Recovery Path message: is sent from a downstream node and carries the
                        content of the last Path message received by the downstream node.

                         NOTE

                        A device can only function as a GR helper to help a neighbor complete an RSVP-TE GR.


Implementation
                 In an RSVP-TE GR, the RSVP-TE Hello extension is used to detect the GR status of a
                 neighbor. For details about the Hello feature, see RSVP-TE Hello.
                 An RSVP-TE GR is implemented as follows:
                 On the network shown in Figure 4-20, after the GR restarter performs a GR due
                 to a fault, it stops sending Hello messages to neighbors. If a GR helper does not
                 receive Hello messages for three consecutive times, it considers that its neighbor is

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     308
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


                 performing a GR and retains all forwarding information. In addition, it continues
                 to transmit services and waits for the GR restarter to complete the GR.
                 If the GR restarter starts and receives Hello messages from the GR helpers, it also
                 sends Hello messages to them. In this case, the upstream and downstream GR
                 helpers process the received messages in different ways.
                 ●      When the upstream GR helper receives the Hello message, it sends a GR Path
                        message downstream to the GR restarter.
                 ●      When the downstream GR helper receives the Hello message, it sends a
                        Recovery Path message upstream to the GR restarter.

                 Figure 4-20 RSVP-TE GR implementation




                 After receiving the GR Path and Recovery Path messages, the GR restarter
                 reestablishes the path state block (PSB) and reservation state block (RSB) of the
                 CR-LSP based on the two messages. In this way, the local control plane
                 information is restored.
                 If the downstream GR helper cannot send Recovery Path messages, the GR
                 restarter reestablishes the local PSB and RSB merely based on GR Path messages.

