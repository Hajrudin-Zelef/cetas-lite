---
id: collect-261001-huawei/huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration-2352000e-167
title: "cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: ["2025-03-03"]
keywords: ["copyright", "distribution"]
source: docs/RAG/collect-261001-huawei/cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e.md
source_anchor: ""
source_lines: [24110, 24284]
sha256: 53ed1b9810bc8cbb0af3318d4e21e78dfbc05e6772d981b0d6fff62d0c8c2de8
---

# cloudengine-s3700-s5700-and-s6700-v600r024c10-configuration--2352000e

                               In inter-AS IPv6 VPN Option B, an ASBR does not need to store VPN
                               instance information, but must store information about all VPNv6 routes
                               and advertise these routes to the peer ASBR. In this case, ASBRs need to
                               import all the received VPNv6 routes without filtering them based on
                               VPN targets.
                    ●    VPN target-based route filtering
                         a.    Enter the system view.
                               system-view

                         b.    Configure VPN targets using either of the following methods:
                               ip extcommunity-filter { basic-extcomm-filter-num | basic basic-extcomm-filter-name } { deny
                               | permit } { rt { as-number:nn | 4as-number:nn | ipv4-address:nn } } &<1-16>


Issue 01 (2025-03-03)              Copyright © Huawei Technologies Co., Ltd.                                           380
VPN Configuration
VPN Configuration                                                                          4 IPv6 L3VPN Configuration

                               ip extcommunity-filter { advanced-extcomm-filter-num | advanced advanced-extcomm-filter-
                               name } { deny | permit } regular-expression
                         c.    Configure a route-policy.
                               route-policy route-policy-name permit node node

                         d.    Configure a matching rule based on the extended community filter for
                               the route-policy.
                               if-match extcommunity-filter { { basic-extcomm-filter-num | adv-extcomm-filter-num }
                               &<1-16> | basic-extcomm-filter-name | advanced-extcomm-filter-name }

                         e.    Return to the system view.
                               quit

                         f.    Enter the BGP view.
                               bgp as-number

                         g.    Enter the BGP-VPNv6 address family view.
                               ipv6-family vpnv6 [ unicast ]

                         h.    Apply the route-policy to control the import and export of VPNv6 routes.
                               peer ipv4-address route-policy route-policy-name { export | import }

                    ----End

4.13.5 (Optional) Configuring One-Label-per-Next-Hop Label
Distribution on an ASBR

Context
                    In inter-AS IPv6 VPN Option B, after one-label-per-next-hop label distribution is
                    configured on an ASBR, the ASBR assigns only one label to VPNv6 routes that
                    share the same next hop and outgoing label. Compared with the on-label-per-
                    route mode, the one-label-per-next-hop mode significantly saves label resources
                    on the ASBR.
                    Perform the following steps on each ASBR.

Procedure
         Step 1 Enter the system view.
                    system-view

         Step 2 Enter the BGP view.
                    bgp as-number

         Step 3 Enter the BGP-VPNv6 address family view.
                    ipv6-family vpnv6 [ unicast ]

         Step 4 Enable one-label-per-next-hop label distribution for VPNv6 routes on the ASBR.
                    apply-label per-nexthop




Issue 01 (2025-03-03)                 Copyright © Huawei Technologies Co., Ltd.                                       381
VPN Configuration
VPN Configuration                                                            4 IPv6 L3VPN Configuration


                        NOTICE

                    After one-label-per-next-hop label distribution is enabled or disabled on an ASBR,
                    the labels assigned by the ASBR to routes change. As a result, temporary packet
                    loss may occur.


                    ----End

4.13.6 Configuring Route Exchange Between the CE and PE

Context
                    BGP, IGP, or static routes (including the default routes) can be used between a CE
                    and a PE. Determine which one to use as required.

Procedure
         Step 1 Determine whether to use BGP, IGP, or static routes between a PE and a CE as
                required. For details, see Configuring an IPv6 VPN Instance.

                    ----End

4.13.7 Verifying the Configuration

Prerequisites
                    All the configurations about inter-AS IPv6 VPN Option B are complete.

Procedure
                    ●   Run the display bgp vpnv6 all peer command on PEs or ASBRs to check the
                        establishment of all BGP peer relationships.
                    ●   Run the display bgp vpnv6 all routing-table command on PEs or ASBRs to
                        check VPNv6 routes.
                    ●   Run the display ipv6 routing-table vpn-instance vpn-instance-name
                        command on PEs to check the VPN routing table.

                    ----End

4.13.8 Example for Configuring IPv6 L3VPN over MPLS Inter-
AS Option B

Networking Requirements
                    A single-hop MP-EBGP peer relationship can be established between ASBRs to
                    exchange VPNv6 routes.

                    On the network shown in Figure 4-10, CE1 and CE2 belong to the same VPN. CE1
                    connects to PE1 in AS100, and CE2 connects to PE2 in AS200. It is required that an

Issue 01 (2025-03-03)            Copyright © Huawei Technologies Co., Ltd.                         382
VPN Configuration
VPN Configuration                                                                   4 IPv6 L3VPN Configuration


                    MP-EBGP peer relationship be established between the ASBRs to transmit VPNv6
                    routes, thereby implementing inter-AS IPv6 VPN Option B.

                    Figure 4-10 Inter-AS IPv6 VPN Option B
                          NOTE

                         In this example, interface 1 and interface 2 represent VLANIF 100 and VLANIF 200,
                         respectively.




Precautions
                    During the configuration, note the following:
                    ●    An MP-EBGP peer relationship is established between ASBR1 and ASBR2, and
                         the ASBRs do not filter received VPNv6 routes based on VPN targets.

Configuration Roadmap
                    The configuration roadmap is as follows:
                    1.   Configure IGP on the backbone network for IP connectivity between the ASBR
                         and PE in the same AS, and set up an MPLS LDP LSP between the ASBR and
                         PE in the same AS.
                    2.   Set up EBGP peer relationships between PEs and CEs and MP-IBGP peer
                         relationships between PEs and ASBRs.
                    3.   Configure VPN instances on the PEs rather than ASBRs.
                    4.   Enable MPLS on the ASBR interfaces connected to each other. Establish an
                         MP-EBGP peer relationship between ASBRs and configure ASBRs not to filter
                         received VPNv6 routes based on VPN targets.




Issue 01 (2025-03-03)             Copyright © Huawei Technologies Co., Ltd.                                  383
VPN Configuration
VPN Configuration                                                                    4 IPv6 L3VPN Configuration


Procedure
         Step 1 Configure IGP on the MPLS backbone networks in AS100 and AS200 for
                communication between PEs on each backbone network. OSPF is used as IGP in
                this example. For detailed configurations, see Configuration Scripts.
                          NOTE

                        The 32-bit IP address of the loopback interface that functions as the LSR ID needs to be
                        advertised using OSPF.

