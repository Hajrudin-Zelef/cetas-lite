---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-166
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "parameters"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [24028, 24209]
sha256: 2ed6255247423126b232a61641c72d2b520e90564b2b03b043acf7f07eb108d1
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      LSR4
                        #
                        sysname LSR4
                        #
                        vlan batch 300 400 500
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif400
                         ip address 10.1.4.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif500
                         ip address 10.1.5.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface 10GE1/0/2
                         port link-type trunk


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      399
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 4.4.4.9 0.0.0.0
                          network 10.1.3.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                          mpls-te enable
                        #
                        return




4.24 Configuring Static BFD for CR-LSP

4.24.1 Understanding Static BFD for CR-LSP
Definition
                 Bidirectional forwarding detection (BFD) is an end-to-end fast fault detection
                 mechanism. It can be implemented in MPLS TE to quickly detect faults in links
                 that a tunnel passes through, triggering a path switchover and enhancing service
                 reliability across the entire network. BFD for MPLS TE can be categorized into BFD
                 for RSVP, BFD for TE tunnel, and BFD for CR-LSP based on detection objects. This
                 section specifically covers BFD for CR-LSP.

Context
                 Conventional detection mechanisms, such as RSVP-TE Hello and Srefresh, detect
                 faults slowly. BFD, however, adopts the fast packet transmission mode and can
                 therefore quickly detect a fault on an MPLS TE tunnel, triggering fast traffic
                 switchover for protection.

                 Figure 4-36 Network diagram of BFD




                 On the network shown in Figure 4-36, if BFD is not enabled and LSR5 is faulty,
                 LSR1 and LSR6 cannot immediately detect the fault because a Layer 2 switch is

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       400
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


                 deployed. The Hello protocol can be used to detect the fault, but the detection
                 takes a long time.

                 After BFD is enabled, if LSR5 fails, LSR1 and LSR6 can quickly detect the fault and
                 switch data flows to the path LSR1 -> LSR2 -> LSR4 -> LSR6.


Related Concepts
                 BFD for CR-LSP is classified into static BFD for CR-LSP and dynamic BFD for CR-
                 LSP based on the establishment mode of BFD sessions.

                 ●      Static BFD for CR-LSP: Local and remote discriminators are manually
                        configured for BFD sessions.
                 ●      Dynamic BFD for CR-LSP: Local and remote discriminators are automatically
                        allocated by the system for BFD sessions.


Fundamentals
                 BFD for CR-LSP can rapidly detect faults on CR-LSPs and notify the forwarding
                 plane of the faults to ensure a fast traffic switchover. BFD for CR-LSP is usually
                 used together with CR-LSP hot standby.

                 BFD for CR-LSP is implemented as follows:

                 1.     Create a BFD session on the ingress and egress of a CR-LSP to bind the
                        session to the LSP. The ingress sends BFD packets to the egress along the CR-
                        LSP, and the egress responds with packets. This allows the ingress to promptly
                        assess the link status on the CR-LSP.
                 2.     If a link fault is detected, BFD notifies the forwarding plane of the fault. The
                        forwarding plane searches for a backup CR-LSP and switches traffic to the
                        backup CR-LSP. Then, the forwarding plane reports the fault information to
                        the control plane. A BFD session can be created for the backup CR-LSP based
                        on the used BFD mode.
                        –   If static BFD for CR-LSP is used and the backup CR-LSP needs to be
                            monitored, manually configure a BFD session for the backup CR-LSP.
                        –   If dynamic BFD for CR-LSP is used, the control plane proactively
                            establishes a BFD session for the backup CR-LSP.

4.24.2 Enabling BFD Globally

Prerequisites
                 Before configuring static BFD for CR-LSP, complete one of the following tasks:

                 ●      Configure a dynamic MPLS TE tunnel.
                 ●      Configure CR-LSP backup.


Context
                 Before configuring static BFD for CR-LSP, enable BFD globally on the ingress and
                 egress of a tunnel.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             401
MPLS Configuration
MPLS Configuration                                                                       4 MPLS TE Configuration


                         NOTE

                        When static BFD for CR-LSP is used and the BFD state is up, the BFD state remains up even
                        after the tunnel interface of the CR-LSP is shut down.

                 Perform the following configuration on the ingress and egress of an MPLS TE
                 tunnel.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enable BFD globally.
                 bfd

                 You can perform BFD-related configuration only after enabling BFD globally using
                 the bfd command.

                 ----End

4.24.3 Setting BFD Parameters on an Ingress
Context
                 BFD parameters that can be configured on a tunnel ingress include the local
                 discriminator, remote discriminator, local interval for sending BFD packets, local
                 interval for receiving BFD packets, and local BFD detection multiplier. These
                 parameters affect the establishment of a BFD session.
                 Perform the following configuration on the ingress of an MPLS TE tunnel.

