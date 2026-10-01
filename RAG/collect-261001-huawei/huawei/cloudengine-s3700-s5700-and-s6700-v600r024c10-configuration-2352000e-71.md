---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-71
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [9526, 9706]
sha256: 74b6494aadcf536a82ce286db12d9cce482ae378d41260dc932144f9c9593439
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                             150
VPN Configuration
VPN Configuration                                                           3 IPv4 L3VPN Configuration


                        Figure 3-28 Route advertisement in an inter-AS VPN Option B scenario




                        The specific process is as follows:
                        a.   CE1 uses BGP, OSPF, or RIP to advertise the route to PE1 in AS 100.
                        b.   PE1 in AS 100 uses MP-IBGP to advertise the labeled VPNv4 route to
                             ASBR1 in AS 100. If an RR is deployed on the network, PE1 advertises the
                             VPNv4 route to the RR, and the RR then reflects the route to ASBR1.
                        c.   ASBR1 uses MP-EBGP to advertise the labeled VPNv4 route to ASBR2.
                             Because MP-EBGP changes the next hop of a route when advertising the
                             route, ASBR1 allocates a new label to the VPNv4 route.
                        d.   ASBR2 uses MP-IBGP to advertise the labeled VPNv4 route to PE2 in AS
                             200. If an RR is deployed on the network, ASBR2 advertises the VPNv4
                             route to the RR, and the RR then reflects the route to PE2. When ASBR2
                             advertises routes to an MP-IBGP peer in the local AS, it changes the next
                             hop of the routes to itself.
                        e.   PE2 in AS 200 uses BGP, OSPF, or RIP to advertise the route to CE2.
                        ASBR1 and ASBR2 both swap the inner labels of VPNv4 routes and use BGP to
                        transmit inter-AS label information. Therefore, LDP does not need to run
                        between ASBRs.
                    ●   Packet Forwarding in an Inter-AS VPN Option B Scenario
                        Figure 3-29 shows packet forwarding over an LSP on the public network.
                        Here, L1, L2, and L3 indicate VPN labels, and Lx and Ly indicate public
                        network labels (outer tunnel labels).

                        Figure 3-29 Packet forwarding in an inter-AS VPN Option B scenario




Issue 01 (2025-03-03)           Copyright © Huawei Technologies Co., Ltd.                          151
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


3.12.2 Configuring Inter-AS VPN Option B (Basic Networking)

Context
                    After an MP-EBGP peer relationship is established between ASBRs, an ASBR can
                    advertise VPNv4 routes in the local AS to the other ASBR.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Establish an IBGP peer relationship between the PE and ASBR in the same AS.
                    peer peer-address as-number as-number

         Step 4 Configure a loopback interface as the outbound interface of the BGP session.
                    peer peer-address connect-interface loopback interface-number

         Step 5 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]

         Step 6 Enable the function to exchange VPNv4 routes between the PE and ASBR in the
                same AS.
                    peer peer-address enable

                    ----End

3.12.3 Configuring MP-EBGP Between ASBRs in Different ASs

Context
                    After an MP-EBGP peer relationship is established between ASBRs, an ASBR can
                    advertise VPNv4 routes in the local AS to the other ASBR.
                    In inter-AS VPN Option B (basic networking), you do not need to create VPN
                    instances on ASBRs. An ASBR does not filter VPNv4 routes received from a PE in
                    the local AS based on VPN targets. Instead, it advertises the received routes to the
                    peer ASBR through MP-EBGP.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the view of the interface connected to the peer ASBR.
                    interface interface-type interface-number

         Step 3 Switch the interface working mode from Layer 2 to Layer 3.
                    undo portswitch

                    Only interfaces on the S6780-H, S6750-H, S6730-H-V2, S6730E-H-V2, S6750-S,
                    S6750E-S, S5732-H-V2, S5755-S, S5755-H, S5755E-H series can be switched from

Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             152
VPN Configuration
VPN Configuration                                                                   3 IPv4 L3VPN Configuration


                    Layer 2 mode to Layer 3 mode using the undo portswitch command. Determine
                    whether to perform this step based on the current interface working mode.

         Step 4 Configure an IP address for the interface.
                    ip address ip-address { mask | mask-length }

         Step 5 Enable MPLS.
                    mpls

         Step 6 Return to the system view.
                    quit

         Step 7 Enter the BGP view.
                    bgp as-number

         Step 8 Establish an IBGP peer relationship between the PE and ASBR in the same AS.
                    peer peer-address as-number as-number

         Step 9 Configure a loopback interface as the outbound interface of the BGP session.
                    peer peer-address connect-interface loopback interface-number

        Step 10 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]

        Step 11 Enable the function to exchange VPNv4 routes between the PE and ASBR in the
                same AS.
                    peer peer-address enable

                    ----End

3.12.4 Configuring ASBRs Not to Filter VPNv4 Routes Based on
VPN Targets

Context
                    By default, an ASBR filters the VPN targets of only the received VPNv4 routes.
                    Eligible routes are imported into the routing table, and other routes are discarded.
                    Therefore, if no VPN instance is configured on the ASBR or no VPN target is
                    configured for the VPN instance, the ASBR discards all the received VPNv4 routes.

                    In inter-AS VPN Option B, an ASBR does not need to store VPN instance
                    information, but must store information about all VPNv4 routes and advertise
                    these routes to the peer ASBR. In this case, ASBRs need to import all the received
                    VPNv4 routes without filtering them based on VPN targets.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Enter the BGP-VPNv4 address family view.
                    ipv4-family vpnv4 [ unicast ]


Issue 01 (2025-03-03)               Copyright © Huawei Technologies Co., Ltd.                             153
VPN Configuration
VPN Configuration                                                                            3 IPv4 L3VPN Configuration


         Step 4 Enable the function to exchange VPNv4 routes between the PE and ASBR in the
                same AS.
                    undo policy vpn-target

                    ----End

3.12.5 (Optional) Using a Route-Policy to Control VPNv4
Route Import and Export on ASBRs

Context
                    ASBRs can use a route-policy to filter VPNv4 routes based on:

                    ●      VPN targets
                    ●      RDs

Procedure
         Step 1 Enter the system view.
                    system-view

