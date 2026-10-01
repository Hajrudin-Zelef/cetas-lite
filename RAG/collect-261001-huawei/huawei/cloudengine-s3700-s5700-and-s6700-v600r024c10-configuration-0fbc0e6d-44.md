---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-0fbc0e6d-44
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d.md
source_anchor: ""
source_lines: [5836, 5996]
sha256: fa102c933d9c7aaafee2826d912bf6057bea4e9a574270f99ce5c09b1f3bad2b
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--0fbc0e6d

3.12 Disabling a Device from Forwarding Unknown
TLVs
Context
                 After a device is configured not to forward unknown TLVs, the device does not
                 transmit unknown TLVs to neighbors.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Disable the device from forwarding unknown TLVs.
                 propagate mapping unknown-tlv disable

                 By default, the device is enabled to forward unknown TLVs.
                 If an upstream device cannot process unknown TLVs, network problems may
                 occur. In this case, you can run this command to disable the local device from
                 forwarding unknown TLVs.

                 ----End


3.13 Configuring LDP Extension for Inter-Area LSPs

3.13.1 Understanding LDP Extension for Inter-Area LSPs
Fundamentals
                 On a large-scale network, multiple IGP areas need to be configured for flexible
                 deployment and fast convergence. To prevent excessive resource consumption
                 caused by a large number of routes, an area border router (ABR) needs to
                 summarize the routes in an area and advertise the summary routes to neighboring
                 IGP areas. LDP extension for inter-area LSPs enables LDP to search for routes
                 based on the longest match rule so that LDP can establish LDP LSPs across
                 multiple IGP areas based on summary routes.




Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                          98
MPLS Configuration
MPLS Configuration                                                         3 MPLS LDP Configuration


                 Figure 3-17 LDP extension for inter-area LSPs




                 The network shown in Figure 3-17 has two IGP areas: Area 10 and Area 20.
                 In the routing table of LSR-D on the edge device of Area 10, there are two host
                 routes to LSR-B and LSR-C. To minimize route resource consumption, you can use
                 IS-IS on LSR-D to summarize the two routes into one. The resulting summary
                 route192.168.3.0/24 can then be sent to Area 20. Consequently, the routing table
                 of LSR-A contains the summary route, but does not contain host routes with 32-
                 bit addresses. By default, when establishing an LSP, LDP searches the routing table
                 for the route that exactly matches the FEC carried in a received Label Mapping
                 message. Table 3-3 describes the routing entry information of LSR-A and the
                 routing information carried in the FEC on the network shown in Figure 3-17.

                 Table 3-3 Routing entry information of LSR-A and routing information carried in
                 the FEC
                  Routing Entry                Routing Information Carried in the FEC
                  Information of LSR-A

                  192.168.3.0/24               192.168.3.1/32

                                               192.168.3.2/32



                 For summary routes, LDP can establish only liberal LSPs (refer to those that have
                 been assigned labels but fail to be established), but cannot establish LDP LSPs
                 across IGP areas. As a result, LDP cannot provide backbone network tunnels for
                 VPN services.
                 To prevent such an issue on the network shown in Figure 3-17, enable LDP to
                 search for routes based on the longest match rule during LSP establishment. The

Issue 01 (2025-03-03)        Copyright © Huawei Technologies Co., Ltd.                            99
MPLS Configuration
MPLS Configuration                                                         3 MPLS LDP Configuration


                 summary route 192.168.3.0/24 already exists in the routing table of LSR-A. When
                 LSR-A receives a Label Mapping message (assuming that the message carries the
                 FEC 192.168.3.1/32) from Area 10, LSR-A can find the summary route
                 192.168.3.0/24 based on the longest match rule. It uses the outbound interface
                 and next hop of the route as the outbound interface and next hop to the FEC
                 192.168.3.1/32. LDP can then establish an LDP LSP across IGP areas.

DoD Support for Inter-Area LDP Extension
                 In a remote DoD LDP session, LDP can use the longest match rule to establish an
                 LSP destined for the peer's LSR ID.

                 Figure 3-18 DoD support for inter-area LDP extension




                 On the network shown in Figure 3-18, only the default route 0.0.0.0 exists
                 between nodes A and C, and a remote DoD session is established between nodes
                 A and C. To establish an LSP between nodes A and C, a downstream session is
                 queried based on the longest match rule, and a Label Request message is sent to
                 the downstream node. Upon receipt of the Label Request message, the transit
                 node B checks whether an exact route exists. If no exact route is found and the
                 longest match function is enabled, node B searches for a route based on the
                 longest match rule and establishes an LSP for the route.
                 A remote LDP session in DoD mode is established on node A, no exact route to
                 the LSR ID of the peer is found, and the longest match function is configured for
                 LDP. In this case, node A can automatically send a DoD Label Request message
                 after the remote IP address is configured for the remote peer.

3.13.2 Configuring LDP Extension for Inter-Area LSPs
Prerequisites
                 Before configuring LDP extension for inter-area LSPs, you have completed the
                 following task:
                 ●      3.6.2 Configuring MPLS LDP Globally

Context
                 On a large-scale network, multiple IGP areas need to be configured for flexible
                 deployment and fast convergence. To prevent excessive resource consumption
                 caused by a large number of routes, an ABR needs to summarize the routes in an
                 area and advertise the summary routes to neighboring IGP areas. By default, when
                 establishing an LSP, LDP searches the routing table for the route that exactly
                 matches the FEC carried in a received Label Mapping message. For summary

Issue 01 (2025-03-03)          Copyright © Huawei Technologies Co., Ltd.                        100
MPLS Configuration
MPLS Configuration                                                                    3 MPLS LDP Configuration


                 routes, LDP can establish only liberal LSPs, but cannot establish LDP LSPs across
                 IGP areas. In this case, you can configure LDP extension for inter-area LSPs to
                 enable LDP to search for routes based on the longest match rule and establish
                 inter-area LDP LSPs.
                 Perform the following configuration on the ingress or transit node.

Procedure
         Step 1 Enter the system view.
                 system-view

         Step 2 Enter the MPLS-LDP view.
                 mpls ldp

         Step 3 Configure LDP extension for inter-area LSPs and enable LDP to search for routes
                based on the longest match rule to establish LSPs.
                 longest-match

                         NOTE

                        This command cannot be run during LDP GR.

                 ----End

Result
                 Run the display mpls lsp command to check information about inter-area LSPs
                 after LDP extension for inter-area LSPs is configured.

