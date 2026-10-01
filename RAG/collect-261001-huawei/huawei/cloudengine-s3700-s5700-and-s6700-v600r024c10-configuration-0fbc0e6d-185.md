---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-185
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [26896, 27071]
sha256: b56dac92f99c0c3f12c6cea77ca55b7ac6c76a50c110d17bf8f6b2249d2eeb32
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
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0004.00
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                      447
MPLS Configuration
MPLS Configuration                                                            4 MPLS TE Configuration

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
                         traffic-eng level-2
                         network-entity 00.0005.0000.0000.0005.00
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




4.27 Configuring MPLS TE Auto FRR



Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      448
MPLS Configuration
MPLS Configuration                                                           4 MPLS TE Configuration


4.27.1 Understanding MPLS TE Auto FRR
                 MPLS TE FRR protects links and nodes on MPLS TE tunnels. If a link or node fails,
                 TE FRR rapidly switches traffic to a backup path, minimizing traffic loss.
                 Auto FRR automatically establishes an eligible bypass tunnel, reducing the
                 configuration workload.

Context
                 Generally, a link or node failure on an MPLS TE tunnel triggers a tunnel switchover
                 from the primary path to the backup path. The switching process involves IGP
                 route convergence on the backup path, CSPF path recalculation, and CR-LSP re-
                 establishment. The switching process is slow, and traffic loss may occur.
                 TE FRR can solve this problem. TE FRR establishes a backup path that bypasses the
                 faulty link or node in advance. When a link or node on an MPLS TE tunnel fails,
                 traffic can be quickly switched to the backup path, preventing traffic loss. When
                 traffic continues to be transmitted over the backup path, the ingress continues to
                 reestablish the primary path.

Benefits
                 TE FRR provides carrier-class local protection capabilities for MPLS TE, improving
                 the reliability of an entire network.

Related Concepts
                 On the network shown in Figure 4-46, TE FRR establishes a bypass CR-LSP for a
                 primary tunnel on each possible faulty link or node. One bypass CR-LSP can
                 protect multiple primary tunnels. This protection mode is also called facility
                 backup.

                 Figure 4-46 Networking diagram of TE FRR local protection




                 For details about TE FRR concepts involved in Figure 4-46, see Table 4-25.




Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                           449
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Table 4-25 Concepts in TE FRR
                  Concept          Description

                  Primary CR-      Primary CR-LSP, which is the protected CR-LSP.
                  LSP

                  Bypass CR-       CR-LSP protecting the primary CR-LSP. A bypass CR-LSP and its
                  LSP              primary CR-LSP belong to different tunnels.
                                   A bypass CR-LSP is usually in idle state and does not forward
                                   service traffics. If a bypass CR-LSP needs to be used to
                                   independently forward service data while protecting the primary
                                   CR-LSP, sufficient bandwidth must be allocated to the bypass
                                   CR-LSP.

                  PLR              Point of local repair. It is the ingress of a bypass CR-LSP. It must
                                   reside on a primary CR-LSP, and can be the ingress or transit
                                   node of a primary CR-LSP, but cannot be the egress of a primary
                                   CR-LSP.

                  MP               Merge point. The egress of a bypass CR-LSP must be on the path
                                   of the primary CR-LSP and cannot be the ingress of the primary
                                   CR-LSP.




                 Table 4-26 describes TE FRR classification.

                 Table 4-26 TE FRR classification
                  Classified      Type                 Description
                  By

