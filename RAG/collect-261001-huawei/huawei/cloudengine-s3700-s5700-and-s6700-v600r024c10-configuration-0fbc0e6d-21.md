---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-21
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [2274, 2469]
sha256: 566bb266c965be7c9203e6e4d8e7bd79f33b10c479edeff8bdf8e255ec64cef5
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                 The next hop and outbound interface of the static LSP 192.168.3.9/32 from LSRA
                 to LSRC are determined by the routing table. In this example, the next-hop IP
                 address is 10.1.1.2/24.
         Step 3 Enable MPLS globally on each node.
                 # Configure LSRA.
                 [LSRA] mpls lsr-id 192.168.1.9
                 [LSRA] mpls
                 [LSRA-mpls] quit

                 # Configure LSRB.
                 [LSRB] mpls lsr-id 192.168.2.9
                 [LSRB] mpls
                 [LSRB-mpls] quit

                 # Configure LSRC.
                 [LSRC] mpls lsr-id 192.168.3.9
                 [LSRC] mpls
                 [LSRC-mpls] quit

         Step 4 Enable MPLS on each interface.
                 # Configure LSRA.
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] quit

                 # Configure LSRB.
                 [LSRB] interface vlanif 100
                 [LSRB-F] mpls
                 [LSRB-Vlanif100] quit
                 [LSRB] interface vlanif 200
                 [LSRB-Vlanif200] mpls
                 [LSRB-Vlanif200] quit


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          39
MPLS Configuration
MPLS Configuration                                                                        3 MPLS LDP Configuration


                 # Configure LSRC.
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] mpls
                 [LSRC-Vlanif100] quit

         Step 5 Create a static LSP from LSRA to LSRC.
                 # Configure LSRA as the ingress.
                 [LSRA] static-lsp ingress AtoC destination 192.168.3.9 32 nexthop 10.1.1.2 out-label 20

                 # Configure LSRB as a transit node.
                 [LSRB] static-lsp transit AtoC in-label 20 outgoing-interface Vlanif200 nexthop 10.2.1.2 out-label 40

                 # Configure LSRC as the egress.
                 [LSRC] static-lsp egress AtoC incoming-interface Vlanif100 in-label 40

         Step 6 Verify the configuration.
                 After completing the configuration, run the display mpls static-lsp or display
                 mpls static-lsp verbose command on each node to check the status of the static
                 LSP. The following example uses the command output on LSRA.
                 [LSRA] display mpls static-lsp
                 TOTAL         :1       STATIC LSP(S)
                 UP           :1       STATIC LSP(S)
                 DOWN             :0      STATIC LSP(S)
                 Name                 FEC           I/O Label   I/O If                 Status
                 AtoC               192.168.3.9/32    NULL/20      -/Vlanif100              Up
                 [LSRA] display mpls static-lsp verbose
                  No           :1
                  LSP-Name           : AtoC
                  LSR-Type        : Ingress
                  FEC          : 192.168.3.9/32
                  In-Label       : NULL
                  Out-Label        : 20
                  In-Interface : -
                  Out-Interface : Vlanif100
                  NextHop           : 10.1.1.2
                  Static-Lsp Type : Normal
                  Lsp Status      : Up

                 ----End

Configuration Scripts
                 ●      LSRA
                        #
                        sysname LSRA
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 192.168.1.9
                        #
                        mpls
                        #
                        interface Vlanif100
                         ip address 10.1.1.1 255.255.255.0
                         mpls
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                              40
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration

                        interface LoopBack1
                         ip address 192.168.1.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 192.168.1.9 0.0.0.0
                        #
                         static-lsp ingress AtoC destination 192.168.3.9 32 nexthop 10.1.1.2 out-label 20
                        #
                        return
                 ●      LSRB
                        #
                        sysname LSRB
                        #
                        vlan batch 100 200
                        #
                        mpls lsr-id 192.168.2.9
                        #
                        mpls
                        #
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.0
                         mpls
                        #
                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.0
                         mpls
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
                         ip address 192.168.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.2.1.0 0.0.0.255
                          network 192.168.2.9 0.0.0.0
                        #
                         static-lsp transit AtoC in-label 20 outgoing-interface Vlanif200 nexthop 10.2.1.2 out-label 40
                        #
                        return
                 ●      LSRC
                        #
                        sysname LSRC
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 192.168.3.9
                        #
                        mpls
                        #
                        interface Vlanif100
                         ip address 10.2.1.2 255.255.255.0
                         mpls
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                               41
MPLS Configuration
MPLS Configuration                                                                         3 MPLS LDP Configuration

                         ip address 192.168.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 10.2.1.0 0.0.0.255
                          network 192.168.3.9 0.0.0.0
                        #
                         static-lsp egress AtoC incoming-interface Vlanif100 in-label 40
                        #
                        return



3.6 Configuring LDP Sessions

