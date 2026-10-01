---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-52
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [7088, 7292]
sha256: 04dcaa43c51c3286b0254ce29f5bd665091e257e7cef04816bd068927aef87cf
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         isis enable 1
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
                        #
                        interface LoopBack0
                         ip address 1.1.1.9 255.255.255.255
                         isis enable 1
                        #
                        return

                 ●      LSRB
                        #
                        sysname LSRB
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 2.2.2.9
                        #
                        mpls
                         lsp-trigger all
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        isis 1
                         network-entity 10.0000.0000.0002.00
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.3.1 255.255.255.0
                         isis enable 1
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
                        #
                        interface LoopBack0
                         ip address 2.2.2.9 255.255.255.255
                         isis enable 1
                        #
                        return

                 ●      LSRC


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       120
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration

                        #
                        sysname LSRC
                        #
                        vlan batch 100 200 300
                        #
                        mpls lsr-id 3.3.3.9
                        #
                        mpls
                         lsp-trigger all
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        isis 1
                         network-entity 10.0000.0000.0003.00
                        #
                        interface Vlanif100
                         ip address 10.1.4.1 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.2.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface Vlanif300
                         ip address 10.1.3.2 255.255.255.0
                         isis enable 1
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
                        #
                        interface 10GE1/0/3
                         port link-type access
                         port default vlan 300
                        #
                        interface LoopBack0
                         ip address 3.3.3.9 255.255.255.255
                         isis enable 1
                        #
                        return
                 ●      LSRD
                        #
                         sysname LSRD
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 4.4.4.9
                        #
                        mpls
                         lsp-trigger all
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        isis 1
                         network-entity 10.0000.0000.0004.00


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                       121
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        #
                        interface Vlanif100
                         ip address 10.1.4.2 255.255.255.0
                         isis enable 1
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack0
                         ip address 4.4.4.9 255.255.255.255
                         isis enable 1
                        #
                        return



3.15 Configuring Static BFD for LDP LSP

3.15.1 Understanding BFD for LDP
                 BFD can rapidly detect faults on LDP LSPs and trigger a primary/backup LSP
                 switchover, improving network reliability.

Context
                 If a node or link along an LDP LSP fails, traffic is switched to a backup LSP. The
                 switchover process involves fault detection and traffic switchover. If either fault
                 detection or traffic switchover is slow, traffic will be lost for a long time in the
                 process. LDP FRR can be used to speed up traffic switchover, but it is ineffective for
                 fault detection. As such, LDP FRR is not enough to resolve the preceding issue.

                 Take the network shown in Figure 3-24 as an example. Each LSR periodically
                 sends Hello messages to notify the neighboring LSRs (peers) of its presence and
                 establish a Hello adjacency with each of the peers. An LSR creates a Hello hold
                 timer for each peer to maintain the Hello adjacency, and updates the Hello hold
                 timer each time the LSR receives a Hello message from the peer. If the Hello hold
                 timer expires before a new Hello message is received, the LSR considers that the
                 Hello adjacency is interrupted. The Hello mechanism has a drawback — it cannot
                 rapidly detect link faults, especially when a Layer 2 device is deployed between an
                 LSR and its peer.

                 Figure 3-24 Primary and backup LDP LSPs




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        122
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration


