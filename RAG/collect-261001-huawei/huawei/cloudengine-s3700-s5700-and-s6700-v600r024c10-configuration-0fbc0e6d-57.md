---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-57
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [7852, 8051]
sha256: 7ca053ec6f2ade60249226013bcb19594dd22193fcbc2e578e70b668ac67b7a4
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Configuration Scripts
                 ●      PE1
                        #
                        sysname PE1
                        #
                        vlan batch 100 200
                        #
                         bfd
                        #
                         mpls lsr-id 1.1.1.1
                        #
                         mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                        #
                        bfd 1to4 bind ldp-lsp peer-ip 4.4.4.4 nexthop 10.1.1.2 interface Vlanif100
                         discriminator local 1
                         discriminator remote 2
                         process-pst
                         #
                        return
                 ●      PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        bfd
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        132
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                         mpls lsr-id 4.4.4.4
                        #
                         mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.5.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.4.1 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.1.5.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                          network 4.4.4.4 0.0.0.0
                        #
                        bfd 4to1 bind peer-ip 1.1.1.1
                         discriminator local 2
                         discriminator remote 1
                         #
                        return
                 ●      P1
                        #
                         sysname P1
                        #
                         vlan batch 100 200
                        #
                         mpls lsr-id 2.2.2.2
                        #
                         mpls
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface vlanif200
                         ip address 10.1.5.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       133
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                        return

                 ●      P2
                        #
                         sysname P2
                        #
                         vlan batch 100 200
                        #
                        interface Vlanif100
                         ip address 10.1.2.2 255.255.255.0
                        #
                        interface Vlanif200
                         ip address 10.1.4.2 255.255.255.0
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.1.4.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                        return



3.16 Configuring Dynamic BFD for LDP LSP

3.16.1 Understanding BFD for LDP
                 BFD can rapidly detect faults on LDP LSPs and trigger a primary/backup LSP
                 switchover, improving network reliability.

Context
                 If a node or link along an LDP LSP fails, traffic is switched to a backup LSP. The
                 switchover process involves fault detection and traffic switchover. If either fault
                 detection or traffic switchover is slow, traffic will be lost for a long time in the
                 process. LDP FRR can be used to speed up traffic switchover, but it is ineffective for
                 fault detection. As such, LDP FRR is not enough to resolve the preceding issue.
                 Take the network shown in Figure 3-26 as an example. Each LSR periodically
                 sends Hello messages to notify the neighboring LSRs (peers) of its presence and
                 establish a Hello adjacency with each of the peers. An LSR creates a Hello hold
                 timer for each peer to maintain the Hello adjacency, and updates the Hello hold
                 timer each time the LSR receives a Hello message from the peer. If the Hello hold

