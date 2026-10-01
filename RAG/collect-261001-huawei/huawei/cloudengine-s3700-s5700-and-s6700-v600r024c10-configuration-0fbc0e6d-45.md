---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-45
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "cost"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [5997, 6155]
sha256: 7f3347a1bc9b7749dffb3fbad5989d5ce482fe4b9379115297b278eeeeedfcf0
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.13.3 Example for Configuring LDP Extension for Inter-Area
LSPs
Networking Requirements
                 The network shown in Figure 3-19 has two IGP areas: Area 10 and Area 20. Inter-
                 area LSPs need to be established from LSRA to LSRB and from LSRA to LSRC. LDP
                 extension for inter-area LSPs needs to be configured on LSRA so that LSRA can
                 search for routes based on the longest match rule to establish LSPs.

                 Figure 3-19 Configuring LDP extension for inter-area LSPs
                         NOTE

                        In this example, interface1, interface2, and interface3 represent VLANIF100, VLANIF200,
                        and VLANIF300, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        101
MPLS Configuration
MPLS Configuration                                                          3 MPLS LDP Configuration




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign IP addresses to interfaces on each node and configure the loopback
                        addresses to be used as LSR IDs.
                 2.     Configure basic IS-IS functions.
                 3.     Configure a policy for summarizing routes.
                 4.     Enable MPLS and MPLS LDP globally and on the interfaces of each node.
                 5.     Configure LDP extension for inter-area LSPs.

Procedure
         Step 1 Assign IP addresses to interfaces on each node and configure the loopback
                addresses to be used as LSR IDs.
                 Assign an IP address and a mask to each interface (including loopback interfaces)
                 according to Figure 3-19.
         Step 2 Configure basic IS-IS functions.
                 # Configure LSRA.
                 <LSRA> system-view
                 [LSRA] isis 1
                 [LSRA-isis-1] is-level level-2
                 [LSRA-isis-1] network-entity 20.0010.0100.0001.00
                 [LSRA-isis-1] quit
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] isis enable 1
                 [LSRA-Vlanif100] quit
                 [*LSRA] interface loopback 0
                 [LSRA-LoopBack0] isis enable 1
                 [LSRA-LoopBack0] quit


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           102
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


                 # Configure LSRD.
                 <LSRD> system-view
                 [LSRD] isis 1
                 [LSRD-isis-1] network-entity 10.0010.0200.0001.00
                 [LSRD-isis-1] quit
                 [LSRD] interface vlanif 100
                 [LSRD-Vlanif100] isis enable 1
                 [LSRD-Vlanif100] isis circuit-level level-2
                 [LSRD-Vlanif100] quit
                 [LSRD] interface vlanif 200
                 [LSRD-Vlanif200] isis enable 1
                 [LSRD-Vlanif200] isis circuit-level level-1
                 [LSRD-Vlanif200] quit
                 [LSRD] interface vlanif 300
                 [LSRD-Vlanif300] isis enable 1
                 [LSRD-Vlanif300] isis circuit-level level-1
                 [LSRD-Vlanif300] quit
                 [LSRD] interface loopback 0
                 [LSRD-LoopBack0] isis enable 1
                 [LSRD-LoopBack0] quit

                 # Configure LSRB.
                 <LSRB> system-view
                 [LSRB] isis 1
                 [LSRB-isis-1] is-level level-1
                 [LSRB-isis-1] network-entity 10.0010.0300.0001.00
                 [LSRB-isis-1] quit
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] isis enable 1
                 [LSRB-Vlanif100] quit
                 [LSRB] interface loopback 0
                 [LSRB-LoopBack0] isis enable 1
                 [LSRB-LoopBack0] quit

                 # Configure LSRC.
                 <LSRC> system-view
                 [LSRC] isis 1
                 [LSRC-isis-1] is-level level-1
                 [LSRC-isis-1] network-entity 10.0010.0300.0002.00
                 [LSRC-isis-1] quit
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] isis enable 1
                 [LSRC-Vlanif100] quit
                 [LSRC] interface loopback 0
                 [LSRC-LoopBack0] isis enable 1
                 [LSRC-LoopBack0] quit

                 # Run the display ip routing-table command on LSRA to check route
                 information.
                 [LSRA] display ip routing-table
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                  ------------------------------------------------------------------------------
                  Routing Table: Public
                         Destinations : 9        Routes : 9

                  Destination/Mask    Proto    Pre Cost       Flags NextHop        Interface

                        10.10.1.1/32 Direct 0 0           D 127.0.0.1      LoopBack0
                        10.10.2.2/32 ISIS-L1 15 10         D 10.1.1.2       Vlanif100
                        10.10.3.1/32 ISIS-L1 15 20         D 10.1.1.2       Vlanif100
                        10.10.3.2/32 ISIS-L1 15 20         D 10.1.1.2       Vlanif100
                         10.1.1.0/24 Direct 0 0           D 10.1.1.1      Vlanif100
                         10.1.1.1/32 Direct 0 0           D 127.0.0.1     Vlanif100
                         10.1.1.2/32 Direct 0 0           D 10.1.1.2      Vlanif100
                        127.0.0.0/8 Direct 0 0            D 127.0.0.1     InLoopBack0
                        127.0.0.1/32 Direct 0 0           D 127.0.0.1      InLoopBack0


Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                        103
MPLS Configuration
MPLS Configuration                                                                             3 MPLS LDP Configuration


         Step 3 Configure a policy for summarizing routes.
                 # On LSRD, run the summary command to summarize host routes to LSRB and
                 LSRC.
                 [LSRD] isis 1
                 [LSRD-isis-1] summary 10.10.3.0 255.255.255.0 avoid-feedback
                 [LSRD-isis-1] quit

                 # Run the display ip routing-table command on LSRA to check route
                 information.
                 Route Flags: R - relay, D - download to fib, T - to vpn-instance, B - black hole route
                 ------------------------------------------------------------------------------
                 Routing Table: Public
                        Destinations : 8        Routes : 8

                 Destination/Mask     Proto   Pre Cost       Flags NextHop        Interface

                     10.10.1.1/32 Direct 0 0             D 127.0.0.1       LoopBack0
                     10.10.2.2/32 ISIS-L1 15 10           D 10.1.1.2        Vlanif100
                     10.10.3.0/24 ISIS-L1 15 20           D 10.1.1.2         Vlanif100
                      10.1.1.0/24 Direct 0 0             D 10.1.1.1       Vlanif100
                      10.1.1.1/32 Direct 0 0             D 127.0.0.1      Vlanif100
                      10.1.1.2/32 Direct 0 0             D 10.1.1.2       Vlanif100
                     127.0.0.0/8 Direct 0 0              D 127.0.0.1      InLoopBack0
                     127.0.0.1/32 Direct 0 0             D 127.0.0.1       InLoopBack0

