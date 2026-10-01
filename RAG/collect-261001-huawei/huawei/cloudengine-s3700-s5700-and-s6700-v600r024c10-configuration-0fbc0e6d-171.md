---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-171
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [24786, 24953]
sha256: daa39e89a66995844610d235a1f23d3e2cf2ab29abb779843d19c3d55c56d228
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                          mpls-te enable
                        #
                        bfd reversebac2lsra bind peer-ip 1.1.1.9
                         discriminator local 439
                         discriminator remote 339
                         min-tx-interval 500
                         min-rx-interval 500
                        #
                        bfd reversepri2lsra bind peer-ip 1.1.1.9
                         discriminator local 239
                         discriminator remote 139
                         min-tx-interval 500
                         min-rx-interval 500
                        #
                        return

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
                         port trunk allow-pass vlan 300
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 500
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 400
                        #
                        interface LoopBack1
                         ip address 4.4.4.9 255.255.255.255


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      412
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

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




4.25 Configuring Dynamic BFD for CR-LSP

4.25.1 Understanding Dynamic BFD for CR-LSP

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

                 Figure 4-38 Network diagram of BFD




                 On the network shown in Figure 4-38, if BFD is not enabled and LSR5 is faulty,
                 LSR1 and LSR6 cannot immediately detect the fault because a Layer 2 switch is
                 deployed. The Hello protocol can be used to detect the fault, but the detection
                 takes a long time.

                 After BFD is enabled, if LSR5 fails, LSR1 and LSR6 can quickly detect the fault and
                 switch data flows to the path LSR1 -> LSR2 -> LSR4 -> LSR6.

Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        413
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


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

Application Scenarios
                 On the network shown in Figure 4-39, with dynamic BFD for CR-LSP deployed, a
                 BFD session is set up to detect faults on the link of the primary CR-LSP. If a fault
                 occurs on the link of the primary CR-LSP, the BFD module on the ingress
                 immediately reports the fault, and the ingress switches traffic to the backup CR-
                 LSP. After protocol convergence is complete and the primary CR-LSP is torn down,
                 a new BFD session is automatically set up for the backup CR-LSP to detect faults
                 on the link of the backup CR-LSP.




