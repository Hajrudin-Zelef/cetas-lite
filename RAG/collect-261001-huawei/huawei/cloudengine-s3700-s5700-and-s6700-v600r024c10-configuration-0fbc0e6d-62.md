---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-62
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [8575, 8769]
sha256: 7685056a111d4945d7bebf6be855408961ead640c8d66caa93bc9ac83cfebe31
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.2.1 255.255.255.0
                         ospf cost 2
                         mpls
                         mpls ldp
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
                         frr
                          loop-free-alternate
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                        #
                        return
                 ●      PE2
                        #
                        sysname PE2
                        #
                        vlan batch 100 200
                        #
                        bfd
                         mpls-passive
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                         lsp-trigger all
                         mpls bfd enable
                         mpls bfd-trigger fec-list l2
                         mpls bfd min-tx-interval 100 min-rx-interval 600 detect-multiplier 4
                        #
                        fec-list l2
                         fec-node 1.1.1.1
                        #
                        mpls ldp
                         #
                         ipv4-family
                          auto-frr lsp-trigger all
                        #
                        interface Vlanif100
                         ip address 10.1.5.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.4.1 255.255.255.0
                         ospf cost 2
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface 10GE1/0/2


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                         144
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         frr
                          loop-free-alternate
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.1.4.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
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
                         lsp-trigger all
                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.5.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
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
                          network 2.2.2.2 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.5.0 0.0.0.255
                        #
                        return
                 ●      P2
                        #
                        sysname P2
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                         lsp-trigger all


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                       145
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                        #
                        mpls ldp
                         #
                         ipv4-family
                        #
                        interface Vlanif100
                         ip address 10.1.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.1.4.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
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
                          network 3.3.3.3 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                          network 10.1.4.0 0.0.0.255
                        #
                        return



3.17 Configuring LDP Session Protection

3.17.1 Understanding LDP Session Protection
                 LDP session protection enables a device to start the extended LDP discovery
                 mechanism to maintain established LDP sessions if the basic LDP discovery
                 mechanism fails. This ensures fast LDP convergence when the basic LDP discovery
                 mechanism recovers.

