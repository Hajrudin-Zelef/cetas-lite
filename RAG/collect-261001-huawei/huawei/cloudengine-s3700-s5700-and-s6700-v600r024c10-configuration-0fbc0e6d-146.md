---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-146
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [20902, 21080]
sha256: ceddb9c081379272c4d485ad9319181f2bf29957c294ee4cb3a80bad53aa9094
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 ●      LSR2
                        #
                        sysname LSR2
                        #
                        vlan batch 100 200 300
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls te
                         mpls rsvp-te
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         mpls
                         mpls te
                         mpls te link administrative group 10101
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls rsvp-te
                        #
                        interface Vlanif300
                         ip address 10.1.3.1 255.255.255.0
                         mpls
                         mpls te
                         mpls te link administrative group 10011
                         mpls te bandwidth max-reservable-bandwidth 100000
                         mpls te bandwidth bc0 100000
                         mpls rsvp-te
                        #
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/3
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        mpls lsr-id 2.2.2.2
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          mpls-te enable
                          network 2.2.2.2 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        return




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                      349
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration


                 ●      LSR3
                        #
                        sysname LSR3
                        #
                        vlan batch 200 300
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
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
                        interface 10GE1/0/1
                         port link-type trunk
                         port trunk allow-pass vlan 200
                        #
                        interface 10GE1/0/2
                         port link-type trunk
                         port trunk allow-pass vlan 300
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         mpls te
                         mpls rsvp-te
                        #
                        ospf 1
                         opaque-capability enable
                         area 0.0.0.0
                          mpls-te enable
                          network 3.3.3.3 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        return



4.17 Configuring MPLS TE Explicit Paths

4.17.1 Understanding MPLS TE Explicit Paths
                 An explicit path is a CR-LSP that is established by manually specifying the nodes
                 to traverse or bypass. Explicit paths are classified into the following types:
                 ●      Strict explicit path
                        On a strict explicit path, all the nodes are manually specified and two
                        consecutive hops must be directly connected. A strict explicit path precisely
                        controls the path of an LSP.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                             350
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                        Figure 4-28 Strict explicit path




                        On the network shown in Figure 4-28, an LSP is established over a strict
                        explicit path from the ingress LSR1 to the egress LSR6. In this example, "LSR2
                        Strict" indicates that this LSP must pass through LSR2 and LSR2's previous
                        hop must be the directly connected node LSR1. Similarly, "LSR3 Strict"
                        indicates that this LSP must pass through LSR3 and LSR3's previous hop must
                        be the directly connected node LSR2. This is truth for all LSRs followed by
                        "Strict". In this way, the path that the LSP passes through is precisely
                        controlled.
                 ●      Loose explicit path
                        A loose explicit path can specify the nodes that the path must pass through,
                        but allows for other nodes to exist between the listed nodes.

                        Figure 4-29 Loose explicit path




                        On the network shown in Figure 4-29, an LSP is established over a loose
                        explicit path from the ingress LSR1 to the egress LSR6. "LSR5 Loose" indicates
                        that the LSP must pass through LSR5, but other LSRs may exist between LSR5
                        and LSR1.
                 Strict and loose explicit paths can be used independently or in combination.




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          351
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


4.17.2 Configuring MPLS TE Explicit Paths
Prerequisites
                 Before configuring MPLS TE explicit paths, complete the following task:
                 ●      Configure a dynamic MPLS TE tunnel.

