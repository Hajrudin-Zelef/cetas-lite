---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-33
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [4141, 4319]
sha256: 7f758002db2bb2e7c48f152e07ce4f187126aaad44a4a68fc87d82b7c19a41d8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

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
                        interface Vlanif100
                         ip address 10.1.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface Vlanif200
                         ip address 10.2.1.1 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        70
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration

                         port default vlan 100
                        #
                        interface 10GE1/0/2
                         port link-type access
                         port default vlan 200
                        #
                        interface LoopBack1
                         ip address 2.2.2.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 2.2.2.9 0.0.0.0
                          network 10.1.1.0 0.0.0.3
                          network 10.2.1.0 0.0.0.3
                        #
                        return

                 ●      LSRC
                        #
                        sysname LSRC
                        #
                        vlan batch 100
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
                        interface Vlanif100
                         ip address 10.2.1.2 255.255.255.252
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 3.3.3.9 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.9 0.0.0.0
                          network 10.2.1.0 0.0.0.3
                        #
                        return


3.9.5 Example for Configuring a Policy for Triggering LDP LSP
Establishment (Transit)
Networking Requirements
                 After MPLS LDP is enabled on each interface, LDP LSPs are automatically
                 established, including a large number of unnecessary transit LSPs, which wastes
                 resources. On the network shown in Figure 3-12, configure a policy for triggering
                 transit LSP establishment, allowing LSRB to establish transit LSPs only for the
                 route 4.4.4.4/32. This effectively reduces the number of LSPs to be established and
                 network resource consumption.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        71
MPLS Configuration
MPLS Configuration                                                                  3 MPLS LDP Configuration


                 Figure 3-12 Configuring a policy for triggering transit LSP establishment
                         NOTE

                        In this example, interface1 and interface2 represent VLANIF100 and VLANIF200,
                        respectively.




Precautions
                 During the configuration, note the following:
                 By default, LDP establishes transit LSPs for all routes, without filtering them.

Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign an IP address to each interface and configure OSPF to advertise the
                        route to the network segment to which each interface is connected and the
                        host route to each LSR ID.
                 2.     Configure an IP prefix list to limit the routes for which transit LSPs can be
                        established.
                 3.     Enable MPLS and MPLS LDP globally on each LSR and configure a policy of
                        triggering LSP establishment.
                 4.     Configure LSRB (transit node) to use the IP prefix list to limit the routes for
                        which transit LSPs can be established.
                 5.     Enable MPLS and MPLS LDP on each interface.

Procedure
         Step 1 Assign an IP address to each interface and configure OSPF to advertise the route
                to the network segment to which each interface is connected and the host route
                to each LSR ID.
                 # Assign an IP address to each interface (as shown in Figure 3-12), including the
                 loopback interfaces. Configure OSPF to advertise the route to the network
                 segment to which each interface is connected and the host route to each LSR ID.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                72
MPLS Configuration
MPLS Configuration                                                                                    3 MPLS LDP Configuration


         Step 2 Configure an IP prefix list on the transit node LSRB.
                 # Configure an IP prefix list on LSRB to allow LSRB to establish a transit LSP only
                 for the route 4.4.4.4/32 to LSRD.
                 [LSRB] ip ip-prefix FilterOnTransit permit 4.4.4.4 32

         Step 3 Configure basic MPLS and MPLS LDP functions on each node and interface, and
                configure a policy for triggering LSP establishment.
                 # Configure LSRA.
                 [LSRA] mpls lsr-id 1.1.1.1
                 [LSRA] mpls
                 [LSRA-mpls] quit
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] quit
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] mpls ldp
                 [LSRA-Vlanif100] quit

                 # Configure LSRB.
                 [LSRB] mpls lsr-id 2.2.2.2
                 [LSRB] mpls
                 [LSRB-mpls] quit
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] propagate mapping for ip-prefix FilterOnTransit
                 [LSRB-mpls-ldp] quit
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] mpls
                 [LSRB-Vlanif100] mpls ldp
                 [LSRB-Vlanif100] quit
                 [LSRB] interface vlanif 200
                 [LSRB-Vlanif200] mpls
                 [LSRB-Vlanif200] mpls ldp
                 [LSRB-Vlanif200] quit

