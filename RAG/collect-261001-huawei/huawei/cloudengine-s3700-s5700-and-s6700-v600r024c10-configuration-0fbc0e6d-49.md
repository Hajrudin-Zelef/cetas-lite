---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-49
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [6657, 6818]
sha256: 8abdaa7100525ebf5022802c4f57f05b80517871ab1e56d3fce8b6f76ca48549
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.14.4 Example for Configuring LDP Auto FRR
Networking Requirements
                 Network services, such as VoIP, online gaming, and online video services, have
                 higher requirements on real-time performance. Many services are based on VPNs,
                 and VPN services are usually transmitted over LDP tunnels. Data loss caused by
                 link faults severely affects these services.
                 To address such issues, LDP FRR can be manually configured. If a fault occurs,
                 public network service traffic can be forwarded along the backup LSP before
                 routes re-converge and a new primary LSP is generated. Traffic loss during fault
                 detection and traffic switchover lasts less than 50 ms. However, after route re-
                 convergence is complete, the time for a VPN service to switch to the new primary
                 LSP depends on the VPN implementation. In order to keep the VPN service
                 interruption time within 50 ms, the speed of switching the VPN service to the new
                 primary LSP needs to be improved. Configure LDP Auto FRR to address this need.
                 On the network shown in Figure 3-23, primary and backup LSPs are established
                 from LSRA to LSRC. The LSP over the path LSRA -> LSRC is the primary one, and
                 the LSP over the path LSRA -> LSRB -> LSRC is the backup one. To allow traffic to
                 rapidly switch to the backup LSP if the primary LSP fails, configure LDP Auto FRR
                 on LSRA to enable LSRA to automatically establish a backup LSP. Traffic can then
                 be rapidly switched to the backup LSP if a fault occurs in the primary LSP,
                 minimizing traffic loss.

                 Figure 3-23 Configuring LDP Auto FRR
                         NOTE

                        In this example, interface1, interface2, and interface3 represent VLANIF100, VLANIF200,
                        and VLANIF300, respectively.




Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                                        113
MPLS Configuration
MPLS Configuration                                                            3 MPLS LDP Configuration




Configuration Roadmap
                 The configuration roadmap is as follows:
                 1.     Assign IP addresses to interfaces on each node and configure the loopback
                        addresses to be used as LSR IDs.
                 2.     Configure IS-IS to advertise the route to each network segment of each
                        interface and to advertise the host route to each LSR ID.
                 3.     Enable MPLS and MPLS LDP globally and on the interfaces of each node.
                 4.     Enable IS-IS Auto FRR on the ingress LSR to protect traffic.
                 5.     Configure a policy for triggering LDP LSP establishment based on all routes.
                 6.     Configure a policy for triggering backup LDP LSP establishment on LSRA.

Procedure
         Step 1 Assign an IP address to each interface.
                 Assign an IP address and a mask to each interface (including loopback interfaces)
                 according to Figure 3-23.
         Step 2 Configure IS-IS to advertise the route to each network segment of each interface
                and to advertise the host route to each LSR ID.
                 # Configure LSRA.
                 <LSRA> system-view
                 [LSRA] isis 1
                 [LSRA-isis-1] network-entity 10.0000.0000.0001.00
                 [LSRA-isis-1] quit
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] isis enable 1
                 [LSRA-Vlanif100] quit
                 [LSRA] interface vlanif 200
                 [LSRA-Vlanif200] isis enable 1
                 [LSRA-Vlanif200] quit
                 [LSRA] interface loopBack 0
                 [LSRA-LoopBack0] isis enable 1
                 [LSRA-LoopBack0] quit


Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                           114
MPLS Configuration
MPLS Configuration                                                          3 MPLS LDP Configuration


                 # Configure LSRB.
                 <LSRB> system-view
                 [LSRB] isis 1
                 [LSRB-isis-1] network-entity 10.0000.0000.0002.00
                 [LSRB-isis-1] quit
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] isis enable 1
                 [LSRB-Vlanif100] quit
                 [LSRB] interface vlanif 200
                 [LSRB-Vlanif200] isis enable 1
                 [LSRB-Vlanif200] quit
                 [LSRB] interface loopBack 0
                 [LSRB-LoopBack0] isis enable 1
                 [LSRB-LoopBack0] quit

                 # Configure LSRC.
                 <LSRC> system-view
                 [LSRC] isis 1
                 [LSRC-isis-1] network-entity 10.0000.0000.0003.00
                 [LSRC-isis-1] quit
                 [LSRC] interface vlanif 100
                 [LSRC-Vlanif100] isis enable 1
                 [LSRC-Vlanif100] quit
                 [LSRC] interface vlanif 200
                 [LSRC-Vlanif200] isis enable 1
                 [LSRC-Vlanif200] quit
                 [LSRC] interface vlanif 300
                 [LSRC-Vlanif300] isis enable 1
                 [LSRC-Vlanif300] quit
                 [LSRC] interface loopBack 0
                 [LSRC-LoopBack0] isis enable 1
                 [LSRC-LoopBack0] quit

                 # Configure LSRD.
                 <LSRD> system-view
                 [LSRD] isis 1
                 [LSRD-isis-1] network-entity 10.0000.0000.0004.00
                 [LSRD-isis-1] quit
                 [LSRD] interface vlanif 100
                 [LSRD-Vlanif100] isis enable 1
                 [LSRD-Vlanif100] quit
                 [LSRD] interface loopBack 0
                 [LSRD-LoopBack0] isis enable 1
                 [LSRD-LoopBack0] quit

         Step 3 Configure MPLS and MPLS LDP globally and on interfaces on each node so that
                the network can forward MPLS traffic. Then, check information about established
                LSPs.
                 # Configure LSRA.
                 [LSRA] mpls lsr-id 1.1.1.9
                 [LSRA] mpls
                 [LSRA-mpls] quit
                 [LSRA] mpls ldp
                 [LSRA-mpls-ldp] quit
                 [LSRA] interface vlanif 100
                 [LSRA-Vlanif100] mpls
                 [LSRA-Vlanif100] mpls ldp
                 [LSRA-Vlanif100] quit
                 [LSRA] interface vlanif 200
                 [LSRA-Vlanif200] mpls
                 [LSRA-Vlanif200] mpls ldp
                 [LSRA-Vlanif200] quit

                 # Configure LSRB.

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                       115
MPLS Configuration
MPLS Configuration                                                                            3 MPLS LDP Configuration

                 [LSRB] mpls lsr-id 2.2.2.9
                 [LSRB] mpls
                 [LSRB-mpls] quit
                 [LSRB] mpls ldp
                 [LSRB-mpls-ldp] quit
                 [LSRB] interface vlanif 100
                 [LSRB-Vlanif100] mpls
                 [LSRB-Vlanif100] mpls ldp
                 [LSRB-Vlanif100] quit
                 [*LSRB] interface vlanif 200
                 [LSRB-Vlanif200] mpls
                 [LSRB-Vlanif200] mpls ldp
                 [LSRB-Vlanif200] quit

