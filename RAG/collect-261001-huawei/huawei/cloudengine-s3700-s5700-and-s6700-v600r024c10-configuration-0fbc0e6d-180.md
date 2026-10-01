---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-180
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [26130, 26276]
sha256: 1897adeb73e0c8602f5bba9cf7a29d9f102be492debce6c21058b435cfa306c0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

Procedure
                 ●      Run the display interface tunnel [ interface-number ] command to check the
                        operating status of tunnel interfaces.
                 ●      Run the display mpls te tunnel-interface command to check tunnel
                        interface information on the local node.
                 ●      Run the display mpls lsp command to check CR-LSP information.
                 ●      Run the display mpls te tunnel [ verbose ] command to check MPLS TE
                        tunnel information.
                 ●      Run the display mpls te tunnel { bypass-inuse { inuse | not-exists | exists-
                        not-used } } [ verbose ] command to check information about TE FRR-
                        enabled tunnels.
                 ●      Run the display mpls te tunnel path command to check path information of
                        the primary and bypass tunnels on the local node.

                 ----End

4.26.5 Example for Configuring Manual MPLS TE FRR

Networking Requirements
                 On the network shown in Figure 4-45, the primary CR-LSP is LSR1 -> LSR2 ->
                 LSR3 -> LSR4. FRR is required to protect the link between LSR2 and LSR3.

                 A bypass tunnel is established along the path LSR2 -> LSR5 -> LSR3. LSR2 is the
                 PLR, and LSR3 is an MP.

                 It is required that the explicit path mode be used to establish the primary and
                 bypass MPLS TE tunnels. RSVP-TE is used as a signaling protocol.


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                     436
MPLS Configuration
MPLS Configuration                                                             4 MPLS TE Configuration


                 Figure 4-45 Networking diagram for configuring manual TE FRR




Precautions
                 During the configuration, note the following:

                 ●      LSR IDs must be set before other MPLS commands are run.
                 ●      LSR IDs can only be manually configured, and do not have default values.
                 ●      When specifying an LSR ID, you are advised to use a reachable loopback
                        interface address.
                 ●      To avoid loops in this scenario, ensure that all connected interfaces have STP
                        disabled and are removed from VLAN1. If STP is enabled and VLANIF
                        interfaces of switches are used to construct a Layer 3 ring network, an
                        interface on the network will be blocked. As a result, Layer 3 services on the
                        network cannot run properly.


Configuration Roadmap
                 The configuration roadmap is as follows:

                 1.     Assign an IP address to each interface on each node and configure a loopback
                        interface address as an MPLS LSR ID on each node. Configure IS-IS to ensure
                        that public network routes between nodes are reachable.
                 2.     Configure a dynamic primary MPLS TE tunnel and enable TE FRR on the
                        tunnel interface.
                 3.     Configure a bypass tunnel on the ingress (LSR2) of the protected link, and
                        specify the interface of the protected link in the tunnel interface view.



Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                            437
MPLS Configuration
MPLS Configuration                                                                               4 MPLS TE Configuration


Procedure
         Step 1 Configure interface IP addresses for the nodes.

                 # Configure LSR1.
                 <HUAWEI> system-view
                 [HUAWEI] sysname LSR1
                 [LSR1] vlan batch 100
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] ip address 10.1.1.1 24
                 [LSR1-Vlanif100] quit
                 [LSR1] interface 10ge 1/0/1
                 [LSR1-10GE1/0/1] port link-type trunk
                 [LSR1-10GE1/0/1] port trunk allow-pass vlan 100
                 [LSR1-10GE1/0/1] quit
                 [LSR1] interface loopback 1
                 [LSR1-loopback1] ip address 1.1.1.9 32
                 [LSR1-loopback1] quit

                 The configurations of LSR2, LSR3, LSR4, and LSR5 are similar to the configuration
                 of LSR1. For detailed configurations, see Configuration Scripts.

         Step 2 Configure IS-IS to advertise routes.

                 # Configure LSR1.
                 [LSR1] isis 1
                 [LSR1-isis-1] is-level level-2
                 [LSR1-isis-1] network-entity 00.0005.0000.0000.0001.00
                 [LSR1-isis-1] quit
                 [LSR1] interface vlanif 100
                 [LSR1-Vlanif100] isis enable 1
                 [LSR1-Vlanif100] quit
                 [LSR1] interface loopback 1
                 [LSR1-LoopBack1] isis enable 1
                 [LSR1-LoopBack1] quit

                 The configurations of LSR2, LSR3, LSR4, and LSR5 are similar to the configuration
                 of LSR1. For detailed configurations, see Configuration Scripts.

                 # After the configuration is complete, run the display ip routing-table command
                 on each node to check whether the nodes have learned routes from each other.
                 The following example uses the command output on LSR1.
                 [LSR1] display ip routing-table
                 Proto: Protocol        Pre: Preference
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table : _public_
                        Destinations : 16        Routes : 16

                 Destination/Mask     Proto Pre Cost        Flags NextHop                          Interface

                        1.1.1.9/32 Direct 0 0           D 127.0.0.1                        LoopBack1
                        2.2.2.9/32 ISIS-L2 15 10         D 10.1.1.2                         Vlanif100
                        3.3.3.9/32 ISIS-L2 15 20         D 10.1.1.2                         Vlanif100
                        4.4.4.9/32 ISIS-L2 15 30         D 10.1.1.2                         Vlanif100
                        5.5.5.9/32 ISIS-L2 15 20         D 10.1.1.2                         Vlanif100
                       10.1.1.0/24 Direct 0 0           D 10.1.1.1                         Vlanif100
                       10.1.1.1/32 Direct 0 0           D 127.0.0.1                        Vlanif100
                     10.1.1.255/32 Direct 0 0            D 127.0.0.1                         Vlanif100
                       10.1.2.0/24 ISIS-L2 15 20         D 10.1.1.2                         Vlanif100
                       10.1.3.0/24 ISIS-L2 15 30         D 10.1.1.2                         Vlanif100
                       10.1.4.0/24 ISIS-L2 15 20         D 10.1.1.2                         Vlanif100
                       10.1.5.0/24 ISIS-L2 15 30         D 10.1.1.2                         Vlanif100
                      127.0.0.0/8 Direct 0 0            D 127.0.0.1                        InLoopBack0


Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                          438
MPLS Configuration
MPLS Configuration                                                              4 MPLS TE Configuration

                     127.0.0.1/32 Direct 0 0           D 127.0.0.1           InLoopBack0
                 127.255.255.255/32 Direct 0 0           D 127.0.0.1            InLoopBack0
                 255.255.255.255/32 Direct 0 0           D 127.0.0.1            InLoopBack0

