---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-34
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [4320, 4500]
sha256: c996a86e0c8b8b188d8077bf8c87cb5b279389a6738e07da925a8b25de3ef53c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The configurations of LSRC and LSRD are similar to the configuration of LSRA.
         Step 4 Verify the configuration.
                 Run the display mpls ldp lsp command to check LSP information.
                 # Display LDP LSPs established on LSRA.
                 [LSRA] display mpls ldp lsp
                  LDP LSP Information
                   -------------------------------------------------------------------------------
                   Flag after Out IF: (I) - RLFA Iterated LSP, (I*) - Normal and RLFA Iterated LSP
                   -------------------------------------------------------------------------------
                   DestAddress/Mask In/OutLabel UpstreamPeer NextHop                             OutInterface
                   -------------------------------------------------------------------------------
                   1.1.1.1/32        3/NULL          2.2.2.2        127.0.0.1        LoopBack1
                   2.2.2.2/32        NULL/3          -            192.168.1.2       Vlanif100
                   2.2.2.2/32        1025/3          2.2.2.2       192.168.1.2       Vlanif100
                   4.4.4.4/32        NULL/1025         -            192.168.1.2       Vlanif100
                   4.4.4.4/32        1026/1026         4.4.4.4       192.168.1.2       Vlanif100
                   192.168.1.0/24       3/NULL          2.2.2.2       192.168.1.1       Vlanif100
                  *192.168.1.0/24       Liberal/26                  DS/2.2.2.2
                   192.168.2.0/24       NULL/3          -            192.168.1.2       Vlanif100
                   192.168.2.0/24       1027/3         3.3.3.3        192.168.1.2       Vlanif100
                  --------------------------------------------------------------------------
                   TOTAL: 8 Normal LSP(s) Found.
                   TOTAL: 1 Liberal LSP(s) Found.
                   TOTAL: 0 Frr LSP(s) Found.
                   An asterisk (*) before an LSP means the LSP is not established
                   An asterisk (*) before a Label means the USCB or DSCB is stale


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                               73
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration

                  An asterisk (*) before an UpstreamPeer means the session is stale
                  An asterisk (*) before a DS means the session is stale
                  An asterisk (*) before a NextHop means the LSP is FRR LSP

                 The command output on each node shows that the LDP LSP with LSRB as the
                 transit node is established only for the route 4.4.4.4/32 and that other LDP LSPs
                 not with LSRB as the transit node are established.

                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #
                        vlan batch 100
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
                         ip address 192.168.1.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 1.1.1.1 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 1.1.1.1 0.0.0.0
                          network 192.168.1.0 0.0.0.255
                        #
                        return

                 ●      LSRB
                        #
                        sysname LSRB
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
                           propagate mapping for ip-prefix FilterOnTransit
                        #
                        interface Vlanif100
                         ip address 192.168.1.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 192.168.2.1 255.255.255.0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                74
MPLS Configuration
MPLS Configuration                                                                3 MPLS LDP Configuration

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
                        interface LoopBack1
                         ip address 2.2.2.2 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 192.168.1.0 0.0.0.255
                          network 192.168.2.0 0.0.0.255
                        #
                        ip ip-prefix FilterOnTransit index 10 permit 4.4.4.4 32
                        #
                        return
                 ●      LSRC
                        #
                        sysname LSRC
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls ldp
                         #
                          ipv4-family
                        #
                        interface Vlanif100
                         ip address 192.168.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 192.168.3.1 255.255.255.0
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
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 192.168.2.0 0.0.0.255
                          network 192.168.3.0 0.0.0.255
                        #
                        return
                 ●      LSRD
                        #
                        sysname LSRD


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                            75
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

