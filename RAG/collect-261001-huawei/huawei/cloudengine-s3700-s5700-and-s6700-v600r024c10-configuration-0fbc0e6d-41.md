---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-41
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "memory", "voice"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [5326, 5523]
sha256: d44e5907dd254232d420cfdf3338390668e4e3d2a1d453afc64e7b038cbe5b5c
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

                         area 0.0.0.0
                          network 2.2.2.2 0.0.0.0
                          network 10.1.1.0 0.0.0.255
                          network 10.1.2.0 0.0.0.255
                          network 10.1.3.0 0.0.0.255
                        #
                        return
                 ●      LSRC
                        #
                        sysname LSRC
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 3.3.3.3
                        #
                        mpls
                        #
                        mpls ldp
                        #
                        interface Vlanif100
                         ip address 10.1.2.2 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 3.3.3.3 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 3.3.3.3 0.0.0.0
                          network 10.1.2.0 0.0.0.255
                        #
                        return
                 ●      LSRD
                        #
                        sysname LSRD
                        #
                        vlan batch 100
                        #
                        mpls lsr-id 4.4.4.4
                        #
                        mpls
                        #
                        mpls ldp
                         ipv4-family
                          inbound peer 2.2.2.2 fec ip-prefix prefix1
                        #
                        interface Vlanif100
                         ip address 10.1.3.1 255.255.255.0
                         mpls
                         mpls ldp
                        #
                        interface 10GE1/0/1
                         port link-type access
                         port default vlan 100
                        #
                        interface LoopBack1
                         ip address 4.4.4.4 255.255.255.255
                        #
                        ospf 1
                         area 0.0.0.0
                          network 4.4.4.4 0.0.0.0
                          network 10.1.3.0 0.0.0.255
                        #


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                        90
MPLS Configuration
MPLS Configuration                                                                   3 MPLS LDP Configuration

                        ip ip-prefix prefix1 index 10 permit 3.3.3.3 32
                        #
                        return


3.10.8 Example for Configuring an Outbound LDP Policy

Networking Requirements
                 An IP metro or bearer network uses L2VPN or L3VPN to transmit high speed
                 Internet (HSI) or voice over IP (VoIP) services over end-to-end public network LDP
                 LSPs. Generally, user-side DSLAMs have low performance, and are easily
                 overloaded if a large number of LDP LSPs are established. To prevent this issue,
                 configure an outbound LDP policy to minimize LDP LSPs to be established, reduce
                 DSLAM memory consumption, and relieve the burden of the DSLAMs.


                 Figure 3-16 Configuring an outbound LDP policy
                         NOTE

                        Interfaces 1 and 2 in this example represent VLANIF100 and VLANIF200, respectively.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                   91
MPLS Configuration
MPLS Configuration                                                           3 MPLS LDP Configuration


Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign an IP address to each interface, including the loopback interface on
                        each node.
                 2.     Configure OSPF to advertise the route to the network segment of each
                        interface and to advertise the host route to each LSR ID.
                 3.     Enable MPLS and MPLS LDP globally on each node.
                 4.     Configure an outbound LDP policy on LSRA to send the DSLAM the Label
                        Mapping messages destined for LSRC only. This allows the DSLAM to establish
                        an LSP to LSRC only, reducing memory usage.
                 5.     Configure MPLS and MPLS LDP on each interface.

Procedure
         Step 1 Assign an IP address to each interface and configure an IGP.
                 Assign an IP address and mask to each interface (as shown in Figure 3-16),
                 including the loopback interfaces. Configure OSPF to advertise the route to the
                 network segment to which each interface is connected and the host route to each
                 LSR ID.
         Step 2 Enable MPLS and MPLS LDP globally on each node.
                 # Configure LSRA.
                 <LSRA> system-view
                 [LSRA] mpls lsr-id 3.3.3.9
                 [LSRA] mpls
                 [LSRA-mpls] quit
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] quit

                 # Configure LSRB.
                 <LSRB> system-view
                 [LSRB] mpls lsr-id 2.2.2.9
                 [LSRB] mpls
                 [LSRB-mpls] quit
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] quit

                 # Configure LSRC.
                 <LSRC> system-view
                 [LSRC] mpls lsr-id 1.1.1.9
                 [LSRC] mpls
                 [LSRC-mpls] quit
                 [LSRC] mpls ldp
                 [LSRC-mpls-ldp] quit

                 # Configure the DSLAM.
                 <DSLAM> system-view
                 [DSLAM] mpls lsr-id 4.4.4.9
                 [DSLAM] mpls
                 [DSLAM-mpls] quit
                 [DSLAM] mpls ldp
                 [DSLAM-mpls-ldp] quit

         Step 3 Configure an outbound LDP policy.

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                            92
MPLS Configuration
MPLS Configuration                                                                                     3 MPLS LDP Configuration


                 # Configure an IP prefix list on LSRA to permit only the routes to LSRC.
                 [LSRA] ip ip-prefix prefix1 permit 1.1.1.9 32

                 # Configure an outbound policy on LSRA to send the DSLAM the Label Mapping
                 messages destined for LSRC only.
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] ipv4-family
                 [LSRA-mpls-ldp-ipv4] outbound peer 4.4.4.9 fec ip-prefix prefix1
                 [LSRA-mpls-ldp-ipv4] quit
                 [LSRA-mpls-ldp] quit

         Step 4 Enable MPLS and MPLS LDP on each interface.
                 # Configure LSRA.
                 <LSRA> system-view
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] mpls ldp
                 [LSRA-Vlanif100] quit
                 [LSRA] interface vlanif 200
                 [LSRA-Vlanif200] mpls
                 [LSRA-Vlanif200] mpls ldp
                 [LSRA-Vlanif200] quit

                 # Configure LSRB.
                 <LSRB> system-view
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] mpls
                 [LSRB-Vlanif100] mpls ldp
                 [LSRB-Vlanif100] quit
                 [LSRB] interface vlanif 200
                 [LSRB-Vlanif200] mpls
                 [LSRB-Vlanif200] mpls ldp
                 [LSRB-Vlanif200] quit

